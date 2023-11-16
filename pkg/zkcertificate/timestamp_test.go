package zkcertificate_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
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
