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

// KYCRecordRegistryMetaData contains all meta data concerning the KYCRecordRegistry contract.
var KYCRecordRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"zkKYCRecordLeafHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Guardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"zkKYCRecordAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"zkKYCRecordLeafHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Guardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"zkKYCRecordRevocation\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ZERO_VALUE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ZKKYCRecordToCenter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_GuardianRegistry\",\"outputs\":[{\"internalType\":\"contractGuardianRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zkKYCRecordHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"addZkKYCRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"hashLeftRight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextLeafIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zkKYCRecordHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"revokeZkKYCRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"zeros\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// KYCRecordRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use KYCRecordRegistryMetaData.ABI instead.
var KYCRecordRegistryABI = KYCRecordRegistryMetaData.ABI

// KYCRecordRegistry is an auto generated Go binding around an Ethereum contract.
type KYCRecordRegistry struct {
	KYCRecordRegistryCaller     // Read-only binding to the contract
	KYCRecordRegistryTransactor // Write-only binding to the contract
	KYCRecordRegistryFilterer   // Log filterer for contract events
}

// KYCRecordRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type KYCRecordRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KYCRecordRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KYCRecordRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KYCRecordRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KYCRecordRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KYCRecordRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KYCRecordRegistrySession struct {
	Contract     *KYCRecordRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KYCRecordRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KYCRecordRegistryCallerSession struct {
	Contract *KYCRecordRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// KYCRecordRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KYCRecordRegistryTransactorSession struct {
	Contract     *KYCRecordRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// KYCRecordRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type KYCRecordRegistryRaw struct {
	Contract *KYCRecordRegistry // Generic contract binding to access the raw methods on
}

// KYCRecordRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KYCRecordRegistryCallerRaw struct {
	Contract *KYCRecordRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// KYCRecordRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KYCRecordRegistryTransactorRaw struct {
	Contract *KYCRecordRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKYCRecordRegistry creates a new instance of KYCRecordRegistry, bound to a specific deployed contract.
func NewKYCRecordRegistry(address common.Address, backend bind.ContractBackend) (*KYCRecordRegistry, error) {
	contract, err := bindKYCRecordRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KYCRecordRegistry{KYCRecordRegistryCaller: KYCRecordRegistryCaller{contract: contract}, KYCRecordRegistryTransactor: KYCRecordRegistryTransactor{contract: contract}, KYCRecordRegistryFilterer: KYCRecordRegistryFilterer{contract: contract}}, nil
}

// NewKYCRecordRegistryCaller creates a new read-only instance of KYCRecordRegistry, bound to a specific deployed contract.
func NewKYCRecordRegistryCaller(address common.Address, caller bind.ContractCaller) (*KYCRecordRegistryCaller, error) {
	contract, err := bindKYCRecordRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KYCRecordRegistryCaller{contract: contract}, nil
}

// NewKYCRecordRegistryTransactor creates a new write-only instance of KYCRecordRegistry, bound to a specific deployed contract.
func NewKYCRecordRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*KYCRecordRegistryTransactor, error) {
	contract, err := bindKYCRecordRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KYCRecordRegistryTransactor{contract: contract}, nil
}

// NewKYCRecordRegistryFilterer creates a new log filterer instance of KYCRecordRegistry, bound to a specific deployed contract.
func NewKYCRecordRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*KYCRecordRegistryFilterer, error) {
	contract, err := bindKYCRecordRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KYCRecordRegistryFilterer{contract: contract}, nil
}

// bindKYCRecordRegistry binds a generic wrapper to an already deployed contract.
func bindKYCRecordRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KYCRecordRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KYCRecordRegistry *KYCRecordRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KYCRecordRegistry.Contract.KYCRecordRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KYCRecordRegistry *KYCRecordRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KYCRecordRegistry.Contract.KYCRecordRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KYCRecordRegistry *KYCRecordRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KYCRecordRegistry.Contract.KYCRecordRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KYCRecordRegistry *KYCRecordRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KYCRecordRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KYCRecordRegistry *KYCRecordRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KYCRecordRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KYCRecordRegistry *KYCRecordRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KYCRecordRegistry.Contract.contract.Transact(opts, method, params...)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(bytes32)
func (_KYCRecordRegistry *KYCRecordRegistryCaller) ZEROVALUE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KYCRecordRegistry.contract.Call(opts, &out, "ZERO_VALUE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(bytes32)
func (_KYCRecordRegistry *KYCRecordRegistrySession) ZEROVALUE() ([32]byte, error) {
	return _KYCRecordRegistry.Contract.ZEROVALUE(&_KYCRecordRegistry.CallOpts)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(bytes32)
func (_KYCRecordRegistry *KYCRecordRegistryCallerSession) ZEROVALUE() ([32]byte, error) {
	return _KYCRecordRegistry.Contract.ZEROVALUE(&_KYCRecordRegistry.CallOpts)
}

// ZKKYCRecordToCenter is a free data retrieval call binding the contract method 0x64d4aaaf.
//
// Solidity: function ZKKYCRecordToCenter(bytes32 ) view returns(address)
func (_KYCRecordRegistry *KYCRecordRegistryCaller) ZKKYCRecordToCenter(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _KYCRecordRegistry.contract.Call(opts, &out, "ZKKYCRecordToCenter", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZKKYCRecordToCenter is a free data retrieval call binding the contract method 0x64d4aaaf.
//
// Solidity: function ZKKYCRecordToCenter(bytes32 ) view returns(address)
func (_KYCRecordRegistry *KYCRecordRegistrySession) ZKKYCRecordToCenter(arg0 [32]byte) (common.Address, error) {
	return _KYCRecordRegistry.Contract.ZKKYCRecordToCenter(&_KYCRecordRegistry.CallOpts, arg0)
}

// ZKKYCRecordToCenter is a free data retrieval call binding the contract method 0x64d4aaaf.
//
// Solidity: function ZKKYCRecordToCenter(bytes32 ) view returns(address)
func (_KYCRecordRegistry *KYCRecordRegistryCallerSession) ZKKYCRecordToCenter(arg0 [32]byte) (common.Address, error) {
	return _KYCRecordRegistry.Contract.ZKKYCRecordToCenter(&_KYCRecordRegistry.CallOpts, arg0)
}

// GuardianRegistry is a free data retrieval call binding the contract method 0x0a1a9950.
//
// Solidity: function _GuardianRegistry() view returns(address)
func (_KYCRecordRegistry *KYCRecordRegistryCaller) GuardianRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KYCRecordRegistry.contract.Call(opts, &out, "_GuardianRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GuardianRegistry is a free data retrieval call binding the contract method 0x0a1a9950.
//
// Solidity: function _GuardianRegistry() view returns(address)
func (_KYCRecordRegistry *KYCRecordRegistrySession) GuardianRegistry() (common.Address, error) {
	return _KYCRecordRegistry.Contract.GuardianRegistry(&_KYCRecordRegistry.CallOpts)
}

// GuardianRegistry is a free data retrieval call binding the contract method 0x0a1a9950.
//
// Solidity: function _GuardianRegistry() view returns(address)
func (_KYCRecordRegistry *KYCRecordRegistryCallerSession) GuardianRegistry() (common.Address, error) {
	return _KYCRecordRegistry.Contract.GuardianRegistry(&_KYCRecordRegistry.CallOpts)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x38bf282e.
//
// Solidity: function hashLeftRight(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_KYCRecordRegistry *KYCRecordRegistryCaller) HashLeftRight(opts *bind.CallOpts, _left [32]byte, _right [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _KYCRecordRegistry.contract.Call(opts, &out, "hashLeftRight", _left, _right)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashLeftRight is a free data retrieval call binding the contract method 0x38bf282e.
//
// Solidity: function hashLeftRight(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_KYCRecordRegistry *KYCRecordRegistrySession) HashLeftRight(_left [32]byte, _right [32]byte) ([32]byte, error) {
	return _KYCRecordRegistry.Contract.HashLeftRight(&_KYCRecordRegistry.CallOpts, _left, _right)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x38bf282e.
//
// Solidity: function hashLeftRight(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_KYCRecordRegistry *KYCRecordRegistryCallerSession) HashLeftRight(_left [32]byte, _right [32]byte) ([32]byte, error) {
	return _KYCRecordRegistry.Contract.HashLeftRight(&_KYCRecordRegistry.CallOpts, _left, _right)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_KYCRecordRegistry *KYCRecordRegistryCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KYCRecordRegistry.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_KYCRecordRegistry *KYCRecordRegistrySession) MerkleRoot() ([32]byte, error) {
	return _KYCRecordRegistry.Contract.MerkleRoot(&_KYCRecordRegistry.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_KYCRecordRegistry *KYCRecordRegistryCallerSession) MerkleRoot() ([32]byte, error) {
	return _KYCRecordRegistry.Contract.MerkleRoot(&_KYCRecordRegistry.CallOpts)
}

// NextLeafIndex is a free data retrieval call binding the contract method 0x0be4f422.
//
// Solidity: function nextLeafIndex() view returns(uint256)
func (_KYCRecordRegistry *KYCRecordRegistryCaller) NextLeafIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KYCRecordRegistry.contract.Call(opts, &out, "nextLeafIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextLeafIndex is a free data retrieval call binding the contract method 0x0be4f422.
//
// Solidity: function nextLeafIndex() view returns(uint256)
func (_KYCRecordRegistry *KYCRecordRegistrySession) NextLeafIndex() (*big.Int, error) {
	return _KYCRecordRegistry.Contract.NextLeafIndex(&_KYCRecordRegistry.CallOpts)
}

// NextLeafIndex is a free data retrieval call binding the contract method 0x0be4f422.
//
// Solidity: function nextLeafIndex() view returns(uint256)
func (_KYCRecordRegistry *KYCRecordRegistryCallerSession) NextLeafIndex() (*big.Int, error) {
	return _KYCRecordRegistry.Contract.NextLeafIndex(&_KYCRecordRegistry.CallOpts)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 ) view returns(bytes32)
func (_KYCRecordRegistry *KYCRecordRegistryCaller) Zeros(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _KYCRecordRegistry.contract.Call(opts, &out, "zeros", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 ) view returns(bytes32)
func (_KYCRecordRegistry *KYCRecordRegistrySession) Zeros(arg0 *big.Int) ([32]byte, error) {
	return _KYCRecordRegistry.Contract.Zeros(&_KYCRecordRegistry.CallOpts, arg0)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 ) view returns(bytes32)
func (_KYCRecordRegistry *KYCRecordRegistryCallerSession) Zeros(arg0 *big.Int) ([32]byte, error) {
	return _KYCRecordRegistry.Contract.Zeros(&_KYCRecordRegistry.CallOpts, arg0)
}

// AddZkKYCRecord is a paid mutator transaction binding the contract method 0xb43acadc.
//
// Solidity: function addZkKYCRecord(uint256 leafIndex, bytes32 zkKYCRecordHash, bytes32[] merkleProof) returns()
func (_KYCRecordRegistry *KYCRecordRegistryTransactor) AddZkKYCRecord(opts *bind.TransactOpts, leafIndex *big.Int, zkKYCRecordHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _KYCRecordRegistry.contract.Transact(opts, "addZkKYCRecord", leafIndex, zkKYCRecordHash, merkleProof)
}

// AddZkKYCRecord is a paid mutator transaction binding the contract method 0xb43acadc.
//
// Solidity: function addZkKYCRecord(uint256 leafIndex, bytes32 zkKYCRecordHash, bytes32[] merkleProof) returns()
func (_KYCRecordRegistry *KYCRecordRegistrySession) AddZkKYCRecord(leafIndex *big.Int, zkKYCRecordHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _KYCRecordRegistry.Contract.AddZkKYCRecord(&_KYCRecordRegistry.TransactOpts, leafIndex, zkKYCRecordHash, merkleProof)
}

// AddZkKYCRecord is a paid mutator transaction binding the contract method 0xb43acadc.
//
// Solidity: function addZkKYCRecord(uint256 leafIndex, bytes32 zkKYCRecordHash, bytes32[] merkleProof) returns()
func (_KYCRecordRegistry *KYCRecordRegistryTransactorSession) AddZkKYCRecord(leafIndex *big.Int, zkKYCRecordHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _KYCRecordRegistry.Contract.AddZkKYCRecord(&_KYCRecordRegistry.TransactOpts, leafIndex, zkKYCRecordHash, merkleProof)
}

// RevokeZkKYCRecord is a paid mutator transaction binding the contract method 0x6ac20d41.
//
// Solidity: function revokeZkKYCRecord(uint256 leafIndex, bytes32 zkKYCRecordHash, bytes32[] merkleProof) returns()
func (_KYCRecordRegistry *KYCRecordRegistryTransactor) RevokeZkKYCRecord(opts *bind.TransactOpts, leafIndex *big.Int, zkKYCRecordHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _KYCRecordRegistry.contract.Transact(opts, "revokeZkKYCRecord", leafIndex, zkKYCRecordHash, merkleProof)
}

// RevokeZkKYCRecord is a paid mutator transaction binding the contract method 0x6ac20d41.
//
// Solidity: function revokeZkKYCRecord(uint256 leafIndex, bytes32 zkKYCRecordHash, bytes32[] merkleProof) returns()
func (_KYCRecordRegistry *KYCRecordRegistrySession) RevokeZkKYCRecord(leafIndex *big.Int, zkKYCRecordHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _KYCRecordRegistry.Contract.RevokeZkKYCRecord(&_KYCRecordRegistry.TransactOpts, leafIndex, zkKYCRecordHash, merkleProof)
}

// RevokeZkKYCRecord is a paid mutator transaction binding the contract method 0x6ac20d41.
//
// Solidity: function revokeZkKYCRecord(uint256 leafIndex, bytes32 zkKYCRecordHash, bytes32[] merkleProof) returns()
func (_KYCRecordRegistry *KYCRecordRegistryTransactorSession) RevokeZkKYCRecord(leafIndex *big.Int, zkKYCRecordHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _KYCRecordRegistry.Contract.RevokeZkKYCRecord(&_KYCRecordRegistry.TransactOpts, leafIndex, zkKYCRecordHash, merkleProof)
}

// KYCRecordRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the KYCRecordRegistry contract.
type KYCRecordRegistryInitializedIterator struct {
	Event *KYCRecordRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *KYCRecordRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCRecordRegistryInitialized)
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
		it.Event = new(KYCRecordRegistryInitialized)
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
func (it *KYCRecordRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCRecordRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCRecordRegistryInitialized represents a Initialized event raised by the KYCRecordRegistry contract.
type KYCRecordRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_KYCRecordRegistry *KYCRecordRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*KYCRecordRegistryInitializedIterator, error) {

	logs, sub, err := _KYCRecordRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &KYCRecordRegistryInitializedIterator{contract: _KYCRecordRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_KYCRecordRegistry *KYCRecordRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *KYCRecordRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _KYCRecordRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCRecordRegistryInitialized)
				if err := _KYCRecordRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_KYCRecordRegistry *KYCRecordRegistryFilterer) ParseInitialized(log types.Log) (*KYCRecordRegistryInitialized, error) {
	event := new(KYCRecordRegistryInitialized)
	if err := _KYCRecordRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KYCRecordRegistryZkKYCRecordAdditionIterator is returned from FilterZkKYCRecordAddition and is used to iterate over the raw logs and unpacked data for ZkKYCRecordAddition events raised by the KYCRecordRegistry contract.
type KYCRecordRegistryZkKYCRecordAdditionIterator struct {
	Event *KYCRecordRegistryZkKYCRecordAddition // Event containing the contract specifics and raw log

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
func (it *KYCRecordRegistryZkKYCRecordAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCRecordRegistryZkKYCRecordAddition)
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
		it.Event = new(KYCRecordRegistryZkKYCRecordAddition)
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
func (it *KYCRecordRegistryZkKYCRecordAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCRecordRegistryZkKYCRecordAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCRecordRegistryZkKYCRecordAddition represents a ZkKYCRecordAddition event raised by the KYCRecordRegistry contract.
type KYCRecordRegistryZkKYCRecordAddition struct {
	ZkKYCRecordLeafHash [32]byte
	Guardian            common.Address
	Index               *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterZkKYCRecordAddition is a free log retrieval operation binding the contract event 0xbcb38b7c28c93727b78345fea2700ee13085390a48043d65339e46d9f6fa17d0.
//
// Solidity: event zkKYCRecordAddition(bytes32 indexed zkKYCRecordLeafHash, address indexed Guardian, uint256 index)
func (_KYCRecordRegistry *KYCRecordRegistryFilterer) FilterZkKYCRecordAddition(opts *bind.FilterOpts, zkKYCRecordLeafHash [][32]byte, Guardian []common.Address) (*KYCRecordRegistryZkKYCRecordAdditionIterator, error) {

	var zkKYCRecordLeafHashRule []interface{}
	for _, zkKYCRecordLeafHashItem := range zkKYCRecordLeafHash {
		zkKYCRecordLeafHashRule = append(zkKYCRecordLeafHashRule, zkKYCRecordLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _KYCRecordRegistry.contract.FilterLogs(opts, "zkKYCRecordAddition", zkKYCRecordLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return &KYCRecordRegistryZkKYCRecordAdditionIterator{contract: _KYCRecordRegistry.contract, event: "zkKYCRecordAddition", logs: logs, sub: sub}, nil
}

// WatchZkKYCRecordAddition is a free log subscription operation binding the contract event 0xbcb38b7c28c93727b78345fea2700ee13085390a48043d65339e46d9f6fa17d0.
//
// Solidity: event zkKYCRecordAddition(bytes32 indexed zkKYCRecordLeafHash, address indexed Guardian, uint256 index)
func (_KYCRecordRegistry *KYCRecordRegistryFilterer) WatchZkKYCRecordAddition(opts *bind.WatchOpts, sink chan<- *KYCRecordRegistryZkKYCRecordAddition, zkKYCRecordLeafHash [][32]byte, Guardian []common.Address) (event.Subscription, error) {

	var zkKYCRecordLeafHashRule []interface{}
	for _, zkKYCRecordLeafHashItem := range zkKYCRecordLeafHash {
		zkKYCRecordLeafHashRule = append(zkKYCRecordLeafHashRule, zkKYCRecordLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _KYCRecordRegistry.contract.WatchLogs(opts, "zkKYCRecordAddition", zkKYCRecordLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCRecordRegistryZkKYCRecordAddition)
				if err := _KYCRecordRegistry.contract.UnpackLog(event, "zkKYCRecordAddition", log); err != nil {
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

// ParseZkKYCRecordAddition is a log parse operation binding the contract event 0xbcb38b7c28c93727b78345fea2700ee13085390a48043d65339e46d9f6fa17d0.
//
// Solidity: event zkKYCRecordAddition(bytes32 indexed zkKYCRecordLeafHash, address indexed Guardian, uint256 index)
func (_KYCRecordRegistry *KYCRecordRegistryFilterer) ParseZkKYCRecordAddition(log types.Log) (*KYCRecordRegistryZkKYCRecordAddition, error) {
	event := new(KYCRecordRegistryZkKYCRecordAddition)
	if err := _KYCRecordRegistry.contract.UnpackLog(event, "zkKYCRecordAddition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KYCRecordRegistryZkKYCRecordRevocationIterator is returned from FilterZkKYCRecordRevocation and is used to iterate over the raw logs and unpacked data for ZkKYCRecordRevocation events raised by the KYCRecordRegistry contract.
type KYCRecordRegistryZkKYCRecordRevocationIterator struct {
	Event *KYCRecordRegistryZkKYCRecordRevocation // Event containing the contract specifics and raw log

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
func (it *KYCRecordRegistryZkKYCRecordRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCRecordRegistryZkKYCRecordRevocation)
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
		it.Event = new(KYCRecordRegistryZkKYCRecordRevocation)
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
func (it *KYCRecordRegistryZkKYCRecordRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCRecordRegistryZkKYCRecordRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCRecordRegistryZkKYCRecordRevocation represents a ZkKYCRecordRevocation event raised by the KYCRecordRegistry contract.
type KYCRecordRegistryZkKYCRecordRevocation struct {
	ZkKYCRecordLeafHash [32]byte
	Guardian            common.Address
	Index               *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterZkKYCRecordRevocation is a free log retrieval operation binding the contract event 0x61ca9c88d3b99939f9c40240b2a45041aa97b3913eb5633e7b477e8247b0d8d4.
//
// Solidity: event zkKYCRecordRevocation(bytes32 indexed zkKYCRecordLeafHash, address indexed Guardian, uint256 index)
func (_KYCRecordRegistry *KYCRecordRegistryFilterer) FilterZkKYCRecordRevocation(opts *bind.FilterOpts, zkKYCRecordLeafHash [][32]byte, Guardian []common.Address) (*KYCRecordRegistryZkKYCRecordRevocationIterator, error) {

	var zkKYCRecordLeafHashRule []interface{}
	for _, zkKYCRecordLeafHashItem := range zkKYCRecordLeafHash {
		zkKYCRecordLeafHashRule = append(zkKYCRecordLeafHashRule, zkKYCRecordLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _KYCRecordRegistry.contract.FilterLogs(opts, "zkKYCRecordRevocation", zkKYCRecordLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return &KYCRecordRegistryZkKYCRecordRevocationIterator{contract: _KYCRecordRegistry.contract, event: "zkKYCRecordRevocation", logs: logs, sub: sub}, nil
}

// WatchZkKYCRecordRevocation is a free log subscription operation binding the contract event 0x61ca9c88d3b99939f9c40240b2a45041aa97b3913eb5633e7b477e8247b0d8d4.
//
// Solidity: event zkKYCRecordRevocation(bytes32 indexed zkKYCRecordLeafHash, address indexed Guardian, uint256 index)
func (_KYCRecordRegistry *KYCRecordRegistryFilterer) WatchZkKYCRecordRevocation(opts *bind.WatchOpts, sink chan<- *KYCRecordRegistryZkKYCRecordRevocation, zkKYCRecordLeafHash [][32]byte, Guardian []common.Address) (event.Subscription, error) {

	var zkKYCRecordLeafHashRule []interface{}
	for _, zkKYCRecordLeafHashItem := range zkKYCRecordLeafHash {
		zkKYCRecordLeafHashRule = append(zkKYCRecordLeafHashRule, zkKYCRecordLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _KYCRecordRegistry.contract.WatchLogs(opts, "zkKYCRecordRevocation", zkKYCRecordLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCRecordRegistryZkKYCRecordRevocation)
				if err := _KYCRecordRegistry.contract.UnpackLog(event, "zkKYCRecordRevocation", log); err != nil {
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

// ParseZkKYCRecordRevocation is a log parse operation binding the contract event 0x61ca9c88d3b99939f9c40240b2a45041aa97b3913eb5633e7b477e8247b0d8d4.
//
// Solidity: event zkKYCRecordRevocation(bytes32 indexed zkKYCRecordLeafHash, address indexed Guardian, uint256 index)
func (_KYCRecordRegistry *KYCRecordRegistryFilterer) ParseZkKYCRecordRevocation(log types.Log) (*KYCRecordRegistryZkKYCRecordRevocation, error) {
	event := new(KYCRecordRegistryZkKYCRecordRevocation)
	if err := _KYCRecordRegistry.contract.UnpackLog(event, "zkKYCRecordRevocation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
