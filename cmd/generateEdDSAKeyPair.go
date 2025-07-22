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
	"crypto/ecdsa"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/iden3/go-iden3-crypto/v2/babyjub"
	"github.com/spf13/cobra"

	"github.com/galactica-corp/guardians-sdk/pkg/keymanagement"
)

type generateEdDSAKeyPairFlags struct {
	privateKeyPath string
	outputFilePath string
}

func NewCmdGenerateEdDSAKeyPair() *cobra.Command {
	var f generateEdDSAKeyPairFlags

	cmd := &cobra.Command{
		Use:   "generateEdDSAKeyPair",
		Short: "Generate EdDSA key pairs for managing zero knowledge certificates",
		Long: `The generateEdDSAKeyPair command allows you to generate EdDSA key pairs used for
managing zero knowledge certificates. EdDSA keys are essential for signing and
verifying data in Zero-Knowledge (ZK) circuits. This command provides the
flexibility to generate these key pairs, either by deriving them from an
Ethereum private key or by generating a random key pair.

When a specific Ethereum private key file is provided, the command derives an
EdDSA private key from it. If no Ethereum private key is provided, a random
EdDSA key pair is generated. The resulting EdDSA private key is saved to a
specified output file.

Example Usage:
$ galactica-guardian generateEdDSAKeyPair -k /path/to/ethereum-key.hex -o /path/to/output.hex`,
		RunE: generateEdDSAKeyPairCmd(&f),
	}

	cmd.Flags().StringVarP(&f.privateKeyPath, "private-key-file", "k", "", "path to a file containing a hex-encoded Ethereum (ECDSA) private key")
	cmd.Flags().StringVarP(&f.outputFilePath, "output-file", "o", "eddsa-private-key.hex", "path to a file where generated private key should be saved")

	return cmd
}

func generateEdDSAKeyPairCmd(f *generateEdDSAKeyPairFlags) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return generateEdDSAKeyPair(f)
	}
}

func generateEdDSAKeyPair(f *generateEdDSAKeyPairFlags) error {
	var ethereumPrivateKey *ecdsa.PrivateKey
	if f.privateKeyPath != "" {
		var err error
		ethereumPrivateKey, err = crypto.LoadECDSA(f.privateKeyPath)
		if err != nil {
			return fmt.Errorf("load ethereum private key: %w", err)
		}
	}

	privateKey, err := GenerateEdDSAKeyPair(ethereumPrivateKey)
	if err != nil {
		return fmt.Errorf("generate eddsa: %w", err)
	}

	if err := keymanagement.SaveEdDSA(f.outputFilePath, privateKey); err != nil {
		return fmt.Errorf("save eddsa private key: %w", err)
	}

	_, _ = fmt.Fprintln(os.Stderr, "Saved EdDSA private key to", f.outputFilePath)

	publicKey := privateKey.Public()

	_, _ = fmt.Fprintln(os.Stderr, "EdDSA public key", publicKey.X, publicKey.Y)

	return nil
}

// GenerateEdDSAKeyPair generates an EdDSA key pair.
//
// The function performs the following steps:
//  1. If an Ethereum private key is provided, derives the EdDSA private key from it.
//  2. If no Ethereum private key is provided, generates a random EdDSA private key.
func GenerateEdDSAKeyPair(ethereumPrivateKey *ecdsa.PrivateKey) (babyjub.PrivateKey, error) {
	if ethereumPrivateKey == nil {
		return babyjub.NewRandPrivKey()
	}

	return keymanagement.DeriveEdDSAKeyFromEthereumPrivateKey(ethereumPrivateKey)
}
