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
