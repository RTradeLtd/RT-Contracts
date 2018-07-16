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

// VestingABI is the input ABI used to generate the binding from.
const VestingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"TOKENADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RTI\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"vests\",\"outputs\":[{\"name\":\"totalVest\",\"type\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vester\",\"type\":\"address\"},{\"name\":\"_totalAmountToVest\",\"type\":\"uint256\"},{\"name\":\"_releaseDates\",\"type\":\"uint256[]\"},{\"name\":\"_releaseAmounts\",\"type\":\"uint256[]\"}],\"name\":\"addVest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vestIndex\",\"type\":\"uint256\"}],\"name\":\"withdrawVestedTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// VestingBin is the compiled bytecode used for deploying new contracts.
const VestingBin = `608060405234801561001057600080fd5b50610948806100206000396000f30060806040526004361061006c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663516f8986811461007157806358e1c174146100715780635c712bc0146100a257806398625080146100ee578063a55a9c41146101a2575b600080fd5b34801561007d57600080fd5b506100866101ba565b60408051600160a060020a039092168252519081900360200190f35b3480156100ae57600080fd5b506100c3600160a060020a03600435166101d2565b604051808381526020018260028111156100d957fe5b60ff1681526020019250505060405180910390f35b3480156100fa57600080fd5b50604080516020600460443581810135838102808601850190965280855261018e958335600160a060020a0316956024803596369695606495939492019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506101ee9650505050505050565b604080519115158252519081900360200190f35b3480156101ae57600080fd5b5061018e600435610589565b73b8fe3b2c83014566733b766a27d94cb9ac167dc681565b6000602081905290815260409020805460039091015460ff1682565b60008060006101fb610881565b876000600160a060020a03821660009081526020819052604090206003015460ff16600281111561022857fe5b1461023257600080fd5b855187511461028b576040805160e560020a62461bcd02815260206004820152601b60248201527f6172726179206c656e6774687320617265206e6f7420657175616c0000000000604482015290519081900360640190fd5b600092505b8551831015610362576102c186848151811015156102aa57fe5b60209081029091010151859063ffffffff61085316565b935086838151811015156102d157fe5b602090810290910101514210610357576040805160e560020a62461bcd02815260206004820152602a60248201527f696e76616c69642072656c656173652064617465206d75737420626520696e2060448201527f7468652066757475726500000000000000000000000000000000000000000000606482015290519081900360840190fd5b600190920191610290565b8388146103b9576040805160e560020a62461bcd02815260206004820152601c60248201527f696e76616c696420746f74616c20616d6f756e7420746f207665737400000000604482015290519081900360640190fd5b60408051608081018252898152602081018990529081018790526060810160019052600160a060020a038a166000908152602081815260409091208251815581830151805193955085939192610417926001850192909101906108b4565b50604082015180516104339160028401916020909101906108b4565b50606082015160038201805460ff1916600183600281111561045157fe5b021790555050604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018b9052905173b8fe3b2c83014566733b766a27d94cb9ac167dc692506323b872dd916064808201926020929091908290030181600087803b1580156104d257600080fd5b505af11580156104e6573d6000803e3d6000fd5b505050506040513d60208110156104fc57600080fd5b5051151561057a576040805160e560020a62461bcd02815260206004820152603060248201527f7472616e736665722066726f6d206661696c65642c206d6f7374206c696b656c60448201527f79206e6565647320617070726f76616c00000000000000000000000000000000606482015290519081900360840190fd5b50600198975050505050505050565b600080808060013360009081526020819052604090206003015460ff1660028111156105b157fe5b146105bb57600080fd5b33600090815260208190526040902060010154859081106105db57600080fd5b33600090815260208181526040808320898452600401909152902054869060ff161561060657600080fd5b33600090815260208190526040902060010180548891908290811061062757fe5b9060005260206000200154421015151561064057600080fd5b336000818152602081815260408083208c8452600481018352908320805460ff19166001908117909155938352919052600201546106839163ffffffff61086c16565b88141561070857600094505b336000908152602081905260409020600201548510156106e4573360009081526020818152604080832088845260040190915290205460ff1615156106d757600095506106e4565b600195509385019361068f565b851561070857336000908152602081905260409020600301805460ff191660021790555b33600090815260208190526040902060020180548990811061072657fe5b9060005260206000200154935073b8fe3b2c83014566733b766a27d94cb9ac167dc6600160a060020a031663a9059cbb33866040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156107c357600080fd5b505af11580156107d7573d6000803e3d6000fd5b505050506040513d60208110156107ed57600080fd5b50511515610845576040805160e560020a62461bcd02815260206004820152601260248201527f6661696c656420746f207472616e736665720000000000000000000000000000604482015290519081900360640190fd5b506001979650505050505050565b60008282018381101561086557600080fd5b9392505050565b60008282111561087b57600080fd5b50900390565b608060405190810160405280600081526020016060815260200160608152602001600060028111156108af57fe5b905290565b8280548282559060005260206000209081019282156108ef579160200282015b828111156108ef5782518255916020019190600101906108d4565b506108fb9291506108ff565b5090565b61091991905b808211156108fb5760008155600101610905565b905600a165627a7a72305820305608eb1b7bfcd6e11bbb4a868f53c7e9cf58f3ff6e3cfd6b76df23124a872b0029`

// DeployVesting deploys a new Ethereum contract, binding an instance of Vesting to it.
func DeployVesting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vesting, error) {
	parsed, err := abi.JSON(strings.NewReader(VestingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VestingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vesting{VestingCaller: VestingCaller{contract: contract}, VestingTransactor: VestingTransactor{contract: contract}, VestingFilterer: VestingFilterer{contract: contract}}, nil
}

// Vesting is an auto generated Go binding around an Ethereum contract.
type Vesting struct {
	VestingCaller     // Read-only binding to the contract
	VestingTransactor // Write-only binding to the contract
	VestingFilterer   // Log filterer for contract events
}

// VestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VestingSession struct {
	Contract     *Vesting          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VestingCallerSession struct {
	Contract *VestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VestingTransactorSession struct {
	Contract     *VestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VestingRaw struct {
	Contract *Vesting // Generic contract binding to access the raw methods on
}

// VestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VestingCallerRaw struct {
	Contract *VestingCaller // Generic read-only contract binding to access the raw methods on
}

// VestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VestingTransactorRaw struct {
	Contract *VestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVesting creates a new instance of Vesting, bound to a specific deployed contract.
func NewVesting(address common.Address, backend bind.ContractBackend) (*Vesting, error) {
	contract, err := bindVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vesting{VestingCaller: VestingCaller{contract: contract}, VestingTransactor: VestingTransactor{contract: contract}, VestingFilterer: VestingFilterer{contract: contract}}, nil
}

// NewVestingCaller creates a new read-only instance of Vesting, bound to a specific deployed contract.
func NewVestingCaller(address common.Address, caller bind.ContractCaller) (*VestingCaller, error) {
	contract, err := bindVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VestingCaller{contract: contract}, nil
}

// NewVestingTransactor creates a new write-only instance of Vesting, bound to a specific deployed contract.
func NewVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*VestingTransactor, error) {
	contract, err := bindVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VestingTransactor{contract: contract}, nil
}

// NewVestingFilterer creates a new log filterer instance of Vesting, bound to a specific deployed contract.
func NewVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*VestingFilterer, error) {
	contract, err := bindVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VestingFilterer{contract: contract}, nil
}

// bindVesting binds a generic wrapper to an already deployed contract.
func bindVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.VestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transact(opts, method, params...)
}

// RTI is a free data retrieval call binding the contract method 0x58e1c174.
//
// Solidity: function RTI() constant returns(address)
func (_Vesting *VestingCaller) RTI(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vesting.contract.Call(opts, out, "RTI")
	return *ret0, err
}

// RTI is a free data retrieval call binding the contract method 0x58e1c174.
//
// Solidity: function RTI() constant returns(address)
func (_Vesting *VestingSession) RTI() (common.Address, error) {
	return _Vesting.Contract.RTI(&_Vesting.CallOpts)
}

// RTI is a free data retrieval call binding the contract method 0x58e1c174.
//
// Solidity: function RTI() constant returns(address)
func (_Vesting *VestingCallerSession) RTI() (common.Address, error) {
	return _Vesting.Contract.RTI(&_Vesting.CallOpts)
}

// TOKENADDRESS is a free data retrieval call binding the contract method 0x516f8986.
//
// Solidity: function TOKENADDRESS() constant returns(address)
func (_Vesting *VestingCaller) TOKENADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vesting.contract.Call(opts, out, "TOKENADDRESS")
	return *ret0, err
}

// TOKENADDRESS is a free data retrieval call binding the contract method 0x516f8986.
//
// Solidity: function TOKENADDRESS() constant returns(address)
func (_Vesting *VestingSession) TOKENADDRESS() (common.Address, error) {
	return _Vesting.Contract.TOKENADDRESS(&_Vesting.CallOpts)
}

// TOKENADDRESS is a free data retrieval call binding the contract method 0x516f8986.
//
// Solidity: function TOKENADDRESS() constant returns(address)
func (_Vesting *VestingCallerSession) TOKENADDRESS() (common.Address, error) {
	return _Vesting.Contract.TOKENADDRESS(&_Vesting.CallOpts)
}

// Vests is a free data retrieval call binding the contract method 0x5c712bc0.
//
// Solidity: function vests( address) constant returns(totalVest uint256, state uint8)
func (_Vesting *VestingCaller) Vests(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalVest *big.Int
	State     uint8
}, error) {
	ret := new(struct {
		TotalVest *big.Int
		State     uint8
	})
	out := ret
	err := _Vesting.contract.Call(opts, out, "vests", arg0)
	return *ret, err
}

// Vests is a free data retrieval call binding the contract method 0x5c712bc0.
//
// Solidity: function vests( address) constant returns(totalVest uint256, state uint8)
func (_Vesting *VestingSession) Vests(arg0 common.Address) (struct {
	TotalVest *big.Int
	State     uint8
}, error) {
	return _Vesting.Contract.Vests(&_Vesting.CallOpts, arg0)
}

// Vests is a free data retrieval call binding the contract method 0x5c712bc0.
//
// Solidity: function vests( address) constant returns(totalVest uint256, state uint8)
func (_Vesting *VestingCallerSession) Vests(arg0 common.Address) (struct {
	TotalVest *big.Int
	State     uint8
}, error) {
	return _Vesting.Contract.Vests(&_Vesting.CallOpts, arg0)
}

// AddVest is a paid mutator transaction binding the contract method 0x98625080.
//
// Solidity: function addVest(_vester address, _totalAmountToVest uint256, _releaseDates uint256[], _releaseAmounts uint256[]) returns(bool)
func (_Vesting *VestingTransactor) AddVest(opts *bind.TransactOpts, _vester common.Address, _totalAmountToVest *big.Int, _releaseDates []*big.Int, _releaseAmounts []*big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "addVest", _vester, _totalAmountToVest, _releaseDates, _releaseAmounts)
}

// AddVest is a paid mutator transaction binding the contract method 0x98625080.
//
// Solidity: function addVest(_vester address, _totalAmountToVest uint256, _releaseDates uint256[], _releaseAmounts uint256[]) returns(bool)
func (_Vesting *VestingSession) AddVest(_vester common.Address, _totalAmountToVest *big.Int, _releaseDates []*big.Int, _releaseAmounts []*big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.AddVest(&_Vesting.TransactOpts, _vester, _totalAmountToVest, _releaseDates, _releaseAmounts)
}

// AddVest is a paid mutator transaction binding the contract method 0x98625080.
//
// Solidity: function addVest(_vester address, _totalAmountToVest uint256, _releaseDates uint256[], _releaseAmounts uint256[]) returns(bool)
func (_Vesting *VestingTransactorSession) AddVest(_vester common.Address, _totalAmountToVest *big.Int, _releaseDates []*big.Int, _releaseAmounts []*big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.AddVest(&_Vesting.TransactOpts, _vester, _totalAmountToVest, _releaseDates, _releaseAmounts)
}

// WithdrawVestedTokens is a paid mutator transaction binding the contract method 0xa55a9c41.
//
// Solidity: function withdrawVestedTokens(_vestIndex uint256) returns(bool)
func (_Vesting *VestingTransactor) WithdrawVestedTokens(opts *bind.TransactOpts, _vestIndex *big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "withdrawVestedTokens", _vestIndex)
}

// WithdrawVestedTokens is a paid mutator transaction binding the contract method 0xa55a9c41.
//
// Solidity: function withdrawVestedTokens(_vestIndex uint256) returns(bool)
func (_Vesting *VestingSession) WithdrawVestedTokens(_vestIndex *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.WithdrawVestedTokens(&_Vesting.TransactOpts, _vestIndex)
}

// WithdrawVestedTokens is a paid mutator transaction binding the contract method 0xa55a9c41.
//
// Solidity: function withdrawVestedTokens(_vestIndex uint256) returns(bool)
func (_Vesting *VestingTransactorSession) WithdrawVestedTokens(_vestIndex *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.WithdrawVestedTokens(&_Vesting.TransactOpts, _vestIndex)
}
