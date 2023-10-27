package keymanagement

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

const edDSAGenerationMessage = "Signing this message generates your EdDSA private key. Only do this on pages you trust to manage your zkCertificates."

func GetEdDSAKeyFromEthereumSigner(privateKey *ecdsa.PrivateKey) ([]byte, error) {
	hash := crypto.Keccak256Hash([]byte(edDSAGenerationMessage))
	return crypto.Sign(hash.Bytes(), privateKey)
}
