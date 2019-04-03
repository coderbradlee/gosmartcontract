// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arraydelete

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ArrayDeleteABI is the input ABI used to generate the binding from.
const ArrayDeleteABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"main\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArrayDeleteBin is the compiled bytecode used for deploying new contracts.
const ArrayDeleteBin = `0x608060405234801561001057600080fd5b506101b6806100206000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663dffeadd08114610045575b600080fd5b34801561005157600080fd5b5061005a6100aa565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561009657818101518382015260200161007e565b505050509050019250505060405180910390f35b600080546001818101835582805260647f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563928301558254808201845560c8908301558254808201845561012c908301558254808201845561019090830155825490810183556101f4910155805460609190600290811061012657fe5b9060005260206000200160009055600080548060200260200160405190810160405280929190818152602001828054801561018057602002820191906000526020600020905b81548152602001906001019080831161016c575b50505050509050905600a165627a7a723058208aa9d63499b5c61b7afe28978ee8c5d5b2b55a814a7cbc5ebf426b87115075cf0029`

// DeployArrayDelete deploys a new Ethereum contract, binding an instance of ArrayDelete to it.
func DeployArrayDelete(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArrayDelete, error) {
	parsed, err := abi.JSON(strings.NewReader(ArrayDeleteABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArrayDeleteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArrayDelete{ArrayDeleteCaller: ArrayDeleteCaller{contract: contract}, ArrayDeleteTransactor: ArrayDeleteTransactor{contract: contract}, ArrayDeleteFilterer: ArrayDeleteFilterer{contract: contract}}, nil
}

// ArrayDelete is an auto generated Go binding around an Ethereum contract.
type ArrayDelete struct {
	ArrayDeleteCaller     // Read-only binding to the contract
	ArrayDeleteTransactor // Write-only binding to the contract
	ArrayDeleteFilterer   // Log filterer for contract events
}

// ArrayDeleteCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArrayDeleteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayDeleteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArrayDeleteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayDeleteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArrayDeleteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayDeleteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArrayDeleteSession struct {
	Contract     *ArrayDelete      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArrayDeleteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArrayDeleteCallerSession struct {
	Contract *ArrayDeleteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ArrayDeleteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArrayDeleteTransactorSession struct {
	Contract     *ArrayDeleteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ArrayDeleteRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArrayDeleteRaw struct {
	Contract *ArrayDelete // Generic contract binding to access the raw methods on
}

// ArrayDeleteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArrayDeleteCallerRaw struct {
	Contract *ArrayDeleteCaller // Generic read-only contract binding to access the raw methods on
}

// ArrayDeleteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArrayDeleteTransactorRaw struct {
	Contract *ArrayDeleteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArrayDelete creates a new instance of ArrayDelete, bound to a specific deployed contract.
func NewArrayDelete(address common.Address, backend bind.ContractBackend) (*ArrayDelete, error) {
	contract, err := bindArrayDelete(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArrayDelete{ArrayDeleteCaller: ArrayDeleteCaller{contract: contract}, ArrayDeleteTransactor: ArrayDeleteTransactor{contract: contract}, ArrayDeleteFilterer: ArrayDeleteFilterer{contract: contract}}, nil
}

// NewArrayDeleteCaller creates a new read-only instance of ArrayDelete, bound to a specific deployed contract.
func NewArrayDeleteCaller(address common.Address, caller bind.ContractCaller) (*ArrayDeleteCaller, error) {
	contract, err := bindArrayDelete(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayDeleteCaller{contract: contract}, nil
}

// NewArrayDeleteTransactor creates a new write-only instance of ArrayDelete, bound to a specific deployed contract.
func NewArrayDeleteTransactor(address common.Address, transactor bind.ContractTransactor) (*ArrayDeleteTransactor, error) {
	contract, err := bindArrayDelete(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayDeleteTransactor{contract: contract}, nil
}

// NewArrayDeleteFilterer creates a new log filterer instance of ArrayDelete, bound to a specific deployed contract.
func NewArrayDeleteFilterer(address common.Address, filterer bind.ContractFilterer) (*ArrayDeleteFilterer, error) {
	contract, err := bindArrayDelete(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArrayDeleteFilterer{contract: contract}, nil
}

// bindArrayDelete binds a generic wrapper to an already deployed contract.
func bindArrayDelete(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArrayDeleteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArrayDelete *ArrayDeleteRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArrayDelete.Contract.ArrayDeleteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArrayDelete *ArrayDeleteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArrayDelete.Contract.ArrayDeleteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArrayDelete *ArrayDeleteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArrayDelete.Contract.ArrayDeleteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArrayDelete *ArrayDeleteCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArrayDelete.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArrayDelete *ArrayDeleteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArrayDelete.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArrayDelete *ArrayDeleteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArrayDelete.Contract.contract.Transact(opts, method, params...)
}

// Main is a paid mutator transaction binding the contract method 0xdffeadd0.
//
// Solidity: function main() returns(uint256[])
func (_ArrayDelete *ArrayDeleteTransactor) Main(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArrayDelete.contract.Transact(opts, "main")
}

// Main is a paid mutator transaction binding the contract method 0xdffeadd0.
//
// Solidity: function main() returns(uint256[])
func (_ArrayDelete *ArrayDeleteSession) Main() (*types.Transaction, error) {
	return _ArrayDelete.Contract.Main(&_ArrayDelete.TransactOpts)
}

// Main is a paid mutator transaction binding the contract method 0xdffeadd0.
//
// Solidity: function main() returns(uint256[])
func (_ArrayDelete *ArrayDeleteTransactorSession) Main() (*types.Transaction, error) {
	return _ArrayDelete.Contract.Main(&_ArrayDelete.TransactOpts)
}
