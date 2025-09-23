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
	"github.com/spf13/cobra"

	"github.com/galactica-corp/guardians-sdk/v3/internal/cli"
	"github.com/galactica-corp/guardians-sdk/v3/pkg/contracts"
	"github.com/galactica-corp/guardians-sdk/v3/pkg/merkle"
	"github.com/galactica-corp/guardians-sdk/v3/pkg/zkcertificate"
)

const (
	// Registry operation constants
	OperationAddition   = 0
	OperationRevocation = 1
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

// MerkleProofClient combines both EmptyLeafProver and Prover interfaces
type MerkleProofClient interface {
	merkle.EmptyLeafProver
	merkle.Prover
}

// IssueZKCert registers a zero-knowledge certificate (ZKCert) to the on-chain queue.
//
// The function performs the following steps:
//  1. Verifies that the provider is authorized as a guardian.
//  2. Adds the certificate to the queue for processing.
//  3. Returns the queue position to track the certificate status.
func IssueZKCert[T zkcertificate.Content](
	ctx context.Context,
	cert zkcertificate.Certificate[T],
	ethRPC EthereumIssueClient,
	merkleProofClient MerkleProofClient,
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

	// Add certificate to queue
	auth, err := bind.NewKeyedTransactorWithChainID(providerKey, chainID)
	if err != nil {
		return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("create transaction signer from private key: %w", err)
	}
	auth.Context = ctx

	var tx *types.Transaction

	// Handle different registry types
	if cert.Standard == zkcertificate.StandardKYC {
		// For zkKYC, we need to use the special registry with additional parameters
		kycRegistry, err := contracts.NewZkKYCRegistry(registryAddress, ethRPC)
		if err != nil {
			return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("load kyc registry: %w", err)
		}

		// Get KYC data from content to get the ID hash
		kycContent := any(cert.Content).(zkcertificate.KYCContent)

		idHash, err := kycContent.IDHash()
		if err != nil {
			return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("get id hash: %w", err)
		}

		// Add to queue with KYC-specific parameters
		tx, err = kycRegistry.AddOperationToQueue(
			auth,
			leafHash.Bytes32(),
			OperationAddition,
			idHash.BigInt(),
			cert.HolderCommitment.BigInt(), // salt hash
			big.NewInt(cert.ExpirationDate.Unix()),
		)
		if err != nil {
			return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("add kyc operation to queue: %w", err)
		}
	} else {
		// For other certificate types, use the standard registry
		tx, err = registry.AddOperationToQueue(auth, leafHash.Bytes32(), OperationAddition)
		if err != nil {
			return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("add operation to queue: %w", err)
		}
	}

	if receipt, err := bind.WaitMined(ctx, ethRPC, tx); err != nil {
		return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("wait until transaction is mined: %w", err)
	} else if receipt.Status == 0 {
		return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("transaction %q failed", receipt.TxHash)
	}

	// Get queue position
	var queuePosition *big.Int
	if cert.Standard == zkcertificate.StandardKYC {
		kycRegistry, err := contracts.NewZkKYCRegistry(registryAddress, ethRPC)
		if err != nil {
			return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("load kyc registry for queue position: %w", err)
		}
		queuePosition, err = kycRegistry.GetQueuePosition(&bind.CallOpts{Context: ctx}, leafHash.Bytes32())
		if err != nil {
			return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("get kyc queue position: %w", err)
		}
	} else {
		queuePosition, err = registry.GetQueuePosition(&bind.CallOpts{Context: ctx}, leafHash.Bytes32())
		if err != nil {
			return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("get queue position: %w", err)
		}
	}

	// Wait for certificate to be processed from the queue
	fmt.Fprintf(os.Stderr, "Certificate queued at position %s. Waiting for processing...\n", queuePosition)

	// Poll until the certificate is processed
	processed := false

	for !processed {
		// Get current queue pointer
		var currentPointer *big.Int
		if cert.Standard == zkcertificate.StandardKYC {
			kycRegistry, err := contracts.NewZkKYCRegistry(registryAddress, ethRPC)
			if err != nil {
				return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("load kyc registry for queue check: %w", err)
			}
			currentPointer, err = kycRegistry.CurrentQueuePointer(&bind.CallOpts{Context: ctx})
			if err != nil {
				return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("get current queue pointer: %w", err)
			}
		} else {
			currentPointer, err = registry.CurrentQueuePointer(&bind.CallOpts{Context: ctx})
			if err != nil {
				return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("get current queue pointer: %w", err)
			}
		}

		// Check if our certificate has been processed
		if currentPointer.Cmp(queuePosition) > 0 {
			// Our certificate should have been processed
			// Get the certificate data to confirm
			if cert.Standard == zkcertificate.StandardKYC {
				kycRegistry, err := contracts.NewZkKYCRegistry(registryAddress, ethRPC)
				if err != nil {
					return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("load kyc registry for status check: %w", err)
				}
				certData, err := kycRegistry.ZkCertificateProcessingData(&bind.CallOpts{Context: ctx}, leafHash.Bytes32())
				if err != nil {
					return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("get certificate processing data: %w", err)
				}
				// Check if state is processed
				if certData.State > 0 {
					processed = true
				}
			} else {
				certData, err := registry.ZkCertificateProcessingData(&bind.CallOpts{Context: ctx}, leafHash.Bytes32())
				if err != nil {
					return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("get certificate processing data: %w", err)
				}
				// Check if state is processed
				if certData.State > 0 {
					processed = true
				}
			}
		}

		if !processed {
			fmt.Fprintf(os.Stderr, "Current queue position: %s, waiting for position: %s\n", currentPointer, queuePosition)
			time.Sleep(5 * time.Second)
		}
	}

	// Get Merkle proof for the processed certificate
	proof, err := merkle.GetProof(ctx, merkleProofClient, registryAddress.Hex(), leafHash.String())
	if err != nil {
		return nil, zkcertificate.IssuedCertificate[T]{}, fmt.Errorf("get merkle proof: %w", err)
	}

	return tx, zkcertificate.IssuedCertificate[T]{
		Certificate: cert,
		Registration: zkcertificate.RegistrationDetails{
			Address:       registryAddress,
			ChainID:       chainID,
			Revocable:     true,
			QueuePosition: queuePosition,
			LeafIndex:     proof.LeafIndex,
		},
		MerkleProof: &proof,
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
