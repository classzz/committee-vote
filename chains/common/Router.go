// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package common

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

// CommonABI is the input ABI used to generate the binding from.
const CommonABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"}],\"name\":\"BurnToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mId\",\"type\":\"uint256\"}],\"name\":\"MintItemCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"MintToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"}],\"name\":\"SwapToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCzzTonkenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCzzTonkenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"value\",\"type\":\"uint8\"}],\"name\":\"setMinSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WethAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapAndBurn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountInMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WethAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapAndBurnEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapAndBurnWithPath\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WethAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WethAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokenForEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WethAddr\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokenWithPath\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"}],\"name\":\"swap_burn_get_amount\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"swap_burn_get_getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"}],\"name\":\"swap_mint_get_amount\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Common is an auto generated Go binding around an Ethereum contract.
type Common struct {
	CommonCaller     // Read-only binding to the contract
	CommonTransactor // Write-only binding to the contract
	CommonFilterer   // Log filterer for contract events
}

// CommonCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommonSession struct {
	Contract     *Common           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommonCallerSession struct {
	Contract *CommonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CommonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommonTransactorSession struct {
	Contract     *CommonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommonRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommonRaw struct {
	Contract *Common // Generic contract binding to access the raw methods on
}

// CommonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommonCallerRaw struct {
	Contract *CommonCaller // Generic read-only contract binding to access the raw methods on
}

// CommonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommonTransactorRaw struct {
	Contract *CommonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommon creates a new instance of Common, bound to a specific deployed contract.
func NewCommon(address common.Address, backend bind.ContractBackend) (*Common, error) {
	contract, err := bindCommon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Common{CommonCaller: CommonCaller{contract: contract}, CommonTransactor: CommonTransactor{contract: contract}, CommonFilterer: CommonFilterer{contract: contract}}, nil
}

// NewCommonCaller creates a new read-only instance of Common, bound to a specific deployed contract.
func NewCommonCaller(address common.Address, caller bind.ContractCaller) (*CommonCaller, error) {
	contract, err := bindCommon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommonCaller{contract: contract}, nil
}

// NewCommonTransactor creates a new write-only instance of Common, bound to a specific deployed contract.
func NewCommonTransactor(address common.Address, transactor bind.ContractTransactor) (*CommonTransactor, error) {
	contract, err := bindCommon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommonTransactor{contract: contract}, nil
}

// NewCommonFilterer creates a new log filterer instance of Common, bound to a specific deployed contract.
func NewCommonFilterer(address common.Address, filterer bind.ContractFilterer) (*CommonFilterer, error) {
	contract, err := bindCommon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommonFilterer{contract: contract}, nil
}

// bindCommon binds a generic wrapper to an already deployed contract.
func bindCommon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CommonABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Common *CommonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Common.Contract.CommonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Common *CommonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Common.Contract.CommonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Common *CommonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Common.Contract.CommonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Common *CommonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Common.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Common *CommonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Common.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Common *CommonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Common.Contract.contract.Transact(opts, method, params...)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x53290b44.
//
// Solidity: function getBalanceOf(address token, address account) view returns(uint256)
func (_Common *CommonCaller) GetBalanceOf(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "getBalanceOf", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceOf is a free data retrieval call binding the contract method 0x53290b44.
//
// Solidity: function getBalanceOf(address token, address account) view returns(uint256)
func (_Common *CommonSession) GetBalanceOf(token common.Address, account common.Address) (*big.Int, error) {
	return _Common.Contract.GetBalanceOf(&_Common.CallOpts, token, account)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x53290b44.
//
// Solidity: function getBalanceOf(address token, address account) view returns(uint256)
func (_Common *CommonCallerSession) GetBalanceOf(token common.Address, account common.Address) (*big.Int, error) {
	return _Common.Contract.GetBalanceOf(&_Common.CallOpts, token, account)
}

// GetCzzTonkenAddress is a free data retrieval call binding the contract method 0x9078f506.
//
// Solidity: function getCzzTonkenAddress() view returns(address)
func (_Common *CommonCaller) GetCzzTonkenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "getCzzTonkenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCzzTonkenAddress is a free data retrieval call binding the contract method 0x9078f506.
//
// Solidity: function getCzzTonkenAddress() view returns(address)
func (_Common *CommonSession) GetCzzTonkenAddress() (common.Address, error) {
	return _Common.Contract.GetCzzTonkenAddress(&_Common.CallOpts)
}

// GetCzzTonkenAddress is a free data retrieval call binding the contract method 0x9078f506.
//
// Solidity: function getCzzTonkenAddress() view returns(address)
func (_Common *CommonCallerSession) GetCzzTonkenAddress() (common.Address, error) {
	return _Common.Contract.GetCzzTonkenAddress(&_Common.CallOpts)
}

// GetMinSignatures is a free data retrieval call binding the contract method 0x1eb36178.
//
// Solidity: function getMinSignatures() view returns(uint256)
func (_Common *CommonCaller) GetMinSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "getMinSignatures")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinSignatures is a free data retrieval call binding the contract method 0x1eb36178.
//
// Solidity: function getMinSignatures() view returns(uint256)
func (_Common *CommonSession) GetMinSignatures() (*big.Int, error) {
	return _Common.Contract.GetMinSignatures(&_Common.CallOpts)
}

// GetMinSignatures is a free data retrieval call binding the contract method 0x1eb36178.
//
// Solidity: function getMinSignatures() view returns(uint256)
func (_Common *CommonCallerSession) GetMinSignatures() (*big.Int, error) {
	return _Common.Contract.GetMinSignatures(&_Common.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Common *CommonCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Common *CommonSession) Owner() (common.Address, error) {
	return _Common.Contract.Owner(&_Common.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Common *CommonCallerSession) Owner() (common.Address, error) {
	return _Common.Contract.Owner(&_Common.CallOpts)
}

// SwapBurnGetAmount is a free data retrieval call binding the contract method 0x8509e10b.
//
// Solidity: function swap_burn_get_amount(uint256 amountIn, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Common *CommonCaller) SwapBurnGetAmount(opts *bind.CallOpts, amountIn *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "swap_burn_get_amount", amountIn, path, routerAddr)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SwapBurnGetAmount is a free data retrieval call binding the contract method 0x8509e10b.
//
// Solidity: function swap_burn_get_amount(uint256 amountIn, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Common *CommonSession) SwapBurnGetAmount(amountIn *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	return _Common.Contract.SwapBurnGetAmount(&_Common.CallOpts, amountIn, path, routerAddr)
}

// SwapBurnGetAmount is a free data retrieval call binding the contract method 0x8509e10b.
//
// Solidity: function swap_burn_get_amount(uint256 amountIn, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Common *CommonCallerSession) SwapBurnGetAmount(amountIn *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	return _Common.Contract.SwapBurnGetAmount(&_Common.CallOpts, amountIn, path, routerAddr)
}

// SwapBurnGetGetReserves is a free data retrieval call binding the contract method 0x84ec7ee3.
//
// Solidity: function swap_burn_get_getReserves(address factory, address tokenA, address tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_Common *CommonCaller) SwapBurnGetGetReserves(opts *bind.CallOpts, factory common.Address, tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "swap_burn_get_getReserves", factory, tokenA, tokenB)

	outstruct := new(struct {
		ReserveA *big.Int
		ReserveB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SwapBurnGetGetReserves is a free data retrieval call binding the contract method 0x84ec7ee3.
//
// Solidity: function swap_burn_get_getReserves(address factory, address tokenA, address tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_Common *CommonSession) SwapBurnGetGetReserves(factory common.Address, tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Common.Contract.SwapBurnGetGetReserves(&_Common.CallOpts, factory, tokenA, tokenB)
}

// SwapBurnGetGetReserves is a free data retrieval call binding the contract method 0x84ec7ee3.
//
// Solidity: function swap_burn_get_getReserves(address factory, address tokenA, address tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_Common *CommonCallerSession) SwapBurnGetGetReserves(factory common.Address, tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Common.Contract.SwapBurnGetGetReserves(&_Common.CallOpts, factory, tokenA, tokenB)
}

// SwapMintGetAmount is a free data retrieval call binding the contract method 0x3c28d41f.
//
// Solidity: function swap_mint_get_amount(uint256 amountOut, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Common *CommonCaller) SwapMintGetAmount(opts *bind.CallOpts, amountOut *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "swap_mint_get_amount", amountOut, path, routerAddr)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SwapMintGetAmount is a free data retrieval call binding the contract method 0x3c28d41f.
//
// Solidity: function swap_mint_get_amount(uint256 amountOut, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Common *CommonSession) SwapMintGetAmount(amountOut *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	return _Common.Contract.SwapMintGetAmount(&_Common.CallOpts, amountOut, path, routerAddr)
}

// SwapMintGetAmount is a free data retrieval call binding the contract method 0x3c28d41f.
//
// Solidity: function swap_mint_get_amount(uint256 amountOut, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Common *CommonCallerSession) SwapMintGetAmount(amountOut *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	return _Common.Contract.SwapMintGetAmount(&_Common.CallOpts, amountOut, path, routerAddr)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Common *CommonTransactor) AddManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "addManager", manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Common *CommonSession) AddManager(manager common.Address) (*types.Transaction, error) {
	return _Common.Contract.AddManager(&_Common.TransactOpts, manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Common *CommonTransactorSession) AddManager(manager common.Address) (*types.Transaction, error) {
	return _Common.Contract.AddManager(&_Common.TransactOpts, manager)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 _amount) returns(bool)
func (_Common *CommonTransactor) Approve(opts *bind.TransactOpts, token common.Address, spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "approve", token, spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 _amount) returns(bool)
func (_Common *CommonSession) Approve(token common.Address, spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Common.Contract.Approve(&_Common.TransactOpts, token, spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 _amount) returns(bool)
func (_Common *CommonTransactorSession) Approve(token common.Address, spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Common.Contract.Approve(&_Common.TransactOpts, token, spender, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xce2198a2.
//
// Solidity: function burn(uint256 _amountIn, uint256 ntype, string toToken) payable returns()
func (_Common *CommonTransactor) Burn(opts *bind.TransactOpts, _amountIn *big.Int, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "burn", _amountIn, ntype, toToken)
}

// Burn is a paid mutator transaction binding the contract method 0xce2198a2.
//
// Solidity: function burn(uint256 _amountIn, uint256 ntype, string toToken) payable returns()
func (_Common *CommonSession) Burn(_amountIn *big.Int, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Common.Contract.Burn(&_Common.TransactOpts, _amountIn, ntype, toToken)
}

// Burn is a paid mutator transaction binding the contract method 0xce2198a2.
//
// Solidity: function burn(uint256 _amountIn, uint256 ntype, string toToken) payable returns()
func (_Common *CommonTransactorSession) Burn(_amountIn *big.Int, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Common.Contract.Burn(&_Common.TransactOpts, _amountIn, ntype, toToken)
}

// Mint is a paid mutator transaction binding the contract method 0x836a1040.
//
// Solidity: function mint(uint256 mid, address fromToken, uint256 _amountIn) payable returns()
func (_Common *CommonTransactor) Mint(opts *bind.TransactOpts, mid *big.Int, fromToken common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "mint", mid, fromToken, _amountIn)
}

// Mint is a paid mutator transaction binding the contract method 0x836a1040.
//
// Solidity: function mint(uint256 mid, address fromToken, uint256 _amountIn) payable returns()
func (_Common *CommonSession) Mint(mid *big.Int, fromToken common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Common.Contract.Mint(&_Common.TransactOpts, mid, fromToken, _amountIn)
}

// Mint is a paid mutator transaction binding the contract method 0x836a1040.
//
// Solidity: function mint(uint256 mid, address fromToken, uint256 _amountIn) payable returns()
func (_Common *CommonTransactorSession) Mint(mid *big.Int, fromToken common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Common.Contract.Mint(&_Common.TransactOpts, mid, fromToken, _amountIn)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Common *CommonTransactor) RemoveManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "removeManager", manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Common *CommonSession) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _Common.Contract.RemoveManager(&_Common.TransactOpts, manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Common *CommonTransactorSession) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _Common.Contract.RemoveManager(&_Common.TransactOpts, manager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Common *CommonTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Common *CommonSession) RenounceOwnership() (*types.Transaction, error) {
	return _Common.Contract.RenounceOwnership(&_Common.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Common *CommonTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Common.Contract.RenounceOwnership(&_Common.TransactOpts)
}

// SetCzzTonkenAddress is a paid mutator transaction binding the contract method 0x44ad1e3d.
//
// Solidity: function setCzzTonkenAddress(address addr) returns()
func (_Common *CommonTransactor) SetCzzTonkenAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "setCzzTonkenAddress", addr)
}

// SetCzzTonkenAddress is a paid mutator transaction binding the contract method 0x44ad1e3d.
//
// Solidity: function setCzzTonkenAddress(address addr) returns()
func (_Common *CommonSession) SetCzzTonkenAddress(addr common.Address) (*types.Transaction, error) {
	return _Common.Contract.SetCzzTonkenAddress(&_Common.TransactOpts, addr)
}

// SetCzzTonkenAddress is a paid mutator transaction binding the contract method 0x44ad1e3d.
//
// Solidity: function setCzzTonkenAddress(address addr) returns()
func (_Common *CommonTransactorSession) SetCzzTonkenAddress(addr common.Address) (*types.Transaction, error) {
	return _Common.Contract.SetCzzTonkenAddress(&_Common.TransactOpts, addr)
}

// SetMinSignatures is a paid mutator transaction binding the contract method 0x745d314f.
//
// Solidity: function setMinSignatures(uint8 value) returns()
func (_Common *CommonTransactor) SetMinSignatures(opts *bind.TransactOpts, value uint8) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "setMinSignatures", value)
}

// SetMinSignatures is a paid mutator transaction binding the contract method 0x745d314f.
//
// Solidity: function setMinSignatures(uint8 value) returns()
func (_Common *CommonSession) SetMinSignatures(value uint8) (*types.Transaction, error) {
	return _Common.Contract.SetMinSignatures(&_Common.TransactOpts, value)
}

// SetMinSignatures is a paid mutator transaction binding the contract method 0x745d314f.
//
// Solidity: function setMinSignatures(uint8 value) returns()
func (_Common *CommonTransactorSession) SetMinSignatures(value uint8) (*types.Transaction, error) {
	return _Common.Contract.SetMinSignatures(&_Common.TransactOpts, value)
}

// SwapAndBurn is a paid mutator transaction binding the contract method 0xa6d2962b.
//
// Solidity: function swapAndBurn(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Common *CommonTransactor) SwapAndBurn(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "swapAndBurn", _amountIn, _amountOutMin, fromToken, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurn is a paid mutator transaction binding the contract method 0xa6d2962b.
//
// Solidity: function swapAndBurn(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Common *CommonSession) SwapAndBurn(_amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SwapAndBurn(&_Common.TransactOpts, _amountIn, _amountOutMin, fromToken, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurn is a paid mutator transaction binding the contract method 0xa6d2962b.
//
// Solidity: function swapAndBurn(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Common *CommonTransactorSession) SwapAndBurn(_amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SwapAndBurn(&_Common.TransactOpts, _amountIn, _amountOutMin, fromToken, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurnEth is a paid mutator transaction binding the contract method 0xe11e1a80.
//
// Solidity: function swapAndBurnEth(uint256 _amountInMin, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Common *CommonTransactor) SwapAndBurnEth(opts *bind.TransactOpts, _amountInMin *big.Int, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "swapAndBurnEth", _amountInMin, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurnEth is a paid mutator transaction binding the contract method 0xe11e1a80.
//
// Solidity: function swapAndBurnEth(uint256 _amountInMin, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Common *CommonSession) SwapAndBurnEth(_amountInMin *big.Int, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SwapAndBurnEth(&_Common.TransactOpts, _amountInMin, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurnEth is a paid mutator transaction binding the contract method 0xe11e1a80.
//
// Solidity: function swapAndBurnEth(uint256 _amountInMin, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Common *CommonTransactorSession) SwapAndBurnEth(_amountInMin *big.Int, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SwapAndBurnEth(&_Common.TransactOpts, _amountInMin, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurnWithPath is a paid mutator transaction binding the contract method 0x6f5e73ad.
//
// Solidity: function swapAndBurnWithPath(uint256 _amountIn, uint256 _amountOutMin, uint256 ntype, string toToken, address routerAddr, address[] path, uint256 deadline) payable returns()
func (_Common *CommonTransactor) SwapAndBurnWithPath(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, ntype *big.Int, toToken string, routerAddr common.Address, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "swapAndBurnWithPath", _amountIn, _amountOutMin, ntype, toToken, routerAddr, path, deadline)
}

// SwapAndBurnWithPath is a paid mutator transaction binding the contract method 0x6f5e73ad.
//
// Solidity: function swapAndBurnWithPath(uint256 _amountIn, uint256 _amountOutMin, uint256 ntype, string toToken, address routerAddr, address[] path, uint256 deadline) payable returns()
func (_Common *CommonSession) SwapAndBurnWithPath(_amountIn *big.Int, _amountOutMin *big.Int, ntype *big.Int, toToken string, routerAddr common.Address, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SwapAndBurnWithPath(&_Common.TransactOpts, _amountIn, _amountOutMin, ntype, toToken, routerAddr, path, deadline)
}

// SwapAndBurnWithPath is a paid mutator transaction binding the contract method 0x6f5e73ad.
//
// Solidity: function swapAndBurnWithPath(uint256 _amountIn, uint256 _amountOutMin, uint256 ntype, string toToken, address routerAddr, address[] path, uint256 deadline) payable returns()
func (_Common *CommonTransactorSession) SwapAndBurnWithPath(_amountIn *big.Int, _amountOutMin *big.Int, ntype *big.Int, toToken string, routerAddr common.Address, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SwapAndBurnWithPath(&_Common.TransactOpts, _amountIn, _amountOutMin, ntype, toToken, routerAddr, path, deadline)
}

// SwapToken is a paid mutator transaction binding the contract method 0x1b1046bd.
//
// Solidity: function swapToken(address _to, uint256 _amountIn, uint256 mid, address toToken, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Common *CommonTransactor) SwapToken(opts *bind.TransactOpts, _to common.Address, _amountIn *big.Int, mid *big.Int, toToken common.Address, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "swapToken", _to, _amountIn, mid, toToken, gas, routerAddr, WethAddr, deadline)
}

// SwapToken is a paid mutator transaction binding the contract method 0x1b1046bd.
//
// Solidity: function swapToken(address _to, uint256 _amountIn, uint256 mid, address toToken, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Common *CommonSession) SwapToken(_to common.Address, _amountIn *big.Int, mid *big.Int, toToken common.Address, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SwapToken(&_Common.TransactOpts, _to, _amountIn, mid, toToken, gas, routerAddr, WethAddr, deadline)
}

// SwapToken is a paid mutator transaction binding the contract method 0x1b1046bd.
//
// Solidity: function swapToken(address _to, uint256 _amountIn, uint256 mid, address toToken, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Common *CommonTransactorSession) SwapToken(_to common.Address, _amountIn *big.Int, mid *big.Int, toToken common.Address, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SwapToken(&_Common.TransactOpts, _to, _amountIn, mid, toToken, gas, routerAddr, WethAddr, deadline)
}

// SwapTokenForEth is a paid mutator transaction binding the contract method 0xcabab781.
//
// Solidity: function swapTokenForEth(address _to, uint256 _amountIn, uint256 mid, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Common *CommonTransactor) SwapTokenForEth(opts *bind.TransactOpts, _to common.Address, _amountIn *big.Int, mid *big.Int, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "swapTokenForEth", _to, _amountIn, mid, gas, routerAddr, WethAddr, deadline)
}

// SwapTokenForEth is a paid mutator transaction binding the contract method 0xcabab781.
//
// Solidity: function swapTokenForEth(address _to, uint256 _amountIn, uint256 mid, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Common *CommonSession) SwapTokenForEth(_to common.Address, _amountIn *big.Int, mid *big.Int, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SwapTokenForEth(&_Common.TransactOpts, _to, _amountIn, mid, gas, routerAddr, WethAddr, deadline)
}

// SwapTokenForEth is a paid mutator transaction binding the contract method 0xcabab781.
//
// Solidity: function swapTokenForEth(address _to, uint256 _amountIn, uint256 mid, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Common *CommonTransactorSession) SwapTokenForEth(_to common.Address, _amountIn *big.Int, mid *big.Int, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SwapTokenForEth(&_Common.TransactOpts, _to, _amountIn, mid, gas, routerAddr, WethAddr, deadline)
}

// SwapTokenWithPath is a paid mutator transaction binding the contract method 0xc788a933.
//
// Solidity: function swapTokenWithPath(address _to, uint256 _amountIn, uint256 mid, uint256 gas, address routerAddr, address WethAddr, address[] path, uint256 deadline) payable returns()
func (_Common *CommonTransactor) SwapTokenWithPath(opts *bind.TransactOpts, _to common.Address, _amountIn *big.Int, mid *big.Int, gas *big.Int, routerAddr common.Address, WethAddr common.Address, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "swapTokenWithPath", _to, _amountIn, mid, gas, routerAddr, WethAddr, path, deadline)
}

// SwapTokenWithPath is a paid mutator transaction binding the contract method 0xc788a933.
//
// Solidity: function swapTokenWithPath(address _to, uint256 _amountIn, uint256 mid, uint256 gas, address routerAddr, address WethAddr, address[] path, uint256 deadline) payable returns()
func (_Common *CommonSession) SwapTokenWithPath(_to common.Address, _amountIn *big.Int, mid *big.Int, gas *big.Int, routerAddr common.Address, WethAddr common.Address, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SwapTokenWithPath(&_Common.TransactOpts, _to, _amountIn, mid, gas, routerAddr, WethAddr, path, deadline)
}

// SwapTokenWithPath is a paid mutator transaction binding the contract method 0xc788a933.
//
// Solidity: function swapTokenWithPath(address _to, uint256 _amountIn, uint256 mid, uint256 gas, address routerAddr, address WethAddr, address[] path, uint256 deadline) payable returns()
func (_Common *CommonTransactorSession) SwapTokenWithPath(_to common.Address, _amountIn *big.Int, mid *big.Int, gas *big.Int, routerAddr common.Address, WethAddr common.Address, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SwapTokenWithPath(&_Common.TransactOpts, _to, _amountIn, mid, gas, routerAddr, WethAddr, path, deadline)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Common *CommonTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Common *CommonSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Common.Contract.TransferOwnership(&_Common.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Common *CommonTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Common.Contract.TransferOwnership(&_Common.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Common *CommonTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Common.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Common *CommonSession) Receive() (*types.Transaction, error) {
	return _Common.Contract.Receive(&_Common.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Common *CommonTransactorSession) Receive() (*types.Transaction, error) {
	return _Common.Contract.Receive(&_Common.TransactOpts)
}

// CommonBurnTokenIterator is returned from FilterBurnToken and is used to iterate over the raw logs and unpacked data for BurnToken events raised by the Common contract.
type CommonBurnTokenIterator struct {
	Event *CommonBurnToken // Event containing the contract specifics and raw log

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
func (it *CommonBurnTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonBurnToken)
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
		it.Event = new(CommonBurnToken)
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
func (it *CommonBurnTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonBurnTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonBurnToken represents a BurnToken event raised by the Common contract.
type CommonBurnToken struct {
	To      common.Address
	Amount  *big.Int
	Ntype   *big.Int
	ToToken string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurnToken is a free log retrieval operation binding the contract event 0x86f32d6c7a935bd338ee00610630fcfb6f043a6ad755db62064ce2ad92c45caa.
//
// Solidity: event BurnToken(address indexed to, uint256 amount, uint256 ntype, string toToken)
func (_Common *CommonFilterer) FilterBurnToken(opts *bind.FilterOpts, to []common.Address) (*CommonBurnTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "BurnToken", toRule)
	if err != nil {
		return nil, err
	}
	return &CommonBurnTokenIterator{contract: _Common.contract, event: "BurnToken", logs: logs, sub: sub}, nil
}

// WatchBurnToken is a free log subscription operation binding the contract event 0x86f32d6c7a935bd338ee00610630fcfb6f043a6ad755db62064ce2ad92c45caa.
//
// Solidity: event BurnToken(address indexed to, uint256 amount, uint256 ntype, string toToken)
func (_Common *CommonFilterer) WatchBurnToken(opts *bind.WatchOpts, sink chan<- *CommonBurnToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "BurnToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonBurnToken)
				if err := _Common.contract.UnpackLog(event, "BurnToken", log); err != nil {
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
func (_Common *CommonFilterer) ParseBurnToken(log types.Log) (*CommonBurnToken, error) {
	event := new(CommonBurnToken)
	if err := _Common.contract.UnpackLog(event, "BurnToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonMintItemCreatedIterator is returned from FilterMintItemCreated and is used to iterate over the raw logs and unpacked data for MintItemCreated events raised by the Common contract.
type CommonMintItemCreatedIterator struct {
	Event *CommonMintItemCreated // Event containing the contract specifics and raw log

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
func (it *CommonMintItemCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonMintItemCreated)
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
		it.Event = new(CommonMintItemCreated)
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
func (it *CommonMintItemCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonMintItemCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonMintItemCreated represents a MintItemCreated event raised by the Common contract.
type CommonMintItemCreated struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	MId    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintItemCreated is a free log retrieval operation binding the contract event 0x82911cbfb5e991769299a8d7cf2c610644a7c6ed1349882285349faed8ca7bb8.
//
// Solidity: event MintItemCreated(address indexed from, address to, uint256 amount, uint256 mId)
func (_Common *CommonFilterer) FilterMintItemCreated(opts *bind.FilterOpts, from []common.Address) (*CommonMintItemCreatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "MintItemCreated", fromRule)
	if err != nil {
		return nil, err
	}
	return &CommonMintItemCreatedIterator{contract: _Common.contract, event: "MintItemCreated", logs: logs, sub: sub}, nil
}

// WatchMintItemCreated is a free log subscription operation binding the contract event 0x82911cbfb5e991769299a8d7cf2c610644a7c6ed1349882285349faed8ca7bb8.
//
// Solidity: event MintItemCreated(address indexed from, address to, uint256 amount, uint256 mId)
func (_Common *CommonFilterer) WatchMintItemCreated(opts *bind.WatchOpts, sink chan<- *CommonMintItemCreated, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "MintItemCreated", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonMintItemCreated)
				if err := _Common.contract.UnpackLog(event, "MintItemCreated", log); err != nil {
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
func (_Common *CommonFilterer) ParseMintItemCreated(log types.Log) (*CommonMintItemCreated, error) {
	event := new(CommonMintItemCreated)
	if err := _Common.contract.UnpackLog(event, "MintItemCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonMintTokenIterator is returned from FilterMintToken and is used to iterate over the raw logs and unpacked data for MintToken events raised by the Common contract.
type CommonMintTokenIterator struct {
	Event *CommonMintToken // Event containing the contract specifics and raw log

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
func (it *CommonMintTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonMintToken)
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
		it.Event = new(CommonMintToken)
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
func (it *CommonMintTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonMintTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonMintToken represents a MintToken event raised by the Common contract.
type CommonMintToken struct {
	To       common.Address
	Amount   *big.Int
	Mid      *big.Int
	AmountIn *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintToken is a free log retrieval operation binding the contract event 0x8fb5c7bffbb272c541556c455c74269997b816df24f56dd255c2391d92d4f1e9.
//
// Solidity: event MintToken(address indexed to, uint256 amount, uint256 mid, uint256 amountIn)
func (_Common *CommonFilterer) FilterMintToken(opts *bind.FilterOpts, to []common.Address) (*CommonMintTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "MintToken", toRule)
	if err != nil {
		return nil, err
	}
	return &CommonMintTokenIterator{contract: _Common.contract, event: "MintToken", logs: logs, sub: sub}, nil
}

// WatchMintToken is a free log subscription operation binding the contract event 0x8fb5c7bffbb272c541556c455c74269997b816df24f56dd255c2391d92d4f1e9.
//
// Solidity: event MintToken(address indexed to, uint256 amount, uint256 mid, uint256 amountIn)
func (_Common *CommonFilterer) WatchMintToken(opts *bind.WatchOpts, sink chan<- *CommonMintToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "MintToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonMintToken)
				if err := _Common.contract.UnpackLog(event, "MintToken", log); err != nil {
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
func (_Common *CommonFilterer) ParseMintToken(log types.Log) (*CommonMintToken, error) {
	event := new(CommonMintToken)
	if err := _Common.contract.UnpackLog(event, "MintToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Common contract.
type CommonOwnershipTransferredIterator struct {
	Event *CommonOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CommonOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonOwnershipTransferred)
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
		it.Event = new(CommonOwnershipTransferred)
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
func (it *CommonOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonOwnershipTransferred represents a OwnershipTransferred event raised by the Common contract.
type CommonOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Common *CommonFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CommonOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CommonOwnershipTransferredIterator{contract: _Common.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Common *CommonFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommonOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonOwnershipTransferred)
				if err := _Common.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Common *CommonFilterer) ParseOwnershipTransferred(log types.Log) (*CommonOwnershipTransferred, error) {
	event := new(CommonOwnershipTransferred)
	if err := _Common.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonSwapTokenIterator is returned from FilterSwapToken and is used to iterate over the raw logs and unpacked data for SwapToken events raised by the Common contract.
type CommonSwapTokenIterator struct {
	Event *CommonSwapToken // Event containing the contract specifics and raw log

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
func (it *CommonSwapTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonSwapToken)
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
		it.Event = new(CommonSwapToken)
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
func (it *CommonSwapTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonSwapTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonSwapToken represents a SwapToken event raised by the Common contract.
type CommonSwapToken struct {
	To        common.Address
	InAmount  *big.Int
	OutAmount *big.Int
	Flag      string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapToken is a free log retrieval operation binding the contract event 0x59472f3e5abe9ecfcb0d72f4334caeb814fd3a06d2b049f049b039d14c663f8c.
//
// Solidity: event SwapToken(address indexed to, uint256 inAmount, uint256 outAmount, string flag)
func (_Common *CommonFilterer) FilterSwapToken(opts *bind.FilterOpts, to []common.Address) (*CommonSwapTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "SwapToken", toRule)
	if err != nil {
		return nil, err
	}
	return &CommonSwapTokenIterator{contract: _Common.contract, event: "SwapToken", logs: logs, sub: sub}, nil
}

// WatchSwapToken is a free log subscription operation binding the contract event 0x59472f3e5abe9ecfcb0d72f4334caeb814fd3a06d2b049f049b039d14c663f8c.
//
// Solidity: event SwapToken(address indexed to, uint256 inAmount, uint256 outAmount, string flag)
func (_Common *CommonFilterer) WatchSwapToken(opts *bind.WatchOpts, sink chan<- *CommonSwapToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "SwapToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonSwapToken)
				if err := _Common.contract.UnpackLog(event, "SwapToken", log); err != nil {
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
func (_Common *CommonFilterer) ParseSwapToken(log types.Log) (*CommonSwapToken, error) {
	event := new(CommonSwapToken)
	if err := _Common.contract.UnpackLog(event, "SwapToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonTransferTokenIterator is returned from FilterTransferToken and is used to iterate over the raw logs and unpacked data for TransferToken events raised by the Common contract.
type CommonTransferTokenIterator struct {
	Event *CommonTransferToken // Event containing the contract specifics and raw log

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
func (it *CommonTransferTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonTransferToken)
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
		it.Event = new(CommonTransferToken)
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
func (it *CommonTransferTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonTransferTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonTransferToken represents a TransferToken event raised by the Common contract.
type CommonTransferToken struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferToken is a free log retrieval operation binding the contract event 0x9dd3045b6df532ed81beb2a333cec6249dafd3c2fc54c80c50155cb0e1a0ba1e.
//
// Solidity: event TransferToken(address indexed to, uint256 amount)
func (_Common *CommonFilterer) FilterTransferToken(opts *bind.FilterOpts, to []common.Address) (*CommonTransferTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "TransferToken", toRule)
	if err != nil {
		return nil, err
	}
	return &CommonTransferTokenIterator{contract: _Common.contract, event: "TransferToken", logs: logs, sub: sub}, nil
}

// WatchTransferToken is a free log subscription operation binding the contract event 0x9dd3045b6df532ed81beb2a333cec6249dafd3c2fc54c80c50155cb0e1a0ba1e.
//
// Solidity: event TransferToken(address indexed to, uint256 amount)
func (_Common *CommonFilterer) WatchTransferToken(opts *bind.WatchOpts, sink chan<- *CommonTransferToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "TransferToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonTransferToken)
				if err := _Common.contract.UnpackLog(event, "TransferToken", log); err != nil {
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
func (_Common *CommonFilterer) ParseTransferToken(log types.Log) (*CommonTransferToken, error) {
	event := new(CommonTransferToken)
	if err := _Common.contract.UnpackLog(event, "TransferToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
