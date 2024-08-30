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
	"fmt"
	"math/big"
	"strconv"

	"github.com/iden3/go-iden3-crypto/poseidon"

	"github.com/galactica-corp/guardians-sdk/internal/validation"
)

// KYCInputs represents the input data for Know Your Customer (KYC) verification.
// It contains various fields required for identity verification and validation.
type KYCInputs struct {
	Surname           string               `json:"surname" validate:"required"`
	Forename          string               `json:"forename" validate:"required"`
	MiddleName        string               `json:"middlename" validate:"omitempty"`
	YearOfBirth       uint16               `json:"yearOfBirth" validate:"required"`
	MonthOfBirth      uint8                `json:"monthOfBirth" validate:"required,gte=1,lte=12"`
	DayOfBirth        uint8                `json:"dayOfBirth" validate:"required,gte=1,lte=31"`
	Citizenship       string               `json:"citizenship" validate:"required,iso3166_1_alpha3"`
	VerificationLevel KYCVerificationLevel `json:"verificationLevel"`
	StreetAndNumber   string               `json:"streetAndNumber"`
	Postcode          string               `json:"postcode"`
	Town              string               `json:"town"`
	Region            string               `json:"region" validate:"omitempty,iso3166_2"`
	Country           string               `json:"country" validate:"required,iso3166_1_alpha3"`
}

// FFEncode implements FFEncoder.
func (k KYCInputs) FFEncode() (KYCContent, error) {
	surnameHash, err := poseidon.HashBytes([]byte(k.Surname))
	if err != nil {
		return KYCContent{}, fmt.Errorf("hash surname: %w", err)
	}

	forenameHash, err := poseidon.HashBytes([]byte(k.Forename))
	if err != nil {
		return KYCContent{}, fmt.Errorf("hash forename: %w", err)
	}

	middleNameHash, err := poseidon.HashBytes([]byte(k.MiddleName))
	if err != nil {
		return KYCContent{}, fmt.Errorf("hash middle name: %w", err)
	}

	streetAndNumberHash, err := poseidon.HashBytes([]byte(k.StreetAndNumber))
	if err != nil {
		return KYCContent{}, fmt.Errorf("hash street and number: %w", err)
	}

	postcodeHash, err := poseidon.HashBytes([]byte(k.Postcode))
	if err != nil {
		return KYCContent{}, fmt.Errorf("hash postcode: %w", err)
	}

	townHash, err := poseidon.HashBytes([]byte(k.Town))
	if err != nil {
		return KYCContent{}, fmt.Errorf("hash town: %w", err)
	}

	regionHash, err := poseidon.HashBytes([]byte(k.Region))
	if err != nil {
		return KYCContent{}, fmt.Errorf("hash region: %w", err)
	}

	countryHash, err := poseidon.HashBytes([]byte(k.Country))
	if err != nil {
		return KYCContent{}, fmt.Errorf("hash country: %w", err)
	}

	citizenshipHash, err := poseidon.HashBytes([]byte(k.Citizenship))
	if err != nil {
		return KYCContent{}, fmt.Errorf("hash citizenship: %w", err)
	}

	return KYCContent{
		Surname:           HashFromBigInt(surnameHash),
		Forename:          HashFromBigInt(forenameHash),
		MiddleName:        HashFromBigInt(middleNameHash),
		YearOfBirth:       k.YearOfBirth,
		MonthOfBirth:      k.MonthOfBirth,
		DayOfBirth:        k.DayOfBirth,
		VerificationLevel: k.VerificationLevel,
		StreetAndNumber:   HashFromBigInt(streetAndNumberHash),
		Postcode:          HashFromBigInt(postcodeHash),
		Town:              HashFromBigInt(townHash),
		Region:            HashFromBigInt(regionHash),
		Country:           HashFromBigInt(countryHash),
		Citizenship:       HashFromBigInt(citizenshipHash),
	}, nil
}

// Validate performs validation on the KYCInputs instance using the struct tags specified
// for field validation. It checks that the fields of the KYCInputs struct adhere to
// the defined validation rules.
func (k *KYCInputs) Validate() error {
	return validation.Validate.Struct(k)
}

// UnmarshalJSON implements [json.Unmarshaler].
func (k *KYCInputs) UnmarshalJSON(data []byte) error {
	type Alias KYCInputs

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	if err := (*KYCInputs)(&alias).Validate(); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	*k = KYCInputs(alias)
	return nil
}

// KYCContent represents the hashed content of KYC (Know Your Customer) data.
// It contains hashed values for various fields related to identity and verification.
type KYCContent struct {
	Surname           Hash                 `json:"surname"`
	Forename          Hash                 `json:"forename"`
	MiddleName        Hash                 `json:"middlename"`
	YearOfBirth       uint16               `json:"yearOfBirth"`
	MonthOfBirth      uint8                `json:"monthOfBirth"`
	DayOfBirth        uint8                `json:"dayOfBirth"`
	VerificationLevel KYCVerificationLevel `json:"verificationLevel"`
	StreetAndNumber   Hash                 `json:"streetAndNumber"`
	Postcode          Hash                 `json:"postcode"`
	Town              Hash                 `json:"town"`
	Region            Hash                 `json:"region"`
	Country           Hash                 `json:"country"`
	Citizenship       Hash                 `json:"citizenship"`
}

// Standard returns the standard associated with the KYCContent, which is StandardKYC.
func (c KYCContent) Standard() Standard {
	return StandardKYC
}

// Hash computes and returns the hash of the KYCContent instance.
func (c KYCContent) Hash() (Hash, error) {
	hash, err := poseidon.Hash([]*big.Int{
		c.Surname.BigInt(),
		c.Forename.BigInt(),
		c.MiddleName.BigInt(),
		big.NewInt(int64(c.YearOfBirth)),
		big.NewInt(int64(c.MonthOfBirth)),
		big.NewInt(int64(c.DayOfBirth)),
		big.NewInt(int64(c.VerificationLevel)),
		c.StreetAndNumber.BigInt(),
		c.Postcode.BigInt(),
		c.Town.BigInt(),
		c.Region.BigInt(),
		c.Country.BigInt(),
		c.Citizenship.BigInt(),
	})
	if err != nil {
		return Hash{}, err
	}

	return HashFromBigInt(hash), nil
}

// KYCVerificationLevel represents the different levels of verification in a KYC (Know Your Customer) process.
type KYCVerificationLevel int

const (
	KYCVerificationLevelNoKYC KYCVerificationLevel = iota
	KYCVerificationLevelPassedKYC
	KYCVerificationLevelQualifiedInvestor
)

// UnmarshalText implements [encoding.TextUnmarshaler].
func (v *KYCVerificationLevel) UnmarshalText(text []byte) error {
	res, err := strconv.Atoi(string(text))
	if err != nil {
		return err
	}

	*v = KYCVerificationLevel(res)
	return nil
}

// MarshalText implements [encoding.TextMarshaler].
func (v KYCVerificationLevel) MarshalText() (text []byte, err error) {
	return []byte(strconv.Itoa(int(v))), nil
}

// IDHash computes and returns a user's ID hash for registration of the HumanID salt hash.
func (c *KYCContent) IDHash() (Hash, error) {
	hash, err := poseidon.Hash([]*big.Int{
		c.Surname.BigInt(),
		c.Forename.BigInt(),
		c.MiddleName.BigInt(),
		big.NewInt(int64(c.YearOfBirth)),
		big.NewInt(int64(c.MonthOfBirth)),
		big.NewInt(int64(c.DayOfBirth)),
		c.Citizenship.BigInt(),
	})
	if err != nil {
		return Hash{}, fmt.Errorf("compute id hash: %w", err)
	}

	return HashFromBigInt(hash), nil
}
