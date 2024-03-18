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

	"github.com/iden3/go-iden3-crypto/poseidon"
)

// SimpleJSON represents the input data for data that consists of
// simple JSON fields: strings only.
type SimpleJSON map[string]string

// FFEncode implements FFEncoder.
func (c SimpleJSON) FFEncode() (SimpleJSONContent, error) {
	keys := make([]string, 0, len(c))
	for key := range c {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	hashedContent := make(SimpleJSONContent, len(c))

	for i, key := range keys {
		value := c[key]

		hash, err := poseidon.HashBytes([]byte(value))
		if err != nil {
			return nil, fmt.Errorf("hash field %s: %w", key, err)
		}
		hashedContent[i] = HashFromBigInt(hash)
	}

	return hashedContent, nil
}

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

	*c = SimpleJSON(tempMap)
	return nil
}

// SimpleJSONContent represents the hashed content of SimpleJSON data.
// It ordered by the SimpleJSON data key's natural order.
type SimpleJSONContent []Hash

// Standard returns the standard associated with the SimpleJSONContent, which is StandardSimpleJSON.
func (c SimpleJSONContent) Standard() Standard {
	return StandardSimpleJSON
}

// Hash computes and returns the hash of the SimpleJSONContent instance.
func (c SimpleJSONContent) Hash() (Hash, error) {
	hashInputs := make([]*big.Int, len(c))
	for i, hash := range c {
		hashInputs[i] = hash.BigInt()
	}

	hash, err := poseidon.Hash(hashInputs)
	if err != nil {
		return Hash{}, err
	}

	return HashFromBigInt(hash), nil
}
