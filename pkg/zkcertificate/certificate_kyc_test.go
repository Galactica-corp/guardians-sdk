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
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

func TestKYCInputs_Validate(t *testing.T) {
	kycInputs := zkcertificate.KYCInputs{
		Surname:           "Doe",
		Forename:          "John",
		MiddleName:        "",
		YearOfBirth:       1989,
		MonthOfBirth:      5,
		DayOfBirth:        28,
		Citizenship:       "SWE",
		VerificationLevel: 1,
		StreetAndNumber:   "Bergstrasse 2",
		Postcode:          "9490",
		Town:              "Vaduz",
		Region:            "",
		Country:           "LIE",
	}

	require.NoError(t, kycInputs.Validate())
}

func TestKYCInputs_Validate_invalid(t *testing.T) {
	kycInputs := zkcertificate.KYCInputs{
		Surname:           "",
		Forename:          "John",
		MiddleName:        "",
		YearOfBirth:       1989,
		MonthOfBirth:      5,
		DayOfBirth:        28,
		Citizenship:       "SWE",
		VerificationLevel: 1,
		StreetAndNumber:   "Bergstrasse 2",
		Postcode:          "9490",
		Town:              "Vaduz",
		Region:            "",
		Country:           "LIE",
	}

	require.EqualError(
		t,
		kycInputs.Validate(),
		"Key: 'KYCInputs.Surname' Error:Field validation for 'Surname' failed on the 'required' tag",
	)
}

func TestKYCInputs_FFEncode(t *testing.T) {
	kycInputs := zkcertificate.KYCInputs{
		Surname:           "Doe",
		Forename:          "John",
		MiddleName:        "",
		YearOfBirth:       1989,
		MonthOfBirth:      5,
		DayOfBirth:        28,
		Citizenship:       "SWE",
		VerificationLevel: 1,
		StreetAndNumber:   "Bergstrasse 2",
		Postcode:          "9490",
		Town:              "Vaduz",
		Region:            "",
		Country:           "LIE",
	}

	content, err := kycInputs.FFEncode()
	require.NoError(t, err)

	require.Equal(t, zkcertificate.KYCContent{
		Surname:           mustHashFromString("9394571395257706853209381240690921572244805705054092251193796665368637684518"),
		Forename:          mustHashFromString("12508415106905269703668517379769578980197323180488076414504369770793910945017"),
		MiddleName:        mustHashFromString("1"),
		YearOfBirth:       kycInputs.YearOfBirth,
		MonthOfBirth:      kycInputs.MonthOfBirth,
		DayOfBirth:        kycInputs.DayOfBirth,
		VerificationLevel: kycInputs.VerificationLevel,
		StreetAndNumber:   mustHashFromString("2716050888916668653838085209344746272579684295509938667842974355541145165593"),
		Postcode:          mustHashFromString("15490214341472802343037230413947290042671105911775897810655006802854190405490"),
		Town:              mustHashFromString("20114455475426887223405337417481199280779025506960488506094430597747419448504"),
		Region:            mustHashFromString("1"),
		Country:           mustHashFromString("5206272097955461433415670799568463787856046740863443717385317734933329763238"),
		Citizenship:       mustHashFromString("3704096963533554561380603784545722638824989796177284091907872955048029270902"),
	}, content)
}

func TestKYCContent_Hash(t *testing.T) {
	kycContent := zkcertificate.KYCContent{
		Surname:           mustHashFromString("9394571395257706853209381240690921572244805705054092251193796665368637684518"),
		Forename:          mustHashFromString("12508415106905269703668517379769578980197323180488076414504369770793910945017"),
		MiddleName:        mustHashFromString("1"),
		YearOfBirth:       1989,
		MonthOfBirth:      5,
		DayOfBirth:        28,
		VerificationLevel: zkcertificate.KYCVerificationLevelPassedKYC,
		StreetAndNumber:   mustHashFromString("2716050888916668653838085209344746272579684295509938667842974355541145165593"),
		Postcode:          mustHashFromString("15490214341472802343037230413947290042671105911775897810655006802854190405490"),
		Town:              mustHashFromString("20114455475426887223405337417481199280779025506960488506094430597747419448504"),
		Region:            mustHashFromString("1"),
		Country:           mustHashFromString("5206272097955461433415670799568463787856046740863443717385317734933329763238"),
		Citizenship:       mustHashFromString("3704096963533554561380603784545722638824989796177284091907872955048029270902"),
	}

	hash, err := kycContent.Hash()
	require.NoError(t, err)
	require.Equal(t, mustHashFromString("13498937448046187479975980844060005602014574276619662435996314654414855730267"), hash)
}

func mustHashFromString(s string) zkcertificate.Hash {
	v, valid := new(big.Int).SetString(s, 10)
	if !valid {
		panic(fmt.Sprintf("invalid big int %q", s))
	}

	return zkcertificate.HashFromBigInt(v)
}

func TestKYCContent_IDHash(t *testing.T) {
	kycContent := zkcertificate.KYCContent{
		Surname:           mustHashFromString("9394571395257706853209381240690921572244805705054092251193796665368637684518"),
		Forename:          mustHashFromString("12508415106905269703668517379769578980197323180488076414504369770793910945017"),
		MiddleName:        mustHashFromString("1"),
		YearOfBirth:       1989,
		MonthOfBirth:      5,
		DayOfBirth:        28,
		VerificationLevel: zkcertificate.KYCVerificationLevelPassedKYC,
		StreetAndNumber:   mustHashFromString("2716050888916668653838085209344746272579684295509938667842974355541145165593"),
		Postcode:          mustHashFromString("15490214341472802343037230413947290042671105911775897810655006802854190405490"),
		Town:              mustHashFromString("20114455475426887223405337417481199280779025506960488506094430597747419448504"),
		Region:            mustHashFromString("1"),
		Country:           mustHashFromString("5206272097955461433415670799568463787856046740863443717385317734933329763238"),
		Citizenship:       mustHashFromString("3704096963533554561380603784545722638824989796177284091907872955048029270902"),
	}

	idHash, err := kycContent.IDHash()
	require.NoError(t, err)
	require.Equal(t, mustHashFromString("5606249467514511053794405094500185048153115639175989255899284975583206743707"), idHash)
}
