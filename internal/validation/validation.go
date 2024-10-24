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

package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

var Validate = validator.New(validator.WithRequiredStructEnabled())

// init function to register custom validators when the package is imported
func init() {
	err := Validate.RegisterValidation("decimal_gt_0", customDecimalGreaterThanZero)
	if err != nil {
		fmt.Println("Error registering validation:", err)
	}
}

// customDecimalGreaterThanZero is a custom validation function to check if decimal is greater than zero
func customDecimalGreaterThanZero(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(decimal.Decimal)
	if !ok {
		return false
	}

	return value.GreaterThan(decimal.Zero)
}
