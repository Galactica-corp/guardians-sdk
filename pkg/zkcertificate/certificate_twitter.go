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

// TwitterInputs represents the input data for X/Twitter verification.
type TwitterInputs struct {
	CreatedAt      time.Time `json:"createdAt" validate:"required,lt"`
	ID             string    `json:"id" validate:"required,number"`
	FollowersCount uint      `json:"followersCount"`
	FollowingCount uint      `json:"followingCount"`
	ListedCount    uint      `json:"listedCount"`
	TweetCount     uint      `json:"tweetCount"`
	Username       string    `json:"username" validate:"required,min=4,max=15"`
	Verified       bool      `json:"verified"`
}

// FFEncode implements FFEncoder.
func (t TwitterInputs) FFEncode() (TwitterContent, error) {
	idHash, err := poseidon.HashBytes([]byte(t.ID))
	if err != nil {
		return TwitterContent{}, fmt.Errorf("hash ID: %w", err)
	}

	usernameHash, err := poseidon.HashBytes([]byte(strings.ToLower(t.Username)))
	if err != nil {
		return TwitterContent{}, fmt.Errorf("hash username: %w", err)
	}

	return TwitterContent{
		CreatedAt:      t.CreatedAt.Unix(),
		ID:             HashFromBigInt(idHash),
		FollowersCount: t.FollowersCount,
		FollowingCount: t.FollowingCount,
		ListedCount:    t.ListedCount,
		TweetCount:     t.TweetCount,
		Username:       HashFromBigInt(usernameHash),
		Verified:       t.Verified,
	}, nil
}

func (t *TwitterInputs) Validate() error {
	return validation.Validate.Struct(t)
}

// UnmarshalJSON implements [json.Unmarshaler].
func (t *TwitterInputs) UnmarshalJSON(data []byte) error {
	type Alias TwitterInputs

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	if err := (*TwitterInputs)(&alias).Validate(); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	*t = TwitterInputs(alias)
	return nil
}

// TwitterContent represents the hashed content of TwitterInputs data.
type TwitterContent struct {
	CreatedAt      int64 `json:"createdAt"`
	ID             Hash  `json:"id"`
	FollowersCount uint  `json:"followersCount"`
	FollowingCount uint  `json:"followingCount"`
	ListedCount    uint  `json:"listedCount"`
	TweetCount     uint  `json:"tweetCount"`
	Username       Hash  `json:"username"`
	Verified       bool  `json:"verified"`
}

// Hash implements Content.
func (t TwitterContent) Hash() (Hash, error) {
	var verified *big.Int
	if t.Verified {
		verified = big.NewInt(1)
	} else {
		verified = big.NewInt(0)
	}

	hash, err := poseidon.Hash([]*big.Int{
		// fields ordered alphabetically regarding their JSON key
		big.NewInt(t.CreatedAt),
		new(big.Int).SetUint64(uint64(t.FollowersCount)),
		new(big.Int).SetUint64(uint64(t.FollowingCount)),
		t.ID.BigInt(),
		new(big.Int).SetUint64(uint64(t.ListedCount)),
		new(big.Int).SetUint64(uint64(t.TweetCount)),
		t.Username.BigInt(),
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
