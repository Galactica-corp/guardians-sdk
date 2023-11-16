/*
 Copyright © 2023 Galactica Network

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.

 You should have received a copy of the GNU General Public License
 along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

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
