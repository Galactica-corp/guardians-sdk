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

package encryption_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/v4/pkg/encryption"
)

func TestIsVersion(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    bool
	}{
		{
			name:    "Valid version",
			version: encryption.VersionNaClAuthenticated.String(),
			want:    true,
		},
		{
			name:    "Invalid version",
			version: "invalid",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encryption.IsVersion(tt.version); got != tt.want {
				t.Errorf("IsVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_MarshalText(t *testing.T) {
	text, err := encryption.VersionNaClAuthenticated.MarshalText()
	require.NoError(t, err)
	require.Equal(t, []byte(encryption.VersionNaClAuthenticated.String()), text)
}

func TestVersion_UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		text    []byte
		wantErr bool
	}{
		{
			name:    "Valid value",
			text:    []byte(encryption.VersionNaClAuthenticated.String()),
			wantErr: false,
		},
		{
			name:    "Invalid value",
			text:    []byte("invalid"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var v encryption.Version
			if err := v.UnmarshalText(tt.text); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
