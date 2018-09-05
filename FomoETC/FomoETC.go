// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FomoETC

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

// DiviesInterfaceABI is the input ABI used to generate the binding from.
const DiviesInterfaceABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// DiviesInterfaceBin is the compiled bytecode used for deploying new contracts.
const DiviesInterfaceBin = `0x`

// DeployDiviesInterface deploys a new Ethereum contract, binding an instance of DiviesInterface to it.
func DeployDiviesInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DiviesInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(DiviesInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DiviesInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DiviesInterface{DiviesInterfaceCaller: DiviesInterfaceCaller{contract: contract}, DiviesInterfaceTransactor: DiviesInterfaceTransactor{contract: contract}, DiviesInterfaceFilterer: DiviesInterfaceFilterer{contract: contract}}, nil
}

// DiviesInterface is an auto generated Go binding around an Ethereum contract.
type DiviesInterface struct {
	DiviesInterfaceCaller     // Read-only binding to the contract
	DiviesInterfaceTransactor // Write-only binding to the contract
	DiviesInterfaceFilterer   // Log filterer for contract events
}

// DiviesInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiviesInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiviesInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiviesInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiviesInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiviesInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiviesInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiviesInterfaceSession struct {
	Contract     *DiviesInterface  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiviesInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiviesInterfaceCallerSession struct {
	Contract *DiviesInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DiviesInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiviesInterfaceTransactorSession struct {
	Contract     *DiviesInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DiviesInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiviesInterfaceRaw struct {
	Contract *DiviesInterface // Generic contract binding to access the raw methods on
}

// DiviesInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiviesInterfaceCallerRaw struct {
	Contract *DiviesInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// DiviesInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiviesInterfaceTransactorRaw struct {
	Contract *DiviesInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiviesInterface creates a new instance of DiviesInterface, bound to a specific deployed contract.
func NewDiviesInterface(address common.Address, backend bind.ContractBackend) (*DiviesInterface, error) {
	contract, err := bindDiviesInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiviesInterface{DiviesInterfaceCaller: DiviesInterfaceCaller{contract: contract}, DiviesInterfaceTransactor: DiviesInterfaceTransactor{contract: contract}, DiviesInterfaceFilterer: DiviesInterfaceFilterer{contract: contract}}, nil
}

// NewDiviesInterfaceCaller creates a new read-only instance of DiviesInterface, bound to a specific deployed contract.
func NewDiviesInterfaceCaller(address common.Address, caller bind.ContractCaller) (*DiviesInterfaceCaller, error) {
	contract, err := bindDiviesInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiviesInterfaceCaller{contract: contract}, nil
}

// NewDiviesInterfaceTransactor creates a new write-only instance of DiviesInterface, bound to a specific deployed contract.
func NewDiviesInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*DiviesInterfaceTransactor, error) {
	contract, err := bindDiviesInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiviesInterfaceTransactor{contract: contract}, nil
}

// NewDiviesInterfaceFilterer creates a new log filterer instance of DiviesInterface, bound to a specific deployed contract.
func NewDiviesInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*DiviesInterfaceFilterer, error) {
	contract, err := bindDiviesInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiviesInterfaceFilterer{contract: contract}, nil
}

// bindDiviesInterface binds a generic wrapper to an already deployed contract.
func bindDiviesInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiviesInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiviesInterface *DiviesInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DiviesInterface.Contract.DiviesInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiviesInterface *DiviesInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiviesInterface.Contract.DiviesInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiviesInterface *DiviesInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiviesInterface.Contract.DiviesInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiviesInterface *DiviesInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DiviesInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiviesInterface *DiviesInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiviesInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiviesInterface *DiviesInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiviesInterface.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_DiviesInterface *DiviesInterfaceTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiviesInterface.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_DiviesInterface *DiviesInterfaceSession) Deposit() (*types.Transaction, error) {
	return _DiviesInterface.Contract.Deposit(&_DiviesInterface.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_DiviesInterface *DiviesInterfaceTransactorSession) Deposit() (*types.Transaction, error) {
	return _DiviesInterface.Contract.Deposit(&_DiviesInterface.TransactOpts)
}

// FETCKeysCalcLongABI is the input ABI used to generate the binding from.
const FETCKeysCalcLongABI = "[]"

// FETCKeysCalcLongBin is the compiled bytecode used for deploying new contracts.
const FETCKeysCalcLongBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058200a17fc8be5cabe6d540bf6ac87350c98777233597523e19bf4db379f1b2836640029`

// DeployFETCKeysCalcLong deploys a new Ethereum contract, binding an instance of FETCKeysCalcLong to it.
func DeployFETCKeysCalcLong(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FETCKeysCalcLong, error) {
	parsed, err := abi.JSON(strings.NewReader(FETCKeysCalcLongABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FETCKeysCalcLongBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FETCKeysCalcLong{FETCKeysCalcLongCaller: FETCKeysCalcLongCaller{contract: contract}, FETCKeysCalcLongTransactor: FETCKeysCalcLongTransactor{contract: contract}, FETCKeysCalcLongFilterer: FETCKeysCalcLongFilterer{contract: contract}}, nil
}

// FETCKeysCalcLong is an auto generated Go binding around an Ethereum contract.
type FETCKeysCalcLong struct {
	FETCKeysCalcLongCaller     // Read-only binding to the contract
	FETCKeysCalcLongTransactor // Write-only binding to the contract
	FETCKeysCalcLongFilterer   // Log filterer for contract events
}

// FETCKeysCalcLongCaller is an auto generated read-only Go binding around an Ethereum contract.
type FETCKeysCalcLongCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FETCKeysCalcLongTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FETCKeysCalcLongTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FETCKeysCalcLongFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FETCKeysCalcLongFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FETCKeysCalcLongSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FETCKeysCalcLongSession struct {
	Contract     *FETCKeysCalcLong // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FETCKeysCalcLongCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FETCKeysCalcLongCallerSession struct {
	Contract *FETCKeysCalcLongCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FETCKeysCalcLongTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FETCKeysCalcLongTransactorSession struct {
	Contract     *FETCKeysCalcLongTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FETCKeysCalcLongRaw is an auto generated low-level Go binding around an Ethereum contract.
type FETCKeysCalcLongRaw struct {
	Contract *FETCKeysCalcLong // Generic contract binding to access the raw methods on
}

// FETCKeysCalcLongCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FETCKeysCalcLongCallerRaw struct {
	Contract *FETCKeysCalcLongCaller // Generic read-only contract binding to access the raw methods on
}

// FETCKeysCalcLongTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FETCKeysCalcLongTransactorRaw struct {
	Contract *FETCKeysCalcLongTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFETCKeysCalcLong creates a new instance of FETCKeysCalcLong, bound to a specific deployed contract.
func NewFETCKeysCalcLong(address common.Address, backend bind.ContractBackend) (*FETCKeysCalcLong, error) {
	contract, err := bindFETCKeysCalcLong(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FETCKeysCalcLong{FETCKeysCalcLongCaller: FETCKeysCalcLongCaller{contract: contract}, FETCKeysCalcLongTransactor: FETCKeysCalcLongTransactor{contract: contract}, FETCKeysCalcLongFilterer: FETCKeysCalcLongFilterer{contract: contract}}, nil
}

// NewFETCKeysCalcLongCaller creates a new read-only instance of FETCKeysCalcLong, bound to a specific deployed contract.
func NewFETCKeysCalcLongCaller(address common.Address, caller bind.ContractCaller) (*FETCKeysCalcLongCaller, error) {
	contract, err := bindFETCKeysCalcLong(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FETCKeysCalcLongCaller{contract: contract}, nil
}

// NewFETCKeysCalcLongTransactor creates a new write-only instance of FETCKeysCalcLong, bound to a specific deployed contract.
func NewFETCKeysCalcLongTransactor(address common.Address, transactor bind.ContractTransactor) (*FETCKeysCalcLongTransactor, error) {
	contract, err := bindFETCKeysCalcLong(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FETCKeysCalcLongTransactor{contract: contract}, nil
}

// NewFETCKeysCalcLongFilterer creates a new log filterer instance of FETCKeysCalcLong, bound to a specific deployed contract.
func NewFETCKeysCalcLongFilterer(address common.Address, filterer bind.ContractFilterer) (*FETCKeysCalcLongFilterer, error) {
	contract, err := bindFETCKeysCalcLong(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FETCKeysCalcLongFilterer{contract: contract}, nil
}

// bindFETCKeysCalcLong binds a generic wrapper to an already deployed contract.
func bindFETCKeysCalcLong(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FETCKeysCalcLongABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FETCKeysCalcLong *FETCKeysCalcLongRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FETCKeysCalcLong.Contract.FETCKeysCalcLongCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FETCKeysCalcLong *FETCKeysCalcLongRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FETCKeysCalcLong.Contract.FETCKeysCalcLongTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FETCKeysCalcLong *FETCKeysCalcLongRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FETCKeysCalcLong.Contract.FETCKeysCalcLongTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FETCKeysCalcLong *FETCKeysCalcLongCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FETCKeysCalcLong.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FETCKeysCalcLong *FETCKeysCalcLongTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FETCKeysCalcLong.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FETCKeysCalcLong *FETCKeysCalcLongTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FETCKeysCalcLong.Contract.contract.Transact(opts, method, params...)
}

// FETCdatasetsABI is the input ABI used to generate the binding from.
const FETCdatasetsABI = "[]"

// FETCdatasetsBin is the compiled bytecode used for deploying new contracts.
const FETCdatasetsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582085fd60f8bd9b62fdadf8a7846161bae22343d7260a035e1db96815e4ca5d3a400029`

// DeployFETCdatasets deploys a new Ethereum contract, binding an instance of FETCdatasets to it.
func DeployFETCdatasets(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FETCdatasets, error) {
	parsed, err := abi.JSON(strings.NewReader(FETCdatasetsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FETCdatasetsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FETCdatasets{FETCdatasetsCaller: FETCdatasetsCaller{contract: contract}, FETCdatasetsTransactor: FETCdatasetsTransactor{contract: contract}, FETCdatasetsFilterer: FETCdatasetsFilterer{contract: contract}}, nil
}

// FETCdatasets is an auto generated Go binding around an Ethereum contract.
type FETCdatasets struct {
	FETCdatasetsCaller     // Read-only binding to the contract
	FETCdatasetsTransactor // Write-only binding to the contract
	FETCdatasetsFilterer   // Log filterer for contract events
}

// FETCdatasetsCaller is an auto generated read-only Go binding around an Ethereum contract.
type FETCdatasetsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FETCdatasetsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FETCdatasetsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FETCdatasetsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FETCdatasetsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FETCdatasetsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FETCdatasetsSession struct {
	Contract     *FETCdatasets     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FETCdatasetsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FETCdatasetsCallerSession struct {
	Contract *FETCdatasetsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FETCdatasetsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FETCdatasetsTransactorSession struct {
	Contract     *FETCdatasetsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FETCdatasetsRaw is an auto generated low-level Go binding around an Ethereum contract.
type FETCdatasetsRaw struct {
	Contract *FETCdatasets // Generic contract binding to access the raw methods on
}

// FETCdatasetsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FETCdatasetsCallerRaw struct {
	Contract *FETCdatasetsCaller // Generic read-only contract binding to access the raw methods on
}

// FETCdatasetsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FETCdatasetsTransactorRaw struct {
	Contract *FETCdatasetsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFETCdatasets creates a new instance of FETCdatasets, bound to a specific deployed contract.
func NewFETCdatasets(address common.Address, backend bind.ContractBackend) (*FETCdatasets, error) {
	contract, err := bindFETCdatasets(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FETCdatasets{FETCdatasetsCaller: FETCdatasetsCaller{contract: contract}, FETCdatasetsTransactor: FETCdatasetsTransactor{contract: contract}, FETCdatasetsFilterer: FETCdatasetsFilterer{contract: contract}}, nil
}

// NewFETCdatasetsCaller creates a new read-only instance of FETCdatasets, bound to a specific deployed contract.
func NewFETCdatasetsCaller(address common.Address, caller bind.ContractCaller) (*FETCdatasetsCaller, error) {
	contract, err := bindFETCdatasets(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FETCdatasetsCaller{contract: contract}, nil
}

// NewFETCdatasetsTransactor creates a new write-only instance of FETCdatasets, bound to a specific deployed contract.
func NewFETCdatasetsTransactor(address common.Address, transactor bind.ContractTransactor) (*FETCdatasetsTransactor, error) {
	contract, err := bindFETCdatasets(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FETCdatasetsTransactor{contract: contract}, nil
}

// NewFETCdatasetsFilterer creates a new log filterer instance of FETCdatasets, bound to a specific deployed contract.
func NewFETCdatasetsFilterer(address common.Address, filterer bind.ContractFilterer) (*FETCdatasetsFilterer, error) {
	contract, err := bindFETCdatasets(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FETCdatasetsFilterer{contract: contract}, nil
}

// bindFETCdatasets binds a generic wrapper to an already deployed contract.
func bindFETCdatasets(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FETCdatasetsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FETCdatasets *FETCdatasetsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FETCdatasets.Contract.FETCdatasetsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FETCdatasets *FETCdatasetsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FETCdatasets.Contract.FETCdatasetsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FETCdatasets *FETCdatasetsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FETCdatasets.Contract.FETCdatasetsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FETCdatasets *FETCdatasetsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FETCdatasets.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FETCdatasets *FETCdatasetsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FETCdatasets.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FETCdatasets *FETCdatasetsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FETCdatasets.Contract.contract.Transact(opts, method, params...)
}

// FETCeventsABI is the input ABI used to generate the binding from.
const FETCeventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isNewPlayer\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onNewName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"etcIn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"keysBought\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"PSLAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"potAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"airDropPot\",\"type\":\"uint256\"}],\"name\":\"onEndTx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"etcOut\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"etcOut\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"PSLAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onWithdrawAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"etcIn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"PSLAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onBuyAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"PSLAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onReLoadAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"roundID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"buyerID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onAffiliatePayout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"roundID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountAddedToPot\",\"type\":\"uint256\"}],\"name\":\"onPotSwapDeposit\",\"type\":\"event\"}]"

// FETCeventsBin is the compiled bytecode used for deploying new contracts.
const FETCeventsBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a7230582022cafe60731db20ad7b39f6e550fac2c5e12cf504dec811f11b2311e19cc500d0029`

// DeployFETCevents deploys a new Ethereum contract, binding an instance of FETCevents to it.
func DeployFETCevents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FETCevents, error) {
	parsed, err := abi.JSON(strings.NewReader(FETCeventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FETCeventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FETCevents{FETCeventsCaller: FETCeventsCaller{contract: contract}, FETCeventsTransactor: FETCeventsTransactor{contract: contract}, FETCeventsFilterer: FETCeventsFilterer{contract: contract}}, nil
}

// FETCevents is an auto generated Go binding around an Ethereum contract.
type FETCevents struct {
	FETCeventsCaller     // Read-only binding to the contract
	FETCeventsTransactor // Write-only binding to the contract
	FETCeventsFilterer   // Log filterer for contract events
}

// FETCeventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type FETCeventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FETCeventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FETCeventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FETCeventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FETCeventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FETCeventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FETCeventsSession struct {
	Contract     *FETCevents       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FETCeventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FETCeventsCallerSession struct {
	Contract *FETCeventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FETCeventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FETCeventsTransactorSession struct {
	Contract     *FETCeventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FETCeventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type FETCeventsRaw struct {
	Contract *FETCevents // Generic contract binding to access the raw methods on
}

// FETCeventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FETCeventsCallerRaw struct {
	Contract *FETCeventsCaller // Generic read-only contract binding to access the raw methods on
}

// FETCeventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FETCeventsTransactorRaw struct {
	Contract *FETCeventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFETCevents creates a new instance of FETCevents, bound to a specific deployed contract.
func NewFETCevents(address common.Address, backend bind.ContractBackend) (*FETCevents, error) {
	contract, err := bindFETCevents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FETCevents{FETCeventsCaller: FETCeventsCaller{contract: contract}, FETCeventsTransactor: FETCeventsTransactor{contract: contract}, FETCeventsFilterer: FETCeventsFilterer{contract: contract}}, nil
}

// NewFETCeventsCaller creates a new read-only instance of FETCevents, bound to a specific deployed contract.
func NewFETCeventsCaller(address common.Address, caller bind.ContractCaller) (*FETCeventsCaller, error) {
	contract, err := bindFETCevents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FETCeventsCaller{contract: contract}, nil
}

// NewFETCeventsTransactor creates a new write-only instance of FETCevents, bound to a specific deployed contract.
func NewFETCeventsTransactor(address common.Address, transactor bind.ContractTransactor) (*FETCeventsTransactor, error) {
	contract, err := bindFETCevents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FETCeventsTransactor{contract: contract}, nil
}

// NewFETCeventsFilterer creates a new log filterer instance of FETCevents, bound to a specific deployed contract.
func NewFETCeventsFilterer(address common.Address, filterer bind.ContractFilterer) (*FETCeventsFilterer, error) {
	contract, err := bindFETCevents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FETCeventsFilterer{contract: contract}, nil
}

// bindFETCevents binds a generic wrapper to an already deployed contract.
func bindFETCevents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FETCeventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FETCevents *FETCeventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FETCevents.Contract.FETCeventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FETCevents *FETCeventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FETCevents.Contract.FETCeventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FETCevents *FETCeventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FETCevents.Contract.FETCeventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FETCevents *FETCeventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FETCevents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FETCevents *FETCeventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FETCevents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FETCevents *FETCeventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FETCevents.Contract.contract.Transact(opts, method, params...)
}

// FETCeventsOnAffiliatePayoutIterator is returned from FilterOnAffiliatePayout and is used to iterate over the raw logs and unpacked data for OnAffiliatePayout events raised by the FETCevents contract.
type FETCeventsOnAffiliatePayoutIterator struct {
	Event *FETCeventsOnAffiliatePayout // Event containing the contract specifics and raw log

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
func (it *FETCeventsOnAffiliatePayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FETCeventsOnAffiliatePayout)
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
		it.Event = new(FETCeventsOnAffiliatePayout)
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
func (it *FETCeventsOnAffiliatePayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FETCeventsOnAffiliatePayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FETCeventsOnAffiliatePayout represents a OnAffiliatePayout event raised by the FETCevents contract.
type FETCeventsOnAffiliatePayout struct {
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	RoundID          *big.Int
	BuyerID          *big.Int
	Amount           *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnAffiliatePayout is a free log retrieval operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: event onAffiliatePayout(affiliateID indexed uint256, affiliateAddress address, affiliateName bytes32, roundID indexed uint256, buyerID indexed uint256, amount uint256, timeStamp uint256)
func (_FETCevents *FETCeventsFilterer) FilterOnAffiliatePayout(opts *bind.FilterOpts, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (*FETCeventsOnAffiliatePayoutIterator, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _FETCevents.contract.FilterLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return &FETCeventsOnAffiliatePayoutIterator{contract: _FETCevents.contract, event: "onAffiliatePayout", logs: logs, sub: sub}, nil
}

// WatchOnAffiliatePayout is a free log subscription operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: event onAffiliatePayout(affiliateID indexed uint256, affiliateAddress address, affiliateName bytes32, roundID indexed uint256, buyerID indexed uint256, amount uint256, timeStamp uint256)
func (_FETCevents *FETCeventsFilterer) WatchOnAffiliatePayout(opts *bind.WatchOpts, sink chan<- *FETCeventsOnAffiliatePayout, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (event.Subscription, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _FETCevents.contract.WatchLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FETCeventsOnAffiliatePayout)
				if err := _FETCevents.contract.UnpackLog(event, "onAffiliatePayout", log); err != nil {
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

// FETCeventsOnBuyAndDistributeIterator is returned from FilterOnBuyAndDistribute and is used to iterate over the raw logs and unpacked data for OnBuyAndDistribute events raised by the FETCevents contract.
type FETCeventsOnBuyAndDistributeIterator struct {
	Event *FETCeventsOnBuyAndDistribute // Event containing the contract specifics and raw log

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
func (it *FETCeventsOnBuyAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FETCeventsOnBuyAndDistribute)
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
		it.Event = new(FETCeventsOnBuyAndDistribute)
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
func (it *FETCeventsOnBuyAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FETCeventsOnBuyAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FETCeventsOnBuyAndDistribute represents a OnBuyAndDistribute event raised by the FETCevents contract.
type FETCeventsOnBuyAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EtcIn          *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	PSLAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnBuyAndDistribute is a free log retrieval operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: event onBuyAndDistribute(playerAddress address, playerName bytes32, etcIn uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_FETCevents *FETCeventsFilterer) FilterOnBuyAndDistribute(opts *bind.FilterOpts) (*FETCeventsOnBuyAndDistributeIterator, error) {

	logs, sub, err := _FETCevents.contract.FilterLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return &FETCeventsOnBuyAndDistributeIterator{contract: _FETCevents.contract, event: "onBuyAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnBuyAndDistribute is a free log subscription operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: event onBuyAndDistribute(playerAddress address, playerName bytes32, etcIn uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_FETCevents *FETCeventsFilterer) WatchOnBuyAndDistribute(opts *bind.WatchOpts, sink chan<- *FETCeventsOnBuyAndDistribute) (event.Subscription, error) {

	logs, sub, err := _FETCevents.contract.WatchLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FETCeventsOnBuyAndDistribute)
				if err := _FETCevents.contract.UnpackLog(event, "onBuyAndDistribute", log); err != nil {
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

// FETCeventsOnEndTxIterator is returned from FilterOnEndTx and is used to iterate over the raw logs and unpacked data for OnEndTx events raised by the FETCevents contract.
type FETCeventsOnEndTxIterator struct {
	Event *FETCeventsOnEndTx // Event containing the contract specifics and raw log

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
func (it *FETCeventsOnEndTxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FETCeventsOnEndTx)
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
		it.Event = new(FETCeventsOnEndTx)
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
func (it *FETCeventsOnEndTxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FETCeventsOnEndTxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FETCeventsOnEndTx represents a OnEndTx event raised by the FETCevents contract.
type FETCeventsOnEndTx struct {
	CompressedData *big.Int
	CompressedIDs  *big.Int
	PlayerName     [32]byte
	PlayerAddress  common.Address
	EtcIn          *big.Int
	KeysBought     *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	PSLAmount      *big.Int
	GenAmount      *big.Int
	PotAmount      *big.Int
	AirDropPot     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnEndTx is a free log retrieval operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: event onEndTx(compressedData uint256, compressedIDs uint256, playerName bytes32, playerAddress address, etcIn uint256, keysBought uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256, potAmount uint256, airDropPot uint256)
func (_FETCevents *FETCeventsFilterer) FilterOnEndTx(opts *bind.FilterOpts) (*FETCeventsOnEndTxIterator, error) {

	logs, sub, err := _FETCevents.contract.FilterLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return &FETCeventsOnEndTxIterator{contract: _FETCevents.contract, event: "onEndTx", logs: logs, sub: sub}, nil
}

// WatchOnEndTx is a free log subscription operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: event onEndTx(compressedData uint256, compressedIDs uint256, playerName bytes32, playerAddress address, etcIn uint256, keysBought uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256, potAmount uint256, airDropPot uint256)
func (_FETCevents *FETCeventsFilterer) WatchOnEndTx(opts *bind.WatchOpts, sink chan<- *FETCeventsOnEndTx) (event.Subscription, error) {

	logs, sub, err := _FETCevents.contract.WatchLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FETCeventsOnEndTx)
				if err := _FETCevents.contract.UnpackLog(event, "onEndTx", log); err != nil {
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

// FETCeventsOnNewNameIterator is returned from FilterOnNewName and is used to iterate over the raw logs and unpacked data for OnNewName events raised by the FETCevents contract.
type FETCeventsOnNewNameIterator struct {
	Event *FETCeventsOnNewName // Event containing the contract specifics and raw log

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
func (it *FETCeventsOnNewNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FETCeventsOnNewName)
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
		it.Event = new(FETCeventsOnNewName)
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
func (it *FETCeventsOnNewNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FETCeventsOnNewNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FETCeventsOnNewName represents a OnNewName event raised by the FETCevents contract.
type FETCeventsOnNewName struct {
	PlayerID         *big.Int
	PlayerAddress    common.Address
	PlayerName       [32]byte
	IsNewPlayer      bool
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	AmountPaid       *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnNewName is a free log retrieval operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: event onNewName(playerID indexed uint256, playerAddress indexed address, playerName indexed bytes32, isNewPlayer bool, affiliateID uint256, affiliateAddress address, affiliateName bytes32, amountPaid uint256, timeStamp uint256)
func (_FETCevents *FETCeventsFilterer) FilterOnNewName(opts *bind.FilterOpts, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (*FETCeventsOnNewNameIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _FETCevents.contract.FilterLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return &FETCeventsOnNewNameIterator{contract: _FETCevents.contract, event: "onNewName", logs: logs, sub: sub}, nil
}

// WatchOnNewName is a free log subscription operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: event onNewName(playerID indexed uint256, playerAddress indexed address, playerName indexed bytes32, isNewPlayer bool, affiliateID uint256, affiliateAddress address, affiliateName bytes32, amountPaid uint256, timeStamp uint256)
func (_FETCevents *FETCeventsFilterer) WatchOnNewName(opts *bind.WatchOpts, sink chan<- *FETCeventsOnNewName, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _FETCevents.contract.WatchLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FETCeventsOnNewName)
				if err := _FETCevents.contract.UnpackLog(event, "onNewName", log); err != nil {
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

// FETCeventsOnPotSwapDepositIterator is returned from FilterOnPotSwapDeposit and is used to iterate over the raw logs and unpacked data for OnPotSwapDeposit events raised by the FETCevents contract.
type FETCeventsOnPotSwapDepositIterator struct {
	Event *FETCeventsOnPotSwapDeposit // Event containing the contract specifics and raw log

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
func (it *FETCeventsOnPotSwapDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FETCeventsOnPotSwapDeposit)
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
		it.Event = new(FETCeventsOnPotSwapDeposit)
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
func (it *FETCeventsOnPotSwapDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FETCeventsOnPotSwapDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FETCeventsOnPotSwapDeposit represents a OnPotSwapDeposit event raised by the FETCevents contract.
type FETCeventsOnPotSwapDeposit struct {
	RoundID          *big.Int
	AmountAddedToPot *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnPotSwapDeposit is a free log retrieval operation binding the contract event 0x74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c.
//
// Solidity: event onPotSwapDeposit(roundID uint256, amountAddedToPot uint256)
func (_FETCevents *FETCeventsFilterer) FilterOnPotSwapDeposit(opts *bind.FilterOpts) (*FETCeventsOnPotSwapDepositIterator, error) {

	logs, sub, err := _FETCevents.contract.FilterLogs(opts, "onPotSwapDeposit")
	if err != nil {
		return nil, err
	}
	return &FETCeventsOnPotSwapDepositIterator{contract: _FETCevents.contract, event: "onPotSwapDeposit", logs: logs, sub: sub}, nil
}

// WatchOnPotSwapDeposit is a free log subscription operation binding the contract event 0x74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c.
//
// Solidity: event onPotSwapDeposit(roundID uint256, amountAddedToPot uint256)
func (_FETCevents *FETCeventsFilterer) WatchOnPotSwapDeposit(opts *bind.WatchOpts, sink chan<- *FETCeventsOnPotSwapDeposit) (event.Subscription, error) {

	logs, sub, err := _FETCevents.contract.WatchLogs(opts, "onPotSwapDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FETCeventsOnPotSwapDeposit)
				if err := _FETCevents.contract.UnpackLog(event, "onPotSwapDeposit", log); err != nil {
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

// FETCeventsOnReLoadAndDistributeIterator is returned from FilterOnReLoadAndDistribute and is used to iterate over the raw logs and unpacked data for OnReLoadAndDistribute events raised by the FETCevents contract.
type FETCeventsOnReLoadAndDistributeIterator struct {
	Event *FETCeventsOnReLoadAndDistribute // Event containing the contract specifics and raw log

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
func (it *FETCeventsOnReLoadAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FETCeventsOnReLoadAndDistribute)
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
		it.Event = new(FETCeventsOnReLoadAndDistribute)
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
func (it *FETCeventsOnReLoadAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FETCeventsOnReLoadAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FETCeventsOnReLoadAndDistribute represents a OnReLoadAndDistribute event raised by the FETCevents contract.
type FETCeventsOnReLoadAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	PSLAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnReLoadAndDistribute is a free log retrieval operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: event onReLoadAndDistribute(playerAddress address, playerName bytes32, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_FETCevents *FETCeventsFilterer) FilterOnReLoadAndDistribute(opts *bind.FilterOpts) (*FETCeventsOnReLoadAndDistributeIterator, error) {

	logs, sub, err := _FETCevents.contract.FilterLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return &FETCeventsOnReLoadAndDistributeIterator{contract: _FETCevents.contract, event: "onReLoadAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnReLoadAndDistribute is a free log subscription operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: event onReLoadAndDistribute(playerAddress address, playerName bytes32, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_FETCevents *FETCeventsFilterer) WatchOnReLoadAndDistribute(opts *bind.WatchOpts, sink chan<- *FETCeventsOnReLoadAndDistribute) (event.Subscription, error) {

	logs, sub, err := _FETCevents.contract.WatchLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FETCeventsOnReLoadAndDistribute)
				if err := _FETCevents.contract.UnpackLog(event, "onReLoadAndDistribute", log); err != nil {
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

// FETCeventsOnWithdrawIterator is returned from FilterOnWithdraw and is used to iterate over the raw logs and unpacked data for OnWithdraw events raised by the FETCevents contract.
type FETCeventsOnWithdrawIterator struct {
	Event *FETCeventsOnWithdraw // Event containing the contract specifics and raw log

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
func (it *FETCeventsOnWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FETCeventsOnWithdraw)
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
		it.Event = new(FETCeventsOnWithdraw)
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
func (it *FETCeventsOnWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FETCeventsOnWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FETCeventsOnWithdraw represents a OnWithdraw event raised by the FETCevents contract.
type FETCeventsOnWithdraw struct {
	PlayerID      *big.Int
	PlayerAddress common.Address
	PlayerName    [32]byte
	EtcOut        *big.Int
	TimeStamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOnWithdraw is a free log retrieval operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: event onWithdraw(playerID indexed uint256, playerAddress address, playerName bytes32, etcOut uint256, timeStamp uint256)
func (_FETCevents *FETCeventsFilterer) FilterOnWithdraw(opts *bind.FilterOpts, playerID []*big.Int) (*FETCeventsOnWithdrawIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _FETCevents.contract.FilterLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return &FETCeventsOnWithdrawIterator{contract: _FETCevents.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: event onWithdraw(playerID indexed uint256, playerAddress address, playerName bytes32, etcOut uint256, timeStamp uint256)
func (_FETCevents *FETCeventsFilterer) WatchOnWithdraw(opts *bind.WatchOpts, sink chan<- *FETCeventsOnWithdraw, playerID []*big.Int) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _FETCevents.contract.WatchLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FETCeventsOnWithdraw)
				if err := _FETCevents.contract.UnpackLog(event, "onWithdraw", log); err != nil {
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

// FETCeventsOnWithdrawAndDistributeIterator is returned from FilterOnWithdrawAndDistribute and is used to iterate over the raw logs and unpacked data for OnWithdrawAndDistribute events raised by the FETCevents contract.
type FETCeventsOnWithdrawAndDistributeIterator struct {
	Event *FETCeventsOnWithdrawAndDistribute // Event containing the contract specifics and raw log

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
func (it *FETCeventsOnWithdrawAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FETCeventsOnWithdrawAndDistribute)
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
		it.Event = new(FETCeventsOnWithdrawAndDistribute)
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
func (it *FETCeventsOnWithdrawAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FETCeventsOnWithdrawAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FETCeventsOnWithdrawAndDistribute represents a OnWithdrawAndDistribute event raised by the FETCevents contract.
type FETCeventsOnWithdrawAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EtcOut         *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	PSLAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnWithdrawAndDistribute is a free log retrieval operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: event onWithdrawAndDistribute(playerAddress address, playerName bytes32, etcOut uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_FETCevents *FETCeventsFilterer) FilterOnWithdrawAndDistribute(opts *bind.FilterOpts) (*FETCeventsOnWithdrawAndDistributeIterator, error) {

	logs, sub, err := _FETCevents.contract.FilterLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return &FETCeventsOnWithdrawAndDistributeIterator{contract: _FETCevents.contract, event: "onWithdrawAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnWithdrawAndDistribute is a free log subscription operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: event onWithdrawAndDistribute(playerAddress address, playerName bytes32, etcOut uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_FETCevents *FETCeventsFilterer) WatchOnWithdrawAndDistribute(opts *bind.WatchOpts, sink chan<- *FETCeventsOnWithdrawAndDistribute) (event.Subscription, error) {

	logs, sub, err := _FETCevents.contract.WatchLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FETCeventsOnWithdrawAndDistribute)
				if err := _FETCevents.contract.UnpackLog(event, "onWithdrawAndDistribute", log); err != nil {
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

// FomoETClongABI is the input ABI used to generate the binding from.
const FomoETClongABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBuyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_affCode\",\"type\":\"bytes32\"},{\"name\":\"_team\",\"type\":\"uint256\"},{\"name\":\"_etc\",\"type\":\"uint256\"}],\"name\":\"reLoadXname\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"pIDxAddr_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"airDropTracker_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"round_\",\"outputs\":[{\"name\":\"plyr\",\"type\":\"uint256\"},{\"name\":\"team\",\"type\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\"},{\"name\":\"ended\",\"type\":\"bool\"},{\"name\":\"strt\",\"type\":\"uint256\"},{\"name\":\"keys\",\"type\":\"uint256\"},{\"name\":\"etc\",\"type\":\"uint256\"},{\"name\":\"pot\",\"type\":\"uint256\"},{\"name\":\"mask\",\"type\":\"uint256\"},{\"name\":\"ico\",\"type\":\"uint256\"},{\"name\":\"icoGen\",\"type\":\"uint256\"},{\"name\":\"icoAvg\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"plyrNames_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fees_\",\"outputs\":[{\"name\":\"gen\",\"type\":\"uint256\"},{\"name\":\"psl\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pIDxName_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_affCode\",\"type\":\"uint256\"},{\"name\":\"_team\",\"type\":\"uint256\"},{\"name\":\"_etc\",\"type\":\"uint256\"}],\"name\":\"reLoadXid\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nameString\",\"type\":\"string\"},{\"name\":\"_affCode\",\"type\":\"address\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXaddr\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"},{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_laff\",\"type\":\"uint256\"}],\"name\":\"receivePlayerInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rndTmEth_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rID_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"}],\"name\":\"getPlayerVaults\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nameString\",\"type\":\"string\"},{\"name\":\"_affCode\",\"type\":\"bytes32\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXname\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentRoundInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_affCode\",\"type\":\"address\"},{\"name\":\"_team\",\"type\":\"uint256\"},{\"name\":\"_etc\",\"type\":\"uint256\"}],\"name\":\"reLoadXaddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_affCode\",\"type\":\"uint256\"},{\"name\":\"_team\",\"type\":\"uint256\"}],\"name\":\"buyXid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"receivePlayerNameList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nameString\",\"type\":\"string\"},{\"name\":\"_affCode\",\"type\":\"uint256\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXID\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_affCode\",\"type\":\"address\"},{\"name\":\"_team\",\"type\":\"uint256\"}],\"name\":\"buyXaddr\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plyrRnds_\",\"outputs\":[{\"name\":\"etc\",\"type\":\"uint256\"},{\"name\":\"keys\",\"type\":\"uint256\"},{\"name\":\"mask\",\"type\":\"uint256\"},{\"name\":\"ico\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_affCode\",\"type\":\"bytes32\"},{\"name\":\"_team\",\"type\":\"uint256\"}],\"name\":\"buyXname\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_otherFETC\",\"type\":\"address\"}],\"name\":\"setOtherFomo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"potSplit_\",\"outputs\":[{\"name\":\"gen\",\"type\":\"uint256\"},{\"name\":\"psl\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTimeLeft\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_rID\",\"type\":\"uint256\"},{\"name\":\"_etc\",\"type\":\"uint256\"}],\"name\":\"calcKeysReceived\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_keys\",\"type\":\"uint256\"}],\"name\":\"iWantXKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"activated_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"airDropPot_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plyr_\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"win\",\"type\":\"uint256\"},{\"name\":\"gen\",\"type\":\"uint256\"},{\"name\":\"aff\",\"type\":\"uint256\"},{\"name\":\"lrnd\",\"type\":\"uint256\"},{\"name\":\"laff\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"potSwap\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getPlayerInfoByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isNewPlayer\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onNewName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"etcIn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"keysBought\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"PSLAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"potAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"airDropPot\",\"type\":\"uint256\"}],\"name\":\"onEndTx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"etcOut\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"etcOut\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"PSLAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onWithdrawAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"etcIn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"PSLAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onBuyAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"PSLAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onReLoadAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"roundID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"buyerID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onAffiliatePayout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"roundID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountAddedToPot\",\"type\":\"uint256\"}],\"name\":\"onPotSwapDeposit\",\"type\":\"event\"}]"

// FomoETClongBin is the compiled bytecode used for deploying new contracts.
const FomoETClongBin = `0x6080604052611c20600155610e106002556000600455602e6006556010805460ff1916905534801561003057600080fd5b50604080518082018252601e8152600660208083019182526000808052600e80835293517fe710864318d4a32f37d6ce54cb3fadbef648dd12d8dbdf53973564d56b7f881c5591517fe710864318d4a32f37d6ce54cb3fadbef648dd12d8dbdf53973564d56b7f881d5583518085018552602b808252818301848152600180865286855292517fa7c5ba7114a813b50159add3a36832908dc83db71d0b9a24c2ad0f83be95820755517fa7c5ba7114a813b50159add3a36832908dc83db71d0b9a24c2ad0f83be958208558551808701875260388152600a818501908152600280875287865291517f9adb202b1492743bc00c81d33cdc6423fa8c79109027eb6a845391e8fc1f048155517f9adb202b1492743bc00c81d33cdc6423fa8c79109027eb6a845391e8fc1f048255865180880188529182526008828501908152600380875296855291517fe0283e559c29e31ee7f56467acc9dd307779c843a883aeeb3bf5c6128c9081445590517fe0283e559c29e31ee7f56467acc9dd307779c843a883aeeb3bf5c6128c908145558551808701875260118152600c818501818152868052600f80875292517ff4803e074bd026baaf6ed2e288c9515f68c72fb7216eebdd7cae1718a53ec37555517ff4803e074bd026baaf6ed2e288c9515f68c72fb7216eebdd7cae1718a53ec3765587518089018952601b8152808601848152948752828652517f169f97de0d9a84d840042b17d3c6b9638b3d6fd9024c9eb0c7a306a17b49f88f5592517f169f97de0d9a84d840042b17d3c6b9638b3d6fd9024c9eb0c7a306a17b49f89055865180880188526016808252818601908152928652818552517fa74ba3945261e09fde15ba3db55005b205e61eeb4ad811ac0faa2b315bffeead5590517fa74ba3945261e09fde15ba3db55005b205e61eeb4ad811ac0faa2b315bffeeae558551808701909652828652858301918252939092529190915290517f45f76dafbbad695564362934e24d72eedc57f9fc1a65f39bca62176cc829682855517f45f76dafbbad695564362934e24d72eedc57f9fc1a65f39bca62176cc829682955615655806200034e6000396000f3006080604052600436106101b35763ffffffff60e060020a600035041663018a25e8811461035357806306fdde031461037a578063079ce327146104045780630f15f4c01461042457806310f01eba1461043957806311a09ae71461045a57806324c33d331461046f5780632660316e146104e65780632ce21999146105155780632e19ebdc14610546578063349cdcac1461055e5780633ccfd60b1461057c5780633ddd46981461059157806349cc635d146105ed5780635893d48114610617578063624ae5c0146106325780636306643414610647578063685ffd831461067d578063747dff42146106d057806382bfc7391461075b5780638f38f309146107825780638f7140ea14610790578063921dec21146107ab57806395d89b41146107fe57806398a0871d14610813578063a2bccae91461082a578063a65b37a11461086b578063b483c05414610879578063c519500e1461089a578063c7e284b8146108b2578063ce89c80c146108c7578063cf808000146108e2578063d53b2679146108fa578063d87574e01461090f578063de7874f314610924578063ed78cf4a1461097e578063ee0b5d8b14610986575b6101bb615530565b60105460009060ff16151560011461021f576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155ca833981519152604482015260008051602061558a833981519152606482015290519081900360840190fd5b33803b8015610266576040805160e560020a62461bcd028152602060048201526011602482015260008051602061560a833981519152604482015290519081900360640190fd5b34633b9aca008110156102be576040805160e560020a62461bcd02815260206004820152602160248201526000805160206155aa833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af680000081111561030e576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206155ea833981519152604482015290519081900360640190fd5b610317856109df565b33600090815260076020908152604080832054808452600990925290912060060154919650945061034c908590600288610c93565b5050505050005b34801561035f57600080fd5b50610368610ecd565b60408051918252519081900360200190f35b34801561038657600080fd5b5061038f610f92565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103c95781810151838201526020016103b1565b50505050905090810190601f1680156103f65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561041057600080fd5b50610422600435602435604435610fc9565b005b34801561043057600080fd5b506104226111d5565b34801561044557600080fd5b50610368600160a060020a03600435166113ed565b34801561046657600080fd5b506103686113ff565b34801561047b57600080fd5b50610487600435611405565b604080519c8d5260208d019b909b528b8b019990995296151560608b015260808a019590955260a089019390935260c088019190915260e087015261010086015261012085015261014084015261016083015251908190036101800190f35b3480156104f257600080fd5b50610501600435602435611468565b604080519115158252519081900360200190f35b34801561052157600080fd5b5061052d600435611488565b6040805192835260208301919091528051918290030190f35b34801561055257600080fd5b506103686004356114a1565b34801561056a57600080fd5b506104226004356024356044356114b3565b34801561058857600080fd5b50610422611699565b6040805160206004803580820135601f810184900484028501840190955284845261042294369492936024939284019190819084018382808284375094975050600160a060020a03853516955050505050602001351515611a1a565b3480156105f957600080fd5b50610422600435600160a060020a0360243516604435606435611bd3565b34801561062357600080fd5b50610368600435602435611dc4565b34801561063e57600080fd5b50610368611de1565b34801561065357600080fd5b5061065f600435611de7565b60408051938452602084019290925282820152519081900360600190f35b6040805160206004803580820135601f8101849004840285018401909552848452610422943694929360249392840191908190840183828082843750949750508435955050505050602001351515611f8e565b3480156106dc57600080fd5b506106e561206e565b604080519e8f5260208f019d909d528d8d019b909b5260608d019990995260808c019790975260a08b019590955260c08a0193909352600160a060020a0390911660e08901526101008801526101208701526101408601526101608501526101808401526101a083015251908190036101c00190f35b34801561076757600080fd5b50610422600160a060020a036004351660243560443561226c565b610422600435602435612468565b34801561079c57600080fd5b5061042260043560243561264f565b6040805160206004803580820135601f810184900484028501840190955284845261042294369492936024939284019190819084018382808284375094975050843595505050505060200135151561272c565b34801561080a57600080fd5b5061038f61280c565b610422600160a060020a0360043516602435612843565b34801561083657600080fd5b50610845600435602435612a5a565b604080519485526020850193909352838301919091526060830152519081900360800190f35b610422600435602435612a8c565b34801561088557600080fd5b50610422600160a060020a0360043516612c89565b3480156108a657600080fd5b5061052d600435612dfe565b3480156108be57600080fd5b50610368612e17565b3480156108d357600080fd5b50610368600435602435612ea6565b3480156108ee57600080fd5b50610368600435612f4e565b34801561090657600080fd5b50610501613000565b34801561091b57600080fd5b50610368613009565b34801561093057600080fd5b5061093c60043561300f565b60408051600160a060020a0390981688526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b610422613056565b34801561099257600080fd5b506109a7600160a060020a03600435166130d3565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b6109e7615530565b336000908152600760205260408120549080821515610c8a57604080517fe56556a9000000000000000000000000000000000000000000000000000000008152336004820152905173ae5223564809853e636116a453c79f07311a26849163e56556a99160248083019260209291908290030181600087803b158015610a6c57600080fd5b505af1158015610a80573d6000803e3d6000fd5b505050506040513d6020811015610a9657600080fd5b5051604080517f82e37b2c00000000000000000000000000000000000000000000000000000000815260048101839052905191945073ae5223564809853e636116a453c79f07311a2684916382e37b2c916024808201926020929091908290030181600087803b158015610b0957600080fd5b505af1158015610b1d573d6000803e3d6000fd5b505050506040513d6020811015610b3357600080fd5b5051604080517fe3c08adf00000000000000000000000000000000000000000000000000000000815260048101869052905191935073ae5223564809853e636116a453c79f07311a26849163e3c08adf916024808201926020929091908290030181600087803b158015610ba657600080fd5b505af1158015610bba573d6000803e3d6000fd5b505050506040513d6020811015610bd057600080fd5b505133600081815260076020908152604080832088905587835260099091529020805473ffffffffffffffffffffffffffffffffffffffff1916909117905590508115610c59576000828152600860209081526040808320869055858352600982528083206001908101869055600b8352818420868552909252909120805460ff191690911790555b8015801590610c685750828114155b15610c825760008381526009602052604090206006018190555b845160010185525b50929392505050565b6005546002546000828152600c602052604090206004015442910181118015610cfe57506000828152600c602052604090206002015481111580610cfe57506000828152600c602052604090206002015481118015610cfe57506000828152600c6020526040902054155b15610d1657610d118287348888886131a8565b610ec5565b6000828152600c602052604090206002015481118015610d4857506000828152600c602052604090206003015460ff16155b15610e90576000828152600c60205260409020600301805460ff19166001179055610d72836136f9565b925080670de0b6b3a764000002836000015101836000018181525050858360200151018360200181815250507fa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a3360096000898152602001908152602001600020600101543486600001518760200151886040015189606001518a608001518b60a001518c60c001518d60e00151604051808c600160a060020a0316600160a060020a031681526020018b600019166000191681526020018a815260200189815260200188815260200187600160a060020a0316600160a060020a0316815260200186600019166000191681526020018581526020018481526020018381526020018281526020019b50505050505050505050505060405180910390a15b600086815260096020526040902060030154610eb2903463ffffffff613b8a16565b6000878152600960205260409020600301555b505050505050565b6005546002546000828152600c602052604081206004015490929142910181118015610f3b57506000828152600c602052604090206002015481111580610f3b57506000828152600c602052604090206002015481118015610f3b57506000828152600c6020526040902054155b15610f83576000828152600c6020526040902060050154610f7c90670de0b6b3a764000090610f70908263ffffffff613b8a16565b9063ffffffff613beb16565b9250610f8d565b6544364c5bb00092505b505090565b60408051808201909152601581527f466f6d6f455443204c6f6e67204f6666696369616c0000000000000000000000602082015281565b610fd1615530565b601054600090819060ff161515600114611037576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155ca833981519152604482015260008051602061558a833981519152606482015290519081900360840190fd5b33803b801561107e576040805160e560020a62461bcd028152602060048201526011602482015260008051602061560a833981519152604482015290519081900360640190fd5b85633b9aca008110156110d6576040805160e560020a62461bcd02815260206004820152602160248201526000805160206155aa833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af6800000811115611126576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206155ea833981519152604482015290519081900360640190fd5b336000908152600760205260409020549450881580611155575060008581526009602052604090206001015489145b156111735760008581526009602052604090206006015493506111b2565b60008981526008602090815260408083205488845260099092529091206006015490945084146111b25760008581526009602052604090206006018490555b6111bb88613c18565b97506111ca85858a8a8a613c3d565b505050505050505050565b73e0e1b1e5d63e0ead6bfefcfc5a9dce543913cc1533148061120a5750738b4da1827932d71759687f925d17f81fc94e3a9d33145b806112285750738e0d985f3ec1857bec39b76aaabdea6b31b67d5333145b806112465750737ac74fcc1a71b106f12c55ee8f802c9f672ce40c33145b80611264575073f39e044e1ab204460e06e87c6dca2c6319fc69e333145b15156112ba576040805160e560020a62461bcd02815260206004820152601b60248201527f6f6e6c79207465616d206a7573742063616e2061637469766174650000000000604482015290519081900360640190fd5b600054600160a060020a0316151561131c576040805160e560020a62461bcd02815260206004820181905260248201527f6d757374206c696e6b20746f206f7468657220466f6d6f455443206669727374604482015290519081900360640190fd5b60105460ff1615611377576040805160e560020a62461bcd02815260206004820152601960248201527f466f6d6f45544320616c72656164792061637469766174656400000000000000604482015290519081900360640190fd5b6010805460ff1916600190811790915560058190556002548154600092909252600c602052429091019081037fd421a5181c571bba3f01190c922c3b2a896fc1d84e86c9f17ac10e67ebef8b6055610e10017fd421a5181c571bba3f01190c922c3b2a896fc1d84e86c9f17ac10e67ebef8b5e55565b60076020526000908152604090205481565b60045481565b600c60205260009081526040902080546001820154600283015460038401546004850154600586015460068701546007880154600889015460098a0154600a8b0154600b909b0154999a9899979860ff909716979596949593949293919290918c565b600b60209081526000928352604080842090915290825290205460ff1681565b600e602052600090815260409020805460019091015482565b60086020526000908152604090205481565b6114bb615530565b60105460009060ff16151560011461151f576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155ca833981519152604482015260008051602061558a833981519152606482015290519081900360840190fd5b33803b8015611566576040805160e560020a62461bcd028152602060048201526011602482015260008051602061560a833981519152604482015290519081900360640190fd5b84633b9aca008110156115be576040805160e560020a62461bcd02815260206004820152602160248201526000805160206155aa833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af680000081111561160e576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206155ea833981519152604482015290519081900360640190fd5b33600090815260076020526040902054935087158061162c57508388145b1561164a576000848152600960205260409020600601549750611677565b60008481526009602052604090206006015488146116775760008481526009602052604090206006018890555b61168087613c18565b965061168f8489898989613c3d565b5050505050505050565b6000806000806116a7615530565b60105460ff161515600114611708576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155ca833981519152604482015260008051602061558a833981519152606482015290519081900360840190fd5b33803b801561174f576040805160e560020a62461bcd028152602060048201526011602482015260008051602061560a833981519152604482015290519081900360640190fd5b60055433600090815260076020908152604080832054848452600c9092529091206002015491985042975095508611801561179c57506000878152600c602052604090206003015460ff16155b80156117b557506000878152600c602052604090205415155b1561195b576000878152600c60205260409020600301805460ff191660011790556117df836136f9565b92506117ea85613e59565b9350600084111561183b57600085815260096020526040808220549051600160a060020a039091169186156108fc02918791818181858888f19350505050158015611839573d6000803e3d6000fd5b505b85670de0b6b3a764000002836000015101836000018181525050848360200151018360200181815250507f0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc3360096000888152602001908152602001600020600101548686600001518760200151886040015189606001518a608001518b60a001518c60c001518d60e00151604051808c600160a060020a0316600160a060020a031681526020018b600019166000191681526020018a815260200189815260200188815260200187600160a060020a0316600160a060020a0316815260200186600019166000191681526020018581526020018481526020018381526020018281526020019b50505050505050505050505060405180910390a1611a11565b61196485613e59565b935060008411156119b557600085815260096020526040808220549051600160a060020a039091169186156108fc02918791818181858888f193505050501580156119b3573d6000803e3d6000fd5b505b6000858152600960209081526040918290206001015482513381529182015280820186905260608101889052905186917f8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a919081900360800190a25b50505050505050565b6000808080808033803b8015611a68576040805160e560020a62461bcd028152602060048201526011602482015260008051602061560a833981519152604482015290519081900360640190fd5b611a718b613ee0565b604080517faa4d490b000000000000000000000000000000000000000000000000000000008152336004820181905260248201849052600160a060020a038e1660448301528c151560648301528251939b50995034985073ae5223564809853e636116a453c79f07311a26849263aa4d490b928a926084808201939182900301818588803b158015611b0257600080fd5b505af1158015611b16573d6000803e3d6000fd5b50505050506040513d6040811015611b2d57600080fd5b508051602091820151600160a060020a03808b1660008181526007865260408082205485835260098852918190208054600190910154825188151581529889018790529416878201526060870193909352608086018c90524260a0870152915193995091975095508a92909186917fdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442919081900360c00190a45050505050505050505050565b3373ae5223564809853e636116a453c79f07311a268414611c64576040805160e560020a62461bcd02815260206004820152602760248201527f796f7572206e6f7420706c617965724e616d657320636f6e74726163742e2e2e60448201527f20686d6d6d2e2e00000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0383166000908152600760205260409020548414611c9f57600160a060020a03831660009081526007602052604090208490555b6000828152600860205260409020548414611cc65760008281526008602052604090208490555b600084815260096020526040902054600160a060020a03848116911614611d1c576000848152600960205260409020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385161790555b6000848152600960205260409020600101548214611d495760008481526009602052604090206001018290555b6000848152600960205260409020600601548114611d765760008481526009602052604090206006018190555b6000848152600b6020908152604080832085845290915290205460ff161515611dbe576000848152600b602090815260408083208584529091529020805460ff191660011790555b50505050565b600d60209081526000928352604080842090915290825290205481565b60055481565b6005546000818152600c60205260408120600201549091829182919042118015611e2357506000818152600c602052604090206003015460ff16155b8015611e3c57506000818152600c602052604090205415155b15611f5e576000818152600c6020526040902054851415611f22576006546000828152600c6020526040902060070154611eab91606491611e829163ffffffff6146f316565b811515611e8b57fe5b60008881526009602052604090206002015491900463ffffffff613b8a16565b6000868152600a60209081526040808320858452909152902060020154611f0490611ee690611eda898661476a565b9063ffffffff61483816565b6000888152600960205260409020600301549063ffffffff613b8a16565b60008781526009602052604090206004015491955093509150611f86565b6000858152600960209081526040808320600290810154600a8452828520868652909352922090910154611f0490611ee690611eda898661476a565b60008581526009602052604090206002810154600590910154611f0490611ee6908890614898565b509193909250565b6000808080808033803b8015611fdc576040805160e560020a62461bcd028152602060048201526011602482015260008051602061560a833981519152604482015290519081900360640190fd5b611fe58b613ee0565b604080517f745ea0c1000000000000000000000000000000000000000000000000000000008152336004820181905260248201849052604482018e90528c151560648301528251939b50995034985073ae5223564809853e636116a453c79f07311a26849263745ea0c1928a926084808201939182900301818588803b158015611b0257600080fd5b60008060008060008060008060008060008060008060006005549050600c60008281526020019081526020016000206009015481600c600084815260200190815260200160002060050154600c600085815260200190815260200160002060020154600c600086815260200190815260200160002060040154600c600087815260200190815260200160002060070154600c600088815260200190815260200160002060000154600a02600c6000898152602001908152602001600020600101540160096000600c60008b815260200190815260200160002060000154815260200190815260200160002060000160009054906101000a9004600160a060020a031660096000600c60008c815260200190815260200160002060000154815260200190815260200160002060010154600d60008b8152602001908152602001600020600080815260200190815260200160002054600d60008c815260200190815260200160002060006001815260200190815260200160002054600d60008d815260200190815260200160002060006002815260200190815260200160002054600d60008e8152602001908152602001600020600060038152602001908152602001600020546003546103e802600454019e509e509e509e509e509e509e509e509e509e509e509e509e509e5050909192939495969798999a9b9c9d565b612274615530565b601054600090819060ff1615156001146122da576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155ca833981519152604482015260008051602061558a833981519152606482015290519081900360840190fd5b33803b8015612321576040805160e560020a62461bcd028152602060048201526011602482015260008051602061560a833981519152604482015290519081900360640190fd5b85633b9aca00811015612379576040805160e560020a62461bcd02815260206004820152602160248201526000805160206155aa833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af68000008111156123c9576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206155ea833981519152604482015290519081900360640190fd5b336000908152600760205260409020549450600160a060020a03891615806123f95750600160a060020a03891633145b156124175760008581526009602052604090206006015493506111b2565b600160a060020a03891660009081526007602090815260408083205488845260099092529091206006015490945084146111b25760008581526009602052604090206006018490556111bb88613c18565b612470615530565b60105460009060ff1615156001146124d4576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155ca833981519152604482015260008051602061558a833981519152606482015290519081900360840190fd5b33803b801561251b576040805160e560020a62461bcd028152602060048201526011602482015260008051602061560a833981519152604482015290519081900360640190fd5b34633b9aca00811015612573576040805160e560020a62461bcd02815260206004820152602160248201526000805160206155aa833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af68000008111156125c3576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206155ea833981519152604482015290519081900360640190fd5b6125cc856109df565b3360009081526007602052604090205490955093508615806125ed57508387145b1561260b576000848152600960205260409020600601549650612638565b60008481526009602052604090206006015487146126385760008481526009602052604090206006018790555b61264186613c18565b9550611a1184888888610c93565b3373ae5223564809853e636116a453c79f07311a2684146126e0576040805160e560020a62461bcd02815260206004820152602760248201527f796f7572206e6f7420706c617965724e616d657320636f6e74726163742e2e2e60448201527f20686d6d6d2e2e00000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000828152600b6020908152604080832084845290915290205460ff161515612728576000828152600b602090815260408083208484529091529020805460ff191660011790555b5050565b6000808080808033803b801561277a576040805160e560020a62461bcd028152602060048201526011602482015260008051602061560a833981519152604482015290519081900360640190fd5b6127838b613ee0565b604080517fc0942dfd000000000000000000000000000000000000000000000000000000008152336004820181905260248201849052604482018e90528c151560648301528251939b50995034985073ae5223564809853e636116a453c79f07311a26849263c0942dfd928a926084808201939182900301818588803b158015611b0257600080fd5b60408051808201909152600481527f4645544300000000000000000000000000000000000000000000000000000000602082015281565b61284b615530565b601054600090819060ff1615156001146128b1576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155ca833981519152604482015260008051602061558a833981519152606482015290519081900360840190fd5b33803b80156128f8576040805160e560020a62461bcd028152602060048201526011602482015260008051602061560a833981519152604482015290519081900360640190fd5b34633b9aca00811015612950576040805160e560020a62461bcd02815260206004820152602160248201526000805160206155aa833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af68000008111156129a0576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206155ea833981519152604482015290519081900360640190fd5b6129a9866109df565b336000908152600760205260409020549096509450600160a060020a03881615806129dc5750600160a060020a03881633145b156129fa576000858152600960205260409020600601549350612a43565b600160a060020a0388166000908152600760209081526040808320548884526009909252909120600601549094508414612a435760008581526009602052604090206006018490555b612a4c87613c18565b965061168f85858989610c93565b600a60209081526000928352604080842090915290825290208054600182015460028301546003909301549192909184565b612a94615530565b601054600090819060ff161515600114612afa576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155ca833981519152604482015260008051602061558a833981519152606482015290519081900360840190fd5b33803b8015612b41576040805160e560020a62461bcd028152602060048201526011602482015260008051602061560a833981519152604482015290519081900360640190fd5b34633b9aca00811015612b99576040805160e560020a62461bcd02815260206004820152602160248201526000805160206155aa833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af6800000811115612be9576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206155ea833981519152604482015290519081900360640190fd5b612bf2866109df565b336000908152600760205260409020549096509450871580612c24575060008581526009602052604090206001015488145b15612c42576000858152600960205260409020600601549350612a43565b6000888152600860209081526040808320548884526009909252909120600601549094508414612a43576000858152600960205260409020600601849055612a4c87613c18565b73e0e1b1e5d63e0ead6bfefcfc5a9dce543913cc15331480612cbe575073270e87f1e481b9509f8bff4229e736b9758b1c1433145b80612cdc57507364f6fbffb5b064af32758c79081c03b2f1ba3b1633145b80612cfa57507362148c76ee3ea4f5e92f1e87e361bf903fc4a1ad33145b80612d185750737e9312307cf7493256d5c0574e9dc81c4bebe5d033145b1515612d6e576040805160e560020a62461bcd02815260206004820152601b60248201527f6f6e6c79207465616d206a7573742063616e2061637469766174650000000000604482015290519081900360640190fd5b600054600160a060020a031615612dcf576040805160e560020a62461bcd02815260206004820152601f60248201527f73696c6c79206465762c20796f7520616c726561647920646964207468617400604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600f602052600090815260409020805460019091015482565b6005546000818152600c60205260408120600201549091904290811015612e9d576002546000838152600c602052604090206004015401811115612e77576000828152600c6020526040902060020154610f7c908263ffffffff61483816565b6002546000838152600c6020526040902060040154610f7c91018263ffffffff61483816565b60009250610f8d565b6002546000838152600c6020526040812060040154909142910181118015612f1057506000848152600c602052604090206002015481111580612f1057506000848152600c602052604090206002015481118015612f1057506000848152600c6020526040902054155b15612f3e576000848152600c6020526040902060060154612f37908463ffffffff6148f516565b9150612f47565b612f3783614916565b5092915050565b6005546002546000828152600c602052604081206004015490929142910181118015612fbc57506000828152600c602052604090206002015481111580612fbc57506000828152600c602052604090206002015481118015612fbc57506000828152600c6020526040902054155b15612ff0576000828152600c6020526040902060050154612fe9908590610f70908263ffffffff613b8a16565b9250612ff9565b612fe98461498e565b5050919050565b60105460ff1681565b60035481565b6009602052600090815260409020805460018201546002830154600384015460048501546005860154600690960154600160a060020a039095169593949293919290919087565b6005546001016000818152600c602052604090206007015461307e903463ffffffff613b8a16565b6000828152600c6020908152604091829020600701929092558051838152349281019290925280517f74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c9281900390910190a150565b6000806000806000806000806000600554915050600160a060020a0389166000908152600760209081526040808320548084526009808452828520600180820154600a87528588208989528752948720015495839052935260028301546005909301549093849390916131699061314b908690614898565b6000878152600960205260409020600301549063ffffffff613b8a16565b600095865260096020908152604080882060040154600a83528189209989529890915290952054939e929d50909b509950919750919550909350915050565b6000858152600a6020908152604080832089845290915281206001015481908190819015156131de576131db89866149fb565b94505b60008a8152600c602052604090206006015468056bc75e2d6310000011801561323857506000898152600a602090815260408083208d8452909152902054670de0b6b3a764000090613236908a63ffffffff613b8a16565b115b156132bf576000898152600a602090815260408083208d845290915290205461327090670de0b6b3a76400009063ffffffff61483816565b9350613282888563ffffffff61483816565b60008a8152600960205260409020600301549093506132a7908463ffffffff613b8a16565b60008a81526009602052604090206003015592965086925b633b9aca008811156136ed5760008a8152600c60205260409020600601546132ed908963ffffffff6148f516565b9150670de0b6b3a7640000821061336457613308828b614a5a565b60008a8152600c6020526040902054891461332f5760008a8152600c602052604090208990555b60008a8152600c6020526040902060010154861461335c5760008a8152600c602052604090206001018690555b845160640185525b67016345785d8a000088106135a457600480546001019055613384614b38565b1515600114156135a457678ac7230489e800008810613425576003546064906133b490604b63ffffffff6146f316565b8115156133bd57fe5b60008b81526009602052604090206002015491900491506133e4908263ffffffff613b8a16565b60008a815260096020526040902060020155600354613409908263ffffffff61483816565b60035584516d0eca8847c4129106ce8300000000018552613579565b670de0b6b3a764000088101580156134445750678ac7230489e8000088105b156134d15760035460649061346090603263ffffffff6146f316565b81151561346957fe5b60008b8152600960205260409020600201549190049150613490908263ffffffff613b8a16565b60008a8152600960205260409020600201556003546134b5908263ffffffff61483816565b60035584516d09dc5ada82b70b59df0200000000018552613579565b67016345785d8a000088101580156134f05750670de0b6b3a764000088105b156135795760035460649061350c90601963ffffffff6146f316565b81151561351557fe5b60008b815260096020526040902060020154919004915061353c908263ffffffff613b8a16565b60008a815260096020526040902060020155600354613561908263ffffffff61483816565b60035584516d0eca8847c4129106ce83000000000185525b84516d314dc6448d9338c15b0a000000008202016c7e37be2022c0914b268000000001855260006004555b60045485516103e89091020185526000898152600a602090815260408083208d84529091529020600101546135e090839063ffffffff613b8a16565b60008a8152600a602090815260408083208e8452909152902060018101919091555461360d908990613b8a565b60008a8152600a602090815260408083208e8452825280832093909355600c9052206005015461364490839063ffffffff613b8a16565b60008b8152600c6020526040902060058101919091556006015461366f90899063ffffffff613b8a16565b60008b8152600c6020908152604080832060060193909355600d8152828220898352905220546136a690899063ffffffff613b8a16565b60008b8152600d602090815260408083208a84529091529020556136ce8a8a8a8a8a8a614d4f565b94506136de8a8a8a89868a615025565b94506136ed89878a8589615193565b50505050505050505050565b613701615530565b6005546000818152600c602052604081208054600182015460079092015460065491939091819081908190819081908190606490613746908a9063ffffffff6146f316565b81151561374f57fe5b04965060328860008b8152600f6020526040902054919004965060649061377d908a9063ffffffff6146f316565b81151561378657fe5b60008b8152600f602052604090206001015491900495506064906137b1908a9063ffffffff6146f316565b8115156137ba57fe5b0493506137d584611eda87818a818e8e63ffffffff61483816565b60008c8152600c602052604090206005015490935061380286670de0b6b3a764000063ffffffff6146f316565b81151561380b57fe5b60008d8152600c6020526040902060050154919004925061385990670de0b6b3a76400009061384190859063ffffffff6146f316565b81151561384a57fe5b8791900463ffffffff61483816565b9050600081111561388957613874858263ffffffff61483816565b9450613886838263ffffffff613b8a16565b92505b60008a8152600960205260409020600201546138ac90889063ffffffff613b8a16565b600960008c8152602001908152602001600020600201819055507386ad1c4b193bd892b5e7d8c43e656f5f3b35559b600160a060020a03168660405180807f6465706f736974282900000000000000000000000000000000000000000000008152506009019050604051809103902060e060020a9004906040518263ffffffff1660e060020a02815260040160006040518083038185885af19350505050151561396757613960848763ffffffff613b8a16565b9350600095505b60008b8152600c602052604090206008015461398a90839063ffffffff613b8a16565b60008c8152600c6020526040812060080191909155841115613a125773b0827c27361f714bc11eaff7401bf7776c9dcab5600160a060020a031663d0e30db0856040518263ffffffff1660e060020a0281526004016000604051808303818588803b1580156139f857600080fd5b505af1158015613a0c573d6000803e3d6000fd5b50505050505b600c60008c815260200190815260200160002060020154620f4240028d60000151018d60000181815250508867016345785d8a0000028a6a52b7d2dcc80cd2e4000000028e6020015101018d6020018181525050600960008b815260200190815260200160002060000160009054906101000a9004600160a060020a03168d60400190600160a060020a03169081600160a060020a031681525050600960008b8152602001908152602001600020600101548d606001906000191690816000191681525050868d6080018181525050848d60e0018181525050838d60c0018181525050828d60a00181815250506005600081548092919060010191905055508a806001019b505042600c60008d815260200190815260200160002060040181905550613b5b600254613b4f610e1042613b8a90919063ffffffff16565b9063ffffffff613b8a16565b60008c8152600c6020526040902060028101919091556007018390558c9b505050505050505050505050919050565b81810182811015613be5576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d61746820616464206661696c656400000000000000000000000000604482015290519081900360640190fd5b92915050565b6000613c11613c08613c03858563ffffffff61483816565b61498e565b611eda8561498e565b9392505050565b600080821080613c285750600382115b15613c3557506002613c38565b50805b919050565b6005546002546000828152600c602052604090206004015442910181118015613ca857506000828152600c602052604090206002015481111580613ca857506000828152600c602052604090206002015481118015613ca857506000828152600c6020526040902054155b15613cdf57613cba84611eda89613e59565b600088815260096020526040902060030155613cda8288868989886131a8565b611a11565b6000828152600c602052604090206002015481118015613d1157506000828152600c602052604090206003015460ff16155b15611a11576000828152600c60205260409020600301805460ff19166001179055613d3b836136f9565b925080670de0b6b3a764000002836000015101836000018181525050868360200151018360200181815250507f88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd33600960008a815260200190815260200160002060010154856000015186602001518760400151886060015189608001518a60a001518b60c001518c60e00151604051808b600160a060020a0316600160a060020a031681526020018a6000191660001916815260200189815260200188815260200187600160a060020a0316600160a060020a0316815260200186600019166000191681526020018581526020018481526020018381526020018281526020019a505050505050505050505060405180910390a150505050505050565b6000818152600960205260408120600501548190613e78908490615301565b600083815260096020526040902060048101546003820154600290920154613eaa92613b4f919063ffffffff613b8a16565b90506000811115613ed65760008381526009602052604081206002810182905560038101829055600401555b8091505b50919050565b8051600090829082808060208411801590613efb5750600084115b1515613f77576040805160e560020a62461bcd02815260206004820152602a60248201527f737472696e67206d757374206265206265747765656e203120616e642033322060448201527f6368617261637465727300000000000000000000000000000000000000000000606482015290519081900360840190fd5b846000815181101515613f8657fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a0214158015613fed57508460018503815181101515613fc557fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a0214155b1515614069576040805160e560020a62461bcd02815260206004820152602560248201527f737472696e672063616e6e6f74207374617274206f7220656e6420776974682060448201527f7370616365000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b84600081518110151561407857fe5b90602001015160f860020a900460f860020a02600160f860020a031916603060f860020a0214156141bb578460018151811015156140b257fe5b90602001015160f860020a900460f860020a02600160f860020a031916607860f860020a021415151561412f576040805160e560020a62461bcd02815260206004820152601b60248201527f737472696e672063616e6e6f7420737461727420776974682030780000000000604482015290519081900360640190fd5b84600181518110151561413e57fe5b90602001015160f860020a900460f860020a02600160f860020a031916605860f860020a02141515156141bb576040805160e560020a62461bcd02815260206004820152601b60248201527f737472696e672063616e6e6f7420737461727420776974682030580000000000604482015290519081900360640190fd5b600091505b8382101561468b5784517f4000000000000000000000000000000000000000000000000000000000000000908690849081106141f857fe5b90602001015160f860020a900460f860020a02600160f860020a03191611801561426c575084517f5b000000000000000000000000000000000000000000000000000000000000009086908490811061424d57fe5b90602001015160f860020a900460f860020a02600160f860020a031916105b156142d957848281518110151561427f57fe5b90602001015160f860020a900460f860020a0260f860020a900460200160f860020a0285838151811015156142b057fe5b906020010190600160f860020a031916908160001a9053508215156142d457600192505b614680565b84828151811015156142e757fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a0214806143b7575084517f60000000000000000000000000000000000000000000000000000000000000009086908490811061434357fe5b90602001015160f860020a900460f860020a02600160f860020a0319161180156143b7575084517f7b000000000000000000000000000000000000000000000000000000000000009086908490811061439857fe5b90602001015160f860020a900460f860020a02600160f860020a031916105b80614461575084517f2f00000000000000000000000000000000000000000000000000000000000000908690849081106143ed57fe5b90602001015160f860020a900460f860020a02600160f860020a031916118015614461575084517f3a000000000000000000000000000000000000000000000000000000000000009086908490811061444257fe5b90602001015160f860020a900460f860020a02600160f860020a031916105b15156144dd576040805160e560020a62461bcd02815260206004820152602260248201527f737472696e6720636f6e7461696e7320696e76616c696420636861726163746560448201527f7273000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b84828151811015156144eb57fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a0214156145ca57848260010181518110151561452757fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a02141515156145ca576040805160e560020a62461bcd02815260206004820152602860248201527f737472696e672063616e6e6f7420636f6e7461696e20636f6e7365637574697660448201527f6520737061636573000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b82158015614676575084517f30000000000000000000000000000000000000000000000000000000000000009086908490811061460357fe5b90602001015160f860020a900460f860020a02600160f860020a0319161080614676575084517f39000000000000000000000000000000000000000000000000000000000000009086908490811061465757fe5b90602001015160f860020a900460f860020a02600160f860020a031916115b1561468057600192505b6001909101906141c0565b6001831515146146e5576040805160e560020a62461bcd02815260206004820152601d60248201527f737472696e672063616e6e6f74206265206f6e6c79206e756d62657273000000604482015290519081900360640190fd5b505050506020015192915050565b600082151561470457506000613be5565b5081810281838281151561471457fe5b0414613be5576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d617468206d756c206661696c656400000000000000000000000000604482015290519081900360640190fd5b6000828152600a602090815260408083208484528252808320600190810154600c8085528386206005810154938101548752600f8652938620548787529452600790920154670de0b6b3a764000093614827939261481b9290916147f29187916064916147dc9163ffffffff6146f316565b8115156147e557fe5b049063ffffffff6146f316565b8115156147fb57fe5b6000888152600c602052604090206008015491900463ffffffff613b8a16565b9063ffffffff6146f316565b81151561483057fe5b049392505050565b600082821115614892576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d61746820737562206661696c656400000000000000000000000000604482015290519081900360640190fd5b50900390565b6000828152600a6020908152604080832084845282528083206002810154600190910154600c90935290832060080154613c1192670de0b6b3a7640000916148df916146f3565b8115156148e857fe5b049063ffffffff61483816565b6000613c1161490384614916565b611eda614916868663ffffffff613b8a16565b60006309502f9061497e6d03b2a1d15167e7c5699bfde00000611eda6149797a0dac7055469777a6122ee4310dd6c14410500f2904840000000000613b4f6b01027e72f1f128130880000061481b8a670de0b6b3a764000063ffffffff6146f316565b615398565b81151561498757fe5b0492915050565b60006149a1670de0b6b3a76400006153eb565b61497e60026149d46149c186670de0b6b3a764000063ffffffff6146f316565b65886c8f6730709063ffffffff6146f316565b8115156149dd57fe5b04613b4f6149ea866153eb565b6304a817c89063ffffffff6146f316565b614a03615530565b60008381526009602052604090206005015415614a3757600083815260096020526040902060050154614a37908490615301565b506005805460009384526009602052604090932001919091558051600a01815290565b6000818152600c602052604081206002015442919082118015614a8957506000838152600c6020526040902054155b15614aad57614aa682613b4f601e670de0b6b3a7640000886147e5565b9050614ada565b6000838152600c6020526040902060020154614ad790613b4f601e670de0b6b3a7640000886147e5565b90505b614aed620151808363ffffffff613b8a16565b811015614b0d576000838152600c60205260409020600201819055611dbe565b614b20620151808363ffffffff613b8a16565b6000848152600c602052604090206002015550505050565b600080614ca943613b4f42336040516020018082600160a060020a0316600160a060020a03166c010000000000000000000000000281526014019150506040516020818303038152906040526040518082805190602001908083835b60208310614bb35780518252601f199092019160209182019101614b94565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912092505050811515614be957fe5b04613b4f45613b4f42416040516020018082600160a060020a0316600160a060020a03166c010000000000000000000000000281526014019150506040516020818303038152906040526040518082805190602001908083835b60208310614c625780518252601f199092019160209182019101614c43565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912092505050811515614c9857fe5b04613b4f424463ffffffff613b8a16565b604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b60208310614cf75780518252601f199092019160209182019101614cd8565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912060045490945092506103e89150839050046103e80282031015614d465760019150614d4b565b600091505b5090565b614d57615530565b60008080806032890493507386ad1c4b193bd892b5e7d8c43e656f5f3b35559b600160a060020a03168460405180807f6465706f736974282900000000000000000000000000000000000000000000008152506009019050604051809103902060e060020a9004906040518263ffffffff1660e060020a02815260040160006040518083038185885af193505050501515614df25760009392505b60008054604080517fed78cf4a000000000000000000000000000000000000000000000000000000008152905160648d049550600160a060020a039092169263ed78cf4a928692600480820193929182900301818588803b158015614e5657600080fd5b505af1158015614e6a573d6000803e3d6000fd5b5050505050600a89811515614e7b57fe5b049050898814158015614e9e575060008881526009602052604090206001015415155b15614f3e57600088815260096020526040902060040154614ec690829063ffffffff613b8a16565b600089815260096020908152604091829020600481019390935582546001909301548251600160a060020a03909416845290830152818101839052426060830152518b918d918b917f590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331919081900360800190a4614f42565b8092505b6000878152600e6020526040902060010154614f8490606490614f6c908c9063ffffffff6146f316565b811515614f7557fe5b8591900463ffffffff613b8a16565b925060008311156150165773b0827c27361f714bc11eaff7401bf7776c9dcab5600160a060020a031663d0e30db0846040518263ffffffff1660e060020a0281526004016000604051808303818588803b158015614fe157600080fd5b505af1158015614ff5573d6000803e3d6000fd5b50505060c0880151615010925085915063ffffffff613b8a16565b60c08701525b50939998505050505050505050565b61502d615530565b6000848152600e6020526040812054819081908190606490615056908b9063ffffffff6146f316565b81151561505f57fe5b04935060648904925061507d83600354613b8a90919063ffffffff16565b6003556000888152600e60205260409020600101546150ee906150e1906064906150ae908d9063ffffffff6146f316565b8115156150b757fe5b0460646150cb8d600e63ffffffff6146f316565b8115156150d457fe5b049063ffffffff613b8a16565b8a9063ffffffff61483816565b9850615100898563ffffffff61483816565b915061510e8b8b868a6153f7565b9050600081111561512c57615129848263ffffffff61483816565b93505b60008b8152600c602052604090206007015461515290613b4f848463ffffffff613b8a16565b60008c8152600c602052604090206007015560e086015161517a90859063ffffffff613b8a16565b60e0870152506101008501525091979650505050505050565b836c01431e0fae6d7217caa00000000242670de0b6b3a76400000282600001510101816000018181525050600554751aba4714957d300d0e549208b31adb100000000000000285826020015101018160200181815250507f500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c7468160000151826020015160096000898152602001908152602001600020600101543387878760400151886060015189608001518a60a001518b60c001518c60e001518d6101000151600354604051808f81526020018e81526020018d600019166000191681526020018c600160a060020a0316600160a060020a031681526020018b81526020018a815260200189600160a060020a0316600160a060020a0316815260200188600019166000191681526020018781526020018681526020018581526020018481526020018381526020018281526020019e50505050505050505050505050505060405180910390a15050505050565b600061530d8383614898565b905060008111156153935760008381526009602052604090206003015461533b90829063ffffffff613b8a16565b600084815260096020908152604080832060030193909355600a81528282208583529052206002015461537590829063ffffffff613b8a16565b6000848152600a602090815260408083208684529091529020600201555b505050565b60008060026153a8846001613b8a565b8115156153b157fe5b0490508291505b81811015613eda5780915060026153da82858115156153d357fe5b0483613b8a565b8115156153e357fe5b0490506153b8565b6000613be582836146f3565b6000848152600c60205260408120600501548190819061542586670de0b6b3a764000063ffffffff6146f316565b81151561542e57fe5b6000898152600c6020526040902060080154919004925061545690839063ffffffff613b8a16565b6000888152600c6020526040902060080155670de0b6b3a7640000615481838663ffffffff6146f316565b81151561548a57fe5b6000888152600a602090815260408083208c8452825280832060020154600c909252909120600801549290910492506154dd91613b4f908490670de0b6b3a7640000906148df908a63ffffffff6146f316565b6000878152600a602090815260408083208b8452825280832060020193909355600c9052206005015461552590670de0b6b3a76400009061384190859063ffffffff6146f316565b979650505050505050565b6101206040519081016040528060008152602001600081526020016000600160a060020a03168152602001600080191681526020016000815260200160008152602001600081526020016000815260200160008152509056006e20646973636f72640000000000000000000000000000000000000000000000706f636b6574206c696e743a206e6f7420612076616c69642063757272656e63697473206e6f74207265616479207965742e2020636865636b203f65746120696e6f20766974616c696b2c206e6f000000000000000000000000000000000000736f7272792068756d616e73206f6e6c79000000000000000000000000000000a165627a7a723058208b119895b042efa0f110971e80ecbf8defb290893651370ba63540e10bea162e0029`

// DeployFomoETClong deploys a new Ethereum contract, binding an instance of FomoETClong to it.
func DeployFomoETClong(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FomoETClong, error) {
	parsed, err := abi.JSON(strings.NewReader(FomoETClongABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FomoETClongBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FomoETClong{FomoETClongCaller: FomoETClongCaller{contract: contract}, FomoETClongTransactor: FomoETClongTransactor{contract: contract}, FomoETClongFilterer: FomoETClongFilterer{contract: contract}}, nil
}

// FomoETClong is an auto generated Go binding around an Ethereum contract.
type FomoETClong struct {
	FomoETClongCaller     // Read-only binding to the contract
	FomoETClongTransactor // Write-only binding to the contract
	FomoETClongFilterer   // Log filterer for contract events
}

// FomoETClongCaller is an auto generated read-only Go binding around an Ethereum contract.
type FomoETClongCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FomoETClongTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FomoETClongTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FomoETClongFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FomoETClongFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FomoETClongSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FomoETClongSession struct {
	Contract     *FomoETClong      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FomoETClongCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FomoETClongCallerSession struct {
	Contract *FomoETClongCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FomoETClongTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FomoETClongTransactorSession struct {
	Contract     *FomoETClongTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FomoETClongRaw is an auto generated low-level Go binding around an Ethereum contract.
type FomoETClongRaw struct {
	Contract *FomoETClong // Generic contract binding to access the raw methods on
}

// FomoETClongCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FomoETClongCallerRaw struct {
	Contract *FomoETClongCaller // Generic read-only contract binding to access the raw methods on
}

// FomoETClongTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FomoETClongTransactorRaw struct {
	Contract *FomoETClongTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFomoETClong creates a new instance of FomoETClong, bound to a specific deployed contract.
func NewFomoETClong(address common.Address, backend bind.ContractBackend) (*FomoETClong, error) {
	contract, err := bindFomoETClong(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FomoETClong{FomoETClongCaller: FomoETClongCaller{contract: contract}, FomoETClongTransactor: FomoETClongTransactor{contract: contract}, FomoETClongFilterer: FomoETClongFilterer{contract: contract}}, nil
}

// NewFomoETClongCaller creates a new read-only instance of FomoETClong, bound to a specific deployed contract.
func NewFomoETClongCaller(address common.Address, caller bind.ContractCaller) (*FomoETClongCaller, error) {
	contract, err := bindFomoETClong(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FomoETClongCaller{contract: contract}, nil
}

// NewFomoETClongTransactor creates a new write-only instance of FomoETClong, bound to a specific deployed contract.
func NewFomoETClongTransactor(address common.Address, transactor bind.ContractTransactor) (*FomoETClongTransactor, error) {
	contract, err := bindFomoETClong(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FomoETClongTransactor{contract: contract}, nil
}

// NewFomoETClongFilterer creates a new log filterer instance of FomoETClong, bound to a specific deployed contract.
func NewFomoETClongFilterer(address common.Address, filterer bind.ContractFilterer) (*FomoETClongFilterer, error) {
	contract, err := bindFomoETClong(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FomoETClongFilterer{contract: contract}, nil
}

// bindFomoETClong binds a generic wrapper to an already deployed contract.
func bindFomoETClong(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FomoETClongABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FomoETClong *FomoETClongRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FomoETClong.Contract.FomoETClongCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FomoETClong *FomoETClongRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FomoETClong.Contract.FomoETClongTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FomoETClong *FomoETClongRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FomoETClong.Contract.FomoETClongTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FomoETClong *FomoETClongCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FomoETClong.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FomoETClong *FomoETClongTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FomoETClong.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FomoETClong *FomoETClongTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FomoETClong.Contract.contract.Transact(opts, method, params...)
}

// Activated is a free data retrieval call binding the contract method 0xd53b2679.
//
// Solidity: function activated_() constant returns(bool)
func (_FomoETClong *FomoETClongCaller) Activated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "activated_")
	return *ret0, err
}

// Activated is a free data retrieval call binding the contract method 0xd53b2679.
//
// Solidity: function activated_() constant returns(bool)
func (_FomoETClong *FomoETClongSession) Activated() (bool, error) {
	return _FomoETClong.Contract.Activated(&_FomoETClong.CallOpts)
}

// Activated is a free data retrieval call binding the contract method 0xd53b2679.
//
// Solidity: function activated_() constant returns(bool)
func (_FomoETClong *FomoETClongCallerSession) Activated() (bool, error) {
	return _FomoETClong.Contract.Activated(&_FomoETClong.CallOpts)
}

// AirDropPot is a free data retrieval call binding the contract method 0xd87574e0.
//
// Solidity: function airDropPot_() constant returns(uint256)
func (_FomoETClong *FomoETClongCaller) AirDropPot(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "airDropPot_")
	return *ret0, err
}

// AirDropPot is a free data retrieval call binding the contract method 0xd87574e0.
//
// Solidity: function airDropPot_() constant returns(uint256)
func (_FomoETClong *FomoETClongSession) AirDropPot() (*big.Int, error) {
	return _FomoETClong.Contract.AirDropPot(&_FomoETClong.CallOpts)
}

// AirDropPot is a free data retrieval call binding the contract method 0xd87574e0.
//
// Solidity: function airDropPot_() constant returns(uint256)
func (_FomoETClong *FomoETClongCallerSession) AirDropPot() (*big.Int, error) {
	return _FomoETClong.Contract.AirDropPot(&_FomoETClong.CallOpts)
}

// AirDropTracker is a free data retrieval call binding the contract method 0x11a09ae7.
//
// Solidity: function airDropTracker_() constant returns(uint256)
func (_FomoETClong *FomoETClongCaller) AirDropTracker(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "airDropTracker_")
	return *ret0, err
}

// AirDropTracker is a free data retrieval call binding the contract method 0x11a09ae7.
//
// Solidity: function airDropTracker_() constant returns(uint256)
func (_FomoETClong *FomoETClongSession) AirDropTracker() (*big.Int, error) {
	return _FomoETClong.Contract.AirDropTracker(&_FomoETClong.CallOpts)
}

// AirDropTracker is a free data retrieval call binding the contract method 0x11a09ae7.
//
// Solidity: function airDropTracker_() constant returns(uint256)
func (_FomoETClong *FomoETClongCallerSession) AirDropTracker() (*big.Int, error) {
	return _FomoETClong.Contract.AirDropTracker(&_FomoETClong.CallOpts)
}

// CalcKeysReceived is a free data retrieval call binding the contract method 0xce89c80c.
//
// Solidity: function calcKeysReceived(_rID uint256, _etc uint256) constant returns(uint256)
func (_FomoETClong *FomoETClongCaller) CalcKeysReceived(opts *bind.CallOpts, _rID *big.Int, _etc *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "calcKeysReceived", _rID, _etc)
	return *ret0, err
}

// CalcKeysReceived is a free data retrieval call binding the contract method 0xce89c80c.
//
// Solidity: function calcKeysReceived(_rID uint256, _etc uint256) constant returns(uint256)
func (_FomoETClong *FomoETClongSession) CalcKeysReceived(_rID *big.Int, _etc *big.Int) (*big.Int, error) {
	return _FomoETClong.Contract.CalcKeysReceived(&_FomoETClong.CallOpts, _rID, _etc)
}

// CalcKeysReceived is a free data retrieval call binding the contract method 0xce89c80c.
//
// Solidity: function calcKeysReceived(_rID uint256, _etc uint256) constant returns(uint256)
func (_FomoETClong *FomoETClongCallerSession) CalcKeysReceived(_rID *big.Int, _etc *big.Int) (*big.Int, error) {
	return _FomoETClong.Contract.CalcKeysReceived(&_FomoETClong.CallOpts, _rID, _etc)
}

// Fees is a free data retrieval call binding the contract method 0x2ce21999.
//
// Solidity: function fees_( uint256) constant returns(gen uint256, psl uint256)
func (_FomoETClong *FomoETClongCaller) Fees(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Gen *big.Int
	Psl *big.Int
}, error) {
	ret := new(struct {
		Gen *big.Int
		Psl *big.Int
	})
	out := ret
	err := _FomoETClong.contract.Call(opts, out, "fees_", arg0)
	return *ret, err
}

// Fees is a free data retrieval call binding the contract method 0x2ce21999.
//
// Solidity: function fees_( uint256) constant returns(gen uint256, psl uint256)
func (_FomoETClong *FomoETClongSession) Fees(arg0 *big.Int) (struct {
	Gen *big.Int
	Psl *big.Int
}, error) {
	return _FomoETClong.Contract.Fees(&_FomoETClong.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0x2ce21999.
//
// Solidity: function fees_( uint256) constant returns(gen uint256, psl uint256)
func (_FomoETClong *FomoETClongCallerSession) Fees(arg0 *big.Int) (struct {
	Gen *big.Int
	Psl *big.Int
}, error) {
	return _FomoETClong.Contract.Fees(&_FomoETClong.CallOpts, arg0)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() constant returns(uint256)
func (_FomoETClong *FomoETClongCaller) GetBuyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "getBuyPrice")
	return *ret0, err
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() constant returns(uint256)
func (_FomoETClong *FomoETClongSession) GetBuyPrice() (*big.Int, error) {
	return _FomoETClong.Contract.GetBuyPrice(&_FomoETClong.CallOpts)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() constant returns(uint256)
func (_FomoETClong *FomoETClongCallerSession) GetBuyPrice() (*big.Int, error) {
	return _FomoETClong.Contract.GetBuyPrice(&_FomoETClong.CallOpts)
}

// GetCurrentRoundInfo is a free data retrieval call binding the contract method 0x747dff42.
//
// Solidity: function getCurrentRoundInfo() constant returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, address, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_FomoETClong *FomoETClongCaller) GetCurrentRoundInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0  = new(*big.Int)
		ret1  = new(*big.Int)
		ret2  = new(*big.Int)
		ret3  = new(*big.Int)
		ret4  = new(*big.Int)
		ret5  = new(*big.Int)
		ret6  = new(*big.Int)
		ret7  = new(common.Address)
		ret8  = new([32]byte)
		ret9  = new(*big.Int)
		ret10 = new(*big.Int)
		ret11 = new(*big.Int)
		ret12 = new(*big.Int)
		ret13 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
		ret9,
		ret10,
		ret11,
		ret12,
		ret13,
	}
	err := _FomoETClong.contract.Call(opts, out, "getCurrentRoundInfo")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, *ret9, *ret10, *ret11, *ret12, *ret13, err
}

// GetCurrentRoundInfo is a free data retrieval call binding the contract method 0x747dff42.
//
// Solidity: function getCurrentRoundInfo() constant returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, address, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_FomoETClong *FomoETClongSession) GetCurrentRoundInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FomoETClong.Contract.GetCurrentRoundInfo(&_FomoETClong.CallOpts)
}

// GetCurrentRoundInfo is a free data retrieval call binding the contract method 0x747dff42.
//
// Solidity: function getCurrentRoundInfo() constant returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, address, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_FomoETClong *FomoETClongCallerSession) GetCurrentRoundInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FomoETClong.Contract.GetCurrentRoundInfo(&_FomoETClong.CallOpts)
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(_addr address) constant returns(uint256, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_FomoETClong *FomoETClongCaller) GetPlayerInfoByAddress(opts *bind.CallOpts, _addr common.Address) (*big.Int, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _FomoETClong.contract.Call(opts, out, "getPlayerInfoByAddress", _addr)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(_addr address) constant returns(uint256, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_FomoETClong *FomoETClongSession) GetPlayerInfoByAddress(_addr common.Address) (*big.Int, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FomoETClong.Contract.GetPlayerInfoByAddress(&_FomoETClong.CallOpts, _addr)
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(_addr address) constant returns(uint256, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_FomoETClong *FomoETClongCallerSession) GetPlayerInfoByAddress(_addr common.Address) (*big.Int, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FomoETClong.Contract.GetPlayerInfoByAddress(&_FomoETClong.CallOpts, _addr)
}

// GetPlayerVaults is a free data retrieval call binding the contract method 0x63066434.
//
// Solidity: function getPlayerVaults(_pID uint256) constant returns(uint256, uint256, uint256)
func (_FomoETClong *FomoETClongCaller) GetPlayerVaults(opts *bind.CallOpts, _pID *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _FomoETClong.contract.Call(opts, out, "getPlayerVaults", _pID)
	return *ret0, *ret1, *ret2, err
}

// GetPlayerVaults is a free data retrieval call binding the contract method 0x63066434.
//
// Solidity: function getPlayerVaults(_pID uint256) constant returns(uint256, uint256, uint256)
func (_FomoETClong *FomoETClongSession) GetPlayerVaults(_pID *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _FomoETClong.Contract.GetPlayerVaults(&_FomoETClong.CallOpts, _pID)
}

// GetPlayerVaults is a free data retrieval call binding the contract method 0x63066434.
//
// Solidity: function getPlayerVaults(_pID uint256) constant returns(uint256, uint256, uint256)
func (_FomoETClong *FomoETClongCallerSession) GetPlayerVaults(_pID *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _FomoETClong.Contract.GetPlayerVaults(&_FomoETClong.CallOpts, _pID)
}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() constant returns(uint256)
func (_FomoETClong *FomoETClongCaller) GetTimeLeft(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "getTimeLeft")
	return *ret0, err
}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() constant returns(uint256)
func (_FomoETClong *FomoETClongSession) GetTimeLeft() (*big.Int, error) {
	return _FomoETClong.Contract.GetTimeLeft(&_FomoETClong.CallOpts)
}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() constant returns(uint256)
func (_FomoETClong *FomoETClongCallerSession) GetTimeLeft() (*big.Int, error) {
	return _FomoETClong.Contract.GetTimeLeft(&_FomoETClong.CallOpts)
}

// IWantXKeys is a free data retrieval call binding the contract method 0xcf808000.
//
// Solidity: function iWantXKeys(_keys uint256) constant returns(uint256)
func (_FomoETClong *FomoETClongCaller) IWantXKeys(opts *bind.CallOpts, _keys *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "iWantXKeys", _keys)
	return *ret0, err
}

// IWantXKeys is a free data retrieval call binding the contract method 0xcf808000.
//
// Solidity: function iWantXKeys(_keys uint256) constant returns(uint256)
func (_FomoETClong *FomoETClongSession) IWantXKeys(_keys *big.Int) (*big.Int, error) {
	return _FomoETClong.Contract.IWantXKeys(&_FomoETClong.CallOpts, _keys)
}

// IWantXKeys is a free data retrieval call binding the contract method 0xcf808000.
//
// Solidity: function iWantXKeys(_keys uint256) constant returns(uint256)
func (_FomoETClong *FomoETClongCallerSession) IWantXKeys(_keys *big.Int) (*big.Int, error) {
	return _FomoETClong.Contract.IWantXKeys(&_FomoETClong.CallOpts, _keys)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FomoETClong *FomoETClongCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FomoETClong *FomoETClongSession) Name() (string, error) {
	return _FomoETClong.Contract.Name(&_FomoETClong.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FomoETClong *FomoETClongCallerSession) Name() (string, error) {
	return _FomoETClong.Contract.Name(&_FomoETClong.CallOpts)
}

// PIDxAddr is a free data retrieval call binding the contract method 0x10f01eba.
//
// Solidity: function pIDxAddr_( address) constant returns(uint256)
func (_FomoETClong *FomoETClongCaller) PIDxAddr(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "pIDxAddr_", arg0)
	return *ret0, err
}

// PIDxAddr is a free data retrieval call binding the contract method 0x10f01eba.
//
// Solidity: function pIDxAddr_( address) constant returns(uint256)
func (_FomoETClong *FomoETClongSession) PIDxAddr(arg0 common.Address) (*big.Int, error) {
	return _FomoETClong.Contract.PIDxAddr(&_FomoETClong.CallOpts, arg0)
}

// PIDxAddr is a free data retrieval call binding the contract method 0x10f01eba.
//
// Solidity: function pIDxAddr_( address) constant returns(uint256)
func (_FomoETClong *FomoETClongCallerSession) PIDxAddr(arg0 common.Address) (*big.Int, error) {
	return _FomoETClong.Contract.PIDxAddr(&_FomoETClong.CallOpts, arg0)
}

// PIDxName is a free data retrieval call binding the contract method 0x2e19ebdc.
//
// Solidity: function pIDxName_( bytes32) constant returns(uint256)
func (_FomoETClong *FomoETClongCaller) PIDxName(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "pIDxName_", arg0)
	return *ret0, err
}

// PIDxName is a free data retrieval call binding the contract method 0x2e19ebdc.
//
// Solidity: function pIDxName_( bytes32) constant returns(uint256)
func (_FomoETClong *FomoETClongSession) PIDxName(arg0 [32]byte) (*big.Int, error) {
	return _FomoETClong.Contract.PIDxName(&_FomoETClong.CallOpts, arg0)
}

// PIDxName is a free data retrieval call binding the contract method 0x2e19ebdc.
//
// Solidity: function pIDxName_( bytes32) constant returns(uint256)
func (_FomoETClong *FomoETClongCallerSession) PIDxName(arg0 [32]byte) (*big.Int, error) {
	return _FomoETClong.Contract.PIDxName(&_FomoETClong.CallOpts, arg0)
}

// PlyrNames is a free data retrieval call binding the contract method 0x2660316e.
//
// Solidity: function plyrNames_( uint256,  bytes32) constant returns(bool)
func (_FomoETClong *FomoETClongCaller) PlyrNames(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "plyrNames_", arg0, arg1)
	return *ret0, err
}

// PlyrNames is a free data retrieval call binding the contract method 0x2660316e.
//
// Solidity: function plyrNames_( uint256,  bytes32) constant returns(bool)
func (_FomoETClong *FomoETClongSession) PlyrNames(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _FomoETClong.Contract.PlyrNames(&_FomoETClong.CallOpts, arg0, arg1)
}

// PlyrNames is a free data retrieval call binding the contract method 0x2660316e.
//
// Solidity: function plyrNames_( uint256,  bytes32) constant returns(bool)
func (_FomoETClong *FomoETClongCallerSession) PlyrNames(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _FomoETClong.Contract.PlyrNames(&_FomoETClong.CallOpts, arg0, arg1)
}

// PlyrRnds is a free data retrieval call binding the contract method 0xa2bccae9.
//
// Solidity: function plyrRnds_( uint256,  uint256) constant returns(etc uint256, keys uint256, mask uint256, ico uint256)
func (_FomoETClong *FomoETClongCaller) PlyrRnds(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Etc  *big.Int
	Keys *big.Int
	Mask *big.Int
	Ico  *big.Int
}, error) {
	ret := new(struct {
		Etc  *big.Int
		Keys *big.Int
		Mask *big.Int
		Ico  *big.Int
	})
	out := ret
	err := _FomoETClong.contract.Call(opts, out, "plyrRnds_", arg0, arg1)
	return *ret, err
}

// PlyrRnds is a free data retrieval call binding the contract method 0xa2bccae9.
//
// Solidity: function plyrRnds_( uint256,  uint256) constant returns(etc uint256, keys uint256, mask uint256, ico uint256)
func (_FomoETClong *FomoETClongSession) PlyrRnds(arg0 *big.Int, arg1 *big.Int) (struct {
	Etc  *big.Int
	Keys *big.Int
	Mask *big.Int
	Ico  *big.Int
}, error) {
	return _FomoETClong.Contract.PlyrRnds(&_FomoETClong.CallOpts, arg0, arg1)
}

// PlyrRnds is a free data retrieval call binding the contract method 0xa2bccae9.
//
// Solidity: function plyrRnds_( uint256,  uint256) constant returns(etc uint256, keys uint256, mask uint256, ico uint256)
func (_FomoETClong *FomoETClongCallerSession) PlyrRnds(arg0 *big.Int, arg1 *big.Int) (struct {
	Etc  *big.Int
	Keys *big.Int
	Mask *big.Int
	Ico  *big.Int
}, error) {
	return _FomoETClong.Contract.PlyrRnds(&_FomoETClong.CallOpts, arg0, arg1)
}

// Plyr is a free data retrieval call binding the contract method 0xde7874f3.
//
// Solidity: function plyr_( uint256) constant returns(addr address, name bytes32, win uint256, gen uint256, aff uint256, lrnd uint256, laff uint256)
func (_FomoETClong *FomoETClongCaller) Plyr(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr common.Address
	Name [32]byte
	Win  *big.Int
	Gen  *big.Int
	Aff  *big.Int
	Lrnd *big.Int
	Laff *big.Int
}, error) {
	ret := new(struct {
		Addr common.Address
		Name [32]byte
		Win  *big.Int
		Gen  *big.Int
		Aff  *big.Int
		Lrnd *big.Int
		Laff *big.Int
	})
	out := ret
	err := _FomoETClong.contract.Call(opts, out, "plyr_", arg0)
	return *ret, err
}

// Plyr is a free data retrieval call binding the contract method 0xde7874f3.
//
// Solidity: function plyr_( uint256) constant returns(addr address, name bytes32, win uint256, gen uint256, aff uint256, lrnd uint256, laff uint256)
func (_FomoETClong *FomoETClongSession) Plyr(arg0 *big.Int) (struct {
	Addr common.Address
	Name [32]byte
	Win  *big.Int
	Gen  *big.Int
	Aff  *big.Int
	Lrnd *big.Int
	Laff *big.Int
}, error) {
	return _FomoETClong.Contract.Plyr(&_FomoETClong.CallOpts, arg0)
}

// Plyr is a free data retrieval call binding the contract method 0xde7874f3.
//
// Solidity: function plyr_( uint256) constant returns(addr address, name bytes32, win uint256, gen uint256, aff uint256, lrnd uint256, laff uint256)
func (_FomoETClong *FomoETClongCallerSession) Plyr(arg0 *big.Int) (struct {
	Addr common.Address
	Name [32]byte
	Win  *big.Int
	Gen  *big.Int
	Aff  *big.Int
	Lrnd *big.Int
	Laff *big.Int
}, error) {
	return _FomoETClong.Contract.Plyr(&_FomoETClong.CallOpts, arg0)
}

// PotSplit is a free data retrieval call binding the contract method 0xc519500e.
//
// Solidity: function potSplit_( uint256) constant returns(gen uint256, psl uint256)
func (_FomoETClong *FomoETClongCaller) PotSplit(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Gen *big.Int
	Psl *big.Int
}, error) {
	ret := new(struct {
		Gen *big.Int
		Psl *big.Int
	})
	out := ret
	err := _FomoETClong.contract.Call(opts, out, "potSplit_", arg0)
	return *ret, err
}

// PotSplit is a free data retrieval call binding the contract method 0xc519500e.
//
// Solidity: function potSplit_( uint256) constant returns(gen uint256, psl uint256)
func (_FomoETClong *FomoETClongSession) PotSplit(arg0 *big.Int) (struct {
	Gen *big.Int
	Psl *big.Int
}, error) {
	return _FomoETClong.Contract.PotSplit(&_FomoETClong.CallOpts, arg0)
}

// PotSplit is a free data retrieval call binding the contract method 0xc519500e.
//
// Solidity: function potSplit_( uint256) constant returns(gen uint256, psl uint256)
func (_FomoETClong *FomoETClongCallerSession) PotSplit(arg0 *big.Int) (struct {
	Gen *big.Int
	Psl *big.Int
}, error) {
	return _FomoETClong.Contract.PotSplit(&_FomoETClong.CallOpts, arg0)
}

// RID is a free data retrieval call binding the contract method 0x624ae5c0.
//
// Solidity: function rID_() constant returns(uint256)
func (_FomoETClong *FomoETClongCaller) RID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "rID_")
	return *ret0, err
}

// RID is a free data retrieval call binding the contract method 0x624ae5c0.
//
// Solidity: function rID_() constant returns(uint256)
func (_FomoETClong *FomoETClongSession) RID() (*big.Int, error) {
	return _FomoETClong.Contract.RID(&_FomoETClong.CallOpts)
}

// RID is a free data retrieval call binding the contract method 0x624ae5c0.
//
// Solidity: function rID_() constant returns(uint256)
func (_FomoETClong *FomoETClongCallerSession) RID() (*big.Int, error) {
	return _FomoETClong.Contract.RID(&_FomoETClong.CallOpts)
}

// RndTmEth is a free data retrieval call binding the contract method 0x5893d481.
//
// Solidity: function rndTmEth_( uint256,  uint256) constant returns(uint256)
func (_FomoETClong *FomoETClongCaller) RndTmEth(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "rndTmEth_", arg0, arg1)
	return *ret0, err
}

// RndTmEth is a free data retrieval call binding the contract method 0x5893d481.
//
// Solidity: function rndTmEth_( uint256,  uint256) constant returns(uint256)
func (_FomoETClong *FomoETClongSession) RndTmEth(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _FomoETClong.Contract.RndTmEth(&_FomoETClong.CallOpts, arg0, arg1)
}

// RndTmEth is a free data retrieval call binding the contract method 0x5893d481.
//
// Solidity: function rndTmEth_( uint256,  uint256) constant returns(uint256)
func (_FomoETClong *FomoETClongCallerSession) RndTmEth(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _FomoETClong.Contract.RndTmEth(&_FomoETClong.CallOpts, arg0, arg1)
}

// Round is a free data retrieval call binding the contract method 0x24c33d33.
//
// Solidity: function round_( uint256) constant returns(plyr uint256, team uint256, end uint256, ended bool, strt uint256, keys uint256, etc uint256, pot uint256, mask uint256, ico uint256, icoGen uint256, icoAvg uint256)
func (_FomoETClong *FomoETClongCaller) Round(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Plyr   *big.Int
	Team   *big.Int
	End    *big.Int
	Ended  bool
	Strt   *big.Int
	Keys   *big.Int
	Etc    *big.Int
	Pot    *big.Int
	Mask   *big.Int
	Ico    *big.Int
	IcoGen *big.Int
	IcoAvg *big.Int
}, error) {
	ret := new(struct {
		Plyr   *big.Int
		Team   *big.Int
		End    *big.Int
		Ended  bool
		Strt   *big.Int
		Keys   *big.Int
		Etc    *big.Int
		Pot    *big.Int
		Mask   *big.Int
		Ico    *big.Int
		IcoGen *big.Int
		IcoAvg *big.Int
	})
	out := ret
	err := _FomoETClong.contract.Call(opts, out, "round_", arg0)
	return *ret, err
}

// Round is a free data retrieval call binding the contract method 0x24c33d33.
//
// Solidity: function round_( uint256) constant returns(plyr uint256, team uint256, end uint256, ended bool, strt uint256, keys uint256, etc uint256, pot uint256, mask uint256, ico uint256, icoGen uint256, icoAvg uint256)
func (_FomoETClong *FomoETClongSession) Round(arg0 *big.Int) (struct {
	Plyr   *big.Int
	Team   *big.Int
	End    *big.Int
	Ended  bool
	Strt   *big.Int
	Keys   *big.Int
	Etc    *big.Int
	Pot    *big.Int
	Mask   *big.Int
	Ico    *big.Int
	IcoGen *big.Int
	IcoAvg *big.Int
}, error) {
	return _FomoETClong.Contract.Round(&_FomoETClong.CallOpts, arg0)
}

// Round is a free data retrieval call binding the contract method 0x24c33d33.
//
// Solidity: function round_( uint256) constant returns(plyr uint256, team uint256, end uint256, ended bool, strt uint256, keys uint256, etc uint256, pot uint256, mask uint256, ico uint256, icoGen uint256, icoAvg uint256)
func (_FomoETClong *FomoETClongCallerSession) Round(arg0 *big.Int) (struct {
	Plyr   *big.Int
	Team   *big.Int
	End    *big.Int
	Ended  bool
	Strt   *big.Int
	Keys   *big.Int
	Etc    *big.Int
	Pot    *big.Int
	Mask   *big.Int
	Ico    *big.Int
	IcoGen *big.Int
	IcoAvg *big.Int
}, error) {
	return _FomoETClong.Contract.Round(&_FomoETClong.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FomoETClong *FomoETClongCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FomoETClong.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FomoETClong *FomoETClongSession) Symbol() (string, error) {
	return _FomoETClong.Contract.Symbol(&_FomoETClong.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FomoETClong *FomoETClongCallerSession) Symbol() (string, error) {
	return _FomoETClong.Contract.Symbol(&_FomoETClong.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_FomoETClong *FomoETClongTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_FomoETClong *FomoETClongSession) Activate() (*types.Transaction, error) {
	return _FomoETClong.Contract.Activate(&_FomoETClong.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_FomoETClong *FomoETClongTransactorSession) Activate() (*types.Transaction, error) {
	return _FomoETClong.Contract.Activate(&_FomoETClong.TransactOpts)
}

// BuyXaddr is a paid mutator transaction binding the contract method 0x98a0871d.
//
// Solidity: function buyXaddr(_affCode address, _team uint256) returns()
func (_FomoETClong *FomoETClongTransactor) BuyXaddr(opts *bind.TransactOpts, _affCode common.Address, _team *big.Int) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "buyXaddr", _affCode, _team)
}

// BuyXaddr is a paid mutator transaction binding the contract method 0x98a0871d.
//
// Solidity: function buyXaddr(_affCode address, _team uint256) returns()
func (_FomoETClong *FomoETClongSession) BuyXaddr(_affCode common.Address, _team *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.BuyXaddr(&_FomoETClong.TransactOpts, _affCode, _team)
}

// BuyXaddr is a paid mutator transaction binding the contract method 0x98a0871d.
//
// Solidity: function buyXaddr(_affCode address, _team uint256) returns()
func (_FomoETClong *FomoETClongTransactorSession) BuyXaddr(_affCode common.Address, _team *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.BuyXaddr(&_FomoETClong.TransactOpts, _affCode, _team)
}

// BuyXid is a paid mutator transaction binding the contract method 0x8f38f309.
//
// Solidity: function buyXid(_affCode uint256, _team uint256) returns()
func (_FomoETClong *FomoETClongTransactor) BuyXid(opts *bind.TransactOpts, _affCode *big.Int, _team *big.Int) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "buyXid", _affCode, _team)
}

// BuyXid is a paid mutator transaction binding the contract method 0x8f38f309.
//
// Solidity: function buyXid(_affCode uint256, _team uint256) returns()
func (_FomoETClong *FomoETClongSession) BuyXid(_affCode *big.Int, _team *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.BuyXid(&_FomoETClong.TransactOpts, _affCode, _team)
}

// BuyXid is a paid mutator transaction binding the contract method 0x8f38f309.
//
// Solidity: function buyXid(_affCode uint256, _team uint256) returns()
func (_FomoETClong *FomoETClongTransactorSession) BuyXid(_affCode *big.Int, _team *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.BuyXid(&_FomoETClong.TransactOpts, _affCode, _team)
}

// BuyXname is a paid mutator transaction binding the contract method 0xa65b37a1.
//
// Solidity: function buyXname(_affCode bytes32, _team uint256) returns()
func (_FomoETClong *FomoETClongTransactor) BuyXname(opts *bind.TransactOpts, _affCode [32]byte, _team *big.Int) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "buyXname", _affCode, _team)
}

// BuyXname is a paid mutator transaction binding the contract method 0xa65b37a1.
//
// Solidity: function buyXname(_affCode bytes32, _team uint256) returns()
func (_FomoETClong *FomoETClongSession) BuyXname(_affCode [32]byte, _team *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.BuyXname(&_FomoETClong.TransactOpts, _affCode, _team)
}

// BuyXname is a paid mutator transaction binding the contract method 0xa65b37a1.
//
// Solidity: function buyXname(_affCode bytes32, _team uint256) returns()
func (_FomoETClong *FomoETClongTransactorSession) BuyXname(_affCode [32]byte, _team *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.BuyXname(&_FomoETClong.TransactOpts, _affCode, _team)
}

// PotSwap is a paid mutator transaction binding the contract method 0xed78cf4a.
//
// Solidity: function potSwap() returns()
func (_FomoETClong *FomoETClongTransactor) PotSwap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "potSwap")
}

// PotSwap is a paid mutator transaction binding the contract method 0xed78cf4a.
//
// Solidity: function potSwap() returns()
func (_FomoETClong *FomoETClongSession) PotSwap() (*types.Transaction, error) {
	return _FomoETClong.Contract.PotSwap(&_FomoETClong.TransactOpts)
}

// PotSwap is a paid mutator transaction binding the contract method 0xed78cf4a.
//
// Solidity: function potSwap() returns()
func (_FomoETClong *FomoETClongTransactorSession) PotSwap() (*types.Transaction, error) {
	return _FomoETClong.Contract.PotSwap(&_FomoETClong.TransactOpts)
}

// ReLoadXaddr is a paid mutator transaction binding the contract method 0x82bfc739.
//
// Solidity: function reLoadXaddr(_affCode address, _team uint256, _etc uint256) returns()
func (_FomoETClong *FomoETClongTransactor) ReLoadXaddr(opts *bind.TransactOpts, _affCode common.Address, _team *big.Int, _etc *big.Int) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "reLoadXaddr", _affCode, _team, _etc)
}

// ReLoadXaddr is a paid mutator transaction binding the contract method 0x82bfc739.
//
// Solidity: function reLoadXaddr(_affCode address, _team uint256, _etc uint256) returns()
func (_FomoETClong *FomoETClongSession) ReLoadXaddr(_affCode common.Address, _team *big.Int, _etc *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.ReLoadXaddr(&_FomoETClong.TransactOpts, _affCode, _team, _etc)
}

// ReLoadXaddr is a paid mutator transaction binding the contract method 0x82bfc739.
//
// Solidity: function reLoadXaddr(_affCode address, _team uint256, _etc uint256) returns()
func (_FomoETClong *FomoETClongTransactorSession) ReLoadXaddr(_affCode common.Address, _team *big.Int, _etc *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.ReLoadXaddr(&_FomoETClong.TransactOpts, _affCode, _team, _etc)
}

// ReLoadXid is a paid mutator transaction binding the contract method 0x349cdcac.
//
// Solidity: function reLoadXid(_affCode uint256, _team uint256, _etc uint256) returns()
func (_FomoETClong *FomoETClongTransactor) ReLoadXid(opts *bind.TransactOpts, _affCode *big.Int, _team *big.Int, _etc *big.Int) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "reLoadXid", _affCode, _team, _etc)
}

// ReLoadXid is a paid mutator transaction binding the contract method 0x349cdcac.
//
// Solidity: function reLoadXid(_affCode uint256, _team uint256, _etc uint256) returns()
func (_FomoETClong *FomoETClongSession) ReLoadXid(_affCode *big.Int, _team *big.Int, _etc *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.ReLoadXid(&_FomoETClong.TransactOpts, _affCode, _team, _etc)
}

// ReLoadXid is a paid mutator transaction binding the contract method 0x349cdcac.
//
// Solidity: function reLoadXid(_affCode uint256, _team uint256, _etc uint256) returns()
func (_FomoETClong *FomoETClongTransactorSession) ReLoadXid(_affCode *big.Int, _team *big.Int, _etc *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.ReLoadXid(&_FomoETClong.TransactOpts, _affCode, _team, _etc)
}

// ReLoadXname is a paid mutator transaction binding the contract method 0x079ce327.
//
// Solidity: function reLoadXname(_affCode bytes32, _team uint256, _etc uint256) returns()
func (_FomoETClong *FomoETClongTransactor) ReLoadXname(opts *bind.TransactOpts, _affCode [32]byte, _team *big.Int, _etc *big.Int) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "reLoadXname", _affCode, _team, _etc)
}

// ReLoadXname is a paid mutator transaction binding the contract method 0x079ce327.
//
// Solidity: function reLoadXname(_affCode bytes32, _team uint256, _etc uint256) returns()
func (_FomoETClong *FomoETClongSession) ReLoadXname(_affCode [32]byte, _team *big.Int, _etc *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.ReLoadXname(&_FomoETClong.TransactOpts, _affCode, _team, _etc)
}

// ReLoadXname is a paid mutator transaction binding the contract method 0x079ce327.
//
// Solidity: function reLoadXname(_affCode bytes32, _team uint256, _etc uint256) returns()
func (_FomoETClong *FomoETClongTransactorSession) ReLoadXname(_affCode [32]byte, _team *big.Int, _etc *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.ReLoadXname(&_FomoETClong.TransactOpts, _affCode, _team, _etc)
}

// ReceivePlayerInfo is a paid mutator transaction binding the contract method 0x49cc635d.
//
// Solidity: function receivePlayerInfo(_pID uint256, _addr address, _name bytes32, _laff uint256) returns()
func (_FomoETClong *FomoETClongTransactor) ReceivePlayerInfo(opts *bind.TransactOpts, _pID *big.Int, _addr common.Address, _name [32]byte, _laff *big.Int) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "receivePlayerInfo", _pID, _addr, _name, _laff)
}

// ReceivePlayerInfo is a paid mutator transaction binding the contract method 0x49cc635d.
//
// Solidity: function receivePlayerInfo(_pID uint256, _addr address, _name bytes32, _laff uint256) returns()
func (_FomoETClong *FomoETClongSession) ReceivePlayerInfo(_pID *big.Int, _addr common.Address, _name [32]byte, _laff *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.ReceivePlayerInfo(&_FomoETClong.TransactOpts, _pID, _addr, _name, _laff)
}

// ReceivePlayerInfo is a paid mutator transaction binding the contract method 0x49cc635d.
//
// Solidity: function receivePlayerInfo(_pID uint256, _addr address, _name bytes32, _laff uint256) returns()
func (_FomoETClong *FomoETClongTransactorSession) ReceivePlayerInfo(_pID *big.Int, _addr common.Address, _name [32]byte, _laff *big.Int) (*types.Transaction, error) {
	return _FomoETClong.Contract.ReceivePlayerInfo(&_FomoETClong.TransactOpts, _pID, _addr, _name, _laff)
}

// ReceivePlayerNameList is a paid mutator transaction binding the contract method 0x8f7140ea.
//
// Solidity: function receivePlayerNameList(_pID uint256, _name bytes32) returns()
func (_FomoETClong *FomoETClongTransactor) ReceivePlayerNameList(opts *bind.TransactOpts, _pID *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "receivePlayerNameList", _pID, _name)
}

// ReceivePlayerNameList is a paid mutator transaction binding the contract method 0x8f7140ea.
//
// Solidity: function receivePlayerNameList(_pID uint256, _name bytes32) returns()
func (_FomoETClong *FomoETClongSession) ReceivePlayerNameList(_pID *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _FomoETClong.Contract.ReceivePlayerNameList(&_FomoETClong.TransactOpts, _pID, _name)
}

// ReceivePlayerNameList is a paid mutator transaction binding the contract method 0x8f7140ea.
//
// Solidity: function receivePlayerNameList(_pID uint256, _name bytes32) returns()
func (_FomoETClong *FomoETClongTransactorSession) ReceivePlayerNameList(_pID *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _FomoETClong.Contract.ReceivePlayerNameList(&_FomoETClong.TransactOpts, _pID, _name)
}

// RegisterNameXID is a paid mutator transaction binding the contract method 0x921dec21.
//
// Solidity: function registerNameXID(_nameString string, _affCode uint256, _all bool) returns()
func (_FomoETClong *FomoETClongTransactor) RegisterNameXID(opts *bind.TransactOpts, _nameString string, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "registerNameXID", _nameString, _affCode, _all)
}

// RegisterNameXID is a paid mutator transaction binding the contract method 0x921dec21.
//
// Solidity: function registerNameXID(_nameString string, _affCode uint256, _all bool) returns()
func (_FomoETClong *FomoETClongSession) RegisterNameXID(_nameString string, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _FomoETClong.Contract.RegisterNameXID(&_FomoETClong.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXID is a paid mutator transaction binding the contract method 0x921dec21.
//
// Solidity: function registerNameXID(_nameString string, _affCode uint256, _all bool) returns()
func (_FomoETClong *FomoETClongTransactorSession) RegisterNameXID(_nameString string, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _FomoETClong.Contract.RegisterNameXID(&_FomoETClong.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXaddr is a paid mutator transaction binding the contract method 0x3ddd4698.
//
// Solidity: function registerNameXaddr(_nameString string, _affCode address, _all bool) returns()
func (_FomoETClong *FomoETClongTransactor) RegisterNameXaddr(opts *bind.TransactOpts, _nameString string, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "registerNameXaddr", _nameString, _affCode, _all)
}

// RegisterNameXaddr is a paid mutator transaction binding the contract method 0x3ddd4698.
//
// Solidity: function registerNameXaddr(_nameString string, _affCode address, _all bool) returns()
func (_FomoETClong *FomoETClongSession) RegisterNameXaddr(_nameString string, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _FomoETClong.Contract.RegisterNameXaddr(&_FomoETClong.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXaddr is a paid mutator transaction binding the contract method 0x3ddd4698.
//
// Solidity: function registerNameXaddr(_nameString string, _affCode address, _all bool) returns()
func (_FomoETClong *FomoETClongTransactorSession) RegisterNameXaddr(_nameString string, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _FomoETClong.Contract.RegisterNameXaddr(&_FomoETClong.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXname is a paid mutator transaction binding the contract method 0x685ffd83.
//
// Solidity: function registerNameXname(_nameString string, _affCode bytes32, _all bool) returns()
func (_FomoETClong *FomoETClongTransactor) RegisterNameXname(opts *bind.TransactOpts, _nameString string, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "registerNameXname", _nameString, _affCode, _all)
}

// RegisterNameXname is a paid mutator transaction binding the contract method 0x685ffd83.
//
// Solidity: function registerNameXname(_nameString string, _affCode bytes32, _all bool) returns()
func (_FomoETClong *FomoETClongSession) RegisterNameXname(_nameString string, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _FomoETClong.Contract.RegisterNameXname(&_FomoETClong.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXname is a paid mutator transaction binding the contract method 0x685ffd83.
//
// Solidity: function registerNameXname(_nameString string, _affCode bytes32, _all bool) returns()
func (_FomoETClong *FomoETClongTransactorSession) RegisterNameXname(_nameString string, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _FomoETClong.Contract.RegisterNameXname(&_FomoETClong.TransactOpts, _nameString, _affCode, _all)
}

// SetOtherFomo is a paid mutator transaction binding the contract method 0xb483c054.
//
// Solidity: function setOtherFomo(_otherFETC address) returns()
func (_FomoETClong *FomoETClongTransactor) SetOtherFomo(opts *bind.TransactOpts, _otherFETC common.Address) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "setOtherFomo", _otherFETC)
}

// SetOtherFomo is a paid mutator transaction binding the contract method 0xb483c054.
//
// Solidity: function setOtherFomo(_otherFETC address) returns()
func (_FomoETClong *FomoETClongSession) SetOtherFomo(_otherFETC common.Address) (*types.Transaction, error) {
	return _FomoETClong.Contract.SetOtherFomo(&_FomoETClong.TransactOpts, _otherFETC)
}

// SetOtherFomo is a paid mutator transaction binding the contract method 0xb483c054.
//
// Solidity: function setOtherFomo(_otherFETC address) returns()
func (_FomoETClong *FomoETClongTransactorSession) SetOtherFomo(_otherFETC common.Address) (*types.Transaction, error) {
	return _FomoETClong.Contract.SetOtherFomo(&_FomoETClong.TransactOpts, _otherFETC)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FomoETClong *FomoETClongTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FomoETClong.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FomoETClong *FomoETClongSession) Withdraw() (*types.Transaction, error) {
	return _FomoETClong.Contract.Withdraw(&_FomoETClong.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FomoETClong *FomoETClongTransactorSession) Withdraw() (*types.Transaction, error) {
	return _FomoETClong.Contract.Withdraw(&_FomoETClong.TransactOpts)
}

// FomoETClongOnAffiliatePayoutIterator is returned from FilterOnAffiliatePayout and is used to iterate over the raw logs and unpacked data for OnAffiliatePayout events raised by the FomoETClong contract.
type FomoETClongOnAffiliatePayoutIterator struct {
	Event *FomoETClongOnAffiliatePayout // Event containing the contract specifics and raw log

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
func (it *FomoETClongOnAffiliatePayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FomoETClongOnAffiliatePayout)
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
		it.Event = new(FomoETClongOnAffiliatePayout)
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
func (it *FomoETClongOnAffiliatePayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FomoETClongOnAffiliatePayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FomoETClongOnAffiliatePayout represents a OnAffiliatePayout event raised by the FomoETClong contract.
type FomoETClongOnAffiliatePayout struct {
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	RoundID          *big.Int
	BuyerID          *big.Int
	Amount           *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnAffiliatePayout is a free log retrieval operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: event onAffiliatePayout(affiliateID indexed uint256, affiliateAddress address, affiliateName bytes32, roundID indexed uint256, buyerID indexed uint256, amount uint256, timeStamp uint256)
func (_FomoETClong *FomoETClongFilterer) FilterOnAffiliatePayout(opts *bind.FilterOpts, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (*FomoETClongOnAffiliatePayoutIterator, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _FomoETClong.contract.FilterLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return &FomoETClongOnAffiliatePayoutIterator{contract: _FomoETClong.contract, event: "onAffiliatePayout", logs: logs, sub: sub}, nil
}

// WatchOnAffiliatePayout is a free log subscription operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: event onAffiliatePayout(affiliateID indexed uint256, affiliateAddress address, affiliateName bytes32, roundID indexed uint256, buyerID indexed uint256, amount uint256, timeStamp uint256)
func (_FomoETClong *FomoETClongFilterer) WatchOnAffiliatePayout(opts *bind.WatchOpts, sink chan<- *FomoETClongOnAffiliatePayout, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (event.Subscription, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _FomoETClong.contract.WatchLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FomoETClongOnAffiliatePayout)
				if err := _FomoETClong.contract.UnpackLog(event, "onAffiliatePayout", log); err != nil {
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

// FomoETClongOnBuyAndDistributeIterator is returned from FilterOnBuyAndDistribute and is used to iterate over the raw logs and unpacked data for OnBuyAndDistribute events raised by the FomoETClong contract.
type FomoETClongOnBuyAndDistributeIterator struct {
	Event *FomoETClongOnBuyAndDistribute // Event containing the contract specifics and raw log

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
func (it *FomoETClongOnBuyAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FomoETClongOnBuyAndDistribute)
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
		it.Event = new(FomoETClongOnBuyAndDistribute)
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
func (it *FomoETClongOnBuyAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FomoETClongOnBuyAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FomoETClongOnBuyAndDistribute represents a OnBuyAndDistribute event raised by the FomoETClong contract.
type FomoETClongOnBuyAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EtcIn          *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	PSLAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnBuyAndDistribute is a free log retrieval operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: event onBuyAndDistribute(playerAddress address, playerName bytes32, etcIn uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_FomoETClong *FomoETClongFilterer) FilterOnBuyAndDistribute(opts *bind.FilterOpts) (*FomoETClongOnBuyAndDistributeIterator, error) {

	logs, sub, err := _FomoETClong.contract.FilterLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return &FomoETClongOnBuyAndDistributeIterator{contract: _FomoETClong.contract, event: "onBuyAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnBuyAndDistribute is a free log subscription operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: event onBuyAndDistribute(playerAddress address, playerName bytes32, etcIn uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_FomoETClong *FomoETClongFilterer) WatchOnBuyAndDistribute(opts *bind.WatchOpts, sink chan<- *FomoETClongOnBuyAndDistribute) (event.Subscription, error) {

	logs, sub, err := _FomoETClong.contract.WatchLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FomoETClongOnBuyAndDistribute)
				if err := _FomoETClong.contract.UnpackLog(event, "onBuyAndDistribute", log); err != nil {
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

// FomoETClongOnEndTxIterator is returned from FilterOnEndTx and is used to iterate over the raw logs and unpacked data for OnEndTx events raised by the FomoETClong contract.
type FomoETClongOnEndTxIterator struct {
	Event *FomoETClongOnEndTx // Event containing the contract specifics and raw log

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
func (it *FomoETClongOnEndTxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FomoETClongOnEndTx)
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
		it.Event = new(FomoETClongOnEndTx)
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
func (it *FomoETClongOnEndTxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FomoETClongOnEndTxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FomoETClongOnEndTx represents a OnEndTx event raised by the FomoETClong contract.
type FomoETClongOnEndTx struct {
	CompressedData *big.Int
	CompressedIDs  *big.Int
	PlayerName     [32]byte
	PlayerAddress  common.Address
	EtcIn          *big.Int
	KeysBought     *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	PSLAmount      *big.Int
	GenAmount      *big.Int
	PotAmount      *big.Int
	AirDropPot     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnEndTx is a free log retrieval operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: event onEndTx(compressedData uint256, compressedIDs uint256, playerName bytes32, playerAddress address, etcIn uint256, keysBought uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256, potAmount uint256, airDropPot uint256)
func (_FomoETClong *FomoETClongFilterer) FilterOnEndTx(opts *bind.FilterOpts) (*FomoETClongOnEndTxIterator, error) {

	logs, sub, err := _FomoETClong.contract.FilterLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return &FomoETClongOnEndTxIterator{contract: _FomoETClong.contract, event: "onEndTx", logs: logs, sub: sub}, nil
}

// WatchOnEndTx is a free log subscription operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: event onEndTx(compressedData uint256, compressedIDs uint256, playerName bytes32, playerAddress address, etcIn uint256, keysBought uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256, potAmount uint256, airDropPot uint256)
func (_FomoETClong *FomoETClongFilterer) WatchOnEndTx(opts *bind.WatchOpts, sink chan<- *FomoETClongOnEndTx) (event.Subscription, error) {

	logs, sub, err := _FomoETClong.contract.WatchLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FomoETClongOnEndTx)
				if err := _FomoETClong.contract.UnpackLog(event, "onEndTx", log); err != nil {
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

// FomoETClongOnNewNameIterator is returned from FilterOnNewName and is used to iterate over the raw logs and unpacked data for OnNewName events raised by the FomoETClong contract.
type FomoETClongOnNewNameIterator struct {
	Event *FomoETClongOnNewName // Event containing the contract specifics and raw log

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
func (it *FomoETClongOnNewNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FomoETClongOnNewName)
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
		it.Event = new(FomoETClongOnNewName)
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
func (it *FomoETClongOnNewNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FomoETClongOnNewNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FomoETClongOnNewName represents a OnNewName event raised by the FomoETClong contract.
type FomoETClongOnNewName struct {
	PlayerID         *big.Int
	PlayerAddress    common.Address
	PlayerName       [32]byte
	IsNewPlayer      bool
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	AmountPaid       *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnNewName is a free log retrieval operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: event onNewName(playerID indexed uint256, playerAddress indexed address, playerName indexed bytes32, isNewPlayer bool, affiliateID uint256, affiliateAddress address, affiliateName bytes32, amountPaid uint256, timeStamp uint256)
func (_FomoETClong *FomoETClongFilterer) FilterOnNewName(opts *bind.FilterOpts, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (*FomoETClongOnNewNameIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _FomoETClong.contract.FilterLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return &FomoETClongOnNewNameIterator{contract: _FomoETClong.contract, event: "onNewName", logs: logs, sub: sub}, nil
}

// WatchOnNewName is a free log subscription operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: event onNewName(playerID indexed uint256, playerAddress indexed address, playerName indexed bytes32, isNewPlayer bool, affiliateID uint256, affiliateAddress address, affiliateName bytes32, amountPaid uint256, timeStamp uint256)
func (_FomoETClong *FomoETClongFilterer) WatchOnNewName(opts *bind.WatchOpts, sink chan<- *FomoETClongOnNewName, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _FomoETClong.contract.WatchLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FomoETClongOnNewName)
				if err := _FomoETClong.contract.UnpackLog(event, "onNewName", log); err != nil {
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

// FomoETClongOnPotSwapDepositIterator is returned from FilterOnPotSwapDeposit and is used to iterate over the raw logs and unpacked data for OnPotSwapDeposit events raised by the FomoETClong contract.
type FomoETClongOnPotSwapDepositIterator struct {
	Event *FomoETClongOnPotSwapDeposit // Event containing the contract specifics and raw log

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
func (it *FomoETClongOnPotSwapDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FomoETClongOnPotSwapDeposit)
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
		it.Event = new(FomoETClongOnPotSwapDeposit)
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
func (it *FomoETClongOnPotSwapDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FomoETClongOnPotSwapDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FomoETClongOnPotSwapDeposit represents a OnPotSwapDeposit event raised by the FomoETClong contract.
type FomoETClongOnPotSwapDeposit struct {
	RoundID          *big.Int
	AmountAddedToPot *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnPotSwapDeposit is a free log retrieval operation binding the contract event 0x74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c.
//
// Solidity: event onPotSwapDeposit(roundID uint256, amountAddedToPot uint256)
func (_FomoETClong *FomoETClongFilterer) FilterOnPotSwapDeposit(opts *bind.FilterOpts) (*FomoETClongOnPotSwapDepositIterator, error) {

	logs, sub, err := _FomoETClong.contract.FilterLogs(opts, "onPotSwapDeposit")
	if err != nil {
		return nil, err
	}
	return &FomoETClongOnPotSwapDepositIterator{contract: _FomoETClong.contract, event: "onPotSwapDeposit", logs: logs, sub: sub}, nil
}

// WatchOnPotSwapDeposit is a free log subscription operation binding the contract event 0x74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c.
//
// Solidity: event onPotSwapDeposit(roundID uint256, amountAddedToPot uint256)
func (_FomoETClong *FomoETClongFilterer) WatchOnPotSwapDeposit(opts *bind.WatchOpts, sink chan<- *FomoETClongOnPotSwapDeposit) (event.Subscription, error) {

	logs, sub, err := _FomoETClong.contract.WatchLogs(opts, "onPotSwapDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FomoETClongOnPotSwapDeposit)
				if err := _FomoETClong.contract.UnpackLog(event, "onPotSwapDeposit", log); err != nil {
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

// FomoETClongOnReLoadAndDistributeIterator is returned from FilterOnReLoadAndDistribute and is used to iterate over the raw logs and unpacked data for OnReLoadAndDistribute events raised by the FomoETClong contract.
type FomoETClongOnReLoadAndDistributeIterator struct {
	Event *FomoETClongOnReLoadAndDistribute // Event containing the contract specifics and raw log

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
func (it *FomoETClongOnReLoadAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FomoETClongOnReLoadAndDistribute)
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
		it.Event = new(FomoETClongOnReLoadAndDistribute)
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
func (it *FomoETClongOnReLoadAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FomoETClongOnReLoadAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FomoETClongOnReLoadAndDistribute represents a OnReLoadAndDistribute event raised by the FomoETClong contract.
type FomoETClongOnReLoadAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	PSLAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnReLoadAndDistribute is a free log retrieval operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: event onReLoadAndDistribute(playerAddress address, playerName bytes32, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_FomoETClong *FomoETClongFilterer) FilterOnReLoadAndDistribute(opts *bind.FilterOpts) (*FomoETClongOnReLoadAndDistributeIterator, error) {

	logs, sub, err := _FomoETClong.contract.FilterLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return &FomoETClongOnReLoadAndDistributeIterator{contract: _FomoETClong.contract, event: "onReLoadAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnReLoadAndDistribute is a free log subscription operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: event onReLoadAndDistribute(playerAddress address, playerName bytes32, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_FomoETClong *FomoETClongFilterer) WatchOnReLoadAndDistribute(opts *bind.WatchOpts, sink chan<- *FomoETClongOnReLoadAndDistribute) (event.Subscription, error) {

	logs, sub, err := _FomoETClong.contract.WatchLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FomoETClongOnReLoadAndDistribute)
				if err := _FomoETClong.contract.UnpackLog(event, "onReLoadAndDistribute", log); err != nil {
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

// FomoETClongOnWithdrawIterator is returned from FilterOnWithdraw and is used to iterate over the raw logs and unpacked data for OnWithdraw events raised by the FomoETClong contract.
type FomoETClongOnWithdrawIterator struct {
	Event *FomoETClongOnWithdraw // Event containing the contract specifics and raw log

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
func (it *FomoETClongOnWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FomoETClongOnWithdraw)
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
		it.Event = new(FomoETClongOnWithdraw)
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
func (it *FomoETClongOnWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FomoETClongOnWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FomoETClongOnWithdraw represents a OnWithdraw event raised by the FomoETClong contract.
type FomoETClongOnWithdraw struct {
	PlayerID      *big.Int
	PlayerAddress common.Address
	PlayerName    [32]byte
	EtcOut        *big.Int
	TimeStamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOnWithdraw is a free log retrieval operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: event onWithdraw(playerID indexed uint256, playerAddress address, playerName bytes32, etcOut uint256, timeStamp uint256)
func (_FomoETClong *FomoETClongFilterer) FilterOnWithdraw(opts *bind.FilterOpts, playerID []*big.Int) (*FomoETClongOnWithdrawIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _FomoETClong.contract.FilterLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return &FomoETClongOnWithdrawIterator{contract: _FomoETClong.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: event onWithdraw(playerID indexed uint256, playerAddress address, playerName bytes32, etcOut uint256, timeStamp uint256)
func (_FomoETClong *FomoETClongFilterer) WatchOnWithdraw(opts *bind.WatchOpts, sink chan<- *FomoETClongOnWithdraw, playerID []*big.Int) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _FomoETClong.contract.WatchLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FomoETClongOnWithdraw)
				if err := _FomoETClong.contract.UnpackLog(event, "onWithdraw", log); err != nil {
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

// FomoETClongOnWithdrawAndDistributeIterator is returned from FilterOnWithdrawAndDistribute and is used to iterate over the raw logs and unpacked data for OnWithdrawAndDistribute events raised by the FomoETClong contract.
type FomoETClongOnWithdrawAndDistributeIterator struct {
	Event *FomoETClongOnWithdrawAndDistribute // Event containing the contract specifics and raw log

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
func (it *FomoETClongOnWithdrawAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FomoETClongOnWithdrawAndDistribute)
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
		it.Event = new(FomoETClongOnWithdrawAndDistribute)
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
func (it *FomoETClongOnWithdrawAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FomoETClongOnWithdrawAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FomoETClongOnWithdrawAndDistribute represents a OnWithdrawAndDistribute event raised by the FomoETClong contract.
type FomoETClongOnWithdrawAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EtcOut         *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	PSLAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnWithdrawAndDistribute is a free log retrieval operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: event onWithdrawAndDistribute(playerAddress address, playerName bytes32, etcOut uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_FomoETClong *FomoETClongFilterer) FilterOnWithdrawAndDistribute(opts *bind.FilterOpts) (*FomoETClongOnWithdrawAndDistributeIterator, error) {

	logs, sub, err := _FomoETClong.contract.FilterLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return &FomoETClongOnWithdrawAndDistributeIterator{contract: _FomoETClong.contract, event: "onWithdrawAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnWithdrawAndDistribute is a free log subscription operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: event onWithdrawAndDistribute(playerAddress address, playerName bytes32, etcOut uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_FomoETClong *FomoETClongFilterer) WatchOnWithdrawAndDistribute(opts *bind.WatchOpts, sink chan<- *FomoETClongOnWithdrawAndDistribute) (event.Subscription, error) {

	logs, sub, err := _FomoETClong.contract.WatchLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FomoETClongOnWithdrawAndDistribute)
				if err := _FomoETClong.contract.UnpackLog(event, "onWithdrawAndDistribute", log); err != nil {
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

// MagicPowerForwarderInterfaceABI is the input ABI used to generate the binding from.
const MagicPowerForwarderInterfaceABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"cancelMigration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_firstCorpBank\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMigration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCorpBank\",\"type\":\"address\"}],\"name\":\"startMigration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// MagicPowerForwarderInterfaceBin is the compiled bytecode used for deploying new contracts.
const MagicPowerForwarderInterfaceBin = `0x`

// DeployMagicPowerForwarderInterface deploys a new Ethereum contract, binding an instance of MagicPowerForwarderInterface to it.
func DeployMagicPowerForwarderInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MagicPowerForwarderInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(MagicPowerForwarderInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MagicPowerForwarderInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MagicPowerForwarderInterface{MagicPowerForwarderInterfaceCaller: MagicPowerForwarderInterfaceCaller{contract: contract}, MagicPowerForwarderInterfaceTransactor: MagicPowerForwarderInterfaceTransactor{contract: contract}, MagicPowerForwarderInterfaceFilterer: MagicPowerForwarderInterfaceFilterer{contract: contract}}, nil
}

// MagicPowerForwarderInterface is an auto generated Go binding around an Ethereum contract.
type MagicPowerForwarderInterface struct {
	MagicPowerForwarderInterfaceCaller     // Read-only binding to the contract
	MagicPowerForwarderInterfaceTransactor // Write-only binding to the contract
	MagicPowerForwarderInterfaceFilterer   // Log filterer for contract events
}

// MagicPowerForwarderInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MagicPowerForwarderInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagicPowerForwarderInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MagicPowerForwarderInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagicPowerForwarderInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MagicPowerForwarderInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagicPowerForwarderInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MagicPowerForwarderInterfaceSession struct {
	Contract     *MagicPowerForwarderInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MagicPowerForwarderInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MagicPowerForwarderInterfaceCallerSession struct {
	Contract *MagicPowerForwarderInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// MagicPowerForwarderInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MagicPowerForwarderInterfaceTransactorSession struct {
	Contract     *MagicPowerForwarderInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// MagicPowerForwarderInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MagicPowerForwarderInterfaceRaw struct {
	Contract *MagicPowerForwarderInterface // Generic contract binding to access the raw methods on
}

// MagicPowerForwarderInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MagicPowerForwarderInterfaceCallerRaw struct {
	Contract *MagicPowerForwarderInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// MagicPowerForwarderInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MagicPowerForwarderInterfaceTransactorRaw struct {
	Contract *MagicPowerForwarderInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMagicPowerForwarderInterface creates a new instance of MagicPowerForwarderInterface, bound to a specific deployed contract.
func NewMagicPowerForwarderInterface(address common.Address, backend bind.ContractBackend) (*MagicPowerForwarderInterface, error) {
	contract, err := bindMagicPowerForwarderInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MagicPowerForwarderInterface{MagicPowerForwarderInterfaceCaller: MagicPowerForwarderInterfaceCaller{contract: contract}, MagicPowerForwarderInterfaceTransactor: MagicPowerForwarderInterfaceTransactor{contract: contract}, MagicPowerForwarderInterfaceFilterer: MagicPowerForwarderInterfaceFilterer{contract: contract}}, nil
}

// NewMagicPowerForwarderInterfaceCaller creates a new read-only instance of MagicPowerForwarderInterface, bound to a specific deployed contract.
func NewMagicPowerForwarderInterfaceCaller(address common.Address, caller bind.ContractCaller) (*MagicPowerForwarderInterfaceCaller, error) {
	contract, err := bindMagicPowerForwarderInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MagicPowerForwarderInterfaceCaller{contract: contract}, nil
}

// NewMagicPowerForwarderInterfaceTransactor creates a new write-only instance of MagicPowerForwarderInterface, bound to a specific deployed contract.
func NewMagicPowerForwarderInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MagicPowerForwarderInterfaceTransactor, error) {
	contract, err := bindMagicPowerForwarderInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MagicPowerForwarderInterfaceTransactor{contract: contract}, nil
}

// NewMagicPowerForwarderInterfaceFilterer creates a new log filterer instance of MagicPowerForwarderInterface, bound to a specific deployed contract.
func NewMagicPowerForwarderInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MagicPowerForwarderInterfaceFilterer, error) {
	contract, err := bindMagicPowerForwarderInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MagicPowerForwarderInterfaceFilterer{contract: contract}, nil
}

// bindMagicPowerForwarderInterface binds a generic wrapper to an already deployed contract.
func bindMagicPowerForwarderInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MagicPowerForwarderInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MagicPowerForwarderInterface.Contract.MagicPowerForwarderInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.MagicPowerForwarderInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.MagicPowerForwarderInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MagicPowerForwarderInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.contract.Transact(opts, method, params...)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(address, address, bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceCaller) Status(opts *bind.CallOpts) (common.Address, common.Address, bool, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _MagicPowerForwarderInterface.contract.Call(opts, out, "status")
	return *ret0, *ret1, *ret2, err
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(address, address, bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceSession) Status() (common.Address, common.Address, bool, error) {
	return _MagicPowerForwarderInterface.Contract.Status(&_MagicPowerForwarderInterface.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(address, address, bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceCallerSession) Status() (common.Address, common.Address, bool, error) {
	return _MagicPowerForwarderInterface.Contract.Status(&_MagicPowerForwarderInterface.CallOpts)
}

// CancelMigration is a paid mutator transaction binding the contract method 0x10639ea0.
//
// Solidity: function cancelMigration() returns(bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceTransactor) CancelMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.contract.Transact(opts, "cancelMigration")
}

// CancelMigration is a paid mutator transaction binding the contract method 0x10639ea0.
//
// Solidity: function cancelMigration() returns(bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceSession) CancelMigration() (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.CancelMigration(&_MagicPowerForwarderInterface.TransactOpts)
}

// CancelMigration is a paid mutator transaction binding the contract method 0x10639ea0.
//
// Solidity: function cancelMigration() returns(bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceTransactorSession) CancelMigration() (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.CancelMigration(&_MagicPowerForwarderInterface.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceSession) Deposit() (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.Deposit(&_MagicPowerForwarderInterface.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceTransactorSession) Deposit() (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.Deposit(&_MagicPowerForwarderInterface.TransactOpts)
}

// FinishMigration is a paid mutator transaction binding the contract method 0x88d761f2.
//
// Solidity: function finishMigration() returns(bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceTransactor) FinishMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.contract.Transact(opts, "finishMigration")
}

// FinishMigration is a paid mutator transaction binding the contract method 0x88d761f2.
//
// Solidity: function finishMigration() returns(bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceSession) FinishMigration() (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.FinishMigration(&_MagicPowerForwarderInterface.TransactOpts)
}

// FinishMigration is a paid mutator transaction binding the contract method 0x88d761f2.
//
// Solidity: function finishMigration() returns(bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceTransactorSession) FinishMigration() (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.FinishMigration(&_MagicPowerForwarderInterface.TransactOpts)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(_firstCorpBank address) returns()
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceTransactor) Setup(opts *bind.TransactOpts, _firstCorpBank common.Address) (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.contract.Transact(opts, "setup", _firstCorpBank)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(_firstCorpBank address) returns()
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceSession) Setup(_firstCorpBank common.Address) (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.Setup(&_MagicPowerForwarderInterface.TransactOpts, _firstCorpBank)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(_firstCorpBank address) returns()
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceTransactorSession) Setup(_firstCorpBank common.Address) (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.Setup(&_MagicPowerForwarderInterface.TransactOpts, _firstCorpBank)
}

// StartMigration is a paid mutator transaction binding the contract method 0xa0f52da0.
//
// Solidity: function startMigration(_newCorpBank address) returns(bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceTransactor) StartMigration(opts *bind.TransactOpts, _newCorpBank common.Address) (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.contract.Transact(opts, "startMigration", _newCorpBank)
}

// StartMigration is a paid mutator transaction binding the contract method 0xa0f52da0.
//
// Solidity: function startMigration(_newCorpBank address) returns(bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceSession) StartMigration(_newCorpBank common.Address) (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.StartMigration(&_MagicPowerForwarderInterface.TransactOpts, _newCorpBank)
}

// StartMigration is a paid mutator transaction binding the contract method 0xa0f52da0.
//
// Solidity: function startMigration(_newCorpBank address) returns(bool)
func (_MagicPowerForwarderInterface *MagicPowerForwarderInterfaceTransactorSession) StartMigration(_newCorpBank common.Address) (*types.Transaction, error) {
	return _MagicPowerForwarderInterface.Contract.StartMigration(&_MagicPowerForwarderInterface.TransactOpts, _newCorpBank)
}

// NameFilterABI is the input ABI used to generate the binding from.
const NameFilterABI = "[]"

// NameFilterBin is the compiled bytecode used for deploying new contracts.
const NameFilterBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058203efbcfe6793d535828ef05b4ad0bb651cb265ad29c00ca355da7b028201e06c50029`

// DeployNameFilter deploys a new Ethereum contract, binding an instance of NameFilter to it.
func DeployNameFilter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NameFilter, error) {
	parsed, err := abi.JSON(strings.NewReader(NameFilterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NameFilterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NameFilter{NameFilterCaller: NameFilterCaller{contract: contract}, NameFilterTransactor: NameFilterTransactor{contract: contract}, NameFilterFilterer: NameFilterFilterer{contract: contract}}, nil
}

// NameFilter is an auto generated Go binding around an Ethereum contract.
type NameFilter struct {
	NameFilterCaller     // Read-only binding to the contract
	NameFilterTransactor // Write-only binding to the contract
	NameFilterFilterer   // Log filterer for contract events
}

// NameFilterCaller is an auto generated read-only Go binding around an Ethereum contract.
type NameFilterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameFilterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NameFilterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameFilterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NameFilterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameFilterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NameFilterSession struct {
	Contract     *NameFilter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NameFilterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NameFilterCallerSession struct {
	Contract *NameFilterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NameFilterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NameFilterTransactorSession struct {
	Contract     *NameFilterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NameFilterRaw is an auto generated low-level Go binding around an Ethereum contract.
type NameFilterRaw struct {
	Contract *NameFilter // Generic contract binding to access the raw methods on
}

// NameFilterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NameFilterCallerRaw struct {
	Contract *NameFilterCaller // Generic read-only contract binding to access the raw methods on
}

// NameFilterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NameFilterTransactorRaw struct {
	Contract *NameFilterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNameFilter creates a new instance of NameFilter, bound to a specific deployed contract.
func NewNameFilter(address common.Address, backend bind.ContractBackend) (*NameFilter, error) {
	contract, err := bindNameFilter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NameFilter{NameFilterCaller: NameFilterCaller{contract: contract}, NameFilterTransactor: NameFilterTransactor{contract: contract}, NameFilterFilterer: NameFilterFilterer{contract: contract}}, nil
}

// NewNameFilterCaller creates a new read-only instance of NameFilter, bound to a specific deployed contract.
func NewNameFilterCaller(address common.Address, caller bind.ContractCaller) (*NameFilterCaller, error) {
	contract, err := bindNameFilter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NameFilterCaller{contract: contract}, nil
}

// NewNameFilterTransactor creates a new write-only instance of NameFilter, bound to a specific deployed contract.
func NewNameFilterTransactor(address common.Address, transactor bind.ContractTransactor) (*NameFilterTransactor, error) {
	contract, err := bindNameFilter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NameFilterTransactor{contract: contract}, nil
}

// NewNameFilterFilterer creates a new log filterer instance of NameFilter, bound to a specific deployed contract.
func NewNameFilterFilterer(address common.Address, filterer bind.ContractFilterer) (*NameFilterFilterer, error) {
	contract, err := bindNameFilter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NameFilterFilterer{contract: contract}, nil
}

// bindNameFilter binds a generic wrapper to an already deployed contract.
func bindNameFilter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NameFilterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameFilter *NameFilterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NameFilter.Contract.NameFilterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameFilter *NameFilterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameFilter.Contract.NameFilterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameFilter *NameFilterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameFilter.Contract.NameFilterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameFilter *NameFilterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NameFilter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameFilter *NameFilterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameFilter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameFilter *NameFilterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameFilter.Contract.contract.Transact(opts, method, params...)
}

// PlayerBookInterfaceABI is the input ABI used to generate the binding from.
const PlayerBookInterfaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getNameFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"}],\"name\":\"getPlayerAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_affCode\",\"type\":\"bytes32\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXnameFromDapp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"}],\"name\":\"getPlayerName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_affCode\",\"type\":\"address\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXaddrFromDapp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_affCode\",\"type\":\"uint256\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXIDFromDapp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"}],\"name\":\"getPlayerLAff\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getPlayerID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PlayerBookInterfaceBin is the compiled bytecode used for deploying new contracts.
const PlayerBookInterfaceBin = `0x`

// DeployPlayerBookInterface deploys a new Ethereum contract, binding an instance of PlayerBookInterface to it.
func DeployPlayerBookInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PlayerBookInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(PlayerBookInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PlayerBookInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PlayerBookInterface{PlayerBookInterfaceCaller: PlayerBookInterfaceCaller{contract: contract}, PlayerBookInterfaceTransactor: PlayerBookInterfaceTransactor{contract: contract}, PlayerBookInterfaceFilterer: PlayerBookInterfaceFilterer{contract: contract}}, nil
}

// PlayerBookInterface is an auto generated Go binding around an Ethereum contract.
type PlayerBookInterface struct {
	PlayerBookInterfaceCaller     // Read-only binding to the contract
	PlayerBookInterfaceTransactor // Write-only binding to the contract
	PlayerBookInterfaceFilterer   // Log filterer for contract events
}

// PlayerBookInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlayerBookInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayerBookInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlayerBookInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayerBookInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlayerBookInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayerBookInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlayerBookInterfaceSession struct {
	Contract     *PlayerBookInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PlayerBookInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlayerBookInterfaceCallerSession struct {
	Contract *PlayerBookInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PlayerBookInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlayerBookInterfaceTransactorSession struct {
	Contract     *PlayerBookInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PlayerBookInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlayerBookInterfaceRaw struct {
	Contract *PlayerBookInterface // Generic contract binding to access the raw methods on
}

// PlayerBookInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlayerBookInterfaceCallerRaw struct {
	Contract *PlayerBookInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// PlayerBookInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlayerBookInterfaceTransactorRaw struct {
	Contract *PlayerBookInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlayerBookInterface creates a new instance of PlayerBookInterface, bound to a specific deployed contract.
func NewPlayerBookInterface(address common.Address, backend bind.ContractBackend) (*PlayerBookInterface, error) {
	contract, err := bindPlayerBookInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlayerBookInterface{PlayerBookInterfaceCaller: PlayerBookInterfaceCaller{contract: contract}, PlayerBookInterfaceTransactor: PlayerBookInterfaceTransactor{contract: contract}, PlayerBookInterfaceFilterer: PlayerBookInterfaceFilterer{contract: contract}}, nil
}

// NewPlayerBookInterfaceCaller creates a new read-only instance of PlayerBookInterface, bound to a specific deployed contract.
func NewPlayerBookInterfaceCaller(address common.Address, caller bind.ContractCaller) (*PlayerBookInterfaceCaller, error) {
	contract, err := bindPlayerBookInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlayerBookInterfaceCaller{contract: contract}, nil
}

// NewPlayerBookInterfaceTransactor creates a new write-only instance of PlayerBookInterface, bound to a specific deployed contract.
func NewPlayerBookInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*PlayerBookInterfaceTransactor, error) {
	contract, err := bindPlayerBookInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlayerBookInterfaceTransactor{contract: contract}, nil
}

// NewPlayerBookInterfaceFilterer creates a new log filterer instance of PlayerBookInterface, bound to a specific deployed contract.
func NewPlayerBookInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*PlayerBookInterfaceFilterer, error) {
	contract, err := bindPlayerBookInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlayerBookInterfaceFilterer{contract: contract}, nil
}

// bindPlayerBookInterface binds a generic wrapper to an already deployed contract.
func bindPlayerBookInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlayerBookInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlayerBookInterface *PlayerBookInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlayerBookInterface.Contract.PlayerBookInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlayerBookInterface *PlayerBookInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.PlayerBookInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlayerBookInterface *PlayerBookInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.PlayerBookInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlayerBookInterface *PlayerBookInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlayerBookInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlayerBookInterface *PlayerBookInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlayerBookInterface *PlayerBookInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.contract.Transact(opts, method, params...)
}

// GetNameFee is a free data retrieval call binding the contract method 0x2614195f.
//
// Solidity: function getNameFee() constant returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceCaller) GetNameFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlayerBookInterface.contract.Call(opts, out, "getNameFee")
	return *ret0, err
}

// GetNameFee is a free data retrieval call binding the contract method 0x2614195f.
//
// Solidity: function getNameFee() constant returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceSession) GetNameFee() (*big.Int, error) {
	return _PlayerBookInterface.Contract.GetNameFee(&_PlayerBookInterface.CallOpts)
}

// GetNameFee is a free data retrieval call binding the contract method 0x2614195f.
//
// Solidity: function getNameFee() constant returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceCallerSession) GetNameFee() (*big.Int, error) {
	return _PlayerBookInterface.Contract.GetNameFee(&_PlayerBookInterface.CallOpts)
}

// GetPlayerAddr is a free data retrieval call binding the contract method 0x4d0d35ff.
//
// Solidity: function getPlayerAddr(_pID uint256) constant returns(address)
func (_PlayerBookInterface *PlayerBookInterfaceCaller) GetPlayerAddr(opts *bind.CallOpts, _pID *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlayerBookInterface.contract.Call(opts, out, "getPlayerAddr", _pID)
	return *ret0, err
}

// GetPlayerAddr is a free data retrieval call binding the contract method 0x4d0d35ff.
//
// Solidity: function getPlayerAddr(_pID uint256) constant returns(address)
func (_PlayerBookInterface *PlayerBookInterfaceSession) GetPlayerAddr(_pID *big.Int) (common.Address, error) {
	return _PlayerBookInterface.Contract.GetPlayerAddr(&_PlayerBookInterface.CallOpts, _pID)
}

// GetPlayerAddr is a free data retrieval call binding the contract method 0x4d0d35ff.
//
// Solidity: function getPlayerAddr(_pID uint256) constant returns(address)
func (_PlayerBookInterface *PlayerBookInterfaceCallerSession) GetPlayerAddr(_pID *big.Int) (common.Address, error) {
	return _PlayerBookInterface.Contract.GetPlayerAddr(&_PlayerBookInterface.CallOpts, _pID)
}

// GetPlayerLAff is a free data retrieval call binding the contract method 0xe3c08adf.
//
// Solidity: function getPlayerLAff(_pID uint256) constant returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceCaller) GetPlayerLAff(opts *bind.CallOpts, _pID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlayerBookInterface.contract.Call(opts, out, "getPlayerLAff", _pID)
	return *ret0, err
}

// GetPlayerLAff is a free data retrieval call binding the contract method 0xe3c08adf.
//
// Solidity: function getPlayerLAff(_pID uint256) constant returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceSession) GetPlayerLAff(_pID *big.Int) (*big.Int, error) {
	return _PlayerBookInterface.Contract.GetPlayerLAff(&_PlayerBookInterface.CallOpts, _pID)
}

// GetPlayerLAff is a free data retrieval call binding the contract method 0xe3c08adf.
//
// Solidity: function getPlayerLAff(_pID uint256) constant returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceCallerSession) GetPlayerLAff(_pID *big.Int) (*big.Int, error) {
	return _PlayerBookInterface.Contract.GetPlayerLAff(&_PlayerBookInterface.CallOpts, _pID)
}

// GetPlayerName is a free data retrieval call binding the contract method 0x82e37b2c.
//
// Solidity: function getPlayerName(_pID uint256) constant returns(bytes32)
func (_PlayerBookInterface *PlayerBookInterfaceCaller) GetPlayerName(opts *bind.CallOpts, _pID *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PlayerBookInterface.contract.Call(opts, out, "getPlayerName", _pID)
	return *ret0, err
}

// GetPlayerName is a free data retrieval call binding the contract method 0x82e37b2c.
//
// Solidity: function getPlayerName(_pID uint256) constant returns(bytes32)
func (_PlayerBookInterface *PlayerBookInterfaceSession) GetPlayerName(_pID *big.Int) ([32]byte, error) {
	return _PlayerBookInterface.Contract.GetPlayerName(&_PlayerBookInterface.CallOpts, _pID)
}

// GetPlayerName is a free data retrieval call binding the contract method 0x82e37b2c.
//
// Solidity: function getPlayerName(_pID uint256) constant returns(bytes32)
func (_PlayerBookInterface *PlayerBookInterfaceCallerSession) GetPlayerName(_pID *big.Int) ([32]byte, error) {
	return _PlayerBookInterface.Contract.GetPlayerName(&_PlayerBookInterface.CallOpts, _pID)
}

// GetPlayerID is a paid mutator transaction binding the contract method 0xe56556a9.
//
// Solidity: function getPlayerID(_addr address) returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactor) GetPlayerID(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _PlayerBookInterface.contract.Transact(opts, "getPlayerID", _addr)
}

// GetPlayerID is a paid mutator transaction binding the contract method 0xe56556a9.
//
// Solidity: function getPlayerID(_addr address) returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceSession) GetPlayerID(_addr common.Address) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.GetPlayerID(&_PlayerBookInterface.TransactOpts, _addr)
}

// GetPlayerID is a paid mutator transaction binding the contract method 0xe56556a9.
//
// Solidity: function getPlayerID(_addr address) returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactorSession) GetPlayerID(_addr common.Address) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.GetPlayerID(&_PlayerBookInterface.TransactOpts, _addr)
}

// RegisterNameXIDFromDapp is a paid mutator transaction binding the contract method 0xc0942dfd.
//
// Solidity: function registerNameXIDFromDapp(_addr address, _name bytes32, _affCode uint256, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactor) RegisterNameXIDFromDapp(opts *bind.TransactOpts, _addr common.Address, _name [32]byte, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.contract.Transact(opts, "registerNameXIDFromDapp", _addr, _name, _affCode, _all)
}

// RegisterNameXIDFromDapp is a paid mutator transaction binding the contract method 0xc0942dfd.
//
// Solidity: function registerNameXIDFromDapp(_addr address, _name bytes32, _affCode uint256, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceSession) RegisterNameXIDFromDapp(_addr common.Address, _name [32]byte, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.RegisterNameXIDFromDapp(&_PlayerBookInterface.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXIDFromDapp is a paid mutator transaction binding the contract method 0xc0942dfd.
//
// Solidity: function registerNameXIDFromDapp(_addr address, _name bytes32, _affCode uint256, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactorSession) RegisterNameXIDFromDapp(_addr common.Address, _name [32]byte, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.RegisterNameXIDFromDapp(&_PlayerBookInterface.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXaddrFromDapp is a paid mutator transaction binding the contract method 0xaa4d490b.
//
// Solidity: function registerNameXaddrFromDapp(_addr address, _name bytes32, _affCode address, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactor) RegisterNameXaddrFromDapp(opts *bind.TransactOpts, _addr common.Address, _name [32]byte, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.contract.Transact(opts, "registerNameXaddrFromDapp", _addr, _name, _affCode, _all)
}

// RegisterNameXaddrFromDapp is a paid mutator transaction binding the contract method 0xaa4d490b.
//
// Solidity: function registerNameXaddrFromDapp(_addr address, _name bytes32, _affCode address, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceSession) RegisterNameXaddrFromDapp(_addr common.Address, _name [32]byte, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.RegisterNameXaddrFromDapp(&_PlayerBookInterface.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXaddrFromDapp is a paid mutator transaction binding the contract method 0xaa4d490b.
//
// Solidity: function registerNameXaddrFromDapp(_addr address, _name bytes32, _affCode address, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactorSession) RegisterNameXaddrFromDapp(_addr common.Address, _name [32]byte, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.RegisterNameXaddrFromDapp(&_PlayerBookInterface.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXnameFromDapp is a paid mutator transaction binding the contract method 0x745ea0c1.
//
// Solidity: function registerNameXnameFromDapp(_addr address, _name bytes32, _affCode bytes32, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactor) RegisterNameXnameFromDapp(opts *bind.TransactOpts, _addr common.Address, _name [32]byte, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.contract.Transact(opts, "registerNameXnameFromDapp", _addr, _name, _affCode, _all)
}

// RegisterNameXnameFromDapp is a paid mutator transaction binding the contract method 0x745ea0c1.
//
// Solidity: function registerNameXnameFromDapp(_addr address, _name bytes32, _affCode bytes32, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceSession) RegisterNameXnameFromDapp(_addr common.Address, _name [32]byte, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.RegisterNameXnameFromDapp(&_PlayerBookInterface.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXnameFromDapp is a paid mutator transaction binding the contract method 0x745ea0c1.
//
// Solidity: function registerNameXnameFromDapp(_addr address, _name bytes32, _affCode bytes32, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactorSession) RegisterNameXnameFromDapp(_addr common.Address, _name [32]byte, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.RegisterNameXnameFromDapp(&_PlayerBookInterface.TransactOpts, _addr, _name, _affCode, _all)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582075c9674936a3eebcaa9e01abd9f37e4eb94f4becfe2b269a138c2a13a4ba08d10029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// ModularLongABI is the input ABI used to generate the binding from.
const ModularLongABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isNewPlayer\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onNewName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"etcIn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"keysBought\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"PSLAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"potAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"airDropPot\",\"type\":\"uint256\"}],\"name\":\"onEndTx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"etcOut\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"etcOut\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"PSLAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onWithdrawAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"etcIn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"PSLAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onBuyAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"PSLAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onReLoadAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"roundID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"buyerID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onAffiliatePayout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"roundID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountAddedToPot\",\"type\":\"uint256\"}],\"name\":\"onPotSwapDeposit\",\"type\":\"event\"}]"

// ModularLongBin is the compiled bytecode used for deploying new contracts.
const ModularLongBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a72305820a195b2aa576afba3349e1865429440ea89161f8a10e40392b02d4deabfd64f6e0029`

// DeployModularLong deploys a new Ethereum contract, binding an instance of ModularLong to it.
func DeployModularLong(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ModularLong, error) {
	parsed, err := abi.JSON(strings.NewReader(ModularLongABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ModularLongBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ModularLong{ModularLongCaller: ModularLongCaller{contract: contract}, ModularLongTransactor: ModularLongTransactor{contract: contract}, ModularLongFilterer: ModularLongFilterer{contract: contract}}, nil
}

// ModularLong is an auto generated Go binding around an Ethereum contract.
type ModularLong struct {
	ModularLongCaller     // Read-only binding to the contract
	ModularLongTransactor // Write-only binding to the contract
	ModularLongFilterer   // Log filterer for contract events
}

// ModularLongCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModularLongCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModularLongTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModularLongTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModularLongFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModularLongFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModularLongSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModularLongSession struct {
	Contract     *ModularLong      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModularLongCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModularLongCallerSession struct {
	Contract *ModularLongCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ModularLongTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModularLongTransactorSession struct {
	Contract     *ModularLongTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ModularLongRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModularLongRaw struct {
	Contract *ModularLong // Generic contract binding to access the raw methods on
}

// ModularLongCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModularLongCallerRaw struct {
	Contract *ModularLongCaller // Generic read-only contract binding to access the raw methods on
}

// ModularLongTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModularLongTransactorRaw struct {
	Contract *ModularLongTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModularLong creates a new instance of ModularLong, bound to a specific deployed contract.
func NewModularLong(address common.Address, backend bind.ContractBackend) (*ModularLong, error) {
	contract, err := bindModularLong(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ModularLong{ModularLongCaller: ModularLongCaller{contract: contract}, ModularLongTransactor: ModularLongTransactor{contract: contract}, ModularLongFilterer: ModularLongFilterer{contract: contract}}, nil
}

// NewModularLongCaller creates a new read-only instance of ModularLong, bound to a specific deployed contract.
func NewModularLongCaller(address common.Address, caller bind.ContractCaller) (*ModularLongCaller, error) {
	contract, err := bindModularLong(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModularLongCaller{contract: contract}, nil
}

// NewModularLongTransactor creates a new write-only instance of ModularLong, bound to a specific deployed contract.
func NewModularLongTransactor(address common.Address, transactor bind.ContractTransactor) (*ModularLongTransactor, error) {
	contract, err := bindModularLong(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModularLongTransactor{contract: contract}, nil
}

// NewModularLongFilterer creates a new log filterer instance of ModularLong, bound to a specific deployed contract.
func NewModularLongFilterer(address common.Address, filterer bind.ContractFilterer) (*ModularLongFilterer, error) {
	contract, err := bindModularLong(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModularLongFilterer{contract: contract}, nil
}

// bindModularLong binds a generic wrapper to an already deployed contract.
func bindModularLong(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ModularLongABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModularLong *ModularLongRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ModularLong.Contract.ModularLongCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModularLong *ModularLongRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModularLong.Contract.ModularLongTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModularLong *ModularLongRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModularLong.Contract.ModularLongTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModularLong *ModularLongCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ModularLong.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModularLong *ModularLongTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModularLong.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModularLong *ModularLongTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModularLong.Contract.contract.Transact(opts, method, params...)
}

// ModularLongOnAffiliatePayoutIterator is returned from FilterOnAffiliatePayout and is used to iterate over the raw logs and unpacked data for OnAffiliatePayout events raised by the ModularLong contract.
type ModularLongOnAffiliatePayoutIterator struct {
	Event *ModularLongOnAffiliatePayout // Event containing the contract specifics and raw log

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
func (it *ModularLongOnAffiliatePayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnAffiliatePayout)
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
		it.Event = new(ModularLongOnAffiliatePayout)
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
func (it *ModularLongOnAffiliatePayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnAffiliatePayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnAffiliatePayout represents a OnAffiliatePayout event raised by the ModularLong contract.
type ModularLongOnAffiliatePayout struct {
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	RoundID          *big.Int
	BuyerID          *big.Int
	Amount           *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnAffiliatePayout is a free log retrieval operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: event onAffiliatePayout(affiliateID indexed uint256, affiliateAddress address, affiliateName bytes32, roundID indexed uint256, buyerID indexed uint256, amount uint256, timeStamp uint256)
func (_ModularLong *ModularLongFilterer) FilterOnAffiliatePayout(opts *bind.FilterOpts, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (*ModularLongOnAffiliatePayoutIterator, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return &ModularLongOnAffiliatePayoutIterator{contract: _ModularLong.contract, event: "onAffiliatePayout", logs: logs, sub: sub}, nil
}

// WatchOnAffiliatePayout is a free log subscription operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: event onAffiliatePayout(affiliateID indexed uint256, affiliateAddress address, affiliateName bytes32, roundID indexed uint256, buyerID indexed uint256, amount uint256, timeStamp uint256)
func (_ModularLong *ModularLongFilterer) WatchOnAffiliatePayout(opts *bind.WatchOpts, sink chan<- *ModularLongOnAffiliatePayout, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (event.Subscription, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnAffiliatePayout)
				if err := _ModularLong.contract.UnpackLog(event, "onAffiliatePayout", log); err != nil {
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

// ModularLongOnBuyAndDistributeIterator is returned from FilterOnBuyAndDistribute and is used to iterate over the raw logs and unpacked data for OnBuyAndDistribute events raised by the ModularLong contract.
type ModularLongOnBuyAndDistributeIterator struct {
	Event *ModularLongOnBuyAndDistribute // Event containing the contract specifics and raw log

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
func (it *ModularLongOnBuyAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnBuyAndDistribute)
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
		it.Event = new(ModularLongOnBuyAndDistribute)
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
func (it *ModularLongOnBuyAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnBuyAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnBuyAndDistribute represents a OnBuyAndDistribute event raised by the ModularLong contract.
type ModularLongOnBuyAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EtcIn          *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	PSLAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnBuyAndDistribute is a free log retrieval operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: event onBuyAndDistribute(playerAddress address, playerName bytes32, etcIn uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_ModularLong *ModularLongFilterer) FilterOnBuyAndDistribute(opts *bind.FilterOpts) (*ModularLongOnBuyAndDistributeIterator, error) {

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return &ModularLongOnBuyAndDistributeIterator{contract: _ModularLong.contract, event: "onBuyAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnBuyAndDistribute is a free log subscription operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: event onBuyAndDistribute(playerAddress address, playerName bytes32, etcIn uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_ModularLong *ModularLongFilterer) WatchOnBuyAndDistribute(opts *bind.WatchOpts, sink chan<- *ModularLongOnBuyAndDistribute) (event.Subscription, error) {

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnBuyAndDistribute)
				if err := _ModularLong.contract.UnpackLog(event, "onBuyAndDistribute", log); err != nil {
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

// ModularLongOnEndTxIterator is returned from FilterOnEndTx and is used to iterate over the raw logs and unpacked data for OnEndTx events raised by the ModularLong contract.
type ModularLongOnEndTxIterator struct {
	Event *ModularLongOnEndTx // Event containing the contract specifics and raw log

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
func (it *ModularLongOnEndTxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnEndTx)
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
		it.Event = new(ModularLongOnEndTx)
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
func (it *ModularLongOnEndTxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnEndTxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnEndTx represents a OnEndTx event raised by the ModularLong contract.
type ModularLongOnEndTx struct {
	CompressedData *big.Int
	CompressedIDs  *big.Int
	PlayerName     [32]byte
	PlayerAddress  common.Address
	EtcIn          *big.Int
	KeysBought     *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	PSLAmount      *big.Int
	GenAmount      *big.Int
	PotAmount      *big.Int
	AirDropPot     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnEndTx is a free log retrieval operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: event onEndTx(compressedData uint256, compressedIDs uint256, playerName bytes32, playerAddress address, etcIn uint256, keysBought uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256, potAmount uint256, airDropPot uint256)
func (_ModularLong *ModularLongFilterer) FilterOnEndTx(opts *bind.FilterOpts) (*ModularLongOnEndTxIterator, error) {

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return &ModularLongOnEndTxIterator{contract: _ModularLong.contract, event: "onEndTx", logs: logs, sub: sub}, nil
}

// WatchOnEndTx is a free log subscription operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: event onEndTx(compressedData uint256, compressedIDs uint256, playerName bytes32, playerAddress address, etcIn uint256, keysBought uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256, potAmount uint256, airDropPot uint256)
func (_ModularLong *ModularLongFilterer) WatchOnEndTx(opts *bind.WatchOpts, sink chan<- *ModularLongOnEndTx) (event.Subscription, error) {

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnEndTx)
				if err := _ModularLong.contract.UnpackLog(event, "onEndTx", log); err != nil {
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

// ModularLongOnNewNameIterator is returned from FilterOnNewName and is used to iterate over the raw logs and unpacked data for OnNewName events raised by the ModularLong contract.
type ModularLongOnNewNameIterator struct {
	Event *ModularLongOnNewName // Event containing the contract specifics and raw log

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
func (it *ModularLongOnNewNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnNewName)
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
		it.Event = new(ModularLongOnNewName)
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
func (it *ModularLongOnNewNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnNewNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnNewName represents a OnNewName event raised by the ModularLong contract.
type ModularLongOnNewName struct {
	PlayerID         *big.Int
	PlayerAddress    common.Address
	PlayerName       [32]byte
	IsNewPlayer      bool
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	AmountPaid       *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnNewName is a free log retrieval operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: event onNewName(playerID indexed uint256, playerAddress indexed address, playerName indexed bytes32, isNewPlayer bool, affiliateID uint256, affiliateAddress address, affiliateName bytes32, amountPaid uint256, timeStamp uint256)
func (_ModularLong *ModularLongFilterer) FilterOnNewName(opts *bind.FilterOpts, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (*ModularLongOnNewNameIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return &ModularLongOnNewNameIterator{contract: _ModularLong.contract, event: "onNewName", logs: logs, sub: sub}, nil
}

// WatchOnNewName is a free log subscription operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: event onNewName(playerID indexed uint256, playerAddress indexed address, playerName indexed bytes32, isNewPlayer bool, affiliateID uint256, affiliateAddress address, affiliateName bytes32, amountPaid uint256, timeStamp uint256)
func (_ModularLong *ModularLongFilterer) WatchOnNewName(opts *bind.WatchOpts, sink chan<- *ModularLongOnNewName, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnNewName)
				if err := _ModularLong.contract.UnpackLog(event, "onNewName", log); err != nil {
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

// ModularLongOnPotSwapDepositIterator is returned from FilterOnPotSwapDeposit and is used to iterate over the raw logs and unpacked data for OnPotSwapDeposit events raised by the ModularLong contract.
type ModularLongOnPotSwapDepositIterator struct {
	Event *ModularLongOnPotSwapDeposit // Event containing the contract specifics and raw log

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
func (it *ModularLongOnPotSwapDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnPotSwapDeposit)
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
		it.Event = new(ModularLongOnPotSwapDeposit)
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
func (it *ModularLongOnPotSwapDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnPotSwapDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnPotSwapDeposit represents a OnPotSwapDeposit event raised by the ModularLong contract.
type ModularLongOnPotSwapDeposit struct {
	RoundID          *big.Int
	AmountAddedToPot *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnPotSwapDeposit is a free log retrieval operation binding the contract event 0x74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c.
//
// Solidity: event onPotSwapDeposit(roundID uint256, amountAddedToPot uint256)
func (_ModularLong *ModularLongFilterer) FilterOnPotSwapDeposit(opts *bind.FilterOpts) (*ModularLongOnPotSwapDepositIterator, error) {

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onPotSwapDeposit")
	if err != nil {
		return nil, err
	}
	return &ModularLongOnPotSwapDepositIterator{contract: _ModularLong.contract, event: "onPotSwapDeposit", logs: logs, sub: sub}, nil
}

// WatchOnPotSwapDeposit is a free log subscription operation binding the contract event 0x74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c.
//
// Solidity: event onPotSwapDeposit(roundID uint256, amountAddedToPot uint256)
func (_ModularLong *ModularLongFilterer) WatchOnPotSwapDeposit(opts *bind.WatchOpts, sink chan<- *ModularLongOnPotSwapDeposit) (event.Subscription, error) {

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onPotSwapDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnPotSwapDeposit)
				if err := _ModularLong.contract.UnpackLog(event, "onPotSwapDeposit", log); err != nil {
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

// ModularLongOnReLoadAndDistributeIterator is returned from FilterOnReLoadAndDistribute and is used to iterate over the raw logs and unpacked data for OnReLoadAndDistribute events raised by the ModularLong contract.
type ModularLongOnReLoadAndDistributeIterator struct {
	Event *ModularLongOnReLoadAndDistribute // Event containing the contract specifics and raw log

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
func (it *ModularLongOnReLoadAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnReLoadAndDistribute)
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
		it.Event = new(ModularLongOnReLoadAndDistribute)
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
func (it *ModularLongOnReLoadAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnReLoadAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnReLoadAndDistribute represents a OnReLoadAndDistribute event raised by the ModularLong contract.
type ModularLongOnReLoadAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	PSLAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnReLoadAndDistribute is a free log retrieval operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: event onReLoadAndDistribute(playerAddress address, playerName bytes32, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_ModularLong *ModularLongFilterer) FilterOnReLoadAndDistribute(opts *bind.FilterOpts) (*ModularLongOnReLoadAndDistributeIterator, error) {

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return &ModularLongOnReLoadAndDistributeIterator{contract: _ModularLong.contract, event: "onReLoadAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnReLoadAndDistribute is a free log subscription operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: event onReLoadAndDistribute(playerAddress address, playerName bytes32, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_ModularLong *ModularLongFilterer) WatchOnReLoadAndDistribute(opts *bind.WatchOpts, sink chan<- *ModularLongOnReLoadAndDistribute) (event.Subscription, error) {

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnReLoadAndDistribute)
				if err := _ModularLong.contract.UnpackLog(event, "onReLoadAndDistribute", log); err != nil {
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

// ModularLongOnWithdrawIterator is returned from FilterOnWithdraw and is used to iterate over the raw logs and unpacked data for OnWithdraw events raised by the ModularLong contract.
type ModularLongOnWithdrawIterator struct {
	Event *ModularLongOnWithdraw // Event containing the contract specifics and raw log

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
func (it *ModularLongOnWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnWithdraw)
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
		it.Event = new(ModularLongOnWithdraw)
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
func (it *ModularLongOnWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnWithdraw represents a OnWithdraw event raised by the ModularLong contract.
type ModularLongOnWithdraw struct {
	PlayerID      *big.Int
	PlayerAddress common.Address
	PlayerName    [32]byte
	EtcOut        *big.Int
	TimeStamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOnWithdraw is a free log retrieval operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: event onWithdraw(playerID indexed uint256, playerAddress address, playerName bytes32, etcOut uint256, timeStamp uint256)
func (_ModularLong *ModularLongFilterer) FilterOnWithdraw(opts *bind.FilterOpts, playerID []*big.Int) (*ModularLongOnWithdrawIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return &ModularLongOnWithdrawIterator{contract: _ModularLong.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: event onWithdraw(playerID indexed uint256, playerAddress address, playerName bytes32, etcOut uint256, timeStamp uint256)
func (_ModularLong *ModularLongFilterer) WatchOnWithdraw(opts *bind.WatchOpts, sink chan<- *ModularLongOnWithdraw, playerID []*big.Int) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnWithdraw)
				if err := _ModularLong.contract.UnpackLog(event, "onWithdraw", log); err != nil {
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

// ModularLongOnWithdrawAndDistributeIterator is returned from FilterOnWithdrawAndDistribute and is used to iterate over the raw logs and unpacked data for OnWithdrawAndDistribute events raised by the ModularLong contract.
type ModularLongOnWithdrawAndDistributeIterator struct {
	Event *ModularLongOnWithdrawAndDistribute // Event containing the contract specifics and raw log

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
func (it *ModularLongOnWithdrawAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnWithdrawAndDistribute)
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
		it.Event = new(ModularLongOnWithdrawAndDistribute)
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
func (it *ModularLongOnWithdrawAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnWithdrawAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnWithdrawAndDistribute represents a OnWithdrawAndDistribute event raised by the ModularLong contract.
type ModularLongOnWithdrawAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EtcOut         *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	PSLAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnWithdrawAndDistribute is a free log retrieval operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: event onWithdrawAndDistribute(playerAddress address, playerName bytes32, etcOut uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_ModularLong *ModularLongFilterer) FilterOnWithdrawAndDistribute(opts *bind.FilterOpts) (*ModularLongOnWithdrawAndDistributeIterator, error) {

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return &ModularLongOnWithdrawAndDistributeIterator{contract: _ModularLong.contract, event: "onWithdrawAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnWithdrawAndDistribute is a free log subscription operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: event onWithdrawAndDistribute(playerAddress address, playerName bytes32, etcOut uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, PSLAmount uint256, genAmount uint256)
func (_ModularLong *ModularLongFilterer) WatchOnWithdrawAndDistribute(opts *bind.WatchOpts, sink chan<- *ModularLongOnWithdrawAndDistribute) (event.Subscription, error) {

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnWithdrawAndDistribute)
				if err := _ModularLong.contract.UnpackLog(event, "onWithdrawAndDistribute", log); err != nil {
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

// OtherFomoETCABI is the input ABI used to generate the binding from.
const OtherFomoETCABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"potSwap\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// OtherFomoETCBin is the compiled bytecode used for deploying new contracts.
const OtherFomoETCBin = `0x`

// DeployOtherFomoETC deploys a new Ethereum contract, binding an instance of OtherFomoETC to it.
func DeployOtherFomoETC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OtherFomoETC, error) {
	parsed, err := abi.JSON(strings.NewReader(OtherFomoETCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OtherFomoETCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OtherFomoETC{OtherFomoETCCaller: OtherFomoETCCaller{contract: contract}, OtherFomoETCTransactor: OtherFomoETCTransactor{contract: contract}, OtherFomoETCFilterer: OtherFomoETCFilterer{contract: contract}}, nil
}

// OtherFomoETC is an auto generated Go binding around an Ethereum contract.
type OtherFomoETC struct {
	OtherFomoETCCaller     // Read-only binding to the contract
	OtherFomoETCTransactor // Write-only binding to the contract
	OtherFomoETCFilterer   // Log filterer for contract events
}

// OtherFomoETCCaller is an auto generated read-only Go binding around an Ethereum contract.
type OtherFomoETCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtherFomoETCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OtherFomoETCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtherFomoETCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OtherFomoETCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtherFomoETCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OtherFomoETCSession struct {
	Contract     *OtherFomoETC     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OtherFomoETCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OtherFomoETCCallerSession struct {
	Contract *OtherFomoETCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OtherFomoETCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OtherFomoETCTransactorSession struct {
	Contract     *OtherFomoETCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OtherFomoETCRaw is an auto generated low-level Go binding around an Ethereum contract.
type OtherFomoETCRaw struct {
	Contract *OtherFomoETC // Generic contract binding to access the raw methods on
}

// OtherFomoETCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OtherFomoETCCallerRaw struct {
	Contract *OtherFomoETCCaller // Generic read-only contract binding to access the raw methods on
}

// OtherFomoETCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OtherFomoETCTransactorRaw struct {
	Contract *OtherFomoETCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOtherFomoETC creates a new instance of OtherFomoETC, bound to a specific deployed contract.
func NewOtherFomoETC(address common.Address, backend bind.ContractBackend) (*OtherFomoETC, error) {
	contract, err := bindOtherFomoETC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OtherFomoETC{OtherFomoETCCaller: OtherFomoETCCaller{contract: contract}, OtherFomoETCTransactor: OtherFomoETCTransactor{contract: contract}, OtherFomoETCFilterer: OtherFomoETCFilterer{contract: contract}}, nil
}

// NewOtherFomoETCCaller creates a new read-only instance of OtherFomoETC, bound to a specific deployed contract.
func NewOtherFomoETCCaller(address common.Address, caller bind.ContractCaller) (*OtherFomoETCCaller, error) {
	contract, err := bindOtherFomoETC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OtherFomoETCCaller{contract: contract}, nil
}

// NewOtherFomoETCTransactor creates a new write-only instance of OtherFomoETC, bound to a specific deployed contract.
func NewOtherFomoETCTransactor(address common.Address, transactor bind.ContractTransactor) (*OtherFomoETCTransactor, error) {
	contract, err := bindOtherFomoETC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OtherFomoETCTransactor{contract: contract}, nil
}

// NewOtherFomoETCFilterer creates a new log filterer instance of OtherFomoETC, bound to a specific deployed contract.
func NewOtherFomoETCFilterer(address common.Address, filterer bind.ContractFilterer) (*OtherFomoETCFilterer, error) {
	contract, err := bindOtherFomoETC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OtherFomoETCFilterer{contract: contract}, nil
}

// bindOtherFomoETC binds a generic wrapper to an already deployed contract.
func bindOtherFomoETC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OtherFomoETCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OtherFomoETC *OtherFomoETCRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OtherFomoETC.Contract.OtherFomoETCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OtherFomoETC *OtherFomoETCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtherFomoETC.Contract.OtherFomoETCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OtherFomoETC *OtherFomoETCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OtherFomoETC.Contract.OtherFomoETCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OtherFomoETC *OtherFomoETCCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OtherFomoETC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OtherFomoETC *OtherFomoETCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtherFomoETC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OtherFomoETC *OtherFomoETCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OtherFomoETC.Contract.contract.Transact(opts, method, params...)
}

// PotSwap is a paid mutator transaction binding the contract method 0xed78cf4a.
//
// Solidity: function potSwap() returns()
func (_OtherFomoETC *OtherFomoETCTransactor) PotSwap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtherFomoETC.contract.Transact(opts, "potSwap")
}

// PotSwap is a paid mutator transaction binding the contract method 0xed78cf4a.
//
// Solidity: function potSwap() returns()
func (_OtherFomoETC *OtherFomoETCSession) PotSwap() (*types.Transaction, error) {
	return _OtherFomoETC.Contract.PotSwap(&_OtherFomoETC.TransactOpts)
}

// PotSwap is a paid mutator transaction binding the contract method 0xed78cf4a.
//
// Solidity: function potSwap() returns()
func (_OtherFomoETC *OtherFomoETCTransactorSession) PotSwap() (*types.Transaction, error) {
	return _OtherFomoETC.Contract.PotSwap(&_OtherFomoETC.TransactOpts)
}
