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

// UnmarshalText implements [encoding.TextUnmarshaler].
func (h *Hash) UnmarshalText(text []byte) error {
	var res big.Int
	if err := res.UnmarshalText(text); err != nil {
		return err
	}

	*h = Hash(res)
	return nil
}

// MarshalText implements [encoding.TextMarshaler].
func (h Hash) MarshalText() (text []byte, err error) {
	return h.BigInt().MarshalText()
}
