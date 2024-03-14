package zkcertificate_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

func TestSimpleJSON_Validate(t *testing.T) {
	simpleJSON := zkcertificate.SimpleJSON{
		"name":      "John Doe",
		"age":       30,
		"isMarried": true,
		"height":    1.75,
		"birthday":  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	require.NoError(t, simpleJSON.Validate())
}

func TestSimpleJSON_FFEncode(t *testing.T) {
	simpleJSON := zkcertificate.SimpleJSON{
		"name":      "John Doe",
		"age":       30,
		"isMarried": true,
		"height":    1.75,
		"birthday":  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	content, err := simpleJSON.FFEncode()
	require.NoError(t, err)

	require.Equal(t, zkcertificate.SimpleJSONContent{
		"name":      mustHashFromString("1375746358235177439421528943531890815720286426595868754203705803877498795804"),
		"age":       mustHashFromString("15354103676486910392335921695264068620680648317671388604584937927146958457868"),
		"isMarried": mustHashFromString("13428808271822965993111337918515956004850801358226726664783893758011128965986"),
		"height":    mustHashFromString("10390072210063895581917865108194000409658292591398372775041595820377437637311"),
		"birthday":  mustHashFromString("15224306929871956365071811469757107548610844039491110436456644707663745462126"),
	}, content)
}

func TestSimpleJSONContent_Hash(t *testing.T) {
	simpleJSONContent := zkcertificate.SimpleJSONContent{
		"name":      mustHashFromString("1375746358235177439421528943531890815720286426595868754203705803877498795804"),
		"age":       mustHashFromString("15354103676486910392335921695264068620680648317671388604584937927146958457868"),
		"isMarried": mustHashFromString("13428808271822965993111337918515956004850801358226726664783893758011128965986"),
		"height":    mustHashFromString("10390072210063895581917865108194000409658292591398372775041595820377437637311"),
		"birthday":  mustHashFromString("15224306929871956365071811469757107548610844039491110436456644707663745462126"),
	}

	hash, err := simpleJSONContent.Hash()
	require.NoError(t, err)
	require.Equal(t, "8797979680535626597025031675791027936583836473963438577360708138820501712641", hash.String())
}

func TestSimpleJSON_UnmarshalJSON(t *testing.T) {
	jsonData := `{
		"name": "John Doe",
		"age": 30,
		"isMarried": true,
		"height": 1.75,
		"birthday": "1990-01-01T00:00:00Z"
	}`

	var simpleJSON zkcertificate.SimpleJSON
	err := json.Unmarshal([]byte(jsonData), &simpleJSON)
	require.NoError(t, err)

	expected := zkcertificate.SimpleJSON{
		"name":      "John Doe",
		"age":       float64(30),
		"isMarried": true,
		"height":    1.75,
		"birthday":  "1990-01-01T00:00:00Z",
	}
	require.Equal(t, expected, simpleJSON)
}
