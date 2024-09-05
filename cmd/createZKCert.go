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
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/galactica-corp/guardians-sdk/pkg/keymanagement"
	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

type createZKCertFlags struct {
	holderFilePath            string
	certificateStandard       string
	certificateInputsFilePath string
	expirationDate            string
	providerPrivateKeyPath    string
	outputFilePath            string
}

func NewCmdCreateZKCert() *cobra.Command {
	var f createZKCertFlags

	cmd := &cobra.Command{
		Use:   "createZKCert",
		Short: "Create Zero Knowledge Certificate (ZKCert) by signing certificate data with an EdDSA key",
		Long: `The createZKCert command enables you to create Zero Knowledge Certificates (ZKCert)
by signing certificate data with an EdDSA key. ZKCert is a crucial component for
providing trust and privacy in the Galactica ecosystem.

ZKCert creation involves multiple steps, including specifying a certificate
standard, providing a holder commitment, and supplying certificate inputs.
The chosen certificate standard defines the required fields in the certificate,
ensuring adherence to specific criteria. The holder commitment links the ZKCert
to a holder's address without revealing this connection to the holder.

The signature is a fundamental requirement, and it must be generated using a
private key. The EdDSA key used for signing the certificate data is loaded from
a specified private key file. Additionally, a random salt is used to enhance
security.

Once all the necessary components are in place, the ZKCert is created and saved
to a JSON file for further use in the Galactica ecosystem.

Example Usage:
$ galactica-guardian createZKCert -s standardA -H holder_commitment.json -i certificate_inputs.json -e 2024-12-31T00:00:00.000Z -k provider_private_key.hex -o output.json`,
		RunE: createZKCertCmd(&f),
	}

	cmd.Flags().StringVarP(&f.holderFilePath, "holder-commitment-file", "H", "", "path to a file containing holder commitment")
	cmd.Flags().StringVarP(&f.certificateStandard, "standard", "s", zkcertificate.StandardKYC.String(), `standard identifies the type of zkCert`)
	cmd.Flags().StringVarP(&f.certificateInputsFilePath, "certificate-inputs-file", "i", "", "path to a JSON file with inputs for a zkCert. Which specific input fields are required depends on the certificate type")
	cmd.Flags().StringVarP(&f.expirationDate, "expiration-date", "e", "", "expiration date for the certificate in RFC3339 format")
	cmd.Flags().StringVarP(&f.providerPrivateKeyPath, "provider-private-key", "k", "", "path to a file containing provider's hex-encoded EdDSA private key")
	cmd.Flags().StringVarP(&f.outputFilePath, "output-file", "o", "certificate.json", "path to a file where the certificate in JSON format should be saved")

	_ = cmd.MarkFlagRequired("holder-commitment-file")
	_ = cmd.MarkFlagRequired("certificate-inputs-file")
	_ = cmd.MarkFlagRequired("expiration-date")
	_ = cmd.MarkFlagRequired("provider-private-key")

	return cmd
}

func createZKCertCmd(f *createZKCertFlags) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return createZKCert(f)
	}
}

func createZKCert(f *createZKCertFlags) error {
	var standard zkcertificate.Standard
	if err := standard.UnmarshalText([]byte(f.certificateStandard)); err != nil {
		return fmt.Errorf("parse certificate standard: %w", err)
	}

	var holderCommitment zkcertificate.HolderCommitment
	if err := decodeJSONFile(f.holderFilePath, &holderCommitment); err != nil {
		return fmt.Errorf("read holder commitment: %w", err)
	}

	expirationDate, err := time.Parse(time.RFC3339, f.expirationDate)
	if err != nil {
		return fmt.Errorf("invalid expiration date: %w", err)
	}

	certificateContent, err := readCertificateContent(f.certificateInputsFilePath, standard)
	if err != nil {
		return fmt.Errorf("read certificate content: %w", err)
	}

	contentHash, err := certificateContent.Hash()
	if err != nil {
		return fmt.Errorf("hash certificate content: %w", err)
	}

	providerKey, err := keymanagement.LoadEdDSA(f.providerPrivateKeyPath)
	if err != nil {
		return fmt.Errorf("load provider private key: %w", err)
	}

	signature, err := zkcertificate.SignCertificate(providerKey, contentHash, holderCommitment.CommitmentHash)
	if err != nil {
		return fmt.Errorf("sign certificate: %w", err)
	}

	salt, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64)) // [0, MaxInt64)
	if err != nil {
		return fmt.Errorf("generate random salt: %w", err)
	}

	randomSalt := salt.Int64() + 1 // [1, MaxInt64]

	certificate, err := zkcertificate.New(
		holderCommitment.CommitmentHash,
		certificateContent,
		providerKey.Public(),
		signature,
		randomSalt,
		expirationDate,
	)
	if err != nil {
		return fmt.Errorf("create certificate: %w", err)
	}

	if err := encodeToJSONFile(f.outputFilePath, certificate); err != nil {
		return fmt.Errorf("save certificate: %w", err)
	}

	_, _ = fmt.Fprintln(os.Stderr, "Saved certificate JSON to", f.outputFilePath)

	return nil
}

func readCertificateContent(filePath string, standard zkcertificate.Standard) (zkcertificate.Content, error) {
	switch standard {
	case zkcertificate.StandardKYC:
		var inputs zkcertificate.KYCInputs
		if err := decodeJSONFile(filePath, &inputs); err != nil {
			return nil, fmt.Errorf("read kyc inputs: %w", err)
		}

		certificateContent, err := inputs.FFEncode()
		if err != nil {
			return nil, fmt.Errorf("encode inputs to finite field: %w", err)
		}

		return certificateContent, nil
	case zkcertificate.StandardSimpleJSON:
		var inputs zkcertificate.SimpleJSON
		if err := decodeJSONFile(filePath, &inputs); err != nil {
			return nil, fmt.Errorf("read simple json: %w", err)
		}

		certificateContent, err := inputs.FFEncode()
		if err != nil {
			return nil, fmt.Errorf("encode inputs to finite field: %w", err)
		}

		return certificateContent, nil
	case zkcertificate.StandardREY:
		var inputs zkcertificate.REYInputs
		if err := decodeJSONFile(filePath, &inputs); err != nil {
			return nil, fmt.Errorf("read rey inputs: %w", err)
		}

		certificateContent, err := inputs.FFEncode()
		if err != nil {
			return nil, fmt.Errorf("encode inputs to finite field: %w", err)
		}

		return certificateContent, nil
	default:
		return nil, fmt.Errorf("standard %q is not supported", standard)
	}
}
