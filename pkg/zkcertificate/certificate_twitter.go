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
	"strings"
	"time"

	"github.com/iden3/go-iden3-crypto/poseidon"

	"github.com/galactica-corp/guardians-sdk/internal/validation"
)

// TwitterContent represents the data for X/Twitter verification.
type TwitterContent struct {
	CreatedAt      time.Time `json:"createdAt" validate:"required,lt"`
	ID             string    `json:"id" validate:"required,number"`
	FollowersCount uint      `json:"followersCount"`
	FollowingCount uint      `json:"followingCount"`
	ListedCount    uint      `json:"listedCount"`
	TweetCount     uint      `json:"tweetCount"`
	Username       string    `json:"username" validate:"required,min=4,max=15"`
	Verified       bool      `json:"verified"`
}

func (t *TwitterContent) Validate() error {
	return validation.Validate.Struct(t)
}

// UnmarshalJSON implements [json.Unmarshaler].
func (t *TwitterContent) UnmarshalJSON(data []byte) error {
	type Alias TwitterContent

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	if err := (*TwitterContent)(&alias).Validate(); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	*t = TwitterContent(alias)
	return nil
}

// Hash implements Content.
func (t TwitterContent) Hash() (Hash, error) {
	idHash, err := poseidon.HashBytes([]byte(t.ID))
	if err != nil {
		return Hash{}, fmt.Errorf("hash ID: %w", err)
	}

	usernameHash, err := poseidon.HashBytes([]byte(strings.ToLower(t.Username)))
	if err != nil {
		return Hash{}, fmt.Errorf("hash username: %w", err)
	}

	var verified *big.Int
	if t.Verified {
		verified = big.NewInt(1)
	} else {
		verified = big.NewInt(0)
	}

	hash, err := poseidon.Hash([]*big.Int{
		// fields ordered alphabetically regarding their JSON key
		big.NewInt(t.CreatedAt.Unix()),
		new(big.Int).SetUint64(uint64(t.FollowersCount)),
		new(big.Int).SetUint64(uint64(t.FollowingCount)),
		idHash,
		new(big.Int).SetUint64(uint64(t.ListedCount)),
		new(big.Int).SetUint64(uint64(t.TweetCount)),
		usernameHash,
		verified,
	})
	if err != nil {
		return Hash{}, err
	}

	return HashFromBigInt(hash), nil
}

// Standard implements Content. It always returns StandardTwitter.
func (t TwitterContent) Standard() Standard {
	return StandardTwitter
}
