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

package keymanagement

import (
	"bufio"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/iden3/go-iden3-crypto/v2/babyjub"
)

const edDSAGenerationMessage = "Signing this message generates your EdDSA private key. Only do this on pages you trust to manage your zkCertificates."

var edDSAGenerationMessageHash = crypto.Keccak256Hash([]byte(edDSAGenerationMessage)).Bytes()

// DeriveEdDSAKeyFromEthereumPrivateKey derives an EdDSA (Poseidon hash over BN128 elliptic curve signature scheme)
// private key from an Ethereum private key. It signs a predefined message hash using the Ethereum private key
// and converts the resulting signature into an EdDSA private key for further use.
func DeriveEdDSAKeyFromEthereumPrivateKey(privateKey *ecdsa.PrivateKey) (babyjub.PrivateKey, error) {
	signature, err := crypto.Sign(edDSAGenerationMessageHash, privateKey)
	if err != nil {
		return babyjub.PrivateKey{}, fmt.Errorf("sign message: %w", err)
	}

	return babyjub.PrivateKey(signature), err
}

// LoadEdDSA loads a Baby Jubjub private key from the given file.
//
// Source: https://github.com/ethereum/go-ethereum/blob/285202aae2a580b82f30ebd909c1819b22d90066/crypto/crypto.go#L194
func LoadEdDSA(file string) (babyjub.PrivateKey, error) {
	fd, err := os.Open(file)
	if err != nil {
		return babyjub.PrivateKey{}, err
	}
	defer fd.Close()

	r := bufio.NewReader(fd)
	buf := make([]byte, 64)
	n, err := readASCII(buf, r)
	if err != nil {
		return babyjub.PrivateKey{}, err
	} else if n != len(buf) {
		return babyjub.PrivateKey{}, errors.New("key file too short, want 64 hex characters")
	}
	if err := checkKeyFileEnd(r); err != nil {
		return babyjub.PrivateKey{}, err
	}

	res := make([]byte, hex.DecodedLen(len(buf)))

	var byteErr hex.InvalidByteError
	if _, err = hex.Decode(res, buf); errors.As(err, &byteErr) {
		return babyjub.PrivateKey{}, fmt.Errorf("invalid hex character %q in private key", byte(byteErr))
	} else if err != nil {
		return babyjub.PrivateKey{}, errors.New("invalid hex data for private key")
	}

	return babyjub.PrivateKey(res), nil
}

// readASCII reads into 'buf', stopping when the buffer is full or
// when a non-printable control character is encountered.
//
// Source: https://github.com/ethereum/go-ethereum/blob/285202aae2a580b82f30ebd909c1819b22d90066/crypto/crypto.go#L218
func readASCII(buf []byte, r io.ByteReader) (n int, err error) {
	for ; n < len(buf); n++ {
		buf[n], err = r.ReadByte()
		switch {
		case errors.Is(err, io.EOF) || buf[n] < '!':
			return n, nil
		case err != nil:
			return n, err
		}
	}
	return n, nil
}

// checkKeyFileEnd skips over additional newlines at the end of a key file.
//
// Source: https://github.com/ethereum/go-ethereum/blob/285202aae2a580b82f30ebd909c1819b22d90066/crypto/crypto.go#L232
func checkKeyFileEnd(r io.ByteReader) error {
	for i := 0; ; i++ {
		b, err := r.ReadByte()
		switch {
		case errors.Is(err, io.EOF):
			return nil
		case err != nil:
			return err
		case b != '\n' && b != '\r':
			return fmt.Errorf("invalid character %q at end of key file", b)
		case i >= 2:
			return errors.New("key file too long, want 64 hex characters")
		}
	}
}

// SaveEdDSA saves a Baby Jubjub private key to the given file with
// restrictive permissions. The key data is saved hex-encoded.
//
// Source: https://github.com/ethereum/go-ethereum/blob/285202aae2a580b82f30ebd909c1819b22d90066/crypto/crypto.go#L250
func SaveEdDSA(file string, key babyjub.PrivateKey) error {
	fd, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()

	_, err = hex.NewEncoder(fd).Write(key[:])
	return err
}
