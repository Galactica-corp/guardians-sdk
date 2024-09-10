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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/shopspring/decimal"

	"github.com/galactica-corp/guardians-sdk/internal/validation"
)

// ExchangeInputs represent the input data for verification of trading on an exchange.
type ExchangeInputs struct {
	Address            string          `json:"address" validate:"required,eth_addr"`
	TotalSwapVolume    decimal.Decimal `json:"totalSwapVolume"`
	SwapVolumeYear     decimal.Decimal `json:"swapVolumeYear"`
	SwapVolumeHalfYear decimal.Decimal `json:"swapVolumeHalfYear"`
}

// FFEncode implements FFEncoder.
func (u *ExchangeInputs) FFEncode() (ExchangeContent, error) {
	addressBytes, err := hex.DecodeString(u.Address[2:])
	if err != nil {
		return ExchangeContent{}, fmt.Errorf("invalid address: %v", err)
	}

	addressHash, err := poseidon.HashBytes(addressBytes)
	if err != nil {
		return ExchangeContent{}, fmt.Errorf("hash address: %v", err)
	}

	return ExchangeContent{
		Address:            HashFromBigInt(addressHash),
		TotalSwapVolume:    dollarsToCentsTruncated(u.TotalSwapVolume),
		SwapVolumeYear:     dollarsToCentsTruncated(u.SwapVolumeYear),
		SwapVolumeHalfYear: dollarsToCentsTruncated(u.SwapVolumeHalfYear),
	}, nil
}

func (u *ExchangeInputs) Validate() error {
	return validation.Validate.Struct(u)
}

// UnmarshalJSON implements [json.Unmarshaler].
func (u *ExchangeInputs) UnmarshalJSON(data []byte) error {
	type Alias ExchangeInputs

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	if err := (*ExchangeInputs)(&alias).Validate(); err != nil {
		return err
	}

	*u = ExchangeInputs(alias)
	return nil
}

// ExchangeContent represent the hashed content of ExchangeInputs data.
type ExchangeContent struct {
	Address            Hash     `json:"address"`
	TotalSwapVolume    *big.Int `json:"totalSwapVolume"`
	SwapVolumeYear     *big.Int `json:"swapVolumeYear"`
	SwapVolumeHalfYear *big.Int `json:"swapVolumeHalfYear"`
}

// Hash implements Content.
func (u ExchangeContent) Hash() (Hash, error) {
	hash, err := poseidon.Hash([]*big.Int{
		u.Address.BigInt(),
		u.TotalSwapVolume,
		u.SwapVolumeYear,
		u.SwapVolumeHalfYear,
	})
	if err != nil {
		return Hash{}, err
	}

	return HashFromBigInt(hash), nil
}

// Standard implements Content. It always returns StandardExchange.
func (u ExchangeContent) Standard() Standard {
	return StandardExchange
}
