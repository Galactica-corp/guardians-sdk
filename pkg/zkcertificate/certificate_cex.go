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
	"math/big"

	"github.com/iden3/go-iden3-crypto/v2/poseidon"
	"github.com/shopspring/decimal"

	"github.com/galactica-corp/guardians-sdk/internal/validation"
)

// CEXContent represent the data for verification of trading on a centralized exchange.
type CEXContent struct {
	TotalSwapVolume    decimal.Decimal `json:"totalSwapVolume" validate:"required,decimal_gt_0"`
	SwapVolumeYear     decimal.Decimal `json:"swapVolumeYear"`
	SwapVolumeHalfYear decimal.Decimal `json:"swapVolumeHalfYear"`
}

func (x CEXContent) Validate() error {
	return validation.Validate.Struct(x)
}

// UnmarshalJSON implements [json.Unmarshaler].
func (x *CEXContent) UnmarshalJSON(data []byte) error {
	type Alias CEXContent

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	if err := (*CEXContent)(&alias).Validate(); err != nil {
		return err
	}

	*x = CEXContent(alias)
	return nil
}

// Hash implements Content.
func (x CEXContent) Hash() (Hash, error) {
	hash, err := poseidon.Hash([]*big.Int{
		// fields ordered alphabetically regarding their JSON key
		dollarsToCentsTruncated(x.SwapVolumeHalfYear),
		dollarsToCentsTruncated(x.SwapVolumeYear),
		dollarsToCentsTruncated(x.TotalSwapVolume),
	})
	if err != nil {
		return Hash{}, err
	}

	return HashFromBigInt(hash), nil
}

// Standard implements Content. It always returns StandardCEX.
func (x CEXContent) Standard() Standard {
	return StandardCEX
}
