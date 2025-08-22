// Copyright © 2025 Galactica Network
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

	"github.com/galactica-corp/guardians-sdk/v3/internal/validation"
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
	return validation.Validate.Struct(c)
}

// UnmarshalJSON implements [json.Unmarshaler].
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
