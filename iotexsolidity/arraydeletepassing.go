// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arrayofstrings

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

// AABI is the input ABI used to generate the binding from.
const AABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"numbers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_numbers\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ABin is the compiled bytecode used for deploying new contracts.
const ABin = `0x608060405234801561001057600080fd5b506040516102003803806102008339810160405280510160005b815181101561006d576000828281518110151561004357fe5b6020908102909101810151825460018181018555600094855292909320909201919091550161002a565b50506101828061007e6000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416636d4ce63c8114610050578063d39fa233146100b5575b600080fd5b34801561005c57600080fd5b506100656100df565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100a1578181015183820152602001610089565b505050509050019250505060405180910390f35b3480156100c157600080fd5b506100cd600435610137565b60408051918252519081900360200190f35b6060600080548060200260200160405190810160405280929190818152602001828054801561012d57602002820191906000526020600020905b815481526020019060010190808311610119575b5050505050905090565b600080548290811061014557fe5b6000918252602090912001549050815600a165627a7a72305820cb930a35f45254ad9c5e09f2df8fe27f2bbde5892a4728ce33ff2497a8974dae0029`

// DeployA deploys a new Ethereum contract, binding an instance of A to it.
func DeployA(auth *bind.TransactOpts, backend bind.ContractBackend, _numbers []*big.Int) (common.Address, *types.Transaction, *A, error) {
	parsed, err := abi.JSON(strings.NewReader(AABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ABin), backend, _numbers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &A{ACaller: ACaller{contract: contract}, ATransactor: ATransactor{contract: contract}, AFilterer: AFilterer{contract: contract}}, nil
}

// A is an auto generated Go binding around an Ethereum contract.
type A struct {
	ACaller     // Read-only binding to the contract
	ATransactor // Write-only binding to the contract
	AFilterer   // Log filterer for contract events
}

// ACaller is an auto generated read-only Go binding around an Ethereum contract.
type ACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ASession struct {
	Contract     *A                // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ACallerSession struct {
	Contract *ACaller      // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ATransactorSession struct {
	Contract     *ATransactor      // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ARaw is an auto generated low-level Go binding around an Ethereum contract.
type ARaw struct {
	Contract *A // Generic contract binding to access the raw methods on
}

// ACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ACallerRaw struct {
	Contract *ACaller // Generic read-only contract binding to access the raw methods on
}

// ATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ATransactorRaw struct {
	Contract *ATransactor // Generic write-only contract binding to access the raw methods on
}

// NewA creates a new instance of A, bound to a specific deployed contract.
func NewA(address common.Address, backend bind.ContractBackend) (*A, error) {
	contract, err := bindA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &A{ACaller: ACaller{contract: contract}, ATransactor: ATransactor{contract: contract}, AFilterer: AFilterer{contract: contract}}, nil
}

// NewACaller creates a new read-only instance of A, bound to a specific deployed contract.
func NewACaller(address common.Address, caller bind.ContractCaller) (*ACaller, error) {
	contract, err := bindA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ACaller{contract: contract}, nil
}

// NewATransactor creates a new write-only instance of A, bound to a specific deployed contract.
func NewATransactor(address common.Address, transactor bind.ContractTransactor) (*ATransactor, error) {
	contract, err := bindA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ATransactor{contract: contract}, nil
}

// NewAFilterer creates a new log filterer instance of A, bound to a specific deployed contract.
func NewAFilterer(address common.Address, filterer bind.ContractFilterer) (*AFilterer, error) {
	contract, err := bindA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AFilterer{contract: contract}, nil
}

// bindA binds a generic wrapper to an already deployed contract.
func bindA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_A *ARaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _A.Contract.ACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_A *ARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _A.Contract.ATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_A *ARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _A.Contract.ATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_A *ACallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _A.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_A *ATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _A.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_A *ATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _A.Contract.contract.Transact(opts, method, params...)
}

// Numbers is a free data retrieval call binding the contract method 0xd39fa233.
//
// Solidity: function numbers(uint256 ) constant returns(uint256)
func (_A *ACaller) Numbers(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _A.contract.Call(opts, out, "numbers", arg0)
	return *ret0, err
}

// Numbers is a free data retrieval call binding the contract method 0xd39fa233.
//
// Solidity: function numbers(uint256 ) constant returns(uint256)
func (_A *ASession) Numbers(arg0 *big.Int) (*big.Int, error) {
	return _A.Contract.Numbers(&_A.CallOpts, arg0)
}

// Numbers is a free data retrieval call binding the contract method 0xd39fa233.
//
// Solidity: function numbers(uint256 ) constant returns(uint256)
func (_A *ACallerSession) Numbers(arg0 *big.Int) (*big.Int, error) {
	return _A.Contract.Numbers(&_A.CallOpts, arg0)
}

// Get is a paid mutator transaction binding the contract method 0x6d4ce63c.
//
// Solidity: function get() returns(uint256[])
func (_A *ATransactor) Get(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _A.contract.Transact(opts, "get")
}

// Get is a paid mutator transaction binding the contract method 0x6d4ce63c.
//
// Solidity: function get() returns(uint256[])
func (_A *ASession) Get() (*types.Transaction, error) {
	return _A.Contract.Get(&_A.TransactOpts)
}

// Get is a paid mutator transaction binding the contract method 0x6d4ce63c.
//
// Solidity: function get() returns(uint256[])
func (_A *ATransactorSession) Get() (*types.Transaction, error) {
	return _A.Contract.Get(&_A.TransactOpts)
}

// ManagerABI is the input ABI used to generate the binding from.
const ManagerABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"makeA\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getArray\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ManagerBin is the compiled bytecode used for deploying new contracts.
const ManagerBin = `0x608060405234801561001057600080fd5b50610540806100206000396000f30060806040526004361061004b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663cb8522718114610050578063d504ea1d14610077575b600080fd5b34801561005c57600080fd5b506100656100dc565b60408051918252519081900360200190f35b34801561008357600080fd5b5061008c610224565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100c85781810151838201526020016100b0565b505050509050019250505060405180910390f35b60008054600181018255818052600a7f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5639091015580808061011b610304565b60208082528254908201819052819060408201908490801561015c57602002820191906000526020600020905b815481526020019060010190808311610148575b505092505050604051809103906000f08015801561017e573d6000803e3d6000fd5b5090508073ffffffffffffffffffffffffffffffffffffffff1663d39fa23360006040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1580156101f157600080fd5b505af1158015610205573d6000803e3d6000fd5b505050506040513d602081101561021b57600080fd5b50519250505090565b600080546001818101835582805260647f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563928301558254808201845560c8908301558254808201845561012c908301558254808201845561019090830155825490810183556101f491015580546060919060029081106102a057fe5b906000526020600020016000905560008054806020026020016040519081016040528092919081815260200182805480156102fa57602002820191906000526020600020905b8154815260200190600101908083116102e6575b5050505050905090565b60405161020080610315833901905600608060405234801561001057600080fd5b506040516102003803806102008339810160405280510160005b815181101561006d576000828281518110151561004357fe5b6020908102909101810151825460018181018555600094855292909320909201919091550161002a565b50506101828061007e6000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416636d4ce63c8114610050578063d39fa233146100b5575b600080fd5b34801561005c57600080fd5b506100656100df565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100a1578181015183820152602001610089565b505050509050019250505060405180910390f35b3480156100c157600080fd5b506100cd600435610137565b60408051918252519081900360200190f35b6060600080548060200260200160405190810160405280929190818152602001828054801561012d57602002820191906000526020600020905b815481526020019060010190808311610119575b5050505050905090565b600080548290811061014557fe5b6000918252602090912001549050815600a165627a7a72305820cb930a35f45254ad9c5e09f2df8fe27f2bbde5892a4728ce33ff2497a8974dae0029a165627a7a7230582073763adcca47bd8c5b5901ac474775c59724b9bb17012d7b3909bc644226652c0029`

// DeployManager deploys a new Ethereum contract, binding an instance of Manager to it.
func DeployManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Manager, error) {
	parsed, err := abi.JSON(strings.NewReader(ManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Manager{ManagerCaller: ManagerCaller{contract: contract}, ManagerTransactor: ManagerTransactor{contract: contract}, ManagerFilterer: ManagerFilterer{contract: contract}}, nil
}

// Manager is an auto generated Go binding around an Ethereum contract.
type Manager struct {
	ManagerCaller     // Read-only binding to the contract
	ManagerTransactor // Write-only binding to the contract
	ManagerFilterer   // Log filterer for contract events
}

// ManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManagerSession struct {
	Contract     *Manager          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManagerCallerSession struct {
	Contract *ManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManagerTransactorSession struct {
	Contract     *ManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManagerRaw struct {
	Contract *Manager // Generic contract binding to access the raw methods on
}

// ManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManagerCallerRaw struct {
	Contract *ManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManagerTransactorRaw struct {
	Contract *ManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManager creates a new instance of Manager, bound to a specific deployed contract.
func NewManager(address common.Address, backend bind.ContractBackend) (*Manager, error) {
	contract, err := bindManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Manager{ManagerCaller: ManagerCaller{contract: contract}, ManagerTransactor: ManagerTransactor{contract: contract}, ManagerFilterer: ManagerFilterer{contract: contract}}, nil
}

// NewManagerCaller creates a new read-only instance of Manager, bound to a specific deployed contract.
func NewManagerCaller(address common.Address, caller bind.ContractCaller) (*ManagerCaller, error) {
	contract, err := bindManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerCaller{contract: contract}, nil
}

// NewManagerTransactor creates a new write-only instance of Manager, bound to a specific deployed contract.
func NewManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ManagerTransactor, error) {
	contract, err := bindManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerTransactor{contract: contract}, nil
}

// NewManagerFilterer creates a new log filterer instance of Manager, bound to a specific deployed contract.
func NewManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ManagerFilterer, error) {
	contract, err := bindManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManagerFilterer{contract: contract}, nil
}

// bindManager binds a generic wrapper to an already deployed contract.
func bindManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.ManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transact(opts, method, params...)
}

// GetArray is a paid mutator transaction binding the contract method 0xd504ea1d.
//
// Solidity: function getArray() returns(uint256[])
func (_Manager *ManagerTransactor) GetArray(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "getArray")
}

// GetArray is a paid mutator transaction binding the contract method 0xd504ea1d.
//
// Solidity: function getArray() returns(uint256[])
func (_Manager *ManagerSession) GetArray() (*types.Transaction, error) {
	return _Manager.Contract.GetArray(&_Manager.TransactOpts)
}

// GetArray is a paid mutator transaction binding the contract method 0xd504ea1d.
//
// Solidity: function getArray() returns(uint256[])
func (_Manager *ManagerTransactorSession) GetArray() (*types.Transaction, error) {
	return _Manager.Contract.GetArray(&_Manager.TransactOpts)
}

// MakeA is a paid mutator transaction binding the contract method 0xcb852271.
//
// Solidity: function makeA() returns(uint256)
func (_Manager *ManagerTransactor) MakeA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "makeA")
}

// MakeA is a paid mutator transaction binding the contract method 0xcb852271.
//
// Solidity: function makeA() returns(uint256)
func (_Manager *ManagerSession) MakeA() (*types.Transaction, error) {
	return _Manager.Contract.MakeA(&_Manager.TransactOpts)
}

// MakeA is a paid mutator transaction binding the contract method 0xcb852271.
//
// Solidity: function makeA() returns(uint256)
func (_Manager *ManagerTransactorSession) MakeA() (*types.Transaction, error) {
	return _Manager.Contract.MakeA(&_Manager.TransactOpts)
}
