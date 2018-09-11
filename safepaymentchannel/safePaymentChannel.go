// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ReceiverPaysABI is the input ABI used to generate the binding from.
const ReceiverPaysABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"claimPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// ReceiverPaysBin is the compiled bytecode used for deploying new contracts.
const ReceiverPaysBin = `0x608060405260008054600160a060020a031916331790556103d2806100256000396000f30060806040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166341c0e1b58114610050578063a90ae88714610067575b600080fd5b34801561005c57600080fd5b506100656100c7565b005b34801561007357600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526100659482359460248035953695946064949201919081908401838280828437509497506100ee9650505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146100eb57600080fd5b33ff5b60008281526001602052604081205460ff161561010a57600080fd5b600083815260016020818152604092839020805460ff191690921790915581516c01000000000000000000000000338102828401526034820188905260548201879052300260748201528251606881830301815260889091019283905280516101ca9391928291908401908083835b602083106101985780518252601f199092019160209182019101610179565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020610244565b60005490915073ffffffffffffffffffffffffffffffffffffffff166101f082846102ee565b73ffffffffffffffffffffffffffffffffffffffff161461021057600080fd5b604051339085156108fc029086906000818181858888f1935050505015801561023d573d6000803e3d6000fd5b5050505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c80830185905283518084039091018152605c909201928390528151600093918291908401908083835b602083106102bc5780518252601f19909201916020918201910161029d565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912095945050505050565b6000806000806102fd85610375565b60408051600080825260208083018085528d905260ff8716838501526060830186905260808301859052925195985093965091945060019360a0808401949293601f19830193908390039091019190865af1158015610360573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60008060008351604114151561038a57600080fd5b5050506020810151604082015160609092015160001a929091905600a165627a7a723058204d7c84438a6c706fcb835e715d0e88f1853b12d4fd3a7484095ea7ae0b9aff020029`

// DeployReceiverPays deploys a new Ethereum contract, binding an instance of ReceiverPays to it.
func DeployReceiverPays(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReceiverPays, error) {
	parsed, err := abi.JSON(strings.NewReader(ReceiverPaysABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReceiverPaysBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReceiverPays{ReceiverPaysCaller: ReceiverPaysCaller{contract: contract}, ReceiverPaysTransactor: ReceiverPaysTransactor{contract: contract}, ReceiverPaysFilterer: ReceiverPaysFilterer{contract: contract}}, nil
}

// ReceiverPays is an auto generated Go binding around an Ethereum contract.
type ReceiverPays struct {
	ReceiverPaysCaller     // Read-only binding to the contract
	ReceiverPaysTransactor // Write-only binding to the contract
	ReceiverPaysFilterer   // Log filterer for contract events
}

// ReceiverPaysCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiverPaysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverPaysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiverPaysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverPaysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiverPaysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverPaysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiverPaysSession struct {
	Contract     *ReceiverPays     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiverPaysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiverPaysCallerSession struct {
	Contract *ReceiverPaysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ReceiverPaysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiverPaysTransactorSession struct {
	Contract     *ReceiverPaysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ReceiverPaysRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiverPaysRaw struct {
	Contract *ReceiverPays // Generic contract binding to access the raw methods on
}

// ReceiverPaysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiverPaysCallerRaw struct {
	Contract *ReceiverPaysCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiverPaysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiverPaysTransactorRaw struct {
	Contract *ReceiverPaysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiverPays creates a new instance of ReceiverPays, bound to a specific deployed contract.
func NewReceiverPays(address common.Address, backend bind.ContractBackend) (*ReceiverPays, error) {
	contract, err := bindReceiverPays(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiverPays{ReceiverPaysCaller: ReceiverPaysCaller{contract: contract}, ReceiverPaysTransactor: ReceiverPaysTransactor{contract: contract}, ReceiverPaysFilterer: ReceiverPaysFilterer{contract: contract}}, nil
}

// NewReceiverPaysCaller creates a new read-only instance of ReceiverPays, bound to a specific deployed contract.
func NewReceiverPaysCaller(address common.Address, caller bind.ContractCaller) (*ReceiverPaysCaller, error) {
	contract, err := bindReceiverPays(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverPaysCaller{contract: contract}, nil
}

// NewReceiverPaysTransactor creates a new write-only instance of ReceiverPays, bound to a specific deployed contract.
func NewReceiverPaysTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiverPaysTransactor, error) {
	contract, err := bindReceiverPays(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverPaysTransactor{contract: contract}, nil
}

// NewReceiverPaysFilterer creates a new log filterer instance of ReceiverPays, bound to a specific deployed contract.
func NewReceiverPaysFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiverPaysFilterer, error) {
	contract, err := bindReceiverPays(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiverPaysFilterer{contract: contract}, nil
}

// bindReceiverPays binds a generic wrapper to an already deployed contract.
func bindReceiverPays(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReceiverPaysABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiverPays *ReceiverPaysRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReceiverPays.Contract.ReceiverPaysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiverPays *ReceiverPaysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiverPays.Contract.ReceiverPaysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiverPays *ReceiverPaysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiverPays.Contract.ReceiverPaysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiverPays *ReceiverPaysCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReceiverPays.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiverPays *ReceiverPaysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiverPays.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiverPays *ReceiverPaysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiverPays.Contract.contract.Transact(opts, method, params...)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0xa90ae887.
//
// Solidity: function claimPayment(amount uint256, nonce uint256, signature bytes) returns()
func (_ReceiverPays *ReceiverPaysTransactor) ClaimPayment(opts *bind.TransactOpts, amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _ReceiverPays.contract.Transact(opts, "claimPayment", amount, nonce, signature)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0xa90ae887.
//
// Solidity: function claimPayment(amount uint256, nonce uint256, signature bytes) returns()
func (_ReceiverPays *ReceiverPaysSession) ClaimPayment(amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _ReceiverPays.Contract.ClaimPayment(&_ReceiverPays.TransactOpts, amount, nonce, signature)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0xa90ae887.
//
// Solidity: function claimPayment(amount uint256, nonce uint256, signature bytes) returns()
func (_ReceiverPays *ReceiverPaysTransactorSession) ClaimPayment(amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _ReceiverPays.Contract.ClaimPayment(&_ReceiverPays.TransactOpts, amount, nonce, signature)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ReceiverPays *ReceiverPaysTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiverPays.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ReceiverPays *ReceiverPaysSession) Kill() (*types.Transaction, error) {
	return _ReceiverPays.Contract.Kill(&_ReceiverPays.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ReceiverPays *ReceiverPaysTransactorSession) Kill() (*types.Transaction, error) {
	return _ReceiverPays.Contract.Kill(&_ReceiverPays.TransactOpts)
}
