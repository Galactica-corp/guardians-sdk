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
	"strconv"
	"time"
)

// Timestamp represents a type that holds a time.Time value that is serialized as Unix timestamp.
type Timestamp time.Time

// Time returns the Timestamp as [time.Time].
func (t Timestamp) Time() time.Time {
	return time.Time(t)
}

// Unix returns the Unix timestamp of the Timestamp.
func (t Timestamp) Unix() int64 {
	return time.Time(t).Unix()
}

// MarshalJSON implements [json.Marshaler].
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Unix())
}

// UnmarshalJSON implements [json.Unmarshaler].
func (t *Timestamp) UnmarshalJSON(bytes []byte) error {
	text := string(bytes)

	if text == "null" {
		return nil
	}

	unix, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return err
	}

	*t = Timestamp(time.Unix(unix, 0))
	return nil
}
