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
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/v3/pkg/zkcertificate"
)

func TestIsStandard(t *testing.T) {
	tests := []struct {
		name     string
		standard string
		want     bool
	}{
		{
			name:     "Valid standard",
			standard: zkcertificate.StandardKYC.String(),
			want:     true,
		},
		{
			name:     "Invalid standard",
			standard: "invalid",
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zkcertificate.IsStandard(tt.standard); got != tt.want {
				t.Errorf("IsStandard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStandard_MarshalText(t *testing.T) {
	text, err := zkcertificate.StandardKYC.MarshalText()
	require.NoError(t, err)
	require.Equal(t, []byte(zkcertificate.StandardKYC.String()), text)
}

func TestStandard_UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		text    []byte
		wantErr bool
	}{
		{
			name:    "Valid value",
			text:    []byte(zkcertificate.StandardKYC.String()),
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
			var v zkcertificate.Standard
			if err := v.UnmarshalText(tt.text); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
