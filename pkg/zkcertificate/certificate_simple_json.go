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

package zkcertificate

import (
	"encoding/json"
	"fmt"
	"math/big"
	"sort"

	"github.com/iden3/go-iden3-crypto/v2/poseidon"

	"github.com/galactica-corp/guardians-sdk/pkg/hashing"
)

// SimpleJSON represents the input data for data that consists of
// simple JSON fields: strings only.
type SimpleJSON map[string]string

// Validate performs validation on the SimpleJSON instance.
func (c *SimpleJSON) Validate() error {
	return nil
}

// UnmarshalJSON implements [json.Unmarshaler].
func (c *SimpleJSON) UnmarshalJSON(data []byte) error {
	var tempMap map[string]string
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	*c = tempMap
	return nil
}

// Standard returns the standard associated with the SimpleJSONContent, which is StandardSimpleJSON.
func (c SimpleJSON) Standard() Standard {
	return StandardSimpleJSON
}

// Hash computes and returns the hash of the SimpleJSON instance.
func (c SimpleJSON) Hash() (Hash, error) {
	keys := make([]string, 0, len(c))
	for key := range c {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	hashedContent := make([]*big.Int, len(c))
	for i, key := range keys {
		value := c[key]

		hash, err := hashing.HashBytes([]byte(value))
		if err != nil {
			return Hash{}, fmt.Errorf("hash field %s: %w", key, err)
		}

		hashedContent[i] = hash
	}

	hash, err := poseidon.Hash(hashedContent)
	if err != nil {
		return Hash{}, err
	}

	return HashFromBigInt(hash), nil
}
