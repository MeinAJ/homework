// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DonationReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"donationMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getDonation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDonations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100845760006040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161007b91906101bb565b60405180910390fd5b610093816100b660201b60201c565b50426001819055506312cc0300426100ab919061020f565b600281905550610243565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101a58261017a565b9050919050565b6101b58161019a565b82525050565b60006020820190506101d060008301846101ac565b92915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061021a826101d6565b9150610225836101d6565b925082820190508082111561023d5761023c6101e0565b5b92915050565b610d23806102526000396000f3fe6080604052600436106100865760003560e01c8063c184a6ab11610059578063c184a6ab14610121578063de2ed8931461015e578063ed88c68e14610189578063f2fde38b14610193578063f8626af8146101bc57610086565b80633ccfd60b1461008b578063410a1d32146100a2578063715018a6146100df5780638da5cb5b146100f6575b600080fd5b34801561009757600080fd5b506100a06101fb565b005b3480156100ae57600080fd5b506100c960048036038101906100c49190610887565b610302565b6040516100d691906108cd565b60405180910390f35b3480156100eb57600080fd5b506100f461034b565b005b34801561010257600080fd5b5061010b61035f565b60405161011891906108f7565b60405180910390f35b34801561012d57600080fd5b5061014860048036038101906101439190610887565b610388565b60405161015591906108cd565b60405180910390f35b34801561016a57600080fd5b506101736103a0565b60405161018091906108cd565b60405180910390f35b6101916103a8565b005b34801561019f57600080fd5b506101ba60048036038101906101b59190610887565b61051f565b005b3480156101c857600080fd5b506101e360048036038101906101de919061093e565b6105a5565b6040516101f29392919061096b565b60405180910390f35b6102036105ff565b60004790506000811161024b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610242906109ff565b60405180910390fd5b600061025561035f565b73ffffffffffffffffffffffffffffffffffffffff168260405161027890610a50565b60006040518083038185875af1925050503d80600081146102b5576040519150601f19603f3d011682016040523d82523d6000602084013e6102ba565b606091505b50509050806102fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f590610ab1565b60405180910390fd5b5050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6103536105ff565b61035d6000610686565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60036020528060005260406000206000915090505481565b600047905090565b6103b061074a565b6103b86107d6565b600060405180606001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001348152602001428152509050600481908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155505034600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104c59190610b00565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f4b0304ab2f313234b1cea012339c8f8c3bf2c8bb357d81a579cfdb35d2e3d0303442604051610514929190610b34565b60405180910390a250565b6105276105ff565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036105995760006040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161059091906108f7565b60405180910390fd5b6105a281610686565b50565b600481815481106105b557600080fd5b90600052602060002090600302016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154905083565b61060761081c565b73ffffffffffffffffffffffffffffffffffffffff1661062561035f565b73ffffffffffffffffffffffffffffffffffffffff16146106845761064861081c565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161067b91906108f7565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60015442101561078f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078690610bcf565b60405180910390fd5b6002544211156107d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107cb90610c3b565b60405180910390fd5b565b600134101561081a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081190610ccd565b60405180910390fd5b565b600033905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061085482610829565b9050919050565b61086481610849565b811461086f57600080fd5b50565b6000813590506108818161085b565b92915050565b60006020828403121561089d5761089c610824565b5b60006108ab84828501610872565b91505092915050565b6000819050919050565b6108c7816108b4565b82525050565b60006020820190506108e260008301846108be565b92915050565b6108f181610849565b82525050565b600060208201905061090c60008301846108e8565b92915050565b61091b816108b4565b811461092657600080fd5b50565b60008135905061093881610912565b92915050565b60006020828403121561095457610953610824565b5b600061096284828501610929565b91505092915050565b600060608201905061098060008301866108e8565b61098d60208301856108be565b61099a60408301846108be565b949350505050565b600082825260208201905092915050565b7f4e6f2066756e647320746f207769746864726177000000000000000000000000600082015250565b60006109e96014836109a2565b91506109f4826109b3565b602082019050919050565b60006020820190508181036000830152610a18816109dc565b9050919050565b600081905092915050565b50565b6000610a3a600083610a1f565b9150610a4582610a2a565b600082019050919050565b6000610a5b82610a2d565b9150819050919050565b7f5472616e73666572206661696c65640000000000000000000000000000000000600082015250565b6000610a9b600f836109a2565b9150610aa682610a65565b602082019050919050565b60006020820190508181036000830152610aca81610a8e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610b0b826108b4565b9150610b16836108b4565b9250828201905080821115610b2e57610b2d610ad1565b5b92915050565b6000604082019050610b4960008301856108be565b610b5660208301846108be565b9392505050565b7f42656767696e67436f6e74726163743a20646f6e6174696f6e206e6f7420737460008201527f6172746564207965740000000000000000000000000000000000000000000000602082015250565b6000610bb96029836109a2565b9150610bc482610b5d565b604082019050919050565b60006020820190508181036000830152610be881610bac565b9050919050565b7f42656767696e67436f6e74726163743a20646f6e6174696f6e20656e64656400600082015250565b6000610c25601f836109a2565b9150610c3082610bef565b602082019050919050565b60006020820190508181036000830152610c5481610c18565b9050919050565b7f42656767696e67436f6e74726163743a20646f6e6174696f6e20616d6f756e7460008201527f2073686f756c64206265206174206c6561737420312077656900000000000000602082015250565b6000610cb76039836109a2565b9150610cc282610c5b565b604082019050919050565b60006020820190508181036000830152610ce681610caa565b905091905056fea2646970667358221220567d1c77b2bfe5227b528a9b2eb83f19fdf4d094624955efbc0ae4d86ffeda2f64736f6c634300081c0033",
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
