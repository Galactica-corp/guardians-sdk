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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"

	"github.com/galactica-corp/guardians-sdk/pkg/contracts"
	"github.com/galactica-corp/guardians-sdk/pkg/merkle"
	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

type revokeZKCertFlags struct {
	certificateFilePath    string
	providerPrivateKeyPath string
	rpcURL                 string
	merkleProofServiceURL  string
	merkleProofServiceTLS  bool
}

func NewCmdRevokeZKCert() *cobra.Command {
	var f revokeZKCertFlags

	cmd := &cobra.Command{
		Use:   "revokeZKCert",
		Short: "Revoke a Zero Knowledge Certificate (ZKCert) from the Galactica blockchain registry",
		Long: `The revokeZKCert command allows you to revoke a Zero Knowledge Certificate (ZKCert)
from the Galactica blockchain registry. Revoking a ZKCert is an essential step in
maintaining the integrity of the registry and responding to counter fraudulent
requests.

This command involves connecting to the blockchain RPC, loading the relevant
record registry, and ensuring that the provider initiating the revocation is a
verified guardian. The ZKCert to be revoked is identified by its leaf index and
hash, and a Merkle proof is constructed for validation.

Upon successful execution, the command outputs the constructed revocation
transaction, that provider needs to send manualy.

Example Usage:
$ galactica-guardian revokeZKCert -c zkcert.json -k provider_private_key.hex`,
		RunE: revokeZKCertCmd(&f),
	}

	cmd.Flags().StringVarP(&f.certificateFilePath, "certificate-file", "c", "", "path to a file containing issued zkCert obtained using issueZKCert command")
	cmd.Flags().StringVarP(&f.providerPrivateKeyPath, "provider-private-key", "k", "", "path to a file containing provider's hex-encoded Ethereum (ECDSA) private key to sign the transaction")
	cmd.Flags().StringVarP(&f.rpcURL, "rpc-url", "", "", "url of Ethereum compatible RPC endpoint")
	cmd.Flags().StringVarP(&f.merkleProofServiceURL, "merkle-proof-service-url", "m", "", "Merkle Proof Service gRPC endpoint url")
	cmd.Flags().BoolVar(&f.merkleProofServiceTLS, "merkle-proof-service-tls", false, "enable TLS for Merkle Proof Service gRPC connection")

	_ = cmd.MarkFlagRequired("certificate-file")
	_ = cmd.MarkFlagRequired("provider-private-key")
	_ = cmd.MarkFlagRequired("rpc-url")
	_ = cmd.MarkFlagRequired("merkle-proof-service-url")

	return cmd
}

func revokeZKCertCmd(f *revokeZKCertFlags) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return revokeZKCert(f)
	}
}

func revokeZKCert(f *revokeZKCertFlags) error {
	ctx := context.Background()

	certificate, err := deserializeIssuedCertificateJSON(f.certificateFilePath)
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

	providerKey, err := crypto.LoadECDSA(f.providerPrivateKeyPath)
	if err != nil {
		return fmt.Errorf("load provider's ethereum private key: %w", err)
	}

	tx, err := RevokeZKCert(ctx, certificate, client, merkleProofClient, providerKey)
	if err != nil {
		return fmt.Errorf("revoke zk cert: %w", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(tx); err != nil {
		return fmt.Errorf("encode revocation transaction to json: %w", err)
	}

	return nil
}

type EthereumRevokeClient interface {
	bind.ContractBackend
	bind.DeployBackend

	ChainID(ctx context.Context) (*big.Int, error)
}

// RevokeZKCert revokes an issued zero-knowledge certificate (ZKCert) by removing it from the blockchain registry.
//
// The function performs the following steps:
//  1. Verifies that the provider is authorized as a guardian.
//  2. Retrieves the Merkle proof for the certificate to be revoked.
//  3. Registers the provider for revocation and waits for its turn.
//  4. Constructs and sends the transaction to revoke the certificate.
func RevokeZKCert[T zkcertificate.Content](
	ctx context.Context,
	certificate zkcertificate.IssuedCertificate[T],
	client EthereumRevokeClient,
	merkleProofClient merkle.Prover,
	providerKey *ecdsa.PrivateKey,
) (*types.Transaction, error) {
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain id from blockchain rpc: %w", err)
	}

	registryAddress := certificate.Registration.Address

	registry, err := contracts.NewZkCertificateRegistry(registryAddress, client)
	if err != nil {
		return nil, fmt.Errorf("load record registry: %w", err)
	}

	providerAddress := crypto.PubkeyToAddress(providerKey.PublicKey)

	if err := ensureProviderIsGuardian(ctx, client, registry, providerAddress); err != nil {
		return nil, fmt.Errorf("ensure provider is guardian: %w", err)
	}

	leafHash := certificate.LeafHash

	if err := registerAndWaitForZkCertificateTurn(ctx, client, chainID, providerKey, registry, leafHash); err != nil {
		return nil, fmt.Errorf("register and wait for zkCertificate turn: %w", err)
	}

	proof, err := merkle.GetProof(ctx, merkleProofClient, registryAddress.Hex(), leafHash.String())
	if err != nil {
		return nil, fmt.Errorf("get merkle proof: %w", err)
	}

	leafIndex := certificate.Registration.LeafIndex

	tx, err := constructRevokeZKCertTx(ctx, chainID, providerKey, registry, leafIndex, leafHash, proof)
	if err != nil {
		return nil, fmt.Errorf("construct transaction to revoke record from registry: %w", err)
	}

	if receipt, err := bind.WaitMined(ctx, client, tx); err != nil {
		return nil, fmt.Errorf("wait until transaction is mined: %w", err)
	} else if receipt.Status == 0 {
		return nil, fmt.Errorf("transaction %q failed", receipt.TxHash)
	}

	return tx, nil
}

type RecordRegistryCertificateRevoker interface {
	RevokeZkCertificate(
		opts *bind.TransactOpts,
		leafIndex *big.Int,
		zkCertificateHash [32]byte,
		merkleProof [][32]byte,
	) (*types.Transaction, error)
}

func constructRevokeZKCertTx(
	ctx context.Context,
	chainID *big.Int,
	providerKey *ecdsa.PrivateKey,
	recordRegistry RecordRegistryCertificateRevoker,
	leafIndex int,
	leafHash zkcertificate.Hash,
	proof merkle.Proof,
) (*types.Transaction, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(providerKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("create transaction signer from private key: %w", err)
	}

	auth.Context = ctx

	return recordRegistry.RevokeZkCertificate(
		auth,
		big.NewInt(int64(leafIndex)),
		leafHash.Bytes32(),
		encodeMerkleProof(proof),
	)
}
