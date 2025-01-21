// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// CarRentingCar is an auto generated low-level Go binding around an user-defined struct.
type CarRentingCar struct {
	Hp        uint16
	FuelLevel uint16
	IsRented  bool
	RentedTo  common.Address
}

// CarRentingMetaData contains all meta data concerning the CarRenting contract.
var CarRentingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addCar\",\"inputs\":[{\"name\":\"car\",\"type\":\"tuple\",\"internalType\":\"structCarRenting.Car\",\"components\":[{\"name\":\"hp\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fuelLevel\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isRented\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"rentedTo\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rentCar\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rents\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"hp\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fuelLevel\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isRented\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"rentedTo\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
}

// CarRentingABI is the input ABI used to generate the binding from.
// Deprecated: Use CarRentingMetaData.ABI instead.
var CarRentingABI = CarRentingMetaData.ABI

// CarRenting is an auto generated Go binding around an Ethereum contract.
type CarRenting struct {
	CarRentingCaller     // Read-only binding to the contract
	CarRentingTransactor // Write-only binding to the contract
	CarRentingFilterer   // Log filterer for contract events
}

// CarRentingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CarRentingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarRentingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CarRentingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarRentingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CarRentingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarRentingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CarRentingSession struct {
	Contract     *CarRenting       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CarRentingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CarRentingCallerSession struct {
	Contract *CarRentingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CarRentingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CarRentingTransactorSession struct {
	Contract     *CarRentingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CarRentingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CarRentingRaw struct {
	Contract *CarRenting // Generic contract binding to access the raw methods on
}

// CarRentingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CarRentingCallerRaw struct {
	Contract *CarRentingCaller // Generic read-only contract binding to access the raw methods on
}

// CarRentingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CarRentingTransactorRaw struct {
	Contract *CarRentingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCarRenting creates a new instance of CarRenting, bound to a specific deployed contract.
func NewCarRenting(address common.Address, backend bind.ContractBackend) (*CarRenting, error) {
	contract, err := bindCarRenting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CarRenting{CarRentingCaller: CarRentingCaller{contract: contract}, CarRentingTransactor: CarRentingTransactor{contract: contract}, CarRentingFilterer: CarRentingFilterer{contract: contract}}, nil
}

// NewCarRentingCaller creates a new read-only instance of CarRenting, bound to a specific deployed contract.
func NewCarRentingCaller(address common.Address, caller bind.ContractCaller) (*CarRentingCaller, error) {
	contract, err := bindCarRenting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CarRentingCaller{contract: contract}, nil
}

// NewCarRentingTransactor creates a new write-only instance of CarRenting, bound to a specific deployed contract.
func NewCarRentingTransactor(address common.Address, transactor bind.ContractTransactor) (*CarRentingTransactor, error) {
	contract, err := bindCarRenting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CarRentingTransactor{contract: contract}, nil
}

// NewCarRentingFilterer creates a new log filterer instance of CarRenting, bound to a specific deployed contract.
func NewCarRentingFilterer(address common.Address, filterer bind.ContractFilterer) (*CarRentingFilterer, error) {
	contract, err := bindCarRenting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CarRentingFilterer{contract: contract}, nil
}

// bindCarRenting binds a generic wrapper to an already deployed contract.
func bindCarRenting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CarRentingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CarRenting *CarRentingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CarRenting.Contract.CarRentingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CarRenting *CarRentingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CarRenting.Contract.CarRentingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CarRenting *CarRentingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CarRenting.Contract.CarRentingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CarRenting *CarRentingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CarRenting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CarRenting *CarRentingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CarRenting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CarRenting *CarRentingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CarRenting.Contract.contract.Transact(opts, method, params...)
}

// Rents is a free data retrieval call binding the contract method 0x9fdc8f05.
//
// Solidity: function rents(address ) view returns(uint16 hp, uint16 fuelLevel, bool isRented, address rentedTo)
func (_CarRenting *CarRentingCaller) Rents(opts *bind.CallOpts, arg0 common.Address) (struct {
	Hp        uint16
	FuelLevel uint16
	IsRented  bool
	RentedTo  common.Address
}, error) {
	var out []interface{}
	err := _CarRenting.contract.Call(opts, &out, "rents", arg0)

	outstruct := new(struct {
		Hp        uint16
		FuelLevel uint16
		IsRented  bool
		RentedTo  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hp = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.FuelLevel = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.IsRented = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.RentedTo = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Rents is a free data retrieval call binding the contract method 0x9fdc8f05.
//
// Solidity: function rents(address ) view returns(uint16 hp, uint16 fuelLevel, bool isRented, address rentedTo)
func (_CarRenting *CarRentingSession) Rents(arg0 common.Address) (struct {
	Hp        uint16
	FuelLevel uint16
	IsRented  bool
	RentedTo  common.Address
}, error) {
	return _CarRenting.Contract.Rents(&_CarRenting.CallOpts, arg0)
}

// Rents is a free data retrieval call binding the contract method 0x9fdc8f05.
//
// Solidity: function rents(address ) view returns(uint16 hp, uint16 fuelLevel, bool isRented, address rentedTo)
func (_CarRenting *CarRentingCallerSession) Rents(arg0 common.Address) (struct {
	Hp        uint16
	FuelLevel uint16
	IsRented  bool
	RentedTo  common.Address
}, error) {
	return _CarRenting.Contract.Rents(&_CarRenting.CallOpts, arg0)
}

// AddCar is a paid mutator transaction binding the contract method 0x4f1d221f.
//
// Solidity: function addCar((uint16,uint16,bool,address) car) returns()
func (_CarRenting *CarRentingTransactor) AddCar(opts *bind.TransactOpts, car CarRentingCar) (*types.Transaction, error) {
	return _CarRenting.contract.Transact(opts, "addCar", car)
}

// AddCar is a paid mutator transaction binding the contract method 0x4f1d221f.
//
// Solidity: function addCar((uint16,uint16,bool,address) car) returns()
func (_CarRenting *CarRentingSession) AddCar(car CarRentingCar) (*types.Transaction, error) {
	return _CarRenting.Contract.AddCar(&_CarRenting.TransactOpts, car)
}

// AddCar is a paid mutator transaction binding the contract method 0x4f1d221f.
//
// Solidity: function addCar((uint16,uint16,bool,address) car) returns()
func (_CarRenting *CarRentingTransactorSession) AddCar(car CarRentingCar) (*types.Transaction, error) {
	return _CarRenting.Contract.AddCar(&_CarRenting.TransactOpts, car)
}

// RentCar is a paid mutator transaction binding the contract method 0xc7067645.
//
// Solidity: function rentCar(address owner) returns()
func (_CarRenting *CarRentingTransactor) RentCar(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _CarRenting.contract.Transact(opts, "rentCar", owner)
}

// RentCar is a paid mutator transaction binding the contract method 0xc7067645.
//
// Solidity: function rentCar(address owner) returns()
func (_CarRenting *CarRentingSession) RentCar(owner common.Address) (*types.Transaction, error) {
	return _CarRenting.Contract.RentCar(&_CarRenting.TransactOpts, owner)
}

// RentCar is a paid mutator transaction binding the contract method 0xc7067645.
//
// Solidity: function rentCar(address owner) returns()
func (_CarRenting *CarRentingTransactorSession) RentCar(owner common.Address) (*types.Transaction, error) {
	return _CarRenting.Contract.RentCar(&_CarRenting.TransactOpts, owner)
}
