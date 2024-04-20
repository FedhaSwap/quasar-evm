// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// QuasarCurrency is an auto generated low-level Go binding around an user-defined struct.
type QuasarCurrency struct {
	Name   string
	Symbol string
}

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"CurrencyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"CurrencyStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"CurrencyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"addCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"changeCurrencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"getCurrencyID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"getCurrencyMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structQuasar.Currency\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedCurrencies\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structQuasar.Currency[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"isCurrencySupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"pushPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"updateCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// GetCurrencyID is a free data retrieval call binding the contract method 0x1c3fcee0.
//
// Solidity: function getCurrencyID(string symbol) view returns(uint64, bool)
func (_Oracle *OracleCaller) GetCurrencyID(opts *bind.CallOpts, symbol string) (uint64, bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getCurrencyID", symbol)

	if err != nil {
		return *new(uint64), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetCurrencyID is a free data retrieval call binding the contract method 0x1c3fcee0.
//
// Solidity: function getCurrencyID(string symbol) view returns(uint64, bool)
func (_Oracle *OracleSession) GetCurrencyID(symbol string) (uint64, bool, error) {
	return _Oracle.Contract.GetCurrencyID(&_Oracle.CallOpts, symbol)
}

// GetCurrencyID is a free data retrieval call binding the contract method 0x1c3fcee0.
//
// Solidity: function getCurrencyID(string symbol) view returns(uint64, bool)
func (_Oracle *OracleCallerSession) GetCurrencyID(symbol string) (uint64, bool, error) {
	return _Oracle.Contract.GetCurrencyID(&_Oracle.CallOpts, symbol)
}

// GetCurrencyMetadata is a free data retrieval call binding the contract method 0x184ca4f2.
//
// Solidity: function getCurrencyMetadata(uint64 id) view returns((string,string))
func (_Oracle *OracleCaller) GetCurrencyMetadata(opts *bind.CallOpts, id uint64) (QuasarCurrency, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getCurrencyMetadata", id)

	if err != nil {
		return *new(QuasarCurrency), err
	}

	out0 := *abi.ConvertType(out[0], new(QuasarCurrency)).(*QuasarCurrency)

	return out0, err

}

// GetCurrencyMetadata is a free data retrieval call binding the contract method 0x184ca4f2.
//
// Solidity: function getCurrencyMetadata(uint64 id) view returns((string,string))
func (_Oracle *OracleSession) GetCurrencyMetadata(id uint64) (QuasarCurrency, error) {
	return _Oracle.Contract.GetCurrencyMetadata(&_Oracle.CallOpts, id)
}

// GetCurrencyMetadata is a free data retrieval call binding the contract method 0x184ca4f2.
//
// Solidity: function getCurrencyMetadata(uint64 id) view returns((string,string))
func (_Oracle *OracleCallerSession) GetCurrencyMetadata(id uint64) (QuasarCurrency, error) {
	return _Oracle.Contract.GetCurrencyMetadata(&_Oracle.CallOpts, id)
}

// GetNextID is a free data retrieval call binding the contract method 0x2b0b086e.
//
// Solidity: function getNextID() view returns(uint64)
func (_Oracle *OracleCaller) GetNextID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getNextID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextID is a free data retrieval call binding the contract method 0x2b0b086e.
//
// Solidity: function getNextID() view returns(uint64)
func (_Oracle *OracleSession) GetNextID() (uint64, error) {
	return _Oracle.Contract.GetNextID(&_Oracle.CallOpts)
}

// GetNextID is a free data retrieval call binding the contract method 0x2b0b086e.
//
// Solidity: function getNextID() view returns(uint64)
func (_Oracle *OracleCallerSession) GetNextID() (uint64, error) {
	return _Oracle.Contract.GetNextID(&_Oracle.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x520bcf4b.
//
// Solidity: function getPrice(uint64 id) view returns(uint256)
func (_Oracle *OracleCaller) GetPrice(opts *bind.CallOpts, id uint64) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getPrice", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x520bcf4b.
//
// Solidity: function getPrice(uint64 id) view returns(uint256)
func (_Oracle *OracleSession) GetPrice(id uint64) (*big.Int, error) {
	return _Oracle.Contract.GetPrice(&_Oracle.CallOpts, id)
}

// GetPrice is a free data retrieval call binding the contract method 0x520bcf4b.
//
// Solidity: function getPrice(uint64 id) view returns(uint256)
func (_Oracle *OracleCallerSession) GetPrice(id uint64) (*big.Int, error) {
	return _Oracle.Contract.GetPrice(&_Oracle.CallOpts, id)
}

// GetSupportedCurrencies is a free data retrieval call binding the contract method 0x2ce75863.
//
// Solidity: function getSupportedCurrencies() view returns((string,string)[], bool[])
func (_Oracle *OracleCaller) GetSupportedCurrencies(opts *bind.CallOpts) ([]QuasarCurrency, []bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getSupportedCurrencies")

	if err != nil {
		return *new([]QuasarCurrency), *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]QuasarCurrency)).(*[]QuasarCurrency)
	out1 := *abi.ConvertType(out[1], new([]bool)).(*[]bool)

	return out0, out1, err

}

// GetSupportedCurrencies is a free data retrieval call binding the contract method 0x2ce75863.
//
// Solidity: function getSupportedCurrencies() view returns((string,string)[], bool[])
func (_Oracle *OracleSession) GetSupportedCurrencies() ([]QuasarCurrency, []bool, error) {
	return _Oracle.Contract.GetSupportedCurrencies(&_Oracle.CallOpts)
}

// GetSupportedCurrencies is a free data retrieval call binding the contract method 0x2ce75863.
//
// Solidity: function getSupportedCurrencies() view returns((string,string)[], bool[])
func (_Oracle *OracleCallerSession) GetSupportedCurrencies() ([]QuasarCurrency, []bool, error) {
	return _Oracle.Contract.GetSupportedCurrencies(&_Oracle.CallOpts)
}

// IsCurrencySupported is a free data retrieval call binding the contract method 0xfd30e0a0.
//
// Solidity: function isCurrencySupported(uint64 id) view returns(bool)
func (_Oracle *OracleCaller) IsCurrencySupported(opts *bind.CallOpts, id uint64) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "isCurrencySupported", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrencySupported is a free data retrieval call binding the contract method 0xfd30e0a0.
//
// Solidity: function isCurrencySupported(uint64 id) view returns(bool)
func (_Oracle *OracleSession) IsCurrencySupported(id uint64) (bool, error) {
	return _Oracle.Contract.IsCurrencySupported(&_Oracle.CallOpts, id)
}

// IsCurrencySupported is a free data retrieval call binding the contract method 0xfd30e0a0.
//
// Solidity: function isCurrencySupported(uint64 id) view returns(bool)
func (_Oracle *OracleCallerSession) IsCurrencySupported(id uint64) (bool, error) {
	return _Oracle.Contract.IsCurrencySupported(&_Oracle.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleCallerSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// AddCurrency is a paid mutator transaction binding the contract method 0xce36ad73.
//
// Solidity: function addCurrency(string name, string symbol) returns()
func (_Oracle *OracleTransactor) AddCurrency(opts *bind.TransactOpts, name string, symbol string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "addCurrency", name, symbol)
}

// AddCurrency is a paid mutator transaction binding the contract method 0xce36ad73.
//
// Solidity: function addCurrency(string name, string symbol) returns()
func (_Oracle *OracleSession) AddCurrency(name string, symbol string) (*types.Transaction, error) {
	return _Oracle.Contract.AddCurrency(&_Oracle.TransactOpts, name, symbol)
}

// AddCurrency is a paid mutator transaction binding the contract method 0xce36ad73.
//
// Solidity: function addCurrency(string name, string symbol) returns()
func (_Oracle *OracleTransactorSession) AddCurrency(name string, symbol string) (*types.Transaction, error) {
	return _Oracle.Contract.AddCurrency(&_Oracle.TransactOpts, name, symbol)
}

// ChangeCurrencyState is a paid mutator transaction binding the contract method 0x861d0d35.
//
// Solidity: function changeCurrencyState(uint64 id, bool state) returns()
func (_Oracle *OracleTransactor) ChangeCurrencyState(opts *bind.TransactOpts, id uint64, state bool) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "changeCurrencyState", id, state)
}

// ChangeCurrencyState is a paid mutator transaction binding the contract method 0x861d0d35.
//
// Solidity: function changeCurrencyState(uint64 id, bool state) returns()
func (_Oracle *OracleSession) ChangeCurrencyState(id uint64, state bool) (*types.Transaction, error) {
	return _Oracle.Contract.ChangeCurrencyState(&_Oracle.TransactOpts, id, state)
}

// ChangeCurrencyState is a paid mutator transaction binding the contract method 0x861d0d35.
//
// Solidity: function changeCurrencyState(uint64 id, bool state) returns()
func (_Oracle *OracleTransactorSession) ChangeCurrencyState(id uint64, state bool) (*types.Transaction, error) {
	return _Oracle.Contract.ChangeCurrencyState(&_Oracle.TransactOpts, id, state)
}

// PushPrice is a paid mutator transaction binding the contract method 0x79e6c42a.
//
// Solidity: function pushPrice(uint64 id, uint256 price) returns()
func (_Oracle *OracleTransactor) PushPrice(opts *bind.TransactOpts, id uint64, price *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "pushPrice", id, price)
}

// PushPrice is a paid mutator transaction binding the contract method 0x79e6c42a.
//
// Solidity: function pushPrice(uint64 id, uint256 price) returns()
func (_Oracle *OracleSession) PushPrice(id uint64, price *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.PushPrice(&_Oracle.TransactOpts, id, price)
}

// PushPrice is a paid mutator transaction binding the contract method 0x79e6c42a.
//
// Solidity: function pushPrice(uint64 id, uint256 price) returns()
func (_Oracle *OracleTransactorSession) PushPrice(id uint64, price *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.PushPrice(&_Oracle.TransactOpts, id, price)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oracle *OracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oracle *OracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oracle.Contract.RenounceOwnership(&_Oracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oracle *OracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oracle.Contract.RenounceOwnership(&_Oracle.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oracle *OracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oracle *OracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.TransferOwnership(&_Oracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oracle *OracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.TransferOwnership(&_Oracle.TransactOpts, newOwner)
}

// UpdateCurrency is a paid mutator transaction binding the contract method 0xabc5e6a6.
//
// Solidity: function updateCurrency(uint64 id, string name, string symbol) returns()
func (_Oracle *OracleTransactor) UpdateCurrency(opts *bind.TransactOpts, id uint64, name string, symbol string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "updateCurrency", id, name, symbol)
}

// UpdateCurrency is a paid mutator transaction binding the contract method 0xabc5e6a6.
//
// Solidity: function updateCurrency(uint64 id, string name, string symbol) returns()
func (_Oracle *OracleSession) UpdateCurrency(id uint64, name string, symbol string) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateCurrency(&_Oracle.TransactOpts, id, name, symbol)
}

// UpdateCurrency is a paid mutator transaction binding the contract method 0xabc5e6a6.
//
// Solidity: function updateCurrency(uint64 id, string name, string symbol) returns()
func (_Oracle *OracleTransactorSession) UpdateCurrency(id uint64, name string, symbol string) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateCurrency(&_Oracle.TransactOpts, id, name, symbol)
}

// OracleCurrencyAddedIterator is returned from FilterCurrencyAdded and is used to iterate over the raw logs and unpacked data for CurrencyAdded events raised by the Oracle contract.
type OracleCurrencyAddedIterator struct {
	Event *OracleCurrencyAdded // Event containing the contract specifics and raw log

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
func (it *OracleCurrencyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleCurrencyAdded)
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
		it.Event = new(OracleCurrencyAdded)
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
func (it *OracleCurrencyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleCurrencyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleCurrencyAdded represents a CurrencyAdded event raised by the Oracle contract.
type OracleCurrencyAdded struct {
	Id     uint64
	Name   string
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCurrencyAdded is a free log retrieval operation binding the contract event 0xad343e4211b41810d43dbf7ba860cd361ec4e7ecfeba134b4306a3251a35e3e8.
//
// Solidity: event CurrencyAdded(uint64 indexed id, string name, string symbol)
func (_Oracle *OracleFilterer) FilterCurrencyAdded(opts *bind.FilterOpts, id []uint64) (*OracleCurrencyAddedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "CurrencyAdded", idRule)
	if err != nil {
		return nil, err
	}
	return &OracleCurrencyAddedIterator{contract: _Oracle.contract, event: "CurrencyAdded", logs: logs, sub: sub}, nil
}

// WatchCurrencyAdded is a free log subscription operation binding the contract event 0xad343e4211b41810d43dbf7ba860cd361ec4e7ecfeba134b4306a3251a35e3e8.
//
// Solidity: event CurrencyAdded(uint64 indexed id, string name, string symbol)
func (_Oracle *OracleFilterer) WatchCurrencyAdded(opts *bind.WatchOpts, sink chan<- *OracleCurrencyAdded, id []uint64) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "CurrencyAdded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleCurrencyAdded)
				if err := _Oracle.contract.UnpackLog(event, "CurrencyAdded", log); err != nil {
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

// ParseCurrencyAdded is a log parse operation binding the contract event 0xad343e4211b41810d43dbf7ba860cd361ec4e7ecfeba134b4306a3251a35e3e8.
//
// Solidity: event CurrencyAdded(uint64 indexed id, string name, string symbol)
func (_Oracle *OracleFilterer) ParseCurrencyAdded(log types.Log) (*OracleCurrencyAdded, error) {
	event := new(OracleCurrencyAdded)
	if err := _Oracle.contract.UnpackLog(event, "CurrencyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleCurrencyStateChangedIterator is returned from FilterCurrencyStateChanged and is used to iterate over the raw logs and unpacked data for CurrencyStateChanged events raised by the Oracle contract.
type OracleCurrencyStateChangedIterator struct {
	Event *OracleCurrencyStateChanged // Event containing the contract specifics and raw log

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
func (it *OracleCurrencyStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleCurrencyStateChanged)
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
		it.Event = new(OracleCurrencyStateChanged)
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
func (it *OracleCurrencyStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleCurrencyStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleCurrencyStateChanged represents a CurrencyStateChanged event raised by the Oracle contract.
type OracleCurrencyStateChanged struct {
	Id    uint64
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCurrencyStateChanged is a free log retrieval operation binding the contract event 0xc491c163d3b6cfb2f97ec0acd164a01b48db3856f076d6079a1b7bd345251976.
//
// Solidity: event CurrencyStateChanged(uint64 indexed id, bool state)
func (_Oracle *OracleFilterer) FilterCurrencyStateChanged(opts *bind.FilterOpts, id []uint64) (*OracleCurrencyStateChangedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "CurrencyStateChanged", idRule)
	if err != nil {
		return nil, err
	}
	return &OracleCurrencyStateChangedIterator{contract: _Oracle.contract, event: "CurrencyStateChanged", logs: logs, sub: sub}, nil
}

// WatchCurrencyStateChanged is a free log subscription operation binding the contract event 0xc491c163d3b6cfb2f97ec0acd164a01b48db3856f076d6079a1b7bd345251976.
//
// Solidity: event CurrencyStateChanged(uint64 indexed id, bool state)
func (_Oracle *OracleFilterer) WatchCurrencyStateChanged(opts *bind.WatchOpts, sink chan<- *OracleCurrencyStateChanged, id []uint64) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "CurrencyStateChanged", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleCurrencyStateChanged)
				if err := _Oracle.contract.UnpackLog(event, "CurrencyStateChanged", log); err != nil {
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

// ParseCurrencyStateChanged is a log parse operation binding the contract event 0xc491c163d3b6cfb2f97ec0acd164a01b48db3856f076d6079a1b7bd345251976.
//
// Solidity: event CurrencyStateChanged(uint64 indexed id, bool state)
func (_Oracle *OracleFilterer) ParseCurrencyStateChanged(log types.Log) (*OracleCurrencyStateChanged, error) {
	event := new(OracleCurrencyStateChanged)
	if err := _Oracle.contract.UnpackLog(event, "CurrencyStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleCurrencyUpdatedIterator is returned from FilterCurrencyUpdated and is used to iterate over the raw logs and unpacked data for CurrencyUpdated events raised by the Oracle contract.
type OracleCurrencyUpdatedIterator struct {
	Event *OracleCurrencyUpdated // Event containing the contract specifics and raw log

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
func (it *OracleCurrencyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleCurrencyUpdated)
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
		it.Event = new(OracleCurrencyUpdated)
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
func (it *OracleCurrencyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleCurrencyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleCurrencyUpdated represents a CurrencyUpdated event raised by the Oracle contract.
type OracleCurrencyUpdated struct {
	Id     uint64
	Name   string
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCurrencyUpdated is a free log retrieval operation binding the contract event 0xc6b670b7657fbec139696e713de6494d3f00f0dcd52fba8b368bc306a33a6b68.
//
// Solidity: event CurrencyUpdated(uint64 indexed id, string name, string symbol)
func (_Oracle *OracleFilterer) FilterCurrencyUpdated(opts *bind.FilterOpts, id []uint64) (*OracleCurrencyUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "CurrencyUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &OracleCurrencyUpdatedIterator{contract: _Oracle.contract, event: "CurrencyUpdated", logs: logs, sub: sub}, nil
}

// WatchCurrencyUpdated is a free log subscription operation binding the contract event 0xc6b670b7657fbec139696e713de6494d3f00f0dcd52fba8b368bc306a33a6b68.
//
// Solidity: event CurrencyUpdated(uint64 indexed id, string name, string symbol)
func (_Oracle *OracleFilterer) WatchCurrencyUpdated(opts *bind.WatchOpts, sink chan<- *OracleCurrencyUpdated, id []uint64) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "CurrencyUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleCurrencyUpdated)
				if err := _Oracle.contract.UnpackLog(event, "CurrencyUpdated", log); err != nil {
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

// ParseCurrencyUpdated is a log parse operation binding the contract event 0xc6b670b7657fbec139696e713de6494d3f00f0dcd52fba8b368bc306a33a6b68.
//
// Solidity: event CurrencyUpdated(uint64 indexed id, string name, string symbol)
func (_Oracle *OracleFilterer) ParseCurrencyUpdated(log types.Log) (*OracleCurrencyUpdated, error) {
	event := new(OracleCurrencyUpdated)
	if err := _Oracle.contract.UnpackLog(event, "CurrencyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oracle contract.
type OracleOwnershipTransferredIterator struct {
	Event *OracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleOwnershipTransferred)
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
		it.Event = new(OracleOwnershipTransferred)
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
func (it *OracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleOwnershipTransferred represents a OwnershipTransferred event raised by the Oracle contract.
type OracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oracle *OracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OracleOwnershipTransferredIterator{contract: _Oracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oracle *OracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleOwnershipTransferred)
				if err := _Oracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseOwnershipTransferred(log types.Log) (*OracleOwnershipTransferred, error) {
	event := new(OracleOwnershipTransferred)
	if err := _Oracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OraclePriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the Oracle contract.
type OraclePriceUpdatedIterator struct {
	Event *OraclePriceUpdated // Event containing the contract specifics and raw log

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
func (it *OraclePriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OraclePriceUpdated)
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
		it.Event = new(OraclePriceUpdated)
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
func (it *OraclePriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OraclePriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OraclePriceUpdated represents a PriceUpdated event raised by the Oracle contract.
type OraclePriceUpdated struct {
	Id    uint64
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0x07ff9b8a0601277d4417d6e639b80e205b57cb801379edfda2ab6820e3b7b30e.
//
// Solidity: event PriceUpdated(uint64 indexed id, uint256 price)
func (_Oracle *OracleFilterer) FilterPriceUpdated(opts *bind.FilterOpts, id []uint64) (*OraclePriceUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "PriceUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &OraclePriceUpdatedIterator{contract: _Oracle.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0x07ff9b8a0601277d4417d6e639b80e205b57cb801379edfda2ab6820e3b7b30e.
//
// Solidity: event PriceUpdated(uint64 indexed id, uint256 price)
func (_Oracle *OracleFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *OraclePriceUpdated, id []uint64) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "PriceUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OraclePriceUpdated)
				if err := _Oracle.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
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

// ParsePriceUpdated is a log parse operation binding the contract event 0x07ff9b8a0601277d4417d6e639b80e205b57cb801379edfda2ab6820e3b7b30e.
//
// Solidity: event PriceUpdated(uint64 indexed id, uint256 price)
func (_Oracle *OracleFilterer) ParsePriceUpdated(log types.Log) (*OraclePriceUpdated, error) {
	event := new(OraclePriceUpdated)
	if err := _Oracle.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
