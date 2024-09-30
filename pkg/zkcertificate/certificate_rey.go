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

	"github.com/iden3/go-iden3-crypto/poseidon"

	"github.com/galactica-corp/guardians-sdk/internal/validation"
)

// REYInputs represents the input data for the REY verification.
type REYInputs struct {
	XID               string `json:"x_id" validate:"required,number"`
	XUsername         string `json:"x_username" validate:"required,min=4,max=15"`
	REYScoreAll       uint   `json:"rey_score_all"`
	REYScoreGalactica uint   `json:"rey_score_galactica"`
	REYFaction        uint   `json:"rey_faction"`
}

// FFEncode implements FFEncoder.
func (r *REYInputs) FFEncode() (REYContent, error) {
	idHash, err := poseidon.HashBytes([]byte(r.XID))
	if err != nil {
		return REYContent{}, fmt.Errorf("hash xid: %w", err)
	}

	usernameHash, err := poseidon.HashBytes([]byte(r.XUsername))
	if err != nil {
		return REYContent{}, fmt.Errorf("hash username: %w", err)
	}

	return REYContent{
		XID:               HashFromBigInt(idHash),
		XUsername:         HashFromBigInt(usernameHash),
		REYScoreAll:       r.REYScoreAll,
		REYScoreGalactica: r.REYScoreGalactica,
		REYFaction:        r.REYFaction,
	}, nil
}

func (r *REYInputs) Validate() error {
	return validation.Validate.Struct(r)
}

// UnmarshalJSON implements [json.Unmarshaler].
func (r *REYInputs) UnmarshalJSON(data []byte) error {
	type Alias REYInputs

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	if err := (*REYInputs)(&alias).Validate(); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	*r = REYInputs(alias)
	return nil
}

// REYContent represents the hashed content of REYInputs data.
type REYContent struct {
	XID               Hash `json:"x_id"`
	XUsername         Hash `json:"x_username"`
	REYScoreAll       uint `json:"rey_score_all"`
	REYScoreGalactica uint `json:"rey_score_galactica"`
	REYFaction        uint `json:"rey_faction"`
}

// Hash implements Content.
func (c REYContent) Hash() (Hash, error) {
	hash, err := poseidon.Hash([]*big.Int{
		// fields ordered alphabetically regarding their JSON key
		new(big.Int).SetUint64(uint64(c.REYFaction)),
		new(big.Int).SetUint64(uint64(c.REYScoreAll)),
		new(big.Int).SetUint64(uint64(c.REYScoreGalactica)),
		c.XID.BigInt(),
		c.XUsername.BigInt(),
	})
	if err != nil {
		return Hash{}, err
	}

	return HashFromBigInt(hash), nil
}

// Standard implements Content. It always returns StandardREY.
func (c REYContent) Standard() Standard {
	return StandardREY
}
