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
func (k SimpleJSON) FFEncode() (SimpleJSONContent, error) {
	keys := make([]string, 0, len(k))
	for key := range k {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	hashedContent := make(SimpleJSONContent, len(k))

	for i, key := range keys {
		value := k[key]
		var valueStr string

		switch v := value.(type) {
		case string:
			valueStr = v
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
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
		return fmt.Errorf("cannot validate nil SimpleJSON")
	}

	keys := make([]string, 0, len(*c))
	for key := range *c {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := (*c)[key]
		var valueStr string

		switch v := value.(type) {
		case string:
			valueStr = v
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			valueStr = fmt.Sprintf("%v", v)
		case bool:
			valueStr = fmt.Sprintf("%t", v)
		default:
			return fmt.Errorf("unsupported type for field %s: %T", key, v)
		}

		_, err := poseidon.HashBytes([]byte(valueStr))
		if err != nil {
			return fmt.Errorf("hash field %s: %w", key, err)
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

	// Create a new SimpleJSON map to store the decoded values
	decoded := make(SimpleJSON)

	for key, value := range tempMap {
		switch v := value.(type) {
		case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
			decoded[key] = v
		default:
			return fmt.Errorf("unsupported type for field %s: %T", key, v)
		}
	}

	*c = decoded
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
