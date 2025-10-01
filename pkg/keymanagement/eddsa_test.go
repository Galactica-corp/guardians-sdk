// Copyright © 2025 Galactica Network
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

package keymanagement_test

import (
	"encoding/hex"
	"os"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/iden3/go-iden3-crypto/v2/babyjub"
	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/v4/pkg/keymanagement"
)

func TestDeriveEdDSAKeyFromEthereumPrivateKey(t *testing.T) {
	ethereumPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Skip(err)
	}

	edDSAPrivateKey, err := keymanagement.DeriveEdDSAKeyFromEthereumPrivateKey(ethereumPrivateKey)
	require.NoError(t, err)
	require.NotEmpty(t, edDSAPrivateKey)
}

func TestLoadEdDSA(t *testing.T) {
	tests := []struct {
		name  string
		input string
		err   string
	}{
		// good
		{name: "good", input: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"},
		{name: "good", input: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef\n"},
		{name: "good", input: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef\n\r"},
		{name: "good", input: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef\r\n"},
		{name: "good", input: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef\n\n"},
		{name: "good", input: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef\n\r"},
		// bad
		{
			name:  "bad",
			input: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcde",
			err:   "key file too short, want 64 hex characters",
		},
		{
			name:  "bad",
			input: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcde\n",
			err:   "key file too short, want 64 hex characters",
		},
		{
			name:  "bad",
			input: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdeX",
			err:   "invalid hex character 'X' in private key",
		},
		{
			name:  "bad",
			input: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdefX",
			err:   "invalid character 'X' at end of key file",
		},
		{
			name:  "bad",
			input: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef\n\n\n",
			err:   "key file too long, want 64 hex characters",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := os.CreateTemp("", "loadeddsa_test.*.txt")
			if err != nil {
				t.Fatal(err)
			}
			filename := f.Name()
			_, _ = f.WriteString(test.input)
			_ = f.Close()
			defer os.Remove(filename)

			_, err = keymanagement.LoadEdDSA(filename)
			switch {
			case err != nil && test.err == "":
				t.Fatalf("unexpected error for input %q:\n  %v", test.input, err)
			case err != nil && err.Error() != test.err:
				t.Fatalf("wrong error for input %q:\n  %v", test.input, err)
			case err == nil && test.err != "":
				t.Fatalf("LoadECDSA did not return error for input %q", test.input)
			}
		})
	}
}

func TestSaveEdDSA(t *testing.T) {
	f, err := os.CreateTemp("", "saveeddsa_test.*.txt")
	if err != nil {
		t.Fatal(err)
	}
	file := f.Name()
	_ = f.Close()
	defer os.Remove(file)

	testPrivBytes, _ := hex.DecodeString("289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032")
	key := babyjub.PrivateKey(testPrivBytes)

	if err := keymanagement.SaveEdDSA(file, key); err != nil {
		t.Fatal(err)
	}
	loaded, err := keymanagement.LoadEdDSA(file)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(key, loaded) {
		t.Fatal("loaded key not equal to saved key")
	}
}
