// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arrayreturn

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

// AABI is the input ABI used to generate the binding from.
const AABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"foo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ABin is the compiled bytecode used for deploying new contracts.
const ABin = `0x608060405234801561001057600080fd5b50600080546001818101835582805260647f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563928301558254808201845560c890830155825490810190925561012c910155610125806100706000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663c298557881146043575b600080fd5b348015604e57600080fd5b50605560a3565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015608f5781810151838201526020016079565b505050509050019250505060405180910390f35b6060600080548060200260200160405190810160405280929190818152602001828054801560ef57602002820191906000526020600020905b81548152602001906001019080831160dc575b50505050509050905600a165627a7a723058202b92ee77c566f1cb20c90f2216e68746355bac330751bbfe5ff7799b10e2e0440029`

// DeployA deploys a new Ethereum contract, binding an instance of A to it.
func DeployA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *A, error) {
	parsed, err := abi.JSON(strings.NewReader(AABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &A{ACaller: ACaller{contract: contract}, ATransactor: ATransactor{contract: contract}, AFilterer: AFilterer{contract: contract}}, nil
}

// A is an auto generated Go binding around an Ethereum contract.
type A struct {
	ACaller     // Read-only binding to the contract
	ATransactor // Write-only binding to the contract
	AFilterer   // Log filterer for contract events
}

// ACaller is an auto generated read-only Go binding around an Ethereum contract.
type ACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ASession struct {
	Contract     *A                // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ACallerSession struct {
	Contract *ACaller      // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ATransactorSession struct {
	Contract     *ATransactor      // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ARaw is an auto generated low-level Go binding around an Ethereum contract.
type ARaw struct {
	Contract *A // Generic contract binding to access the raw methods on
}

// ACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ACallerRaw struct {
	Contract *ACaller // Generic read-only contract binding to access the raw methods on
}

// ATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ATransactorRaw struct {
	Contract *ATransactor // Generic write-only contract binding to access the raw methods on
}

// NewA creates a new instance of A, bound to a specific deployed contract.
func NewA(address common.Address, backend bind.ContractBackend) (*A, error) {
	contract, err := bindA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &A{ACaller: ACaller{contract: contract}, ATransactor: ATransactor{contract: contract}, AFilterer: AFilterer{contract: contract}}, nil
}

// NewACaller creates a new read-only instance of A, bound to a specific deployed contract.
func NewACaller(address common.Address, caller bind.ContractCaller) (*ACaller, error) {
	contract, err := bindA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ACaller{contract: contract}, nil
}

// NewATransactor creates a new write-only instance of A, bound to a specific deployed contract.
func NewATransactor(address common.Address, transactor bind.ContractTransactor) (*ATransactor, error) {
	contract, err := bindA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ATransactor{contract: contract}, nil
}

// NewAFilterer creates a new log filterer instance of A, bound to a specific deployed contract.
func NewAFilterer(address common.Address, filterer bind.ContractFilterer) (*AFilterer, error) {
	contract, err := bindA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AFilterer{contract: contract}, nil
}

// bindA binds a generic wrapper to an already deployed contract.
func bindA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_A *ARaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _A.Contract.ACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_A *ARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _A.Contract.ATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_A *ARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _A.Contract.ATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_A *ACallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _A.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_A *ATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _A.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_A *ATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _A.Contract.contract.Transact(opts, method, params...)
}

// Foo is a free data retrieval call binding the contract method 0xc2985578.
//
// Solidity: function foo() constant returns(uint256[])
func (_A *ACaller) Foo(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _A.contract.Call(opts, out, "foo")
	return *ret0, err
}

// Foo is a free data retrieval call binding the contract method 0xc2985578.
//
// Solidity: function foo() constant returns(uint256[])
func (_A *ASession) Foo() ([]*big.Int, error) {
	return _A.Contract.Foo(&_A.CallOpts)
}

// Foo is a free data retrieval call binding the contract method 0xc2985578.
//
// Solidity: function foo() constant returns(uint256[])
func (_A *ACallerSession) Foo() ([]*big.Int, error) {
	return _A.Contract.Foo(&_A.CallOpts)
}

// BABI is the input ABI used to generate the binding from.
const BABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"bar\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BBin is the compiled bytecode used for deploying new contracts.
const BBin = `0x608060405234801561001057600080fd5b5061001961005b565b604051809103906000f080158015610035573d6000803e3d6000fd5b5060008054600160a060020a031916600160a060020a039290921691909117905561006b565b6040516101958061024b83390190565b6101d18061007a6000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663febb0f7e8114610045575b600080fd5b34801561005157600080fd5b5061005a6100aa565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561009657818101518382015260200161007e565b505050509050019250505060405180910390f35b60008054604080517fc2985578000000000000000000000000000000000000000000000000000000008152905160609373ffffffffffffffffffffffffffffffffffffffff9093169263c2985578926004808201939182900301818387803b15801561011557600080fd5b505af1158015610129573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561015257600080fd5b81019080805164010000000081111561016a57600080fd5b8201602081018481111561017d57600080fd5b815185602082028301116401000000008211171561019a57600080fd5b5090945050505050905600a165627a7a72305820200c3fa7019d44689e19cdc6c76a74b03cab23c5952e6af02d6ec334d08e2abc0029608060405234801561001057600080fd5b50600080546001818101835582805260647f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563928301558254808201845560c890830155825490810190925561012c910155610125806100706000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663c298557881146043575b600080fd5b348015604e57600080fd5b50605560a3565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015608f5781810151838201526020016079565b505050509050019250505060405180910390f35b6060600080548060200260200160405190810160405280929190818152602001828054801560ef57602002820191906000526020600020905b81548152602001906001019080831160dc575b50505050509050905600a165627a7a723058202b92ee77c566f1cb20c90f2216e68746355bac330751bbfe5ff7799b10e2e0440029`

// DeployB deploys a new Ethereum contract, binding an instance of B to it.
func DeployB(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *B, error) {
	parsed, err := abi.JSON(strings.NewReader(BABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &B{BCaller: BCaller{contract: contract}, BTransactor: BTransactor{contract: contract}, BFilterer: BFilterer{contract: contract}}, nil
}

// B is an auto generated Go binding around an Ethereum contract.
type B struct {
	BCaller     // Read-only binding to the contract
	BTransactor // Write-only binding to the contract
	BFilterer   // Log filterer for contract events
}

// BCaller is an auto generated read-only Go binding around an Ethereum contract.
type BCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BSession struct {
	Contract     *B                // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BCallerSession struct {
	Contract *BCaller      // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BTransactorSession struct {
	Contract     *BTransactor      // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BRaw is an auto generated low-level Go binding around an Ethereum contract.
type BRaw struct {
	Contract *B // Generic contract binding to access the raw methods on
}

// BCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BCallerRaw struct {
	Contract *BCaller // Generic read-only contract binding to access the raw methods on
}

// BTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BTransactorRaw struct {
	Contract *BTransactor // Generic write-only contract binding to access the raw methods on
}

// NewB creates a new instance of B, bound to a specific deployed contract.
func NewB(address common.Address, backend bind.ContractBackend) (*B, error) {
	contract, err := bindB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &B{BCaller: BCaller{contract: contract}, BTransactor: BTransactor{contract: contract}, BFilterer: BFilterer{contract: contract}}, nil
}

// NewBCaller creates a new read-only instance of B, bound to a specific deployed contract.
func NewBCaller(address common.Address, caller bind.ContractCaller) (*BCaller, error) {
	contract, err := bindB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BCaller{contract: contract}, nil
}

// NewBTransactor creates a new write-only instance of B, bound to a specific deployed contract.
func NewBTransactor(address common.Address, transactor bind.ContractTransactor) (*BTransactor, error) {
	contract, err := bindB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BTransactor{contract: contract}, nil
}

// NewBFilterer creates a new log filterer instance of B, bound to a specific deployed contract.
func NewBFilterer(address common.Address, filterer bind.ContractFilterer) (*BFilterer, error) {
	contract, err := bindB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BFilterer{contract: contract}, nil
}

// bindB binds a generic wrapper to an already deployed contract.
func bindB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_B *BRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _B.Contract.BCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_B *BRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _B.Contract.BTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_B *BRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _B.Contract.BTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_B *BCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _B.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_B *BTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _B.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_B *BTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _B.Contract.contract.Transact(opts, method, params...)
}

// Bar is a free data retrieval call binding the contract method 0xfebb0f7e.
//
// Solidity: function bar() constant returns(uint256[])
func (_B *BCaller) Bar(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _B.contract.Call(opts, out, "bar")
	return *ret0, err
}

// Bar is a free data retrieval call binding the contract method 0xfebb0f7e.
//
// Solidity: function bar() constant returns(uint256[])
func (_B *BSession) Bar() ([]*big.Int, error) {
	return _B.Contract.Bar(&_B.CallOpts)
}

// Bar is a free data retrieval call binding the contract method 0xfebb0f7e.
//
// Solidity: function bar() constant returns(uint256[])
func (_B *BCallerSession) Bar() ([]*big.Int, error) {
	return _B.Contract.Bar(&_B.CallOpts)
}
