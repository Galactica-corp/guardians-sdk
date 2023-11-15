/*
Copyright © 2023 Galactica Network

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/galactica-corp/guardians-sdk/internal/cli"
	"github.com/galactica-corp/guardians-sdk/pkg/merkle"
	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

type revokeZKCertFlags struct {
	certificateFilePath    string
	firstBlock             int64
	providerPrivateKeyPath string
	registryAddress        cli.Address
	rpcURL                 string
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
$ galactica-guardian revokeZKCert -c zkcert.json -k provider_private_key.hex -r registry_address`,
		RunE: revokeZKCertCmd(&f),
	}

	cmd.Flags().StringVarP(&f.certificateFilePath, "certificate-file", "c", "", "path to a file containing issued zkCert obtained using issueZKCert command")
	cmd.Flags().VarP(&f.registryAddress, "registry-address", "r", "Ethereum address of the registry contract on-chain")
	cmd.Flags().Int64VarP(&f.firstBlock, "registry-events-start", "", 0, "block number in which first event was emitted by the registry. This block might be before the first event, but if it will be after, then it will lead to incorrect result. It greatly improves time to build a merkle tree, because RPC requests are limited to inspect at most 10'000 blocks at once")
	cmd.Flags().StringVarP(&f.providerPrivateKeyPath, "provider-private-key", "k", "", "path to a file containing provider's hex-encoded Ethereum (ECDSA) private key to sign the transaction")
	cmd.Flags().StringVarP(&f.rpcURL, "rpc-url", "", "", "url of Ethereum compatible RPC endpoint")

	_ = cmd.MarkFlagRequired("certificate-file")
	_ = cmd.MarkFlagRequired("provider-private-key")
	_ = cmd.MarkFlagRequired("registry-address")
	_ = cmd.MarkFlagRequired("rpc-url")

	return cmd
}

func revokeZKCertCmd(f *revokeZKCertFlags) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return revokeZKCert(f)
	}
}

func revokeZKCert(f *revokeZKCertFlags) error {
	ctx := context.Background()

	var certificate zkcertificate.IssuedCertificate[json.RawMessage]
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

	providerAddress := crypto.PubkeyToAddress(providerKey.PublicKey)

	if err := ensureProviderIsGuardian(client, registry, providerAddress); err != nil {
		return fmt.Errorf("ensure provider is guardian: %w", err)
	}

	leafIndex := certificate.Registration.LeafIndex
	leafHash := certificate.LeafHash

	proof, err := buildMerkleProof(ctx, client, registryAddress, registry, leafIndex, leafHash, f.firstBlock)
	if err != nil {
		return fmt.Errorf("build merkle proof: %w", err)
	}

	tx, err := constructRevokeZKCertTx(ctx, client, providerKey, registry, leafIndex, leafHash, proof)
	if err != nil {
		return fmt.Errorf("construct transaction to revoke record from registry: %w", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(tx); err != nil {
		return fmt.Errorf("encode revocation transaction to json: %w", err)
	}

	return nil
}

func constructRevokeZKCertTx(
	ctx context.Context,
	client *ethclient.Client,
	providerKey *ecdsa.PrivateKey,
	recordRegistry RecordRegistry,
	leafIndex int,
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

	auth.NoSend = true

	return recordRegistry.RevokeZkKYCRecord(
		auth,
		big.NewInt(int64(leafIndex)),
		leafHash.Bytes32(),
		encodeMerkleProof(proof),
	)
}

func buildMerkleProof(
	ctx context.Context,
	client *ethclient.Client,
	registryAddress common.Address,
	registry RegistryEventParser,
	leafIndex int,
	leafHash zkcertificate.Hash,
	firstBlock int64,
) (merkle.Proof, error) {
	tree, err := buildMerkleTreeFromEvents(ctx, client, registryAddress, registry, firstBlock)
	if err != nil {
		return merkle.Proof{}, fmt.Errorf("build merkle tree from events: %w", err)
	}

	proof, err := tree.GetProof(leafIndex)
	if err != nil {
		return merkle.Proof{}, err
	}

	if proof.Path[0].Value.ToBig().Cmp(leafHash.BigInt()) != 0 {
		return merkle.Proof{}, fmt.Errorf("incorrect leaf hash at specfied leaf index")
	}

	return proof, nil
}
