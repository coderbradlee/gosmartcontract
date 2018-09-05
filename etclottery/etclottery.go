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
const DatasetsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058205e91c3154cc84a2dbaefe7771843479324630ce4a91e83ea1e7cb492efc460a90029`

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
const LotteryeventsBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a723058206e2976961ec0db319b6468ec0708c05bdc4a81abb9d262a4cabc47cf87ba952b0029`

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
// Solidity: event onBuyAndDistribute(rid uint256, strt uint256, end uint256, etc0 uint256, etc1 uint256, win int256)
func (_Lotteryevents *LotteryeventsFilterer) FilterOnBuyAndDistribute(opts *bind.FilterOpts) (*LotteryeventsOnBuyAndDistributeIterator, error) {

	logs, sub, err := _Lotteryevents.contract.FilterLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return &LotteryeventsOnBuyAndDistributeIterator{contract: _Lotteryevents.contract, event: "onBuyAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnBuyAndDistribute is a free log subscription operation binding the contract event 0x755811ed4d5c4c398a7bfa8123f92c98e628b975243e845e607f1c649f9cf7ac.
//
// Solidity: event onBuyAndDistribute(rid uint256, strt uint256, end uint256, etc0 uint256, etc1 uint256, win int256)
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
// Solidity: event onBuys(addr address, amount uint256, _team uint8)
func (_Lotteryevents *LotteryeventsFilterer) FilterOnBuys(opts *bind.FilterOpts) (*LotteryeventsOnBuysIterator, error) {

	logs, sub, err := _Lotteryevents.contract.FilterLogs(opts, "onBuys")
	if err != nil {
		return nil, err
	}
	return &LotteryeventsOnBuysIterator{contract: _Lotteryevents.contract, event: "onBuys", logs: logs, sub: sub}, nil
}

// WatchOnBuys is a free log subscription operation binding the contract event 0xd3262b557966269a237416d28e9e625027d870723fa202fc25a2df9912b34d60.
//
// Solidity: event onBuys(addr address, amount uint256, _team uint8)
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
// Solidity: event onWithdraw(playerAddress address, out uint256, timeStamp uint256)
func (_Lotteryevents *LotteryeventsFilterer) FilterOnWithdraw(opts *bind.FilterOpts) (*LotteryeventsOnWithdrawIterator, error) {

	logs, sub, err := _Lotteryevents.contract.FilterLogs(opts, "onWithdraw")
	if err != nil {
		return nil, err
	}
	return &LotteryeventsOnWithdrawIterator{contract: _Lotteryevents.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0x1b091269e929df55d64d6ea7e9cadbe4fb38dce5ccdc995767bc515030dbfbbf.
//
// Solidity: event onWithdraw(playerAddress address, out uint256, timeStamp uint256)
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820e1bcd0b8bed676b1126ea028e77202980388efadca741b1b1c86d7dea475cf340029`

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

// EtclotteryABI is the input ABI used to generate the binding from.
const EtclotteryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCreator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_team\",\"type\":\"uint8\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAddressLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"playerWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEndowmentBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastRoundInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAddressArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getPlayerInfoByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"_team\",\"type\":\"uint8\"}],\"name\":\"reinvest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentRoundLeft\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_team\",\"type\":\"uint8\"}],\"name\":\"onBuys\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"out\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"rid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"strt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"etc0\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"etc1\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"win\",\"type\":\"int256\"}],\"name\":\"onBuyAndDistribute\",\"type\":\"event\"}]"

// EtclotteryBin is the compiled bytecode used for deploying new contracts.
const EtclotteryBin = `0x6080604052600080556001600e5534801561001957600080fd5b5060018054600160a060020a031916331790554360028190556028016003556004805460ff19169055600060075561141c806100566000396000f3006080604052600436106100e55763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ee2cb10811461021e57806314107f3c1461024f578063285109d11461025f5780632e97766d1461029a5780633b3672bd146102af57806341c0e1b5146102d65780635629365b146102eb5780635acce36b14610303578063705e798e14610318578063a5b0930d14610362578063b93ab1651461039d578063be357616146103ec578063ced72f8714610404578063ee0b5d8b14610419578063f0a3843d1461046d578063fb7ae31e1461048b575b3360003282146100f457600080fd5b50803b801561014d576040805160e560020a62461bcd02815260206004820152601160248201527f736f7272792068756d616e73206f6e6c79000000000000000000000000000000604482015290519081900360640190fd5b34662386f26fc100008110156101ad576040805160e560020a62461bcd02815260206004820152600a60248201527f746f6f206c6974746c6500000000000000000000000000000000000000000000604482015290519081900360640190fd5b68056bc75e2d6310000081111561020e576040805160e560020a62461bcd02815260206004820152600860248201527f746f6f206d756368000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6102193460006104a0565b505050005b34801561022a57600080fd5b5061023361058a565b60408051600160a060020a039092168252519081900360200190f35b61025d60ff6004351661059a565b005b34801561026b57600080fd5b506102746106d3565b604080519485526020850193909352838301919091526060830152519081900360800190f35b3480156102a657600080fd5b5061025d6106e5565b3480156102bb57600080fd5b506102c46106ef565b60408051918252519081900360200190f35b3480156102e257600080fd5b5061025d6106f5565b3480156102f757600080fd5b5061025d600435610716565b34801561030f57600080fd5b506102c4610804565b34801561032457600080fd5b5061032d610809565b6040805196875260208701959095528585019390935260608501919091521515608084015260a0830152519081900360c00190f35b34801561036e57600080fd5b50610377610827565b60408051600160a060020a03938416815291909216602082015281519081900390910190f35b3480156103a957600080fd5b506103b261087f565b6040805197885260208801969096528686019490945260608601929092526080850152151560a084015260c0830152519081900360e00190f35b3480156103f857600080fd5b5061025d6004356108a1565b34801561041057600080fd5b506102c4610952565b34801561042557600080fd5b5061043a600160a060020a03600435166109be565b604080519687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b34801561047957600080fd5b5061025d60043560ff60243516610a20565b34801561049757600080fd5b506102c4610b1a565b60ff811615806104b357508060ff166001145b1515610509576040805160e560020a62461bcd02815260206004820152600b60248201527f7465616d2030206f722031000000000000000000000000000000000000000000604482015290519081900360640190fd5b610514338383610b4c565b33600090815260106020526040902060020154610537908363ffffffff610ba716565b3360009081526010602052604090206002015560ff8116151561056f57601154610567908363ffffffff610ba716565b601155610586565b601254610582908363ffffffff610ba716565b6012555b5050565b600154600160a060020a03165b90565b3360003282146105a957600080fd5b50803b8015610602576040805160e560020a62461bcd02815260206004820152601160248201527f736f7272792068756d616e73206f6e6c79000000000000000000000000000000604482015290519081900360640190fd5b34662386f26fc10000811015610662576040805160e560020a62461bcd02815260206004820152600a60248201527f746f6f206c6974746c6500000000000000000000000000000000000000000000604482015290519081900360640190fd5b68056bc75e2d631000008111156106c3576040805160e560020a62461bcd02815260206004820152600860248201527f746f6f206d756368000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6106cd34856104a0565b50505050565b60115460125460135460145490919293565b6106ed610c08565b565b600f5490565b600154600160a060020a03163314156106ed57600154600160a060020a0316ff5b336000818152601060205260408120600581015460039091015461073f9163ffffffff61101016565b90508281106107ff57600160a060020a038216600090815260106020526040902060050154610774908463ffffffff610ba716565b600160a060020a038316600081815260106020526040808220600501939093559151909185156108fc02918691818181858888f193505050501580156107be573d6000803e3d6000fd5b506040805133815260208101859052428183015290517f1b091269e929df55d64d6ea7e9cadbe4fb38dce5ccdc995767bc515030dbfbbf9181900360600190a15b505050565b303190565b600854600954600b54600c54600a54600d5460ff9091169192939495565b600080600f600081548110151561083a57fe5b600091825260209091200154600f8054600160a060020a0390921691600190811061086157fe5b600091825260209091200154909250600160a060020a031690509091565b600e5460025460035460055460065460045460075460ff909116919293949596565b600154600160a060020a03163314610903576040805160e560020a62461bcd02815260206004820152600b60248201527f6e6f742063726561746f72000000000000000000000000000000000000000000604482015290519081900360640190fd5b600054811161094f5760005461091f908263ffffffff61101016565b6000908155604051339183156108fc02918491818181858888f19350505050158015610586573d6000803e3d6000fd5b50565b600154600090600160a060020a031633146109b7576040805160e560020a62461bcd02815260206004820152600b60248201527f6e6f742063726561746f72000000000000000000000000000000000000000000604482015290519081900360640190fd5b5060005490565b6000808080808086600160a060020a03811615156109d95750335b600160a060020a0316600090815260106020526040902080546001820154600283015460038401546004850154600590950154939c929b5090995097509195509350915050565b6000803381328214610a3157600080fd5b50803b8015610a8a576040805160e560020a62461bcd02815260206004820152601160248201527f736f7272792068756d616e73206f6e6c79000000000000000000000000000000604482015290519081900360640190fd5b3360008181526010602052604090206005810154600390910154919550610ab7919063ffffffff61101016565b9250858310610b1257600160a060020a038416600090815260106020526040902060050154610aec908763ffffffff610ba716565b600160a060020a038516600090815260106020526040902060050155610b1286866104a0565b505050505050565b6003546000904390811015610b4357600354610b3c908263ffffffff61101016565b9150610b48565b600091505b5090565b610b57838383611070565b60408051600160a060020a03851681526020810184905260ff83168183015290517fd3262b557966269a237416d28e9e625027d870723fa202fc25a2df9912b34d609181900360600190a1505050565b81810182811015610c02576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d61746820616464206661696c656400000000000000000000000000604482015290519081900360640190fd5b92915050565b60008060008060008060008060008043600a60ff166002600101540311151515610c7c576040805160e560020a62461bcd02815260206004820152600f60248201527f63616e6e6f7420636c656172696e670000000000000000000000000000000000604482015290519081900360640190fd5b60045460ff1615610cd7576040805160e560020a62461bcd02815260206004820152600f60248201527f616c726561647920636c65617265640000000000000000000000000000000000604482015290519081900360640190fd5b60065460055401995066b1a2bc2ec500008a1015610cfb5743602801600355611004565b600354431015610d0a57611004565b600554600010610d205743602801600355611004565b600654600010610d365743602801600355611004565b600354610d429061127e565b60ff81166007559850610d6d6064610d618c600863ffffffff61130a16565b9063ffffffff61138116565b9750610d7f8a8963ffffffff61101016565b600054909750610d95908963ffffffff610ba716565b60005560ff89161515610dac576005549550610db2565b60065495505b600094505b600f54851015610ed157600f805486908110610dcf57fe5b600091825260209091200154600160a060020a0316935060ff89161515610e1057600160a060020a0384166000908152601060205260409020549250610e2f565b600160a060020a03841660009081526010602052604090206001015492505b600091506000831115610e9c57610e5086610d61898663ffffffff61130a16565b600160a060020a038516600090815260106020526040902060030154909250610e7f908363ffffffff610ba716565b600160a060020a0385166000908152601060205260409020600301555b600160a060020a0384166000908152601060205260408120600481018490558181556001908101919091559490940193610db7565b60048054600160ff19918216179182905560025460088190556003546009819055600a805490931660ff909416151593909317909155600554600b819055600654600c819055600754600d819055600e54604080519182526020820195909552808501959095526060850192909252608084015260a0830152517f755811ed4d5c4c398a7bfa8123f92c98e628b975243e845e607f1c649f9cf7ac9160c0908290030190a150600e8054600190810190915543908101600281905560299091016003556004805460ff1916905560006007819055600581905560065560ff89161515610fd257601354610fca908863ffffffff610ba716565b601355610fe9565b601454610fe5908863ffffffff610ba716565b6014555b610ff5600f6000611398565b6000611002600f826113b6565b505b50505050505050505050565b60008282111561106a576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d61746820737562206661696c656400000000000000000000000000604482015290519081900360640190fd5b50900390565b60035460009081906009190143106110d6576003544310156110b05760065460055401915066b1a2bc2ec500008210156110b057436028016003556110d6565b600160a060020a0385166000908152601060205260409020600301805485019055611277565b600f546064101561110757600160a060020a0385166000908152601060205260409020600301805485019055611277565b5060005b600f5481101561114f57600f80548290811061112357fe5b600091825260209091200154600160a060020a03868116911614156111475761114f565b60010161110b565b600f5481106111b157600f80546001810182556000919091527f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac80201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387161790555b60ff8316151561121857600160a060020a0385166000908152601060205260409020546111e4908563ffffffff610ba716565b600160a060020a038616600090815260106020526040902055600554611210908563ffffffff610ba716565b600555611277565b600160a060020a038516600090815260106020526040902060010154611244908563ffffffff610ba716565b600160a060020a038616600090815260106020526040902060010155600654611273908563ffffffff610ba716565b6006555b5050505050565b60006009198201407f08000000000000000000000000000000000000000000000000000000000000007f0f000000000000000000000000000000000000000000000000000000000000007f0100000000000000000000000000000000000000000000000000000000000000601f84901a021611156112ff5760019150611304565b600091505b50919050565b600082151561131b57506000610c02565b5081810281838281151561132b57fe5b0414610c02576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d617468206d756c206661696c656400000000000000000000000000604482015290519081900360640190fd5b600080828481151561138f57fe5b04949350505050565b508054600082559060005260206000209081019061094f91906113d6565b8154818355818111156107ff576000838152602090206107ff9181019083015b61059791905b80821115610b4857600081556001016113dc5600a165627a7a723058209abf2f3546796d6e7b8ba94392ad0b4727bade1179c1a5d104fad58c56fca0dc0029`

// DeployEtclottery deploys a new Ethereum contract, binding an instance of Etclottery to it.
func DeployEtclottery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Etclottery, error) {
	parsed, err := abi.JSON(strings.NewReader(EtclotteryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EtclotteryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Etclottery{EtclotteryCaller: EtclotteryCaller{contract: contract}, EtclotteryTransactor: EtclotteryTransactor{contract: contract}, EtclotteryFilterer: EtclotteryFilterer{contract: contract}}, nil
}

// Etclottery is an auto generated Go binding around an Ethereum contract.
type Etclottery struct {
	EtclotteryCaller     // Read-only binding to the contract
	EtclotteryTransactor // Write-only binding to the contract
	EtclotteryFilterer   // Log filterer for contract events
}

// EtclotteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtclotteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtclotteryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtclotteryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtclotteryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtclotteryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtclotterySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtclotterySession struct {
	Contract     *Etclottery       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtclotteryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtclotteryCallerSession struct {
	Contract *EtclotteryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EtclotteryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtclotteryTransactorSession struct {
	Contract     *EtclotteryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EtclotteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtclotteryRaw struct {
	Contract *Etclottery // Generic contract binding to access the raw methods on
}

// EtclotteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtclotteryCallerRaw struct {
	Contract *EtclotteryCaller // Generic read-only contract binding to access the raw methods on
}

// EtclotteryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtclotteryTransactorRaw struct {
	Contract *EtclotteryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtclottery creates a new instance of Etclottery, bound to a specific deployed contract.
func NewEtclottery(address common.Address, backend bind.ContractBackend) (*Etclottery, error) {
	contract, err := bindEtclottery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Etclottery{EtclotteryCaller: EtclotteryCaller{contract: contract}, EtclotteryTransactor: EtclotteryTransactor{contract: contract}, EtclotteryFilterer: EtclotteryFilterer{contract: contract}}, nil
}

// NewEtclotteryCaller creates a new read-only instance of Etclottery, bound to a specific deployed contract.
func NewEtclotteryCaller(address common.Address, caller bind.ContractCaller) (*EtclotteryCaller, error) {
	contract, err := bindEtclottery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtclotteryCaller{contract: contract}, nil
}

// NewEtclotteryTransactor creates a new write-only instance of Etclottery, bound to a specific deployed contract.
func NewEtclotteryTransactor(address common.Address, transactor bind.ContractTransactor) (*EtclotteryTransactor, error) {
	contract, err := bindEtclottery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtclotteryTransactor{contract: contract}, nil
}

// NewEtclotteryFilterer creates a new log filterer instance of Etclottery, bound to a specific deployed contract.
func NewEtclotteryFilterer(address common.Address, filterer bind.ContractFilterer) (*EtclotteryFilterer, error) {
	contract, err := bindEtclottery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtclotteryFilterer{contract: contract}, nil
}

// bindEtclottery binds a generic wrapper to an already deployed contract.
func bindEtclottery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtclotteryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Etclottery *EtclotteryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Etclottery.Contract.EtclotteryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Etclottery *EtclotteryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etclottery.Contract.EtclotteryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Etclottery *EtclotteryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Etclottery.Contract.EtclotteryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Etclottery *EtclotteryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Etclottery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Etclottery *EtclotteryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etclottery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Etclottery *EtclotteryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Etclottery.Contract.contract.Transact(opts, method, params...)
}

// GetAddressArray is a free data retrieval call binding the contract method 0xa5b0930d.
//
// Solidity: function getAddressArray() constant returns(address, address)
func (_Etclottery *EtclotteryCaller) GetAddressArray(opts *bind.CallOpts) (common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Etclottery.contract.Call(opts, out, "getAddressArray")
	return *ret0, *ret1, err
}

// GetAddressArray is a free data retrieval call binding the contract method 0xa5b0930d.
//
// Solidity: function getAddressArray() constant returns(address, address)
func (_Etclottery *EtclotterySession) GetAddressArray() (common.Address, common.Address, error) {
	return _Etclottery.Contract.GetAddressArray(&_Etclottery.CallOpts)
}

// GetAddressArray is a free data retrieval call binding the contract method 0xa5b0930d.
//
// Solidity: function getAddressArray() constant returns(address, address)
func (_Etclottery *EtclotteryCallerSession) GetAddressArray() (common.Address, common.Address, error) {
	return _Etclottery.Contract.GetAddressArray(&_Etclottery.CallOpts)
}

// GetAddressLength is a free data retrieval call binding the contract method 0x3b3672bd.
//
// Solidity: function getAddressLength() constant returns(uint256)
func (_Etclottery *EtclotteryCaller) GetAddressLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Etclottery.contract.Call(opts, out, "getAddressLength")
	return *ret0, err
}

// GetAddressLength is a free data retrieval call binding the contract method 0x3b3672bd.
//
// Solidity: function getAddressLength() constant returns(uint256)
func (_Etclottery *EtclotterySession) GetAddressLength() (*big.Int, error) {
	return _Etclottery.Contract.GetAddressLength(&_Etclottery.CallOpts)
}

// GetAddressLength is a free data retrieval call binding the contract method 0x3b3672bd.
//
// Solidity: function getAddressLength() constant returns(uint256)
func (_Etclottery *EtclotteryCallerSession) GetAddressLength() (*big.Int, error) {
	return _Etclottery.Contract.GetAddressLength(&_Etclottery.CallOpts)
}

// GetCreator is a free data retrieval call binding the contract method 0x0ee2cb10.
//
// Solidity: function getCreator() constant returns(address)
func (_Etclottery *EtclotteryCaller) GetCreator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Etclottery.contract.Call(opts, out, "getCreator")
	return *ret0, err
}

// GetCreator is a free data retrieval call binding the contract method 0x0ee2cb10.
//
// Solidity: function getCreator() constant returns(address)
func (_Etclottery *EtclotterySession) GetCreator() (common.Address, error) {
	return _Etclottery.Contract.GetCreator(&_Etclottery.CallOpts)
}

// GetCreator is a free data retrieval call binding the contract method 0x0ee2cb10.
//
// Solidity: function getCreator() constant returns(address)
func (_Etclottery *EtclotteryCallerSession) GetCreator() (common.Address, error) {
	return _Etclottery.Contract.GetCreator(&_Etclottery.CallOpts)
}

// GetCurrentInfo is a free data retrieval call binding the contract method 0xb93ab165.
//
// Solidity: function getCurrentInfo() constant returns(uint256, uint256, uint256, uint256, uint256, bool, int256)
func (_Etclottery *EtclotteryCaller) GetCurrentInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
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
	err := _Etclottery.contract.Call(opts, out, "getCurrentInfo")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetCurrentInfo is a free data retrieval call binding the contract method 0xb93ab165.
//
// Solidity: function getCurrentInfo() constant returns(uint256, uint256, uint256, uint256, uint256, bool, int256)
func (_Etclottery *EtclotterySession) GetCurrentInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Etclottery.Contract.GetCurrentInfo(&_Etclottery.CallOpts)
}

// GetCurrentInfo is a free data retrieval call binding the contract method 0xb93ab165.
//
// Solidity: function getCurrentInfo() constant returns(uint256, uint256, uint256, uint256, uint256, bool, int256)
func (_Etclottery *EtclotteryCallerSession) GetCurrentInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Etclottery.Contract.GetCurrentInfo(&_Etclottery.CallOpts)
}

// GetCurrentRoundLeft is a free data retrieval call binding the contract method 0xfb7ae31e.
//
// Solidity: function getCurrentRoundLeft() constant returns(uint256)
func (_Etclottery *EtclotteryCaller) GetCurrentRoundLeft(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Etclottery.contract.Call(opts, out, "getCurrentRoundLeft")
	return *ret0, err
}

// GetCurrentRoundLeft is a free data retrieval call binding the contract method 0xfb7ae31e.
//
// Solidity: function getCurrentRoundLeft() constant returns(uint256)
func (_Etclottery *EtclotterySession) GetCurrentRoundLeft() (*big.Int, error) {
	return _Etclottery.Contract.GetCurrentRoundLeft(&_Etclottery.CallOpts)
}

// GetCurrentRoundLeft is a free data retrieval call binding the contract method 0xfb7ae31e.
//
// Solidity: function getCurrentRoundLeft() constant returns(uint256)
func (_Etclottery *EtclotteryCallerSession) GetCurrentRoundLeft() (*big.Int, error) {
	return _Etclottery.Contract.GetCurrentRoundLeft(&_Etclottery.CallOpts)
}

// GetEndowmentBalance is a free data retrieval call binding the contract method 0x5acce36b.
//
// Solidity: function getEndowmentBalance() constant returns(uint256)
func (_Etclottery *EtclotteryCaller) GetEndowmentBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Etclottery.contract.Call(opts, out, "getEndowmentBalance")
	return *ret0, err
}

// GetEndowmentBalance is a free data retrieval call binding the contract method 0x5acce36b.
//
// Solidity: function getEndowmentBalance() constant returns(uint256)
func (_Etclottery *EtclotterySession) GetEndowmentBalance() (*big.Int, error) {
	return _Etclottery.Contract.GetEndowmentBalance(&_Etclottery.CallOpts)
}

// GetEndowmentBalance is a free data retrieval call binding the contract method 0x5acce36b.
//
// Solidity: function getEndowmentBalance() constant returns(uint256)
func (_Etclottery *EtclotteryCallerSession) GetEndowmentBalance() (*big.Int, error) {
	return _Etclottery.Contract.GetEndowmentBalance(&_Etclottery.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() constant returns(uint256)
func (_Etclottery *EtclotteryCaller) GetFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Etclottery.contract.Call(opts, out, "getFee")
	return *ret0, err
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() constant returns(uint256)
func (_Etclottery *EtclotterySession) GetFee() (*big.Int, error) {
	return _Etclottery.Contract.GetFee(&_Etclottery.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() constant returns(uint256)
func (_Etclottery *EtclotteryCallerSession) GetFee() (*big.Int, error) {
	return _Etclottery.Contract.GetFee(&_Etclottery.CallOpts)
}

// GetLastRoundInfo is a free data retrieval call binding the contract method 0x705e798e.
//
// Solidity: function getLastRoundInfo() constant returns(uint256, uint256, uint256, uint256, bool, int256)
func (_Etclottery *EtclotteryCaller) GetLastRoundInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
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
	err := _Etclottery.contract.Call(opts, out, "getLastRoundInfo")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetLastRoundInfo is a free data retrieval call binding the contract method 0x705e798e.
//
// Solidity: function getLastRoundInfo() constant returns(uint256, uint256, uint256, uint256, bool, int256)
func (_Etclottery *EtclotterySession) GetLastRoundInfo() (*big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Etclottery.Contract.GetLastRoundInfo(&_Etclottery.CallOpts)
}

// GetLastRoundInfo is a free data retrieval call binding the contract method 0x705e798e.
//
// Solidity: function getLastRoundInfo() constant returns(uint256, uint256, uint256, uint256, bool, int256)
func (_Etclottery *EtclotteryCallerSession) GetLastRoundInfo() (*big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Etclottery.Contract.GetLastRoundInfo(&_Etclottery.CallOpts)
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(addr address) constant returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_Etclottery *EtclotteryCaller) GetPlayerInfoByAddress(opts *bind.CallOpts, addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
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
	err := _Etclottery.contract.Call(opts, out, "getPlayerInfoByAddress", addr)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(addr address) constant returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_Etclottery *EtclotterySession) GetPlayerInfoByAddress(addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Etclottery.Contract.GetPlayerInfoByAddress(&_Etclottery.CallOpts, addr)
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(addr address) constant returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_Etclottery *EtclotteryCallerSession) GetPlayerInfoByAddress(addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Etclottery.Contract.GetPlayerInfoByAddress(&_Etclottery.CallOpts, addr)
}

// GetTotalInfo is a free data retrieval call binding the contract method 0x285109d1.
//
// Solidity: function getTotalInfo() constant returns(uint256, uint256, uint256, uint256)
func (_Etclottery *EtclotteryCaller) GetTotalInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
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
	err := _Etclottery.contract.Call(opts, out, "getTotalInfo")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetTotalInfo is a free data retrieval call binding the contract method 0x285109d1.
//
// Solidity: function getTotalInfo() constant returns(uint256, uint256, uint256, uint256)
func (_Etclottery *EtclotterySession) GetTotalInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Etclottery.Contract.GetTotalInfo(&_Etclottery.CallOpts)
}

// GetTotalInfo is a free data retrieval call binding the contract method 0x285109d1.
//
// Solidity: function getTotalInfo() constant returns(uint256, uint256, uint256, uint256)
func (_Etclottery *EtclotteryCallerSession) GetTotalInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Etclottery.Contract.GetTotalInfo(&_Etclottery.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0x14107f3c.
//
// Solidity: function buy(_team uint8) returns()
func (_Etclottery *EtclotteryTransactor) Buy(opts *bind.TransactOpts, _team uint8) (*types.Transaction, error) {
	return _Etclottery.contract.Transact(opts, "buy", _team)
}

// Buy is a paid mutator transaction binding the contract method 0x14107f3c.
//
// Solidity: function buy(_team uint8) returns()
func (_Etclottery *EtclotterySession) Buy(_team uint8) (*types.Transaction, error) {
	return _Etclottery.Contract.Buy(&_Etclottery.TransactOpts, _team)
}

// Buy is a paid mutator transaction binding the contract method 0x14107f3c.
//
// Solidity: function buy(_team uint8) returns()
func (_Etclottery *EtclotteryTransactorSession) Buy(_team uint8) (*types.Transaction, error) {
	return _Etclottery.Contract.Buy(&_Etclottery.TransactOpts, _team)
}

// GetBlock is a paid mutator transaction binding the contract method 0x2e97766d.
//
// Solidity: function getBlock() returns()
func (_Etclottery *EtclotteryTransactor) GetBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etclottery.contract.Transact(opts, "getBlock")
}

// GetBlock is a paid mutator transaction binding the contract method 0x2e97766d.
//
// Solidity: function getBlock() returns()
func (_Etclottery *EtclotterySession) GetBlock() (*types.Transaction, error) {
	return _Etclottery.Contract.GetBlock(&_Etclottery.TransactOpts)
}

// GetBlock is a paid mutator transaction binding the contract method 0x2e97766d.
//
// Solidity: function getBlock() returns()
func (_Etclottery *EtclotteryTransactorSession) GetBlock() (*types.Transaction, error) {
	return _Etclottery.Contract.GetBlock(&_Etclottery.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Etclottery *EtclotteryTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etclottery.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Etclottery *EtclotterySession) Kill() (*types.Transaction, error) {
	return _Etclottery.Contract.Kill(&_Etclottery.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Etclottery *EtclotteryTransactorSession) Kill() (*types.Transaction, error) {
	return _Etclottery.Contract.Kill(&_Etclottery.TransactOpts)
}

// PlayerWithdraw is a paid mutator transaction binding the contract method 0x5629365b.
//
// Solidity: function playerWithdraw(amount uint256) returns()
func (_Etclottery *EtclotteryTransactor) PlayerWithdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Etclottery.contract.Transact(opts, "playerWithdraw", amount)
}

// PlayerWithdraw is a paid mutator transaction binding the contract method 0x5629365b.
//
// Solidity: function playerWithdraw(amount uint256) returns()
func (_Etclottery *EtclotterySession) PlayerWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _Etclottery.Contract.PlayerWithdraw(&_Etclottery.TransactOpts, amount)
}

// PlayerWithdraw is a paid mutator transaction binding the contract method 0x5629365b.
//
// Solidity: function playerWithdraw(amount uint256) returns()
func (_Etclottery *EtclotteryTransactorSession) PlayerWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _Etclottery.Contract.PlayerWithdraw(&_Etclottery.TransactOpts, amount)
}

// Reinvest is a paid mutator transaction binding the contract method 0xf0a3843d.
//
// Solidity: function reinvest(amount uint256, _team uint8) returns()
func (_Etclottery *EtclotteryTransactor) Reinvest(opts *bind.TransactOpts, amount *big.Int, _team uint8) (*types.Transaction, error) {
	return _Etclottery.contract.Transact(opts, "reinvest", amount, _team)
}

// Reinvest is a paid mutator transaction binding the contract method 0xf0a3843d.
//
// Solidity: function reinvest(amount uint256, _team uint8) returns()
func (_Etclottery *EtclotterySession) Reinvest(amount *big.Int, _team uint8) (*types.Transaction, error) {
	return _Etclottery.Contract.Reinvest(&_Etclottery.TransactOpts, amount, _team)
}

// Reinvest is a paid mutator transaction binding the contract method 0xf0a3843d.
//
// Solidity: function reinvest(amount uint256, _team uint8) returns()
func (_Etclottery *EtclotteryTransactorSession) Reinvest(amount *big.Int, _team uint8) (*types.Transaction, error) {
	return _Etclottery.Contract.Reinvest(&_Etclottery.TransactOpts, amount, _team)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(amount uint256) returns()
func (_Etclottery *EtclotteryTransactor) WithdrawFee(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Etclottery.contract.Transact(opts, "withdrawFee", amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(amount uint256) returns()
func (_Etclottery *EtclotterySession) WithdrawFee(amount *big.Int) (*types.Transaction, error) {
	return _Etclottery.Contract.WithdrawFee(&_Etclottery.TransactOpts, amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(amount uint256) returns()
func (_Etclottery *EtclotteryTransactorSession) WithdrawFee(amount *big.Int) (*types.Transaction, error) {
	return _Etclottery.Contract.WithdrawFee(&_Etclottery.TransactOpts, amount)
}

// EtclotteryOnBuyAndDistributeIterator is returned from FilterOnBuyAndDistribute and is used to iterate over the raw logs and unpacked data for OnBuyAndDistribute events raised by the Etclottery contract.
type EtclotteryOnBuyAndDistributeIterator struct {
	Event *EtclotteryOnBuyAndDistribute // Event containing the contract specifics and raw log

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
func (it *EtclotteryOnBuyAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtclotteryOnBuyAndDistribute)
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
		it.Event = new(EtclotteryOnBuyAndDistribute)
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
func (it *EtclotteryOnBuyAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtclotteryOnBuyAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtclotteryOnBuyAndDistribute represents a OnBuyAndDistribute event raised by the Etclottery contract.
type EtclotteryOnBuyAndDistribute struct {
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
// Solidity: event onBuyAndDistribute(rid uint256, strt uint256, end uint256, etc0 uint256, etc1 uint256, win int256)
func (_Etclottery *EtclotteryFilterer) FilterOnBuyAndDistribute(opts *bind.FilterOpts) (*EtclotteryOnBuyAndDistributeIterator, error) {

	logs, sub, err := _Etclottery.contract.FilterLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return &EtclotteryOnBuyAndDistributeIterator{contract: _Etclottery.contract, event: "onBuyAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnBuyAndDistribute is a free log subscription operation binding the contract event 0x755811ed4d5c4c398a7bfa8123f92c98e628b975243e845e607f1c649f9cf7ac.
//
// Solidity: event onBuyAndDistribute(rid uint256, strt uint256, end uint256, etc0 uint256, etc1 uint256, win int256)
func (_Etclottery *EtclotteryFilterer) WatchOnBuyAndDistribute(opts *bind.WatchOpts, sink chan<- *EtclotteryOnBuyAndDistribute) (event.Subscription, error) {

	logs, sub, err := _Etclottery.contract.WatchLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtclotteryOnBuyAndDistribute)
				if err := _Etclottery.contract.UnpackLog(event, "onBuyAndDistribute", log); err != nil {
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

// EtclotteryOnBuysIterator is returned from FilterOnBuys and is used to iterate over the raw logs and unpacked data for OnBuys events raised by the Etclottery contract.
type EtclotteryOnBuysIterator struct {
	Event *EtclotteryOnBuys // Event containing the contract specifics and raw log

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
func (it *EtclotteryOnBuysIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtclotteryOnBuys)
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
		it.Event = new(EtclotteryOnBuys)
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
func (it *EtclotteryOnBuysIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtclotteryOnBuysIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtclotteryOnBuys represents a OnBuys event raised by the Etclottery contract.
type EtclotteryOnBuys struct {
	Addr   common.Address
	Amount *big.Int
	Team   uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOnBuys is a free log retrieval operation binding the contract event 0xd3262b557966269a237416d28e9e625027d870723fa202fc25a2df9912b34d60.
//
// Solidity: event onBuys(addr address, amount uint256, _team uint8)
func (_Etclottery *EtclotteryFilterer) FilterOnBuys(opts *bind.FilterOpts) (*EtclotteryOnBuysIterator, error) {

	logs, sub, err := _Etclottery.contract.FilterLogs(opts, "onBuys")
	if err != nil {
		return nil, err
	}
	return &EtclotteryOnBuysIterator{contract: _Etclottery.contract, event: "onBuys", logs: logs, sub: sub}, nil
}

// WatchOnBuys is a free log subscription operation binding the contract event 0xd3262b557966269a237416d28e9e625027d870723fa202fc25a2df9912b34d60.
//
// Solidity: event onBuys(addr address, amount uint256, _team uint8)
func (_Etclottery *EtclotteryFilterer) WatchOnBuys(opts *bind.WatchOpts, sink chan<- *EtclotteryOnBuys) (event.Subscription, error) {

	logs, sub, err := _Etclottery.contract.WatchLogs(opts, "onBuys")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtclotteryOnBuys)
				if err := _Etclottery.contract.UnpackLog(event, "onBuys", log); err != nil {
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

// EtclotteryOnWithdrawIterator is returned from FilterOnWithdraw and is used to iterate over the raw logs and unpacked data for OnWithdraw events raised by the Etclottery contract.
type EtclotteryOnWithdrawIterator struct {
	Event *EtclotteryOnWithdraw // Event containing the contract specifics and raw log

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
func (it *EtclotteryOnWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtclotteryOnWithdraw)
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
		it.Event = new(EtclotteryOnWithdraw)
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
func (it *EtclotteryOnWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtclotteryOnWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtclotteryOnWithdraw represents a OnWithdraw event raised by the Etclottery contract.
type EtclotteryOnWithdraw struct {
	PlayerAddress common.Address
	Out           *big.Int
	TimeStamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOnWithdraw is a free log retrieval operation binding the contract event 0x1b091269e929df55d64d6ea7e9cadbe4fb38dce5ccdc995767bc515030dbfbbf.
//
// Solidity: event onWithdraw(playerAddress address, out uint256, timeStamp uint256)
func (_Etclottery *EtclotteryFilterer) FilterOnWithdraw(opts *bind.FilterOpts) (*EtclotteryOnWithdrawIterator, error) {

	logs, sub, err := _Etclottery.contract.FilterLogs(opts, "onWithdraw")
	if err != nil {
		return nil, err
	}
	return &EtclotteryOnWithdrawIterator{contract: _Etclottery.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0x1b091269e929df55d64d6ea7e9cadbe4fb38dce5ccdc995767bc515030dbfbbf.
//
// Solidity: event onWithdraw(playerAddress address, out uint256, timeStamp uint256)
func (_Etclottery *EtclotteryFilterer) WatchOnWithdraw(opts *bind.WatchOpts, sink chan<- *EtclotteryOnWithdraw) (event.Subscription, error) {

	logs, sub, err := _Etclottery.contract.WatchLogs(opts, "onWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtclotteryOnWithdraw)
				if err := _Etclottery.contract.UnpackLog(event, "onWithdraw", log); err != nil {
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
