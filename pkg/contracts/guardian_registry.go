// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// GuardianRegistryMetaData contains all meta data concerning the GuardianRegistry contract.
var GuardianRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"pubKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"grantGuardianRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"guardians\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pubKeyToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGuardianRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"revokeGuardianRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GuardianRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use GuardianRegistryMetaData.ABI instead.
var GuardianRegistryABI = GuardianRegistryMetaData.ABI

// GuardianRegistry is an auto generated Go binding around an Ethereum contract.
type GuardianRegistry struct {
	GuardianRegistryCaller     // Read-only binding to the contract
	GuardianRegistryTransactor // Write-only binding to the contract
	GuardianRegistryFilterer   // Log filterer for contract events
}

// GuardianRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GuardianRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardianRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GuardianRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardianRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GuardianRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardianRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GuardianRegistrySession struct {
	Contract     *GuardianRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GuardianRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GuardianRegistryCallerSession struct {
	Contract *GuardianRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// GuardianRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GuardianRegistryTransactorSession struct {
	Contract     *GuardianRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// GuardianRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GuardianRegistryRaw struct {
	Contract *GuardianRegistry // Generic contract binding to access the raw methods on
}

// GuardianRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GuardianRegistryCallerRaw struct {
	Contract *GuardianRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// GuardianRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GuardianRegistryTransactorRaw struct {
	Contract *GuardianRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGuardianRegistry creates a new instance of GuardianRegistry, bound to a specific deployed contract.
func NewGuardianRegistry(address common.Address, backend bind.ContractBackend) (*GuardianRegistry, error) {
	contract, err := bindGuardianRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GuardianRegistry{GuardianRegistryCaller: GuardianRegistryCaller{contract: contract}, GuardianRegistryTransactor: GuardianRegistryTransactor{contract: contract}, GuardianRegistryFilterer: GuardianRegistryFilterer{contract: contract}}, nil
}

// NewGuardianRegistryCaller creates a new read-only instance of GuardianRegistry, bound to a specific deployed contract.
func NewGuardianRegistryCaller(address common.Address, caller bind.ContractCaller) (*GuardianRegistryCaller, error) {
	contract, err := bindGuardianRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GuardianRegistryCaller{contract: contract}, nil
}

// NewGuardianRegistryTransactor creates a new write-only instance of GuardianRegistry, bound to a specific deployed contract.
func NewGuardianRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*GuardianRegistryTransactor, error) {
	contract, err := bindGuardianRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GuardianRegistryTransactor{contract: contract}, nil
}

// NewGuardianRegistryFilterer creates a new log filterer instance of GuardianRegistry, bound to a specific deployed contract.
func NewGuardianRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*GuardianRegistryFilterer, error) {
	contract, err := bindGuardianRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GuardianRegistryFilterer{contract: contract}, nil
}

// bindGuardianRegistry binds a generic wrapper to an already deployed contract.
func bindGuardianRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GuardianRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuardianRegistry *GuardianRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuardianRegistry.Contract.GuardianRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuardianRegistry *GuardianRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuardianRegistry.Contract.GuardianRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuardianRegistry *GuardianRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuardianRegistry.Contract.GuardianRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuardianRegistry *GuardianRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuardianRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuardianRegistry *GuardianRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuardianRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuardianRegistry *GuardianRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuardianRegistry.Contract.contract.Transact(opts, method, params...)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address ) view returns(bool whitelisted, string name)
func (_GuardianRegistry *GuardianRegistryCaller) Guardians(opts *bind.CallOpts, arg0 common.Address) (struct {
	Whitelisted bool
	Name        string
}, error) {
	var out []interface{}
	err := _GuardianRegistry.contract.Call(opts, &out, "guardians", arg0)

	outstruct := new(struct {
		Whitelisted bool
		Name        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Whitelisted = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address ) view returns(bool whitelisted, string name)
func (_GuardianRegistry *GuardianRegistrySession) Guardians(arg0 common.Address) (struct {
	Whitelisted bool
	Name        string
}, error) {
	return _GuardianRegistry.Contract.Guardians(&_GuardianRegistry.CallOpts, arg0)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address ) view returns(bool whitelisted, string name)
func (_GuardianRegistry *GuardianRegistryCallerSession) Guardians(arg0 common.Address) (struct {
	Whitelisted bool
	Name        string
}, error) {
	return _GuardianRegistry.Contract.Guardians(&_GuardianRegistry.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address guardian) view returns(bool)
func (_GuardianRegistry *GuardianRegistryCaller) IsWhitelisted(opts *bind.CallOpts, guardian common.Address) (bool, error) {
	var out []interface{}
	err := _GuardianRegistry.contract.Call(opts, &out, "isWhitelisted", guardian)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address guardian) view returns(bool)
func (_GuardianRegistry *GuardianRegistrySession) IsWhitelisted(guardian common.Address) (bool, error) {
	return _GuardianRegistry.Contract.IsWhitelisted(&_GuardianRegistry.CallOpts, guardian)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address guardian) view returns(bool)
func (_GuardianRegistry *GuardianRegistryCallerSession) IsWhitelisted(guardian common.Address) (bool, error) {
	return _GuardianRegistry.Contract.IsWhitelisted(&_GuardianRegistry.CallOpts, guardian)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_GuardianRegistry *GuardianRegistryCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GuardianRegistry.contract.Call(opts, &out, "newOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_GuardianRegistry *GuardianRegistrySession) NewOwner() (common.Address, error) {
	return _GuardianRegistry.Contract.NewOwner(&_GuardianRegistry.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_GuardianRegistry *GuardianRegistryCallerSession) NewOwner() (common.Address, error) {
	return _GuardianRegistry.Contract.NewOwner(&_GuardianRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GuardianRegistry *GuardianRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GuardianRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GuardianRegistry *GuardianRegistrySession) Owner() (common.Address, error) {
	return _GuardianRegistry.Contract.Owner(&_GuardianRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GuardianRegistry *GuardianRegistryCallerSession) Owner() (common.Address, error) {
	return _GuardianRegistry.Contract.Owner(&_GuardianRegistry.CallOpts)
}

// PubKeyToAddress is a free data retrieval call binding the contract method 0x7b9bff31.
//
// Solidity: function pubKeyToAddress(uint256 , uint256 ) view returns(address)
func (_GuardianRegistry *GuardianRegistryCaller) PubKeyToAddress(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GuardianRegistry.contract.Call(opts, &out, "pubKeyToAddress", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubKeyToAddress is a free data retrieval call binding the contract method 0x7b9bff31.
//
// Solidity: function pubKeyToAddress(uint256 , uint256 ) view returns(address)
func (_GuardianRegistry *GuardianRegistrySession) PubKeyToAddress(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _GuardianRegistry.Contract.PubKeyToAddress(&_GuardianRegistry.CallOpts, arg0, arg1)
}

// PubKeyToAddress is a free data retrieval call binding the contract method 0x7b9bff31.
//
// Solidity: function pubKeyToAddress(uint256 , uint256 ) view returns(address)
func (_GuardianRegistry *GuardianRegistryCallerSession) PubKeyToAddress(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _GuardianRegistry.Contract.PubKeyToAddress(&_GuardianRegistry.CallOpts, arg0, arg1)
}

// GrantGuardianRole is a paid mutator transaction binding the contract method 0x06b7e1e8.
//
// Solidity: function grantGuardianRole(address guardian, uint256[2] pubKey, string name) returns()
func (_GuardianRegistry *GuardianRegistryTransactor) GrantGuardianRole(opts *bind.TransactOpts, guardian common.Address, pubKey [2]*big.Int, name string) (*types.Transaction, error) {
	return _GuardianRegistry.contract.Transact(opts, "grantGuardianRole", guardian, pubKey, name)
}

// GrantGuardianRole is a paid mutator transaction binding the contract method 0x06b7e1e8.
//
// Solidity: function grantGuardianRole(address guardian, uint256[2] pubKey, string name) returns()
func (_GuardianRegistry *GuardianRegistrySession) GrantGuardianRole(guardian common.Address, pubKey [2]*big.Int, name string) (*types.Transaction, error) {
	return _GuardianRegistry.Contract.GrantGuardianRole(&_GuardianRegistry.TransactOpts, guardian, pubKey, name)
}

// GrantGuardianRole is a paid mutator transaction binding the contract method 0x06b7e1e8.
//
// Solidity: function grantGuardianRole(address guardian, uint256[2] pubKey, string name) returns()
func (_GuardianRegistry *GuardianRegistryTransactorSession) GrantGuardianRole(guardian common.Address, pubKey [2]*big.Int, name string) (*types.Transaction, error) {
	return _GuardianRegistry.Contract.GrantGuardianRole(&_GuardianRegistry.TransactOpts, guardian, pubKey, name)
}

// RenounceGuardianRole is a paid mutator transaction binding the contract method 0x34d03190.
//
// Solidity: function renounceGuardianRole() returns()
func (_GuardianRegistry *GuardianRegistryTransactor) RenounceGuardianRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuardianRegistry.contract.Transact(opts, "renounceGuardianRole")
}

// RenounceGuardianRole is a paid mutator transaction binding the contract method 0x34d03190.
//
// Solidity: function renounceGuardianRole() returns()
func (_GuardianRegistry *GuardianRegistrySession) RenounceGuardianRole() (*types.Transaction, error) {
	return _GuardianRegistry.Contract.RenounceGuardianRole(&_GuardianRegistry.TransactOpts)
}

// RenounceGuardianRole is a paid mutator transaction binding the contract method 0x34d03190.
//
// Solidity: function renounceGuardianRole() returns()
func (_GuardianRegistry *GuardianRegistryTransactorSession) RenounceGuardianRole() (*types.Transaction, error) {
	return _GuardianRegistry.Contract.RenounceGuardianRole(&_GuardianRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GuardianRegistry *GuardianRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuardianRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GuardianRegistry *GuardianRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _GuardianRegistry.Contract.RenounceOwnership(&_GuardianRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GuardianRegistry *GuardianRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GuardianRegistry.Contract.RenounceOwnership(&_GuardianRegistry.TransactOpts)
}

// RevokeGuardianRole is a paid mutator transaction binding the contract method 0x4b4b5389.
//
// Solidity: function revokeGuardianRole(address guardian) returns()
func (_GuardianRegistry *GuardianRegistryTransactor) RevokeGuardianRole(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _GuardianRegistry.contract.Transact(opts, "revokeGuardianRole", guardian)
}

// RevokeGuardianRole is a paid mutator transaction binding the contract method 0x4b4b5389.
//
// Solidity: function revokeGuardianRole(address guardian) returns()
func (_GuardianRegistry *GuardianRegistrySession) RevokeGuardianRole(guardian common.Address) (*types.Transaction, error) {
	return _GuardianRegistry.Contract.RevokeGuardianRole(&_GuardianRegistry.TransactOpts, guardian)
}

// RevokeGuardianRole is a paid mutator transaction binding the contract method 0x4b4b5389.
//
// Solidity: function revokeGuardianRole(address guardian) returns()
func (_GuardianRegistry *GuardianRegistryTransactorSession) RevokeGuardianRole(guardian common.Address) (*types.Transaction, error) {
	return _GuardianRegistry.Contract.RevokeGuardianRole(&_GuardianRegistry.TransactOpts, guardian)
}

// SetNewOwner is a paid mutator transaction binding the contract method 0xf5a1f5b4.
//
// Solidity: function setNewOwner(address _newOwner) returns()
func (_GuardianRegistry *GuardianRegistryTransactor) SetNewOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _GuardianRegistry.contract.Transact(opts, "setNewOwner", _newOwner)
}

// SetNewOwner is a paid mutator transaction binding the contract method 0xf5a1f5b4.
//
// Solidity: function setNewOwner(address _newOwner) returns()
func (_GuardianRegistry *GuardianRegistrySession) SetNewOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _GuardianRegistry.Contract.SetNewOwner(&_GuardianRegistry.TransactOpts, _newOwner)
}

// SetNewOwner is a paid mutator transaction binding the contract method 0xf5a1f5b4.
//
// Solidity: function setNewOwner(address _newOwner) returns()
func (_GuardianRegistry *GuardianRegistryTransactorSession) SetNewOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _GuardianRegistry.Contract.SetNewOwner(&_GuardianRegistry.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x880ad0af.
//
// Solidity: function transferOwnership() returns()
func (_GuardianRegistry *GuardianRegistryTransactor) TransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuardianRegistry.contract.Transact(opts, "transferOwnership")
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x880ad0af.
//
// Solidity: function transferOwnership() returns()
func (_GuardianRegistry *GuardianRegistrySession) TransferOwnership() (*types.Transaction, error) {
	return _GuardianRegistry.Contract.TransferOwnership(&_GuardianRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x880ad0af.
//
// Solidity: function transferOwnership() returns()
func (_GuardianRegistry *GuardianRegistryTransactorSession) TransferOwnership() (*types.Transaction, error) {
	return _GuardianRegistry.Contract.TransferOwnership(&_GuardianRegistry.TransactOpts)
}

// GuardianRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GuardianRegistry contract.
type GuardianRegistryOwnershipTransferredIterator struct {
	Event *GuardianRegistryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GuardianRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardianRegistryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GuardianRegistryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GuardianRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardianRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardianRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the GuardianRegistry contract.
type GuardianRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GuardianRegistry *GuardianRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GuardianRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GuardianRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GuardianRegistryOwnershipTransferredIterator{contract: _GuardianRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GuardianRegistry *GuardianRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GuardianRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GuardianRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardianRegistryOwnershipTransferred)
				if err := _GuardianRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GuardianRegistry *GuardianRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*GuardianRegistryOwnershipTransferred, error) {
	event := new(GuardianRegistryOwnershipTransferred)
	if err := _GuardianRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
