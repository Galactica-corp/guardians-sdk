// Copyright © 2024 Galactica Network
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"

	"github.com/galactica-corp/guardians-sdk/internal/cli"
	"github.com/galactica-corp/guardians-sdk/pkg/contracts"
	"github.com/galactica-corp/guardians-sdk/pkg/merkle"
	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

type issueZKCertFlags struct {
	certificateFilePath    string
	outputFilePath         string
	providerPrivateKeyPath string
	rpcURL                 string
	registryAddress        cli.Address
	firstBlock             int64
}

func NewCmdIssueZKCert() *cobra.Command {
	var f issueZKCertFlags

	cmd := &cobra.Command{
		Use:   "issueZKCert",
		Short: "Issue a Zero Knowledge Certificate (ZKCert) to the Galactica blockchain registry",
		Long: `The issueZKCert command facilitates the issuance of a Zero Knowledge Certificate
(ZKCert) to the Galactica blockchain registry. ZKCert is a crucial component for
providing trust and privacy in the Galactica ecosystem.

This command involves several steps, including connecting to the blockchain RPC,
retrieving the relevant registry contract, verifying the certificate standard,
and checking the guardian whitelist status. The ZKCert issuance process requires
a provider's Ethereum private key for signing the transaction.

Upon successful issuance, the command outputs transaction details, registration
information, and the Merkle proof for verification purposes. After submitting
the transaction, the issued ZKCert will be added to the blockchain registry,
ensuring its validity and accessibility.

Example Usage:
$ galactica-guardian issueZKCert -c zkcert.json -k provider_private_key.hex -o output.json`,
		RunE: issueZKCertCmd(&f),
	}

	cmd.Flags().StringVarP(&f.certificateFilePath, "certificate-file", "c", "", "path to a file containing zkCert created using createZKCert command")
	cmd.Flags().StringVarP(&f.outputFilePath, "output-file", "o", "issued-certificate.json", "path to a file where the issued certificate in JSON format should be saved")
	cmd.Flags().StringVarP(&f.providerPrivateKeyPath, "provider-private-key", "k", "", "path to a file containing provider's hex-encoded Ethereum (ECDSA) private key to sign the transaction")
	cmd.Flags().VarP(&f.registryAddress, "registry-address", "r", "Ethereum address of the registry contract on-chain")
	cmd.Flags().Int64VarP(&f.firstBlock, "registry-events-start", "", 0, "block number in which first event was emitted by the registry. This block might be before the first event, but if it will be after, then it will lead to incorrect result. It greatly improves time to build a merkle tree, because RPC requests are limited to inspect at most 10'000 blocks at once")
	cmd.Flags().StringVarP(&f.rpcURL, "rpc-url", "", "", "url of Ethereum compatible RPC endpoint")

	_ = cmd.MarkFlagRequired("certificate-file")
	_ = cmd.MarkFlagRequired("provider-private-key")
	_ = cmd.MarkFlagRequired("registry-address")
	_ = cmd.MarkFlagRequired("rpc-url")

	return cmd
}

func issueZKCertCmd(f *issueZKCertFlags) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return issueZKCert(f)
	}
}

func issueZKCert(f *issueZKCertFlags) error {
	ctx := context.Background()

	var certificate zkcertificate.Certificate[json.RawMessage]
	if err := decodeJSONFile(f.certificateFilePath, &certificate); err != nil {
		return fmt.Errorf("read certificate: %w", err)
	}

	client, err := connectToBlockchainRPC(ctx, f.rpcURL)
	if err != nil {
		return fmt.Errorf("connect to blockchain rpc: %w", err)
	}

	registryAddress := f.registryAddress.Address()

	registry, err := loadRecordRegistry(client, registryAddress, certificate.Standard)
	if err != nil {
		return fmt.Errorf("load record registry: %w", err)
	}

	providerKey, err := crypto.LoadECDSA(f.providerPrivateKeyPath)
	if err != nil {
		return fmt.Errorf("load provider's ethereum private key: %w", err)
	}

	if err := ensureProviderIsGuardian(client, registry, crypto.PubkeyToAddress(providerKey.PublicKey)); err != nil {
		return fmt.Errorf("ensure provider is guardian: %w", err)
	}

	emptyLeafIndex, proof, err := findEmptyTreeLeaf(ctx, client, registryAddress, registry, f.firstBlock)
	if err != nil {
		return fmt.Errorf("find empty tree leaf: %w", err)
	}

	tx, err := constructIssueZKCertTx(ctx, client, providerKey, registry, emptyLeafIndex, certificate.LeafHash, proof)
	if err != nil {
		return fmt.Errorf("construct transaction to add record to registry: %w", err)
	}

	if receipt, err := bind.WaitMined(ctx, client, tx); err != nil {
		return fmt.Errorf("wait until transaction is mined: %w", err)
	} else if receipt.Status == 0 {
		// TODO: Inspect a transaction that is really failed
		return fmt.Errorf("transaction %q falied", receipt.TxHash)
	}

	if err := json.NewEncoder(os.Stdout).Encode(tx); err != nil {
		return fmt.Errorf("encode registration transaction to json: %w", err)
	}

	if err := buildAndSaveOutput(f.outputFilePath, certificate, registryAddress, emptyLeafIndex, proof); err != nil {
		return fmt.Errorf("collect output: %w", err)
	}

	_, _ = fmt.Fprintln(os.Stderr, "Saved issued certificate to", f.outputFilePath)

	return nil
}

func connectToBlockchainRPC(ctx context.Context, rawURL string) (*ethclient.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	return ethclient.DialContext(ctx, rawURL)
}

// TODO: It would be nice to use the same ABI for every contract
type (
	RegistryEventParser interface {
		ParseZkKYCRecordAddition(log types.Log) (*contracts.KYCRecordRegistryZkKYCRecordAddition, error)
		ParseZkKYCRecordRevocation(log types.Log) (*contracts.KYCRecordRegistryZkKYCRecordRevocation, error)
	}

	RecordRegistryTransactor interface {
		AddZkKYCRecord(
			opts *bind.TransactOpts,
			leafIndex *big.Int,
			zkKYCRecordHash [32]byte,
			merkleProof [][32]byte,
		) (*types.Transaction, error)
		RevokeZkKYCRecord(
			opts *bind.TransactOpts,
			leafIndex *big.Int,
			zkKYCRecordHash [32]byte,
			merkleProof [][32]byte,
		) (*types.Transaction, error)
	}

	RecordRegistryCaller interface {
		GuardianRegistry(opts *bind.CallOpts) (common.Address, error)
	}

	RecordRegistry interface {
		RegistryEventParser
		RecordRegistryTransactor
		RecordRegistryCaller
	}
)

func loadRecordRegistry(
	client bind.ContractBackend,
	registryAddress common.Address,
	standard zkcertificate.Standard,
) (RecordRegistry, error) {
	switch standard {
	case zkcertificate.StandardKYC:
		recordRegistry, err := contracts.NewKYCRecordRegistry(registryAddress, client)
		if err != nil {
			return nil, fmt.Errorf("bind kyc record registry contract: %w", err)
		}

		return recordRegistry, nil
	default:
		return nil, fmt.Errorf("standard %s is not supported", standard)
	}
}

func ensureProviderIsGuardian(
	client bind.ContractBackend,
	registry RecordRegistryCaller,
	providerAddress common.Address,
) error {
	guardianRegistryAddress, err := registry.GuardianRegistry(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("retrieve guardian registry address: %w", err)
	}

	guardianRegistry, err := contracts.NewGuardianRegistry(guardianRegistryAddress, client)
	if err != nil {
		return fmt.Errorf("bind guardian registry contract: %w", err)
	}

	guardian, err := guardianRegistry.Guardians(&bind.CallOpts{}, providerAddress)
	if err != nil {
		return fmt.Errorf("retrieve guardian whitelist status: %w", err)
	}
	if !guardian.Whitelisted {
		return fmt.Errorf("provider %s is not a guardian yet", providerAddress)
	}

	return nil
}

func findEmptyTreeLeaf(
	ctx context.Context,
	client *ethclient.Client,
	registryAddress common.Address,
	recordRegistry RegistryEventParser,
	registryFirstEventBlockNumber int64,
) (int, merkle.Proof, error) {
	tree, err := buildMerkleTreeFromEvents(ctx, client, registryAddress, recordRegistry, registryFirstEventBlockNumber)
	if err != nil {
		return 0, merkle.Proof{}, fmt.Errorf("build merkle tree from events: %w", err)
	}

	emptyLeafIndex, err := findFirstEmptyLeafIndex(tree)
	if err != nil {
		return 0, merkle.Proof{}, fmt.Errorf("find first empty leaf index: %w", err)
	}

	proof, err := tree.GetProof(emptyLeafIndex)
	if err != nil {
		return 0, merkle.Proof{}, fmt.Errorf("compute merkle proof: %w", err)
	}

	return emptyLeafIndex, proof, nil
}

func constructIssueZKCertTx(
	ctx context.Context,
	client *ethclient.Client,
	providerKey *ecdsa.PrivateKey,
	recordRegistry RecordRegistryTransactor,
	emptyLeafIndex int,
	leafHash zkcertificate.Hash,
	proof merkle.Proof,
) (*types.Transaction, error) {
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieve chain id: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(providerKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("create transaction signer from private key: %w", err)
	}

	return recordRegistry.AddZkKYCRecord(
		auth,
		big.NewInt(int64(emptyLeafIndex)),
		leafHash.Bytes32(),
		encodeMerkleProof(proof),
	)
}

func buildAndSaveOutput[T any](
	outputFilePath string,
	certificate zkcertificate.Certificate[T],
	registryAddress common.Address,
	leafIndex int,
	proof merkle.Proof,
) error {
	if err := encodeToJSONFile(outputFilePath, zkcertificate.IssuedCertificate[T]{
		Certificate: certificate,
		Registration: zkcertificate.RegistrationDetails{
			Address:   registryAddress,
			Revocable: true,
			LeafIndex: leafIndex,
		},
		MerkleProof: proof,
	}); err != nil {
		return fmt.Errorf("save issued certificate: %w", err)
	}

	return nil
}

// TODO: Deployed contract doesn't emit these events, it emits an older version of them
var (
	signatureRecordAddition   = crypto.Keccak256Hash([]byte("zkKYCRecordAddition(bytes32,address,uint)"))
	signatureRecordRevocation = crypto.Keccak256Hash([]byte("zkKYCRecordRevocation(bytes32,address,uint)"))
)

var blocksDistance = big.NewInt(10_000)

var blockDistancePlusOne = new(big.Int).Add(blocksDistance, big.NewInt(1))

func newProgressBar(max int64) *progressbar.ProgressBar {
	return progressbar.NewOptions64(
		max,
		progressbar.OptionSetDescription("Query blockchain events to build the Merkle tree"),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionSetWidth(10),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionSetItsString("block"),
		progressbar.OptionOnCompletion(func() {
			_, _ = fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetRenderBlankState(true),
	)
}

func buildMerkleTreeFromEvents(
	ctx context.Context,
	client *ethclient.Client,
	registryAddress common.Address,
	registryEventParser RegistryEventParser,
	firstBlock int64,
) (*merkle.Tree, error) {
	tree, err := merkle.NewEmptyTree(merkle.TreeDepth, merkle.EmptyLeafValue)
	if err != nil {
		return nil, fmt.Errorf("initialize empty tree: %w", err)
	}

	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieve head block number: %w", err)
	}

	headBlock := big.NewInt(int64(head))

	bar := newProgressBar(int64(head) - firstBlock)

	for i := big.NewInt(firstBlock); i.Cmp(headBlock) == -1; i.Add(i, blockDistancePlusOne) {
		fromBlock := new(big.Int).Set(i)
		toBlock := new(big.Int).Add(fromBlock, blocksDistance)
		if toBlock.Cmp(headBlock) == 1 {
			toBlock = new(big.Int).Set(headBlock)
		}

		logs, err := client.FilterLogs(ctx, ethereum.FilterQuery{
			FromBlock: fromBlock,
			ToBlock:   toBlock,
			Addresses: []common.Address{registryAddress},
			Topics:    [][]common.Hash{{signatureRecordAddition, signatureRecordRevocation}},
		})
		if err != nil {
			return nil, fmt.Errorf("execute filter query: %w", err)
		}

		for _, logEntry := range logs {
			if err := processEvent(logEntry, registryEventParser, tree); err != nil {
				return nil, err
			}
		}

		blockRangeLength := new(big.Int).Sub(toBlock, fromBlock).Int64()
		if toBlock.Cmp(headBlock) != 0 {
			blockRangeLength++ // because it is a closed range
		}

		_ = bar.Add64(blockRangeLength)
	}

	_ = bar.Finish()

	return tree, nil
}

func processEvent(logEntry types.Log, registryEventParser RegistryEventParser, tree *merkle.Tree) error {
	if logEntry.Removed {
		return fmt.Errorf("not supported: log is removed due to chain reorganisation")
	}

	var index int
	var leafHash *uint256.Int

	switch logEntry.Topics[0] {
	case signatureRecordAddition:
		eventAddition, err := registryEventParser.ParseZkKYCRecordAddition(logEntry)
		if err != nil {
			return fmt.Errorf("parse addition record event: %w", err)
		}

		index = int(eventAddition.Index.Int64())

		var isOverflow bool
		leafHash, isOverflow = uint256.FromBig(new(big.Int).SetBytes(eventAddition.ZkKYCRecordLeafHash[:]))
		if isOverflow {
			return fmt.Errorf("invalid leaf hash")
		}
	case signatureRecordRevocation:
		eventRevocation, err := registryEventParser.ParseZkKYCRecordRevocation(logEntry)
		if err != nil {
			return fmt.Errorf("parse revocation record event: %w", err)
		}

		index = int(eventRevocation.Index.Int64())
		leafHash = merkle.EmptyLeafValue
	}

	if err := tree.SetLeaf(index, merkle.TreeNode{Value: leafHash}); err != nil {
		return fmt.Errorf("set leaf: %w", err)
	}

	return nil
}

func findFirstEmptyLeafIndex(tree *merkle.Tree) (int, error) {
	leavesAmount := tree.GetLeavesAmount()
	offset := len(tree.Nodes) - leavesAmount

	for i := 0; i < leavesAmount; i++ {
		if tree.Nodes[offset+i].Value.Eq(merkle.EmptyLeafValue) {
			return i, nil
		}
	}

	return 0, fmt.Errorf("tree is full")
}

func encodeMerkleProof(proof merkle.Proof) [][32]byte {
	res := make([][32]byte, len(proof.Path))

	for i, node := range proof.Path {
		res[i] = node.Value.Bytes32()
	}

	return res
}
