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

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

// TestBlumContent_UnmarshalJSON tests various scenarios for unmarshalling JSON into BlumContent
func TestBlumContent_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name        string
		jsonData    string
		wantErr     bool
		errContains string
	}{
		{
			name: "Missing Required Field - TelegramID",
			jsonData: `{
				"sybilScore": 85,
				"activityScore": 90
			}`,
			wantErr:     true,
			errContains: "required",
		},
		{
			name: "Missing Required Field - SybilScore",
			jsonData: `{
				"telegramId": 123456789,
				"activityScore": "90.5"
			}`,
			wantErr:     true,
			errContains: "required",
		},
		{
			name: "Missing Required Field - ActivityScore",
			jsonData: `{
				"telegramId": 123456789,
				"sybilScore": "85.3"
			}`,
			wantErr:     true,
			errContains: "required",
		},
		{
			name: "Valid JSON - All Fields",
			jsonData: `{
				"telegramId": 123456789,
				"sybilScore": "85.5",
				"activityScore": "90.75"
			}`,
			wantErr: false,
		},
		{
			name: "Valid JSON - Zero Scores",
			jsonData: `{
				"telegramId": 123456789,
				"sybilScore": "0.0",
				"activityScore": "0.0"
			}`,
			wantErr: false,
		},
		{
			name: "Valid JSON - Large TelegramID",
			jsonData: `{
				"telegramId": 9999999999,
				"sybilScore": "100.0",
				"activityScore": "100.0"
			}`,
			wantErr: false,
		},
		{
			name: "Valid JSON - Decimal Precision",
			jsonData: `{
				"telegramId": 1305035200,
				"activityScore": "40.06289271058492",
				"sybilScore": "0.20000000298023224"
			}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var content BlumContent
			err := json.Unmarshal([]byte(tt.jsonData), &content)

			if tt.wantErr {
				require.Error(t, err)
				require.ErrorContains(t, err, tt.errContains)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// TestBlumContent_Standard tests the Standard method.
func TestBlumContent_Standard(t *testing.T) {
	content := BlumContent{}
	standard := content.Standard()
	require.Equal(t, StandardBlum, standard)
}

// TestBlumContent_Hash tests the Hash method with different input combinations.
func TestBlumContent_Hash(t *testing.T) {
	tests := []struct {
		name          string
		telegramID    int64
		sybilScore    string
		activityScore string
	}{
		{
			name:          "Basic Values",
			telegramID:    123456789,
			sybilScore:    "85.5",
			activityScore: "90.75",
		},
		{
			name:          "Different TelegramID",
			telegramID:    987654321,
			sybilScore:    "85.5",
			activityScore: "90.75",
		},
		{
			name:          "Different SybilScore",
			telegramID:    123456789,
			sybilScore:    "95.3",
			activityScore: "90.75",
		},
		{
			name:          "Different ActivityScore",
			telegramID:    123456789,
			sybilScore:    "85.5",
			activityScore: "95.2",
		},
		{
			name:          "Low Scores",
			telegramID:    123456789,
			sybilScore:    "1.1",
			activityScore: "1.5",
		},
		{
			name:          "High Scores",
			telegramID:    123456789,
			sybilScore:    "100.0",
			activityScore: "100.0",
		},
		{
			name:          "Large TelegramID",
			telegramID:    9999999999,
			sybilScore:    "50.5",
			activityScore: "75.25",
		},
		{
			name:          "Decimal Precision",
			telegramID:    1305035200,
			sybilScore:    "0.20000000298023224",
			activityScore: "40.06289271058492",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := BlumContent{
				TelegramID:    tt.telegramID,
				SybilScore:    decimal.RequireFromString(tt.sybilScore),
				ActivityScore: decimal.RequireFromString(tt.activityScore),
			}

			hash, err := content.Hash()
			require.NoError(t, err)

			// The hash value will depend on the implementation of the Hash method
			// We're just checking that it doesn't return an error and produces a non-zero hash
			require.NotEqual(t, Hash{}, hash)

			// Verify that different inputs produce different hashes
			// This is a basic check to ensure the hash function is working as expected
			for _, otherTT := range tests {
				if otherTT.name != tt.name {
					otherContent := BlumContent{
						TelegramID:    otherTT.telegramID,
						SybilScore:    decimal.RequireFromString(otherTT.sybilScore),
						ActivityScore: decimal.RequireFromString(otherTT.activityScore),
					}

					otherHash, err := otherContent.Hash()
					require.NoError(t, err)

					// If any field is different, the hash should be different
					if tt.telegramID != otherTT.telegramID ||
						tt.sybilScore != otherTT.sybilScore ||
						tt.activityScore != otherTT.activityScore {
						require.NotEqual(t, hash, otherHash,
							"Hash for %s should be different from hash for %s", tt.name, otherTT.name)
					}
				}
			}
		})
	}
}

// TestBlumContent_JSON_RoundTrip tests that marshaling and then unmarshalling
// a BlumContent struct results in the same struct for different combinations of field values.
func TestBlumContent_JSON_RoundTrip(t *testing.T) {
	tests := []struct {
		name          string
		telegramID    int64
		sybilScore    string
		activityScore string
	}{
		{
			name:          "Basic Values",
			telegramID:    123456789,
			sybilScore:    "85.5",
			activityScore: "90.75",
		},
		{
			name:          "Low Scores",
			telegramID:    123456789,
			sybilScore:    "1.1",
			activityScore: "1.5",
		},
		{
			name:          "High Scores",
			telegramID:    123456789,
			sybilScore:    "100.0",
			activityScore: "100.0",
		},
		{
			name:          "Large TelegramID",
			telegramID:    9999999999,
			sybilScore:    "50.5",
			activityScore: "75.25",
		},
		{
			name:          "Decimal Precision",
			telegramID:    1305035200,
			sybilScore:    "0.20000000298023224",
			activityScore: "40.06289271058492",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := BlumContent{
				TelegramID:    tt.telegramID,
				SybilScore:    decimal.RequireFromString(tt.sybilScore),
				ActivityScore: decimal.RequireFromString(tt.activityScore),
			}

			jsonData, err := json.Marshal(original)
			require.NoError(t, err)

			var unmarshalled BlumContent
			err = json.Unmarshal(jsonData, &unmarshalled)
			require.NoError(t, err)

			require.Equal(t, original.TelegramID, unmarshalled.TelegramID)
			require.True(t, original.SybilScore.Equal(unmarshalled.SybilScore))
			require.True(t, original.ActivityScore.Equal(unmarshalled.ActivityScore))
		})
	}
}
