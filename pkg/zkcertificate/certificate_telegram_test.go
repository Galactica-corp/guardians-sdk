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
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestTelegramContent_UnmarshalJSON tests various scenarios for unmarshalling JSON into TelegramContent
func TestTelegramContent_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name        string
		jsonData    string
		wantErr     bool
		errContains string
	}{
		{
			name: "Invalid Time Format",
			jsonData: `{
				"activeDaysCount": 120,
				"createdAt": "invalid-time-format"
			}`,
			wantErr:     true,
			errContains: "parsing time",
		},
		{
			name: "Missing Required Field - ActiveDaysCount",
			jsonData: `{
				"createdAt": "2023-01-01T00:00:00Z"
			}`,
			wantErr:     true,
			errContains: "required",
		},
		{
			name: "ActiveDaysCount is 0",
			jsonData: `{
				"activeDaysCount": 0,
				"createdAt": "2023-01-01T00:00:00Z"
			}`,
			wantErr:     true,
			errContains: "required",
		},
		{
			name: "Missing Required Field - CreatedAt",
			jsonData: `{
				"activeDaysCount": 120
			}`,
			wantErr:     true,
			errContains: "required",
		},
		{
			name: "Valid JSON - All Fields",
			jsonData: `{
				"activeDaysCount": 120,
				"contactWithAtLeast10MessagesCount": 25,
				"createdAt": "2023-01-01T00:00:00Z",
				"meanMonthlyMessageCount": 50
			}`,
			wantErr: false,
		},
		{
			name: "Valid JSON - Only Required Fields",
			jsonData: `{
				"activeDaysCount": 120,
				"createdAt": "2023-01-01T00:00:00Z"
			}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var content TelegramContent
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

// TestTelegramContent_Standard tests the Standard method.
func TestTelegramContent_Standard(t *testing.T) {
	content := TelegramContent{}
	standard := content.Standard()
	require.Equal(t, StandardTelegram, standard)
}

// TestTelegramContent_Hash tests the Hash method with different input combinations.
func TestTelegramContent_Hash(t *testing.T) {
	baseTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name                              string
		activeDaysCount                   uint
		contactWithAtLeast10MessagesCount uint
		createdAt                         time.Time
		meanMonthlyMessageCount           uint
	}{
		{
			name:                              "All Fields Set",
			activeDaysCount:                   120,
			contactWithAtLeast10MessagesCount: 25,
			createdAt:                         baseTime,
			meanMonthlyMessageCount:           50,
		},
		{
			name:                              "Minimum Required Fields",
			activeDaysCount:                   1,
			contactWithAtLeast10MessagesCount: 0,
			createdAt:                         baseTime,
			meanMonthlyMessageCount:           0,
		},
		{
			name:                              "High Activity Values",
			activeDaysCount:                   365,
			contactWithAtLeast10MessagesCount: 100,
			createdAt:                         baseTime,
			meanMonthlyMessageCount:           500,
		},
		{
			name:                              "Different Creation Date",
			activeDaysCount:                   120,
			contactWithAtLeast10MessagesCount: 25,
			createdAt:                         time.Date(2020, 6, 15, 12, 30, 45, 0, time.UTC),
			meanMonthlyMessageCount:           50,
		},
		{
			name:                              "Zero Optional Fields",
			activeDaysCount:                   120,
			contactWithAtLeast10MessagesCount: 0,
			createdAt:                         baseTime,
			meanMonthlyMessageCount:           0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := TelegramContent{
				ActiveDaysCount:                   tt.activeDaysCount,
				ContactWithAtLeast10MessagesCount: tt.contactWithAtLeast10MessagesCount,
				CreatedAt:                         tt.createdAt,
				MeanMonthlyMessageCount:           tt.meanMonthlyMessageCount,
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
					otherContent := TelegramContent{
						ActiveDaysCount:                   otherTT.activeDaysCount,
						ContactWithAtLeast10MessagesCount: otherTT.contactWithAtLeast10MessagesCount,
						CreatedAt:                         otherTT.createdAt,
						MeanMonthlyMessageCount:           otherTT.meanMonthlyMessageCount,
					}

					otherHash, err := otherContent.Hash()
					require.NoError(t, err)

					// If any field is different, the hash should be different
					if tt.activeDaysCount != otherTT.activeDaysCount ||
						tt.contactWithAtLeast10MessagesCount != otherTT.contactWithAtLeast10MessagesCount ||
						!tt.createdAt.Equal(otherTT.createdAt) ||
						tt.meanMonthlyMessageCount != otherTT.meanMonthlyMessageCount {
						require.NotEqual(t, hash, otherHash,
							"Hash for %s should be different from hash for %s", tt.name, otherTT.name)
					}
				}
			}
		})
	}
}

// TestTelegramContent_JSON_RoundTrip tests that marshaling and then unmarshalling
// a TelegramContent struct results in the same struct for different combinations of field values.
func TestTelegramContent_JSON_RoundTrip(t *testing.T) {
	baseTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name                              string
		activeDaysCount                   uint
		contactWithAtLeast10MessagesCount uint
		createdAt                         time.Time
		meanMonthlyMessageCount           uint
	}{
		{
			name:                              "All Fields Set",
			activeDaysCount:                   120,
			contactWithAtLeast10MessagesCount: 25,
			createdAt:                         baseTime,
			meanMonthlyMessageCount:           50,
		},
		{
			name:                              "Only Required Fields",
			activeDaysCount:                   120,
			contactWithAtLeast10MessagesCount: 0,
			createdAt:                         baseTime,
			meanMonthlyMessageCount:           0,
		},
		{
			name:                              "High Activity Values",
			activeDaysCount:                   365,
			contactWithAtLeast10MessagesCount: 100,
			createdAt:                         baseTime,
			meanMonthlyMessageCount:           500,
		},
		{
			name:                              "Different Creation Date",
			activeDaysCount:                   120,
			contactWithAtLeast10MessagesCount: 25,
			createdAt:                         time.Date(2020, 6, 15, 12, 30, 45, 0, time.UTC),
			meanMonthlyMessageCount:           50,
		},
		{
			name:                              "Maximum Values",
			activeDaysCount:                   math.MaxUint64,
			contactWithAtLeast10MessagesCount: math.MaxUint64,
			createdAt:                         baseTime,
			meanMonthlyMessageCount:           math.MaxUint64,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := TelegramContent{
				ActiveDaysCount:                   tt.activeDaysCount,
				ContactWithAtLeast10MessagesCount: tt.contactWithAtLeast10MessagesCount,
				CreatedAt:                         tt.createdAt,
				MeanMonthlyMessageCount:           tt.meanMonthlyMessageCount,
			}

			jsonData, err := json.Marshal(original)
			require.NoError(t, err)

			var unmarshalled TelegramContent
			err = json.Unmarshal(jsonData, &unmarshalled)
			require.NoError(t, err)

			require.Equal(t, original, unmarshalled)
		})
	}
}
