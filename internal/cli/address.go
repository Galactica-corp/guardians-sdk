/*
 Copyright © 2023 Galactica Network

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.

 You should have received a copy of the GNU General Public License
 along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package cli

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Address common.Address

func (a Address) Address() common.Address {
	return common.Address(a)
}

func (a Address) String() string {
	return a.Address().String()
}

func (a *Address) Set(s string) error {
	if !common.IsHexAddress(s) {
		return fmt.Errorf("invalid address")
	}

	*a = Address(common.HexToAddress(s))
	return nil
}

func (a Address) Type() string {
	return "address"
}
