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

package zkcertificate_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/v4/pkg/zkcertificate"
)

func TestTimestamp_Unix(t *testing.T) {
	now := time.Now()
	timestamp := zkcertificate.Timestamp(now)

	require.Equal(t, now.Unix(), timestamp.Unix())
}

func TestTimestamp_JSON(t *testing.T) {
	unix := time.Unix(1700000000, 0)
	timestamp := zkcertificate.Timestamp(unix)

	data, err := json.Marshal(timestamp)
	require.NoError(t, err)
	require.Equal(t, []byte(`1700000000`), data)

	var deserialized zkcertificate.Timestamp
	require.NoError(t, json.Unmarshal(data, &deserialized))
	require.Equal(t, timestamp, deserialized)
}

func TestTimestamp_JSON_null(t *testing.T) {
	var deserialized zkcertificate.Timestamp
	require.NoError(t, json.Unmarshal([]byte(`null`), &deserialized))
	require.Zero(t, deserialized)
}
