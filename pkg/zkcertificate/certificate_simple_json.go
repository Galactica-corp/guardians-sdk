package zkcertificate

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

// SimpleJSON represents the input data for data that consists of
// simple JSON fields: strings, numbers, and booleans.
type SimpleJSON map[string]interface{}

// FFEncode implements FFEncoder.
func (k SimpleJSON) FFEncode() (SimpleJSONContent, error) {
	hashedContent := make(SimpleJSONContent)

	for key, value := range k {
		hash, err := poseidon.HashBytes([]byte(fmt.Sprintf("%v", value)))
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
	*k = SimpleJSON(tempMap)
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
