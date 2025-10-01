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
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/v4/pkg/zkcertificate"
)

func TestKYCInputs_Validate(t *testing.T) {
	kycInputs := zkcertificate.KYCContent{
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
	kycInputs := zkcertificate.KYCContent{
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
		"Key: 'KYCContent.Surname' Error:Field validation for 'Surname' failed on the 'required' tag",
	)
}

func TestKYCContent_Hash(t *testing.T) {
	kycInputs := zkcertificate.KYCContent{
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

	hash, err := kycInputs.Hash()
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
	kycInputs := zkcertificate.KYCContent{
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

	idHash, err := kycInputs.IDHash()
	require.NoError(t, err)
	require.Equal(t, mustHashFromString("5606249467514511053794405094500185048153115639175989255899284975583206743707"), idHash)
}
