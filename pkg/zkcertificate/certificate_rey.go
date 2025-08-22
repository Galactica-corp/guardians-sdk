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
	"math/big"
	"strings"

	"github.com/iden3/go-iden3-crypto/v2/poseidon"

	"github.com/galactica-corp/guardians-sdk/v3/internal/validation"
	"github.com/galactica-corp/guardians-sdk/v3/pkg/hashing"
)

// REYContent represents the data for the REY verification.
type REYContent struct {
	XID               string `json:"xID" validate:"required,number"`
	XUsername         string `json:"xUsername" validate:"required,min=4,max=15"`
	REYScoreAll       uint   `json:"reyScoreAll"`
	REYScoreGalactica uint   `json:"reyScoreGalactica"`
	REYFaction        uint   `json:"reyFaction"`
}

func (r REYContent) Validate() error {
	return validation.Validate.Struct(r)
}

// UnmarshalJSON implements [json.Unmarshaler].
func (r *REYContent) UnmarshalJSON(data []byte) error {
	type Alias REYContent

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	if err := (*REYContent)(&alias).Validate(); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	*r = REYContent(alias)
	return nil
}

// Hash implements Content.
func (r REYContent) Hash() (Hash, error) {
	idHash, err := hashing.HashBytes([]byte(r.XID))
	if err != nil {
		return Hash{}, fmt.Errorf("hash xid: %w", err)
	}

	usernameHash, err := hashing.HashBytes([]byte(strings.ToLower(r.XUsername)))
	if err != nil {
		return Hash{}, fmt.Errorf("hash username: %w", err)
	}

	hash, err := poseidon.Hash([]*big.Int{
		// fields ordered alphabetically regarding their JSON key
		new(big.Int).SetUint64(uint64(r.REYFaction)),
		new(big.Int).SetUint64(uint64(r.REYScoreAll)),
		new(big.Int).SetUint64(uint64(r.REYScoreGalactica)),
		idHash,
		usernameHash,
	})
	if err != nil {
		return Hash{}, err
	}

	return HashFromBigInt(hash), nil
}

// Standard implements Content. It always returns StandardREY.
func (r REYContent) Standard() Standard {
	return StandardREY
}
