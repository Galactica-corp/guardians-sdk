package zkcertificate_test

import (
	"fmt"
	"math/big"
	"testing"
	"time"

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
		ExpirationDate:    zkcertificate.Timestamp(time.Unix(1779736098, 0)),
		StreetAndNumber:   "Bergstrasse 2",
		Postcode:          "9490",
		Town:              "Vaduz",
		Region:            "",
		Country:           "LIE",
		PassportID:        "F43G45EE6N67KL",
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
		ExpirationDate:    zkcertificate.Timestamp(time.Unix(1779736098, 0)),
		StreetAndNumber:   "Bergstrasse 2",
		Postcode:          "9490",
		Town:              "Vaduz",
		Region:            "",
		Country:           "LIE",
		PassportID:        "F43G45EE6N67KL",
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
		ExpirationDate:    zkcertificate.Timestamp(time.Unix(1779736098, 0)),
		StreetAndNumber:   "Bergstrasse 2",
		Postcode:          "9490",
		Town:              "Vaduz",
		Region:            "",
		Country:           "LIE",
		PassportID:        "F43G45EE6N67KL",
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
		ExpirationDate:    kycInputs.ExpirationDate,
		StreetAndNumber:   mustHashFromString("2716050888916668653838085209344746272579684295509938667842974355541145165593"),
		Postcode:          mustHashFromString("15490214341472802343037230413947290042671105911775897810655006802854190405490"),
		Town:              mustHashFromString("20114455475426887223405337417481199280779025506960488506094430597747419448504"),
		Region:            mustHashFromString("1"),
		Country:           mustHashFromString("5206272097955461433415670799568463787856046740863443717385317734933329763238"),
		Citizenship:       mustHashFromString("3704096963533554561380603784545722638824989796177284091907872955048029270902"),
		PassportID:        mustHashFromString("18400658382545519200829231663023181102885291022219012542155740714578393219901"),
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
		ExpirationDate:    zkcertificate.Timestamp(time.Unix(1779736098, 0)),
		StreetAndNumber:   mustHashFromString("2716050888916668653838085209344746272579684295509938667842974355541145165593"),
		Postcode:          mustHashFromString("15490214341472802343037230413947290042671105911775897810655006802854190405490"),
		Town:              mustHashFromString("20114455475426887223405337417481199280779025506960488506094430597747419448504"),
		Region:            mustHashFromString("1"),
		Country:           mustHashFromString("5206272097955461433415670799568463787856046740863443717385317734933329763238"),
		Citizenship:       mustHashFromString("3704096963533554561380603784545722638824989796177284091907872955048029270902"),
		PassportID:        mustHashFromString("18400658382545519200829231663023181102885291022219012542155740714578393219901"),
	}

	hash, err := kycContent.Hash()
	require.NoError(t, err)
	require.Equal(t, mustHashFromString("6485519903590620859988164613335857332142930963704620586195124368699722613961"), hash)
}

func mustHashFromString(s string) zkcertificate.Hash {
	v, valid := new(big.Int).SetString(s, 10)
	if !valid {
		panic(fmt.Sprintf("invalid big int %q", s))
	}

	return zkcertificate.HashFromBigInt(v)
}
