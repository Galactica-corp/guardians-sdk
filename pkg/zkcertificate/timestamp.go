package zkcertificate

import (
	"encoding/json"
	"strconv"
	"time"
)

// Timestamp represents a type that holds a time.Time value that is serialized as Unix timestamp.
type Timestamp time.Time

// Unix returns the Unix timestamp of the Timestamp.
func (t Timestamp) Unix() int64 {
	return time.Time(t).Unix()
}

// MarshalJSON implements json.Marshaler interface.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Unix())
}

// UnmarshalJSON implements json.Unmarshaler interface.
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
