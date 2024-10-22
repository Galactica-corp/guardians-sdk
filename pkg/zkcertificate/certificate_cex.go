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
	"math/big"

	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/shopspring/decimal"

	"github.com/galactica-corp/guardians-sdk/internal/validation"
)

// CEXInputs represent the input data for verification of trading on a centralized exchange.
type CEXInputs struct {
	TotalSwapVolume    decimal.Decimal `json:"totalSwapVolume" validate:"required,decimal_gt_0"`
	SwapVolumeYear     decimal.Decimal `json:"swapVolumeYear"`
	SwapVolumeHalfYear decimal.Decimal `json:"swapVolumeHalfYear"`
}

// FFEncode implements FFEncoder.
func (u *CEXInputs) FFEncode() (ExchangeContent, error) {
	return ExchangeContent{
		TotalSwapVolume:    dollarsToCentsTruncated(u.TotalSwapVolume),
		SwapVolumeYear:     dollarsToCentsTruncated(u.SwapVolumeYear),
		SwapVolumeHalfYear: dollarsToCentsTruncated(u.SwapVolumeHalfYear),
	}, nil
}

func (u *CEXInputs) Validate() error {
	return validation.Validate.Struct(u)
}

// UnmarshalJSON implements [json.Unmarshaler].
func (u *CEXInputs) UnmarshalJSON(data []byte) error {
	type Alias CEXInputs

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	if err := (*CEXInputs)(&alias).Validate(); err != nil {
		return err
	}

	*u = CEXInputs(alias)
	return nil
}

// CEXContent represent the hashed content of CEXInputs data.
type CEXContent struct {
	TotalSwapVolume    *big.Int `json:"totalSwapVolume"`
	SwapVolumeYear     *big.Int `json:"swapVolumeYear"`
	SwapVolumeHalfYear *big.Int `json:"swapVolumeHalfYear"`
}

// Hash implements Content.
func (u CEXContent) Hash() (Hash, error) {
	hash, err := poseidon.Hash([]*big.Int{
		// fields ordered alphabetically regarding their JSON key
		u.SwapVolumeHalfYear,
		u.SwapVolumeYear,
		u.TotalSwapVolume,
	})
	if err != nil {
		return Hash{}, err
	}

	return HashFromBigInt(hash), nil
}

// Standard implements Content. It always returns StandardCEX.
func (u CEXContent) Standard() Standard {
	return StandardCEX
}
