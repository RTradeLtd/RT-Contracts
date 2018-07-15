// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ValidatorABI is the input ABI used to generate the binding from.
const ValidatorABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"submitBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SUBMISSIONREWARD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKENADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RTI\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINWITHDRAWAL\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BLOCKREWARD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"name\":\"number\",\"type\":\"uint256\"},{\"name\":\"coinbase\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"set\",\"type\":\"bool\"},{\"name\":\"claimed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ValidatorBin is the compiled bytecode used for deploying new contracts.
const ValidatorBin = `608060405234801561001057600080fd5b50610619806100206000396000f30060806040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166325ceb4b281146100925780633218ebcc146100bb578063516f8986146100e257806358e1c174146100e2578063942a050c146100bb578063ae169a5014610120578063f2477f7c146100bb578063f25b3f9914610138575b600080fd5b34801561009e57600080fd5b506100a7610194565b604080519115158252519081900360200190f35b3480156100c757600080fd5b506100d06103a8565b60408051918252519081900360200190f35b3480156100ee57600080fd5b506100f76103ad565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561012c57600080fd5b506100a76004356103c5565b34801561014457600080fd5b50610150600435610575565b6040805195865273ffffffffffffffffffffffffffffffffffffffff9094166020860152848401929092521515606084015215156080830152519081900360a00190f35b600061019e6105bf565b4360008181526020819052604090206003015460ff16156101be57600080fd5b6040805160a081018252438082524160208084019182528240848601908152600160608601818152600060808801818152968152808552888120885181559551868401805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff909216919091179055925160028601555160039094018054955160ff199096169415159490941761ff001916610100951515959095029490941790925584517fa9059cbb0000000000000000000000000000000000000000000000000000000081523360048201526024810193909352935192955073185ae6a87bbb02097923e859d742747bb979ae9a9363a9059cbb93604480850194929391928390030190829087803b1580156102e057600080fd5b505af11580156102f4573d6000803e3d6000fd5b505050506040513d602081101561030a57600080fd5b5051151561039f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f6661696c656420746f207472616e736665722072657761726420746f20626c6f60448201527f636b207375626d69747465720000000000000000000000000000000000000000606482015290519081900360840190fd5b60019250505090565b600181565b73185ae6a87bbb02097923e859d742747bb979ae9a81565b600081815260208190526040812060010154829073ffffffffffffffffffffffffffffffffffffffff1633146103fa57600080fd5b6000838152602081905260409020600301548390610100900460ff161561042057600080fd5b600084815260208190526040902060030154849060ff16151561044257600080fd5b600085815260208181526040808320600301805461ff00191661010017905580517fa9059cbb00000000000000000000000000000000000000000000000000000000815233600482015260016024820152905173185ae6a87bbb02097923e859d742747bb979ae9a9363a9059cbb93604480850194919392918390030190829087803b1580156104d157600080fd5b505af11580156104e5573d6000803e3d6000fd5b505050506040513d60208110156104fb57600080fd5b5051151561056a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f6661696c656420746f207472616e7366657220626c6f636b2072657761726400604482015290519081900360640190fd5b506001949350505050565b6000602081905290815260409020805460018201546002830154600390930154919273ffffffffffffffffffffffffffffffffffffffff9091169160ff8082169161010090041685565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152905600a165627a7a7230582078f1597517bdcde6745120e43a179e347de0a4697c65a40090dc80e3c21915100029`

// DeployValidator deploys a new Ethereum contract, binding an instance of Validator to it.
func DeployValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validator, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// Validator is an auto generated Go binding around an Ethereum contract.
type Validator struct {
	ValidatorCaller     // Read-only binding to the contract
	ValidatorTransactor // Write-only binding to the contract
	ValidatorFilterer   // Log filterer for contract events
}

// ValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSession struct {
	Contract     *Validator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorCallerSession struct {
	Contract *ValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorTransactorSession struct {
	Contract     *ValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRaw struct {
	Contract *Validator // Generic contract binding to access the raw methods on
}

// ValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorCallerRaw struct {
	Contract *ValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorTransactorRaw struct {
	Contract *ValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidator creates a new instance of Validator, bound to a specific deployed contract.
func NewValidator(address common.Address, backend bind.ContractBackend) (*Validator, error) {
	contract, err := bindValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// NewValidatorCaller creates a new read-only instance of Validator, bound to a specific deployed contract.
func NewValidatorCaller(address common.Address, caller bind.ContractCaller) (*ValidatorCaller, error) {
	contract, err := bindValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorCaller{contract: contract}, nil
}

// NewValidatorTransactor creates a new write-only instance of Validator, bound to a specific deployed contract.
func NewValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorTransactor, error) {
	contract, err := bindValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorTransactor{contract: contract}, nil
}

// NewValidatorFilterer creates a new log filterer instance of Validator, bound to a specific deployed contract.
func NewValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorFilterer, error) {
	contract, err := bindValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorFilterer{contract: contract}, nil
}

// bindValidator binds a generic wrapper to an already deployed contract.
func bindValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.ValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transact(opts, method, params...)
}

// BLOCKREWARD is a free data retrieval call binding the contract method 0xf2477f7c.
//
// Solidity: function BLOCKREWARD() constant returns(uint256)
func (_Validator *ValidatorCaller) BLOCKREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "BLOCKREWARD")
	return *ret0, err
}

// BLOCKREWARD is a free data retrieval call binding the contract method 0xf2477f7c.
//
// Solidity: function BLOCKREWARD() constant returns(uint256)
func (_Validator *ValidatorSession) BLOCKREWARD() (*big.Int, error) {
	return _Validator.Contract.BLOCKREWARD(&_Validator.CallOpts)
}

// BLOCKREWARD is a free data retrieval call binding the contract method 0xf2477f7c.
//
// Solidity: function BLOCKREWARD() constant returns(uint256)
func (_Validator *ValidatorCallerSession) BLOCKREWARD() (*big.Int, error) {
	return _Validator.Contract.BLOCKREWARD(&_Validator.CallOpts)
}

// MINWITHDRAWAL is a free data retrieval call binding the contract method 0x942a050c.
//
// Solidity: function MINWITHDRAWAL() constant returns(uint256)
func (_Validator *ValidatorCaller) MINWITHDRAWAL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "MINWITHDRAWAL")
	return *ret0, err
}

// MINWITHDRAWAL is a free data retrieval call binding the contract method 0x942a050c.
//
// Solidity: function MINWITHDRAWAL() constant returns(uint256)
func (_Validator *ValidatorSession) MINWITHDRAWAL() (*big.Int, error) {
	return _Validator.Contract.MINWITHDRAWAL(&_Validator.CallOpts)
}

// MINWITHDRAWAL is a free data retrieval call binding the contract method 0x942a050c.
//
// Solidity: function MINWITHDRAWAL() constant returns(uint256)
func (_Validator *ValidatorCallerSession) MINWITHDRAWAL() (*big.Int, error) {
	return _Validator.Contract.MINWITHDRAWAL(&_Validator.CallOpts)
}

// RTI is a free data retrieval call binding the contract method 0x58e1c174.
//
// Solidity: function RTI() constant returns(address)
func (_Validator *ValidatorCaller) RTI(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "RTI")
	return *ret0, err
}

// RTI is a free data retrieval call binding the contract method 0x58e1c174.
//
// Solidity: function RTI() constant returns(address)
func (_Validator *ValidatorSession) RTI() (common.Address, error) {
	return _Validator.Contract.RTI(&_Validator.CallOpts)
}

// RTI is a free data retrieval call binding the contract method 0x58e1c174.
//
// Solidity: function RTI() constant returns(address)
func (_Validator *ValidatorCallerSession) RTI() (common.Address, error) {
	return _Validator.Contract.RTI(&_Validator.CallOpts)
}

// SUBMISSIONREWARD is a free data retrieval call binding the contract method 0x3218ebcc.
//
// Solidity: function SUBMISSIONREWARD() constant returns(uint256)
func (_Validator *ValidatorCaller) SUBMISSIONREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "SUBMISSIONREWARD")
	return *ret0, err
}

// SUBMISSIONREWARD is a free data retrieval call binding the contract method 0x3218ebcc.
//
// Solidity: function SUBMISSIONREWARD() constant returns(uint256)
func (_Validator *ValidatorSession) SUBMISSIONREWARD() (*big.Int, error) {
	return _Validator.Contract.SUBMISSIONREWARD(&_Validator.CallOpts)
}

// SUBMISSIONREWARD is a free data retrieval call binding the contract method 0x3218ebcc.
//
// Solidity: function SUBMISSIONREWARD() constant returns(uint256)
func (_Validator *ValidatorCallerSession) SUBMISSIONREWARD() (*big.Int, error) {
	return _Validator.Contract.SUBMISSIONREWARD(&_Validator.CallOpts)
}

// TOKENADDRESS is a free data retrieval call binding the contract method 0x516f8986.
//
// Solidity: function TOKENADDRESS() constant returns(address)
func (_Validator *ValidatorCaller) TOKENADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "TOKENADDRESS")
	return *ret0, err
}

// TOKENADDRESS is a free data retrieval call binding the contract method 0x516f8986.
//
// Solidity: function TOKENADDRESS() constant returns(address)
func (_Validator *ValidatorSession) TOKENADDRESS() (common.Address, error) {
	return _Validator.Contract.TOKENADDRESS(&_Validator.CallOpts)
}

// TOKENADDRESS is a free data retrieval call binding the contract method 0x516f8986.
//
// Solidity: function TOKENADDRESS() constant returns(address)
func (_Validator *ValidatorCallerSession) TOKENADDRESS() (common.Address, error) {
	return _Validator.Contract.TOKENADDRESS(&_Validator.CallOpts)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks( uint256) constant returns(number uint256, coinbase address, hash bytes32, set bool, claimed bool)
func (_Validator *ValidatorCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Number   *big.Int
	Coinbase common.Address
	Hash     [32]byte
	Set      bool
	Claimed  bool
}, error) {
	ret := new(struct {
		Number   *big.Int
		Coinbase common.Address
		Hash     [32]byte
		Set      bool
		Claimed  bool
	})
	out := ret
	err := _Validator.contract.Call(opts, out, "blocks", arg0)
	return *ret, err
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks( uint256) constant returns(number uint256, coinbase address, hash bytes32, set bool, claimed bool)
func (_Validator *ValidatorSession) Blocks(arg0 *big.Int) (struct {
	Number   *big.Int
	Coinbase common.Address
	Hash     [32]byte
	Set      bool
	Claimed  bool
}, error) {
	return _Validator.Contract.Blocks(&_Validator.CallOpts, arg0)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks( uint256) constant returns(number uint256, coinbase address, hash bytes32, set bool, claimed bool)
func (_Validator *ValidatorCallerSession) Blocks(arg0 *big.Int) (struct {
	Number   *big.Int
	Coinbase common.Address
	Hash     [32]byte
	Set      bool
	Claimed  bool
}, error) {
	return _Validator.Contract.Blocks(&_Validator.CallOpts, arg0)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(_blockNumber uint256) returns(bool)
func (_Validator *ValidatorTransactor) ClaimReward(opts *bind.TransactOpts, _blockNumber *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "claimReward", _blockNumber)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(_blockNumber uint256) returns(bool)
func (_Validator *ValidatorSession) ClaimReward(_blockNumber *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ClaimReward(&_Validator.TransactOpts, _blockNumber)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(_blockNumber uint256) returns(bool)
func (_Validator *ValidatorTransactorSession) ClaimReward(_blockNumber *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ClaimReward(&_Validator.TransactOpts, _blockNumber)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0x25ceb4b2.
//
// Solidity: function submitBlock() returns(bool)
func (_Validator *ValidatorTransactor) SubmitBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "submitBlock")
}

// SubmitBlock is a paid mutator transaction binding the contract method 0x25ceb4b2.
//
// Solidity: function submitBlock() returns(bool)
func (_Validator *ValidatorSession) SubmitBlock() (*types.Transaction, error) {
	return _Validator.Contract.SubmitBlock(&_Validator.TransactOpts)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0x25ceb4b2.
//
// Solidity: function submitBlock() returns(bool)
func (_Validator *ValidatorTransactorSession) SubmitBlock() (*types.Transaction, error) {
	return _Validator.Contract.SubmitBlock(&_Validator.TransactOpts)
}
