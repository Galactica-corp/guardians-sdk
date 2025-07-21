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

package zkcertificate

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/iden3/go-iden3-crypto/utils"

	"github.com/galactica-corp/guardians-sdk/pkg/encryption"
	"github.com/galactica-corp/guardians-sdk/pkg/merkle"
)

var (
	eddsaPrimeFieldMod = utils.NewIntFromString("2736030358979909402780800718157159386076813972158567259200215660948447373040")
)

// Certificate represents a zero knowledge certificate structure that can hold different types of content.
// It is parameterized by the type T for the content field.
// Certificate content must be directly determined by the certificate Standard.
type Certificate[T Content] struct {
	HolderCommitment Hash         `json:"holderCommitment"`
	LeafHash         Hash         `json:"leafHash"`
	DID              string       `json:"did"`
	Standard         Standard     `json:"zkCertStandard"`
	Content          T            `json:"content"`
	ContentHash      Hash         `json:"contentHash"`
	ExpirationDate   Timestamp    `json:"expirationDate"`
	Provider         ProviderData `json:"providerData"`
	RandomSalt       string       `json:"randomSalt"`
}

// ProviderData represents the public key and signature data of a certificate provider.
type ProviderData struct {
	PublicKey babyjub.PublicKey
	Signature babyjub.Signature
}

// Content is an interface that represents the content of a certificate.
// It defines methods for calculating the content's hash and obtaining the standard it adheres to.
type Content interface {
	// Hash computes and returns the Poseidon hash of the certificate content.
	Hash() (Hash, error)
	// Standard returns the standard to which the certificate content adheres.
	Standard() Standard
}

// New creates a new certificate instance with the provided parameters and content.
// It computes the content hash, verifies if the content was actually signed with providers public key,
// and generates a leaf hash.
func New[T Content](
	holderCommitment Hash,
	content T,
	providerPublicKey *babyjub.PublicKey,
	providerSignature *babyjub.Signature,
	salt int64,
	expirationDate time.Time,
) (*Certificate[T], error) {
	contentHash, err := content.Hash()
	if err != nil {
		return nil, fmt.Errorf("hash certificate content: %w", err)
	}

	signatureValid, err := VerifySignature(providerPublicKey, contentHash, holderCommitment, providerSignature)
	if err != nil {
		return nil, fmt.Errorf("verify signature: %w", err)
	}
	if !signatureValid {
		return nil, fmt.Errorf("invalid signature")
	}

	leafHash, err := LeafHash(contentHash, providerPublicKey, providerSignature, holderCommitment, salt, expirationDate)
	if err != nil {
		return nil, fmt.Errorf("compute leaf hash: %w", err)
	}

	standard := content.Standard()

	return &Certificate[T]{
		HolderCommitment: holderCommitment,
		LeafHash:         leafHash,
		DID:              DID(standard, leafHash),
		Standard:         standard,
		Content:          content,
		ContentHash:      contentHash,
		ExpirationDate:   Timestamp(expirationDate),
		Provider: ProviderData{
			PublicKey: *providerPublicKey,
			Signature: *providerSignature,
		},
		RandomSalt: strconv.FormatInt(salt, 10),
	}, nil
}

type providerDataDTO struct {
	Ax  string `json:"ax"`
	Ay  string `json:"ay"`
	S   string `json:"s"`
	R8x string `json:"r8x"`
	R8y string `json:"r8y"`
}

// MarshalJSON implements [json.Marshaler].
func (p ProviderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(providerDataDTO{
		Ax:  p.PublicKey.X.String(),
		Ay:  p.PublicKey.Y.String(),
		S:   p.Signature.S.String(),
		R8x: p.Signature.R8.X.String(),
		R8y: p.Signature.R8.Y.String(),
	})
}

// UnmarshalJSON implements [json.Unmarshaler].
func (p *ProviderData) UnmarshalJSON(data []byte) error {
	var dto providerDataDTO
	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	var ok bool

	p.PublicKey.X, ok = new(big.Int).SetString(dto.Ax, 10)
	if !ok {
		return fmt.Errorf("invalid x coordinate of public key point")
	}

	p.PublicKey.Y, ok = new(big.Int).SetString(dto.Ay, 10)
	if !ok {
		return fmt.Errorf("invalid y coordinate of public key point")
	}

	signatureR8Point := &babyjub.Point{}

	signatureR8Point.X, ok = new(big.Int).SetString(dto.R8x, 10)
	if !ok {
		return fmt.Errorf("invalid x coordinate of signature r8 point")
	}

	signatureR8Point.Y, ok = new(big.Int).SetString(dto.R8y, 10)
	if !ok {
		return fmt.Errorf("invalid y coordinate of signature r8 point")
	}

	p.Signature.R8 = signatureR8Point

	p.Signature.S, ok = new(big.Int).SetString(dto.S, 10)
	if !ok {
		return fmt.Errorf("invalid s component of signature")
	}

	return nil
}

// IssuedCertificate represents a certificate that has been issued and includes registration details.
type IssuedCertificate[T Content] struct {
	Certificate[T] `json:",inline"`
	Registration   RegistrationDetails `json:"registration"`
	MerkleProof    merkle.Proof        `json:"merkleProof"`
}

// RegistrationDetails represents details related to the registration of a certificate.
type RegistrationDetails struct {
	Address   common.Address `json:"address"`
	ChainID   *big.Int       `json:"chainID"`
	Revocable bool           `json:"revocable"`
	LeafIndex int            `json:"leafIndex"`
}

// prepareForEdDSA computes the Poseidon hash of the given inputs and reduces it to the field supported by EdDSA.
func prepareForEdDSA(inputs ...*big.Int) (*big.Int, error) {
	hash, err := poseidon.Hash(inputs)
	if err != nil {
		return nil, fmt.Errorf("compute poseidon hash: %w", err)
	}
	return hash.Mod(hash, eddsaPrimeFieldMod), nil
}

// SignCertificate generates a digital signature for a certificate using the provider's private key.
func SignCertificate(
	providerKey babyjub.PrivateKey,
	contentHash Hash,
	commitmentHash Hash,
) (*babyjub.Signature, error) {
	message, err := prepareForEdDSA(contentHash.BigInt(), commitmentHash.BigInt())
	if err != nil {
		return nil, fmt.Errorf("hash message: %w", err)
	}

	return providerKey.SignPoseidon(message), nil
}

// VerifySignature verifies the digital signature of a certificate using the provider's public key.
func VerifySignature(
	providerKey *babyjub.PublicKey,
	contentHash Hash,
	commitmentHash Hash,
	signature *babyjub.Signature,
) (bool, error) {
	message, err := prepareForEdDSA(contentHash.BigInt(), commitmentHash.BigInt())
	if err != nil {
		return false, fmt.Errorf("hash message: %w", err)
	}

	return providerKey.VerifyPoseidon(message, signature), nil
}

// LeafHash computes the hash of a certificate's components and additional data to create a leaf hash.
func LeafHash(
	contentHash Hash,
	providerPublicKey *babyjub.PublicKey,
	signature *babyjub.Signature,
	commitmentHash Hash,
	salt int64,
	expirationDate time.Time,
) (Hash, error) {
	hash, err := poseidon.Hash([]*big.Int{
		// fields ordered alphabetically regarding their JSON key in the circuit
		contentHash.BigInt(),
		big.NewInt(expirationDate.Unix()),
		commitmentHash.BigInt(),
		providerPublicKey.X,
		providerPublicKey.Y,
		signature.R8.X,
		signature.R8.Y,
		signature.S,
		big.NewInt(salt),
	})
	if err != nil {
		return Hash{}, fmt.Errorf("compute leaf hash: %w", err)
	}

	return HashFromBigInt(hash), nil
}

// DID is a method to generate a Decentralized Identifier (DID) by combining a given standard and leaf hash.
func DID(standard Standard, leafHash Hash) string {
	return fmt.Sprintf("did:%s:%s", standard, leafHash)
}

// EncryptedCertificate is a Certificate that has been encrypted.
type EncryptedCertificate struct {
	encryption.EncryptedMessage `json:",inline"`
	HolderCommitment            Hash `json:"holderCommitment"`
}

// DeserializeCertificateJSON deserializes Certificate from the given [io.Reader].
// As it is not possible to determine type of content in compile time, the returned content type is always Content.
func DeserializeCertificateJSON(r io.Reader) (Certificate[Content], error) {
	var alias struct {
		HolderCommitment Hash            `json:"holderCommitment"`
		LeafHash         Hash            `json:"leafHash"`
		DID              string          `json:"did"`
		Standard         Standard        `json:"zkCertStandard"`
		Content          json.RawMessage `json:"content"`
		ContentHash      Hash            `json:"contentHash"`
		ExpirationDate   Timestamp       `json:"expirationDate"`
		Provider         ProviderData    `json:"providerData"`
		RandomSalt       string          `json:"randomSalt"`
	}
	if err := json.NewDecoder(r).Decode(&alias); err != nil {
		return Certificate[Content]{}, err
	}

	content, err := decodeCertificateContent(alias.Standard, alias.Content)
	if err != nil {
		return Certificate[Content]{}, fmt.Errorf("decode alias content: %w", err)
	}

	return Certificate[Content]{
		HolderCommitment: alias.HolderCommitment,
		LeafHash:         alias.LeafHash,
		DID:              alias.DID,
		Standard:         alias.Standard,
		Content:          content,
		ContentHash:      alias.ContentHash,
		ExpirationDate:   alias.ExpirationDate,
		Provider:         alias.Provider,
		RandomSalt:       alias.RandomSalt,
	}, nil
}

// DeserializeIssuedCertificateJSON deserializes IssuedCertificate from the given [io.Reader].
// As it is not possible to determine type of content in compile time, the returned content type is always Content.
func DeserializeIssuedCertificateJSON(r io.Reader) (IssuedCertificate[Content], error) {
	var alias struct {
		HolderCommitment Hash                `json:"holderCommitment"`
		LeafHash         Hash                `json:"leafHash"`
		DID              string              `json:"did"`
		Standard         Standard            `json:"zkCertStandard"`
		Content          json.RawMessage     `json:"content"`
		ContentHash      Hash                `json:"contentHash"`
		ExpirationDate   Timestamp           `json:"expirationDate"`
		Provider         ProviderData        `json:"providerData"`
		RandomSalt       string              `json:"randomSalt"`
		Registration     RegistrationDetails `json:"registration"`
		MerkleProof      merkle.Proof        `json:"merkleProof"`
	}
	if err := json.NewDecoder(r).Decode(&alias); err != nil {
		return IssuedCertificate[Content]{}, err
	}

	content, err := decodeCertificateContent(alias.Standard, alias.Content)
	if err != nil {
		return IssuedCertificate[Content]{}, fmt.Errorf("decode alias content: %w", err)
	}

	return IssuedCertificate[Content]{
		Certificate: Certificate[Content]{
			HolderCommitment: alias.HolderCommitment,
			LeafHash:         alias.LeafHash,
			DID:              alias.DID,
			Standard:         alias.Standard,
			Content:          content,
			ContentHash:      alias.ContentHash,
			ExpirationDate:   alias.ExpirationDate,
			Provider:         alias.Provider,
			RandomSalt:       alias.RandomSalt,
		},
		Registration: alias.Registration,
		MerkleProof:  alias.MerkleProof,
	}, nil
}

func decodeCertificateContent(
	certificateStandard Standard,
	certificateContent json.RawMessage,
) (Content, error) {
	switch certificateStandard {
	case StandardKYC:
		var content KYCContent
		if err := json.Unmarshal(certificateContent, &content); err != nil {
			return nil, fmt.Errorf("decode kyc certificate content json: %w", err)
		}

		return content, nil
	case StandardSimpleJSON:
		var content SimpleJSON
		if err := json.Unmarshal(certificateContent, &content); err != nil {
			return nil, fmt.Errorf("decode kyc certificate content json: %w", err)
		}

		return content, nil
	case StandardTwitter:
		var content TwitterContent
		if err := json.Unmarshal(certificateContent, &content); err != nil {
			return nil, fmt.Errorf("decode twitter certificate content json: %w", err)
		}

		return content, nil
	case StandardREY:
		var content REYContent
		if err := json.Unmarshal(certificateContent, &content); err != nil {
			return nil, fmt.Errorf("decode rey certificate content json: %w", err)
		}

		return content, nil
	case StandardDEX:
		var content DEXContent
		if err := json.Unmarshal(certificateContent, &content); err != nil {
			return nil, fmt.Errorf("decode decentralized exchange certificate content json: %w", err)
		}

		return content, nil
	case StandardCEX:
		var content CEXContent
		if err := json.Unmarshal(certificateContent, &content); err != nil {
			return nil, fmt.Errorf("decode centrilized exchange certificate content json: %w", err)
		}

		return content, nil
	case StandardTelegram:
		var content TelegramContent
		if err := json.Unmarshal(certificateContent, &content); err != nil {
			return nil, fmt.Errorf("decode telegram certificate content json: %w", err)
		}

		return content, nil
	default:
		return nil, fmt.Errorf("unsupported certificate standard %s", certificateStandard)
	}
}
