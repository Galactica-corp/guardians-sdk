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

package encryption

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"golang.org/x/crypto/nacl/box"
)

const defaultPaddingLength = 2 << 10

type dataWithPadding struct {
	Data    any
	Padding string
}

// EncryptedMessage represents an encrypted message with additional meta-information.
type EncryptedMessage struct {
	Version            Version `json:"version"`
	Nonce              []byte  `json:"nonce"`
	EphemeralPublicKey []byte  `json:"ephemPublicKey"`
	Ciphertext         []byte  `json:"ciphertext"`
}

// EncryptWithPadding encrypts a message in a way that obscures the message length.
//
// The message is padded to a multiple of 2048 before being encrypted so that the length of the
// resulting encrypted message can't be used to guess the exact length of the original message.
//
// Source: https://github.com/MetaMask/eth-sig-util/blob/10206bf2f16f0b47b1f2da9a9cfbb39c6a7a7800/src/encryption.ts#L94
func EncryptWithPadding(encryptionKey [32]byte, message any) (EncryptedMessage, error) {
	data, err := json.Marshal(dataWithPadding{
		Data:    message,
		Padding: "",
	})
	if err != nil {
		return EncryptedMessage{}, fmt.Errorf("encode message to json: %w", err)
	}

	padLength := 0
	if m := len(data) % defaultPaddingLength; m > 0 {
		padLength = defaultPaddingLength - m - box.Overhead
	}

	data, err = json.Marshal(dataWithPadding{
		Data:    message,
		Padding: strings.Repeat("0", padLength),
	})
	if err != nil {
		return EncryptedMessage{}, fmt.Errorf("encode padded message to json: %w", err)
	}

	return Encrypt(encryptionKey, data)
}

// Encrypt encrypts a message using VersionNaClAuthenticated encryption version.
func Encrypt(encryptionKey [32]byte, message []byte) (EncryptedMessage, error) {
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return EncryptedMessage{}, fmt.Errorf("generate nonce: %w", err)
	}

	ephemeralPublicKey, ephemeralPrivateKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return EncryptedMessage{}, fmt.Errorf("generate ephemeral key pair: %w", err)
	}

	encryptedMessage := box.Seal(nil, message, &nonce, &encryptionKey, ephemeralPrivateKey)

	return EncryptedMessage{
		Version:            VersionNaClAuthenticated,
		Nonce:              nonce[:],
		EphemeralPublicKey: (*ephemeralPublicKey)[:],
		Ciphertext:         encryptedMessage,
	}, nil
}
