// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// VestingABI is the input ABI used to generate the binding from.
const VestingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"TOKENADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RTI\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"vests\",\"outputs\":[{\"name\":\"totalVest\",\"type\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vester\",\"type\":\"address\"},{\"name\":\"_totalAmountToVest\",\"type\":\"uint256\"},{\"name\":\"_releaseDates\",\"type\":\"uint256[]\"},{\"name\":\"_releaseAmounts\",\"type\":\"uint256[]\"}],\"name\":\"addVest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vestIndex\",\"type\":\"uint256\"}],\"name\":\"withdrawVestedTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// VestingBin is the compiled bytecode used for deploying new contracts.
const VestingBin = `608060405234801561001057600080fd5b50604051602080610e42833981016040525160008054600160a060020a03909216600160a060020a0319909216919091179055610df0806100526000396000f3006080604052600436106100825763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663516f8986811461008757806358e1c174146100875780635c712bc0146100b85780639862508014610104578063a55a9c41146101b8578063f851a440146101d0578063ffa1ad74146101e5575b600080fd5b34801561009357600080fd5b5061009c61026f565b60408051600160a060020a039092168252519081900360200190f35b3480156100c457600080fd5b506100d9600160a060020a0360043516610287565b604051808381526020018260028111156100ef57fe5b60ff1681526020019250505060405180910390f35b34801561011057600080fd5b5060408051602060046044358181013583810280860185019096528085526101a4958335600160a060020a0316956024803596369695606495939492019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506102a39650505050505050565b604080519115158252519081900360200190f35b3480156101c457600080fd5b506101a46004356107a9565b3480156101dc57600080fd5b5061009c610cb5565b3480156101f157600080fd5b506101fa610cc4565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561023457818101518382015260200161021c565b50505050905090810190601f1680156102615780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b73ecc043b92834c1ebde65f2181b59597a6588d61681565b6001602052600090815260409020805460039091015460ff1682565b60008060006102b0610d29565b876000600160a060020a03821660009081526001602052604090206003015460ff1660028111156102dd57fe5b14610357576040805160e560020a62461bcd028152602060048201526024808201527f61646472657373206d757374206e6f74206861766520616e206163746976652060448201527f7665737400000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600054600160a060020a031633146103b9576040805160e560020a62461bcd02815260206004820152601460248201527f73656e646572206d7573742062652061646d696e000000000000000000000000604482015290519081900360640190fd5b600087511180156103cb575060008651115b80156103d75750600088115b1515610453576040805160e560020a62461bcd02815260206004820152602160248201527f617474656d7074696e6720746f20757365206e6f6e207a65726f2076616c756560448201527f7300000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b85518751146104ac576040805160e560020a62461bcd02815260206004820152601b60248201527f6172726179206c656e6774687320617265206e6f7420657175616c0000000000604482015290519081900360640190fd5b600092505b8551831015610583576104e286848151811015156104cb57fe5b60209081029091010151859063ffffffff610cfb16565b935086838151811015156104f257fe5b602090810290910101514210610578576040805160e560020a62461bcd02815260206004820152602260248201527f72656c656173652064617465206d75737420626520696e20746865206675747560448201527f7265000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6001909201916104b1565b8388146105da576040805160e560020a62461bcd02815260206004820152601c60248201527f696e76616c696420746f74616c20616d6f756e7420746f207665737400000000604482015290519081900360640190fd5b60408051608081018252898152602081018990529081018790526060810160019052600160a060020a038a166000908152600160208181526040909220835181558284015180519496508694919361063793850192910190610d5c565b5060408201518051610653916002840191602090910190610d5c565b50606082015160038201805460ff1916600183600281111561067157fe5b021790555050604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018b9052905173ecc043b92834c1ebde65f2181b59597a6588d61692506323b872dd916064808201926020929091908290030181600087803b1580156106f257600080fd5b505af1158015610706573d6000803e3d6000fd5b505050506040513d602081101561071c57600080fd5b5051151561079a576040805160e560020a62461bcd02815260206004820152603060248201527f7472616e736665722066726f6d206661696c65642c206d6f7374206c696b656c60448201527f79206e6565647320617070726f76616c00000000000000000000000000000000606482015290519081900360840190fd5b50600198975050505050505050565b600080808060013360009081526001602052604090206003015460ff1660028111156107d157fe5b14610826576040805160e560020a62461bcd02815260206004820152601360248201527f76657374206d7573742062652061637469766500000000000000000000000000604482015290519081900360640190fd5b3360009081526001602081905260409091200154859081106108de576040805160e560020a62461bcd02815260206004820152604960248201527f617474656d7074696e6720746f2061636365737320696e76616c69642076657360448201527f7420696e646578206d757374206265206c657373207468616e206c656e67746860648201527f206f662061727261790000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b336000908152600160209081526040808320898452600401909152902054869060ff1615610956576040805160e560020a62461bcd02815260206004820152601660248201527f76657374206d75737420626520756e636c61696d656400000000000000000000604482015290519081900360640190fd5b3360009081526001602081905260409091200180548891908290811061097857fe5b90600052602060002001544210151515610a02576040805160e560020a62461bcd02815260206004820152602c60248201527f617474656d7074696e6720746f20636c61696d2076657374206265666f72652060448201527f72656c6561736520646174650000000000000000000000000000000000000000606482015290519081900360840190fd5b33600090815260016020819052604090912060020154610a279163ffffffff610d1416565b881415610b4b57600094505b33600090815260016020526040902060020154851015610a8a5733600090815260016020908152604080832088845260040190915290205460ff161515610a7d5760009550610a8a565b6001955093850193610a33565b851515610b2d576040805160e560020a62461bcd02815260206004820152604a60248201527f6e6f7420616c6c2076657374732068617665206265656e20776974686472617760448201527f6e206265666f726520617474656d7074696e6720746f2077697468647261772060648201527f66696e616c207665737400000000000000000000000000000000000000000000608482015290519081900360a40190fd5b336000908152600160205260409020600301805460ff191660021790555b3360008181526001602081815260408084208d8552600481018352908420805460ff191684179055939092529052600201805489908110610b8857fe5b9060005260206000200154935073ecc043b92834c1ebde65f2181b59597a6588d616600160a060020a031663a9059cbb33866040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015610c2557600080fd5b505af1158015610c39573d6000803e3d6000fd5b505050506040513d6020811015610c4f57600080fd5b50511515610ca7576040805160e560020a62461bcd02815260206004820152601260248201527f6661696c656420746f207472616e736665720000000000000000000000000000604482015290519081900360640190fd5b506001979650505050505050565b600054600160a060020a031681565b60408051808201909152600a81527f70726f64756374696f6e00000000000000000000000000000000000000000000602082015281565b600082820183811015610d0d57600080fd5b9392505050565b600082821115610d2357600080fd5b50900390565b60806040519081016040528060008152602001606081526020016060815260200160006002811115610d5757fe5b905290565b828054828255906000526020600020908101928215610d97579160200282015b82811115610d97578251825591602001919060010190610d7c565b50610da3929150610da7565b5090565b610dc191905b80821115610da35760008155600101610dad565b905600a165627a7a723058205a82bf4ccdf569b0fa99a1a64f0f308dbac760b61acbef0767f8383e0490eb830029`

// DeployVesting deploys a new Ethereum contract, binding an instance of Vesting to it.
func DeployVesting(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address) (common.Address, *types.Transaction, *Vesting, error) {
	parsed, err := abi.JSON(strings.NewReader(VestingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VestingBin), backend, _admin)
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

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Vesting *VestingCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Vesting.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Vesting *VestingSession) VERSION() (string, error) {
	return _Vesting.Contract.VERSION(&_Vesting.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Vesting *VestingCallerSession) VERSION() (string, error) {
	return _Vesting.Contract.VERSION(&_Vesting.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Vesting *VestingCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vesting.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Vesting *VestingSession) Admin() (common.Address, error) {
	return _Vesting.Contract.Admin(&_Vesting.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Vesting *VestingCallerSession) Admin() (common.Address, error) {
	return _Vesting.Contract.Admin(&_Vesting.CallOpts)
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
