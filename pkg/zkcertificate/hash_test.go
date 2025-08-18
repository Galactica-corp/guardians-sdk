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

package zkcertificate_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/v2/pkg/zkcertificate"
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
