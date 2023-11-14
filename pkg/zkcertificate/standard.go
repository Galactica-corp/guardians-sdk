package zkcertificate

import (
	"fmt"
	"slices"
)

// Standard represents a string that indicates the standard of Zero Knowledge certificates.
type Standard string

const StandardKYC Standard = "gip69"

var allStandards = []string{StandardKYC.String()}

// IsStandard returns true if given value is a valid Standard.
func IsStandard(value string) bool {
	return slices.Contains(allStandards, value)
}

// String returns the string representation of the Standard value.
func (s Standard) String() string {
	return string(s)
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (s *Standard) UnmarshalText(value []byte) error {
	text := string(value)

	if !IsStandard(text) {
		return fmt.Errorf("invalid value")
	}

	*s = Standard(value)
	return nil
}

// MarshalText implements encoding.TextMarshaler interface.
func (s Standard) MarshalText() (text []byte, err error) {
	return []byte(s), nil
}
