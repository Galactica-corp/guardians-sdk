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

// ZkKYCRegistryMetaData contains all meta data concerning the ZkKYCRegistry contract.
var ZkKYCRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"GuardianRegistry_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"treeDepth_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"zkCertificateLeafHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Guardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"zkCertificateAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"zkCertificateLeafHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Guardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"zkCertificateRevocation\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ZERO_VALUE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ZkCertificateHashToCommitedGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ZkCertificateHashToIndexInQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ZkCertificateHashToQueueTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ZkCertificateQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ZkCertificateToGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_GuardianRegistry\",\"outputs\":[{\"internalType\":\"contractGuardianRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"name\":\"addZkCertificate\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zkCertificateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"idHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saltHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"addZkKYC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"changeQueueExpirationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"zkCertificateHash\",\"type\":\"bytes32\"}],\"name\":\"checkZkCertificateHashInQueue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentQueuePointer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"zkCertificateHash\",\"type\":\"bytes32\"}],\"name\":\"getTimeParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"hashLeftRight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"humanIDSaltRegistry\",\"outputs\":[{\"internalType\":\"contractHumanIDSaltRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initBlockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"merkleRootIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRootValidIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextLeafIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueExpirationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"zkCertificateHash\",\"type\":\"bytes32\"}],\"name\":\"registerToQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zkCertificateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"revokeZkCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeDepth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZkKYCRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkKYCRegistryMetaData.ABI instead.
var ZkKYCRegistryABI = ZkKYCRegistryMetaData.ABI

// ZkKYCRegistry is an auto generated Go binding around an Ethereum contract.
type ZkKYCRegistry struct {
	ZkKYCRegistryCaller     // Read-only binding to the contract
	ZkKYCRegistryTransactor // Write-only binding to the contract
	ZkKYCRegistryFilterer   // Log filterer for contract events
}

// ZkKYCRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkKYCRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkKYCRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkKYCRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkKYCRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkKYCRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkKYCRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkKYCRegistrySession struct {
	Contract     *ZkKYCRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkKYCRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkKYCRegistryCallerSession struct {
	Contract *ZkKYCRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ZkKYCRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkKYCRegistryTransactorSession struct {
	Contract     *ZkKYCRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ZkKYCRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkKYCRegistryRaw struct {
	Contract *ZkKYCRegistry // Generic contract binding to access the raw methods on
}

// ZkKYCRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkKYCRegistryCallerRaw struct {
	Contract *ZkKYCRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ZkKYCRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkKYCRegistryTransactorRaw struct {
	Contract *ZkKYCRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkKYCRegistry creates a new instance of ZkKYCRegistry, bound to a specific deployed contract.
func NewZkKYCRegistry(address common.Address, backend bind.ContractBackend) (*ZkKYCRegistry, error) {
	contract, err := bindZkKYCRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZkKYCRegistry{ZkKYCRegistryCaller: ZkKYCRegistryCaller{contract: contract}, ZkKYCRegistryTransactor: ZkKYCRegistryTransactor{contract: contract}, ZkKYCRegistryFilterer: ZkKYCRegistryFilterer{contract: contract}}, nil
}

// NewZkKYCRegistryCaller creates a new read-only instance of ZkKYCRegistry, bound to a specific deployed contract.
func NewZkKYCRegistryCaller(address common.Address, caller bind.ContractCaller) (*ZkKYCRegistryCaller, error) {
	contract, err := bindZkKYCRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkKYCRegistryCaller{contract: contract}, nil
}

// NewZkKYCRegistryTransactor creates a new write-only instance of ZkKYCRegistry, bound to a specific deployed contract.
func NewZkKYCRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkKYCRegistryTransactor, error) {
	contract, err := bindZkKYCRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkKYCRegistryTransactor{contract: contract}, nil
}

// NewZkKYCRegistryFilterer creates a new log filterer instance of ZkKYCRegistry, bound to a specific deployed contract.
func NewZkKYCRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkKYCRegistryFilterer, error) {
	contract, err := bindZkKYCRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkKYCRegistryFilterer{contract: contract}, nil
}

// bindZkKYCRegistry binds a generic wrapper to an already deployed contract.
func bindZkKYCRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZkKYCRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkKYCRegistry *ZkKYCRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkKYCRegistry.Contract.ZkKYCRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkKYCRegistry *ZkKYCRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.ZkKYCRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkKYCRegistry *ZkKYCRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.ZkKYCRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkKYCRegistry *ZkKYCRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkKYCRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkKYCRegistry *ZkKYCRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkKYCRegistry *ZkKYCRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.contract.Transact(opts, method, params...)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) ZEROVALUE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "ZERO_VALUE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistrySession) ZEROVALUE() ([32]byte, error) {
	return _ZkKYCRegistry.Contract.ZEROVALUE(&_ZkKYCRegistry.CallOpts)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) ZEROVALUE() ([32]byte, error) {
	return _ZkKYCRegistry.Contract.ZEROVALUE(&_ZkKYCRegistry.CallOpts)
}

// ZkCertificateHashToCommitedGuardian is a free data retrieval call binding the contract method 0xc36600fe.
//
// Solidity: function ZkCertificateHashToCommitedGuardian(bytes32 ) view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) ZkCertificateHashToCommitedGuardian(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "ZkCertificateHashToCommitedGuardian", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkCertificateHashToCommitedGuardian is a free data retrieval call binding the contract method 0xc36600fe.
//
// Solidity: function ZkCertificateHashToCommitedGuardian(bytes32 ) view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistrySession) ZkCertificateHashToCommitedGuardian(arg0 [32]byte) (common.Address, error) {
	return _ZkKYCRegistry.Contract.ZkCertificateHashToCommitedGuardian(&_ZkKYCRegistry.CallOpts, arg0)
}

// ZkCertificateHashToCommitedGuardian is a free data retrieval call binding the contract method 0xc36600fe.
//
// Solidity: function ZkCertificateHashToCommitedGuardian(bytes32 ) view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) ZkCertificateHashToCommitedGuardian(arg0 [32]byte) (common.Address, error) {
	return _ZkKYCRegistry.Contract.ZkCertificateHashToCommitedGuardian(&_ZkKYCRegistry.CallOpts, arg0)
}

// ZkCertificateHashToIndexInQueue is a free data retrieval call binding the contract method 0xccdc65d8.
//
// Solidity: function ZkCertificateHashToIndexInQueue(bytes32 ) view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) ZkCertificateHashToIndexInQueue(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "ZkCertificateHashToIndexInQueue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZkCertificateHashToIndexInQueue is a free data retrieval call binding the contract method 0xccdc65d8.
//
// Solidity: function ZkCertificateHashToIndexInQueue(bytes32 ) view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistrySession) ZkCertificateHashToIndexInQueue(arg0 [32]byte) (*big.Int, error) {
	return _ZkKYCRegistry.Contract.ZkCertificateHashToIndexInQueue(&_ZkKYCRegistry.CallOpts, arg0)
}

// ZkCertificateHashToIndexInQueue is a free data retrieval call binding the contract method 0xccdc65d8.
//
// Solidity: function ZkCertificateHashToIndexInQueue(bytes32 ) view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) ZkCertificateHashToIndexInQueue(arg0 [32]byte) (*big.Int, error) {
	return _ZkKYCRegistry.Contract.ZkCertificateHashToIndexInQueue(&_ZkKYCRegistry.CallOpts, arg0)
}

// ZkCertificateHashToQueueTime is a free data retrieval call binding the contract method 0x4e86710b.
//
// Solidity: function ZkCertificateHashToQueueTime(bytes32 ) view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) ZkCertificateHashToQueueTime(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "ZkCertificateHashToQueueTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZkCertificateHashToQueueTime is a free data retrieval call binding the contract method 0x4e86710b.
//
// Solidity: function ZkCertificateHashToQueueTime(bytes32 ) view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistrySession) ZkCertificateHashToQueueTime(arg0 [32]byte) (*big.Int, error) {
	return _ZkKYCRegistry.Contract.ZkCertificateHashToQueueTime(&_ZkKYCRegistry.CallOpts, arg0)
}

// ZkCertificateHashToQueueTime is a free data retrieval call binding the contract method 0x4e86710b.
//
// Solidity: function ZkCertificateHashToQueueTime(bytes32 ) view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) ZkCertificateHashToQueueTime(arg0 [32]byte) (*big.Int, error) {
	return _ZkKYCRegistry.Contract.ZkCertificateHashToQueueTime(&_ZkKYCRegistry.CallOpts, arg0)
}

// ZkCertificateQueue is a free data retrieval call binding the contract method 0x37b1b263.
//
// Solidity: function ZkCertificateQueue(uint256 ) view returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) ZkCertificateQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "ZkCertificateQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZkCertificateQueue is a free data retrieval call binding the contract method 0x37b1b263.
//
// Solidity: function ZkCertificateQueue(uint256 ) view returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistrySession) ZkCertificateQueue(arg0 *big.Int) ([32]byte, error) {
	return _ZkKYCRegistry.Contract.ZkCertificateQueue(&_ZkKYCRegistry.CallOpts, arg0)
}

// ZkCertificateQueue is a free data retrieval call binding the contract method 0x37b1b263.
//
// Solidity: function ZkCertificateQueue(uint256 ) view returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) ZkCertificateQueue(arg0 *big.Int) ([32]byte, error) {
	return _ZkKYCRegistry.Contract.ZkCertificateQueue(&_ZkKYCRegistry.CallOpts, arg0)
}

// ZkCertificateToGuardian is a free data retrieval call binding the contract method 0xe09c8a15.
//
// Solidity: function ZkCertificateToGuardian(bytes32 ) view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) ZkCertificateToGuardian(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "ZkCertificateToGuardian", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkCertificateToGuardian is a free data retrieval call binding the contract method 0xe09c8a15.
//
// Solidity: function ZkCertificateToGuardian(bytes32 ) view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistrySession) ZkCertificateToGuardian(arg0 [32]byte) (common.Address, error) {
	return _ZkKYCRegistry.Contract.ZkCertificateToGuardian(&_ZkKYCRegistry.CallOpts, arg0)
}

// ZkCertificateToGuardian is a free data retrieval call binding the contract method 0xe09c8a15.
//
// Solidity: function ZkCertificateToGuardian(bytes32 ) view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) ZkCertificateToGuardian(arg0 [32]byte) (common.Address, error) {
	return _ZkKYCRegistry.Contract.ZkCertificateToGuardian(&_ZkKYCRegistry.CallOpts, arg0)
}

// GuardianRegistry is a free data retrieval call binding the contract method 0x0a1a9950.
//
// Solidity: function _GuardianRegistry() view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) GuardianRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "_GuardianRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GuardianRegistry is a free data retrieval call binding the contract method 0x0a1a9950.
//
// Solidity: function _GuardianRegistry() view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistrySession) GuardianRegistry() (common.Address, error) {
	return _ZkKYCRegistry.Contract.GuardianRegistry(&_ZkKYCRegistry.CallOpts)
}

// GuardianRegistry is a free data retrieval call binding the contract method 0x0a1a9950.
//
// Solidity: function _GuardianRegistry() view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) GuardianRegistry() (common.Address, error) {
	return _ZkKYCRegistry.Contract.GuardianRegistry(&_ZkKYCRegistry.CallOpts)
}

// AddZkCertificate is a free data retrieval call binding the contract method 0x89a64453.
//
// Solidity: function addZkCertificate(uint256 , bytes32 , bytes32[] ) pure returns()
func (_ZkKYCRegistry *ZkKYCRegistryCaller) AddZkCertificate(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte, arg2 [][32]byte) error {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "addZkCertificate", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// AddZkCertificate is a free data retrieval call binding the contract method 0x89a64453.
//
// Solidity: function addZkCertificate(uint256 , bytes32 , bytes32[] ) pure returns()
func (_ZkKYCRegistry *ZkKYCRegistrySession) AddZkCertificate(arg0 *big.Int, arg1 [32]byte, arg2 [][32]byte) error {
	return _ZkKYCRegistry.Contract.AddZkCertificate(&_ZkKYCRegistry.CallOpts, arg0, arg1, arg2)
}

// AddZkCertificate is a free data retrieval call binding the contract method 0x89a64453.
//
// Solidity: function addZkCertificate(uint256 , bytes32 , bytes32[] ) pure returns()
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) AddZkCertificate(arg0 *big.Int, arg1 [32]byte, arg2 [][32]byte) error {
	return _ZkKYCRegistry.Contract.AddZkCertificate(&_ZkKYCRegistry.CallOpts, arg0, arg1, arg2)
}

// CheckZkCertificateHashInQueue is a free data retrieval call binding the contract method 0x574aeb69.
//
// Solidity: function checkZkCertificateHashInQueue(bytes32 zkCertificateHash) view returns(bool)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) CheckZkCertificateHashInQueue(opts *bind.CallOpts, zkCertificateHash [32]byte) (bool, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "checkZkCertificateHashInQueue", zkCertificateHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckZkCertificateHashInQueue is a free data retrieval call binding the contract method 0x574aeb69.
//
// Solidity: function checkZkCertificateHashInQueue(bytes32 zkCertificateHash) view returns(bool)
func (_ZkKYCRegistry *ZkKYCRegistrySession) CheckZkCertificateHashInQueue(zkCertificateHash [32]byte) (bool, error) {
	return _ZkKYCRegistry.Contract.CheckZkCertificateHashInQueue(&_ZkKYCRegistry.CallOpts, zkCertificateHash)
}

// CheckZkCertificateHashInQueue is a free data retrieval call binding the contract method 0x574aeb69.
//
// Solidity: function checkZkCertificateHashInQueue(bytes32 zkCertificateHash) view returns(bool)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) CheckZkCertificateHashInQueue(zkCertificateHash [32]byte) (bool, error) {
	return _ZkKYCRegistry.Contract.CheckZkCertificateHashInQueue(&_ZkKYCRegistry.CallOpts, zkCertificateHash)
}

// CurrentQueuePointer is a free data retrieval call binding the contract method 0xb0e7f5bc.
//
// Solidity: function currentQueuePointer() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) CurrentQueuePointer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "currentQueuePointer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentQueuePointer is a free data retrieval call binding the contract method 0xb0e7f5bc.
//
// Solidity: function currentQueuePointer() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistrySession) CurrentQueuePointer() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.CurrentQueuePointer(&_ZkKYCRegistry.CallOpts)
}

// CurrentQueuePointer is a free data retrieval call binding the contract method 0xb0e7f5bc.
//
// Solidity: function currentQueuePointer() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) CurrentQueuePointer() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.CurrentQueuePointer(&_ZkKYCRegistry.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ZkKYCRegistry *ZkKYCRegistrySession) Description() (string, error) {
	return _ZkKYCRegistry.Contract.Description(&_ZkKYCRegistry.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) Description() (string, error) {
	return _ZkKYCRegistry.Contract.Description(&_ZkKYCRegistry.CallOpts)
}

// GetMerkleRoots is a free data retrieval call binding the contract method 0x85135a1a.
//
// Solidity: function getMerkleRoots() view returns(bytes32[])
func (_ZkKYCRegistry *ZkKYCRegistryCaller) GetMerkleRoots(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "getMerkleRoots")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetMerkleRoots is a free data retrieval call binding the contract method 0x85135a1a.
//
// Solidity: function getMerkleRoots() view returns(bytes32[])
func (_ZkKYCRegistry *ZkKYCRegistrySession) GetMerkleRoots() ([][32]byte, error) {
	return _ZkKYCRegistry.Contract.GetMerkleRoots(&_ZkKYCRegistry.CallOpts)
}

// GetMerkleRoots is a free data retrieval call binding the contract method 0x85135a1a.
//
// Solidity: function getMerkleRoots() view returns(bytes32[])
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) GetMerkleRoots() ([][32]byte, error) {
	return _ZkKYCRegistry.Contract.GetMerkleRoots(&_ZkKYCRegistry.CallOpts)
}

// GetTimeParameters is a free data retrieval call binding the contract method 0xc7b6c5f5.
//
// Solidity: function getTimeParameters(bytes32 zkCertificateHash) view returns(uint256, uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) GetTimeParameters(opts *bind.CallOpts, zkCertificateHash [32]byte) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "getTimeParameters", zkCertificateHash)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTimeParameters is a free data retrieval call binding the contract method 0xc7b6c5f5.
//
// Solidity: function getTimeParameters(bytes32 zkCertificateHash) view returns(uint256, uint256)
func (_ZkKYCRegistry *ZkKYCRegistrySession) GetTimeParameters(zkCertificateHash [32]byte) (*big.Int, *big.Int, error) {
	return _ZkKYCRegistry.Contract.GetTimeParameters(&_ZkKYCRegistry.CallOpts, zkCertificateHash)
}

// GetTimeParameters is a free data retrieval call binding the contract method 0xc7b6c5f5.
//
// Solidity: function getTimeParameters(bytes32 zkCertificateHash) view returns(uint256, uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) GetTimeParameters(zkCertificateHash [32]byte) (*big.Int, *big.Int, error) {
	return _ZkKYCRegistry.Contract.GetTimeParameters(&_ZkKYCRegistry.CallOpts, zkCertificateHash)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x38bf282e.
//
// Solidity: function hashLeftRight(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) HashLeftRight(opts *bind.CallOpts, _left [32]byte, _right [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "hashLeftRight", _left, _right)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashLeftRight is a free data retrieval call binding the contract method 0x38bf282e.
//
// Solidity: function hashLeftRight(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistrySession) HashLeftRight(_left [32]byte, _right [32]byte) ([32]byte, error) {
	return _ZkKYCRegistry.Contract.HashLeftRight(&_ZkKYCRegistry.CallOpts, _left, _right)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x38bf282e.
//
// Solidity: function hashLeftRight(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) HashLeftRight(_left [32]byte, _right [32]byte) ([32]byte, error) {
	return _ZkKYCRegistry.Contract.HashLeftRight(&_ZkKYCRegistry.CallOpts, _left, _right)
}

// HumanIDSaltRegistry is a free data retrieval call binding the contract method 0x290939b6.
//
// Solidity: function humanIDSaltRegistry() view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) HumanIDSaltRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "humanIDSaltRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HumanIDSaltRegistry is a free data retrieval call binding the contract method 0x290939b6.
//
// Solidity: function humanIDSaltRegistry() view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistrySession) HumanIDSaltRegistry() (common.Address, error) {
	return _ZkKYCRegistry.Contract.HumanIDSaltRegistry(&_ZkKYCRegistry.CallOpts)
}

// HumanIDSaltRegistry is a free data retrieval call binding the contract method 0x290939b6.
//
// Solidity: function humanIDSaltRegistry() view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) HumanIDSaltRegistry() (common.Address, error) {
	return _ZkKYCRegistry.Contract.HumanIDSaltRegistry(&_ZkKYCRegistry.CallOpts)
}

// InitBlockHeight is a free data retrieval call binding the contract method 0xef4b4119.
//
// Solidity: function initBlockHeight() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) InitBlockHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "initBlockHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitBlockHeight is a free data retrieval call binding the contract method 0xef4b4119.
//
// Solidity: function initBlockHeight() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistrySession) InitBlockHeight() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.InitBlockHeight(&_ZkKYCRegistry.CallOpts)
}

// InitBlockHeight is a free data retrieval call binding the contract method 0xef4b4119.
//
// Solidity: function initBlockHeight() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) InitBlockHeight() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.InitBlockHeight(&_ZkKYCRegistry.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistrySession) MerkleRoot() ([32]byte, error) {
	return _ZkKYCRegistry.Contract.MerkleRoot(&_ZkKYCRegistry.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) MerkleRoot() ([32]byte, error) {
	return _ZkKYCRegistry.Contract.MerkleRoot(&_ZkKYCRegistry.CallOpts)
}

// MerkleRootIndex is a free data retrieval call binding the contract method 0x390f2b65.
//
// Solidity: function merkleRootIndex(bytes32 ) view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) MerkleRootIndex(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "merkleRootIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MerkleRootIndex is a free data retrieval call binding the contract method 0x390f2b65.
//
// Solidity: function merkleRootIndex(bytes32 ) view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistrySession) MerkleRootIndex(arg0 [32]byte) (*big.Int, error) {
	return _ZkKYCRegistry.Contract.MerkleRootIndex(&_ZkKYCRegistry.CallOpts, arg0)
}

// MerkleRootIndex is a free data retrieval call binding the contract method 0x390f2b65.
//
// Solidity: function merkleRootIndex(bytes32 ) view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) MerkleRootIndex(arg0 [32]byte) (*big.Int, error) {
	return _ZkKYCRegistry.Contract.MerkleRootIndex(&_ZkKYCRegistry.CallOpts, arg0)
}

// MerkleRootValidIndex is a free data retrieval call binding the contract method 0xb0033450.
//
// Solidity: function merkleRootValidIndex() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) MerkleRootValidIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "merkleRootValidIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MerkleRootValidIndex is a free data retrieval call binding the contract method 0xb0033450.
//
// Solidity: function merkleRootValidIndex() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistrySession) MerkleRootValidIndex() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.MerkleRootValidIndex(&_ZkKYCRegistry.CallOpts)
}

// MerkleRootValidIndex is a free data retrieval call binding the contract method 0xb0033450.
//
// Solidity: function merkleRootValidIndex() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) MerkleRootValidIndex() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.MerkleRootValidIndex(&_ZkKYCRegistry.CallOpts)
}

// MerkleRoots is a free data retrieval call binding the contract method 0x71c5ecb1.
//
// Solidity: function merkleRoots(uint256 ) view returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) MerkleRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "merkleRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoots is a free data retrieval call binding the contract method 0x71c5ecb1.
//
// Solidity: function merkleRoots(uint256 ) view returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistrySession) MerkleRoots(arg0 *big.Int) ([32]byte, error) {
	return _ZkKYCRegistry.Contract.MerkleRoots(&_ZkKYCRegistry.CallOpts, arg0)
}

// MerkleRoots is a free data retrieval call binding the contract method 0x71c5ecb1.
//
// Solidity: function merkleRoots(uint256 ) view returns(bytes32)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) MerkleRoots(arg0 *big.Int) ([32]byte, error) {
	return _ZkKYCRegistry.Contract.MerkleRoots(&_ZkKYCRegistry.CallOpts, arg0)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "newOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistrySession) NewOwner() (common.Address, error) {
	return _ZkKYCRegistry.Contract.NewOwner(&_ZkKYCRegistry.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) NewOwner() (common.Address, error) {
	return _ZkKYCRegistry.Contract.NewOwner(&_ZkKYCRegistry.CallOpts)
}

// NextLeafIndex is a free data retrieval call binding the contract method 0x0be4f422.
//
// Solidity: function nextLeafIndex() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) NextLeafIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "nextLeafIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextLeafIndex is a free data retrieval call binding the contract method 0x0be4f422.
//
// Solidity: function nextLeafIndex() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistrySession) NextLeafIndex() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.NextLeafIndex(&_ZkKYCRegistry.CallOpts)
}

// NextLeafIndex is a free data retrieval call binding the contract method 0x0be4f422.
//
// Solidity: function nextLeafIndex() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) NextLeafIndex() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.NextLeafIndex(&_ZkKYCRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistrySession) Owner() (common.Address, error) {
	return _ZkKYCRegistry.Contract.Owner(&_ZkKYCRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) Owner() (common.Address, error) {
	return _ZkKYCRegistry.Contract.Owner(&_ZkKYCRegistry.CallOpts)
}

// QueueExpirationTime is a free data retrieval call binding the contract method 0x03f2ee35.
//
// Solidity: function queueExpirationTime() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) QueueExpirationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "queueExpirationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueExpirationTime is a free data retrieval call binding the contract method 0x03f2ee35.
//
// Solidity: function queueExpirationTime() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistrySession) QueueExpirationTime() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.QueueExpirationTime(&_ZkKYCRegistry.CallOpts)
}

// QueueExpirationTime is a free data retrieval call binding the contract method 0x03f2ee35.
//
// Solidity: function queueExpirationTime() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) QueueExpirationTime() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.QueueExpirationTime(&_ZkKYCRegistry.CallOpts)
}

// TreeDepth is a free data retrieval call binding the contract method 0x16a56c41.
//
// Solidity: function treeDepth() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) TreeDepth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "treeDepth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreeDepth is a free data retrieval call binding the contract method 0x16a56c41.
//
// Solidity: function treeDepth() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistrySession) TreeDepth() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.TreeDepth(&_ZkKYCRegistry.CallOpts)
}

// TreeDepth is a free data retrieval call binding the contract method 0x16a56c41.
//
// Solidity: function treeDepth() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) TreeDepth() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.TreeDepth(&_ZkKYCRegistry.CallOpts)
}

// TreeSize is a free data retrieval call binding the contract method 0x8d1ddfb1.
//
// Solidity: function treeSize() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) TreeSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "treeSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreeSize is a free data retrieval call binding the contract method 0x8d1ddfb1.
//
// Solidity: function treeSize() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistrySession) TreeSize() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.TreeSize(&_ZkKYCRegistry.CallOpts)
}

// TreeSize is a free data retrieval call binding the contract method 0x8d1ddfb1.
//
// Solidity: function treeSize() view returns(uint256)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) TreeSize() (*big.Int, error) {
	return _ZkKYCRegistry.Contract.TreeSize(&_ZkKYCRegistry.CallOpts)
}

// VerifyMerkleRoot is a free data retrieval call binding the contract method 0x0820dc2f.
//
// Solidity: function verifyMerkleRoot(bytes32 _merkleRoot) view returns(bool)
func (_ZkKYCRegistry *ZkKYCRegistryCaller) VerifyMerkleRoot(opts *bind.CallOpts, _merkleRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _ZkKYCRegistry.contract.Call(opts, &out, "verifyMerkleRoot", _merkleRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleRoot is a free data retrieval call binding the contract method 0x0820dc2f.
//
// Solidity: function verifyMerkleRoot(bytes32 _merkleRoot) view returns(bool)
func (_ZkKYCRegistry *ZkKYCRegistrySession) VerifyMerkleRoot(_merkleRoot [32]byte) (bool, error) {
	return _ZkKYCRegistry.Contract.VerifyMerkleRoot(&_ZkKYCRegistry.CallOpts, _merkleRoot)
}

// VerifyMerkleRoot is a free data retrieval call binding the contract method 0x0820dc2f.
//
// Solidity: function verifyMerkleRoot(bytes32 _merkleRoot) view returns(bool)
func (_ZkKYCRegistry *ZkKYCRegistryCallerSession) VerifyMerkleRoot(_merkleRoot [32]byte) (bool, error) {
	return _ZkKYCRegistry.Contract.VerifyMerkleRoot(&_ZkKYCRegistry.CallOpts, _merkleRoot)
}

// AddZkKYC is a paid mutator transaction binding the contract method 0xcb9508ec.
//
// Solidity: function addZkKYC(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof, uint256 idHash, uint256 saltHash, uint256 expirationTime) returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactor) AddZkKYC(opts *bind.TransactOpts, leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte, idHash *big.Int, saltHash *big.Int, expirationTime *big.Int) (*types.Transaction, error) {
	return _ZkKYCRegistry.contract.Transact(opts, "addZkKYC", leafIndex, zkCertificateHash, merkleProof, idHash, saltHash, expirationTime)
}

// AddZkKYC is a paid mutator transaction binding the contract method 0xcb9508ec.
//
// Solidity: function addZkKYC(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof, uint256 idHash, uint256 saltHash, uint256 expirationTime) returns()
func (_ZkKYCRegistry *ZkKYCRegistrySession) AddZkKYC(leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte, idHash *big.Int, saltHash *big.Int, expirationTime *big.Int) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.AddZkKYC(&_ZkKYCRegistry.TransactOpts, leafIndex, zkCertificateHash, merkleProof, idHash, saltHash, expirationTime)
}

// AddZkKYC is a paid mutator transaction binding the contract method 0xcb9508ec.
//
// Solidity: function addZkKYC(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof, uint256 idHash, uint256 saltHash, uint256 expirationTime) returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactorSession) AddZkKYC(leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte, idHash *big.Int, saltHash *big.Int, expirationTime *big.Int) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.AddZkKYC(&_ZkKYCRegistry.TransactOpts, leafIndex, zkCertificateHash, merkleProof, idHash, saltHash, expirationTime)
}

// ChangeQueueExpirationTime is a paid mutator transaction binding the contract method 0x6a7d1e44.
//
// Solidity: function changeQueueExpirationTime(uint256 newTime) returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactor) ChangeQueueExpirationTime(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _ZkKYCRegistry.contract.Transact(opts, "changeQueueExpirationTime", newTime)
}

// ChangeQueueExpirationTime is a paid mutator transaction binding the contract method 0x6a7d1e44.
//
// Solidity: function changeQueueExpirationTime(uint256 newTime) returns()
func (_ZkKYCRegistry *ZkKYCRegistrySession) ChangeQueueExpirationTime(newTime *big.Int) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.ChangeQueueExpirationTime(&_ZkKYCRegistry.TransactOpts, newTime)
}

// ChangeQueueExpirationTime is a paid mutator transaction binding the contract method 0x6a7d1e44.
//
// Solidity: function changeQueueExpirationTime(uint256 newTime) returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactorSession) ChangeQueueExpirationTime(newTime *big.Int) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.ChangeQueueExpirationTime(&_ZkKYCRegistry.TransactOpts, newTime)
}

// RegisterToQueue is a paid mutator transaction binding the contract method 0x9de27abe.
//
// Solidity: function registerToQueue(bytes32 zkCertificateHash) returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactor) RegisterToQueue(opts *bind.TransactOpts, zkCertificateHash [32]byte) (*types.Transaction, error) {
	return _ZkKYCRegistry.contract.Transact(opts, "registerToQueue", zkCertificateHash)
}

// RegisterToQueue is a paid mutator transaction binding the contract method 0x9de27abe.
//
// Solidity: function registerToQueue(bytes32 zkCertificateHash) returns()
func (_ZkKYCRegistry *ZkKYCRegistrySession) RegisterToQueue(zkCertificateHash [32]byte) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.RegisterToQueue(&_ZkKYCRegistry.TransactOpts, zkCertificateHash)
}

// RegisterToQueue is a paid mutator transaction binding the contract method 0x9de27abe.
//
// Solidity: function registerToQueue(bytes32 zkCertificateHash) returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactorSession) RegisterToQueue(zkCertificateHash [32]byte) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.RegisterToQueue(&_ZkKYCRegistry.TransactOpts, zkCertificateHash)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkKYCRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZkKYCRegistry *ZkKYCRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.RenounceOwnership(&_ZkKYCRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.RenounceOwnership(&_ZkKYCRegistry.TransactOpts)
}

// RevokeZkCertificate is a paid mutator transaction binding the contract method 0xa2873348.
//
// Solidity: function revokeZkCertificate(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof) returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactor) RevokeZkCertificate(opts *bind.TransactOpts, leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkKYCRegistry.contract.Transact(opts, "revokeZkCertificate", leafIndex, zkCertificateHash, merkleProof)
}

// RevokeZkCertificate is a paid mutator transaction binding the contract method 0xa2873348.
//
// Solidity: function revokeZkCertificate(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof) returns()
func (_ZkKYCRegistry *ZkKYCRegistrySession) RevokeZkCertificate(leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.RevokeZkCertificate(&_ZkKYCRegistry.TransactOpts, leafIndex, zkCertificateHash, merkleProof)
}

// RevokeZkCertificate is a paid mutator transaction binding the contract method 0xa2873348.
//
// Solidity: function revokeZkCertificate(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof) returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactorSession) RevokeZkCertificate(leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.RevokeZkCertificate(&_ZkKYCRegistry.TransactOpts, leafIndex, zkCertificateHash, merkleProof)
}

// SetNewOwner is a paid mutator transaction binding the contract method 0xf5a1f5b4.
//
// Solidity: function setNewOwner(address _newOwner) returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactor) SetNewOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ZkKYCRegistry.contract.Transact(opts, "setNewOwner", _newOwner)
}

// SetNewOwner is a paid mutator transaction binding the contract method 0xf5a1f5b4.
//
// Solidity: function setNewOwner(address _newOwner) returns()
func (_ZkKYCRegistry *ZkKYCRegistrySession) SetNewOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.SetNewOwner(&_ZkKYCRegistry.TransactOpts, _newOwner)
}

// SetNewOwner is a paid mutator transaction binding the contract method 0xf5a1f5b4.
//
// Solidity: function setNewOwner(address _newOwner) returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactorSession) SetNewOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.SetNewOwner(&_ZkKYCRegistry.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x880ad0af.
//
// Solidity: function transferOwnership() returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactor) TransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkKYCRegistry.contract.Transact(opts, "transferOwnership")
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x880ad0af.
//
// Solidity: function transferOwnership() returns()
func (_ZkKYCRegistry *ZkKYCRegistrySession) TransferOwnership() (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.TransferOwnership(&_ZkKYCRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x880ad0af.
//
// Solidity: function transferOwnership() returns()
func (_ZkKYCRegistry *ZkKYCRegistryTransactorSession) TransferOwnership() (*types.Transaction, error) {
	return _ZkKYCRegistry.Contract.TransferOwnership(&_ZkKYCRegistry.TransactOpts)
}

// ZkKYCRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZkKYCRegistry contract.
type ZkKYCRegistryInitializedIterator struct {
	Event *ZkKYCRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ZkKYCRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkKYCRegistryInitialized)
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
		it.Event = new(ZkKYCRegistryInitialized)
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
func (it *ZkKYCRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkKYCRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkKYCRegistryInitialized represents a Initialized event raised by the ZkKYCRegistry contract.
type ZkKYCRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZkKYCRegistry *ZkKYCRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZkKYCRegistryInitializedIterator, error) {

	logs, sub, err := _ZkKYCRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZkKYCRegistryInitializedIterator{contract: _ZkKYCRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZkKYCRegistry *ZkKYCRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZkKYCRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ZkKYCRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkKYCRegistryInitialized)
				if err := _ZkKYCRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ZkKYCRegistry *ZkKYCRegistryFilterer) ParseInitialized(log types.Log) (*ZkKYCRegistryInitialized, error) {
	event := new(ZkKYCRegistryInitialized)
	if err := _ZkKYCRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkKYCRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZkKYCRegistry contract.
type ZkKYCRegistryOwnershipTransferredIterator struct {
	Event *ZkKYCRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZkKYCRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkKYCRegistryOwnershipTransferred)
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
		it.Event = new(ZkKYCRegistryOwnershipTransferred)
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
func (it *ZkKYCRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkKYCRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkKYCRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ZkKYCRegistry contract.
type ZkKYCRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZkKYCRegistry *ZkKYCRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZkKYCRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZkKYCRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZkKYCRegistryOwnershipTransferredIterator{contract: _ZkKYCRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZkKYCRegistry *ZkKYCRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZkKYCRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZkKYCRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkKYCRegistryOwnershipTransferred)
				if err := _ZkKYCRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZkKYCRegistry *ZkKYCRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ZkKYCRegistryOwnershipTransferred, error) {
	event := new(ZkKYCRegistryOwnershipTransferred)
	if err := _ZkKYCRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkKYCRegistryZkCertificateAdditionIterator is returned from FilterZkCertificateAddition and is used to iterate over the raw logs and unpacked data for ZkCertificateAddition events raised by the ZkKYCRegistry contract.
type ZkKYCRegistryZkCertificateAdditionIterator struct {
	Event *ZkKYCRegistryZkCertificateAddition // Event containing the contract specifics and raw log

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
func (it *ZkKYCRegistryZkCertificateAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkKYCRegistryZkCertificateAddition)
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
		it.Event = new(ZkKYCRegistryZkCertificateAddition)
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
func (it *ZkKYCRegistryZkCertificateAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkKYCRegistryZkCertificateAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkKYCRegistryZkCertificateAddition represents a ZkCertificateAddition event raised by the ZkKYCRegistry contract.
type ZkKYCRegistryZkCertificateAddition struct {
	ZkCertificateLeafHash [32]byte
	Guardian              common.Address
	Index                 *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterZkCertificateAddition is a free log retrieval operation binding the contract event 0xd8ed161990bf70dd13139e68384461b9fffbd1d3433c3bd56a9ccd22e0118cb6.
//
// Solidity: event zkCertificateAddition(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint256 index)
func (_ZkKYCRegistry *ZkKYCRegistryFilterer) FilterZkCertificateAddition(opts *bind.FilterOpts, zkCertificateLeafHash [][32]byte, Guardian []common.Address) (*ZkKYCRegistryZkCertificateAdditionIterator, error) {

	var zkCertificateLeafHashRule []interface{}
	for _, zkCertificateLeafHashItem := range zkCertificateLeafHash {
		zkCertificateLeafHashRule = append(zkCertificateLeafHashRule, zkCertificateLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _ZkKYCRegistry.contract.FilterLogs(opts, "zkCertificateAddition", zkCertificateLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return &ZkKYCRegistryZkCertificateAdditionIterator{contract: _ZkKYCRegistry.contract, event: "zkCertificateAddition", logs: logs, sub: sub}, nil
}

// WatchZkCertificateAddition is a free log subscription operation binding the contract event 0xd8ed161990bf70dd13139e68384461b9fffbd1d3433c3bd56a9ccd22e0118cb6.
//
// Solidity: event zkCertificateAddition(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint256 index)
func (_ZkKYCRegistry *ZkKYCRegistryFilterer) WatchZkCertificateAddition(opts *bind.WatchOpts, sink chan<- *ZkKYCRegistryZkCertificateAddition, zkCertificateLeafHash [][32]byte, Guardian []common.Address) (event.Subscription, error) {

	var zkCertificateLeafHashRule []interface{}
	for _, zkCertificateLeafHashItem := range zkCertificateLeafHash {
		zkCertificateLeafHashRule = append(zkCertificateLeafHashRule, zkCertificateLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _ZkKYCRegistry.contract.WatchLogs(opts, "zkCertificateAddition", zkCertificateLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkKYCRegistryZkCertificateAddition)
				if err := _ZkKYCRegistry.contract.UnpackLog(event, "zkCertificateAddition", log); err != nil {
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
func (_ZkKYCRegistry *ZkKYCRegistryFilterer) ParseZkCertificateAddition(log types.Log) (*ZkKYCRegistryZkCertificateAddition, error) {
	event := new(ZkKYCRegistryZkCertificateAddition)
	if err := _ZkKYCRegistry.contract.UnpackLog(event, "zkCertificateAddition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkKYCRegistryZkCertificateRevocationIterator is returned from FilterZkCertificateRevocation and is used to iterate over the raw logs and unpacked data for ZkCertificateRevocation events raised by the ZkKYCRegistry contract.
type ZkKYCRegistryZkCertificateRevocationIterator struct {
	Event *ZkKYCRegistryZkCertificateRevocation // Event containing the contract specifics and raw log

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
func (it *ZkKYCRegistryZkCertificateRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkKYCRegistryZkCertificateRevocation)
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
		it.Event = new(ZkKYCRegistryZkCertificateRevocation)
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
func (it *ZkKYCRegistryZkCertificateRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkKYCRegistryZkCertificateRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkKYCRegistryZkCertificateRevocation represents a ZkCertificateRevocation event raised by the ZkKYCRegistry contract.
type ZkKYCRegistryZkCertificateRevocation struct {
	ZkCertificateLeafHash [32]byte
	Guardian              common.Address
	Index                 *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterZkCertificateRevocation is a free log retrieval operation binding the contract event 0x2d6db1db44b3ed99e9b90701d664345ef76dd37336779182d6251abf294840a6.
//
// Solidity: event zkCertificateRevocation(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint256 index)
func (_ZkKYCRegistry *ZkKYCRegistryFilterer) FilterZkCertificateRevocation(opts *bind.FilterOpts, zkCertificateLeafHash [][32]byte, Guardian []common.Address) (*ZkKYCRegistryZkCertificateRevocationIterator, error) {

	var zkCertificateLeafHashRule []interface{}
	for _, zkCertificateLeafHashItem := range zkCertificateLeafHash {
		zkCertificateLeafHashRule = append(zkCertificateLeafHashRule, zkCertificateLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _ZkKYCRegistry.contract.FilterLogs(opts, "zkCertificateRevocation", zkCertificateLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return &ZkKYCRegistryZkCertificateRevocationIterator{contract: _ZkKYCRegistry.contract, event: "zkCertificateRevocation", logs: logs, sub: sub}, nil
}

// WatchZkCertificateRevocation is a free log subscription operation binding the contract event 0x2d6db1db44b3ed99e9b90701d664345ef76dd37336779182d6251abf294840a6.
//
// Solidity: event zkCertificateRevocation(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint256 index)
func (_ZkKYCRegistry *ZkKYCRegistryFilterer) WatchZkCertificateRevocation(opts *bind.WatchOpts, sink chan<- *ZkKYCRegistryZkCertificateRevocation, zkCertificateLeafHash [][32]byte, Guardian []common.Address) (event.Subscription, error) {

	var zkCertificateLeafHashRule []interface{}
	for _, zkCertificateLeafHashItem := range zkCertificateLeafHash {
		zkCertificateLeafHashRule = append(zkCertificateLeafHashRule, zkCertificateLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _ZkKYCRegistry.contract.WatchLogs(opts, "zkCertificateRevocation", zkCertificateLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkKYCRegistryZkCertificateRevocation)
				if err := _ZkKYCRegistry.contract.UnpackLog(event, "zkCertificateRevocation", log); err != nil {
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
func (_ZkKYCRegistry *ZkKYCRegistryFilterer) ParseZkCertificateRevocation(log types.Log) (*ZkKYCRegistryZkCertificateRevocation, error) {
	event := new(ZkKYCRegistryZkCertificateRevocation)
	if err := _ZkKYCRegistry.contract.UnpackLog(event, "zkCertificateRevocation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
