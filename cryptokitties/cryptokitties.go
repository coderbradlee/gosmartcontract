// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cryptokitties

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

// ClockAuctionABI is the input ABI used to generate the binding from.
const ClockAuctionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"startedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuctionWhenPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_nftAddress\",\"type\":\"address\"},{\"name\":\"_cut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// ClockAuctionBin is the compiled bytecode used for deploying new contracts.
const ClockAuctionBin = `0x60806040526000805460a060020a60ff021916905534801561002057600080fd5b50604051604080610e8483398101604052805160209091015160008054600160a060020a0319163317815561271082111561005a57600080fd5b506002819055604080517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f9a20483d00000000000000000000000000000000000000000000000000000000600482015290518391600160a060020a038316916301ffc9a7916024808201926020929091908290030181600087803b1580156100e457600080fd5b505af11580156100f8573d6000803e3d6000fd5b505050506040513d602081101561010e57600080fd5b5051151561011b57600080fd5b60018054600160a060020a03909216600160a060020a03199092169190911790555050610d378061014d6000396000f3006080604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166327ebe40a81146100d45780633f4ba83a14610103578063454a2ab31461012c5780635c975abb146101375780635fd8c7101461014c57806378bd79351461016157806383b5ff8b146101ae5780638456cb59146101d5578063878eb368146101ea5780638da5cb5b1461020257806396b5a75514610233578063c55d0f561461024b578063dd1b7a0f14610263578063f2fde38b14610278575b600080fd5b3480156100e057600080fd5b50610101600435602435604435606435600160a060020a0360843516610299565b005b34801561010f57600080fd5b506101186103a6565b604080519115158252519081900360200190f35b610101600435610421565b34801561014357600080fd5b50610118610450565b34801561015857600080fd5b50610101610460565b34801561016d57600080fd5b506101796004356104bd565b60408051600160a060020a03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b3480156101ba57600080fd5b506101c3610553565b60408051918252519081900360200190f35b3480156101e157600080fd5b50610118610559565b3480156101f657600080fd5b506101016004356105d9565b34801561020e57600080fd5b50610217610646565b60408051600160a060020a039092168252519081900360200190f35b34801561023f57600080fd5b50610101600435610655565b34801561025757600080fd5b506101c360043561069f565b34801561026f57600080fd5b506102176106d1565b34801561028457600080fd5b50610101600160a060020a03600435166106e0565b6102a1610cdd565b60005460a060020a900460ff16156102b857600080fd5b6fffffffffffffffffffffffffffffffff851685146102d657600080fd5b6fffffffffffffffffffffffffffffffff841684146102f457600080fd5b67ffffffffffffffff8316831461030a57600080fd5b6103143387610733565b151561031f57600080fd5b61032933876107dc565b60a06040519081016040528083600160a060020a03168152602001866fffffffffffffffffffffffffffffffff168152602001856fffffffffffffffffffffffffffffffff1681526020018467ffffffffffffffff1681526020014267ffffffffffffffff16815250905061039e8682610864565b505050505050565b60008054600160a060020a031633146103be57600080fd5b60005460a060020a900460ff1615156103d657600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a150600190565b60005460a060020a900460ff161561043857600080fd5b61044281346109b8565b5061044d3382610ade565b50565b60005460a060020a900460ff1681565b60015460008054600160a060020a039283169216331480610489575033600160a060020a038316145b151561049457600080fd5b604051600160a060020a03831690303180156108fc02916000818181858888f150505050505050565b600081815260036020526040812081908190819081906104dc81610b4c565b15156104e757600080fd5b80546001820154600290920154600160a060020a03909116986fffffffffffffffffffffffffffffffff8084169950700100000000000000000000000000000000909304909216965067ffffffffffffffff808216965068010000000000000000909104169350915050565b60025481565b60008054600160a060020a0316331461057157600080fd5b60005460a060020a900460ff161561058857600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a150600190565b6000805460a060020a900460ff1615156105f257600080fd5b600054600160a060020a0316331461060957600080fd5b50600081815260036020526040902061062181610b4c565b151561062c57600080fd5b8054610642908390600160a060020a0316610b6d565b5050565b600054600160a060020a031681565b60008181526003602052604081209061066d82610b4c565b151561067857600080fd5b508054600160a060020a031633811461069057600080fd5b61069a8382610b6d565b505050565b60008181526003602052604081206106b681610b4c565b15156106c157600080fd5b6106ca81610bb7565b9392505050565b600154600160a060020a031681565b600054600160a060020a031633146106f757600080fd5b600160a060020a0381161561044d5760008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1990911617905550565b600154604080517f6352211e000000000000000000000000000000000000000000000000000000008152600481018490529051600092600160a060020a0380871693911691636352211e9160248082019260209290919082900301818887803b15801561079f57600080fd5b505af11580156107b3573d6000803e3d6000fd5b505050506040513d60208110156107c957600080fd5b5051600160a060020a0316149392505050565b600154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015230602483015260448201859052915191909216916323b872dd91606480830192600092919082900301818387803b15801561085057600080fd5b505af115801561039e573d6000803e3d6000fd5b603c816060015167ffffffffffffffff161015151561088257600080fd5b60008281526003602090815260409182902083518154600160a060020a0390911673ffffffffffffffffffffffffffffffffffffffff1990911617815581840151600182018054858701516fffffffffffffffffffffffffffffffff90811670010000000000000000000000000000000081029482166fffffffffffffffffffffffffffffffff19909316831790911693909317909155606080870151600290940180546080808a015167ffffffffffffffff90811668010000000000000000026fffffffffffffffff0000000000000000199190981667ffffffffffffffff1990931683171696909617909155865189815295860192909252848601929092529083015291517fa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7929181900390910190a15050565b600082815260036020526040812081808080806109d486610b4c565b15156109df57600080fd5b6109e886610bb7565b9450848810156109f757600080fd5b8554600160a060020a03169350610a0d89610c47565b6000851115610a5f57610a1f85610c94565b6040519093508386039250600160a060020a0385169083156108fc029084906000818181858888f19350505050158015610a5d573d6000803e3d6000fd5b505b5060405184880390339082156108fc029083906000818181858888f19350505050158015610a91573d6000803e3d6000fd5b50604080518a815260208101879052338183015290517f4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd29181900360600190a15092979650505050505050565b600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038581166004830152602482018590529151919092169163a9059cbb91604480830192600092919082900301818387803b15801561085057600080fd5b6002015460006801000000000000000090910467ffffffffffffffff161190565b610b7682610c47565b610b808183610ade565b6040805183815290517f2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df9181900360200190a15050565b6002810154600090819068010000000000000000900467ffffffffffffffff16421115610bfd5750600282015468010000000000000000900467ffffffffffffffff1642035b600183015460028401546106ca916fffffffffffffffffffffffffffffffff80821692700100000000000000000000000000000000909204169067ffffffffffffffff1684610ca0565b6000908152600360205260408120805473ffffffffffffffffffffffffffffffffffffffff19168155600181019190915560020180546fffffffffffffffffffffffffffffffff19169055565b60025461271091020490565b6000808080858510610cb457869350610cd2565b878703925085858402811515610cc657fe5b05915081880190508093505b505050949350505050565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152905600a165627a7a723058200d6984a66d98c11cb33ad8951249f7e800cd8cf57e7344af27cf3c34c05e3bb20029`

// DeployClockAuction deploys a new Ethereum contract, binding an instance of ClockAuction to it.
func DeployClockAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _nftAddress common.Address, _cut *big.Int) (common.Address, *types.Transaction, *ClockAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(ClockAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClockAuctionBin), backend, _nftAddress, _cut)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClockAuction{ClockAuctionCaller: ClockAuctionCaller{contract: contract}, ClockAuctionTransactor: ClockAuctionTransactor{contract: contract}, ClockAuctionFilterer: ClockAuctionFilterer{contract: contract}}, nil
}

// ClockAuction is an auto generated Go binding around an Ethereum contract.
type ClockAuction struct {
	ClockAuctionCaller     // Read-only binding to the contract
	ClockAuctionTransactor // Write-only binding to the contract
	ClockAuctionFilterer   // Log filterer for contract events
}

// ClockAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClockAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClockAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClockAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClockAuctionSession struct {
	Contract     *ClockAuction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClockAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClockAuctionCallerSession struct {
	Contract *ClockAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ClockAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClockAuctionTransactorSession struct {
	Contract     *ClockAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ClockAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClockAuctionRaw struct {
	Contract *ClockAuction // Generic contract binding to access the raw methods on
}

// ClockAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClockAuctionCallerRaw struct {
	Contract *ClockAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// ClockAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClockAuctionTransactorRaw struct {
	Contract *ClockAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClockAuction creates a new instance of ClockAuction, bound to a specific deployed contract.
func NewClockAuction(address common.Address, backend bind.ContractBackend) (*ClockAuction, error) {
	contract, err := bindClockAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClockAuction{ClockAuctionCaller: ClockAuctionCaller{contract: contract}, ClockAuctionTransactor: ClockAuctionTransactor{contract: contract}, ClockAuctionFilterer: ClockAuctionFilterer{contract: contract}}, nil
}

// NewClockAuctionCaller creates a new read-only instance of ClockAuction, bound to a specific deployed contract.
func NewClockAuctionCaller(address common.Address, caller bind.ContractCaller) (*ClockAuctionCaller, error) {
	contract, err := bindClockAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionCaller{contract: contract}, nil
}

// NewClockAuctionTransactor creates a new write-only instance of ClockAuction, bound to a specific deployed contract.
func NewClockAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*ClockAuctionTransactor, error) {
	contract, err := bindClockAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionTransactor{contract: contract}, nil
}

// NewClockAuctionFilterer creates a new log filterer instance of ClockAuction, bound to a specific deployed contract.
func NewClockAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*ClockAuctionFilterer, error) {
	contract, err := bindClockAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionFilterer{contract: contract}, nil
}

// bindClockAuction binds a generic wrapper to an already deployed contract.
func bindClockAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClockAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClockAuction *ClockAuctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClockAuction.Contract.ClockAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClockAuction *ClockAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.Contract.ClockAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClockAuction *ClockAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClockAuction.Contract.ClockAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClockAuction *ClockAuctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClockAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClockAuction *ClockAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClockAuction *ClockAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClockAuction.Contract.contract.Transact(opts, method, params...)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_ClockAuction *ClockAuctionCaller) GetAuction(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	ret := new(struct {
		Seller        common.Address
		StartingPrice *big.Int
		EndingPrice   *big.Int
		Duration      *big.Int
		StartedAt     *big.Int
	})
	out := ret
	err := _ClockAuction.contract.Call(opts, out, "getAuction", _tokenId)
	return *ret, err
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_ClockAuction *ClockAuctionSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _ClockAuction.Contract.GetAuction(&_ClockAuction.CallOpts, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_ClockAuction *ClockAuctionCallerSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _ClockAuction.Contract.GetAuction(&_ClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_ClockAuction *ClockAuctionCaller) GetCurrentPrice(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClockAuction.contract.Call(opts, out, "getCurrentPrice", _tokenId)
	return *ret0, err
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_ClockAuction *ClockAuctionSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _ClockAuction.Contract.GetCurrentPrice(&_ClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_ClockAuction *ClockAuctionCallerSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _ClockAuction.Contract.GetCurrentPrice(&_ClockAuction.CallOpts, _tokenId)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_ClockAuction *ClockAuctionCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClockAuction.contract.Call(opts, out, "nonFungibleContract")
	return *ret0, err
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_ClockAuction *ClockAuctionSession) NonFungibleContract() (common.Address, error) {
	return _ClockAuction.Contract.NonFungibleContract(&_ClockAuction.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_ClockAuction *ClockAuctionCallerSession) NonFungibleContract() (common.Address, error) {
	return _ClockAuction.Contract.NonFungibleContract(&_ClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClockAuction *ClockAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClockAuction.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClockAuction *ClockAuctionSession) Owner() (common.Address, error) {
	return _ClockAuction.Contract.Owner(&_ClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClockAuction *ClockAuctionCallerSession) Owner() (common.Address, error) {
	return _ClockAuction.Contract.Owner(&_ClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_ClockAuction *ClockAuctionCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClockAuction.contract.Call(opts, out, "ownerCut")
	return *ret0, err
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_ClockAuction *ClockAuctionSession) OwnerCut() (*big.Int, error) {
	return _ClockAuction.Contract.OwnerCut(&_ClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_ClockAuction *ClockAuctionCallerSession) OwnerCut() (*big.Int, error) {
	return _ClockAuction.Contract.OwnerCut(&_ClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ClockAuction *ClockAuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ClockAuction.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ClockAuction *ClockAuctionSession) Paused() (bool, error) {
	return _ClockAuction.Contract.Paused(&_ClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ClockAuction *ClockAuctionCallerSession) Paused() (bool, error) {
	return _ClockAuction.Contract.Paused(&_ClockAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionTransactor) Bid(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "bid", _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.Bid(&_ClockAuction.TransactOpts, _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionTransactorSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.Bid(&_ClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionTransactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "cancelAuction", _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.CancelAuction(&_ClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionTransactorSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.CancelAuction(&_ClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionTransactor) CancelAuctionWhenPaused(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "cancelAuctionWhenPaused", _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.CancelAuctionWhenPaused(&_ClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionTransactorSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.CancelAuctionWhenPaused(&_ClockAuction.TransactOpts, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_ClockAuction *ClockAuctionTransactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "createAuction", _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_ClockAuction *ClockAuctionSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _ClockAuction.Contract.CreateAuction(&_ClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_ClockAuction *ClockAuctionTransactorSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _ClockAuction.Contract.CreateAuction(&_ClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ClockAuction *ClockAuctionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ClockAuction *ClockAuctionSession) Pause() (*types.Transaction, error) {
	return _ClockAuction.Contract.Pause(&_ClockAuction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ClockAuction *ClockAuctionTransactorSession) Pause() (*types.Transaction, error) {
	return _ClockAuction.Contract.Pause(&_ClockAuction.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ClockAuction *ClockAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ClockAuction *ClockAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClockAuction.Contract.TransferOwnership(&_ClockAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ClockAuction *ClockAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClockAuction.Contract.TransferOwnership(&_ClockAuction.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ClockAuction *ClockAuctionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ClockAuction *ClockAuctionSession) Unpause() (*types.Transaction, error) {
	return _ClockAuction.Contract.Unpause(&_ClockAuction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ClockAuction *ClockAuctionTransactorSession) Unpause() (*types.Transaction, error) {
	return _ClockAuction.Contract.Unpause(&_ClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_ClockAuction *ClockAuctionTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_ClockAuction *ClockAuctionSession) WithdrawBalance() (*types.Transaction, error) {
	return _ClockAuction.Contract.WithdrawBalance(&_ClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_ClockAuction *ClockAuctionTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _ClockAuction.Contract.WithdrawBalance(&_ClockAuction.TransactOpts)
}

// ClockAuctionAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the ClockAuction contract.
type ClockAuctionAuctionCancelledIterator struct {
	Event *ClockAuctionAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *ClockAuctionAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionAuctionCancelled)
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
		it.Event = new(ClockAuctionAuctionCancelled)
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
func (it *ClockAuctionAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionAuctionCancelled represents a AuctionCancelled event raised by the ClockAuction contract.
type ClockAuctionAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_ClockAuction *ClockAuctionFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*ClockAuctionAuctionCancelledIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionAuctionCancelledIterator{contract: _ClockAuction.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_ClockAuction *ClockAuctionFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *ClockAuctionAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionAuctionCancelled)
				if err := _ClockAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ClockAuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the ClockAuction contract.
type ClockAuctionAuctionCreatedIterator struct {
	Event *ClockAuctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *ClockAuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionAuctionCreated)
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
		it.Event = new(ClockAuctionAuctionCreated)
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
func (it *ClockAuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionAuctionCreated represents a AuctionCreated event raised by the ClockAuction contract.
type ClockAuctionAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_ClockAuction *ClockAuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*ClockAuctionAuctionCreatedIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionAuctionCreatedIterator{contract: _ClockAuction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_ClockAuction *ClockAuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *ClockAuctionAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionAuctionCreated)
				if err := _ClockAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ClockAuctionAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the ClockAuction contract.
type ClockAuctionAuctionSuccessfulIterator struct {
	Event *ClockAuctionAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *ClockAuctionAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionAuctionSuccessful)
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
		it.Event = new(ClockAuctionAuctionSuccessful)
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
func (it *ClockAuctionAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionAuctionSuccessful represents a AuctionSuccessful event raised by the ClockAuction contract.
type ClockAuctionAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_ClockAuction *ClockAuctionFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*ClockAuctionAuctionSuccessfulIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionAuctionSuccessfulIterator{contract: _ClockAuction.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_ClockAuction *ClockAuctionFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *ClockAuctionAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionAuctionSuccessful)
				if err := _ClockAuction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// ClockAuctionPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the ClockAuction contract.
type ClockAuctionPauseIterator struct {
	Event *ClockAuctionPause // Event containing the contract specifics and raw log

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
func (it *ClockAuctionPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionPause)
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
		it.Event = new(ClockAuctionPause)
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
func (it *ClockAuctionPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionPause represents a Pause event raised by the ClockAuction contract.
type ClockAuctionPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_ClockAuction *ClockAuctionFilterer) FilterPause(opts *bind.FilterOpts) (*ClockAuctionPauseIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionPauseIterator{contract: _ClockAuction.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_ClockAuction *ClockAuctionFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *ClockAuctionPause) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionPause)
				if err := _ClockAuction.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ClockAuctionUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the ClockAuction contract.
type ClockAuctionUnpauseIterator struct {
	Event *ClockAuctionUnpause // Event containing the contract specifics and raw log

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
func (it *ClockAuctionUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionUnpause)
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
		it.Event = new(ClockAuctionUnpause)
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
func (it *ClockAuctionUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionUnpause represents a Unpause event raised by the ClockAuction contract.
type ClockAuctionUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_ClockAuction *ClockAuctionFilterer) FilterUnpause(opts *bind.FilterOpts) (*ClockAuctionUnpauseIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionUnpauseIterator{contract: _ClockAuction.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_ClockAuction *ClockAuctionFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *ClockAuctionUnpause) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionUnpause)
				if err := _ClockAuction.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ClockAuctionBaseABI is the input ABI used to generate the binding from.
const ClockAuctionBaseABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"}]"

// ClockAuctionBaseBin is the compiled bytecode used for deploying new contracts.
const ClockAuctionBaseBin = `0x608060405234801561001057600080fd5b5060fa8061001f6000396000f30060806040526004361060485763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166383b5ff8b8114604d578063dd1b7a0f146071575b600080fd5b348015605857600080fd5b50605f60ac565b60408051918252519081900360200190f35b348015607c57600080fd5b50608360b2565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60015481565b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820742f9ff7ac52448dfd23ee7bc2e98ebdbe40183698de7010f64d4185930726cd0029`

// DeployClockAuctionBase deploys a new Ethereum contract, binding an instance of ClockAuctionBase to it.
func DeployClockAuctionBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ClockAuctionBase, error) {
	parsed, err := abi.JSON(strings.NewReader(ClockAuctionBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClockAuctionBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClockAuctionBase{ClockAuctionBaseCaller: ClockAuctionBaseCaller{contract: contract}, ClockAuctionBaseTransactor: ClockAuctionBaseTransactor{contract: contract}, ClockAuctionBaseFilterer: ClockAuctionBaseFilterer{contract: contract}}, nil
}

// ClockAuctionBase is an auto generated Go binding around an Ethereum contract.
type ClockAuctionBase struct {
	ClockAuctionBaseCaller     // Read-only binding to the contract
	ClockAuctionBaseTransactor // Write-only binding to the contract
	ClockAuctionBaseFilterer   // Log filterer for contract events
}

// ClockAuctionBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClockAuctionBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClockAuctionBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClockAuctionBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClockAuctionBaseSession struct {
	Contract     *ClockAuctionBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClockAuctionBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClockAuctionBaseCallerSession struct {
	Contract *ClockAuctionBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ClockAuctionBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClockAuctionBaseTransactorSession struct {
	Contract     *ClockAuctionBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ClockAuctionBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClockAuctionBaseRaw struct {
	Contract *ClockAuctionBase // Generic contract binding to access the raw methods on
}

// ClockAuctionBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClockAuctionBaseCallerRaw struct {
	Contract *ClockAuctionBaseCaller // Generic read-only contract binding to access the raw methods on
}

// ClockAuctionBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClockAuctionBaseTransactorRaw struct {
	Contract *ClockAuctionBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClockAuctionBase creates a new instance of ClockAuctionBase, bound to a specific deployed contract.
func NewClockAuctionBase(address common.Address, backend bind.ContractBackend) (*ClockAuctionBase, error) {
	contract, err := bindClockAuctionBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBase{ClockAuctionBaseCaller: ClockAuctionBaseCaller{contract: contract}, ClockAuctionBaseTransactor: ClockAuctionBaseTransactor{contract: contract}, ClockAuctionBaseFilterer: ClockAuctionBaseFilterer{contract: contract}}, nil
}

// NewClockAuctionBaseCaller creates a new read-only instance of ClockAuctionBase, bound to a specific deployed contract.
func NewClockAuctionBaseCaller(address common.Address, caller bind.ContractCaller) (*ClockAuctionBaseCaller, error) {
	contract, err := bindClockAuctionBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseCaller{contract: contract}, nil
}

// NewClockAuctionBaseTransactor creates a new write-only instance of ClockAuctionBase, bound to a specific deployed contract.
func NewClockAuctionBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ClockAuctionBaseTransactor, error) {
	contract, err := bindClockAuctionBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseTransactor{contract: contract}, nil
}

// NewClockAuctionBaseFilterer creates a new log filterer instance of ClockAuctionBase, bound to a specific deployed contract.
func NewClockAuctionBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ClockAuctionBaseFilterer, error) {
	contract, err := bindClockAuctionBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseFilterer{contract: contract}, nil
}

// bindClockAuctionBase binds a generic wrapper to an already deployed contract.
func bindClockAuctionBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClockAuctionBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClockAuctionBase *ClockAuctionBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClockAuctionBase.Contract.ClockAuctionBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClockAuctionBase *ClockAuctionBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuctionBase.Contract.ClockAuctionBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClockAuctionBase *ClockAuctionBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClockAuctionBase.Contract.ClockAuctionBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClockAuctionBase *ClockAuctionBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClockAuctionBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClockAuctionBase *ClockAuctionBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuctionBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClockAuctionBase *ClockAuctionBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClockAuctionBase.Contract.contract.Transact(opts, method, params...)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_ClockAuctionBase *ClockAuctionBaseCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClockAuctionBase.contract.Call(opts, out, "nonFungibleContract")
	return *ret0, err
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_ClockAuctionBase *ClockAuctionBaseSession) NonFungibleContract() (common.Address, error) {
	return _ClockAuctionBase.Contract.NonFungibleContract(&_ClockAuctionBase.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_ClockAuctionBase *ClockAuctionBaseCallerSession) NonFungibleContract() (common.Address, error) {
	return _ClockAuctionBase.Contract.NonFungibleContract(&_ClockAuctionBase.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_ClockAuctionBase *ClockAuctionBaseCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClockAuctionBase.contract.Call(opts, out, "ownerCut")
	return *ret0, err
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_ClockAuctionBase *ClockAuctionBaseSession) OwnerCut() (*big.Int, error) {
	return _ClockAuctionBase.Contract.OwnerCut(&_ClockAuctionBase.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_ClockAuctionBase *ClockAuctionBaseCallerSession) OwnerCut() (*big.Int, error) {
	return _ClockAuctionBase.Contract.OwnerCut(&_ClockAuctionBase.CallOpts)
}

// ClockAuctionBaseAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionCancelledIterator struct {
	Event *ClockAuctionBaseAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *ClockAuctionBaseAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionBaseAuctionCancelled)
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
		it.Event = new(ClockAuctionBaseAuctionCancelled)
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
func (it *ClockAuctionBaseAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionBaseAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionBaseAuctionCancelled represents a AuctionCancelled event raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*ClockAuctionBaseAuctionCancelledIterator, error) {

	logs, sub, err := _ClockAuctionBase.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseAuctionCancelledIterator{contract: _ClockAuctionBase.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *ClockAuctionBaseAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _ClockAuctionBase.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionBaseAuctionCancelled)
				if err := _ClockAuctionBase.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ClockAuctionBaseAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionCreatedIterator struct {
	Event *ClockAuctionBaseAuctionCreated // Event containing the contract specifics and raw log

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
func (it *ClockAuctionBaseAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionBaseAuctionCreated)
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
		it.Event = new(ClockAuctionBaseAuctionCreated)
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
func (it *ClockAuctionBaseAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionBaseAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionBaseAuctionCreated represents a AuctionCreated event raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*ClockAuctionBaseAuctionCreatedIterator, error) {

	logs, sub, err := _ClockAuctionBase.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseAuctionCreatedIterator{contract: _ClockAuctionBase.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *ClockAuctionBaseAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _ClockAuctionBase.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionBaseAuctionCreated)
				if err := _ClockAuctionBase.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ClockAuctionBaseAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionSuccessfulIterator struct {
	Event *ClockAuctionBaseAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *ClockAuctionBaseAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionBaseAuctionSuccessful)
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
		it.Event = new(ClockAuctionBaseAuctionSuccessful)
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
func (it *ClockAuctionBaseAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionBaseAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionBaseAuctionSuccessful represents a AuctionSuccessful event raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*ClockAuctionBaseAuctionSuccessfulIterator, error) {

	logs, sub, err := _ClockAuctionBase.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseAuctionSuccessfulIterator{contract: _ClockAuctionBase.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *ClockAuctionBaseAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _ClockAuctionBase.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionBaseAuctionSuccessful)
				if err := _ClockAuctionBase.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// ERC721ABI is the input ABI used to generate the binding from.
const ERC721ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"total\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

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
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
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
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC721 *ERC721Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC721 *ERC721CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, _owner)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
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
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_ERC721 *ERC721Session) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_ERC721 *ERC721CallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_ERC721 *ERC721Caller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "supportsInterface", _interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_ERC721 *ERC721Session) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_ERC721 *ERC721CallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, _interfaceID)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(total uint256)
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
// Solidity: function totalSupply() constant returns(total uint256)
func (_ERC721 *ERC721Session) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(total uint256)
func (_ERC721 *ERC721CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Session) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721 *ERC721TransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Session) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Transfer(&_ERC721.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_ERC721 *ERC721TransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Transfer(&_ERC721.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Session) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
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
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts) (*ERC721ApprovalIterator, error) {

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval) (event.Subscription, error) {

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval")
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
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts) (*ERC721TransferIterator, error) {

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer) (event.Subscription, error) {

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer")
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

// ERC721MetadataABI is the input ABI used to generate the binding from.
const ERC721MetadataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"getMetadata\",\"outputs\":[{\"name\":\"buffer\",\"type\":\"bytes32[4]\"},{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC721MetadataBin is the compiled bytecode used for deploying new contracts.
const ERC721MetadataBin = `0x608060405234801561001057600080fd5b5061026e806100206000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663cb4799f28114610045575b600080fd5b34801561005157600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100a39583359536956044949193909101919081908401838280828437509497506100e29650505050505050565b6040518083608080838360005b838110156100c85781810151838201526020016100b0565b505050509050018281526020019250505060405180910390f35b6100ea610223565b6000836001141561012057507f48656c6c6f20576f726c6421203a4400000000000000000000000000000000008152600f61021c565b836002141561017a57507f4920776f756c6420646566696e6974656c792063686f6f73652061206d65646981527f756d206c656e67746820737472696e672e0000000000000000000000000000006020820152603161021c565b836003141561021c57507f4c6f72656d20697073756d20646f6c6f722073697420616d65742c206d69206581527f737420616363756d73616e2064617069627573206175677565206c6f72656d2c60208201527f2074726973746971756520766573746962756c756d2069642c206c696265726f60408201527f207375736369706974207661726975732073617069656e20616c697175616d2e606082015260805b9250929050565b60806040519081016040528060049060208202803883395091929150505600a165627a7a72305820f8b37c482208c7267d69699cbcb8432265a00ad18e5ccfe10528d840a14155d40029`

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

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(_tokenId uint256,  string) constant returns(buffer bytes32[4], count uint256)
func (_ERC721Metadata *ERC721MetadataCaller) GetMetadata(opts *bind.CallOpts, _tokenId *big.Int, arg1 string) (struct {
	Buffer [4][32]byte
	Count  *big.Int
}, error) {
	ret := new(struct {
		Buffer [4][32]byte
		Count  *big.Int
	})
	out := ret
	err := _ERC721Metadata.contract.Call(opts, out, "getMetadata", _tokenId, arg1)
	return *ret, err
}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(_tokenId uint256,  string) constant returns(buffer bytes32[4], count uint256)
func (_ERC721Metadata *ERC721MetadataSession) GetMetadata(_tokenId *big.Int, arg1 string) (struct {
	Buffer [4][32]byte
	Count  *big.Int
}, error) {
	return _ERC721Metadata.Contract.GetMetadata(&_ERC721Metadata.CallOpts, _tokenId, arg1)
}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(_tokenId uint256,  string) constant returns(buffer bytes32[4], count uint256)
func (_ERC721Metadata *ERC721MetadataCallerSession) GetMetadata(_tokenId *big.Int, arg1 string) (struct {
	Buffer [4][32]byte
	Count  *big.Int
}, error) {
	return _ERC721Metadata.Contract.GetMetadata(&_ERC721Metadata.CallOpts, _tokenId, arg1)
}

// GeneScienceInterfaceABI is the input ABI used to generate the binding from.
const GeneScienceInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"genes1\",\"type\":\"uint256\"},{\"name\":\"genes2\",\"type\":\"uint256\"},{\"name\":\"targetBlock\",\"type\":\"uint256\"}],\"name\":\"mixGenes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isGeneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// GeneScienceInterfaceBin is the compiled bytecode used for deploying new contracts.
const GeneScienceInterfaceBin = `0x`

// DeployGeneScienceInterface deploys a new Ethereum contract, binding an instance of GeneScienceInterface to it.
func DeployGeneScienceInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GeneScienceInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(GeneScienceInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GeneScienceInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GeneScienceInterface{GeneScienceInterfaceCaller: GeneScienceInterfaceCaller{contract: contract}, GeneScienceInterfaceTransactor: GeneScienceInterfaceTransactor{contract: contract}, GeneScienceInterfaceFilterer: GeneScienceInterfaceFilterer{contract: contract}}, nil
}

// GeneScienceInterface is an auto generated Go binding around an Ethereum contract.
type GeneScienceInterface struct {
	GeneScienceInterfaceCaller     // Read-only binding to the contract
	GeneScienceInterfaceTransactor // Write-only binding to the contract
	GeneScienceInterfaceFilterer   // Log filterer for contract events
}

// GeneScienceInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GeneScienceInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneScienceInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GeneScienceInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneScienceInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GeneScienceInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneScienceInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GeneScienceInterfaceSession struct {
	Contract     *GeneScienceInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GeneScienceInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GeneScienceInterfaceCallerSession struct {
	Contract *GeneScienceInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// GeneScienceInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GeneScienceInterfaceTransactorSession struct {
	Contract     *GeneScienceInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// GeneScienceInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GeneScienceInterfaceRaw struct {
	Contract *GeneScienceInterface // Generic contract binding to access the raw methods on
}

// GeneScienceInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GeneScienceInterfaceCallerRaw struct {
	Contract *GeneScienceInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// GeneScienceInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GeneScienceInterfaceTransactorRaw struct {
	Contract *GeneScienceInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGeneScienceInterface creates a new instance of GeneScienceInterface, bound to a specific deployed contract.
func NewGeneScienceInterface(address common.Address, backend bind.ContractBackend) (*GeneScienceInterface, error) {
	contract, err := bindGeneScienceInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GeneScienceInterface{GeneScienceInterfaceCaller: GeneScienceInterfaceCaller{contract: contract}, GeneScienceInterfaceTransactor: GeneScienceInterfaceTransactor{contract: contract}, GeneScienceInterfaceFilterer: GeneScienceInterfaceFilterer{contract: contract}}, nil
}

// NewGeneScienceInterfaceCaller creates a new read-only instance of GeneScienceInterface, bound to a specific deployed contract.
func NewGeneScienceInterfaceCaller(address common.Address, caller bind.ContractCaller) (*GeneScienceInterfaceCaller, error) {
	contract, err := bindGeneScienceInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GeneScienceInterfaceCaller{contract: contract}, nil
}

// NewGeneScienceInterfaceTransactor creates a new write-only instance of GeneScienceInterface, bound to a specific deployed contract.
func NewGeneScienceInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*GeneScienceInterfaceTransactor, error) {
	contract, err := bindGeneScienceInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GeneScienceInterfaceTransactor{contract: contract}, nil
}

// NewGeneScienceInterfaceFilterer creates a new log filterer instance of GeneScienceInterface, bound to a specific deployed contract.
func NewGeneScienceInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*GeneScienceInterfaceFilterer, error) {
	contract, err := bindGeneScienceInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GeneScienceInterfaceFilterer{contract: contract}, nil
}

// bindGeneScienceInterface binds a generic wrapper to an already deployed contract.
func bindGeneScienceInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GeneScienceInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneScienceInterface *GeneScienceInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GeneScienceInterface.Contract.GeneScienceInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneScienceInterface *GeneScienceInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.GeneScienceInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneScienceInterface *GeneScienceInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.GeneScienceInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneScienceInterface *GeneScienceInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GeneScienceInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneScienceInterface *GeneScienceInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneScienceInterface *GeneScienceInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.contract.Transact(opts, method, params...)
}

// IsGeneScience is a free data retrieval call binding the contract method 0x54c15b82.
//
// Solidity: function isGeneScience() constant returns(bool)
func (_GeneScienceInterface *GeneScienceInterfaceCaller) IsGeneScience(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GeneScienceInterface.contract.Call(opts, out, "isGeneScience")
	return *ret0, err
}

// IsGeneScience is a free data retrieval call binding the contract method 0x54c15b82.
//
// Solidity: function isGeneScience() constant returns(bool)
func (_GeneScienceInterface *GeneScienceInterfaceSession) IsGeneScience() (bool, error) {
	return _GeneScienceInterface.Contract.IsGeneScience(&_GeneScienceInterface.CallOpts)
}

// IsGeneScience is a free data retrieval call binding the contract method 0x54c15b82.
//
// Solidity: function isGeneScience() constant returns(bool)
func (_GeneScienceInterface *GeneScienceInterfaceCallerSession) IsGeneScience() (bool, error) {
	return _GeneScienceInterface.Contract.IsGeneScience(&_GeneScienceInterface.CallOpts)
}

// MixGenes is a paid mutator transaction binding the contract method 0x0d9f5aed.
//
// Solidity: function mixGenes(genes1 uint256, genes2 uint256, targetBlock uint256) returns(uint256)
func (_GeneScienceInterface *GeneScienceInterfaceTransactor) MixGenes(opts *bind.TransactOpts, genes1 *big.Int, genes2 *big.Int, targetBlock *big.Int) (*types.Transaction, error) {
	return _GeneScienceInterface.contract.Transact(opts, "mixGenes", genes1, genes2, targetBlock)
}

// MixGenes is a paid mutator transaction binding the contract method 0x0d9f5aed.
//
// Solidity: function mixGenes(genes1 uint256, genes2 uint256, targetBlock uint256) returns(uint256)
func (_GeneScienceInterface *GeneScienceInterfaceSession) MixGenes(genes1 *big.Int, genes2 *big.Int, targetBlock *big.Int) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.MixGenes(&_GeneScienceInterface.TransactOpts, genes1, genes2, targetBlock)
}

// MixGenes is a paid mutator transaction binding the contract method 0x0d9f5aed.
//
// Solidity: function mixGenes(genes1 uint256, genes2 uint256, targetBlock uint256) returns(uint256)
func (_GeneScienceInterface *GeneScienceInterfaceTransactorSession) MixGenes(genes1 *big.Int, genes2 *big.Int, targetBlock *big.Int) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.MixGenes(&_GeneScienceInterface.TransactOpts, genes1, genes2, targetBlock)
}

// KittyAccessControlABI is the input ABI used to generate the binding from.
const KittyAccessControlABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyAccessControlBin is the compiled bytecode used for deploying new contracts.
const KittyAccessControlBin = `0x60806040526002805460a060020a60ff021916905534801561002057600080fd5b5061043d806100306000396000f3006080604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630519ce79811461009d5780630a0f8168146100ce57806327d7874c146100e35780632ba73c15146101065780633f4ba83a146101275780634e0a33791461013c5780635c975abb1461015d5780638456cb5914610186578063b047fb501461019b575b600080fd5b3480156100a957600080fd5b506100b26101b0565b60408051600160a060020a039092168252519081900360200190f35b3480156100da57600080fd5b506100b26101bf565b3480156100ef57600080fd5b50610104600160a060020a03600435166101ce565b005b34801561011257600080fd5b50610104600160a060020a0360043516610229565b34801561013357600080fd5b50610104610284565b34801561014857600080fd5b50610104600160a060020a03600435166102e4565b34801561016957600080fd5b5061017261033f565b604080519115158252519081900360200190f35b34801561019257600080fd5b50610104610360565b3480156101a757600080fd5b506100b2610402565b600154600160a060020a031681565b600054600160a060020a031681565b600054600160a060020a031633146101e557600080fd5b600160a060020a03811615156101fa57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461024057600080fd5b600160a060020a038116151561025557600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461029b57600080fd5b60025474010000000000000000000000000000000000000000900460ff1615156102c457600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600054600160a060020a031633146102fb57600080fd5b600160a060020a038116151561031057600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60025474010000000000000000000000000000000000000000900460ff1681565b600254600160a060020a03163314806103835750600054600160a060020a031633145b806103985750600154600160a060020a031633145b15156103a357600080fd5b60025474010000000000000000000000000000000000000000900460ff16156103cb57600080fd5b6002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600254600160a060020a0316815600a165627a7a723058201be81598567bc2737e05ea6a6eb7009bcb39e218ae4952fdd4e16126540b6f250029`

// DeployKittyAccessControl deploys a new Ethereum contract, binding an instance of KittyAccessControl to it.
func DeployKittyAccessControl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyAccessControl, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyAccessControlABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyAccessControlBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyAccessControl{KittyAccessControlCaller: KittyAccessControlCaller{contract: contract}, KittyAccessControlTransactor: KittyAccessControlTransactor{contract: contract}, KittyAccessControlFilterer: KittyAccessControlFilterer{contract: contract}}, nil
}

// KittyAccessControl is an auto generated Go binding around an Ethereum contract.
type KittyAccessControl struct {
	KittyAccessControlCaller     // Read-only binding to the contract
	KittyAccessControlTransactor // Write-only binding to the contract
	KittyAccessControlFilterer   // Log filterer for contract events
}

// KittyAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyAccessControlSession struct {
	Contract     *KittyAccessControl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KittyAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyAccessControlCallerSession struct {
	Contract *KittyAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// KittyAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyAccessControlTransactorSession struct {
	Contract     *KittyAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// KittyAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyAccessControlRaw struct {
	Contract *KittyAccessControl // Generic contract binding to access the raw methods on
}

// KittyAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyAccessControlCallerRaw struct {
	Contract *KittyAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// KittyAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyAccessControlTransactorRaw struct {
	Contract *KittyAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyAccessControl creates a new instance of KittyAccessControl, bound to a specific deployed contract.
func NewKittyAccessControl(address common.Address, backend bind.ContractBackend) (*KittyAccessControl, error) {
	contract, err := bindKittyAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyAccessControl{KittyAccessControlCaller: KittyAccessControlCaller{contract: contract}, KittyAccessControlTransactor: KittyAccessControlTransactor{contract: contract}, KittyAccessControlFilterer: KittyAccessControlFilterer{contract: contract}}, nil
}

// NewKittyAccessControlCaller creates a new read-only instance of KittyAccessControl, bound to a specific deployed contract.
func NewKittyAccessControlCaller(address common.Address, caller bind.ContractCaller) (*KittyAccessControlCaller, error) {
	contract, err := bindKittyAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyAccessControlCaller{contract: contract}, nil
}

// NewKittyAccessControlTransactor creates a new write-only instance of KittyAccessControl, bound to a specific deployed contract.
func NewKittyAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyAccessControlTransactor, error) {
	contract, err := bindKittyAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyAccessControlTransactor{contract: contract}, nil
}

// NewKittyAccessControlFilterer creates a new log filterer instance of KittyAccessControl, bound to a specific deployed contract.
func NewKittyAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyAccessControlFilterer, error) {
	contract, err := bindKittyAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyAccessControlFilterer{contract: contract}, nil
}

// bindKittyAccessControl binds a generic wrapper to an already deployed contract.
func bindKittyAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyAccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyAccessControl *KittyAccessControlRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyAccessControl.Contract.KittyAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyAccessControl *KittyAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.KittyAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyAccessControl *KittyAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.KittyAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyAccessControl *KittyAccessControlCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyAccessControl *KittyAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyAccessControl *KittyAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.contract.Transact(opts, method, params...)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyAccessControl *KittyAccessControlCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAccessControl.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyAccessControl *KittyAccessControlSession) CeoAddress() (common.Address, error) {
	return _KittyAccessControl.Contract.CeoAddress(&_KittyAccessControl.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyAccessControl *KittyAccessControlCallerSession) CeoAddress() (common.Address, error) {
	return _KittyAccessControl.Contract.CeoAddress(&_KittyAccessControl.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyAccessControl *KittyAccessControlCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAccessControl.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyAccessControl *KittyAccessControlSession) CfoAddress() (common.Address, error) {
	return _KittyAccessControl.Contract.CfoAddress(&_KittyAccessControl.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyAccessControl *KittyAccessControlCallerSession) CfoAddress() (common.Address, error) {
	return _KittyAccessControl.Contract.CfoAddress(&_KittyAccessControl.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyAccessControl *KittyAccessControlCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAccessControl.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyAccessControl *KittyAccessControlSession) CooAddress() (common.Address, error) {
	return _KittyAccessControl.Contract.CooAddress(&_KittyAccessControl.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyAccessControl *KittyAccessControlCallerSession) CooAddress() (common.Address, error) {
	return _KittyAccessControl.Contract.CooAddress(&_KittyAccessControl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyAccessControl *KittyAccessControlCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyAccessControl.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyAccessControl *KittyAccessControlSession) Paused() (bool, error) {
	return _KittyAccessControl.Contract.Paused(&_KittyAccessControl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyAccessControl *KittyAccessControlCallerSession) Paused() (bool, error) {
	return _KittyAccessControl.Contract.Paused(&_KittyAccessControl.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyAccessControl *KittyAccessControlTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAccessControl.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyAccessControl *KittyAccessControlSession) Pause() (*types.Transaction, error) {
	return _KittyAccessControl.Contract.Pause(&_KittyAccessControl.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyAccessControl *KittyAccessControlTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyAccessControl.Contract.Pause(&_KittyAccessControl.TransactOpts)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyAccessControl *KittyAccessControlTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyAccessControl *KittyAccessControlSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.SetCEO(&_KittyAccessControl.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyAccessControl *KittyAccessControlTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.SetCEO(&_KittyAccessControl.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyAccessControl *KittyAccessControlTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyAccessControl *KittyAccessControlSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.SetCFO(&_KittyAccessControl.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyAccessControl *KittyAccessControlTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.SetCFO(&_KittyAccessControl.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyAccessControl *KittyAccessControlTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyAccessControl *KittyAccessControlSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.SetCOO(&_KittyAccessControl.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyAccessControl *KittyAccessControlTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.SetCOO(&_KittyAccessControl.TransactOpts, _newCOO)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyAccessControl *KittyAccessControlTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAccessControl.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyAccessControl *KittyAccessControlSession) Unpause() (*types.Transaction, error) {
	return _KittyAccessControl.Contract.Unpause(&_KittyAccessControl.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyAccessControl *KittyAccessControlTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyAccessControl.Contract.Unpause(&_KittyAccessControl.TransactOpts)
}

// KittyAccessControlContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyAccessControl contract.
type KittyAccessControlContractUpgradeIterator struct {
	Event *KittyAccessControlContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyAccessControlContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyAccessControlContractUpgrade)
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
		it.Event = new(KittyAccessControlContractUpgrade)
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
func (it *KittyAccessControlContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyAccessControlContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyAccessControlContractUpgrade represents a ContractUpgrade event raised by the KittyAccessControl contract.
type KittyAccessControlContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyAccessControl *KittyAccessControlFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyAccessControlContractUpgradeIterator, error) {

	logs, sub, err := _KittyAccessControl.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyAccessControlContractUpgradeIterator{contract: _KittyAccessControl.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyAccessControl *KittyAccessControlFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyAccessControlContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyAccessControl.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyAccessControlContractUpgrade)
				if err := _KittyAccessControl.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// KittyAuctionABI is the input ABI used to generate the binding from.
const KittyAuctionABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSiringAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pregnantKitties\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isPregnant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setGeneScienceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSaleAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"canBreedWith\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSiringAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAutoBirthFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"approveSiring\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSaleAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"giveBirth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAuctionBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoBirthFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isReadyToBreed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sireId\",\"type\":\"uint256\"},{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"bidOnSiringAuction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"geneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"breedWithAuto\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cooldownEndBlock\",\"type\":\"uint256\"}],\"name\":\"Pregnant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyAuctionBin is the compiled bytecode used for deploying new contracts.
const KittyAuctionBin = `0x6002805460a060020a60ff0219169055610240604052603c6080908152607860a05261012c60c05261025860e05261070861010052610e1061012052611c2061014052613840610160526170806101805261e1006101a052620151806101c0526202a3006101e052620546006102005262093a80610220526200008790600390600e620000ac565b50600f60055566071afd498d0000600e55348015620000a557600080fd5b5062000176565b6002830191839082156200013d5791602002820160005b838211156200010957835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302620000c3565b80156200013b5782816101000a81549063ffffffff021916905560040160208160030104928301926001030262000109565b505b506200014b9291506200014f565b5090565b6200017391905b808211156200014b57805463ffffffff1916815560010162000156565b90565b6129bd80620001866000396000f3006080604052600436106102215763ffffffff60e060020a60003504166301ffc9a781146102265780630519ce79146102715780630560ff44146102a257806306fdde031461033b578063095ea7b3146103505780630a0f81681461037657806314001f4c1461038b57806318160ddd146103ac578063183a7947146103d35780631940a936146103e857806321717ebf1461040057806323b872dd1461041557806324e7a38a1461043f57806327d7874c146104605780632ba73c15146104815780633d7d3f5a146104a25780633f4ba83a146104c357806346116e6f146104d857806346d22c70146104f0578063481af3d31461050b5780634ad8c938146105235780634b85fd55146105445780634dfff04f1461055c5780634e0a3379146105805780635663896e146105a15780635c975abb146105b95780636352211e146105ce5780636fbde40d146105e657806370a08231146106075780637a7d4937146106285780638456cb591461063d5780638462151c1461065257806388c2a0bf146106c357806391876e57146106db57806395d89b41146106f05780639d6fac6f14610705578063a45f4bfc14610736578063a9059cbb1461074e578063b047fb5014610772578063b0c35c0514610787578063bc4006f51461079c578063d3e6f49f146107b1578063e17b25af146107c9578063e6cbe351146107ea578063ed60ade6146107ff578063f2b47d521461080d578063f7d8c88314610822575b600080fd5b34801561023257600080fd5b5061025d7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1960043516610830565b604080519115158252519081900360200190f35b34801561027d57600080fd5b50610286610ac3565b60408051600160a060020a039092168252519081900360200190f35b3480156102ae57600080fd5b506102c6600480359060248035908101910135610ad2565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103005781810151838201526020016102e8565b50505050905090810190601f16801561032d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561034757600080fd5b506102c6610bd5565b34801561035c57600080fd5b50610374600160a060020a0360043516602435610c0c565b005b34801561038257600080fd5b50610286610c8e565b34801561039757600080fd5b50610374600160a060020a0360043516610c9d565b3480156103b857600080fd5b506103c1610d50565b60408051918252519081900360200190f35b3480156103df57600080fd5b506103c1610d5a565b3480156103f457600080fd5b5061025d600435610d60565b34801561040c57600080fd5b50610286610da5565b34801561042157600080fd5b50610374600160a060020a0360043581169060243516604435610db4565b34801561044b57600080fd5b50610374600160a060020a0360043516610e30565b34801561046c57600080fd5b50610374600160a060020a0360043516610ee3565b34801561048d57600080fd5b50610374600160a060020a0360043516610f31565b3480156104ae57600080fd5b50610374600435602435604435606435610f7f565b3480156104cf57600080fd5b50610374611074565b3480156104e457600080fd5b506102866004356110c3565b3480156104fc57600080fd5b5061025d6004356024356110de565b34801561051757600080fd5b5061028660043561115e565b34801561052f57600080fd5b50610374600435602435604435606435611179565b34801561055057600080fd5b50610374600435611251565b34801561056857600080fd5b50610374600160a060020a036004351660243561126d565b34801561058c57600080fd5b50610374600160a060020a03600435166112c7565b3480156105ad57600080fd5b50610374600435611315565b3480156105c557600080fd5b5061025d611371565b3480156105da57600080fd5b50610286600435611381565b3480156105f257600080fd5b50610374600160a060020a03600435166113a5565b34801561061357600080fd5b506103c1600160a060020a0360043516611458565b34801561063457600080fd5b506103c1611473565b34801561064957600080fd5b50610374611479565b34801561065e57600080fd5b50610673600160a060020a03600435166114f9565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156106af578181015183820152602001610697565b505050509050019250505060405180910390f35b3480156106cf57600080fd5b506103c16004356115cb565b3480156106e757600080fd5b506103746118a7565b3480156106fc57600080fd5b506102c66119c2565b34801561071157600080fd5b5061071d6004356119f9565b6040805163ffffffff9092168252519081900360200190f35b34801561074257600080fd5b50610286600435611a26565b34801561075a57600080fd5b50610374600160a060020a0360043516602435611a41565b34801561077e57600080fd5b50610286611add565b34801561079357600080fd5b506103c1611aec565b3480156107a857600080fd5b50610286611af2565b3480156107bd57600080fd5b5061025d600435611b01565b3480156107d557600080fd5b50610374600160a060020a0360043516611bd0565b3480156107f657600080fd5b50610286611c09565b610374600435602435611c18565b34801561081957600080fd5b50610286611db2565b610374600435602435611dc1565b604080517f737570706f727473496e74657266616365286279746573342900000000000000815290519081900360190190206000907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1983811691161480610abb5750604080517f746f6b656e4d657461646174612875696e743235362c737472696e67290000008152815190819003601d0181207f746f6b656e734f664f776e657228616464726573732900000000000000000000825282519182900360160182207f7472616e7366657246726f6d28616464726573732c616464726573732c75696e83527f7432353629000000000000000000000000000000000000000000000000000000602084015283519283900360250183207f7472616e7366657228616464726573732c75696e743235362900000000000000845284519384900360190184207f617070726f766528616464726573732c75696e74323536290000000000000000855285519485900360180185207f6f776e65724f662875696e743235362900000000000000000000000000000000865286519586900360100186207f62616c616e63654f662861646472657373290000000000000000000000000000875287519687900360120187207f746f74616c537570706c792829000000000000000000000000000000000000008852885197889003600d0188207f73796d626f6c2829000000000000000000000000000000000000000000000000895289519889900360080189207f6e616d65282900000000000000000000000000000000000000000000000000008a529951988990036006019098207bffffffffffffffffffffffffffffffffffffffffffffffffffffffff198c811691909a189098181818181818181891909116145b90505b919050565b600154600160a060020a031681565b6060610adc61292e565b600d54600090600160a060020a03161515610af657600080fd5b600d54604080517fcb4799f2000000000000000000000000000000000000000000000000000000008152600481018981526024820192835260448201889052600160a060020a039093169263cb4799f2928a928a928a929091606401848480828437820191505094505050505060a060405180830381600087803b158015610b7d57600080fd5b505af1158015610b91573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060a0811015610bb657600080fd5b5060808101519092509050610bcb8282611fbf565b9695505050505050565b60408051808201909152600d81527f43727970746f4b69747469657300000000000000000000000000000000000000602082015281565b60025460a060020a900460ff1615610c2357600080fd5b610c2d3382612013565b1515610c3857600080fd5b610c428183612033565b60408051338152600160a060020a038416602082015280820183905290517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360600190a15050565b600054600160a060020a031681565b60008054600160a060020a03163314610cb557600080fd5b81905080600160a060020a03166376190f8f6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610cf657600080fd5b505af1158015610d0a573d6000803e3d6000fd5b505050506040513d6020811015610d2057600080fd5b50511515610d2d57600080fd5b600c8054600160a060020a031916600160a060020a039290921691909117905550565b6006546000190190565b600f5481565b6000808211610d6e57600080fd5b6006805483908110610d7c57fe5b600091825260209091206002909102016001015460c060020a900463ffffffff16151592915050565b600c54600160a060020a031681565b60025460a060020a900460ff1615610dcb57600080fd5b600160a060020a0382161515610de057600080fd5b600160a060020a038216301415610df657600080fd5b610e003382612061565b1515610e0b57600080fd5b610e158382612013565b1515610e2057600080fd5b610e2b838383612081565b505050565b60008054600160a060020a03163314610e4857600080fd5b81905080600160a060020a03166354c15b826040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610e8957600080fd5b505af1158015610e9d573d6000803e3d6000fd5b505050506040513d6020811015610eb357600080fd5b50511515610ec057600080fd5b60108054600160a060020a031916600160a060020a039290921691909117905550565b600054600160a060020a03163314610efa57600080fd5b600160a060020a0381161515610f0f57600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314610f4857600080fd5b600160a060020a0381161515610f5d57600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b60025460a060020a900460ff1615610f9657600080fd5b610fa03385612013565b1515610fab57600080fd5b610fb484610d60565b15610fbe57600080fd5b600b54610fd5908590600160a060020a0316612033565b600b54604080517f27ebe40a000000000000000000000000000000000000000000000000000000008152600481018790526024810186905260448101859052606481018490523360848201529051600160a060020a03909216916327ebe40a9160a48082019260009290919082900301818387803b15801561105657600080fd5b505af115801561106a573d6000803e3d6000fd5b5050505050505050565b600054600160a060020a0316331461108b57600080fd5b60025460a060020a900460ff1615156110a357600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600a60205260009081526040902054600160a060020a031681565b600080808085116110ee57600080fd5b600084116110fb57600080fd5b600680548690811061110957fe5b9060005260206000209060020201915060068481548110151561112857fe5b9060005260206000209060020201905061114482868387612163565b8015611155575061115584866122e3565b95945050505050565b600960205260009081526040902054600160a060020a031681565b60025460a060020a900460ff161561119057600080fd5b61119a3385612013565b15156111a557600080fd5b6111ae84611b01565b15156111b957600080fd5b600c546111d0908590600160a060020a0316612033565b600c54604080517f27ebe40a000000000000000000000000000000000000000000000000000000008152600481018790526024810186905260448101859052606481018490523360848201529051600160a060020a03909216916327ebe40a9160a48082019260009290919082900301818387803b15801561105657600080fd5b600254600160a060020a0316331461126857600080fd5b600e55565b60025460a060020a900460ff161561128457600080fd5b61128e3382612013565b151561129957600080fd5b6000908152600a602052604090208054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a031633146112de57600080fd5b600160a060020a03811615156112f357600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b600254600160a060020a03163314806113385750600054600160a060020a031633145b8061134d5750600154600160a060020a031633145b151561135857600080fd5b60035463ffffffff16811061136c57600080fd5b600555565b60025460a060020a900460ff1681565b600081815260076020526040902054600160a060020a0316801515610abe57600080fd5b60008054600160a060020a031633146113bd57600080fd5b81905080600160a060020a03166385b861886040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156113fe57600080fd5b505af1158015611412573d6000803e3d6000fd5b505050506040513d602081101561142857600080fd5b5051151561143557600080fd5b600b8054600160a060020a031916600160a060020a039290921691909117905550565b600160a060020a031660009081526008602052604090205490565b60055481565b600254600160a060020a031633148061149c5750600054600160a060020a031633145b806114b15750600154600160a060020a031633145b15156114bc57600080fd5b60025460a060020a900460ff16156114d357600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a179055565b606060006060600080600061150d87611458565b945084151561152c5760408051600081526020810190915295506115c1565b84604051908082528060200260200182016040528015611556578160200160208202803883390190505b509350611561610d50565b925060009150600190505b8281116115bd57600081815260076020526040902054600160a060020a03888116911614156115b5578084838151811015156115a457fe5b602090810290910101526001909101905b60010161156c565b8395505b5050505050919050565b600080600080600080600080600260149054906101000a900460ff161515156115f357600080fd5b600680548a90811061160157fe5b60009182526020909120600290910201600181015490975067ffffffffffffffff16151561162e57600080fd5b604080516101008101825288548152600189015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e08201526116bf90612338565b15156116ca57600080fd5b60018701546006805460c060020a90920463ffffffff16975090879081106116ee57fe5b600091825260209091206001808a015460029093029091019081015490965061ffff60f060020a92839004811696509190041684101561173b57600185015460f060020a900461ffff1693505b6010548754865460018a0154604080517f0d9f5aed0000000000000000000000000000000000000000000000000000000081526004810194909452602484019290925260001967ffffffffffffffff6801000000000000000090920482160116604483015251600160a060020a0390921691630d9f5aed916064808201926020929091908290030181600087803b1580156117d557600080fd5b505af11580156117e9573d6000803e3d6000fd5b505050506040513d60208110156117ff57600080fd5b505160008a815260076020526040902054600189810154929550600160a060020a039091169350611848918b9160c060020a90910463ffffffff1690870161ffff168686612368565b6001880180547bffffffff00000000000000000000000000000000000000000000000019169055600f8054600019019055600e54604051919250339181156108fc0291906000818181858888f150939c9b505050505050505050505050565b600254600160a060020a03163314806118ca5750600054600160a060020a031633145b806118df5750600154600160a060020a031633145b15156118ea57600080fd5b600b60009054906101000a9004600160a060020a0316600160a060020a0316635fd8c7106040518163ffffffff1660e060020a028152600401600060405180830381600087803b15801561193d57600080fd5b505af1158015611951573d6000803e3d6000fd5b50505050600c60009054906101000a9004600160a060020a0316600160a060020a0316635fd8c7106040518163ffffffff1660e060020a028152600401600060405180830381600087803b1580156119a857600080fd5b505af11580156119bc573d6000803e3d6000fd5b50505050565b60408051808201909152600281527f434b000000000000000000000000000000000000000000000000000000000000602082015281565b600381600e8110611a0657fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b600760205260009081526040902054600160a060020a031681565b60025460a060020a900460ff1615611a5857600080fd5b600160a060020a0382161515611a6d57600080fd5b600160a060020a038216301415611a8357600080fd5b600b54600160a060020a0383811691161415611a9e57600080fd5b600c54600160a060020a0383811691161415611ab957600080fd5b611ac33382612013565b1515611ace57600080fd5b611ad9338383612081565b5050565b600254600160a060020a031681565b600e5481565b600d54600160a060020a031681565b600080808311611b1057600080fd5b6006805484908110611b1e57fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e0820152909150611bc990612645565b9392505050565b600054600160a060020a03163314611be757600080fd5b600d8054600160a060020a031916600160a060020a0392909216919091179055565b600b54600160a060020a031681565b60025460009060a060020a900460ff1615611c3257600080fd5b611c3c3383612013565b1515611c4757600080fd5b611c5082611b01565b1515611c5b57600080fd5b611c658284612674565b1515611c7057600080fd5b600c54604080517fc55d0f56000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a039092169163c55d0f56916024808201926020929091908290030181600087803b158015611cd757600080fd5b505af1158015611ceb573d6000803e3d6000fd5b505050506040513d6020811015611d0157600080fd5b5051600e549091508101341015611d1757600080fd5b600c54600e54604080517f454a2ab3000000000000000000000000000000000000000000000000000000008152600481018790529051600160a060020a039093169263454a2ab39234039160248082019260009290919082900301818588803b158015611d8357600080fd5b505af1158015611d97573d6000803e3d6000fd5b5050505050610e2b8263ffffffff168463ffffffff166126c3565b601054600160a060020a031681565b600254600090819060a060020a900460ff1615611ddd57600080fd5b600e54341015611dec57600080fd5b611df63385612013565b1515611e0157600080fd5b611e0b83856122e3565b1515611e1657600080fd5b6006805485908110611e2457fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e0820152909250611ecf90612645565b1515611eda57600080fd5b6006805484908110611ee857fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e0820152909150611f9390612645565b1515611f9e57600080fd5b611faa82858386612163565b1515611fb557600080fd5b6119bc84846126c3565b606080600080846040519080825280601f01601f191660200182016040528015611ff3578160200160208202803883390190505b5092505060208201905084612009828287612801565b5090949350505050565b600090815260076020526040902054600160a060020a0391821691161490565b6000918252600960205260409091208054600160a060020a031916600160a060020a03909216919091179055565b600090815260096020526040902054600160a060020a0391821691161490565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a031916909117905583161561211457600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b60408051600160a060020a0380861682528416602082015280820183905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a1505050565b600081841415612175575060006122db565b6001850154608060020a900463ffffffff168214806121a45750600185015460a060020a900463ffffffff1682145b156121b1575060006122db565b6001830154608060020a900463ffffffff168414806121e05750600183015460a060020a900463ffffffff1684145b156121ed575060006122db565b6001830154608060020a900463ffffffff16158061221a57506001850154608060020a900463ffffffff16155b15612227575060016122db565b60018581015490840154608060020a9182900463ffffffff90811692909104161480612272575060018086015490840154608060020a900463ffffffff90811660a060020a90920416145b1561227f575060006122db565b6001808601549084015460a060020a900463ffffffff908116608060020a9092041614806122ca57506001858101549084015460a060020a9182900463ffffffff9081169290910416145b156122d7575060006122db565b5060015b949350505050565b6000818152600760205260408082205484835290822054600160a060020a0391821691168082148061115557506000858152600a6020526040902054600160a060020a03908116908316149250505092915050565b60008160a0015163ffffffff16600014158015610abb5750506040015167ffffffffffffffff4381169116111590565b60008061237361294d565b600063ffffffff8916891461238757600080fd5b63ffffffff8816881461239957600080fd5b61ffff871687146123a957600080fd5b600287049250600d8361ffff1611156123c157600d92505b505060408051610100810182528581524267ffffffffffffffff90811660208301908152600093830184815263ffffffff8c8116606086019081528c82166080870190815260a0870188815261ffff8a811660c08a019081528f821660e08b01908152600680546001810182559c528a517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f60028e029081019190915598517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d40909901805498519651955194519251915167ffffffffffffffff19909916998b16999099176fffffffffffffffff000000000000000019166801000000000000000096909a16959095029890981773ffffffff000000000000000000000000000000001916608060020a938616939093029290921777ffffffff0000000000000000000000000000000000000000191660a060020a91851691909102177bffffffff000000000000000000000000000000000000000000000000191660c060020a96841696909602959095177fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e060020a91861691909102177dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f060020a9290941691909102929092179055909190811681146125bf57600080fd5b606080830151608080850151855160408051600160a060020a038c1681526020810188905263ffffffff95861681830152929094169482019490945290810192909252517f0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad59181900360a00190a161263960008683612081565b98975050505050505050565b60008160a0015163ffffffff166000148015610abb5750506040015167ffffffffffffffff4381169116111590565b600080600060068581548110151561268857fe5b906000526020600020906002020191506006848154811015156126a757fe5b9060005260206000209060020201905061115582868387612163565b6000806006838154811015156126d557fe5b906000526020600020906002020191506006848154811015156126f457fe5b600091825260209091206002909102016001810180547bffffffff000000000000000000000000000000000000000000000000191660c060020a63ffffffff871602179055905061274482612845565b61274d81612845565b6000848152600a602090815260408083208054600160a060020a031990811690915586845281842080549091169055600f80546001908101909155878452600783529281902054928401548151600160a060020a0390941684529183018790528281018690526801000000000000000090910467ffffffffffffffff166060830152517f241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80916080908290030190a150505050565b60005b60208210612826578251845260209384019390920191601f1990910190612804565b50905182516020929092036101000a6000190180199091169116179052565b600554600182015443919060039060e060020a900461ffff16600e811061286857fe5b600891828204019190066004029054906101000a900463ffffffff1663ffffffff1681151561289357fe5b6001840180546fffffffffffffffff0000000000000000191668010000000000000000939092049390930167ffffffffffffffff16919091021790819055600d60e060020a90910461ffff16101561292b576001818101805461ffff60e060020a8083048216909401169092027fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790555b50565b6080604051908101604052806004906020820280388339509192915050565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810191909152905600a165627a7a723058205dc2e850f832820ea857401c1ffec616a93a07b6555f5f951bc7c91825e4c0e30029`

// DeployKittyAuction deploys a new Ethereum contract, binding an instance of KittyAuction to it.
func DeployKittyAuction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyAuctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyAuction{KittyAuctionCaller: KittyAuctionCaller{contract: contract}, KittyAuctionTransactor: KittyAuctionTransactor{contract: contract}, KittyAuctionFilterer: KittyAuctionFilterer{contract: contract}}, nil
}

// KittyAuction is an auto generated Go binding around an Ethereum contract.
type KittyAuction struct {
	KittyAuctionCaller     // Read-only binding to the contract
	KittyAuctionTransactor // Write-only binding to the contract
	KittyAuctionFilterer   // Log filterer for contract events
}

// KittyAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyAuctionSession struct {
	Contract     *KittyAuction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittyAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyAuctionCallerSession struct {
	Contract *KittyAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KittyAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyAuctionTransactorSession struct {
	Contract     *KittyAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KittyAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyAuctionRaw struct {
	Contract *KittyAuction // Generic contract binding to access the raw methods on
}

// KittyAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyAuctionCallerRaw struct {
	Contract *KittyAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// KittyAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyAuctionTransactorRaw struct {
	Contract *KittyAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyAuction creates a new instance of KittyAuction, bound to a specific deployed contract.
func NewKittyAuction(address common.Address, backend bind.ContractBackend) (*KittyAuction, error) {
	contract, err := bindKittyAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyAuction{KittyAuctionCaller: KittyAuctionCaller{contract: contract}, KittyAuctionTransactor: KittyAuctionTransactor{contract: contract}, KittyAuctionFilterer: KittyAuctionFilterer{contract: contract}}, nil
}

// NewKittyAuctionCaller creates a new read-only instance of KittyAuction, bound to a specific deployed contract.
func NewKittyAuctionCaller(address common.Address, caller bind.ContractCaller) (*KittyAuctionCaller, error) {
	contract, err := bindKittyAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyAuctionCaller{contract: contract}, nil
}

// NewKittyAuctionTransactor creates a new write-only instance of KittyAuction, bound to a specific deployed contract.
func NewKittyAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyAuctionTransactor, error) {
	contract, err := bindKittyAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyAuctionTransactor{contract: contract}, nil
}

// NewKittyAuctionFilterer creates a new log filterer instance of KittyAuction, bound to a specific deployed contract.
func NewKittyAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyAuctionFilterer, error) {
	contract, err := bindKittyAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyAuctionFilterer{contract: contract}, nil
}

// bindKittyAuction binds a generic wrapper to an already deployed contract.
func bindKittyAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyAuction *KittyAuctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyAuction.Contract.KittyAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyAuction *KittyAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAuction.Contract.KittyAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyAuction *KittyAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyAuction.Contract.KittyAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyAuction *KittyAuctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyAuction *KittyAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyAuction *KittyAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyAuction.Contract.contract.Transact(opts, method, params...)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_KittyAuction *KittyAuctionCaller) AutoBirthFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "autoBirthFee")
	return *ret0, err
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_KittyAuction *KittyAuctionSession) AutoBirthFee() (*big.Int, error) {
	return _KittyAuction.Contract.AutoBirthFee(&_KittyAuction.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_KittyAuction *KittyAuctionCallerSession) AutoBirthFee() (*big.Int, error) {
	return _KittyAuction.Contract.AutoBirthFee(&_KittyAuction.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyAuction *KittyAuctionCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyAuction *KittyAuctionSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyAuction.Contract.BalanceOf(&_KittyAuction.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyAuction *KittyAuctionCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyAuction.Contract.BalanceOf(&_KittyAuction.CallOpts, _owner)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_KittyAuction *KittyAuctionCaller) CanBreedWith(opts *bind.CallOpts, _matronId *big.Int, _sireId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "canBreedWith", _matronId, _sireId)
	return *ret0, err
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_KittyAuction *KittyAuctionSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyAuction.Contract.CanBreedWith(&_KittyAuction.CallOpts, _matronId, _sireId)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_KittyAuction *KittyAuctionCallerSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyAuction.Contract.CanBreedWith(&_KittyAuction.CallOpts, _matronId, _sireId)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyAuction *KittyAuctionCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyAuction *KittyAuctionSession) CeoAddress() (common.Address, error) {
	return _KittyAuction.Contract.CeoAddress(&_KittyAuction.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyAuction *KittyAuctionCallerSession) CeoAddress() (common.Address, error) {
	return _KittyAuction.Contract.CeoAddress(&_KittyAuction.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyAuction *KittyAuctionCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyAuction *KittyAuctionSession) CfoAddress() (common.Address, error) {
	return _KittyAuction.Contract.CfoAddress(&_KittyAuction.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyAuction *KittyAuctionCallerSession) CfoAddress() (common.Address, error) {
	return _KittyAuction.Contract.CfoAddress(&_KittyAuction.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyAuction *KittyAuctionCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyAuction *KittyAuctionSession) CooAddress() (common.Address, error) {
	return _KittyAuction.Contract.CooAddress(&_KittyAuction.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyAuction *KittyAuctionCallerSession) CooAddress() (common.Address, error) {
	return _KittyAuction.Contract.CooAddress(&_KittyAuction.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyAuction *KittyAuctionCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "cooldowns", arg0)
	return *ret0, err
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyAuction *KittyAuctionSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyAuction.Contract.Cooldowns(&_KittyAuction.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyAuction *KittyAuctionCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyAuction.Contract.Cooldowns(&_KittyAuction.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyAuction *KittyAuctionCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "erc721Metadata")
	return *ret0, err
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyAuction *KittyAuctionSession) Erc721Metadata() (common.Address, error) {
	return _KittyAuction.Contract.Erc721Metadata(&_KittyAuction.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyAuction *KittyAuctionCallerSession) Erc721Metadata() (common.Address, error) {
	return _KittyAuction.Contract.Erc721Metadata(&_KittyAuction.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_KittyAuction *KittyAuctionCaller) GeneScience(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "geneScience")
	return *ret0, err
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_KittyAuction *KittyAuctionSession) GeneScience() (common.Address, error) {
	return _KittyAuction.Contract.GeneScience(&_KittyAuction.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_KittyAuction *KittyAuctionCallerSession) GeneScience() (common.Address, error) {
	return _KittyAuction.Contract.GeneScience(&_KittyAuction.CallOpts)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_KittyAuction *KittyAuctionCaller) IsPregnant(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "isPregnant", _kittyId)
	return *ret0, err
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_KittyAuction *KittyAuctionSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyAuction.Contract.IsPregnant(&_KittyAuction.CallOpts, _kittyId)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_KittyAuction *KittyAuctionCallerSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyAuction.Contract.IsPregnant(&_KittyAuction.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_KittyAuction *KittyAuctionCaller) IsReadyToBreed(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "isReadyToBreed", _kittyId)
	return *ret0, err
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_KittyAuction *KittyAuctionSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyAuction.Contract.IsReadyToBreed(&_KittyAuction.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_KittyAuction *KittyAuctionCallerSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyAuction.Contract.IsReadyToBreed(&_KittyAuction.CallOpts, _kittyId)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyAuction *KittyAuctionCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "kittyIndexToApproved", arg0)
	return *ret0, err
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyAuction *KittyAuctionSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.KittyIndexToApproved(&_KittyAuction.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyAuction *KittyAuctionCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.KittyIndexToApproved(&_KittyAuction.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyAuction *KittyAuctionCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "kittyIndexToOwner", arg0)
	return *ret0, err
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyAuction *KittyAuctionSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.KittyIndexToOwner(&_KittyAuction.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyAuction *KittyAuctionCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.KittyIndexToOwner(&_KittyAuction.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyAuction *KittyAuctionCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyAuction *KittyAuctionSession) Name() (string, error) {
	return _KittyAuction.Contract.Name(&_KittyAuction.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyAuction *KittyAuctionCallerSession) Name() (string, error) {
	return _KittyAuction.Contract.Name(&_KittyAuction.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyAuction *KittyAuctionCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyAuction *KittyAuctionSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.OwnerOf(&_KittyAuction.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyAuction *KittyAuctionCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.OwnerOf(&_KittyAuction.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyAuction *KittyAuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyAuction *KittyAuctionSession) Paused() (bool, error) {
	return _KittyAuction.Contract.Paused(&_KittyAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyAuction *KittyAuctionCallerSession) Paused() (bool, error) {
	return _KittyAuction.Contract.Paused(&_KittyAuction.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_KittyAuction *KittyAuctionCaller) PregnantKitties(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "pregnantKitties")
	return *ret0, err
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_KittyAuction *KittyAuctionSession) PregnantKitties() (*big.Int, error) {
	return _KittyAuction.Contract.PregnantKitties(&_KittyAuction.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_KittyAuction *KittyAuctionCallerSession) PregnantKitties() (*big.Int, error) {
	return _KittyAuction.Contract.PregnantKitties(&_KittyAuction.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyAuction *KittyAuctionCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "saleAuction")
	return *ret0, err
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyAuction *KittyAuctionSession) SaleAuction() (common.Address, error) {
	return _KittyAuction.Contract.SaleAuction(&_KittyAuction.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyAuction *KittyAuctionCallerSession) SaleAuction() (common.Address, error) {
	return _KittyAuction.Contract.SaleAuction(&_KittyAuction.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyAuction *KittyAuctionCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "secondsPerBlock")
	return *ret0, err
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyAuction *KittyAuctionSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyAuction.Contract.SecondsPerBlock(&_KittyAuction.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyAuction *KittyAuctionCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyAuction.Contract.SecondsPerBlock(&_KittyAuction.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyAuction *KittyAuctionCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "sireAllowedToAddress", arg0)
	return *ret0, err
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyAuction *KittyAuctionSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.SireAllowedToAddress(&_KittyAuction.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyAuction *KittyAuctionCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.SireAllowedToAddress(&_KittyAuction.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyAuction *KittyAuctionCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "siringAuction")
	return *ret0, err
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyAuction *KittyAuctionSession) SiringAuction() (common.Address, error) {
	return _KittyAuction.Contract.SiringAuction(&_KittyAuction.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyAuction *KittyAuctionCallerSession) SiringAuction() (common.Address, error) {
	return _KittyAuction.Contract.SiringAuction(&_KittyAuction.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyAuction *KittyAuctionCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "supportsInterface", _interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyAuction *KittyAuctionSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyAuction.Contract.SupportsInterface(&_KittyAuction.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyAuction *KittyAuctionCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyAuction.Contract.SupportsInterface(&_KittyAuction.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyAuction *KittyAuctionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyAuction *KittyAuctionSession) Symbol() (string, error) {
	return _KittyAuction.Contract.Symbol(&_KittyAuction.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyAuction *KittyAuctionCallerSession) Symbol() (string, error) {
	return _KittyAuction.Contract.Symbol(&_KittyAuction.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyAuction *KittyAuctionCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "tokenMetadata", _tokenId, _preferredTransport)
	return *ret0, err
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyAuction *KittyAuctionSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyAuction.Contract.TokenMetadata(&_KittyAuction.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyAuction *KittyAuctionCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyAuction.Contract.TokenMetadata(&_KittyAuction.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyAuction *KittyAuctionCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyAuction *KittyAuctionSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyAuction.Contract.TokensOfOwner(&_KittyAuction.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyAuction *KittyAuctionCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyAuction.Contract.TokensOfOwner(&_KittyAuction.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyAuction *KittyAuctionCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyAuction.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyAuction *KittyAuctionSession) TotalSupply() (*big.Int, error) {
	return _KittyAuction.Contract.TotalSupply(&_KittyAuction.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyAuction *KittyAuctionCallerSession) TotalSupply() (*big.Int, error) {
	return _KittyAuction.Contract.TotalSupply(&_KittyAuction.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyAuction *KittyAuctionTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyAuction *KittyAuctionSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.Approve(&_KittyAuction.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyAuction *KittyAuctionTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.Approve(&_KittyAuction.TransactOpts, _to, _tokenId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_KittyAuction *KittyAuctionTransactor) ApproveSiring(opts *bind.TransactOpts, _addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "approveSiring", _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_KittyAuction *KittyAuctionSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.ApproveSiring(&_KittyAuction.TransactOpts, _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_KittyAuction *KittyAuctionTransactorSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.ApproveSiring(&_KittyAuction.TransactOpts, _addr, _sireId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(_sireId uint256, _matronId uint256) returns()
func (_KittyAuction *KittyAuctionTransactor) BidOnSiringAuction(opts *bind.TransactOpts, _sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "bidOnSiringAuction", _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(_sireId uint256, _matronId uint256) returns()
func (_KittyAuction *KittyAuctionSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.BidOnSiringAuction(&_KittyAuction.TransactOpts, _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(_sireId uint256, _matronId uint256) returns()
func (_KittyAuction *KittyAuctionTransactorSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.BidOnSiringAuction(&_KittyAuction.TransactOpts, _sireId, _matronId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_KittyAuction *KittyAuctionTransactor) BreedWithAuto(opts *bind.TransactOpts, _matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "breedWithAuto", _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_KittyAuction *KittyAuctionSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.BreedWithAuto(&_KittyAuction.TransactOpts, _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_KittyAuction *KittyAuctionTransactorSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.BreedWithAuto(&_KittyAuction.TransactOpts, _matronId, _sireId)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyAuction *KittyAuctionTransactor) CreateSaleAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "createSaleAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyAuction *KittyAuctionSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.CreateSaleAuction(&_KittyAuction.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyAuction *KittyAuctionTransactorSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.CreateSaleAuction(&_KittyAuction.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyAuction *KittyAuctionTransactor) CreateSiringAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "createSiringAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyAuction *KittyAuctionSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.CreateSiringAuction(&_KittyAuction.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyAuction *KittyAuctionTransactorSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.CreateSiringAuction(&_KittyAuction.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_KittyAuction *KittyAuctionTransactor) GiveBirth(opts *bind.TransactOpts, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "giveBirth", _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_KittyAuction *KittyAuctionSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.GiveBirth(&_KittyAuction.TransactOpts, _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_KittyAuction *KittyAuctionTransactorSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.GiveBirth(&_KittyAuction.TransactOpts, _matronId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyAuction *KittyAuctionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyAuction *KittyAuctionSession) Pause() (*types.Transaction, error) {
	return _KittyAuction.Contract.Pause(&_KittyAuction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyAuction *KittyAuctionTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyAuction.Contract.Pause(&_KittyAuction.TransactOpts)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_KittyAuction *KittyAuctionTransactor) SetAutoBirthFee(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setAutoBirthFee", val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_KittyAuction *KittyAuctionSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetAutoBirthFee(&_KittyAuction.TransactOpts, val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetAutoBirthFee(&_KittyAuction.TransactOpts, val)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyAuction *KittyAuctionTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyAuction *KittyAuctionSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetCEO(&_KittyAuction.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetCEO(&_KittyAuction.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyAuction *KittyAuctionTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyAuction *KittyAuctionSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetCFO(&_KittyAuction.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetCFO(&_KittyAuction.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyAuction *KittyAuctionTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyAuction *KittyAuctionSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetCOO(&_KittyAuction.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetCOO(&_KittyAuction.TransactOpts, _newCOO)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_KittyAuction *KittyAuctionTransactor) SetGeneScienceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setGeneScienceAddress", _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_KittyAuction *KittyAuctionSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetGeneScienceAddress(&_KittyAuction.TransactOpts, _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetGeneScienceAddress(&_KittyAuction.TransactOpts, _address)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyAuction *KittyAuctionTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyAuction *KittyAuctionSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetMetadataAddress(&_KittyAuction.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetMetadataAddress(&_KittyAuction.TransactOpts, _contractAddress)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(_address address) returns()
func (_KittyAuction *KittyAuctionTransactor) SetSaleAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setSaleAuctionAddress", _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(_address address) returns()
func (_KittyAuction *KittyAuctionSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetSaleAuctionAddress(&_KittyAuction.TransactOpts, _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(_address address) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetSaleAuctionAddress(&_KittyAuction.TransactOpts, _address)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyAuction *KittyAuctionTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyAuction *KittyAuctionSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetSecondsPerBlock(&_KittyAuction.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetSecondsPerBlock(&_KittyAuction.TransactOpts, secs)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(_address address) returns()
func (_KittyAuction *KittyAuctionTransactor) SetSiringAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setSiringAuctionAddress", _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(_address address) returns()
func (_KittyAuction *KittyAuctionSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetSiringAuctionAddress(&_KittyAuction.TransactOpts, _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(_address address) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetSiringAuctionAddress(&_KittyAuction.TransactOpts, _address)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyAuction *KittyAuctionTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyAuction *KittyAuctionSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.Transfer(&_KittyAuction.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyAuction *KittyAuctionTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.Transfer(&_KittyAuction.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyAuction *KittyAuctionTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyAuction *KittyAuctionSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.TransferFrom(&_KittyAuction.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyAuction *KittyAuctionTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.TransferFrom(&_KittyAuction.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyAuction *KittyAuctionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyAuction *KittyAuctionSession) Unpause() (*types.Transaction, error) {
	return _KittyAuction.Contract.Unpause(&_KittyAuction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyAuction *KittyAuctionTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyAuction.Contract.Unpause(&_KittyAuction.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyAuction *KittyAuctionTransactor) WithdrawAuctionBalances(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "withdrawAuctionBalances")
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyAuction *KittyAuctionSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _KittyAuction.Contract.WithdrawAuctionBalances(&_KittyAuction.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyAuction *KittyAuctionTransactorSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _KittyAuction.Contract.WithdrawAuctionBalances(&_KittyAuction.TransactOpts)
}

// KittyAuctionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KittyAuction contract.
type KittyAuctionApprovalIterator struct {
	Event *KittyAuctionApproval // Event containing the contract specifics and raw log

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
func (it *KittyAuctionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyAuctionApproval)
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
		it.Event = new(KittyAuctionApproval)
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
func (it *KittyAuctionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyAuctionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyAuctionApproval represents a Approval event raised by the KittyAuction contract.
type KittyAuctionApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_KittyAuction *KittyAuctionFilterer) FilterApproval(opts *bind.FilterOpts) (*KittyAuctionApprovalIterator, error) {

	logs, sub, err := _KittyAuction.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &KittyAuctionApprovalIterator{contract: _KittyAuction.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_KittyAuction *KittyAuctionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KittyAuctionApproval) (event.Subscription, error) {

	logs, sub, err := _KittyAuction.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyAuctionApproval)
				if err := _KittyAuction.contract.UnpackLog(event, "Approval", log); err != nil {
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

// KittyAuctionBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the KittyAuction contract.
type KittyAuctionBirthIterator struct {
	Event *KittyAuctionBirth // Event containing the contract specifics and raw log

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
func (it *KittyAuctionBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyAuctionBirth)
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
		it.Event = new(KittyAuctionBirth)
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
func (it *KittyAuctionBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyAuctionBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyAuctionBirth represents a Birth event raised by the KittyAuction contract.
type KittyAuctionBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_KittyAuction *KittyAuctionFilterer) FilterBirth(opts *bind.FilterOpts) (*KittyAuctionBirthIterator, error) {

	logs, sub, err := _KittyAuction.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &KittyAuctionBirthIterator{contract: _KittyAuction.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_KittyAuction *KittyAuctionFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *KittyAuctionBirth) (event.Subscription, error) {

	logs, sub, err := _KittyAuction.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyAuctionBirth)
				if err := _KittyAuction.contract.UnpackLog(event, "Birth", log); err != nil {
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

// KittyAuctionContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyAuction contract.
type KittyAuctionContractUpgradeIterator struct {
	Event *KittyAuctionContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyAuctionContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyAuctionContractUpgrade)
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
		it.Event = new(KittyAuctionContractUpgrade)
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
func (it *KittyAuctionContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyAuctionContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyAuctionContractUpgrade represents a ContractUpgrade event raised by the KittyAuction contract.
type KittyAuctionContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyAuction *KittyAuctionFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyAuctionContractUpgradeIterator, error) {

	logs, sub, err := _KittyAuction.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyAuctionContractUpgradeIterator{contract: _KittyAuction.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyAuction *KittyAuctionFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyAuctionContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyAuction.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyAuctionContractUpgrade)
				if err := _KittyAuction.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// KittyAuctionPregnantIterator is returned from FilterPregnant and is used to iterate over the raw logs and unpacked data for Pregnant events raised by the KittyAuction contract.
type KittyAuctionPregnantIterator struct {
	Event *KittyAuctionPregnant // Event containing the contract specifics and raw log

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
func (it *KittyAuctionPregnantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyAuctionPregnant)
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
		it.Event = new(KittyAuctionPregnant)
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
func (it *KittyAuctionPregnantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyAuctionPregnantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyAuctionPregnant represents a Pregnant event raised by the KittyAuction contract.
type KittyAuctionPregnant struct {
	Owner            common.Address
	MatronId         *big.Int
	SireId           *big.Int
	CooldownEndBlock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPregnant is a free log retrieval operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: e Pregnant(owner address, matronId uint256, sireId uint256, cooldownEndBlock uint256)
func (_KittyAuction *KittyAuctionFilterer) FilterPregnant(opts *bind.FilterOpts) (*KittyAuctionPregnantIterator, error) {

	logs, sub, err := _KittyAuction.contract.FilterLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return &KittyAuctionPregnantIterator{contract: _KittyAuction.contract, event: "Pregnant", logs: logs, sub: sub}, nil
}

// WatchPregnant is a free log subscription operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: e Pregnant(owner address, matronId uint256, sireId uint256, cooldownEndBlock uint256)
func (_KittyAuction *KittyAuctionFilterer) WatchPregnant(opts *bind.WatchOpts, sink chan<- *KittyAuctionPregnant) (event.Subscription, error) {

	logs, sub, err := _KittyAuction.contract.WatchLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyAuctionPregnant)
				if err := _KittyAuction.contract.UnpackLog(event, "Pregnant", log); err != nil {
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

// KittyAuctionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KittyAuction contract.
type KittyAuctionTransferIterator struct {
	Event *KittyAuctionTransfer // Event containing the contract specifics and raw log

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
func (it *KittyAuctionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyAuctionTransfer)
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
		it.Event = new(KittyAuctionTransfer)
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
func (it *KittyAuctionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyAuctionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyAuctionTransfer represents a Transfer event raised by the KittyAuction contract.
type KittyAuctionTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_KittyAuction *KittyAuctionFilterer) FilterTransfer(opts *bind.FilterOpts) (*KittyAuctionTransferIterator, error) {

	logs, sub, err := _KittyAuction.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &KittyAuctionTransferIterator{contract: _KittyAuction.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_KittyAuction *KittyAuctionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KittyAuctionTransfer) (event.Subscription, error) {

	logs, sub, err := _KittyAuction.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyAuctionTransfer)
				if err := _KittyAuction.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// KittyBaseABI is the input ABI used to generate the binding from.
const KittyBaseABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyBaseBin is the compiled bytecode used for deploying new contracts.
const KittyBaseBin = `0x6002805460a060020a60ff0219169055610240604052603c6080908152607860a05261012c60c05261025860e05261070861010052610e1061012052611c2061014052613840610160526170806101805261e1006101a052620151806101c0526202a3006101e052620546006102005262093a806102205261008590600390600e61009d565b50600f60055534801561009757600080fd5b5061015d565b6002830191839082156101295791602002820160005b838211156100f757835183826101000a81548163ffffffff021916908363ffffffff16021790555092602001926004016020816003010492830192600103026100b3565b80156101275782816101000a81549063ffffffff02191690556004016020816003010492830192600103026100f7565b505b50610135929150610139565b5090565b61015a91905b8082111561013557805463ffffffff1916815560010161013f565b90565b6106758061016c6000396000f3006080604052600436106100f05763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630519ce7981146100f55780630a0f81681461012657806321717ebf1461013b57806327d7874c146101505780632ba73c15146101735780633f4ba83a1461019457806346116e6f146101a9578063481af3d3146101c15780634e0a3379146101d95780635663896e146101fa5780635c975abb146102125780637a7d49371461023b5780638456cb59146102625780639d6fac6f14610277578063a45f4bfc146102a8578063b047fb50146102c0578063e6cbe351146102d5575b600080fd5b34801561010157600080fd5b5061010a6102ea565b60408051600160a060020a039092168252519081900360200190f35b34801561013257600080fd5b5061010a6102f9565b34801561014757600080fd5b5061010a610308565b34801561015c57600080fd5b50610171600160a060020a0360043516610317565b005b34801561017f57600080fd5b50610171600160a060020a0360043516610372565b3480156101a057600080fd5b506101716103cd565b3480156101b557600080fd5b5061010a60043561042d565b3480156101cd57600080fd5b5061010a600435610448565b3480156101e557600080fd5b50610171600160a060020a0360043516610463565b34801561020657600080fd5b506101716004356104be565b34801561021e57600080fd5b5061022761051a565b604080519115158252519081900360200190f35b34801561024757600080fd5b5061025061053b565b60408051918252519081900360200190f35b34801561026e57600080fd5b50610171610541565b34801561028357600080fd5b5061028f6004356105e3565b6040805163ffffffff9092168252519081900360200190f35b3480156102b457600080fd5b5061010a600435610610565b3480156102cc57600080fd5b5061010a61062b565b3480156102e157600080fd5b5061010a61063a565b600154600160a060020a031681565b600054600160a060020a031681565b600c54600160a060020a031681565b600054600160a060020a0316331461032e57600080fd5b600160a060020a038116151561034357600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461038957600080fd5b600160a060020a038116151561039e57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031633146103e457600080fd5b60025474010000000000000000000000000000000000000000900460ff16151561040d57600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600a60205260009081526040902054600160a060020a031681565b600960205260009081526040902054600160a060020a031681565b600054600160a060020a0316331461047a57600080fd5b600160a060020a038116151561048f57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a03163314806104e15750600054600160a060020a031633145b806104f65750600154600160a060020a031633145b151561050157600080fd5b60035463ffffffff16811061051557600080fd5b600555565b60025474010000000000000000000000000000000000000000900460ff1681565b60055481565b600254600160a060020a03163314806105645750600054600160a060020a031633145b806105795750600154600160a060020a031633145b151561058457600080fd5b60025474010000000000000000000000000000000000000000900460ff16156105ac57600080fd5b6002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600381600e81106105f057fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b600760205260009081526040902054600160a060020a031681565b600254600160a060020a031681565b600b54600160a060020a0316815600a165627a7a72305820e1563fceedcafbd10073d64d35e5b0c8954c0e8c0b7f2c954c259e6d90055c920029`

// DeployKittyBase deploys a new Ethereum contract, binding an instance of KittyBase to it.
func DeployKittyBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyBase, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyBase{KittyBaseCaller: KittyBaseCaller{contract: contract}, KittyBaseTransactor: KittyBaseTransactor{contract: contract}, KittyBaseFilterer: KittyBaseFilterer{contract: contract}}, nil
}

// KittyBase is an auto generated Go binding around an Ethereum contract.
type KittyBase struct {
	KittyBaseCaller     // Read-only binding to the contract
	KittyBaseTransactor // Write-only binding to the contract
	KittyBaseFilterer   // Log filterer for contract events
}

// KittyBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyBaseSession struct {
	Contract     *KittyBase        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittyBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyBaseCallerSession struct {
	Contract *KittyBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KittyBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyBaseTransactorSession struct {
	Contract     *KittyBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KittyBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyBaseRaw struct {
	Contract *KittyBase // Generic contract binding to access the raw methods on
}

// KittyBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyBaseCallerRaw struct {
	Contract *KittyBaseCaller // Generic read-only contract binding to access the raw methods on
}

// KittyBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyBaseTransactorRaw struct {
	Contract *KittyBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyBase creates a new instance of KittyBase, bound to a specific deployed contract.
func NewKittyBase(address common.Address, backend bind.ContractBackend) (*KittyBase, error) {
	contract, err := bindKittyBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyBase{KittyBaseCaller: KittyBaseCaller{contract: contract}, KittyBaseTransactor: KittyBaseTransactor{contract: contract}, KittyBaseFilterer: KittyBaseFilterer{contract: contract}}, nil
}

// NewKittyBaseCaller creates a new read-only instance of KittyBase, bound to a specific deployed contract.
func NewKittyBaseCaller(address common.Address, caller bind.ContractCaller) (*KittyBaseCaller, error) {
	contract, err := bindKittyBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyBaseCaller{contract: contract}, nil
}

// NewKittyBaseTransactor creates a new write-only instance of KittyBase, bound to a specific deployed contract.
func NewKittyBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyBaseTransactor, error) {
	contract, err := bindKittyBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyBaseTransactor{contract: contract}, nil
}

// NewKittyBaseFilterer creates a new log filterer instance of KittyBase, bound to a specific deployed contract.
func NewKittyBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyBaseFilterer, error) {
	contract, err := bindKittyBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyBaseFilterer{contract: contract}, nil
}

// bindKittyBase binds a generic wrapper to an already deployed contract.
func bindKittyBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyBase *KittyBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyBase.Contract.KittyBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyBase *KittyBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBase.Contract.KittyBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyBase *KittyBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyBase.Contract.KittyBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyBase *KittyBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyBase *KittyBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyBase *KittyBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyBase.Contract.contract.Transact(opts, method, params...)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyBase *KittyBaseCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBase.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyBase *KittyBaseSession) CeoAddress() (common.Address, error) {
	return _KittyBase.Contract.CeoAddress(&_KittyBase.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyBase *KittyBaseCallerSession) CeoAddress() (common.Address, error) {
	return _KittyBase.Contract.CeoAddress(&_KittyBase.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyBase *KittyBaseCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBase.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyBase *KittyBaseSession) CfoAddress() (common.Address, error) {
	return _KittyBase.Contract.CfoAddress(&_KittyBase.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyBase *KittyBaseCallerSession) CfoAddress() (common.Address, error) {
	return _KittyBase.Contract.CfoAddress(&_KittyBase.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyBase *KittyBaseCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBase.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyBase *KittyBaseSession) CooAddress() (common.Address, error) {
	return _KittyBase.Contract.CooAddress(&_KittyBase.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyBase *KittyBaseCallerSession) CooAddress() (common.Address, error) {
	return _KittyBase.Contract.CooAddress(&_KittyBase.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyBase *KittyBaseCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _KittyBase.contract.Call(opts, out, "cooldowns", arg0)
	return *ret0, err
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyBase *KittyBaseSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyBase.Contract.Cooldowns(&_KittyBase.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyBase *KittyBaseCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyBase.Contract.Cooldowns(&_KittyBase.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyBase *KittyBaseCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBase.contract.Call(opts, out, "kittyIndexToApproved", arg0)
	return *ret0, err
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyBase *KittyBaseSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyBase.Contract.KittyIndexToApproved(&_KittyBase.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyBase *KittyBaseCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyBase.Contract.KittyIndexToApproved(&_KittyBase.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyBase *KittyBaseCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBase.contract.Call(opts, out, "kittyIndexToOwner", arg0)
	return *ret0, err
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyBase *KittyBaseSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyBase.Contract.KittyIndexToOwner(&_KittyBase.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyBase *KittyBaseCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyBase.Contract.KittyIndexToOwner(&_KittyBase.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyBase *KittyBaseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyBase.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyBase *KittyBaseSession) Paused() (bool, error) {
	return _KittyBase.Contract.Paused(&_KittyBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyBase *KittyBaseCallerSession) Paused() (bool, error) {
	return _KittyBase.Contract.Paused(&_KittyBase.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyBase *KittyBaseCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBase.contract.Call(opts, out, "saleAuction")
	return *ret0, err
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyBase *KittyBaseSession) SaleAuction() (common.Address, error) {
	return _KittyBase.Contract.SaleAuction(&_KittyBase.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyBase *KittyBaseCallerSession) SaleAuction() (common.Address, error) {
	return _KittyBase.Contract.SaleAuction(&_KittyBase.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyBase *KittyBaseCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyBase.contract.Call(opts, out, "secondsPerBlock")
	return *ret0, err
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyBase *KittyBaseSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyBase.Contract.SecondsPerBlock(&_KittyBase.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyBase *KittyBaseCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyBase.Contract.SecondsPerBlock(&_KittyBase.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyBase *KittyBaseCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBase.contract.Call(opts, out, "sireAllowedToAddress", arg0)
	return *ret0, err
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyBase *KittyBaseSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyBase.Contract.SireAllowedToAddress(&_KittyBase.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyBase *KittyBaseCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyBase.Contract.SireAllowedToAddress(&_KittyBase.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyBase *KittyBaseCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBase.contract.Call(opts, out, "siringAuction")
	return *ret0, err
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyBase *KittyBaseSession) SiringAuction() (common.Address, error) {
	return _KittyBase.Contract.SiringAuction(&_KittyBase.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyBase *KittyBaseCallerSession) SiringAuction() (common.Address, error) {
	return _KittyBase.Contract.SiringAuction(&_KittyBase.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyBase *KittyBaseTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBase.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyBase *KittyBaseSession) Pause() (*types.Transaction, error) {
	return _KittyBase.Contract.Pause(&_KittyBase.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyBase *KittyBaseTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyBase.Contract.Pause(&_KittyBase.TransactOpts)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyBase *KittyBaseTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyBase.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyBase *KittyBaseSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyBase.Contract.SetCEO(&_KittyBase.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyBase *KittyBaseTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyBase.Contract.SetCEO(&_KittyBase.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyBase *KittyBaseTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyBase.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyBase *KittyBaseSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyBase.Contract.SetCFO(&_KittyBase.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyBase *KittyBaseTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyBase.Contract.SetCFO(&_KittyBase.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyBase *KittyBaseTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyBase.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyBase *KittyBaseSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyBase.Contract.SetCOO(&_KittyBase.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyBase *KittyBaseTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyBase.Contract.SetCOO(&_KittyBase.TransactOpts, _newCOO)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyBase *KittyBaseTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _KittyBase.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyBase *KittyBaseSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyBase.Contract.SetSecondsPerBlock(&_KittyBase.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyBase *KittyBaseTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyBase.Contract.SetSecondsPerBlock(&_KittyBase.TransactOpts, secs)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyBase *KittyBaseTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBase.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyBase *KittyBaseSession) Unpause() (*types.Transaction, error) {
	return _KittyBase.Contract.Unpause(&_KittyBase.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyBase *KittyBaseTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyBase.Contract.Unpause(&_KittyBase.TransactOpts)
}

// KittyBaseBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the KittyBase contract.
type KittyBaseBirthIterator struct {
	Event *KittyBaseBirth // Event containing the contract specifics and raw log

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
func (it *KittyBaseBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBaseBirth)
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
		it.Event = new(KittyBaseBirth)
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
func (it *KittyBaseBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBaseBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBaseBirth represents a Birth event raised by the KittyBase contract.
type KittyBaseBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_KittyBase *KittyBaseFilterer) FilterBirth(opts *bind.FilterOpts) (*KittyBaseBirthIterator, error) {

	logs, sub, err := _KittyBase.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &KittyBaseBirthIterator{contract: _KittyBase.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_KittyBase *KittyBaseFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *KittyBaseBirth) (event.Subscription, error) {

	logs, sub, err := _KittyBase.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBaseBirth)
				if err := _KittyBase.contract.UnpackLog(event, "Birth", log); err != nil {
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

// KittyBaseContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyBase contract.
type KittyBaseContractUpgradeIterator struct {
	Event *KittyBaseContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyBaseContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBaseContractUpgrade)
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
		it.Event = new(KittyBaseContractUpgrade)
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
func (it *KittyBaseContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBaseContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBaseContractUpgrade represents a ContractUpgrade event raised by the KittyBase contract.
type KittyBaseContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyBase *KittyBaseFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyBaseContractUpgradeIterator, error) {

	logs, sub, err := _KittyBase.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyBaseContractUpgradeIterator{contract: _KittyBase.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyBase *KittyBaseFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyBaseContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyBase.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBaseContractUpgrade)
				if err := _KittyBase.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// KittyBaseTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KittyBase contract.
type KittyBaseTransferIterator struct {
	Event *KittyBaseTransfer // Event containing the contract specifics and raw log

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
func (it *KittyBaseTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBaseTransfer)
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
		it.Event = new(KittyBaseTransfer)
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
func (it *KittyBaseTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBaseTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBaseTransfer represents a Transfer event raised by the KittyBase contract.
type KittyBaseTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_KittyBase *KittyBaseFilterer) FilterTransfer(opts *bind.FilterOpts) (*KittyBaseTransferIterator, error) {

	logs, sub, err := _KittyBase.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &KittyBaseTransferIterator{contract: _KittyBase.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_KittyBase *KittyBaseFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KittyBaseTransfer) (event.Subscription, error) {

	logs, sub, err := _KittyBase.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBaseTransfer)
				if err := _KittyBase.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// KittyBreedingABI is the input ABI used to generate the binding from.
const KittyBreedingABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pregnantKitties\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isPregnant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setGeneScienceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"canBreedWith\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAutoBirthFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"approveSiring\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"giveBirth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoBirthFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isReadyToBreed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"geneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"breedWithAuto\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cooldownEndBlock\",\"type\":\"uint256\"}],\"name\":\"Pregnant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyBreedingBin is the compiled bytecode used for deploying new contracts.
const KittyBreedingBin = `0x6002805460a060020a60ff0219169055610240604052603c6080908152607860a05261012c60c05261025860e05261070861010052610e1061012052611c2061014052613840610160526170806101805261e1006101a052620151806101c0526202a3006101e052620546006102005262093a80610220526200008790600390600e620000ac565b50600f60055566071afd498d0000600e55348015620000a557600080fd5b5062000176565b6002830191839082156200013d5791602002820160005b838211156200010957835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302620000c3565b80156200013b5782816101000a81549063ffffffff021916905560040160208160030104928301926001030262000109565b505b506200014b9291506200014f565b5090565b6200017391905b808211156200014b57805463ffffffff1916815560010162000156565b90565b6122a380620001866000396000f3006080604052600436106101df5763ffffffff60e060020a60003504166301ffc9a781146101e45780630519ce791461022f5780630560ff441461026057806306fdde03146102f9578063095ea7b31461030e5780630a0f81681461033457806318160ddd14610349578063183a7947146103705780631940a9361461038557806321717ebf1461039d57806323b872dd146103b257806324e7a38a146103dc57806327d7874c146103fd5780632ba73c151461041e5780633f4ba83a1461043f57806346116e6f1461045457806346d22c701461046c578063481af3d3146104875780634b85fd551461049f5780634dfff04f146104b75780634e0a3379146104db5780635663896e146104fc5780635c975abb146105145780636352211e1461052957806370a08231146105415780637a7d4937146105625780638456cb59146105775780638462151c1461058c57806388c2a0bf146105fd57806395d89b41146106155780639d6fac6f1461062a578063a45f4bfc1461065b578063a9059cbb14610673578063b047fb5014610697578063b0c35c05146106ac578063bc4006f5146106c1578063d3e6f49f146106d6578063e17b25af146106ee578063e6cbe3511461070f578063f2b47d5214610724578063f7d8c88314610739575b600080fd5b3480156101f057600080fd5b5061021b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1960043516610747565b604080519115158252519081900360200190f35b34801561023b57600080fd5b506102446109da565b60408051600160a060020a039092168252519081900360200190f35b34801561026c57600080fd5b506102846004803590602480359081019101356109e9565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102be5781810151838201526020016102a6565b50505050905090810190601f1680156102eb5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561030557600080fd5b50610284610aec565b34801561031a57600080fd5b50610332600160a060020a0360043516602435610b23565b005b34801561034057600080fd5b50610244610ba5565b34801561035557600080fd5b5061035e610bb4565b60408051918252519081900360200190f35b34801561037c57600080fd5b5061035e610bbe565b34801561039157600080fd5b5061021b600435610bc4565b3480156103a957600080fd5b50610244610c09565b3480156103be57600080fd5b50610332600160a060020a0360043581169060243516604435610c18565b3480156103e857600080fd5b50610332600160a060020a0360043516610c94565b34801561040957600080fd5b50610332600160a060020a0360043516610d47565b34801561042a57600080fd5b50610332600160a060020a0360043516610d95565b34801561044b57600080fd5b50610332610de3565b34801561046057600080fd5b50610244600435610e32565b34801561047857600080fd5b5061021b600435602435610e4d565b34801561049357600080fd5b50610244600435610ecd565b3480156104ab57600080fd5b50610332600435610ee8565b3480156104c357600080fd5b50610332600160a060020a0360043516602435610f04565b3480156104e757600080fd5b50610332600160a060020a0360043516610f5e565b34801561050857600080fd5b50610332600435610fac565b34801561052057600080fd5b5061021b611008565b34801561053557600080fd5b50610244600435611018565b34801561054d57600080fd5b5061035e600160a060020a036004351661103c565b34801561056e57600080fd5b5061035e611057565b34801561058357600080fd5b5061033261105d565b34801561059857600080fd5b506105ad600160a060020a03600435166110dd565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156105e95781810151838201526020016105d1565b505050509050019250505060405180910390f35b34801561060957600080fd5b5061035e6004356111af565b34801561062157600080fd5b5061028461148b565b34801561063657600080fd5b506106426004356114c2565b6040805163ffffffff9092168252519081900360200190f35b34801561066757600080fd5b506102446004356114ef565b34801561067f57600080fd5b50610332600160a060020a036004351660243561150a565b3480156106a357600080fd5b506102446115a6565b3480156106b857600080fd5b5061035e6115b5565b3480156106cd57600080fd5b506102446115bb565b3480156106e257600080fd5b5061021b6004356115ca565b3480156106fa57600080fd5b50610332600160a060020a0360043516611699565b34801561071b57600080fd5b506102446116d2565b34801561073057600080fd5b506102446116e1565b6103326004356024356116f0565b604080517f737570706f727473496e74657266616365286279746573342900000000000000815290519081900360190190206000907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19838116911614806109d25750604080517f746f6b656e4d657461646174612875696e743235362c737472696e67290000008152815190819003601d0181207f746f6b656e734f664f776e657228616464726573732900000000000000000000825282519182900360160182207f7472616e7366657246726f6d28616464726573732c616464726573732c75696e83527f7432353629000000000000000000000000000000000000000000000000000000602084015283519283900360250183207f7472616e7366657228616464726573732c75696e743235362900000000000000845284519384900360190184207f617070726f766528616464726573732c75696e74323536290000000000000000855285519485900360180185207f6f776e65724f662875696e743235362900000000000000000000000000000000865286519586900360100186207f62616c616e63654f662861646472657373290000000000000000000000000000875287519687900360120187207f746f74616c537570706c792829000000000000000000000000000000000000008852885197889003600d0188207f73796d626f6c2829000000000000000000000000000000000000000000000000895289519889900360080189207f6e616d65282900000000000000000000000000000000000000000000000000008a529951988990036006019098207bffffffffffffffffffffffffffffffffffffffffffffffffffffffff198c811691909a189098181818181818181891909116145b90505b919050565b600154600160a060020a031681565b60606109f3612214565b600d54600090600160a060020a03161515610a0d57600080fd5b600d54604080517fcb4799f2000000000000000000000000000000000000000000000000000000008152600481018981526024820192835260448201889052600160a060020a039093169263cb4799f2928a928a928a929091606401848480828437820191505094505050505060a060405180830381600087803b158015610a9457600080fd5b505af1158015610aa8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060a0811015610acd57600080fd5b5060808101519092509050610ae282826118f4565b9695505050505050565b60408051808201909152600d81527f43727970746f4b69747469657300000000000000000000000000000000000000602082015281565b60025460a060020a900460ff1615610b3a57600080fd5b610b443382611948565b1515610b4f57600080fd5b610b598183611968565b60408051338152600160a060020a038416602082015280820183905290517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360600190a15050565b600054600160a060020a031681565b6006546000190190565b600f5481565b6000808211610bd257600080fd5b6006805483908110610be057fe5b600091825260209091206002909102016001015460c060020a900463ffffffff16151592915050565b600c54600160a060020a031681565b60025460a060020a900460ff1615610c2f57600080fd5b600160a060020a0382161515610c4457600080fd5b600160a060020a038216301415610c5a57600080fd5b610c643382611996565b1515610c6f57600080fd5b610c798382611948565b1515610c8457600080fd5b610c8f8383836119b6565b505050565b60008054600160a060020a03163314610cac57600080fd5b81905080600160a060020a03166354c15b826040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610ced57600080fd5b505af1158015610d01573d6000803e3d6000fd5b505050506040513d6020811015610d1757600080fd5b50511515610d2457600080fd5b60108054600160a060020a031916600160a060020a039290921691909117905550565b600054600160a060020a03163314610d5e57600080fd5b600160a060020a0381161515610d7357600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314610dac57600080fd5b600160a060020a0381161515610dc157600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314610dfa57600080fd5b60025460a060020a900460ff161515610e1257600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600a60205260009081526040902054600160a060020a031681565b60008080808511610e5d57600080fd5b60008411610e6a57600080fd5b6006805486908110610e7857fe5b90600052602060002090600202019150600684815481101515610e9757fe5b90600052602060002090600202019050610eb382868387611a98565b8015610ec45750610ec48486611c18565b95945050505050565b600960205260009081526040902054600160a060020a031681565b600254600160a060020a03163314610eff57600080fd5b600e55565b60025460a060020a900460ff1615610f1b57600080fd5b610f253382611948565b1515610f3057600080fd5b6000908152600a602052604090208054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314610f7557600080fd5b600160a060020a0381161515610f8a57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b600254600160a060020a0316331480610fcf5750600054600160a060020a031633145b80610fe45750600154600160a060020a031633145b1515610fef57600080fd5b60035463ffffffff16811061100357600080fd5b600555565b60025460a060020a900460ff1681565b600081815260076020526040902054600160a060020a03168015156109d557600080fd5b600160a060020a031660009081526008602052604090205490565b60055481565b600254600160a060020a03163314806110805750600054600160a060020a031633145b806110955750600154600160a060020a031633145b15156110a057600080fd5b60025460a060020a900460ff16156110b757600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a179055565b60606000606060008060006110f18761103c565b94508415156111105760408051600081526020810190915295506111a5565b8460405190808252806020026020018201604052801561113a578160200160208202803883390190505b509350611145610bb4565b925060009150600190505b8281116111a157600081815260076020526040902054600160a060020a03888116911614156111995780848381518110151561118857fe5b602090810290910101526001909101905b600101611150565b8395505b5050505050919050565b600080600080600080600080600260149054906101000a900460ff161515156111d757600080fd5b600680548a9081106111e557fe5b60009182526020909120600290910201600181015490975067ffffffffffffffff16151561121257600080fd5b604080516101008101825288548152600189015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e08201526112a390611c6d565b15156112ae57600080fd5b60018701546006805460c060020a90920463ffffffff16975090879081106112d257fe5b600091825260209091206001808a015460029093029091019081015490965061ffff60f060020a92839004811696509190041684101561131f57600185015460f060020a900461ffff1693505b6010548754865460018a0154604080517f0d9f5aed0000000000000000000000000000000000000000000000000000000081526004810194909452602484019290925260001967ffffffffffffffff6801000000000000000090920482160116604483015251600160a060020a0390921691630d9f5aed916064808201926020929091908290030181600087803b1580156113b957600080fd5b505af11580156113cd573d6000803e3d6000fd5b505050506040513d60208110156113e357600080fd5b505160008a815260076020526040902054600189810154929550600160a060020a03909116935061142c918b9160c060020a90910463ffffffff1690870161ffff168686611c9d565b6001880180547bffffffff00000000000000000000000000000000000000000000000019169055600f8054600019019055600e54604051919250339181156108fc0291906000818181858888f150939c9b505050505050505050505050565b60408051808201909152600281527f434b000000000000000000000000000000000000000000000000000000000000602082015281565b600381600e81106114cf57fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b600760205260009081526040902054600160a060020a031681565b60025460a060020a900460ff161561152157600080fd5b600160a060020a038216151561153657600080fd5b600160a060020a03821630141561154c57600080fd5b600b54600160a060020a038381169116141561156757600080fd5b600c54600160a060020a038381169116141561158257600080fd5b61158c3382611948565b151561159757600080fd5b6115a23383836119b6565b5050565b600254600160a060020a031681565b600e5481565b600d54600160a060020a031681565b6000808083116115d957600080fd5b60068054849081106115e757fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e082015290915061169290611f7a565b9392505050565b600054600160a060020a031633146116b057600080fd5b600d8054600160a060020a031916600160a060020a0392909216919091179055565b600b54600160a060020a031681565b601054600160a060020a031681565b600254600090819060a060020a900460ff161561170c57600080fd5b600e5434101561171b57600080fd5b6117253385611948565b151561173057600080fd5b61173a8385611c18565b151561174557600080fd5b600680548590811061175357fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e08201529092506117fe90611f7a565b151561180957600080fd5b600680548490811061181757fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e08201529091506118c290611f7a565b15156118cd57600080fd5b6118d982858386611a98565b15156118e457600080fd5b6118ee8484611fa9565b50505050565b606080600080846040519080825280601f01601f191660200182016040528015611928578160200160208202803883390190505b509250506020820190508461193e8282876120e7565b5090949350505050565b600090815260076020526040902054600160a060020a0391821691161490565b6000918252600960205260409091208054600160a060020a031916600160a060020a03909216919091179055565b600090815260096020526040902054600160a060020a0391821691161490565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a0319169091179055831615611a4957600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b60408051600160a060020a0380861682528416602082015280820183905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a1505050565b600081841415611aaa57506000611c10565b6001850154608060020a900463ffffffff16821480611ad95750600185015460a060020a900463ffffffff1682145b15611ae657506000611c10565b6001830154608060020a900463ffffffff16841480611b155750600183015460a060020a900463ffffffff1684145b15611b2257506000611c10565b6001830154608060020a900463ffffffff161580611b4f57506001850154608060020a900463ffffffff16155b15611b5c57506001611c10565b60018581015490840154608060020a9182900463ffffffff90811692909104161480611ba7575060018086015490840154608060020a900463ffffffff90811660a060020a90920416145b15611bb457506000611c10565b6001808601549084015460a060020a900463ffffffff908116608060020a909204161480611bff57506001858101549084015460a060020a9182900463ffffffff9081169290910416145b15611c0c57506000611c10565b5060015b949350505050565b6000818152600760205260408082205484835290822054600160a060020a03918216911680821480610ec457506000858152600a6020526040902054600160a060020a03908116908316149250505092915050565b60008160a0015163ffffffff166000141580156109d25750506040015167ffffffffffffffff4381169116111590565b600080611ca8612233565b600063ffffffff89168914611cbc57600080fd5b63ffffffff88168814611cce57600080fd5b61ffff87168714611cde57600080fd5b600287049250600d8361ffff161115611cf657600d92505b505060408051610100810182528581524267ffffffffffffffff90811660208301908152600093830184815263ffffffff8c8116606086019081528c82166080870190815260a0870188815261ffff8a811660c08a019081528f821660e08b01908152600680546001810182559c528a517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f60028e029081019190915598517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d40909901805498519651955194519251915167ffffffffffffffff19909916998b16999099176fffffffffffffffff000000000000000019166801000000000000000096909a16959095029890981773ffffffff000000000000000000000000000000001916608060020a938616939093029290921777ffffffff0000000000000000000000000000000000000000191660a060020a91851691909102177bffffffff000000000000000000000000000000000000000000000000191660c060020a96841696909602959095177fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e060020a91861691909102177dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f060020a929094169190910292909217905590919081168114611ef457600080fd5b606080830151608080850151855160408051600160a060020a038c1681526020810188905263ffffffff95861681830152929094169482019490945290810192909252517f0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad59181900360a00190a1611f6e600086836119b6565b98975050505050505050565b60008160a0015163ffffffff1660001480156109d25750506040015167ffffffffffffffff4381169116111590565b600080600683815481101515611fbb57fe5b90600052602060002090600202019150600684815481101515611fda57fe5b600091825260209091206002909102016001810180547bffffffff000000000000000000000000000000000000000000000000191660c060020a63ffffffff871602179055905061202a8261212b565b6120338161212b565b6000848152600a602090815260408083208054600160a060020a031990811690915586845281842080549091169055600f80546001908101909155878452600783529281902054928401548151600160a060020a0390941684529183018790528281018690526801000000000000000090910467ffffffffffffffff166060830152517f241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80916080908290030190a150505050565b60005b6020821061210c578251845260209384019390920191601f19909101906120ea565b50905182516020929092036101000a6000190180199091169116179052565b600554600182015443919060039060e060020a900461ffff16600e811061214e57fe5b600891828204019190066004029054906101000a900463ffffffff1663ffffffff1681151561217957fe5b6001840180546fffffffffffffffff0000000000000000191668010000000000000000939092049390930167ffffffffffffffff16919091021790819055600d60e060020a90910461ffff161015612211576001818101805461ffff60e060020a8083048216909401169092027fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790555b50565b6080604051908101604052806004906020820280388339509192915050565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810191909152905600a165627a7a7230582006eb054987107c4b21eb82a592f14cf7cb1e288c24322070a04faf0bfa68e4380029`

// DeployKittyBreeding deploys a new Ethereum contract, binding an instance of KittyBreeding to it.
func DeployKittyBreeding(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyBreeding, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyBreedingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyBreedingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyBreeding{KittyBreedingCaller: KittyBreedingCaller{contract: contract}, KittyBreedingTransactor: KittyBreedingTransactor{contract: contract}, KittyBreedingFilterer: KittyBreedingFilterer{contract: contract}}, nil
}

// KittyBreeding is an auto generated Go binding around an Ethereum contract.
type KittyBreeding struct {
	KittyBreedingCaller     // Read-only binding to the contract
	KittyBreedingTransactor // Write-only binding to the contract
	KittyBreedingFilterer   // Log filterer for contract events
}

// KittyBreedingCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyBreedingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyBreedingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyBreedingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyBreedingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyBreedingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyBreedingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyBreedingSession struct {
	Contract     *KittyBreeding    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittyBreedingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyBreedingCallerSession struct {
	Contract *KittyBreedingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// KittyBreedingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyBreedingTransactorSession struct {
	Contract     *KittyBreedingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// KittyBreedingRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyBreedingRaw struct {
	Contract *KittyBreeding // Generic contract binding to access the raw methods on
}

// KittyBreedingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyBreedingCallerRaw struct {
	Contract *KittyBreedingCaller // Generic read-only contract binding to access the raw methods on
}

// KittyBreedingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyBreedingTransactorRaw struct {
	Contract *KittyBreedingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyBreeding creates a new instance of KittyBreeding, bound to a specific deployed contract.
func NewKittyBreeding(address common.Address, backend bind.ContractBackend) (*KittyBreeding, error) {
	contract, err := bindKittyBreeding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyBreeding{KittyBreedingCaller: KittyBreedingCaller{contract: contract}, KittyBreedingTransactor: KittyBreedingTransactor{contract: contract}, KittyBreedingFilterer: KittyBreedingFilterer{contract: contract}}, nil
}

// NewKittyBreedingCaller creates a new read-only instance of KittyBreeding, bound to a specific deployed contract.
func NewKittyBreedingCaller(address common.Address, caller bind.ContractCaller) (*KittyBreedingCaller, error) {
	contract, err := bindKittyBreeding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyBreedingCaller{contract: contract}, nil
}

// NewKittyBreedingTransactor creates a new write-only instance of KittyBreeding, bound to a specific deployed contract.
func NewKittyBreedingTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyBreedingTransactor, error) {
	contract, err := bindKittyBreeding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyBreedingTransactor{contract: contract}, nil
}

// NewKittyBreedingFilterer creates a new log filterer instance of KittyBreeding, bound to a specific deployed contract.
func NewKittyBreedingFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyBreedingFilterer, error) {
	contract, err := bindKittyBreeding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyBreedingFilterer{contract: contract}, nil
}

// bindKittyBreeding binds a generic wrapper to an already deployed contract.
func bindKittyBreeding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyBreedingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyBreeding *KittyBreedingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyBreeding.Contract.KittyBreedingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyBreeding *KittyBreedingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBreeding.Contract.KittyBreedingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyBreeding *KittyBreedingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyBreeding.Contract.KittyBreedingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyBreeding *KittyBreedingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyBreeding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyBreeding *KittyBreedingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBreeding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyBreeding *KittyBreedingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyBreeding.Contract.contract.Transact(opts, method, params...)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_KittyBreeding *KittyBreedingCaller) AutoBirthFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "autoBirthFee")
	return *ret0, err
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_KittyBreeding *KittyBreedingSession) AutoBirthFee() (*big.Int, error) {
	return _KittyBreeding.Contract.AutoBirthFee(&_KittyBreeding.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_KittyBreeding *KittyBreedingCallerSession) AutoBirthFee() (*big.Int, error) {
	return _KittyBreeding.Contract.AutoBirthFee(&_KittyBreeding.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyBreeding *KittyBreedingCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyBreeding *KittyBreedingSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyBreeding.Contract.BalanceOf(&_KittyBreeding.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyBreeding *KittyBreedingCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyBreeding.Contract.BalanceOf(&_KittyBreeding.CallOpts, _owner)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_KittyBreeding *KittyBreedingCaller) CanBreedWith(opts *bind.CallOpts, _matronId *big.Int, _sireId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "canBreedWith", _matronId, _sireId)
	return *ret0, err
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_KittyBreeding *KittyBreedingSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyBreeding.Contract.CanBreedWith(&_KittyBreeding.CallOpts, _matronId, _sireId)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_KittyBreeding *KittyBreedingCallerSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyBreeding.Contract.CanBreedWith(&_KittyBreeding.CallOpts, _matronId, _sireId)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyBreeding *KittyBreedingCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyBreeding *KittyBreedingSession) CeoAddress() (common.Address, error) {
	return _KittyBreeding.Contract.CeoAddress(&_KittyBreeding.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) CeoAddress() (common.Address, error) {
	return _KittyBreeding.Contract.CeoAddress(&_KittyBreeding.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyBreeding *KittyBreedingCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyBreeding *KittyBreedingSession) CfoAddress() (common.Address, error) {
	return _KittyBreeding.Contract.CfoAddress(&_KittyBreeding.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) CfoAddress() (common.Address, error) {
	return _KittyBreeding.Contract.CfoAddress(&_KittyBreeding.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyBreeding *KittyBreedingCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyBreeding *KittyBreedingSession) CooAddress() (common.Address, error) {
	return _KittyBreeding.Contract.CooAddress(&_KittyBreeding.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) CooAddress() (common.Address, error) {
	return _KittyBreeding.Contract.CooAddress(&_KittyBreeding.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyBreeding *KittyBreedingCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "cooldowns", arg0)
	return *ret0, err
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyBreeding *KittyBreedingSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyBreeding.Contract.Cooldowns(&_KittyBreeding.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyBreeding *KittyBreedingCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyBreeding.Contract.Cooldowns(&_KittyBreeding.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyBreeding *KittyBreedingCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "erc721Metadata")
	return *ret0, err
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyBreeding *KittyBreedingSession) Erc721Metadata() (common.Address, error) {
	return _KittyBreeding.Contract.Erc721Metadata(&_KittyBreeding.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) Erc721Metadata() (common.Address, error) {
	return _KittyBreeding.Contract.Erc721Metadata(&_KittyBreeding.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_KittyBreeding *KittyBreedingCaller) GeneScience(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "geneScience")
	return *ret0, err
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_KittyBreeding *KittyBreedingSession) GeneScience() (common.Address, error) {
	return _KittyBreeding.Contract.GeneScience(&_KittyBreeding.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) GeneScience() (common.Address, error) {
	return _KittyBreeding.Contract.GeneScience(&_KittyBreeding.CallOpts)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_KittyBreeding *KittyBreedingCaller) IsPregnant(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "isPregnant", _kittyId)
	return *ret0, err
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_KittyBreeding *KittyBreedingSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyBreeding.Contract.IsPregnant(&_KittyBreeding.CallOpts, _kittyId)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_KittyBreeding *KittyBreedingCallerSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyBreeding.Contract.IsPregnant(&_KittyBreeding.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_KittyBreeding *KittyBreedingCaller) IsReadyToBreed(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "isReadyToBreed", _kittyId)
	return *ret0, err
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_KittyBreeding *KittyBreedingSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyBreeding.Contract.IsReadyToBreed(&_KittyBreeding.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_KittyBreeding *KittyBreedingCallerSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyBreeding.Contract.IsReadyToBreed(&_KittyBreeding.CallOpts, _kittyId)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyBreeding *KittyBreedingCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "kittyIndexToApproved", arg0)
	return *ret0, err
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyBreeding *KittyBreedingSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.KittyIndexToApproved(&_KittyBreeding.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.KittyIndexToApproved(&_KittyBreeding.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyBreeding *KittyBreedingCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "kittyIndexToOwner", arg0)
	return *ret0, err
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyBreeding *KittyBreedingSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.KittyIndexToOwner(&_KittyBreeding.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.KittyIndexToOwner(&_KittyBreeding.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyBreeding *KittyBreedingCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyBreeding *KittyBreedingSession) Name() (string, error) {
	return _KittyBreeding.Contract.Name(&_KittyBreeding.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyBreeding *KittyBreedingCallerSession) Name() (string, error) {
	return _KittyBreeding.Contract.Name(&_KittyBreeding.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyBreeding *KittyBreedingCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyBreeding *KittyBreedingSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.OwnerOf(&_KittyBreeding.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyBreeding *KittyBreedingCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.OwnerOf(&_KittyBreeding.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyBreeding *KittyBreedingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyBreeding *KittyBreedingSession) Paused() (bool, error) {
	return _KittyBreeding.Contract.Paused(&_KittyBreeding.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyBreeding *KittyBreedingCallerSession) Paused() (bool, error) {
	return _KittyBreeding.Contract.Paused(&_KittyBreeding.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_KittyBreeding *KittyBreedingCaller) PregnantKitties(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "pregnantKitties")
	return *ret0, err
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_KittyBreeding *KittyBreedingSession) PregnantKitties() (*big.Int, error) {
	return _KittyBreeding.Contract.PregnantKitties(&_KittyBreeding.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_KittyBreeding *KittyBreedingCallerSession) PregnantKitties() (*big.Int, error) {
	return _KittyBreeding.Contract.PregnantKitties(&_KittyBreeding.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyBreeding *KittyBreedingCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "saleAuction")
	return *ret0, err
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyBreeding *KittyBreedingSession) SaleAuction() (common.Address, error) {
	return _KittyBreeding.Contract.SaleAuction(&_KittyBreeding.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) SaleAuction() (common.Address, error) {
	return _KittyBreeding.Contract.SaleAuction(&_KittyBreeding.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyBreeding *KittyBreedingCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "secondsPerBlock")
	return *ret0, err
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyBreeding *KittyBreedingSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyBreeding.Contract.SecondsPerBlock(&_KittyBreeding.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyBreeding *KittyBreedingCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyBreeding.Contract.SecondsPerBlock(&_KittyBreeding.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyBreeding *KittyBreedingCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "sireAllowedToAddress", arg0)
	return *ret0, err
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyBreeding *KittyBreedingSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.SireAllowedToAddress(&_KittyBreeding.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.SireAllowedToAddress(&_KittyBreeding.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyBreeding *KittyBreedingCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "siringAuction")
	return *ret0, err
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyBreeding *KittyBreedingSession) SiringAuction() (common.Address, error) {
	return _KittyBreeding.Contract.SiringAuction(&_KittyBreeding.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) SiringAuction() (common.Address, error) {
	return _KittyBreeding.Contract.SiringAuction(&_KittyBreeding.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyBreeding *KittyBreedingCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "supportsInterface", _interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyBreeding *KittyBreedingSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyBreeding.Contract.SupportsInterface(&_KittyBreeding.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyBreeding *KittyBreedingCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyBreeding.Contract.SupportsInterface(&_KittyBreeding.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyBreeding *KittyBreedingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyBreeding *KittyBreedingSession) Symbol() (string, error) {
	return _KittyBreeding.Contract.Symbol(&_KittyBreeding.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyBreeding *KittyBreedingCallerSession) Symbol() (string, error) {
	return _KittyBreeding.Contract.Symbol(&_KittyBreeding.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyBreeding *KittyBreedingCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "tokenMetadata", _tokenId, _preferredTransport)
	return *ret0, err
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyBreeding *KittyBreedingSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyBreeding.Contract.TokenMetadata(&_KittyBreeding.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyBreeding *KittyBreedingCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyBreeding.Contract.TokenMetadata(&_KittyBreeding.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyBreeding *KittyBreedingCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyBreeding *KittyBreedingSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyBreeding.Contract.TokensOfOwner(&_KittyBreeding.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyBreeding *KittyBreedingCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyBreeding.Contract.TokensOfOwner(&_KittyBreeding.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyBreeding *KittyBreedingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyBreeding.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyBreeding *KittyBreedingSession) TotalSupply() (*big.Int, error) {
	return _KittyBreeding.Contract.TotalSupply(&_KittyBreeding.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyBreeding *KittyBreedingCallerSession) TotalSupply() (*big.Int, error) {
	return _KittyBreeding.Contract.TotalSupply(&_KittyBreeding.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyBreeding *KittyBreedingTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyBreeding *KittyBreedingSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.Approve(&_KittyBreeding.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.Approve(&_KittyBreeding.TransactOpts, _to, _tokenId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_KittyBreeding *KittyBreedingTransactor) ApproveSiring(opts *bind.TransactOpts, _addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "approveSiring", _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_KittyBreeding *KittyBreedingSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.ApproveSiring(&_KittyBreeding.TransactOpts, _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.ApproveSiring(&_KittyBreeding.TransactOpts, _addr, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_KittyBreeding *KittyBreedingTransactor) BreedWithAuto(opts *bind.TransactOpts, _matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "breedWithAuto", _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_KittyBreeding *KittyBreedingSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.BreedWithAuto(&_KittyBreeding.TransactOpts, _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.BreedWithAuto(&_KittyBreeding.TransactOpts, _matronId, _sireId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_KittyBreeding *KittyBreedingTransactor) GiveBirth(opts *bind.TransactOpts, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "giveBirth", _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_KittyBreeding *KittyBreedingSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.GiveBirth(&_KittyBreeding.TransactOpts, _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_KittyBreeding *KittyBreedingTransactorSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.GiveBirth(&_KittyBreeding.TransactOpts, _matronId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyBreeding *KittyBreedingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyBreeding *KittyBreedingSession) Pause() (*types.Transaction, error) {
	return _KittyBreeding.Contract.Pause(&_KittyBreeding.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyBreeding *KittyBreedingTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyBreeding.Contract.Pause(&_KittyBreeding.TransactOpts)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetAutoBirthFee(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setAutoBirthFee", val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_KittyBreeding *KittyBreedingSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetAutoBirthFee(&_KittyBreeding.TransactOpts, val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetAutoBirthFee(&_KittyBreeding.TransactOpts, val)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyBreeding *KittyBreedingSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetCEO(&_KittyBreeding.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetCEO(&_KittyBreeding.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyBreeding *KittyBreedingSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetCFO(&_KittyBreeding.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetCFO(&_KittyBreeding.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyBreeding *KittyBreedingSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetCOO(&_KittyBreeding.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetCOO(&_KittyBreeding.TransactOpts, _newCOO)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetGeneScienceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setGeneScienceAddress", _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_KittyBreeding *KittyBreedingSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetGeneScienceAddress(&_KittyBreeding.TransactOpts, _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetGeneScienceAddress(&_KittyBreeding.TransactOpts, _address)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyBreeding *KittyBreedingSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetMetadataAddress(&_KittyBreeding.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetMetadataAddress(&_KittyBreeding.TransactOpts, _contractAddress)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyBreeding *KittyBreedingSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetSecondsPerBlock(&_KittyBreeding.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetSecondsPerBlock(&_KittyBreeding.TransactOpts, secs)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyBreeding *KittyBreedingTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyBreeding *KittyBreedingSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.Transfer(&_KittyBreeding.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.Transfer(&_KittyBreeding.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyBreeding *KittyBreedingTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyBreeding *KittyBreedingSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.TransferFrom(&_KittyBreeding.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.TransferFrom(&_KittyBreeding.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyBreeding *KittyBreedingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyBreeding *KittyBreedingSession) Unpause() (*types.Transaction, error) {
	return _KittyBreeding.Contract.Unpause(&_KittyBreeding.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyBreeding *KittyBreedingTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyBreeding.Contract.Unpause(&_KittyBreeding.TransactOpts)
}

// KittyBreedingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KittyBreeding contract.
type KittyBreedingApprovalIterator struct {
	Event *KittyBreedingApproval // Event containing the contract specifics and raw log

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
func (it *KittyBreedingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBreedingApproval)
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
		it.Event = new(KittyBreedingApproval)
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
func (it *KittyBreedingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBreedingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBreedingApproval represents a Approval event raised by the KittyBreeding contract.
type KittyBreedingApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_KittyBreeding *KittyBreedingFilterer) FilterApproval(opts *bind.FilterOpts) (*KittyBreedingApprovalIterator, error) {

	logs, sub, err := _KittyBreeding.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &KittyBreedingApprovalIterator{contract: _KittyBreeding.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_KittyBreeding *KittyBreedingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KittyBreedingApproval) (event.Subscription, error) {

	logs, sub, err := _KittyBreeding.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBreedingApproval)
				if err := _KittyBreeding.contract.UnpackLog(event, "Approval", log); err != nil {
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

// KittyBreedingBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the KittyBreeding contract.
type KittyBreedingBirthIterator struct {
	Event *KittyBreedingBirth // Event containing the contract specifics and raw log

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
func (it *KittyBreedingBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBreedingBirth)
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
		it.Event = new(KittyBreedingBirth)
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
func (it *KittyBreedingBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBreedingBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBreedingBirth represents a Birth event raised by the KittyBreeding contract.
type KittyBreedingBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_KittyBreeding *KittyBreedingFilterer) FilterBirth(opts *bind.FilterOpts) (*KittyBreedingBirthIterator, error) {

	logs, sub, err := _KittyBreeding.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &KittyBreedingBirthIterator{contract: _KittyBreeding.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_KittyBreeding *KittyBreedingFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *KittyBreedingBirth) (event.Subscription, error) {

	logs, sub, err := _KittyBreeding.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBreedingBirth)
				if err := _KittyBreeding.contract.UnpackLog(event, "Birth", log); err != nil {
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

// KittyBreedingContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyBreeding contract.
type KittyBreedingContractUpgradeIterator struct {
	Event *KittyBreedingContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyBreedingContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBreedingContractUpgrade)
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
		it.Event = new(KittyBreedingContractUpgrade)
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
func (it *KittyBreedingContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBreedingContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBreedingContractUpgrade represents a ContractUpgrade event raised by the KittyBreeding contract.
type KittyBreedingContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyBreeding *KittyBreedingFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyBreedingContractUpgradeIterator, error) {

	logs, sub, err := _KittyBreeding.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyBreedingContractUpgradeIterator{contract: _KittyBreeding.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyBreeding *KittyBreedingFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyBreedingContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyBreeding.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBreedingContractUpgrade)
				if err := _KittyBreeding.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// KittyBreedingPregnantIterator is returned from FilterPregnant and is used to iterate over the raw logs and unpacked data for Pregnant events raised by the KittyBreeding contract.
type KittyBreedingPregnantIterator struct {
	Event *KittyBreedingPregnant // Event containing the contract specifics and raw log

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
func (it *KittyBreedingPregnantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBreedingPregnant)
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
		it.Event = new(KittyBreedingPregnant)
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
func (it *KittyBreedingPregnantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBreedingPregnantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBreedingPregnant represents a Pregnant event raised by the KittyBreeding contract.
type KittyBreedingPregnant struct {
	Owner            common.Address
	MatronId         *big.Int
	SireId           *big.Int
	CooldownEndBlock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPregnant is a free log retrieval operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: e Pregnant(owner address, matronId uint256, sireId uint256, cooldownEndBlock uint256)
func (_KittyBreeding *KittyBreedingFilterer) FilterPregnant(opts *bind.FilterOpts) (*KittyBreedingPregnantIterator, error) {

	logs, sub, err := _KittyBreeding.contract.FilterLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return &KittyBreedingPregnantIterator{contract: _KittyBreeding.contract, event: "Pregnant", logs: logs, sub: sub}, nil
}

// WatchPregnant is a free log subscription operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: e Pregnant(owner address, matronId uint256, sireId uint256, cooldownEndBlock uint256)
func (_KittyBreeding *KittyBreedingFilterer) WatchPregnant(opts *bind.WatchOpts, sink chan<- *KittyBreedingPregnant) (event.Subscription, error) {

	logs, sub, err := _KittyBreeding.contract.WatchLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBreedingPregnant)
				if err := _KittyBreeding.contract.UnpackLog(event, "Pregnant", log); err != nil {
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

// KittyBreedingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KittyBreeding contract.
type KittyBreedingTransferIterator struct {
	Event *KittyBreedingTransfer // Event containing the contract specifics and raw log

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
func (it *KittyBreedingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBreedingTransfer)
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
		it.Event = new(KittyBreedingTransfer)
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
func (it *KittyBreedingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBreedingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBreedingTransfer represents a Transfer event raised by the KittyBreeding contract.
type KittyBreedingTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_KittyBreeding *KittyBreedingFilterer) FilterTransfer(opts *bind.FilterOpts) (*KittyBreedingTransferIterator, error) {

	logs, sub, err := _KittyBreeding.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &KittyBreedingTransferIterator{contract: _KittyBreeding.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_KittyBreeding *KittyBreedingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KittyBreedingTransfer) (event.Subscription, error) {

	logs, sub, err := _KittyBreeding.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBreedingTransfer)
				if err := _KittyBreeding.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// KittyCoreABI is the input ABI used to generate the binding from.
const KittyCoreABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"promoCreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_STARTING_PRICE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSiringAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pregnantKitties\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isPregnant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_AUCTION_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setGeneScienceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSaleAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"canBreedWith\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSiringAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAutoBirthFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"approveSiring\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genes\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createPromoKitty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSaleAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_v2Address\",\"type\":\"address\"}],\"name\":\"setNewAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"giveBirth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAuctionBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoBirthFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genes\",\"type\":\"uint256\"}],\"name\":\"createGen0Auction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isReadyToBreed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PROMO_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getKitty\",\"outputs\":[{\"name\":\"isGestating\",\"type\":\"bool\"},{\"name\":\"isReady\",\"type\":\"bool\"},{\"name\":\"cooldownIndex\",\"type\":\"uint256\"},{\"name\":\"nextActionAt\",\"type\":\"uint256\"},{\"name\":\"siringWithId\",\"type\":\"uint256\"},{\"name\":\"birthTime\",\"type\":\"uint256\"},{\"name\":\"matronId\",\"type\":\"uint256\"},{\"name\":\"sireId\",\"type\":\"uint256\"},{\"name\":\"generation\",\"type\":\"uint256\"},{\"name\":\"genes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sireId\",\"type\":\"uint256\"},{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"bidOnSiringAuction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gen0CreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"geneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"breedWithAuto\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cooldownEndBlock\",\"type\":\"uint256\"}],\"name\":\"Pregnant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyCoreBin is the compiled bytecode used for deploying new contracts.
const KittyCoreBin = `0x6002805460a060020a60ff0219169055610240604052603c6080908152607860a05261012c60c05261025860e05261070861010052610e1061012052611c2061014052613840610160526170806101805261e1006101a052620151806101c0526202a3006101e052620546006102005262093a80610220526200008790600390600e620004e6565b50600f60055566071afd498d0000600e55348015620000a557600080fd5b506002805460008054600160a060020a031990811633908117835560a060020a60ff021990931674010000000000000000000000000000000000000000171690911790915562000104908080600019816401000000006200010b810204565b50620005f4565b6000806200011862000589565b600063ffffffff891689146200012d57600080fd5b63ffffffff881688146200014057600080fd5b61ffff871687146200015157600080fd5b600287049250600d8361ffff1611156200016a57600d92505b505060408051610100810182528581524267ffffffffffffffff90811660208301908152600093830184815263ffffffff8c8116606086019081528c82166080870190815260a0870188815261ffff8a811660c08a019081528f821660e08b01908152600680546001810182559c528a517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f60028e029081019190915598517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d40909901805498519651955194519251915167ffffffffffffffff19909916998b1699909917604060020a608060020a0319166801000000000000000096909a169590950298909817608060020a63ffffffff021916700100000000000000000000000000000000938616939093029290921760a060020a63ffffffff02191674010000000000000000000000000000000000000000918516919091021760c060020a63ffffffff0219167801000000000000000000000000000000000000000000000000968416969096029590951760e060020a61ffff0219167c01000000000000000000000000000000000000000000000000000000009186169190910217600160f060020a03167e010000000000000000000000000000000000000000000000000000000000009290941691909102929092179055909190811681146200037257600080fd5b606080830151608080850151855160408051600160a060020a038c1681526020810188905263ffffffff95861681830152929094169482019490945290810192909252517f0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad59181900360a00190a1620003f76000868364010000000062000403810204565b98975050505050505050565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a03191690911790558316156200049757600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b60408051600160a060020a0380861682528416602082015280820183905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a1505050565b600283019183908215620005775791602002820160005b838211156200054357835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302620004fd565b8015620005755782816101000a81549063ffffffff021916905560040160208160030104928301926001030262000543565b505b5062000585929150620005cd565b5090565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915290565b620005f191905b808211156200058557805463ffffffff19168155600101620005d4565b90565b6130ff80620006046000396000f3006080604052600436106102a55763ffffffff60e060020a60003504166301ffc9a781146102d55780630519ce79146103205780630560ff441461035157806305e45546146103ea57806306fdde0314610411578063095ea7b3146104265780630a0f81681461044a5780630e583df01461045f57806314001f4c1461047457806318160ddd14610495578063183a7947146104aa5780631940a936146104bf57806319c2f201146104d757806321717ebf146104ec57806323b872dd1461050157806324e7a38a1461052b57806327d7874c1461054c5780632ba73c151461056d5780633d7d3f5a1461058e5780633f4ba83a146105af57806346116e6f146105c457806346d22c70146105dc578063481af3d3146105f75780634ad8c9381461060f5780634b85fd55146106305780634dfff04f146106485780634e0a33791461066c578063561291341461068d5780635663896e146106b15780635c975abb146106c95780635fd8c710146106de5780636352211e146106f3578063680eba271461070b5780636af04a57146107205780636fbde40d1461073557806370a082311461075657806371587988146107775780637a7d4937146107985780638456cb59146107ad5780638462151c146107c257806388c2a0bf1461083357806391876e571461084b57806395d89b41146108605780639d6fac6f14610875578063a45f4bfc146108a6578063a9059cbb146108be578063b047fb50146108e2578063b0c35c05146108f7578063bc4006f51461090c578063c3bea9af14610921578063d3e6f49f14610939578063defb958414610951578063e17b25af14610966578063e6cbe35114610987578063e98b7f4d1461099c578063ed60ade614610a06578063f1ca941014610a14578063f2b47d5214610a29578063f7d8c88314610a3e575b600b54600160a060020a03163314806102c85750600c54600160a060020a031633145b15156102d357600080fd5b005b3480156102e157600080fd5b5061030c7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1960043516610a4c565b604080519115158252519081900360200190f35b34801561032c57600080fd5b50610335610cdf565b60408051600160a060020a039092168252519081900360200190f35b34801561035d57600080fd5b50610375600480359060248035908101910135610cee565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103af578181015183820152602001610397565b50505050905090810190601f1680156103dc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103f657600080fd5b506103ff610df1565b60408051918252519081900360200190f35b34801561041d57600080fd5b50610375610df7565b34801561043257600080fd5b506102d3600160a060020a0360043516602435610e2e565b34801561045657600080fd5b50610335610eb0565b34801561046b57600080fd5b506103ff610ebf565b34801561048057600080fd5b506102d3600160a060020a0360043516610eca565b3480156104a157600080fd5b506103ff610f7d565b3480156104b657600080fd5b506103ff610f87565b3480156104cb57600080fd5b5061030c600435610f8d565b3480156104e357600080fd5b506103ff610fd2565b3480156104f857600080fd5b50610335610fd9565b34801561050d57600080fd5b506102d3600160a060020a0360043581169060243516604435610fe8565b34801561053757600080fd5b506102d3600160a060020a0360043516611064565b34801561055857600080fd5b506102d3600160a060020a0360043516611117565b34801561057957600080fd5b506102d3600160a060020a0360043516611165565b34801561059a57600080fd5b506102d36004356024356044356064356111b3565b3480156105bb57600080fd5b506102d36112a8565b3480156105d057600080fd5b5061033560043561133c565b3480156105e857600080fd5b5061030c600435602435611357565b34801561060357600080fd5b506103356004356113d7565b34801561061b57600080fd5b506102d36004356024356044356064356113f2565b34801561063c57600080fd5b506102d36004356114ca565b34801561065457600080fd5b506102d3600160a060020a03600435166024356114e6565b34801561067857600080fd5b506102d3600160a060020a0360043516611540565b34801561069957600080fd5b506102d3600435600160a060020a036024351661158e565b3480156106bd57600080fd5b506102d36004356115f5565b3480156106d557600080fd5b5061030c611651565b3480156106ea57600080fd5b506102d3611661565b3480156106ff57600080fd5b506103356004356116c5565b34801561071757600080fd5b506103ff6116e9565b34801561072c57600080fd5b506103356116ef565b34801561074157600080fd5b506102d3600160a060020a03600435166116fe565b34801561076257600080fd5b506103ff600160a060020a03600435166117b1565b34801561078357600080fd5b506102d3600160a060020a03600435166117cc565b3480156107a457600080fd5b506103ff61184f565b3480156107b957600080fd5b506102d3611855565b3480156107ce57600080fd5b506107e3600160a060020a03600435166118d5565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561081f578181015183820152602001610807565b505050509050019250505060405180910390f35b34801561083f57600080fd5b506103ff6004356119a7565b34801561085757600080fd5b506102d3611c83565b34801561086c57600080fd5b50610375611d98565b34801561088157600080fd5b5061088d600435611dcf565b6040805163ffffffff9092168252519081900360200190f35b3480156108b257600080fd5b50610335600435611dfc565b3480156108ca57600080fd5b506102d3600160a060020a0360043516602435611e17565b3480156108ee57600080fd5b50610335611eaf565b34801561090357600080fd5b506103ff611ebe565b34801561091857600080fd5b50610335611ec4565b34801561092d57600080fd5b506102d3600435611ed3565b34801561094557600080fd5b5061030c600435611fbd565b34801561095d57600080fd5b506103ff61208c565b34801561097257600080fd5b506102d3600160a060020a0360043516612092565b34801561099357600080fd5b506103356120cb565b3480156109a857600080fd5b506109b46004356120da565b604080519a15158b5298151560208b0152898901979097526060890195909552608088019390935260a087019190915260c086015260e085015261010084015261012083015251908190036101400190f35b6102d360043560243561223b565b348015610a2057600080fd5b506103ff6123d5565b348015610a3557600080fd5b506103356123db565b6102d36004356024356123ea565b604080517f737570706f727473496e74657266616365286279746573342900000000000000815290519081900360190190206000907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1983811691161480610cd75750604080517f746f6b656e4d657461646174612875696e743235362c737472696e67290000008152815190819003601d0181207f746f6b656e734f664f776e657228616464726573732900000000000000000000825282519182900360160182207f7472616e7366657246726f6d28616464726573732c616464726573732c75696e83527f7432353629000000000000000000000000000000000000000000000000000000602084015283519283900360250183207f7472616e7366657228616464726573732c75696e743235362900000000000000845284519384900360190184207f617070726f766528616464726573732c75696e74323536290000000000000000855285519485900360180185207f6f776e65724f662875696e743235362900000000000000000000000000000000865286519586900360100186207f62616c616e63654f662861646472657373290000000000000000000000000000875287519687900360120187207f746f74616c537570706c792829000000000000000000000000000000000000008852885197889003600d0188207f73796d626f6c2829000000000000000000000000000000000000000000000000895289519889900360080189207f6e616d65282900000000000000000000000000000000000000000000000000008a529951988990036006019098207bffffffffffffffffffffffffffffffffffffffffffffffffffffffff198c811691909a189098181818181818181891909116145b90505b919050565b600154600160a060020a031681565b6060610cf8613070565b600d54600090600160a060020a03161515610d1257600080fd5b600d54604080517fcb4799f2000000000000000000000000000000000000000000000000000000008152600481018981526024820192835260448201889052600160a060020a039093169263cb4799f2928a928a928a929091606401848480828437820191505094505050505060a060405180830381600087803b158015610d9957600080fd5b505af1158015610dad573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060a0811015610dd257600080fd5b5060808101519092509050610de782826125e8565b9695505050505050565b60115481565b60408051808201909152600d81527f43727970746f4b69747469657300000000000000000000000000000000000000602082015281565b60025460a060020a900460ff1615610e4557600080fd5b610e4f338261263c565b1515610e5a57600080fd5b610e64818361265c565b60408051338152600160a060020a038416602082015280820183905290517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360600190a15050565b600054600160a060020a031681565b662386f26fc1000081565b60008054600160a060020a03163314610ee257600080fd5b81905080600160a060020a03166376190f8f6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610f2357600080fd5b505af1158015610f37573d6000803e3d6000fd5b505050506040513d6020811015610f4d57600080fd5b50511515610f5a57600080fd5b600c8054600160a060020a031916600160a060020a039290921691909117905550565b6006546000190190565b600f5481565b6000808211610f9b57600080fd5b6006805483908110610fa957fe5b600091825260209091206002909102016001015460c060020a900463ffffffff16151592915050565b6201518081565b600c54600160a060020a031681565b60025460a060020a900460ff1615610fff57600080fd5b600160a060020a038216151561101457600080fd5b600160a060020a03821630141561102a57600080fd5b611034338261268a565b151561103f57600080fd5b611049838261263c565b151561105457600080fd5b61105f8383836126aa565b505050565b60008054600160a060020a0316331461107c57600080fd5b81905080600160a060020a03166354c15b826040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156110bd57600080fd5b505af11580156110d1573d6000803e3d6000fd5b505050506040513d60208110156110e757600080fd5b505115156110f457600080fd5b60108054600160a060020a031916600160a060020a039290921691909117905550565b600054600160a060020a0316331461112e57600080fd5b600160a060020a038116151561114357600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a0316331461117c57600080fd5b600160a060020a038116151561119157600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b60025460a060020a900460ff16156111ca57600080fd5b6111d4338561263c565b15156111df57600080fd5b6111e884610f8d565b156111f257600080fd5b600b54611209908590600160a060020a031661265c565b600b54604080517f27ebe40a000000000000000000000000000000000000000000000000000000008152600481018790526024810186905260448101859052606481018490523360848201529051600160a060020a03909216916327ebe40a9160a48082019260009290919082900301818387803b15801561128a57600080fd5b505af115801561129e573d6000803e3d6000fd5b5050505050505050565b600054600160a060020a031633146112bf57600080fd5b60025460a060020a900460ff1615156112d757600080fd5b600b54600160a060020a031615156112ee57600080fd5b600c54600160a060020a0316151561130557600080fd5b601054600160a060020a0316151561131c57600080fd5b601354600160a060020a03161561133257600080fd5b61133a61278c565b565b600a60205260009081526040902054600160a060020a031681565b6000808080851161136757600080fd5b6000841161137457600080fd5b600680548690811061138257fe5b906000526020600020906002020191506006848154811015156113a157fe5b906000526020600020906002020190506113bd828683876127db565b80156113ce57506113ce848661295b565b95945050505050565b600960205260009081526040902054600160a060020a031681565b60025460a060020a900460ff161561140957600080fd5b611413338561263c565b151561141e57600080fd5b61142784611fbd565b151561143257600080fd5b600c54611449908590600160a060020a031661265c565b600c54604080517f27ebe40a000000000000000000000000000000000000000000000000000000008152600481018790526024810186905260448101859052606481018490523360848201529051600160a060020a03909216916327ebe40a9160a48082019260009290919082900301818387803b15801561128a57600080fd5b600254600160a060020a031633146114e157600080fd5b600e55565b60025460a060020a900460ff16156114fd57600080fd5b611507338261263c565b151561151257600080fd5b6000908152600a602052604090208054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a0316331461155757600080fd5b600160a060020a038116151561156c57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b600254600090600160a060020a031633146115a857600080fd5b5080600160a060020a03811615156115c85750600254600160a060020a03165b601154611388116115d857600080fd5b6011805460010190556115ef6000808086856129b0565b50505050565b600254600160a060020a03163314806116185750600054600160a060020a031633145b8061162d5750600154600160a060020a031633145b151561163857600080fd5b60035463ffffffff16811061164c57600080fd5b600555565b60025460a060020a900460ff1681565b6001546000908190600160a060020a0316331461167d57600080fd5b5050600e54600f54303191600190910102808211156116c157600154604051600160a060020a039091169082840380156108fc02916000818181858888f150505050505b5050565b600081815260076020526040902054600160a060020a0316801515610cda57600080fd5b61afc881565b601354600160a060020a031681565b60008054600160a060020a0316331461171657600080fd5b81905080600160a060020a03166385b861886040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561175757600080fd5b505af115801561176b573d6000803e3d6000fd5b505050506040513d602081101561178157600080fd5b5051151561178e57600080fd5b600b8054600160a060020a031916600160a060020a039290921691909117905550565b600160a060020a031660009081526008602052604090205490565b600054600160a060020a031633146117e357600080fd5b60025460a060020a900460ff1615156117fb57600080fd5b60138054600160a060020a038316600160a060020a0319909116811790915560408051918252517f450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa4461993059181900360200190a150565b60055481565b600254600160a060020a03163314806118785750600054600160a060020a031633145b8061188d5750600154600160a060020a031633145b151561189857600080fd5b60025460a060020a900460ff16156118af57600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a179055565b60606000606060008060006118e9876117b1565b945084151561190857604080516000815260208101909152955061199d565b84604051908082528060200260200182016040528015611932578160200160208202803883390190505b50935061193d610f7d565b925060009150600190505b82811161199957600081815260076020526040902054600160a060020a03888116911614156119915780848381518110151561198057fe5b602090810290910101526001909101905b600101611948565b8395505b5050505050919050565b600080600080600080600080600260149054906101000a900460ff161515156119cf57600080fd5b600680548a9081106119dd57fe5b60009182526020909120600290910201600181015490975067ffffffffffffffff161515611a0a57600080fd5b604080516101008101825288548152600189015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e0820152611a9b90612c8d565b1515611aa657600080fd5b60018701546006805460c060020a90920463ffffffff1697509087908110611aca57fe5b600091825260209091206001808a015460029093029091019081015490965061ffff60f060020a928390048116965091900416841015611b1757600185015460f060020a900461ffff1693505b6010548754865460018a0154604080517f0d9f5aed0000000000000000000000000000000000000000000000000000000081526004810194909452602484019290925260001967ffffffffffffffff6801000000000000000090920482160116604483015251600160a060020a0390921691630d9f5aed916064808201926020929091908290030181600087803b158015611bb157600080fd5b505af1158015611bc5573d6000803e3d6000fd5b505050506040513d6020811015611bdb57600080fd5b505160008a815260076020526040902054600189810154929550600160a060020a039091169350611c24918b9160c060020a90910463ffffffff1690870161ffff1686866129b0565b6001880180547bffffffff00000000000000000000000000000000000000000000000019169055600f8054600019019055600e54604051919250339181156108fc0291906000818181858888f150939c9b505050505050505050505050565b600254600160a060020a0316331480611ca65750600054600160a060020a031633145b80611cbb5750600154600160a060020a031633145b1515611cc657600080fd5b600b60009054906101000a9004600160a060020a0316600160a060020a0316635fd8c7106040518163ffffffff1660e060020a028152600401600060405180830381600087803b158015611d1957600080fd5b505af1158015611d2d573d6000803e3d6000fd5b50505050600c60009054906101000a9004600160a060020a0316600160a060020a0316635fd8c7106040518163ffffffff1660e060020a028152600401600060405180830381600087803b158015611d8457600080fd5b505af11580156115ef573d6000803e3d6000fd5b60408051808201909152600281527f434b000000000000000000000000000000000000000000000000000000000000602082015281565b600381600e8110611ddc57fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b600760205260009081526040902054600160a060020a031681565b60025460a060020a900460ff1615611e2e57600080fd5b600160a060020a0382161515611e4357600080fd5b600160a060020a038216301415611e5957600080fd5b600b54600160a060020a0383811691161415611e7457600080fd5b600c54600160a060020a0383811691161415611e8f57600080fd5b611e99338261263c565b1515611ea457600080fd5b6116c13383836126aa565b600254600160a060020a031681565b600e5481565b600d54600160a060020a031681565b600254600090600160a060020a03163314611eed57600080fd5b60125461afc811611efd57600080fd5b611f0c600080600085306129b0565b600b54909150611f26908290600160a060020a031661265c565b600b54600160a060020a03166327ebe40a82611f40612cbd565b6040805160e060020a63ffffffff861602815260048101939093526024830191909152600060448301819052620151806064840152306084840152905160a48084019382900301818387803b158015611f9857600080fd5b505af1158015611fac573d6000803e3d6000fd5b505060128054600101905550505050565b600080808311611fcc57600080fd5b6006805484908110611fda57fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e082015290915061208590612d87565b9392505050565b61138881565b600054600160a060020a031633146120a957600080fd5b600d8054600160a060020a031916600160a060020a0392909216919091179055565b600b54600160a060020a031681565b600080600080600080600080600080600060068c8154811015156120fa57fe5b906000526020600020906002020190508060010160189054906101000a900463ffffffff1663ffffffff16600014159a50438160010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff161115995080600101601c9054906101000a900461ffff1661ffff1698508060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1697508060010160189054906101000a900463ffffffff1663ffffffff1696508060010160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1695508060010160109054906101000a900463ffffffff1663ffffffff1694508060010160149054906101000a900463ffffffff1663ffffffff16935080600101601e9054906101000a900461ffff1661ffff16925080600001549150509193959799509193959799565b60025460009060a060020a900460ff161561225557600080fd5b61225f338361263c565b151561226a57600080fd5b61227382611fbd565b151561227e57600080fd5b6122888284612db6565b151561229357600080fd5b600c54604080517fc55d0f56000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a039092169163c55d0f56916024808201926020929091908290030181600087803b1580156122fa57600080fd5b505af115801561230e573d6000803e3d6000fd5b505050506040513d602081101561232457600080fd5b5051600e54909150810134101561233a57600080fd5b600c54600e54604080517f454a2ab3000000000000000000000000000000000000000000000000000000008152600481018790529051600160a060020a039093169263454a2ab39234039160248082019260009290919082900301818588803b1580156123a657600080fd5b505af11580156123ba573d6000803e3d6000fd5b505050505061105f8263ffffffff168463ffffffff16612e05565b60125481565b601054600160a060020a031681565b600254600090819060a060020a900460ff161561240657600080fd5b600e5434101561241557600080fd5b61241f338561263c565b151561242a57600080fd5b612434838561295b565b151561243f57600080fd5b600680548590811061244d57fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e08201529092506124f890612d87565b151561250357600080fd5b600680548490811061251157fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e08201529091506125bc90612d87565b15156125c757600080fd5b6125d3828583866127db565b15156125de57600080fd5b6115ef8484612e05565b606080600080846040519080825280601f01601f19166020018201604052801561261c578160200160208202803883390190505b5092505060208201905084612632828287612f43565b5090949350505050565b600090815260076020526040902054600160a060020a0391821691161490565b6000918252600960205260409091208054600160a060020a031916600160a060020a03909216919091179055565b600090815260096020526040902054600160a060020a0391821691161490565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a031916909117905583161561273d57600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b60408051600160a060020a0380861682528416602082015280820183905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a1505050565b600054600160a060020a031633146127a357600080fd5b60025460a060020a900460ff1615156127bb57600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b6000818414156127ed57506000612953565b6001850154608060020a900463ffffffff1682148061281c5750600185015460a060020a900463ffffffff1682145b1561282957506000612953565b6001830154608060020a900463ffffffff168414806128585750600183015460a060020a900463ffffffff1684145b1561286557506000612953565b6001830154608060020a900463ffffffff16158061289257506001850154608060020a900463ffffffff16155b1561289f57506001612953565b60018581015490840154608060020a9182900463ffffffff908116929091041614806128ea575060018086015490840154608060020a900463ffffffff90811660a060020a90920416145b156128f757506000612953565b6001808601549084015460a060020a900463ffffffff908116608060020a90920416148061294257506001858101549084015460a060020a9182900463ffffffff9081169290910416145b1561294f57506000612953565b5060015b949350505050565b6000818152600760205260408082205484835290822054600160a060020a039182169116808214806113ce57506000858152600a6020526040902054600160a060020a03908116908316149250505092915050565b6000806129bb61308f565b600063ffffffff891689146129cf57600080fd5b63ffffffff881688146129e157600080fd5b61ffff871687146129f157600080fd5b600287049250600d8361ffff161115612a0957600d92505b505060408051610100810182528581524267ffffffffffffffff90811660208301908152600093830184815263ffffffff8c8116606086019081528c82166080870190815260a0870188815261ffff8a811660c08a019081528f821660e08b01908152600680546001810182559c528a517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f60028e029081019190915598517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d40909901805498519651955194519251915167ffffffffffffffff19909916998b16999099176fffffffffffffffff000000000000000019166801000000000000000096909a16959095029890981773ffffffff000000000000000000000000000000001916608060020a938616939093029290921777ffffffff0000000000000000000000000000000000000000191660a060020a91851691909102177bffffffff000000000000000000000000000000000000000000000000191660c060020a96841696909602959095177fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e060020a91861691909102177dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f060020a929094169190910292909217905590919081168114612c0757600080fd5b606080830151608080850151855160408051600160a060020a038c1681526020810188905263ffffffff95861681830152929094169482019490945290810192909252517f0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad59181900360a00190a1612c81600086836126aa565b98975050505050505050565b60008160a0015163ffffffff16600014158015610cd75750506040015167ffffffffffffffff4381169116111590565b6000806000600b60009054906101000a9004600160a060020a0316600160a060020a031663eac9d94c6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612d1557600080fd5b505af1158015612d29573d6000803e3d6000fd5b505050506040513d6020811015612d3f57600080fd5b505191506fffffffffffffffffffffffffffffffff82168214612d6157600080fd5b50600281048101662386f26fc10000811015612d815750662386f26fc100005b92915050565b60008160a0015163ffffffff166000148015610cd75750506040015167ffffffffffffffff4381169116111590565b6000806000600685815481101515612dca57fe5b90600052602060002090600202019150600684815481101515612de957fe5b906000526020600020906002020190506113ce828683876127db565b600080600683815481101515612e1757fe5b90600052602060002090600202019150600684815481101515612e3657fe5b600091825260209091206002909102016001810180547bffffffff000000000000000000000000000000000000000000000000191660c060020a63ffffffff8716021790559050612e8682612f87565b612e8f81612f87565b6000848152600a602090815260408083208054600160a060020a031990811690915586845281842080549091169055600f80546001908101909155878452600783529281902054928401548151600160a060020a0390941684529183018790528281018690526801000000000000000090910467ffffffffffffffff166060830152517f241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80916080908290030190a150505050565b60005b60208210612f68578251845260209384019390920191601f1990910190612f46565b50905182516020929092036101000a6000190180199091169116179052565b600554600182015443919060039060e060020a900461ffff16600e8110612faa57fe5b600891828204019190066004029054906101000a900463ffffffff1663ffffffff16811515612fd557fe5b6001840180546fffffffffffffffff0000000000000000191668010000000000000000939092049390930167ffffffffffffffff16919091021790819055600d60e060020a90910461ffff16101561306d576001818101805461ffff60e060020a8083048216909401169092027fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790555b50565b6080604051908101604052806004906020820280388339509192915050565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810191909152905600a165627a7a72305820c0025073c2476997a4c60fae1679a6fe680980aab39f45bf9f22de24cea376ae0029`

// DeployKittyCore deploys a new Ethereum contract, binding an instance of KittyCore to it.
func DeployKittyCore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyCore, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyCoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyCoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyCore{KittyCoreCaller: KittyCoreCaller{contract: contract}, KittyCoreTransactor: KittyCoreTransactor{contract: contract}, KittyCoreFilterer: KittyCoreFilterer{contract: contract}}, nil
}

// KittyCore is an auto generated Go binding around an Ethereum contract.
type KittyCore struct {
	KittyCoreCaller     // Read-only binding to the contract
	KittyCoreTransactor // Write-only binding to the contract
	KittyCoreFilterer   // Log filterer for contract events
}

// KittyCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyCoreSession struct {
	Contract     *KittyCore        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittyCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyCoreCallerSession struct {
	Contract *KittyCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KittyCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyCoreTransactorSession struct {
	Contract     *KittyCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KittyCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyCoreRaw struct {
	Contract *KittyCore // Generic contract binding to access the raw methods on
}

// KittyCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyCoreCallerRaw struct {
	Contract *KittyCoreCaller // Generic read-only contract binding to access the raw methods on
}

// KittyCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyCoreTransactorRaw struct {
	Contract *KittyCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyCore creates a new instance of KittyCore, bound to a specific deployed contract.
func NewKittyCore(address common.Address, backend bind.ContractBackend) (*KittyCore, error) {
	contract, err := bindKittyCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyCore{KittyCoreCaller: KittyCoreCaller{contract: contract}, KittyCoreTransactor: KittyCoreTransactor{contract: contract}, KittyCoreFilterer: KittyCoreFilterer{contract: contract}}, nil
}

// NewKittyCoreCaller creates a new read-only instance of KittyCore, bound to a specific deployed contract.
func NewKittyCoreCaller(address common.Address, caller bind.ContractCaller) (*KittyCoreCaller, error) {
	contract, err := bindKittyCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyCoreCaller{contract: contract}, nil
}

// NewKittyCoreTransactor creates a new write-only instance of KittyCore, bound to a specific deployed contract.
func NewKittyCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyCoreTransactor, error) {
	contract, err := bindKittyCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyCoreTransactor{contract: contract}, nil
}

// NewKittyCoreFilterer creates a new log filterer instance of KittyCore, bound to a specific deployed contract.
func NewKittyCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyCoreFilterer, error) {
	contract, err := bindKittyCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyCoreFilterer{contract: contract}, nil
}

// bindKittyCore binds a generic wrapper to an already deployed contract.
func bindKittyCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyCoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyCore *KittyCoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyCore.Contract.KittyCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyCore *KittyCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyCore.Contract.KittyCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyCore *KittyCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyCore.Contract.KittyCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyCore *KittyCoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyCore *KittyCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyCore *KittyCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyCore.Contract.contract.Transact(opts, method, params...)
}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() constant returns(uint256)
func (_KittyCore *KittyCoreCaller) GEN0AUCTIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "GEN0_AUCTION_DURATION")
	return *ret0, err
}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() constant returns(uint256)
func (_KittyCore *KittyCoreSession) GEN0AUCTIONDURATION() (*big.Int, error) {
	return _KittyCore.Contract.GEN0AUCTIONDURATION(&_KittyCore.CallOpts)
}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() constant returns(uint256)
func (_KittyCore *KittyCoreCallerSession) GEN0AUCTIONDURATION() (*big.Int, error) {
	return _KittyCore.Contract.GEN0AUCTIONDURATION(&_KittyCore.CallOpts)
}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() constant returns(uint256)
func (_KittyCore *KittyCoreCaller) GEN0CREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "GEN0_CREATION_LIMIT")
	return *ret0, err
}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() constant returns(uint256)
func (_KittyCore *KittyCoreSession) GEN0CREATIONLIMIT() (*big.Int, error) {
	return _KittyCore.Contract.GEN0CREATIONLIMIT(&_KittyCore.CallOpts)
}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() constant returns(uint256)
func (_KittyCore *KittyCoreCallerSession) GEN0CREATIONLIMIT() (*big.Int, error) {
	return _KittyCore.Contract.GEN0CREATIONLIMIT(&_KittyCore.CallOpts)
}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() constant returns(uint256)
func (_KittyCore *KittyCoreCaller) GEN0STARTINGPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "GEN0_STARTING_PRICE")
	return *ret0, err
}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() constant returns(uint256)
func (_KittyCore *KittyCoreSession) GEN0STARTINGPRICE() (*big.Int, error) {
	return _KittyCore.Contract.GEN0STARTINGPRICE(&_KittyCore.CallOpts)
}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() constant returns(uint256)
func (_KittyCore *KittyCoreCallerSession) GEN0STARTINGPRICE() (*big.Int, error) {
	return _KittyCore.Contract.GEN0STARTINGPRICE(&_KittyCore.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_KittyCore *KittyCoreCaller) PROMOCREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "PROMO_CREATION_LIMIT")
	return *ret0, err
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_KittyCore *KittyCoreSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _KittyCore.Contract.PROMOCREATIONLIMIT(&_KittyCore.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_KittyCore *KittyCoreCallerSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _KittyCore.Contract.PROMOCREATIONLIMIT(&_KittyCore.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_KittyCore *KittyCoreCaller) AutoBirthFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "autoBirthFee")
	return *ret0, err
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_KittyCore *KittyCoreSession) AutoBirthFee() (*big.Int, error) {
	return _KittyCore.Contract.AutoBirthFee(&_KittyCore.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_KittyCore *KittyCoreCallerSession) AutoBirthFee() (*big.Int, error) {
	return _KittyCore.Contract.AutoBirthFee(&_KittyCore.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyCore *KittyCoreCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyCore *KittyCoreSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyCore.Contract.BalanceOf(&_KittyCore.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyCore *KittyCoreCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyCore.Contract.BalanceOf(&_KittyCore.CallOpts, _owner)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_KittyCore *KittyCoreCaller) CanBreedWith(opts *bind.CallOpts, _matronId *big.Int, _sireId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "canBreedWith", _matronId, _sireId)
	return *ret0, err
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_KittyCore *KittyCoreSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyCore.Contract.CanBreedWith(&_KittyCore.CallOpts, _matronId, _sireId)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_KittyCore *KittyCoreCallerSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyCore.Contract.CanBreedWith(&_KittyCore.CallOpts, _matronId, _sireId)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyCore *KittyCoreCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyCore *KittyCoreSession) CeoAddress() (common.Address, error) {
	return _KittyCore.Contract.CeoAddress(&_KittyCore.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyCore *KittyCoreCallerSession) CeoAddress() (common.Address, error) {
	return _KittyCore.Contract.CeoAddress(&_KittyCore.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyCore *KittyCoreCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyCore *KittyCoreSession) CfoAddress() (common.Address, error) {
	return _KittyCore.Contract.CfoAddress(&_KittyCore.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyCore *KittyCoreCallerSession) CfoAddress() (common.Address, error) {
	return _KittyCore.Contract.CfoAddress(&_KittyCore.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyCore *KittyCoreCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyCore *KittyCoreSession) CooAddress() (common.Address, error) {
	return _KittyCore.Contract.CooAddress(&_KittyCore.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyCore *KittyCoreCallerSession) CooAddress() (common.Address, error) {
	return _KittyCore.Contract.CooAddress(&_KittyCore.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyCore *KittyCoreCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "cooldowns", arg0)
	return *ret0, err
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyCore *KittyCoreSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyCore.Contract.Cooldowns(&_KittyCore.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyCore *KittyCoreCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyCore.Contract.Cooldowns(&_KittyCore.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyCore *KittyCoreCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "erc721Metadata")
	return *ret0, err
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyCore *KittyCoreSession) Erc721Metadata() (common.Address, error) {
	return _KittyCore.Contract.Erc721Metadata(&_KittyCore.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyCore *KittyCoreCallerSession) Erc721Metadata() (common.Address, error) {
	return _KittyCore.Contract.Erc721Metadata(&_KittyCore.CallOpts)
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() constant returns(uint256)
func (_KittyCore *KittyCoreCaller) Gen0CreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "gen0CreatedCount")
	return *ret0, err
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() constant returns(uint256)
func (_KittyCore *KittyCoreSession) Gen0CreatedCount() (*big.Int, error) {
	return _KittyCore.Contract.Gen0CreatedCount(&_KittyCore.CallOpts)
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() constant returns(uint256)
func (_KittyCore *KittyCoreCallerSession) Gen0CreatedCount() (*big.Int, error) {
	return _KittyCore.Contract.Gen0CreatedCount(&_KittyCore.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_KittyCore *KittyCoreCaller) GeneScience(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "geneScience")
	return *ret0, err
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_KittyCore *KittyCoreSession) GeneScience() (common.Address, error) {
	return _KittyCore.Contract.GeneScience(&_KittyCore.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_KittyCore *KittyCoreCallerSession) GeneScience() (common.Address, error) {
	return _KittyCore.Contract.GeneScience(&_KittyCore.CallOpts)
}

// GetKitty is a free data retrieval call binding the contract method 0xe98b7f4d.
//
// Solidity: function getKitty(_id uint256) constant returns(isGestating bool, isReady bool, cooldownIndex uint256, nextActionAt uint256, siringWithId uint256, birthTime uint256, matronId uint256, sireId uint256, generation uint256, genes uint256)
func (_KittyCore *KittyCoreCaller) GetKitty(opts *bind.CallOpts, _id *big.Int) (struct {
	IsGestating   bool
	IsReady       bool
	CooldownIndex *big.Int
	NextActionAt  *big.Int
	SiringWithId  *big.Int
	BirthTime     *big.Int
	MatronId      *big.Int
	SireId        *big.Int
	Generation    *big.Int
	Genes         *big.Int
}, error) {
	ret := new(struct {
		IsGestating   bool
		IsReady       bool
		CooldownIndex *big.Int
		NextActionAt  *big.Int
		SiringWithId  *big.Int
		BirthTime     *big.Int
		MatronId      *big.Int
		SireId        *big.Int
		Generation    *big.Int
		Genes         *big.Int
	})
	out := ret
	err := _KittyCore.contract.Call(opts, out, "getKitty", _id)
	return *ret, err
}

// GetKitty is a free data retrieval call binding the contract method 0xe98b7f4d.
//
// Solidity: function getKitty(_id uint256) constant returns(isGestating bool, isReady bool, cooldownIndex uint256, nextActionAt uint256, siringWithId uint256, birthTime uint256, matronId uint256, sireId uint256, generation uint256, genes uint256)
func (_KittyCore *KittyCoreSession) GetKitty(_id *big.Int) (struct {
	IsGestating   bool
	IsReady       bool
	CooldownIndex *big.Int
	NextActionAt  *big.Int
	SiringWithId  *big.Int
	BirthTime     *big.Int
	MatronId      *big.Int
	SireId        *big.Int
	Generation    *big.Int
	Genes         *big.Int
}, error) {
	return _KittyCore.Contract.GetKitty(&_KittyCore.CallOpts, _id)
}

// GetKitty is a free data retrieval call binding the contract method 0xe98b7f4d.
//
// Solidity: function getKitty(_id uint256) constant returns(isGestating bool, isReady bool, cooldownIndex uint256, nextActionAt uint256, siringWithId uint256, birthTime uint256, matronId uint256, sireId uint256, generation uint256, genes uint256)
func (_KittyCore *KittyCoreCallerSession) GetKitty(_id *big.Int) (struct {
	IsGestating   bool
	IsReady       bool
	CooldownIndex *big.Int
	NextActionAt  *big.Int
	SiringWithId  *big.Int
	BirthTime     *big.Int
	MatronId      *big.Int
	SireId        *big.Int
	Generation    *big.Int
	Genes         *big.Int
}, error) {
	return _KittyCore.Contract.GetKitty(&_KittyCore.CallOpts, _id)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_KittyCore *KittyCoreCaller) IsPregnant(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "isPregnant", _kittyId)
	return *ret0, err
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_KittyCore *KittyCoreSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyCore.Contract.IsPregnant(&_KittyCore.CallOpts, _kittyId)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_KittyCore *KittyCoreCallerSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyCore.Contract.IsPregnant(&_KittyCore.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_KittyCore *KittyCoreCaller) IsReadyToBreed(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "isReadyToBreed", _kittyId)
	return *ret0, err
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_KittyCore *KittyCoreSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyCore.Contract.IsReadyToBreed(&_KittyCore.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_KittyCore *KittyCoreCallerSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyCore.Contract.IsReadyToBreed(&_KittyCore.CallOpts, _kittyId)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyCore *KittyCoreCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "kittyIndexToApproved", arg0)
	return *ret0, err
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyCore *KittyCoreSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyCore.Contract.KittyIndexToApproved(&_KittyCore.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyCore *KittyCoreCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyCore.Contract.KittyIndexToApproved(&_KittyCore.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyCore *KittyCoreCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "kittyIndexToOwner", arg0)
	return *ret0, err
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyCore *KittyCoreSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyCore.Contract.KittyIndexToOwner(&_KittyCore.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyCore *KittyCoreCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyCore.Contract.KittyIndexToOwner(&_KittyCore.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyCore *KittyCoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyCore *KittyCoreSession) Name() (string, error) {
	return _KittyCore.Contract.Name(&_KittyCore.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyCore *KittyCoreCallerSession) Name() (string, error) {
	return _KittyCore.Contract.Name(&_KittyCore.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_KittyCore *KittyCoreCaller) NewContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "newContractAddress")
	return *ret0, err
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_KittyCore *KittyCoreSession) NewContractAddress() (common.Address, error) {
	return _KittyCore.Contract.NewContractAddress(&_KittyCore.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_KittyCore *KittyCoreCallerSession) NewContractAddress() (common.Address, error) {
	return _KittyCore.Contract.NewContractAddress(&_KittyCore.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyCore *KittyCoreCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyCore *KittyCoreSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyCore.Contract.OwnerOf(&_KittyCore.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyCore *KittyCoreCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyCore.Contract.OwnerOf(&_KittyCore.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyCore *KittyCoreCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyCore *KittyCoreSession) Paused() (bool, error) {
	return _KittyCore.Contract.Paused(&_KittyCore.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyCore *KittyCoreCallerSession) Paused() (bool, error) {
	return _KittyCore.Contract.Paused(&_KittyCore.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_KittyCore *KittyCoreCaller) PregnantKitties(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "pregnantKitties")
	return *ret0, err
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_KittyCore *KittyCoreSession) PregnantKitties() (*big.Int, error) {
	return _KittyCore.Contract.PregnantKitties(&_KittyCore.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_KittyCore *KittyCoreCallerSession) PregnantKitties() (*big.Int, error) {
	return _KittyCore.Contract.PregnantKitties(&_KittyCore.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_KittyCore *KittyCoreCaller) PromoCreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "promoCreatedCount")
	return *ret0, err
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_KittyCore *KittyCoreSession) PromoCreatedCount() (*big.Int, error) {
	return _KittyCore.Contract.PromoCreatedCount(&_KittyCore.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_KittyCore *KittyCoreCallerSession) PromoCreatedCount() (*big.Int, error) {
	return _KittyCore.Contract.PromoCreatedCount(&_KittyCore.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyCore *KittyCoreCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "saleAuction")
	return *ret0, err
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyCore *KittyCoreSession) SaleAuction() (common.Address, error) {
	return _KittyCore.Contract.SaleAuction(&_KittyCore.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyCore *KittyCoreCallerSession) SaleAuction() (common.Address, error) {
	return _KittyCore.Contract.SaleAuction(&_KittyCore.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyCore *KittyCoreCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "secondsPerBlock")
	return *ret0, err
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyCore *KittyCoreSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyCore.Contract.SecondsPerBlock(&_KittyCore.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyCore *KittyCoreCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyCore.Contract.SecondsPerBlock(&_KittyCore.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyCore *KittyCoreCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "sireAllowedToAddress", arg0)
	return *ret0, err
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyCore *KittyCoreSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyCore.Contract.SireAllowedToAddress(&_KittyCore.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyCore *KittyCoreCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyCore.Contract.SireAllowedToAddress(&_KittyCore.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyCore *KittyCoreCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "siringAuction")
	return *ret0, err
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyCore *KittyCoreSession) SiringAuction() (common.Address, error) {
	return _KittyCore.Contract.SiringAuction(&_KittyCore.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyCore *KittyCoreCallerSession) SiringAuction() (common.Address, error) {
	return _KittyCore.Contract.SiringAuction(&_KittyCore.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyCore *KittyCoreCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "supportsInterface", _interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyCore *KittyCoreSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyCore.Contract.SupportsInterface(&_KittyCore.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyCore *KittyCoreCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyCore.Contract.SupportsInterface(&_KittyCore.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyCore *KittyCoreCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyCore *KittyCoreSession) Symbol() (string, error) {
	return _KittyCore.Contract.Symbol(&_KittyCore.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyCore *KittyCoreCallerSession) Symbol() (string, error) {
	return _KittyCore.Contract.Symbol(&_KittyCore.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyCore *KittyCoreCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "tokenMetadata", _tokenId, _preferredTransport)
	return *ret0, err
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyCore *KittyCoreSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyCore.Contract.TokenMetadata(&_KittyCore.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyCore *KittyCoreCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyCore.Contract.TokenMetadata(&_KittyCore.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyCore *KittyCoreCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyCore *KittyCoreSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyCore.Contract.TokensOfOwner(&_KittyCore.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyCore *KittyCoreCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyCore.Contract.TokensOfOwner(&_KittyCore.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyCore *KittyCoreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyCore.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyCore *KittyCoreSession) TotalSupply() (*big.Int, error) {
	return _KittyCore.Contract.TotalSupply(&_KittyCore.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyCore *KittyCoreCallerSession) TotalSupply() (*big.Int, error) {
	return _KittyCore.Contract.TotalSupply(&_KittyCore.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyCore *KittyCoreTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyCore *KittyCoreSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.Approve(&_KittyCore.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyCore *KittyCoreTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.Approve(&_KittyCore.TransactOpts, _to, _tokenId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_KittyCore *KittyCoreTransactor) ApproveSiring(opts *bind.TransactOpts, _addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "approveSiring", _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_KittyCore *KittyCoreSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.ApproveSiring(&_KittyCore.TransactOpts, _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_KittyCore *KittyCoreTransactorSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.ApproveSiring(&_KittyCore.TransactOpts, _addr, _sireId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(_sireId uint256, _matronId uint256) returns()
func (_KittyCore *KittyCoreTransactor) BidOnSiringAuction(opts *bind.TransactOpts, _sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "bidOnSiringAuction", _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(_sireId uint256, _matronId uint256) returns()
func (_KittyCore *KittyCoreSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.BidOnSiringAuction(&_KittyCore.TransactOpts, _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(_sireId uint256, _matronId uint256) returns()
func (_KittyCore *KittyCoreTransactorSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.BidOnSiringAuction(&_KittyCore.TransactOpts, _sireId, _matronId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_KittyCore *KittyCoreTransactor) BreedWithAuto(opts *bind.TransactOpts, _matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "breedWithAuto", _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_KittyCore *KittyCoreSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.BreedWithAuto(&_KittyCore.TransactOpts, _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_KittyCore *KittyCoreTransactorSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.BreedWithAuto(&_KittyCore.TransactOpts, _matronId, _sireId)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(_genes uint256) returns()
func (_KittyCore *KittyCoreTransactor) CreateGen0Auction(opts *bind.TransactOpts, _genes *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "createGen0Auction", _genes)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(_genes uint256) returns()
func (_KittyCore *KittyCoreSession) CreateGen0Auction(_genes *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.CreateGen0Auction(&_KittyCore.TransactOpts, _genes)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(_genes uint256) returns()
func (_KittyCore *KittyCoreTransactorSession) CreateGen0Auction(_genes *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.CreateGen0Auction(&_KittyCore.TransactOpts, _genes)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(_genes uint256, _owner address) returns()
func (_KittyCore *KittyCoreTransactor) CreatePromoKitty(opts *bind.TransactOpts, _genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "createPromoKitty", _genes, _owner)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(_genes uint256, _owner address) returns()
func (_KittyCore *KittyCoreSession) CreatePromoKitty(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.CreatePromoKitty(&_KittyCore.TransactOpts, _genes, _owner)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(_genes uint256, _owner address) returns()
func (_KittyCore *KittyCoreTransactorSession) CreatePromoKitty(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.CreatePromoKitty(&_KittyCore.TransactOpts, _genes, _owner)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyCore *KittyCoreTransactor) CreateSaleAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "createSaleAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyCore *KittyCoreSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.CreateSaleAuction(&_KittyCore.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyCore *KittyCoreTransactorSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.CreateSaleAuction(&_KittyCore.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyCore *KittyCoreTransactor) CreateSiringAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "createSiringAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyCore *KittyCoreSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.CreateSiringAuction(&_KittyCore.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyCore *KittyCoreTransactorSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.CreateSiringAuction(&_KittyCore.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_KittyCore *KittyCoreTransactor) GiveBirth(opts *bind.TransactOpts, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "giveBirth", _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_KittyCore *KittyCoreSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.GiveBirth(&_KittyCore.TransactOpts, _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_KittyCore *KittyCoreTransactorSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.GiveBirth(&_KittyCore.TransactOpts, _matronId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyCore *KittyCoreTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyCore *KittyCoreSession) Pause() (*types.Transaction, error) {
	return _KittyCore.Contract.Pause(&_KittyCore.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyCore *KittyCoreTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyCore.Contract.Pause(&_KittyCore.TransactOpts)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_KittyCore *KittyCoreTransactor) SetAutoBirthFee(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setAutoBirthFee", val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_KittyCore *KittyCoreSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.SetAutoBirthFee(&_KittyCore.TransactOpts, val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_KittyCore *KittyCoreTransactorSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.SetAutoBirthFee(&_KittyCore.TransactOpts, val)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyCore *KittyCoreTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyCore *KittyCoreSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetCEO(&_KittyCore.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyCore *KittyCoreTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetCEO(&_KittyCore.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyCore *KittyCoreTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyCore *KittyCoreSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetCFO(&_KittyCore.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyCore *KittyCoreTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetCFO(&_KittyCore.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyCore *KittyCoreTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyCore *KittyCoreSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetCOO(&_KittyCore.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyCore *KittyCoreTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetCOO(&_KittyCore.TransactOpts, _newCOO)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_KittyCore *KittyCoreTransactor) SetGeneScienceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setGeneScienceAddress", _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_KittyCore *KittyCoreSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetGeneScienceAddress(&_KittyCore.TransactOpts, _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_KittyCore *KittyCoreTransactorSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetGeneScienceAddress(&_KittyCore.TransactOpts, _address)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyCore *KittyCoreTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyCore *KittyCoreSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetMetadataAddress(&_KittyCore.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyCore *KittyCoreTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetMetadataAddress(&_KittyCore.TransactOpts, _contractAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_v2Address address) returns()
func (_KittyCore *KittyCoreTransactor) SetNewAddress(opts *bind.TransactOpts, _v2Address common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setNewAddress", _v2Address)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_v2Address address) returns()
func (_KittyCore *KittyCoreSession) SetNewAddress(_v2Address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetNewAddress(&_KittyCore.TransactOpts, _v2Address)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_v2Address address) returns()
func (_KittyCore *KittyCoreTransactorSession) SetNewAddress(_v2Address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetNewAddress(&_KittyCore.TransactOpts, _v2Address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(_address address) returns()
func (_KittyCore *KittyCoreTransactor) SetSaleAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setSaleAuctionAddress", _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(_address address) returns()
func (_KittyCore *KittyCoreSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetSaleAuctionAddress(&_KittyCore.TransactOpts, _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(_address address) returns()
func (_KittyCore *KittyCoreTransactorSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetSaleAuctionAddress(&_KittyCore.TransactOpts, _address)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyCore *KittyCoreTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyCore *KittyCoreSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.SetSecondsPerBlock(&_KittyCore.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyCore *KittyCoreTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.SetSecondsPerBlock(&_KittyCore.TransactOpts, secs)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(_address address) returns()
func (_KittyCore *KittyCoreTransactor) SetSiringAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setSiringAuctionAddress", _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(_address address) returns()
func (_KittyCore *KittyCoreSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetSiringAuctionAddress(&_KittyCore.TransactOpts, _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(_address address) returns()
func (_KittyCore *KittyCoreTransactorSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetSiringAuctionAddress(&_KittyCore.TransactOpts, _address)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyCore *KittyCoreTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyCore *KittyCoreSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.Transfer(&_KittyCore.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyCore *KittyCoreTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.Transfer(&_KittyCore.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyCore *KittyCoreTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyCore *KittyCoreSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.TransferFrom(&_KittyCore.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyCore *KittyCoreTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.TransferFrom(&_KittyCore.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyCore *KittyCoreTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyCore *KittyCoreSession) Unpause() (*types.Transaction, error) {
	return _KittyCore.Contract.Unpause(&_KittyCore.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyCore *KittyCoreTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyCore.Contract.Unpause(&_KittyCore.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyCore *KittyCoreTransactor) WithdrawAuctionBalances(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "withdrawAuctionBalances")
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyCore *KittyCoreSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _KittyCore.Contract.WithdrawAuctionBalances(&_KittyCore.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyCore *KittyCoreTransactorSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _KittyCore.Contract.WithdrawAuctionBalances(&_KittyCore.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_KittyCore *KittyCoreTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_KittyCore *KittyCoreSession) WithdrawBalance() (*types.Transaction, error) {
	return _KittyCore.Contract.WithdrawBalance(&_KittyCore.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_KittyCore *KittyCoreTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _KittyCore.Contract.WithdrawBalance(&_KittyCore.TransactOpts)
}

// KittyCoreApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KittyCore contract.
type KittyCoreApprovalIterator struct {
	Event *KittyCoreApproval // Event containing the contract specifics and raw log

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
func (it *KittyCoreApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyCoreApproval)
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
		it.Event = new(KittyCoreApproval)
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
func (it *KittyCoreApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyCoreApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyCoreApproval represents a Approval event raised by the KittyCore contract.
type KittyCoreApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_KittyCore *KittyCoreFilterer) FilterApproval(opts *bind.FilterOpts) (*KittyCoreApprovalIterator, error) {

	logs, sub, err := _KittyCore.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &KittyCoreApprovalIterator{contract: _KittyCore.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_KittyCore *KittyCoreFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KittyCoreApproval) (event.Subscription, error) {

	logs, sub, err := _KittyCore.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyCoreApproval)
				if err := _KittyCore.contract.UnpackLog(event, "Approval", log); err != nil {
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

// KittyCoreBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the KittyCore contract.
type KittyCoreBirthIterator struct {
	Event *KittyCoreBirth // Event containing the contract specifics and raw log

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
func (it *KittyCoreBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyCoreBirth)
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
		it.Event = new(KittyCoreBirth)
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
func (it *KittyCoreBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyCoreBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyCoreBirth represents a Birth event raised by the KittyCore contract.
type KittyCoreBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_KittyCore *KittyCoreFilterer) FilterBirth(opts *bind.FilterOpts) (*KittyCoreBirthIterator, error) {

	logs, sub, err := _KittyCore.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &KittyCoreBirthIterator{contract: _KittyCore.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_KittyCore *KittyCoreFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *KittyCoreBirth) (event.Subscription, error) {

	logs, sub, err := _KittyCore.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyCoreBirth)
				if err := _KittyCore.contract.UnpackLog(event, "Birth", log); err != nil {
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

// KittyCoreContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyCore contract.
type KittyCoreContractUpgradeIterator struct {
	Event *KittyCoreContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyCoreContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyCoreContractUpgrade)
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
		it.Event = new(KittyCoreContractUpgrade)
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
func (it *KittyCoreContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyCoreContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyCoreContractUpgrade represents a ContractUpgrade event raised by the KittyCore contract.
type KittyCoreContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyCore *KittyCoreFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyCoreContractUpgradeIterator, error) {

	logs, sub, err := _KittyCore.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyCoreContractUpgradeIterator{contract: _KittyCore.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyCore *KittyCoreFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyCoreContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyCore.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyCoreContractUpgrade)
				if err := _KittyCore.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// KittyCorePregnantIterator is returned from FilterPregnant and is used to iterate over the raw logs and unpacked data for Pregnant events raised by the KittyCore contract.
type KittyCorePregnantIterator struct {
	Event *KittyCorePregnant // Event containing the contract specifics and raw log

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
func (it *KittyCorePregnantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyCorePregnant)
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
		it.Event = new(KittyCorePregnant)
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
func (it *KittyCorePregnantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyCorePregnantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyCorePregnant represents a Pregnant event raised by the KittyCore contract.
type KittyCorePregnant struct {
	Owner            common.Address
	MatronId         *big.Int
	SireId           *big.Int
	CooldownEndBlock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPregnant is a free log retrieval operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: e Pregnant(owner address, matronId uint256, sireId uint256, cooldownEndBlock uint256)
func (_KittyCore *KittyCoreFilterer) FilterPregnant(opts *bind.FilterOpts) (*KittyCorePregnantIterator, error) {

	logs, sub, err := _KittyCore.contract.FilterLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return &KittyCorePregnantIterator{contract: _KittyCore.contract, event: "Pregnant", logs: logs, sub: sub}, nil
}

// WatchPregnant is a free log subscription operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: e Pregnant(owner address, matronId uint256, sireId uint256, cooldownEndBlock uint256)
func (_KittyCore *KittyCoreFilterer) WatchPregnant(opts *bind.WatchOpts, sink chan<- *KittyCorePregnant) (event.Subscription, error) {

	logs, sub, err := _KittyCore.contract.WatchLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyCorePregnant)
				if err := _KittyCore.contract.UnpackLog(event, "Pregnant", log); err != nil {
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

// KittyCoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KittyCore contract.
type KittyCoreTransferIterator struct {
	Event *KittyCoreTransfer // Event containing the contract specifics and raw log

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
func (it *KittyCoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyCoreTransfer)
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
		it.Event = new(KittyCoreTransfer)
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
func (it *KittyCoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyCoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyCoreTransfer represents a Transfer event raised by the KittyCore contract.
type KittyCoreTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_KittyCore *KittyCoreFilterer) FilterTransfer(opts *bind.FilterOpts) (*KittyCoreTransferIterator, error) {

	logs, sub, err := _KittyCore.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &KittyCoreTransferIterator{contract: _KittyCore.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_KittyCore *KittyCoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KittyCoreTransfer) (event.Subscription, error) {

	logs, sub, err := _KittyCore.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyCoreTransfer)
				if err := _KittyCore.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// KittyMintingABI is the input ABI used to generate the binding from.
const KittyMintingABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"promoCreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_STARTING_PRICE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSiringAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pregnantKitties\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isPregnant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_AUCTION_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setGeneScienceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSaleAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"canBreedWith\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSiringAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAutoBirthFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"approveSiring\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genes\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createPromoKitty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSaleAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"giveBirth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAuctionBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoBirthFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genes\",\"type\":\"uint256\"}],\"name\":\"createGen0Auction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isReadyToBreed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PROMO_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sireId\",\"type\":\"uint256\"},{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"bidOnSiringAuction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gen0CreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"geneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"breedWithAuto\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cooldownEndBlock\",\"type\":\"uint256\"}],\"name\":\"Pregnant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyMintingBin is the compiled bytecode used for deploying new contracts.
const KittyMintingBin = `0x6002805460a060020a60ff0219169055610240604052603c6080908152607860a05261012c60c05261025860e05261070861010052610e1061012052611c2061014052613840610160526170806101805261e1006101a052620151806101c0526202a3006101e052620546006102005262093a80610220526200008790600390600e620000ac565b50600f60055566071afd498d0000600e55348015620000a557600080fd5b5062000176565b6002830191839082156200013d5791602002820160005b838211156200010957835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302620000c3565b80156200013b5782816101000a81549063ffffffff021916905560040160208160030104928301926001030262000109565b505b506200014b9291506200014f565b5090565b6200017391905b808211156200014b57805463ffffffff1916815560010162000156565b90565b612d0e80620001866000396000f3006080604052600436106102795763ffffffff60e060020a60003504166301ffc9a7811461027e5780630519ce79146102c95780630560ff44146102fa57806305e455461461039357806306fdde03146103ba578063095ea7b3146103cf5780630a0f8168146103f55780630e583df01461040a57806314001f4c1461041f57806318160ddd14610440578063183a7947146104555780631940a9361461046a57806319c2f2011461048257806321717ebf1461049757806323b872dd146104ac57806324e7a38a146104d657806327d7874c146104f75780632ba73c15146105185780633d7d3f5a146105395780633f4ba83a1461055a57806346116e6f1461056f57806346d22c7014610587578063481af3d3146105a25780634ad8c938146105ba5780634b85fd55146105db5780634dfff04f146105f35780634e0a33791461061757806356129134146106385780635663896e1461065c5780635c975abb146106745780636352211e14610689578063680eba27146106a15780636fbde40d146106b657806370a08231146106d75780637a7d4937146106f85780638456cb591461070d5780638462151c1461072257806388c2a0bf1461079357806391876e57146107ab57806395d89b41146107c05780639d6fac6f146107d5578063a45f4bfc14610806578063a9059cbb1461081e578063b047fb5014610842578063b0c35c0514610857578063bc4006f51461086c578063c3bea9af14610881578063d3e6f49f14610899578063defb9584146108b1578063e17b25af146108c6578063e6cbe351146108e7578063ed60ade6146108fc578063f1ca94101461090a578063f2b47d521461091f578063f7d8c88314610934575b600080fd5b34801561028a57600080fd5b506102b57bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1960043516610942565b604080519115158252519081900360200190f35b3480156102d557600080fd5b506102de610bd5565b60408051600160a060020a039092168252519081900360200190f35b34801561030657600080fd5b5061031e600480359060248035908101910135610be4565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610358578181015183820152602001610340565b50505050905090810190601f1680156103855780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561039f57600080fd5b506103a8610ce7565b60408051918252519081900360200190f35b3480156103c657600080fd5b5061031e610ced565b3480156103db57600080fd5b506103f3600160a060020a0360043516602435610d24565b005b34801561040157600080fd5b506102de610da6565b34801561041657600080fd5b506103a8610db5565b34801561042b57600080fd5b506103f3600160a060020a0360043516610dc0565b34801561044c57600080fd5b506103a8610e73565b34801561046157600080fd5b506103a8610e7d565b34801561047657600080fd5b506102b5600435610e83565b34801561048e57600080fd5b506103a8610ec8565b3480156104a357600080fd5b506102de610ecf565b3480156104b857600080fd5b506103f3600160a060020a0360043581169060243516604435610ede565b3480156104e257600080fd5b506103f3600160a060020a0360043516610f5a565b34801561050357600080fd5b506103f3600160a060020a036004351661100d565b34801561052457600080fd5b506103f3600160a060020a036004351661105b565b34801561054557600080fd5b506103f36004356024356044356064356110a9565b34801561056657600080fd5b506103f361119e565b34801561057b57600080fd5b506102de6004356111ed565b34801561059357600080fd5b506102b5600435602435611208565b3480156105ae57600080fd5b506102de600435611288565b3480156105c657600080fd5b506103f36004356024356044356064356112a3565b3480156105e757600080fd5b506103f360043561137b565b3480156105ff57600080fd5b506103f3600160a060020a0360043516602435611397565b34801561062357600080fd5b506103f3600160a060020a03600435166113f1565b34801561064457600080fd5b506103f3600435600160a060020a036024351661143f565b34801561066857600080fd5b506103f36004356114a6565b34801561068057600080fd5b506102b5611502565b34801561069557600080fd5b506102de600435611512565b3480156106ad57600080fd5b506103a8611536565b3480156106c257600080fd5b506103f3600160a060020a036004351661153c565b3480156106e357600080fd5b506103a8600160a060020a03600435166115ef565b34801561070457600080fd5b506103a861160a565b34801561071957600080fd5b506103f3611610565b34801561072e57600080fd5b50610743600160a060020a0360043516611690565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561077f578181015183820152602001610767565b505050509050019250505060405180910390f35b34801561079f57600080fd5b506103a8600435611762565b3480156107b757600080fd5b506103f3611a3e565b3480156107cc57600080fd5b5061031e611b53565b3480156107e157600080fd5b506107ed600435611b8a565b6040805163ffffffff9092168252519081900360200190f35b34801561081257600080fd5b506102de600435611bb7565b34801561082a57600080fd5b506103f3600160a060020a0360043516602435611bd2565b34801561084e57600080fd5b506102de611c6e565b34801561086357600080fd5b506103a8611c7d565b34801561087857600080fd5b506102de611c83565b34801561088d57600080fd5b506103f3600435611c92565b3480156108a557600080fd5b506102b5600435611d7c565b3480156108bd57600080fd5b506103a8611e4b565b3480156108d257600080fd5b506103f3600160a060020a0360043516611e51565b3480156108f357600080fd5b506102de611e8a565b6103f3600435602435611e99565b34801561091657600080fd5b506103a8612033565b34801561092b57600080fd5b506102de612039565b6103f3600435602435612048565b604080517f737570706f727473496e74657266616365286279746573342900000000000000815290519081900360190190206000907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1983811691161480610bcd5750604080517f746f6b656e4d657461646174612875696e743235362c737472696e67290000008152815190819003601d0181207f746f6b656e734f664f776e657228616464726573732900000000000000000000825282519182900360160182207f7472616e7366657246726f6d28616464726573732c616464726573732c75696e83527f7432353629000000000000000000000000000000000000000000000000000000602084015283519283900360250183207f7472616e7366657228616464726573732c75696e743235362900000000000000845284519384900360190184207f617070726f766528616464726573732c75696e74323536290000000000000000855285519485900360180185207f6f776e65724f662875696e743235362900000000000000000000000000000000865286519586900360100186207f62616c616e63654f662861646472657373290000000000000000000000000000875287519687900360120187207f746f74616c537570706c792829000000000000000000000000000000000000008852885197889003600d0188207f73796d626f6c2829000000000000000000000000000000000000000000000000895289519889900360080189207f6e616d65282900000000000000000000000000000000000000000000000000008a529951988990036006019098207bffffffffffffffffffffffffffffffffffffffffffffffffffffffff198c811691909a189098181818181818181891909116145b90505b919050565b600154600160a060020a031681565b6060610bee612c7f565b600d54600090600160a060020a03161515610c0857600080fd5b600d54604080517fcb4799f2000000000000000000000000000000000000000000000000000000008152600481018981526024820192835260448201889052600160a060020a039093169263cb4799f2928a928a928a929091606401848480828437820191505094505050505060a060405180830381600087803b158015610c8f57600080fd5b505af1158015610ca3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060a0811015610cc857600080fd5b5060808101519092509050610cdd8282612246565b9695505050505050565b60115481565b60408051808201909152600d81527f43727970746f4b69747469657300000000000000000000000000000000000000602082015281565b60025460a060020a900460ff1615610d3b57600080fd5b610d45338261229a565b1515610d5057600080fd5b610d5a81836122ba565b60408051338152600160a060020a038416602082015280820183905290517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360600190a15050565b600054600160a060020a031681565b662386f26fc1000081565b60008054600160a060020a03163314610dd857600080fd5b81905080600160a060020a03166376190f8f6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610e1957600080fd5b505af1158015610e2d573d6000803e3d6000fd5b505050506040513d6020811015610e4357600080fd5b50511515610e5057600080fd5b600c8054600160a060020a031916600160a060020a039290921691909117905550565b6006546000190190565b600f5481565b6000808211610e9157600080fd5b6006805483908110610e9f57fe5b600091825260209091206002909102016001015460c060020a900463ffffffff16151592915050565b6201518081565b600c54600160a060020a031681565b60025460a060020a900460ff1615610ef557600080fd5b600160a060020a0382161515610f0a57600080fd5b600160a060020a038216301415610f2057600080fd5b610f2a33826122e8565b1515610f3557600080fd5b610f3f838261229a565b1515610f4a57600080fd5b610f55838383612308565b505050565b60008054600160a060020a03163314610f7257600080fd5b81905080600160a060020a03166354c15b826040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610fb357600080fd5b505af1158015610fc7573d6000803e3d6000fd5b505050506040513d6020811015610fdd57600080fd5b50511515610fea57600080fd5b60108054600160a060020a031916600160a060020a039290921691909117905550565b600054600160a060020a0316331461102457600080fd5b600160a060020a038116151561103957600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a0316331461107257600080fd5b600160a060020a038116151561108757600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b60025460a060020a900460ff16156110c057600080fd5b6110ca338561229a565b15156110d557600080fd5b6110de84610e83565b156110e857600080fd5b600b546110ff908590600160a060020a03166122ba565b600b54604080517f27ebe40a000000000000000000000000000000000000000000000000000000008152600481018790526024810186905260448101859052606481018490523360848201529051600160a060020a03909216916327ebe40a9160a48082019260009290919082900301818387803b15801561118057600080fd5b505af1158015611194573d6000803e3d6000fd5b5050505050505050565b600054600160a060020a031633146111b557600080fd5b60025460a060020a900460ff1615156111cd57600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600a60205260009081526040902054600160a060020a031681565b6000808080851161121857600080fd5b6000841161122557600080fd5b600680548690811061123357fe5b9060005260206000209060020201915060068481548110151561125257fe5b9060005260206000209060020201905061126e828683876123ea565b801561127f575061127f848661256a565b95945050505050565b600960205260009081526040902054600160a060020a031681565b60025460a060020a900460ff16156112ba57600080fd5b6112c4338561229a565b15156112cf57600080fd5b6112d884611d7c565b15156112e357600080fd5b600c546112fa908590600160a060020a03166122ba565b600c54604080517f27ebe40a000000000000000000000000000000000000000000000000000000008152600481018790526024810186905260448101859052606481018490523360848201529051600160a060020a03909216916327ebe40a9160a48082019260009290919082900301818387803b15801561118057600080fd5b600254600160a060020a0316331461139257600080fd5b600e55565b60025460a060020a900460ff16156113ae57600080fd5b6113b8338261229a565b15156113c357600080fd5b6000908152600a602052604090208054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a0316331461140857600080fd5b600160a060020a038116151561141d57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b600254600090600160a060020a0316331461145957600080fd5b5080600160a060020a03811615156114795750600254600160a060020a03165b6011546113881161148957600080fd5b6011805460010190556114a06000808086856125bf565b50505050565b600254600160a060020a03163314806114c95750600054600160a060020a031633145b806114de5750600154600160a060020a031633145b15156114e957600080fd5b60035463ffffffff1681106114fd57600080fd5b600555565b60025460a060020a900460ff1681565b600081815260076020526040902054600160a060020a0316801515610bd057600080fd5b61afc881565b60008054600160a060020a0316331461155457600080fd5b81905080600160a060020a03166385b861886040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561159557600080fd5b505af11580156115a9573d6000803e3d6000fd5b505050506040513d60208110156115bf57600080fd5b505115156115cc57600080fd5b600b8054600160a060020a031916600160a060020a039290921691909117905550565b600160a060020a031660009081526008602052604090205490565b60055481565b600254600160a060020a03163314806116335750600054600160a060020a031633145b806116485750600154600160a060020a031633145b151561165357600080fd5b60025460a060020a900460ff161561166a57600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a179055565b60606000606060008060006116a4876115ef565b94508415156116c3576040805160008152602081019091529550611758565b846040519080825280602002602001820160405280156116ed578160200160208202803883390190505b5093506116f8610e73565b925060009150600190505b82811161175457600081815260076020526040902054600160a060020a038881169116141561174c5780848381518110151561173b57fe5b602090810290910101526001909101905b600101611703565b8395505b5050505050919050565b600080600080600080600080600260149054906101000a900460ff1615151561178a57600080fd5b600680548a90811061179857fe5b60009182526020909120600290910201600181015490975067ffffffffffffffff1615156117c557600080fd5b604080516101008101825288548152600189015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e08201526118569061289c565b151561186157600080fd5b60018701546006805460c060020a90920463ffffffff169750908790811061188557fe5b600091825260209091206001808a015460029093029091019081015490965061ffff60f060020a9283900481169650919004168410156118d257600185015460f060020a900461ffff1693505b6010548754865460018a0154604080517f0d9f5aed0000000000000000000000000000000000000000000000000000000081526004810194909452602484019290925260001967ffffffffffffffff6801000000000000000090920482160116604483015251600160a060020a0390921691630d9f5aed916064808201926020929091908290030181600087803b15801561196c57600080fd5b505af1158015611980573d6000803e3d6000fd5b505050506040513d602081101561199657600080fd5b505160008a815260076020526040902054600189810154929550600160a060020a0390911693506119df918b9160c060020a90910463ffffffff1690870161ffff1686866125bf565b6001880180547bffffffff00000000000000000000000000000000000000000000000019169055600f8054600019019055600e54604051919250339181156108fc0291906000818181858888f150939c9b505050505050505050505050565b600254600160a060020a0316331480611a615750600054600160a060020a031633145b80611a765750600154600160a060020a031633145b1515611a8157600080fd5b600b60009054906101000a9004600160a060020a0316600160a060020a0316635fd8c7106040518163ffffffff1660e060020a028152600401600060405180830381600087803b158015611ad457600080fd5b505af1158015611ae8573d6000803e3d6000fd5b50505050600c60009054906101000a9004600160a060020a0316600160a060020a0316635fd8c7106040518163ffffffff1660e060020a028152600401600060405180830381600087803b158015611b3f57600080fd5b505af11580156114a0573d6000803e3d6000fd5b60408051808201909152600281527f434b000000000000000000000000000000000000000000000000000000000000602082015281565b600381600e8110611b9757fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b600760205260009081526040902054600160a060020a031681565b60025460a060020a900460ff1615611be957600080fd5b600160a060020a0382161515611bfe57600080fd5b600160a060020a038216301415611c1457600080fd5b600b54600160a060020a0383811691161415611c2f57600080fd5b600c54600160a060020a0383811691161415611c4a57600080fd5b611c54338261229a565b1515611c5f57600080fd5b611c6a338383612308565b5050565b600254600160a060020a031681565b600e5481565b600d54600160a060020a031681565b600254600090600160a060020a03163314611cac57600080fd5b60125461afc811611cbc57600080fd5b611ccb600080600085306125bf565b600b54909150611ce5908290600160a060020a03166122ba565b600b54600160a060020a03166327ebe40a82611cff6128cc565b6040805160e060020a63ffffffff861602815260048101939093526024830191909152600060448301819052620151806064840152306084840152905160a48084019382900301818387803b158015611d5757600080fd5b505af1158015611d6b573d6000803e3d6000fd5b505060128054600101905550505050565b600080808311611d8b57600080fd5b6006805484908110611d9957fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e0820152909150611e4490612996565b9392505050565b61138881565b600054600160a060020a03163314611e6857600080fd5b600d8054600160a060020a031916600160a060020a0392909216919091179055565b600b54600160a060020a031681565b60025460009060a060020a900460ff1615611eb357600080fd5b611ebd338361229a565b1515611ec857600080fd5b611ed182611d7c565b1515611edc57600080fd5b611ee682846129c5565b1515611ef157600080fd5b600c54604080517fc55d0f56000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a039092169163c55d0f56916024808201926020929091908290030181600087803b158015611f5857600080fd5b505af1158015611f6c573d6000803e3d6000fd5b505050506040513d6020811015611f8257600080fd5b5051600e549091508101341015611f9857600080fd5b600c54600e54604080517f454a2ab3000000000000000000000000000000000000000000000000000000008152600481018790529051600160a060020a039093169263454a2ab39234039160248082019260009290919082900301818588803b15801561200457600080fd5b505af1158015612018573d6000803e3d6000fd5b5050505050610f558263ffffffff168463ffffffff16612a14565b60125481565b601054600160a060020a031681565b600254600090819060a060020a900460ff161561206457600080fd5b600e5434101561207357600080fd5b61207d338561229a565b151561208857600080fd5b612092838561256a565b151561209d57600080fd5b60068054859081106120ab57fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e082015290925061215690612996565b151561216157600080fd5b600680548490811061216f57fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e082015290915061221a90612996565b151561222557600080fd5b612231828583866123ea565b151561223c57600080fd5b6114a08484612a14565b606080600080846040519080825280601f01601f19166020018201604052801561227a578160200160208202803883390190505b5092505060208201905084612290828287612b52565b5090949350505050565b600090815260076020526040902054600160a060020a0391821691161490565b6000918252600960205260409091208054600160a060020a031916600160a060020a03909216919091179055565b600090815260096020526040902054600160a060020a0391821691161490565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a031916909117905583161561239b57600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b60408051600160a060020a0380861682528416602082015280820183905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a1505050565b6000818414156123fc57506000612562565b6001850154608060020a900463ffffffff1682148061242b5750600185015460a060020a900463ffffffff1682145b1561243857506000612562565b6001830154608060020a900463ffffffff168414806124675750600183015460a060020a900463ffffffff1684145b1561247457506000612562565b6001830154608060020a900463ffffffff1615806124a157506001850154608060020a900463ffffffff16155b156124ae57506001612562565b60018581015490840154608060020a9182900463ffffffff908116929091041614806124f9575060018086015490840154608060020a900463ffffffff90811660a060020a90920416145b1561250657506000612562565b6001808601549084015460a060020a900463ffffffff908116608060020a90920416148061255157506001858101549084015460a060020a9182900463ffffffff9081169290910416145b1561255e57506000612562565b5060015b949350505050565b6000818152600760205260408082205484835290822054600160a060020a0391821691168082148061127f57506000858152600a6020526040902054600160a060020a03908116908316149250505092915050565b6000806125ca612c9e565b600063ffffffff891689146125de57600080fd5b63ffffffff881688146125f057600080fd5b61ffff8716871461260057600080fd5b600287049250600d8361ffff16111561261857600d92505b505060408051610100810182528581524267ffffffffffffffff90811660208301908152600093830184815263ffffffff8c8116606086019081528c82166080870190815260a0870188815261ffff8a811660c08a019081528f821660e08b01908152600680546001810182559c528a517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f60028e029081019190915598517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d40909901805498519651955194519251915167ffffffffffffffff19909916998b16999099176fffffffffffffffff000000000000000019166801000000000000000096909a16959095029890981773ffffffff000000000000000000000000000000001916608060020a938616939093029290921777ffffffff0000000000000000000000000000000000000000191660a060020a91851691909102177bffffffff000000000000000000000000000000000000000000000000191660c060020a96841696909602959095177fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e060020a91861691909102177dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f060020a92909416919091029290921790559091908116811461281657600080fd5b606080830151608080850151855160408051600160a060020a038c1681526020810188905263ffffffff95861681830152929094169482019490945290810192909252517f0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad59181900360a00190a161289060008683612308565b98975050505050505050565b60008160a0015163ffffffff16600014158015610bcd5750506040015167ffffffffffffffff4381169116111590565b6000806000600b60009054906101000a9004600160a060020a0316600160a060020a031663eac9d94c6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561292457600080fd5b505af1158015612938573d6000803e3d6000fd5b505050506040513d602081101561294e57600080fd5b505191506fffffffffffffffffffffffffffffffff8216821461297057600080fd5b50600281048101662386f26fc100008110156129905750662386f26fc100005b92915050565b60008160a0015163ffffffff166000148015610bcd5750506040015167ffffffffffffffff4381169116111590565b60008060006006858154811015156129d957fe5b906000526020600020906002020191506006848154811015156129f857fe5b9060005260206000209060020201905061127f828683876123ea565b600080600683815481101515612a2657fe5b90600052602060002090600202019150600684815481101515612a4557fe5b600091825260209091206002909102016001810180547bffffffff000000000000000000000000000000000000000000000000191660c060020a63ffffffff8716021790559050612a9582612b96565b612a9e81612b96565b6000848152600a602090815260408083208054600160a060020a031990811690915586845281842080549091169055600f80546001908101909155878452600783529281902054928401548151600160a060020a0390941684529183018790528281018690526801000000000000000090910467ffffffffffffffff166060830152517f241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80916080908290030190a150505050565b60005b60208210612b77578251845260209384019390920191601f1990910190612b55565b50905182516020929092036101000a6000190180199091169116179052565b600554600182015443919060039060e060020a900461ffff16600e8110612bb957fe5b600891828204019190066004029054906101000a900463ffffffff1663ffffffff16811515612be457fe5b6001840180546fffffffffffffffff0000000000000000191668010000000000000000939092049390930167ffffffffffffffff16919091021790819055600d60e060020a90910461ffff161015612c7c576001818101805461ffff60e060020a8083048216909401169092027fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790555b50565b6080604051908101604052806004906020820280388339509192915050565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810191909152905600a165627a7a72305820d8531d9a3f811656e3c381ca06aa261f9fe24f242b7e416335eac0bd438de5e90029`

// DeployKittyMinting deploys a new Ethereum contract, binding an instance of KittyMinting to it.
func DeployKittyMinting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyMinting, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyMintingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyMintingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyMinting{KittyMintingCaller: KittyMintingCaller{contract: contract}, KittyMintingTransactor: KittyMintingTransactor{contract: contract}, KittyMintingFilterer: KittyMintingFilterer{contract: contract}}, nil
}

// KittyMinting is an auto generated Go binding around an Ethereum contract.
type KittyMinting struct {
	KittyMintingCaller     // Read-only binding to the contract
	KittyMintingTransactor // Write-only binding to the contract
	KittyMintingFilterer   // Log filterer for contract events
}

// KittyMintingCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyMintingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyMintingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyMintingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyMintingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyMintingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyMintingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyMintingSession struct {
	Contract     *KittyMinting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittyMintingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyMintingCallerSession struct {
	Contract *KittyMintingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KittyMintingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyMintingTransactorSession struct {
	Contract     *KittyMintingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KittyMintingRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyMintingRaw struct {
	Contract *KittyMinting // Generic contract binding to access the raw methods on
}

// KittyMintingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyMintingCallerRaw struct {
	Contract *KittyMintingCaller // Generic read-only contract binding to access the raw methods on
}

// KittyMintingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyMintingTransactorRaw struct {
	Contract *KittyMintingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyMinting creates a new instance of KittyMinting, bound to a specific deployed contract.
func NewKittyMinting(address common.Address, backend bind.ContractBackend) (*KittyMinting, error) {
	contract, err := bindKittyMinting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyMinting{KittyMintingCaller: KittyMintingCaller{contract: contract}, KittyMintingTransactor: KittyMintingTransactor{contract: contract}, KittyMintingFilterer: KittyMintingFilterer{contract: contract}}, nil
}

// NewKittyMintingCaller creates a new read-only instance of KittyMinting, bound to a specific deployed contract.
func NewKittyMintingCaller(address common.Address, caller bind.ContractCaller) (*KittyMintingCaller, error) {
	contract, err := bindKittyMinting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyMintingCaller{contract: contract}, nil
}

// NewKittyMintingTransactor creates a new write-only instance of KittyMinting, bound to a specific deployed contract.
func NewKittyMintingTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyMintingTransactor, error) {
	contract, err := bindKittyMinting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyMintingTransactor{contract: contract}, nil
}

// NewKittyMintingFilterer creates a new log filterer instance of KittyMinting, bound to a specific deployed contract.
func NewKittyMintingFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyMintingFilterer, error) {
	contract, err := bindKittyMinting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyMintingFilterer{contract: contract}, nil
}

// bindKittyMinting binds a generic wrapper to an already deployed contract.
func bindKittyMinting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyMintingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyMinting *KittyMintingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyMinting.Contract.KittyMintingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyMinting *KittyMintingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyMinting.Contract.KittyMintingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyMinting *KittyMintingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyMinting.Contract.KittyMintingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyMinting *KittyMintingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyMinting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyMinting *KittyMintingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyMinting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyMinting *KittyMintingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyMinting.Contract.contract.Transact(opts, method, params...)
}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() constant returns(uint256)
func (_KittyMinting *KittyMintingCaller) GEN0AUCTIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "GEN0_AUCTION_DURATION")
	return *ret0, err
}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() constant returns(uint256)
func (_KittyMinting *KittyMintingSession) GEN0AUCTIONDURATION() (*big.Int, error) {
	return _KittyMinting.Contract.GEN0AUCTIONDURATION(&_KittyMinting.CallOpts)
}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() constant returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) GEN0AUCTIONDURATION() (*big.Int, error) {
	return _KittyMinting.Contract.GEN0AUCTIONDURATION(&_KittyMinting.CallOpts)
}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() constant returns(uint256)
func (_KittyMinting *KittyMintingCaller) GEN0CREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "GEN0_CREATION_LIMIT")
	return *ret0, err
}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() constant returns(uint256)
func (_KittyMinting *KittyMintingSession) GEN0CREATIONLIMIT() (*big.Int, error) {
	return _KittyMinting.Contract.GEN0CREATIONLIMIT(&_KittyMinting.CallOpts)
}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() constant returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) GEN0CREATIONLIMIT() (*big.Int, error) {
	return _KittyMinting.Contract.GEN0CREATIONLIMIT(&_KittyMinting.CallOpts)
}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() constant returns(uint256)
func (_KittyMinting *KittyMintingCaller) GEN0STARTINGPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "GEN0_STARTING_PRICE")
	return *ret0, err
}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() constant returns(uint256)
func (_KittyMinting *KittyMintingSession) GEN0STARTINGPRICE() (*big.Int, error) {
	return _KittyMinting.Contract.GEN0STARTINGPRICE(&_KittyMinting.CallOpts)
}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() constant returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) GEN0STARTINGPRICE() (*big.Int, error) {
	return _KittyMinting.Contract.GEN0STARTINGPRICE(&_KittyMinting.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_KittyMinting *KittyMintingCaller) PROMOCREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "PROMO_CREATION_LIMIT")
	return *ret0, err
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_KittyMinting *KittyMintingSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _KittyMinting.Contract.PROMOCREATIONLIMIT(&_KittyMinting.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _KittyMinting.Contract.PROMOCREATIONLIMIT(&_KittyMinting.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_KittyMinting *KittyMintingCaller) AutoBirthFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "autoBirthFee")
	return *ret0, err
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_KittyMinting *KittyMintingSession) AutoBirthFee() (*big.Int, error) {
	return _KittyMinting.Contract.AutoBirthFee(&_KittyMinting.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) AutoBirthFee() (*big.Int, error) {
	return _KittyMinting.Contract.AutoBirthFee(&_KittyMinting.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyMinting *KittyMintingCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyMinting *KittyMintingSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyMinting.Contract.BalanceOf(&_KittyMinting.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyMinting *KittyMintingCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyMinting.Contract.BalanceOf(&_KittyMinting.CallOpts, _owner)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_KittyMinting *KittyMintingCaller) CanBreedWith(opts *bind.CallOpts, _matronId *big.Int, _sireId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "canBreedWith", _matronId, _sireId)
	return *ret0, err
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_KittyMinting *KittyMintingSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyMinting.Contract.CanBreedWith(&_KittyMinting.CallOpts, _matronId, _sireId)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_KittyMinting *KittyMintingCallerSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyMinting.Contract.CanBreedWith(&_KittyMinting.CallOpts, _matronId, _sireId)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyMinting *KittyMintingCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyMinting *KittyMintingSession) CeoAddress() (common.Address, error) {
	return _KittyMinting.Contract.CeoAddress(&_KittyMinting.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyMinting *KittyMintingCallerSession) CeoAddress() (common.Address, error) {
	return _KittyMinting.Contract.CeoAddress(&_KittyMinting.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyMinting *KittyMintingCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyMinting *KittyMintingSession) CfoAddress() (common.Address, error) {
	return _KittyMinting.Contract.CfoAddress(&_KittyMinting.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyMinting *KittyMintingCallerSession) CfoAddress() (common.Address, error) {
	return _KittyMinting.Contract.CfoAddress(&_KittyMinting.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyMinting *KittyMintingCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyMinting *KittyMintingSession) CooAddress() (common.Address, error) {
	return _KittyMinting.Contract.CooAddress(&_KittyMinting.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyMinting *KittyMintingCallerSession) CooAddress() (common.Address, error) {
	return _KittyMinting.Contract.CooAddress(&_KittyMinting.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyMinting *KittyMintingCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "cooldowns", arg0)
	return *ret0, err
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyMinting *KittyMintingSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyMinting.Contract.Cooldowns(&_KittyMinting.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyMinting *KittyMintingCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyMinting.Contract.Cooldowns(&_KittyMinting.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyMinting *KittyMintingCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "erc721Metadata")
	return *ret0, err
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyMinting *KittyMintingSession) Erc721Metadata() (common.Address, error) {
	return _KittyMinting.Contract.Erc721Metadata(&_KittyMinting.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyMinting *KittyMintingCallerSession) Erc721Metadata() (common.Address, error) {
	return _KittyMinting.Contract.Erc721Metadata(&_KittyMinting.CallOpts)
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() constant returns(uint256)
func (_KittyMinting *KittyMintingCaller) Gen0CreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "gen0CreatedCount")
	return *ret0, err
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() constant returns(uint256)
func (_KittyMinting *KittyMintingSession) Gen0CreatedCount() (*big.Int, error) {
	return _KittyMinting.Contract.Gen0CreatedCount(&_KittyMinting.CallOpts)
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() constant returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) Gen0CreatedCount() (*big.Int, error) {
	return _KittyMinting.Contract.Gen0CreatedCount(&_KittyMinting.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_KittyMinting *KittyMintingCaller) GeneScience(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "geneScience")
	return *ret0, err
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_KittyMinting *KittyMintingSession) GeneScience() (common.Address, error) {
	return _KittyMinting.Contract.GeneScience(&_KittyMinting.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_KittyMinting *KittyMintingCallerSession) GeneScience() (common.Address, error) {
	return _KittyMinting.Contract.GeneScience(&_KittyMinting.CallOpts)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_KittyMinting *KittyMintingCaller) IsPregnant(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "isPregnant", _kittyId)
	return *ret0, err
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_KittyMinting *KittyMintingSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyMinting.Contract.IsPregnant(&_KittyMinting.CallOpts, _kittyId)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_KittyMinting *KittyMintingCallerSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyMinting.Contract.IsPregnant(&_KittyMinting.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_KittyMinting *KittyMintingCaller) IsReadyToBreed(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "isReadyToBreed", _kittyId)
	return *ret0, err
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_KittyMinting *KittyMintingSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyMinting.Contract.IsReadyToBreed(&_KittyMinting.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_KittyMinting *KittyMintingCallerSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyMinting.Contract.IsReadyToBreed(&_KittyMinting.CallOpts, _kittyId)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyMinting *KittyMintingCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "kittyIndexToApproved", arg0)
	return *ret0, err
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyMinting *KittyMintingSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.KittyIndexToApproved(&_KittyMinting.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyMinting *KittyMintingCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.KittyIndexToApproved(&_KittyMinting.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyMinting *KittyMintingCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "kittyIndexToOwner", arg0)
	return *ret0, err
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyMinting *KittyMintingSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.KittyIndexToOwner(&_KittyMinting.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyMinting *KittyMintingCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.KittyIndexToOwner(&_KittyMinting.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyMinting *KittyMintingCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyMinting *KittyMintingSession) Name() (string, error) {
	return _KittyMinting.Contract.Name(&_KittyMinting.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyMinting *KittyMintingCallerSession) Name() (string, error) {
	return _KittyMinting.Contract.Name(&_KittyMinting.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyMinting *KittyMintingCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyMinting *KittyMintingSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.OwnerOf(&_KittyMinting.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyMinting *KittyMintingCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.OwnerOf(&_KittyMinting.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyMinting *KittyMintingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyMinting *KittyMintingSession) Paused() (bool, error) {
	return _KittyMinting.Contract.Paused(&_KittyMinting.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyMinting *KittyMintingCallerSession) Paused() (bool, error) {
	return _KittyMinting.Contract.Paused(&_KittyMinting.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_KittyMinting *KittyMintingCaller) PregnantKitties(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "pregnantKitties")
	return *ret0, err
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_KittyMinting *KittyMintingSession) PregnantKitties() (*big.Int, error) {
	return _KittyMinting.Contract.PregnantKitties(&_KittyMinting.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) PregnantKitties() (*big.Int, error) {
	return _KittyMinting.Contract.PregnantKitties(&_KittyMinting.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_KittyMinting *KittyMintingCaller) PromoCreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "promoCreatedCount")
	return *ret0, err
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_KittyMinting *KittyMintingSession) PromoCreatedCount() (*big.Int, error) {
	return _KittyMinting.Contract.PromoCreatedCount(&_KittyMinting.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) PromoCreatedCount() (*big.Int, error) {
	return _KittyMinting.Contract.PromoCreatedCount(&_KittyMinting.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyMinting *KittyMintingCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "saleAuction")
	return *ret0, err
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyMinting *KittyMintingSession) SaleAuction() (common.Address, error) {
	return _KittyMinting.Contract.SaleAuction(&_KittyMinting.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyMinting *KittyMintingCallerSession) SaleAuction() (common.Address, error) {
	return _KittyMinting.Contract.SaleAuction(&_KittyMinting.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyMinting *KittyMintingCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "secondsPerBlock")
	return *ret0, err
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyMinting *KittyMintingSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyMinting.Contract.SecondsPerBlock(&_KittyMinting.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyMinting.Contract.SecondsPerBlock(&_KittyMinting.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyMinting *KittyMintingCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "sireAllowedToAddress", arg0)
	return *ret0, err
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyMinting *KittyMintingSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.SireAllowedToAddress(&_KittyMinting.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyMinting *KittyMintingCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.SireAllowedToAddress(&_KittyMinting.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyMinting *KittyMintingCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "siringAuction")
	return *ret0, err
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyMinting *KittyMintingSession) SiringAuction() (common.Address, error) {
	return _KittyMinting.Contract.SiringAuction(&_KittyMinting.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyMinting *KittyMintingCallerSession) SiringAuction() (common.Address, error) {
	return _KittyMinting.Contract.SiringAuction(&_KittyMinting.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyMinting *KittyMintingCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "supportsInterface", _interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyMinting *KittyMintingSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyMinting.Contract.SupportsInterface(&_KittyMinting.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyMinting *KittyMintingCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyMinting.Contract.SupportsInterface(&_KittyMinting.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyMinting *KittyMintingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyMinting *KittyMintingSession) Symbol() (string, error) {
	return _KittyMinting.Contract.Symbol(&_KittyMinting.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyMinting *KittyMintingCallerSession) Symbol() (string, error) {
	return _KittyMinting.Contract.Symbol(&_KittyMinting.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyMinting *KittyMintingCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "tokenMetadata", _tokenId, _preferredTransport)
	return *ret0, err
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyMinting *KittyMintingSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyMinting.Contract.TokenMetadata(&_KittyMinting.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyMinting *KittyMintingCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyMinting.Contract.TokenMetadata(&_KittyMinting.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyMinting *KittyMintingCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyMinting *KittyMintingSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyMinting.Contract.TokensOfOwner(&_KittyMinting.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyMinting *KittyMintingCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyMinting.Contract.TokensOfOwner(&_KittyMinting.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyMinting *KittyMintingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyMinting.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyMinting *KittyMintingSession) TotalSupply() (*big.Int, error) {
	return _KittyMinting.Contract.TotalSupply(&_KittyMinting.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) TotalSupply() (*big.Int, error) {
	return _KittyMinting.Contract.TotalSupply(&_KittyMinting.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyMinting *KittyMintingTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyMinting *KittyMintingSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.Approve(&_KittyMinting.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyMinting *KittyMintingTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.Approve(&_KittyMinting.TransactOpts, _to, _tokenId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_KittyMinting *KittyMintingTransactor) ApproveSiring(opts *bind.TransactOpts, _addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "approveSiring", _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_KittyMinting *KittyMintingSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.ApproveSiring(&_KittyMinting.TransactOpts, _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_KittyMinting *KittyMintingTransactorSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.ApproveSiring(&_KittyMinting.TransactOpts, _addr, _sireId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(_sireId uint256, _matronId uint256) returns()
func (_KittyMinting *KittyMintingTransactor) BidOnSiringAuction(opts *bind.TransactOpts, _sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "bidOnSiringAuction", _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(_sireId uint256, _matronId uint256) returns()
func (_KittyMinting *KittyMintingSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.BidOnSiringAuction(&_KittyMinting.TransactOpts, _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(_sireId uint256, _matronId uint256) returns()
func (_KittyMinting *KittyMintingTransactorSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.BidOnSiringAuction(&_KittyMinting.TransactOpts, _sireId, _matronId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_KittyMinting *KittyMintingTransactor) BreedWithAuto(opts *bind.TransactOpts, _matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "breedWithAuto", _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_KittyMinting *KittyMintingSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.BreedWithAuto(&_KittyMinting.TransactOpts, _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_KittyMinting *KittyMintingTransactorSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.BreedWithAuto(&_KittyMinting.TransactOpts, _matronId, _sireId)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(_genes uint256) returns()
func (_KittyMinting *KittyMintingTransactor) CreateGen0Auction(opts *bind.TransactOpts, _genes *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "createGen0Auction", _genes)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(_genes uint256) returns()
func (_KittyMinting *KittyMintingSession) CreateGen0Auction(_genes *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreateGen0Auction(&_KittyMinting.TransactOpts, _genes)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(_genes uint256) returns()
func (_KittyMinting *KittyMintingTransactorSession) CreateGen0Auction(_genes *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreateGen0Auction(&_KittyMinting.TransactOpts, _genes)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(_genes uint256, _owner address) returns()
func (_KittyMinting *KittyMintingTransactor) CreatePromoKitty(opts *bind.TransactOpts, _genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "createPromoKitty", _genes, _owner)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(_genes uint256, _owner address) returns()
func (_KittyMinting *KittyMintingSession) CreatePromoKitty(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreatePromoKitty(&_KittyMinting.TransactOpts, _genes, _owner)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(_genes uint256, _owner address) returns()
func (_KittyMinting *KittyMintingTransactorSession) CreatePromoKitty(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreatePromoKitty(&_KittyMinting.TransactOpts, _genes, _owner)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyMinting *KittyMintingTransactor) CreateSaleAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "createSaleAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyMinting *KittyMintingSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreateSaleAuction(&_KittyMinting.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyMinting *KittyMintingTransactorSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreateSaleAuction(&_KittyMinting.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyMinting *KittyMintingTransactor) CreateSiringAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "createSiringAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyMinting *KittyMintingSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreateSiringAuction(&_KittyMinting.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_KittyMinting *KittyMintingTransactorSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreateSiringAuction(&_KittyMinting.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_KittyMinting *KittyMintingTransactor) GiveBirth(opts *bind.TransactOpts, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "giveBirth", _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_KittyMinting *KittyMintingSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.GiveBirth(&_KittyMinting.TransactOpts, _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_KittyMinting *KittyMintingTransactorSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.GiveBirth(&_KittyMinting.TransactOpts, _matronId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyMinting *KittyMintingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyMinting *KittyMintingSession) Pause() (*types.Transaction, error) {
	return _KittyMinting.Contract.Pause(&_KittyMinting.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyMinting *KittyMintingTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyMinting.Contract.Pause(&_KittyMinting.TransactOpts)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_KittyMinting *KittyMintingTransactor) SetAutoBirthFee(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setAutoBirthFee", val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_KittyMinting *KittyMintingSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetAutoBirthFee(&_KittyMinting.TransactOpts, val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetAutoBirthFee(&_KittyMinting.TransactOpts, val)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyMinting *KittyMintingTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyMinting *KittyMintingSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetCEO(&_KittyMinting.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetCEO(&_KittyMinting.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyMinting *KittyMintingTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyMinting *KittyMintingSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetCFO(&_KittyMinting.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetCFO(&_KittyMinting.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyMinting *KittyMintingTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyMinting *KittyMintingSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetCOO(&_KittyMinting.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetCOO(&_KittyMinting.TransactOpts, _newCOO)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_KittyMinting *KittyMintingTransactor) SetGeneScienceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setGeneScienceAddress", _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_KittyMinting *KittyMintingSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetGeneScienceAddress(&_KittyMinting.TransactOpts, _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetGeneScienceAddress(&_KittyMinting.TransactOpts, _address)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyMinting *KittyMintingTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyMinting *KittyMintingSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetMetadataAddress(&_KittyMinting.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetMetadataAddress(&_KittyMinting.TransactOpts, _contractAddress)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(_address address) returns()
func (_KittyMinting *KittyMintingTransactor) SetSaleAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setSaleAuctionAddress", _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(_address address) returns()
func (_KittyMinting *KittyMintingSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetSaleAuctionAddress(&_KittyMinting.TransactOpts, _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(_address address) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetSaleAuctionAddress(&_KittyMinting.TransactOpts, _address)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyMinting *KittyMintingTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyMinting *KittyMintingSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetSecondsPerBlock(&_KittyMinting.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetSecondsPerBlock(&_KittyMinting.TransactOpts, secs)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(_address address) returns()
func (_KittyMinting *KittyMintingTransactor) SetSiringAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setSiringAuctionAddress", _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(_address address) returns()
func (_KittyMinting *KittyMintingSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetSiringAuctionAddress(&_KittyMinting.TransactOpts, _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(_address address) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetSiringAuctionAddress(&_KittyMinting.TransactOpts, _address)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyMinting *KittyMintingTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyMinting *KittyMintingSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.Transfer(&_KittyMinting.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyMinting *KittyMintingTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.Transfer(&_KittyMinting.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyMinting *KittyMintingTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyMinting *KittyMintingSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.TransferFrom(&_KittyMinting.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyMinting *KittyMintingTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.TransferFrom(&_KittyMinting.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyMinting *KittyMintingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyMinting *KittyMintingSession) Unpause() (*types.Transaction, error) {
	return _KittyMinting.Contract.Unpause(&_KittyMinting.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyMinting *KittyMintingTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyMinting.Contract.Unpause(&_KittyMinting.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyMinting *KittyMintingTransactor) WithdrawAuctionBalances(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "withdrawAuctionBalances")
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyMinting *KittyMintingSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _KittyMinting.Contract.WithdrawAuctionBalances(&_KittyMinting.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyMinting *KittyMintingTransactorSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _KittyMinting.Contract.WithdrawAuctionBalances(&_KittyMinting.TransactOpts)
}

// KittyMintingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KittyMinting contract.
type KittyMintingApprovalIterator struct {
	Event *KittyMintingApproval // Event containing the contract specifics and raw log

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
func (it *KittyMintingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyMintingApproval)
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
		it.Event = new(KittyMintingApproval)
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
func (it *KittyMintingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyMintingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyMintingApproval represents a Approval event raised by the KittyMinting contract.
type KittyMintingApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_KittyMinting *KittyMintingFilterer) FilterApproval(opts *bind.FilterOpts) (*KittyMintingApprovalIterator, error) {

	logs, sub, err := _KittyMinting.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &KittyMintingApprovalIterator{contract: _KittyMinting.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_KittyMinting *KittyMintingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KittyMintingApproval) (event.Subscription, error) {

	logs, sub, err := _KittyMinting.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyMintingApproval)
				if err := _KittyMinting.contract.UnpackLog(event, "Approval", log); err != nil {
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

// KittyMintingBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the KittyMinting contract.
type KittyMintingBirthIterator struct {
	Event *KittyMintingBirth // Event containing the contract specifics and raw log

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
func (it *KittyMintingBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyMintingBirth)
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
		it.Event = new(KittyMintingBirth)
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
func (it *KittyMintingBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyMintingBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyMintingBirth represents a Birth event raised by the KittyMinting contract.
type KittyMintingBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_KittyMinting *KittyMintingFilterer) FilterBirth(opts *bind.FilterOpts) (*KittyMintingBirthIterator, error) {

	logs, sub, err := _KittyMinting.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &KittyMintingBirthIterator{contract: _KittyMinting.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_KittyMinting *KittyMintingFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *KittyMintingBirth) (event.Subscription, error) {

	logs, sub, err := _KittyMinting.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyMintingBirth)
				if err := _KittyMinting.contract.UnpackLog(event, "Birth", log); err != nil {
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

// KittyMintingContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyMinting contract.
type KittyMintingContractUpgradeIterator struct {
	Event *KittyMintingContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyMintingContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyMintingContractUpgrade)
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
		it.Event = new(KittyMintingContractUpgrade)
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
func (it *KittyMintingContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyMintingContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyMintingContractUpgrade represents a ContractUpgrade event raised by the KittyMinting contract.
type KittyMintingContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyMinting *KittyMintingFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyMintingContractUpgradeIterator, error) {

	logs, sub, err := _KittyMinting.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyMintingContractUpgradeIterator{contract: _KittyMinting.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyMinting *KittyMintingFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyMintingContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyMinting.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyMintingContractUpgrade)
				if err := _KittyMinting.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// KittyMintingPregnantIterator is returned from FilterPregnant and is used to iterate over the raw logs and unpacked data for Pregnant events raised by the KittyMinting contract.
type KittyMintingPregnantIterator struct {
	Event *KittyMintingPregnant // Event containing the contract specifics and raw log

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
func (it *KittyMintingPregnantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyMintingPregnant)
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
		it.Event = new(KittyMintingPregnant)
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
func (it *KittyMintingPregnantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyMintingPregnantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyMintingPregnant represents a Pregnant event raised by the KittyMinting contract.
type KittyMintingPregnant struct {
	Owner            common.Address
	MatronId         *big.Int
	SireId           *big.Int
	CooldownEndBlock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPregnant is a free log retrieval operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: e Pregnant(owner address, matronId uint256, sireId uint256, cooldownEndBlock uint256)
func (_KittyMinting *KittyMintingFilterer) FilterPregnant(opts *bind.FilterOpts) (*KittyMintingPregnantIterator, error) {

	logs, sub, err := _KittyMinting.contract.FilterLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return &KittyMintingPregnantIterator{contract: _KittyMinting.contract, event: "Pregnant", logs: logs, sub: sub}, nil
}

// WatchPregnant is a free log subscription operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: e Pregnant(owner address, matronId uint256, sireId uint256, cooldownEndBlock uint256)
func (_KittyMinting *KittyMintingFilterer) WatchPregnant(opts *bind.WatchOpts, sink chan<- *KittyMintingPregnant) (event.Subscription, error) {

	logs, sub, err := _KittyMinting.contract.WatchLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyMintingPregnant)
				if err := _KittyMinting.contract.UnpackLog(event, "Pregnant", log); err != nil {
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

// KittyMintingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KittyMinting contract.
type KittyMintingTransferIterator struct {
	Event *KittyMintingTransfer // Event containing the contract specifics and raw log

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
func (it *KittyMintingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyMintingTransfer)
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
		it.Event = new(KittyMintingTransfer)
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
func (it *KittyMintingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyMintingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyMintingTransfer represents a Transfer event raised by the KittyMinting contract.
type KittyMintingTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_KittyMinting *KittyMintingFilterer) FilterTransfer(opts *bind.FilterOpts) (*KittyMintingTransferIterator, error) {

	logs, sub, err := _KittyMinting.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &KittyMintingTransferIterator{contract: _KittyMinting.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_KittyMinting *KittyMintingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KittyMintingTransfer) (event.Subscription, error) {

	logs, sub, err := _KittyMinting.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyMintingTransfer)
				if err := _KittyMinting.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// KittyOwnershipABI is the input ABI used to generate the binding from.
const KittyOwnershipABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyOwnershipBin is the compiled bytecode used for deploying new contracts.
const KittyOwnershipBin = `0x6002805460a060020a60ff0219169055610240604052603c6080908152607860a05261012c60c05261025860e05261070861010052610e1061012052611c2061014052613840610160526170806101805261e1006101a052620151806101c0526202a3006101e052620546006102005262093a80610220526200008790600390600e620000a1565b50600f6005553480156200009a57600080fd5b506200016b565b600283019183908215620001325791602002820160005b83821115620000fe57835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302620000b8565b8015620001305782816101000a81549063ffffffff0219169055600401602081600301049283019260010302620000fe565b505b506200014092915062000144565b5090565b6200016891905b808211156200014057805463ffffffff191681556001016200014b565b90565b611202806200017b6000396000f30060806040526004361061017f5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a781146101845780630519ce79146101cf5780630560ff441461020057806306fdde0314610299578063095ea7b3146102ae5780630a0f8168146102d457806318160ddd146102e957806321717ebf1461031057806323b872dd1461032557806327d7874c1461034f5780632ba73c15146103705780633f4ba83a1461039157806346116e6f146103a6578063481af3d3146103be5780634e0a3379146103d65780635663896e146103f75780635c975abb1461040f5780636352211e1461042457806370a082311461043c5780637a7d49371461045d5780638456cb59146104725780638462151c1461048757806395d89b41146104f85780639d6fac6f1461050d578063a45f4bfc1461053e578063a9059cbb14610556578063b047fb501461057a578063bc4006f51461058f578063e17b25af146105a4578063e6cbe351146105c5575b600080fd5b34801561019057600080fd5b506101bb7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19600435166105da565b604080519115158252519081900360200190f35b3480156101db57600080fd5b506101e461086d565b60408051600160a060020a039092168252519081900360200190f35b34801561020c57600080fd5b5061022460048035906024803590810191013561087c565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561025e578181015183820152602001610246565b50505050905090810190601f16801561028b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102a557600080fd5b5061022461097f565b3480156102ba57600080fd5b506102d2600160a060020a03600435166024356109b6565b005b3480156102e057600080fd5b506101e4610a38565b3480156102f557600080fd5b506102fe610a47565b60408051918252519081900360200190f35b34801561031c57600080fd5b506101e4610a51565b34801561033157600080fd5b506102d2600160a060020a0360043581169060243516604435610a60565b34801561035b57600080fd5b506102d2600160a060020a0360043516610adc565b34801561037c57600080fd5b506102d2600160a060020a0360043516610b2a565b34801561039d57600080fd5b506102d2610b78565b3480156103b257600080fd5b506101e4600435610bc7565b3480156103ca57600080fd5b506101e4600435610be2565b3480156103e257600080fd5b506102d2600160a060020a0360043516610bfd565b34801561040357600080fd5b506102d2600435610c4b565b34801561041b57600080fd5b506101bb610ca7565b34801561043057600080fd5b506101e4600435610cb7565b34801561044857600080fd5b506102fe600160a060020a0360043516610cdb565b34801561046957600080fd5b506102fe610cf6565b34801561047e57600080fd5b506102d2610cfc565b34801561049357600080fd5b506104a8600160a060020a0360043516610d7c565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156104e45781810151838201526020016104cc565b505050509050019250505060405180910390f35b34801561050457600080fd5b50610224610e4e565b34801561051957600080fd5b50610525600435610e85565b6040805163ffffffff9092168252519081900360200190f35b34801561054a57600080fd5b506101e4600435610eb2565b34801561056257600080fd5b506102d2600160a060020a0360043516602435610ecd565b34801561058657600080fd5b506101e4610f69565b34801561059b57600080fd5b506101e4610f78565b3480156105b057600080fd5b506102d2600160a060020a0360043516610f87565b3480156105d157600080fd5b506101e4610fc0565b604080517f737570706f727473496e74657266616365286279746573342900000000000000815290519081900360190190206000907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19838116911614806108655750604080517f746f6b656e4d657461646174612875696e743235362c737472696e67290000008152815190819003601d0181207f746f6b656e734f664f776e657228616464726573732900000000000000000000825282519182900360160182207f7472616e7366657246726f6d28616464726573732c616464726573732c75696e83527f7432353629000000000000000000000000000000000000000000000000000000602084015283519283900360250183207f7472616e7366657228616464726573732c75696e743235362900000000000000845284519384900360190184207f617070726f766528616464726573732c75696e74323536290000000000000000855285519485900360180185207f6f776e65724f662875696e743235362900000000000000000000000000000000865286519586900360100186207f62616c616e63654f662861646472657373290000000000000000000000000000875287519687900360120187207f746f74616c537570706c792829000000000000000000000000000000000000008852885197889003600d0188207f73796d626f6c2829000000000000000000000000000000000000000000000000895289519889900360080189207f6e616d65282900000000000000000000000000000000000000000000000000008a529951988990036006019098207bffffffffffffffffffffffffffffffffffffffffffffffffffffffff198c811691909a189098181818181818181891909116145b90505b919050565b600154600160a060020a031681565b60606108866111b7565b600d54600090600160a060020a031615156108a057600080fd5b600d54604080517fcb4799f2000000000000000000000000000000000000000000000000000000008152600481018981526024820192835260448201889052600160a060020a039093169263cb4799f2928a928a928a929091606401848480828437820191505094505050505060a060405180830381600087803b15801561092757600080fd5b505af115801561093b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060a081101561096057600080fd5b50608081015190925090506109758282610fcf565b9695505050505050565b60408051808201909152600d81527f43727970746f4b69747469657300000000000000000000000000000000000000602082015281565b60025460a060020a900460ff16156109cd57600080fd5b6109d73382611023565b15156109e257600080fd5b6109ec8183611043565b60408051338152600160a060020a038416602082015280820183905290517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360600190a15050565b600054600160a060020a031681565b6006546000190190565b600c54600160a060020a031681565b60025460a060020a900460ff1615610a7757600080fd5b600160a060020a0382161515610a8c57600080fd5b600160a060020a038216301415610aa257600080fd5b610aac3382611071565b1515610ab757600080fd5b610ac18382611023565b1515610acc57600080fd5b610ad7838383611091565b505050565b600054600160a060020a03163314610af357600080fd5b600160a060020a0381161515610b0857600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314610b4157600080fd5b600160a060020a0381161515610b5657600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314610b8f57600080fd5b60025460a060020a900460ff161515610ba757600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600a60205260009081526040902054600160a060020a031681565b600960205260009081526040902054600160a060020a031681565b600054600160a060020a03163314610c1457600080fd5b600160a060020a0381161515610c2957600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b600254600160a060020a0316331480610c6e5750600054600160a060020a031633145b80610c835750600154600160a060020a031633145b1515610c8e57600080fd5b60035463ffffffff168110610ca257600080fd5b600555565b60025460a060020a900460ff1681565b600081815260076020526040902054600160a060020a031680151561086857600080fd5b600160a060020a031660009081526008602052604090205490565b60055481565b600254600160a060020a0316331480610d1f5750600054600160a060020a031633145b80610d345750600154600160a060020a031633145b1515610d3f57600080fd5b60025460a060020a900460ff1615610d5657600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a179055565b6060600060606000806000610d9087610cdb565b9450841515610daf576040805160008152602081019091529550610e44565b84604051908082528060200260200182016040528015610dd9578160200160208202803883390190505b509350610de4610a47565b925060009150600190505b828111610e4057600081815260076020526040902054600160a060020a0388811691161415610e3857808483815181101515610e2757fe5b602090810290910101526001909101905b600101610def565b8395505b5050505050919050565b60408051808201909152600281527f434b000000000000000000000000000000000000000000000000000000000000602082015281565b600381600e8110610e9257fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b600760205260009081526040902054600160a060020a031681565b60025460a060020a900460ff1615610ee457600080fd5b600160a060020a0382161515610ef957600080fd5b600160a060020a038216301415610f0f57600080fd5b600b54600160a060020a0383811691161415610f2a57600080fd5b600c54600160a060020a0383811691161415610f4557600080fd5b610f4f3382611023565b1515610f5a57600080fd5b610f65338383611091565b5050565b600254600160a060020a031681565b600d54600160a060020a031681565b600054600160a060020a03163314610f9e57600080fd5b600d8054600160a060020a031916600160a060020a0392909216919091179055565b600b54600160a060020a031681565b606080600080846040519080825280601f01601f191660200182016040528015611003578160200160208202803883390190505b5092505060208201905084611019828287611173565b5090949350505050565b600090815260076020526040902054600160a060020a0391821691161490565b6000918252600960205260409091208054600160a060020a031916600160a060020a03909216919091179055565b600090815260096020526040902054600160a060020a0391821691161490565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a031916909117905583161561112457600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b60408051600160a060020a0380861682528416602082015280820183905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a1505050565b60005b60208210611198578251845260209384019390920191601f1990910190611176565b50905182516020929092036101000a6000190180199091169116179052565b60806040519081016040528060049060208202803883395091929150505600a165627a7a723058209af868021e88d693a86326a394ed134db171018037302602393d32e0c1896d980029`

// DeployKittyOwnership deploys a new Ethereum contract, binding an instance of KittyOwnership to it.
func DeployKittyOwnership(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyOwnership, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyOwnershipABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyOwnershipBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyOwnership{KittyOwnershipCaller: KittyOwnershipCaller{contract: contract}, KittyOwnershipTransactor: KittyOwnershipTransactor{contract: contract}, KittyOwnershipFilterer: KittyOwnershipFilterer{contract: contract}}, nil
}

// KittyOwnership is an auto generated Go binding around an Ethereum contract.
type KittyOwnership struct {
	KittyOwnershipCaller     // Read-only binding to the contract
	KittyOwnershipTransactor // Write-only binding to the contract
	KittyOwnershipFilterer   // Log filterer for contract events
}

// KittyOwnershipCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyOwnershipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyOwnershipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyOwnershipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyOwnershipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyOwnershipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyOwnershipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyOwnershipSession struct {
	Contract     *KittyOwnership   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittyOwnershipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyOwnershipCallerSession struct {
	Contract *KittyOwnershipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// KittyOwnershipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyOwnershipTransactorSession struct {
	Contract     *KittyOwnershipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// KittyOwnershipRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyOwnershipRaw struct {
	Contract *KittyOwnership // Generic contract binding to access the raw methods on
}

// KittyOwnershipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyOwnershipCallerRaw struct {
	Contract *KittyOwnershipCaller // Generic read-only contract binding to access the raw methods on
}

// KittyOwnershipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyOwnershipTransactorRaw struct {
	Contract *KittyOwnershipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyOwnership creates a new instance of KittyOwnership, bound to a specific deployed contract.
func NewKittyOwnership(address common.Address, backend bind.ContractBackend) (*KittyOwnership, error) {
	contract, err := bindKittyOwnership(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyOwnership{KittyOwnershipCaller: KittyOwnershipCaller{contract: contract}, KittyOwnershipTransactor: KittyOwnershipTransactor{contract: contract}, KittyOwnershipFilterer: KittyOwnershipFilterer{contract: contract}}, nil
}

// NewKittyOwnershipCaller creates a new read-only instance of KittyOwnership, bound to a specific deployed contract.
func NewKittyOwnershipCaller(address common.Address, caller bind.ContractCaller) (*KittyOwnershipCaller, error) {
	contract, err := bindKittyOwnership(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipCaller{contract: contract}, nil
}

// NewKittyOwnershipTransactor creates a new write-only instance of KittyOwnership, bound to a specific deployed contract.
func NewKittyOwnershipTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyOwnershipTransactor, error) {
	contract, err := bindKittyOwnership(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipTransactor{contract: contract}, nil
}

// NewKittyOwnershipFilterer creates a new log filterer instance of KittyOwnership, bound to a specific deployed contract.
func NewKittyOwnershipFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyOwnershipFilterer, error) {
	contract, err := bindKittyOwnership(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipFilterer{contract: contract}, nil
}

// bindKittyOwnership binds a generic wrapper to an already deployed contract.
func bindKittyOwnership(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyOwnershipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyOwnership *KittyOwnershipRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyOwnership.Contract.KittyOwnershipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyOwnership *KittyOwnershipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyOwnership.Contract.KittyOwnershipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyOwnership *KittyOwnershipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyOwnership.Contract.KittyOwnershipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyOwnership *KittyOwnershipCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KittyOwnership.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyOwnership *KittyOwnershipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyOwnership.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyOwnership *KittyOwnershipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyOwnership.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyOwnership *KittyOwnershipCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyOwnership *KittyOwnershipSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyOwnership.Contract.BalanceOf(&_KittyOwnership.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_KittyOwnership *KittyOwnershipCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyOwnership.Contract.BalanceOf(&_KittyOwnership.CallOpts, _owner)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyOwnership *KittyOwnershipCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyOwnership *KittyOwnershipSession) CeoAddress() (common.Address, error) {
	return _KittyOwnership.Contract.CeoAddress(&_KittyOwnership.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) CeoAddress() (common.Address, error) {
	return _KittyOwnership.Contract.CeoAddress(&_KittyOwnership.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyOwnership *KittyOwnershipCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyOwnership *KittyOwnershipSession) CfoAddress() (common.Address, error) {
	return _KittyOwnership.Contract.CfoAddress(&_KittyOwnership.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) CfoAddress() (common.Address, error) {
	return _KittyOwnership.Contract.CfoAddress(&_KittyOwnership.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyOwnership *KittyOwnershipCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyOwnership *KittyOwnershipSession) CooAddress() (common.Address, error) {
	return _KittyOwnership.Contract.CooAddress(&_KittyOwnership.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) CooAddress() (common.Address, error) {
	return _KittyOwnership.Contract.CooAddress(&_KittyOwnership.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyOwnership *KittyOwnershipCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "cooldowns", arg0)
	return *ret0, err
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyOwnership *KittyOwnershipSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyOwnership.Contract.Cooldowns(&_KittyOwnership.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_KittyOwnership *KittyOwnershipCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyOwnership.Contract.Cooldowns(&_KittyOwnership.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyOwnership *KittyOwnershipCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "erc721Metadata")
	return *ret0, err
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyOwnership *KittyOwnershipSession) Erc721Metadata() (common.Address, error) {
	return _KittyOwnership.Contract.Erc721Metadata(&_KittyOwnership.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) Erc721Metadata() (common.Address, error) {
	return _KittyOwnership.Contract.Erc721Metadata(&_KittyOwnership.CallOpts)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyOwnership *KittyOwnershipCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "kittyIndexToApproved", arg0)
	return *ret0, err
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyOwnership *KittyOwnershipSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.KittyIndexToApproved(&_KittyOwnership.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.KittyIndexToApproved(&_KittyOwnership.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyOwnership *KittyOwnershipCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "kittyIndexToOwner", arg0)
	return *ret0, err
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyOwnership *KittyOwnershipSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.KittyIndexToOwner(&_KittyOwnership.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.KittyIndexToOwner(&_KittyOwnership.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyOwnership *KittyOwnershipCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyOwnership *KittyOwnershipSession) Name() (string, error) {
	return _KittyOwnership.Contract.Name(&_KittyOwnership.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KittyOwnership *KittyOwnershipCallerSession) Name() (string, error) {
	return _KittyOwnership.Contract.Name(&_KittyOwnership.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyOwnership *KittyOwnershipCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyOwnership *KittyOwnershipSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.OwnerOf(&_KittyOwnership.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_KittyOwnership *KittyOwnershipCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.OwnerOf(&_KittyOwnership.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyOwnership *KittyOwnershipCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyOwnership *KittyOwnershipSession) Paused() (bool, error) {
	return _KittyOwnership.Contract.Paused(&_KittyOwnership.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_KittyOwnership *KittyOwnershipCallerSession) Paused() (bool, error) {
	return _KittyOwnership.Contract.Paused(&_KittyOwnership.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyOwnership *KittyOwnershipCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "saleAuction")
	return *ret0, err
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyOwnership *KittyOwnershipSession) SaleAuction() (common.Address, error) {
	return _KittyOwnership.Contract.SaleAuction(&_KittyOwnership.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) SaleAuction() (common.Address, error) {
	return _KittyOwnership.Contract.SaleAuction(&_KittyOwnership.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyOwnership *KittyOwnershipCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "secondsPerBlock")
	return *ret0, err
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyOwnership *KittyOwnershipSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyOwnership.Contract.SecondsPerBlock(&_KittyOwnership.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_KittyOwnership *KittyOwnershipCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyOwnership.Contract.SecondsPerBlock(&_KittyOwnership.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyOwnership *KittyOwnershipCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "sireAllowedToAddress", arg0)
	return *ret0, err
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyOwnership *KittyOwnershipSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.SireAllowedToAddress(&_KittyOwnership.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.SireAllowedToAddress(&_KittyOwnership.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyOwnership *KittyOwnershipCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "siringAuction")
	return *ret0, err
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyOwnership *KittyOwnershipSession) SiringAuction() (common.Address, error) {
	return _KittyOwnership.Contract.SiringAuction(&_KittyOwnership.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) SiringAuction() (common.Address, error) {
	return _KittyOwnership.Contract.SiringAuction(&_KittyOwnership.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyOwnership *KittyOwnershipCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "supportsInterface", _interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyOwnership *KittyOwnershipSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyOwnership.Contract.SupportsInterface(&_KittyOwnership.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_KittyOwnership *KittyOwnershipCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyOwnership.Contract.SupportsInterface(&_KittyOwnership.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyOwnership *KittyOwnershipCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyOwnership *KittyOwnershipSession) Symbol() (string, error) {
	return _KittyOwnership.Contract.Symbol(&_KittyOwnership.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KittyOwnership *KittyOwnershipCallerSession) Symbol() (string, error) {
	return _KittyOwnership.Contract.Symbol(&_KittyOwnership.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyOwnership *KittyOwnershipCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "tokenMetadata", _tokenId, _preferredTransport)
	return *ret0, err
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyOwnership *KittyOwnershipSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyOwnership.Contract.TokenMetadata(&_KittyOwnership.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_KittyOwnership *KittyOwnershipCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyOwnership.Contract.TokenMetadata(&_KittyOwnership.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyOwnership *KittyOwnershipCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyOwnership *KittyOwnershipSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyOwnership.Contract.TokensOfOwner(&_KittyOwnership.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_KittyOwnership *KittyOwnershipCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyOwnership.Contract.TokensOfOwner(&_KittyOwnership.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyOwnership *KittyOwnershipCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KittyOwnership.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyOwnership *KittyOwnershipSession) TotalSupply() (*big.Int, error) {
	return _KittyOwnership.Contract.TotalSupply(&_KittyOwnership.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KittyOwnership *KittyOwnershipCallerSession) TotalSupply() (*big.Int, error) {
	return _KittyOwnership.Contract.TotalSupply(&_KittyOwnership.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyOwnership *KittyOwnershipTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyOwnership *KittyOwnershipSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.Approve(&_KittyOwnership.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.Approve(&_KittyOwnership.TransactOpts, _to, _tokenId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyOwnership *KittyOwnershipTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyOwnership *KittyOwnershipSession) Pause() (*types.Transaction, error) {
	return _KittyOwnership.Contract.Pause(&_KittyOwnership.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyOwnership.Contract.Pause(&_KittyOwnership.TransactOpts)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyOwnership *KittyOwnershipTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyOwnership *KittyOwnershipSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetCEO(&_KittyOwnership.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetCEO(&_KittyOwnership.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyOwnership *KittyOwnershipTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyOwnership *KittyOwnershipSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetCFO(&_KittyOwnership.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetCFO(&_KittyOwnership.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyOwnership *KittyOwnershipTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyOwnership *KittyOwnershipSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetCOO(&_KittyOwnership.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetCOO(&_KittyOwnership.TransactOpts, _newCOO)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyOwnership *KittyOwnershipTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyOwnership *KittyOwnershipSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetMetadataAddress(&_KittyOwnership.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetMetadataAddress(&_KittyOwnership.TransactOpts, _contractAddress)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyOwnership *KittyOwnershipTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyOwnership *KittyOwnershipSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetSecondsPerBlock(&_KittyOwnership.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetSecondsPerBlock(&_KittyOwnership.TransactOpts, secs)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyOwnership *KittyOwnershipTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyOwnership *KittyOwnershipSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.Transfer(&_KittyOwnership.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.Transfer(&_KittyOwnership.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyOwnership *KittyOwnershipTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyOwnership *KittyOwnershipSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.TransferFrom(&_KittyOwnership.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.TransferFrom(&_KittyOwnership.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyOwnership *KittyOwnershipTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyOwnership *KittyOwnershipSession) Unpause() (*types.Transaction, error) {
	return _KittyOwnership.Contract.Unpause(&_KittyOwnership.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyOwnership.Contract.Unpause(&_KittyOwnership.TransactOpts)
}

// KittyOwnershipApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KittyOwnership contract.
type KittyOwnershipApprovalIterator struct {
	Event *KittyOwnershipApproval // Event containing the contract specifics and raw log

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
func (it *KittyOwnershipApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyOwnershipApproval)
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
		it.Event = new(KittyOwnershipApproval)
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
func (it *KittyOwnershipApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyOwnershipApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyOwnershipApproval represents a Approval event raised by the KittyOwnership contract.
type KittyOwnershipApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_KittyOwnership *KittyOwnershipFilterer) FilterApproval(opts *bind.FilterOpts) (*KittyOwnershipApprovalIterator, error) {

	logs, sub, err := _KittyOwnership.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipApprovalIterator{contract: _KittyOwnership.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_KittyOwnership *KittyOwnershipFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KittyOwnershipApproval) (event.Subscription, error) {

	logs, sub, err := _KittyOwnership.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyOwnershipApproval)
				if err := _KittyOwnership.contract.UnpackLog(event, "Approval", log); err != nil {
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

// KittyOwnershipBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the KittyOwnership contract.
type KittyOwnershipBirthIterator struct {
	Event *KittyOwnershipBirth // Event containing the contract specifics and raw log

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
func (it *KittyOwnershipBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyOwnershipBirth)
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
		it.Event = new(KittyOwnershipBirth)
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
func (it *KittyOwnershipBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyOwnershipBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyOwnershipBirth represents a Birth event raised by the KittyOwnership contract.
type KittyOwnershipBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_KittyOwnership *KittyOwnershipFilterer) FilterBirth(opts *bind.FilterOpts) (*KittyOwnershipBirthIterator, error) {

	logs, sub, err := _KittyOwnership.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipBirthIterator{contract: _KittyOwnership.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_KittyOwnership *KittyOwnershipFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *KittyOwnershipBirth) (event.Subscription, error) {

	logs, sub, err := _KittyOwnership.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyOwnershipBirth)
				if err := _KittyOwnership.contract.UnpackLog(event, "Birth", log); err != nil {
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

// KittyOwnershipContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyOwnership contract.
type KittyOwnershipContractUpgradeIterator struct {
	Event *KittyOwnershipContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyOwnershipContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyOwnershipContractUpgrade)
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
		it.Event = new(KittyOwnershipContractUpgrade)
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
func (it *KittyOwnershipContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyOwnershipContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyOwnershipContractUpgrade represents a ContractUpgrade event raised by the KittyOwnership contract.
type KittyOwnershipContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyOwnership *KittyOwnershipFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyOwnershipContractUpgradeIterator, error) {

	logs, sub, err := _KittyOwnership.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipContractUpgradeIterator{contract: _KittyOwnership.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_KittyOwnership *KittyOwnershipFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyOwnershipContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyOwnership.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyOwnershipContractUpgrade)
				if err := _KittyOwnership.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// KittyOwnershipTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KittyOwnership contract.
type KittyOwnershipTransferIterator struct {
	Event *KittyOwnershipTransfer // Event containing the contract specifics and raw log

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
func (it *KittyOwnershipTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyOwnershipTransfer)
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
		it.Event = new(KittyOwnershipTransfer)
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
func (it *KittyOwnershipTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyOwnershipTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyOwnershipTransfer represents a Transfer event raised by the KittyOwnership contract.
type KittyOwnershipTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_KittyOwnership *KittyOwnershipFilterer) FilterTransfer(opts *bind.FilterOpts) (*KittyOwnershipTransferIterator, error) {

	logs, sub, err := _KittyOwnership.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipTransferIterator{contract: _KittyOwnership.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_KittyOwnership *KittyOwnershipFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KittyOwnershipTransfer) (event.Subscription, error) {

	logs, sub, err := _KittyOwnership.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyOwnershipTransfer)
				if err := _KittyOwnership.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a0319163317905561017f806100326000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b1461008e575b600080fd5b34801561005c57600080fd5b506100656100be565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561009a57600080fd5b506100bc73ffffffffffffffffffffffffffffffffffffffff600435166100da565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff1633146100fe57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615610150576000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b505600a165627a7a72305820eb3fda94c72fbc71524dbc2fa6f5b752a395570596ef5ae24d8063eca7488b7a0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// PausableBin is the compiled bytecode used for deploying new contracts.
const PausableBin = `0x608060405260008054600160a860020a031916331790556102f4806100256000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f4ba83a81146100715780635c975abb1461009a5780638456cb59146100af5780638da5cb5b146100c4578063f2fde38b146100f5575b600080fd5b34801561007d57600080fd5b50610086610118565b604080519115158252519081900360200190f35b3480156100a657600080fd5b506100866101a4565b3480156100bb57600080fd5b506100866101c5565b3480156100d057600080fd5b506100d9610267565b60408051600160a060020a039092168252519081900360200190f35b34801561010157600080fd5b50610116600160a060020a0360043516610276565b005b60008054600160a060020a0316331461013057600080fd5b60005474010000000000000000000000000000000000000000900460ff16151561015957600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a150600190565b60005474010000000000000000000000000000000000000000900460ff1681565b60008054600160a060020a031633146101dd57600080fd5b60005474010000000000000000000000000000000000000000900460ff161561020557600080fd5b6000805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a150600190565b600054600160a060020a031681565b600054600160a060020a0316331461028d57600080fd5b600160a060020a038116156102c5576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b505600a165627a7a723058202f309e438dd1eacb85cd366c3bd406288da146f8867563d5c484440e44b248c60029`

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCallerSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// PausablePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Pausable contract.
type PausablePauseIterator struct {
	Event *PausablePause // Event containing the contract specifics and raw log

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
func (it *PausablePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePause)
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
		it.Event = new(PausablePause)
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
func (it *PausablePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePause represents a Pause event raised by the Pausable contract.
type PausablePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) FilterPause(opts *bind.FilterOpts) (*PausablePauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausablePauseIterator{contract: _Pausable.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausablePause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePause)
				if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PausableUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Pausable contract.
type PausableUnpauseIterator struct {
	Event *PausableUnpause // Event containing the contract specifics and raw log

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
func (it *PausableUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpause)
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
		it.Event = new(PausableUnpause)
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
func (it *PausableUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpause represents a Unpause event raised by the Pausable contract.
type PausableUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableUnpauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableUnpauseIterator{contract: _Pausable.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableUnpause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpause)
				if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// SaleClockAuctionABI is the input ABI used to generate the binding from.
const SaleClockAuctionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastGen0SalePrices\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"startedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSaleClockAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuctionWhenPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gen0SaleCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"averageGen0SalePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_nftAddr\",\"type\":\"address\"},{\"name\":\"_cut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// SaleClockAuctionBin is the compiled bytecode used for deploying new contracts.
const SaleClockAuctionBin = `0x60806040526000805460a060020a60ff02191690556004805460ff1916600117905534801561002d57600080fd5b50604051604080610ee883398101604052805160209091015160008054600160a060020a031916331781558290829061271082111561006b57600080fd5b506002819055604080517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f9a20483d00000000000000000000000000000000000000000000000000000000600482015290518391600160a060020a038316916301ffc9a7916024808201926020929091908290030181600087803b1580156100f557600080fd5b505af1158015610109573d6000803e3d6000fd5b505050506040513d602081101561011f57600080fd5b5051151561012c57600080fd5b60018054600160a060020a03909216600160a060020a031990921691909117905550505050610d88806101606000396000f3006080604052600436106100fb5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166327ebe40a81146101005780633f4ba83a1461012f578063454a2ab314610158578063484eccb4146101635780635c975abb1461018d5780635fd8c710146101a257806378bd7935146101b757806383b5ff8b146102045780638456cb591461021957806385b861881461022e578063878eb368146102435780638a98a9cc1461025b5780638da5cb5b1461027057806396b5a755146102a1578063c55d0f56146102b9578063dd1b7a0f146102d1578063eac9d94c146102e6578063f2fde38b146102fb575b600080fd5b34801561010c57600080fd5b5061012d600435602435604435606435600160a060020a036084351661031c565b005b34801561013b57600080fd5b50610144610414565b604080519115158252519081900360200190f35b61012d60043561048f565b34801561016f57600080fd5b5061017b6004356104f9565b60408051918252519081900360200190f35b34801561019957600080fd5b5061014461050d565b3480156101ae57600080fd5b5061012d61051d565b3480156101c357600080fd5b506101cf60043561057a565b60408051600160a060020a03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b34801561021057600080fd5b5061017b610610565b34801561022557600080fd5b50610144610616565b34801561023a57600080fd5b50610144610696565b34801561024f57600080fd5b5061012d60043561069f565b34801561026757600080fd5b5061017b61070c565b34801561027c57600080fd5b50610285610712565b60408051600160a060020a039092168252519081900360200190f35b3480156102ad57600080fd5b5061012d600435610721565b3480156102c557600080fd5b5061017b600435610766565b3480156102dd57600080fd5b50610285610798565b3480156102f257600080fd5b5061017b6107a7565b34801561030757600080fd5b5061012d600160a060020a03600435166107db565b610324610d2e565b6fffffffffffffffffffffffffffffffff8516851461034257600080fd5b6fffffffffffffffffffffffffffffffff8416841461036057600080fd5b67ffffffffffffffff8316831461037657600080fd5b600154600160a060020a0316331461038d57600080fd5b610397828761082d565b60a06040519081016040528083600160a060020a03168152602001866fffffffffffffffffffffffffffffffff168152602001856fffffffffffffffffffffffffffffffff1681526020018467ffffffffffffffff1681526020014267ffffffffffffffff16815250905061040c86826108b5565b505050505050565b60008054600160a060020a0316331461042c57600080fd5b60005460a060020a900460ff16151561044457600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a150600190565b600081815260036020526040812054600160a060020a0316906104b28334610a09565b90506104be3384610b2f565b600154600160a060020a03838116911614156104f45760058054829160069106600581106104e857fe5b01556005805460010190555b505050565b6006816005811061050657fe5b0154905081565b60005460a060020a900460ff1681565b60015460008054600160a060020a039283169216331480610546575033600160a060020a038316145b151561055157600080fd5b604051600160a060020a03831690303180156108fc02916000818181858888f150505050505050565b6000818152600360205260408120819081908190819061059981610b9d565b15156105a457600080fd5b80546001820154600290920154600160a060020a03909116986fffffffffffffffffffffffffffffffff8084169950700100000000000000000000000000000000909304909216965067ffffffffffffffff808216965068010000000000000000909104169350915050565b60025481565b60008054600160a060020a0316331461062e57600080fd5b60005460a060020a900460ff161561064557600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a150600190565b60045460ff1681565b6000805460a060020a900460ff1615156106b857600080fd5b600054600160a060020a031633146106cf57600080fd5b5060008181526003602052604090206106e781610b9d565b15156106f257600080fd5b8054610708908390600160a060020a0316610bbe565b5050565b60055481565b600054600160a060020a031681565b60008181526003602052604081209061073982610b9d565b151561074457600080fd5b508054600160a060020a031633811461075c57600080fd5b6104f48382610bbe565b600081815260036020526040812061077d81610b9d565b151561078857600080fd5b61079181610c08565b9392505050565b600154600160a060020a031681565b600080805b60058110156107d157600681600581106107c257fe5b015491909101906001016107ac565b5060059004919050565b600054600160a060020a031633146107f257600080fd5b600160a060020a0381161561082a576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015230602483015260448201859052915191909216916323b872dd91606480830192600092919082900301818387803b1580156108a157600080fd5b505af115801561040c573d6000803e3d6000fd5b603c816060015167ffffffffffffffff16101515156108d357600080fd5b60008281526003602090815260409182902083518154600160a060020a0390911673ffffffffffffffffffffffffffffffffffffffff1990911617815581840151600182018054858701516fffffffffffffffffffffffffffffffff90811670010000000000000000000000000000000081029482166fffffffffffffffffffffffffffffffff19909316831790911693909317909155606080870151600290940180546080808a015167ffffffffffffffff90811668010000000000000000026fffffffffffffffff0000000000000000199190981667ffffffffffffffff1990931683171696909617909155865189815295860192909252848601929092529083015291517fa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7929181900390910190a15050565b60008281526003602052604081208180808080610a2586610b9d565b1515610a3057600080fd5b610a3986610c08565b945084881015610a4857600080fd5b8554600160a060020a03169350610a5e89610c98565b6000851115610ab057610a7085610ce5565b6040519093508386039250600160a060020a0385169083156108fc029084906000818181858888f19350505050158015610aae573d6000803e3d6000fd5b505b5060405184880390339082156108fc029083906000818181858888f19350505050158015610ae2573d6000803e3d6000fd5b50604080518a815260208101879052338183015290517f4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd29181900360600190a15092979650505050505050565b600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038581166004830152602482018590529151919092169163a9059cbb91604480830192600092919082900301818387803b1580156108a157600080fd5b6002015460006801000000000000000090910467ffffffffffffffff161190565b610bc782610c98565b610bd18183610b2f565b6040805183815290517f2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df9181900360200190a15050565b6002810154600090819068010000000000000000900467ffffffffffffffff16421115610c4e5750600282015468010000000000000000900467ffffffffffffffff1642035b60018301546002840154610791916fffffffffffffffffffffffffffffffff80821692700100000000000000000000000000000000909204169067ffffffffffffffff1684610cf1565b6000908152600360205260408120805473ffffffffffffffffffffffffffffffffffffffff19168155600181019190915560020180546fffffffffffffffffffffffffffffffff19169055565b60025461271091020490565b6000808080858510610d0557869350610d23565b878703925085858402811515610d1757fe5b05915081880190508093505b505050949350505050565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152905600a165627a7a72305820dbcffead7a92f0b00dc6814f66433946804db25f10d746465e3e4f0d5f7c93c70029`

// DeploySaleClockAuction deploys a new Ethereum contract, binding an instance of SaleClockAuction to it.
func DeploySaleClockAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _nftAddr common.Address, _cut *big.Int) (common.Address, *types.Transaction, *SaleClockAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleClockAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SaleClockAuctionBin), backend, _nftAddr, _cut)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SaleClockAuction{SaleClockAuctionCaller: SaleClockAuctionCaller{contract: contract}, SaleClockAuctionTransactor: SaleClockAuctionTransactor{contract: contract}, SaleClockAuctionFilterer: SaleClockAuctionFilterer{contract: contract}}, nil
}

// SaleClockAuction is an auto generated Go binding around an Ethereum contract.
type SaleClockAuction struct {
	SaleClockAuctionCaller     // Read-only binding to the contract
	SaleClockAuctionTransactor // Write-only binding to the contract
	SaleClockAuctionFilterer   // Log filterer for contract events
}

// SaleClockAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaleClockAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleClockAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaleClockAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleClockAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaleClockAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleClockAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaleClockAuctionSession struct {
	Contract     *SaleClockAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleClockAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaleClockAuctionCallerSession struct {
	Contract *SaleClockAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SaleClockAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaleClockAuctionTransactorSession struct {
	Contract     *SaleClockAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SaleClockAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaleClockAuctionRaw struct {
	Contract *SaleClockAuction // Generic contract binding to access the raw methods on
}

// SaleClockAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaleClockAuctionCallerRaw struct {
	Contract *SaleClockAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SaleClockAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaleClockAuctionTransactorRaw struct {
	Contract *SaleClockAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSaleClockAuction creates a new instance of SaleClockAuction, bound to a specific deployed contract.
func NewSaleClockAuction(address common.Address, backend bind.ContractBackend) (*SaleClockAuction, error) {
	contract, err := bindSaleClockAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SaleClockAuction{SaleClockAuctionCaller: SaleClockAuctionCaller{contract: contract}, SaleClockAuctionTransactor: SaleClockAuctionTransactor{contract: contract}, SaleClockAuctionFilterer: SaleClockAuctionFilterer{contract: contract}}, nil
}

// NewSaleClockAuctionCaller creates a new read-only instance of SaleClockAuction, bound to a specific deployed contract.
func NewSaleClockAuctionCaller(address common.Address, caller bind.ContractCaller) (*SaleClockAuctionCaller, error) {
	contract, err := bindSaleClockAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionCaller{contract: contract}, nil
}

// NewSaleClockAuctionTransactor creates a new write-only instance of SaleClockAuction, bound to a specific deployed contract.
func NewSaleClockAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SaleClockAuctionTransactor, error) {
	contract, err := bindSaleClockAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionTransactor{contract: contract}, nil
}

// NewSaleClockAuctionFilterer creates a new log filterer instance of SaleClockAuction, bound to a specific deployed contract.
func NewSaleClockAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SaleClockAuctionFilterer, error) {
	contract, err := bindSaleClockAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionFilterer{contract: contract}, nil
}

// bindSaleClockAuction binds a generic wrapper to an already deployed contract.
func bindSaleClockAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleClockAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaleClockAuction *SaleClockAuctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SaleClockAuction.Contract.SaleClockAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaleClockAuction *SaleClockAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.SaleClockAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaleClockAuction *SaleClockAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.SaleClockAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaleClockAuction *SaleClockAuctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SaleClockAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaleClockAuction *SaleClockAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaleClockAuction *SaleClockAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.contract.Transact(opts, method, params...)
}

// AverageGen0SalePrice is a free data retrieval call binding the contract method 0xeac9d94c.
//
// Solidity: function averageGen0SalePrice() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) AverageGen0SalePrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "averageGen0SalePrice")
	return *ret0, err
}

// AverageGen0SalePrice is a free data retrieval call binding the contract method 0xeac9d94c.
//
// Solidity: function averageGen0SalePrice() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) AverageGen0SalePrice() (*big.Int, error) {
	return _SaleClockAuction.Contract.AverageGen0SalePrice(&_SaleClockAuction.CallOpts)
}

// AverageGen0SalePrice is a free data retrieval call binding the contract method 0xeac9d94c.
//
// Solidity: function averageGen0SalePrice() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) AverageGen0SalePrice() (*big.Int, error) {
	return _SaleClockAuction.Contract.AverageGen0SalePrice(&_SaleClockAuction.CallOpts)
}

// Gen0SaleCount is a free data retrieval call binding the contract method 0x8a98a9cc.
//
// Solidity: function gen0SaleCount() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) Gen0SaleCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "gen0SaleCount")
	return *ret0, err
}

// Gen0SaleCount is a free data retrieval call binding the contract method 0x8a98a9cc.
//
// Solidity: function gen0SaleCount() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) Gen0SaleCount() (*big.Int, error) {
	return _SaleClockAuction.Contract.Gen0SaleCount(&_SaleClockAuction.CallOpts)
}

// Gen0SaleCount is a free data retrieval call binding the contract method 0x8a98a9cc.
//
// Solidity: function gen0SaleCount() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) Gen0SaleCount() (*big.Int, error) {
	return _SaleClockAuction.Contract.Gen0SaleCount(&_SaleClockAuction.CallOpts)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) GetAuction(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	ret := new(struct {
		Seller        common.Address
		StartingPrice *big.Int
		EndingPrice   *big.Int
		Duration      *big.Int
		StartedAt     *big.Int
	})
	out := ret
	err := _SaleClockAuction.contract.Call(opts, out, "getAuction", _tokenId)
	return *ret, err
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_SaleClockAuction *SaleClockAuctionSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _SaleClockAuction.Contract.GetAuction(&_SaleClockAuction.CallOpts, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _SaleClockAuction.Contract.GetAuction(&_SaleClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) GetCurrentPrice(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "getCurrentPrice", _tokenId)
	return *ret0, err
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _SaleClockAuction.Contract.GetCurrentPrice(&_SaleClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _SaleClockAuction.Contract.GetCurrentPrice(&_SaleClockAuction.CallOpts, _tokenId)
}

// IsSaleClockAuction is a free data retrieval call binding the contract method 0x85b86188.
//
// Solidity: function isSaleClockAuction() constant returns(bool)
func (_SaleClockAuction *SaleClockAuctionCaller) IsSaleClockAuction(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "isSaleClockAuction")
	return *ret0, err
}

// IsSaleClockAuction is a free data retrieval call binding the contract method 0x85b86188.
//
// Solidity: function isSaleClockAuction() constant returns(bool)
func (_SaleClockAuction *SaleClockAuctionSession) IsSaleClockAuction() (bool, error) {
	return _SaleClockAuction.Contract.IsSaleClockAuction(&_SaleClockAuction.CallOpts)
}

// IsSaleClockAuction is a free data retrieval call binding the contract method 0x85b86188.
//
// Solidity: function isSaleClockAuction() constant returns(bool)
func (_SaleClockAuction *SaleClockAuctionCallerSession) IsSaleClockAuction() (bool, error) {
	return _SaleClockAuction.Contract.IsSaleClockAuction(&_SaleClockAuction.CallOpts)
}

// LastGen0SalePrices is a free data retrieval call binding the contract method 0x484eccb4.
//
// Solidity: function lastGen0SalePrices( uint256) constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) LastGen0SalePrices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "lastGen0SalePrices", arg0)
	return *ret0, err
}

// LastGen0SalePrices is a free data retrieval call binding the contract method 0x484eccb4.
//
// Solidity: function lastGen0SalePrices( uint256) constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) LastGen0SalePrices(arg0 *big.Int) (*big.Int, error) {
	return _SaleClockAuction.Contract.LastGen0SalePrices(&_SaleClockAuction.CallOpts, arg0)
}

// LastGen0SalePrices is a free data retrieval call binding the contract method 0x484eccb4.
//
// Solidity: function lastGen0SalePrices( uint256) constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) LastGen0SalePrices(arg0 *big.Int) (*big.Int, error) {
	return _SaleClockAuction.Contract.LastGen0SalePrices(&_SaleClockAuction.CallOpts, arg0)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_SaleClockAuction *SaleClockAuctionCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "nonFungibleContract")
	return *ret0, err
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_SaleClockAuction *SaleClockAuctionSession) NonFungibleContract() (common.Address, error) {
	return _SaleClockAuction.Contract.NonFungibleContract(&_SaleClockAuction.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_SaleClockAuction *SaleClockAuctionCallerSession) NonFungibleContract() (common.Address, error) {
	return _SaleClockAuction.Contract.NonFungibleContract(&_SaleClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SaleClockAuction *SaleClockAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SaleClockAuction *SaleClockAuctionSession) Owner() (common.Address, error) {
	return _SaleClockAuction.Contract.Owner(&_SaleClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SaleClockAuction *SaleClockAuctionCallerSession) Owner() (common.Address, error) {
	return _SaleClockAuction.Contract.Owner(&_SaleClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "ownerCut")
	return *ret0, err
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) OwnerCut() (*big.Int, error) {
	return _SaleClockAuction.Contract.OwnerCut(&_SaleClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) OwnerCut() (*big.Int, error) {
	return _SaleClockAuction.Contract.OwnerCut(&_SaleClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_SaleClockAuction *SaleClockAuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_SaleClockAuction *SaleClockAuctionSession) Paused() (bool, error) {
	return _SaleClockAuction.Contract.Paused(&_SaleClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_SaleClockAuction *SaleClockAuctionCallerSession) Paused() (bool, error) {
	return _SaleClockAuction.Contract.Paused(&_SaleClockAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) Bid(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "bid", _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Bid(&_SaleClockAuction.TransactOpts, _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Bid(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "cancelAuction", _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CancelAuction(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CancelAuction(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) CancelAuctionWhenPaused(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "cancelAuctionWhenPaused", _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CancelAuctionWhenPaused(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CancelAuctionWhenPaused(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "createAuction", _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_SaleClockAuction *SaleClockAuctionSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CreateAuction(&_SaleClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CreateAuction(&_SaleClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionSession) Pause() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Pause(&_SaleClockAuction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionTransactorSession) Pause() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Pause(&_SaleClockAuction.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SaleClockAuction *SaleClockAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.TransferOwnership(&_SaleClockAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.TransferOwnership(&_SaleClockAuction.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionSession) Unpause() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Unpause(&_SaleClockAuction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionTransactorSession) Unpause() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Unpause(&_SaleClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SaleClockAuction *SaleClockAuctionSession) WithdrawBalance() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.WithdrawBalance(&_SaleClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.WithdrawBalance(&_SaleClockAuction.TransactOpts)
}

// SaleClockAuctionAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionCancelledIterator struct {
	Event *SaleClockAuctionAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionAuctionCancelled)
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
		it.Event = new(SaleClockAuctionAuctionCancelled)
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
func (it *SaleClockAuctionAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionAuctionCancelled represents a AuctionCancelled event raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*SaleClockAuctionAuctionCancelledIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionAuctionCancelledIterator{contract: _SaleClockAuction.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionAuctionCancelled)
				if err := _SaleClockAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// SaleClockAuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionCreatedIterator struct {
	Event *SaleClockAuctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionAuctionCreated)
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
		it.Event = new(SaleClockAuctionAuctionCreated)
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
func (it *SaleClockAuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionAuctionCreated represents a AuctionCreated event raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*SaleClockAuctionAuctionCreatedIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionAuctionCreatedIterator{contract: _SaleClockAuction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionAuctionCreated)
				if err := _SaleClockAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// SaleClockAuctionAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionSuccessfulIterator struct {
	Event *SaleClockAuctionAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionAuctionSuccessful)
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
		it.Event = new(SaleClockAuctionAuctionSuccessful)
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
func (it *SaleClockAuctionAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionAuctionSuccessful represents a AuctionSuccessful event raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*SaleClockAuctionAuctionSuccessfulIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionAuctionSuccessfulIterator{contract: _SaleClockAuction.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionAuctionSuccessful)
				if err := _SaleClockAuction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// SaleClockAuctionPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the SaleClockAuction contract.
type SaleClockAuctionPauseIterator struct {
	Event *SaleClockAuctionPause // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionPause)
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
		it.Event = new(SaleClockAuctionPause)
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
func (it *SaleClockAuctionPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionPause represents a Pause event raised by the SaleClockAuction contract.
type SaleClockAuctionPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterPause(opts *bind.FilterOpts) (*SaleClockAuctionPauseIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionPauseIterator{contract: _SaleClockAuction.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionPause) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionPause)
				if err := _SaleClockAuction.contract.UnpackLog(event, "Pause", log); err != nil {
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

// SaleClockAuctionUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the SaleClockAuction contract.
type SaleClockAuctionUnpauseIterator struct {
	Event *SaleClockAuctionUnpause // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionUnpause)
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
		it.Event = new(SaleClockAuctionUnpause)
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
func (it *SaleClockAuctionUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionUnpause represents a Unpause event raised by the SaleClockAuction contract.
type SaleClockAuctionUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterUnpause(opts *bind.FilterOpts) (*SaleClockAuctionUnpauseIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionUnpauseIterator{contract: _SaleClockAuction.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionUnpause) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionUnpause)
				if err := _SaleClockAuction.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// SiringClockAuctionABI is the input ABI used to generate the binding from.
const SiringClockAuctionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSiringClockAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"startedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuctionWhenPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_nftAddr\",\"type\":\"address\"},{\"name\":\"_cut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// SiringClockAuctionBin is the compiled bytecode used for deploying new contracts.
const SiringClockAuctionBin = `0x60806040526000805460a060020a60ff02191690556004805460ff1916600117905534801561002d57600080fd5b50604051604080610e1a83398101604052805160209091015160008054600160a060020a031916331781558290829061271082111561006b57600080fd5b506002819055604080517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f9a20483d00000000000000000000000000000000000000000000000000000000600482015290518391600160a060020a038316916301ffc9a7916024808201926020929091908290030181600087803b1580156100f557600080fd5b505af1158015610109573d6000803e3d6000fd5b505050506040513d602081101561011f57600080fd5b5051151561012c57600080fd5b60018054600160a060020a03909216600160a060020a031990921691909117905550505050610cba806101606000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166327ebe40a81146100df5780633f4ba83a1461010e578063454a2ab3146101375780635c975abb146101425780635fd8c7101461015757806376190f8f1461016c57806378bd79351461018157806383b5ff8b146101ce5780638456cb59146101f5578063878eb3681461020a5780638da5cb5b1461022257806396b5a75514610253578063c55d0f561461026b578063dd1b7a0f14610283578063f2fde38b14610298575b600080fd5b3480156100eb57600080fd5b5061010c600435602435604435606435600160a060020a03608435166102b9565b005b34801561011a57600080fd5b506101236103b1565b604080519115158252519081900360200190f35b61010c60043561042c565b34801561014e57600080fd5b50610123610478565b34801561016357600080fd5b5061010c610488565b34801561017857600080fd5b506101236104e5565b34801561018d57600080fd5b506101996004356104ee565b60408051600160a060020a03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b3480156101da57600080fd5b506101e3610584565b60408051918252519081900360200190f35b34801561020157600080fd5b5061012361058a565b34801561021657600080fd5b5061010c60043561060a565b34801561022e57600080fd5b50610237610673565b60408051600160a060020a039092168252519081900360200190f35b34801561025f57600080fd5b5061010c600435610682565b34801561027757600080fd5b506101e36004356106cc565b34801561028f57600080fd5b506102376106fe565b3480156102a457600080fd5b5061010c600160a060020a036004351661070d565b6102c1610c60565b6fffffffffffffffffffffffffffffffff851685146102df57600080fd5b6fffffffffffffffffffffffffffffffff841684146102fd57600080fd5b67ffffffffffffffff8316831461031357600080fd5b600154600160a060020a0316331461032a57600080fd5b610334828761075f565b60a06040519081016040528083600160a060020a03168152602001866fffffffffffffffffffffffffffffffff168152602001856fffffffffffffffffffffffffffffffff1681526020018467ffffffffffffffff1681526020014267ffffffffffffffff1681525090506103a986826107e7565b505050505050565b60008054600160a060020a031633146103c957600080fd5b60005460a060020a900460ff1615156103e157600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a150600190565b600154600090600160a060020a0316331461044657600080fd5b50600081815260036020526040902054600160a060020a0316610469823461093b565b506104748183610a61565b5050565b60005460a060020a900460ff1681565b60015460008054600160a060020a0392831692163314806104b1575033600160a060020a038316145b15156104bc57600080fd5b604051600160a060020a03831690303180156108fc02916000818181858888f150505050505050565b60045460ff1681565b6000818152600360205260408120819081908190819061050d81610acf565b151561051857600080fd5b80546001820154600290920154600160a060020a03909116986fffffffffffffffffffffffffffffffff8084169950700100000000000000000000000000000000909304909216965067ffffffffffffffff808216965068010000000000000000909104169350915050565b60025481565b60008054600160a060020a031633146105a257600080fd5b60005460a060020a900460ff16156105b957600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a150600190565b6000805460a060020a900460ff16151561062357600080fd5b600054600160a060020a0316331461063a57600080fd5b50600081815260036020526040902061065281610acf565b151561065d57600080fd5b8054610474908390600160a060020a0316610af0565b600054600160a060020a031681565b60008181526003602052604081209061069a82610acf565b15156106a557600080fd5b508054600160a060020a03163381146106bd57600080fd5b6106c78382610af0565b505050565b60008181526003602052604081206106e381610acf565b15156106ee57600080fd5b6106f781610b3a565b9392505050565b600154600160a060020a031681565b600054600160a060020a0316331461072457600080fd5b600160a060020a0381161561075c576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015230602483015260448201859052915191909216916323b872dd91606480830192600092919082900301818387803b1580156107d357600080fd5b505af11580156103a9573d6000803e3d6000fd5b603c816060015167ffffffffffffffff161015151561080557600080fd5b60008281526003602090815260409182902083518154600160a060020a0390911673ffffffffffffffffffffffffffffffffffffffff1990911617815581840151600182018054858701516fffffffffffffffffffffffffffffffff90811670010000000000000000000000000000000081029482166fffffffffffffffffffffffffffffffff19909316831790911693909317909155606080870151600290940180546080808a015167ffffffffffffffff90811668010000000000000000026fffffffffffffffff0000000000000000199190981667ffffffffffffffff1990931683171696909617909155865189815295860192909252848601929092529083015291517fa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7929181900390910190a15050565b6000828152600360205260408120818080808061095786610acf565b151561096257600080fd5b61096b86610b3a565b94508488101561097a57600080fd5b8554600160a060020a0316935061099089610bca565b60008511156109e2576109a285610c17565b6040519093508386039250600160a060020a0385169083156108fc029084906000818181858888f193505050501580156109e0573d6000803e3d6000fd5b505b5060405184880390339082156108fc029083906000818181858888f19350505050158015610a14573d6000803e3d6000fd5b50604080518a815260208101879052338183015290517f4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd29181900360600190a15092979650505050505050565b600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038581166004830152602482018590529151919092169163a9059cbb91604480830192600092919082900301818387803b1580156107d357600080fd5b6002015460006801000000000000000090910467ffffffffffffffff161190565b610af982610bca565b610b038183610a61565b6040805183815290517f2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df9181900360200190a15050565b6002810154600090819068010000000000000000900467ffffffffffffffff16421115610b805750600282015468010000000000000000900467ffffffffffffffff1642035b600183015460028401546106f7916fffffffffffffffffffffffffffffffff80821692700100000000000000000000000000000000909204169067ffffffffffffffff1684610c23565b6000908152600360205260408120805473ffffffffffffffffffffffffffffffffffffffff19168155600181019190915560020180546fffffffffffffffffffffffffffffffff19169055565b60025461271091020490565b6000808080858510610c3757869350610c55565b878703925085858402811515610c4957fe5b05915081880190508093505b505050949350505050565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152905600a165627a7a72305820c50d0d71382a875aaa93aa4b7035fe7bcda20c486404f09d19b73247e601a0e00029`

// DeploySiringClockAuction deploys a new Ethereum contract, binding an instance of SiringClockAuction to it.
func DeploySiringClockAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _nftAddr common.Address, _cut *big.Int) (common.Address, *types.Transaction, *SiringClockAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(SiringClockAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SiringClockAuctionBin), backend, _nftAddr, _cut)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SiringClockAuction{SiringClockAuctionCaller: SiringClockAuctionCaller{contract: contract}, SiringClockAuctionTransactor: SiringClockAuctionTransactor{contract: contract}, SiringClockAuctionFilterer: SiringClockAuctionFilterer{contract: contract}}, nil
}

// SiringClockAuction is an auto generated Go binding around an Ethereum contract.
type SiringClockAuction struct {
	SiringClockAuctionCaller     // Read-only binding to the contract
	SiringClockAuctionTransactor // Write-only binding to the contract
	SiringClockAuctionFilterer   // Log filterer for contract events
}

// SiringClockAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SiringClockAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiringClockAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SiringClockAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiringClockAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SiringClockAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiringClockAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SiringClockAuctionSession struct {
	Contract     *SiringClockAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SiringClockAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SiringClockAuctionCallerSession struct {
	Contract *SiringClockAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SiringClockAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SiringClockAuctionTransactorSession struct {
	Contract     *SiringClockAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SiringClockAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SiringClockAuctionRaw struct {
	Contract *SiringClockAuction // Generic contract binding to access the raw methods on
}

// SiringClockAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SiringClockAuctionCallerRaw struct {
	Contract *SiringClockAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SiringClockAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SiringClockAuctionTransactorRaw struct {
	Contract *SiringClockAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSiringClockAuction creates a new instance of SiringClockAuction, bound to a specific deployed contract.
func NewSiringClockAuction(address common.Address, backend bind.ContractBackend) (*SiringClockAuction, error) {
	contract, err := bindSiringClockAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SiringClockAuction{SiringClockAuctionCaller: SiringClockAuctionCaller{contract: contract}, SiringClockAuctionTransactor: SiringClockAuctionTransactor{contract: contract}, SiringClockAuctionFilterer: SiringClockAuctionFilterer{contract: contract}}, nil
}

// NewSiringClockAuctionCaller creates a new read-only instance of SiringClockAuction, bound to a specific deployed contract.
func NewSiringClockAuctionCaller(address common.Address, caller bind.ContractCaller) (*SiringClockAuctionCaller, error) {
	contract, err := bindSiringClockAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionCaller{contract: contract}, nil
}

// NewSiringClockAuctionTransactor creates a new write-only instance of SiringClockAuction, bound to a specific deployed contract.
func NewSiringClockAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SiringClockAuctionTransactor, error) {
	contract, err := bindSiringClockAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionTransactor{contract: contract}, nil
}

// NewSiringClockAuctionFilterer creates a new log filterer instance of SiringClockAuction, bound to a specific deployed contract.
func NewSiringClockAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SiringClockAuctionFilterer, error) {
	contract, err := bindSiringClockAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionFilterer{contract: contract}, nil
}

// bindSiringClockAuction binds a generic wrapper to an already deployed contract.
func bindSiringClockAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SiringClockAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiringClockAuction *SiringClockAuctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SiringClockAuction.Contract.SiringClockAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiringClockAuction *SiringClockAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.SiringClockAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiringClockAuction *SiringClockAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.SiringClockAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiringClockAuction *SiringClockAuctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SiringClockAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiringClockAuction *SiringClockAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiringClockAuction *SiringClockAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.contract.Transact(opts, method, params...)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_SiringClockAuction *SiringClockAuctionCaller) GetAuction(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	ret := new(struct {
		Seller        common.Address
		StartingPrice *big.Int
		EndingPrice   *big.Int
		Duration      *big.Int
		StartedAt     *big.Int
	})
	out := ret
	err := _SiringClockAuction.contract.Call(opts, out, "getAuction", _tokenId)
	return *ret, err
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_SiringClockAuction *SiringClockAuctionSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _SiringClockAuction.Contract.GetAuction(&_SiringClockAuction.CallOpts, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_SiringClockAuction *SiringClockAuctionCallerSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _SiringClockAuction.Contract.GetAuction(&_SiringClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_SiringClockAuction *SiringClockAuctionCaller) GetCurrentPrice(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SiringClockAuction.contract.Call(opts, out, "getCurrentPrice", _tokenId)
	return *ret0, err
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_SiringClockAuction *SiringClockAuctionSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _SiringClockAuction.Contract.GetCurrentPrice(&_SiringClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_SiringClockAuction *SiringClockAuctionCallerSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _SiringClockAuction.Contract.GetCurrentPrice(&_SiringClockAuction.CallOpts, _tokenId)
}

// IsSiringClockAuction is a free data retrieval call binding the contract method 0x76190f8f.
//
// Solidity: function isSiringClockAuction() constant returns(bool)
func (_SiringClockAuction *SiringClockAuctionCaller) IsSiringClockAuction(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SiringClockAuction.contract.Call(opts, out, "isSiringClockAuction")
	return *ret0, err
}

// IsSiringClockAuction is a free data retrieval call binding the contract method 0x76190f8f.
//
// Solidity: function isSiringClockAuction() constant returns(bool)
func (_SiringClockAuction *SiringClockAuctionSession) IsSiringClockAuction() (bool, error) {
	return _SiringClockAuction.Contract.IsSiringClockAuction(&_SiringClockAuction.CallOpts)
}

// IsSiringClockAuction is a free data retrieval call binding the contract method 0x76190f8f.
//
// Solidity: function isSiringClockAuction() constant returns(bool)
func (_SiringClockAuction *SiringClockAuctionCallerSession) IsSiringClockAuction() (bool, error) {
	return _SiringClockAuction.Contract.IsSiringClockAuction(&_SiringClockAuction.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_SiringClockAuction *SiringClockAuctionCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SiringClockAuction.contract.Call(opts, out, "nonFungibleContract")
	return *ret0, err
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_SiringClockAuction *SiringClockAuctionSession) NonFungibleContract() (common.Address, error) {
	return _SiringClockAuction.Contract.NonFungibleContract(&_SiringClockAuction.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_SiringClockAuction *SiringClockAuctionCallerSession) NonFungibleContract() (common.Address, error) {
	return _SiringClockAuction.Contract.NonFungibleContract(&_SiringClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SiringClockAuction *SiringClockAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SiringClockAuction.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SiringClockAuction *SiringClockAuctionSession) Owner() (common.Address, error) {
	return _SiringClockAuction.Contract.Owner(&_SiringClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SiringClockAuction *SiringClockAuctionCallerSession) Owner() (common.Address, error) {
	return _SiringClockAuction.Contract.Owner(&_SiringClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_SiringClockAuction *SiringClockAuctionCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SiringClockAuction.contract.Call(opts, out, "ownerCut")
	return *ret0, err
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_SiringClockAuction *SiringClockAuctionSession) OwnerCut() (*big.Int, error) {
	return _SiringClockAuction.Contract.OwnerCut(&_SiringClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_SiringClockAuction *SiringClockAuctionCallerSession) OwnerCut() (*big.Int, error) {
	return _SiringClockAuction.Contract.OwnerCut(&_SiringClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_SiringClockAuction *SiringClockAuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SiringClockAuction.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_SiringClockAuction *SiringClockAuctionSession) Paused() (bool, error) {
	return _SiringClockAuction.Contract.Paused(&_SiringClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_SiringClockAuction *SiringClockAuctionCallerSession) Paused() (bool, error) {
	return _SiringClockAuction.Contract.Paused(&_SiringClockAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) Bid(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "bid", _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Bid(&_SiringClockAuction.TransactOpts, _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Bid(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "cancelAuction", _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CancelAuction(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CancelAuction(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) CancelAuctionWhenPaused(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "cancelAuctionWhenPaused", _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CancelAuctionWhenPaused(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CancelAuctionWhenPaused(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "createAuction", _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_SiringClockAuction *SiringClockAuctionSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CreateAuction(&_SiringClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CreateAuction(&_SiringClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionSession) Pause() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Pause(&_SiringClockAuction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionTransactorSession) Pause() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Pause(&_SiringClockAuction.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SiringClockAuction *SiringClockAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.TransferOwnership(&_SiringClockAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.TransferOwnership(&_SiringClockAuction.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionSession) Unpause() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Unpause(&_SiringClockAuction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionTransactorSession) Unpause() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Unpause(&_SiringClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SiringClockAuction *SiringClockAuctionSession) WithdrawBalance() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.WithdrawBalance(&_SiringClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.WithdrawBalance(&_SiringClockAuction.TransactOpts)
}

// SiringClockAuctionAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionCancelledIterator struct {
	Event *SiringClockAuctionAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionAuctionCancelled)
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
		it.Event = new(SiringClockAuctionAuctionCancelled)
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
func (it *SiringClockAuctionAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionAuctionCancelled represents a AuctionCancelled event raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*SiringClockAuctionAuctionCancelledIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionAuctionCancelledIterator{contract: _SiringClockAuction.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionAuctionCancelled)
				if err := _SiringClockAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// SiringClockAuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionCreatedIterator struct {
	Event *SiringClockAuctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionAuctionCreated)
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
		it.Event = new(SiringClockAuctionAuctionCreated)
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
func (it *SiringClockAuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionAuctionCreated represents a AuctionCreated event raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*SiringClockAuctionAuctionCreatedIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionAuctionCreatedIterator{contract: _SiringClockAuction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionAuctionCreated)
				if err := _SiringClockAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// SiringClockAuctionAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionSuccessfulIterator struct {
	Event *SiringClockAuctionAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionAuctionSuccessful)
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
		it.Event = new(SiringClockAuctionAuctionSuccessful)
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
func (it *SiringClockAuctionAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionAuctionSuccessful represents a AuctionSuccessful event raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*SiringClockAuctionAuctionSuccessfulIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionAuctionSuccessfulIterator{contract: _SiringClockAuction.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionAuctionSuccessful)
				if err := _SiringClockAuction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// SiringClockAuctionPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the SiringClockAuction contract.
type SiringClockAuctionPauseIterator struct {
	Event *SiringClockAuctionPause // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionPause)
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
		it.Event = new(SiringClockAuctionPause)
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
func (it *SiringClockAuctionPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionPause represents a Pause event raised by the SiringClockAuction contract.
type SiringClockAuctionPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterPause(opts *bind.FilterOpts) (*SiringClockAuctionPauseIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionPauseIterator{contract: _SiringClockAuction.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionPause) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionPause)
				if err := _SiringClockAuction.contract.UnpackLog(event, "Pause", log); err != nil {
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

// SiringClockAuctionUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the SiringClockAuction contract.
type SiringClockAuctionUnpauseIterator struct {
	Event *SiringClockAuctionUnpause // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionUnpause)
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
		it.Event = new(SiringClockAuctionUnpause)
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
func (it *SiringClockAuctionUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionUnpause represents a Unpause event raised by the SiringClockAuction contract.
type SiringClockAuctionUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterUnpause(opts *bind.FilterOpts) (*SiringClockAuctionUnpauseIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionUnpauseIterator{contract: _SiringClockAuction.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionUnpause) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionUnpause)
				if err := _SiringClockAuction.contract.UnpackLog(event, "Unpause", log); err != nil {
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
