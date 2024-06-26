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
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

func TestSimpleJSON_Validate(t *testing.T) {
	simpleJSON := zkcertificate.SimpleJSON{
		"name":      "John Doe",
		"age":       "30",
		"isMarried": "true",
		"height":    "1.75",
		"birthday":  "1990-01-01T00:00:00Z",
	}

	require.NoError(t, simpleJSON.Validate())
}

func TestSimpleJSON_FFEncode(t *testing.T) {
	simpleJSON := zkcertificate.SimpleJSON{
		"name":      "John Doe",
		"age":       "30",
		"isMarried": "true",
		"height":    "1.75",
		"birthday":  "1990-01-01T00:00:00Z",
	}

	content, err := simpleJSON.FFEncode()
	require.NoError(t, err)

	require.Equal(t, zkcertificate.SimpleJSONContent{
		mustHashFromString("15354103676486910392335921695264068620680648317671388604584937927146958457868"),
		mustHashFromString("15224306929871956365071811469757107548610844039491110436456644707663745462126"),
		mustHashFromString("10390072210063895581917865108194000409658292591398372775041595820377437637311"),
		mustHashFromString("13428808271822965993111337918515956004850801358226726664783893758011128965986"),
		mustHashFromString("1375746358235177439421528943531890815720286426595868754203705803877498795804"),
	}, content)
}

func TestSimpleJSONContent_Hash(t *testing.T) {
	simpleJSONContent := zkcertificate.SimpleJSONContent{
		mustHashFromString("15354103676486910392335921695264068620680648317671388604584937927146958457868"),
		mustHashFromString("15224306929871956365071811469757107548610844039491110436456644707663745462126"),
		mustHashFromString("10390072210063895581917865108194000409658292591398372775041595820377437637311"),
		mustHashFromString("13428808271822965993111337918515956004850801358226726664783893758011128965986"),
		mustHashFromString("1375746358235177439421528943531890815720286426595868754203705803877498795804"),
	}

	hash, err := simpleJSONContent.Hash()
	require.NoError(t, err)
	require.Equal(t, "8797979680535626597025031675791027936583836473963438577360708138820501712641", hash.String())
}

func TestSimpleJSON_UnmarshalJSON(t *testing.T) {
	jsonData := `{
		"name": "John Doe",
		"age": "30",
		"isMarried": "true",
		"height": "1.75",
		"birthday": "1990-01-01T00:00:00Z"
	}`

	var simpleJSON zkcertificate.SimpleJSON
	err := json.Unmarshal([]byte(jsonData), &simpleJSON)
	require.NoError(t, err)

	expected := zkcertificate.SimpleJSON{
		"name":      "John Doe",
		"age":       "30",
		"isMarried": "true",
		"height":    "1.75",
		"birthday":  "1990-01-01T00:00:00Z",
	}
	require.Equal(t, expected, simpleJSON)
}

func TestSimpleJSON_UnmarshalJSON_Negative(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "Invalid JSON Format", input: "{name: John Doe}", wantErr: true},
		{name: "Unsupported Types", input: `{"unsupported": {"nested": "value"}}`, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var js zkcertificate.SimpleJSON
			err := json.Unmarshal([]byte(tt.input), &js)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimpleJSON.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimpleJSON_UnmarshalJSON_InvalidStructure(t *testing.T) {
	jsonData := `["not", "a", "json", "object"]`

	var simpleJSON zkcertificate.SimpleJSON
	err := json.Unmarshal([]byte(jsonData), &simpleJSON)
	require.Error(t, err)
}

func TestSimpleJSONContent_Standard(t *testing.T) {
	content := zkcertificate.SimpleJSONContent{}

	standard := content.Standard()

	require.Equal(t, zkcertificate.StandardSimpleJSON, standard)
}
