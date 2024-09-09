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

package zkcertificate_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

func TestTwitterInputs_FFEncode(t *testing.T) {
	kycInputs := zkcertificate.TwitterInputs{
		CreatedAt:      time.Unix(1725914135, 0),
		ID:             "54291165",
		FollowersCount: 171,
		FollowingCount: 150,
		ListedCount:    2,
		TweetCount:     4351,
		Username:       "UnruhSpine",
		Verified:       true,
	}

	content, err := kycInputs.FFEncode()
	require.NoError(t, err)

	require.Equal(t, zkcertificate.TwitterContent{
		CreatedAt:      1725914135,
		ID:             mustHashFromString("6915188239842894578259451304274431896530928159819786989862125634701278515445"),
		FollowersCount: 171,
		FollowingCount: 150,
		ListedCount:    2,
		TweetCount:     4351,
		Username:       mustHashFromString("5886417250391205166336501621081802088080159499387114611838309812253065846588"),
		Verified:       true,
	}, content)
}
