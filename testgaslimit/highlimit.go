// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package highlimit

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

// HighlimitABI is the input ABI used to generate the binding from.
const HighlimitABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"del\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"run\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"a1\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"a2\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"a3\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"a4\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"a5\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"a6\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"a17\",\"type\":\"uint256\"}],\"name\":\"onDel\",\"type\":\"event\"}]"

// HighlimitBin is the compiled bytecode used for deploying new contracts.
const HighlimitBin = `0x608060405234801561001057600080fd5b50610261806100206000396000f30060806040526004361061004b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663b6588ffd8114610050578063c04062261461009d575b600080fd5b34801561005c57600080fd5b506100656100b4565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b3480156100a957600080fd5b506100b2610233565b005b60008080808080808080805b600a8210156100ee5782546001818101855560008581526020902060028501920191909155909101906100c0565b50508054600090600419015b8082101561012657828281548110151561011057fe5b60009182526020822001556001909101906100fa565b7f52f847c6da443e437b0d57a554be808022a7e69637c9ed63a1e16d4b6341341183600481548110151561015657fe5b906000526020600020015484600581548110151561017057fe5b906000526020600020015485600681548110151561018a57fe5b90600052602060002001548660078154811015156101a457fe5b90600052602060002001548760088154811015156101be57fe5b90600052602060002001548860098154811015156101d857fe5b906000526020600020015489805490506040518088815260200187815260200186815260200185815260200184815260200183815260200182815260200197505050505050505060405180910390a150505090919293949596565bfe00a165627a7a7230582088e9a16da9f9ca6ad6ce4ba678afb7eca1b6266dbb9117b3820f8caaaa443e940029`

// DeployHighlimit deploys a new Ethereum contract, binding an instance of Highlimit to it.
func DeployHighlimit(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Highlimit, error) {
	parsed, err := abi.JSON(strings.NewReader(HighlimitABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HighlimitBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Highlimit{HighlimitCaller: HighlimitCaller{contract: contract}, HighlimitTransactor: HighlimitTransactor{contract: contract}, HighlimitFilterer: HighlimitFilterer{contract: contract}}, nil
}

// Highlimit is an auto generated Go binding around an Ethereum contract.
type Highlimit struct {
	HighlimitCaller     // Read-only binding to the contract
	HighlimitTransactor // Write-only binding to the contract
	HighlimitFilterer   // Log filterer for contract events
}

// HighlimitCaller is an auto generated read-only Go binding around an Ethereum contract.
type HighlimitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HighlimitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HighlimitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HighlimitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HighlimitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HighlimitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HighlimitSession struct {
	Contract     *Highlimit        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HighlimitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HighlimitCallerSession struct {
	Contract *HighlimitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HighlimitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HighlimitTransactorSession struct {
	Contract     *HighlimitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HighlimitRaw is an auto generated low-level Go binding around an Ethereum contract.
type HighlimitRaw struct {
	Contract *Highlimit // Generic contract binding to access the raw methods on
}

// HighlimitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HighlimitCallerRaw struct {
	Contract *HighlimitCaller // Generic read-only contract binding to access the raw methods on
}

// HighlimitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HighlimitTransactorRaw struct {
	Contract *HighlimitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHighlimit creates a new instance of Highlimit, bound to a specific deployed contract.
func NewHighlimit(address common.Address, backend bind.ContractBackend) (*Highlimit, error) {
	contract, err := bindHighlimit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Highlimit{HighlimitCaller: HighlimitCaller{contract: contract}, HighlimitTransactor: HighlimitTransactor{contract: contract}, HighlimitFilterer: HighlimitFilterer{contract: contract}}, nil
}

// NewHighlimitCaller creates a new read-only instance of Highlimit, bound to a specific deployed contract.
func NewHighlimitCaller(address common.Address, caller bind.ContractCaller) (*HighlimitCaller, error) {
	contract, err := bindHighlimit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HighlimitCaller{contract: contract}, nil
}

// NewHighlimitTransactor creates a new write-only instance of Highlimit, bound to a specific deployed contract.
func NewHighlimitTransactor(address common.Address, transactor bind.ContractTransactor) (*HighlimitTransactor, error) {
	contract, err := bindHighlimit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HighlimitTransactor{contract: contract}, nil
}

// NewHighlimitFilterer creates a new log filterer instance of Highlimit, bound to a specific deployed contract.
func NewHighlimitFilterer(address common.Address, filterer bind.ContractFilterer) (*HighlimitFilterer, error) {
	contract, err := bindHighlimit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HighlimitFilterer{contract: contract}, nil
}

// bindHighlimit binds a generic wrapper to an already deployed contract.
func bindHighlimit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HighlimitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Highlimit *HighlimitRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Highlimit.Contract.HighlimitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Highlimit *HighlimitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Highlimit.Contract.HighlimitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Highlimit *HighlimitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Highlimit.Contract.HighlimitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Highlimit *HighlimitCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Highlimit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Highlimit *HighlimitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Highlimit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Highlimit *HighlimitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Highlimit.Contract.contract.Transact(opts, method, params...)
}

// Del is a paid mutator transaction binding the contract method 0xb6588ffd.
//
// Solidity: function del() returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_Highlimit *HighlimitTransactor) Del(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Highlimit.contract.Transact(opts, "del")
}

// Del is a paid mutator transaction binding the contract method 0xb6588ffd.
//
// Solidity: function del() returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_Highlimit *HighlimitSession) Del() (*types.Transaction, error) {
	return _Highlimit.Contract.Del(&_Highlimit.TransactOpts)
}

// Del is a paid mutator transaction binding the contract method 0xb6588ffd.
//
// Solidity: function del() returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_Highlimit *HighlimitTransactorSession) Del() (*types.Transaction, error) {
	return _Highlimit.Contract.Del(&_Highlimit.TransactOpts)
}

// Run is a paid mutator transaction binding the contract method 0xc0406226.
//
// Solidity: function run() returns()
func (_Highlimit *HighlimitTransactor) Run(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Highlimit.contract.Transact(opts, "run")
}

// Run is a paid mutator transaction binding the contract method 0xc0406226.
//
// Solidity: function run() returns()
func (_Highlimit *HighlimitSession) Run() (*types.Transaction, error) {
	return _Highlimit.Contract.Run(&_Highlimit.TransactOpts)
}

// Run is a paid mutator transaction binding the contract method 0xc0406226.
//
// Solidity: function run() returns()
func (_Highlimit *HighlimitTransactorSession) Run() (*types.Transaction, error) {
	return _Highlimit.Contract.Run(&_Highlimit.TransactOpts)
}

// HighlimitOnDelIterator is returned from FilterOnDel and is used to iterate over the raw logs and unpacked data for OnDel events raised by the Highlimit contract.
type HighlimitOnDelIterator struct {
	Event *HighlimitOnDel // Event containing the contract specifics and raw log

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
func (it *HighlimitOnDelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HighlimitOnDel)
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
		it.Event = new(HighlimitOnDel)
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
func (it *HighlimitOnDelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HighlimitOnDelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HighlimitOnDel represents a OnDel event raised by the Highlimit contract.
type HighlimitOnDel struct {
	A1  *big.Int
	A2  *big.Int
	A3  *big.Int
	A4  *big.Int
	A5  *big.Int
	A6  *big.Int
	A17 *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOnDel is a free log retrieval operation binding the contract event 0x52f847c6da443e437b0d57a554be808022a7e69637c9ed63a1e16d4b63413411.
//
// Solidity: event onDel(a1 uint256, a2 uint256, a3 uint256, a4 uint256, a5 uint256, a6 uint256, a17 uint256)
func (_Highlimit *HighlimitFilterer) FilterOnDel(opts *bind.FilterOpts) (*HighlimitOnDelIterator, error) {

	logs, sub, err := _Highlimit.contract.FilterLogs(opts, "onDel")
	if err != nil {
		return nil, err
	}
	return &HighlimitOnDelIterator{contract: _Highlimit.contract, event: "onDel", logs: logs, sub: sub}, nil
}

// WatchOnDel is a free log subscription operation binding the contract event 0x52f847c6da443e437b0d57a554be808022a7e69637c9ed63a1e16d4b63413411.
//
// Solidity: event onDel(a1 uint256, a2 uint256, a3 uint256, a4 uint256, a5 uint256, a6 uint256, a17 uint256)
func (_Highlimit *HighlimitFilterer) WatchOnDel(opts *bind.WatchOpts, sink chan<- *HighlimitOnDel) (event.Subscription, error) {

	logs, sub, err := _Highlimit.contract.WatchLogs(opts, "onDel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HighlimitOnDel)
				if err := _Highlimit.contract.UnpackLog(event, "onDel", log); err != nil {
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
