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

package zkcertificate

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestDEXContent_UnmarshalJSON_validJSON(t *testing.T) {
	jsonData := `{
			"address": "0x1234567890123456789012345678901234567890",
			"totalSwapVolume": "100.50",
			"swapVolumeYear": "75.25",
			"swapVolumeHalfYear": "50.75"
		}`

	var content DEXContent
	err := json.Unmarshal([]byte(jsonData), &content)
	require.NoError(t, err)

	expected := DEXContent{
		Address:            common.HexToAddress("0x1234567890123456789012345678901234567890"),
		TotalSwapVolume:    decimal.RequireFromString("100.50"),
		SwapVolumeYear:     decimal.RequireFromString("75.25"),
		SwapVolumeHalfYear: decimal.RequireFromString("50.75"),
	}
	require.Equal(t, expected.Address, content.Address)
	require.True(t, expected.TotalSwapVolume.Equal(content.TotalSwapVolume))
	require.True(t, expected.SwapVolumeYear.Equal(content.SwapVolumeYear))
	require.True(t, expected.SwapVolumeHalfYear.Equal(content.SwapVolumeHalfYear))
}

func TestDEXContent_UnmarshalJSON_missingRequiredFields(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name: "Missing Address",
			input: `{
					"totalSwapVolume": "100.50"
				}`,
			wantErr: true,
		},
		{
			name: "Missing TotalSwapVolume",
			input: `{
					"address": "0x1234567890123456789012345678901234567890"
				}`,
			wantErr: true,
		},
		{
			name: "TotalSwapVolume Zero",
			input: `{
					"address": "0x1234567890123456789012345678901234567890",
					"totalSwapVolume": "0"
				}`,
			wantErr: true,
		},
		{
			name: "TotalSwapVolume Negative",
			input: `{
					"address": "0x1234567890123456789012345678901234567890",
					"totalSwapVolume": "-10.5"
				}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var content DEXContent
			err := json.Unmarshal([]byte(tt.input), &content)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
