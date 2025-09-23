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

// CertificateData is an auto generated low-level Go binding around an user-defined struct.
type CertificateData struct {
	Guardian   common.Address
	QueueIndex *big.Int
	State      uint8
}

// ZkCertificateRegistryMetaData contains all meta data concerning the ZkCertificateRegistry contract.
var ZkCertificateRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"GuardianRegistry_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"treeDepth_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"zkCertificateLeafHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Guardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumRegistryOperation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"}],\"name\":\"CertificateProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"zkCertificateLeafHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Guardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumRegistryOperation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"}],\"name\":\"OperationQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"ZERO_VALUE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ZkCertificateQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"zkCertificateHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumRegistryOperation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"addOperationToQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentQueuePointer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL1BlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"l1BlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"}],\"name\":\"getMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endIndex\",\"type\":\"uint256\"}],\"name\":\"getMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"zkCertificateHash\",\"type\":\"bytes32\"}],\"name\":\"getQueuePosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getZkCertificateQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardianRegistry\",\"outputs\":[{\"internalType\":\"contractIGuardianRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"hashLeftRight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initBlockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"zkCertificateHash\",\"type\":\"bytes32\"}],\"name\":\"isZkCertificateInTurn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"merkleRootIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRootValidIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRootsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zkCertificateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"processNextOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeDepth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"zkCertificateHashToData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumCertificateState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"zkCertificateHash\",\"type\":\"bytes32\"}],\"name\":\"zkCertificateProcessingData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumCertificateState\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structCertificateData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c0604052600160035534801561001557600080fd5b5060405161247538038061247583398101604081905261003491610493565b338061005a57604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b61006381610180565b50600061006e6101d0565b805490915060ff6801000000000000000082041615906001600160401b031660008115801561009a5750825b90506000826001600160401b031660011480156100b65750303b155b9050811580156100c4575080155b156100e25760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b0319166001178555831561011057845460ff60401b1916680100000000000000001785555b6080879052610120876002610685565b60a05261012d88876101fb565b831561017357845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050610857565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005b92915050565b610203610306565b600161020f828261071a565b50600061025c7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000017fc9211cf8a2ecf2d9ff7e3f783b959c25e2a209cb6cc7b15ff6d264cbc8a296326107d8565b905060005b60805181101561028957610275828061032d565b91506102826001826107fa565b9050610261565b50600280546001808201835560007f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9283018190558354808301909455929091018390558282526004602052604090912055600980546001600160a01b0319166001600160a01b0385161790556102fe6103bf565b600555505050565b61030e61045e565b61032b57604051631afcd79f60e31b815260040160405180910390fd5b565b60408051808201825283815260208101839052905163014cf2b360e51b815260009173__$fd9c4ca1c9b5545e1f17f9e12436e2ab86$__9163299e5660916103779160040161080d565b602060405180830381865af4158015610394573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b8919061083e565b9392505050565b60004661a4b114806103d357504662066eee145b806103df57504661a4ba145b806103ec575046620ce043145b156104595760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610430573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610454919061083e565b905090565b504390565b60006104686101d0565b5468010000000000000000900460ff16919050565b634e487b7160e01b600052604160045260246000fd5b6000806000606084860312156104a857600080fd5b83516001600160a01b03811681146104bf57600080fd5b6020850151604086015191945092506001600160401b038111156104e257600080fd5b8401601f810186136104f357600080fd5b80516001600160401b0381111561050c5761050c61047d565b604051601f8201601f19908116603f011681016001600160401b038111828210171561053a5761053a61047d565b60405281815282820160200188101561055257600080fd5b60005b8281101561057157602081850181015183830182015201610555565b506000602083830101528093505050509250925092565b634e487b7160e01b600052601160045260246000fd5b6001815b60018411156105d9578085048111156105bd576105bd610588565b60018416156105cb57908102905b60019390931c9280026105a2565b935093915050565b6000826105f0575060016101f5565b816105fd575060006101f5565b8160018114610613576002811461061d57610639565b60019150506101f5565b60ff84111561062e5761062e610588565b50506001821b6101f5565b5060208310610133831016604e8410600b841016171561065c575081810a6101f5565b610669600019848461059e565b806000190482111561067d5761067d610588565b029392505050565b60006103b883836105e1565b600181811c908216806106a557607f821691505b6020821081036106c557634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561071557806000526020600020601f840160051c810160208510156106f25750805b601f840160051c820191505b8181101561071257600081556001016106fe565b50505b505050565b81516001600160401b038111156107335761073361047d565b610747816107418454610691565b846106cb565b6020601f82116001811461077b57600083156107635750848201515b600019600385901b1c1916600184901b178455610712565b600084815260208120601f198516915b828110156107ab578785015182556020948501946001909201910161078b565b50848210156107c95786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6000826107f557634e487b7160e01b600052601260045260246000fd5b500690565b808201808211156101f5576101f5610588565b60408101818360005b6002811015610835578151835260209283019290910190600101610816565b50505092915050565b60006020828403121561085057600080fd5b5051919050565b60805160a051611be4610891600039600081816103cf015261145b015260008181610256015281816114ca01526115360152611be46000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c8063715018a611610104578063b0033450116100a2578063e6f5d30811610071578063e6f5d30814610480578063ec732959146104a8578063ef4b4119146104b0578063f2fde38b146104b9576101cf565b8063b00334501461041d578063b0e7f5bc14610426578063b9b3efe91461042f578063d71f152414610435576101cf565b80638d1ddfb1116100de5780638d1ddfb1146103ca5780638da5cb5b146103f1578063a9a8ce3614610402578063ade2ab6214610415576101cf565b8063715018a61461039a57806371c5ecb1146103a25780637284e416146103b5576101cf565b80634232a3c3116101715780634e60f5571161014b5780634e60f5571461033c578063600ed33814610344578063670a1ebb14610364578063685c40dd14610387576101cf565b80634232a3c3146102f457806342895b3d1461031f57806342cbb15c14610334576101cf565b806337b1b263116101ad57806337b1b2631461028e57806338bf282e146102a1578063390f2b65146102b45780633ed9b835146102d4576101cf565b80630820dc2f1461021757806316a56c41146102515780632eb4a7ab14610286575b6101d8306104cc565b6040516020016101e891906115f0565b60408051601f198184030181529082905262461bcd60e51b825261020e9160040161162b565b60405180910390fd5b61023c61022536600461165e565b600090815260046020526040902054600354111590565b60405190151581526020015b60405180910390f35b6102787f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610248565b610278610675565b61027861029c36600461165e565b6106a9565b6102786102af366004611677565b6106ca565b6102786102c236600461165e565b60046020526000908152604090205481565b6102e76102e236600461165e565b61075e565b6040516102489190611699565b600954610307906001600160a01b031681565b6040516001600160a01b039091168152602001610248565b61033261032d3660046116f2565b6107c0565b005b610278610aed565b600254610278565b61035761035236600461165e565b610b8c565b60405161024891906117fe565b61027861037236600461165e565b60009081526008602052604090206001015490565b6102e7610395366004611677565b610c1b565b610332610d83565b6102786103b036600461165e565b610d97565b6103bd610da7565b604051610248919061162b565b6102787f000000000000000000000000000000000000000000000000000000000000000081565b6000546001600160a01b0316610307565b61033261041036600461182c565b610e35565b600654610278565b61027860035481565b61027860075481565b43610278565b61047161044336600461165e565b6008602052600090815260409020805460018201546002909201546001600160a01b03909116919060ff1683565b60405161024893929190611860565b61023c61048e36600461165e565b600754600091825260086020526040909120600101541490565b610278611238565b61027860055481565b6103326104c7366004611884565b611285565b60408051602a808252606082810190935260009190602082018180368337019050509050600360fc1b81600081518110610508576105086118b4565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110610537576105376118b4565b60200101906001600160f81b031916908160001a90535060005b601481101561066e5760006105678260136118e0565b6105729060086118f3565b61057d9060026119f1565b610590906001600160a01b038716611a13565b60f81b9050600060108260f81c6105a79190611a27565b60f81b905060008160f81c60106105be9190611a49565b8360f81c6105cc9190611a65565b60f81b90506105da826112c3565b856105e68660026118f3565b6105f1906002611a7e565b81518110610601576106016118b4565b60200101906001600160f81b031916908160001a905350610621816112c3565b8561062d8660026118f3565b610638906003611a7e565b81518110610648576106486118b4565b60200101906001600160f81b031916908160001a90535050600190920191506105519050565b5092915050565b6002805460009190610689906001906118e0565b81548110610699576106996118b4565b9060005260206000200154905090565b600681815481106106b957600080fd5b600091825260209091200154905081565b60408051808201825283815260208101839052905163014cf2b360e51b815260009173__$fd9c4ca1c9b5545e1f17f9e12436e2ab86$__9163299e56609161071491600401611a91565b602060405180830381865af4158015610731573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107559190611ac2565b90505b92915050565b60025460609082106107b25760405162461bcd60e51b815260206004820152601960248201527f537461727420696e646578206f7574206f6620626f756e647300000000000000604482015260640161020e565b600254610758908390610c1b565b600754600083815260086020526040902060010154146108545760405162461bcd60e51b815260206004820152604360248201527f5a6b436572746966696361746552656769737472793a207a6b4365727469666960448201527f63617465206973206e6f7420696e207475726e20746f2062652070726f6365736064820152621cd95960ea1b608482015260a40161020e565b600160008381526008602052604090206002015460ff16600481111561087c5761087c6117d4565b0361094f5760006108cd7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000017fc9211cf8a2ecf2d9ff7e3f783b959c25e2a209cb6cc7b15ff6d264cbc8a29632611adb565b90506108db848285856112f9565b6000838152600860205260408082206002808201805460ff19169091179055805460019091015491516001600160a01b039091169286927f9cb9155d4b0c129d1b924032efb4ac74da991e0b9970a7d01ba98e18213b7681926109419291908a90611aff565b60405180910390a350610ad0565b600360008381526008602052604090206002015460ff166004811115610977576109776117d4565b03610a4d5760006109c87f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000017fc9211cf8a2ecf2d9ff7e3f783b959c25e2a209cb6cc7b15ff6d264cbc8a29632611adb565b90506109d6848483856112f9565b6002546109e5906001906118e0565b6003556000838152600860205260409081902060028101805460ff19166004179055805460019182015492516001600160a01b039091169286927f9cb9155d4b0c129d1b924032efb4ac74da991e0b9970a7d01ba98e18213b76819261094192908a90611aff565b60405162461bcd60e51b815260206004820152604c60248201527f5a6b436572746966696361746552656769737472793a2070726f63657373696e60448201527f6720696e76616c6964206f7065726174696f6e2e20546869732073686f756c6460648201526b103737ba103430b83832b71760a11b608482015260a40161020e565b600160076000828254610ae39190611a7e565b9091555050505050565b60004661a4b11480610b0157504662066eee145b80610b0d57504661a4ba145b80610b1a575046620ce043145b15610b875760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b829190611ac2565b905090565b504390565b610bad60408051606081018252600080825260208201819052909182015290565b600082815260086020908152604091829020825160608101845281546001600160a01b031681526001820154928101929092526002810154919290919083019060ff166004811115610c0157610c016117d4565b6004811115610c1257610c126117d4565b90525092915050565b6060818310610c7c5760405162461bcd60e51b815260206004820152602760248201527f537461727420696e646578206d757374206265206c657373207468616e20656e6044820152660c840d2dcc8caf60cb1b606482015260840161020e565b600254821115610cce5760405162461bcd60e51b815260206004820152601760248201527f456e6420696e646578206f7574206f6620626f756e6473000000000000000000604482015260640161020e565b6000610cda84846118e0565b905060008167ffffffffffffffff811115610cf757610cf76116dc565b604051908082528060200260200182016040528015610d20578160200160208202803683370190505b50905060005b82811015610d7a576002610d3a8288611a7e565b81548110610d4a57610d4a6118b4565b9060005260206000200154828281518110610d6757610d676118b4565b6020908102919091010152600101610d26565b50949350505050565b610d8b6113c0565b610d9560006113ed565b565b600281815481106106b957600080fd5b60018054610db490611b1e565b80601f0160208091040260200160405190810160405280929190818152602001828054610de090611b1e565b8015610e2d5780601f10610e0257610100808354040283529160200191610e2d565b820191906000526020600020905b815481529060010190602001808311610e1057829003601f168201915b505050505081565b600954604051633af32abf60e01b81523360048201526001600160a01b0390911690633af32abf90602401602060405180830381865afa158015610e7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea19190611b58565b610efb5760405162461bcd60e51b815260206004820152602560248201527f5a6b436572746966696361746552656769737472793a206e6f74206120477561604482015264393234b0b760d91b606482015260840161020e565b6000816001811115610f0f57610f0f6117d4565b03610fcb5760008281526008602052604081206002015460ff166004811115610f3a57610f3a6117d4565b14610fa35760405162461bcd60e51b815260206004820152603360248201527f5a6b436572746966696361746552656769737472793a207a6b436572746966696044820152726361746520616c72656164792065786973747360681b606482015260840161020e565b600082815260086020526040902060020180546001919060ff191682805b021790555061119b565b6001816001811115610fdf57610fdf6117d4565b0361114257600260008381526008602052604090206002015460ff16600481111561100c5761100c6117d4565b1461109b5760405162461bcd60e51b815260206004820152605360248201527f5a6b436572746966696361746552656769737472793a207a6b4365727469666960448201527f63617465206973206e6f742069737375656420616e642063616e207468657265606482015272199bdc99481b9bdd081899481c995d9bdad959606a1b608482015260a40161020e565b6000828152600860205260409020546001600160a01b0316331461111f5760405162461bcd60e51b815260206004820152603560248201527f5a6b436572746966696361746552656769737472793a206e6f7420746865206360448201527437b93932b9b837b73234b7339023bab0b93234b0b760591b606482015260840161020e565b600082815260086020526040902060020180546003919060ff1916600183610fc1565b60405162461bcd60e51b815260206004820152602860248201527f5a6b436572746966696361746552656769737472793a20696e76616c6964206f6044820152673832b930ba34b7b760c11b606482015260840161020e565b6006805460008481526008602052604090819020600180820193845581546001600160a01b0319163390811790925584549081019094557ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f9093018590559054905184917fb1acf3fe87405cb3d477043a120b6933cf2364d4826ab90b112545cecc628e8f9161122c918691611b7a565b60405180910390a35050565b6112827f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000017fc9211cf8a2ecf2d9ff7e3f783b959c25e2a209cb6cc7b15ff6d264cbc8a29632611adb565b81565b61128d6113c0565b6001600160a01b0381166112b757604051631e4fbdf760e01b81526000600482015260240161020e565b6112c0816113ed565b50565b6000600a60f883901c10156112ea576112e160f883901c6030611b95565b60f81b92915050565b6112e160f883901c6057611b95565b61130c818585611307610675565b61143d565b6113585760405162461bcd60e51b815260206004820152601960248201527f6d65726b6c652070726f6f66206973206e6f742076616c696400000000000000604482015260640161020e565b6000611365828685611457565b600280546001818101835560008390527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace90910183905590549192506113aa916118e0565b6000918252600460205260409091205550505050565b6000546001600160a01b03163314610d955760405163118cdaa760e01b815233600482015260240161020e565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60008161144b868686611457565b1490505b949350505050565b60007f000000000000000000000000000000000000000000000000000000000000000083106114c85760405162461bcd60e51b815260206004820152601c60248201527f5f696e64657820626967676572207468616e20747265652073697a6500000000604482015260640161020e565b7f00000000000000000000000000000000000000000000000000000000000000008451146115315760405162461bcd60e51b8152602060048201526016602482015275092dcecc2d8d2c840bee0e4dedecce640d8cadccee8d60531b604482015260640161020e565b60005b7f00000000000000000000000000000000000000000000000000000000000000008110156115c357836001166001036115915761158a85828151811061157c5761157c6118b4565b6020026020010151846106ca565b92506115b7565b6115b4838683815181106115a7576115a76118b4565b60200260200101516106ca565b92505b600193841c9301611534565b50909392505050565b60005b838110156115e75781810151838201526020016115cf565b50506000910152565b7203ab739bab83837b93a32b21036b2ba3437b21606d1b81526000825161161e8160138501602087016115cc565b9190910160130192915050565b602081526000825180602084015261164a8160408501602087016115cc565b601f01601f19169190910160400192915050565b60006020828403121561167057600080fd5b5035919050565b6000806040838503121561168a57600080fd5b50508035926020909101359150565b602080825282518282018190526000918401906040840190835b818110156116d15783518352602093840193909201916001016116b3565b509095945050505050565b634e487b7160e01b600052604160045260246000fd5b60008060006060848603121561170757600080fd5b8335925060208401359150604084013567ffffffffffffffff81111561172c57600080fd5b8401601f8101861361173d57600080fd5b803567ffffffffffffffff811115611757576117576116dc565b8060051b604051601f19603f830116810181811067ffffffffffffffff82111715611784576117846116dc565b6040529182526020818401810192908101898411156117a257600080fd5b6020850194505b838510156117c5578435808252602095860195909350016117a9565b50809450505050509250925092565b634e487b7160e01b600052602160045260246000fd5b600581106117fa576117fa6117d4565b9052565b81516001600160a01b0316815260208083015190820152604080830151606083019161066e908401826117ea565b6000806040838503121561183f57600080fd5b8235915060208301356002811061185557600080fd5b809150509250929050565b6001600160a01b0384168152602081018390526060810161144f60408301846117ea565b60006020828403121561189657600080fd5b81356001600160a01b03811681146118ad57600080fd5b9392505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b81810381811115610758576107586118ca565b8082028115828204841417610758576107586118ca565b6001815b600184111561194557808504811115611929576119296118ca565b600184161561193757908102905b60019390931c92800261190e565b935093915050565b60008261195c57506001610758565b8161196957506000610758565b816001811461197f5760028114611989576119a5565b6001915050610758565b60ff84111561199a5761199a6118ca565b50506001821b610758565b5060208310610133831016604e8410600b84101617156119c8575081810a610758565b6119d5600019848461190a565b80600019048211156119e9576119e96118ca565b029392505050565b6000610755838361194d565b634e487b7160e01b600052601260045260246000fd5b600082611a2257611a226119fd565b500490565b600060ff831680611a3a57611a3a6119fd565b8060ff84160491505092915050565b60ff818116838216029081169081811461066e5761066e6118ca565b60ff8281168282160390811115610758576107586118ca565b80820180821115610758576107586118ca565b60408101818360005b6002811015611ab9578151835260209283019290910190600101611a9a565b50505092915050565b600060208284031215611ad457600080fd5b5051919050565b600082611aea57611aea6119fd565b500690565b600281106117fa576117fa6117d4565b60608101611b0d8286611aef565b602082019390935260400152919050565b600181811c90821680611b3257607f821691505b602082108103611b5257634e487b7160e01b600052602260045260246000fd5b50919050565b600060208284031215611b6a57600080fd5b815180151581146118ad57600080fd5b60408101611b888285611aef565b8260208301529392505050565b60ff8181168382160190811115610758576107586118ca56fea2646970667358221220c092ff1a4fadf596cbba92c6c003b01e9984adbee17ff3b7304e4f453e41068664736f6c634300081c0033",
}

// ZkCertificateRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkCertificateRegistryMetaData.ABI instead.
var ZkCertificateRegistryABI = ZkCertificateRegistryMetaData.ABI

// ZkCertificateRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZkCertificateRegistryMetaData.Bin instead.
var ZkCertificateRegistryBin = ZkCertificateRegistryMetaData.Bin

// DeployZkCertificateRegistry deploys a new Ethereum contract, binding an instance of ZkCertificateRegistry to it.
func DeployZkCertificateRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, GuardianRegistry_ common.Address, treeDepth_ *big.Int, description_ string) (common.Address, *types.Transaction, *ZkCertificateRegistry, error) {
	parsed, err := ZkCertificateRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZkCertificateRegistryBin), backend, GuardianRegistry_, treeDepth_, description_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZkCertificateRegistry{ZkCertificateRegistryCaller: ZkCertificateRegistryCaller{contract: contract}, ZkCertificateRegistryTransactor: ZkCertificateRegistryTransactor{contract: contract}, ZkCertificateRegistryFilterer: ZkCertificateRegistryFilterer{contract: contract}}, nil
}

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

// ZkCertificateQueue is a free data retrieval call binding the contract method 0x37b1b263.
//
// Solidity: function ZkCertificateQueue(uint256 ) view returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) ZkCertificateQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "ZkCertificateQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZkCertificateQueue is a free data retrieval call binding the contract method 0x37b1b263.
//
// Solidity: function ZkCertificateQueue(uint256 ) view returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) ZkCertificateQueue(arg0 *big.Int) ([32]byte, error) {
	return _ZkCertificateRegistry.Contract.ZkCertificateQueue(&_ZkCertificateRegistry.CallOpts, arg0)
}

// ZkCertificateQueue is a free data retrieval call binding the contract method 0x37b1b263.
//
// Solidity: function ZkCertificateQueue(uint256 ) view returns(bytes32)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) ZkCertificateQueue(arg0 *big.Int) ([32]byte, error) {
	return _ZkCertificateRegistry.Contract.ZkCertificateQueue(&_ZkCertificateRegistry.CallOpts, arg0)
}

// CurrentQueuePointer is a free data retrieval call binding the contract method 0xb0e7f5bc.
//
// Solidity: function currentQueuePointer() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) CurrentQueuePointer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "currentQueuePointer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentQueuePointer is a free data retrieval call binding the contract method 0xb0e7f5bc.
//
// Solidity: function currentQueuePointer() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) CurrentQueuePointer() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.CurrentQueuePointer(&_ZkCertificateRegistry.CallOpts)
}

// CurrentQueuePointer is a free data retrieval call binding the contract method 0xb0e7f5bc.
//
// Solidity: function currentQueuePointer() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) CurrentQueuePointer() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.CurrentQueuePointer(&_ZkCertificateRegistry.CallOpts)
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

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) GetBlockNumber() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.GetBlockNumber(&_ZkCertificateRegistry.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) GetBlockNumber() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.GetBlockNumber(&_ZkCertificateRegistry.CallOpts)
}

// GetL1BlockNumber is a free data retrieval call binding the contract method 0xb9b3efe9.
//
// Solidity: function getL1BlockNumber() view returns(uint256 l1BlockNumber)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) GetL1BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "getL1BlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL1BlockNumber is a free data retrieval call binding the contract method 0xb9b3efe9.
//
// Solidity: function getL1BlockNumber() view returns(uint256 l1BlockNumber)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) GetL1BlockNumber() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.GetL1BlockNumber(&_ZkCertificateRegistry.CallOpts)
}

// GetL1BlockNumber is a free data retrieval call binding the contract method 0xb9b3efe9.
//
// Solidity: function getL1BlockNumber() view returns(uint256 l1BlockNumber)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) GetL1BlockNumber() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.GetL1BlockNumber(&_ZkCertificateRegistry.CallOpts)
}

// GetMerkleRoots is a free data retrieval call binding the contract method 0x3ed9b835.
//
// Solidity: function getMerkleRoots(uint256 _startIndex) view returns(bytes32[])
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) GetMerkleRoots(opts *bind.CallOpts, _startIndex *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "getMerkleRoots", _startIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetMerkleRoots is a free data retrieval call binding the contract method 0x3ed9b835.
//
// Solidity: function getMerkleRoots(uint256 _startIndex) view returns(bytes32[])
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) GetMerkleRoots(_startIndex *big.Int) ([][32]byte, error) {
	return _ZkCertificateRegistry.Contract.GetMerkleRoots(&_ZkCertificateRegistry.CallOpts, _startIndex)
}

// GetMerkleRoots is a free data retrieval call binding the contract method 0x3ed9b835.
//
// Solidity: function getMerkleRoots(uint256 _startIndex) view returns(bytes32[])
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) GetMerkleRoots(_startIndex *big.Int) ([][32]byte, error) {
	return _ZkCertificateRegistry.Contract.GetMerkleRoots(&_ZkCertificateRegistry.CallOpts, _startIndex)
}

// GetMerkleRoots0 is a free data retrieval call binding the contract method 0x685c40dd.
//
// Solidity: function getMerkleRoots(uint256 _startIndex, uint256 _endIndex) view returns(bytes32[])
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) GetMerkleRoots0(opts *bind.CallOpts, _startIndex *big.Int, _endIndex *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "getMerkleRoots0", _startIndex, _endIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetMerkleRoots0 is a free data retrieval call binding the contract method 0x685c40dd.
//
// Solidity: function getMerkleRoots(uint256 _startIndex, uint256 _endIndex) view returns(bytes32[])
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) GetMerkleRoots0(_startIndex *big.Int, _endIndex *big.Int) ([][32]byte, error) {
	return _ZkCertificateRegistry.Contract.GetMerkleRoots0(&_ZkCertificateRegistry.CallOpts, _startIndex, _endIndex)
}

// GetMerkleRoots0 is a free data retrieval call binding the contract method 0x685c40dd.
//
// Solidity: function getMerkleRoots(uint256 _startIndex, uint256 _endIndex) view returns(bytes32[])
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) GetMerkleRoots0(_startIndex *big.Int, _endIndex *big.Int) ([][32]byte, error) {
	return _ZkCertificateRegistry.Contract.GetMerkleRoots0(&_ZkCertificateRegistry.CallOpts, _startIndex, _endIndex)
}

// GetQueuePosition is a free data retrieval call binding the contract method 0x670a1ebb.
//
// Solidity: function getQueuePosition(bytes32 zkCertificateHash) view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) GetQueuePosition(opts *bind.CallOpts, zkCertificateHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "getQueuePosition", zkCertificateHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQueuePosition is a free data retrieval call binding the contract method 0x670a1ebb.
//
// Solidity: function getQueuePosition(bytes32 zkCertificateHash) view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) GetQueuePosition(zkCertificateHash [32]byte) (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.GetQueuePosition(&_ZkCertificateRegistry.CallOpts, zkCertificateHash)
}

// GetQueuePosition is a free data retrieval call binding the contract method 0x670a1ebb.
//
// Solidity: function getQueuePosition(bytes32 zkCertificateHash) view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) GetQueuePosition(zkCertificateHash [32]byte) (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.GetQueuePosition(&_ZkCertificateRegistry.CallOpts, zkCertificateHash)
}

// GetZkCertificateQueueLength is a free data retrieval call binding the contract method 0xade2ab62.
//
// Solidity: function getZkCertificateQueueLength() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) GetZkCertificateQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "getZkCertificateQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetZkCertificateQueueLength is a free data retrieval call binding the contract method 0xade2ab62.
//
// Solidity: function getZkCertificateQueueLength() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) GetZkCertificateQueueLength() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.GetZkCertificateQueueLength(&_ZkCertificateRegistry.CallOpts)
}

// GetZkCertificateQueueLength is a free data retrieval call binding the contract method 0xade2ab62.
//
// Solidity: function getZkCertificateQueueLength() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) GetZkCertificateQueueLength() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.GetZkCertificateQueueLength(&_ZkCertificateRegistry.CallOpts)
}

// GuardianRegistry is a free data retrieval call binding the contract method 0x4232a3c3.
//
// Solidity: function guardianRegistry() view returns(address)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) GuardianRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "guardianRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GuardianRegistry is a free data retrieval call binding the contract method 0x4232a3c3.
//
// Solidity: function guardianRegistry() view returns(address)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) GuardianRegistry() (common.Address, error) {
	return _ZkCertificateRegistry.Contract.GuardianRegistry(&_ZkCertificateRegistry.CallOpts)
}

// GuardianRegistry is a free data retrieval call binding the contract method 0x4232a3c3.
//
// Solidity: function guardianRegistry() view returns(address)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) GuardianRegistry() (common.Address, error) {
	return _ZkCertificateRegistry.Contract.GuardianRegistry(&_ZkCertificateRegistry.CallOpts)
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

// IsZkCertificateInTurn is a free data retrieval call binding the contract method 0xe6f5d308.
//
// Solidity: function isZkCertificateInTurn(bytes32 zkCertificateHash) view returns(bool)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) IsZkCertificateInTurn(opts *bind.CallOpts, zkCertificateHash [32]byte) (bool, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "isZkCertificateInTurn", zkCertificateHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsZkCertificateInTurn is a free data retrieval call binding the contract method 0xe6f5d308.
//
// Solidity: function isZkCertificateInTurn(bytes32 zkCertificateHash) view returns(bool)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) IsZkCertificateInTurn(zkCertificateHash [32]byte) (bool, error) {
	return _ZkCertificateRegistry.Contract.IsZkCertificateInTurn(&_ZkCertificateRegistry.CallOpts, zkCertificateHash)
}

// IsZkCertificateInTurn is a free data retrieval call binding the contract method 0xe6f5d308.
//
// Solidity: function isZkCertificateInTurn(bytes32 zkCertificateHash) view returns(bool)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) IsZkCertificateInTurn(zkCertificateHash [32]byte) (bool, error) {
	return _ZkCertificateRegistry.Contract.IsZkCertificateInTurn(&_ZkCertificateRegistry.CallOpts, zkCertificateHash)
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

// MerkleRootsLength is a free data retrieval call binding the contract method 0x4e60f557.
//
// Solidity: function merkleRootsLength() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) MerkleRootsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "merkleRootsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MerkleRootsLength is a free data retrieval call binding the contract method 0x4e60f557.
//
// Solidity: function merkleRootsLength() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) MerkleRootsLength() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.MerkleRootsLength(&_ZkCertificateRegistry.CallOpts)
}

// MerkleRootsLength is a free data retrieval call binding the contract method 0x4e60f557.
//
// Solidity: function merkleRootsLength() view returns(uint256)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) MerkleRootsLength() (*big.Int, error) {
	return _ZkCertificateRegistry.Contract.MerkleRootsLength(&_ZkCertificateRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) Owner() (common.Address, error) {
	return _ZkCertificateRegistry.Contract.Owner(&_ZkCertificateRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) Owner() (common.Address, error) {
	return _ZkCertificateRegistry.Contract.Owner(&_ZkCertificateRegistry.CallOpts)
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

// ZkCertificateHashToData is a free data retrieval call binding the contract method 0xd71f1524.
//
// Solidity: function zkCertificateHashToData(bytes32 ) view returns(address guardian, uint256 queueIndex, uint8 state)
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) ZkCertificateHashToData(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Guardian   common.Address
	QueueIndex *big.Int
	State      uint8
}, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "zkCertificateHashToData", arg0)

	outstruct := new(struct {
		Guardian   common.Address
		QueueIndex *big.Int
		State      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Guardian = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.QueueIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.State = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// ZkCertificateHashToData is a free data retrieval call binding the contract method 0xd71f1524.
//
// Solidity: function zkCertificateHashToData(bytes32 ) view returns(address guardian, uint256 queueIndex, uint8 state)
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) ZkCertificateHashToData(arg0 [32]byte) (struct {
	Guardian   common.Address
	QueueIndex *big.Int
	State      uint8
}, error) {
	return _ZkCertificateRegistry.Contract.ZkCertificateHashToData(&_ZkCertificateRegistry.CallOpts, arg0)
}

// ZkCertificateHashToData is a free data retrieval call binding the contract method 0xd71f1524.
//
// Solidity: function zkCertificateHashToData(bytes32 ) view returns(address guardian, uint256 queueIndex, uint8 state)
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) ZkCertificateHashToData(arg0 [32]byte) (struct {
	Guardian   common.Address
	QueueIndex *big.Int
	State      uint8
}, error) {
	return _ZkCertificateRegistry.Contract.ZkCertificateHashToData(&_ZkCertificateRegistry.CallOpts, arg0)
}

// ZkCertificateProcessingData is a free data retrieval call binding the contract method 0x600ed338.
//
// Solidity: function zkCertificateProcessingData(bytes32 zkCertificateHash) view returns((address,uint256,uint8))
func (_ZkCertificateRegistry *ZkCertificateRegistryCaller) ZkCertificateProcessingData(opts *bind.CallOpts, zkCertificateHash [32]byte) (KycCertificateData, error) {
	var out []interface{}
	err := _ZkCertificateRegistry.contract.Call(opts, &out, "zkCertificateProcessingData", zkCertificateHash)

	if err != nil {
		return *new(KycCertificateData), err
	}

	out0 := *abi.ConvertType(out[0], new(KycCertificateData)).(*KycCertificateData)

	return out0, err

}

// ZkCertificateProcessingData is a free data retrieval call binding the contract method 0x600ed338.
//
// Solidity: function zkCertificateProcessingData(bytes32 zkCertificateHash) view returns((address,uint256,uint8))
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) ZkCertificateProcessingData(zkCertificateHash [32]byte) (KycCertificateData, error) {
	return _ZkCertificateRegistry.Contract.ZkCertificateProcessingData(&_ZkCertificateRegistry.CallOpts, zkCertificateHash)
}

// ZkCertificateProcessingData is a free data retrieval call binding the contract method 0x600ed338.
//
// Solidity: function zkCertificateProcessingData(bytes32 zkCertificateHash) view returns((address,uint256,uint8))
func (_ZkCertificateRegistry *ZkCertificateRegistryCallerSession) ZkCertificateProcessingData(zkCertificateHash [32]byte) (KycCertificateData, error) {
	return _ZkCertificateRegistry.Contract.ZkCertificateProcessingData(&_ZkCertificateRegistry.CallOpts, zkCertificateHash)
}

// AddOperationToQueue is a paid mutator transaction binding the contract method 0xa9a8ce36.
//
// Solidity: function addOperationToQueue(bytes32 zkCertificateHash, uint8 operation) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactor) AddOperationToQueue(opts *bind.TransactOpts, zkCertificateHash [32]byte, operation uint8) (*types.Transaction, error) {
	return _ZkCertificateRegistry.contract.Transact(opts, "addOperationToQueue", zkCertificateHash, operation)
}

// AddOperationToQueue is a paid mutator transaction binding the contract method 0xa9a8ce36.
//
// Solidity: function addOperationToQueue(bytes32 zkCertificateHash, uint8 operation) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) AddOperationToQueue(zkCertificateHash [32]byte, operation uint8) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.AddOperationToQueue(&_ZkCertificateRegistry.TransactOpts, zkCertificateHash, operation)
}

// AddOperationToQueue is a paid mutator transaction binding the contract method 0xa9a8ce36.
//
// Solidity: function addOperationToQueue(bytes32 zkCertificateHash, uint8 operation) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactorSession) AddOperationToQueue(zkCertificateHash [32]byte, operation uint8) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.AddOperationToQueue(&_ZkCertificateRegistry.TransactOpts, zkCertificateHash, operation)
}

// ProcessNextOperation is a paid mutator transaction binding the contract method 0x42895b3d.
//
// Solidity: function processNextOperation(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactor) ProcessNextOperation(opts *bind.TransactOpts, leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkCertificateRegistry.contract.Transact(opts, "processNextOperation", leafIndex, zkCertificateHash, merkleProof)
}

// ProcessNextOperation is a paid mutator transaction binding the contract method 0x42895b3d.
//
// Solidity: function processNextOperation(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) ProcessNextOperation(leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.ProcessNextOperation(&_ZkCertificateRegistry.TransactOpts, leafIndex, zkCertificateHash, merkleProof)
}

// ProcessNextOperation is a paid mutator transaction binding the contract method 0x42895b3d.
//
// Solidity: function processNextOperation(uint256 leafIndex, bytes32 zkCertificateHash, bytes32[] merkleProof) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactorSession) ProcessNextOperation(leafIndex *big.Int, zkCertificateHash [32]byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.ProcessNextOperation(&_ZkCertificateRegistry.TransactOpts, leafIndex, zkCertificateHash, merkleProof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkCertificateRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.RenounceOwnership(&_ZkCertificateRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.RenounceOwnership(&_ZkCertificateRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZkCertificateRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.TransferOwnership(&_ZkCertificateRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.TransferOwnership(&_ZkCertificateRegistry.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ZkCertificateRegistry.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ZkCertificateRegistry *ZkCertificateRegistrySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.Fallback(&_ZkCertificateRegistry.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ZkCertificateRegistry *ZkCertificateRegistryTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ZkCertificateRegistry.Contract.Fallback(&_ZkCertificateRegistry.TransactOpts, calldata)
}

// ZkCertificateRegistryCertificateProcessedIterator is returned from FilterCertificateProcessed and is used to iterate over the raw logs and unpacked data for CertificateProcessed events raised by the ZkCertificateRegistry contract.
type ZkCertificateRegistryCertificateProcessedIterator struct {
	Event *ZkCertificateRegistryCertificateProcessed // Event containing the contract specifics and raw log

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
func (it *ZkCertificateRegistryCertificateProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkCertificateRegistryCertificateProcessed)
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
		it.Event = new(ZkCertificateRegistryCertificateProcessed)
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
func (it *ZkCertificateRegistryCertificateProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkCertificateRegistryCertificateProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkCertificateRegistryCertificateProcessed represents a CertificateProcessed event raised by the ZkCertificateRegistry contract.
type ZkCertificateRegistryCertificateProcessed struct {
	ZkCertificateLeafHash [32]byte
	Guardian              common.Address
	Operation             uint8
	QueueIndex            *big.Int
	LeafIndex             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCertificateProcessed is a free log retrieval operation binding the contract event 0x9cb9155d4b0c129d1b924032efb4ac74da991e0b9970a7d01ba98e18213b7681.
//
// Solidity: event CertificateProcessed(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint8 operation, uint256 queueIndex, uint256 leafIndex)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) FilterCertificateProcessed(opts *bind.FilterOpts, zkCertificateLeafHash [][32]byte, Guardian []common.Address) (*ZkCertificateRegistryCertificateProcessedIterator, error) {

	var zkCertificateLeafHashRule []interface{}
	for _, zkCertificateLeafHashItem := range zkCertificateLeafHash {
		zkCertificateLeafHashRule = append(zkCertificateLeafHashRule, zkCertificateLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _ZkCertificateRegistry.contract.FilterLogs(opts, "CertificateProcessed", zkCertificateLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return &ZkCertificateRegistryCertificateProcessedIterator{contract: _ZkCertificateRegistry.contract, event: "CertificateProcessed", logs: logs, sub: sub}, nil
}

// WatchCertificateProcessed is a free log subscription operation binding the contract event 0x9cb9155d4b0c129d1b924032efb4ac74da991e0b9970a7d01ba98e18213b7681.
//
// Solidity: event CertificateProcessed(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint8 operation, uint256 queueIndex, uint256 leafIndex)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) WatchCertificateProcessed(opts *bind.WatchOpts, sink chan<- *ZkCertificateRegistryCertificateProcessed, zkCertificateLeafHash [][32]byte, Guardian []common.Address) (event.Subscription, error) {

	var zkCertificateLeafHashRule []interface{}
	for _, zkCertificateLeafHashItem := range zkCertificateLeafHash {
		zkCertificateLeafHashRule = append(zkCertificateLeafHashRule, zkCertificateLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _ZkCertificateRegistry.contract.WatchLogs(opts, "CertificateProcessed", zkCertificateLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkCertificateRegistryCertificateProcessed)
				if err := _ZkCertificateRegistry.contract.UnpackLog(event, "CertificateProcessed", log); err != nil {
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

// ParseCertificateProcessed is a log parse operation binding the contract event 0x9cb9155d4b0c129d1b924032efb4ac74da991e0b9970a7d01ba98e18213b7681.
//
// Solidity: event CertificateProcessed(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint8 operation, uint256 queueIndex, uint256 leafIndex)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) ParseCertificateProcessed(log types.Log) (*ZkCertificateRegistryCertificateProcessed, error) {
	event := new(ZkCertificateRegistryCertificateProcessed)
	if err := _ZkCertificateRegistry.contract.UnpackLog(event, "CertificateProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZkCertificateRegistryInitializedIterator, error) {

	logs, sub, err := _ZkCertificateRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZkCertificateRegistryInitializedIterator{contract: _ZkCertificateRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) ParseInitialized(log types.Log) (*ZkCertificateRegistryInitialized, error) {
	event := new(ZkCertificateRegistryInitialized)
	if err := _ZkCertificateRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkCertificateRegistryOperationQueuedIterator is returned from FilterOperationQueued and is used to iterate over the raw logs and unpacked data for OperationQueued events raised by the ZkCertificateRegistry contract.
type ZkCertificateRegistryOperationQueuedIterator struct {
	Event *ZkCertificateRegistryOperationQueued // Event containing the contract specifics and raw log

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
func (it *ZkCertificateRegistryOperationQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkCertificateRegistryOperationQueued)
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
		it.Event = new(ZkCertificateRegistryOperationQueued)
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
func (it *ZkCertificateRegistryOperationQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkCertificateRegistryOperationQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkCertificateRegistryOperationQueued represents a OperationQueued event raised by the ZkCertificateRegistry contract.
type ZkCertificateRegistryOperationQueued struct {
	ZkCertificateLeafHash [32]byte
	Guardian              common.Address
	Operation             uint8
	QueueIndex            *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterOperationQueued is a free log retrieval operation binding the contract event 0xb1acf3fe87405cb3d477043a120b6933cf2364d4826ab90b112545cecc628e8f.
//
// Solidity: event OperationQueued(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint8 operation, uint256 queueIndex)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) FilterOperationQueued(opts *bind.FilterOpts, zkCertificateLeafHash [][32]byte, Guardian []common.Address) (*ZkCertificateRegistryOperationQueuedIterator, error) {

	var zkCertificateLeafHashRule []interface{}
	for _, zkCertificateLeafHashItem := range zkCertificateLeafHash {
		zkCertificateLeafHashRule = append(zkCertificateLeafHashRule, zkCertificateLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _ZkCertificateRegistry.contract.FilterLogs(opts, "OperationQueued", zkCertificateLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return &ZkCertificateRegistryOperationQueuedIterator{contract: _ZkCertificateRegistry.contract, event: "OperationQueued", logs: logs, sub: sub}, nil
}

// WatchOperationQueued is a free log subscription operation binding the contract event 0xb1acf3fe87405cb3d477043a120b6933cf2364d4826ab90b112545cecc628e8f.
//
// Solidity: event OperationQueued(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint8 operation, uint256 queueIndex)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) WatchOperationQueued(opts *bind.WatchOpts, sink chan<- *ZkCertificateRegistryOperationQueued, zkCertificateLeafHash [][32]byte, Guardian []common.Address) (event.Subscription, error) {

	var zkCertificateLeafHashRule []interface{}
	for _, zkCertificateLeafHashItem := range zkCertificateLeafHash {
		zkCertificateLeafHashRule = append(zkCertificateLeafHashRule, zkCertificateLeafHashItem)
	}
	var GuardianRule []interface{}
	for _, GuardianItem := range Guardian {
		GuardianRule = append(GuardianRule, GuardianItem)
	}

	logs, sub, err := _ZkCertificateRegistry.contract.WatchLogs(opts, "OperationQueued", zkCertificateLeafHashRule, GuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkCertificateRegistryOperationQueued)
				if err := _ZkCertificateRegistry.contract.UnpackLog(event, "OperationQueued", log); err != nil {
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

// ParseOperationQueued is a log parse operation binding the contract event 0xb1acf3fe87405cb3d477043a120b6933cf2364d4826ab90b112545cecc628e8f.
//
// Solidity: event OperationQueued(bytes32 indexed zkCertificateLeafHash, address indexed Guardian, uint8 operation, uint256 queueIndex)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) ParseOperationQueued(log types.Log) (*ZkCertificateRegistryOperationQueued, error) {
	event := new(ZkCertificateRegistryOperationQueued)
	if err := _ZkCertificateRegistry.contract.UnpackLog(event, "OperationQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkCertificateRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZkCertificateRegistry contract.
type ZkCertificateRegistryOwnershipTransferredIterator struct {
	Event *ZkCertificateRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZkCertificateRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkCertificateRegistryOwnershipTransferred)
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
		it.Event = new(ZkCertificateRegistryOwnershipTransferred)
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
func (it *ZkCertificateRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkCertificateRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkCertificateRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ZkCertificateRegistry contract.
type ZkCertificateRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZkCertificateRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZkCertificateRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZkCertificateRegistryOwnershipTransferredIterator{contract: _ZkCertificateRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZkCertificateRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZkCertificateRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkCertificateRegistryOwnershipTransferred)
				if err := _ZkCertificateRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZkCertificateRegistry *ZkCertificateRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ZkCertificateRegistryOwnershipTransferred, error) {
	event := new(ZkCertificateRegistryOwnershipTransferred)
	if err := _ZkCertificateRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
