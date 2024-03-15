package zkcertificate

import (
	"encoding/json"
	"fmt"
	"math/big"
	"sort"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

// SimpleJSON represents the input data for data that consists of
// simple JSON fields: strings, numbers, and booleans.
type SimpleJSON map[string]interface{}

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
		var valueStr string

		switch v := value.(type) {
		case string:
			valueStr = v
		case float64:
			valueStr = fmt.Sprintf("%v", v)
		case bool:
			valueStr = fmt.Sprintf("%t", v)
		default:
			return nil, fmt.Errorf("unsupported type for field %s: %T", key, v)
		}

		hash, err := poseidon.HashBytes([]byte(valueStr))
		if err != nil {
			return nil, fmt.Errorf("hash field %s: %w", key, err)
		}
		hashedContent[i] = HashFromBigInt(hash)
	}

	return hashedContent, nil
}

// Validate performs validation on the SimpleJSON instance.
func (c *SimpleJSON) Validate() error {
	if c == nil {
		return nil
	}

	for key, value := range *c {
		switch value.(type) {
		case string, float64, bool:
			// If the value is of an expected type, no action is needed.
		default:
			return fmt.Errorf("unsupported type for field %s: %T", key, value)
		}
	}

	return nil
}

// UnmarshalJSON implements [json.Unmarshaler].
func (c *SimpleJSON) UnmarshalJSON(data []byte) error {
	var tempMap map[string]interface{}
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	tempSimpleJSON := SimpleJSON(tempMap)

	if err := tempSimpleJSON.Validate(); err != nil {
		return err
	}

	*c = tempSimpleJSON
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
