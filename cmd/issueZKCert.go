// Copyright © 2025 Galactica Network
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
	"github.com/holiman/uint256"
	"github.com/spf13/cobra"

	"github.com/galactica-corp/guardians-sdk/v2/internal/cli"
	"github.com/galactica-corp/guardians-sdk/v2/pkg/contracts"
	"github.com/galactica-corp/guardians-sdk/v2/pkg/merkle"
	"github.com/galactica-corp/guardians-sdk/v2/pkg/zkcertificate"
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

	cert, err := deserializeCertificateJSON(f.certificateFilePath)
	if err != nil {
		return fmt.Errorf("read certificate: %w", err)
	}

	client, err := connectToBlockchainRPC(ctx, f.rpcURL)
	if err != nil {
		return fmt.Errorf("connect to blockchain rpc: %w", err)
	}

	merkleProofClient, err := merkle.ConnectToMerkleProofService(f.merkleProofServiceURL, f.merkleProofServiceTLS)
	if err != nil {
		return fmt.Errorf("connect to merkle proof service: %w", err)
	}

	registryAddress := f.registryAddress.Address()

	providerKey, err := crypto.LoadECDSA(f.providerPrivateKeyPath)
	if err != nil {
		return fmt.Errorf("load provider's ethereum private key: %w", err)
	}

	tx, issuedCertificate, err := IssueZKCert(ctx, cert, client, merkleProofClient, registryAddress, providerKey)
	if err != nil {
		return fmt.Errorf("issue zk cert: %w", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(tx); err != nil {
		return fmt.Errorf("encode registration transaction to json: %w", err)
	}

	if err := encodeToJSONFile(f.outputFilePath, issuedCertificate); err != nil {
		return fmt.Errorf("save issued certificate: %w", err)
	}

	_, _ = fmt.Fprintln(os.Stderr, "Saved issued certificate to", f.outputFilePath)

	return nil
}

type EthereumIssueClient interface {
	bind.ContractBackend
	bind.DeployBackend

	ChainID(ctx context.Context) (*big.Int, error)
}

// IssueZKCert registers and issues a zero-knowledge certificate (ZKCert) on-chain.
//
// The function performs the following steps:
//  1. Verifies that the provider is authorized as a guardian.
//  2. Finds an empty leaf in the Merkle tree for certificate registration.
//  3. Registers the certificate and waits for the provider's turn to issue it.
//  4. Constructs and sends the transaction to add the certificate on-chain.
func IssueZKCert[T zkcertificate.Content](
	ctx context.Context,
	cert zkcertificate.Certificate[T],
	ethRPC EthereumIssueClient,
	merkleProofClient merkle.EmptyLeafProver,
	registryAddress common.Address,
	providerKey *ecdsa.PrivateKey,
) (*types.Transaction, zkcertificate.IssuedCertificate[T], error) {
	chainID, err := ethRPC.ChainID(ctx)
	if err != nil {
		return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("retrieve chain-id: %w", err)
	}

	registry, err := contracts.NewZkCertificateRegistry(registryAddress, ethRPC)
	if err != nil {
		return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("load record registry: %w", err)
	}

	providerAddress := crypto.PubkeyToAddress(providerKey.PublicKey)

	if err := ensureProviderIsGuardian(ctx, ethRPC, registry, providerAddress); err != nil {
		return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("ensure provider is guardian: %w", err)
	}

	if cert.Standard == zkcertificate.StandardKYC {
		content := any(cert.Content).(zkcertificate.KYCContent)

		if err := ensureKYCSaltIsCompatible(ctx, content, cert.HolderCommitment, ethRPC, registryAddress, providerAddress); err != nil {
			return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("ensure KYC salt is compatible: %w", err)
		}
	}

	leafHash := cert.LeafHash

	if err := registerAndWaitForZkCertificateTurn(ctx, ethRPC, chainID, providerKey, registry, leafHash); err != nil {
		return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("register and wait for issue turn: %w", err)
	}

	index, proof, err := findEmptyTreeLeaf(ctx, merkleProofClient, registryAddress)
	if err != nil {
		return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("find empty tree leaf: %w", err)
	}

	tx, err := constructIssueZKCertTx(ctx, ethRPC, chainID, providerKey, registryAddress, index, cert, proof)
	if err != nil {
		return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("construct add record tx: %w", err)
	}

	if receipt, err := bind.WaitMined(ctx, ethRPC, tx); err != nil {
		return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("wait until transaction is mined: %w", err)
	} else if receipt.Status == 0 {
		return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("transaction %q failed", receipt.TxHash)
	}

	proof.Leaf = merkle.TreeNode{Value: uint256.MustFromBig(cert.LeafHash.BigInt())}

	return tx, zkcertificate.IssuedCertificate[T]{
		Certificate: cert,
		Registration: zkcertificate.RegistrationDetails{
			Address:   registryAddress,
			ChainID:   chainID,
			Revocable: true,
			LeafIndex: index,
		},
		MerkleProof: proof,
	}, nil
}

func ensureKYCSaltIsCompatible(
	ctx context.Context,
	content zkcertificate.KYCContent,
	salt zkcertificate.Hash,
	ethRPC EthereumIssueClient,
	registryAddress common.Address,
	providerAddress common.Address,
) error {
	recordRegistry, err := contracts.NewZkKYCRegistry(registryAddress, ethRPC)
	if err != nil {
		return fmt.Errorf("load kyc registry: %w", err)
	}

	saltRegistryAddress, err := recordRegistry.HumanIDSaltRegistry(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("get salt registry: %w", err)
	}

	saltRegistry, err := contracts.NewHumanIDSaltRegistry(saltRegistryAddress, ethRPC)
	if err != nil {
		return fmt.Errorf("load salt registry: %w", err)
	}

	idHash, err := content.IDHash()
	if err != nil {
		return fmt.Errorf("get id hash: %w", err)
	}

	registeredHash, err := saltRegistry.GetSaltHash(&bind.CallOpts{
		Context: ctx,
		From:    providerAddress,
	}, idHash.BigInt())
	if err != nil {
		return fmt.Errorf("get registered salt: %w", err)
	}

	if registeredHash.Cmp(&big.Int{}) != 0 && registeredHash.Cmp(salt.BigInt()) != 0 {
		return ErrSaltIncompatible
	}

	return nil
}

func connectToBlockchainRPC(ctx context.Context, rawURL string) (*ethclient.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	return ethclient.DialContext(ctx, rawURL)
}

type RecordRegistryAddressCaller interface {
	GuardianRegistry(opts *bind.CallOpts) (common.Address, error)
}

func ensureProviderIsGuardian(
	ctx context.Context,
	client bind.ContractBackend,
	registry RecordRegistryAddressCaller,
	providerAddress common.Address,
) error {
	guardianRegistryAddress, err := registry.GuardianRegistry(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("retrieve guardian registry address: %w", err)
	}

	guardianRegistry, err := contracts.NewGuardianRegistry(guardianRegistryAddress, client)
	if err != nil {
		return fmt.Errorf("bind guardian registry contract: %w", err)
	}

	whitelisted, err := guardianRegistry.IsWhitelisted(&bind.CallOpts{Context: ctx}, providerAddress)
	if err != nil {
		return fmt.Errorf("retrieve guardian whitelist status: %w", err)
	}

	if !whitelisted {
		return fmt.Errorf("provider %s is not a guardian yet", providerAddress)
	}

	return nil
}

func findEmptyTreeLeaf(
	ctx context.Context,
	client merkle.EmptyLeafProver,
	registryAddress common.Address,
) (int, merkle.Proof, error) {
	emptyLeafIndex, proof, err := merkle.GetEmptyLeafProof(ctx, client, registryAddress.Hex())
	if err != nil {
		return 0, merkle.Proof{}, fmt.Errorf("get empty leaf proof: %w", err)
	}

	return int(emptyLeafIndex), proof, nil
}

func constructIssueZKCertTx[T zkcertificate.Content](
	ctx context.Context,
	client bind.ContractBackend,
	chainID *big.Int,
	providerKey *ecdsa.PrivateKey,
	registryAddress common.Address,
	leafIndex int,
	cert zkcertificate.Certificate[T],
	proof merkle.Proof,
) (*types.Transaction, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(providerKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("create transaction signer from private key: %w", err)
	}

	auth.Context = ctx

	if cert.Standard == zkcertificate.StandardKYC {
		// For zkCertificate of type zkKYC, the smart contract interface expects a few more parameters
		recordRegistry, err := contracts.NewZkKYCRegistry(registryAddress, client)
		if err != nil {
			return nil, fmt.Errorf("load record registry: %w", err)
		}

		// get KYC data from content to get the ID hash
		kycContent := any(cert.Content).(zkcertificate.KYCContent)

		idHash, err := kycContent.IDHash()
		if err != nil {
			return nil, fmt.Errorf("get id hash: %w", err)
		}

		return recordRegistry.AddZkKYC(
			auth,
			big.NewInt(int64(leafIndex)),
			cert.LeafHash.Bytes32(),
			encodeMerkleProof(proof),
			idHash.BigInt(),
			cert.HolderCommitment.BigInt(),
			big.NewInt(cert.ExpirationDate.Unix()),
		)
	} else {
		recordRegistry, err := contracts.NewZkCertificateRegistry(registryAddress, client)
		if err != nil {
			return nil, fmt.Errorf("load record registry: %w", err)
		}

		return recordRegistry.AddZkCertificate(
			auth,
			big.NewInt(int64(leafIndex)),
			cert.LeafHash.Bytes32(),
			encodeMerkleProof(proof),
		)
	}
}

func encodeMerkleProof(proof merkle.Proof) [][32]byte {
	res := make([][32]byte, len(proof.Path))

	for i, node := range proof.Path {
		res[i] = node.Value.Bytes32()
	}

	return res
}

type RecordRegistryQueue interface {
	CheckZkCertificateHashInQueue(opts *bind.CallOpts, zkCertificateHash [32]byte) (bool, error)
	RegisterToQueue(opts *bind.TransactOpts, zkCertificateHash [32]byte) (*types.Transaction, error)
}

func registerToQueue(
	ctx context.Context,
	chainID *big.Int,
	providerKey *ecdsa.PrivateKey,
	registry RecordRegistryQueue,
	leafHash zkcertificate.Hash,
) (*types.Transaction, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(providerKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("create transaction signer from private key: %w", err)
	}

	auth.Context = ctx

	tx, err := registry.RegisterToQueue(auth, leafHash.Bytes32())
	if err != nil {
		exists, checkErr := registry.CheckZkCertificateHashInQueue(&bind.CallOpts{Context: ctx}, leafHash.Bytes32())
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
	client bind.DeployBackend,
	chainID *big.Int,
	providerKey *ecdsa.PrivateKey,
	registry RecordRegistryQueue,
	leafHash zkcertificate.Hash,
) error {
	registerTx, err := registerToQueue(ctx, chainID, providerKey, registry, leafHash)
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
		myTurn, err := registry.CheckZkCertificateHashInQueue(&bind.CallOpts{Context: ctx}, leafHash.Bytes32())
		if err != nil {
			return fmt.Errorf("retrieve zkCertificate hash to check: %w", err)
		}

		if myTurn {
			break
		}

		select {
		case <-time.After(5 * time.Second):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}
