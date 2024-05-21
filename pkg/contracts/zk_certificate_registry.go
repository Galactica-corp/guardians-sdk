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

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
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

// ZkCertificateRegistryMetaData contains all meta data concerning the ZkCertificateRegistry contract.
var ZkCertificateRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"GuardianRegistry_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"treeDepth_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"zkCertificateLeafHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Guardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"zkCertificateAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"zkCertificateLeafHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Guardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"zkCertificateRevocation\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ZERO_VALUE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ZkCertificateToGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_GuardianRegistry\",\"outputs\":[{\"internalType\":\"contractGuardianRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zkCertificateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"addZkCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"hashLeftRight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initBlockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"merkleRootIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRootValidIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextLeafIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zkCertificateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"revokeZkCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeDepth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZkCertificateRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkCertificateRegistryMetaData.ABI instead.
var ZkCertificateRegistryABI = ZkCertificateRegistryMetaData.ABI

// ZkCertificateRegistry is an auto generated Go binding around an Ethereum contract.
type ZkCertificateRegistry struct {
	ZkCertificateRegistryCaller     // Read-only binding to the contract
	ZkCertificateRegistryTransactor // Write-only binding to the contract
	ZkCertificateRegistryFilterer   // Log filterer for contract events
}

// ZkCertificateRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkCertificateRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkCertificateRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkCertificateRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkCertificateRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkCertificateRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkCertificateRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkCertificateRegistrySession struct {
	Contract     *ZkCertificateRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZkCertificateRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkCertificateRegistryCallerSession struct {
	Contract *ZkCertificateRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ZkCertificateRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkCertificateRegistryTransactorSession struct {
	Contract     *ZkCertificateRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ZkCertificateRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkCertificateRegistryRaw struct {
	Contract *ZkCertificateRegistry // Generic contract binding to access the raw methods on
}

// ZkCertificateRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkCertificateRegistryCallerRaw struct {
	Contract *ZkCertificateRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ZkCertificateRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkCertificateRegistryTransactorRaw struct {
	Contract *ZkCertificateRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkCertificateRegistry creates a new instance of ZkCertificateRegistry, bound to a specific deployed contract.
func NewZkCertificateRegistry(address common.Address, backend bind.ContractBackend) (*ZkCertificateRegistry, error) {
	contract, err := bindZkCertificateRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZkCertificateRegistry{ZkCertificateRegistryCaller: ZkCertificateRegistryCaller{contract: contract}, ZkCertificateRegistryTransactor: ZkCertificateRegistryTransactor{contract: contract}, ZkCertificateRegistryFilterer: ZkCertificateRegistryFilterer{contract: contract}}, nil
}

// NewZkCertificateRegistryCaller creates a new read-only instance of ZkCertificateRegistry, bound to a specific deployed contract.
func NewZkCertificateRegistryCaller(address common.Address, caller bind.ContractCaller) (*ZkCertificateRegistryCaller, error) {
	contract, err := bindZkCertificateRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkCertificateRegistryCaller{contract: contract}, nil
}

// NewZkCertificateRegistryTransactor creates a new write-only instance of ZkCertificateRegistry, bound to a specific deployed contract.
func NewZkCertificateRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkCertificateRegistryTransactor, error) {
	contract, err := bindZkCertificateRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkCertificateRegistryTransactor{contract: contract}, nil
}

// NewZkCertificateRegistryFilterer creates a new log filterer instance of ZkCertificateRegistry, bound to a specific deployed contract.
func NewZkCertificateRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkCertificateRegistryFilterer, error) {
	contract, err := bindZkCertificateRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkCertificateRegistryFilterer{contract: contract}, nil
}

// bindZkCertificateRegistry binds a generic wrapper to an already deployed contract.
func bindZkCertificateRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZkCertificateRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkCertificateRegistry *ZkCertificateRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkCertificateRegistry.Contract.ZkCertificateRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkCertificateRegistry *ZkCertificateRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.ZkCertificateRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkCertificateRegistry *ZkCertificateRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.ZkCertificateRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkCertificateRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.contract.Transact(opts, method, params...)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) ZEROVALUE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "ZERO_VALUE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) ZEROVALUE() ([32]byte, error) {
	return _ZkCertificateRegistry.Contract.ZEROVALUE(&_ZkCertificateRegistry.CallOpts)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) ZEROVALUE() ([32]byte, error) {
	return _ZkCertificateRegistry.Contract.ZEROVALUE(&_ZkCertificateRegistry.CallOpts)
}

// ZkCertificateToGuardian is a free data retrieval call binding the contract method 0xe09c8a15.
//
// Solidity: function ZkCertificateToGuardian(bytes32 ) view returns(address)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) ZkCertificateToGuardian(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "ZkCertificateToGuardian", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkCertificateToGuardian is a free data retrieval call binding the contract method 0xe09c8a15.
//
// Solidity: function ZkCertificateToGuardian(bytes32 ) view returns(address)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) ZkCertificateToGuardian(arg0 [32]byte) (common.Address, error) {
	return _ZkCertificateRegistry.Contract.ZkCertificateToGuardian(&_ZkCertificateRegistry.CallOpts, arg0)
}

// ZkCertificateToGuardian is a free data retrieval call binding the contract method 0xe09c8a15.
//
// Solidity: function ZkCertificateToGuardian(bytes32 ) view returns(address)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) ZkCertificateToGuardian(arg0 [32]byte) (common.Address, error) {
	return _ZkCertificateRegistry.Contract.ZkCertificateToGuardian(&_ZkCertificateRegistry.CallOpts, arg0)
}

// GuardianRegistry is a free data retrieval call binding the contract method 0x0a1a9950.
//
// Solidity: function _GuardianRegistry() view returns(address)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) GuardianRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "_GuardianRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GuardianRegistry is a free data retrieval call binding the contract method 0x0a1a9950.
//
// Solidity: function _GuardianRegistry() view returns(address)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) GuardianRegistry() (common.Address, error) {
	return _ZkCertificateRegistry.Contract.GuardianRegistry(&_ZkCertificateRegistry.CallOpts)
}

// GuardianRegistry is a free data retrieval call binding the contract method 0x0a1a9950.
//
// Solidity: function _GuardianRegistry() view returns(address)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) GuardianRegistry() (common.Address, error) {
	return _ZkCertificateRegistry.Contract.GuardianRegistry(&_ZkCertificateRegistry.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) Description() (string, error) {
	return _ZkCertificateRegistry.Contract.Description(&_ZkCertificateRegistry.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) Description() (string, error) {
	return _ZkCertificateRegistry.Contract.Description(&_ZkCertificateRegistry.CallOpts)
}

// GetMerkleRoots is a free data retrieval call binding the contract method 0x85135a1a.
//
// Solidity: function getMerkleRoots() view returns(bytes32[])
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) GetMerkleRoots(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "getMerkleRoots")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetMerkleRoots is a free data retrieval call binding the contract method 0x85135a1a.
//
// Solidity: function getMerkleRoots() view returns(bytes32[])
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) GetMerkleRoots() ([][32]byte, error) {
	return _ZkCertificateRegistry.Contract.GetMerkleRoots(&_ZkCertificateRegistry.CallOpts)
}

// GetMerkleRoots is a free data retrieval call binding the contract method 0x85135a1a.
//
// Solidity: function getMerkleRoots() view returns(bytes32[])
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) GetMerkleRoots() ([][32]byte, error) {
	return _ZkCertificateRegistry.Contract.GetMerkleRoots(&_ZkCertificateRegistry.CallOpts)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x38bf282e.
//
// Solidity: function hashLeftRight(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) HashLeftRight(opts *bind.CallOpts, _left [32]byte, _right [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "hashLeftRight", _left, _right)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashLeftRight is a free data retrieval call binding the contract method 0x38bf282e.
//
// Solidity: function hashLeftRight(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) HashLeftRight(_left [32]byte, _right [32]byte) ([32]byte, error) {
	return _ZkCertificateRegistry.Contract.HashLeftRight(&_ZkCertificateRegistry.CallOpts, _left, _right)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x38bf282e.
//
// Solidity: function hashLeftRight(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) HashLeftRight(_left [32]byte, _right [32]byte) ([32]byte, error) {
	return _ZkCertificateRegistry.Contract.HashLeftRight(&_ZkCertificateRegistry.CallOpts, _left, _right)
}

// InitBlockHeight is a free data retrieval call binding the contract method 0xef4b4119.
//
// Solidity: function initBlockHeight() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) InitBlockHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "initBlockHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitBlockHeight is a free data retrieval call binding the contract method 0xef4b4119.
//
// Solidity: function initBlockHeight() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) InitBlockHeight() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.InitBlockHeight(&_ZkCertificateRegistry.CallOpts)
}

// InitBlockHeight is a free data retrieval call binding the contract method 0xef4b4119.
//
// Solidity: function initBlockHeight() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) InitBlockHeight() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.InitBlockHeight(&_ZkCertificateRegistry.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) MerkleRoot() ([32]byte, error) {
	return _ZkCertificateRegistry.Contract.MerkleRoot(&_ZkCertificateRegistry.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) MerkleRoot() ([32]byte, error) {
	return _ZkCertificateRegistry.Contract.MerkleRoot(&_ZkCertificateRegistry.CallOpts)
}

// MerkleRootIndex is a free data retrieval call binding the contract method 0x390f2b65.
//
// Solidity: function merkleRootIndex(bytes32 ) view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) MerkleRootIndex(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "merkleRootIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MerkleRootIndex is a free data retrieval call binding the contract method 0x390f2b65.
//
// Solidity: function merkleRootIndex(bytes32 ) view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) MerkleRootIndex(arg0 [32]byte) (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.MerkleRootIndex(&_ZkCertificateRegistry.CallOpts, arg0)
}

// MerkleRootIndex is a free data retrieval call binding the contract method 0x390f2b65.
//
// Solidity: function merkleRootIndex(bytes32 ) view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) MerkleRootIndex(arg0 [32]byte) (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.MerkleRootIndex(&_ZkCertificateRegistry.CallOpts, arg0)
}

// MerkleRootValidIndex is a free data retrieval call binding the contract method 0xb0033450.
//
// Solidity: function merkleRootValidIndex() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) MerkleRootValidIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "merkleRootValidIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MerkleRootValidIndex is a free data retrieval call binding the contract method 0xb0033450.
//
// Solidity: function merkleRootValidIndex() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) MerkleRootValidIndex() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.MerkleRootValidIndex(&_ZkCertificateRegistry.CallOpts)
}

// MerkleRootValidIndex is a free data retrieval call binding the contract method 0xb0033450.
//
// Solidity: function merkleRootValidIndex() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) MerkleRootValidIndex() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.MerkleRootValidIndex(&_ZkCertificateRegistry.CallOpts)
}

// MerkleRoots is a free data retrieval call binding the contract method 0x71c5ecb1.
//
// Solidity: function merkleRoots(uint256 ) view returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) MerkleRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "merkleRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoots is a free data retrieval call binding the contract method 0x71c5ecb1.
//
// Solidity: function merkleRoots(uint256 ) view returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) MerkleRoots(arg0 *big.Int) ([32]byte, error) {
	return _ZkCertificateRegistry.Contract.MerkleRoots(&_ZkCertificateRegistry.CallOpts, arg0)
}

// MerkleRoots is a free data retrieval call binding the contract method 0x71c5ecb1.
//
// Solidity: function merkleRoots(uint256 ) view returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) MerkleRoots(arg0 *big.Int) ([32]byte, error) {
	return _ZkCertificateRegistry.Contract.MerkleRoots(&_ZkCertificateRegistry.CallOpts, arg0)
}

// NextLeafIndex is a free data retrieval call binding the contract method 0x0be4f422.
//
// Solidity: function nextLeafIndex() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) NextLeafIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "nextLeafIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextLeafIndex is a free data retrieval call binding the contract method 0x0be4f422.
//
// Solidity: function nextLeafIndex() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) NextLeafIndex() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.NextLeafIndex(&_ZkCertificateRegistry.CallOpts)
}

// NextLeafIndex is a free data retrieval call binding the contract method 0x0be4f422.
//
// Solidity: function nextLeafIndex() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) NextLeafIndex() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.NextLeafIndex(&_ZkCertificateRegistry.CallOpts)
}

// TreeDepth is a free data retrieval call binding the contract method 0x16a56c41.
//
// Solidity: function treeDepth() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) TreeDepth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "treeDepth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreeDepth is a free data retrieval call binding the contract method 0x16a56c41.
//
// Solidity: function treeDepth() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) TreeDepth() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.TreeDepth(&_ZkCertificateRegistry.CallOpts)
}

// TreeDepth is a free data retrieval call binding the contract method 0x16a56c41.
//
// Solidity: function treeDepth() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) TreeDepth() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.TreeDepth(&_ZkCertificateRegistry.CallOpts)
}

// TreeSize is a free data retrieval call binding the contract method 0x8d1ddfb1.
//
// Solidity: function treeSize() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) TreeSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "treeSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreeSize is a free data retrieval call binding the contract method 0x8d1ddfb1.
//
// Solidity: function treeSize() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) TreeSize() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.TreeSize(&_ZkCertificateRegistry.CallOpts)
}

// TreeSize is a free data retrieval call binding the contract method 0x8d1ddfb1.
//
// Solidity: function treeSize() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) TreeSize() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.TreeSize(&_ZkCertificateRegistry.CallOpts)
}

// VerifyMerkleRoot is a free data retrieval call binding the contract method 0x0820dc2f.
//
// Solidity: function verifyMerkleRoot(bytes32 _merkleRoot) view returns(bool)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) VerifyMerkleRoot(opts *bind.CallOpts, _merkleRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "verifyMerkleRoot", _merkleRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleRoot is a free data retrieval call binding the contract method 0x0820dc2f.
//
// Solidity: function verifyMerkleRoot(bytes32 _merkleRoot) view returns(bool)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) VerifyMerkleRoot(_merkleRoot [32]byte) (bool, error) {
	return _ZkCertificateRegistry.Contract.VerifyMerkleRoot(&_ZkCertificateRegistry.CallOpts, _merkleRoot)
}

// VerifyMerkleRoot is a free data retrieval call binding the contract method 0x0820dc2f.
//
// Solidity: function verifyMerkleRoot(bytes32 _merkleRoot) view returns(bool)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) VerifyMerkleRoot(_merkleRoot [32]byte) (bool, error) {
	return _ZkCertificateRegistry.Contract.VerifyMerkleRoot(&_ZkCertificateRegistry.CallOpts, _merkleRoot)
}

// AddZkCertificate is a paid mutator transaction binding the contract method 0x89a64453.
//
// Solidity: function addZkCertificate(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactor) AddZkCertificate(opts *bind.TransactOpts, leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkCertificateRegistry.contract.Transact(opts, "addZkCertificate", leafIndex, zkCertificateHash, merkleProof)
}

// AddZkCertificate is a paid mutator transaction binding the contract method 0x89a64453.
//
// Solidity: function addZkCertificate(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) AddZkCertificate(leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.AddZkCertificate(&_ZkCertificateRegistry.TransactOpts, leafIndex, zkCertificateHash, merkleProof)
}

// AddZkCertificate is a paid mutator transaction binding the contract method 0x89a64453.
//
// Solidity: function addZkCertificate(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactorSession) AddZkCertificate(leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.AddZkCertificate(&_ZkCertificateRegistry.TransactOpts, leafIndex, zkCertificateHash, merkleProof)
}

// RevokeZkCertificate is a paid mutator transaction binding the contract method 0xa2873348.
//
// Solidity: function revokeZkCertificate(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactor) RevokeZkCertificate(opts *bind.TransactOpts, leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkCertificateRegistry.contract.Transact(opts, "revokeZkCertificate", leafIndex, zkCertificateHash, merkleProof)
}

// RevokeZkCertificate is a paid mutator transaction binding the contract method 0xa2873348.
//
// Solidity: function revokeZkCertificate(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) RevokeZkCertificate(leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.RevokeZkCertificate(&_ZkCertificateRegistry.TransactOpts, leafIndex, zkCertificateHash, merkleProof)
}

// RevokeZkCertificate is a paid mutator transaction binding the contract method 0xa2873348.
//
// Solidity: function revokeZkCertificate(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactorSession) RevokeZkCertificate(leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.RevokeZkCertificate(&_ZkCertificateRegistry.TransactOpts, leafIndex, zkCertificateHash, merkleProof)
}

// ZkCertificateRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZkCertificateRegistry contract.
type ZkCertificateRegistryInitializedIterator struct {
	Event *ZkCertificateRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ZkCertificateRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkCertificateRegistryInitialized)
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
		it.Event = new(ZkCertificateRegistryInitialized)
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
func (it *ZkCertificateRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkCertificateRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkCertificateRegistryInitialized represents a Initialized event raised by the ZkCertificateRegistry contract.
type ZkCertificateRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZkCertificateRegistryInitializedIterator, error) {

	logs, sub, err := _ZkCertificateRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZkCertificateRegistryInitializedIterator{contract: _ZkCertificateRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZkCertificateRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ZkCertificateRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkCertificateRegistryInitialized)
				if err := _ZkCertificateRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) ParseInitialized(log types.Log) (*ZkCertificateRegistryInitialized, error) {
	event := new(ZkCertificateRegistryInitialized)
	if err := _ZkCertificateRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkCertificateRegistryZkCertificateAdditionIterator is returned from FilterZkCertificateAddition and is used to iterate over the raw logs and unpacked data for ZkCertificateAddition events raised by the ZkCertificateRegistry contract.
type ZkCertificateRegistryZkCertificateAdditionIterator struct {
	Event *ZkCertificateRegistryZkCertificateAddition // Event containing the contract specifics and raw log

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
func (it *ZkCertificateRegistryZkCertificateAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkCertificateRegistryZkCertificateAddition)
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
		it.Event = new(ZkCertificateRegistryZkCertificateAddition)
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
func (it *ZkCertificateRegistryZkCertificateAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkCertificateRegistryZkCertificateAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkCertificateRegistryZkCertificateAddition represents a ZkCertificateAddition event raised by the ZkCertificateRegistry contract.
type ZkCertificateRegistryZkCertificateAddition struct {
	ZkCertificateLeafHash [32]byte
	Guardian              common.Address
	Index                 *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterZkCertificateAddition is a free log retrieval operation binding the contract event 0xd8ed161990bf70dd13139e68384461b9fffbd1d3433c3bd56a9ccd22e0118cb6.
//
// Solidity: event zkCertificateAddition(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint256 index)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) FilterZkCertificateAddition(opts *bind.FilterOpts, zkCertificateLeafHash [][32]byte, Guardian []common.Address) (*ZkCertificateRegistryZkCertificateAdditionIterator, error) {

	var zkCertificateLeafHashRule []interface{}
	for _, zkCertificateLeafHashItem := range zkCertificateLeafHash {
		zkCertificateLeafHashRule = append(zkCertificateLeafHashRule, zkCertificateLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _ZkCertificateRegistry.contract.FilterLogs(opts, "zkCertificateAddition", zkCertificateLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return &ZkCertificateRegistryZkCertificateAdditionIterator{contract: _ZkCertificateRegistry.contract, event: "zkCertificateAddition", logs: logs, sub: sub}, nil
}

// WatchZkCertificateAddition is a free log subscription operation binding the contract event 0xd8ed161990bf70dd13139e68384461b9fffbd1d3433c3bd56a9ccd22e0118cb6.
//
// Solidity: event zkCertificateAddition(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint256 index)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) WatchZkCertificateAddition(opts *bind.WatchOpts, sink chan<- *ZkCertificateRegistryZkCertificateAddition, zkCertificateLeafHash [][32]byte, Guardian []common.Address) (event.Subscription, error) {

	var zkCertificateLeafHashRule []interface{}
	for _, zkCertificateLeafHashItem := range zkCertificateLeafHash {
		zkCertificateLeafHashRule = append(zkCertificateLeafHashRule, zkCertificateLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _ZkCertificateRegistry.contract.WatchLogs(opts, "zkCertificateAddition", zkCertificateLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkCertificateRegistryZkCertificateAddition)
				if err := _ZkCertificateRegistry.contract.UnpackLog(event, "zkCertificateAddition", log); err != nil {
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

// ParseZkCertificateAddition is a log parse operation binding the contract event 0xd8ed161990bf70dd13139e68384461b9fffbd1d3433c3bd56a9ccd22e0118cb6.
//
// Solidity: event zkCertificateAddition(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint256 index)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) ParseZkCertificateAddition(log types.Log) (*ZkCertificateRegistryZkCertificateAddition, error) {
	event := new(ZkCertificateRegistryZkCertificateAddition)
	if err := _ZkCertificateRegistry.contract.UnpackLog(event, "zkCertificateAddition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkCertificateRegistryZkCertificateRevocationIterator is returned from FilterZkCertificateRevocation and is used to iterate over the raw logs and unpacked data for ZkCertificateRevocation events raised by the ZkCertificateRegistry contract.
type ZkCertificateRegistryZkCertificateRevocationIterator struct {
	Event *ZkCertificateRegistryZkCertificateRevocation // Event containing the contract specifics and raw log

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
func (it *ZkCertificateRegistryZkCertificateRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkCertificateRegistryZkCertificateRevocation)
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
		it.Event = new(ZkCertificateRegistryZkCertificateRevocation)
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
func (it *ZkCertificateRegistryZkCertificateRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkCertificateRegistryZkCertificateRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkCertificateRegistryZkCertificateRevocation represents a ZkCertificateRevocation event raised by the ZkCertificateRegistry contract.
type ZkCertificateRegistryZkCertificateRevocation struct {
	ZkCertificateLeafHash [32]byte
	Guardian              common.Address
	Index                 *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterZkCertificateRevocation is a free log retrieval operation binding the contract event 0x2d6db1db44b3ed99e9b90701d664345ef76dd37336779182d6251abf294840a6.
//
// Solidity: event zkCertificateRevocation(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint256 index)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) FilterZkCertificateRevocation(opts *bind.FilterOpts, zkCertificateLeafHash [][32]byte, Guardian []common.Address) (*ZkCertificateRegistryZkCertificateRevocationIterator, error) {

	var zkCertificateLeafHashRule []interface{}
	for _, zkCertificateLeafHashItem := range zkCertificateLeafHash {
		zkCertificateLeafHashRule = append(zkCertificateLeafHashRule, zkCertificateLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _ZkCertificateRegistry.contract.FilterLogs(opts, "zkCertificateRevocation", zkCertificateLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return &ZkCertificateRegistryZkCertificateRevocationIterator{contract: _ZkCertificateRegistry.contract, event: "zkCertificateRevocation", logs: logs, sub: sub}, nil
}

// WatchZkCertificateRevocation is a free log subscription operation binding the contract event 0x2d6db1db44b3ed99e9b90701d664345ef76dd37336779182d6251abf294840a6.
//
// Solidity: event zkCertificateRevocation(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint256 index)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) WatchZkCertificateRevocation(opts *bind.WatchOpts, sink chan<- *ZkCertificateRegistryZkCertificateRevocation, zkCertificateLeafHash [][32]byte, Guardian []common.Address) (event.Subscription, error) {

	var zkCertificateLeafHashRule []interface{}
	for _, zkCertificateLeafHashItem := range zkCertificateLeafHash {
		zkCertificateLeafHashRule = append(zkCertificateLeafHashRule, zkCertificateLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _ZkCertificateRegistry.contract.WatchLogs(opts, "zkCertificateRevocation", zkCertificateLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkCertificateRegistryZkCertificateRevocation)
				if err := _ZkCertificateRegistry.contract.UnpackLog(event, "zkCertificateRevocation", log); err != nil {
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

// ParseZkCertificateRevocation is a log parse operation binding the contract event 0x2d6db1db44b3ed99e9b90701d664345ef76dd37336779182d6251abf294840a6.
//
// Solidity: event zkCertificateRevocation(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint256 index)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) ParseZkCertificateRevocation(log types.Log) (*ZkCertificateRegistryZkCertificateRevocation, error) {
	event := new(ZkCertificateRegistryZkCertificateRevocation)
	if err := _ZkCertificateRegistry.contract.UnpackLog(event, "zkCertificateRevocation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
