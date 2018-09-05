// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testinterface

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DiviesABI is the input ABI used to generate the binding from.
const DiviesABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// DiviesBin is the compiled bytecode used for deploying new contracts.
const DiviesBin = `0x6080604052348015600f57600080fd5b50608e8061001e6000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663d0e30db081146043575b600080fd5b6049605b565b60408051918252519081900360200190f35b62abcc53905600a165627a7a723058205ad7b065444ae42622bb9505a1e04aa1b5bc6b22096f3e3271c284ef5006fe040029`

// DeployDivies deploys a new Ethereum contract, binding an instance of Divies to it.
func DeployDivies(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Divies, error) {
	parsed, err := abi.JSON(strings.NewReader(DiviesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DiviesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Divies{DiviesCaller: DiviesCaller{contract: contract}, DiviesTransactor: DiviesTransactor{contract: contract}, DiviesFilterer: DiviesFilterer{contract: contract}}, nil
}

// Divies is an auto generated Go binding around an Ethereum contract.
type Divies struct {
	DiviesCaller     // Read-only binding to the contract
	DiviesTransactor // Write-only binding to the contract
	DiviesFilterer   // Log filterer for contract events
}

// DiviesCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiviesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiviesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiviesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiviesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiviesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiviesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiviesSession struct {
	Contract     *Divies           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiviesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiviesCallerSession struct {
	Contract *DiviesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DiviesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiviesTransactorSession struct {
	Contract     *DiviesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiviesRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiviesRaw struct {
	Contract *Divies // Generic contract binding to access the raw methods on
}

// DiviesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiviesCallerRaw struct {
	Contract *DiviesCaller // Generic read-only contract binding to access the raw methods on
}

// DiviesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiviesTransactorRaw struct {
	Contract *DiviesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDivies creates a new instance of Divies, bound to a specific deployed contract.
func NewDivies(address common.Address, backend bind.ContractBackend) (*Divies, error) {
	contract, err := bindDivies(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Divies{DiviesCaller: DiviesCaller{contract: contract}, DiviesTransactor: DiviesTransactor{contract: contract}, DiviesFilterer: DiviesFilterer{contract: contract}}, nil
}

// NewDiviesCaller creates a new read-only instance of Divies, bound to a specific deployed contract.
func NewDiviesCaller(address common.Address, caller bind.ContractCaller) (*DiviesCaller, error) {
	contract, err := bindDivies(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiviesCaller{contract: contract}, nil
}

// NewDiviesTransactor creates a new write-only instance of Divies, bound to a specific deployed contract.
func NewDiviesTransactor(address common.Address, transactor bind.ContractTransactor) (*DiviesTransactor, error) {
	contract, err := bindDivies(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiviesTransactor{contract: contract}, nil
}

// NewDiviesFilterer creates a new log filterer instance of Divies, bound to a specific deployed contract.
func NewDiviesFilterer(address common.Address, filterer bind.ContractFilterer) (*DiviesFilterer, error) {
	contract, err := bindDivies(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiviesFilterer{contract: contract}, nil
}

// bindDivies binds a generic wrapper to an already deployed contract.
func bindDivies(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiviesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Divies *DiviesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Divies.Contract.DiviesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Divies *DiviesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Divies.Contract.DiviesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Divies *DiviesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Divies.Contract.DiviesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Divies *DiviesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Divies.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Divies *DiviesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Divies.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Divies *DiviesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Divies.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Divies *DiviesTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Divies.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Divies *DiviesSession) Deposit() (*types.Transaction, error) {
	return _Divies.Contract.Deposit(&_Divies.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_Divies *DiviesTransactorSession) Deposit() (*types.Transaction, error) {
	return _Divies.Contract.Deposit(&_Divies.TransactOpts)
}

// DiviesInterfaceABI is the input ABI used to generate the binding from.
const DiviesInterfaceABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

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
// Solidity: function deposit() returns(uint256)
func (_DiviesInterface *DiviesInterfaceTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiviesInterface.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_DiviesInterface *DiviesInterfaceSession) Deposit() (*types.Transaction, error) {
	return _DiviesInterface.Contract.Deposit(&_DiviesInterface.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_DiviesInterface *DiviesInterfaceTransactorSession) Deposit() (*types.Transaction, error) {
	return _DiviesInterface.Contract.Deposit(&_DiviesInterface.TransactOpts)
}

// FoMo3DlongABI is the input ABI used to generate the binding from.
const FoMo3DlongABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"what\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FoMo3DlongBin is the compiled bytecode used for deploying new contracts.
const FoMo3DlongBin = `0x608060405234801561001057600080fd5b50610127806100206000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663b24bb84581146043575b600080fd5b348015604e57600080fd5b5060556067565b60408051918252519081900360200190f35b60008073ffffffffffffffffffffffffffffffffffffffff1663d0e30db06040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801560cc57600080fd5b505af115801560df573d6000803e3d6000fd5b505050506040513d602081101560f457600080fd5b50519050905600a165627a7a72305820ddce7b3b24e91aed3ed0c0c97a5b0b2108bda1cf32503d628bb811b08961f4d70029`

// DeployFoMo3Dlong deploys a new Ethereum contract, binding an instance of FoMo3Dlong to it.
func DeployFoMo3Dlong(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FoMo3Dlong, error) {
	parsed, err := abi.JSON(strings.NewReader(FoMo3DlongABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FoMo3DlongBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FoMo3Dlong{FoMo3DlongCaller: FoMo3DlongCaller{contract: contract}, FoMo3DlongTransactor: FoMo3DlongTransactor{contract: contract}, FoMo3DlongFilterer: FoMo3DlongFilterer{contract: contract}}, nil
}

// FoMo3Dlong is an auto generated Go binding around an Ethereum contract.
type FoMo3Dlong struct {
	FoMo3DlongCaller     // Read-only binding to the contract
	FoMo3DlongTransactor // Write-only binding to the contract
	FoMo3DlongFilterer   // Log filterer for contract events
}

// FoMo3DlongCaller is an auto generated read-only Go binding around an Ethereum contract.
type FoMo3DlongCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoMo3DlongTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FoMo3DlongTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoMo3DlongFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FoMo3DlongFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoMo3DlongSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FoMo3DlongSession struct {
	Contract     *FoMo3Dlong       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FoMo3DlongCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FoMo3DlongCallerSession struct {
	Contract *FoMo3DlongCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FoMo3DlongTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FoMo3DlongTransactorSession struct {
	Contract     *FoMo3DlongTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FoMo3DlongRaw is an auto generated low-level Go binding around an Ethereum contract.
type FoMo3DlongRaw struct {
	Contract *FoMo3Dlong // Generic contract binding to access the raw methods on
}

// FoMo3DlongCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FoMo3DlongCallerRaw struct {
	Contract *FoMo3DlongCaller // Generic read-only contract binding to access the raw methods on
}

// FoMo3DlongTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FoMo3DlongTransactorRaw struct {
	Contract *FoMo3DlongTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFoMo3Dlong creates a new instance of FoMo3Dlong, bound to a specific deployed contract.
func NewFoMo3Dlong(address common.Address, backend bind.ContractBackend) (*FoMo3Dlong, error) {
	contract, err := bindFoMo3Dlong(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FoMo3Dlong{FoMo3DlongCaller: FoMo3DlongCaller{contract: contract}, FoMo3DlongTransactor: FoMo3DlongTransactor{contract: contract}, FoMo3DlongFilterer: FoMo3DlongFilterer{contract: contract}}, nil
}

// NewFoMo3DlongCaller creates a new read-only instance of FoMo3Dlong, bound to a specific deployed contract.
func NewFoMo3DlongCaller(address common.Address, caller bind.ContractCaller) (*FoMo3DlongCaller, error) {
	contract, err := bindFoMo3Dlong(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongCaller{contract: contract}, nil
}

// NewFoMo3DlongTransactor creates a new write-only instance of FoMo3Dlong, bound to a specific deployed contract.
func NewFoMo3DlongTransactor(address common.Address, transactor bind.ContractTransactor) (*FoMo3DlongTransactor, error) {
	contract, err := bindFoMo3Dlong(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongTransactor{contract: contract}, nil
}

// NewFoMo3DlongFilterer creates a new log filterer instance of FoMo3Dlong, bound to a specific deployed contract.
func NewFoMo3DlongFilterer(address common.Address, filterer bind.ContractFilterer) (*FoMo3DlongFilterer, error) {
	contract, err := bindFoMo3Dlong(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongFilterer{contract: contract}, nil
}

// bindFoMo3Dlong binds a generic wrapper to an already deployed contract.
func bindFoMo3Dlong(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FoMo3DlongABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FoMo3Dlong *FoMo3DlongRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FoMo3Dlong.Contract.FoMo3DlongCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FoMo3Dlong *FoMo3DlongRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.FoMo3DlongTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FoMo3Dlong *FoMo3DlongRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.FoMo3DlongTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FoMo3Dlong *FoMo3DlongCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FoMo3Dlong.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FoMo3Dlong *FoMo3DlongTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FoMo3Dlong *FoMo3DlongTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.contract.Transact(opts, method, params...)
}

// What is a paid mutator transaction binding the contract method 0xb24bb845.
//
// Solidity: function what() returns(uint256)
func (_FoMo3Dlong *FoMo3DlongTransactor) What(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "what")
}

// What is a paid mutator transaction binding the contract method 0xb24bb845.
//
// Solidity: function what() returns(uint256)
func (_FoMo3Dlong *FoMo3DlongSession) What() (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.What(&_FoMo3Dlong.TransactOpts)
}

// What is a paid mutator transaction binding the contract method 0xb24bb845.
//
// Solidity: function what() returns(uint256)
func (_FoMo3Dlong *FoMo3DlongTransactorSession) What() (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.What(&_FoMo3Dlong.TransactOpts)
}
