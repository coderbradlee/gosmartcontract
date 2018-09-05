// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Divies

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

// DiviesABI is the input ABI used to generate the binding from.
const DiviesABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"pushers_\",\"outputs\":[{\"name\":\"tracker\",\"type\":\"uint256\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pusherTracker_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_percent\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rateLimiter_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pusher\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"startingBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"masternodePayout\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"finalBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"}],\"name\":\"onDistribute\",\"type\":\"event\"}]"

// DiviesBin is the compiled bytecode used for deploying new contracts.
const DiviesBin = `0x6080604052606460005534801561001557600080fd5b50610c8d806100256000396000f30060806040526004361061005e5763ffffffff60e060020a6000350416631f75c3908114610060578063521ceba71461009a5780637bb98a68146100c157806391c05b0b146100d6578063b519cf31146100ee578063d0e30db014610103575b005b34801561006c57600080fd5b50610081600160a060020a036004351661010b565b6040805192835260208301919091528051918290030190f35b3480156100a657600080fd5b506100af610124565b60408051918252519081900360200190f35b3480156100cd57600080fd5b506100af61012a565b3480156100e257600080fd5b5061005e60043561012f565b3480156100fa57600080fd5b506100af61094b565b61005e610951565b6001602081905260009182526040909120805491015482565b60005481565b303190565b6000808080808033803b801561018f576040805160e560020a62461bcd02815260206004820152601160248201527f736f7272792068756d616e73206f6e6c79000000000000000000000000000000604482015290519081900360640190fd5b60008911801561019f5750606489105b151561021b576040805160e560020a62461bcd02815260206004820152602660248201527f706c65617365207069636b20612070657263656e74206265747765656e20312060448201527f616e642039390000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000543398503031975061023690606463ffffffff61095316565b600160a060020a0389166000908152600160205260409020541180159061028b5750600160a060020a03881660009081526001602081905260409091200154429061028990610e1063ffffffff6109b816565b105b156108385760008054600160a060020a038a16825260016020818152604080852084905591909201835580517f56d399e8000000000000000000000000000000000000000000000000000000008152905173b3775fb83f7d12a36e0475abdd1fca35c091efbe936356d399e893600480850194919392918390030190829087803b15801561031857600080fd5b505af115801561032c573d6000803e3d6000fd5b505050506040513d602081101561034257600080fd5b5051604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038b166004820152905173b3775fb83f7d12a36e0475abdd1fca35c091efbe916370a082319160248083019260209291908290030181600087803b1580156103b957600080fd5b505af11580156103cd573d6000803e3d6000fd5b505050506040513d60208110156103e357600080fd5b5051106103f4576003600a88040495505b6064610408888b830363ffffffff610a1316565b81151561041157fe5b04935073b3775fb83f7d12a36e0475abdd1fca35c091efbe600160a060020a031663f088d547888a6040518363ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a031681526020019150506020604051808303818588803b15801561048357600080fd5b505af1158015610497573d6000803e3d6000fd5b50505050506040513d60208110156104ae57600080fd5b5050604080517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152905173b3775fb83f7d12a36e0475abdd1fca35c091efbe9163e4849b329183916370a082319160248083019260209291908290030181600087803b15801561052457600080fd5b505af1158015610538573d6000803e3d6000fd5b505050506040513d602081101561054e57600080fd5b50516040805160e060020a63ffffffff8516028152600481019290925251602480830192600092919082900301818387803b15801561058c57600080fd5b505af11580156105a0573d6000803e3d6000fd5b5050604080517e65318b000000000000000000000000000000000000000000000000000000008152306004820152905173b3775fb83f7d12a36e0475abdd1fca35c091efbe93506265318b925060248083019260209291908290030181600087803b15801561060e57600080fd5b505af1158015610622573d6000803e3d6000fd5b505050506040513d602081101561063857600080fd5b505192505b8383106107c95773b3775fb83f7d12a36e0475abdd1fca35c091efbe600160a060020a031663fdb5a03e6040518163ffffffff1660e060020a028152600401600060405180830381600087803b15801561069657600080fd5b505af11580156106aa573d6000803e3d6000fd5b5050604080517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152905173b3775fb83f7d12a36e0475abdd1fca35c091efbe935063e4849b32925083916370a082319160248083019260209291908290030181600087803b15801561072257600080fd5b505af1158015610736573d6000803e3d6000fd5b505050506040513d602081101561074c57600080fd5b50516040805160e060020a63ffffffff8516028152600481019290925251602480830192600092919082900301818387803b15801561078a57600080fd5b505af115801561079e573d6000803e3d6000fd5b5050505060646107b8605185610a1390919063ffffffff16565b8115156107c157fe5b04925061063d565b73b3775fb83f7d12a36e0475abdd1fca35c091efbe600160a060020a0316633ccfd60b6040518163ffffffff1660e060020a028152600401600060405180830381600087803b15801561081b57600080fd5b505af115801561082f573d6000803e3d6000fd5b5050505061084f565b61084c856001602f8063ffffffff610a8a16565b94505b600160a060020a0388166000908152600160208190526040822042910181905561088391879190600e63ffffffff610a8a16565b600160a060020a0389166000908152600160205260409020549095506108b4908690600f601d63ffffffff610a8a16565b6000549095506108cf908690601e602c63ffffffff610a8a16565b94506108e5858a602d602e63ffffffff610a8a16565b60408051600160a060020a038b168152602081018a9052808201899052303160608201526080810183905290519196507fa0e5739b8aab1e9b0469db4a982bc2211bdf4b0ed8b8aebe2e0d89d10057657d919081900360a00190a1505050505050505050565b60025481565b565b6000828211156109ad576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d61746820737562206661696c656400000000000000000000000000604482015290519081900360640190fd5b508082035b92915050565b818101828110156109b2576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d61746820616464206661696c656400000000000000000000000000604482015290519081900360640190fd5b6000821515610a24575060006109b2565b50818102818382811515610a3457fe5b04146109b2576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d617468206d756c206661696c656400000000000000000000000000604482015290519081900360640190fd5b6000604d82108015610a9c5750604d83105b1515610af2576040805160e560020a62461bcd02815260206004820152601e60248201527f73746172742f656e64206d757374206265206c657373207468616e2037370000604482015290519081900360640190fd5b82821015610b4a576040805160e560020a62461bcd02815260206004820152601460248201527f656e64206d757374206265203e3d207374617274000000000000000000000000604482015290519081900360640190fd5b610b64600a610b5884610bfb565b9063ffffffff610a1316565b9150610b6f83610bfb565b92508282811515610b7c57fe5b048410610b8857600080fd5b6000841115610ba457610ba1848463ffffffff610a1316565b93505b610bf2610bc3838488811515610bb657fe5b049063ffffffff610a1316565b610be686610be6610bd988898c811515610bb657fe5b8a9063ffffffff61095316565b9063ffffffff6109b816565b95945050505050565b60006109b2600a8363ffffffff610c0e16565b60008080841515610c225760009250610c59565b831515610c325760019250610c59565b5083905060015b83811015610c5557610c4b8286610a13565b9150600101610c39565b8192505b5050929150505600a165627a7a723058203e860075bb397063ab17030f460020c7619f2afceed020fdf9bdc12378193f7b0029`

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

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(uint256)
func (_Divies *DiviesCaller) Balances(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Divies.contract.Call(opts, out, "balances")
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(uint256)
func (_Divies *DiviesSession) Balances() (*big.Int, error) {
	return _Divies.Contract.Balances(&_Divies.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(uint256)
func (_Divies *DiviesCallerSession) Balances() (*big.Int, error) {
	return _Divies.Contract.Balances(&_Divies.CallOpts)
}

// PusherTracker is a free data retrieval call binding the contract method 0x521ceba7.
//
// Solidity: function pusherTracker_() constant returns(uint256)
func (_Divies *DiviesCaller) PusherTracker(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Divies.contract.Call(opts, out, "pusherTracker_")
	return *ret0, err
}

// PusherTracker is a free data retrieval call binding the contract method 0x521ceba7.
//
// Solidity: function pusherTracker_() constant returns(uint256)
func (_Divies *DiviesSession) PusherTracker() (*big.Int, error) {
	return _Divies.Contract.PusherTracker(&_Divies.CallOpts)
}

// PusherTracker is a free data retrieval call binding the contract method 0x521ceba7.
//
// Solidity: function pusherTracker_() constant returns(uint256)
func (_Divies *DiviesCallerSession) PusherTracker() (*big.Int, error) {
	return _Divies.Contract.PusherTracker(&_Divies.CallOpts)
}

// Pushers is a free data retrieval call binding the contract method 0x1f75c390.
//
// Solidity: function pushers_( address) constant returns(tracker uint256, time uint256)
func (_Divies *DiviesCaller) Pushers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Tracker *big.Int
	Time    *big.Int
}, error) {
	ret := new(struct {
		Tracker *big.Int
		Time    *big.Int
	})
	out := ret
	err := _Divies.contract.Call(opts, out, "pushers_", arg0)
	return *ret, err
}

// Pushers is a free data retrieval call binding the contract method 0x1f75c390.
//
// Solidity: function pushers_( address) constant returns(tracker uint256, time uint256)
func (_Divies *DiviesSession) Pushers(arg0 common.Address) (struct {
	Tracker *big.Int
	Time    *big.Int
}, error) {
	return _Divies.Contract.Pushers(&_Divies.CallOpts, arg0)
}

// Pushers is a free data retrieval call binding the contract method 0x1f75c390.
//
// Solidity: function pushers_( address) constant returns(tracker uint256, time uint256)
func (_Divies *DiviesCallerSession) Pushers(arg0 common.Address) (struct {
	Tracker *big.Int
	Time    *big.Int
}, error) {
	return _Divies.Contract.Pushers(&_Divies.CallOpts, arg0)
}

// RateLimiter is a free data retrieval call binding the contract method 0xb519cf31.
//
// Solidity: function rateLimiter_() constant returns(uint256)
func (_Divies *DiviesCaller) RateLimiter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Divies.contract.Call(opts, out, "rateLimiter_")
	return *ret0, err
}

// RateLimiter is a free data retrieval call binding the contract method 0xb519cf31.
//
// Solidity: function rateLimiter_() constant returns(uint256)
func (_Divies *DiviesSession) RateLimiter() (*big.Int, error) {
	return _Divies.Contract.RateLimiter(&_Divies.CallOpts)
}

// RateLimiter is a free data retrieval call binding the contract method 0xb519cf31.
//
// Solidity: function rateLimiter_() constant returns(uint256)
func (_Divies *DiviesCallerSession) RateLimiter() (*big.Int, error) {
	return _Divies.Contract.RateLimiter(&_Divies.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Divies *DiviesTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Divies.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Divies *DiviesSession) Deposit() (*types.Transaction, error) {
	return _Divies.Contract.Deposit(&_Divies.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Divies *DiviesTransactorSession) Deposit() (*types.Transaction, error) {
	return _Divies.Contract.Deposit(&_Divies.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0x91c05b0b.
//
// Solidity: function distribute(_percent uint256) returns()
func (_Divies *DiviesTransactor) Distribute(opts *bind.TransactOpts, _percent *big.Int) (*types.Transaction, error) {
	return _Divies.contract.Transact(opts, "distribute", _percent)
}

// Distribute is a paid mutator transaction binding the contract method 0x91c05b0b.
//
// Solidity: function distribute(_percent uint256) returns()
func (_Divies *DiviesSession) Distribute(_percent *big.Int) (*types.Transaction, error) {
	return _Divies.Contract.Distribute(&_Divies.TransactOpts, _percent)
}

// Distribute is a paid mutator transaction binding the contract method 0x91c05b0b.
//
// Solidity: function distribute(_percent uint256) returns()
func (_Divies *DiviesTransactorSession) Distribute(_percent *big.Int) (*types.Transaction, error) {
	return _Divies.Contract.Distribute(&_Divies.TransactOpts, _percent)
}

// DiviesOnDistributeIterator is returned from FilterOnDistribute and is used to iterate over the raw logs and unpacked data for OnDistribute events raised by the Divies contract.
type DiviesOnDistributeIterator struct {
	Event *DiviesOnDistribute // Event containing the contract specifics and raw log

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
func (it *DiviesOnDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiviesOnDistribute)
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
		it.Event = new(DiviesOnDistribute)
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
func (it *DiviesOnDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiviesOnDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiviesOnDistribute represents a OnDistribute event raised by the Divies contract.
type DiviesOnDistribute struct {
	Pusher           common.Address
	StartingBalance  *big.Int
	MasternodePayout *big.Int
	FinalBalance     *big.Int
	CompressedData   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnDistribute is a free log retrieval operation binding the contract event 0xa0e5739b8aab1e9b0469db4a982bc2211bdf4b0ed8b8aebe2e0d89d10057657d.
//
// Solidity: e onDistribute(pusher address, startingBalance uint256, masternodePayout uint256, finalBalance uint256, compressedData uint256)
func (_Divies *DiviesFilterer) FilterOnDistribute(opts *bind.FilterOpts) (*DiviesOnDistributeIterator, error) {

	logs, sub, err := _Divies.contract.FilterLogs(opts, "onDistribute")
	if err != nil {
		return nil, err
	}
	return &DiviesOnDistributeIterator{contract: _Divies.contract, event: "onDistribute", logs: logs, sub: sub}, nil
}

// WatchOnDistribute is a free log subscription operation binding the contract event 0xa0e5739b8aab1e9b0469db4a982bc2211bdf4b0ed8b8aebe2e0d89d10057657d.
//
// Solidity: e onDistribute(pusher address, startingBalance uint256, masternodePayout uint256, finalBalance uint256, compressedData uint256)
func (_Divies *DiviesFilterer) WatchOnDistribute(opts *bind.WatchOpts, sink chan<- *DiviesOnDistribute) (event.Subscription, error) {

	logs, sub, err := _Divies.contract.WatchLogs(opts, "onDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiviesOnDistribute)
				if err := _Divies.contract.UnpackLog(event, "onDistribute", log); err != nil {
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

// HourglassInterfaceABI is the input ABI used to generate the binding from.
const HourglassInterfaceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_playerAddress\",\"type\":\"address\"}],\"name\":\"dividendsOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingRequirement\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_playerAddress\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_toAddress\",\"type\":\"address\"},{\"name\":\"_amountOfTokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amountOfTokens\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_playerAddress\",\"type\":\"address\"}],\"name\":\"buy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reinvest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// HourglassInterfaceBin is the compiled bytecode used for deploying new contracts.
const HourglassInterfaceBin = `0x`

// DeployHourglassInterface deploys a new Ethereum contract, binding an instance of HourglassInterface to it.
func DeployHourglassInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HourglassInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(HourglassInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HourglassInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HourglassInterface{HourglassInterfaceCaller: HourglassInterfaceCaller{contract: contract}, HourglassInterfaceTransactor: HourglassInterfaceTransactor{contract: contract}, HourglassInterfaceFilterer: HourglassInterfaceFilterer{contract: contract}}, nil
}

// HourglassInterface is an auto generated Go binding around an Ethereum contract.
type HourglassInterface struct {
	HourglassInterfaceCaller     // Read-only binding to the contract
	HourglassInterfaceTransactor // Write-only binding to the contract
	HourglassInterfaceFilterer   // Log filterer for contract events
}

// HourglassInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type HourglassInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HourglassInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HourglassInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HourglassInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HourglassInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HourglassInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HourglassInterfaceSession struct {
	Contract     *HourglassInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HourglassInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HourglassInterfaceCallerSession struct {
	Contract *HourglassInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// HourglassInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HourglassInterfaceTransactorSession struct {
	Contract     *HourglassInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// HourglassInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type HourglassInterfaceRaw struct {
	Contract *HourglassInterface // Generic contract binding to access the raw methods on
}

// HourglassInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HourglassInterfaceCallerRaw struct {
	Contract *HourglassInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// HourglassInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HourglassInterfaceTransactorRaw struct {
	Contract *HourglassInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHourglassInterface creates a new instance of HourglassInterface, bound to a specific deployed contract.
func NewHourglassInterface(address common.Address, backend bind.ContractBackend) (*HourglassInterface, error) {
	contract, err := bindHourglassInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HourglassInterface{HourglassInterfaceCaller: HourglassInterfaceCaller{contract: contract}, HourglassInterfaceTransactor: HourglassInterfaceTransactor{contract: contract}, HourglassInterfaceFilterer: HourglassInterfaceFilterer{contract: contract}}, nil
}

// NewHourglassInterfaceCaller creates a new read-only instance of HourglassInterface, bound to a specific deployed contract.
func NewHourglassInterfaceCaller(address common.Address, caller bind.ContractCaller) (*HourglassInterfaceCaller, error) {
	contract, err := bindHourglassInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HourglassInterfaceCaller{contract: contract}, nil
}

// NewHourglassInterfaceTransactor creates a new write-only instance of HourglassInterface, bound to a specific deployed contract.
func NewHourglassInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*HourglassInterfaceTransactor, error) {
	contract, err := bindHourglassInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HourglassInterfaceTransactor{contract: contract}, nil
}

// NewHourglassInterfaceFilterer creates a new log filterer instance of HourglassInterface, bound to a specific deployed contract.
func NewHourglassInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*HourglassInterfaceFilterer, error) {
	contract, err := bindHourglassInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HourglassInterfaceFilterer{contract: contract}, nil
}

// bindHourglassInterface binds a generic wrapper to an already deployed contract.
func bindHourglassInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HourglassInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HourglassInterface *HourglassInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HourglassInterface.Contract.HourglassInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HourglassInterface *HourglassInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HourglassInterface.Contract.HourglassInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HourglassInterface *HourglassInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HourglassInterface.Contract.HourglassInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HourglassInterface *HourglassInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HourglassInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HourglassInterface *HourglassInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HourglassInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HourglassInterface *HourglassInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HourglassInterface.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_playerAddress address) constant returns(uint256)
func (_HourglassInterface *HourglassInterfaceCaller) BalanceOf(opts *bind.CallOpts, _playerAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HourglassInterface.contract.Call(opts, out, "balanceOf", _playerAddress)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_playerAddress address) constant returns(uint256)
func (_HourglassInterface *HourglassInterfaceSession) BalanceOf(_playerAddress common.Address) (*big.Int, error) {
	return _HourglassInterface.Contract.BalanceOf(&_HourglassInterface.CallOpts, _playerAddress)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_playerAddress address) constant returns(uint256)
func (_HourglassInterface *HourglassInterfaceCallerSession) BalanceOf(_playerAddress common.Address) (*big.Int, error) {
	return _HourglassInterface.Contract.BalanceOf(&_HourglassInterface.CallOpts, _playerAddress)
}

// DividendsOf is a free data retrieval call binding the contract method 0x0065318b.
//
// Solidity: function dividendsOf(_playerAddress address) constant returns(uint256)
func (_HourglassInterface *HourglassInterfaceCaller) DividendsOf(opts *bind.CallOpts, _playerAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HourglassInterface.contract.Call(opts, out, "dividendsOf", _playerAddress)
	return *ret0, err
}

// DividendsOf is a free data retrieval call binding the contract method 0x0065318b.
//
// Solidity: function dividendsOf(_playerAddress address) constant returns(uint256)
func (_HourglassInterface *HourglassInterfaceSession) DividendsOf(_playerAddress common.Address) (*big.Int, error) {
	return _HourglassInterface.Contract.DividendsOf(&_HourglassInterface.CallOpts, _playerAddress)
}

// DividendsOf is a free data retrieval call binding the contract method 0x0065318b.
//
// Solidity: function dividendsOf(_playerAddress address) constant returns(uint256)
func (_HourglassInterface *HourglassInterfaceCallerSession) DividendsOf(_playerAddress common.Address) (*big.Int, error) {
	return _HourglassInterface.Contract.DividendsOf(&_HourglassInterface.CallOpts, _playerAddress)
}

// StakingRequirement is a free data retrieval call binding the contract method 0x56d399e8.
//
// Solidity: function stakingRequirement() constant returns(uint256)
func (_HourglassInterface *HourglassInterfaceCaller) StakingRequirement(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HourglassInterface.contract.Call(opts, out, "stakingRequirement")
	return *ret0, err
}

// StakingRequirement is a free data retrieval call binding the contract method 0x56d399e8.
//
// Solidity: function stakingRequirement() constant returns(uint256)
func (_HourglassInterface *HourglassInterfaceSession) StakingRequirement() (*big.Int, error) {
	return _HourglassInterface.Contract.StakingRequirement(&_HourglassInterface.CallOpts)
}

// StakingRequirement is a free data retrieval call binding the contract method 0x56d399e8.
//
// Solidity: function stakingRequirement() constant returns(uint256)
func (_HourglassInterface *HourglassInterfaceCallerSession) StakingRequirement() (*big.Int, error) {
	return _HourglassInterface.Contract.StakingRequirement(&_HourglassInterface.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(_playerAddress address) returns(uint256)
func (_HourglassInterface *HourglassInterfaceTransactor) Buy(opts *bind.TransactOpts, _playerAddress common.Address) (*types.Transaction, error) {
	return _HourglassInterface.contract.Transact(opts, "buy", _playerAddress)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(_playerAddress address) returns(uint256)
func (_HourglassInterface *HourglassInterfaceSession) Buy(_playerAddress common.Address) (*types.Transaction, error) {
	return _HourglassInterface.Contract.Buy(&_HourglassInterface.TransactOpts, _playerAddress)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(_playerAddress address) returns(uint256)
func (_HourglassInterface *HourglassInterfaceTransactorSession) Buy(_playerAddress common.Address) (*types.Transaction, error) {
	return _HourglassInterface.Contract.Buy(&_HourglassInterface.TransactOpts, _playerAddress)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_HourglassInterface *HourglassInterfaceTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HourglassInterface.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_HourglassInterface *HourglassInterfaceSession) Exit() (*types.Transaction, error) {
	return _HourglassInterface.Contract.Exit(&_HourglassInterface.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_HourglassInterface *HourglassInterfaceTransactorSession) Exit() (*types.Transaction, error) {
	return _HourglassInterface.Contract.Exit(&_HourglassInterface.TransactOpts)
}

// Reinvest is a paid mutator transaction binding the contract method 0xfdb5a03e.
//
// Solidity: function reinvest() returns()
func (_HourglassInterface *HourglassInterfaceTransactor) Reinvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HourglassInterface.contract.Transact(opts, "reinvest")
}

// Reinvest is a paid mutator transaction binding the contract method 0xfdb5a03e.
//
// Solidity: function reinvest() returns()
func (_HourglassInterface *HourglassInterfaceSession) Reinvest() (*types.Transaction, error) {
	return _HourglassInterface.Contract.Reinvest(&_HourglassInterface.TransactOpts)
}

// Reinvest is a paid mutator transaction binding the contract method 0xfdb5a03e.
//
// Solidity: function reinvest() returns()
func (_HourglassInterface *HourglassInterfaceTransactorSession) Reinvest() (*types.Transaction, error) {
	return _HourglassInterface.Contract.Reinvest(&_HourglassInterface.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(_amountOfTokens uint256) returns()
func (_HourglassInterface *HourglassInterfaceTransactor) Sell(opts *bind.TransactOpts, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _HourglassInterface.contract.Transact(opts, "sell", _amountOfTokens)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(_amountOfTokens uint256) returns()
func (_HourglassInterface *HourglassInterfaceSession) Sell(_amountOfTokens *big.Int) (*types.Transaction, error) {
	return _HourglassInterface.Contract.Sell(&_HourglassInterface.TransactOpts, _amountOfTokens)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(_amountOfTokens uint256) returns()
func (_HourglassInterface *HourglassInterfaceTransactorSession) Sell(_amountOfTokens *big.Int) (*types.Transaction, error) {
	return _HourglassInterface.Contract.Sell(&_HourglassInterface.TransactOpts, _amountOfTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_toAddress address, _amountOfTokens uint256) returns(bool)
func (_HourglassInterface *HourglassInterfaceTransactor) Transfer(opts *bind.TransactOpts, _toAddress common.Address, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _HourglassInterface.contract.Transact(opts, "transfer", _toAddress, _amountOfTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_toAddress address, _amountOfTokens uint256) returns(bool)
func (_HourglassInterface *HourglassInterfaceSession) Transfer(_toAddress common.Address, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _HourglassInterface.Contract.Transfer(&_HourglassInterface.TransactOpts, _toAddress, _amountOfTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_toAddress address, _amountOfTokens uint256) returns(bool)
func (_HourglassInterface *HourglassInterfaceTransactorSession) Transfer(_toAddress common.Address, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _HourglassInterface.Contract.Transfer(&_HourglassInterface.TransactOpts, _toAddress, _amountOfTokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_HourglassInterface *HourglassInterfaceTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HourglassInterface.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_HourglassInterface *HourglassInterfaceSession) Withdraw() (*types.Transaction, error) {
	return _HourglassInterface.Contract.Withdraw(&_HourglassInterface.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_HourglassInterface *HourglassInterfaceTransactorSession) Withdraw() (*types.Transaction, error) {
	return _HourglassInterface.Contract.Withdraw(&_HourglassInterface.TransactOpts)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820b2b4268532fa3d2414dbbca1f66036082f14f242e83e36e1d8c3014c2b1d6fe40029`

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

// UintCompressorABI is the input ABI used to generate the binding from.
const UintCompressorABI = "[]"

// UintCompressorBin is the compiled bytecode used for deploying new contracts.
const UintCompressorBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058204e1d089931ebe659f94a6baf3b11f6945f4b3c9e0de04346ff9d1a3f6574b5550029`

// DeployUintCompressor deploys a new Ethereum contract, binding an instance of UintCompressor to it.
func DeployUintCompressor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UintCompressor, error) {
	parsed, err := abi.JSON(strings.NewReader(UintCompressorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UintCompressorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UintCompressor{UintCompressorCaller: UintCompressorCaller{contract: contract}, UintCompressorTransactor: UintCompressorTransactor{contract: contract}, UintCompressorFilterer: UintCompressorFilterer{contract: contract}}, nil
}

// UintCompressor is an auto generated Go binding around an Ethereum contract.
type UintCompressor struct {
	UintCompressorCaller     // Read-only binding to the contract
	UintCompressorTransactor // Write-only binding to the contract
	UintCompressorFilterer   // Log filterer for contract events
}

// UintCompressorCaller is an auto generated read-only Go binding around an Ethereum contract.
type UintCompressorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UintCompressorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UintCompressorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UintCompressorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UintCompressorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UintCompressorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UintCompressorSession struct {
	Contract     *UintCompressor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UintCompressorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UintCompressorCallerSession struct {
	Contract *UintCompressorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UintCompressorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UintCompressorTransactorSession struct {
	Contract     *UintCompressorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UintCompressorRaw is an auto generated low-level Go binding around an Ethereum contract.
type UintCompressorRaw struct {
	Contract *UintCompressor // Generic contract binding to access the raw methods on
}

// UintCompressorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UintCompressorCallerRaw struct {
	Contract *UintCompressorCaller // Generic read-only contract binding to access the raw methods on
}

// UintCompressorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UintCompressorTransactorRaw struct {
	Contract *UintCompressorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUintCompressor creates a new instance of UintCompressor, bound to a specific deployed contract.
func NewUintCompressor(address common.Address, backend bind.ContractBackend) (*UintCompressor, error) {
	contract, err := bindUintCompressor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UintCompressor{UintCompressorCaller: UintCompressorCaller{contract: contract}, UintCompressorTransactor: UintCompressorTransactor{contract: contract}, UintCompressorFilterer: UintCompressorFilterer{contract: contract}}, nil
}

// NewUintCompressorCaller creates a new read-only instance of UintCompressor, bound to a specific deployed contract.
func NewUintCompressorCaller(address common.Address, caller bind.ContractCaller) (*UintCompressorCaller, error) {
	contract, err := bindUintCompressor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UintCompressorCaller{contract: contract}, nil
}

// NewUintCompressorTransactor creates a new write-only instance of UintCompressor, bound to a specific deployed contract.
func NewUintCompressorTransactor(address common.Address, transactor bind.ContractTransactor) (*UintCompressorTransactor, error) {
	contract, err := bindUintCompressor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UintCompressorTransactor{contract: contract}, nil
}

// NewUintCompressorFilterer creates a new log filterer instance of UintCompressor, bound to a specific deployed contract.
func NewUintCompressorFilterer(address common.Address, filterer bind.ContractFilterer) (*UintCompressorFilterer, error) {
	contract, err := bindUintCompressor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UintCompressorFilterer{contract: contract}, nil
}

// bindUintCompressor binds a generic wrapper to an already deployed contract.
func bindUintCompressor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UintCompressorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UintCompressor *UintCompressorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UintCompressor.Contract.UintCompressorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UintCompressor *UintCompressorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UintCompressor.Contract.UintCompressorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UintCompressor *UintCompressorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UintCompressor.Contract.UintCompressorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UintCompressor *UintCompressorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UintCompressor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UintCompressor *UintCompressorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UintCompressor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UintCompressor *UintCompressorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UintCompressor.Contract.contract.Transact(opts, method, params...)
}
