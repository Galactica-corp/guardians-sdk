package encryption_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/pkg/encryption"
)

func TestEncryptWithPadding(t *testing.T) {
	tests := []struct {
		name    string
		message any
	}{
		{
			name:    "Non-empty message",
			message: "Hello, world!",
		},
		{
			name:    "Empty message",
			message: "",
		},
		{
			name:    "Exactly 2048 bytes",
			message: strings.Repeat("A", 2048-len(`{"data":"","padding":""}`)),
		},
		{
			name:    "More than 2048 bytes",
			message: strings.Repeat("A", 3000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := encryption.EncryptWithPadding([32]byte{}, tt.message)
			require.NoError(t, err)
			require.NotEmpty(t, res.Ciphertext)
		})
	}
}

func TestEncrypt(t *testing.T) {
	res, err := encryption.Encrypt([32]byte{}, []byte("Hello, world!"))
	require.NoError(t, err)
	require.Equal(t, encryption.VersionNaClAuthenticated, res.Version)
	require.NotEmpty(t, res.Nonce)
	require.NotEmpty(t, res.EphemeralPublicKey)
	require.NotEmpty(t, res.Ciphertext)
}
