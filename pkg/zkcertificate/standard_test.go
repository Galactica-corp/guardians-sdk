package zkcertificate_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
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
