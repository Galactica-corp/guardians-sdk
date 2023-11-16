package zkcertificate_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

func TestHolderCommitment_Validate(t *testing.T) {
	holderCommitment := zkcertificate.HolderCommitment{
		CommitmentHash: zkcertificate.HashFromBigInt(big.NewInt(1)),
		EncryptionKey: []byte{
			0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
			0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20,
			0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30,
			0x31, 0x32,
		},
	}

	require.NoError(t, holderCommitment.Validate())
}

func TestHolderCommitment_Validate_invalid(t *testing.T) {
	holderCommitment := zkcertificate.HolderCommitment{
		CommitmentHash: zkcertificate.HashFromBigInt(big.NewInt(1)),
		EncryptionKey: []byte{
			0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
			0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20,
			0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30,
		},
	}

	require.EqualError(
		t,
		holderCommitment.Validate(),
		"Key: 'HolderCommitment.EncryptionKey' Error:Field validation for 'EncryptionKey' failed on the 'len' tag",
	)
}

func TestHolderCommitment_JSON(t *testing.T) {
	holderCommitment := zkcertificate.HolderCommitment{
		CommitmentHash: mustHashFromString("21299951605992408668949924562963568070883824906758011123350028140304929514899"),
		EncryptionKey:  mustDecodeBase64("85AgkcUzbC3hZQyE5911ZfZMhy5OyEZPFd6l5H6uqEw="),
	}

	data, err := json.Marshal(holderCommitment)
	require.NoError(t, err)

	expected := []byte(`{"holderCommitment":"21299951605992408668949924562963568070883824906758011123350028140304929514899","encryptionPubKey":"85AgkcUzbC3hZQyE5911ZfZMhy5OyEZPFd6l5H6uqEw="}`)
	require.Equal(t, expected, data)

	var deserialized zkcertificate.HolderCommitment
	require.NoError(t, json.Unmarshal(data, &deserialized))
	require.Equal(t, holderCommitment, deserialized)
}

func mustDecodeBase64(s string) []byte {
	res, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(fmt.Sprintf("invalid base64 string %q: %s", s, err.Error()))
	}

	return res
}
