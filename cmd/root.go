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
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "galactica-guardian",
		Short: "galactica-guardian is a CLI tool for managing and verifying user data on the Galactica blockchain",
		Long: `galactica-guardian is a powerful command-line interface (CLI) tool designed for
guardians in the Galactica blockchain ecosystem. Guardians play a pivotal role
in verifying and managing user data on the blockchain. 

In the initial stage, galactica-guardian primarily focuses on Know Your
Customer (KYC) verification. Users can visit the Galactica website or app to
input their personal information. Guardians then perform document verification
using national IDs, passports, or driver's licenses. Once verified, guardians
issue zkKYC attestations on-chain, which are stored in a Merkle tree registry.

As Galactica evolves, guardians will have the capability to issue various
other attestations, such as ZkCerts for education diplomas, gym memberships,
conference tickets, and more. This CLI empowers guardians to efficiently
manage these verification processes and maintain a secure and trusted
ecosystem.

For detailed usage instructions, available commands, and options, please refer
to the respective sections in the documentation.
`,
	}

	cmd.AddCommand(
		NewCmdGenerateEdDSAKeyPair(),
		NewCmdCreateZKCert(),
		NewCmdIssueZKCert(),
		NewCmdEncryptZKCert(),
		NewCmdRevokeZKCert(),
		NewCmdRenewZKCert(),
		NewCmdPrintEdDSAPublicKey(),
	)

	return cmd
}

func encodeToJSONFile(filePath string, target any) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(target); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

func decodeJSONFile(filePath string, target any) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(target); err != nil {
		return fmt.Errorf("decode json: %w", err)
	}

	return nil
}

func deserializeCertificateJSON(filePath string) (zkcertificate.Certificate[zkcertificate.Content], error) {
	f, err := os.Open(filePath)
	if err != nil {
		return zkcertificate.Certificate[zkcertificate.Content]{}, fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	return zkcertificate.DeserializeCertificateJSON(f)
}

func deserializeIssuedCertificateJSON(filePath string) (zkcertificate.IssuedCertificate[zkcertificate.Content], error) {
	f, err := os.Open(filePath)
	if err != nil {
		return zkcertificate.IssuedCertificate[zkcertificate.Content]{}, fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	return zkcertificate.DeserializeIssuedCertificateJSON(f)
}
