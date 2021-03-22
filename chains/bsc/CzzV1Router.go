// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bsc

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BscABI is the input ABI used to generate the binding from.
const BscABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"}],\"name\":\"BurnToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mId\",\"type\":\"uint256\"}],\"name\":\"MintItemCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"MintToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"}],\"name\":\"SwapToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCzzTonkenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCzzTonkenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"value\",\"type\":\"uint8\"}],\"name\":\"setMinSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WethAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapAndBurn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountInMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WethAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapAndBurnEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WethAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WethAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokenForEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"}],\"name\":\"swap_burn_get_amount\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"swap_burn_get_getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"}],\"name\":\"swap_mint_get_amount\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Bsc is an auto generated Go binding around an Ethereum contract.
type Bsc struct {
	BscCaller     // Read-only binding to the contract
	BscTransactor // Write-only binding to the contract
	BscFilterer   // Log filterer for contract events
}

// BscCaller is an auto generated read-only Go binding around an Ethereum contract.
type BscCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BscTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscSession struct {
	Contract     *Bsc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BscCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscCallerSession struct {
	Contract *BscCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BscTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscTransactorSession struct {
	Contract     *BscTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BscRaw is an auto generated low-level Go binding around an Ethereum contract.
type BscRaw struct {
	Contract *Bsc // Generic contract binding to access the raw methods on
}

// BscCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscCallerRaw struct {
	Contract *BscCaller // Generic read-only contract binding to access the raw methods on
}

// BscTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscTransactorRaw struct {
	Contract *BscTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBsc creates a new instance of Bsc, bound to a specific deployed contract.
func NewBsc(address common.Address, backend bind.ContractBackend) (*Bsc, error) {
	contract, err := bindBsc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bsc{BscCaller: BscCaller{contract: contract}, BscTransactor: BscTransactor{contract: contract}, BscFilterer: BscFilterer{contract: contract}}, nil
}

// NewBscCaller creates a new read-only instance of Bsc, bound to a specific deployed contract.
func NewBscCaller(address common.Address, caller bind.ContractCaller) (*BscCaller, error) {
	contract, err := bindBsc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscCaller{contract: contract}, nil
}

// NewBscTransactor creates a new write-only instance of Bsc, bound to a specific deployed contract.
func NewBscTransactor(address common.Address, transactor bind.ContractTransactor) (*BscTransactor, error) {
	contract, err := bindBsc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscTransactor{contract: contract}, nil
}

// NewBscFilterer creates a new log filterer instance of Bsc, bound to a specific deployed contract.
func NewBscFilterer(address common.Address, filterer bind.ContractFilterer) (*BscFilterer, error) {
	contract, err := bindBsc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscFilterer{contract: contract}, nil
}

// bindBsc binds a generic wrapper to an already deployed contract.
func bindBsc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BscABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bsc *BscRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bsc.Contract.BscCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bsc *BscRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bsc.Contract.BscTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bsc *BscRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bsc.Contract.BscTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bsc *BscCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bsc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bsc *BscTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bsc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bsc *BscTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bsc.Contract.contract.Transact(opts, method, params...)
}

// GetCzzTonkenAddress is a free data retrieval call binding the contract method 0x9078f506.
//
// Solidity: function getCzzTonkenAddress() view returns(address)
func (_Bsc *BscCaller) GetCzzTonkenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bsc.contract.Call(opts, &out, "getCzzTonkenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCzzTonkenAddress is a free data retrieval call binding the contract method 0x9078f506.
//
// Solidity: function getCzzTonkenAddress() view returns(address)
func (_Bsc *BscSession) GetCzzTonkenAddress() (common.Address, error) {
	return _Bsc.Contract.GetCzzTonkenAddress(&_Bsc.CallOpts)
}

// GetCzzTonkenAddress is a free data retrieval call binding the contract method 0x9078f506.
//
// Solidity: function getCzzTonkenAddress() view returns(address)
func (_Bsc *BscCallerSession) GetCzzTonkenAddress() (common.Address, error) {
	return _Bsc.Contract.GetCzzTonkenAddress(&_Bsc.CallOpts)
}

// GetMinSignatures is a free data retrieval call binding the contract method 0x1eb36178.
//
// Solidity: function getMinSignatures() view returns(uint256)
func (_Bsc *BscCaller) GetMinSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bsc.contract.Call(opts, &out, "getMinSignatures")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinSignatures is a free data retrieval call binding the contract method 0x1eb36178.
//
// Solidity: function getMinSignatures() view returns(uint256)
func (_Bsc *BscSession) GetMinSignatures() (*big.Int, error) {
	return _Bsc.Contract.GetMinSignatures(&_Bsc.CallOpts)
}

// GetMinSignatures is a free data retrieval call binding the contract method 0x1eb36178.
//
// Solidity: function getMinSignatures() view returns(uint256)
func (_Bsc *BscCallerSession) GetMinSignatures() (*big.Int, error) {
	return _Bsc.Contract.GetMinSignatures(&_Bsc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bsc *BscCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bsc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bsc *BscSession) Owner() (common.Address, error) {
	return _Bsc.Contract.Owner(&_Bsc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bsc *BscCallerSession) Owner() (common.Address, error) {
	return _Bsc.Contract.Owner(&_Bsc.CallOpts)
}

// SwapBurnGetAmount is a free data retrieval call binding the contract method 0x8509e10b.
//
// Solidity: function swap_burn_get_amount(uint256 amountIn, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Bsc *BscCaller) SwapBurnGetAmount(opts *bind.CallOpts, amountIn *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Bsc.contract.Call(opts, &out, "swap_burn_get_amount", amountIn, path, routerAddr)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SwapBurnGetAmount is a free data retrieval call binding the contract method 0x8509e10b.
//
// Solidity: function swap_burn_get_amount(uint256 amountIn, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Bsc *BscSession) SwapBurnGetAmount(amountIn *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	return _Bsc.Contract.SwapBurnGetAmount(&_Bsc.CallOpts, amountIn, path, routerAddr)
}

// SwapBurnGetAmount is a free data retrieval call binding the contract method 0x8509e10b.
//
// Solidity: function swap_burn_get_amount(uint256 amountIn, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Bsc *BscCallerSession) SwapBurnGetAmount(amountIn *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	return _Bsc.Contract.SwapBurnGetAmount(&_Bsc.CallOpts, amountIn, path, routerAddr)
}

// SwapBurnGetGetReserves is a free data retrieval call binding the contract method 0x84ec7ee3.
//
// Solidity: function swap_burn_get_getReserves(address factory, address tokenA, address tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_Bsc *BscCaller) SwapBurnGetGetReserves(opts *bind.CallOpts, factory common.Address, tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _Bsc.contract.Call(opts, &out, "swap_burn_get_getReserves", factory, tokenA, tokenB)

	outstruct := new(struct {
		ReserveA *big.Int
		ReserveB *big.Int
	})

	outstruct.ReserveA = out[0].(*big.Int)
	outstruct.ReserveB = out[1].(*big.Int)

	return *outstruct, err

}

// SwapBurnGetGetReserves is a free data retrieval call binding the contract method 0x84ec7ee3.
//
// Solidity: function swap_burn_get_getReserves(address factory, address tokenA, address tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_Bsc *BscSession) SwapBurnGetGetReserves(factory common.Address, tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Bsc.Contract.SwapBurnGetGetReserves(&_Bsc.CallOpts, factory, tokenA, tokenB)
}

// SwapBurnGetGetReserves is a free data retrieval call binding the contract method 0x84ec7ee3.
//
// Solidity: function swap_burn_get_getReserves(address factory, address tokenA, address tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_Bsc *BscCallerSession) SwapBurnGetGetReserves(factory common.Address, tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Bsc.Contract.SwapBurnGetGetReserves(&_Bsc.CallOpts, factory, tokenA, tokenB)
}

// SwapMintGetAmount is a free data retrieval call binding the contract method 0x3c28d41f.
//
// Solidity: function swap_mint_get_amount(uint256 amountOut, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Bsc *BscCaller) SwapMintGetAmount(opts *bind.CallOpts, amountOut *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Bsc.contract.Call(opts, &out, "swap_mint_get_amount", amountOut, path, routerAddr)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SwapMintGetAmount is a free data retrieval call binding the contract method 0x3c28d41f.
//
// Solidity: function swap_mint_get_amount(uint256 amountOut, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Bsc *BscSession) SwapMintGetAmount(amountOut *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	return _Bsc.Contract.SwapMintGetAmount(&_Bsc.CallOpts, amountOut, path, routerAddr)
}

// SwapMintGetAmount is a free data retrieval call binding the contract method 0x3c28d41f.
//
// Solidity: function swap_mint_get_amount(uint256 amountOut, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Bsc *BscCallerSession) SwapMintGetAmount(amountOut *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	return _Bsc.Contract.SwapMintGetAmount(&_Bsc.CallOpts, amountOut, path, routerAddr)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Bsc *BscTransactor) AddManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Bsc.contract.Transact(opts, "addManager", manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Bsc *BscSession) AddManager(manager common.Address) (*types.Transaction, error) {
	return _Bsc.Contract.AddManager(&_Bsc.TransactOpts, manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Bsc *BscTransactorSession) AddManager(manager common.Address) (*types.Transaction, error) {
	return _Bsc.Contract.AddManager(&_Bsc.TransactOpts, manager)
}

// Burn is a paid mutator transaction binding the contract method 0xce2198a2.
//
// Solidity: function burn(uint256 _amountIn, uint256 ntype, string toToken) payable returns()
func (_Bsc *BscTransactor) Burn(opts *bind.TransactOpts, _amountIn *big.Int, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Bsc.contract.Transact(opts, "burn", _amountIn, ntype, toToken)
}

// Burn is a paid mutator transaction binding the contract method 0xce2198a2.
//
// Solidity: function burn(uint256 _amountIn, uint256 ntype, string toToken) payable returns()
func (_Bsc *BscSession) Burn(_amountIn *big.Int, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Bsc.Contract.Burn(&_Bsc.TransactOpts, _amountIn, ntype, toToken)
}

// Burn is a paid mutator transaction binding the contract method 0xce2198a2.
//
// Solidity: function burn(uint256 _amountIn, uint256 ntype, string toToken) payable returns()
func (_Bsc *BscTransactorSession) Burn(_amountIn *big.Int, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Bsc.Contract.Burn(&_Bsc.TransactOpts, _amountIn, ntype, toToken)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address fromToken, uint256 _amountIn) payable returns()
func (_Bsc *BscTransactor) Mint(opts *bind.TransactOpts, fromToken common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Bsc.contract.Transact(opts, "mint", fromToken, _amountIn)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address fromToken, uint256 _amountIn) payable returns()
func (_Bsc *BscSession) Mint(fromToken common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Bsc.Contract.Mint(&_Bsc.TransactOpts, fromToken, _amountIn)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address fromToken, uint256 _amountIn) payable returns()
func (_Bsc *BscTransactorSession) Mint(fromToken common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Bsc.Contract.Mint(&_Bsc.TransactOpts, fromToken, _amountIn)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Bsc *BscTransactor) RemoveManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Bsc.contract.Transact(opts, "removeManager", manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Bsc *BscSession) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _Bsc.Contract.RemoveManager(&_Bsc.TransactOpts, manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Bsc *BscTransactorSession) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _Bsc.Contract.RemoveManager(&_Bsc.TransactOpts, manager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bsc *BscTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bsc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bsc *BscSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bsc.Contract.RenounceOwnership(&_Bsc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bsc *BscTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bsc.Contract.RenounceOwnership(&_Bsc.TransactOpts)
}

// SetCzzTonkenAddress is a paid mutator transaction binding the contract method 0x44ad1e3d.
//
// Solidity: function setCzzTonkenAddress(address addr) returns()
func (_Bsc *BscTransactor) SetCzzTonkenAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Bsc.contract.Transact(opts, "setCzzTonkenAddress", addr)
}

// SetCzzTonkenAddress is a paid mutator transaction binding the contract method 0x44ad1e3d.
//
// Solidity: function setCzzTonkenAddress(address addr) returns()
func (_Bsc *BscSession) SetCzzTonkenAddress(addr common.Address) (*types.Transaction, error) {
	return _Bsc.Contract.SetCzzTonkenAddress(&_Bsc.TransactOpts, addr)
}

// SetCzzTonkenAddress is a paid mutator transaction binding the contract method 0x44ad1e3d.
//
// Solidity: function setCzzTonkenAddress(address addr) returns()
func (_Bsc *BscTransactorSession) SetCzzTonkenAddress(addr common.Address) (*types.Transaction, error) {
	return _Bsc.Contract.SetCzzTonkenAddress(&_Bsc.TransactOpts, addr)
}

// SetMinSignatures is a paid mutator transaction binding the contract method 0x745d314f.
//
// Solidity: function setMinSignatures(uint8 value) returns()
func (_Bsc *BscTransactor) SetMinSignatures(opts *bind.TransactOpts, value uint8) (*types.Transaction, error) {
	return _Bsc.contract.Transact(opts, "setMinSignatures", value)
}

// SetMinSignatures is a paid mutator transaction binding the contract method 0x745d314f.
//
// Solidity: function setMinSignatures(uint8 value) returns()
func (_Bsc *BscSession) SetMinSignatures(value uint8) (*types.Transaction, error) {
	return _Bsc.Contract.SetMinSignatures(&_Bsc.TransactOpts, value)
}

// SetMinSignatures is a paid mutator transaction binding the contract method 0x745d314f.
//
// Solidity: function setMinSignatures(uint8 value) returns()
func (_Bsc *BscTransactorSession) SetMinSignatures(value uint8) (*types.Transaction, error) {
	return _Bsc.Contract.SetMinSignatures(&_Bsc.TransactOpts, value)
}

// SwapAndBurn is a paid mutator transaction binding the contract method 0xa6d2962b.
//
// Solidity: function swapAndBurn(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Bsc *BscTransactor) SwapAndBurn(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bsc.contract.Transact(opts, "swapAndBurn", _amountIn, _amountOutMin, fromToken, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurn is a paid mutator transaction binding the contract method 0xa6d2962b.
//
// Solidity: function swapAndBurn(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Bsc *BscSession) SwapAndBurn(_amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bsc.Contract.SwapAndBurn(&_Bsc.TransactOpts, _amountIn, _amountOutMin, fromToken, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurn is a paid mutator transaction binding the contract method 0xa6d2962b.
//
// Solidity: function swapAndBurn(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Bsc *BscTransactorSession) SwapAndBurn(_amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bsc.Contract.SwapAndBurn(&_Bsc.TransactOpts, _amountIn, _amountOutMin, fromToken, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurnEth is a paid mutator transaction binding the contract method 0xe11e1a80.
//
// Solidity: function swapAndBurnEth(uint256 _amountInMin, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Bsc *BscTransactor) SwapAndBurnEth(opts *bind.TransactOpts, _amountInMin *big.Int, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bsc.contract.Transact(opts, "swapAndBurnEth", _amountInMin, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurnEth is a paid mutator transaction binding the contract method 0xe11e1a80.
//
// Solidity: function swapAndBurnEth(uint256 _amountInMin, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Bsc *BscSession) SwapAndBurnEth(_amountInMin *big.Int, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bsc.Contract.SwapAndBurnEth(&_Bsc.TransactOpts, _amountInMin, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurnEth is a paid mutator transaction binding the contract method 0xe11e1a80.
//
// Solidity: function swapAndBurnEth(uint256 _amountInMin, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Bsc *BscTransactorSession) SwapAndBurnEth(_amountInMin *big.Int, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bsc.Contract.SwapAndBurnEth(&_Bsc.TransactOpts, _amountInMin, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapToken is a paid mutator transaction binding the contract method 0x1b1046bd.
//
// Solidity: function swapToken(address _to, uint256 _amountIn, uint256 mid, address toToken, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Bsc *BscTransactor) SwapToken(opts *bind.TransactOpts, _to common.Address, _amountIn *big.Int, mid *big.Int, toToken common.Address, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bsc.contract.Transact(opts, "swapToken", _to, _amountIn, mid, toToken, gas, routerAddr, WethAddr, deadline)
}

// SwapToken is a paid mutator transaction binding the contract method 0x1b1046bd.
//
// Solidity: function swapToken(address _to, uint256 _amountIn, uint256 mid, address toToken, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Bsc *BscSession) SwapToken(_to common.Address, _amountIn *big.Int, mid *big.Int, toToken common.Address, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bsc.Contract.SwapToken(&_Bsc.TransactOpts, _to, _amountIn, mid, toToken, gas, routerAddr, WethAddr, deadline)
}

// SwapToken is a paid mutator transaction binding the contract method 0x1b1046bd.
//
// Solidity: function swapToken(address _to, uint256 _amountIn, uint256 mid, address toToken, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Bsc *BscTransactorSession) SwapToken(_to common.Address, _amountIn *big.Int, mid *big.Int, toToken common.Address, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bsc.Contract.SwapToken(&_Bsc.TransactOpts, _to, _amountIn, mid, toToken, gas, routerAddr, WethAddr, deadline)
}

// SwapTokenForEth is a paid mutator transaction binding the contract method 0xcabab781.
//
// Solidity: function swapTokenForEth(address _to, uint256 _amountIn, uint256 mid, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Bsc *BscTransactor) SwapTokenForEth(opts *bind.TransactOpts, _to common.Address, _amountIn *big.Int, mid *big.Int, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bsc.contract.Transact(opts, "swapTokenForEth", _to, _amountIn, mid, gas, routerAddr, WethAddr, deadline)
}

// SwapTokenForEth is a paid mutator transaction binding the contract method 0xcabab781.
//
// Solidity: function swapTokenForEth(address _to, uint256 _amountIn, uint256 mid, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Bsc *BscSession) SwapTokenForEth(_to common.Address, _amountIn *big.Int, mid *big.Int, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bsc.Contract.SwapTokenForEth(&_Bsc.TransactOpts, _to, _amountIn, mid, gas, routerAddr, WethAddr, deadline)
}

// SwapTokenForEth is a paid mutator transaction binding the contract method 0xcabab781.
//
// Solidity: function swapTokenForEth(address _to, uint256 _amountIn, uint256 mid, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Bsc *BscTransactorSession) SwapTokenForEth(_to common.Address, _amountIn *big.Int, mid *big.Int, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bsc.Contract.SwapTokenForEth(&_Bsc.TransactOpts, _to, _amountIn, mid, gas, routerAddr, WethAddr, deadline)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bsc *BscTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bsc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bsc *BscSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bsc.Contract.TransferOwnership(&_Bsc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bsc *BscTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bsc.Contract.TransferOwnership(&_Bsc.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bsc *BscTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bsc.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bsc *BscSession) Receive() (*types.Transaction, error) {
	return _Bsc.Contract.Receive(&_Bsc.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bsc *BscTransactorSession) Receive() (*types.Transaction, error) {
	return _Bsc.Contract.Receive(&_Bsc.TransactOpts)
}

// BscBurnTokenIterator is returned from FilterBurnToken and is used to iterate over the raw logs and unpacked data for BurnToken events raised by the Bsc contract.
type BscBurnTokenIterator struct {
	Event *BscBurnToken // Event containing the contract specifics and raw log

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
func (it *BscBurnTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscBurnToken)
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
		it.Event = new(BscBurnToken)
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
func (it *BscBurnTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscBurnTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscBurnToken represents a BurnToken event raised by the Bsc contract.
type BscBurnToken struct {
	To      common.Address
	Amount  *big.Int
	Ntype   *big.Int
	ToToken string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurnToken is a free log retrieval operation binding the contract event 0x86f32d6c7a935bd338ee00610630fcfb6f043a6ad755db62064ce2ad92c45caa.
//
// Solidity: event BurnToken(address indexed to, uint256 amount, uint256 ntype, string toToken)
func (_Bsc *BscFilterer) FilterBurnToken(opts *bind.FilterOpts, to []common.Address) (*BscBurnTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bsc.contract.FilterLogs(opts, "BurnToken", toRule)
	if err != nil {
		return nil, err
	}
	return &BscBurnTokenIterator{contract: _Bsc.contract, event: "BurnToken", logs: logs, sub: sub}, nil
}

// WatchBurnToken is a free log subscription operation binding the contract event 0x86f32d6c7a935bd338ee00610630fcfb6f043a6ad755db62064ce2ad92c45caa.
//
// Solidity: event BurnToken(address indexed to, uint256 amount, uint256 ntype, string toToken)
func (_Bsc *BscFilterer) WatchBurnToken(opts *bind.WatchOpts, sink chan<- *BscBurnToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bsc.contract.WatchLogs(opts, "BurnToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscBurnToken)
				if err := _Bsc.contract.UnpackLog(event, "BurnToken", log); err != nil {
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

// ParseBurnToken is a log parse operation binding the contract event 0x86f32d6c7a935bd338ee00610630fcfb6f043a6ad755db62064ce2ad92c45caa.
//
// Solidity: event BurnToken(address indexed to, uint256 amount, uint256 ntype, string toToken)
func (_Bsc *BscFilterer) ParseBurnToken(log types.Log) (*BscBurnToken, error) {
	event := new(BscBurnToken)
	if err := _Bsc.contract.UnpackLog(event, "BurnToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscMintItemCreatedIterator is returned from FilterMintItemCreated and is used to iterate over the raw logs and unpacked data for MintItemCreated events raised by the Bsc contract.
type BscMintItemCreatedIterator struct {
	Event *BscMintItemCreated // Event containing the contract specifics and raw log

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
func (it *BscMintItemCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscMintItemCreated)
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
		it.Event = new(BscMintItemCreated)
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
func (it *BscMintItemCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscMintItemCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscMintItemCreated represents a MintItemCreated event raised by the Bsc contract.
type BscMintItemCreated struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	MId    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintItemCreated is a free log retrieval operation binding the contract event 0x82911cbfb5e991769299a8d7cf2c610644a7c6ed1349882285349faed8ca7bb8.
//
// Solidity: event MintItemCreated(address indexed from, address to, uint256 amount, uint256 mId)
func (_Bsc *BscFilterer) FilterMintItemCreated(opts *bind.FilterOpts, from []common.Address) (*BscMintItemCreatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bsc.contract.FilterLogs(opts, "MintItemCreated", fromRule)
	if err != nil {
		return nil, err
	}
	return &BscMintItemCreatedIterator{contract: _Bsc.contract, event: "MintItemCreated", logs: logs, sub: sub}, nil
}

// WatchMintItemCreated is a free log subscription operation binding the contract event 0x82911cbfb5e991769299a8d7cf2c610644a7c6ed1349882285349faed8ca7bb8.
//
// Solidity: event MintItemCreated(address indexed from, address to, uint256 amount, uint256 mId)
func (_Bsc *BscFilterer) WatchMintItemCreated(opts *bind.WatchOpts, sink chan<- *BscMintItemCreated, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bsc.contract.WatchLogs(opts, "MintItemCreated", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscMintItemCreated)
				if err := _Bsc.contract.UnpackLog(event, "MintItemCreated", log); err != nil {
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

// ParseMintItemCreated is a log parse operation binding the contract event 0x82911cbfb5e991769299a8d7cf2c610644a7c6ed1349882285349faed8ca7bb8.
//
// Solidity: event MintItemCreated(address indexed from, address to, uint256 amount, uint256 mId)
func (_Bsc *BscFilterer) ParseMintItemCreated(log types.Log) (*BscMintItemCreated, error) {
	event := new(BscMintItemCreated)
	if err := _Bsc.contract.UnpackLog(event, "MintItemCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscMintTokenIterator is returned from FilterMintToken and is used to iterate over the raw logs and unpacked data for MintToken events raised by the Bsc contract.
type BscMintTokenIterator struct {
	Event *BscMintToken // Event containing the contract specifics and raw log

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
func (it *BscMintTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscMintToken)
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
		it.Event = new(BscMintToken)
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
func (it *BscMintTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscMintTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscMintToken represents a MintToken event raised by the Bsc contract.
type BscMintToken struct {
	To       common.Address
	Amount   *big.Int
	Mid      *big.Int
	AmountIn *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintToken is a free log retrieval operation binding the contract event 0x8fb5c7bffbb272c541556c455c74269997b816df24f56dd255c2391d92d4f1e9.
//
// Solidity: event MintToken(address indexed to, uint256 amount, uint256 mid, uint256 amountIn)
func (_Bsc *BscFilterer) FilterMintToken(opts *bind.FilterOpts, to []common.Address) (*BscMintTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bsc.contract.FilterLogs(opts, "MintToken", toRule)
	if err != nil {
		return nil, err
	}
	return &BscMintTokenIterator{contract: _Bsc.contract, event: "MintToken", logs: logs, sub: sub}, nil
}

// WatchMintToken is a free log subscription operation binding the contract event 0x8fb5c7bffbb272c541556c455c74269997b816df24f56dd255c2391d92d4f1e9.
//
// Solidity: event MintToken(address indexed to, uint256 amount, uint256 mid, uint256 amountIn)
func (_Bsc *BscFilterer) WatchMintToken(opts *bind.WatchOpts, sink chan<- *BscMintToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bsc.contract.WatchLogs(opts, "MintToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscMintToken)
				if err := _Bsc.contract.UnpackLog(event, "MintToken", log); err != nil {
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

// ParseMintToken is a log parse operation binding the contract event 0x8fb5c7bffbb272c541556c455c74269997b816df24f56dd255c2391d92d4f1e9.
//
// Solidity: event MintToken(address indexed to, uint256 amount, uint256 mid, uint256 amountIn)
func (_Bsc *BscFilterer) ParseMintToken(log types.Log) (*BscMintToken, error) {
	event := new(BscMintToken)
	if err := _Bsc.contract.UnpackLog(event, "MintToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bsc contract.
type BscOwnershipTransferredIterator struct {
	Event *BscOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BscOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscOwnershipTransferred)
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
		it.Event = new(BscOwnershipTransferred)
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
func (it *BscOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscOwnershipTransferred represents a OwnershipTransferred event raised by the Bsc contract.
type BscOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bsc *BscFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BscOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bsc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BscOwnershipTransferredIterator{contract: _Bsc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bsc *BscFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BscOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bsc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscOwnershipTransferred)
				if err := _Bsc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bsc *BscFilterer) ParseOwnershipTransferred(log types.Log) (*BscOwnershipTransferred, error) {
	event := new(BscOwnershipTransferred)
	if err := _Bsc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscSwapTokenIterator is returned from FilterSwapToken and is used to iterate over the raw logs and unpacked data for SwapToken events raised by the Bsc contract.
type BscSwapTokenIterator struct {
	Event *BscSwapToken // Event containing the contract specifics and raw log

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
func (it *BscSwapTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscSwapToken)
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
		it.Event = new(BscSwapToken)
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
func (it *BscSwapTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscSwapTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscSwapToken represents a SwapToken event raised by the Bsc contract.
type BscSwapToken struct {
	To        common.Address
	InAmount  *big.Int
	OutAmount *big.Int
	Flag      string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapToken is a free log retrieval operation binding the contract event 0x59472f3e5abe9ecfcb0d72f4334caeb814fd3a06d2b049f049b039d14c663f8c.
//
// Solidity: event SwapToken(address indexed to, uint256 inAmount, uint256 outAmount, string flag)
func (_Bsc *BscFilterer) FilterSwapToken(opts *bind.FilterOpts, to []common.Address) (*BscSwapTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bsc.contract.FilterLogs(opts, "SwapToken", toRule)
	if err != nil {
		return nil, err
	}
	return &BscSwapTokenIterator{contract: _Bsc.contract, event: "SwapToken", logs: logs, sub: sub}, nil
}

// WatchSwapToken is a free log subscription operation binding the contract event 0x59472f3e5abe9ecfcb0d72f4334caeb814fd3a06d2b049f049b039d14c663f8c.
//
// Solidity: event SwapToken(address indexed to, uint256 inAmount, uint256 outAmount, string flag)
func (_Bsc *BscFilterer) WatchSwapToken(opts *bind.WatchOpts, sink chan<- *BscSwapToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bsc.contract.WatchLogs(opts, "SwapToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscSwapToken)
				if err := _Bsc.contract.UnpackLog(event, "SwapToken", log); err != nil {
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

// ParseSwapToken is a log parse operation binding the contract event 0x59472f3e5abe9ecfcb0d72f4334caeb814fd3a06d2b049f049b039d14c663f8c.
//
// Solidity: event SwapToken(address indexed to, uint256 inAmount, uint256 outAmount, string flag)
func (_Bsc *BscFilterer) ParseSwapToken(log types.Log) (*BscSwapToken, error) {
	event := new(BscSwapToken)
	if err := _Bsc.contract.UnpackLog(event, "SwapToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscTransferTokenIterator is returned from FilterTransferToken and is used to iterate over the raw logs and unpacked data for TransferToken events raised by the Bsc contract.
type BscTransferTokenIterator struct {
	Event *BscTransferToken // Event containing the contract specifics and raw log

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
func (it *BscTransferTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscTransferToken)
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
		it.Event = new(BscTransferToken)
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
func (it *BscTransferTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscTransferTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscTransferToken represents a TransferToken event raised by the Bsc contract.
type BscTransferToken struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferToken is a free log retrieval operation binding the contract event 0x9dd3045b6df532ed81beb2a333cec6249dafd3c2fc54c80c50155cb0e1a0ba1e.
//
// Solidity: event TransferToken(address indexed to, uint256 amount)
func (_Bsc *BscFilterer) FilterTransferToken(opts *bind.FilterOpts, to []common.Address) (*BscTransferTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bsc.contract.FilterLogs(opts, "TransferToken", toRule)
	if err != nil {
		return nil, err
	}
	return &BscTransferTokenIterator{contract: _Bsc.contract, event: "TransferToken", logs: logs, sub: sub}, nil
}

// WatchTransferToken is a free log subscription operation binding the contract event 0x9dd3045b6df532ed81beb2a333cec6249dafd3c2fc54c80c50155cb0e1a0ba1e.
//
// Solidity: event TransferToken(address indexed to, uint256 amount)
func (_Bsc *BscFilterer) WatchTransferToken(opts *bind.WatchOpts, sink chan<- *BscTransferToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bsc.contract.WatchLogs(opts, "TransferToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscTransferToken)
				if err := _Bsc.contract.UnpackLog(event, "TransferToken", log); err != nil {
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

// ParseTransferToken is a log parse operation binding the contract event 0x9dd3045b6df532ed81beb2a333cec6249dafd3c2fc54c80c50155cb0e1a0ba1e.
//
// Solidity: event TransferToken(address indexed to, uint256 amount)
func (_Bsc *BscFilterer) ParseTransferToken(log types.Log) (*BscTransferToken, error) {
	event := new(BscTransferToken)
	if err := _Bsc.contract.UnpackLog(event, "TransferToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
