// Copyright © 2024 Galactica Network
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

// Package keymanagement provides utilities for managing cryptographic keys,
// specifically focusing on EdDSA (Poseidon hash over BN128 elliptic curve
// signature scheme) keys derived from Ethereum private keys.
//
// It includes functions for deriving EdDSA keys from Ethereum private keys,
// loading EdDSA keys from files, and saving EdDSA keys to files securely.
package keymanagement
