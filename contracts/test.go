// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// TestABI is the input ABI used to generate the binding from.
const TestABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMyVariable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"myPublicVariable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"}],\"name\":\"ValueSet\",\"type\":\"event\"}]"

// TestFuncSigs maps the 4-byte function signature to its string representation.
var TestFuncSigs = map[string]string{
	"c2cb7544": "myPublicVariable()",
	"6ec9c7ae": "setMyVariable(uint256)",
}

// TestBin is the compiled bytecode used for deploying new contracts.
var TestBin = "0x608060405234801561001057600080fd5b5060e88061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80636ec9c7ae146037578063c2cb7544146053575b600080fd5b605160048036036020811015604b57600080fd5b5035606b565b005b605960ad565b60408051918252519081900360200190f35b60008190556040805182815242602082015281517f69be06033bef8d755e18606a27d6d07393aabbd1800776e503af2c8a03b7c681929181900390910190a150565b6000548156fea265627a7a7231582039d81082c98748de0fffe3a79e92315db86ed0e6a5e42e6411f0bc610da7969964736f6c634300050b0032"

// DeployTest deploys a new Ethereum contract, binding an instance of Test to it.
func DeployTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Test, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// Test is an auto generated Go binding around an Ethereum contract.
type Test struct {
	TestCaller     // Read-only binding to the contract
	TestTransactor // Write-only binding to the contract
	TestFilterer   // Log filterer for contract events
}

// TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSession struct {
	Contract     *Test             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCallerSession struct {
	Contract *TestCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTransactorSession struct {
	Contract     *TestTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRaw struct {
	Contract *Test // Generic contract binding to access the raw methods on
}

// TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCallerRaw struct {
	Contract *TestCaller // Generic read-only contract binding to access the raw methods on
}

// TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTransactorRaw struct {
	Contract *TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest creates a new instance of Test, bound to a specific deployed contract.
func NewTest(address common.Address, backend bind.ContractBackend) (*Test, error) {
	contract, err := bindTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// NewTestCaller creates a new read-only instance of Test, bound to a specific deployed contract.
func NewTestCaller(address common.Address, caller bind.ContractCaller) (*TestCaller, error) {
	contract, err := bindTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCaller{contract: contract}, nil
}

// NewTestTransactor creates a new write-only instance of Test, bound to a specific deployed contract.
func NewTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTransactor, error) {
	contract, err := bindTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTransactor{contract: contract}, nil
}

// NewTestFilterer creates a new log filterer instance of Test, bound to a specific deployed contract.
func NewTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TestFilterer, error) {
	contract, err := bindTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestFilterer{contract: contract}, nil
}

// bindTest binds a generic wrapper to an already deployed contract.
func bindTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Test.Contract.TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.contract.Transact(opts, method, params...)
}

// MyPublicVariable is a free data retrieval call binding the contract method 0xc2cb7544.
//
// Solidity: function myPublicVariable() constant returns(uint256)
func (_Test *TestCaller) MyPublicVariable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Test.contract.Call(opts, out, "myPublicVariable")
	return *ret0, err
}

// MyPublicVariable is a free data retrieval call binding the contract method 0xc2cb7544.
//
// Solidity: function myPublicVariable() constant returns(uint256)
func (_Test *TestSession) MyPublicVariable() (*big.Int, error) {
	return _Test.Contract.MyPublicVariable(&_Test.CallOpts)
}

// MyPublicVariable is a free data retrieval call binding the contract method 0xc2cb7544.
//
// Solidity: function myPublicVariable() constant returns(uint256)
func (_Test *TestCallerSession) MyPublicVariable() (*big.Int, error) {
	return _Test.Contract.MyPublicVariable(&_Test.CallOpts)
}

// SetMyVariable is a paid mutator transaction binding the contract method 0x6ec9c7ae.
//
// Solidity: function setMyVariable(uint256 newValue) returns()
func (_Test *TestTransactor) SetMyVariable(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "setMyVariable", newValue)
}

// SetMyVariable is a paid mutator transaction binding the contract method 0x6ec9c7ae.
//
// Solidity: function setMyVariable(uint256 newValue) returns()
func (_Test *TestSession) SetMyVariable(newValue *big.Int) (*types.Transaction, error) {
	return _Test.Contract.SetMyVariable(&_Test.TransactOpts, newValue)
}

// SetMyVariable is a paid mutator transaction binding the contract method 0x6ec9c7ae.
//
// Solidity: function setMyVariable(uint256 newValue) returns()
func (_Test *TestTransactorSession) SetMyVariable(newValue *big.Int) (*types.Transaction, error) {
	return _Test.Contract.SetMyVariable(&_Test.TransactOpts, newValue)
}

// TestValueSetIterator is returned from FilterValueSet and is used to iterate over the raw logs and unpacked data for ValueSet events raised by the Test contract.
type TestValueSetIterator struct {
	Event *TestValueSet // Event containing the contract specifics and raw log

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
func (it *TestValueSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestValueSet)
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
		it.Event = new(TestValueSet)
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
func (it *TestValueSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestValueSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestValueSet represents a ValueSet event raised by the Test contract.
type TestValueSet struct {
	NewValue *big.Int
	Date     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValueSet is a free log retrieval operation binding the contract event 0x69be06033bef8d755e18606a27d6d07393aabbd1800776e503af2c8a03b7c681.
//
// Solidity: event ValueSet(uint256 newValue, uint256 date)
func (_Test *TestFilterer) FilterValueSet(opts *bind.FilterOpts) (*TestValueSetIterator, error) {

	logs, sub, err := _Test.contract.FilterLogs(opts, "ValueSet")
	if err != nil {
		return nil, err
	}
	return &TestValueSetIterator{contract: _Test.contract, event: "ValueSet", logs: logs, sub: sub}, nil
}

// WatchValueSet is a free log subscription operation binding the contract event 0x69be06033bef8d755e18606a27d6d07393aabbd1800776e503af2c8a03b7c681.
//
// Solidity: event ValueSet(uint256 newValue, uint256 date)
func (_Test *TestFilterer) WatchValueSet(opts *bind.WatchOpts, sink chan<- *TestValueSet) (event.Subscription, error) {

	logs, sub, err := _Test.contract.WatchLogs(opts, "ValueSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestValueSet)
				if err := _Test.contract.UnpackLog(event, "ValueSet", log); err != nil {
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

// ParseValueSet is a log parse operation binding the contract event 0x69be06033bef8d755e18606a27d6d07393aabbd1800776e503af2c8a03b7c681.
//
// Solidity: event ValueSet(uint256 newValue, uint256 date)
func (_Test *TestFilterer) ParseValueSet(log types.Log) (*TestValueSet, error) {
	event := new(TestValueSet)
	if err := _Test.contract.UnpackLog(event, "ValueSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
