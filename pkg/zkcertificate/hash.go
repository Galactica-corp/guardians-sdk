package zkcertificate

import (
	"math/big"
)

// Hash represents a cryptographic hash value obtained by Poseidon algorithm.
type Hash big.Int

// HashFromBigInt creates a Hash from a given big.Int value.
func HashFromBigInt(n *big.Int) Hash {
	if n == nil {
		n = big.NewInt(1)
	}

	return Hash(*n)
}

// BigInt converts a Hash value to a big.Int.
func (h Hash) BigInt() *big.Int {
	n := big.Int(h)
	return &n
}

// Bytes32 converts a Hash value to a 32-byte array.
func (h Hash) Bytes32() [32]byte {
	var res [32]byte
	h.BigInt().FillBytes(res[:])

	return res
}

// String returns the string representation of the Hash value.
func (h Hash) String() string {
	return h.BigInt().String()
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (h *Hash) UnmarshalText(text []byte) error {
	var res big.Int
	if err := res.UnmarshalText(text); err != nil {
		return err
	}

	*h = Hash(res)
	return nil
}

// MarshalText implements encoding.TextMarshaler interface.
func (h Hash) MarshalText() (text []byte, err error) {
	return h.BigInt().MarshalText()
}
