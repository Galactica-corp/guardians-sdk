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
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/galactica-corp/guardians-sdk/pkg/keymanagement"
	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

type renewZKCertFlags struct {
	certificateFilePath    string
	expirationDate         string
	outputFilePath         string
	providerPrivateKeyPath string
	randomSalt             int64
}

func NewCmdRenewZKCert() *cobra.Command {
	var f renewZKCertFlags

	cmd := &cobra.Command{
		Use:   "renewZKCert",
		Short: "Renew a Zero Knowledge Certificate (ZKCert) with an updated expiration date",
		Long: `The renewZKCert command enables you to renew a Zero Knowledge Certificate (ZKCert)
with an updated expiration date, ensuring the continued validity of the certificate
within the Galactica ecosystem.

This command involves specifying a new expiration date for the ZKCert and signing
the updated certificate with the provider's private key. The renewal process
requires loading the existing ZKCert, setting the new expiration date, and
constructing a new certificate with the updated information.

Upon successful renewal, the command outputs the prolonged ZKCert in JSON 
format. Additionally, a notification is provided to guide users on handling the
updated certificate.

Example Usage:
$ galactica-guardian renewZKCert -c zkcert.json -k provider_private_key.hex -e 2024-12-31T00:00:00.000Z -o prolonged_zkcert.json`,
		RunE: renewZKCertCmd(&f),
	}

	cmd.Flags().StringVarP(&f.certificateFilePath, "certificate-file", "c", "", "path to a file containing issued zkCert obtained using issueZKCert command")
	cmd.Flags().StringVarP(&f.expirationDate, "expiration-date", "e", "", "new expiration date for the certificate in RFC3339 format")
	cmd.Flags().StringVarP(&f.outputFilePath, "output-file", "o", "prolonged-certificate.json", "path to a file where the prolonged certificate in JSON format should be saved")
	cmd.Flags().StringVarP(&f.providerPrivateKeyPath, "provider-private-key", "k", "", "path to a file containing provider's hex-encoded EdDSA private key")
	cmd.Flags().Int64VarP(&f.randomSalt, "random-salt", "", 0, "random salt to input into zkCert hashing")

	_ = cmd.MarkFlagRequired("certificate-file")
	_ = cmd.MarkFlagRequired("expiration-date")
	_ = cmd.MarkFlagRequired("provider-private-key")

	return cmd
}

func renewZKCertCmd(f *renewZKCertFlags) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return renewZKCert(f)
	}
}

func renewZKCert(f *renewZKCertFlags) error {
	if f.randomSalt == 0 {
		salt, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64)) // [0, MaxInt64)
		if err != nil {
			return fmt.Errorf("generate random salt: %w", err)
		}

		f.randomSalt = salt.Int64() + 1 // [1, MaxInt64]
	}

	expirationDate, err := time.Parse(time.RFC3339, f.expirationDate)
	if err != nil {
		return fmt.Errorf("invalid expiration date: %w", err)
	}

	var certificate zkcertificate.Certificate[json.RawMessage]
	if err := decodeJSONFile(f.certificateFilePath, &certificate); err != nil {
		return fmt.Errorf("read certificate: %w", err)
	}

	certificateContent, err := setNewExpirationDate(certificate.Standard, certificate.Content, expirationDate)
	if err != nil {
		return fmt.Errorf("set new expiration date: %w", err)
	}

	contentHash, err := certificateContent.Hash()
	if err != nil {
		return fmt.Errorf("hash certificate content: %w", err)
	}

	providerKey, err := keymanagement.LoadEdDSA(f.providerPrivateKeyPath)
	if err != nil {
		return fmt.Errorf("load provider private key: %w", err)
	}

	signature, err := zkcertificate.SignCertificate(providerKey, contentHash, certificate.HolderCommitment)
	if err != nil {
		return fmt.Errorf("sign certificate: %w", err)
	}

	newCertificate, err := zkcertificate.NewCertificate(
		certificate.HolderCommitment,
		certificateContent,
		providerKey.Public(), // TODO: Can different provider renew the certificate?
		signature,
		f.randomSalt, // TODO: Should the salt be the same? It better not
	)
	if err != nil {
		return fmt.Errorf("create new certificate: %w", err)
	}

	if err := encodeToJSONFile(f.outputFilePath, newCertificate); err != nil {
		return fmt.Errorf("save prolonged certificate: %w", err)
	}

	_, _ = fmt.Fprintln(os.Stderr, "Saved certificate JSON to", f.outputFilePath)
	printRenewalInstruction(f.certificateFilePath, f.outputFilePath)

	return nil
}

func printRenewalInstruction(oldCertificatePath string, newCertificatePath string) {
	_, _ = fmt.Fprintf(
		os.Stderr,
		`Please, run the following commands to complete the renewal process:

galactica-guardian revokeZKCert -c %s -k provider_private_key.hex -r registry_address
galactica-guardian issueZKCert -c %s -k provider_private_key.hex -o issued-prolonger-certificate.json`,
		oldCertificatePath,
		newCertificatePath,
	)
}

func setNewExpirationDate(
	certificateStandard zkcertificate.Standard,
	certificateContent json.RawMessage,
	expirationDate time.Time,
) (zkcertificate.Content, error) {
	switch certificateStandard {
	case zkcertificate.StandardKYC:
		var content zkcertificate.KYCContent
		if err := json.Unmarshal(certificateContent, &content); err != nil {
			return nil, fmt.Errorf("decode kyc certificate content json: %w", err)
		}

		// TODO: Maybe check that the new expiration date is after the old one?
		content.ExpirationDate = zkcertificate.Timestamp(expirationDate)
		return content, nil
	default:
		return nil, fmt.Errorf("unsupported certificate standard %s", certificateStandard)
	}
}
