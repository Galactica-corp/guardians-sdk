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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	merkleproofservice "github.com/Galactica-corp/merkle-proof-service/gen/galactica/merkle"
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
	merkleProofServiceURL  string
	merkleProofServiceTLS  bool
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
	cmd.Flags().StringVarP(&f.rpcURL, "rpc-url", "", "", "url of Ethereum compatible RPC endpoint")
	cmd.Flags().StringVarP(&f.merkleProofServiceURL, "merkle-proof-service-url", "m", "", "Merkle Proof Service gRPC endpoint url")
	cmd.Flags().BoolVar(&f.merkleProofServiceTLS, "merkle-proof-service-tls", false, "enable TLS for Merkle Proof Service gRPC connection")

	_ = cmd.MarkFlagRequired("certificate-file")
	_ = cmd.MarkFlagRequired("provider-private-key")
	_ = cmd.MarkFlagRequired("registry-address")
	_ = cmd.MarkFlagRequired("rpc-url")
	_ = cmd.MarkFlagRequired("merkle-proof-service-url")

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

	merkleProofClient, err := merkle.ConnectToMerkleProofService(ctx, f.merkleProofServiceURL, f.merkleProofServiceTLS)
	if err != nil {
		return fmt.Errorf("connect to merkle proof service: %w", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("retrieve chain-id: %w", err)
	}

	registryAddress := f.registryAddress.Address()

	registry, err := contracts.NewZkCertificateRegistry(registryAddress, client)
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

	emptyLeafIndex, proof, err := findEmptyTreeLeaf(ctx, merkleProofClient, registryAddress)
	if err != nil {
		return fmt.Errorf("find empty tree leaf: %w", err)
	}

	if err := registerAndWaitForZkCertificateTurn(ctx, client, providerKey, registry, certificate.LeafHash); err != nil {
		return fmt.Errorf("register and wait for zkCertificate turn: %w", err)
	}

	tx, err := constructIssueZKCertTx(ctx, client, providerKey, registry, emptyLeafIndex, certificate.LeafHash, proof)
	if err != nil {
		return fmt.Errorf("construct transaction to add record to registry: %w", err)
	}

	if receipt, err := bind.WaitMined(ctx, client, tx); err != nil {
		return fmt.Errorf("wait until transaction is mined: %w", err)
	} else if receipt.Status == 0 {
		return fmt.Errorf("transaction %q failed", receipt.TxHash)
	}

	if err := json.NewEncoder(os.Stdout).Encode(tx); err != nil {
		return fmt.Errorf("encode registration transaction to json: %w", err)
	}

	if err := buildAndSaveOutput(f.outputFilePath, certificate, registryAddress, chainID, emptyLeafIndex, proof); err != nil {
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

type (
	RegistryEventParser interface {
		ParseZkCertificateAddition(log types.Log) (*contracts.ZkCertificateRegistryZkCertificateAddition, error)
		ParseZkCertificateRevocation(log types.Log) (*contracts.ZkCertificateRegistryZkCertificateRevocation, error)
	}

	RecordRegistryTransactor interface {
		AddZkCertificate(
			opts *bind.TransactOpts,
			leafIndex *big.Int,
			zkCertificateHash [32]byte,
			merkleProof [][32]byte,
		) (*types.Transaction, error)
		RevokeZkCertificate(
			opts *bind.TransactOpts,
			leafIndex *big.Int,
			zkCertificateHash [32]byte,
			merkleProof [][32]byte,
		) (*types.Transaction, error)
		RegisterToQueue(opts *bind.TransactOpts, zkCertificateHash [32]byte) (*types.Transaction, error)
		CheckZkCertificateHashInQueue(opts *bind.CallOpts, zkCertificateHash [32]byte) (bool, error)
	}

	RecordRegistryCaller interface {
		GuardianRegistry(opts *bind.CallOpts) (common.Address, error)
		GetTimeParameters(opts *bind.CallOpts, zkCertificateHash [32]byte) (*big.Int, *big.Int, error)
	}

	RecordRegistry interface {
		RegistryEventParser
		RecordRegistryTransactor
		RecordRegistryCaller
	}
)

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
	client merkleproofservice.QueryClient,
	registryAddress common.Address,
) (int, merkle.Proof, error) {
	emptyLeafIndex, proof, err := merkle.GetEmptyLeafProof(ctx, client, registryAddress.Hex())
	if err != nil {
		return 0, merkle.Proof{}, fmt.Errorf("get empty leaf proof: %w", err)
	}

	return int(emptyLeafIndex), proof, nil
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

	return recordRegistry.AddZkCertificate(
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
	chainID *big.Int,
	leafIndex int,
	proof merkle.Proof,
) error {
	if err := encodeToJSONFile(outputFilePath, zkcertificate.IssuedCertificate[T]{
		Certificate: certificate,
		Registration: zkcertificate.RegistrationDetails{
			Address:   registryAddress,
			ChainID:   chainID,
			Revocable: true,
			LeafIndex: leafIndex,
		},
		MerkleProof: proof,
	}); err != nil {
		return fmt.Errorf("save issued certificate: %w", err)
	}

	return nil
}

func encodeMerkleProof(proof merkle.Proof) [][32]byte {
	res := make([][32]byte, len(proof.Path))

	for i, node := range proof.Path {
		res[i] = node.Value.Bytes32()
	}

	return res
}

func registerToQueue(
	ctx context.Context,
	client *ethclient.Client,
	providerKey *ecdsa.PrivateKey,
	recordRegistry RecordRegistryTransactor,
	leafHash zkcertificate.Hash,
) (*types.Transaction, error) {
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieve chain id: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(providerKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("create transaction signer from private key: %w", err)
	}

	tx, err := recordRegistry.RegisterToQueue(auth, leafHash.Bytes32())
	if err != nil {
		exists, checkErr := recordRegistry.CheckZkCertificateHashInQueue(&bind.CallOpts{}, leafHash.Bytes32())
		if checkErr != nil {
			return nil, fmt.Errorf("register to queue failed: %w, also failed to check if zkCertificateHash is in queue: %w", err, checkErr)
		}
		if exists {
			return nil, nil
		}
		return nil, fmt.Errorf("register to queue failed: %w", err)
	}

	return tx, nil
}

func registerAndWaitForZkCertificateTurn(
	ctx context.Context,
	client *ethclient.Client,
	providerKey *ecdsa.PrivateKey,
	registry RecordRegistry,
	leafHash zkcertificate.Hash,
) error {
	registerTx, err := registerToQueue(ctx, client, providerKey, registry, leafHash)
	if err != nil {
		return fmt.Errorf("register zkCertificate hash to queue: %w", err)
	}

	if registerTx != nil {
		receipt, err := bind.WaitMined(ctx, client, registerTx)
		if err != nil {
			return fmt.Errorf("wait until queue registration transaction is mined: %w", err)
		}
		if receipt.Status == 0 {
			return fmt.Errorf("queue registration transaction %q failed", receipt.TxHash)
		}
	}

	for {
		startTime, expirationTime, err := registry.GetTimeParameters(&bind.CallOpts{}, leafHash.Bytes32())
		if err != nil {
			return fmt.Errorf("retrieve time parameters for zkCertificate hash: %w", err)
		}

		currentTime := big.NewInt(time.Now().Unix())
		if startTime.Cmp(currentTime) <= 0 {
			break
		}
		if expirationTime.Cmp(currentTime) <= 0 {
			return fmt.Errorf("queue wait time expired for zkCertificate hash %s", leafHash)
		}

		select {
		case <-time.After(5 * time.Second):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}
