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

	"github.com/iden3/go-iden3-crypto/v2/poseidon"
	"github.com/shopspring/decimal"

	"github.com/galactica-corp/guardians-sdk/v4/internal/validation"
)

// BlumContent represents the data for verification of a Blum account.
type BlumContent struct {
	TelegramID    int64           `json:"telegramId" validate:"required"`
	ActivityScore decimal.Decimal `json:"activityScore" validate:"required"`
	SybilScore    decimal.Decimal `json:"sybilScore" validate:"required"`
}

func (b BlumContent) Validate() error {
	return validation.Validate.Struct(b)
}

// UnmarshalJSON implements [json.Unmarshaler].
func (b *BlumContent) UnmarshalJSON(data []byte) error {
	type Alias BlumContent

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	if err := (*BlumContent)(&alias).Validate(); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	*b = BlumContent(alias)
	return nil
}

// Hash implements Content.
func (b BlumContent) Hash() (Hash, error) {
	hash, err := poseidon.Hash([]*big.Int{
		// fields ordered alphabetically regarding their JSON key
		scoreToFixedPoint(b.ActivityScore, 18),
		scoreToFixedPoint(b.SybilScore, 18),
		big.NewInt(b.TelegramID),
	})
	if err != nil {
		return Hash{}, err
	}

	return HashFromBigInt(hash), nil
}

// scoreToFixedPoint converts a decimal score to a big.Int by scaling with 10^decimals
// to preserve decimal precision for ZK-proof compatibility.
func scoreToFixedPoint(score decimal.Decimal, decimals int) *big.Int {
	return score.Shift(int32(decimals)).BigInt()
}

// Standard implements Content. It always returns StandardBlum.
func (b BlumContent) Standard() Standard {
	return StandardBlum
}
