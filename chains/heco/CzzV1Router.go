// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package heco

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

// HecoABI is the input ABI used to generate the binding from.
const HecoABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"}],\"name\":\"BurnToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mId\",\"type\":\"uint256\"}],\"name\":\"MintItemCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MintToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"}],\"name\":\"SwapToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"czzToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"value\",\"type\":\"uint8\"}],\"name\":\"setMinSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"}],\"name\":\"swapAndBurn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"}],\"name\":\"swapAndBurn1\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"}],\"name\":\"swapAndmint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"swap_burn_get_amount\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swap_mint_get_amount\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"swap_test\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Heco is an auto generated Go binding around an Ethereum contract.
type Heco struct {
	HecoCaller     // Read-only binding to the contract
	HecoTransactor // Write-only binding to the contract
	HecoFilterer   // Log filterer for contract events
}

// HecoCaller is an auto generated read-only Go binding around an Ethereum contract.
type HecoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HecoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HecoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HecoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HecoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HecoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HecoSession struct {
	Contract     *Heco             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HecoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HecoCallerSession struct {
	Contract *HecoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HecoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HecoTransactorSession struct {
	Contract     *HecoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HecoRaw is an auto generated low-level Go binding around an Ethereum contract.
type HecoRaw struct {
	Contract *Heco // Generic contract binding to access the raw methods on
}

// HecoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HecoCallerRaw struct {
	Contract *HecoCaller // Generic read-only contract binding to access the raw methods on
}

// HecoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HecoTransactorRaw struct {
	Contract *HecoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHeco creates a new instance of Heco, bound to a specific deployed contract.
func NewHeco(address common.Address, backend bind.ContractBackend) (*Heco, error) {
	contract, err := bindHeco(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Heco{HecoCaller: HecoCaller{contract: contract}, HecoTransactor: HecoTransactor{contract: contract}, HecoFilterer: HecoFilterer{contract: contract}}, nil
}

// NewHecoCaller creates a new read-only instance of Heco, bound to a specific deployed contract.
func NewHecoCaller(address common.Address, caller bind.ContractCaller) (*HecoCaller, error) {
	contract, err := bindHeco(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HecoCaller{contract: contract}, nil
}

// NewHecoTransactor creates a new write-only instance of Heco, bound to a specific deployed contract.
func NewHecoTransactor(address common.Address, transactor bind.ContractTransactor) (*HecoTransactor, error) {
	contract, err := bindHeco(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HecoTransactor{contract: contract}, nil
}

// NewHecoFilterer creates a new log filterer instance of Heco, bound to a specific deployed contract.
func NewHecoFilterer(address common.Address, filterer bind.ContractFilterer) (*HecoFilterer, error) {
	contract, err := bindHeco(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HecoFilterer{contract: contract}, nil
}

// bindHeco binds a generic wrapper to an already deployed contract.
func bindHeco(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HecoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Heco *HecoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Heco.Contract.HecoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Heco *HecoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heco.Contract.HecoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Heco *HecoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Heco.Contract.HecoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Heco *HecoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Heco.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Heco *HecoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heco.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Heco *HecoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Heco.Contract.contract.Transact(opts, method, params...)
}

// CzzToken is a free data retrieval call binding the contract method 0xd7bca30d.
//
// Solidity: function czzToken() view returns(address)
func (_Heco *HecoCaller) CzzToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Heco.contract.Call(opts, &out, "czzToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CzzToken is a free data retrieval call binding the contract method 0xd7bca30d.
//
// Solidity: function czzToken() view returns(address)
func (_Heco *HecoSession) CzzToken() (common.Address, error) {
	return _Heco.Contract.CzzToken(&_Heco.CallOpts)
}

// CzzToken is a free data retrieval call binding the contract method 0xd7bca30d.
//
// Solidity: function czzToken() view returns(address)
func (_Heco *HecoCallerSession) CzzToken() (common.Address, error) {
	return _Heco.Contract.CzzToken(&_Heco.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Heco *HecoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Heco.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Heco *HecoSession) Owner() (common.Address, error) {
	return _Heco.Contract.Owner(&_Heco.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Heco *HecoCallerSession) Owner() (common.Address, error) {
	return _Heco.Contract.Owner(&_Heco.CallOpts)
}

// SwapBurnGetAmount is a free data retrieval call binding the contract method 0x3bc5a841.
//
// Solidity: function swap_burn_get_amount(uint256 amountIn, address from) view returns(uint256[] amounts)
func (_Heco *HecoCaller) SwapBurnGetAmount(opts *bind.CallOpts, amountIn *big.Int, from common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Heco.contract.Call(opts, &out, "swap_burn_get_amount", amountIn, from)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SwapBurnGetAmount is a free data retrieval call binding the contract method 0x3bc5a841.
//
// Solidity: function swap_burn_get_amount(uint256 amountIn, address from) view returns(uint256[] amounts)
func (_Heco *HecoSession) SwapBurnGetAmount(amountIn *big.Int, from common.Address) ([]*big.Int, error) {
	return _Heco.Contract.SwapBurnGetAmount(&_Heco.CallOpts, amountIn, from)
}

// SwapBurnGetAmount is a free data retrieval call binding the contract method 0x3bc5a841.
//
// Solidity: function swap_burn_get_amount(uint256 amountIn, address from) view returns(uint256[] amounts)
func (_Heco *HecoCallerSession) SwapBurnGetAmount(amountIn *big.Int, from common.Address) ([]*big.Int, error) {
	return _Heco.Contract.SwapBurnGetAmount(&_Heco.CallOpts, amountIn, from)
}

// SwapMintGetAmount is a free data retrieval call binding the contract method 0xa6baa82b.
//
// Solidity: function swap_mint_get_amount(uint256 amountIn, address to) view returns(uint256[] amounts)
func (_Heco *HecoCaller) SwapMintGetAmount(opts *bind.CallOpts, amountIn *big.Int, to common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Heco.contract.Call(opts, &out, "swap_mint_get_amount", amountIn, to)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SwapMintGetAmount is a free data retrieval call binding the contract method 0xa6baa82b.
//
// Solidity: function swap_mint_get_amount(uint256 amountIn, address to) view returns(uint256[] amounts)
func (_Heco *HecoSession) SwapMintGetAmount(amountIn *big.Int, to common.Address) ([]*big.Int, error) {
	return _Heco.Contract.SwapMintGetAmount(&_Heco.CallOpts, amountIn, to)
}

// SwapMintGetAmount is a free data retrieval call binding the contract method 0xa6baa82b.
//
// Solidity: function swap_mint_get_amount(uint256 amountIn, address to) view returns(uint256[] amounts)
func (_Heco *HecoCallerSession) SwapMintGetAmount(amountIn *big.Int, to common.Address) ([]*big.Int, error) {
	return _Heco.Contract.SwapMintGetAmount(&_Heco.CallOpts, amountIn, to)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xca3d6539.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to) returns()
func (_Heco *HecoTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _Heco.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xca3d6539.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to) returns()
func (_Heco *HecoSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _Heco.Contract.AddLiquidity(&_Heco.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xca3d6539.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to) returns()
func (_Heco *HecoTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _Heco.Contract.AddLiquidity(&_Heco.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Heco *HecoTransactor) AddManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Heco.contract.Transact(opts, "addManager", manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Heco *HecoSession) AddManager(manager common.Address) (*types.Transaction, error) {
	return _Heco.Contract.AddManager(&_Heco.TransactOpts, manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Heco *HecoTransactorSession) AddManager(manager common.Address) (*types.Transaction, error) {
	return _Heco.Contract.AddManager(&_Heco.TransactOpts, manager)
}

// Burn is a paid mutator transaction binding the contract method 0xce2198a2.
//
// Solidity: function burn(uint256 _amountIn, uint256 ntype, string toToken) payable returns()
func (_Heco *HecoTransactor) Burn(opts *bind.TransactOpts, _amountIn *big.Int, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Heco.contract.Transact(opts, "burn", _amountIn, ntype, toToken)
}

// Burn is a paid mutator transaction binding the contract method 0xce2198a2.
//
// Solidity: function burn(uint256 _amountIn, uint256 ntype, string toToken) payable returns()
func (_Heco *HecoSession) Burn(_amountIn *big.Int, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Heco.Contract.Burn(&_Heco.TransactOpts, _amountIn, ntype, toToken)
}

// Burn is a paid mutator transaction binding the contract method 0xce2198a2.
//
// Solidity: function burn(uint256 _amountIn, uint256 ntype, string toToken) payable returns()
func (_Heco *HecoTransactorSession) Burn(_amountIn *big.Int, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Heco.Contract.Burn(&_Heco.TransactOpts, _amountIn, ntype, toToken)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address fromToken, uint256 _amountIn) payable returns()
func (_Heco *HecoTransactor) Mint(opts *bind.TransactOpts, fromToken common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Heco.contract.Transact(opts, "mint", fromToken, _amountIn)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address fromToken, uint256 _amountIn) payable returns()
func (_Heco *HecoSession) Mint(fromToken common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Heco.Contract.Mint(&_Heco.TransactOpts, fromToken, _amountIn)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address fromToken, uint256 _amountIn) payable returns()
func (_Heco *HecoTransactorSession) Mint(fromToken common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Heco.Contract.Mint(&_Heco.TransactOpts, fromToken, _amountIn)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Heco *HecoTransactor) RemoveManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Heco.contract.Transact(opts, "removeManager", manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Heco *HecoSession) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _Heco.Contract.RemoveManager(&_Heco.TransactOpts, manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Heco *HecoTransactorSession) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _Heco.Contract.RemoveManager(&_Heco.TransactOpts, manager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Heco *HecoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heco.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Heco *HecoSession) RenounceOwnership() (*types.Transaction, error) {
	return _Heco.Contract.RenounceOwnership(&_Heco.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Heco *HecoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Heco.Contract.RenounceOwnership(&_Heco.TransactOpts)
}

// SetMinSignatures is a paid mutator transaction binding the contract method 0x745d314f.
//
// Solidity: function setMinSignatures(uint8 value) returns()
func (_Heco *HecoTransactor) SetMinSignatures(opts *bind.TransactOpts, value uint8) (*types.Transaction, error) {
	return _Heco.contract.Transact(opts, "setMinSignatures", value)
}

// SetMinSignatures is a paid mutator transaction binding the contract method 0x745d314f.
//
// Solidity: function setMinSignatures(uint8 value) returns()
func (_Heco *HecoSession) SetMinSignatures(value uint8) (*types.Transaction, error) {
	return _Heco.Contract.SetMinSignatures(&_Heco.TransactOpts, value)
}

// SetMinSignatures is a paid mutator transaction binding the contract method 0x745d314f.
//
// Solidity: function setMinSignatures(uint8 value) returns()
func (_Heco *HecoTransactorSession) SetMinSignatures(value uint8) (*types.Transaction, error) {
	return _Heco.Contract.SetMinSignatures(&_Heco.TransactOpts, value)
}

// SwapAndBurn is a paid mutator transaction binding the contract method 0xe45fab1d.
//
// Solidity: function swapAndBurn(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken) payable returns()
func (_Heco *HecoTransactor) SwapAndBurn(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Heco.contract.Transact(opts, "swapAndBurn", _amountIn, _amountOutMin, fromToken, ntype, toToken)
}

// SwapAndBurn is a paid mutator transaction binding the contract method 0xe45fab1d.
//
// Solidity: function swapAndBurn(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken) payable returns()
func (_Heco *HecoSession) SwapAndBurn(_amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Heco.Contract.SwapAndBurn(&_Heco.TransactOpts, _amountIn, _amountOutMin, fromToken, ntype, toToken)
}

// SwapAndBurn is a paid mutator transaction binding the contract method 0xe45fab1d.
//
// Solidity: function swapAndBurn(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken) payable returns()
func (_Heco *HecoTransactorSession) SwapAndBurn(_amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Heco.Contract.SwapAndBurn(&_Heco.TransactOpts, _amountIn, _amountOutMin, fromToken, ntype, toToken)
}

// SwapAndBurn1 is a paid mutator transaction binding the contract method 0x832f74f7.
//
// Solidity: function swapAndBurn1(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken) payable returns()
func (_Heco *HecoTransactor) SwapAndBurn1(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Heco.contract.Transact(opts, "swapAndBurn1", _amountIn, _amountOutMin, fromToken, ntype, toToken)
}

// SwapAndBurn1 is a paid mutator transaction binding the contract method 0x832f74f7.
//
// Solidity: function swapAndBurn1(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken) payable returns()
func (_Heco *HecoSession) SwapAndBurn1(_amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Heco.Contract.SwapAndBurn1(&_Heco.TransactOpts, _amountIn, _amountOutMin, fromToken, ntype, toToken)
}

// SwapAndBurn1 is a paid mutator transaction binding the contract method 0x832f74f7.
//
// Solidity: function swapAndBurn1(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken) payable returns()
func (_Heco *HecoTransactorSession) SwapAndBurn1(_amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Heco.Contract.SwapAndBurn1(&_Heco.TransactOpts, _amountIn, _amountOutMin, fromToken, ntype, toToken)
}

// SwapAndmint is a paid mutator transaction binding the contract method 0x5c7fa3f5.
//
// Solidity: function swapAndmint(address _to, uint256 _amountIn, uint256 mid, address toToken) payable returns()
func (_Heco *HecoTransactor) SwapAndmint(opts *bind.TransactOpts, _to common.Address, _amountIn *big.Int, mid *big.Int, toToken common.Address) (*types.Transaction, error) {
	return _Heco.contract.Transact(opts, "swapAndmint", _to, _amountIn, mid, toToken)
}

// SwapAndmint is a paid mutator transaction binding the contract method 0x5c7fa3f5.
//
// Solidity: function swapAndmint(address _to, uint256 _amountIn, uint256 mid, address toToken) payable returns()
func (_Heco *HecoSession) SwapAndmint(_to common.Address, _amountIn *big.Int, mid *big.Int, toToken common.Address) (*types.Transaction, error) {
	return _Heco.Contract.SwapAndmint(&_Heco.TransactOpts, _to, _amountIn, mid, toToken)
}

// SwapAndmint is a paid mutator transaction binding the contract method 0x5c7fa3f5.
//
// Solidity: function swapAndmint(address _to, uint256 _amountIn, uint256 mid, address toToken) payable returns()
func (_Heco *HecoTransactorSession) SwapAndmint(_to common.Address, _amountIn *big.Int, mid *big.Int, toToken common.Address) (*types.Transaction, error) {
	return _Heco.Contract.SwapAndmint(&_Heco.TransactOpts, _to, _amountIn, mid, toToken)
}

// SwapTest is a paid mutator transaction binding the contract method 0x064136f5.
//
// Solidity: function swap_test(uint256 amountIn, uint256 amountOutMin, address from) payable returns()
func (_Heco *HecoTransactor) SwapTest(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, from common.Address) (*types.Transaction, error) {
	return _Heco.contract.Transact(opts, "swap_test", amountIn, amountOutMin, from)
}

// SwapTest is a paid mutator transaction binding the contract method 0x064136f5.
//
// Solidity: function swap_test(uint256 amountIn, uint256 amountOutMin, address from) payable returns()
func (_Heco *HecoSession) SwapTest(amountIn *big.Int, amountOutMin *big.Int, from common.Address) (*types.Transaction, error) {
	return _Heco.Contract.SwapTest(&_Heco.TransactOpts, amountIn, amountOutMin, from)
}

// SwapTest is a paid mutator transaction binding the contract method 0x064136f5.
//
// Solidity: function swap_test(uint256 amountIn, uint256 amountOutMin, address from) payable returns()
func (_Heco *HecoTransactorSession) SwapTest(amountIn *big.Int, amountOutMin *big.Int, from common.Address) (*types.Transaction, error) {
	return _Heco.Contract.SwapTest(&_Heco.TransactOpts, amountIn, amountOutMin, from)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Heco *HecoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Heco.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Heco *HecoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Heco.Contract.TransferOwnership(&_Heco.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Heco *HecoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Heco.Contract.TransferOwnership(&_Heco.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Heco *HecoTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heco.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Heco *HecoSession) Receive() (*types.Transaction, error) {
	return _Heco.Contract.Receive(&_Heco.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Heco *HecoTransactorSession) Receive() (*types.Transaction, error) {
	return _Heco.Contract.Receive(&_Heco.TransactOpts)
}

// HecoBurnTokenIterator is returned from FilterBurnToken and is used to iterate over the raw logs and unpacked data for BurnToken events raised by the Heco contract.
type HecoBurnTokenIterator struct {
	Event *HecoBurnToken // Event containing the contract specifics and raw log

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
func (it *HecoBurnTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HecoBurnToken)
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
		it.Event = new(HecoBurnToken)
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
func (it *HecoBurnTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HecoBurnTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HecoBurnToken represents a BurnToken event raised by the Heco contract.
type HecoBurnToken struct {
	To      common.Address
	Amount  *big.Int
	Ntype   *big.Int
	ToToken string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurnToken is a free log retrieval operation binding the contract event 0x86f32d6c7a935bd338ee00610630fcfb6f043a6ad755db62064ce2ad92c45caa.
//
// Solidity: event BurnToken(address indexed to, uint256 amount, uint256 ntype, string toToken)
func (_Heco *HecoFilterer) FilterBurnToken(opts *bind.FilterOpts, to []common.Address) (*HecoBurnTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Heco.contract.FilterLogs(opts, "BurnToken", toRule)
	if err != nil {
		return nil, err
	}
	return &HecoBurnTokenIterator{contract: _Heco.contract, event: "BurnToken", logs: logs, sub: sub}, nil
}

// WatchBurnToken is a free log subscription operation binding the contract event 0x86f32d6c7a935bd338ee00610630fcfb6f043a6ad755db62064ce2ad92c45caa.
//
// Solidity: event BurnToken(address indexed to, uint256 amount, uint256 ntype, string toToken)
func (_Heco *HecoFilterer) WatchBurnToken(opts *bind.WatchOpts, sink chan<- *HecoBurnToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Heco.contract.WatchLogs(opts, "BurnToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HecoBurnToken)
				if err := _Heco.contract.UnpackLog(event, "BurnToken", log); err != nil {
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
func (_Heco *HecoFilterer) ParseBurnToken(log types.Log) (*HecoBurnToken, error) {
	event := new(HecoBurnToken)
	if err := _Heco.contract.UnpackLog(event, "BurnToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HecoMintItemCreatedIterator is returned from FilterMintItemCreated and is used to iterate over the raw logs and unpacked data for MintItemCreated events raised by the Heco contract.
type HecoMintItemCreatedIterator struct {
	Event *HecoMintItemCreated // Event containing the contract specifics and raw log

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
func (it *HecoMintItemCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HecoMintItemCreated)
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
		it.Event = new(HecoMintItemCreated)
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
func (it *HecoMintItemCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HecoMintItemCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HecoMintItemCreated represents a MintItemCreated event raised by the Heco contract.
type HecoMintItemCreated struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	MId    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintItemCreated is a free log retrieval operation binding the contract event 0x82911cbfb5e991769299a8d7cf2c610644a7c6ed1349882285349faed8ca7bb8.
//
// Solidity: event MintItemCreated(address indexed from, address to, uint256 amount, uint256 mId)
func (_Heco *HecoFilterer) FilterMintItemCreated(opts *bind.FilterOpts, from []common.Address) (*HecoMintItemCreatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Heco.contract.FilterLogs(opts, "MintItemCreated", fromRule)
	if err != nil {
		return nil, err
	}
	return &HecoMintItemCreatedIterator{contract: _Heco.contract, event: "MintItemCreated", logs: logs, sub: sub}, nil
}

// WatchMintItemCreated is a free log subscription operation binding the contract event 0x82911cbfb5e991769299a8d7cf2c610644a7c6ed1349882285349faed8ca7bb8.
//
// Solidity: event MintItemCreated(address indexed from, address to, uint256 amount, uint256 mId)
func (_Heco *HecoFilterer) WatchMintItemCreated(opts *bind.WatchOpts, sink chan<- *HecoMintItemCreated, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Heco.contract.WatchLogs(opts, "MintItemCreated", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HecoMintItemCreated)
				if err := _Heco.contract.UnpackLog(event, "MintItemCreated", log); err != nil {
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
func (_Heco *HecoFilterer) ParseMintItemCreated(log types.Log) (*HecoMintItemCreated, error) {
	event := new(HecoMintItemCreated)
	if err := _Heco.contract.UnpackLog(event, "MintItemCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HecoMintTokenIterator is returned from FilterMintToken and is used to iterate over the raw logs and unpacked data for MintToken events raised by the Heco contract.
type HecoMintTokenIterator struct {
	Event *HecoMintToken // Event containing the contract specifics and raw log

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
func (it *HecoMintTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HecoMintToken)
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
		it.Event = new(HecoMintToken)
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
func (it *HecoMintTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HecoMintTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HecoMintToken represents a MintToken event raised by the Heco contract.
type HecoMintToken struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintToken is a free log retrieval operation binding the contract event 0xdcdea898caf5576419f89caf69301592a4758349a0bd62300b58849213420a72.
//
// Solidity: event MintToken(address indexed to, uint256 amount)
func (_Heco *HecoFilterer) FilterMintToken(opts *bind.FilterOpts, to []common.Address) (*HecoMintTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Heco.contract.FilterLogs(opts, "MintToken", toRule)
	if err != nil {
		return nil, err
	}
	return &HecoMintTokenIterator{contract: _Heco.contract, event: "MintToken", logs: logs, sub: sub}, nil
}

// WatchMintToken is a free log subscription operation binding the contract event 0xdcdea898caf5576419f89caf69301592a4758349a0bd62300b58849213420a72.
//
// Solidity: event MintToken(address indexed to, uint256 amount)
func (_Heco *HecoFilterer) WatchMintToken(opts *bind.WatchOpts, sink chan<- *HecoMintToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Heco.contract.WatchLogs(opts, "MintToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HecoMintToken)
				if err := _Heco.contract.UnpackLog(event, "MintToken", log); err != nil {
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

// ParseMintToken is a log parse operation binding the contract event 0xdcdea898caf5576419f89caf69301592a4758349a0bd62300b58849213420a72.
//
// Solidity: event MintToken(address indexed to, uint256 amount)
func (_Heco *HecoFilterer) ParseMintToken(log types.Log) (*HecoMintToken, error) {
	event := new(HecoMintToken)
	if err := _Heco.contract.UnpackLog(event, "MintToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HecoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Heco contract.
type HecoOwnershipTransferredIterator struct {
	Event *HecoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HecoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HecoOwnershipTransferred)
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
		it.Event = new(HecoOwnershipTransferred)
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
func (it *HecoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HecoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HecoOwnershipTransferred represents a OwnershipTransferred event raised by the Heco contract.
type HecoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Heco *HecoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HecoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Heco.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HecoOwnershipTransferredIterator{contract: _Heco.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Heco *HecoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HecoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Heco.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HecoOwnershipTransferred)
				if err := _Heco.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Heco *HecoFilterer) ParseOwnershipTransferred(log types.Log) (*HecoOwnershipTransferred, error) {
	event := new(HecoOwnershipTransferred)
	if err := _Heco.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HecoSwapTokenIterator is returned from FilterSwapToken and is used to iterate over the raw logs and unpacked data for SwapToken events raised by the Heco contract.
type HecoSwapTokenIterator struct {
	Event *HecoSwapToken // Event containing the contract specifics and raw log

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
func (it *HecoSwapTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HecoSwapToken)
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
		it.Event = new(HecoSwapToken)
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
func (it *HecoSwapTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HecoSwapTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HecoSwapToken represents a SwapToken event raised by the Heco contract.
type HecoSwapToken struct {
	To        common.Address
	InAmount  *big.Int
	OutAmount *big.Int
	Flag      string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapToken is a free log retrieval operation binding the contract event 0x59472f3e5abe9ecfcb0d72f4334caeb814fd3a06d2b049f049b039d14c663f8c.
//
// Solidity: event SwapToken(address indexed to, uint256 inAmount, uint256 outAmount, string flag)
func (_Heco *HecoFilterer) FilterSwapToken(opts *bind.FilterOpts, to []common.Address) (*HecoSwapTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Heco.contract.FilterLogs(opts, "SwapToken", toRule)
	if err != nil {
		return nil, err
	}
	return &HecoSwapTokenIterator{contract: _Heco.contract, event: "SwapToken", logs: logs, sub: sub}, nil
}

// WatchSwapToken is a free log subscription operation binding the contract event 0x59472f3e5abe9ecfcb0d72f4334caeb814fd3a06d2b049f049b039d14c663f8c.
//
// Solidity: event SwapToken(address indexed to, uint256 inAmount, uint256 outAmount, string flag)
func (_Heco *HecoFilterer) WatchSwapToken(opts *bind.WatchOpts, sink chan<- *HecoSwapToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Heco.contract.WatchLogs(opts, "SwapToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HecoSwapToken)
				if err := _Heco.contract.UnpackLog(event, "SwapToken", log); err != nil {
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
func (_Heco *HecoFilterer) ParseSwapToken(log types.Log) (*HecoSwapToken, error) {
	event := new(HecoSwapToken)
	if err := _Heco.contract.UnpackLog(event, "SwapToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HecoTransferTokenIterator is returned from FilterTransferToken and is used to iterate over the raw logs and unpacked data for TransferToken events raised by the Heco contract.
type HecoTransferTokenIterator struct {
	Event *HecoTransferToken // Event containing the contract specifics and raw log

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
func (it *HecoTransferTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HecoTransferToken)
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
		it.Event = new(HecoTransferToken)
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
func (it *HecoTransferTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HecoTransferTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HecoTransferToken represents a TransferToken event raised by the Heco contract.
type HecoTransferToken struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferToken is a free log retrieval operation binding the contract event 0x9dd3045b6df532ed81beb2a333cec6249dafd3c2fc54c80c50155cb0e1a0ba1e.
//
// Solidity: event TransferToken(address indexed to, uint256 amount)
func (_Heco *HecoFilterer) FilterTransferToken(opts *bind.FilterOpts, to []common.Address) (*HecoTransferTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Heco.contract.FilterLogs(opts, "TransferToken", toRule)
	if err != nil {
		return nil, err
	}
	return &HecoTransferTokenIterator{contract: _Heco.contract, event: "TransferToken", logs: logs, sub: sub}, nil
}

// WatchTransferToken is a free log subscription operation binding the contract event 0x9dd3045b6df532ed81beb2a333cec6249dafd3c2fc54c80c50155cb0e1a0ba1e.
//
// Solidity: event TransferToken(address indexed to, uint256 amount)
func (_Heco *HecoFilterer) WatchTransferToken(opts *bind.WatchOpts, sink chan<- *HecoTransferToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Heco.contract.WatchLogs(opts, "TransferToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HecoTransferToken)
				if err := _Heco.contract.UnpackLog(event, "TransferToken", log); err != nil {
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
func (_Heco *HecoFilterer) ParseTransferToken(log types.Log) (*HecoTransferToken, error) {
	event := new(HecoTransferToken)
	if err := _Heco.contract.UnpackLog(event, "TransferToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
