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
	"time"

	"github.com/iden3/go-iden3-crypto/poseidon"

	"github.com/galactica-corp/guardians-sdk/internal/validation"
)

// TelegramInputs represents the input data for verification of Telegram account.
type TelegramInputs struct {
	ActiveDaysCount                   uint      `json:"activeDaysCount" validate:"required"`
	ContactWithAtLeast10MessagesCount uint      `json:"contactWithAtLeast10MessagesCount"`
	CreatedAt                         time.Time `json:"createdAt" validate:"required,lt"`
	MeanMonthlyMessageCount           uint      `json:"meanMonthlyMessageCount" validate:"required"`
}

// FFEncode implements FFEncoder.
func (t TelegramInputs) FFEncode() (TelegramContent, error) {
	return TelegramContent{
		ActiveDaysCount:                   t.ActiveDaysCount,
		ContactWithAtLeast10MessagesCount: t.ContactWithAtLeast10MessagesCount,
		CreatedAt:                         t.CreatedAt.Unix(),
		MeanMonthlyMessageCount:           t.MeanMonthlyMessageCount,
	}, nil
}

func (t *TelegramInputs) Validate() error {
	return validation.Validate.Struct(&t)
}

// UnmarshalJSON implements [json.Unmarshaler].
func (t *TelegramInputs) UnmarshalJSON(data []byte) error {
	type Alias TelegramInputs

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	if err := (*TelegramInputs)(&alias).Validate(); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	*t = TelegramInputs(alias)
	return nil
}

// TelegramContent represents the hashed content of TelegramInputs data.
type TelegramContent struct {
	ActiveDaysCount                   uint  `json:"activeDaysCount"`
	ContactWithAtLeast10MessagesCount uint  `json:"contactWithAtLeast10MessagesCount"`
	CreatedAt                         int64 `json:"createdAt"`
	MeanMonthlyMessageCount           uint  `json:"meanMonthlyMessageCount"`
}

// Hash implements Content.
func (t TelegramContent) Hash() (Hash, error) {
	hash, err := poseidon.Hash([]*big.Int{
		// fields ordered alphabetically regarding their JSON key
		new(big.Int).SetUint64(uint64(t.ActiveDaysCount)),
		new(big.Int).SetUint64(uint64(t.ContactWithAtLeast10MessagesCount)),
		big.NewInt(t.CreatedAt),
		new(big.Int).SetUint64(uint64(t.MeanMonthlyMessageCount)),
	})
	if err != nil {
		return Hash{}, err
	}

	return HashFromBigInt(hash), nil
}

// Standard implements Content. It always returns StandardTelegram.
func (t TelegramContent) Standard() Standard {
	return StandardTelegram
}
