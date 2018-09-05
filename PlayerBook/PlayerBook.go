// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PlayerBook

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

// MSFunABI is the input ABI used to generate the binding from.
const MSFunABI = "[]"

// MSFunBin is the compiled bytecode used for deploying new contracts.
const MSFunBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206faef2894e67db21632231e53a460cd6c6624ade52c9d847d76990d64b726bbd0029`

// DeployMSFun deploys a new Ethereum contract, binding an instance of MSFun to it.
func DeployMSFun(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MSFun, error) {
	parsed, err := abi.JSON(strings.NewReader(MSFunABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MSFunBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MSFun{MSFunCaller: MSFunCaller{contract: contract}, MSFunTransactor: MSFunTransactor{contract: contract}, MSFunFilterer: MSFunFilterer{contract: contract}}, nil
}

// MSFun is an auto generated Go binding around an Ethereum contract.
type MSFun struct {
	MSFunCaller     // Read-only binding to the contract
	MSFunTransactor // Write-only binding to the contract
	MSFunFilterer   // Log filterer for contract events
}

// MSFunCaller is an auto generated read-only Go binding around an Ethereum contract.
type MSFunCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MSFunTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MSFunTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MSFunFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MSFunFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MSFunSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MSFunSession struct {
	Contract     *MSFun            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MSFunCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MSFunCallerSession struct {
	Contract *MSFunCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MSFunTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MSFunTransactorSession struct {
	Contract     *MSFunTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MSFunRaw is an auto generated low-level Go binding around an Ethereum contract.
type MSFunRaw struct {
	Contract *MSFun // Generic contract binding to access the raw methods on
}

// MSFunCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MSFunCallerRaw struct {
	Contract *MSFunCaller // Generic read-only contract binding to access the raw methods on
}

// MSFunTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MSFunTransactorRaw struct {
	Contract *MSFunTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMSFun creates a new instance of MSFun, bound to a specific deployed contract.
func NewMSFun(address common.Address, backend bind.ContractBackend) (*MSFun, error) {
	contract, err := bindMSFun(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MSFun{MSFunCaller: MSFunCaller{contract: contract}, MSFunTransactor: MSFunTransactor{contract: contract}, MSFunFilterer: MSFunFilterer{contract: contract}}, nil
}

// NewMSFunCaller creates a new read-only instance of MSFun, bound to a specific deployed contract.
func NewMSFunCaller(address common.Address, caller bind.ContractCaller) (*MSFunCaller, error) {
	contract, err := bindMSFun(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MSFunCaller{contract: contract}, nil
}

// NewMSFunTransactor creates a new write-only instance of MSFun, bound to a specific deployed contract.
func NewMSFunTransactor(address common.Address, transactor bind.ContractTransactor) (*MSFunTransactor, error) {
	contract, err := bindMSFun(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MSFunTransactor{contract: contract}, nil
}

// NewMSFunFilterer creates a new log filterer instance of MSFun, bound to a specific deployed contract.
func NewMSFunFilterer(address common.Address, filterer bind.ContractFilterer) (*MSFunFilterer, error) {
	contract, err := bindMSFun(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MSFunFilterer{contract: contract}, nil
}

// bindMSFun binds a generic wrapper to an already deployed contract.
func bindMSFun(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MSFunABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MSFun *MSFunRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MSFun.Contract.MSFunCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MSFun *MSFunRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MSFun.Contract.MSFunTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MSFun *MSFunRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MSFun.Contract.MSFunTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MSFun *MSFunCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MSFun.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MSFun *MSFunTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MSFun.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MSFun *MSFunTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MSFun.Contract.contract.Transact(opts, method, params...)
}

// MagicPowerDevsInterfaceABI is the input ABI used to generate the binding from.
const MagicPowerDevsInterfaceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"isDev\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"adminCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requiredSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"adminName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"devCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requiredDevSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MagicPowerDevsInterfaceBin is the compiled bytecode used for deploying new contracts.
const MagicPowerDevsInterfaceBin = `0x`

// DeployMagicPowerDevsInterface deploys a new Ethereum contract, binding an instance of MagicPowerDevsInterface to it.
func DeployMagicPowerDevsInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MagicPowerDevsInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(MagicPowerDevsInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MagicPowerDevsInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MagicPowerDevsInterface{MagicPowerDevsInterfaceCaller: MagicPowerDevsInterfaceCaller{contract: contract}, MagicPowerDevsInterfaceTransactor: MagicPowerDevsInterfaceTransactor{contract: contract}, MagicPowerDevsInterfaceFilterer: MagicPowerDevsInterfaceFilterer{contract: contract}}, nil
}

// MagicPowerDevsInterface is an auto generated Go binding around an Ethereum contract.
type MagicPowerDevsInterface struct {
	MagicPowerDevsInterfaceCaller     // Read-only binding to the contract
	MagicPowerDevsInterfaceTransactor // Write-only binding to the contract
	MagicPowerDevsInterfaceFilterer   // Log filterer for contract events
}

// MagicPowerDevsInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MagicPowerDevsInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagicPowerDevsInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MagicPowerDevsInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagicPowerDevsInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MagicPowerDevsInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagicPowerDevsInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MagicPowerDevsInterfaceSession struct {
	Contract     *MagicPowerDevsInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MagicPowerDevsInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MagicPowerDevsInterfaceCallerSession struct {
	Contract *MagicPowerDevsInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// MagicPowerDevsInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MagicPowerDevsInterfaceTransactorSession struct {
	Contract     *MagicPowerDevsInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// MagicPowerDevsInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MagicPowerDevsInterfaceRaw struct {
	Contract *MagicPowerDevsInterface // Generic contract binding to access the raw methods on
}

// MagicPowerDevsInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MagicPowerDevsInterfaceCallerRaw struct {
	Contract *MagicPowerDevsInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// MagicPowerDevsInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MagicPowerDevsInterfaceTransactorRaw struct {
	Contract *MagicPowerDevsInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMagicPowerDevsInterface creates a new instance of MagicPowerDevsInterface, bound to a specific deployed contract.
func NewMagicPowerDevsInterface(address common.Address, backend bind.ContractBackend) (*MagicPowerDevsInterface, error) {
	contract, err := bindMagicPowerDevsInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MagicPowerDevsInterface{MagicPowerDevsInterfaceCaller: MagicPowerDevsInterfaceCaller{contract: contract}, MagicPowerDevsInterfaceTransactor: MagicPowerDevsInterfaceTransactor{contract: contract}, MagicPowerDevsInterfaceFilterer: MagicPowerDevsInterfaceFilterer{contract: contract}}, nil
}

// NewMagicPowerDevsInterfaceCaller creates a new read-only instance of MagicPowerDevsInterface, bound to a specific deployed contract.
func NewMagicPowerDevsInterfaceCaller(address common.Address, caller bind.ContractCaller) (*MagicPowerDevsInterfaceCaller, error) {
	contract, err := bindMagicPowerDevsInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MagicPowerDevsInterfaceCaller{contract: contract}, nil
}

// NewMagicPowerDevsInterfaceTransactor creates a new write-only instance of MagicPowerDevsInterface, bound to a specific deployed contract.
func NewMagicPowerDevsInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MagicPowerDevsInterfaceTransactor, error) {
	contract, err := bindMagicPowerDevsInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MagicPowerDevsInterfaceTransactor{contract: contract}, nil
}

// NewMagicPowerDevsInterfaceFilterer creates a new log filterer instance of MagicPowerDevsInterface, bound to a specific deployed contract.
func NewMagicPowerDevsInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MagicPowerDevsInterfaceFilterer, error) {
	contract, err := bindMagicPowerDevsInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MagicPowerDevsInterfaceFilterer{contract: contract}, nil
}

// bindMagicPowerDevsInterface binds a generic wrapper to an already deployed contract.
func bindMagicPowerDevsInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MagicPowerDevsInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MagicPowerDevsInterface.Contract.MagicPowerDevsInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagicPowerDevsInterface.Contract.MagicPowerDevsInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MagicPowerDevsInterface.Contract.MagicPowerDevsInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MagicPowerDevsInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagicPowerDevsInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MagicPowerDevsInterface.Contract.contract.Transact(opts, method, params...)
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() constant returns(uint256)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCaller) AdminCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MagicPowerDevsInterface.contract.Call(opts, out, "adminCount")
	return *ret0, err
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() constant returns(uint256)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceSession) AdminCount() (*big.Int, error) {
	return _MagicPowerDevsInterface.Contract.AdminCount(&_MagicPowerDevsInterface.CallOpts)
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() constant returns(uint256)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCallerSession) AdminCount() (*big.Int, error) {
	return _MagicPowerDevsInterface.Contract.AdminCount(&_MagicPowerDevsInterface.CallOpts)
}

// AdminName is a free data retrieval call binding the contract method 0xaf1c084d.
//
// Solidity: function adminName(_who address) constant returns(bytes32)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCaller) AdminName(opts *bind.CallOpts, _who common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MagicPowerDevsInterface.contract.Call(opts, out, "adminName", _who)
	return *ret0, err
}

// AdminName is a free data retrieval call binding the contract method 0xaf1c084d.
//
// Solidity: function adminName(_who address) constant returns(bytes32)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceSession) AdminName(_who common.Address) ([32]byte, error) {
	return _MagicPowerDevsInterface.Contract.AdminName(&_MagicPowerDevsInterface.CallOpts, _who)
}

// AdminName is a free data retrieval call binding the contract method 0xaf1c084d.
//
// Solidity: function adminName(_who address) constant returns(bytes32)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCallerSession) AdminName(_who common.Address) ([32]byte, error) {
	return _MagicPowerDevsInterface.Contract.AdminName(&_MagicPowerDevsInterface.CallOpts, _who)
}

// DevCount is a free data retrieval call binding the contract method 0xcebc141a.
//
// Solidity: function devCount() constant returns(uint256)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCaller) DevCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MagicPowerDevsInterface.contract.Call(opts, out, "devCount")
	return *ret0, err
}

// DevCount is a free data retrieval call binding the contract method 0xcebc141a.
//
// Solidity: function devCount() constant returns(uint256)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceSession) DevCount() (*big.Int, error) {
	return _MagicPowerDevsInterface.Contract.DevCount(&_MagicPowerDevsInterface.CallOpts)
}

// DevCount is a free data retrieval call binding the contract method 0xcebc141a.
//
// Solidity: function devCount() constant returns(uint256)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCallerSession) DevCount() (*big.Int, error) {
	return _MagicPowerDevsInterface.Contract.DevCount(&_MagicPowerDevsInterface.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(_who address) constant returns(bool)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCaller) IsAdmin(opts *bind.CallOpts, _who common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MagicPowerDevsInterface.contract.Call(opts, out, "isAdmin", _who)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(_who address) constant returns(bool)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceSession) IsAdmin(_who common.Address) (bool, error) {
	return _MagicPowerDevsInterface.Contract.IsAdmin(&_MagicPowerDevsInterface.CallOpts, _who)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(_who address) constant returns(bool)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCallerSession) IsAdmin(_who common.Address) (bool, error) {
	return _MagicPowerDevsInterface.Contract.IsAdmin(&_MagicPowerDevsInterface.CallOpts, _who)
}

// IsDev is a free data retrieval call binding the contract method 0x0c3f64bf.
//
// Solidity: function isDev(_who address) constant returns(bool)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCaller) IsDev(opts *bind.CallOpts, _who common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MagicPowerDevsInterface.contract.Call(opts, out, "isDev", _who)
	return *ret0, err
}

// IsDev is a free data retrieval call binding the contract method 0x0c3f64bf.
//
// Solidity: function isDev(_who address) constant returns(bool)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceSession) IsDev(_who common.Address) (bool, error) {
	return _MagicPowerDevsInterface.Contract.IsDev(&_MagicPowerDevsInterface.CallOpts, _who)
}

// IsDev is a free data retrieval call binding the contract method 0x0c3f64bf.
//
// Solidity: function isDev(_who address) constant returns(bool)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCallerSession) IsDev(_who common.Address) (bool, error) {
	return _MagicPowerDevsInterface.Contract.IsDev(&_MagicPowerDevsInterface.CallOpts, _who)
}

// RequiredDevSignatures is a free data retrieval call binding the contract method 0xfcf2f85f.
//
// Solidity: function requiredDevSignatures() constant returns(uint256)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCaller) RequiredDevSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MagicPowerDevsInterface.contract.Call(opts, out, "requiredDevSignatures")
	return *ret0, err
}

// RequiredDevSignatures is a free data retrieval call binding the contract method 0xfcf2f85f.
//
// Solidity: function requiredDevSignatures() constant returns(uint256)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceSession) RequiredDevSignatures() (*big.Int, error) {
	return _MagicPowerDevsInterface.Contract.RequiredDevSignatures(&_MagicPowerDevsInterface.CallOpts)
}

// RequiredDevSignatures is a free data retrieval call binding the contract method 0xfcf2f85f.
//
// Solidity: function requiredDevSignatures() constant returns(uint256)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCallerSession) RequiredDevSignatures() (*big.Int, error) {
	return _MagicPowerDevsInterface.Contract.RequiredDevSignatures(&_MagicPowerDevsInterface.CallOpts)
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCaller) RequiredSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MagicPowerDevsInterface.contract.Call(opts, out, "requiredSignatures")
	return *ret0, err
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceSession) RequiredSignatures() (*big.Int, error) {
	return _MagicPowerDevsInterface.Contract.RequiredSignatures(&_MagicPowerDevsInterface.CallOpts)
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_MagicPowerDevsInterface *MagicPowerDevsInterfaceCallerSession) RequiredSignatures() (*big.Int, error) {
	return _MagicPowerDevsInterface.Contract.RequiredSignatures(&_MagicPowerDevsInterface.CallOpts)
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
const NameFilterBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206987648e5fb35952b4eec3ab938c1bf9f5e59d186c68c51a07c30ec19c7100790029`

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

// PlayerBookABI is the input ABI used to generate the binding from.
const PlayerBookABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"addMeToAllGames\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_whatFunction\",\"type\":\"bytes32\"}],\"name\":\"deleteAnyProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"pIDxAddr_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registrationFee_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNameFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"plyrNames_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"gameNames_\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pIDxName_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nameString\",\"type\":\"string\"},{\"name\":\"_affCode\",\"type\":\"address\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXaddr\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gameAddress\",\"type\":\"address\"},{\"name\":\"_gameNameStr\",\"type\":\"string\"}],\"name\":\"addGame\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pID_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"}],\"name\":\"getPlayerAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nameString\",\"type\":\"string\"},{\"name\":\"_affCode\",\"type\":\"bytes32\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXname\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nameStr\",\"type\":\"string\"}],\"name\":\"checkIfNameValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_affCode\",\"type\":\"bytes32\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXnameFromDapp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_whatFunction\",\"type\":\"bytes32\"},{\"name\":\"_signerA\",\"type\":\"uint256\"},{\"name\":\"_signerB\",\"type\":\"uint256\"},{\"name\":\"_signerC\",\"type\":\"uint256\"}],\"name\":\"checkSignersByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gameID\",\"type\":\"uint256\"}],\"name\":\"addMeToGame\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"}],\"name\":\"getPlayerName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nameString\",\"type\":\"string\"},{\"name\":\"_affCode\",\"type\":\"uint256\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXID\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plyrNameList_\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_whatFunction\",\"type\":\"bytes32\"}],\"name\":\"checkData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_affCode\",\"type\":\"address\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXaddrFromDapp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nameString\",\"type\":\"string\"}],\"name\":\"useMyOldName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gID_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_affCode\",\"type\":\"uint256\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXIDFromDapp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setRegistrationFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"games_\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"gameIDs_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plyr_\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"laff\",\"type\":\"uint256\"},{\"name\":\"names\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"}],\"name\":\"getPlayerLAff\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getPlayerID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_whatFunction\",\"type\":\"bytes32\"},{\"name\":\"_signerA\",\"type\":\"uint256\"},{\"name\":\"_signerB\",\"type\":\"uint256\"},{\"name\":\"_signerC\",\"type\":\"uint256\"}],\"name\":\"checkSignersByName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isNewPlayer\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onNewName\",\"type\":\"event\"}]"

// PlayerBookBin is the compiled bytecode used for deploying new contracts.
const PlayerBookBin = `0x60806040526706f05b59d3b2000060015534801561001c57600080fd5b507f92e85d02570a8092d09a6e3a57665bc3815a2699a4074001bf1ccabf660f5a368054600160a060020a0319908116737d6dce6ef707e78131211a45234dd8ac7d09a033179091557f62726164000000000000000000000000000000000000000000000000000000007f92e85d02570a8092d09a6e3a57665bc3815a2699a4074001bf1ccabf660f5a3781905560017f92e85d02570a8092d09a6e3a57665bc3815a2699a4074001bf1ccabf660f5a398190557f7eadfce641e63468d03873b22af4bd98d969c78709219afc022fb2f8ea94fd168190557fd7c4f62a8a635c3262b4da55c90ddafb67a5638de96459f6d5f314e710ae650f8190557f7de85bc73ec2afb9f14a09dedced4fce3e0fdb26019e6dcf86383c340b1726f8805460ff1990811683179091557f6bd9ce7c7c44b510ee825857e9b7d0dbc77ffc0dcbb487d9763a092e420c9a18929092557f6cde3cea4b3a3fb2488b2808bae7556f4a405e50f65e1794383bc026131b13c380548416735a3bae6f3e6704334d6713eb0290dea11730c2381790557f7a616368617279000000000000000000000000000000000000000000000000007f6cde3cea4b3a3fb2488b2808bae7556f4a405e50f65e1794383bc026131b13c48190557f6cde3cea4b3a3fb2488b2808bae7556f4a405e50f65e1794383bc026131b13c682905560027ff2f601776e6b8fa32893180d9161ee447777059aa60710ff93cbf45090d8a3b78190557f83153f74e0fe417fc226321d5f77503a98b6369156a2d5d7924444c02925286a557f763ae76480ce8990664f34c22a169e6b4bc1feafe6ebf5857f959cc5a28340ef80548416831790557f5ec873723cdc64ab6619ea2fcbca33beef9a820510f97aa9ecb9e3ff0f638853557fc575c31fea594a6eb97c8e9d3f9caee4c16218c6ef37e923234c0fe9014a61e78054841673125e0fa6d73dcd198081752295aac83867fc36e21790557f62727563650000000000000000000000000000000000000000000000000000007fc575c31fea594a6eb97c8e9d3f9caee4c16218c6ef37e923234c0fe9014a61e88190557fc575c31fea594a6eb97c8e9d3f9caee4c16218c6ef37e923234c0fe9014a61ea82905560037f59d2ad8351a3b718a36a9e3026854bc38fe60231a068af47d498ec48d8b730fc8190557f09f2fcec8a9f030c1bb715fc2f310dcd14abdaa7ae22f12c155b651f01516240557f618016b21c1f4e30ef280123d5595957e17de79c6c8e8dc0927f8c76296d15d980548416831790557f960add1d2a4af75b2cd16bfb54d92ee5d4c29427a1bf605fa9b667906b27d701557f8dc18c4ccfd75f5c815b63770fa542fd953e8fef7e0e44bbdd4913470ce7e9cb80549093167335f0a9270393b24c826d37914948ec12ec51ffb3179092557f70657465720000000000000000000000000000000000000000000000000000007f8dc18c4ccfd75f5c815b63770fa542fd953e8fef7e0e44bbdd4913470ce7e9cc8190557f8dc18c4ccfd75f5c815b63770fa542fd953e8fef7e0e44bbdd4913470ce7e9ce83905560047fb93738e4eb83cbb70b3b97e139dd8457a07e7f7d6fafa7d77372e3e787b717aa8190557fb1cbaf80101f739e52711eb38805bf3f134cfdc1c37abd02355d183fb09652a88190557f57bfbbc01617fdd3cc632290d91976bf432b89e2f9a58d68ab3732b683c5dde5805490931684179092556000929092527f12d0c11577e2f0950f57c455c117796550b79f444811db8ba2f69c57b646c7846020527f2c4ff72984b96d6e24887abcd47b1c2fb2b1bd600b7f7d037d1a2ac35568baab919091556006556133008061055b6000396000f30060806040526004361061017c5763ffffffff60e060020a6000350416630c6940ea81146101815780630efcf2951461019857806310f01eba146101b0578063180603eb146101e35780632614195f146101f85780632660316e1461020d57806327249e611461023c5780632e19ebdc1461025d5780633ddd4698146102755780633fda926e146102d15780634b227176146103385780634d0d35ff1461034d578063685ffd83146103815780636c52660d146103d4578063745ea0c11461042d578063768f930d1461046757806381c5b206146104b357806382e37b2c146104cb578063921dec21146104e3578063a448ed4b14610536578063a553506e14610551578063aa4d490b14610582578063b9291296146105a5578063b9eca0c8146105fe578063c0942dfd14610613578063c320c72714610632578063d52412791461064a578063dbbcaa9714610662578063de7874f314610683578063e3c08adf146106cb578063e56556a9146106e3578063ed3643d614610704575b600080fd5b34801561018d57600080fd5b50610196610743565b005b3480156101a457600080fd5b506101966004356109b1565b3480156101bc57600080fd5b506101d1600160a060020a0360043516610a88565b60408051918252519081900360200190f35b3480156101ef57600080fd5b506101d1610a9a565b34801561020457600080fd5b506101d1610aa0565b34801561021957600080fd5b50610228600435602435610aa6565b604080519115158252519081900360200190f35b34801561024857600080fd5b506101d1600160a060020a0360043516610ac6565b34801561026957600080fd5b506101d1600435610ad8565b6040805160206004803580820135601f810184900484028501840190955284845261019694369492936024939284019190819084018382808284375094975050600160a060020a03853516955050505050602001351515610aea565b3480156102dd57600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610196958335600160a060020a0316953695604494919390910191908190840183828082843750949750610c549650505050505050565b34801561034457600080fd5b506101d16111c0565b34801561035957600080fd5b506103656004356111c6565b60408051600160a060020a039092168252519081900360200190f35b6040805160206004803580820135601f81018490048402850184019095528484526101969436949293602493928401919081908401838280828437509497505084359550505050506020013515156111e4565b3480156103e057600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526102289436949293602493928401919081908401838280828437509497506113229650505050505050565b61044c600160a060020a0360043516602435604435606435151561135a565b60408051921515835260208301919091528051918290030190f35b34801561047357600080fd5b5061048860043560243560443560643561146e565b60408051600160a060020a039485168152928416602084015292168183015290519081900360600190f35b3480156104bf57600080fd5b50610196600435611572565b3480156104d757600080fd5b506101d160043561183f565b6040805160206004803580820135601f8101849004840285018401909552848452610196943694929360249392840191908190840183828082843750949750508435955050505050602001351515611854565b34801561054257600080fd5b506101d1600435602435611996565b34801561055d57600080fd5b506105696004356119b3565b6040805192835260208301919091528051918290030190f35b61044c600160a060020a0360043581169060243590604435166064351515611a9f565b3480156105b157600080fd5b506040805160206004803580820135601f8101849004840285018401909552848452610196943694929360249392840191908190840183828082843750949750611bc29650505050505050565b34801561060a57600080fd5b506101d1611cb4565b61044c600160a060020a03600435166024356044356064351515611cba565b34801561063e57600080fd5b50610196600435611dc6565b34801561065657600080fd5b50610365600435611ef0565b34801561066e57600080fd5b506101d1600160a060020a0360043516611f0b565b34801561068f57600080fd5b5061069b600435611f1d565b60408051600160a060020a0390951685526020850193909352838301919091526060830152519081900360800190f35b3480156106d757600080fd5b506101d1600435611f4e565b3480156106ef57600080fd5b506101d1600160a060020a0360043516611f63565b34801561071057600080fd5b50610725600435602435604435606435611fa4565b60408051938452602084019290925282820152519081900360600190f35b600080808080808033803b8015610792576040805160e560020a62461bcd02815260206004820152601160248201526000805160206132b5833981519152604482015290519081900360640190fd5b336000818152600760205260409020549099509750871515610824576040805160e560020a62461bcd02815260206004820152602e60248201527f6865792074686572652062756464792c20796f7520646f6e74206576656e206860448201527f61766520616e206163636f756e74000000000000000000000000000000000000606482015290519081900360840190fd5b6000888152600960205260409020600281015460038201546001928301549199509750955093505b60055484116109a65760008481526002602052604080822054815160e060020a6349cc635d028152600481018c9052600160a060020a038d81166024830152604482018a9052606482018c9052925192909116926349cc635d9260848084019382900301818387803b1580156108c157600080fd5b505af11580156108d5573d6000803e3d6000fd5b50505050600186111561099b57600192505b85831161099b576000848152600260209081526040808320548b8452600b83528184208785529092528083205481517f8f7140ea000000000000000000000000000000000000000000000000000000008152600481018d905260248101919091529051600160a060020a0390921692638f7140ea9260448084019382900301818387803b15801561097757600080fd5b505af115801561098b573d6000803e3d6000fd5b5050600190940193506108e79050565b60019093019261084c565b505050505050505050565b6040805160e060020a630c3f64bf0281523360048201529051731a1837c6d259be9c77080d6e931c8fce5709c4a091630c3f64bf9160248083019260209291908290030181600087803b158015610a0757600080fd5b505af1158015610a1b573d6000803e3d6000fd5b505050506040513d6020811015610a3157600080fd5b50511515600114610a7a576040805160e560020a62461bcd0281526020600482015260176024820152600080516020613275833981519152604482015290519081900360640190fd5b610a8560008261225b565b50565b60076020526000908152604090205481565b60015481565b60015490565b600a60209081526000928352604080842090915290825290205460ff1681565b60036020526000908152604090205481565b60086020526000908152604090205481565b60008080808033803b8015610b37576040805160e560020a62461bcd02815260206004820152601160248201526000805160206132b5833981519152604482015290519081900360640190fd5b600154341015610b93576040805160e560020a62461bcd02815260206004820152602660248201526000805160206132558339815191526044820152600080516020613295833981519152606482015290519081900360840190fd5b610b9c8a61230d565b9650339550610baa86612b20565b600160a060020a03808816600090815260076020526040902054919650909450891615801590610bec575085600160a060020a031689600160a060020a031614155b15610c3a57600160a060020a0389166000908152600760209081526040808320548784526009909252909120600201549093508314610c3a5760008481526009602052604090206002018390555b610c488487858a898d612ba2565b50505050505050505050565b6040805160e060020a630c3f64bf0281523360048201529051600091731a1837c6d259be9c77080d6e931c8fce5709c4a091630c3f64bf9160248082019260209290919082900301818787803b158015610cad57600080fd5b505af1158015610cc1573d6000803e3d6000fd5b505050506040513d6020811015610cd757600080fd5b50511515600114610d20576040805160e560020a62461bcd0281526020600482015260176024820152600080516020613275833981519152604482015290519081900360640190fd5b600160a060020a03831660009081526004602052604090205415610db4576040805160e560020a62461bcd02815260206004820152602860248201527f646572702c20746861742067616d657320616c7265616479206265656e20726560448201527f6769737465726564000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b610ddd7f61646447616d6500000000000000000000000000000000000000000000000000612e70565b1515600114156111bb57610e107f61646447616d6500000000000000000000000000000000000000000000000000610a7a565b600580546001019055610e228261230d565b60058054600160a060020a03808716600081815260046020818152604080842096909655600381528583208890558654835260028152858320805473ffffffffffffffffffffffffffffffffffffffff19169094179093559454815283812054600180835260099093527f92e85d02570a8092d09a6e3a57665bc3815a2699a4074001bf1ccabf660f5a36547f92e85d02570a8092d09a6e3a57665bc3815a2699a4074001bf1ccabf660f5a3754865160e060020a6349cc635d0281529788019490945284166024870152604486019290925260648501819052925194955016926349cc635d92608480820193929182900301818387803b158015610f2657600080fd5b505af1158015610f3a573d6000803e3d6000fd5b505060055460009081526002602081815260408084205483855260099092527f6cde3cea4b3a3fb2488b2808bae7556f4a405e50f65e1794383bc026131b13c3547f6cde3cea4b3a3fb2488b2808bae7556f4a405e50f65e1794383bc026131b13c454825160e060020a6349cc635d0281526004810195909552600160a060020a0391821660248601526044850152606484018590529051911694506349cc635d93506084808301939282900301818387803b158015610ff957600080fd5b505af115801561100d573d6000803e3d6000fd5b5050600554600090815260026020908152604080832054600380855260099093527fc575c31fea594a6eb97c8e9d3f9caee4c16218c6ef37e923234c0fe9014a61e7547fc575c31fea594a6eb97c8e9d3f9caee4c16218c6ef37e923234c0fe9014a61e854835160e060020a6349cc635d0281526004810195909552600160a060020a0391821660248601526044850152606484018590529151911694506349cc635d93506084808301939282900301818387803b1580156110ce57600080fd5b505af11580156110e2573d6000803e3d6000fd5b5050600554600090815260026020908152604080832054600480855260099093527f8dc18c4ccfd75f5c815b63770fa542fd953e8fef7e0e44bbdd4913470ce7e9cb547f8dc18c4ccfd75f5c815b63770fa542fd953e8fef7e0e44bbdd4913470ce7e9cc54835160e060020a6349cc635d02815280860195909552600160a060020a0391821660248601526044850152606484018590529151911694506349cc635d93506084808301939282900301818387803b1580156111a257600080fd5b505af11580156111b6573d6000803e3d6000fd5b505050505b505050565b60065481565b600081815260096020526040902054600160a060020a03165b919050565b60008080808033803b8015611231576040805160e560020a62461bcd02815260206004820152601160248201526000805160206132b5833981519152604482015290519081900360640190fd5b60015434101561128d576040805160e560020a62461bcd02815260206004820152602660248201526000805160206132558339815191526044820152600080516020613295833981519152606482015290519081900360840190fd5b6112968a61230d565b96503395506112a486612b20565b600160a060020a038716600090815260076020526040902054909550935088158015906112d15750888714155b15610c3a576000898152600860209081526040808320548784526009909252909120600201549093508314610c3a576000848152600960205260409020600201839055610c488487858a898d612ba2565b60008061132e8361230d565b600081815260086020526040902054909150151561134f5760019150611354565b600091505b50919050565b336000908152600460205260408120548190819081908190151561137d57600080fd5b6001543410156113d9576040805160e560020a62461bcd02815260206004820152602660248201526000805160206132558339815191526044820152600080516020613295833981519152606482015290519081900360840190fd5b6113e289612b20565b600160a060020a038a166000908152600760205260409020549093509150861580159061140f5750868814155b15611451575060008681526008602090815260408083205484845260099092529091206002015481146114515760008281526009602052604090206002018190555b61145f828a838b878b612ba2565b91989197509095505050505050565b6040805160e060020a630c3f64bf028152336004820152905160009182918291731a1837c6d259be9c77080d6e931c8fce5709c4a091630c3f64bf9160248082019260209290919082900301818787803b1580156114cb57600080fd5b505af11580156114df573d6000803e3d6000fd5b505050506040513d60208110156114f557600080fd5b5051151560011461153e576040805160e560020a62461bcd0281526020600482015260176024820152600080516020613275833981519152604482015290519081900360640190fd5b61154a60008888612f01565b61155660008988612f01565b61156260008a88612f01565b9250925092509450945094915050565b600080808033803b80156115be576040805160e560020a62461bcd02815260206004820152601160248201526000805160206132b5833981519152604482015290519081900360640190fd5b60055487111561163e576040805160e560020a62461bcd02815260206004820152602960248201527f73696c6c7920706c617965722c20746861742067616d6520646f65736e27742060448201527f6578697374207965740000000000000000000000000000000000000000000000606482015290519081900360840190fd5b3360008181526007602052604090205490965094508415156116d0576040805160e560020a62461bcd02815260206004820152602e60248201527f6865792074686572652062756464792c20796f7520646f6e74206576656e206860448201527f61766520616e206163636f756e74000000000000000000000000000000000000606482015290519081900360840190fd5b600085815260096020818152604080842060038101548c86526002808552838720548c88529590945260018201549390910154825160e060020a6349cc635d028152600481018c9052600160a060020a038d8116602483015260448201959095526064810191909152915190985091909216926349cc635d926084808201939182900301818387803b15801561176557600080fd5b505af1158015611779573d6000803e3d6000fd5b5050505060018411156111b657600192505b8383116111b657600087815260026020908152604080832054888452600b83528184208785529092528083205481517f8f7140ea000000000000000000000000000000000000000000000000000000008152600481018a905260248101919091529051600160a060020a0390921692638f7140ea9260448084019382900301818387803b15801561181b57600080fd5b505af115801561182f573d6000803e3d6000fd5b50506001909401935061178b9050565b60009081526009602052604090206001015490565b600080808033803b80156118a0576040805160e560020a62461bcd02815260206004820152601160248201526000805160206132b5833981519152604482015290519081900360640190fd5b6001543410156118fc576040805160e560020a62461bcd02815260206004820152602660248201526000805160206132558339815191526044820152600080516020613295833981519152606482015290519081900360840190fd5b6119058961230d565b955033945061191385612b20565b600160a060020a0386166000908152600760205260409020549094509250871580159061195157506000838152600960205260409020600201548814155b801561195d5750828814155b1561197b576000838152600960205260409020600201889055611988565b8288141561198857600097505b6109a683868a89888c612ba2565b600b60209081526000928352604080842090915290825290205481565b6040805160e060020a630c3f64bf02815233600482015290516000918291731a1837c6d259be9c77080d6e931c8fce5709c4a091630c3f64bf91602480830192602092919082900301818787803b158015611a0d57600080fd5b505af1158015611a21573d6000803e3d6000fd5b505050506040513d6020811015611a3757600080fd5b50511515600114611a80576040805160e560020a62461bcd0281526020600482015260176024820152600080516020613275833981519152604482015290519081900360640190fd5b611a8b600084612fbe565b611a96600085612fe2565b91509150915091565b3360009081526004602052604081205481908190819081901515611ac257600080fd5b600154341015611b1e576040805160e560020a62461bcd02815260206004820152602660248201526000805160206132558339815191526044820152600080516020613295833981519152606482015290519081900360840190fd5b611b2789612b20565b600160a060020a03808b16600090815260076020526040902054919450909250871615801590611b69575088600160a060020a031687600160a060020a031614155b156114515750600160a060020a038616600090815260076020908152604080832054848452600990925290912060020154811461145157600082815260096020526040902060020181905561145f828a838b878b612ba2565b60008033803b8015611c0c576040805160e560020a62461bcd02815260206004820152601160248201526000805160206132b5833981519152604482015290519081900360640190fd5b611c158561230d565b33600090815260076020908152604080832054808452600a835281842085855290925290912054919550935060ff161515600114611c9d576040805160e560020a62461bcd02815260206004820152601f60248201527f756d6d2e2e2e207468617473206e6f742061206e616d6520796f75206f776e00604482015290519081900360640190fd5b505060009081526009602052604090206001015550565b60055481565b3360009081526004602052604081205481908190819081901515611cdd57600080fd5b600154341015611d39576040805160e560020a62461bcd02815260206004820152602660248201526000805160206132558339815191526044820152600080516020613295833981519152606482015290519081900360840190fd5b611d4289612b20565b600160a060020a038a1660009081526007602052604090205490935091508690508015801590611d8357506000828152600960205260409020600201548114155b8015611d8f5750818114155b15611dad576000828152600960205260409020600201819055611451565b818114156114515750600061145f828a838b878b612ba2565b6040805160e060020a630c3f64bf0281523360048201529051731a1837c6d259be9c77080d6e931c8fce5709c4a091630c3f64bf9160248083019260209291908290030181600087803b158015611e1c57600080fd5b505af1158015611e30573d6000803e3d6000fd5b505050506040513d6020811015611e4657600080fd5b50511515600114611e8f576040805160e560020a62461bcd0281526020600482015260176024820152600080516020613275833981519152604482015290519081900360640190fd5b611eb87f736574526567697374726174696f6e4665650000000000000000000000000000612e70565b151560011415610a8557611eeb7f736574526567697374726174696f6e4665650000000000000000000000000000610a7a565b600155565b600260205260009081526040902054600160a060020a031681565b60046020526000908152604090205481565b6009602052600090815260409020805460018201546002830154600390930154600160a060020a0390921692909184565b60009081526009602052604090206002015490565b336000908152600460205260408120541515611f7e57600080fd5b611f8782612b20565b5050600160a060020a031660009081526007602052604090205490565b6040805160e060020a630c3f64bf028152336004820152905160009182918291731a1837c6d259be9c77080d6e931c8fce5709c4a091630c3f64bf9160248082019260209290919082900301818787803b15801561200157600080fd5b505af1158015612015573d6000803e3d6000fd5b505050506040513d602081101561202b57600080fd5b50511515600114612074576040805160e560020a62461bcd0281526020600482015260176024820152600080516020613275833981519152604482015290519081900360640190fd5b731a1837c6d259be9c77080d6e931c8fce5709c4a063af1c084d61209a60008a8a612f01565b6040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156120e557600080fd5b505af11580156120f9573d6000803e3d6000fd5b505050506040513d602081101561210f57600080fd5b5051731a1837c6d259be9c77080d6e931c8fce5709c4a063af1c084d61213760008b8a612f01565b6040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b15801561218257600080fd5b505af1158015612196573d6000803e3d6000fd5b505050506040513d60208110156121ac57600080fd5b5051731a1837c6d259be9c77080d6e931c8fce5709c4a063af1c084d6121d460008c8a612f01565b6040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b15801561221f57600080fd5b505af1158015612233573d6000803e3d6000fd5b505050506040513d602081101561224957600080fd5b50519199909850909650945050505050565b600080600061226984613009565b9250600090505b6000838152602086905260409020600101548110156122f257600083815260208681526040808320848452600381018084528285208054600160a060020a031680875260029093018552928520805460ff191690559385905292909152805473ffffffffffffffffffffffffffffffffffffffff191690559150600101612270565b50506000908152602092909252506040812081815560010155565b80516000908290828080602084118015906123285750600084115b15156123a4576040805160e560020a62461bcd02815260206004820152602a60248201527f737472696e67206d757374206265206265747765656e203120616e642033322060448201527f6368617261637465727300000000000000000000000000000000000000000000606482015290519081900360840190fd5b8460008151811015156123b357fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a021415801561241a575084600185038151811015156123f257fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a0214155b1515612496576040805160e560020a62461bcd02815260206004820152602560248201527f737472696e672063616e6e6f74207374617274206f7220656e6420776974682060448201527f7370616365000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b8460008151811015156124a557fe5b90602001015160f860020a900460f860020a02600160f860020a031916603060f860020a0214156125e8578460018151811015156124df57fe5b90602001015160f860020a900460f860020a02600160f860020a031916607860f860020a021415151561255c576040805160e560020a62461bcd02815260206004820152601b60248201527f737472696e672063616e6e6f7420737461727420776974682030780000000000604482015290519081900360640190fd5b84600181518110151561256b57fe5b90602001015160f860020a900460f860020a02600160f860020a031916605860f860020a02141515156125e8576040805160e560020a62461bcd02815260206004820152601b60248201527f737472696e672063616e6e6f7420737461727420776974682030580000000000604482015290519081900360640190fd5b600091505b83821015612ab85784517f40000000000000000000000000000000000000000000000000000000000000009086908490811061262557fe5b90602001015160f860020a900460f860020a02600160f860020a031916118015612699575084517f5b000000000000000000000000000000000000000000000000000000000000009086908490811061267a57fe5b90602001015160f860020a900460f860020a02600160f860020a031916105b156127065784828151811015156126ac57fe5b90602001015160f860020a900460f860020a0260f860020a900460200160f860020a0285838151811015156126dd57fe5b906020010190600160f860020a031916908160001a90535082151561270157600192505b612aad565b848281518110151561271457fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a0214806127e4575084517f60000000000000000000000000000000000000000000000000000000000000009086908490811061277057fe5b90602001015160f860020a900460f860020a02600160f860020a0319161180156127e4575084517f7b00000000000000000000000000000000000000000000000000000000000000908690849081106127c557fe5b90602001015160f860020a900460f860020a02600160f860020a031916105b8061288e575084517f2f000000000000000000000000000000000000000000000000000000000000009086908490811061281a57fe5b90602001015160f860020a900460f860020a02600160f860020a03191611801561288e575084517f3a000000000000000000000000000000000000000000000000000000000000009086908490811061286f57fe5b90602001015160f860020a900460f860020a02600160f860020a031916105b151561290a576040805160e560020a62461bcd02815260206004820152602260248201527f737472696e6720636f6e7461696e7320696e76616c696420636861726163746560448201527f7273000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b848281518110151561291857fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a0214156129f757848260010181518110151561295457fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a02141515156129f7576040805160e560020a62461bcd02815260206004820152602860248201527f737472696e672063616e6e6f7420636f6e7461696e20636f6e7365637574697660448201527f6520737061636573000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b82158015612aa3575084517f300000000000000000000000000000000000000000000000000000000000000090869084908110612a3057fe5b90602001015160f860020a900460f860020a02600160f860020a0319161080612aa3575084517f390000000000000000000000000000000000000000000000000000000000000090869084908110612a8457fe5b90602001015160f860020a900460f860020a02600160f860020a031916115b15612aad57600192505b6001909101906125ed565b600183151514612b12576040805160e560020a62461bcd02815260206004820152601d60248201527f737472696e672063616e6e6f74206265206f6e6c79206e756d62657273000000604482015290519081900360640190fd5b505050506020015192915050565b600160a060020a0381166000908152600760205260408120541515612b9a575060068054600190810191829055600160a060020a03831660008181526007602090815260408083208690559482526009905292909220805473ffffffffffffffffffffffffffffffffffffffff19169092179091556111df565b5060006111df565b60008381526008602052604081205415612c2c576000878152600a6020908152604080832087845290915290205460ff161515600114612c2c576040805160e560020a62461bcd02815260206004820152601e60248201527f736f7272792074686174206e616d657320616c72656164792074616b656e0000604482015290519081900360640190fd5b6000878152600960209081526040808320600101879055868352600882528083208a9055898352600a825280832087845290915290205460ff161515612cbc576000878152600a602090815260408083208784528252808320805460ff191660019081179091558a845260098352818420600301805490910190819055600b835281842090845290915290208490555b73ddc5f4eb3540cefc8684be2f2331f4fe1d7393d8600160a060020a031663d0e30db030600160a060020a0316316040518263ffffffff1660e060020a0281526004016020604051808303818588803b158015612d1857600080fd5b505af1158015612d2c573d6000803e3d6000fd5b50505050506040513d6020811015612d4357600080fd5b505060018215151415612ded575060015b6005548111612ded5760008181526002602052604080822054815160e060020a6349cc635d028152600481018b9052600160a060020a038a8116602483015260448201899052606482018a9052925192909116926349cc635d9260848084019382900301818387803b158015612dc957600080fd5b505af1158015612ddd573d6000803e3d6000fd5b505060019092019150612d549050565b600085815260096020908152604091829020805460019091015483518715158152928301899052600160a060020a039182168385015260608301523460808301524260a0830152915186928916918a917fdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e644429181900360c00190a450505050505050565b6000612efb6000731a1837c6d259be9c77080d6e931c8fce5709c4a0600160a060020a031663fcf2f85f6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612ec957600080fd5b505af1158015612edd573d6000803e3d6000fd5b505050506040513d6020811015612ef357600080fd5b50518461309d565b92915050565b600080808311612f81576040805160e560020a62461bcd02815260206004820152602860248201527f4d5346756e20636865636b5369676e6572206661696c6564202d2030206e6f7460448201527f20616c6c6f776564000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b612f8a84613009565b60008181526020878152604080832060001988018452600301909152902054600160a060020a031692509050509392505050565b600080612fca83613009565b60009081526020949094525050604090912054919050565b600080612fee83613009565b60009081526020949094525050604090912060010154919050565b6040805160208082018490526c01000000000000000000000000300282840152825160348184030181526054909201928390528151600093918291908401908083835b6020831061306b5780518252601f19909201916020918201910161304c565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912095945050505050565b60008060008060006130ae86613009565b600081815260208a905260408082206001015490519296509450339350903690808383808284376040519201829003909120945050508415159150613178905057600084815260208981526040808320848155600160a060020a038616808552600282018452828520805460ff19166001908117909155888652600383018552928520805473ffffffffffffffffffffffffffffffffffffffff1916909117905592879052908a90529081018054909101908190558714156131735760019450613249565b613249565b60008481526020899052604090205481141561324957600084815260208981526040808320600160a060020a038616845260020190915290205460ff16151561322b57600084815260208981526040808320600160a060020a038616808552600282018452828520805460ff19166001908117909155888652600383018552928520805473ffffffffffffffffffffffffffffffffffffffff1916909117905592879052908a9052908101805490910190555b60008481526020899052604090206001015487141561324957600194505b5050505093925050505600756d6d2e2e2e2e2e2020796f75206861766520746f2070617920746865206e616d73672073656e646572206973206e6f742061206465760000000000000000006d65206665650000000000000000000000000000000000000000000000000000736f7272792068756d616e73206f6e6c79000000000000000000000000000000a165627a7a723058204b809f199f78238856958bb84222429165555cd9337965a09bc43b839788c0fd0029`

// DeployPlayerBook deploys a new Ethereum contract, binding an instance of PlayerBook to it.
func DeployPlayerBook(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PlayerBook, error) {
	parsed, err := abi.JSON(strings.NewReader(PlayerBookABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PlayerBookBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PlayerBook{PlayerBookCaller: PlayerBookCaller{contract: contract}, PlayerBookTransactor: PlayerBookTransactor{contract: contract}, PlayerBookFilterer: PlayerBookFilterer{contract: contract}}, nil
}

// PlayerBook is an auto generated Go binding around an Ethereum contract.
type PlayerBook struct {
	PlayerBookCaller     // Read-only binding to the contract
	PlayerBookTransactor // Write-only binding to the contract
	PlayerBookFilterer   // Log filterer for contract events
}

// PlayerBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlayerBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayerBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlayerBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayerBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlayerBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayerBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlayerBookSession struct {
	Contract     *PlayerBook       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlayerBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlayerBookCallerSession struct {
	Contract *PlayerBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PlayerBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlayerBookTransactorSession struct {
	Contract     *PlayerBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PlayerBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlayerBookRaw struct {
	Contract *PlayerBook // Generic contract binding to access the raw methods on
}

// PlayerBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlayerBookCallerRaw struct {
	Contract *PlayerBookCaller // Generic read-only contract binding to access the raw methods on
}

// PlayerBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlayerBookTransactorRaw struct {
	Contract *PlayerBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlayerBook creates a new instance of PlayerBook, bound to a specific deployed contract.
func NewPlayerBook(address common.Address, backend bind.ContractBackend) (*PlayerBook, error) {
	contract, err := bindPlayerBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlayerBook{PlayerBookCaller: PlayerBookCaller{contract: contract}, PlayerBookTransactor: PlayerBookTransactor{contract: contract}, PlayerBookFilterer: PlayerBookFilterer{contract: contract}}, nil
}

// NewPlayerBookCaller creates a new read-only instance of PlayerBook, bound to a specific deployed contract.
func NewPlayerBookCaller(address common.Address, caller bind.ContractCaller) (*PlayerBookCaller, error) {
	contract, err := bindPlayerBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlayerBookCaller{contract: contract}, nil
}

// NewPlayerBookTransactor creates a new write-only instance of PlayerBook, bound to a specific deployed contract.
func NewPlayerBookTransactor(address common.Address, transactor bind.ContractTransactor) (*PlayerBookTransactor, error) {
	contract, err := bindPlayerBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlayerBookTransactor{contract: contract}, nil
}

// NewPlayerBookFilterer creates a new log filterer instance of PlayerBook, bound to a specific deployed contract.
func NewPlayerBookFilterer(address common.Address, filterer bind.ContractFilterer) (*PlayerBookFilterer, error) {
	contract, err := bindPlayerBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlayerBookFilterer{contract: contract}, nil
}

// bindPlayerBook binds a generic wrapper to an already deployed contract.
func bindPlayerBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlayerBookABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlayerBook *PlayerBookRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlayerBook.Contract.PlayerBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlayerBook *PlayerBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlayerBook.Contract.PlayerBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlayerBook *PlayerBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlayerBook.Contract.PlayerBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlayerBook *PlayerBookCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlayerBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlayerBook *PlayerBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlayerBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlayerBook *PlayerBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlayerBook.Contract.contract.Transact(opts, method, params...)
}

// CheckData is a free data retrieval call binding the contract method 0xa553506e.
//
// Solidity: function checkData(_whatFunction bytes32) constant returns(bytes32, uint256)
func (_PlayerBook *PlayerBookCaller) CheckData(opts *bind.CallOpts, _whatFunction [32]byte) ([32]byte, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _PlayerBook.contract.Call(opts, out, "checkData", _whatFunction)
	return *ret0, *ret1, err
}

// CheckData is a free data retrieval call binding the contract method 0xa553506e.
//
// Solidity: function checkData(_whatFunction bytes32) constant returns(bytes32, uint256)
func (_PlayerBook *PlayerBookSession) CheckData(_whatFunction [32]byte) ([32]byte, *big.Int, error) {
	return _PlayerBook.Contract.CheckData(&_PlayerBook.CallOpts, _whatFunction)
}

// CheckData is a free data retrieval call binding the contract method 0xa553506e.
//
// Solidity: function checkData(_whatFunction bytes32) constant returns(bytes32, uint256)
func (_PlayerBook *PlayerBookCallerSession) CheckData(_whatFunction [32]byte) ([32]byte, *big.Int, error) {
	return _PlayerBook.Contract.CheckData(&_PlayerBook.CallOpts, _whatFunction)
}

// CheckIfNameValid is a free data retrieval call binding the contract method 0x6c52660d.
//
// Solidity: function checkIfNameValid(_nameStr string) constant returns(bool)
func (_PlayerBook *PlayerBookCaller) CheckIfNameValid(opts *bind.CallOpts, _nameStr string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "checkIfNameValid", _nameStr)
	return *ret0, err
}

// CheckIfNameValid is a free data retrieval call binding the contract method 0x6c52660d.
//
// Solidity: function checkIfNameValid(_nameStr string) constant returns(bool)
func (_PlayerBook *PlayerBookSession) CheckIfNameValid(_nameStr string) (bool, error) {
	return _PlayerBook.Contract.CheckIfNameValid(&_PlayerBook.CallOpts, _nameStr)
}

// CheckIfNameValid is a free data retrieval call binding the contract method 0x6c52660d.
//
// Solidity: function checkIfNameValid(_nameStr string) constant returns(bool)
func (_PlayerBook *PlayerBookCallerSession) CheckIfNameValid(_nameStr string) (bool, error) {
	return _PlayerBook.Contract.CheckIfNameValid(&_PlayerBook.CallOpts, _nameStr)
}

// CheckSignersByAddress is a free data retrieval call binding the contract method 0x768f930d.
//
// Solidity: function checkSignersByAddress(_whatFunction bytes32, _signerA uint256, _signerB uint256, _signerC uint256) constant returns(address, address, address)
func (_PlayerBook *PlayerBookCaller) CheckSignersByAddress(opts *bind.CallOpts, _whatFunction [32]byte, _signerA *big.Int, _signerB *big.Int, _signerC *big.Int) (common.Address, common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _PlayerBook.contract.Call(opts, out, "checkSignersByAddress", _whatFunction, _signerA, _signerB, _signerC)
	return *ret0, *ret1, *ret2, err
}

// CheckSignersByAddress is a free data retrieval call binding the contract method 0x768f930d.
//
// Solidity: function checkSignersByAddress(_whatFunction bytes32, _signerA uint256, _signerB uint256, _signerC uint256) constant returns(address, address, address)
func (_PlayerBook *PlayerBookSession) CheckSignersByAddress(_whatFunction [32]byte, _signerA *big.Int, _signerB *big.Int, _signerC *big.Int) (common.Address, common.Address, common.Address, error) {
	return _PlayerBook.Contract.CheckSignersByAddress(&_PlayerBook.CallOpts, _whatFunction, _signerA, _signerB, _signerC)
}

// CheckSignersByAddress is a free data retrieval call binding the contract method 0x768f930d.
//
// Solidity: function checkSignersByAddress(_whatFunction bytes32, _signerA uint256, _signerB uint256, _signerC uint256) constant returns(address, address, address)
func (_PlayerBook *PlayerBookCallerSession) CheckSignersByAddress(_whatFunction [32]byte, _signerA *big.Int, _signerB *big.Int, _signerC *big.Int) (common.Address, common.Address, common.Address, error) {
	return _PlayerBook.Contract.CheckSignersByAddress(&_PlayerBook.CallOpts, _whatFunction, _signerA, _signerB, _signerC)
}

// CheckSignersByName is a free data retrieval call binding the contract method 0xed3643d6.
//
// Solidity: function checkSignersByName(_whatFunction bytes32, _signerA uint256, _signerB uint256, _signerC uint256) constant returns(bytes32, bytes32, bytes32)
func (_PlayerBook *PlayerBookCaller) CheckSignersByName(opts *bind.CallOpts, _whatFunction [32]byte, _signerA *big.Int, _signerB *big.Int, _signerC *big.Int) ([32]byte, [32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _PlayerBook.contract.Call(opts, out, "checkSignersByName", _whatFunction, _signerA, _signerB, _signerC)
	return *ret0, *ret1, *ret2, err
}

// CheckSignersByName is a free data retrieval call binding the contract method 0xed3643d6.
//
// Solidity: function checkSignersByName(_whatFunction bytes32, _signerA uint256, _signerB uint256, _signerC uint256) constant returns(bytes32, bytes32, bytes32)
func (_PlayerBook *PlayerBookSession) CheckSignersByName(_whatFunction [32]byte, _signerA *big.Int, _signerB *big.Int, _signerC *big.Int) ([32]byte, [32]byte, [32]byte, error) {
	return _PlayerBook.Contract.CheckSignersByName(&_PlayerBook.CallOpts, _whatFunction, _signerA, _signerB, _signerC)
}

// CheckSignersByName is a free data retrieval call binding the contract method 0xed3643d6.
//
// Solidity: function checkSignersByName(_whatFunction bytes32, _signerA uint256, _signerB uint256, _signerC uint256) constant returns(bytes32, bytes32, bytes32)
func (_PlayerBook *PlayerBookCallerSession) CheckSignersByName(_whatFunction [32]byte, _signerA *big.Int, _signerB *big.Int, _signerC *big.Int) ([32]byte, [32]byte, [32]byte, error) {
	return _PlayerBook.Contract.CheckSignersByName(&_PlayerBook.CallOpts, _whatFunction, _signerA, _signerB, _signerC)
}

// GID is a free data retrieval call binding the contract method 0xb9eca0c8.
//
// Solidity: function gID_() constant returns(uint256)
func (_PlayerBook *PlayerBookCaller) GID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "gID_")
	return *ret0, err
}

// GID is a free data retrieval call binding the contract method 0xb9eca0c8.
//
// Solidity: function gID_() constant returns(uint256)
func (_PlayerBook *PlayerBookSession) GID() (*big.Int, error) {
	return _PlayerBook.Contract.GID(&_PlayerBook.CallOpts)
}

// GID is a free data retrieval call binding the contract method 0xb9eca0c8.
//
// Solidity: function gID_() constant returns(uint256)
func (_PlayerBook *PlayerBookCallerSession) GID() (*big.Int, error) {
	return _PlayerBook.Contract.GID(&_PlayerBook.CallOpts)
}

// GameIDs is a free data retrieval call binding the contract method 0xdbbcaa97.
//
// Solidity: function gameIDs_( address) constant returns(uint256)
func (_PlayerBook *PlayerBookCaller) GameIDs(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "gameIDs_", arg0)
	return *ret0, err
}

// GameIDs is a free data retrieval call binding the contract method 0xdbbcaa97.
//
// Solidity: function gameIDs_( address) constant returns(uint256)
func (_PlayerBook *PlayerBookSession) GameIDs(arg0 common.Address) (*big.Int, error) {
	return _PlayerBook.Contract.GameIDs(&_PlayerBook.CallOpts, arg0)
}

// GameIDs is a free data retrieval call binding the contract method 0xdbbcaa97.
//
// Solidity: function gameIDs_( address) constant returns(uint256)
func (_PlayerBook *PlayerBookCallerSession) GameIDs(arg0 common.Address) (*big.Int, error) {
	return _PlayerBook.Contract.GameIDs(&_PlayerBook.CallOpts, arg0)
}

// GameNames is a free data retrieval call binding the contract method 0x27249e61.
//
// Solidity: function gameNames_( address) constant returns(bytes32)
func (_PlayerBook *PlayerBookCaller) GameNames(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "gameNames_", arg0)
	return *ret0, err
}

// GameNames is a free data retrieval call binding the contract method 0x27249e61.
//
// Solidity: function gameNames_( address) constant returns(bytes32)
func (_PlayerBook *PlayerBookSession) GameNames(arg0 common.Address) ([32]byte, error) {
	return _PlayerBook.Contract.GameNames(&_PlayerBook.CallOpts, arg0)
}

// GameNames is a free data retrieval call binding the contract method 0x27249e61.
//
// Solidity: function gameNames_( address) constant returns(bytes32)
func (_PlayerBook *PlayerBookCallerSession) GameNames(arg0 common.Address) ([32]byte, error) {
	return _PlayerBook.Contract.GameNames(&_PlayerBook.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0xd5241279.
//
// Solidity: function games_( uint256) constant returns(address)
func (_PlayerBook *PlayerBookCaller) Games(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "games_", arg0)
	return *ret0, err
}

// Games is a free data retrieval call binding the contract method 0xd5241279.
//
// Solidity: function games_( uint256) constant returns(address)
func (_PlayerBook *PlayerBookSession) Games(arg0 *big.Int) (common.Address, error) {
	return _PlayerBook.Contract.Games(&_PlayerBook.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0xd5241279.
//
// Solidity: function games_( uint256) constant returns(address)
func (_PlayerBook *PlayerBookCallerSession) Games(arg0 *big.Int) (common.Address, error) {
	return _PlayerBook.Contract.Games(&_PlayerBook.CallOpts, arg0)
}

// GetNameFee is a free data retrieval call binding the contract method 0x2614195f.
//
// Solidity: function getNameFee() constant returns(uint256)
func (_PlayerBook *PlayerBookCaller) GetNameFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "getNameFee")
	return *ret0, err
}

// GetNameFee is a free data retrieval call binding the contract method 0x2614195f.
//
// Solidity: function getNameFee() constant returns(uint256)
func (_PlayerBook *PlayerBookSession) GetNameFee() (*big.Int, error) {
	return _PlayerBook.Contract.GetNameFee(&_PlayerBook.CallOpts)
}

// GetNameFee is a free data retrieval call binding the contract method 0x2614195f.
//
// Solidity: function getNameFee() constant returns(uint256)
func (_PlayerBook *PlayerBookCallerSession) GetNameFee() (*big.Int, error) {
	return _PlayerBook.Contract.GetNameFee(&_PlayerBook.CallOpts)
}

// GetPlayerAddr is a free data retrieval call binding the contract method 0x4d0d35ff.
//
// Solidity: function getPlayerAddr(_pID uint256) constant returns(address)
func (_PlayerBook *PlayerBookCaller) GetPlayerAddr(opts *bind.CallOpts, _pID *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "getPlayerAddr", _pID)
	return *ret0, err
}

// GetPlayerAddr is a free data retrieval call binding the contract method 0x4d0d35ff.
//
// Solidity: function getPlayerAddr(_pID uint256) constant returns(address)
func (_PlayerBook *PlayerBookSession) GetPlayerAddr(_pID *big.Int) (common.Address, error) {
	return _PlayerBook.Contract.GetPlayerAddr(&_PlayerBook.CallOpts, _pID)
}

// GetPlayerAddr is a free data retrieval call binding the contract method 0x4d0d35ff.
//
// Solidity: function getPlayerAddr(_pID uint256) constant returns(address)
func (_PlayerBook *PlayerBookCallerSession) GetPlayerAddr(_pID *big.Int) (common.Address, error) {
	return _PlayerBook.Contract.GetPlayerAddr(&_PlayerBook.CallOpts, _pID)
}

// GetPlayerLAff is a free data retrieval call binding the contract method 0xe3c08adf.
//
// Solidity: function getPlayerLAff(_pID uint256) constant returns(uint256)
func (_PlayerBook *PlayerBookCaller) GetPlayerLAff(opts *bind.CallOpts, _pID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "getPlayerLAff", _pID)
	return *ret0, err
}

// GetPlayerLAff is a free data retrieval call binding the contract method 0xe3c08adf.
//
// Solidity: function getPlayerLAff(_pID uint256) constant returns(uint256)
func (_PlayerBook *PlayerBookSession) GetPlayerLAff(_pID *big.Int) (*big.Int, error) {
	return _PlayerBook.Contract.GetPlayerLAff(&_PlayerBook.CallOpts, _pID)
}

// GetPlayerLAff is a free data retrieval call binding the contract method 0xe3c08adf.
//
// Solidity: function getPlayerLAff(_pID uint256) constant returns(uint256)
func (_PlayerBook *PlayerBookCallerSession) GetPlayerLAff(_pID *big.Int) (*big.Int, error) {
	return _PlayerBook.Contract.GetPlayerLAff(&_PlayerBook.CallOpts, _pID)
}

// GetPlayerName is a free data retrieval call binding the contract method 0x82e37b2c.
//
// Solidity: function getPlayerName(_pID uint256) constant returns(bytes32)
func (_PlayerBook *PlayerBookCaller) GetPlayerName(opts *bind.CallOpts, _pID *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "getPlayerName", _pID)
	return *ret0, err
}

// GetPlayerName is a free data retrieval call binding the contract method 0x82e37b2c.
//
// Solidity: function getPlayerName(_pID uint256) constant returns(bytes32)
func (_PlayerBook *PlayerBookSession) GetPlayerName(_pID *big.Int) ([32]byte, error) {
	return _PlayerBook.Contract.GetPlayerName(&_PlayerBook.CallOpts, _pID)
}

// GetPlayerName is a free data retrieval call binding the contract method 0x82e37b2c.
//
// Solidity: function getPlayerName(_pID uint256) constant returns(bytes32)
func (_PlayerBook *PlayerBookCallerSession) GetPlayerName(_pID *big.Int) ([32]byte, error) {
	return _PlayerBook.Contract.GetPlayerName(&_PlayerBook.CallOpts, _pID)
}

// PID is a free data retrieval call binding the contract method 0x4b227176.
//
// Solidity: function pID_() constant returns(uint256)
func (_PlayerBook *PlayerBookCaller) PID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "pID_")
	return *ret0, err
}

// PID is a free data retrieval call binding the contract method 0x4b227176.
//
// Solidity: function pID_() constant returns(uint256)
func (_PlayerBook *PlayerBookSession) PID() (*big.Int, error) {
	return _PlayerBook.Contract.PID(&_PlayerBook.CallOpts)
}

// PID is a free data retrieval call binding the contract method 0x4b227176.
//
// Solidity: function pID_() constant returns(uint256)
func (_PlayerBook *PlayerBookCallerSession) PID() (*big.Int, error) {
	return _PlayerBook.Contract.PID(&_PlayerBook.CallOpts)
}

// PIDxAddr is a free data retrieval call binding the contract method 0x10f01eba.
//
// Solidity: function pIDxAddr_( address) constant returns(uint256)
func (_PlayerBook *PlayerBookCaller) PIDxAddr(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "pIDxAddr_", arg0)
	return *ret0, err
}

// PIDxAddr is a free data retrieval call binding the contract method 0x10f01eba.
//
// Solidity: function pIDxAddr_( address) constant returns(uint256)
func (_PlayerBook *PlayerBookSession) PIDxAddr(arg0 common.Address) (*big.Int, error) {
	return _PlayerBook.Contract.PIDxAddr(&_PlayerBook.CallOpts, arg0)
}

// PIDxAddr is a free data retrieval call binding the contract method 0x10f01eba.
//
// Solidity: function pIDxAddr_( address) constant returns(uint256)
func (_PlayerBook *PlayerBookCallerSession) PIDxAddr(arg0 common.Address) (*big.Int, error) {
	return _PlayerBook.Contract.PIDxAddr(&_PlayerBook.CallOpts, arg0)
}

// PIDxName is a free data retrieval call binding the contract method 0x2e19ebdc.
//
// Solidity: function pIDxName_( bytes32) constant returns(uint256)
func (_PlayerBook *PlayerBookCaller) PIDxName(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "pIDxName_", arg0)
	return *ret0, err
}

// PIDxName is a free data retrieval call binding the contract method 0x2e19ebdc.
//
// Solidity: function pIDxName_( bytes32) constant returns(uint256)
func (_PlayerBook *PlayerBookSession) PIDxName(arg0 [32]byte) (*big.Int, error) {
	return _PlayerBook.Contract.PIDxName(&_PlayerBook.CallOpts, arg0)
}

// PIDxName is a free data retrieval call binding the contract method 0x2e19ebdc.
//
// Solidity: function pIDxName_( bytes32) constant returns(uint256)
func (_PlayerBook *PlayerBookCallerSession) PIDxName(arg0 [32]byte) (*big.Int, error) {
	return _PlayerBook.Contract.PIDxName(&_PlayerBook.CallOpts, arg0)
}

// PlyrNameList is a free data retrieval call binding the contract method 0xa448ed4b.
//
// Solidity: function plyrNameList_( uint256,  uint256) constant returns(bytes32)
func (_PlayerBook *PlayerBookCaller) PlyrNameList(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "plyrNameList_", arg0, arg1)
	return *ret0, err
}

// PlyrNameList is a free data retrieval call binding the contract method 0xa448ed4b.
//
// Solidity: function plyrNameList_( uint256,  uint256) constant returns(bytes32)
func (_PlayerBook *PlayerBookSession) PlyrNameList(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _PlayerBook.Contract.PlyrNameList(&_PlayerBook.CallOpts, arg0, arg1)
}

// PlyrNameList is a free data retrieval call binding the contract method 0xa448ed4b.
//
// Solidity: function plyrNameList_( uint256,  uint256) constant returns(bytes32)
func (_PlayerBook *PlayerBookCallerSession) PlyrNameList(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _PlayerBook.Contract.PlyrNameList(&_PlayerBook.CallOpts, arg0, arg1)
}

// PlyrNames is a free data retrieval call binding the contract method 0x2660316e.
//
// Solidity: function plyrNames_( uint256,  bytes32) constant returns(bool)
func (_PlayerBook *PlayerBookCaller) PlyrNames(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "plyrNames_", arg0, arg1)
	return *ret0, err
}

// PlyrNames is a free data retrieval call binding the contract method 0x2660316e.
//
// Solidity: function plyrNames_( uint256,  bytes32) constant returns(bool)
func (_PlayerBook *PlayerBookSession) PlyrNames(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _PlayerBook.Contract.PlyrNames(&_PlayerBook.CallOpts, arg0, arg1)
}

// PlyrNames is a free data retrieval call binding the contract method 0x2660316e.
//
// Solidity: function plyrNames_( uint256,  bytes32) constant returns(bool)
func (_PlayerBook *PlayerBookCallerSession) PlyrNames(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _PlayerBook.Contract.PlyrNames(&_PlayerBook.CallOpts, arg0, arg1)
}

// Plyr is a free data retrieval call binding the contract method 0xde7874f3.
//
// Solidity: function plyr_( uint256) constant returns(addr address, name bytes32, laff uint256, names uint256)
func (_PlayerBook *PlayerBookCaller) Plyr(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr  common.Address
	Name  [32]byte
	Laff  *big.Int
	Names *big.Int
}, error) {
	ret := new(struct {
		Addr  common.Address
		Name  [32]byte
		Laff  *big.Int
		Names *big.Int
	})
	out := ret
	err := _PlayerBook.contract.Call(opts, out, "plyr_", arg0)
	return *ret, err
}

// Plyr is a free data retrieval call binding the contract method 0xde7874f3.
//
// Solidity: function plyr_( uint256) constant returns(addr address, name bytes32, laff uint256, names uint256)
func (_PlayerBook *PlayerBookSession) Plyr(arg0 *big.Int) (struct {
	Addr  common.Address
	Name  [32]byte
	Laff  *big.Int
	Names *big.Int
}, error) {
	return _PlayerBook.Contract.Plyr(&_PlayerBook.CallOpts, arg0)
}

// Plyr is a free data retrieval call binding the contract method 0xde7874f3.
//
// Solidity: function plyr_( uint256) constant returns(addr address, name bytes32, laff uint256, names uint256)
func (_PlayerBook *PlayerBookCallerSession) Plyr(arg0 *big.Int) (struct {
	Addr  common.Address
	Name  [32]byte
	Laff  *big.Int
	Names *big.Int
}, error) {
	return _PlayerBook.Contract.Plyr(&_PlayerBook.CallOpts, arg0)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x180603eb.
//
// Solidity: function registrationFee_() constant returns(uint256)
func (_PlayerBook *PlayerBookCaller) RegistrationFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlayerBook.contract.Call(opts, out, "registrationFee_")
	return *ret0, err
}

// RegistrationFee is a free data retrieval call binding the contract method 0x180603eb.
//
// Solidity: function registrationFee_() constant returns(uint256)
func (_PlayerBook *PlayerBookSession) RegistrationFee() (*big.Int, error) {
	return _PlayerBook.Contract.RegistrationFee(&_PlayerBook.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x180603eb.
//
// Solidity: function registrationFee_() constant returns(uint256)
func (_PlayerBook *PlayerBookCallerSession) RegistrationFee() (*big.Int, error) {
	return _PlayerBook.Contract.RegistrationFee(&_PlayerBook.CallOpts)
}

// AddGame is a paid mutator transaction binding the contract method 0x3fda926e.
//
// Solidity: function addGame(_gameAddress address, _gameNameStr string) returns()
func (_PlayerBook *PlayerBookTransactor) AddGame(opts *bind.TransactOpts, _gameAddress common.Address, _gameNameStr string) (*types.Transaction, error) {
	return _PlayerBook.contract.Transact(opts, "addGame", _gameAddress, _gameNameStr)
}

// AddGame is a paid mutator transaction binding the contract method 0x3fda926e.
//
// Solidity: function addGame(_gameAddress address, _gameNameStr string) returns()
func (_PlayerBook *PlayerBookSession) AddGame(_gameAddress common.Address, _gameNameStr string) (*types.Transaction, error) {
	return _PlayerBook.Contract.AddGame(&_PlayerBook.TransactOpts, _gameAddress, _gameNameStr)
}

// AddGame is a paid mutator transaction binding the contract method 0x3fda926e.
//
// Solidity: function addGame(_gameAddress address, _gameNameStr string) returns()
func (_PlayerBook *PlayerBookTransactorSession) AddGame(_gameAddress common.Address, _gameNameStr string) (*types.Transaction, error) {
	return _PlayerBook.Contract.AddGame(&_PlayerBook.TransactOpts, _gameAddress, _gameNameStr)
}

// AddMeToAllGames is a paid mutator transaction binding the contract method 0x0c6940ea.
//
// Solidity: function addMeToAllGames() returns()
func (_PlayerBook *PlayerBookTransactor) AddMeToAllGames(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlayerBook.contract.Transact(opts, "addMeToAllGames")
}

// AddMeToAllGames is a paid mutator transaction binding the contract method 0x0c6940ea.
//
// Solidity: function addMeToAllGames() returns()
func (_PlayerBook *PlayerBookSession) AddMeToAllGames() (*types.Transaction, error) {
	return _PlayerBook.Contract.AddMeToAllGames(&_PlayerBook.TransactOpts)
}

// AddMeToAllGames is a paid mutator transaction binding the contract method 0x0c6940ea.
//
// Solidity: function addMeToAllGames() returns()
func (_PlayerBook *PlayerBookTransactorSession) AddMeToAllGames() (*types.Transaction, error) {
	return _PlayerBook.Contract.AddMeToAllGames(&_PlayerBook.TransactOpts)
}

// AddMeToGame is a paid mutator transaction binding the contract method 0x81c5b206.
//
// Solidity: function addMeToGame(_gameID uint256) returns()
func (_PlayerBook *PlayerBookTransactor) AddMeToGame(opts *bind.TransactOpts, _gameID *big.Int) (*types.Transaction, error) {
	return _PlayerBook.contract.Transact(opts, "addMeToGame", _gameID)
}

// AddMeToGame is a paid mutator transaction binding the contract method 0x81c5b206.
//
// Solidity: function addMeToGame(_gameID uint256) returns()
func (_PlayerBook *PlayerBookSession) AddMeToGame(_gameID *big.Int) (*types.Transaction, error) {
	return _PlayerBook.Contract.AddMeToGame(&_PlayerBook.TransactOpts, _gameID)
}

// AddMeToGame is a paid mutator transaction binding the contract method 0x81c5b206.
//
// Solidity: function addMeToGame(_gameID uint256) returns()
func (_PlayerBook *PlayerBookTransactorSession) AddMeToGame(_gameID *big.Int) (*types.Transaction, error) {
	return _PlayerBook.Contract.AddMeToGame(&_PlayerBook.TransactOpts, _gameID)
}

// DeleteAnyProposal is a paid mutator transaction binding the contract method 0x0efcf295.
//
// Solidity: function deleteAnyProposal(_whatFunction bytes32) returns()
func (_PlayerBook *PlayerBookTransactor) DeleteAnyProposal(opts *bind.TransactOpts, _whatFunction [32]byte) (*types.Transaction, error) {
	return _PlayerBook.contract.Transact(opts, "deleteAnyProposal", _whatFunction)
}

// DeleteAnyProposal is a paid mutator transaction binding the contract method 0x0efcf295.
//
// Solidity: function deleteAnyProposal(_whatFunction bytes32) returns()
func (_PlayerBook *PlayerBookSession) DeleteAnyProposal(_whatFunction [32]byte) (*types.Transaction, error) {
	return _PlayerBook.Contract.DeleteAnyProposal(&_PlayerBook.TransactOpts, _whatFunction)
}

// DeleteAnyProposal is a paid mutator transaction binding the contract method 0x0efcf295.
//
// Solidity: function deleteAnyProposal(_whatFunction bytes32) returns()
func (_PlayerBook *PlayerBookTransactorSession) DeleteAnyProposal(_whatFunction [32]byte) (*types.Transaction, error) {
	return _PlayerBook.Contract.DeleteAnyProposal(&_PlayerBook.TransactOpts, _whatFunction)
}

// GetPlayerID is a paid mutator transaction binding the contract method 0xe56556a9.
//
// Solidity: function getPlayerID(_addr address) returns(uint256)
func (_PlayerBook *PlayerBookTransactor) GetPlayerID(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _PlayerBook.contract.Transact(opts, "getPlayerID", _addr)
}

// GetPlayerID is a paid mutator transaction binding the contract method 0xe56556a9.
//
// Solidity: function getPlayerID(_addr address) returns(uint256)
func (_PlayerBook *PlayerBookSession) GetPlayerID(_addr common.Address) (*types.Transaction, error) {
	return _PlayerBook.Contract.GetPlayerID(&_PlayerBook.TransactOpts, _addr)
}

// GetPlayerID is a paid mutator transaction binding the contract method 0xe56556a9.
//
// Solidity: function getPlayerID(_addr address) returns(uint256)
func (_PlayerBook *PlayerBookTransactorSession) GetPlayerID(_addr common.Address) (*types.Transaction, error) {
	return _PlayerBook.Contract.GetPlayerID(&_PlayerBook.TransactOpts, _addr)
}

// RegisterNameXID is a paid mutator transaction binding the contract method 0x921dec21.
//
// Solidity: function registerNameXID(_nameString string, _affCode uint256, _all bool) returns()
func (_PlayerBook *PlayerBookTransactor) RegisterNameXID(opts *bind.TransactOpts, _nameString string, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _PlayerBook.contract.Transact(opts, "registerNameXID", _nameString, _affCode, _all)
}

// RegisterNameXID is a paid mutator transaction binding the contract method 0x921dec21.
//
// Solidity: function registerNameXID(_nameString string, _affCode uint256, _all bool) returns()
func (_PlayerBook *PlayerBookSession) RegisterNameXID(_nameString string, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _PlayerBook.Contract.RegisterNameXID(&_PlayerBook.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXID is a paid mutator transaction binding the contract method 0x921dec21.
//
// Solidity: function registerNameXID(_nameString string, _affCode uint256, _all bool) returns()
func (_PlayerBook *PlayerBookTransactorSession) RegisterNameXID(_nameString string, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _PlayerBook.Contract.RegisterNameXID(&_PlayerBook.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXIDFromDapp is a paid mutator transaction binding the contract method 0xc0942dfd.
//
// Solidity: function registerNameXIDFromDapp(_addr address, _name bytes32, _affCode uint256, _all bool) returns(bool, uint256)
func (_PlayerBook *PlayerBookTransactor) RegisterNameXIDFromDapp(opts *bind.TransactOpts, _addr common.Address, _name [32]byte, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _PlayerBook.contract.Transact(opts, "registerNameXIDFromDapp", _addr, _name, _affCode, _all)
}

// RegisterNameXIDFromDapp is a paid mutator transaction binding the contract method 0xc0942dfd.
//
// Solidity: function registerNameXIDFromDapp(_addr address, _name bytes32, _affCode uint256, _all bool) returns(bool, uint256)
func (_PlayerBook *PlayerBookSession) RegisterNameXIDFromDapp(_addr common.Address, _name [32]byte, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _PlayerBook.Contract.RegisterNameXIDFromDapp(&_PlayerBook.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXIDFromDapp is a paid mutator transaction binding the contract method 0xc0942dfd.
//
// Solidity: function registerNameXIDFromDapp(_addr address, _name bytes32, _affCode uint256, _all bool) returns(bool, uint256)
func (_PlayerBook *PlayerBookTransactorSession) RegisterNameXIDFromDapp(_addr common.Address, _name [32]byte, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _PlayerBook.Contract.RegisterNameXIDFromDapp(&_PlayerBook.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXaddr is a paid mutator transaction binding the contract method 0x3ddd4698.
//
// Solidity: function registerNameXaddr(_nameString string, _affCode address, _all bool) returns()
func (_PlayerBook *PlayerBookTransactor) RegisterNameXaddr(opts *bind.TransactOpts, _nameString string, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _PlayerBook.contract.Transact(opts, "registerNameXaddr", _nameString, _affCode, _all)
}

// RegisterNameXaddr is a paid mutator transaction binding the contract method 0x3ddd4698.
//
// Solidity: function registerNameXaddr(_nameString string, _affCode address, _all bool) returns()
func (_PlayerBook *PlayerBookSession) RegisterNameXaddr(_nameString string, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _PlayerBook.Contract.RegisterNameXaddr(&_PlayerBook.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXaddr is a paid mutator transaction binding the contract method 0x3ddd4698.
//
// Solidity: function registerNameXaddr(_nameString string, _affCode address, _all bool) returns()
func (_PlayerBook *PlayerBookTransactorSession) RegisterNameXaddr(_nameString string, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _PlayerBook.Contract.RegisterNameXaddr(&_PlayerBook.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXaddrFromDapp is a paid mutator transaction binding the contract method 0xaa4d490b.
//
// Solidity: function registerNameXaddrFromDapp(_addr address, _name bytes32, _affCode address, _all bool) returns(bool, uint256)
func (_PlayerBook *PlayerBookTransactor) RegisterNameXaddrFromDapp(opts *bind.TransactOpts, _addr common.Address, _name [32]byte, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _PlayerBook.contract.Transact(opts, "registerNameXaddrFromDapp", _addr, _name, _affCode, _all)
}

// RegisterNameXaddrFromDapp is a paid mutator transaction binding the contract method 0xaa4d490b.
//
// Solidity: function registerNameXaddrFromDapp(_addr address, _name bytes32, _affCode address, _all bool) returns(bool, uint256)
func (_PlayerBook *PlayerBookSession) RegisterNameXaddrFromDapp(_addr common.Address, _name [32]byte, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _PlayerBook.Contract.RegisterNameXaddrFromDapp(&_PlayerBook.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXaddrFromDapp is a paid mutator transaction binding the contract method 0xaa4d490b.
//
// Solidity: function registerNameXaddrFromDapp(_addr address, _name bytes32, _affCode address, _all bool) returns(bool, uint256)
func (_PlayerBook *PlayerBookTransactorSession) RegisterNameXaddrFromDapp(_addr common.Address, _name [32]byte, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _PlayerBook.Contract.RegisterNameXaddrFromDapp(&_PlayerBook.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXname is a paid mutator transaction binding the contract method 0x685ffd83.
//
// Solidity: function registerNameXname(_nameString string, _affCode bytes32, _all bool) returns()
func (_PlayerBook *PlayerBookTransactor) RegisterNameXname(opts *bind.TransactOpts, _nameString string, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _PlayerBook.contract.Transact(opts, "registerNameXname", _nameString, _affCode, _all)
}

// RegisterNameXname is a paid mutator transaction binding the contract method 0x685ffd83.
//
// Solidity: function registerNameXname(_nameString string, _affCode bytes32, _all bool) returns()
func (_PlayerBook *PlayerBookSession) RegisterNameXname(_nameString string, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _PlayerBook.Contract.RegisterNameXname(&_PlayerBook.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXname is a paid mutator transaction binding the contract method 0x685ffd83.
//
// Solidity: function registerNameXname(_nameString string, _affCode bytes32, _all bool) returns()
func (_PlayerBook *PlayerBookTransactorSession) RegisterNameXname(_nameString string, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _PlayerBook.Contract.RegisterNameXname(&_PlayerBook.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXnameFromDapp is a paid mutator transaction binding the contract method 0x745ea0c1.
//
// Solidity: function registerNameXnameFromDapp(_addr address, _name bytes32, _affCode bytes32, _all bool) returns(bool, uint256)
func (_PlayerBook *PlayerBookTransactor) RegisterNameXnameFromDapp(opts *bind.TransactOpts, _addr common.Address, _name [32]byte, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _PlayerBook.contract.Transact(opts, "registerNameXnameFromDapp", _addr, _name, _affCode, _all)
}

// RegisterNameXnameFromDapp is a paid mutator transaction binding the contract method 0x745ea0c1.
//
// Solidity: function registerNameXnameFromDapp(_addr address, _name bytes32, _affCode bytes32, _all bool) returns(bool, uint256)
func (_PlayerBook *PlayerBookSession) RegisterNameXnameFromDapp(_addr common.Address, _name [32]byte, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _PlayerBook.Contract.RegisterNameXnameFromDapp(&_PlayerBook.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXnameFromDapp is a paid mutator transaction binding the contract method 0x745ea0c1.
//
// Solidity: function registerNameXnameFromDapp(_addr address, _name bytes32, _affCode bytes32, _all bool) returns(bool, uint256)
func (_PlayerBook *PlayerBookTransactorSession) RegisterNameXnameFromDapp(_addr common.Address, _name [32]byte, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _PlayerBook.Contract.RegisterNameXnameFromDapp(&_PlayerBook.TransactOpts, _addr, _name, _affCode, _all)
}

// SetRegistrationFee is a paid mutator transaction binding the contract method 0xc320c727.
//
// Solidity: function setRegistrationFee(_fee uint256) returns()
func (_PlayerBook *PlayerBookTransactor) SetRegistrationFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _PlayerBook.contract.Transact(opts, "setRegistrationFee", _fee)
}

// SetRegistrationFee is a paid mutator transaction binding the contract method 0xc320c727.
//
// Solidity: function setRegistrationFee(_fee uint256) returns()
func (_PlayerBook *PlayerBookSession) SetRegistrationFee(_fee *big.Int) (*types.Transaction, error) {
	return _PlayerBook.Contract.SetRegistrationFee(&_PlayerBook.TransactOpts, _fee)
}

// SetRegistrationFee is a paid mutator transaction binding the contract method 0xc320c727.
//
// Solidity: function setRegistrationFee(_fee uint256) returns()
func (_PlayerBook *PlayerBookTransactorSession) SetRegistrationFee(_fee *big.Int) (*types.Transaction, error) {
	return _PlayerBook.Contract.SetRegistrationFee(&_PlayerBook.TransactOpts, _fee)
}

// UseMyOldName is a paid mutator transaction binding the contract method 0xb9291296.
//
// Solidity: function useMyOldName(_nameString string) returns()
func (_PlayerBook *PlayerBookTransactor) UseMyOldName(opts *bind.TransactOpts, _nameString string) (*types.Transaction, error) {
	return _PlayerBook.contract.Transact(opts, "useMyOldName", _nameString)
}

// UseMyOldName is a paid mutator transaction binding the contract method 0xb9291296.
//
// Solidity: function useMyOldName(_nameString string) returns()
func (_PlayerBook *PlayerBookSession) UseMyOldName(_nameString string) (*types.Transaction, error) {
	return _PlayerBook.Contract.UseMyOldName(&_PlayerBook.TransactOpts, _nameString)
}

// UseMyOldName is a paid mutator transaction binding the contract method 0xb9291296.
//
// Solidity: function useMyOldName(_nameString string) returns()
func (_PlayerBook *PlayerBookTransactorSession) UseMyOldName(_nameString string) (*types.Transaction, error) {
	return _PlayerBook.Contract.UseMyOldName(&_PlayerBook.TransactOpts, _nameString)
}

// PlayerBookOnNewNameIterator is returned from FilterOnNewName and is used to iterate over the raw logs and unpacked data for OnNewName events raised by the PlayerBook contract.
type PlayerBookOnNewNameIterator struct {
	Event *PlayerBookOnNewName // Event containing the contract specifics and raw log

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
func (it *PlayerBookOnNewNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlayerBookOnNewName)
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
		it.Event = new(PlayerBookOnNewName)
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
func (it *PlayerBookOnNewNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlayerBookOnNewNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlayerBookOnNewName represents a OnNewName event raised by the PlayerBook contract.
type PlayerBookOnNewName struct {
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
func (_PlayerBook *PlayerBookFilterer) FilterOnNewName(opts *bind.FilterOpts, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (*PlayerBookOnNewNameIterator, error) {

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

	logs, sub, err := _PlayerBook.contract.FilterLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return &PlayerBookOnNewNameIterator{contract: _PlayerBook.contract, event: "onNewName", logs: logs, sub: sub}, nil
}

// WatchOnNewName is a free log subscription operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: event onNewName(playerID indexed uint256, playerAddress indexed address, playerName indexed bytes32, isNewPlayer bool, affiliateID uint256, affiliateAddress address, affiliateName bytes32, amountPaid uint256, timeStamp uint256)
func (_PlayerBook *PlayerBookFilterer) WatchOnNewName(opts *bind.WatchOpts, sink chan<- *PlayerBookOnNewName, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PlayerBook.contract.WatchLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlayerBookOnNewName)
				if err := _PlayerBook.contract.UnpackLog(event, "onNewName", log); err != nil {
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

// PlayerBookReceiverInterfaceABI is the input ABI used to generate the binding from.
const PlayerBookReceiverInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"},{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_laff\",\"type\":\"uint256\"}],\"name\":\"receivePlayerInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"receivePlayerNameList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PlayerBookReceiverInterfaceBin is the compiled bytecode used for deploying new contracts.
const PlayerBookReceiverInterfaceBin = `0x`

// DeployPlayerBookReceiverInterface deploys a new Ethereum contract, binding an instance of PlayerBookReceiverInterface to it.
func DeployPlayerBookReceiverInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PlayerBookReceiverInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(PlayerBookReceiverInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PlayerBookReceiverInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PlayerBookReceiverInterface{PlayerBookReceiverInterfaceCaller: PlayerBookReceiverInterfaceCaller{contract: contract}, PlayerBookReceiverInterfaceTransactor: PlayerBookReceiverInterfaceTransactor{contract: contract}, PlayerBookReceiverInterfaceFilterer: PlayerBookReceiverInterfaceFilterer{contract: contract}}, nil
}

// PlayerBookReceiverInterface is an auto generated Go binding around an Ethereum contract.
type PlayerBookReceiverInterface struct {
	PlayerBookReceiverInterfaceCaller     // Read-only binding to the contract
	PlayerBookReceiverInterfaceTransactor // Write-only binding to the contract
	PlayerBookReceiverInterfaceFilterer   // Log filterer for contract events
}

// PlayerBookReceiverInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlayerBookReceiverInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayerBookReceiverInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlayerBookReceiverInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayerBookReceiverInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlayerBookReceiverInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayerBookReceiverInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlayerBookReceiverInterfaceSession struct {
	Contract     *PlayerBookReceiverInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PlayerBookReceiverInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlayerBookReceiverInterfaceCallerSession struct {
	Contract *PlayerBookReceiverInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// PlayerBookReceiverInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlayerBookReceiverInterfaceTransactorSession struct {
	Contract     *PlayerBookReceiverInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// PlayerBookReceiverInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlayerBookReceiverInterfaceRaw struct {
	Contract *PlayerBookReceiverInterface // Generic contract binding to access the raw methods on
}

// PlayerBookReceiverInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlayerBookReceiverInterfaceCallerRaw struct {
	Contract *PlayerBookReceiverInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// PlayerBookReceiverInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlayerBookReceiverInterfaceTransactorRaw struct {
	Contract *PlayerBookReceiverInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlayerBookReceiverInterface creates a new instance of PlayerBookReceiverInterface, bound to a specific deployed contract.
func NewPlayerBookReceiverInterface(address common.Address, backend bind.ContractBackend) (*PlayerBookReceiverInterface, error) {
	contract, err := bindPlayerBookReceiverInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlayerBookReceiverInterface{PlayerBookReceiverInterfaceCaller: PlayerBookReceiverInterfaceCaller{contract: contract}, PlayerBookReceiverInterfaceTransactor: PlayerBookReceiverInterfaceTransactor{contract: contract}, PlayerBookReceiverInterfaceFilterer: PlayerBookReceiverInterfaceFilterer{contract: contract}}, nil
}

// NewPlayerBookReceiverInterfaceCaller creates a new read-only instance of PlayerBookReceiverInterface, bound to a specific deployed contract.
func NewPlayerBookReceiverInterfaceCaller(address common.Address, caller bind.ContractCaller) (*PlayerBookReceiverInterfaceCaller, error) {
	contract, err := bindPlayerBookReceiverInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlayerBookReceiverInterfaceCaller{contract: contract}, nil
}

// NewPlayerBookReceiverInterfaceTransactor creates a new write-only instance of PlayerBookReceiverInterface, bound to a specific deployed contract.
func NewPlayerBookReceiverInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*PlayerBookReceiverInterfaceTransactor, error) {
	contract, err := bindPlayerBookReceiverInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlayerBookReceiverInterfaceTransactor{contract: contract}, nil
}

// NewPlayerBookReceiverInterfaceFilterer creates a new log filterer instance of PlayerBookReceiverInterface, bound to a specific deployed contract.
func NewPlayerBookReceiverInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*PlayerBookReceiverInterfaceFilterer, error) {
	contract, err := bindPlayerBookReceiverInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlayerBookReceiverInterfaceFilterer{contract: contract}, nil
}

// bindPlayerBookReceiverInterface binds a generic wrapper to an already deployed contract.
func bindPlayerBookReceiverInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlayerBookReceiverInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlayerBookReceiverInterface *PlayerBookReceiverInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlayerBookReceiverInterface.Contract.PlayerBookReceiverInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlayerBookReceiverInterface *PlayerBookReceiverInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlayerBookReceiverInterface.Contract.PlayerBookReceiverInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlayerBookReceiverInterface *PlayerBookReceiverInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlayerBookReceiverInterface.Contract.PlayerBookReceiverInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlayerBookReceiverInterface *PlayerBookReceiverInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlayerBookReceiverInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlayerBookReceiverInterface *PlayerBookReceiverInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlayerBookReceiverInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlayerBookReceiverInterface *PlayerBookReceiverInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlayerBookReceiverInterface.Contract.contract.Transact(opts, method, params...)
}

// ReceivePlayerInfo is a paid mutator transaction binding the contract method 0x49cc635d.
//
// Solidity: function receivePlayerInfo(_pID uint256, _addr address, _name bytes32, _laff uint256) returns()
func (_PlayerBookReceiverInterface *PlayerBookReceiverInterfaceTransactor) ReceivePlayerInfo(opts *bind.TransactOpts, _pID *big.Int, _addr common.Address, _name [32]byte, _laff *big.Int) (*types.Transaction, error) {
	return _PlayerBookReceiverInterface.contract.Transact(opts, "receivePlayerInfo", _pID, _addr, _name, _laff)
}

// ReceivePlayerInfo is a paid mutator transaction binding the contract method 0x49cc635d.
//
// Solidity: function receivePlayerInfo(_pID uint256, _addr address, _name bytes32, _laff uint256) returns()
func (_PlayerBookReceiverInterface *PlayerBookReceiverInterfaceSession) ReceivePlayerInfo(_pID *big.Int, _addr common.Address, _name [32]byte, _laff *big.Int) (*types.Transaction, error) {
	return _PlayerBookReceiverInterface.Contract.ReceivePlayerInfo(&_PlayerBookReceiverInterface.TransactOpts, _pID, _addr, _name, _laff)
}

// ReceivePlayerInfo is a paid mutator transaction binding the contract method 0x49cc635d.
//
// Solidity: function receivePlayerInfo(_pID uint256, _addr address, _name bytes32, _laff uint256) returns()
func (_PlayerBookReceiverInterface *PlayerBookReceiverInterfaceTransactorSession) ReceivePlayerInfo(_pID *big.Int, _addr common.Address, _name [32]byte, _laff *big.Int) (*types.Transaction, error) {
	return _PlayerBookReceiverInterface.Contract.ReceivePlayerInfo(&_PlayerBookReceiverInterface.TransactOpts, _pID, _addr, _name, _laff)
}

// ReceivePlayerNameList is a paid mutator transaction binding the contract method 0x8f7140ea.
//
// Solidity: function receivePlayerNameList(_pID uint256, _name bytes32) returns()
func (_PlayerBookReceiverInterface *PlayerBookReceiverInterfaceTransactor) ReceivePlayerNameList(opts *bind.TransactOpts, _pID *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _PlayerBookReceiverInterface.contract.Transact(opts, "receivePlayerNameList", _pID, _name)
}

// ReceivePlayerNameList is a paid mutator transaction binding the contract method 0x8f7140ea.
//
// Solidity: function receivePlayerNameList(_pID uint256, _name bytes32) returns()
func (_PlayerBookReceiverInterface *PlayerBookReceiverInterfaceSession) ReceivePlayerNameList(_pID *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _PlayerBookReceiverInterface.Contract.ReceivePlayerNameList(&_PlayerBookReceiverInterface.TransactOpts, _pID, _name)
}

// ReceivePlayerNameList is a paid mutator transaction binding the contract method 0x8f7140ea.
//
// Solidity: function receivePlayerNameList(_pID uint256, _name bytes32) returns()
func (_PlayerBookReceiverInterface *PlayerBookReceiverInterfaceTransactorSession) ReceivePlayerNameList(_pID *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _PlayerBookReceiverInterface.Contract.ReceivePlayerNameList(&_PlayerBookReceiverInterface.TransactOpts, _pID, _name)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058201c83ae6d406eb3f13294d375b31284baea86c945ffad421edc1d0afd7eebd6cc0029`

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
