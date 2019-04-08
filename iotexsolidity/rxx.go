// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iotexsolidity

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

// AttackerABI is the input ABI used to generate the binding from.
const AttackerABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"attack\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"daoAddress\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// AttackerBin is the compiled bytecode used for deploying new contracts.
const AttackerBin = `0x60806040526000805560405160208061026a833981016040818152915160028054600160a060020a031916600160a060020a0380841691909117918290553460018190557fd0e30db00000000000000000000000000000000000000000000000000000000085529451929491169263d0e30db092600480830192600092919082900301818588803b15801561009357600080fd5b505af11580156100a7573d6000803e3d6000fd5b5050505050506101ae806100bc6000396000f3006080604052600436106100405763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416639e5faafc81146100e4575b6000805460018101909155600a11156100e257600254600154604080517f2e1a7d4d00000000000000000000000000000000000000000000000000000000815260048101929092525173ffffffffffffffffffffffffffffffffffffffff90921691632e1a7d4d9160248082019260009290919082900301818387803b1580156100c957600080fd5b505af11580156100dd573d6000803e3d6000fd5b505050505b005b3480156100f057600080fd5b50600254600154604080517f2e1a7d4d0000000000000000000000000000000000000000000000000000000081526004810192909252516100e29273ffffffffffffffffffffffffffffffffffffffff1691632e1a7d4d91602480830192600092919082900301818387803b15801561016857600080fd5b505af115801561017c573d6000803e3d6000fd5b505050505600a165627a7a72305820fbdac2e242043ca628eb2389600eae825efe284e2bfb251a03b12a5a113570340029`

// DeployAttacker deploys a new Ethereum contract, binding an instance of Attacker to it.
func DeployAttacker(auth *bind.TransactOpts, backend bind.ContractBackend, daoAddress common.Address) (common.Address, *types.Transaction, *Attacker, error) {
	parsed, err := abi.JSON(strings.NewReader(AttackerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AttackerBin), backend, daoAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Attacker{AttackerCaller: AttackerCaller{contract: contract}, AttackerTransactor: AttackerTransactor{contract: contract}, AttackerFilterer: AttackerFilterer{contract: contract}}, nil
}

// Attacker is an auto generated Go binding around an Ethereum contract.
type Attacker struct {
	AttackerCaller     // Read-only binding to the contract
	AttackerTransactor // Write-only binding to the contract
	AttackerFilterer   // Log filterer for contract events
}

// AttackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttackerSession struct {
	Contract     *Attacker         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttackerCallerSession struct {
	Contract *AttackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AttackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttackerTransactorSession struct {
	Contract     *AttackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AttackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttackerRaw struct {
	Contract *Attacker // Generic contract binding to access the raw methods on
}

// AttackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttackerCallerRaw struct {
	Contract *AttackerCaller // Generic read-only contract binding to access the raw methods on
}

// AttackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttackerTransactorRaw struct {
	Contract *AttackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttacker creates a new instance of Attacker, bound to a specific deployed contract.
func NewAttacker(address common.Address, backend bind.ContractBackend) (*Attacker, error) {
	contract, err := bindAttacker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Attacker{AttackerCaller: AttackerCaller{contract: contract}, AttackerTransactor: AttackerTransactor{contract: contract}, AttackerFilterer: AttackerFilterer{contract: contract}}, nil
}

// NewAttackerCaller creates a new read-only instance of Attacker, bound to a specific deployed contract.
func NewAttackerCaller(address common.Address, caller bind.ContractCaller) (*AttackerCaller, error) {
	contract, err := bindAttacker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttackerCaller{contract: contract}, nil
}

// NewAttackerTransactor creates a new write-only instance of Attacker, bound to a specific deployed contract.
func NewAttackerTransactor(address common.Address, transactor bind.ContractTransactor) (*AttackerTransactor, error) {
	contract, err := bindAttacker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttackerTransactor{contract: contract}, nil
}

// NewAttackerFilterer creates a new log filterer instance of Attacker, bound to a specific deployed contract.
func NewAttackerFilterer(address common.Address, filterer bind.ContractFilterer) (*AttackerFilterer, error) {
	contract, err := bindAttacker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttackerFilterer{contract: contract}, nil
}

// bindAttacker binds a generic wrapper to an already deployed contract.
func bindAttacker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attacker *AttackerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Attacker.Contract.AttackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attacker *AttackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attacker.Contract.AttackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attacker *AttackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attacker.Contract.AttackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attacker *AttackerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Attacker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attacker *AttackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attacker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attacker *AttackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attacker.Contract.contract.Transact(opts, method, params...)
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_Attacker *AttackerTransactor) Attack(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attacker.contract.Transact(opts, "attack")
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_Attacker *AttackerSession) Attack() (*types.Transaction, error) {
	return _Attacker.Contract.Attack(&_Attacker.TransactOpts)
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_Attacker *AttackerTransactorSession) Attack() (*types.Transaction, error) {
	return _Attacker.Contract.Attack(&_Attacker.TransactOpts)
}

// MiniDAOABI is the input ABI used to generate the binding from.
const MiniDAOABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// MiniDAOBin is the compiled bytecode used for deploying new contracts.
const MiniDAOBin = `0x608060405234801561001057600080fd5b5060f68061001f6000396000f30060806040526004361060485763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632e1a7d4d8114604d578063d0e30db0146064575b600080fd5b348015605857600080fd5b506062600435606a565b005b606260b3565b33600090815260208190526040902054811115608557600080fd5b60405133908290600081818185875af150503360009081526020819052604090208054939093039092555050565b3360009081526020819052604090208054340190555600a165627a7a72305820e506c8c25ab2327072e450f7b4f65648426593c47d22cf14b09a31226e1b8b6b0029`

// DeployMiniDAO deploys a new Ethereum contract, binding an instance of MiniDAO to it.
func DeployMiniDAO(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MiniDAO, error) {
	parsed, err := abi.JSON(strings.NewReader(MiniDAOABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MiniDAOBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MiniDAO{MiniDAOCaller: MiniDAOCaller{contract: contract}, MiniDAOTransactor: MiniDAOTransactor{contract: contract}, MiniDAOFilterer: MiniDAOFilterer{contract: contract}}, nil
}

// MiniDAO is an auto generated Go binding around an Ethereum contract.
type MiniDAO struct {
	MiniDAOCaller     // Read-only binding to the contract
	MiniDAOTransactor // Write-only binding to the contract
	MiniDAOFilterer   // Log filterer for contract events
}

// MiniDAOCaller is an auto generated read-only Go binding around an Ethereum contract.
type MiniDAOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiniDAOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MiniDAOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiniDAOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MiniDAOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiniDAOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MiniDAOSession struct {
	Contract     *MiniDAO          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MiniDAOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MiniDAOCallerSession struct {
	Contract *MiniDAOCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MiniDAOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MiniDAOTransactorSession struct {
	Contract     *MiniDAOTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MiniDAORaw is an auto generated low-level Go binding around an Ethereum contract.
type MiniDAORaw struct {
	Contract *MiniDAO // Generic contract binding to access the raw methods on
}

// MiniDAOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MiniDAOCallerRaw struct {
	Contract *MiniDAOCaller // Generic read-only contract binding to access the raw methods on
}

// MiniDAOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MiniDAOTransactorRaw struct {
	Contract *MiniDAOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMiniDAO creates a new instance of MiniDAO, bound to a specific deployed contract.
func NewMiniDAO(address common.Address, backend bind.ContractBackend) (*MiniDAO, error) {
	contract, err := bindMiniDAO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MiniDAO{MiniDAOCaller: MiniDAOCaller{contract: contract}, MiniDAOTransactor: MiniDAOTransactor{contract: contract}, MiniDAOFilterer: MiniDAOFilterer{contract: contract}}, nil
}

// NewMiniDAOCaller creates a new read-only instance of MiniDAO, bound to a specific deployed contract.
func NewMiniDAOCaller(address common.Address, caller bind.ContractCaller) (*MiniDAOCaller, error) {
	contract, err := bindMiniDAO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MiniDAOCaller{contract: contract}, nil
}

// NewMiniDAOTransactor creates a new write-only instance of MiniDAO, bound to a specific deployed contract.
func NewMiniDAOTransactor(address common.Address, transactor bind.ContractTransactor) (*MiniDAOTransactor, error) {
	contract, err := bindMiniDAO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MiniDAOTransactor{contract: contract}, nil
}

// NewMiniDAOFilterer creates a new log filterer instance of MiniDAO, bound to a specific deployed contract.
func NewMiniDAOFilterer(address common.Address, filterer bind.ContractFilterer) (*MiniDAOFilterer, error) {
	contract, err := bindMiniDAO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MiniDAOFilterer{contract: contract}, nil
}

// bindMiniDAO binds a generic wrapper to an already deployed contract.
func bindMiniDAO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MiniDAOABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiniDAO *MiniDAORaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MiniDAO.Contract.MiniDAOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiniDAO *MiniDAORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiniDAO.Contract.MiniDAOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiniDAO *MiniDAORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiniDAO.Contract.MiniDAOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiniDAO *MiniDAOCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MiniDAO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiniDAO *MiniDAOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiniDAO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiniDAO *MiniDAOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiniDAO.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_MiniDAO *MiniDAOTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiniDAO.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_MiniDAO *MiniDAOSession) Deposit() (*types.Transaction, error) {
	return _MiniDAO.Contract.Deposit(&_MiniDAO.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_MiniDAO *MiniDAOTransactorSession) Deposit() (*types.Transaction, error) {
	return _MiniDAO.Contract.Deposit(&_MiniDAO.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_MiniDAO *MiniDAOTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MiniDAO.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_MiniDAO *MiniDAOSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _MiniDAO.Contract.Withdraw(&_MiniDAO.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_MiniDAO *MiniDAOTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _MiniDAO.Contract.Withdraw(&_MiniDAO.TransactOpts, amount)
}
