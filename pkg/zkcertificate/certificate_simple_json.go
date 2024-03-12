package zkcertificate

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

// SimpleJSON represents the input data for data that consists of
// simple JSON fields: strings, numbers, and booleans.
type SimpleJSON map[string]interface{}

// FFEncode implements FFEncoder.
func (k SimpleJSON) FFEncode() (SimpleJSONContent, error) {
	hashedContent := make(SimpleJSONContent)

	for key, value := range k {
		var valueStr string

		switch v := value.(type) {
		case string:
			valueStr = v
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			valueStr = fmt.Sprintf("%v", v)
		case bool:
			if v {
				valueStr = "true"
			} else {
				valueStr = "false"
			}
		case time.Time:
			valueStr = v.Format(time.RFC3339)
		default:
			return nil, fmt.Errorf("unsupported type for field %s: %T", key, v)
		}

		hash, err := poseidon.HashBytes([]byte(valueStr))
		if err != nil {
			return nil, fmt.Errorf("hash field %s: %w", key, err)
		}
		hashedContent[key] = HashFromBigInt(hash)
	}

	return hashedContent, nil
}

// Validate performs validation on the SimpleJSON instance.
// Currently, it does not perform any validation as SimpleJSON can contain any type of data.
func (k *SimpleJSON) Validate() error {
	// Currently, SimpleJSON does not have specific validation requirements.
	// We may add validation rules here if needed in the future.
	return nil
}

// UnmarshalJSON implements [json.Unmarshaler].
func (k *SimpleJSON) UnmarshalJSON(data []byte) error {
	var tempMap map[string]interface{}
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	// Create a new SimpleJSON map to store the decoded values
	decoded := make(SimpleJSON)

	for key, value := range tempMap {
		switch v := value.(type) {
		case string:
			// Try to parse the string as a time.Time value
			if t, err := time.Parse(time.RFC3339, v); err == nil {
				decoded[key] = t
			} else {
				decoded[key] = v
			}
		case float64:
			// Check if the float64 value has a decimal part
			if v == float64(int64(v)) {
				// If the value doesn't have a decimal part, convert it to int64
				decoded[key] = int64(v)
			} else {
				decoded[key] = v
			}
		case bool:
			decoded[key] = v
		default:
			return fmt.Errorf("unsupported type for field %s: %T", key, v)
		}
	}

	*k = decoded
	return nil
}

// SimpleJSONContent represents the hashed content of SimpleJSON data.
// It contains hashed values for various fields related to identity and verification.
type SimpleJSONContent map[string]Hash

// Standard returns the standard associated with the SimpleJSONContent, which is StandardSimpleJSON.
func (c SimpleJSONContent) Standard() Standard {
	return StandardSimpleJSON
}

// Hash computes and returns the hash of the SimpleJSONContent instance.
func (c SimpleJSONContent) Hash() (Hash, error) {
	var hashInputs []*big.Int

	for _, value := range c {
		hashInputs = append(hashInputs, value.BigInt())
	}

	hash, err := poseidon.Hash(hashInputs)
	if err != nil {
		return Hash{}, err
	}

	return HashFromBigInt(hash), nil
}
