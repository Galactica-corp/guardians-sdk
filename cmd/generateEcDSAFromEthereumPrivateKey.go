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
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"

	"github.com/galactica-corp/guardians-sdk/pkg/keymanagement"
)

var (
	privateKeyPath string
)

// generateEcDSAFromEthereumPrivateKeyCmd represents the generateEcDSAFromEthereumPrivateKey command
var generateEcDSAFromEthereumPrivateKeyCmd = &cobra.Command{
	Use:   "generateEcDSAFromEthereumPrivateKey",
	Short: "Generate EdDSA key for signing data verified in ZK circuits from an Ethereum private key",
	Long: `The generateEcDSAFromEthereumPrivateKeyCmd command allows you to generate EdDSA
keys used for signing data that is subsequently verified in Zero-Knowledge (ZK)
circuits. This type of key generation is less common, but the SDK provides the
capability to create these keys. Guardians, in particular, may find it convenient
to derive EdDSA keys from a seed that corresponds to a message signed with an
Ethereum account.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var privateKey *ecdsa.PrivateKey

		if privateKeyPath != "" {
			data, err := os.ReadFile(privateKeyPath)
			if err != nil {
				return fmt.Errorf("read private key: %w", err)
			}

			data = bytes.TrimPrefix(data, []byte("0x"))

			privateKey, err = crypto.HexToECDSA(string(data))
			if err != nil {
				return fmt.Errorf("parse private key: %w", err)
			}
		} else {
			var err error
			privateKey, err = crypto.GenerateKey()
			if err != nil {
				return fmt.Errorf("generate random private key: %w", err)
			}
		}

		edDSAKey, err := keymanagement.GetEdDSAKeyFromEthereumSigner(privateKey)
		if err != nil {
			return fmt.Errorf("get eddsa key from ethereum signer: %w", err)
		}

		fmt.Printf("0x%x", edDSAKey)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateEcDSAFromEthereumPrivateKeyCmd)

	generateEcDSAFromEthereumPrivateKeyCmd.Flags().StringVarP(&privateKeyPath, "private-key-file", "k", "", "path to a file containing a hex-encoded Ethereum private key")
}
