// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// DatasetsABI is the input ABI used to generate the binding from.
const DatasetsABI = "[]"

// DatasetsBin is the compiled bytecode used for deploying new contracts.
const DatasetsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820f9d36c82d281fd45930594e52752a1869f7d4d46446d9067e33fe89c8326f0f90029`

// DeployDatasets deploys a new Ethereum contract, binding an instance of Datasets to it.
func DeployDatasets(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Datasets, error) {
	parsed, err := abi.JSON(strings.NewReader(DatasetsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DatasetsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Datasets{DatasetsCaller: DatasetsCaller{contract: contract}, DatasetsTransactor: DatasetsTransactor{contract: contract}, DatasetsFilterer: DatasetsFilterer{contract: contract}}, nil
}

// Datasets is an auto generated Go binding around an Ethereum contract.
type Datasets struct {
	DatasetsCaller     // Read-only binding to the contract
	DatasetsTransactor // Write-only binding to the contract
	DatasetsFilterer   // Log filterer for contract events
}

// DatasetsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DatasetsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatasetsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DatasetsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatasetsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DatasetsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatasetsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DatasetsSession struct {
	Contract     *Datasets         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DatasetsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DatasetsCallerSession struct {
	Contract *DatasetsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DatasetsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DatasetsTransactorSession struct {
	Contract     *DatasetsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DatasetsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DatasetsRaw struct {
	Contract *Datasets // Generic contract binding to access the raw methods on
}

// DatasetsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DatasetsCallerRaw struct {
	Contract *DatasetsCaller // Generic read-only contract binding to access the raw methods on
}

// DatasetsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DatasetsTransactorRaw struct {
	Contract *DatasetsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDatasets creates a new instance of Datasets, bound to a specific deployed contract.
func NewDatasets(address common.Address, backend bind.ContractBackend) (*Datasets, error) {
	contract, err := bindDatasets(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Datasets{DatasetsCaller: DatasetsCaller{contract: contract}, DatasetsTransactor: DatasetsTransactor{contract: contract}, DatasetsFilterer: DatasetsFilterer{contract: contract}}, nil
}

// NewDatasetsCaller creates a new read-only instance of Datasets, bound to a specific deployed contract.
func NewDatasetsCaller(address common.Address, caller bind.ContractCaller) (*DatasetsCaller, error) {
	contract, err := bindDatasets(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DatasetsCaller{contract: contract}, nil
}

// NewDatasetsTransactor creates a new write-only instance of Datasets, bound to a specific deployed contract.
func NewDatasetsTransactor(address common.Address, transactor bind.ContractTransactor) (*DatasetsTransactor, error) {
	contract, err := bindDatasets(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DatasetsTransactor{contract: contract}, nil
}

// NewDatasetsFilterer creates a new log filterer instance of Datasets, bound to a specific deployed contract.
func NewDatasetsFilterer(address common.Address, filterer bind.ContractFilterer) (*DatasetsFilterer, error) {
	contract, err := bindDatasets(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DatasetsFilterer{contract: contract}, nil
}

// bindDatasets binds a generic wrapper to an already deployed contract.
func bindDatasets(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DatasetsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datasets *DatasetsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Datasets.Contract.DatasetsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datasets *DatasetsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datasets.Contract.DatasetsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datasets *DatasetsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datasets.Contract.DatasetsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datasets *DatasetsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Datasets.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datasets *DatasetsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datasets.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datasets *DatasetsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datasets.Contract.contract.Transact(opts, method, params...)
}

// LotteryeventsABI is the input ABI used to generate the binding from.
const LotteryeventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_team\",\"type\":\"uint8\"}],\"name\":\"onBuys\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"out\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"rid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"strt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"etc0\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"etc1\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"win\",\"type\":\"int256\"}],\"name\":\"onBuyAndDistribute\",\"type\":\"event\"}]"

// LotteryeventsBin is the compiled bytecode used for deploying new contracts.
const LotteryeventsBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a7230582050e48192f02a360a91ecac7118d768677a4df80f4bbf9d2916778af16c61d16b0029`

// DeployLotteryevents deploys a new Ethereum contract, binding an instance of Lotteryevents to it.
func DeployLotteryevents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lotteryevents, error) {
	parsed, err := abi.JSON(strings.NewReader(LotteryeventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LotteryeventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lotteryevents{LotteryeventsCaller: LotteryeventsCaller{contract: contract}, LotteryeventsTransactor: LotteryeventsTransactor{contract: contract}, LotteryeventsFilterer: LotteryeventsFilterer{contract: contract}}, nil
}

// Lotteryevents is an auto generated Go binding around an Ethereum contract.
type Lotteryevents struct {
	LotteryeventsCaller     // Read-only binding to the contract
	LotteryeventsTransactor // Write-only binding to the contract
	LotteryeventsFilterer   // Log filterer for contract events
}

// LotteryeventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type LotteryeventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryeventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LotteryeventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryeventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LotteryeventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryeventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LotteryeventsSession struct {
	Contract     *Lotteryevents    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LotteryeventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LotteryeventsCallerSession struct {
	Contract *LotteryeventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LotteryeventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LotteryeventsTransactorSession struct {
	Contract     *LotteryeventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LotteryeventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type LotteryeventsRaw struct {
	Contract *Lotteryevents // Generic contract binding to access the raw methods on
}

// LotteryeventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LotteryeventsCallerRaw struct {
	Contract *LotteryeventsCaller // Generic read-only contract binding to access the raw methods on
}

// LotteryeventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LotteryeventsTransactorRaw struct {
	Contract *LotteryeventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLotteryevents creates a new instance of Lotteryevents, bound to a specific deployed contract.
func NewLotteryevents(address common.Address, backend bind.ContractBackend) (*Lotteryevents, error) {
	contract, err := bindLotteryevents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lotteryevents{LotteryeventsCaller: LotteryeventsCaller{contract: contract}, LotteryeventsTransactor: LotteryeventsTransactor{contract: contract}, LotteryeventsFilterer: LotteryeventsFilterer{contract: contract}}, nil
}

// NewLotteryeventsCaller creates a new read-only instance of Lotteryevents, bound to a specific deployed contract.
func NewLotteryeventsCaller(address common.Address, caller bind.ContractCaller) (*LotteryeventsCaller, error) {
	contract, err := bindLotteryevents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryeventsCaller{contract: contract}, nil
}

// NewLotteryeventsTransactor creates a new write-only instance of Lotteryevents, bound to a specific deployed contract.
func NewLotteryeventsTransactor(address common.Address, transactor bind.ContractTransactor) (*LotteryeventsTransactor, error) {
	contract, err := bindLotteryevents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryeventsTransactor{contract: contract}, nil
}

// NewLotteryeventsFilterer creates a new log filterer instance of Lotteryevents, bound to a specific deployed contract.
func NewLotteryeventsFilterer(address common.Address, filterer bind.ContractFilterer) (*LotteryeventsFilterer, error) {
	contract, err := bindLotteryevents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LotteryeventsFilterer{contract: contract}, nil
}

// bindLotteryevents binds a generic wrapper to an already deployed contract.
func bindLotteryevents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LotteryeventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lotteryevents *LotteryeventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lotteryevents.Contract.LotteryeventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lotteryevents *LotteryeventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lotteryevents.Contract.LotteryeventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lotteryevents *LotteryeventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lotteryevents.Contract.LotteryeventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lotteryevents *LotteryeventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lotteryevents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lotteryevents *LotteryeventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lotteryevents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lotteryevents *LotteryeventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lotteryevents.Contract.contract.Transact(opts, method, params...)
}

// LotteryeventsOnBuyAndDistributeIterator is returned from FilterOnBuyAndDistribute and is used to iterate over the raw logs and unpacked data for OnBuyAndDistribute events raised by the Lotteryevents contract.
type LotteryeventsOnBuyAndDistributeIterator struct {
	Event *LotteryeventsOnBuyAndDistribute // Event containing the contract specifics and raw log

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
func (it *LotteryeventsOnBuyAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryeventsOnBuyAndDistribute)
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
		it.Event = new(LotteryeventsOnBuyAndDistribute)
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
func (it *LotteryeventsOnBuyAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryeventsOnBuyAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryeventsOnBuyAndDistribute represents a OnBuyAndDistribute event raised by the Lotteryevents contract.
type LotteryeventsOnBuyAndDistribute struct {
	Rid  *big.Int
	Strt *big.Int
	End  *big.Int
	Etc0 *big.Int
	Etc1 *big.Int
	Win  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnBuyAndDistribute is a free log retrieval operation binding the contract event 0x755811ed4d5c4c398a7bfa8123f92c98e628b975243e845e607f1c649f9cf7ac.
//
// Solidity: e onBuyAndDistribute(rid uint256, strt uint256, end uint256, etc0 uint256, etc1 uint256, win int256)
func (_Lotteryevents *LotteryeventsFilterer) FilterOnBuyAndDistribute(opts *bind.FilterOpts) (*LotteryeventsOnBuyAndDistributeIterator, error) {

	logs, sub, err := _Lotteryevents.contract.FilterLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return &LotteryeventsOnBuyAndDistributeIterator{contract: _Lotteryevents.contract, event: "onBuyAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnBuyAndDistribute is a free log subscription operation binding the contract event 0x755811ed4d5c4c398a7bfa8123f92c98e628b975243e845e607f1c649f9cf7ac.
//
// Solidity: e onBuyAndDistribute(rid uint256, strt uint256, end uint256, etc0 uint256, etc1 uint256, win int256)
func (_Lotteryevents *LotteryeventsFilterer) WatchOnBuyAndDistribute(opts *bind.WatchOpts, sink chan<- *LotteryeventsOnBuyAndDistribute) (event.Subscription, error) {

	logs, sub, err := _Lotteryevents.contract.WatchLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryeventsOnBuyAndDistribute)
				if err := _Lotteryevents.contract.UnpackLog(event, "onBuyAndDistribute", log); err != nil {
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

// LotteryeventsOnBuysIterator is returned from FilterOnBuys and is used to iterate over the raw logs and unpacked data for OnBuys events raised by the Lotteryevents contract.
type LotteryeventsOnBuysIterator struct {
	Event *LotteryeventsOnBuys // Event containing the contract specifics and raw log

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
func (it *LotteryeventsOnBuysIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryeventsOnBuys)
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
		it.Event = new(LotteryeventsOnBuys)
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
func (it *LotteryeventsOnBuysIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryeventsOnBuysIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryeventsOnBuys represents a OnBuys event raised by the Lotteryevents contract.
type LotteryeventsOnBuys struct {
	Addr   common.Address
	Amount *big.Int
	Team   uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOnBuys is a free log retrieval operation binding the contract event 0xd3262b557966269a237416d28e9e625027d870723fa202fc25a2df9912b34d60.
//
// Solidity: e onBuys(addr address, amount uint256, _team uint8)
func (_Lotteryevents *LotteryeventsFilterer) FilterOnBuys(opts *bind.FilterOpts) (*LotteryeventsOnBuysIterator, error) {

	logs, sub, err := _Lotteryevents.contract.FilterLogs(opts, "onBuys")
	if err != nil {
		return nil, err
	}
	return &LotteryeventsOnBuysIterator{contract: _Lotteryevents.contract, event: "onBuys", logs: logs, sub: sub}, nil
}

// WatchOnBuys is a free log subscription operation binding the contract event 0xd3262b557966269a237416d28e9e625027d870723fa202fc25a2df9912b34d60.
//
// Solidity: e onBuys(addr address, amount uint256, _team uint8)
func (_Lotteryevents *LotteryeventsFilterer) WatchOnBuys(opts *bind.WatchOpts, sink chan<- *LotteryeventsOnBuys) (event.Subscription, error) {

	logs, sub, err := _Lotteryevents.contract.WatchLogs(opts, "onBuys")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryeventsOnBuys)
				if err := _Lotteryevents.contract.UnpackLog(event, "onBuys", log); err != nil {
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

// LotteryeventsOnWithdrawIterator is returned from FilterOnWithdraw and is used to iterate over the raw logs and unpacked data for OnWithdraw events raised by the Lotteryevents contract.
type LotteryeventsOnWithdrawIterator struct {
	Event *LotteryeventsOnWithdraw // Event containing the contract specifics and raw log

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
func (it *LotteryeventsOnWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryeventsOnWithdraw)
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
		it.Event = new(LotteryeventsOnWithdraw)
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
func (it *LotteryeventsOnWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryeventsOnWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryeventsOnWithdraw represents a OnWithdraw event raised by the Lotteryevents contract.
type LotteryeventsOnWithdraw struct {
	PlayerAddress common.Address
	Out           *big.Int
	TimeStamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOnWithdraw is a free log retrieval operation binding the contract event 0x1b091269e929df55d64d6ea7e9cadbe4fb38dce5ccdc995767bc515030dbfbbf.
//
// Solidity: e onWithdraw(playerAddress address, out uint256, timeStamp uint256)
func (_Lotteryevents *LotteryeventsFilterer) FilterOnWithdraw(opts *bind.FilterOpts) (*LotteryeventsOnWithdrawIterator, error) {

	logs, sub, err := _Lotteryevents.contract.FilterLogs(opts, "onWithdraw")
	if err != nil {
		return nil, err
	}
	return &LotteryeventsOnWithdrawIterator{contract: _Lotteryevents.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0x1b091269e929df55d64d6ea7e9cadbe4fb38dce5ccdc995767bc515030dbfbbf.
//
// Solidity: e onWithdraw(playerAddress address, out uint256, timeStamp uint256)
func (_Lotteryevents *LotteryeventsFilterer) WatchOnWithdraw(opts *bind.WatchOpts, sink chan<- *LotteryeventsOnWithdraw) (event.Subscription, error) {

	logs, sub, err := _Lotteryevents.contract.WatchLogs(opts, "onWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryeventsOnWithdraw)
				if err := _Lotteryevents.contract.UnpackLog(event, "onWithdraw", log); err != nil {
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

// NXlotteryABI is the input ABI used to generate the binding from.
const NXlotteryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCreator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_team\",\"type\":\"uint8\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAddressLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"playerWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEndowmentBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastRoundInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAddressArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getPlayerInfoByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"_team\",\"type\":\"uint8\"}],\"name\":\"reinvest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentRoundLeft\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_team\",\"type\":\"uint8\"}],\"name\":\"onBuys\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"out\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"rid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"strt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"etc0\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"etc1\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"win\",\"type\":\"int256\"}],\"name\":\"onBuyAndDistribute\",\"type\":\"event\"}]"

// NXlotteryBin is the compiled bytecode used for deploying new contracts.
const NXlotteryBin = `0x6080604052600080556001600e5534801561001957600080fd5b5060018054600160a060020a031916331790554360028190556064016003556004805460ff19169055600060075561140d806100566000396000f3006080604052600436106100e55763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ee2cb10811461021e57806314107f3c1461024f578063285109d11461025f5780632e97766d1461029a5780633b3672bd146102af57806341c0e1b5146102d65780635629365b146102eb5780635acce36b14610303578063705e798e14610318578063a5b0930d14610362578063b93ab1651461039d578063be357616146103ec578063ced72f8714610404578063ee0b5d8b14610419578063f0a3843d1461046d578063fb7ae31e1461048b575b3360003282146100f457600080fd5b50803b801561014d576040805160e560020a62461bcd02815260206004820152601160248201527f736f7272792068756d616e73206f6e6c79000000000000000000000000000000604482015290519081900360640190fd5b34662386f26fc100008110156101ad576040805160e560020a62461bcd02815260206004820152600a60248201527f746f6f206c6974746c6500000000000000000000000000000000000000000000604482015290519081900360640190fd5b68056bc75e2d6310000081111561020e576040805160e560020a62461bcd02815260206004820152600860248201527f746f6f206d756368000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6102193460006104a0565b505050005b34801561022a57600080fd5b5061023361055a565b60408051600160a060020a039092168252519081900360200190f35b61025d60ff6004351661056a565b005b34801561026b57600080fd5b506102746106a3565b604080519485526020850193909352838301919091526060830152519081900360800190f35b3480156102a657600080fd5b5061025d6106b5565b3480156102bb57600080fd5b506102c46106bf565b60408051918252519081900360200190f35b3480156102e257600080fd5b5061025d6106c5565b3480156102f757600080fd5b5061025d6004356106e6565b34801561030f57600080fd5b506102c46107d4565b34801561032457600080fd5b5061032d6107d9565b6040805196875260208701959095528585019390935260608501919091521515608084015260a0830152519081900360c00190f35b34801561036e57600080fd5b506103776107f7565b60408051600160a060020a03938416815291909216602082015281519081900390910190f35b3480156103a957600080fd5b506103b261084f565b6040805197885260208801969096528686019490945260608601929092526080850152151560a084015260c0830152519081900360e00190f35b3480156103f857600080fd5b5061025d600435610871565b34801561041057600080fd5b506102c4610924565b34801561042557600080fd5b5061043a600160a060020a0360043516610990565b604080519687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b34801561047957600080fd5b5061025d60043560ff602435166109f2565b34801561049757600080fd5b506102c4610aec565b60ff811615806104b357508060ff166001145b1515610509576040805160e560020a62461bcd02815260206004820152600b60248201527f7465616d2030206f722031000000000000000000000000000000000000000000604482015290519081900360640190fd5b610514338383610b1e565b604080513381526020810184905260ff83168183015290517fd3262b557966269a237416d28e9e625027d870723fa202fc25a2df9912b34d609181900360600190a15050565b600154600160a060020a03165b90565b33600032821461057957600080fd5b50803b80156105d2576040805160e560020a62461bcd02815260206004820152601160248201527f736f7272792068756d616e73206f6e6c79000000000000000000000000000000604482015290519081900360640190fd5b34662386f26fc10000811015610632576040805160e560020a62461bcd02815260206004820152600a60248201527f746f6f206c6974746c6500000000000000000000000000000000000000000000604482015290519081900360640190fd5b68056bc75e2d63100000811115610693576040805160e560020a62461bcd02815260206004820152600860248201527f746f6f206d756368000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61069d34856104a0565b50505050565b60115460125460135460145490919293565b6106bd610d8f565b565b600f5490565b600154600160a060020a03163314156106bd57600154600160a060020a0316ff5b336000818152601060205260408120600581015460039091015461070f9163ffffffff6111af16565b90508281106107cf57600160a060020a038216600090815260106020526040902060050154610744908463ffffffff61121416565b600160a060020a038316600081815260106020526040808220600501939093559151909185156108fc02918691818181858888f1935050505015801561078e573d6000803e3d6000fd5b506040805133815260208101859052428183015290517f1b091269e929df55d64d6ea7e9cadbe4fb38dce5ccdc995767bc515030dbfbbf9181900360600190a15b505050565b303190565b600854600954600b54600c54600a54600d5460ff9091169192939495565b600080600f600081548110151561080a57fe5b600091825260209091200154600f8054600160a060020a0390921691600190811061083157fe5b600091825260209091200154909250600160a060020a031690509091565b600e5460025460035460055460065460045460075460ff909116919293949596565b600154600160a060020a031633146108d3576040805160e560020a62461bcd02815260206004820152600b60248201527f6e6f742063726561746f72000000000000000000000000000000000000000000604482015290519081900360640190fd5b6000548111610921576000546108ef908263ffffffff6111af16565b6000908155604051339183156108fc02918491818181858888f1935050505015801561091f573d6000803e3d6000fd5b505b50565b600154600090600160a060020a03163314610989576040805160e560020a62461bcd02815260206004820152600b60248201527f6e6f742063726561746f72000000000000000000000000000000000000000000604482015290519081900360640190fd5b5060005490565b6000808080808086600160a060020a03811615156109ab5750335b600160a060020a0316600090815260106020526040902080546001820154600283015460038401546004850154600590950154939c929b5090995097509195509350915050565b6000803381328214610a0357600080fd5b50803b8015610a5c576040805160e560020a62461bcd02815260206004820152601160248201527f736f7272792068756d616e73206f6e6c79000000000000000000000000000000604482015290519081900360640190fd5b3360008181526010602052604090206005810154600390910154919550610a89919063ffffffff6111af16565b9250858310610ae457600160a060020a038416600090815260106020526040902060050154610abe908763ffffffff61121416565b600160a060020a038516600090815260106020526040902060050155610ae486866104a0565b505050505050565b6003546000904390811015610b1557600354610b0e908263ffffffff6111af16565b9150610b1a565b600091505b5090565b6003546000908190600919014310610b8557600354431015610b5f576006546005540191506729a2241af62c0000821015610b5f5743606401600355610b85565b600160a060020a0385166000908152601060205260409020600301805485019055610d88565b600f5460c811610bb557600160a060020a0385166000908152601060205260409020600301805485019055610d88565b5060005b600f54811015610bfd57600f805482908110610bd157fe5b600091825260209091200154600160a060020a0386811691161415610bf557610bfd565b600101610bb9565b600f548110610c5f57600f80546001810182556000919091527f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac80201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387161790555b60ff83161515610cdc57600160a060020a038516600090815260106020526040902054610c92908563ffffffff61121416565b600160a060020a038616600090815260106020526040902055600554610cbe908563ffffffff61121416565b600555601154610cd4908563ffffffff61121416565b601155610d51565b600160a060020a038516600090815260106020526040902060010154610d08908563ffffffff61121416565b600160a060020a038616600090815260106020526040902060010155600654610d37908563ffffffff61121416565b600655601254610d4d908563ffffffff61121416565b6012555b33600090815260106020526040902060020154610d74908563ffffffff61121416565b336000908152601060205260409020600201555b5050505050565b60008060008060008060008060008043600a60ff166002600101540311151515610e03576040805160e560020a62461bcd02815260206004820152600f60248201527f63616e6e6f7420636c656172696e670000000000000000000000000000000000604482015290519081900360640190fd5b60045460ff1615610e5e576040805160e560020a62461bcd02815260206004820152600f60248201527f616c726561647920636c65617265640000000000000000000000000000000000604482015290519081900360640190fd5b6006546005540199506729a2241af62c00008a1015610e8357436064016003556111a3565b600354431015610e92576111a3565b600354610e9e9061126f565b600554909950600010610eca57600f5460c811610ebe5760019850610eca565b436064016003556111a3565b600654600010610ee357600f5460c811610ebe57600098505b60ff8916600755610f0c6064610f008c600563ffffffff6112fb16565b9063ffffffff61137216565b9750610f1e8a8963ffffffff6111af16565b600054909750610f34908963ffffffff61121416565b60005560ff89161515610f4b576005549550610f51565b60065495505b600094505b600f5485101561107057600f805486908110610f6e57fe5b600091825260209091200154600160a060020a0316935060ff89161515610faf57600160a060020a0384166000908152601060205260409020549250610fce565b600160a060020a03841660009081526010602052604090206001015492505b60009150600083111561103b57610fef86610f00898663ffffffff6112fb16565b600160a060020a03851660009081526010602052604090206003015490925061101e908363ffffffff61121416565b600160a060020a0385166000908152601060205260409020600301555b600160a060020a0384166000908152601060205260408120600481018490558181556001908101919091559490940193610f56565b60048054600160ff19918216179182905560025460088190556003546009819055600a805490931660ff909416151593909317909155600554600b819055600654600c819055600754600d819055600e54604080519182526020820195909552808501959095526060850192909252608084015260a0830152517f755811ed4d5c4c398a7bfa8123f92c98e628b975243e845e607f1c649f9cf7ac9160c0908290030190a150600e8054600190810190915543908101600281905560659091016003556004805460ff1916905560006007819055600581905560065560ff8916151561117157601354611169908863ffffffff61121416565b601355611188565b601454611184908863ffffffff61121416565b6014555b611194600f6000611389565b60006111a1600f826113a7565b505b50505050505050505050565b600082821115611209576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d61746820737562206661696c656400000000000000000000000000604482015290519081900360640190fd5b508082035b92915050565b8181018281101561120e576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d61746820616464206661696c656400000000000000000000000000604482015290519081900360640190fd5b60006009198201407f08000000000000000000000000000000000000000000000000000000000000007f0f000000000000000000000000000000000000000000000000000000000000007f0100000000000000000000000000000000000000000000000000000000000000601f84901a021611156112f057600191506112f5565b600091505b50919050565b600082151561130c5750600061120e565b5081810281838281151561131c57fe5b041461120e576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d617468206d756c206661696c656400000000000000000000000000604482015290519081900360640190fd5b600080828481151561138057fe5b04949350505050565b508054600082559060005260206000209081019061092191906113c7565b8154818355818111156107cf576000838152602090206107cf9181019083015b61056791905b80821115610b1a57600081556001016113cd5600a165627a7a723058206aa03dedbe09bb0bb816784d2cdc98f70c7bdeb25cc1f08c4b1d47ae72f99e7b0029`

// DeployNXlottery deploys a new Ethereum contract, binding an instance of NXlottery to it.
func DeployNXlottery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NXlottery, error) {
	parsed, err := abi.JSON(strings.NewReader(NXlotteryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NXlotteryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NXlottery{NXlotteryCaller: NXlotteryCaller{contract: contract}, NXlotteryTransactor: NXlotteryTransactor{contract: contract}, NXlotteryFilterer: NXlotteryFilterer{contract: contract}}, nil
}

// NXlottery is an auto generated Go binding around an Ethereum contract.
type NXlottery struct {
	NXlotteryCaller     // Read-only binding to the contract
	NXlotteryTransactor // Write-only binding to the contract
	NXlotteryFilterer   // Log filterer for contract events
}

// NXlotteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NXlotteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NXlotteryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NXlotteryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NXlotteryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NXlotteryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NXlotterySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NXlotterySession struct {
	Contract     *NXlottery        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NXlotteryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NXlotteryCallerSession struct {
	Contract *NXlotteryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NXlotteryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NXlotteryTransactorSession struct {
	Contract     *NXlotteryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NXlotteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NXlotteryRaw struct {
	Contract *NXlottery // Generic contract binding to access the raw methods on
}

// NXlotteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NXlotteryCallerRaw struct {
	Contract *NXlotteryCaller // Generic read-only contract binding to access the raw methods on
}

// NXlotteryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NXlotteryTransactorRaw struct {
	Contract *NXlotteryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNXlottery creates a new instance of NXlottery, bound to a specific deployed contract.
func NewNXlottery(address common.Address, backend bind.ContractBackend) (*NXlottery, error) {
	contract, err := bindNXlottery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NXlottery{NXlotteryCaller: NXlotteryCaller{contract: contract}, NXlotteryTransactor: NXlotteryTransactor{contract: contract}, NXlotteryFilterer: NXlotteryFilterer{contract: contract}}, nil
}

// NewNXlotteryCaller creates a new read-only instance of NXlottery, bound to a specific deployed contract.
func NewNXlotteryCaller(address common.Address, caller bind.ContractCaller) (*NXlotteryCaller, error) {
	contract, err := bindNXlottery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NXlotteryCaller{contract: contract}, nil
}

// NewNXlotteryTransactor creates a new write-only instance of NXlottery, bound to a specific deployed contract.
func NewNXlotteryTransactor(address common.Address, transactor bind.ContractTransactor) (*NXlotteryTransactor, error) {
	contract, err := bindNXlottery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NXlotteryTransactor{contract: contract}, nil
}

// NewNXlotteryFilterer creates a new log filterer instance of NXlottery, bound to a specific deployed contract.
func NewNXlotteryFilterer(address common.Address, filterer bind.ContractFilterer) (*NXlotteryFilterer, error) {
	contract, err := bindNXlottery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NXlotteryFilterer{contract: contract}, nil
}

// bindNXlottery binds a generic wrapper to an already deployed contract.
func bindNXlottery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NXlotteryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NXlottery *NXlotteryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NXlottery.Contract.NXlotteryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NXlottery *NXlotteryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NXlottery.Contract.NXlotteryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NXlottery *NXlotteryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NXlottery.Contract.NXlotteryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NXlottery *NXlotteryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NXlottery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NXlottery *NXlotteryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NXlottery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NXlottery *NXlotteryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NXlottery.Contract.contract.Transact(opts, method, params...)
}

// GetAddressArray is a free data retrieval call binding the contract method 0xa5b0930d.
//
// Solidity: function getAddressArray() constant returns(address, address)
func (_NXlottery *NXlotteryCaller) GetAddressArray(opts *bind.CallOpts) (common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _NXlottery.contract.Call(opts, out, "getAddressArray")
	return *ret0, *ret1, err
}

// GetAddressArray is a free data retrieval call binding the contract method 0xa5b0930d.
//
// Solidity: function getAddressArray() constant returns(address, address)
func (_NXlottery *NXlotterySession) GetAddressArray() (common.Address, common.Address, error) {
	return _NXlottery.Contract.GetAddressArray(&_NXlottery.CallOpts)
}

// GetAddressArray is a free data retrieval call binding the contract method 0xa5b0930d.
//
// Solidity: function getAddressArray() constant returns(address, address)
func (_NXlottery *NXlotteryCallerSession) GetAddressArray() (common.Address, common.Address, error) {
	return _NXlottery.Contract.GetAddressArray(&_NXlottery.CallOpts)
}

// GetAddressLength is a free data retrieval call binding the contract method 0x3b3672bd.
//
// Solidity: function getAddressLength() constant returns(uint256)
func (_NXlottery *NXlotteryCaller) GetAddressLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NXlottery.contract.Call(opts, out, "getAddressLength")
	return *ret0, err
}

// GetAddressLength is a free data retrieval call binding the contract method 0x3b3672bd.
//
// Solidity: function getAddressLength() constant returns(uint256)
func (_NXlottery *NXlotterySession) GetAddressLength() (*big.Int, error) {
	return _NXlottery.Contract.GetAddressLength(&_NXlottery.CallOpts)
}

// GetAddressLength is a free data retrieval call binding the contract method 0x3b3672bd.
//
// Solidity: function getAddressLength() constant returns(uint256)
func (_NXlottery *NXlotteryCallerSession) GetAddressLength() (*big.Int, error) {
	return _NXlottery.Contract.GetAddressLength(&_NXlottery.CallOpts)
}

// GetCreator is a free data retrieval call binding the contract method 0x0ee2cb10.
//
// Solidity: function getCreator() constant returns(address)
func (_NXlottery *NXlotteryCaller) GetCreator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NXlottery.contract.Call(opts, out, "getCreator")
	return *ret0, err
}

// GetCreator is a free data retrieval call binding the contract method 0x0ee2cb10.
//
// Solidity: function getCreator() constant returns(address)
func (_NXlottery *NXlotterySession) GetCreator() (common.Address, error) {
	return _NXlottery.Contract.GetCreator(&_NXlottery.CallOpts)
}

// GetCreator is a free data retrieval call binding the contract method 0x0ee2cb10.
//
// Solidity: function getCreator() constant returns(address)
func (_NXlottery *NXlotteryCallerSession) GetCreator() (common.Address, error) {
	return _NXlottery.Contract.GetCreator(&_NXlottery.CallOpts)
}

// GetCurrentInfo is a free data retrieval call binding the contract method 0xb93ab165.
//
// Solidity: function getCurrentInfo() constant returns(uint256, uint256, uint256, uint256, uint256, bool, int256)
func (_NXlottery *NXlotteryCaller) GetCurrentInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(bool)
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
	err := _NXlottery.contract.Call(opts, out, "getCurrentInfo")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetCurrentInfo is a free data retrieval call binding the contract method 0xb93ab165.
//
// Solidity: function getCurrentInfo() constant returns(uint256, uint256, uint256, uint256, uint256, bool, int256)
func (_NXlottery *NXlotterySession) GetCurrentInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _NXlottery.Contract.GetCurrentInfo(&_NXlottery.CallOpts)
}

// GetCurrentInfo is a free data retrieval call binding the contract method 0xb93ab165.
//
// Solidity: function getCurrentInfo() constant returns(uint256, uint256, uint256, uint256, uint256, bool, int256)
func (_NXlottery *NXlotteryCallerSession) GetCurrentInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _NXlottery.Contract.GetCurrentInfo(&_NXlottery.CallOpts)
}

// GetCurrentRoundLeft is a free data retrieval call binding the contract method 0xfb7ae31e.
//
// Solidity: function getCurrentRoundLeft() constant returns(uint256)
func (_NXlottery *NXlotteryCaller) GetCurrentRoundLeft(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NXlottery.contract.Call(opts, out, "getCurrentRoundLeft")
	return *ret0, err
}

// GetCurrentRoundLeft is a free data retrieval call binding the contract method 0xfb7ae31e.
//
// Solidity: function getCurrentRoundLeft() constant returns(uint256)
func (_NXlottery *NXlotterySession) GetCurrentRoundLeft() (*big.Int, error) {
	return _NXlottery.Contract.GetCurrentRoundLeft(&_NXlottery.CallOpts)
}

// GetCurrentRoundLeft is a free data retrieval call binding the contract method 0xfb7ae31e.
//
// Solidity: function getCurrentRoundLeft() constant returns(uint256)
func (_NXlottery *NXlotteryCallerSession) GetCurrentRoundLeft() (*big.Int, error) {
	return _NXlottery.Contract.GetCurrentRoundLeft(&_NXlottery.CallOpts)
}

// GetEndowmentBalance is a free data retrieval call binding the contract method 0x5acce36b.
//
// Solidity: function getEndowmentBalance() constant returns(uint256)
func (_NXlottery *NXlotteryCaller) GetEndowmentBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NXlottery.contract.Call(opts, out, "getEndowmentBalance")
	return *ret0, err
}

// GetEndowmentBalance is a free data retrieval call binding the contract method 0x5acce36b.
//
// Solidity: function getEndowmentBalance() constant returns(uint256)
func (_NXlottery *NXlotterySession) GetEndowmentBalance() (*big.Int, error) {
	return _NXlottery.Contract.GetEndowmentBalance(&_NXlottery.CallOpts)
}

// GetEndowmentBalance is a free data retrieval call binding the contract method 0x5acce36b.
//
// Solidity: function getEndowmentBalance() constant returns(uint256)
func (_NXlottery *NXlotteryCallerSession) GetEndowmentBalance() (*big.Int, error) {
	return _NXlottery.Contract.GetEndowmentBalance(&_NXlottery.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() constant returns(uint256)
func (_NXlottery *NXlotteryCaller) GetFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NXlottery.contract.Call(opts, out, "getFee")
	return *ret0, err
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() constant returns(uint256)
func (_NXlottery *NXlotterySession) GetFee() (*big.Int, error) {
	return _NXlottery.Contract.GetFee(&_NXlottery.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() constant returns(uint256)
func (_NXlottery *NXlotteryCallerSession) GetFee() (*big.Int, error) {
	return _NXlottery.Contract.GetFee(&_NXlottery.CallOpts)
}

// GetLastRoundInfo is a free data retrieval call binding the contract method 0x705e798e.
//
// Solidity: function getLastRoundInfo() constant returns(uint256, uint256, uint256, uint256, bool, int256)
func (_NXlottery *NXlotteryCaller) GetLastRoundInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(bool)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _NXlottery.contract.Call(opts, out, "getLastRoundInfo")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetLastRoundInfo is a free data retrieval call binding the contract method 0x705e798e.
//
// Solidity: function getLastRoundInfo() constant returns(uint256, uint256, uint256, uint256, bool, int256)
func (_NXlottery *NXlotterySession) GetLastRoundInfo() (*big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _NXlottery.Contract.GetLastRoundInfo(&_NXlottery.CallOpts)
}

// GetLastRoundInfo is a free data retrieval call binding the contract method 0x705e798e.
//
// Solidity: function getLastRoundInfo() constant returns(uint256, uint256, uint256, uint256, bool, int256)
func (_NXlottery *NXlotteryCallerSession) GetLastRoundInfo() (*big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _NXlottery.Contract.GetLastRoundInfo(&_NXlottery.CallOpts)
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(addr address) constant returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_NXlottery *NXlotteryCaller) GetPlayerInfoByAddress(opts *bind.CallOpts, addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _NXlottery.contract.Call(opts, out, "getPlayerInfoByAddress", addr)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(addr address) constant returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_NXlottery *NXlotterySession) GetPlayerInfoByAddress(addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _NXlottery.Contract.GetPlayerInfoByAddress(&_NXlottery.CallOpts, addr)
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(addr address) constant returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_NXlottery *NXlotteryCallerSession) GetPlayerInfoByAddress(addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _NXlottery.Contract.GetPlayerInfoByAddress(&_NXlottery.CallOpts, addr)
}

// GetTotalInfo is a free data retrieval call binding the contract method 0x285109d1.
//
// Solidity: function getTotalInfo() constant returns(uint256, uint256, uint256, uint256)
func (_NXlottery *NXlotteryCaller) GetTotalInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _NXlottery.contract.Call(opts, out, "getTotalInfo")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetTotalInfo is a free data retrieval call binding the contract method 0x285109d1.
//
// Solidity: function getTotalInfo() constant returns(uint256, uint256, uint256, uint256)
func (_NXlottery *NXlotterySession) GetTotalInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _NXlottery.Contract.GetTotalInfo(&_NXlottery.CallOpts)
}

// GetTotalInfo is a free data retrieval call binding the contract method 0x285109d1.
//
// Solidity: function getTotalInfo() constant returns(uint256, uint256, uint256, uint256)
func (_NXlottery *NXlotteryCallerSession) GetTotalInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _NXlottery.Contract.GetTotalInfo(&_NXlottery.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0x14107f3c.
//
// Solidity: function buy(_team uint8) returns()
func (_NXlottery *NXlotteryTransactor) Buy(opts *bind.TransactOpts, _team uint8) (*types.Transaction, error) {
	return _NXlottery.contract.Transact(opts, "buy", _team)
}

// Buy is a paid mutator transaction binding the contract method 0x14107f3c.
//
// Solidity: function buy(_team uint8) returns()
func (_NXlottery *NXlotterySession) Buy(_team uint8) (*types.Transaction, error) {
	return _NXlottery.Contract.Buy(&_NXlottery.TransactOpts, _team)
}

// Buy is a paid mutator transaction binding the contract method 0x14107f3c.
//
// Solidity: function buy(_team uint8) returns()
func (_NXlottery *NXlotteryTransactorSession) Buy(_team uint8) (*types.Transaction, error) {
	return _NXlottery.Contract.Buy(&_NXlottery.TransactOpts, _team)
}

// GetBlock is a paid mutator transaction binding the contract method 0x2e97766d.
//
// Solidity: function getBlock() returns()
func (_NXlottery *NXlotteryTransactor) GetBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NXlottery.contract.Transact(opts, "getBlock")
}

// GetBlock is a paid mutator transaction binding the contract method 0x2e97766d.
//
// Solidity: function getBlock() returns()
func (_NXlottery *NXlotterySession) GetBlock() (*types.Transaction, error) {
	return _NXlottery.Contract.GetBlock(&_NXlottery.TransactOpts)
}

// GetBlock is a paid mutator transaction binding the contract method 0x2e97766d.
//
// Solidity: function getBlock() returns()
func (_NXlottery *NXlotteryTransactorSession) GetBlock() (*types.Transaction, error) {
	return _NXlottery.Contract.GetBlock(&_NXlottery.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_NXlottery *NXlotteryTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NXlottery.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_NXlottery *NXlotterySession) Kill() (*types.Transaction, error) {
	return _NXlottery.Contract.Kill(&_NXlottery.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_NXlottery *NXlotteryTransactorSession) Kill() (*types.Transaction, error) {
	return _NXlottery.Contract.Kill(&_NXlottery.TransactOpts)
}

// PlayerWithdraw is a paid mutator transaction binding the contract method 0x5629365b.
//
// Solidity: function playerWithdraw(amount uint256) returns()
func (_NXlottery *NXlotteryTransactor) PlayerWithdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _NXlottery.contract.Transact(opts, "playerWithdraw", amount)
}

// PlayerWithdraw is a paid mutator transaction binding the contract method 0x5629365b.
//
// Solidity: function playerWithdraw(amount uint256) returns()
func (_NXlottery *NXlotterySession) PlayerWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _NXlottery.Contract.PlayerWithdraw(&_NXlottery.TransactOpts, amount)
}

// PlayerWithdraw is a paid mutator transaction binding the contract method 0x5629365b.
//
// Solidity: function playerWithdraw(amount uint256) returns()
func (_NXlottery *NXlotteryTransactorSession) PlayerWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _NXlottery.Contract.PlayerWithdraw(&_NXlottery.TransactOpts, amount)
}

// Reinvest is a paid mutator transaction binding the contract method 0xf0a3843d.
//
// Solidity: function reinvest(amount uint256, _team uint8) returns()
func (_NXlottery *NXlotteryTransactor) Reinvest(opts *bind.TransactOpts, amount *big.Int, _team uint8) (*types.Transaction, error) {
	return _NXlottery.contract.Transact(opts, "reinvest", amount, _team)
}

// Reinvest is a paid mutator transaction binding the contract method 0xf0a3843d.
//
// Solidity: function reinvest(amount uint256, _team uint8) returns()
func (_NXlottery *NXlotterySession) Reinvest(amount *big.Int, _team uint8) (*types.Transaction, error) {
	return _NXlottery.Contract.Reinvest(&_NXlottery.TransactOpts, amount, _team)
}

// Reinvest is a paid mutator transaction binding the contract method 0xf0a3843d.
//
// Solidity: function reinvest(amount uint256, _team uint8) returns()
func (_NXlottery *NXlotteryTransactorSession) Reinvest(amount *big.Int, _team uint8) (*types.Transaction, error) {
	return _NXlottery.Contract.Reinvest(&_NXlottery.TransactOpts, amount, _team)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(amount uint256) returns()
func (_NXlottery *NXlotteryTransactor) WithdrawFee(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _NXlottery.contract.Transact(opts, "withdrawFee", amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(amount uint256) returns()
func (_NXlottery *NXlotterySession) WithdrawFee(amount *big.Int) (*types.Transaction, error) {
	return _NXlottery.Contract.WithdrawFee(&_NXlottery.TransactOpts, amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(amount uint256) returns()
func (_NXlottery *NXlotteryTransactorSession) WithdrawFee(amount *big.Int) (*types.Transaction, error) {
	return _NXlottery.Contract.WithdrawFee(&_NXlottery.TransactOpts, amount)
}

// NXlotteryOnBuyAndDistributeIterator is returned from FilterOnBuyAndDistribute and is used to iterate over the raw logs and unpacked data for OnBuyAndDistribute events raised by the NXlottery contract.
type NXlotteryOnBuyAndDistributeIterator struct {
	Event *NXlotteryOnBuyAndDistribute // Event containing the contract specifics and raw log

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
func (it *NXlotteryOnBuyAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NXlotteryOnBuyAndDistribute)
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
		it.Event = new(NXlotteryOnBuyAndDistribute)
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
func (it *NXlotteryOnBuyAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NXlotteryOnBuyAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NXlotteryOnBuyAndDistribute represents a OnBuyAndDistribute event raised by the NXlottery contract.
type NXlotteryOnBuyAndDistribute struct {
	Rid  *big.Int
	Strt *big.Int
	End  *big.Int
	Etc0 *big.Int
	Etc1 *big.Int
	Win  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnBuyAndDistribute is a free log retrieval operation binding the contract event 0x755811ed4d5c4c398a7bfa8123f92c98e628b975243e845e607f1c649f9cf7ac.
//
// Solidity: e onBuyAndDistribute(rid uint256, strt uint256, end uint256, etc0 uint256, etc1 uint256, win int256)
func (_NXlottery *NXlotteryFilterer) FilterOnBuyAndDistribute(opts *bind.FilterOpts) (*NXlotteryOnBuyAndDistributeIterator, error) {

	logs, sub, err := _NXlottery.contract.FilterLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return &NXlotteryOnBuyAndDistributeIterator{contract: _NXlottery.contract, event: "onBuyAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnBuyAndDistribute is a free log subscription operation binding the contract event 0x755811ed4d5c4c398a7bfa8123f92c98e628b975243e845e607f1c649f9cf7ac.
//
// Solidity: e onBuyAndDistribute(rid uint256, strt uint256, end uint256, etc0 uint256, etc1 uint256, win int256)
func (_NXlottery *NXlotteryFilterer) WatchOnBuyAndDistribute(opts *bind.WatchOpts, sink chan<- *NXlotteryOnBuyAndDistribute) (event.Subscription, error) {

	logs, sub, err := _NXlottery.contract.WatchLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NXlotteryOnBuyAndDistribute)
				if err := _NXlottery.contract.UnpackLog(event, "onBuyAndDistribute", log); err != nil {
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

// NXlotteryOnBuysIterator is returned from FilterOnBuys and is used to iterate over the raw logs and unpacked data for OnBuys events raised by the NXlottery contract.
type NXlotteryOnBuysIterator struct {
	Event *NXlotteryOnBuys // Event containing the contract specifics and raw log

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
func (it *NXlotteryOnBuysIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NXlotteryOnBuys)
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
		it.Event = new(NXlotteryOnBuys)
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
func (it *NXlotteryOnBuysIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NXlotteryOnBuysIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NXlotteryOnBuys represents a OnBuys event raised by the NXlottery contract.
type NXlotteryOnBuys struct {
	Addr   common.Address
	Amount *big.Int
	Team   uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOnBuys is a free log retrieval operation binding the contract event 0xd3262b557966269a237416d28e9e625027d870723fa202fc25a2df9912b34d60.
//
// Solidity: e onBuys(addr address, amount uint256, _team uint8)
func (_NXlottery *NXlotteryFilterer) FilterOnBuys(opts *bind.FilterOpts) (*NXlotteryOnBuysIterator, error) {

	logs, sub, err := _NXlottery.contract.FilterLogs(opts, "onBuys")
	if err != nil {
		return nil, err
	}
	return &NXlotteryOnBuysIterator{contract: _NXlottery.contract, event: "onBuys", logs: logs, sub: sub}, nil
}

// WatchOnBuys is a free log subscription operation binding the contract event 0xd3262b557966269a237416d28e9e625027d870723fa202fc25a2df9912b34d60.
//
// Solidity: e onBuys(addr address, amount uint256, _team uint8)
func (_NXlottery *NXlotteryFilterer) WatchOnBuys(opts *bind.WatchOpts, sink chan<- *NXlotteryOnBuys) (event.Subscription, error) {

	logs, sub, err := _NXlottery.contract.WatchLogs(opts, "onBuys")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NXlotteryOnBuys)
				if err := _NXlottery.contract.UnpackLog(event, "onBuys", log); err != nil {
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

// NXlotteryOnWithdrawIterator is returned from FilterOnWithdraw and is used to iterate over the raw logs and unpacked data for OnWithdraw events raised by the NXlottery contract.
type NXlotteryOnWithdrawIterator struct {
	Event *NXlotteryOnWithdraw // Event containing the contract specifics and raw log

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
func (it *NXlotteryOnWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NXlotteryOnWithdraw)
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
		it.Event = new(NXlotteryOnWithdraw)
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
func (it *NXlotteryOnWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NXlotteryOnWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NXlotteryOnWithdraw represents a OnWithdraw event raised by the NXlottery contract.
type NXlotteryOnWithdraw struct {
	PlayerAddress common.Address
	Out           *big.Int
	TimeStamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOnWithdraw is a free log retrieval operation binding the contract event 0x1b091269e929df55d64d6ea7e9cadbe4fb38dce5ccdc995767bc515030dbfbbf.
//
// Solidity: e onWithdraw(playerAddress address, out uint256, timeStamp uint256)
func (_NXlottery *NXlotteryFilterer) FilterOnWithdraw(opts *bind.FilterOpts) (*NXlotteryOnWithdrawIterator, error) {

	logs, sub, err := _NXlottery.contract.FilterLogs(opts, "onWithdraw")
	if err != nil {
		return nil, err
	}
	return &NXlotteryOnWithdrawIterator{contract: _NXlottery.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0x1b091269e929df55d64d6ea7e9cadbe4fb38dce5ccdc995767bc515030dbfbbf.
//
// Solidity: e onWithdraw(playerAddress address, out uint256, timeStamp uint256)
func (_NXlottery *NXlotteryFilterer) WatchOnWithdraw(opts *bind.WatchOpts, sink chan<- *NXlotteryOnWithdraw) (event.Subscription, error) {

	logs, sub, err := _NXlottery.contract.WatchLogs(opts, "onWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NXlotteryOnWithdraw)
				if err := _NXlottery.contract.UnpackLog(event, "onWithdraw", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820e72b54cb680e0c40d8febfa47f268cd413d2331b3560468a47ddc32035b038050029`

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
