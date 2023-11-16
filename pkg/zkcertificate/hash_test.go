package zkcertificate_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

func TestHash_Bytes32(t *testing.T) {
	expected := [32]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x15}
	actual := zkcertificate.HashFromBigInt(big.NewInt(789)).Bytes32()
	require.Equal(t, expected, actual)
}

func TestHash_String(t *testing.T) {
	actual := zkcertificate.HashFromBigInt(big.NewInt(101112)).String()
	require.Equal(t, "101112", actual)
}

func TestHash_JSON(t *testing.T) {
	hash := zkcertificate.HashFromBigInt(big.NewInt(131415))

	data, err := json.Marshal(hash)
	require.NoError(t, err)
	require.Equal(t, []byte(`"131415"`), data)

	var deserialized zkcertificate.Hash
	require.NoError(t, json.Unmarshal(data, &deserialized))
	require.Equal(t, hash, deserialized)
}
