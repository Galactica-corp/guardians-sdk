package zkcertificate

import (
	"encoding/json"
	"fmt"
)

// HolderCommitment represents a structure containing a commitment hash and an encryption key.
type HolderCommitment struct {
	CommitmentHash Hash   `json:"holderCommitment" validate:"required"`
	EncryptionKey  []byte `json:"encryptionPubKey" validate:"required,len=32"`
}

// Validate performs validation on the HolderCommitment instance using the struct tags specified
// for field validation. It checks that the fields of the HolderCommitment struct adhere to
// the defined validation rules.
func (c HolderCommitment) Validate() error {
	return validate.Struct(c)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (c *HolderCommitment) UnmarshalJSON(data []byte) error {
	type Alias HolderCommitment

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	if err := (*HolderCommitment)(&alias).Validate(); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	*c = HolderCommitment(alias)
	return nil
}
