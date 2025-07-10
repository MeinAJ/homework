// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package go_contract

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

// BeggingContractMetaData contains all meta data concerning the BeggingContract contract.
var BeggingContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"donate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"donationMapping\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"donations\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"donor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDonation\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalDonations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DonationReceived\",\"inputs\":[{\"name\":\"donor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f5ffd5b50335f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610081575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161007891906101b3565b60405180910390fd5b610090816100b360201b60201c565b50426001819055506312cc0300426100a89190610202565b600281905550610235565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61019d82610174565b9050919050565b6101ad81610193565b82525050565b5f6020820190506101c65f8301846101a4565b92915050565b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61020c826101cc565b9150610217836101cc565b925082820190508082111561022f5761022e6101d5565b5b92915050565b610cc2806102425f395ff3fe608060405260043610610085575f3560e01c8063c184a6ab11610058578063c184a6ab1461011b578063de2ed89314610157578063ed88c68e14610181578063f2fde38b1461018b578063f8626af8146101b357610085565b80633ccfd60b14610089578063410a1d321461009f578063715018a6146100db5780638da5cb5b146100f1575b5f5ffd5b348015610094575f5ffd5b5061009d6101f1565b005b3480156100aa575f5ffd5b506100c560048036038101906100c09190610852565b6102f2565b6040516100d29190610895565b60405180910390f35b3480156100e6575f5ffd5b506100ef610338565b005b3480156100fc575f5ffd5b5061010561034b565b60405161011291906108bd565b60405180910390f35b348015610126575f5ffd5b50610141600480360381019061013c9190610852565b610372565b60405161014e9190610895565b60405180910390f35b348015610162575f5ffd5b5061016b610387565b6040516101789190610895565b60405180910390f35b61018961038e565b005b348015610196575f5ffd5b506101b160048036038101906101ac9190610852565b6104fb565b005b3480156101be575f5ffd5b506101d960048036038101906101d49190610900565b61057f565b6040516101e89392919061092b565b60405180910390f35b6101f96105d3565b5f4790505f811161023f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610236906109ba565b60405180910390fd5b5f61024861034b565b73ffffffffffffffffffffffffffffffffffffffff168260405161026b90610a05565b5f6040518083038185875af1925050503d805f81146102a5576040519150601f19603f3d011682016040523d82523d5f602084013e6102aa565b606091505b50509050806102ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e590610a63565b60405180910390fd5b5050565b5f60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6103406105d3565b6103495f61065a565b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6003602052805f5260405f205f915090505481565b5f47905090565b61039661071b565b61039e6107a7565b5f60405180606001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001348152602001428152509050600481908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015550503460035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546104a19190610aae565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f4b0304ab2f313234b1cea012339c8f8c3bf2c8bb357d81a579cfdb35d2e3d03034426040516104f0929190610ae1565b60405180910390a250565b6105036105d3565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610573575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161056a91906108bd565b60405180910390fd5b61057c8161065a565b50565b6004818154811061058e575f80fd5b905f5260205f2090600302015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154905083565b6105db6107ed565b73ffffffffffffffffffffffffffffffffffffffff166105f961034b565b73ffffffffffffffffffffffffffffffffffffffff16146106585761061c6107ed565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161064f91906108bd565b60405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600154421015610760576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075790610b78565b60405180910390fd5b6002544211156107a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079c90610be0565b60405180910390fd5b565b60013410156107eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e290610c6e565b60405180910390fd5b565b5f33905090565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610821826107f8565b9050919050565b61083181610817565b811461083b575f5ffd5b50565b5f8135905061084c81610828565b92915050565b5f60208284031215610867576108666107f4565b5b5f6108748482850161083e565b91505092915050565b5f819050919050565b61088f8161087d565b82525050565b5f6020820190506108a85f830184610886565b92915050565b6108b781610817565b82525050565b5f6020820190506108d05f8301846108ae565b92915050565b6108df8161087d565b81146108e9575f5ffd5b50565b5f813590506108fa816108d6565b92915050565b5f60208284031215610915576109146107f4565b5b5f610922848285016108ec565b91505092915050565b5f60608201905061093e5f8301866108ae565b61094b6020830185610886565b6109586040830184610886565b949350505050565b5f82825260208201905092915050565b7f4e6f2066756e647320746f2077697468647261770000000000000000000000005f82015250565b5f6109a4601483610960565b91506109af82610970565b602082019050919050565b5f6020820190508181035f8301526109d181610998565b9050919050565b5f81905092915050565b50565b5f6109f05f836109d8565b91506109fb826109e2565b5f82019050919050565b5f610a0f826109e5565b9150819050919050565b7f5472616e73666572206661696c656400000000000000000000000000000000005f82015250565b5f610a4d600f83610960565b9150610a5882610a19565b602082019050919050565b5f6020820190508181035f830152610a7a81610a41565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610ab88261087d565b9150610ac38361087d565b9250828201905080821115610adb57610ada610a81565b5b92915050565b5f604082019050610af45f830185610886565b610b016020830184610886565b9392505050565b7f42656767696e67436f6e74726163743a20646f6e6174696f6e206e6f742073745f8201527f6172746564207965740000000000000000000000000000000000000000000000602082015250565b5f610b62602983610960565b9150610b6d82610b08565b604082019050919050565b5f6020820190508181035f830152610b8f81610b56565b9050919050565b7f42656767696e67436f6e74726163743a20646f6e6174696f6e20656e646564005f82015250565b5f610bca601f83610960565b9150610bd582610b96565b602082019050919050565b5f6020820190508181035f830152610bf781610bbe565b9050919050565b7f42656767696e67436f6e74726163743a20646f6e6174696f6e20616d6f756e745f8201527f2073686f756c64206265206174206c6561737420312077656900000000000000602082015250565b5f610c58603983610960565b9150610c6382610bfe565b604082019050919050565b5f6020820190508181035f830152610c8581610c4c565b905091905056fea264697066735822122001592e50f9e6ea864cadb8cdb2ce14a38a2c25ba4d2f76d5fee8fefb0093072e64736f6c634300081c0033",
}

// BeggingContractABI is the input ABI used to generate the binding from.
// Deprecated: Use BeggingContractMetaData.ABI instead.
var BeggingContractABI = BeggingContractMetaData.ABI

// BeggingContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BeggingContractMetaData.Bin instead.
var BeggingContractBin = BeggingContractMetaData.Bin

// DeployBeggingContract deploys a new Ethereum contract, binding an instance of BeggingContract to it.
func DeployBeggingContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BeggingContract, error) {
	parsed, err := BeggingContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BeggingContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BeggingContract{BeggingContractCaller: BeggingContractCaller{contract: contract}, BeggingContractTransactor: BeggingContractTransactor{contract: contract}, BeggingContractFilterer: BeggingContractFilterer{contract: contract}}, nil
}

// BeggingContract is an auto generated Go binding around an Ethereum contract.
type BeggingContract struct {
	BeggingContractCaller     // Read-only binding to the contract
	BeggingContractTransactor // Write-only binding to the contract
	BeggingContractFilterer   // Log filterer for contract events
}

// BeggingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeggingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeggingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeggingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeggingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeggingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeggingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeggingContractSession struct {
	Contract     *BeggingContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeggingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeggingContractCallerSession struct {
	Contract *BeggingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BeggingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeggingContractTransactorSession struct {
	Contract     *BeggingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BeggingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeggingContractRaw struct {
	Contract *BeggingContract // Generic contract binding to access the raw methods on
}

// BeggingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeggingContractCallerRaw struct {
	Contract *BeggingContractCaller // Generic read-only contract binding to access the raw methods on
}

// BeggingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeggingContractTransactorRaw struct {
	Contract *BeggingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeggingContract creates a new instance of BeggingContract, bound to a specific deployed contract.
func NewBeggingContract(address common.Address, backend bind.ContractBackend) (*BeggingContract, error) {
	contract, err := bindBeggingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeggingContract{BeggingContractCaller: BeggingContractCaller{contract: contract}, BeggingContractTransactor: BeggingContractTransactor{contract: contract}, BeggingContractFilterer: BeggingContractFilterer{contract: contract}}, nil
}

// NewBeggingContractCaller creates a new read-only instance of BeggingContract, bound to a specific deployed contract.
func NewBeggingContractCaller(address common.Address, caller bind.ContractCaller) (*BeggingContractCaller, error) {
	contract, err := bindBeggingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeggingContractCaller{contract: contract}, nil
}

// NewBeggingContractTransactor creates a new write-only instance of BeggingContract, bound to a specific deployed contract.
func NewBeggingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BeggingContractTransactor, error) {
	contract, err := bindBeggingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeggingContractTransactor{contract: contract}, nil
}

// NewBeggingContractFilterer creates a new log filterer instance of BeggingContract, bound to a specific deployed contract.
func NewBeggingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BeggingContractFilterer, error) {
	contract, err := bindBeggingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeggingContractFilterer{contract: contract}, nil
}

// bindBeggingContract binds a generic wrapper to an already deployed contract.
func bindBeggingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeggingContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeggingContract *BeggingContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeggingContract.Contract.BeggingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeggingContract *BeggingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContract.Contract.BeggingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeggingContract *BeggingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeggingContract.Contract.BeggingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeggingContract *BeggingContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeggingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeggingContract *BeggingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeggingContract *BeggingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeggingContract.Contract.contract.Transact(opts, method, params...)
}

// DonationMapping is a free data retrieval call binding the contract method 0xc184a6ab.
//
// Solidity: function donationMapping(address ) view returns(uint256)
func (_BeggingContract *BeggingContractCaller) DonationMapping(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeggingContract.contract.Call(opts, &out, "donationMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DonationMapping is a free data retrieval call binding the contract method 0xc184a6ab.
//
// Solidity: function donationMapping(address ) view returns(uint256)
func (_BeggingContract *BeggingContractSession) DonationMapping(arg0 common.Address) (*big.Int, error) {
	return _BeggingContract.Contract.DonationMapping(&_BeggingContract.CallOpts, arg0)
}

// DonationMapping is a free data retrieval call binding the contract method 0xc184a6ab.
//
// Solidity: function donationMapping(address ) view returns(uint256)
func (_BeggingContract *BeggingContractCallerSession) DonationMapping(arg0 common.Address) (*big.Int, error) {
	return _BeggingContract.Contract.DonationMapping(&_BeggingContract.CallOpts, arg0)
}

// Donations is a free data retrieval call binding the contract method 0xf8626af8.
//
// Solidity: function donations(uint256 ) view returns(address donor, uint256 amount, uint256 timestamp)
func (_BeggingContract *BeggingContractCaller) Donations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Donor     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _BeggingContract.contract.Call(opts, &out, "donations", arg0)

	outstruct := new(struct {
		Donor     common.Address
		Amount    *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Donor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Donations is a free data retrieval call binding the contract method 0xf8626af8.
//
// Solidity: function donations(uint256 ) view returns(address donor, uint256 amount, uint256 timestamp)
func (_BeggingContract *BeggingContractSession) Donations(arg0 *big.Int) (struct {
	Donor     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _BeggingContract.Contract.Donations(&_BeggingContract.CallOpts, arg0)
}

// Donations is a free data retrieval call binding the contract method 0xf8626af8.
//
// Solidity: function donations(uint256 ) view returns(address donor, uint256 amount, uint256 timestamp)
func (_BeggingContract *BeggingContractCallerSession) Donations(arg0 *big.Int) (struct {
	Donor     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _BeggingContract.Contract.Donations(&_BeggingContract.CallOpts, arg0)
}

// GetDonation is a free data retrieval call binding the contract method 0x410a1d32.
//
// Solidity: function getDonation(address from) view returns(uint256)
func (_BeggingContract *BeggingContractCaller) GetDonation(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeggingContract.contract.Call(opts, &out, "getDonation", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonation is a free data retrieval call binding the contract method 0x410a1d32.
//
// Solidity: function getDonation(address from) view returns(uint256)
func (_BeggingContract *BeggingContractSession) GetDonation(from common.Address) (*big.Int, error) {
	return _BeggingContract.Contract.GetDonation(&_BeggingContract.CallOpts, from)
}

// GetDonation is a free data retrieval call binding the contract method 0x410a1d32.
//
// Solidity: function getDonation(address from) view returns(uint256)
func (_BeggingContract *BeggingContractCallerSession) GetDonation(from common.Address) (*big.Int, error) {
	return _BeggingContract.Contract.GetDonation(&_BeggingContract.CallOpts, from)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeggingContract *BeggingContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeggingContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeggingContract *BeggingContractSession) Owner() (common.Address, error) {
	return _BeggingContract.Contract.Owner(&_BeggingContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeggingContract *BeggingContractCallerSession) Owner() (common.Address, error) {
	return _BeggingContract.Contract.Owner(&_BeggingContract.CallOpts)
}

// TotalDonations is a free data retrieval call binding the contract method 0xde2ed893.
//
// Solidity: function totalDonations() view returns(uint256)
func (_BeggingContract *BeggingContractCaller) TotalDonations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeggingContract.contract.Call(opts, &out, "totalDonations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDonations is a free data retrieval call binding the contract method 0xde2ed893.
//
// Solidity: function totalDonations() view returns(uint256)
func (_BeggingContract *BeggingContractSession) TotalDonations() (*big.Int, error) {
	return _BeggingContract.Contract.TotalDonations(&_BeggingContract.CallOpts)
}

// TotalDonations is a free data retrieval call binding the contract method 0xde2ed893.
//
// Solidity: function totalDonations() view returns(uint256)
func (_BeggingContract *BeggingContractCallerSession) TotalDonations() (*big.Int, error) {
	return _BeggingContract.Contract.TotalDonations(&_BeggingContract.CallOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_BeggingContract *BeggingContractTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContract.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_BeggingContract *BeggingContractSession) Donate() (*types.Transaction, error) {
	return _BeggingContract.Contract.Donate(&_BeggingContract.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_BeggingContract *BeggingContractTransactorSession) Donate() (*types.Transaction, error) {
	return _BeggingContract.Contract.Donate(&_BeggingContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BeggingContract *BeggingContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BeggingContract *BeggingContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _BeggingContract.Contract.RenounceOwnership(&_BeggingContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BeggingContract *BeggingContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BeggingContract.Contract.RenounceOwnership(&_BeggingContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BeggingContract *BeggingContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BeggingContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BeggingContract *BeggingContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BeggingContract.Contract.TransferOwnership(&_BeggingContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BeggingContract *BeggingContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BeggingContract.Contract.TransferOwnership(&_BeggingContract.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BeggingContract *BeggingContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BeggingContract *BeggingContractSession) Withdraw() (*types.Transaction, error) {
	return _BeggingContract.Contract.Withdraw(&_BeggingContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BeggingContract *BeggingContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _BeggingContract.Contract.Withdraw(&_BeggingContract.TransactOpts)
}

// BeggingContractDonationReceivedIterator is returned from FilterDonationReceived and is used to iterate over the raw logs and unpacked data for DonationReceived events raised by the BeggingContract contract.
type BeggingContractDonationReceivedIterator struct {
	Event *BeggingContractDonationReceived // Event containing the contract specifics and raw log

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
func (it *BeggingContractDonationReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeggingContractDonationReceived)
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
		it.Event = new(BeggingContractDonationReceived)
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
func (it *BeggingContractDonationReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeggingContractDonationReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeggingContractDonationReceived represents a DonationReceived event raised by the BeggingContract contract.
type BeggingContractDonationReceived struct {
	Donor     common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDonationReceived is a free log retrieval operation binding the contract event 0x4b0304ab2f313234b1cea012339c8f8c3bf2c8bb357d81a579cfdb35d2e3d030.
//
// Solidity: event DonationReceived(address indexed donor, uint256 amount, uint256 timestamp)
func (_BeggingContract *BeggingContractFilterer) FilterDonationReceived(opts *bind.FilterOpts, donor []common.Address) (*BeggingContractDonationReceivedIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _BeggingContract.contract.FilterLogs(opts, "DonationReceived", donorRule)
	if err != nil {
		return nil, err
	}
	return &BeggingContractDonationReceivedIterator{contract: _BeggingContract.contract, event: "DonationReceived", logs: logs, sub: sub}, nil
}

// WatchDonationReceived is a free log subscription operation binding the contract event 0x4b0304ab2f313234b1cea012339c8f8c3bf2c8bb357d81a579cfdb35d2e3d030.
//
// Solidity: event DonationReceived(address indexed donor, uint256 amount, uint256 timestamp)
func (_BeggingContract *BeggingContractFilterer) WatchDonationReceived(opts *bind.WatchOpts, sink chan<- *BeggingContractDonationReceived, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _BeggingContract.contract.WatchLogs(opts, "DonationReceived", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeggingContractDonationReceived)
				if err := _BeggingContract.contract.UnpackLog(event, "DonationReceived", log); err != nil {
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

// ParseDonationReceived is a log parse operation binding the contract event 0x4b0304ab2f313234b1cea012339c8f8c3bf2c8bb357d81a579cfdb35d2e3d030.
//
// Solidity: event DonationReceived(address indexed donor, uint256 amount, uint256 timestamp)
func (_BeggingContract *BeggingContractFilterer) ParseDonationReceived(log types.Log) (*BeggingContractDonationReceived, error) {
	event := new(BeggingContractDonationReceived)
	if err := _BeggingContract.contract.UnpackLog(event, "DonationReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeggingContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BeggingContract contract.
type BeggingContractOwnershipTransferredIterator struct {
	Event *BeggingContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BeggingContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeggingContractOwnershipTransferred)
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
		it.Event = new(BeggingContractOwnershipTransferred)
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
func (it *BeggingContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeggingContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeggingContractOwnershipTransferred represents a OwnershipTransferred event raised by the BeggingContract contract.
type BeggingContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BeggingContract *BeggingContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BeggingContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BeggingContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BeggingContractOwnershipTransferredIterator{contract: _BeggingContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BeggingContract *BeggingContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BeggingContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BeggingContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeggingContractOwnershipTransferred)
				if err := _BeggingContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BeggingContract *BeggingContractFilterer) ParseOwnershipTransferred(log types.Log) (*BeggingContractOwnershipTransferred, error) {
	event := new(BeggingContractOwnershipTransferred)
	if err := _BeggingContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
