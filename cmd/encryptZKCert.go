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
	"os"

	"github.com/spf13/cobra"

	"github.com/galactica-corp/guardians-sdk/pkg/encryption"
	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

type encryptZKCertFlags struct {
	certificateFilePath string
	holderFilePath      string
	outputFilePath      string
}

func NewCmdEncryptZKCert() *cobra.Command {
	var f encryptZKCertFlags

	cmd := &cobra.Command{
		Use:   "encryptZKCert",
		Short: "Encrypt a Zero Knowledge Certificate (ZKCert) with a holder's encryption key",
		Long: `The encryptZKCert command allows you to encrypt a Zero Knowledge Certificate (ZKCert)
with a holder's encryption key to the format compatible with MetaMask snap.

This command requires the ZKCert in JSON format, specified by the -c or --certificate-file option,
and the holder's commitment information, provided by the -H or --holder-file option.
The holder's encryption key, expected to be 32 bytes long, is crucial for the encryption process.

Upon successful encryption, the command outputs the encrypted certificate and the
associated holder's commitment. The encrypted ZKCert is saved to a specified output file.

Example Usage:
$ galactica-guardian encryptZKCert -c zkcert.json -H holder_commitment.json -o encrypted.json`,
		RunE: encryptZKCertCmd(&f),
	}

	cmd.Flags().StringVarP(&f.certificateFilePath, "certificate-file", "c", "", "path to a file containing issued zkCert obtained using issueZKCert command")
	cmd.Flags().StringVarP(&f.holderFilePath, "holder-commitment-file", "H", "", "path to a file containing holder commitment")
	cmd.Flags().StringVarP(&f.outputFilePath, "output-file", "o", "encrypted-certificate.json", "path to a file where the encrypted certificate in JSON format should be saved")

	_ = cmd.MarkFlagRequired("certificate-file")
	_ = cmd.MarkFlagRequired("holder-commitment-file")

	return cmd
}

func encryptZKCertCmd(f *encryptZKCertFlags) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return encryptZKCert(f)
	}
}

func encryptZKCert(f *encryptZKCertFlags) error {
	certificate, err := deserializeCertificateJSON(f.certificateFilePath)
	if err != nil {
		return fmt.Errorf("read certificate: %w", err)
	}

	var holderCommitment zkcertificate.HolderCommitment
	if err := decodeJSONFile(f.holderFilePath, &holderCommitment); err != nil {
		return fmt.Errorf("read holder commitment: %w", err)
	}

	encryptedCertificate, err := EncryptZKCert(certificate, holderCommitment)
	if err != nil {
		return fmt.Errorf("encrypt certificate: %w", err)
	}

	if err := encodeToJSONFile(f.outputFilePath, encryptedCertificate); err != nil {
		return fmt.Errorf("save encrypted certificate: %w", err)
	}

	_, _ = fmt.Fprintln(os.Stderr, "Saved encrypted certificate to", f.outputFilePath)

	return nil
}

// EncryptZKCert encrypts a zero-knowledge certificate using the holder's commitment encryption key.
func EncryptZKCert[T zkcertificate.Content](
	certificate zkcertificate.Certificate[T],
	holderCommitment zkcertificate.HolderCommitment,
) (zkcertificate.EncryptedCertificate, error) {
	encryptedMessage, err := encryption.EncryptWithPadding([32]byte(holderCommitment.EncryptionKey), certificate)
	if err != nil {
		return zkcertificate.EncryptedCertificate{}, fmt.Errorf("encrypt certificate: %w", err)
	}

	return zkcertificate.EncryptedCertificate{
		EncryptedMessage: encryptedMessage,
		HolderCommitment: holderCommitment.CommitmentHash,
	}, nil
}
