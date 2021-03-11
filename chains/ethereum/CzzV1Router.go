// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// EthereumABI is the input ABI used to generate the binding from.
const EthereumABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"}],\"name\":\"BurnToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mId\",\"type\":\"uint256\"}],\"name\":\"MintItemCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"MintToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"}],\"name\":\"SwapToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCzzTonkenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCzzTonkenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"value\",\"type\":\"uint8\"}],\"name\":\"setMinSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WethAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapAndBurn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountInMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ntype\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toToken\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WethAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapAndBurnEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WethAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WethAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokenForEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"}],\"name\":\"swap_burn_get_amount\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"swap_burn_get_getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"routerAddr\",\"type\":\"address\"}],\"name\":\"swap_mint_get_amount\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Ethereum is an auto generated Go binding around an Ethereum contract.
type Ethereum struct {
	EthereumCaller     // Read-only binding to the contract
	EthereumTransactor // Write-only binding to the contract
	EthereumFilterer   // Log filterer for contract events
}

// EthereumCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthereumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthereumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthereumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthereumSession struct {
	Contract     *Ethereum         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthereumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthereumCallerSession struct {
	Contract *EthereumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EthereumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthereumTransactorSession struct {
	Contract     *EthereumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthereumRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthereumRaw struct {
	Contract *Ethereum // Generic contract binding to access the raw methods on
}

// EthereumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthereumCallerRaw struct {
	Contract *EthereumCaller // Generic read-only contract binding to access the raw methods on
}

// EthereumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthereumTransactorRaw struct {
	Contract *EthereumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthereum creates a new instance of Ethereum, bound to a specific deployed contract.
func NewEthereum(address common.Address, backend bind.ContractBackend) (*Ethereum, error) {
	contract, err := bindEthereum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethereum{EthereumCaller: EthereumCaller{contract: contract}, EthereumTransactor: EthereumTransactor{contract: contract}, EthereumFilterer: EthereumFilterer{contract: contract}}, nil
}

// NewEthereumCaller creates a new read-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumCaller(address common.Address, caller bind.ContractCaller) (*EthereumCaller, error) {
	contract, err := bindEthereum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumCaller{contract: contract}, nil
}

// NewEthereumTransactor creates a new write-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumTransactor(address common.Address, transactor bind.ContractTransactor) (*EthereumTransactor, error) {
	contract, err := bindEthereum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumTransactor{contract: contract}, nil
}

// NewEthereumFilterer creates a new log filterer instance of Ethereum, bound to a specific deployed contract.
func NewEthereumFilterer(address common.Address, filterer bind.ContractFilterer) (*EthereumFilterer, error) {
	contract, err := bindEthereum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthereumFilterer{contract: contract}, nil
}

// bindEthereum binds a generic wrapper to an already deployed contract.
func bindEthereum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthereumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.EthereumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transact(opts, method, params...)
}

// GetCzzTonkenAddress is a free data retrieval call binding the contract method 0x9078f506.
//
// Solidity: function getCzzTonkenAddress() view returns(address)
func (_Ethereum *EthereumCaller) GetCzzTonkenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "getCzzTonkenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCzzTonkenAddress is a free data retrieval call binding the contract method 0x9078f506.
//
// Solidity: function getCzzTonkenAddress() view returns(address)
func (_Ethereum *EthereumSession) GetCzzTonkenAddress() (common.Address, error) {
	return _Ethereum.Contract.GetCzzTonkenAddress(&_Ethereum.CallOpts)
}

// GetCzzTonkenAddress is a free data retrieval call binding the contract method 0x9078f506.
//
// Solidity: function getCzzTonkenAddress() view returns(address)
func (_Ethereum *EthereumCallerSession) GetCzzTonkenAddress() (common.Address, error) {
	return _Ethereum.Contract.GetCzzTonkenAddress(&_Ethereum.CallOpts)
}

// GetMinSignatures is a free data retrieval call binding the contract method 0x1eb36178.
//
// Solidity: function getMinSignatures() view returns(uint256)
func (_Ethereum *EthereumCaller) GetMinSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "getMinSignatures")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinSignatures is a free data retrieval call binding the contract method 0x1eb36178.
//
// Solidity: function getMinSignatures() view returns(uint256)
func (_Ethereum *EthereumSession) GetMinSignatures() (*big.Int, error) {
	return _Ethereum.Contract.GetMinSignatures(&_Ethereum.CallOpts)
}

// GetMinSignatures is a free data retrieval call binding the contract method 0x1eb36178.
//
// Solidity: function getMinSignatures() view returns(uint256)
func (_Ethereum *EthereumCallerSession) GetMinSignatures() (*big.Int, error) {
	return _Ethereum.Contract.GetMinSignatures(&_Ethereum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethereum *EthereumCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethereum *EthereumSession) Owner() (common.Address, error) {
	return _Ethereum.Contract.Owner(&_Ethereum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethereum *EthereumCallerSession) Owner() (common.Address, error) {
	return _Ethereum.Contract.Owner(&_Ethereum.CallOpts)
}

// SwapBurnGetAmount is a free data retrieval call binding the contract method 0x8509e10b.
//
// Solidity: function swap_burn_get_amount(uint256 amountIn, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Ethereum *EthereumCaller) SwapBurnGetAmount(opts *bind.CallOpts, amountIn *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "swap_burn_get_amount", amountIn, path, routerAddr)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SwapBurnGetAmount is a free data retrieval call binding the contract method 0x8509e10b.
//
// Solidity: function swap_burn_get_amount(uint256 amountIn, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Ethereum *EthereumSession) SwapBurnGetAmount(amountIn *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	return _Ethereum.Contract.SwapBurnGetAmount(&_Ethereum.CallOpts, amountIn, path, routerAddr)
}

// SwapBurnGetAmount is a free data retrieval call binding the contract method 0x8509e10b.
//
// Solidity: function swap_burn_get_amount(uint256 amountIn, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Ethereum *EthereumCallerSession) SwapBurnGetAmount(amountIn *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	return _Ethereum.Contract.SwapBurnGetAmount(&_Ethereum.CallOpts, amountIn, path, routerAddr)
}

// SwapBurnGetGetReserves is a free data retrieval call binding the contract method 0x84ec7ee3.
//
// Solidity: function swap_burn_get_getReserves(address factory, address tokenA, address tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_Ethereum *EthereumCaller) SwapBurnGetGetReserves(opts *bind.CallOpts, factory common.Address, tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "swap_burn_get_getReserves", factory, tokenA, tokenB)

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
func (_Ethereum *EthereumSession) SwapBurnGetGetReserves(factory common.Address, tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Ethereum.Contract.SwapBurnGetGetReserves(&_Ethereum.CallOpts, factory, tokenA, tokenB)
}

// SwapBurnGetGetReserves is a free data retrieval call binding the contract method 0x84ec7ee3.
//
// Solidity: function swap_burn_get_getReserves(address factory, address tokenA, address tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_Ethereum *EthereumCallerSession) SwapBurnGetGetReserves(factory common.Address, tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Ethereum.Contract.SwapBurnGetGetReserves(&_Ethereum.CallOpts, factory, tokenA, tokenB)
}

// SwapMintGetAmount is a free data retrieval call binding the contract method 0x3c28d41f.
//
// Solidity: function swap_mint_get_amount(uint256 amountOut, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Ethereum *EthereumCaller) SwapMintGetAmount(opts *bind.CallOpts, amountOut *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "swap_mint_get_amount", amountOut, path, routerAddr)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SwapMintGetAmount is a free data retrieval call binding the contract method 0x3c28d41f.
//
// Solidity: function swap_mint_get_amount(uint256 amountOut, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Ethereum *EthereumSession) SwapMintGetAmount(amountOut *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	return _Ethereum.Contract.SwapMintGetAmount(&_Ethereum.CallOpts, amountOut, path, routerAddr)
}

// SwapMintGetAmount is a free data retrieval call binding the contract method 0x3c28d41f.
//
// Solidity: function swap_mint_get_amount(uint256 amountOut, address[] path, address routerAddr) view returns(uint256[] amounts)
func (_Ethereum *EthereumCallerSession) SwapMintGetAmount(amountOut *big.Int, path []common.Address, routerAddr common.Address) ([]*big.Int, error) {
	return _Ethereum.Contract.SwapMintGetAmount(&_Ethereum.CallOpts, amountOut, path, routerAddr)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Ethereum *EthereumTransactor) AddManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "addManager", manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Ethereum *EthereumSession) AddManager(manager common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.AddManager(&_Ethereum.TransactOpts, manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Ethereum *EthereumTransactorSession) AddManager(manager common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.AddManager(&_Ethereum.TransactOpts, manager)
}

// Burn is a paid mutator transaction binding the contract method 0xce2198a2.
//
// Solidity: function burn(uint256 _amountIn, uint256 ntype, string toToken) payable returns()
func (_Ethereum *EthereumTransactor) Burn(opts *bind.TransactOpts, _amountIn *big.Int, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "burn", _amountIn, ntype, toToken)
}

// Burn is a paid mutator transaction binding the contract method 0xce2198a2.
//
// Solidity: function burn(uint256 _amountIn, uint256 ntype, string toToken) payable returns()
func (_Ethereum *EthereumSession) Burn(_amountIn *big.Int, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Ethereum.Contract.Burn(&_Ethereum.TransactOpts, _amountIn, ntype, toToken)
}

// Burn is a paid mutator transaction binding the contract method 0xce2198a2.
//
// Solidity: function burn(uint256 _amountIn, uint256 ntype, string toToken) payable returns()
func (_Ethereum *EthereumTransactorSession) Burn(_amountIn *big.Int, ntype *big.Int, toToken string) (*types.Transaction, error) {
	return _Ethereum.Contract.Burn(&_Ethereum.TransactOpts, _amountIn, ntype, toToken)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address fromToken, uint256 _amountIn) payable returns()
func (_Ethereum *EthereumTransactor) Mint(opts *bind.TransactOpts, fromToken common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "mint", fromToken, _amountIn)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address fromToken, uint256 _amountIn) payable returns()
func (_Ethereum *EthereumSession) Mint(fromToken common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.Mint(&_Ethereum.TransactOpts, fromToken, _amountIn)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address fromToken, uint256 _amountIn) payable returns()
func (_Ethereum *EthereumTransactorSession) Mint(fromToken common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.Mint(&_Ethereum.TransactOpts, fromToken, _amountIn)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Ethereum *EthereumTransactor) RemoveManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "removeManager", manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Ethereum *EthereumSession) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.RemoveManager(&_Ethereum.TransactOpts, manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Ethereum *EthereumTransactorSession) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.RemoveManager(&_Ethereum.TransactOpts, manager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethereum *EthereumTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethereum *EthereumSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ethereum.Contract.RenounceOwnership(&_Ethereum.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethereum *EthereumTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ethereum.Contract.RenounceOwnership(&_Ethereum.TransactOpts)
}

// SetCzzTonkenAddress is a paid mutator transaction binding the contract method 0x44ad1e3d.
//
// Solidity: function setCzzTonkenAddress(address addr) returns()
func (_Ethereum *EthereumTransactor) SetCzzTonkenAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "setCzzTonkenAddress", addr)
}

// SetCzzTonkenAddress is a paid mutator transaction binding the contract method 0x44ad1e3d.
//
// Solidity: function setCzzTonkenAddress(address addr) returns()
func (_Ethereum *EthereumSession) SetCzzTonkenAddress(addr common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.SetCzzTonkenAddress(&_Ethereum.TransactOpts, addr)
}

// SetCzzTonkenAddress is a paid mutator transaction binding the contract method 0x44ad1e3d.
//
// Solidity: function setCzzTonkenAddress(address addr) returns()
func (_Ethereum *EthereumTransactorSession) SetCzzTonkenAddress(addr common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.SetCzzTonkenAddress(&_Ethereum.TransactOpts, addr)
}

// SetMinSignatures is a paid mutator transaction binding the contract method 0x745d314f.
//
// Solidity: function setMinSignatures(uint8 value) returns()
func (_Ethereum *EthereumTransactor) SetMinSignatures(opts *bind.TransactOpts, value uint8) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "setMinSignatures", value)
}

// SetMinSignatures is a paid mutator transaction binding the contract method 0x745d314f.
//
// Solidity: function setMinSignatures(uint8 value) returns()
func (_Ethereum *EthereumSession) SetMinSignatures(value uint8) (*types.Transaction, error) {
	return _Ethereum.Contract.SetMinSignatures(&_Ethereum.TransactOpts, value)
}

// SetMinSignatures is a paid mutator transaction binding the contract method 0x745d314f.
//
// Solidity: function setMinSignatures(uint8 value) returns()
func (_Ethereum *EthereumTransactorSession) SetMinSignatures(value uint8) (*types.Transaction, error) {
	return _Ethereum.Contract.SetMinSignatures(&_Ethereum.TransactOpts, value)
}

// SwapAndBurn is a paid mutator transaction binding the contract method 0xa6d2962b.
//
// Solidity: function swapAndBurn(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Ethereum *EthereumTransactor) SwapAndBurn(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "swapAndBurn", _amountIn, _amountOutMin, fromToken, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurn is a paid mutator transaction binding the contract method 0xa6d2962b.
//
// Solidity: function swapAndBurn(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Ethereum *EthereumSession) SwapAndBurn(_amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SwapAndBurn(&_Ethereum.TransactOpts, _amountIn, _amountOutMin, fromToken, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurn is a paid mutator transaction binding the contract method 0xa6d2962b.
//
// Solidity: function swapAndBurn(uint256 _amountIn, uint256 _amountOutMin, address fromToken, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Ethereum *EthereumTransactorSession) SwapAndBurn(_amountIn *big.Int, _amountOutMin *big.Int, fromToken common.Address, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SwapAndBurn(&_Ethereum.TransactOpts, _amountIn, _amountOutMin, fromToken, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurnEth is a paid mutator transaction binding the contract method 0xe11e1a80.
//
// Solidity: function swapAndBurnEth(uint256 _amountInMin, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Ethereum *EthereumTransactor) SwapAndBurnEth(opts *bind.TransactOpts, _amountInMin *big.Int, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "swapAndBurnEth", _amountInMin, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurnEth is a paid mutator transaction binding the contract method 0xe11e1a80.
//
// Solidity: function swapAndBurnEth(uint256 _amountInMin, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Ethereum *EthereumSession) SwapAndBurnEth(_amountInMin *big.Int, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SwapAndBurnEth(&_Ethereum.TransactOpts, _amountInMin, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapAndBurnEth is a paid mutator transaction binding the contract method 0xe11e1a80.
//
// Solidity: function swapAndBurnEth(uint256 _amountInMin, uint256 ntype, string toToken, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Ethereum *EthereumTransactorSession) SwapAndBurnEth(_amountInMin *big.Int, ntype *big.Int, toToken string, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SwapAndBurnEth(&_Ethereum.TransactOpts, _amountInMin, ntype, toToken, routerAddr, WethAddr, deadline)
}

// SwapToken is a paid mutator transaction binding the contract method 0x1b1046bd.
//
// Solidity: function swapToken(address _to, uint256 _amountIn, uint256 mid, address toToken, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Ethereum *EthereumTransactor) SwapToken(opts *bind.TransactOpts, _to common.Address, _amountIn *big.Int, mid *big.Int, toToken common.Address, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "swapToken", _to, _amountIn, mid, toToken, gas, routerAddr, WethAddr, deadline)
}

// SwapToken is a paid mutator transaction binding the contract method 0x1b1046bd.
//
// Solidity: function swapToken(address _to, uint256 _amountIn, uint256 mid, address toToken, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Ethereum *EthereumSession) SwapToken(_to common.Address, _amountIn *big.Int, mid *big.Int, toToken common.Address, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SwapToken(&_Ethereum.TransactOpts, _to, _amountIn, mid, toToken, gas, routerAddr, WethAddr, deadline)
}

// SwapToken is a paid mutator transaction binding the contract method 0x1b1046bd.
//
// Solidity: function swapToken(address _to, uint256 _amountIn, uint256 mid, address toToken, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Ethereum *EthereumTransactorSession) SwapToken(_to common.Address, _amountIn *big.Int, mid *big.Int, toToken common.Address, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SwapToken(&_Ethereum.TransactOpts, _to, _amountIn, mid, toToken, gas, routerAddr, WethAddr, deadline)
}

// SwapTokenForEth is a paid mutator transaction binding the contract method 0xcabab781.
//
// Solidity: function swapTokenForEth(address _to, uint256 _amountIn, uint256 mid, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Ethereum *EthereumTransactor) SwapTokenForEth(opts *bind.TransactOpts, _to common.Address, _amountIn *big.Int, mid *big.Int, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "swapTokenForEth", _to, _amountIn, mid, gas, routerAddr, WethAddr, deadline)
}

// SwapTokenForEth is a paid mutator transaction binding the contract method 0xcabab781.
//
// Solidity: function swapTokenForEth(address _to, uint256 _amountIn, uint256 mid, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Ethereum *EthereumSession) SwapTokenForEth(_to common.Address, _amountIn *big.Int, mid *big.Int, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SwapTokenForEth(&_Ethereum.TransactOpts, _to, _amountIn, mid, gas, routerAddr, WethAddr, deadline)
}

// SwapTokenForEth is a paid mutator transaction binding the contract method 0xcabab781.
//
// Solidity: function swapTokenForEth(address _to, uint256 _amountIn, uint256 mid, uint256 gas, address routerAddr, address WethAddr, uint256 deadline) payable returns()
func (_Ethereum *EthereumTransactorSession) SwapTokenForEth(_to common.Address, _amountIn *big.Int, mid *big.Int, gas *big.Int, routerAddr common.Address, WethAddr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SwapTokenForEth(&_Ethereum.TransactOpts, _to, _amountIn, mid, gas, routerAddr, WethAddr, deadline)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethereum *EthereumTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethereum *EthereumSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.TransferOwnership(&_Ethereum.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethereum *EthereumTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.TransferOwnership(&_Ethereum.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ethereum *EthereumTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ethereum *EthereumSession) Receive() (*types.Transaction, error) {
	return _Ethereum.Contract.Receive(&_Ethereum.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ethereum *EthereumTransactorSession) Receive() (*types.Transaction, error) {
	return _Ethereum.Contract.Receive(&_Ethereum.TransactOpts)
}

// EthereumBurnTokenIterator is returned from FilterBurnToken and is used to iterate over the raw logs and unpacked data for BurnToken events raised by the Ethereum contract.
type EthereumBurnTokenIterator struct {
	Event *EthereumBurnToken // Event containing the contract specifics and raw log

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
func (it *EthereumBurnTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBurnToken)
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
		it.Event = new(EthereumBurnToken)
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
func (it *EthereumBurnTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBurnTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBurnToken represents a BurnToken event raised by the Ethereum contract.
type EthereumBurnToken struct {
	To      common.Address
	Amount  *big.Int
	Ntype   *big.Int
	ToToken string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurnToken is a free log retrieval operation binding the contract event 0x86f32d6c7a935bd338ee00610630fcfb6f043a6ad755db62064ce2ad92c45caa.
//
// Solidity: event BurnToken(address indexed to, uint256 amount, uint256 ntype, string toToken)
func (_Ethereum *EthereumFilterer) FilterBurnToken(opts *bind.FilterOpts, to []common.Address) (*EthereumBurnTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "BurnToken", toRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBurnTokenIterator{contract: _Ethereum.contract, event: "BurnToken", logs: logs, sub: sub}, nil
}

// WatchBurnToken is a free log subscription operation binding the contract event 0x86f32d6c7a935bd338ee00610630fcfb6f043a6ad755db62064ce2ad92c45caa.
//
// Solidity: event BurnToken(address indexed to, uint256 amount, uint256 ntype, string toToken)
func (_Ethereum *EthereumFilterer) WatchBurnToken(opts *bind.WatchOpts, sink chan<- *EthereumBurnToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "BurnToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBurnToken)
				if err := _Ethereum.contract.UnpackLog(event, "BurnToken", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParseBurnToken(log types.Log) (*EthereumBurnToken, error) {
	event := new(EthereumBurnToken)
	if err := _Ethereum.contract.UnpackLog(event, "BurnToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumMintItemCreatedIterator is returned from FilterMintItemCreated and is used to iterate over the raw logs and unpacked data for MintItemCreated events raised by the Ethereum contract.
type EthereumMintItemCreatedIterator struct {
	Event *EthereumMintItemCreated // Event containing the contract specifics and raw log

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
func (it *EthereumMintItemCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumMintItemCreated)
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
		it.Event = new(EthereumMintItemCreated)
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
func (it *EthereumMintItemCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumMintItemCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumMintItemCreated represents a MintItemCreated event raised by the Ethereum contract.
type EthereumMintItemCreated struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	MId    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintItemCreated is a free log retrieval operation binding the contract event 0x82911cbfb5e991769299a8d7cf2c610644a7c6ed1349882285349faed8ca7bb8.
//
// Solidity: event MintItemCreated(address indexed from, address to, uint256 amount, uint256 mId)
func (_Ethereum *EthereumFilterer) FilterMintItemCreated(opts *bind.FilterOpts, from []common.Address) (*EthereumMintItemCreatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "MintItemCreated", fromRule)
	if err != nil {
		return nil, err
	}
	return &EthereumMintItemCreatedIterator{contract: _Ethereum.contract, event: "MintItemCreated", logs: logs, sub: sub}, nil
}

// WatchMintItemCreated is a free log subscription operation binding the contract event 0x82911cbfb5e991769299a8d7cf2c610644a7c6ed1349882285349faed8ca7bb8.
//
// Solidity: event MintItemCreated(address indexed from, address to, uint256 amount, uint256 mId)
func (_Ethereum *EthereumFilterer) WatchMintItemCreated(opts *bind.WatchOpts, sink chan<- *EthereumMintItemCreated, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "MintItemCreated", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumMintItemCreated)
				if err := _Ethereum.contract.UnpackLog(event, "MintItemCreated", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParseMintItemCreated(log types.Log) (*EthereumMintItemCreated, error) {
	event := new(EthereumMintItemCreated)
	if err := _Ethereum.contract.UnpackLog(event, "MintItemCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumMintTokenIterator is returned from FilterMintToken and is used to iterate over the raw logs and unpacked data for MintToken events raised by the Ethereum contract.
type EthereumMintTokenIterator struct {
	Event *EthereumMintToken // Event containing the contract specifics and raw log

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
func (it *EthereumMintTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumMintToken)
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
		it.Event = new(EthereumMintToken)
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
func (it *EthereumMintTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumMintTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumMintToken represents a MintToken event raised by the Ethereum contract.
type EthereumMintToken struct {
	To       common.Address
	Amount   *big.Int
	Mid      *big.Int
	AmountIn *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintToken is a free log retrieval operation binding the contract event 0x8fb5c7bffbb272c541556c455c74269997b816df24f56dd255c2391d92d4f1e9.
//
// Solidity: event MintToken(address indexed to, uint256 amount, uint256 mid, uint256 amountIn)
func (_Ethereum *EthereumFilterer) FilterMintToken(opts *bind.FilterOpts, to []common.Address) (*EthereumMintTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "MintToken", toRule)
	if err != nil {
		return nil, err
	}
	return &EthereumMintTokenIterator{contract: _Ethereum.contract, event: "MintToken", logs: logs, sub: sub}, nil
}

// WatchMintToken is a free log subscription operation binding the contract event 0x8fb5c7bffbb272c541556c455c74269997b816df24f56dd255c2391d92d4f1e9.
//
// Solidity: event MintToken(address indexed to, uint256 amount, uint256 mid, uint256 amountIn)
func (_Ethereum *EthereumFilterer) WatchMintToken(opts *bind.WatchOpts, sink chan<- *EthereumMintToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "MintToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumMintToken)
				if err := _Ethereum.contract.UnpackLog(event, "MintToken", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParseMintToken(log types.Log) (*EthereumMintToken, error) {
	event := new(EthereumMintToken)
	if err := _Ethereum.contract.UnpackLog(event, "MintToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ethereum contract.
type EthereumOwnershipTransferredIterator struct {
	Event *EthereumOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EthereumOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumOwnershipTransferred)
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
		it.Event = new(EthereumOwnershipTransferred)
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
func (it *EthereumOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumOwnershipTransferred represents a OwnershipTransferred event raised by the Ethereum contract.
type EthereumOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethereum *EthereumFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthereumOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthereumOwnershipTransferredIterator{contract: _Ethereum.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethereum *EthereumFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthereumOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumOwnershipTransferred)
				if err := _Ethereum.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParseOwnershipTransferred(log types.Log) (*EthereumOwnershipTransferred, error) {
	event := new(EthereumOwnershipTransferred)
	if err := _Ethereum.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumSwapTokenIterator is returned from FilterSwapToken and is used to iterate over the raw logs and unpacked data for SwapToken events raised by the Ethereum contract.
type EthereumSwapTokenIterator struct {
	Event *EthereumSwapToken // Event containing the contract specifics and raw log

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
func (it *EthereumSwapTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumSwapToken)
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
		it.Event = new(EthereumSwapToken)
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
func (it *EthereumSwapTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumSwapTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumSwapToken represents a SwapToken event raised by the Ethereum contract.
type EthereumSwapToken struct {
	To        common.Address
	InAmount  *big.Int
	OutAmount *big.Int
	Flag      string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapToken is a free log retrieval operation binding the contract event 0x59472f3e5abe9ecfcb0d72f4334caeb814fd3a06d2b049f049b039d14c663f8c.
//
// Solidity: event SwapToken(address indexed to, uint256 inAmount, uint256 outAmount, string flag)
func (_Ethereum *EthereumFilterer) FilterSwapToken(opts *bind.FilterOpts, to []common.Address) (*EthereumSwapTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "SwapToken", toRule)
	if err != nil {
		return nil, err
	}
	return &EthereumSwapTokenIterator{contract: _Ethereum.contract, event: "SwapToken", logs: logs, sub: sub}, nil
}

// WatchSwapToken is a free log subscription operation binding the contract event 0x59472f3e5abe9ecfcb0d72f4334caeb814fd3a06d2b049f049b039d14c663f8c.
//
// Solidity: event SwapToken(address indexed to, uint256 inAmount, uint256 outAmount, string flag)
func (_Ethereum *EthereumFilterer) WatchSwapToken(opts *bind.WatchOpts, sink chan<- *EthereumSwapToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "SwapToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumSwapToken)
				if err := _Ethereum.contract.UnpackLog(event, "SwapToken", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParseSwapToken(log types.Log) (*EthereumSwapToken, error) {
	event := new(EthereumSwapToken)
	if err := _Ethereum.contract.UnpackLog(event, "SwapToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumTransferTokenIterator is returned from FilterTransferToken and is used to iterate over the raw logs and unpacked data for TransferToken events raised by the Ethereum contract.
type EthereumTransferTokenIterator struct {
	Event *EthereumTransferToken // Event containing the contract specifics and raw log

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
func (it *EthereumTransferTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumTransferToken)
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
		it.Event = new(EthereumTransferToken)
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
func (it *EthereumTransferTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumTransferTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumTransferToken represents a TransferToken event raised by the Ethereum contract.
type EthereumTransferToken struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferToken is a free log retrieval operation binding the contract event 0x9dd3045b6df532ed81beb2a333cec6249dafd3c2fc54c80c50155cb0e1a0ba1e.
//
// Solidity: event TransferToken(address indexed to, uint256 amount)
func (_Ethereum *EthereumFilterer) FilterTransferToken(opts *bind.FilterOpts, to []common.Address) (*EthereumTransferTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "TransferToken", toRule)
	if err != nil {
		return nil, err
	}
	return &EthereumTransferTokenIterator{contract: _Ethereum.contract, event: "TransferToken", logs: logs, sub: sub}, nil
}

// WatchTransferToken is a free log subscription operation binding the contract event 0x9dd3045b6df532ed81beb2a333cec6249dafd3c2fc54c80c50155cb0e1a0ba1e.
//
// Solidity: event TransferToken(address indexed to, uint256 amount)
func (_Ethereum *EthereumFilterer) WatchTransferToken(opts *bind.WatchOpts, sink chan<- *EthereumTransferToken, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "TransferToken", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumTransferToken)
				if err := _Ethereum.contract.UnpackLog(event, "TransferToken", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParseTransferToken(log types.Log) (*EthereumTransferToken, error) {
	event := new(EthereumTransferToken)
	if err := _Ethereum.contract.UnpackLog(event, "TransferToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
