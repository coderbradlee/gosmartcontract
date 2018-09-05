// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Jekyll_Island_Inc

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// JIincForwarderABI is the input ABI used to generate the binding from.
const JIincForwarderABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelMigration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_firstCorpBank\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMigration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCorpBank\",\"type\":\"address\"}],\"name\":\"startMigration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// JIincForwarderBin is the compiled bytecode used for deploying new contracts.
const JIincForwarderBin = `0x60c0604052600e60808190527f4a49696e63466f7277617264657200000000000000000000000000000000000060a090815261003e9160009190610078565b506002805460a060020a60ff0219167401000000000000000000000000000000000000000017905534801561007257600080fd5b50610113565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100b957805160ff19168380011785556100e6565b828001600101855582156100e6579182015b828111156100e65782518255916020019190600101906100cb565b506100f29291506100f6565b5090565b61011091905b808211156100f257600081556001016100fc565b90565b610943806101226000396000f3006080604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461011b57806310639ea0146101a5578063200d2ed2146101ce57806366d382031461020f57806388d761f214610232578063a0f52da014610247578063d0e30db014610268575b600154604080517ff340fa01000000000000000000000000000000000000000000000000000000008152600160a060020a03909216600483018190529051909163f340fa019130319160248082019260209290919082900301818588803b1580156100ec57600080fd5b505af1158015610100573d6000803e3d6000fd5b50505050506040513d602081101561011757600080fd5b5050005b34801561012757600080fd5b50610130610270565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561016a578181015183820152602001610152565b50505050905090810190601f1680156101975780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101b157600080fd5b506101ba6102fe565b604080519115158252519081900360200190f35b3480156101da57600080fd5b506101e36103d3565b60408051600160a060020a03948516815292909316602083015215158183015290519081900360600190f35b34801561021b57600080fd5b50610230600160a060020a0360043516610408565b005b34801561023e57600080fd5b506101ba6104f2565b34801561025357600080fd5b506101ba600160a060020a03600435166105db565b6101ba610752565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156102f65780601f106102cb576101008083540402835291602001916102f6565b820191906000526020600020905b8154815290600101906020018083116102d957829003601f168201915b505050505081565b600154600090600160a060020a031633146103af576040805160e560020a62461bcd02815260206004820152604760248201527f466f727761726465722063616e63656c4d6967726174696f6e206661696c656460448201527f202d206d73672e73656e646572206d7573742062652063757272656e7420636f60648201527f72702062616e6b00000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b506002805473ffffffffffffffffffffffffffffffffffffffff1916905560015b90565b600154600254600160a060020a0391821692918116917401000000000000000000000000000000000000000090910460ff1690565b60025474010000000000000000000000000000000000000000900460ff1615156001146104a5576040805160e560020a62461bcd02815260206004820152603560248201527f466f72776172646572207365747570206661696c6564202d20636f727020626160448201527f6e6b20616c726561647920726567697374657265640000000000000000000000606482015290519081900360840190fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790556002805474ff000000000000000000000000000000000000000019169055565b600254600090600160a060020a031633146105a3576040805160e560020a62461bcd02815260206004820152604360248201527f466f727761726465722066696e6973684d6967726174696f6e206661696c656460448201527f202d206d73672e73656e646572206d757374206265206e657720636f7270206260648201527f616e6b0000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b50600280546001805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617825590911690915590565b600154600090600160a060020a0316331461068c576040805160e560020a62461bcd02815260206004820152604660248201527f466f727761726465722073746172744d6967726174696f6e206661696c65642060448201527f2d206d73672e73656e646572206d7573742062652063757272656e7420636f7260648201527f702062616e6b0000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b81600160a060020a0316630839e0fb6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156106e357600080fd5b505af11580156106f7573d6000803e3d6000fd5b505050506040513d602081101561070d57600080fd5b505115156001141561074957506002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038316179055600161074d565b5060005b919050565b60003481106107d1576040805160e560020a62461bcd02815260206004820152603460248201527f466f72776172646572204465706f736974206661696c6564202d207a65726f2060448201527f6465706f73697473206e6f7420616c6c6f776564000000000000000000000000606482015290519081900360840190fd5b60025474010000000000000000000000000000000000000000900460ff161561086a576040805160e560020a62461bcd02815260206004820152602d60248201527f466f72776172646572204465706f736974206661696c6564202d206e6f20726560448201527f67697374657265642062616e6b00000000000000000000000000000000000000606482015290519081900360840190fd5b600154604080517ff340fa010000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163f340fa01913491602480830192602092919082900301818588803b1580156108d057600080fd5b505af11580156108e4573d6000803e3d6000fd5b50505050506040513d60208110156108fb57600080fd5b505115156001141561090f575060016103d0565b5060006103d05600a165627a7a7230582068838f4c20943277fbc4fd16acda60e5e3de3e81beda82ff9ca14df110abecaf0029`

// DeployJIincForwarder deploys a new Ethereum contract, binding an instance of JIincForwarder to it.
func DeployJIincForwarder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *JIincForwarder, error) {
	parsed, err := abi.JSON(strings.NewReader(JIincForwarderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(JIincForwarderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JIincForwarder{JIincForwarderCaller: JIincForwarderCaller{contract: contract}, JIincForwarderTransactor: JIincForwarderTransactor{contract: contract}, JIincForwarderFilterer: JIincForwarderFilterer{contract: contract}}, nil
}

// JIincForwarder is an auto generated Go binding around an Ethereum contract.
type JIincForwarder struct {
	JIincForwarderCaller     // Read-only binding to the contract
	JIincForwarderTransactor // Write-only binding to the contract
	JIincForwarderFilterer   // Log filterer for contract events
}

// JIincForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type JIincForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JIincForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JIincForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JIincForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JIincForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JIincForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JIincForwarderSession struct {
	Contract     *JIincForwarder   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JIincForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JIincForwarderCallerSession struct {
	Contract *JIincForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// JIincForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JIincForwarderTransactorSession struct {
	Contract     *JIincForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// JIincForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type JIincForwarderRaw struct {
	Contract *JIincForwarder // Generic contract binding to access the raw methods on
}

// JIincForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JIincForwarderCallerRaw struct {
	Contract *JIincForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// JIincForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JIincForwarderTransactorRaw struct {
	Contract *JIincForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJIincForwarder creates a new instance of JIincForwarder, bound to a specific deployed contract.
func NewJIincForwarder(address common.Address, backend bind.ContractBackend) (*JIincForwarder, error) {
	contract, err := bindJIincForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JIincForwarder{JIincForwarderCaller: JIincForwarderCaller{contract: contract}, JIincForwarderTransactor: JIincForwarderTransactor{contract: contract}, JIincForwarderFilterer: JIincForwarderFilterer{contract: contract}}, nil
}

// NewJIincForwarderCaller creates a new read-only instance of JIincForwarder, bound to a specific deployed contract.
func NewJIincForwarderCaller(address common.Address, caller bind.ContractCaller) (*JIincForwarderCaller, error) {
	contract, err := bindJIincForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JIincForwarderCaller{contract: contract}, nil
}

// NewJIincForwarderTransactor creates a new write-only instance of JIincForwarder, bound to a specific deployed contract.
func NewJIincForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*JIincForwarderTransactor, error) {
	contract, err := bindJIincForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JIincForwarderTransactor{contract: contract}, nil
}

// NewJIincForwarderFilterer creates a new log filterer instance of JIincForwarder, bound to a specific deployed contract.
func NewJIincForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*JIincForwarderFilterer, error) {
	contract, err := bindJIincForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JIincForwarderFilterer{contract: contract}, nil
}

// bindJIincForwarder binds a generic wrapper to an already deployed contract.
func bindJIincForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JIincForwarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JIincForwarder *JIincForwarderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JIincForwarder.Contract.JIincForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JIincForwarder *JIincForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JIincForwarder.Contract.JIincForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JIincForwarder *JIincForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JIincForwarder.Contract.JIincForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JIincForwarder *JIincForwarderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JIincForwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JIincForwarder *JIincForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JIincForwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JIincForwarder *JIincForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JIincForwarder.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_JIincForwarder *JIincForwarderCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _JIincForwarder.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_JIincForwarder *JIincForwarderSession) Name() (string, error) {
	return _JIincForwarder.Contract.Name(&_JIincForwarder.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_JIincForwarder *JIincForwarderCallerSession) Name() (string, error) {
	return _JIincForwarder.Contract.Name(&_JIincForwarder.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(address, address, bool)
func (_JIincForwarder *JIincForwarderCaller) Status(opts *bind.CallOpts) (common.Address, common.Address, bool, error) {
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
	err := _JIincForwarder.contract.Call(opts, out, "status")
	return *ret0, *ret1, *ret2, err
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(address, address, bool)
func (_JIincForwarder *JIincForwarderSession) Status() (common.Address, common.Address, bool, error) {
	return _JIincForwarder.Contract.Status(&_JIincForwarder.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(address, address, bool)
func (_JIincForwarder *JIincForwarderCallerSession) Status() (common.Address, common.Address, bool, error) {
	return _JIincForwarder.Contract.Status(&_JIincForwarder.CallOpts)
}

// CancelMigration is a paid mutator transaction binding the contract method 0x10639ea0.
//
// Solidity: function cancelMigration() returns(bool)
func (_JIincForwarder *JIincForwarderTransactor) CancelMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JIincForwarder.contract.Transact(opts, "cancelMigration")
}

// CancelMigration is a paid mutator transaction binding the contract method 0x10639ea0.
//
// Solidity: function cancelMigration() returns(bool)
func (_JIincForwarder *JIincForwarderSession) CancelMigration() (*types.Transaction, error) {
	return _JIincForwarder.Contract.CancelMigration(&_JIincForwarder.TransactOpts)
}

// CancelMigration is a paid mutator transaction binding the contract method 0x10639ea0.
//
// Solidity: function cancelMigration() returns(bool)
func (_JIincForwarder *JIincForwarderTransactorSession) CancelMigration() (*types.Transaction, error) {
	return _JIincForwarder.Contract.CancelMigration(&_JIincForwarder.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(bool)
func (_JIincForwarder *JIincForwarderTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JIincForwarder.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(bool)
func (_JIincForwarder *JIincForwarderSession) Deposit() (*types.Transaction, error) {
	return _JIincForwarder.Contract.Deposit(&_JIincForwarder.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(bool)
func (_JIincForwarder *JIincForwarderTransactorSession) Deposit() (*types.Transaction, error) {
	return _JIincForwarder.Contract.Deposit(&_JIincForwarder.TransactOpts)
}

// FinishMigration is a paid mutator transaction binding the contract method 0x88d761f2.
//
// Solidity: function finishMigration() returns(bool)
func (_JIincForwarder *JIincForwarderTransactor) FinishMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JIincForwarder.contract.Transact(opts, "finishMigration")
}

// FinishMigration is a paid mutator transaction binding the contract method 0x88d761f2.
//
// Solidity: function finishMigration() returns(bool)
func (_JIincForwarder *JIincForwarderSession) FinishMigration() (*types.Transaction, error) {
	return _JIincForwarder.Contract.FinishMigration(&_JIincForwarder.TransactOpts)
}

// FinishMigration is a paid mutator transaction binding the contract method 0x88d761f2.
//
// Solidity: function finishMigration() returns(bool)
func (_JIincForwarder *JIincForwarderTransactorSession) FinishMigration() (*types.Transaction, error) {
	return _JIincForwarder.Contract.FinishMigration(&_JIincForwarder.TransactOpts)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(_firstCorpBank address) returns()
func (_JIincForwarder *JIincForwarderTransactor) Setup(opts *bind.TransactOpts, _firstCorpBank common.Address) (*types.Transaction, error) {
	return _JIincForwarder.contract.Transact(opts, "setup", _firstCorpBank)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(_firstCorpBank address) returns()
func (_JIincForwarder *JIincForwarderSession) Setup(_firstCorpBank common.Address) (*types.Transaction, error) {
	return _JIincForwarder.Contract.Setup(&_JIincForwarder.TransactOpts, _firstCorpBank)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(_firstCorpBank address) returns()
func (_JIincForwarder *JIincForwarderTransactorSession) Setup(_firstCorpBank common.Address) (*types.Transaction, error) {
	return _JIincForwarder.Contract.Setup(&_JIincForwarder.TransactOpts, _firstCorpBank)
}

// StartMigration is a paid mutator transaction binding the contract method 0xa0f52da0.
//
// Solidity: function startMigration(_newCorpBank address) returns(bool)
func (_JIincForwarder *JIincForwarderTransactor) StartMigration(opts *bind.TransactOpts, _newCorpBank common.Address) (*types.Transaction, error) {
	return _JIincForwarder.contract.Transact(opts, "startMigration", _newCorpBank)
}

// StartMigration is a paid mutator transaction binding the contract method 0xa0f52da0.
//
// Solidity: function startMigration(_newCorpBank address) returns(bool)
func (_JIincForwarder *JIincForwarderSession) StartMigration(_newCorpBank common.Address) (*types.Transaction, error) {
	return _JIincForwarder.Contract.StartMigration(&_JIincForwarder.TransactOpts, _newCorpBank)
}

// StartMigration is a paid mutator transaction binding the contract method 0xa0f52da0.
//
// Solidity: function startMigration(_newCorpBank address) returns(bool)
func (_JIincForwarder *JIincForwarderTransactorSession) StartMigration(_newCorpBank common.Address) (*types.Transaction, error) {
	return _JIincForwarder.Contract.StartMigration(&_JIincForwarder.TransactOpts, _newCorpBank)
}

// JIincInterfaceForForwarderABI is the input ABI used to generate the binding from.
const JIincInterfaceForForwarderABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"migrationReceiver_setup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// JIincInterfaceForForwarderBin is the compiled bytecode used for deploying new contracts.
const JIincInterfaceForForwarderBin = `0x`

// DeployJIincInterfaceForForwarder deploys a new Ethereum contract, binding an instance of JIincInterfaceForForwarder to it.
func DeployJIincInterfaceForForwarder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *JIincInterfaceForForwarder, error) {
	parsed, err := abi.JSON(strings.NewReader(JIincInterfaceForForwarderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(JIincInterfaceForForwarderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JIincInterfaceForForwarder{JIincInterfaceForForwarderCaller: JIincInterfaceForForwarderCaller{contract: contract}, JIincInterfaceForForwarderTransactor: JIincInterfaceForForwarderTransactor{contract: contract}, JIincInterfaceForForwarderFilterer: JIincInterfaceForForwarderFilterer{contract: contract}}, nil
}

// JIincInterfaceForForwarder is an auto generated Go binding around an Ethereum contract.
type JIincInterfaceForForwarder struct {
	JIincInterfaceForForwarderCaller     // Read-only binding to the contract
	JIincInterfaceForForwarderTransactor // Write-only binding to the contract
	JIincInterfaceForForwarderFilterer   // Log filterer for contract events
}

// JIincInterfaceForForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type JIincInterfaceForForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JIincInterfaceForForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JIincInterfaceForForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JIincInterfaceForForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JIincInterfaceForForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JIincInterfaceForForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JIincInterfaceForForwarderSession struct {
	Contract     *JIincInterfaceForForwarder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// JIincInterfaceForForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JIincInterfaceForForwarderCallerSession struct {
	Contract *JIincInterfaceForForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// JIincInterfaceForForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JIincInterfaceForForwarderTransactorSession struct {
	Contract     *JIincInterfaceForForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// JIincInterfaceForForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type JIincInterfaceForForwarderRaw struct {
	Contract *JIincInterfaceForForwarder // Generic contract binding to access the raw methods on
}

// JIincInterfaceForForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JIincInterfaceForForwarderCallerRaw struct {
	Contract *JIincInterfaceForForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// JIincInterfaceForForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JIincInterfaceForForwarderTransactorRaw struct {
	Contract *JIincInterfaceForForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJIincInterfaceForForwarder creates a new instance of JIincInterfaceForForwarder, bound to a specific deployed contract.
func NewJIincInterfaceForForwarder(address common.Address, backend bind.ContractBackend) (*JIincInterfaceForForwarder, error) {
	contract, err := bindJIincInterfaceForForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JIincInterfaceForForwarder{JIincInterfaceForForwarderCaller: JIincInterfaceForForwarderCaller{contract: contract}, JIincInterfaceForForwarderTransactor: JIincInterfaceForForwarderTransactor{contract: contract}, JIincInterfaceForForwarderFilterer: JIincInterfaceForForwarderFilterer{contract: contract}}, nil
}

// NewJIincInterfaceForForwarderCaller creates a new read-only instance of JIincInterfaceForForwarder, bound to a specific deployed contract.
func NewJIincInterfaceForForwarderCaller(address common.Address, caller bind.ContractCaller) (*JIincInterfaceForForwarderCaller, error) {
	contract, err := bindJIincInterfaceForForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JIincInterfaceForForwarderCaller{contract: contract}, nil
}

// NewJIincInterfaceForForwarderTransactor creates a new write-only instance of JIincInterfaceForForwarder, bound to a specific deployed contract.
func NewJIincInterfaceForForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*JIincInterfaceForForwarderTransactor, error) {
	contract, err := bindJIincInterfaceForForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JIincInterfaceForForwarderTransactor{contract: contract}, nil
}

// NewJIincInterfaceForForwarderFilterer creates a new log filterer instance of JIincInterfaceForForwarder, bound to a specific deployed contract.
func NewJIincInterfaceForForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*JIincInterfaceForForwarderFilterer, error) {
	contract, err := bindJIincInterfaceForForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JIincInterfaceForForwarderFilterer{contract: contract}, nil
}

// bindJIincInterfaceForForwarder binds a generic wrapper to an already deployed contract.
func bindJIincInterfaceForForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JIincInterfaceForForwarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JIincInterfaceForForwarder *JIincInterfaceForForwarderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JIincInterfaceForForwarder.Contract.JIincInterfaceForForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JIincInterfaceForForwarder *JIincInterfaceForForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JIincInterfaceForForwarder.Contract.JIincInterfaceForForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JIincInterfaceForForwarder *JIincInterfaceForForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JIincInterfaceForForwarder.Contract.JIincInterfaceForForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JIincInterfaceForForwarder *JIincInterfaceForForwarderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JIincInterfaceForForwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JIincInterfaceForForwarder *JIincInterfaceForForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JIincInterfaceForForwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JIincInterfaceForForwarder *JIincInterfaceForForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JIincInterfaceForForwarder.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(_addr address) returns(bool)
func (_JIincInterfaceForForwarder *JIincInterfaceForForwarderTransactor) Deposit(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _JIincInterfaceForForwarder.contract.Transact(opts, "deposit", _addr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(_addr address) returns(bool)
func (_JIincInterfaceForForwarder *JIincInterfaceForForwarderSession) Deposit(_addr common.Address) (*types.Transaction, error) {
	return _JIincInterfaceForForwarder.Contract.Deposit(&_JIincInterfaceForForwarder.TransactOpts, _addr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(_addr address) returns(bool)
func (_JIincInterfaceForForwarder *JIincInterfaceForForwarderTransactorSession) Deposit(_addr common.Address) (*types.Transaction, error) {
	return _JIincInterfaceForForwarder.Contract.Deposit(&_JIincInterfaceForForwarder.TransactOpts, _addr)
}

// MigrationReceiverSetup is a paid mutator transaction binding the contract method 0x0839e0fb.
//
// Solidity: function migrationReceiver_setup() returns(bool)
func (_JIincInterfaceForForwarder *JIincInterfaceForForwarderTransactor) MigrationReceiverSetup(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JIincInterfaceForForwarder.contract.Transact(opts, "migrationReceiver_setup")
}

// MigrationReceiverSetup is a paid mutator transaction binding the contract method 0x0839e0fb.
//
// Solidity: function migrationReceiver_setup() returns(bool)
func (_JIincInterfaceForForwarder *JIincInterfaceForForwarderSession) MigrationReceiverSetup() (*types.Transaction, error) {
	return _JIincInterfaceForForwarder.Contract.MigrationReceiverSetup(&_JIincInterfaceForForwarder.TransactOpts)
}

// MigrationReceiverSetup is a paid mutator transaction binding the contract method 0x0839e0fb.
//
// Solidity: function migrationReceiver_setup() returns(bool)
func (_JIincInterfaceForForwarder *JIincInterfaceForForwarderTransactorSession) MigrationReceiverSetup() (*types.Transaction, error) {
	return _JIincInterfaceForForwarder.Contract.MigrationReceiverSetup(&_JIincInterfaceForForwarder.TransactOpts)
}
