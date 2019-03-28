// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc721

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

// AddressUtilsABI is the input ABI used to generate the binding from.
const AddressUtilsABI = "[]"

// AddressUtilsBin is the compiled bytecode used for deploying new contracts.
const AddressUtilsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058204571c6879ed829d5aa042b62e433287448831e99e4303afa87e8e90da27dfcb80029`

// DeployAddressUtils deploys a new Ethereum contract, binding an instance of AddressUtils to it.
func DeployAddressUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUtils{AddressUtilsCaller: AddressUtilsCaller{contract: contract}, AddressUtilsTransactor: AddressUtilsTransactor{contract: contract}, AddressUtilsFilterer: AddressUtilsFilterer{contract: contract}}, nil
}

// AddressUtils is an auto generated Go binding around an Ethereum contract.
type AddressUtils struct {
	AddressUtilsCaller     // Read-only binding to the contract
	AddressUtilsTransactor // Write-only binding to the contract
	AddressUtilsFilterer   // Log filterer for contract events
}

// AddressUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUtilsSession struct {
	Contract     *AddressUtils     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUtilsCallerSession struct {
	Contract *AddressUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AddressUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUtilsTransactorSession struct {
	Contract     *AddressUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AddressUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUtilsRaw struct {
	Contract *AddressUtils // Generic contract binding to access the raw methods on
}

// AddressUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUtilsCallerRaw struct {
	Contract *AddressUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUtilsTransactorRaw struct {
	Contract *AddressUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUtils creates a new instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtils(address common.Address, backend bind.ContractBackend) (*AddressUtils, error) {
	contract, err := bindAddressUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUtils{AddressUtilsCaller: AddressUtilsCaller{contract: contract}, AddressUtilsTransactor: AddressUtilsTransactor{contract: contract}, AddressUtilsFilterer: AddressUtilsFilterer{contract: contract}}, nil
}

// NewAddressUtilsCaller creates a new read-only instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtilsCaller(address common.Address, caller bind.ContractCaller) (*AddressUtilsCaller, error) {
	contract, err := bindAddressUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUtilsCaller{contract: contract}, nil
}

// NewAddressUtilsTransactor creates a new write-only instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUtilsTransactor, error) {
	contract, err := bindAddressUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUtilsTransactor{contract: contract}, nil
}

// NewAddressUtilsFilterer creates a new log filterer instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUtilsFilterer, error) {
	contract, err := bindAddressUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUtilsFilterer{contract: contract}, nil
}

// bindAddressUtils binds a generic wrapper to an already deployed contract.
func bindAddressUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUtils *AddressUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressUtils.Contract.AddressUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUtils *AddressUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUtils.Contract.AddressUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUtils *AddressUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUtils.Contract.AddressUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUtils *AddressUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUtils *AddressUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUtils *AddressUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUtils.Contract.contract.Transact(opts, method, params...)
}

// ERC165ABI is the input ABI used to generate the binding from.
const ERC165ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC165Bin is the compiled bytecode used for deploying new contracts.
const ERC165Bin = `0x`

// DeployERC165 deploys a new Ethereum contract, binding an instance of ERC165 to it.
func DeployERC165(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC165, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC165Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC165.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, _interfaceId)
}

// ERC721ABI is the input ABI used to generate the binding from.
const ERC721ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"_exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721Bin is the compiled bytecode used for deploying new contracts.
const ERC721Bin = `0x`

// DeployERC721 deploys a new Ethereum contract, binding an instance of ERC721 to it.
func DeployERC721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// ERC721 is an auto generated Go binding around an Ethereum contract.
type ERC721 struct {
	ERC721Caller     // Read-only binding to the contract
	ERC721Transactor // Write-only binding to the contract
	ERC721Filterer   // Log filterer for contract events
}

// ERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721Session struct {
	Contract     *ERC721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721CallerSession struct {
	Contract *ERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TransactorSession struct {
	Contract     *ERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721Raw struct {
	Contract *ERC721 // Generic contract binding to access the raw methods on
}

// ERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721CallerRaw struct {
	Contract *ERC721Caller // Generic read-only contract binding to access the raw methods on
}

// ERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TransactorRaw struct {
	Contract *ERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721 creates a new instance of ERC721, bound to a specific deployed contract.
func NewERC721(address common.Address, backend bind.ContractBackend) (*ERC721, error) {
	contract, err := bindERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// NewERC721Caller creates a new read-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Caller(address common.Address, caller bind.ContractCaller) (*ERC721Caller, error) {
	contract, err := bindERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Caller{contract: contract}, nil
}

// NewERC721Transactor creates a new write-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC721Transactor, error) {
	contract, err := bindERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Transactor{contract: contract}, nil
}

// NewERC721Filterer creates a new log filterer instance of ERC721, bound to a specific deployed contract.
func NewERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC721Filterer, error) {
	contract, err := bindERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721Filterer{contract: contract}, nil
}

// bindERC721 binds a generic wrapper to an already deployed contract.
func bindERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.ERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_ERC721 *ERC721Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_ERC721 *ERC721Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_ERC721 *ERC721CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool _exists)
func (_ERC721 *ERC721Caller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool _exists)
func (_ERC721 *ERC721Session) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721.Contract.Exists(&_ERC721.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool _exists)
func (_ERC721 *ERC721CallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721.Contract.Exists(&_ERC721.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _operator)
func (_ERC721 *ERC721Caller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _operator)
func (_ERC721 *ERC721Session) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _operator)
func (_ERC721 *ERC721CallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721 *ERC721Caller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721 *ERC721Session) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721 *ERC721CallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string _name)
func (_ERC721 *ERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string _name)
func (_ERC721 *ERC721Session) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string _name)
func (_ERC721 *ERC721CallerSession) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_ERC721 *ERC721Caller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_ERC721 *ERC721Session) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_ERC721 *ERC721CallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721 *ERC721Caller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721 *ERC721Session) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721 *ERC721CallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string _symbol)
func (_ERC721 *ERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string _symbol)
func (_ERC721 *ERC721Session) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string _symbol)
func (_ERC721 *ERC721CallerSession) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_ERC721 *ERC721Caller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_ERC721 *ERC721Session) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenByIndex(&_ERC721.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_ERC721 *ERC721CallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenByIndex(&_ERC721.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256 _tokenId)
func (_ERC721 *ERC721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256 _tokenId)
func (_ERC721 *ERC721Session) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenOfOwnerByIndex(&_ERC721.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256 _tokenId)
func (_ERC721 *ERC721CallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenOfOwnerByIndex(&_ERC721.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_ERC721 *ERC721Caller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_ERC721 *ERC721Session) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_ERC721 *ERC721CallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721 *ERC721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721 *ERC721Session) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721 *ERC721CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721Session) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721TransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC721 *ERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC721 *ERC721Session) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC721 *ERC721TransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721Session) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721TransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId)
}

// ERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721 contract.
type ERC721ApprovalIterator struct {
	Event *ERC721Approval // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Approval)
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
		it.Event = new(ERC721Approval)
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
func (it *ERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Approval represents a Approval event raised by the ERC721 contract.
type ERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ERC721ApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Approval)
				if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721 contract.
type ERC721ApprovalForAllIterator struct {
	Event *ERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ApprovalForAll)
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
		it.Event = new(ERC721ApprovalForAll)
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
func (it *ERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ApprovalForAll represents a ApprovalForAll event raised by the ERC721 contract.
type ERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ERC721 *ERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721ApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalForAllIterator{contract: _ERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ERC721 *ERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721ApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ApprovalForAll)
				if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721 contract.
type ERC721TransferIterator struct {
	Event *ERC721Transfer // Event containing the contract specifics and raw log

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
func (it *ERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Transfer)
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
		it.Event = new(ERC721Transfer)
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
func (it *ERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
type ERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ERC721TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Transfer)
				if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721BasicABI is the input ABI used to generate the binding from.
const ERC721BasicABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"_exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721BasicBin is the compiled bytecode used for deploying new contracts.
const ERC721BasicBin = `0x`

// DeployERC721Basic deploys a new Ethereum contract, binding an instance of ERC721Basic to it.
func DeployERC721Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Basic{ERC721BasicCaller: ERC721BasicCaller{contract: contract}, ERC721BasicTransactor: ERC721BasicTransactor{contract: contract}, ERC721BasicFilterer: ERC721BasicFilterer{contract: contract}}, nil
}

// ERC721Basic is an auto generated Go binding around an Ethereum contract.
type ERC721Basic struct {
	ERC721BasicCaller     // Read-only binding to the contract
	ERC721BasicTransactor // Write-only binding to the contract
	ERC721BasicFilterer   // Log filterer for contract events
}

// ERC721BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721BasicSession struct {
	Contract     *ERC721Basic      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721BasicCallerSession struct {
	Contract *ERC721BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC721BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721BasicTransactorSession struct {
	Contract     *ERC721BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC721BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721BasicRaw struct {
	Contract *ERC721Basic // Generic contract binding to access the raw methods on
}

// ERC721BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721BasicCallerRaw struct {
	Contract *ERC721BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721BasicTransactorRaw struct {
	Contract *ERC721BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Basic creates a new instance of ERC721Basic, bound to a specific deployed contract.
func NewERC721Basic(address common.Address, backend bind.ContractBackend) (*ERC721Basic, error) {
	contract, err := bindERC721Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Basic{ERC721BasicCaller: ERC721BasicCaller{contract: contract}, ERC721BasicTransactor: ERC721BasicTransactor{contract: contract}, ERC721BasicFilterer: ERC721BasicFilterer{contract: contract}}, nil
}

// NewERC721BasicCaller creates a new read-only instance of ERC721Basic, bound to a specific deployed contract.
func NewERC721BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC721BasicCaller, error) {
	contract, err := bindERC721Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicCaller{contract: contract}, nil
}

// NewERC721BasicTransactor creates a new write-only instance of ERC721Basic, bound to a specific deployed contract.
func NewERC721BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721BasicTransactor, error) {
	contract, err := bindERC721Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTransactor{contract: contract}, nil
}

// NewERC721BasicFilterer creates a new log filterer instance of ERC721Basic, bound to a specific deployed contract.
func NewERC721BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721BasicFilterer, error) {
	contract, err := bindERC721Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicFilterer{contract: contract}, nil
}

// bindERC721Basic binds a generic wrapper to an already deployed contract.
func bindERC721Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Basic *ERC721BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Basic.Contract.ERC721BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Basic *ERC721BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Basic.Contract.ERC721BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Basic *ERC721BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Basic.Contract.ERC721BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Basic *ERC721BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Basic *ERC721BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Basic *ERC721BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_ERC721Basic *ERC721BasicCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_ERC721Basic *ERC721BasicSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Basic.Contract.BalanceOf(&_ERC721Basic.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_ERC721Basic *ERC721BasicCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Basic.Contract.BalanceOf(&_ERC721Basic.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool _exists)
func (_ERC721Basic *ERC721BasicCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool _exists)
func (_ERC721Basic *ERC721BasicSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Basic.Contract.Exists(&_ERC721Basic.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool _exists)
func (_ERC721Basic *ERC721BasicCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Basic.Contract.Exists(&_ERC721Basic.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _operator)
func (_ERC721Basic *ERC721BasicCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _operator)
func (_ERC721Basic *ERC721BasicSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Basic.Contract.GetApproved(&_ERC721Basic.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _operator)
func (_ERC721Basic *ERC721BasicCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Basic.Contract.GetApproved(&_ERC721Basic.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721Basic *ERC721BasicCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721Basic *ERC721BasicSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Basic.Contract.IsApprovedForAll(&_ERC721Basic.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721Basic *ERC721BasicCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Basic.Contract.IsApprovedForAll(&_ERC721Basic.CallOpts, _owner, _operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_ERC721Basic *ERC721BasicCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_ERC721Basic *ERC721BasicSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Basic.Contract.OwnerOf(&_ERC721Basic.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_ERC721Basic *ERC721BasicCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Basic.Contract.OwnerOf(&_ERC721Basic.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721Basic *ERC721BasicCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721Basic *ERC721BasicSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Basic.Contract.SupportsInterface(&_ERC721Basic.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721Basic *ERC721BasicCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Basic.Contract.SupportsInterface(&_ERC721Basic.CallOpts, _interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Basic *ERC721BasicTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Basic *ERC721BasicSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.Contract.Approve(&_ERC721Basic.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Basic *ERC721BasicTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.Contract.Approve(&_ERC721Basic.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721Basic *ERC721BasicTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Basic.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721Basic *ERC721BasicSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Basic.Contract.SafeTransferFrom(&_ERC721Basic.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721Basic *ERC721BasicTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Basic.Contract.SafeTransferFrom(&_ERC721Basic.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC721Basic *ERC721BasicTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Basic.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC721Basic *ERC721BasicSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Basic.Contract.SetApprovalForAll(&_ERC721Basic.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC721Basic *ERC721BasicTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Basic.Contract.SetApprovalForAll(&_ERC721Basic.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Basic *ERC721BasicTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Basic *ERC721BasicSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.Contract.TransferFrom(&_ERC721Basic.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Basic *ERC721BasicTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.Contract.TransferFrom(&_ERC721Basic.TransactOpts, _from, _to, _tokenId)
}

// ERC721BasicApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Basic contract.
type ERC721BasicApprovalIterator struct {
	Event *ERC721BasicApproval // Event containing the contract specifics and raw log

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
func (it *ERC721BasicApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicApproval)
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
		it.Event = new(ERC721BasicApproval)
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
func (it *ERC721BasicApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicApproval represents a Approval event raised by the ERC721Basic contract.
type ERC721BasicApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_ERC721Basic *ERC721BasicFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ERC721BasicApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Basic.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicApprovalIterator{contract: _ERC721Basic.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_ERC721Basic *ERC721BasicFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721BasicApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Basic.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicApproval)
				if err := _ERC721Basic.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721BasicApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Basic contract.
type ERC721BasicApprovalForAllIterator struct {
	Event *ERC721BasicApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721BasicApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicApprovalForAll)
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
		it.Event = new(ERC721BasicApprovalForAll)
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
func (it *ERC721BasicApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicApprovalForAll represents a ApprovalForAll event raised by the ERC721Basic contract.
type ERC721BasicApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ERC721Basic *ERC721BasicFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721BasicApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Basic.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicApprovalForAllIterator{contract: _ERC721Basic.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ERC721Basic *ERC721BasicFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721BasicApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Basic.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicApprovalForAll)
				if err := _ERC721Basic.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Basic contract.
type ERC721BasicTransferIterator struct {
	Event *ERC721BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicTransfer)
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
		it.Event = new(ERC721BasicTransfer)
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
func (it *ERC721BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicTransfer represents a Transfer event raised by the ERC721Basic contract.
type ERC721BasicTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ERC721Basic *ERC721BasicFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ERC721BasicTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Basic.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTransferIterator{contract: _ERC721Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ERC721Basic *ERC721BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721BasicTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Basic.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicTransfer)
				if err := _ERC721Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721BasicTokenABI is the input ABI used to generate the binding from.
const ERC721BasicTokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InterfaceId_ERC165\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721BasicTokenBin is the compiled bytecode used for deploying new contracts.
const ERC721BasicTokenBin = `0x608060405234801561001057600080fd5b506100437f01ffc9a7000000000000000000000000000000000000000000000000000000006401000000006100ac810204565b6100757f80ac58cd000000000000000000000000000000000000000000000000000000006401000000006100ac810204565b6100a77f4f558e79000000000000000000000000000000000000000000000000000000006401000000006100ac810204565b610118565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156100db57600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b6109d5806101276000396000f3006080604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a781146100be578063081812fc146100f4578063095ea7b31461012857806319fa8f501461014e57806323b872dd1461018057806342842e0e146101aa5780634f558e79146101d45780636352211e146101ec57806370a0823114610204578063a22cb46514610237578063b88d4fde1461025d578063e985e9c5146102cc575b600080fd5b3480156100ca57600080fd5b506100e0600160e060020a0319600435166102f3565b604080519115158252519081900360200190f35b34801561010057600080fd5b5061010c600435610312565b60408051600160a060020a039092168252519081900360200190f35b34801561013457600080fd5b5061014c600160a060020a036004351660243561032d565b005b34801561015a57600080fd5b506101636103e3565b60408051600160e060020a03199092168252519081900360200190f35b34801561018c57600080fd5b5061014c600160a060020a0360043581169060243516604435610407565b3480156101b657600080fd5b5061014c600160a060020a03600435811690602435166044356104aa565b3480156101e057600080fd5b506100e06004356104cb565b3480156101f857600080fd5b5061010c6004356104e8565b34801561021057600080fd5b50610225600160a060020a0360043516610512565b60408051918252519081900360200190f35b34801561024357600080fd5b5061014c600160a060020a03600435166024351515610545565b34801561026957600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261014c94600160a060020a0381358116956024803590921695604435953695608494019181908401838280828437509497506105c99650505050505050565b3480156102d857600080fd5b506100e0600160a060020a03600435811690602435166105f1565b600160e060020a03191660009081526020819052604090205460ff1690565b600090815260026020526040902054600160a060020a031690565b6000610338826104e8565b9050600160a060020a03838116908216141561035357600080fd5b33600160a060020a038216148061036f575061036f81336105f1565b151561037a57600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081565b610411338261061f565b151561041c57600080fd5b600160a060020a038316151561043157600080fd5b600160a060020a038216151561044657600080fd5b610450838261067e565b61045a83826106ef565b6104648282610785565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b6104c683838360206040519081016040528060008152506105c9565b505050565b600090815260016020526040902054600160a060020a0316151590565b600081815260016020526040812054600160a060020a031680151561050c57600080fd5b92915050565b6000600160a060020a038216151561052957600080fd5b50600160a060020a031660009081526003602052604090205490565b600160a060020a03821633141561055b57600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6105d4848484610407565b6105e084848484610815565b15156105eb57600080fd5b50505050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b60008061062b836104e8565b905080600160a060020a031684600160a060020a03161480610666575083600160a060020a031661065b84610312565b600160a060020a0316145b80610676575061067681856105f1565b949350505050565b81600160a060020a0316610691826104e8565b600160a060020a0316146106a457600080fd5b600081815260026020526040902054600160a060020a0316156106eb576000818152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050565b81600160a060020a0316610702826104e8565b600160a060020a03161461071557600080fd5b600160a060020a03821660009081526003602052604090205461073f90600163ffffffff61098216565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b600081815260016020526040902054600160a060020a0316156107a757600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03881690811790915584526003909152909120546107f591610994565b600160a060020a0390921660009081526003602052604090209190915550565b60008061082a85600160a060020a03166109a1565b15156108395760019150610979565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b838110156108cc5781810151838201526020016108b4565b50505050905090810190601f1680156108f95780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561091b57600080fd5b505af115801561092f573d6000803e3d6000fd5b505050506040513d602081101561094557600080fd5b5051600160e060020a031981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b60008282111561098e57fe5b50900390565b8181018281101561050c57fe5b6000903b11905600a165627a7a7230582062901e210f7af0f0f1d93cfb7abf74c72eb1e25b60f21aaf162dbcb68daf2f890029`

// DeployERC721BasicToken deploys a new Ethereum contract, binding an instance of ERC721BasicToken to it.
func DeployERC721BasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721BasicToken{ERC721BasicTokenCaller: ERC721BasicTokenCaller{contract: contract}, ERC721BasicTokenTransactor: ERC721BasicTokenTransactor{contract: contract}, ERC721BasicTokenFilterer: ERC721BasicTokenFilterer{contract: contract}}, nil
}

// ERC721BasicToken is an auto generated Go binding around an Ethereum contract.
type ERC721BasicToken struct {
	ERC721BasicTokenCaller     // Read-only binding to the contract
	ERC721BasicTokenTransactor // Write-only binding to the contract
	ERC721BasicTokenFilterer   // Log filterer for contract events
}

// ERC721BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721BasicTokenSession struct {
	Contract     *ERC721BasicToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721BasicTokenCallerSession struct {
	Contract *ERC721BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC721BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721BasicTokenTransactorSession struct {
	Contract     *ERC721BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC721BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721BasicTokenRaw struct {
	Contract *ERC721BasicToken // Generic contract binding to access the raw methods on
}

// ERC721BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721BasicTokenCallerRaw struct {
	Contract *ERC721BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721BasicTokenTransactorRaw struct {
	Contract *ERC721BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721BasicToken creates a new instance of ERC721BasicToken, bound to a specific deployed contract.
func NewERC721BasicToken(address common.Address, backend bind.ContractBackend) (*ERC721BasicToken, error) {
	contract, err := bindERC721BasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicToken{ERC721BasicTokenCaller: ERC721BasicTokenCaller{contract: contract}, ERC721BasicTokenTransactor: ERC721BasicTokenTransactor{contract: contract}, ERC721BasicTokenFilterer: ERC721BasicTokenFilterer{contract: contract}}, nil
}

// NewERC721BasicTokenCaller creates a new read-only instance of ERC721BasicToken, bound to a specific deployed contract.
func NewERC721BasicTokenCaller(address common.Address, caller bind.ContractCaller) (*ERC721BasicTokenCaller, error) {
	contract, err := bindERC721BasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenCaller{contract: contract}, nil
}

// NewERC721BasicTokenTransactor creates a new write-only instance of ERC721BasicToken, bound to a specific deployed contract.
func NewERC721BasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721BasicTokenTransactor, error) {
	contract, err := bindERC721BasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenTransactor{contract: contract}, nil
}

// NewERC721BasicTokenFilterer creates a new log filterer instance of ERC721BasicToken, bound to a specific deployed contract.
func NewERC721BasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721BasicTokenFilterer, error) {
	contract, err := bindERC721BasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenFilterer{contract: contract}, nil
}

// bindERC721BasicToken binds a generic wrapper to an already deployed contract.
func bindERC721BasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721BasicToken *ERC721BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721BasicToken.Contract.ERC721BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721BasicToken *ERC721BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.ERC721BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721BasicToken *ERC721BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.ERC721BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721BasicToken *ERC721BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721BasicToken *ERC721BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721BasicToken *ERC721BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.contract.Transact(opts, method, params...)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721BasicToken *ERC721BasicTokenCaller) InterfaceIdERC165(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "InterfaceId_ERC165")
	return *ret0, err
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721BasicToken *ERC721BasicTokenSession) InterfaceIdERC165() ([4]byte, error) {
	return _ERC721BasicToken.Contract.InterfaceIdERC165(&_ERC721BasicToken.CallOpts)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) InterfaceIdERC165() ([4]byte, error) {
	return _ERC721BasicToken.Contract.InterfaceIdERC165(&_ERC721BasicToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_ERC721BasicToken *ERC721BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_ERC721BasicToken *ERC721BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721BasicToken.Contract.BalanceOf(&_ERC721BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721BasicToken.Contract.BalanceOf(&_ERC721BasicToken.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721BasicToken.Contract.Exists(&_ERC721BasicToken.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721BasicToken.Contract.Exists(&_ERC721BasicToken.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721BasicToken.Contract.GetApproved(&_ERC721BasicToken.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721BasicToken.Contract.GetApproved(&_ERC721BasicToken.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721BasicToken.Contract.IsApprovedForAll(&_ERC721BasicToken.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721BasicToken.Contract.IsApprovedForAll(&_ERC721BasicToken.CallOpts, _owner, _operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721BasicToken.Contract.OwnerOf(&_ERC721BasicToken.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721BasicToken.Contract.OwnerOf(&_ERC721BasicToken.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721BasicToken.Contract.SupportsInterface(&_ERC721BasicToken.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721BasicToken.Contract.SupportsInterface(&_ERC721BasicToken.CallOpts, _interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721BasicToken *ERC721BasicTokenSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.Approve(&_ERC721BasicToken.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.Approve(&_ERC721BasicToken.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721BasicToken.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721BasicToken *ERC721BasicTokenSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.SafeTransferFrom(&_ERC721BasicToken.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.SafeTransferFrom(&_ERC721BasicToken.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721BasicToken.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_ERC721BasicToken *ERC721BasicTokenSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.SetApprovalForAll(&_ERC721BasicToken.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.SetApprovalForAll(&_ERC721BasicToken.TransactOpts, _to, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721BasicToken *ERC721BasicTokenSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.TransferFrom(&_ERC721BasicToken.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.TransferFrom(&_ERC721BasicToken.TransactOpts, _from, _to, _tokenId)
}

// ERC721BasicTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721BasicToken contract.
type ERC721BasicTokenApprovalIterator struct {
	Event *ERC721BasicTokenApproval // Event containing the contract specifics and raw log

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
func (it *ERC721BasicTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicTokenApproval)
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
		it.Event = new(ERC721BasicTokenApproval)
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
func (it *ERC721BasicTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicTokenApproval represents a Approval event raised by the ERC721BasicToken contract.
type ERC721BasicTokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ERC721BasicTokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenApprovalIterator{contract: _ERC721BasicToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721BasicTokenApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicTokenApproval)
				if err := _ERC721BasicToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721BasicTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721BasicToken contract.
type ERC721BasicTokenApprovalForAllIterator struct {
	Event *ERC721BasicTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721BasicTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicTokenApprovalForAll)
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
		it.Event = new(ERC721BasicTokenApprovalForAll)
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
func (it *ERC721BasicTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicTokenApprovalForAll represents a ApprovalForAll event raised by the ERC721BasicToken contract.
type ERC721BasicTokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721BasicTokenApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenApprovalForAllIterator{contract: _ERC721BasicToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721BasicTokenApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicTokenApprovalForAll)
				if err := _ERC721BasicToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721BasicToken contract.
type ERC721BasicTokenTransferIterator struct {
	Event *ERC721BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicTokenTransfer)
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
		it.Event = new(ERC721BasicTokenTransfer)
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
func (it *ERC721BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicTokenTransfer represents a Transfer event raised by the ERC721BasicToken contract.
type ERC721BasicTokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ERC721BasicTokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenTransferIterator{contract: _ERC721BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721BasicTokenTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicTokenTransfer)
				if err := _ERC721BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721EnumerableABI is the input ABI used to generate the binding from.
const ERC721EnumerableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"_exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721EnumerableBin is the compiled bytecode used for deploying new contracts.
const ERC721EnumerableBin = `0x`

// DeployERC721Enumerable deploys a new Ethereum contract, binding an instance of ERC721Enumerable to it.
func DeployERC721Enumerable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Enumerable, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721EnumerableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721EnumerableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Enumerable{ERC721EnumerableCaller: ERC721EnumerableCaller{contract: contract}, ERC721EnumerableTransactor: ERC721EnumerableTransactor{contract: contract}, ERC721EnumerableFilterer: ERC721EnumerableFilterer{contract: contract}}, nil
}

// ERC721Enumerable is an auto generated Go binding around an Ethereum contract.
type ERC721Enumerable struct {
	ERC721EnumerableCaller     // Read-only binding to the contract
	ERC721EnumerableTransactor // Write-only binding to the contract
	ERC721EnumerableFilterer   // Log filterer for contract events
}

// ERC721EnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721EnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721EnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721EnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721EnumerableSession struct {
	Contract     *ERC721Enumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721EnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721EnumerableCallerSession struct {
	Contract *ERC721EnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC721EnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721EnumerableTransactorSession struct {
	Contract     *ERC721EnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC721EnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721EnumerableRaw struct {
	Contract *ERC721Enumerable // Generic contract binding to access the raw methods on
}

// ERC721EnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721EnumerableCallerRaw struct {
	Contract *ERC721EnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721EnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721EnumerableTransactorRaw struct {
	Contract *ERC721EnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Enumerable creates a new instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721Enumerable(address common.Address, backend bind.ContractBackend) (*ERC721Enumerable, error) {
	contract, err := bindERC721Enumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Enumerable{ERC721EnumerableCaller: ERC721EnumerableCaller{contract: contract}, ERC721EnumerableTransactor: ERC721EnumerableTransactor{contract: contract}, ERC721EnumerableFilterer: ERC721EnumerableFilterer{contract: contract}}, nil
}

// NewERC721EnumerableCaller creates a new read-only instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableCaller(address common.Address, caller bind.ContractCaller) (*ERC721EnumerableCaller, error) {
	contract, err := bindERC721Enumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableCaller{contract: contract}, nil
}

// NewERC721EnumerableTransactor creates a new write-only instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721EnumerableTransactor, error) {
	contract, err := bindERC721Enumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableTransactor{contract: contract}, nil
}

// NewERC721EnumerableFilterer creates a new log filterer instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721EnumerableFilterer, error) {
	contract, err := bindERC721Enumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableFilterer{contract: contract}, nil
}

// bindERC721Enumerable binds a generic wrapper to an already deployed contract.
func bindERC721Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721EnumerableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Enumerable *ERC721EnumerableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Enumerable.Contract.ERC721EnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Enumerable *ERC721EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Enumerable *ERC721EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Enumerable *ERC721EnumerableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Enumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_ERC721Enumerable *ERC721EnumerableCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_ERC721Enumerable *ERC721EnumerableSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool _exists)
func (_ERC721Enumerable *ERC721EnumerableCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool _exists)
func (_ERC721Enumerable *ERC721EnumerableSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Enumerable.Contract.Exists(&_ERC721Enumerable.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool _exists)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Enumerable.Contract.Exists(&_ERC721Enumerable.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _operator)
func (_ERC721Enumerable *ERC721EnumerableCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _operator)
func (_ERC721Enumerable *ERC721EnumerableSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _operator)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts, _owner, _operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_ERC721Enumerable *ERC721EnumerableCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_ERC721Enumerable *ERC721EnumerableSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Enumerable.Contract.SupportsInterface(&_ERC721Enumerable.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Enumerable.Contract.SupportsInterface(&_ERC721Enumerable.CallOpts, _interfaceId)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256 _tokenId)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256 _tokenId)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256 _tokenId)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts, _owner, _index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TotalSupply() (*big.Int, error) {
	return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts, _from, _to, _tokenId)
}

// ERC721EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalIterator struct {
	Event *ERC721EnumerableApproval // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableApproval)
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
		it.Event = new(ERC721EnumerableApproval)
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
func (it *ERC721EnumerableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableApproval represents a Approval event raised by the ERC721Enumerable contract.
type ERC721EnumerableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ERC721EnumerableApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableApprovalIterator{contract: _ERC721Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableApproval)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalForAllIterator struct {
	Event *ERC721EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableApprovalForAll)
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
		it.Event = new(ERC721EnumerableApprovalForAll)
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
func (it *ERC721EnumerableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableApprovalForAll represents a ApprovalForAll event raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721EnumerableApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableApprovalForAllIterator{contract: _ERC721Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableApprovalForAll)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Enumerable contract.
type ERC721EnumerableTransferIterator struct {
	Event *ERC721EnumerableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableTransfer)
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
		it.Event = new(ERC721EnumerableTransfer)
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
func (it *ERC721EnumerableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableTransfer represents a Transfer event raised by the ERC721Enumerable contract.
type ERC721EnumerableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ERC721EnumerableTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableTransferIterator{contract: _ERC721Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableTransfer)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721MetadataABI is the input ABI used to generate the binding from.
const ERC721MetadataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"_exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721MetadataBin is the compiled bytecode used for deploying new contracts.
const ERC721MetadataBin = `0x`

// DeployERC721Metadata deploys a new Ethereum contract, binding an instance of ERC721Metadata to it.
func DeployERC721Metadata(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Metadata, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MetadataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721MetadataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Metadata{ERC721MetadataCaller: ERC721MetadataCaller{contract: contract}, ERC721MetadataTransactor: ERC721MetadataTransactor{contract: contract}, ERC721MetadataFilterer: ERC721MetadataFilterer{contract: contract}}, nil
}

// ERC721Metadata is an auto generated Go binding around an Ethereum contract.
type ERC721Metadata struct {
	ERC721MetadataCaller     // Read-only binding to the contract
	ERC721MetadataTransactor // Write-only binding to the contract
	ERC721MetadataFilterer   // Log filterer for contract events
}

// ERC721MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721MetadataSession struct {
	Contract     *ERC721Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721MetadataCallerSession struct {
	Contract *ERC721MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC721MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721MetadataTransactorSession struct {
	Contract     *ERC721MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC721MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721MetadataRaw struct {
	Contract *ERC721Metadata // Generic contract binding to access the raw methods on
}

// ERC721MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721MetadataCallerRaw struct {
	Contract *ERC721MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721MetadataTransactorRaw struct {
	Contract *ERC721MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Metadata creates a new instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721Metadata(address common.Address, backend bind.ContractBackend) (*ERC721Metadata, error) {
	contract, err := bindERC721Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Metadata{ERC721MetadataCaller: ERC721MetadataCaller{contract: contract}, ERC721MetadataTransactor: ERC721MetadataTransactor{contract: contract}, ERC721MetadataFilterer: ERC721MetadataFilterer{contract: contract}}, nil
}

// NewERC721MetadataCaller creates a new read-only instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataCaller(address common.Address, caller bind.ContractCaller) (*ERC721MetadataCaller, error) {
	contract, err := bindERC721Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataCaller{contract: contract}, nil
}

// NewERC721MetadataTransactor creates a new write-only instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721MetadataTransactor, error) {
	contract, err := bindERC721Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataTransactor{contract: contract}, nil
}

// NewERC721MetadataFilterer creates a new log filterer instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721MetadataFilterer, error) {
	contract, err := bindERC721Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataFilterer{contract: contract}, nil
}

// bindERC721Metadata binds a generic wrapper to an already deployed contract.
func bindERC721Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Metadata *ERC721MetadataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Metadata.Contract.ERC721MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Metadata *ERC721MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.ERC721MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Metadata *ERC721MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.ERC721MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Metadata *ERC721MetadataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Metadata *ERC721MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Metadata *ERC721MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_ERC721Metadata *ERC721MetadataCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_ERC721Metadata *ERC721MetadataSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Metadata.Contract.BalanceOf(&_ERC721Metadata.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_ERC721Metadata *ERC721MetadataCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Metadata.Contract.BalanceOf(&_ERC721Metadata.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool _exists)
func (_ERC721Metadata *ERC721MetadataCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool _exists)
func (_ERC721Metadata *ERC721MetadataSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Metadata.Contract.Exists(&_ERC721Metadata.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool _exists)
func (_ERC721Metadata *ERC721MetadataCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Metadata.Contract.Exists(&_ERC721Metadata.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _operator)
func (_ERC721Metadata *ERC721MetadataCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _operator)
func (_ERC721Metadata *ERC721MetadataSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.GetApproved(&_ERC721Metadata.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _operator)
func (_ERC721Metadata *ERC721MetadataCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.GetApproved(&_ERC721Metadata.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Metadata.Contract.IsApprovedForAll(&_ERC721Metadata.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Metadata.Contract.IsApprovedForAll(&_ERC721Metadata.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string _name)
func (_ERC721Metadata *ERC721MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string _name)
func (_ERC721Metadata *ERC721MetadataSession) Name() (string, error) {
	return _ERC721Metadata.Contract.Name(&_ERC721Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string _name)
func (_ERC721Metadata *ERC721MetadataCallerSession) Name() (string, error) {
	return _ERC721Metadata.Contract.Name(&_ERC721Metadata.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_ERC721Metadata *ERC721MetadataCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_ERC721Metadata *ERC721MetadataSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.OwnerOf(&_ERC721Metadata.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_ERC721Metadata *ERC721MetadataCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.OwnerOf(&_ERC721Metadata.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Metadata.Contract.SupportsInterface(&_ERC721Metadata.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Metadata.Contract.SupportsInterface(&_ERC721Metadata.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string _symbol)
func (_ERC721Metadata *ERC721MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string _symbol)
func (_ERC721Metadata *ERC721MetadataSession) Symbol() (string, error) {
	return _ERC721Metadata.Contract.Symbol(&_ERC721Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string _symbol)
func (_ERC721Metadata *ERC721MetadataCallerSession) Symbol() (string, error) {
	return _ERC721Metadata.Contract.Symbol(&_ERC721Metadata.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_ERC721Metadata *ERC721MetadataCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_ERC721Metadata *ERC721MetadataSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721Metadata.Contract.TokenURI(&_ERC721Metadata.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_ERC721Metadata *ERC721MetadataCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721Metadata.Contract.TokenURI(&_ERC721Metadata.CallOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Metadata *ERC721MetadataSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.Approve(&_ERC721Metadata.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.Approve(&_ERC721Metadata.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Metadata.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721Metadata *ERC721MetadataSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.SafeTransferFrom(&_ERC721Metadata.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.SafeTransferFrom(&_ERC721Metadata.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Metadata.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC721Metadata *ERC721MetadataSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.SetApprovalForAll(&_ERC721Metadata.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.SetApprovalForAll(&_ERC721Metadata.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Metadata *ERC721MetadataSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.TransferFrom(&_ERC721Metadata.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.TransferFrom(&_ERC721Metadata.TransactOpts, _from, _to, _tokenId)
}

// ERC721MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Metadata contract.
type ERC721MetadataApprovalIterator struct {
	Event *ERC721MetadataApproval // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataApproval)
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
		it.Event = new(ERC721MetadataApproval)
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
func (it *ERC721MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataApproval represents a Approval event raised by the ERC721Metadata contract.
type ERC721MetadataApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_ERC721Metadata *ERC721MetadataFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ERC721MetadataApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataApprovalIterator{contract: _ERC721Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_ERC721Metadata *ERC721MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721MetadataApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataApproval)
				if err := _ERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721MetadataApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Metadata contract.
type ERC721MetadataApprovalForAllIterator struct {
	Event *ERC721MetadataApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataApprovalForAll)
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
		it.Event = new(ERC721MetadataApprovalForAll)
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
func (it *ERC721MetadataApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataApprovalForAll represents a ApprovalForAll event raised by the ERC721Metadata contract.
type ERC721MetadataApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ERC721Metadata *ERC721MetadataFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721MetadataApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataApprovalForAllIterator{contract: _ERC721Metadata.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ERC721Metadata *ERC721MetadataFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721MetadataApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataApprovalForAll)
				if err := _ERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Metadata contract.
type ERC721MetadataTransferIterator struct {
	Event *ERC721MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataTransfer)
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
		it.Event = new(ERC721MetadataTransfer)
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
func (it *ERC721MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataTransfer represents a Transfer event raised by the ERC721Metadata contract.
type ERC721MetadataTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ERC721Metadata *ERC721MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ERC721MetadataTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataTransferIterator{contract: _ERC721Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ERC721Metadata *ERC721MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721MetadataTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataTransfer)
				if err := _ERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721ReceiverABI is the input ABI used to generate the binding from.
const ERC721ReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721ReceiverBin is the compiled bytecode used for deploying new contracts.
const ERC721ReceiverBin = `0x`

// DeployERC721Receiver deploys a new Ethereum contract, binding an instance of ERC721Receiver to it.
func DeployERC721Receiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Receiver, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ReceiverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721ReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Receiver{ERC721ReceiverCaller: ERC721ReceiverCaller{contract: contract}, ERC721ReceiverTransactor: ERC721ReceiverTransactor{contract: contract}, ERC721ReceiverFilterer: ERC721ReceiverFilterer{contract: contract}}, nil
}

// ERC721Receiver is an auto generated Go binding around an Ethereum contract.
type ERC721Receiver struct {
	ERC721ReceiverCaller     // Read-only binding to the contract
	ERC721ReceiverTransactor // Write-only binding to the contract
	ERC721ReceiverFilterer   // Log filterer for contract events
}

// ERC721ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721ReceiverSession struct {
	Contract     *ERC721Receiver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721ReceiverCallerSession struct {
	Contract *ERC721ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC721ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721ReceiverTransactorSession struct {
	Contract     *ERC721ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC721ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721ReceiverRaw struct {
	Contract *ERC721Receiver // Generic contract binding to access the raw methods on
}

// ERC721ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721ReceiverCallerRaw struct {
	Contract *ERC721ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721ReceiverTransactorRaw struct {
	Contract *ERC721ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Receiver creates a new instance of ERC721Receiver, bound to a specific deployed contract.
func NewERC721Receiver(address common.Address, backend bind.ContractBackend) (*ERC721Receiver, error) {
	contract, err := bindERC721Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Receiver{ERC721ReceiverCaller: ERC721ReceiverCaller{contract: contract}, ERC721ReceiverTransactor: ERC721ReceiverTransactor{contract: contract}, ERC721ReceiverFilterer: ERC721ReceiverFilterer{contract: contract}}, nil
}

// NewERC721ReceiverCaller creates a new read-only instance of ERC721Receiver, bound to a specific deployed contract.
func NewERC721ReceiverCaller(address common.Address, caller bind.ContractCaller) (*ERC721ReceiverCaller, error) {
	contract, err := bindERC721Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721ReceiverCaller{contract: contract}, nil
}

// NewERC721ReceiverTransactor creates a new write-only instance of ERC721Receiver, bound to a specific deployed contract.
func NewERC721ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721ReceiverTransactor, error) {
	contract, err := bindERC721Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721ReceiverTransactor{contract: contract}, nil
}

// NewERC721ReceiverFilterer creates a new log filterer instance of ERC721Receiver, bound to a specific deployed contract.
func NewERC721ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721ReceiverFilterer, error) {
	contract, err := bindERC721Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721ReceiverFilterer{contract: contract}, nil
}

// bindERC721Receiver binds a generic wrapper to an already deployed contract.
func bindERC721Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Receiver *ERC721ReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Receiver.Contract.ERC721ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Receiver *ERC721ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.ERC721ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Receiver *ERC721ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.ERC721ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Receiver *ERC721ReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Receiver *ERC721ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Receiver *ERC721ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_ERC721Receiver *ERC721ReceiverTransactor) OnERC721Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Receiver.contract.Transact(opts, "onERC721Received", _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_ERC721Receiver *ERC721ReceiverSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.OnERC721Received(&_ERC721Receiver.TransactOpts, _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_ERC721Receiver *ERC721ReceiverTransactorSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.OnERC721Received(&_ERC721Receiver.TransactOpts, _operator, _from, _tokenId, _data)
}

// ERC721TokenABI is the input ABI used to generate the binding from.
const ERC721TokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InterfaceId_ERC165\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721TokenBin is the compiled bytecode used for deploying new contracts.
const ERC721TokenBin = `0x60806040523480156200001157600080fd5b506040516200116e3803806200116e83398101604052805160208201519082019101620000677f01ffc9a7000000000000000000000000000000000000000000000000000000006401000000006200016b810204565b6200009b7f80ac58cd000000000000000000000000000000000000000000000000000000006401000000006200016b810204565b620000cf7f4f558e79000000000000000000000000000000000000000000000000000000006401000000006200016b810204565b8151620000e4906005906020850190620001d8565b508051620000fa906006906020840190620001d8565b506200012f7f780e9d63000000000000000000000000000000000000000000000000000000006401000000006200016b810204565b620001637f5b5e139f000000000000000000000000000000000000000000000000000000006401000000006200016b810204565b50506200027d565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156200019b57600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200021b57805160ff19168380011785556200024b565b828001600101855582156200024b579182015b828111156200024b5782518255916020019190600101906200022e565b50620002599291506200025d565b5090565b6200027a91905b8082111562000259576000815560010162000264565b90565b610ee1806200028d6000396000f3006080604052600436106100fb5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a7811461010057806306fdde0314610136578063081812fc146101c0578063095ea7b3146101f457806318160ddd1461021a57806319fa8f501461024157806323b872dd146102735780632f745c591461029d57806342842e0e146102c15780634f558e79146102eb5780634f6ccce7146103035780636352211e1461031b57806370a082311461033357806395d89b4114610354578063a22cb46514610369578063b88d4fde1461038f578063c87b56dd146103fe578063e985e9c514610416575b600080fd5b34801561010c57600080fd5b50610122600160e060020a03196004351661043d565b604080519115158252519081900360200190f35b34801561014257600080fd5b5061014b61045c565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561018557818101518382015260200161016d565b50505050905090810190601f1680156101b25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101cc57600080fd5b506101d86004356104f3565b60408051600160a060020a039092168252519081900360200190f35b34801561020057600080fd5b50610218600160a060020a036004351660243561050e565b005b34801561022657600080fd5b5061022f6105c4565b60408051918252519081900360200190f35b34801561024d57600080fd5b506102566105ca565b60408051600160e060020a03199092168252519081900360200190f35b34801561027f57600080fd5b50610218600160a060020a03600435811690602435166044356105ee565b3480156102a957600080fd5b5061022f600160a060020a0360043516602435610691565b3480156102cd57600080fd5b50610218600160a060020a03600435811690602435166044356106de565b3480156102f757600080fd5b506101226004356106ff565b34801561030f57600080fd5b5061022f60043561071c565b34801561032757600080fd5b506101d8600435610751565b34801561033f57600080fd5b5061022f600160a060020a036004351661077b565b34801561036057600080fd5b5061014b6107ae565b34801561037557600080fd5b50610218600160a060020a0360043516602435151561080f565b34801561039b57600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261021894600160a060020a0381358116956024803590921695604435953695608494019181908401838280828437509497506108939650505050505050565b34801561040a57600080fd5b5061014b6004356108bb565b34801561042257600080fd5b50610122600160a060020a0360043581169060243516610970565b600160e060020a03191660009081526020819052604090205460ff1690565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104e85780601f106104bd576101008083540402835291602001916104e8565b820191906000526020600020905b8154815290600101906020018083116104cb57829003601f168201915b505050505090505b90565b600090815260026020526040902054600160a060020a031690565b600061051982610751565b9050600160a060020a03838116908216141561053457600080fd5b33600160a060020a038216148061055057506105508133610970565b151561055b57600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b60095490565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081565b6105f8338261099e565b151561060357600080fd5b600160a060020a038316151561061857600080fd5b600160a060020a038216151561062d57600080fd5b61063783826109fd565b6106418382610a6e565b61064b8282610b75565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b600061069c8361077b565b82106106a757600080fd5b600160a060020a03831660009081526007602052604090208054839081106106cb57fe5b9060005260206000200154905092915050565b6106fa8383836020604051908101604052806000815250610893565b505050565b600090815260016020526040902054600160a060020a0316151590565b60006107266105c4565b821061073157600080fd5b600980548390811061073f57fe5b90600052602060002001549050919050565b600081815260016020526040812054600160a060020a031680151561077557600080fd5b92915050565b6000600160a060020a038216151561079257600080fd5b50600160a060020a031660009081526003602052604090205490565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104e85780601f106104bd576101008083540402835291602001916104e8565b600160a060020a03821633141561082557600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b61089e8484846105ee565b6108aa84848484610bbe565b15156108b557600080fd5b50505050565b60606108c6826106ff565b15156108d157600080fd5b6000828152600b602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156109645780601f1061093957610100808354040283529160200191610964565b820191906000526020600020905b81548152906001019060200180831161094757829003601f168201915b50505050509050919050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000806109aa83610751565b905080600160a060020a031684600160a060020a031614806109e5575083600160a060020a03166109da846104f3565b600160a060020a0316145b806109f557506109f58185610970565b949350505050565b81600160a060020a0316610a1082610751565b600160a060020a031614610a2357600080fd5b600081815260026020526040902054600160a060020a031615610a6a576000818152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050565b6000806000610a7d8585610d2b565b600084815260086020908152604080832054600160a060020a0389168452600790925290912054909350610ab890600163ffffffff610dc116565b600160a060020a038616600090815260076020526040902080549193509083908110610ae057fe5b90600052602060002001549050806007600087600160a060020a0316600160a060020a0316815260200190815260200160002084815481101515610b2057fe5b6000918252602080832090910192909255600160a060020a0387168152600790915260409020805490610b57906000198301610e78565b50600093845260086020526040808520859055908452909220555050565b6000610b818383610dd3565b50600160a060020a039091166000908152600760209081526040808320805460018101825590845282842081018590559383526008909152902055565b600080610bd385600160a060020a0316610e63565b1515610be25760019150610d22565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b83811015610c75578181015183820152602001610c5d565b50505050905090810190601f168015610ca25780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610cc457600080fd5b505af1158015610cd8573d6000803e3d6000fd5b505050506040513d6020811015610cee57600080fd5b5051600160e060020a031981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b81600160a060020a0316610d3e82610751565b600160a060020a031614610d5157600080fd5b600160a060020a038216600090815260036020526040902054610d7b90600163ffffffff610dc116565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b600082821115610dcd57fe5b50900390565b600081815260016020526040902054600160a060020a031615610df557600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0388169081179091558452600390915290912054610e4391610e6b565b600160a060020a0390921660009081526003602052604090209190915550565b6000903b1190565b8181018281101561077557fe5b8154818355818111156106fa576000838152602090206106fa9181019083016104f091905b80821115610eb15760008155600101610e9d565b50905600a165627a7a7230582044ae314e7185bcc9c05e8d0532b2a8fbd70b2792ffd4dfc5cd482f1ce0f1b63a0029`

// DeployERC721Token deploys a new Ethereum contract, binding an instance of ERC721Token to it.
func DeployERC721Token(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *ERC721Token, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721TokenBin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Token{ERC721TokenCaller: ERC721TokenCaller{contract: contract}, ERC721TokenTransactor: ERC721TokenTransactor{contract: contract}, ERC721TokenFilterer: ERC721TokenFilterer{contract: contract}}, nil
}

// ERC721Token is an auto generated Go binding around an Ethereum contract.
type ERC721Token struct {
	ERC721TokenCaller     // Read-only binding to the contract
	ERC721TokenTransactor // Write-only binding to the contract
	ERC721TokenFilterer   // Log filterer for contract events
}

// ERC721TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721TokenSession struct {
	Contract     *ERC721Token      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721TokenCallerSession struct {
	Contract *ERC721TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC721TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TokenTransactorSession struct {
	Contract     *ERC721TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC721TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721TokenRaw struct {
	Contract *ERC721Token // Generic contract binding to access the raw methods on
}

// ERC721TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721TokenCallerRaw struct {
	Contract *ERC721TokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TokenTransactorRaw struct {
	Contract *ERC721TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Token creates a new instance of ERC721Token, bound to a specific deployed contract.
func NewERC721Token(address common.Address, backend bind.ContractBackend) (*ERC721Token, error) {
	contract, err := bindERC721Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Token{ERC721TokenCaller: ERC721TokenCaller{contract: contract}, ERC721TokenTransactor: ERC721TokenTransactor{contract: contract}, ERC721TokenFilterer: ERC721TokenFilterer{contract: contract}}, nil
}

// NewERC721TokenCaller creates a new read-only instance of ERC721Token, bound to a specific deployed contract.
func NewERC721TokenCaller(address common.Address, caller bind.ContractCaller) (*ERC721TokenCaller, error) {
	contract, err := bindERC721Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenCaller{contract: contract}, nil
}

// NewERC721TokenTransactor creates a new write-only instance of ERC721Token, bound to a specific deployed contract.
func NewERC721TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721TokenTransactor, error) {
	contract, err := bindERC721Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenTransactor{contract: contract}, nil
}

// NewERC721TokenFilterer creates a new log filterer instance of ERC721Token, bound to a specific deployed contract.
func NewERC721TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721TokenFilterer, error) {
	contract, err := bindERC721Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenFilterer{contract: contract}, nil
}

// bindERC721Token binds a generic wrapper to an already deployed contract.
func bindERC721Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Token *ERC721TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Token.Contract.ERC721TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Token *ERC721TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Token.Contract.ERC721TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Token *ERC721TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Token.Contract.ERC721TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Token *ERC721TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Token *ERC721TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Token *ERC721TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Token.Contract.contract.Transact(opts, method, params...)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721Token *ERC721TokenCaller) InterfaceIdERC165(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "InterfaceId_ERC165")
	return *ret0, err
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721Token *ERC721TokenSession) InterfaceIdERC165() ([4]byte, error) {
	return _ERC721Token.Contract.InterfaceIdERC165(&_ERC721Token.CallOpts)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721Token *ERC721TokenCallerSession) InterfaceIdERC165() ([4]byte, error) {
	return _ERC721Token.Contract.InterfaceIdERC165(&_ERC721Token.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_ERC721Token *ERC721TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_ERC721Token *ERC721TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Token.Contract.BalanceOf(&_ERC721Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_ERC721Token *ERC721TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Token.Contract.BalanceOf(&_ERC721Token.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool)
func (_ERC721Token *ERC721TokenCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool)
func (_ERC721Token *ERC721TokenSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Token.Contract.Exists(&_ERC721Token.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool)
func (_ERC721Token *ERC721TokenCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721Token.Contract.Exists(&_ERC721Token.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_ERC721Token *ERC721TokenCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_ERC721Token *ERC721TokenSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Token.Contract.GetApproved(&_ERC721Token.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_ERC721Token *ERC721TokenCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Token.Contract.GetApproved(&_ERC721Token.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721Token *ERC721TokenCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721Token *ERC721TokenSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Token.Contract.IsApprovedForAll(&_ERC721Token.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_ERC721Token *ERC721TokenCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Token.Contract.IsApprovedForAll(&_ERC721Token.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721Token *ERC721TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721Token *ERC721TokenSession) Name() (string, error) {
	return _ERC721Token.Contract.Name(&_ERC721Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721Token *ERC721TokenCallerSession) Name() (string, error) {
	return _ERC721Token.Contract.Name(&_ERC721Token.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address)
func (_ERC721Token *ERC721TokenCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address)
func (_ERC721Token *ERC721TokenSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Token.Contract.OwnerOf(&_ERC721Token.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address)
func (_ERC721Token *ERC721TokenCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Token.Contract.OwnerOf(&_ERC721Token.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721Token *ERC721TokenCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721Token *ERC721TokenSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Token.Contract.SupportsInterface(&_ERC721Token.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_ERC721Token *ERC721TokenCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Token.Contract.SupportsInterface(&_ERC721Token.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721Token *ERC721TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721Token *ERC721TokenSession) Symbol() (string, error) {
	return _ERC721Token.Contract.Symbol(&_ERC721Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721Token *ERC721TokenCallerSession) Symbol() (string, error) {
	return _ERC721Token.Contract.Symbol(&_ERC721Token.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_ERC721Token *ERC721TokenCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_ERC721Token *ERC721TokenSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721Token.Contract.TokenByIndex(&_ERC721Token.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_ERC721Token *ERC721TokenCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721Token.Contract.TokenByIndex(&_ERC721Token.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256)
func (_ERC721Token *ERC721TokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256)
func (_ERC721Token *ERC721TokenSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721Token.Contract.TokenOfOwnerByIndex(&_ERC721Token.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256)
func (_ERC721Token *ERC721TokenCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721Token.Contract.TokenOfOwnerByIndex(&_ERC721Token.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_ERC721Token *ERC721TokenCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_ERC721Token *ERC721TokenSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721Token.Contract.TokenURI(&_ERC721Token.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_ERC721Token *ERC721TokenCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721Token.Contract.TokenURI(&_ERC721Token.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Token *ERC721TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Token *ERC721TokenSession) TotalSupply() (*big.Int, error) {
	return _ERC721Token.Contract.TotalSupply(&_ERC721Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Token *ERC721TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721Token.Contract.TotalSupply(&_ERC721Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Token *ERC721TokenTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Token *ERC721TokenSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.Approve(&_ERC721Token.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Token *ERC721TokenTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.Approve(&_ERC721Token.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721Token *ERC721TokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721Token *ERC721TokenSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Token.Contract.SafeTransferFrom(&_ERC721Token.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_ERC721Token *ERC721TokenTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Token.Contract.SafeTransferFrom(&_ERC721Token.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_ERC721Token *ERC721TokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_ERC721Token *ERC721TokenSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Token.Contract.SetApprovalForAll(&_ERC721Token.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_ERC721Token *ERC721TokenTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Token.Contract.SetApprovalForAll(&_ERC721Token.TransactOpts, _to, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Token *ERC721TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Token *ERC721TokenSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.TransferFrom(&_ERC721Token.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Token *ERC721TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.TransferFrom(&_ERC721Token.TransactOpts, _from, _to, _tokenId)
}

// ERC721TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Token contract.
type ERC721TokenApprovalIterator struct {
	Event *ERC721TokenApproval // Event containing the contract specifics and raw log

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
func (it *ERC721TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721TokenApproval)
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
		it.Event = new(ERC721TokenApproval)
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
func (it *ERC721TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721TokenApproval represents a Approval event raised by the ERC721Token contract.
type ERC721TokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_ERC721Token *ERC721TokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ERC721TokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Token.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenApprovalIterator{contract: _ERC721Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_ERC721Token *ERC721TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721TokenApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Token.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721TokenApproval)
				if err := _ERC721Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721TokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Token contract.
type ERC721TokenApprovalForAllIterator struct {
	Event *ERC721TokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721TokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721TokenApprovalForAll)
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
		it.Event = new(ERC721TokenApprovalForAll)
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
func (it *ERC721TokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721TokenApprovalForAll represents a ApprovalForAll event raised by the ERC721Token contract.
type ERC721TokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ERC721Token *ERC721TokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721TokenApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Token.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenApprovalForAllIterator{contract: _ERC721Token.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ERC721Token *ERC721TokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721TokenApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Token.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721TokenApprovalForAll)
				if err := _ERC721Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Token contract.
type ERC721TokenTransferIterator struct {
	Event *ERC721TokenTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721TokenTransfer)
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
		it.Event = new(ERC721TokenTransfer)
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
func (it *ERC721TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721TokenTransfer represents a Transfer event raised by the ERC721Token contract.
type ERC721TokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ERC721Token *ERC721TokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ERC721TokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenTransferIterator{contract: _ERC721Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ERC721Token *ERC721TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721TokenTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721TokenTransfer)
				if err := _ERC721Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// NftiesABI is the input ABI used to generate the binding from.
const NftiesABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InterfaceId_ERC165\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"body\",\"type\":\"uint8\"},{\"name\":\"feet\",\"type\":\"uint8\"},{\"name\":\"head\",\"type\":\"uint8\"},{\"name\":\"mouth\",\"type\":\"uint8\"},{\"name\":\"extra\",\"type\":\"uint8\"},{\"name\":\"birthBlock\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"create\",\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_body\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_feet\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_head\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_mouth\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_extra\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_birthBlock\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// NftiesBin is the compiled bytecode used for deploying new contracts.
const NftiesBin = `0x60806040523480156200001157600080fd5b5060408051808201825260068082527f4e6674696573000000000000000000000000000000000000000000000000000060208084019190915283518085019094529083527f4e465449455300000000000000000000000000000000000000000000000000009083015290620000af7f01ffc9a700000000000000000000000000000000000000000000000000000000640100000000620001b3810204565b620000e37f80ac58cd00000000000000000000000000000000000000000000000000000000640100000000620001b3810204565b620001177f4f558e7900000000000000000000000000000000000000000000000000000000640100000000620001b3810204565b81516200012c90600590602085019062000220565b5080516200014290600690602084019062000220565b50620001777f780e9d6300000000000000000000000000000000000000000000000000000000640100000000620001b3810204565b620001ab7f5b5e139f00000000000000000000000000000000000000000000000000000000640100000000620001b3810204565b5050620002c5565b7fffffffff000000000000000000000000000000000000000000000000000000008082161415620001e357600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200026357805160ff191683800117855562000293565b8280016001018555821562000293579182015b828111156200029357825182559160200191906001019062000276565b50620002a1929150620002a5565b5090565b620002c291905b80821115620002a15760008155600101620002ac565b90565b611aad80620002d56000396000f30060806040526004361061011c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a7811461012157806306fdde0314610157578063081812fc146101e1578063095ea7b31461021557806318160ddd1461023b57806319fa8f501461026257806323b872dd146102945780632f745c59146102be57806342842e0e146102e25780634f558e791461030c5780634f6ccce7146103245780636352211e1461033c57806370a08231146103545780638462151c146103755780639507d39a146103e657806395d89b4114610454578063a22cb46514610469578063b88d4fde1461048f578063c87b56dd146104fe578063e985e9c514610516578063efc81a8c1461053d575b600080fd5b34801561012d57600080fd5b50610143600160e060020a031960043516610552565b604080519115158252519081900360200190f35b34801561016357600080fd5b5061016c610571565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101a657818101518382015260200161018e565b50505050905090810190601f1680156101d35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ed57600080fd5b506101f9600435610608565b60408051600160a060020a039092168252519081900360200190f35b34801561022157600080fd5b50610239600160a060020a0360043516602435610623565b005b34801561024757600080fd5b506102506106d9565b60408051918252519081900360200190f35b34801561026e57600080fd5b506102776106df565b60408051600160e060020a03199092168252519081900360200190f35b3480156102a057600080fd5b50610239600160a060020a0360043581169060243516604435610703565b3480156102ca57600080fd5b50610250600160a060020a03600435166024356107a6565b3480156102ee57600080fd5b50610239600160a060020a03600435811690602435166044356107f3565b34801561031857600080fd5b50610143600435610814565b34801561033057600080fd5b50610250600435610831565b34801561034857600080fd5b506101f9600435610866565b34801561036057600080fd5b50610250600160a060020a0360043516610890565b34801561038157600080fd5b50610396600160a060020a03600435166108c3565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103d25781810151838201526020016103ba565b505050509050019250505060405180910390f35b3480156103f257600080fd5b506103fe60043561092f565b60408051600160a060020a03909816885260ff9687166020890152948616878601529285166060870152908416608086015290921660a084015267ffffffffffffffff90911660c0830152519081900360e00190f35b34801561046057600080fd5b5061016c610a62565b34801561047557600080fd5b50610239600160a060020a03600435166024351515610ac3565b34801561049b57600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261023994600160a060020a038135811695602480359092169560443595369560849401918190840183828082843750949750610b479650505050505050565b34801561050a57600080fd5b5061016c600435610b6f565b34801561052257600080fd5b50610143600160a060020a0360043581169060243516610c1a565b34801561054957600080fd5b50610250610c48565b600160e060020a03191660009081526020819052604090205460ff1690565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105fd5780601f106105d2576101008083540402835291602001916105fd565b820191906000526020600020905b8154815290600101906020018083116105e057829003601f168201915b505050505090505b90565b600090815260026020526040902054600160a060020a031690565b600061062e82610866565b9050600160a060020a03838116908216141561064957600080fd5b33600160a060020a038216148061066557506106658133610c1a565b151561067057600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b60095490565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081565b61070d3382610fcc565b151561071857600080fd5b600160a060020a038316151561072d57600080fd5b600160a060020a038216151561074257600080fd5b61074c838261102b565b610756838261109c565b61076082826111a3565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b60006107b183610890565b82106107bc57600080fd5b600160a060020a03831660009081526007602052604090208054839081106107e057fe5b9060005260206000200154905092915050565b61080f8383836020604051908101604052806000815250610b47565b505050565b600090815260016020526040902054600160a060020a0316151590565b600061083b6106d9565b821061084657600080fd5b600980548390811061085457fe5b90600052602060002001549050919050565b600081815260016020526040812054600160a060020a031680151561088a57600080fd5b92915050565b6000600160a060020a03821615156108a757600080fd5b50600160a060020a031660009081526003602052604090205490565b600160a060020a03811660009081526007602090815260409182902080548351818402810184019094528084526060939283018282801561092357602002820191906000526020600020905b81548152602001906001019080831161090f575b50505050509050919050565b600081815260016020526040812054600c8054839283928392839283928392600160a060020a03909216918a90811061096457fe5b600091825260209091200154600c805460ff909216918b90811061098457fe5b9060005260206000200160000160019054906101000a900460ff16600c8b8154811015156109ae57fe5b9060005260206000200160000160029054906101000a900460ff16600c8c8154811015156109d857fe5b9060005260206000200160000160039054906101000a900460ff16600c8d815481101515610a0257fe5b9060005260206000200160000160049054906101000a900460ff16600c8e815481101515610a2c57fe5b600091825260209091200154959e949d50929b50909950975095506501000000000090910467ffffffffffffffff169350915050565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105fd5780601f106105d2576101008083540402835291602001916105fd565b600160a060020a038216331415610ad957600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b610b52848484610703565b610b5e848484846111ec565b1515610b6957600080fd5b50505050565b6060610b7a82610814565b1515610b8557600080fd5b6000828152600b602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156109235780601f10610bed57610100808354040283529160200191610923565b820191906000526020600020905b815481529060010190602001808311610bfb5750939695505050505050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b60008060008060008060006060610c5d611990565b610c656106d9565b6040805160208082019390935260001943014081830152815180820383018152606090910191829052805190928291908401908083835b60208310610cbb5780518252601f199092019160209182019101610c9c565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091209a5060059250505060f860020a60008a901a81020460ff16066001908101975060059089901a60f860020a0260f860020a900460ff16811515610d2357fe5b066001019550600560f860020a60028a901a81020460ff16066001019450600560f860020a60038a901a81020460ff16066001019350600560f860020a60048a901a81020460ff16066001019250610d7e8787878787611359565b6040805160c08101825260ff808b16825289811660208301908152898216938301938452888216606084019081528883166080850190815267ffffffffffffffff43811660a08701908152600c805460018101825560009190915287517fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c78201805497519a5196519551935190941665010000000000026cffffffffffffffff0000000000199389166401000000000264ff0000000019968a1663010000000263ff00000019988b16620100000262ff0000199d8c166101000261ff001995909c1660ff19909b169a909a1793909316999099179a909a16969096179490941694909417919091169390931791909116939093179055909a509092509050610ea6338a6114b6565b610eb08983611505565b33600160a060020a03167f1ae41272ced32fa050dc3df49761e279866a7a9378212de7a61104821698f18c8a89898989898860a001518a604051808981526020018860ff1660ff1681526020018760ff1660ff1681526020018660ff1660ff1681526020018560ff1660ff1681526020018460ff1660ff1681526020018367ffffffffffffffff1667ffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610f80578181015183820152602001610f68565b50505050905090810190601f168015610fad5780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390a2505050505050505090565b600080610fd883610866565b905080600160a060020a031684600160a060020a03161480611013575083600160a060020a031661100884610608565b600160a060020a0316145b8061102357506110238185610c1a565b949350505050565b81600160a060020a031661103e82610866565b600160a060020a03161461105157600080fd5b600081815260026020526040902054600160a060020a031615611098576000818152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050565b60008060006110ab8585611538565b600084815260086020908152604080832054600160a060020a03891684526007909252909120549093506110e690600163ffffffff6115ce16565b600160a060020a03861660009081526007602052604090208054919350908390811061110e57fe5b90600052602060002001549050806007600087600160a060020a0316600160a060020a031681526020019081526020016000208481548110151561114e57fe5b6000918252602080832090910192909255600160a060020a03871681526007909152604090208054906111859060001983016119c5565b50600093845260086020526040808520859055908452909220555050565b60006111af83836115e0565b50600160a060020a039091166000908152600760209081526040808320805460018101825590845282842081018590559383526008909152902055565b60008061120185600160a060020a0316611670565b15156112105760019150611350565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b838110156112a357818101518382015260200161128b565b50505050905090810190601f1680156112d05780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b1580156112f257600080fd5b505af1158015611306573d6000803e3d6000fd5b505050506040513d602081101561131c57600080fd5b5051600160e060020a031981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b6040805180820190915260208082527f68747470733a2f2f6e66746965732e696f2f746f6b656e732f6e66746965732d9082015260609061139a8188611678565b90506113c28160408051908101604052806001815260200160f860020a602d02815250611809565b90506113ce8187611678565b90506113f68160408051908101604052806001815260200160f860020a602d02815250611809565b90506114028186611678565b905061142a8160408051908101604052806001815260200160f860020a602d02815250611809565b90506114368185611678565b905061145e8160408051908101604052806001815260200160f860020a602d02815250611809565b905061146a8184611678565b90506114ab816040805190810160405280600481526020017f2e706e6700000000000000000000000000000000000000000000000000000000815250611809565b979650505050505050565b6114c08282611928565b600980546000838152600a60205260408120829055600182018355919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af015550565b61150e82610814565b151561151957600080fd5b6000828152600b60209081526040909120825161080f928401906119e9565b81600160a060020a031661154b82610866565b600160a060020a03161461155e57600080fd5b600160a060020a03821660009081526003602052604090205461158890600163ffffffff6115ce16565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b6000828211156115da57fe5b50900390565b600081815260016020526040902054600160a060020a03161561160257600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038816908117909155845260039091529091205461165091611983565b600160a060020a0390921660009081526003602052604090209190915550565b6000903b1190565b60408051606480825260a0820190925260609190829060009081908390819083908760208201610c8080388339019050509550600094505b60ff891615611705578551600a60ff9a8b168181049b60018901989290910616955060f860020a60308701029188919081106116e857fe5b906020010190600160f860020a031916908160001a9053506116b0565b899250848351016040519080825280601f01601f191660200182016040528015611739578160200160208202803883390190505b509150600090505b825181101561179957828181518110151561175857fe5b90602001015160f860020a900460f860020a02828281518110151561177957fe5b906020010190600160f860020a031916908160001a905350600101611741565b5060005b848110156117fc57858160018703038151811015156117b857fe5b90602001015160f860020a900460f860020a0282845183018151811015156117dc57fe5b906020010190600160f860020a031916908160001a90535060010161179d565b5098975050505050505050565b606080606080606060008088955087945084518651016040519080825280601f01601f19166020018201604052801561184c578160200160208202803883390190505b50935083925060009150600090505b85518110156118b957858181518110151561187257fe5b90602001015160f860020a900460f860020a02838380600101945081518110151561189957fe5b906020010190600160f860020a031916908160001a90535060010161185b565b5060005b845181101561191b5784818151811015156118d457fe5b90602001015160f860020a900460f860020a0283838060010194508151811015156118fb57fe5b906020010190600160f860020a031916908160001a9053506001016118bd565b5090979650505050505050565b600160a060020a038216151561193d57600080fd5b61194782826111a3565b6040518190600160a060020a038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b8181018281101561088a57fe5b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b81548183558181111561080f5760008381526020902061080f918101908301611a67565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611a2a57805160ff1916838001178555611a57565b82800160010185558215611a57579182015b82811115611a57578251825591602001919060010190611a3c565b50611a63929150611a67565b5090565b61060591905b80821115611a635760008155600101611a6d5600a165627a7a723058202e1db455ccf673027962faedc7a01b1c7f7a5f69c7d060a3906020284b5ee8d40029`

// DeployNfties deploys a new Ethereum contract, binding an instance of Nfties to it.
func DeployNfties(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Nfties, error) {
	parsed, err := abi.JSON(strings.NewReader(NftiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NftiesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Nfties{NftiesCaller: NftiesCaller{contract: contract}, NftiesTransactor: NftiesTransactor{contract: contract}, NftiesFilterer: NftiesFilterer{contract: contract}}, nil
}

// Nfties is an auto generated Go binding around an Ethereum contract.
type Nfties struct {
	NftiesCaller     // Read-only binding to the contract
	NftiesTransactor // Write-only binding to the contract
	NftiesFilterer   // Log filterer for contract events
}

// NftiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftiesSession struct {
	Contract     *Nfties           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftiesCallerSession struct {
	Contract *NftiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NftiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftiesTransactorSession struct {
	Contract     *NftiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftiesRaw struct {
	Contract *Nfties // Generic contract binding to access the raw methods on
}

// NftiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftiesCallerRaw struct {
	Contract *NftiesCaller // Generic read-only contract binding to access the raw methods on
}

// NftiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftiesTransactorRaw struct {
	Contract *NftiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNfties creates a new instance of Nfties, bound to a specific deployed contract.
func NewNfties(address common.Address, backend bind.ContractBackend) (*Nfties, error) {
	contract, err := bindNfties(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nfties{NftiesCaller: NftiesCaller{contract: contract}, NftiesTransactor: NftiesTransactor{contract: contract}, NftiesFilterer: NftiesFilterer{contract: contract}}, nil
}

// NewNftiesCaller creates a new read-only instance of Nfties, bound to a specific deployed contract.
func NewNftiesCaller(address common.Address, caller bind.ContractCaller) (*NftiesCaller, error) {
	contract, err := bindNfties(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftiesCaller{contract: contract}, nil
}

// NewNftiesTransactor creates a new write-only instance of Nfties, bound to a specific deployed contract.
func NewNftiesTransactor(address common.Address, transactor bind.ContractTransactor) (*NftiesTransactor, error) {
	contract, err := bindNfties(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftiesTransactor{contract: contract}, nil
}

// NewNftiesFilterer creates a new log filterer instance of Nfties, bound to a specific deployed contract.
func NewNftiesFilterer(address common.Address, filterer bind.ContractFilterer) (*NftiesFilterer, error) {
	contract, err := bindNfties(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftiesFilterer{contract: contract}, nil
}

// bindNfties binds a generic wrapper to an already deployed contract.
func bindNfties(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NftiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nfties *NftiesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Nfties.Contract.NftiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nfties *NftiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nfties.Contract.NftiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nfties *NftiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nfties.Contract.NftiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nfties *NftiesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Nfties.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nfties *NftiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nfties.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nfties *NftiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nfties.Contract.contract.Transact(opts, method, params...)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_Nfties *NftiesCaller) InterfaceIdERC165(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "InterfaceId_ERC165")
	return *ret0, err
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_Nfties *NftiesSession) InterfaceIdERC165() ([4]byte, error) {
	return _Nfties.Contract.InterfaceIdERC165(&_Nfties.CallOpts)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_Nfties *NftiesCallerSession) InterfaceIdERC165() ([4]byte, error) {
	return _Nfties.Contract.InterfaceIdERC165(&_Nfties.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Nfties *NftiesCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Nfties *NftiesSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Nfties.Contract.BalanceOf(&_Nfties.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Nfties *NftiesCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Nfties.Contract.BalanceOf(&_Nfties.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool)
func (_Nfties *NftiesCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool)
func (_Nfties *NftiesSession) Exists(_tokenId *big.Int) (bool, error) {
	return _Nfties.Contract.Exists(&_Nfties.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) constant returns(bool)
func (_Nfties *NftiesCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _Nfties.Contract.Exists(&_Nfties.CallOpts, _tokenId)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) constant returns(address owner, uint8 body, uint8 feet, uint8 head, uint8 mouth, uint8 extra, uint64 birthBlock)
func (_Nfties *NftiesCaller) Get(opts *bind.CallOpts, _id *big.Int) (struct {
	Owner      common.Address
	Body       uint8
	Feet       uint8
	Head       uint8
	Mouth      uint8
	Extra      uint8
	BirthBlock uint64
}, error) {
	ret := new(struct {
		Owner      common.Address
		Body       uint8
		Feet       uint8
		Head       uint8
		Mouth      uint8
		Extra      uint8
		BirthBlock uint64
	})
	out := ret
	err := _Nfties.contract.Call(opts, out, "get", _id)
	return *ret, err
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) constant returns(address owner, uint8 body, uint8 feet, uint8 head, uint8 mouth, uint8 extra, uint64 birthBlock)
func (_Nfties *NftiesSession) Get(_id *big.Int) (struct {
	Owner      common.Address
	Body       uint8
	Feet       uint8
	Head       uint8
	Mouth      uint8
	Extra      uint8
	BirthBlock uint64
}, error) {
	return _Nfties.Contract.Get(&_Nfties.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) constant returns(address owner, uint8 body, uint8 feet, uint8 head, uint8 mouth, uint8 extra, uint64 birthBlock)
func (_Nfties *NftiesCallerSession) Get(_id *big.Int) (struct {
	Owner      common.Address
	Body       uint8
	Feet       uint8
	Head       uint8
	Mouth      uint8
	Extra      uint8
	BirthBlock uint64
}, error) {
	return _Nfties.Contract.Get(&_Nfties.CallOpts, _id)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_Nfties *NftiesCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_Nfties *NftiesSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Nfties.Contract.GetApproved(&_Nfties.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_Nfties *NftiesCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Nfties.Contract.GetApproved(&_Nfties.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_Nfties *NftiesCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_Nfties *NftiesSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Nfties.Contract.IsApprovedForAll(&_Nfties.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_Nfties *NftiesCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Nfties.Contract.IsApprovedForAll(&_Nfties.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Nfties *NftiesCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Nfties *NftiesSession) Name() (string, error) {
	return _Nfties.Contract.Name(&_Nfties.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Nfties *NftiesCallerSession) Name() (string, error) {
	return _Nfties.Contract.Name(&_Nfties.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address)
func (_Nfties *NftiesCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address)
func (_Nfties *NftiesSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Nfties.Contract.OwnerOf(&_Nfties.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address)
func (_Nfties *NftiesCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Nfties.Contract.OwnerOf(&_Nfties.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_Nfties *NftiesCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_Nfties *NftiesSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Nfties.Contract.SupportsInterface(&_Nfties.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_Nfties *NftiesCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Nfties.Contract.SupportsInterface(&_Nfties.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Nfties *NftiesCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Nfties *NftiesSession) Symbol() (string, error) {
	return _Nfties.Contract.Symbol(&_Nfties.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Nfties *NftiesCallerSession) Symbol() (string, error) {
	return _Nfties.Contract.Symbol(&_Nfties.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_Nfties *NftiesCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_Nfties *NftiesSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _Nfties.Contract.TokenByIndex(&_Nfties.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_Nfties *NftiesCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _Nfties.Contract.TokenByIndex(&_Nfties.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256)
func (_Nfties *NftiesCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256)
func (_Nfties *NftiesSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Nfties.Contract.TokenOfOwnerByIndex(&_Nfties.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256)
func (_Nfties *NftiesCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Nfties.Contract.TokenOfOwnerByIndex(&_Nfties.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_Nfties *NftiesCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_Nfties *NftiesSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Nfties.Contract.TokenURI(&_Nfties.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_Nfties *NftiesCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Nfties.Contract.TokenURI(&_Nfties.CallOpts, _tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) constant returns(uint256[])
func (_Nfties *NftiesCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) constant returns(uint256[])
func (_Nfties *NftiesSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Nfties.Contract.TokensOfOwner(&_Nfties.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) constant returns(uint256[])
func (_Nfties *NftiesCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Nfties.Contract.TokensOfOwner(&_Nfties.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Nfties *NftiesCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Nfties.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Nfties *NftiesSession) TotalSupply() (*big.Int, error) {
	return _Nfties.Contract.TotalSupply(&_Nfties.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Nfties *NftiesCallerSession) TotalSupply() (*big.Int, error) {
	return _Nfties.Contract.TotalSupply(&_Nfties.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Nfties *NftiesTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Nfties.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Nfties *NftiesSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Nfties.Contract.Approve(&_Nfties.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Nfties *NftiesTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Nfties.Contract.Approve(&_Nfties.TransactOpts, _to, _tokenId)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(uint256 _tokenId)
func (_Nfties *NftiesTransactor) Create(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nfties.contract.Transact(opts, "create")
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(uint256 _tokenId)
func (_Nfties *NftiesSession) Create() (*types.Transaction, error) {
	return _Nfties.Contract.Create(&_Nfties.TransactOpts)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(uint256 _tokenId)
func (_Nfties *NftiesTransactorSession) Create() (*types.Transaction, error) {
	return _Nfties.Contract.Create(&_Nfties.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Nfties *NftiesTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Nfties.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Nfties *NftiesSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Nfties.Contract.SafeTransferFrom(&_Nfties.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Nfties *NftiesTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Nfties.Contract.SafeTransferFrom(&_Nfties.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_Nfties *NftiesTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _Nfties.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_Nfties *NftiesSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _Nfties.Contract.SetApprovalForAll(&_Nfties.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_Nfties *NftiesTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _Nfties.Contract.SetApprovalForAll(&_Nfties.TransactOpts, _to, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Nfties *NftiesTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Nfties.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Nfties *NftiesSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Nfties.Contract.TransferFrom(&_Nfties.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Nfties *NftiesTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Nfties.Contract.TransferFrom(&_Nfties.TransactOpts, _from, _to, _tokenId)
}

// NftiesApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Nfties contract.
type NftiesApprovalIterator struct {
	Event *NftiesApproval // Event containing the contract specifics and raw log

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
func (it *NftiesApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftiesApproval)
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
		it.Event = new(NftiesApproval)
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
func (it *NftiesApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftiesApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftiesApproval represents a Approval event raised by the Nfties contract.
type NftiesApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Nfties *NftiesFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*NftiesApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Nfties.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftiesApprovalIterator{contract: _Nfties.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Nfties *NftiesFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NftiesApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Nfties.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftiesApproval)
				if err := _Nfties.contract.UnpackLog(event, "Approval", log); err != nil {
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

// NftiesApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Nfties contract.
type NftiesApprovalForAllIterator struct {
	Event *NftiesApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NftiesApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftiesApprovalForAll)
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
		it.Event = new(NftiesApprovalForAll)
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
func (it *NftiesApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftiesApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftiesApprovalForAll represents a ApprovalForAll event raised by the Nfties contract.
type NftiesApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Nfties *NftiesFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*NftiesApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Nfties.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &NftiesApprovalForAllIterator{contract: _Nfties.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Nfties *NftiesFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NftiesApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Nfties.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftiesApprovalForAll)
				if err := _Nfties.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// NftiesCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the Nfties contract.
type NftiesCreateIterator struct {
	Event *NftiesCreate // Event containing the contract specifics and raw log

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
func (it *NftiesCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftiesCreate)
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
		it.Event = new(NftiesCreate)
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
func (it *NftiesCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftiesCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftiesCreate represents a Create event raised by the Nfties contract.
type NftiesCreate struct {
	Id         *big.Int
	Owner      common.Address
	Body       uint8
	Feet       uint8
	Head       uint8
	Mouth      uint8
	Extra      uint8
	BirthBlock uint64
	Uri        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x1ae41272ced32fa050dc3df49761e279866a7a9378212de7a61104821698f18c.
//
// Solidity: event Create(uint256 _id, address indexed _owner, uint8 _body, uint8 _feet, uint8 _head, uint8 _mouth, uint8 _extra, uint64 _birthBlock, string _uri)
func (_Nfties *NftiesFilterer) FilterCreate(opts *bind.FilterOpts, _owner []common.Address) (*NftiesCreateIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Nfties.contract.FilterLogs(opts, "Create", _ownerRule)
	if err != nil {
		return nil, err
	}
	return &NftiesCreateIterator{contract: _Nfties.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x1ae41272ced32fa050dc3df49761e279866a7a9378212de7a61104821698f18c.
//
// Solidity: event Create(uint256 _id, address indexed _owner, uint8 _body, uint8 _feet, uint8 _head, uint8 _mouth, uint8 _extra, uint64 _birthBlock, string _uri)
func (_Nfties *NftiesFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *NftiesCreate, _owner []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Nfties.contract.WatchLogs(opts, "Create", _ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftiesCreate)
				if err := _Nfties.contract.UnpackLog(event, "Create", log); err != nil {
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

// NftiesTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Nfties contract.
type NftiesTransferIterator struct {
	Event *NftiesTransfer // Event containing the contract specifics and raw log

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
func (it *NftiesTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftiesTransfer)
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
		it.Event = new(NftiesTransfer)
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
func (it *NftiesTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftiesTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftiesTransfer represents a Transfer event raised by the Nfties contract.
type NftiesTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Nfties *NftiesFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*NftiesTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Nfties.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftiesTransferIterator{contract: _Nfties.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Nfties *NftiesFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NftiesTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Nfties.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftiesTransfer)
				if err := _Nfties.contract.UnpackLog(event, "Transfer", log); err != nil {
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
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582035b0766d1050e873b63a880abba29e9010e7d823896764aacb93248ec9bf93560029`

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

// SupportsInterfaceWithLookupABI is the input ABI used to generate the binding from.
const SupportsInterfaceWithLookupABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InterfaceId_ERC165\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SupportsInterfaceWithLookupBin is the compiled bytecode used for deploying new contracts.
const SupportsInterfaceWithLookupBin = `0x608060405234801561001057600080fd5b506100437f01ffc9a700000000000000000000000000000000000000000000000000000000640100000000610048810204565b6100b4565b7fffffffff00000000000000000000000000000000000000000000000000000000808216141561007757600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b610166806100c36000396000f30060806040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a7811461005057806319fa8f501461009b575b600080fd5b34801561005c57600080fd5b506100877bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19600435166100e2565b604080519115158252519081900360200190f35b3480156100a757600080fd5b506100b0610116565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff199092168252519081900360200190f35b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191660009081526020819052604090205460ff1690565b7f01ffc9a700000000000000000000000000000000000000000000000000000000815600a165627a7a72305820e9b10beeedf665681e0107c9895f2562e114b99fefa4ee295d46e3722e4aedcf0029`

// DeploySupportsInterfaceWithLookup deploys a new Ethereum contract, binding an instance of SupportsInterfaceWithLookup to it.
func DeploySupportsInterfaceWithLookup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SupportsInterfaceWithLookup, error) {
	parsed, err := abi.JSON(strings.NewReader(SupportsInterfaceWithLookupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SupportsInterfaceWithLookupBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SupportsInterfaceWithLookup{SupportsInterfaceWithLookupCaller: SupportsInterfaceWithLookupCaller{contract: contract}, SupportsInterfaceWithLookupTransactor: SupportsInterfaceWithLookupTransactor{contract: contract}, SupportsInterfaceWithLookupFilterer: SupportsInterfaceWithLookupFilterer{contract: contract}}, nil
}

// SupportsInterfaceWithLookup is an auto generated Go binding around an Ethereum contract.
type SupportsInterfaceWithLookup struct {
	SupportsInterfaceWithLookupCaller     // Read-only binding to the contract
	SupportsInterfaceWithLookupTransactor // Write-only binding to the contract
	SupportsInterfaceWithLookupFilterer   // Log filterer for contract events
}

// SupportsInterfaceWithLookupCaller is an auto generated read-only Go binding around an Ethereum contract.
type SupportsInterfaceWithLookupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupportsInterfaceWithLookupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SupportsInterfaceWithLookupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupportsInterfaceWithLookupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SupportsInterfaceWithLookupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupportsInterfaceWithLookupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SupportsInterfaceWithLookupSession struct {
	Contract     *SupportsInterfaceWithLookup // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SupportsInterfaceWithLookupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SupportsInterfaceWithLookupCallerSession struct {
	Contract *SupportsInterfaceWithLookupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// SupportsInterfaceWithLookupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SupportsInterfaceWithLookupTransactorSession struct {
	Contract     *SupportsInterfaceWithLookupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// SupportsInterfaceWithLookupRaw is an auto generated low-level Go binding around an Ethereum contract.
type SupportsInterfaceWithLookupRaw struct {
	Contract *SupportsInterfaceWithLookup // Generic contract binding to access the raw methods on
}

// SupportsInterfaceWithLookupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SupportsInterfaceWithLookupCallerRaw struct {
	Contract *SupportsInterfaceWithLookupCaller // Generic read-only contract binding to access the raw methods on
}

// SupportsInterfaceWithLookupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SupportsInterfaceWithLookupTransactorRaw struct {
	Contract *SupportsInterfaceWithLookupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSupportsInterfaceWithLookup creates a new instance of SupportsInterfaceWithLookup, bound to a specific deployed contract.
func NewSupportsInterfaceWithLookup(address common.Address, backend bind.ContractBackend) (*SupportsInterfaceWithLookup, error) {
	contract, err := bindSupportsInterfaceWithLookup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SupportsInterfaceWithLookup{SupportsInterfaceWithLookupCaller: SupportsInterfaceWithLookupCaller{contract: contract}, SupportsInterfaceWithLookupTransactor: SupportsInterfaceWithLookupTransactor{contract: contract}, SupportsInterfaceWithLookupFilterer: SupportsInterfaceWithLookupFilterer{contract: contract}}, nil
}

// NewSupportsInterfaceWithLookupCaller creates a new read-only instance of SupportsInterfaceWithLookup, bound to a specific deployed contract.
func NewSupportsInterfaceWithLookupCaller(address common.Address, caller bind.ContractCaller) (*SupportsInterfaceWithLookupCaller, error) {
	contract, err := bindSupportsInterfaceWithLookup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SupportsInterfaceWithLookupCaller{contract: contract}, nil
}

// NewSupportsInterfaceWithLookupTransactor creates a new write-only instance of SupportsInterfaceWithLookup, bound to a specific deployed contract.
func NewSupportsInterfaceWithLookupTransactor(address common.Address, transactor bind.ContractTransactor) (*SupportsInterfaceWithLookupTransactor, error) {
	contract, err := bindSupportsInterfaceWithLookup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SupportsInterfaceWithLookupTransactor{contract: contract}, nil
}

// NewSupportsInterfaceWithLookupFilterer creates a new log filterer instance of SupportsInterfaceWithLookup, bound to a specific deployed contract.
func NewSupportsInterfaceWithLookupFilterer(address common.Address, filterer bind.ContractFilterer) (*SupportsInterfaceWithLookupFilterer, error) {
	contract, err := bindSupportsInterfaceWithLookup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SupportsInterfaceWithLookupFilterer{contract: contract}, nil
}

// bindSupportsInterfaceWithLookup binds a generic wrapper to an already deployed contract.
func bindSupportsInterfaceWithLookup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SupportsInterfaceWithLookupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SupportsInterfaceWithLookup.Contract.SupportsInterfaceWithLookupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupportsInterfaceWithLookup.Contract.SupportsInterfaceWithLookupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupportsInterfaceWithLookup.Contract.SupportsInterfaceWithLookupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SupportsInterfaceWithLookup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupportsInterfaceWithLookup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupportsInterfaceWithLookup.Contract.contract.Transact(opts, method, params...)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupCaller) InterfaceIdERC165(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _SupportsInterfaceWithLookup.contract.Call(opts, out, "InterfaceId_ERC165")
	return *ret0, err
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupSession) InterfaceIdERC165() ([4]byte, error) {
	return _SupportsInterfaceWithLookup.Contract.InterfaceIdERC165(&_SupportsInterfaceWithLookup.CallOpts)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupCallerSession) InterfaceIdERC165() ([4]byte, error) {
	return _SupportsInterfaceWithLookup.Contract.InterfaceIdERC165(&_SupportsInterfaceWithLookup.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SupportsInterfaceWithLookup.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SupportsInterfaceWithLookup.Contract.SupportsInterface(&_SupportsInterfaceWithLookup.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) constant returns(bool)
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SupportsInterfaceWithLookup.Contract.SupportsInterface(&_SupportsInterfaceWithLookup.CallOpts, _interfaceId)
}
