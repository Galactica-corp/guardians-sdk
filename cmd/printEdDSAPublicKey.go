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
	"fmt"
	"io"
	"os"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/spf13/cobra"

	"github.com/galactica-corp/guardians-sdk/pkg/keymanagement"
)

type printEdDSAPublicKeyFlags struct {
	privateKeyPath string
}

func NewCmdPrintEdDSAPublicKey() *cobra.Command {
	var f printEdDSAPublicKeyFlags

	cmd := &cobra.Command{
		Use:   "printEdDSAPublicKey -k eddsa-private-key.hex",
		Short: "Print EdDSA public key for a given private key",
		Long: `The printEdDSAPublicKey command allows you to retrieve and print the EdDSA public key
corresponding to a given EdDSA private key. The command extracts the public key from a specified
private key file.

To use this command, you must provide a file path to an EdDSA private key, from which the public key
will be derived and displayed.

Example Usage:
$ galactica-guardian printEdDSAPublicKey -k eddsa-private-key.hex`,
		RunE: printEdDSAPublicKey(&f),
	}

	cmd.Flags().StringVarP(&f.privateKeyPath, "provider-private-key", "k", "", "path to a file containing provider's hex-encoded EdDSA private key")

	_ = cmd.MarkFlagRequired("provider-private-key")

	return cmd
}

func printEdDSAPublicKey(f *printEdDSAPublicKeyFlags) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		privateKey, err := keymanagement.LoadEdDSA(f.privateKeyPath)
		if err != nil {
			return fmt.Errorf("load private key: %w", err)
		}

		return PrintEdDSAPublicKey(privateKey, os.Stdout)
	}
}

func PrintEdDSAPublicKey(privateKey babyjub.PrivateKey, w io.Writer) error {
	publicKey := privateKey.Public()

	if _, err := fmt.Fprintln(w, publicKey); err != nil {
		return fmt.Errorf("print public key: %w", err)
	}

	return nil
}
