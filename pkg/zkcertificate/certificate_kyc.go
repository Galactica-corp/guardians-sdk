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

package zkcertificate

import (
	"cmp"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"

	"github.com/iden3/go-iden3-crypto/v2/poseidon"

	"github.com/galactica-corp/guardians-sdk/internal/validation"
	"github.com/galactica-corp/guardians-sdk/pkg/hashing"
)

// KYCContent represents the data for Know Your Customer (KYC) verification.
// It contains various fields required for identity verification and validation.
type KYCContent struct {
	Surname           string               `json:"surname" validate:"required"`
	Forename          string               `json:"forename" validate:"required"`
	MiddleName        string               `json:"middlename"`
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

func (k KYCContent) Validate() error {
	return validation.Validate.Struct(k)
}

// UnmarshalJSON implements [json.Unmarshaler].
func (k *KYCContent) UnmarshalJSON(data []byte) error {
	type Alias KYCContent

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	if err := (*KYCContent)(&alias).Validate(); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	*k = KYCContent(alias)
	return nil
}

// Standard returns the standard associated with the KYCContent, which is StandardKYC.
func (k KYCContent) Standard() Standard {
	return StandardKYC
}

// Hash implements Content.
func (k KYCContent) Hash() (Hash, error) {
	surnameHash, err := hashing.HashBytes([]byte(k.Surname))
	if err != nil {
		return Hash{}, fmt.Errorf("hash surname: %w", err)
	}

	forenameHash, err := hashing.HashBytes([]byte(k.Forename))
	if err != nil {
		return Hash{}, fmt.Errorf("hash forename: %w", err)
	}

	middleNameHash, err := hashing.HashBytes([]byte(k.MiddleName))
	if err != nil {
		return Hash{}, fmt.Errorf("hash middle name: %w", err)
	}

	streetAndNumberHash, err := hashing.HashBytes([]byte(k.StreetAndNumber))
	if err != nil {
		return Hash{}, fmt.Errorf("hash street and number: %w", err)
	}

	postcodeHash, err := hashing.HashBytes([]byte(k.Postcode))
	if err != nil {
		return Hash{}, fmt.Errorf("hash postcode: %w", err)
	}

	townHash, err := hashing.HashBytes([]byte(k.Town))
	if err != nil {
		return Hash{}, fmt.Errorf("hash town: %w", err)
	}

	regionHash, err := hashing.HashBytes([]byte(k.Region))
	if err != nil {
		return Hash{}, fmt.Errorf("hash region: %w", err)
	}

	countryHash, err := hashing.HashBytes([]byte(k.Country))
	if err != nil {
		return Hash{}, fmt.Errorf("hash country: %w", err)
	}

	citizenshipHash, err := hashing.HashBytes([]byte(k.Citizenship))
	if err != nil {
		return Hash{}, fmt.Errorf("hash citizenship: %w", err)
	}

	hash, err := poseidon.Hash([]*big.Int{
		// fields ordered alphabetically regarding their JSON key
		citizenshipHash,
		countryHash,
		big.NewInt(int64(k.DayOfBirth)),
		forenameHash,
		cmp.Or(middleNameHash, EmptySequenceHash),
		big.NewInt(int64(k.MonthOfBirth)),
		cmp.Or(postcodeHash, EmptySequenceHash),
		cmp.Or(regionHash, EmptySequenceHash),
		cmp.Or(streetAndNumberHash, EmptySequenceHash),
		surnameHash,
		cmp.Or(townHash, EmptySequenceHash),
		big.NewInt(int64(k.VerificationLevel)),
		big.NewInt(int64(k.YearOfBirth)),
	})
	if err != nil {
		return Hash{}, err
	}

	return HashFromBigInt(hash), nil
}

// IDHash computes and returns a user's ID hash for registration of the HumanID salt hash.
func (k KYCContent) IDHash() (Hash, error) {
	citizenshipHash, err := hashing.HashBytes([]byte(k.Citizenship))
	if err != nil {
		return Hash{}, fmt.Errorf("hash citizenship: %w", err)
	}

	surnameHash, err := hashing.HashBytes([]byte(k.Surname))
	if err != nil {
		return Hash{}, fmt.Errorf("hash surname: %w", err)
	}

	forenameHash, err := hashing.HashBytes([]byte(k.Forename))
	if err != nil {
		return Hash{}, fmt.Errorf("hash forename: %w", err)
	}

	middleNameHash, err := hashing.HashBytes([]byte(k.MiddleName))
	if err != nil {
		return Hash{}, fmt.Errorf("hash middle name: %w", err)
	}

	hash, err := poseidon.Hash([]*big.Int{
		// fields ordered alphabetically regarding their JSON key
		citizenshipHash,
		big.NewInt(int64(k.DayOfBirth)),
		forenameHash,
		cmp.Or(middleNameHash, EmptySequenceHash),
		big.NewInt(int64(k.MonthOfBirth)),
		surnameHash,
		big.NewInt(int64(k.YearOfBirth)),
	})
	if err != nil {
		return Hash{}, fmt.Errorf("compute id hash: %w", err)
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
