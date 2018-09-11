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
const ReceiverPaysABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEndowmentBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"claimPayment\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// ReceiverPaysBin is the compiled bytecode used for deploying new contracts.
const ReceiverPaysBin = `0x608060405260008054600160a060020a0319163317905534801561002257600080fd5b5061048d806100326000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166341c0e1b581146100585780635acce36b1461006d578063a90ae88714610094575b005b34801561006457600080fd5b50610056610173565b34801561007957600080fd5b5061008261019a565b60408051918252519081900360200190f35b3480156100a057600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526100f494823594602480359536959460649492019190819084018382808284375094975061019f9650505050505050565b6040805183815260208082018381528451938301939093528351919291606084019185019080838360005b8381101561013757818101518382015260200161011f565b50505050905090810190601f1680156101645780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b60005473ffffffffffffffffffffffffffffffffffffffff16331461019757600080fd5b33ff5b303190565b600082815260016020526040812054606090829060ff16156101c057600080fd5b600085815260016020818152604092839020805460ff191690921790915581516c0100000000000000000000000033810282840152603482018a905260548201899052300260748201528251606881830301815260889091019283905280516102809391928291908401908083835b6020831061024e5780518252601f19909201916020918201910161022f565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206102ff565b60005490915073ffffffffffffffffffffffffffffffffffffffff166102a682866103a9565b73ffffffffffffffffffffffffffffffffffffffff16146102c657600080fd5b604051339087156108fc029088906000818181858888f193505050501580156102f3573d6000803e3d6000fd5b50959294509192505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c80830185905283518084039091018152605c909201928390528151600093918291908401908083835b602083106103775780518252601f199092019160209182019101610358565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912095945050505050565b6000806000806103b885610430565b60408051600080825260208083018085528d905260ff8716838501526060830186905260808301859052925195985093965091945060019360a0808401949293601f19830193908390039091019190865af115801561041b573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60008060008351604114151561044557600080fd5b5050506020810151604082015160609092015160001a929091905600a165627a7a7230582073a3d81be792ef5e55749a38849a892318a67a3cde4050b2cc19f491ed12520b0029`

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

// GetEndowmentBalance is a free data retrieval call binding the contract method 0x5acce36b.
//
// Solidity: function getEndowmentBalance() constant returns(uint256)
func (_ReceiverPays *ReceiverPaysCaller) GetEndowmentBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReceiverPays.contract.Call(opts, out, "getEndowmentBalance")
	return *ret0, err
}

// GetEndowmentBalance is a free data retrieval call binding the contract method 0x5acce36b.
//
// Solidity: function getEndowmentBalance() constant returns(uint256)
func (_ReceiverPays *ReceiverPaysSession) GetEndowmentBalance() (*big.Int, error) {
	return _ReceiverPays.Contract.GetEndowmentBalance(&_ReceiverPays.CallOpts)
}

// GetEndowmentBalance is a free data retrieval call binding the contract method 0x5acce36b.
//
// Solidity: function getEndowmentBalance() constant returns(uint256)
func (_ReceiverPays *ReceiverPaysCallerSession) GetEndowmentBalance() (*big.Int, error) {
	return _ReceiverPays.Contract.GetEndowmentBalance(&_ReceiverPays.CallOpts)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0xa90ae887.
//
// Solidity: function claimPayment(amount uint256, nonce uint256, signature bytes) returns(bytes32, bytes)
func (_ReceiverPays *ReceiverPaysTransactor) ClaimPayment(opts *bind.TransactOpts, amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _ReceiverPays.contract.Transact(opts, "claimPayment", amount, nonce, signature)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0xa90ae887.
//
// Solidity: function claimPayment(amount uint256, nonce uint256, signature bytes) returns(bytes32, bytes)
func (_ReceiverPays *ReceiverPaysSession) ClaimPayment(amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _ReceiverPays.Contract.ClaimPayment(&_ReceiverPays.TransactOpts, amount, nonce, signature)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0xa90ae887.
//
// Solidity: function claimPayment(amount uint256, nonce uint256, signature bytes) returns(bytes32, bytes)
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
