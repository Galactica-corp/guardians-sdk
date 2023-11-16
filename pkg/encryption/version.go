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

package encryption

import (
	"fmt"
	"slices"
)

// Version represents a string that indicates the set of used cryptography algorithms.
type Version string

// VersionNaClAuthenticated uses Curve25519, XSalsa20 and Poly1305 to encrypt and authenticate messages.
const VersionNaClAuthenticated Version = "x25519-xsalsa20-poly1305"

var allVersions = []string{VersionNaClAuthenticated.String()}

// IsVersion returns true if given value is a valid Version.
func IsVersion(version string) bool {
	return slices.Contains(allVersions, version)
}

// String returns the string representation of the Version value.
func (v Version) String() string {
	return string(v)
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (s *Version) UnmarshalText(value []byte) error {
	text := string(value)

	if !IsVersion(text) {
		return fmt.Errorf("invalid value")
	}

	*s = Version(value)
	return nil
}

// MarshalText implements encoding.TextMarshaler interface.
func (s Version) MarshalText() (text []byte, err error) {
	return []byte(s), nil
}
