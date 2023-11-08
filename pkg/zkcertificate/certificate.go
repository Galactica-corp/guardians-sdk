package zkcertificate

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

// Certificate represents a zero knowledge certificate structure that can hold different types of content.
// It is parameterized by the type T for the content field.
// Certificate content must be directly determined by the certificate Standard.
type Certificate[T any] struct {
	HolderCommitment Hash         `json:"holderCommitment"`
	LeafHash         Hash         `json:"leafHash"`
	DID              string       `json:"did"`
	Standard         Standard     `json:"zkCertStandard"`
	Content          T            `json:"content"`
	ContentHash      Hash         `json:"contentHash"`
	Provider         ProviderData `json:"providerData"`
	RandomSalt       int64        `json:"randomSalt"`
}

// ProviderData represents the public key and signature data of a certificate provider.
type ProviderData struct {
	PublicKey *babyjub.PublicKey
	Signature *babyjub.Signature
}

// CertificateContent is an interface that represents the content of a certificate.
// It defines methods for calculating the content's hash and obtaining the standard it adheres to.
type CertificateContent interface {
	// Hash computes and returns the Poseidon hash of the certificate content.
	Hash() (Hash, error)
	// Standard returns the standard to which the certificate content adheres.
	Standard() Standard
}

// NewCertificate creates a new certificate instance with the provided parameters and content.
// It computes the content hash, verifies if the content was actually signed with providers public key,
// and generates a leaf hash.
func NewCertificate[T CertificateContent](
	holderCommitment Hash,
	content T,
	providerPublicKey *babyjub.PublicKey,
	providerSignature *babyjub.Signature,
	salt int64,
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

	leafHash, err := LeafHash(contentHash, providerPublicKey, providerSignature, holderCommitment, salt)
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
		Provider: ProviderData{
			PublicKey: providerPublicKey,
			Signature: providerSignature,
		},
		RandomSalt: salt,
	}, nil
}

type providerDataDTO struct {
	Ax  string `json:"ax"`
	Bx  string `json:"bx"`
	S   string `json:"s"`
	R8x string `json:"r8x"`
	R8y string `json:"r8y"`
}

// MarshalJSON implements json.Marshaler interface.
func (p ProviderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(providerDataDTO{
		Ax:  p.PublicKey.X.String(),
		Bx:  p.PublicKey.Y.String(),
		S:   p.Signature.S.String(),
		R8x: p.Signature.R8.X.String(),
		R8y: p.Signature.R8.Y.String(),
	})
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (p *ProviderData) UnmarshalJSON(data []byte) error {
	var dto providerDataDTO
	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	publicKeyPoint := &babyjub.Point{}

	var ok bool

	publicKeyPoint.X, ok = new(big.Int).SetString(dto.Ax, 10)
	if !ok {
		return fmt.Errorf("invalid x coordinate of public key point")
	}

	publicKeyPoint.Y, ok = new(big.Int).SetString(dto.Bx, 10)
	if !ok {
		return fmt.Errorf("invalid y coordinate of public key point")
	}

	p.PublicKey = (*babyjub.PublicKey)(publicKeyPoint)

	signature := &babyjub.Signature{}

	signatureR8Point := &babyjub.Point{}

	signatureR8Point.X, ok = new(big.Int).SetString(dto.R8x, 10)
	if !ok {
		return fmt.Errorf("invalid x coordinate of signature r8 point")
	}

	signatureR8Point.Y, ok = new(big.Int).SetString(dto.R8y, 10)
	if !ok {
		return fmt.Errorf("invalid y coordinate of signature r8 point")
	}

	signature.R8 = signatureR8Point

	signature.S, ok = new(big.Int).SetString(dto.S, 10)
	if !ok {
		return fmt.Errorf("invalid s component of signature")
	}

	p.Signature = signature

	return nil
}

// SignCertificate generates a digital signature for a certificate using the provider's private key.
func SignCertificate(
	providerKey babyjub.PrivateKey,
	contentHash Hash,
	commitmentHash Hash,
) (*babyjub.Signature, error) {
	message, err := poseidon.Hash([]*big.Int{contentHash.BigInt(), commitmentHash.BigInt()})
	if err != nil {
		return nil, fmt.Errorf("hash message: %w", err)
	}

	// TODO: Why mod here? It doesn't crash without mod
	// message = message.Mod(message, utils.NewIntFromString("2736030358979909402780800718157159386076813972158567259200215660948447373040"))

	return providerKey.SignPoseidon(message), nil
}

// VerifySignature verifies the digital signature of a certificate using the provider's public key.
func VerifySignature(
	providerKey *babyjub.PublicKey,
	contentHash Hash,
	commitmentHash Hash,
	signature *babyjub.Signature,
) (bool, error) {
	message, err := poseidon.Hash([]*big.Int{contentHash.BigInt(), commitmentHash.BigInt()})
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
) (Hash, error) {
	hash, err := poseidon.Hash([]*big.Int{
		contentHash.BigInt(),
		providerPublicKey.X,
		providerPublicKey.Y,
		signature.S,
		signature.R8.X,
		signature.R8.Y,
		commitmentHash.BigInt(),
		big.NewInt(salt),
	})
	if err != nil {
		return Hash{}, fmt.Errorf("compute hash: %w", err)
	}

	return HashFromBigInt(hash), nil
}

// DID is a method to generate a Decentralized Identifier (DID) by combining a given standard and leaf hash.
func DID(standard Standard, leafHash Hash) string {
	return fmt.Sprintf("did:%s:%s", standard, leafHash)
}

// FFEncoder is an interface for objects that can perform encoding to Finite Field (FF).
type FFEncoder[T CertificateContent] interface {
	// FFEncode performs Finite Field (FF) encoding and returns the result that can be used as certificate content.
	FFEncode() (T, error)
}
