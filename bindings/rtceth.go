// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// RTCETHABI is the input ABI used to generate the binding from.
const RTCETHABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unlockSales\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiPerRtc\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ethUSD\",\"type\":\"uint256\"}],\"name\":\"updateEthPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hotWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RTI\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buyRtc\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferForeignToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hotWalletAddress\",\"type\":\"address\"}],\"name\":\"setHotWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethUSD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawRemainingRtc\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lockSales\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_ethUSD\",\"type\":\"uint256\"}],\"name\":\"EthUsdPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_ethPerRtc\",\"type\":\"uint256\"}],\"name\":\"EthPerRtcUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_rtcPurchased\",\"type\":\"uint256\"}],\"name\":\"RtcPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ForeignTokenTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RTCETHBin is the compiled bytecode used for deploying new contracts.
const RTCETHBin = `608060405234801561001057600080fd5b506000805433600160a060020a031991821681179092556001805490911690911790556005805460ff191660011790556110ab8061004f6000396000f3006080604052600436106100f05763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663094c8bee81146101a65780630d0b8825146101cf57806321370942146101f657806329113bc81461020e57806358e1c1741461023f578063704b6c02146102545780637ed32df6146102755780638da5cb5b1461027d5780639e5fea8a146102925780639fb755d7146102bc578063ac48bd5a146102dd578063cf309012146102f2578063f2fde38b14610307578063f6a5855814610328578063f851a4401461033d578063fa23023b14610352578063ffa1ad7414610367575b3615610146576040805160e560020a62461bcd02815260206004820152601560248201527f64617461206c656e677468206d75737420626520300000000000000000000000604482015290519081900360640190fd5b61014e6103f1565b15156101a4576040805160e560020a62461bcd02815260206004820152601160248201527f627579696e6720727463206661696c6564000000000000000000000000000000604482015290519081900360640190fd5b005b3480156101b257600080fd5b506101bb6106c7565b604080519115158252519081900360200190f35b3480156101db57600080fd5b506101e4610762565b60408051918252519081900360200190f35b34801561020257600080fd5b506101bb600435610768565b34801561021a57600080fd5b50610223610845565b60408051600160a060020a039092168252519081900360200190f35b34801561024b57600080fd5b50610223610854565b34801561026057600080fd5b506101bb600160a060020a036004351661086c565b6101bb6103f1565b34801561028957600080fd5b50610223610967565b34801561029e57600080fd5b506101bb600160a060020a0360043581169060243516604435610976565b3480156102c857600080fd5b506101bb600160a060020a0360043516610bb6565b3480156102e957600080fd5b506101e4610c5c565b3480156102fe57600080fd5b506101bb610c62565b34801561031357600080fd5b506101bb600160a060020a0360043516610c6b565b34801561033457600080fd5b506101bb610d53565b34801561034957600080fd5b50610223610f56565b34801561035e57600080fd5b506101bb610f65565b34801561037357600080fd5b5061037c611003565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103b657818101518382015260200161039e565b50505050905090810190601f1680156103e35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600554600090819060ff1615610451576040805160e560020a62461bcd02815260206004820152601760248201527f73616c65206d757374206e6f74206265206c6f636b6564000000000000000000604482015290519081900360640190fd5b600254600160a060020a031615156104b3576040805160e560020a62461bcd02815260206004820152601860248201527f686f742077616c6c65742063616e7420626520756e7365740000000000000000604482015290519081900360640190fd5b60003411610531576040805160e560020a62461bcd02815260206004820152602360248201527f6d73672076616c7565206d7573742062652067726561746572207468616e207a60448201527f65726f0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60045461055c9061055034670de0b6b3a764000063ffffffff61103a16565b9063ffffffff61106816565b600254604051919250600160a060020a0316903480156108fc02916000818181858888f19350505050158015610596573d6000803e3d6000fd5b50604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815233600482015260248101839052905173ecc043b92834c1ebde65f2181b59597a6588d6169163a9059cbb9160448083019260209291908290030181600087803b15801561060a57600080fd5b505af115801561061e573d6000803e3d6000fd5b505050506040513d602081101561063457600080fd5b5051151561068c576040805160e560020a62461bcd02815260206004820152600f60248201527f7472616e73666572206661696c65640000000000000000000000000000000000604482015290519081900360640190fd5b6040805182815290517fa0a4ba9ef6957a5ef0d6abd66ffc659cc3759a9b34885d1be885b5045d49d0069181900360200190a1600191505090565b60008054600160a060020a03163314806106eb5750600154600160a060020a031633145b15156106f657600080fd5b60055460ff161515610752576040805160e560020a62461bcd02815260206004820152601360248201527f73616c65206d757374206265206c6f636b656400000000000000000000000000604482015290519081900360640190fd5b506005805460ff19169055600190565b60045481565b6000805481908190600160a060020a03163314806107905750600154600160a060020a031633145b151561079b57600080fd5b6003849055670de0b6b3a764000091506107bb828563ffffffff61106816565b90506107ce81600863ffffffff61106816565b60045560035460408051918252517fb8a7d16d8966ae3f48e95e49ed078690c23bd91afb16363bbaaaac00ff99b03f9181900360200190a160045460408051918252517ffffc90eea46c9cbb07fb42de1eaf85778ea803d977cf72f27fca2fbef82ead4f9181900360200190a15060019392505050565b600254600160a060020a031681565b73ecc043b92834c1ebde65f2181b59597a6588d61681565b60008054600160a060020a0316331461088457600080fd5b81600160a060020a03811615156108e5576040805160e560020a62461bcd02815260206004820152601860248201527f6d757374206265206e6f6e207a65726f20616464726573730000000000000000604482015290519081900360640190fd5b600154600160a060020a038481169116141561090057600080fd5b60018054600160a060020a03851673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c9181900360200190a150600192915050565b600054600160a060020a031681565b600080548190600160a060020a031633148061099c5750600154600160a060020a031633145b15156109a757600080fd5b600160a060020a0384161515610a07576040805160e560020a62461bcd02815260206004820181905260248201527f726563697069656e7420616464726573732063616e277420626520656d707479604482015290519081900360640190fd5b600160a060020a03851673ecc043b92834c1ebde65f2181b59597a6588d6161415610a7c576040805160e560020a62461bcd02815260206004820152601260248201527f746f6b656e2063616e2774206265205254430000000000000000000000000000604482015290519081900360640190fd5b50604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038581166004830152602482018590529151869283169163a9059cbb9160448083019260209291908290030181600087803b158015610ae957600080fd5b505af1158015610afd573d6000803e3d6000fd5b505050506040513d6020811015610b1357600080fd5b50511515610b6b576040805160e560020a62461bcd02815260206004820152601560248201527f746f6b656e207472616e73666572206661696c65640000000000000000000000604482015290519081900360640190fd5b604080518481529051600160a060020a0386169133917f10a46ed575affad8e954ae27853b1f89c6da90d8c35f619fc640f8a21bcb78579181900360200190a3506001949350505050565b60008054600160a060020a03163314610bce57600080fd5b60055460ff161515610c2a576040805160e560020a62461bcd02815260206004820152601360248201527f73616c65206d757374206265206c6f636b656400000000000000000000000000604482015290519081900360640190fd5b5060028054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b60035481565b60055460ff1681565b60008054600160a060020a03163314610c8357600080fd5b81600160a060020a0381161515610ce4576040805160e560020a62461bcd02815260206004820152601860248201527f6d757374206265206e6f6e207a65726f20616464726573730000000000000000604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03851690811790915560408051338152602081019290925280517f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09281900390910190a150600192915050565b60008054600160a060020a03163314610d6b57600080fd5b60055460ff161515610dc7576040805160e560020a62461bcd02815260206004820152601360248201527f73616c65206d757374206265206c6f636b656400000000000000000000000000604482015290519081900360640190fd5b604080517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152905173ecc043b92834c1ebde65f2181b59597a6588d6169163a9059cbb91339184916370a0823191602480820192602092909190829003018186803b158015610e3c57600080fd5b505afa158015610e50573d6000803e3d6000fd5b505050506040513d6020811015610e6657600080fd5b5051604080517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a03909316600484015260248301919091525160448083019260209291908290030181600087803b158015610ece57600080fd5b505af1158015610ee2573d6000803e3d6000fd5b505050506040513d6020811015610ef857600080fd5b50511515610f50576040805160e560020a62461bcd02815260206004820152600f60248201527f7472616e73666572206661696c65640000000000000000000000000000000000604482015290519081900360640190fd5b50600190565b600154600160a060020a031681565b60008054600160a060020a0316331480610f895750600154600160a060020a031633145b1515610f9457600080fd5b60055460ff1615610fef576040805160e560020a62461bcd02815260206004820152601760248201527f73616c65206d757374206e6f74206265206c6f636b6564000000000000000000604482015290519081900360640190fd5b506005805460ff1916600190811790915590565b60408051808201909152600a81527f70726f64756374696f6e00000000000000000000000000000000000000000000602082015281565b6000828202831580611056575082848281151561105357fe5b04145b151561106157600080fd5b9392505050565b600080828481151561107657fe5b049493505050505600a165627a7a72305820ca41d479a89ecda0d55eae705e47e46e2fc6a8034efbaf03310e0ad9db2816c40029`

// DeployRTCETH deploys a new Ethereum contract, binding an instance of RTCETH to it.
func DeployRTCETH(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RTCETH, error) {
	parsed, err := abi.JSON(strings.NewReader(RTCETHABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RTCETHBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RTCETH{RTCETHCaller: RTCETHCaller{contract: contract}, RTCETHTransactor: RTCETHTransactor{contract: contract}, RTCETHFilterer: RTCETHFilterer{contract: contract}}, nil
}

// RTCETH is an auto generated Go binding around an Ethereum contract.
type RTCETH struct {
	RTCETHCaller     // Read-only binding to the contract
	RTCETHTransactor // Write-only binding to the contract
	RTCETHFilterer   // Log filterer for contract events
}

// RTCETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type RTCETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTCETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RTCETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTCETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RTCETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTCETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RTCETHSession struct {
	Contract     *RTCETH           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RTCETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RTCETHCallerSession struct {
	Contract *RTCETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RTCETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RTCETHTransactorSession struct {
	Contract     *RTCETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RTCETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type RTCETHRaw struct {
	Contract *RTCETH // Generic contract binding to access the raw methods on
}

// RTCETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RTCETHCallerRaw struct {
	Contract *RTCETHCaller // Generic read-only contract binding to access the raw methods on
}

// RTCETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RTCETHTransactorRaw struct {
	Contract *RTCETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRTCETH creates a new instance of RTCETH, bound to a specific deployed contract.
func NewRTCETH(address common.Address, backend bind.ContractBackend) (*RTCETH, error) {
	contract, err := bindRTCETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RTCETH{RTCETHCaller: RTCETHCaller{contract: contract}, RTCETHTransactor: RTCETHTransactor{contract: contract}, RTCETHFilterer: RTCETHFilterer{contract: contract}}, nil
}

// NewRTCETHCaller creates a new read-only instance of RTCETH, bound to a specific deployed contract.
func NewRTCETHCaller(address common.Address, caller bind.ContractCaller) (*RTCETHCaller, error) {
	contract, err := bindRTCETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RTCETHCaller{contract: contract}, nil
}

// NewRTCETHTransactor creates a new write-only instance of RTCETH, bound to a specific deployed contract.
func NewRTCETHTransactor(address common.Address, transactor bind.ContractTransactor) (*RTCETHTransactor, error) {
	contract, err := bindRTCETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RTCETHTransactor{contract: contract}, nil
}

// NewRTCETHFilterer creates a new log filterer instance of RTCETH, bound to a specific deployed contract.
func NewRTCETHFilterer(address common.Address, filterer bind.ContractFilterer) (*RTCETHFilterer, error) {
	contract, err := bindRTCETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RTCETHFilterer{contract: contract}, nil
}

// bindRTCETH binds a generic wrapper to an already deployed contract.
func bindRTCETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RTCETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RTCETH *RTCETHRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RTCETH.Contract.RTCETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RTCETH *RTCETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCETH.Contract.RTCETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RTCETH *RTCETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RTCETH.Contract.RTCETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RTCETH *RTCETHCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RTCETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RTCETH *RTCETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RTCETH *RTCETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RTCETH.Contract.contract.Transact(opts, method, params...)
}

// RTI is a free data retrieval call binding the contract method 0x58e1c174.
//
// Solidity: function RTI() constant returns(address)
func (_RTCETH *RTCETHCaller) RTI(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RTCETH.contract.Call(opts, out, "RTI")
	return *ret0, err
}

// RTI is a free data retrieval call binding the contract method 0x58e1c174.
//
// Solidity: function RTI() constant returns(address)
func (_RTCETH *RTCETHSession) RTI() (common.Address, error) {
	return _RTCETH.Contract.RTI(&_RTCETH.CallOpts)
}

// RTI is a free data retrieval call binding the contract method 0x58e1c174.
//
// Solidity: function RTI() constant returns(address)
func (_RTCETH *RTCETHCallerSession) RTI() (common.Address, error) {
	return _RTCETH.Contract.RTI(&_RTCETH.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_RTCETH *RTCETHCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RTCETH.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_RTCETH *RTCETHSession) VERSION() (string, error) {
	return _RTCETH.Contract.VERSION(&_RTCETH.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_RTCETH *RTCETHCallerSession) VERSION() (string, error) {
	return _RTCETH.Contract.VERSION(&_RTCETH.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_RTCETH *RTCETHCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RTCETH.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_RTCETH *RTCETHSession) Admin() (common.Address, error) {
	return _RTCETH.Contract.Admin(&_RTCETH.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_RTCETH *RTCETHCallerSession) Admin() (common.Address, error) {
	return _RTCETH.Contract.Admin(&_RTCETH.CallOpts)
}

// EthUSD is a free data retrieval call binding the contract method 0xac48bd5a.
//
// Solidity: function ethUSD() constant returns(uint256)
func (_RTCETH *RTCETHCaller) EthUSD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RTCETH.contract.Call(opts, out, "ethUSD")
	return *ret0, err
}

// EthUSD is a free data retrieval call binding the contract method 0xac48bd5a.
//
// Solidity: function ethUSD() constant returns(uint256)
func (_RTCETH *RTCETHSession) EthUSD() (*big.Int, error) {
	return _RTCETH.Contract.EthUSD(&_RTCETH.CallOpts)
}

// EthUSD is a free data retrieval call binding the contract method 0xac48bd5a.
//
// Solidity: function ethUSD() constant returns(uint256)
func (_RTCETH *RTCETHCallerSession) EthUSD() (*big.Int, error) {
	return _RTCETH.Contract.EthUSD(&_RTCETH.CallOpts)
}

// HotWallet is a free data retrieval call binding the contract method 0x29113bc8.
//
// Solidity: function hotWallet() constant returns(address)
func (_RTCETH *RTCETHCaller) HotWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RTCETH.contract.Call(opts, out, "hotWallet")
	return *ret0, err
}

// HotWallet is a free data retrieval call binding the contract method 0x29113bc8.
//
// Solidity: function hotWallet() constant returns(address)
func (_RTCETH *RTCETHSession) HotWallet() (common.Address, error) {
	return _RTCETH.Contract.HotWallet(&_RTCETH.CallOpts)
}

// HotWallet is a free data retrieval call binding the contract method 0x29113bc8.
//
// Solidity: function hotWallet() constant returns(address)
func (_RTCETH *RTCETHCallerSession) HotWallet() (common.Address, error) {
	return _RTCETH.Contract.HotWallet(&_RTCETH.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_RTCETH *RTCETHCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RTCETH.contract.Call(opts, out, "locked")
	return *ret0, err
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_RTCETH *RTCETHSession) Locked() (bool, error) {
	return _RTCETH.Contract.Locked(&_RTCETH.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_RTCETH *RTCETHCallerSession) Locked() (bool, error) {
	return _RTCETH.Contract.Locked(&_RTCETH.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RTCETH *RTCETHCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RTCETH.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RTCETH *RTCETHSession) Owner() (common.Address, error) {
	return _RTCETH.Contract.Owner(&_RTCETH.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RTCETH *RTCETHCallerSession) Owner() (common.Address, error) {
	return _RTCETH.Contract.Owner(&_RTCETH.CallOpts)
}

// WeiPerRtc is a free data retrieval call binding the contract method 0x0d0b8825.
//
// Solidity: function weiPerRtc() constant returns(uint256)
func (_RTCETH *RTCETHCaller) WeiPerRtc(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RTCETH.contract.Call(opts, out, "weiPerRtc")
	return *ret0, err
}

// WeiPerRtc is a free data retrieval call binding the contract method 0x0d0b8825.
//
// Solidity: function weiPerRtc() constant returns(uint256)
func (_RTCETH *RTCETHSession) WeiPerRtc() (*big.Int, error) {
	return _RTCETH.Contract.WeiPerRtc(&_RTCETH.CallOpts)
}

// WeiPerRtc is a free data retrieval call binding the contract method 0x0d0b8825.
//
// Solidity: function weiPerRtc() constant returns(uint256)
func (_RTCETH *RTCETHCallerSession) WeiPerRtc() (*big.Int, error) {
	return _RTCETH.Contract.WeiPerRtc(&_RTCETH.CallOpts)
}

// BuyRtc is a paid mutator transaction binding the contract method 0x7ed32df6.
//
// Solidity: function buyRtc() returns(bool)
func (_RTCETH *RTCETHTransactor) BuyRtc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCETH.contract.Transact(opts, "buyRtc")
}

// BuyRtc is a paid mutator transaction binding the contract method 0x7ed32df6.
//
// Solidity: function buyRtc() returns(bool)
func (_RTCETH *RTCETHSession) BuyRtc() (*types.Transaction, error) {
	return _RTCETH.Contract.BuyRtc(&_RTCETH.TransactOpts)
}

// BuyRtc is a paid mutator transaction binding the contract method 0x7ed32df6.
//
// Solidity: function buyRtc() returns(bool)
func (_RTCETH *RTCETHTransactorSession) BuyRtc() (*types.Transaction, error) {
	return _RTCETH.Contract.BuyRtc(&_RTCETH.TransactOpts)
}

// LockSales is a paid mutator transaction binding the contract method 0xfa23023b.
//
// Solidity: function lockSales() returns(bool)
func (_RTCETH *RTCETHTransactor) LockSales(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCETH.contract.Transact(opts, "lockSales")
}

// LockSales is a paid mutator transaction binding the contract method 0xfa23023b.
//
// Solidity: function lockSales() returns(bool)
func (_RTCETH *RTCETHSession) LockSales() (*types.Transaction, error) {
	return _RTCETH.Contract.LockSales(&_RTCETH.TransactOpts)
}

// LockSales is a paid mutator transaction binding the contract method 0xfa23023b.
//
// Solidity: function lockSales() returns(bool)
func (_RTCETH *RTCETHTransactorSession) LockSales() (*types.Transaction, error) {
	return _RTCETH.Contract.LockSales(&_RTCETH.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_RTCETH *RTCETHTransactor) SetAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _RTCETH.contract.Transact(opts, "setAdmin", _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_RTCETH *RTCETHSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _RTCETH.Contract.SetAdmin(&_RTCETH.TransactOpts, _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_RTCETH *RTCETHTransactorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _RTCETH.Contract.SetAdmin(&_RTCETH.TransactOpts, _newAdmin)
}

// SetHotWallet is a paid mutator transaction binding the contract method 0x9fb755d7.
//
// Solidity: function setHotWallet(_hotWalletAddress address) returns(bool)
func (_RTCETH *RTCETHTransactor) SetHotWallet(opts *bind.TransactOpts, _hotWalletAddress common.Address) (*types.Transaction, error) {
	return _RTCETH.contract.Transact(opts, "setHotWallet", _hotWalletAddress)
}

// SetHotWallet is a paid mutator transaction binding the contract method 0x9fb755d7.
//
// Solidity: function setHotWallet(_hotWalletAddress address) returns(bool)
func (_RTCETH *RTCETHSession) SetHotWallet(_hotWalletAddress common.Address) (*types.Transaction, error) {
	return _RTCETH.Contract.SetHotWallet(&_RTCETH.TransactOpts, _hotWalletAddress)
}

// SetHotWallet is a paid mutator transaction binding the contract method 0x9fb755d7.
//
// Solidity: function setHotWallet(_hotWalletAddress address) returns(bool)
func (_RTCETH *RTCETHTransactorSession) SetHotWallet(_hotWalletAddress common.Address) (*types.Transaction, error) {
	return _RTCETH.Contract.SetHotWallet(&_RTCETH.TransactOpts, _hotWalletAddress)
}

// TransferForeignToken is a paid mutator transaction binding the contract method 0x9e5fea8a.
//
// Solidity: function transferForeignToken(_tokenAddress address, _recipient address, _amount uint256) returns(bool)
func (_RTCETH *RTCETHTransactor) TransferForeignToken(opts *bind.TransactOpts, _tokenAddress common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCETH.contract.Transact(opts, "transferForeignToken", _tokenAddress, _recipient, _amount)
}

// TransferForeignToken is a paid mutator transaction binding the contract method 0x9e5fea8a.
//
// Solidity: function transferForeignToken(_tokenAddress address, _recipient address, _amount uint256) returns(bool)
func (_RTCETH *RTCETHSession) TransferForeignToken(_tokenAddress common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCETH.Contract.TransferForeignToken(&_RTCETH.TransactOpts, _tokenAddress, _recipient, _amount)
}

// TransferForeignToken is a paid mutator transaction binding the contract method 0x9e5fea8a.
//
// Solidity: function transferForeignToken(_tokenAddress address, _recipient address, _amount uint256) returns(bool)
func (_RTCETH *RTCETHTransactorSession) TransferForeignToken(_tokenAddress common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCETH.Contract.TransferForeignToken(&_RTCETH.TransactOpts, _tokenAddress, _recipient, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_RTCETH *RTCETHTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RTCETH.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_RTCETH *RTCETHSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RTCETH.Contract.TransferOwnership(&_RTCETH.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_RTCETH *RTCETHTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RTCETH.Contract.TransferOwnership(&_RTCETH.TransactOpts, _newOwner)
}

// UnlockSales is a paid mutator transaction binding the contract method 0x094c8bee.
//
// Solidity: function unlockSales() returns(bool)
func (_RTCETH *RTCETHTransactor) UnlockSales(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCETH.contract.Transact(opts, "unlockSales")
}

// UnlockSales is a paid mutator transaction binding the contract method 0x094c8bee.
//
// Solidity: function unlockSales() returns(bool)
func (_RTCETH *RTCETHSession) UnlockSales() (*types.Transaction, error) {
	return _RTCETH.Contract.UnlockSales(&_RTCETH.TransactOpts)
}

// UnlockSales is a paid mutator transaction binding the contract method 0x094c8bee.
//
// Solidity: function unlockSales() returns(bool)
func (_RTCETH *RTCETHTransactorSession) UnlockSales() (*types.Transaction, error) {
	return _RTCETH.Contract.UnlockSales(&_RTCETH.TransactOpts)
}

// UpdateEthPrice is a paid mutator transaction binding the contract method 0x21370942.
//
// Solidity: function updateEthPrice(_ethUSD uint256) returns(bool)
func (_RTCETH *RTCETHTransactor) UpdateEthPrice(opts *bind.TransactOpts, _ethUSD *big.Int) (*types.Transaction, error) {
	return _RTCETH.contract.Transact(opts, "updateEthPrice", _ethUSD)
}

// UpdateEthPrice is a paid mutator transaction binding the contract method 0x21370942.
//
// Solidity: function updateEthPrice(_ethUSD uint256) returns(bool)
func (_RTCETH *RTCETHSession) UpdateEthPrice(_ethUSD *big.Int) (*types.Transaction, error) {
	return _RTCETH.Contract.UpdateEthPrice(&_RTCETH.TransactOpts, _ethUSD)
}

// UpdateEthPrice is a paid mutator transaction binding the contract method 0x21370942.
//
// Solidity: function updateEthPrice(_ethUSD uint256) returns(bool)
func (_RTCETH *RTCETHTransactorSession) UpdateEthPrice(_ethUSD *big.Int) (*types.Transaction, error) {
	return _RTCETH.Contract.UpdateEthPrice(&_RTCETH.TransactOpts, _ethUSD)
}

// WithdrawRemainingRtc is a paid mutator transaction binding the contract method 0xf6a58558.
//
// Solidity: function withdrawRemainingRtc() returns(bool)
func (_RTCETH *RTCETHTransactor) WithdrawRemainingRtc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCETH.contract.Transact(opts, "withdrawRemainingRtc")
}

// WithdrawRemainingRtc is a paid mutator transaction binding the contract method 0xf6a58558.
//
// Solidity: function withdrawRemainingRtc() returns(bool)
func (_RTCETH *RTCETHSession) WithdrawRemainingRtc() (*types.Transaction, error) {
	return _RTCETH.Contract.WithdrawRemainingRtc(&_RTCETH.TransactOpts)
}

// WithdrawRemainingRtc is a paid mutator transaction binding the contract method 0xf6a58558.
//
// Solidity: function withdrawRemainingRtc() returns(bool)
func (_RTCETH *RTCETHTransactorSession) WithdrawRemainingRtc() (*types.Transaction, error) {
	return _RTCETH.Contract.WithdrawRemainingRtc(&_RTCETH.TransactOpts)
}

// RTCETHAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the RTCETH contract.
type RTCETHAdminSetIterator struct {
	Event *RTCETHAdminSet // Event containing the contract specifics and raw log

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
func (it *RTCETHAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCETHAdminSet)
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
		it.Event = new(RTCETHAdminSet)
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
func (it *RTCETHAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCETHAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCETHAdminSet represents a AdminSet event raised by the RTCETH contract.
type RTCETHAdminSet struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: e AdminSet(_admin address)
func (_RTCETH *RTCETHFilterer) FilterAdminSet(opts *bind.FilterOpts) (*RTCETHAdminSetIterator, error) {

	logs, sub, err := _RTCETH.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &RTCETHAdminSetIterator{contract: _RTCETH.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: e AdminSet(_admin address)
func (_RTCETH *RTCETHFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *RTCETHAdminSet) (event.Subscription, error) {

	logs, sub, err := _RTCETH.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCETHAdminSet)
				if err := _RTCETH.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// RTCETHEthPerRtcUpdatedIterator is returned from FilterEthPerRtcUpdated and is used to iterate over the raw logs and unpacked data for EthPerRtcUpdated events raised by the RTCETH contract.
type RTCETHEthPerRtcUpdatedIterator struct {
	Event *RTCETHEthPerRtcUpdated // Event containing the contract specifics and raw log

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
func (it *RTCETHEthPerRtcUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCETHEthPerRtcUpdated)
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
		it.Event = new(RTCETHEthPerRtcUpdated)
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
func (it *RTCETHEthPerRtcUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCETHEthPerRtcUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCETHEthPerRtcUpdated represents a EthPerRtcUpdated event raised by the RTCETH contract.
type RTCETHEthPerRtcUpdated struct {
	EthPerRtc *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthPerRtcUpdated is a free log retrieval operation binding the contract event 0xfffc90eea46c9cbb07fb42de1eaf85778ea803d977cf72f27fca2fbef82ead4f.
//
// Solidity: e EthPerRtcUpdated(_ethPerRtc uint256)
func (_RTCETH *RTCETHFilterer) FilterEthPerRtcUpdated(opts *bind.FilterOpts) (*RTCETHEthPerRtcUpdatedIterator, error) {

	logs, sub, err := _RTCETH.contract.FilterLogs(opts, "EthPerRtcUpdated")
	if err != nil {
		return nil, err
	}
	return &RTCETHEthPerRtcUpdatedIterator{contract: _RTCETH.contract, event: "EthPerRtcUpdated", logs: logs, sub: sub}, nil
}

// WatchEthPerRtcUpdated is a free log subscription operation binding the contract event 0xfffc90eea46c9cbb07fb42de1eaf85778ea803d977cf72f27fca2fbef82ead4f.
//
// Solidity: e EthPerRtcUpdated(_ethPerRtc uint256)
func (_RTCETH *RTCETHFilterer) WatchEthPerRtcUpdated(opts *bind.WatchOpts, sink chan<- *RTCETHEthPerRtcUpdated) (event.Subscription, error) {

	logs, sub, err := _RTCETH.contract.WatchLogs(opts, "EthPerRtcUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCETHEthPerRtcUpdated)
				if err := _RTCETH.contract.UnpackLog(event, "EthPerRtcUpdated", log); err != nil {
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

// RTCETHEthUsdPriceUpdatedIterator is returned from FilterEthUsdPriceUpdated and is used to iterate over the raw logs and unpacked data for EthUsdPriceUpdated events raised by the RTCETH contract.
type RTCETHEthUsdPriceUpdatedIterator struct {
	Event *RTCETHEthUsdPriceUpdated // Event containing the contract specifics and raw log

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
func (it *RTCETHEthUsdPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCETHEthUsdPriceUpdated)
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
		it.Event = new(RTCETHEthUsdPriceUpdated)
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
func (it *RTCETHEthUsdPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCETHEthUsdPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCETHEthUsdPriceUpdated represents a EthUsdPriceUpdated event raised by the RTCETH contract.
type RTCETHEthUsdPriceUpdated struct {
	EthUSD *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthUsdPriceUpdated is a free log retrieval operation binding the contract event 0xb8a7d16d8966ae3f48e95e49ed078690c23bd91afb16363bbaaaac00ff99b03f.
//
// Solidity: e EthUsdPriceUpdated(_ethUSD uint256)
func (_RTCETH *RTCETHFilterer) FilterEthUsdPriceUpdated(opts *bind.FilterOpts) (*RTCETHEthUsdPriceUpdatedIterator, error) {

	logs, sub, err := _RTCETH.contract.FilterLogs(opts, "EthUsdPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &RTCETHEthUsdPriceUpdatedIterator{contract: _RTCETH.contract, event: "EthUsdPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchEthUsdPriceUpdated is a free log subscription operation binding the contract event 0xb8a7d16d8966ae3f48e95e49ed078690c23bd91afb16363bbaaaac00ff99b03f.
//
// Solidity: e EthUsdPriceUpdated(_ethUSD uint256)
func (_RTCETH *RTCETHFilterer) WatchEthUsdPriceUpdated(opts *bind.WatchOpts, sink chan<- *RTCETHEthUsdPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _RTCETH.contract.WatchLogs(opts, "EthUsdPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCETHEthUsdPriceUpdated)
				if err := _RTCETH.contract.UnpackLog(event, "EthUsdPriceUpdated", log); err != nil {
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

// RTCETHForeignTokenTransferIterator is returned from FilterForeignTokenTransfer and is used to iterate over the raw logs and unpacked data for ForeignTokenTransfer events raised by the RTCETH contract.
type RTCETHForeignTokenTransferIterator struct {
	Event *RTCETHForeignTokenTransfer // Event containing the contract specifics and raw log

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
func (it *RTCETHForeignTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCETHForeignTokenTransfer)
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
		it.Event = new(RTCETHForeignTokenTransfer)
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
func (it *RTCETHForeignTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCETHForeignTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCETHForeignTokenTransfer represents a ForeignTokenTransfer event raised by the RTCETH contract.
type RTCETHForeignTokenTransfer struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterForeignTokenTransfer is a free log retrieval operation binding the contract event 0x10a46ed575affad8e954ae27853b1f89c6da90d8c35f619fc640f8a21bcb7857.
//
// Solidity: e ForeignTokenTransfer(_sender indexed address, _recipient indexed address, _amount uint256)
func (_RTCETH *RTCETHFilterer) FilterForeignTokenTransfer(opts *bind.FilterOpts, _sender []common.Address, _recipient []common.Address) (*RTCETHForeignTokenTransferIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _RTCETH.contract.FilterLogs(opts, "ForeignTokenTransfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return &RTCETHForeignTokenTransferIterator{contract: _RTCETH.contract, event: "ForeignTokenTransfer", logs: logs, sub: sub}, nil
}

// WatchForeignTokenTransfer is a free log subscription operation binding the contract event 0x10a46ed575affad8e954ae27853b1f89c6da90d8c35f619fc640f8a21bcb7857.
//
// Solidity: e ForeignTokenTransfer(_sender indexed address, _recipient indexed address, _amount uint256)
func (_RTCETH *RTCETHFilterer) WatchForeignTokenTransfer(opts *bind.WatchOpts, sink chan<- *RTCETHForeignTokenTransfer, _sender []common.Address, _recipient []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _RTCETH.contract.WatchLogs(opts, "ForeignTokenTransfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCETHForeignTokenTransfer)
				if err := _RTCETH.contract.UnpackLog(event, "ForeignTokenTransfer", log); err != nil {
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

// RTCETHOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RTCETH contract.
type RTCETHOwnershipTransferredIterator struct {
	Event *RTCETHOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RTCETHOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCETHOwnershipTransferred)
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
		it.Event = new(RTCETHOwnershipTransferred)
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
func (it *RTCETHOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCETHOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCETHOwnershipTransferred represents a OwnershipTransferred event raised by the RTCETH contract.
type RTCETHOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_previousOwner address, _newOwner address)
func (_RTCETH *RTCETHFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts) (*RTCETHOwnershipTransferredIterator, error) {

	logs, sub, err := _RTCETH.contract.FilterLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &RTCETHOwnershipTransferredIterator{contract: _RTCETH.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_previousOwner address, _newOwner address)
func (_RTCETH *RTCETHFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RTCETHOwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _RTCETH.contract.WatchLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCETHOwnershipTransferred)
				if err := _RTCETH.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RTCETHRtcPurchasedIterator is returned from FilterRtcPurchased and is used to iterate over the raw logs and unpacked data for RtcPurchased events raised by the RTCETH contract.
type RTCETHRtcPurchasedIterator struct {
	Event *RTCETHRtcPurchased // Event containing the contract specifics and raw log

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
func (it *RTCETHRtcPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCETHRtcPurchased)
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
		it.Event = new(RTCETHRtcPurchased)
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
func (it *RTCETHRtcPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCETHRtcPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCETHRtcPurchased represents a RtcPurchased event raised by the RTCETH contract.
type RTCETHRtcPurchased struct {
	RtcPurchased *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRtcPurchased is a free log retrieval operation binding the contract event 0xa0a4ba9ef6957a5ef0d6abd66ffc659cc3759a9b34885d1be885b5045d49d006.
//
// Solidity: e RtcPurchased(_rtcPurchased uint256)
func (_RTCETH *RTCETHFilterer) FilterRtcPurchased(opts *bind.FilterOpts) (*RTCETHRtcPurchasedIterator, error) {

	logs, sub, err := _RTCETH.contract.FilterLogs(opts, "RtcPurchased")
	if err != nil {
		return nil, err
	}
	return &RTCETHRtcPurchasedIterator{contract: _RTCETH.contract, event: "RtcPurchased", logs: logs, sub: sub}, nil
}

// WatchRtcPurchased is a free log subscription operation binding the contract event 0xa0a4ba9ef6957a5ef0d6abd66ffc659cc3759a9b34885d1be885b5045d49d006.
//
// Solidity: e RtcPurchased(_rtcPurchased uint256)
func (_RTCETH *RTCETHFilterer) WatchRtcPurchased(opts *bind.WatchOpts, sink chan<- *RTCETHRtcPurchased) (event.Subscription, error) {

	logs, sub, err := _RTCETH.contract.WatchLogs(opts, "RtcPurchased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCETHRtcPurchased)
				if err := _RTCETH.contract.UnpackLog(event, "RtcPurchased", log); err != nil {
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
