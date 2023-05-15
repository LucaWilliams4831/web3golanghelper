// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IPancakeFactory

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

// PancakeABI is the input ABI used to generate the binding from.
const PancakeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Pancake is an auto generated Go binding around an Ethereum contract.
type Pancake struct {
	PancakeCaller     // Read-only binding to the contract
	PancakeTransactor // Write-only binding to the contract
	PancakeFilterer   // Log filterer for contract events
}

// PancakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeSession struct {
	Contract     *Pancake          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeCallerSession struct {
	Contract *PancakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PancakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeTransactorSession struct {
	Contract     *PancakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PancakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeRaw struct {
	Contract *Pancake // Generic contract binding to access the raw methods on
}

// PancakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeCallerRaw struct {
	Contract *PancakeCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeTransactorRaw struct {
	Contract *PancakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancake creates a new instance of Pancake, bound to a specific deployed contract.
func NewPancake(address common.Address, backend bind.ContractBackend) (*Pancake, error) {
	contract, err := bindPancake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pancake{PancakeCaller: PancakeCaller{contract: contract}, PancakeTransactor: PancakeTransactor{contract: contract}, PancakeFilterer: PancakeFilterer{contract: contract}}, nil
}

// NewPancakeCaller creates a new read-only instance of Pancake, bound to a specific deployed contract.
func NewPancakeCaller(address common.Address, caller bind.ContractCaller) (*PancakeCaller, error) {
	contract, err := bindPancake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeCaller{contract: contract}, nil
}

// NewPancakeTransactor creates a new write-only instance of Pancake, bound to a specific deployed contract.
func NewPancakeTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeTransactor, error) {
	contract, err := bindPancake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeTransactor{contract: contract}, nil
}

// NewPancakeFilterer creates a new log filterer instance of Pancake, bound to a specific deployed contract.
func NewPancakeFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeFilterer, error) {
	contract, err := bindPancake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeFilterer{contract: contract}, nil
}

// bindPancake binds a generic wrapper to an already deployed contract.
func bindPancake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PancakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pancake *PancakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pancake.Contract.PancakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pancake *PancakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancake.Contract.PancakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pancake *PancakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pancake.Contract.PancakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pancake *PancakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pancake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pancake *PancakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pancake *PancakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pancake.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_Pancake *PancakeCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_Pancake *PancakeSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Pancake.Contract.AllPairs(&_Pancake.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_Pancake *PancakeCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Pancake.Contract.AllPairs(&_Pancake.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Pancake *PancakeCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Pancake *PancakeSession) AllPairsLength() (*big.Int, error) {
	return _Pancake.Contract.AllPairsLength(&_Pancake.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Pancake *PancakeCallerSession) AllPairsLength() (*big.Int, error) {
	return _Pancake.Contract.AllPairsLength(&_Pancake.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Pancake *PancakeCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Pancake *PancakeSession) FeeTo() (common.Address, error) {
	return _Pancake.Contract.FeeTo(&_Pancake.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Pancake *PancakeCallerSession) FeeTo() (common.Address, error) {
	return _Pancake.Contract.FeeTo(&_Pancake.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Pancake *PancakeCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Pancake *PancakeSession) FeeToSetter() (common.Address, error) {
	return _Pancake.Contract.FeeToSetter(&_Pancake.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Pancake *PancakeCallerSession) FeeToSetter() (common.Address, error) {
	return _Pancake.Contract.FeeToSetter(&_Pancake.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Pancake *PancakeCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Pancake *PancakeSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Pancake.Contract.GetPair(&_Pancake.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Pancake *PancakeCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Pancake.Contract.GetPair(&_Pancake.CallOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Pancake *PancakeTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Pancake *PancakeSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.CreatePair(&_Pancake.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Pancake *PancakeTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.CreatePair(&_Pancake.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_Pancake *PancakeTransactor) SetFeeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "setFeeTo", arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_Pancake *PancakeSession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.SetFeeTo(&_Pancake.TransactOpts, arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_Pancake *PancakeTransactorSession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.SetFeeTo(&_Pancake.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_Pancake *PancakeTransactor) SetFeeToSetter(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "setFeeToSetter", arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_Pancake *PancakeSession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.SetFeeToSetter(&_Pancake.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_Pancake *PancakeTransactorSession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.SetFeeToSetter(&_Pancake.TransactOpts, arg0)
}

// PancakePairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the Pancake contract.
type PancakePairCreatedIterator struct {
	Event *PancakePairCreated // Event containing the contract specifics and raw log

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
func (it *PancakePairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePairCreated)
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
		it.Event = new(PancakePairCreated)
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
func (it *PancakePairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePairCreated represents a PairCreated event raised by the Pancake contract.
type PancakePairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Pancake *PancakeFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*PancakePairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &PancakePairCreatedIterator{contract: _Pancake.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Pancake *PancakeFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *PancakePairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePairCreated)
				if err := _Pancake.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Pancake *PancakeFilterer) ParsePairCreated(log types.Log) (*PancakePairCreated, error) {
	event := new(PancakePairCreated)
	if err := _Pancake.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
