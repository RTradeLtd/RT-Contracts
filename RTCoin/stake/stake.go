// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// StakeABI is the input ABI used to generate the binding from.
const StakeABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"disableNewStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"allowNewStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stakeNumber\",\"type\":\"uint256\"}],\"name\":\"withdrawInitialStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKENADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakes\",\"outputs\":[{\"name\":\"initialStake\",\"type\":\"uint256\"},{\"name\":\"blockLocked\",\"type\":\"uint256\"},{\"name\":\"blockUnlocked\",\"type\":\"uint256\"},{\"name\":\"releaseDate\",\"type\":\"uint256\"},{\"name\":\"totalCoinsToMint\",\"type\":\"uint256\"},{\"name\":\"coinsMinted\",\"type\":\"uint256\"},{\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"name\":\"lastBlockWithdrawn\",\"type\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RTI\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newStakesAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stakeNumber\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canMint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_numRTC\",\"type\":\"uint256\"}],\"name\":\"depositStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setRTI\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"numberOfStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"activeStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"internalRTCBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakesDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakesEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_stakeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_coinsToMint\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_releaseDate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_releaseBlock\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_stakeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"StakeRewardWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_stakeNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"InitialStakeWithdrawn\",\"type\":\"event\"}]"

// StakeBin is the compiled bytecode used for deploying new contracts.
const StakeBin = `608060405260008054600160a060020a031990811673e9aec23c620681a59e2111785b0d35a90498128f179182905560018054909116600160a060020a039290921691909117905534801561005357600080fd5b50600054600160a060020a031615156100cd57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f746f6b656e2061646472657373206e6f74207365740000000000000000000000604482015290519081900360640190fd5b60038054600160a060020a03191633179055611651806100ee6000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166304b5723d81146100df578063211db50d1461010857806329830ccc1461011d578063516f898614610135578063584b62a11461016657806358e1c174146101e6578063594548d5146101fb578063a0712d6814610210578063beb9716d14610228578063cb82cc8f1461023d578063deb78ab614610255578063dfef667914610276578063ed2f2369146102a9578063f1610821146102be578063f851a440146102df575b600080fd5b3480156100eb57600080fd5b506100f46102f4565b604080519115158252519081900360200190f35b34801561011457600080fd5b506100f461037d565b34801561012957600080fd5b506100f4600435610550565b34801561014157600080fd5b5061014a610907565b60408051600160a060020a039092168252519081900360200190f35b34801561017257600080fd5b5061018a600160a060020a0360043516602435610916565b604051808a81526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020018260028111156101ca57fe5b60ff168152602001995050505050505050505060405180910390f35b3480156101f257600080fd5b5061014a61096c565b34801561020757600080fd5b506100f461097b565b34801561021c57600080fd5b506100f460043561099c565b34801561023457600080fd5b506100f4610db6565b34801561024957600080fd5b506100f4600435610e2c565b34801561026157600080fd5b506100f4600160a060020a03600435166112c2565b34801561028257600080fd5b50610297600160a060020a0360043516611363565b60408051918252519081900360200190f35b3480156102b557600080fd5b50610297611375565b3480156102ca57600080fd5b50610297600160a060020a036004351661137b565b3480156102eb57600080fd5b5061014a61138d565b600354600090600160a060020a03163314610359576040805160e560020a62461bcd02815260206004820152601360248201527f73656e646572206973206e6f742061646d696e00000000000000000000000000604482015290519081900360640190fd5b506003805474ff000000000000000000000000000000000000000019169055600190565b600354600090600160a060020a031633146103e2576040805160e560020a62461bcd02815260206004820152601360248201527f73656e646572206973206e6f742061646d696e00000000000000000000000000604482015290519081900360640190fd5b6003805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055600154604080517f272caf6900000000000000000000000000000000000000000000000000000000815290513092600160a060020a03169163272caf69916004808301926020929190829003018186803b15801561047457600080fd5b505afa158015610488573d6000803e3d6000fd5b505050506040513d602081101561049e57600080fd5b5051600160a060020a03161461054a576040805160e560020a62461bcd02815260206004820152604a60248201527f72746320746f6b656e20636f6e7472616374206973206e6f742073657420746f60448201527f20757365207468697320636f6e747261637420617320746865207374616b696e60648201527f6720636f6e747261637400000000000000000000000000000000000000000000608482015290519081900360a40190fd5b50600190565b60008082600133600090815260046020908152604080832085845290915290206008015460ff16600281111561058257fe5b146105d7576040805160e560020a62461bcd02815260206004820152601360248201527f7374616b65206973206e6f742061637469766500000000000000000000000000604482015290519081900360640190fd5b336000908152600460209081526040808320848452909152902060030154421080159061062157503360009081526004602090815260408083208484529091529020600201544310155b15156106c3576040805160e560020a62461bcd02815260206004820152604160248201527f617474656d7074696e6720746f20776974686472617720696e697469616c207360448201527f74616b65206265666f726520756e6c6f636b20626c6f636b20616e642064617460648201527f6500000000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b3360008181526004602090815260408083208584528252808320549383526006909152902054101561073f576040805160e560020a62461bcd02815260206004820152601c60248201527f696e76616c696420696e7465726e616c207274632062616c616e636500000000604482015290519081900360640190fd5b336000908152600460209081526040808320878452909152902080546008909101805460ff191660029081179091555490925061077d90600161139c565b600255336000908152600660205260409020546107a0908363ffffffff61139c16565b33600081815260066020908152604091829020939093558051858152905187937f7d252c33d474583922a2f7a0c2f4d04631095dbd4e35b09adc7f801ec3e743f7928290030190a3600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018590529051600160a060020a039092169163a9059cbb916044808201926020929091908290030181600087803b15801561085557600080fd5b505af1158015610869573d6000803e3d6000fd5b505050506040513d602081101561087f57600080fd5b505115156108fd576040805160e560020a62461bcd02815260206004820152603960248201527f756e61626c6520746f207472616e7366657220746f6b656e73206c696b656c7960448201527f2064756520746f20696e636f72726563742062616c616e636500000000000000606482015290519081900360840190fd5b5060019392505050565b600054600160a060020a031681565b600460208181526000938452604080852090915291835291208054600182015460028301546003840154948401546005850154600686015460078701546008909701549597949693959293919290919060ff1689565b600154600160a060020a031681565b60035474010000000000000000000000000000000000000000900460ff1681565b600080828180600133600090815260046020908152604080832087845290915290206008015460ff1660028111156109d057fe5b1480610a075750600233600090815260046020908152604080832087845290915290206008015460ff166002811115610a0557fe5b145b1515610a83576040805160e560020a62461bcd02815260206004820152603860248201527f7374616b65206d75737420626520616374697665206f7220696e61637469766560448201527f20696e206f7264657220746f206d696e7420746f6b656e730000000000000000606482015290519081900360840190fd5b3360009081526004602081815260408084208785529091529091209081015460059091015410610b23576040805160e560020a62461bcd02815260206004820152602c60248201527f63757272656e7420636f696e73206d696e746564206d757374206265206c657360448201527f73207468616e20746f74616c0000000000000000000000000000000000000000606482015290519081900360840190fd5b50503360009081526004602090815260408083208484529091529020600701544390808211610bc2576040805160e560020a62461bcd02815260206004820152603560248201527f63757272656e7420626c6f636b206d757374206265206f6e652068696768657260448201527f207468616e206c617374207769746864726177616c0000000000000000000000606482015290519081900360840190fd5b610bcb866113b1565b3360009081526004602081815260408084208b85529091529091209081015460059091015491955090610bfe9086611488565b1115610c7a576040805160e560020a62461bcd02815260206004820152602260248201527f746f74616c20636f696e73206d696e74656420646f6573206e6f74206164642060448201527f7570000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b336000908152600460209081526040808320898452909152902060050154610ca8908563ffffffff61148816565b3360008181526004602090815260408083208b84528252918290206005810194909455436007909401939093558051878152905189937f275541ddbc93a3fb1e5e94000231500252d2ba460de93bd1cf285e68563c1a64928290030190a3600154604080517f40c10f19000000000000000000000000000000000000000000000000000000008152336004820152602481018790529051600160a060020a03909216916340c10f19916044808201926020929091908290030181600087803b158015610d7357600080fd5b505af1158015610d87573d6000803e3d6000fd5b505050506040513d6020811015610d9d57600080fd5b50511515610daa57600080fd5b50600195945050505050565b600030600160a060020a0316600160009054906101000a9004600160a060020a0316600160a060020a031663272caf696040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561047457600080fd5b6000806000806000806000610e3f6115ce565b88610e48610db6565b1515610ec4576040805160e560020a62461bcd02815260206004820152602960248201527f7374616b696e6720636f6e747261637420697320756e61626c6520746f206d6960448201527f6e7420746f6b656e730000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60035474010000000000000000000000000000000000000000900460ff161515610f38576040805160e560020a62461bcd02815260206004820152601a60248201527f6e6577207374616b657320617265206e6f7420616c6c6f776564000000000000604482015290519081900360640190fd5b670de0b6b3a7640000811015610fbe576040805160e560020a62461bcd02815260206004820152602c60248201527f737065636966696564207374616b65206973206c6f776572207468616e206d6960448201527f6e696d756d20616d6f756e740000000000000000000000000000000000000000606482015290519081900360840190fd5b610fc7336114a1565b9750610fd28a6114bc565b96509650965096509650610120604051908101604052808b8152602001888152602001878152602001868152602001858152602001600081526020018481526020018881526020016001600281111561102757fe5b8152509150816004600033600160a060020a0316600160a060020a0316815260200190815260200160002060008a8152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701556101008201518160080160006101000a81548160ff021916908360028111156110d857fe5b021790555050336000908152600560205260409020546111009150600163ffffffff61148816565b3360009081526005602090815260408083209390935560069052205461112c908b63ffffffff61148816565b3360009081526006602052604090205560025461115090600163ffffffff61148816565b60025560408051858152602081018790528082018890529051899133917f1a325385f16807e99fb688b597db78b00faee313dcf02e882dd16daab6fc3e1f9181900360600190a3600154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018d90529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b15801561120a57600080fd5b505af115801561121e573d6000803e3d6000fd5b505050506040513d602081101561123457600080fd5b505115156112b2576040805160e560020a62461bcd02815260206004820152602b60248201527f7472616e736665722066726f6d206661696c65642c206c696b656c79206e656560448201527f647320617070726f76616c000000000000000000000000000000000000000000606482015290519081900360840190fd5b5060019998505050505050505050565b600354600090600160a060020a03163314611327576040805160e560020a62461bcd02815260206004820152601360248201527f73656e646572206973206e6f742061646d696e00000000000000000000000000604482015290519081900360640190fd5b5060018054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1991821681178355600080549092161790555b919050565b60056020526000908152604090205481565b60025481565b60066020526000908152604090205481565b600354600160a060020a031681565b6000828211156113ab57600080fd5b50900390565b60008060008060008060006113c588611549565b3360009081526004602090815260408083208c845290915290206007015490965094506113f8868663ffffffff61139c16565b3360009081526004602090815260408083208c845290915290206006015490945061142a90859063ffffffff61159016565b3360009081526004602081815260408084208d855290915290912090810154600590910154919850935091506114608288611488565b90508281111561147d5761147a818463ffffffff61139c16565b96505b505050505050919050565b60008282018381101561149a57600080fd5b9392505050565b600160a060020a031660009081526005602052604090205490565b4360008080806114d385600563ffffffff61148816565b93506114f76114ea6005600f63ffffffff61159016565b429063ffffffff61148816565b92506115118667016345785d8a000063ffffffff61159016565b915061152b82670de0b6b3a764000063ffffffff6115b716565b915061153e82600563ffffffff6115b716565b905091939590929450565b3360009081526004602090815260408083208484529091529020600201544390811061135e5750336000908152600460209081526040808320938352929052206002015490565b60008282028315806115ac57508284828115156115a957fe5b04145b151561149a57600080fd5b60008082848115156115c557fe5b04949350505050565b6101206040519081016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000600281111561162057fe5b9052905600a165627a7a723058208cd3ab648c4c933b6c482821593f25b4d09fa5ce9f28a5f74d70ee1e4ad13f220029`

// DeployStake deploys a new Ethereum contract, binding an instance of Stake to it.
func DeployStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Stake, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// Stake is an auto generated Go binding around an Ethereum contract.
type Stake struct {
	StakeCaller     // Read-only binding to the contract
	StakeTransactor // Write-only binding to the contract
	StakeFilterer   // Log filterer for contract events
}

// StakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeSession struct {
	Contract     *Stake            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeCallerSession struct {
	Contract *StakeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeTransactorSession struct {
	Contract     *StakeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeRaw struct {
	Contract *Stake // Generic contract binding to access the raw methods on
}

// StakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeCallerRaw struct {
	Contract *StakeCaller // Generic read-only contract binding to access the raw methods on
}

// StakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeTransactorRaw struct {
	Contract *StakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStake creates a new instance of Stake, bound to a specific deployed contract.
func NewStake(address common.Address, backend bind.ContractBackend) (*Stake, error) {
	contract, err := bindStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// NewStakeCaller creates a new read-only instance of Stake, bound to a specific deployed contract.
func NewStakeCaller(address common.Address, caller bind.ContractCaller) (*StakeCaller, error) {
	contract, err := bindStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeCaller{contract: contract}, nil
}

// NewStakeTransactor creates a new write-only instance of Stake, bound to a specific deployed contract.
func NewStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeTransactor, error) {
	contract, err := bindStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTransactor{contract: contract}, nil
}

// NewStakeFilterer creates a new log filterer instance of Stake, bound to a specific deployed contract.
func NewStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeFilterer, error) {
	contract, err := bindStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeFilterer{contract: contract}, nil
}

// bindStake binds a generic wrapper to an already deployed contract.
func bindStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.StakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transact(opts, method, params...)
}

// RTI is a free data retrieval call binding the contract method 0x58e1c174.
//
// Solidity: function RTI() constant returns(address)
func (_Stake *StakeCaller) RTI(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "RTI")
	return *ret0, err
}

// RTI is a free data retrieval call binding the contract method 0x58e1c174.
//
// Solidity: function RTI() constant returns(address)
func (_Stake *StakeSession) RTI() (common.Address, error) {
	return _Stake.Contract.RTI(&_Stake.CallOpts)
}

// RTI is a free data retrieval call binding the contract method 0x58e1c174.
//
// Solidity: function RTI() constant returns(address)
func (_Stake *StakeCallerSession) RTI() (common.Address, error) {
	return _Stake.Contract.RTI(&_Stake.CallOpts)
}

// TOKENADDRESS is a free data retrieval call binding the contract method 0x516f8986.
//
// Solidity: function TOKENADDRESS() constant returns(address)
func (_Stake *StakeCaller) TOKENADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "TOKENADDRESS")
	return *ret0, err
}

// TOKENADDRESS is a free data retrieval call binding the contract method 0x516f8986.
//
// Solidity: function TOKENADDRESS() constant returns(address)
func (_Stake *StakeSession) TOKENADDRESS() (common.Address, error) {
	return _Stake.Contract.TOKENADDRESS(&_Stake.CallOpts)
}

// TOKENADDRESS is a free data retrieval call binding the contract method 0x516f8986.
//
// Solidity: function TOKENADDRESS() constant returns(address)
func (_Stake *StakeCallerSession) TOKENADDRESS() (common.Address, error) {
	return _Stake.Contract.TOKENADDRESS(&_Stake.CallOpts)
}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() constant returns(uint256)
func (_Stake *StakeCaller) ActiveStakes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "activeStakes")
	return *ret0, err
}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() constant returns(uint256)
func (_Stake *StakeSession) ActiveStakes() (*big.Int, error) {
	return _Stake.Contract.ActiveStakes(&_Stake.CallOpts)
}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() constant returns(uint256)
func (_Stake *StakeCallerSession) ActiveStakes() (*big.Int, error) {
	return _Stake.Contract.ActiveStakes(&_Stake.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Stake *StakeCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Stake *StakeSession) Admin() (common.Address, error) {
	return _Stake.Contract.Admin(&_Stake.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Stake *StakeCallerSession) Admin() (common.Address, error) {
	return _Stake.Contract.Admin(&_Stake.CallOpts)
}

// CanMint is a free data retrieval call binding the contract method 0xbeb9716d.
//
// Solidity: function canMint() constant returns(bool)
func (_Stake *StakeCaller) CanMint(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "canMint")
	return *ret0, err
}

// CanMint is a free data retrieval call binding the contract method 0xbeb9716d.
//
// Solidity: function canMint() constant returns(bool)
func (_Stake *StakeSession) CanMint() (bool, error) {
	return _Stake.Contract.CanMint(&_Stake.CallOpts)
}

// CanMint is a free data retrieval call binding the contract method 0xbeb9716d.
//
// Solidity: function canMint() constant returns(bool)
func (_Stake *StakeCallerSession) CanMint() (bool, error) {
	return _Stake.Contract.CanMint(&_Stake.CallOpts)
}

// InternalRTCBalances is a free data retrieval call binding the contract method 0xf1610821.
//
// Solidity: function internalRTCBalances( address) constant returns(uint256)
func (_Stake *StakeCaller) InternalRTCBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "internalRTCBalances", arg0)
	return *ret0, err
}

// InternalRTCBalances is a free data retrieval call binding the contract method 0xf1610821.
//
// Solidity: function internalRTCBalances( address) constant returns(uint256)
func (_Stake *StakeSession) InternalRTCBalances(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.InternalRTCBalances(&_Stake.CallOpts, arg0)
}

// InternalRTCBalances is a free data retrieval call binding the contract method 0xf1610821.
//
// Solidity: function internalRTCBalances( address) constant returns(uint256)
func (_Stake *StakeCallerSession) InternalRTCBalances(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.InternalRTCBalances(&_Stake.CallOpts, arg0)
}

// NewStakesAllowed is a free data retrieval call binding the contract method 0x594548d5.
//
// Solidity: function newStakesAllowed() constant returns(bool)
func (_Stake *StakeCaller) NewStakesAllowed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "newStakesAllowed")
	return *ret0, err
}

// NewStakesAllowed is a free data retrieval call binding the contract method 0x594548d5.
//
// Solidity: function newStakesAllowed() constant returns(bool)
func (_Stake *StakeSession) NewStakesAllowed() (bool, error) {
	return _Stake.Contract.NewStakesAllowed(&_Stake.CallOpts)
}

// NewStakesAllowed is a free data retrieval call binding the contract method 0x594548d5.
//
// Solidity: function newStakesAllowed() constant returns(bool)
func (_Stake *StakeCallerSession) NewStakesAllowed() (bool, error) {
	return _Stake.Contract.NewStakesAllowed(&_Stake.CallOpts)
}

// NumberOfStakes is a free data retrieval call binding the contract method 0xdfef6679.
//
// Solidity: function numberOfStakes( address) constant returns(uint256)
func (_Stake *StakeCaller) NumberOfStakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "numberOfStakes", arg0)
	return *ret0, err
}

// NumberOfStakes is a free data retrieval call binding the contract method 0xdfef6679.
//
// Solidity: function numberOfStakes( address) constant returns(uint256)
func (_Stake *StakeSession) NumberOfStakes(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.NumberOfStakes(&_Stake.CallOpts, arg0)
}

// NumberOfStakes is a free data retrieval call binding the contract method 0xdfef6679.
//
// Solidity: function numberOfStakes( address) constant returns(uint256)
func (_Stake *StakeCallerSession) NumberOfStakes(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.NumberOfStakes(&_Stake.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x584b62a1.
//
// Solidity: function stakes( address,  uint256) constant returns(initialStake uint256, blockLocked uint256, blockUnlocked uint256, releaseDate uint256, totalCoinsToMint uint256, coinsMinted uint256, rewardPerBlock uint256, lastBlockWithdrawn uint256, state uint8)
func (_Stake *StakeCaller) Stakes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	InitialStake       *big.Int
	BlockLocked        *big.Int
	BlockUnlocked      *big.Int
	ReleaseDate        *big.Int
	TotalCoinsToMint   *big.Int
	CoinsMinted        *big.Int
	RewardPerBlock     *big.Int
	LastBlockWithdrawn *big.Int
	State              uint8
}, error) {
	ret := new(struct {
		InitialStake       *big.Int
		BlockLocked        *big.Int
		BlockUnlocked      *big.Int
		ReleaseDate        *big.Int
		TotalCoinsToMint   *big.Int
		CoinsMinted        *big.Int
		RewardPerBlock     *big.Int
		LastBlockWithdrawn *big.Int
		State              uint8
	})
	out := ret
	err := _Stake.contract.Call(opts, out, "stakes", arg0, arg1)
	return *ret, err
}

// Stakes is a free data retrieval call binding the contract method 0x584b62a1.
//
// Solidity: function stakes( address,  uint256) constant returns(initialStake uint256, blockLocked uint256, blockUnlocked uint256, releaseDate uint256, totalCoinsToMint uint256, coinsMinted uint256, rewardPerBlock uint256, lastBlockWithdrawn uint256, state uint8)
func (_Stake *StakeSession) Stakes(arg0 common.Address, arg1 *big.Int) (struct {
	InitialStake       *big.Int
	BlockLocked        *big.Int
	BlockUnlocked      *big.Int
	ReleaseDate        *big.Int
	TotalCoinsToMint   *big.Int
	CoinsMinted        *big.Int
	RewardPerBlock     *big.Int
	LastBlockWithdrawn *big.Int
	State              uint8
}, error) {
	return _Stake.Contract.Stakes(&_Stake.CallOpts, arg0, arg1)
}

// Stakes is a free data retrieval call binding the contract method 0x584b62a1.
//
// Solidity: function stakes( address,  uint256) constant returns(initialStake uint256, blockLocked uint256, blockUnlocked uint256, releaseDate uint256, totalCoinsToMint uint256, coinsMinted uint256, rewardPerBlock uint256, lastBlockWithdrawn uint256, state uint8)
func (_Stake *StakeCallerSession) Stakes(arg0 common.Address, arg1 *big.Int) (struct {
	InitialStake       *big.Int
	BlockLocked        *big.Int
	BlockUnlocked      *big.Int
	ReleaseDate        *big.Int
	TotalCoinsToMint   *big.Int
	CoinsMinted        *big.Int
	RewardPerBlock     *big.Int
	LastBlockWithdrawn *big.Int
	State              uint8
}, error) {
	return _Stake.Contract.Stakes(&_Stake.CallOpts, arg0, arg1)
}

// AllowNewStakes is a paid mutator transaction binding the contract method 0x211db50d.
//
// Solidity: function allowNewStakes() returns(bool)
func (_Stake *StakeTransactor) AllowNewStakes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "allowNewStakes")
}

// AllowNewStakes is a paid mutator transaction binding the contract method 0x211db50d.
//
// Solidity: function allowNewStakes() returns(bool)
func (_Stake *StakeSession) AllowNewStakes() (*types.Transaction, error) {
	return _Stake.Contract.AllowNewStakes(&_Stake.TransactOpts)
}

// AllowNewStakes is a paid mutator transaction binding the contract method 0x211db50d.
//
// Solidity: function allowNewStakes() returns(bool)
func (_Stake *StakeTransactorSession) AllowNewStakes() (*types.Transaction, error) {
	return _Stake.Contract.AllowNewStakes(&_Stake.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0xcb82cc8f.
//
// Solidity: function depositStake(_numRTC uint256) returns(bool)
func (_Stake *StakeTransactor) DepositStake(opts *bind.TransactOpts, _numRTC *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "depositStake", _numRTC)
}

// DepositStake is a paid mutator transaction binding the contract method 0xcb82cc8f.
//
// Solidity: function depositStake(_numRTC uint256) returns(bool)
func (_Stake *StakeSession) DepositStake(_numRTC *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.DepositStake(&_Stake.TransactOpts, _numRTC)
}

// DepositStake is a paid mutator transaction binding the contract method 0xcb82cc8f.
//
// Solidity: function depositStake(_numRTC uint256) returns(bool)
func (_Stake *StakeTransactorSession) DepositStake(_numRTC *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.DepositStake(&_Stake.TransactOpts, _numRTC)
}

// DisableNewStakes is a paid mutator transaction binding the contract method 0x04b5723d.
//
// Solidity: function disableNewStakes() returns(bool)
func (_Stake *StakeTransactor) DisableNewStakes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "disableNewStakes")
}

// DisableNewStakes is a paid mutator transaction binding the contract method 0x04b5723d.
//
// Solidity: function disableNewStakes() returns(bool)
func (_Stake *StakeSession) DisableNewStakes() (*types.Transaction, error) {
	return _Stake.Contract.DisableNewStakes(&_Stake.TransactOpts)
}

// DisableNewStakes is a paid mutator transaction binding the contract method 0x04b5723d.
//
// Solidity: function disableNewStakes() returns(bool)
func (_Stake *StakeTransactorSession) DisableNewStakes() (*types.Transaction, error) {
	return _Stake.Contract.DisableNewStakes(&_Stake.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(_stakeNumber uint256) returns(bool)
func (_Stake *StakeTransactor) Mint(opts *bind.TransactOpts, _stakeNumber *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "mint", _stakeNumber)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(_stakeNumber uint256) returns(bool)
func (_Stake *StakeSession) Mint(_stakeNumber *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Mint(&_Stake.TransactOpts, _stakeNumber)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(_stakeNumber uint256) returns(bool)
func (_Stake *StakeTransactorSession) Mint(_stakeNumber *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Mint(&_Stake.TransactOpts, _stakeNumber)
}

// SetRTI is a paid mutator transaction binding the contract method 0xdeb78ab6.
//
// Solidity: function setRTI(_contract address) returns(bool)
func (_Stake *StakeTransactor) SetRTI(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setRTI", _contract)
}

// SetRTI is a paid mutator transaction binding the contract method 0xdeb78ab6.
//
// Solidity: function setRTI(_contract address) returns(bool)
func (_Stake *StakeSession) SetRTI(_contract common.Address) (*types.Transaction, error) {
	return _Stake.Contract.SetRTI(&_Stake.TransactOpts, _contract)
}

// SetRTI is a paid mutator transaction binding the contract method 0xdeb78ab6.
//
// Solidity: function setRTI(_contract address) returns(bool)
func (_Stake *StakeTransactorSession) SetRTI(_contract common.Address) (*types.Transaction, error) {
	return _Stake.Contract.SetRTI(&_Stake.TransactOpts, _contract)
}

// WithdrawInitialStake is a paid mutator transaction binding the contract method 0x29830ccc.
//
// Solidity: function withdrawInitialStake(_stakeNumber uint256) returns(bool)
func (_Stake *StakeTransactor) WithdrawInitialStake(opts *bind.TransactOpts, _stakeNumber *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "withdrawInitialStake", _stakeNumber)
}

// WithdrawInitialStake is a paid mutator transaction binding the contract method 0x29830ccc.
//
// Solidity: function withdrawInitialStake(_stakeNumber uint256) returns(bool)
func (_Stake *StakeSession) WithdrawInitialStake(_stakeNumber *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.WithdrawInitialStake(&_Stake.TransactOpts, _stakeNumber)
}

// WithdrawInitialStake is a paid mutator transaction binding the contract method 0x29830ccc.
//
// Solidity: function withdrawInitialStake(_stakeNumber uint256) returns(bool)
func (_Stake *StakeTransactorSession) WithdrawInitialStake(_stakeNumber *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.WithdrawInitialStake(&_Stake.TransactOpts, _stakeNumber)
}

// StakeInitialStakeWithdrawnIterator is returned from FilterInitialStakeWithdrawn and is used to iterate over the raw logs and unpacked data for InitialStakeWithdrawn events raised by the Stake contract.
type StakeInitialStakeWithdrawnIterator struct {
	Event *StakeInitialStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakeInitialStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeInitialStakeWithdrawn)
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
		it.Event = new(StakeInitialStakeWithdrawn)
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
func (it *StakeInitialStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeInitialStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeInitialStakeWithdrawn represents a InitialStakeWithdrawn event raised by the Stake contract.
type StakeInitialStakeWithdrawn struct {
	Staker      common.Address
	StakeNumber *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInitialStakeWithdrawn is a free log retrieval operation binding the contract event 0x7d252c33d474583922a2f7a0c2f4d04631095dbd4e35b09adc7f801ec3e743f7.
//
// Solidity: e InitialStakeWithdrawn(_staker indexed address, _stakeNumber indexed uint256, _amount uint256)
func (_Stake *StakeFilterer) FilterInitialStakeWithdrawn(opts *bind.FilterOpts, _staker []common.Address, _stakeNumber []*big.Int) (*StakeInitialStakeWithdrawnIterator, error) {

	var _stakerRule []interface{}
	for _, _stakerItem := range _staker {
		_stakerRule = append(_stakerRule, _stakerItem)
	}
	var _stakeNumberRule []interface{}
	for _, _stakeNumberItem := range _stakeNumber {
		_stakeNumberRule = append(_stakeNumberRule, _stakeNumberItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "InitialStakeWithdrawn", _stakerRule, _stakeNumberRule)
	if err != nil {
		return nil, err
	}
	return &StakeInitialStakeWithdrawnIterator{contract: _Stake.contract, event: "InitialStakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchInitialStakeWithdrawn is a free log subscription operation binding the contract event 0x7d252c33d474583922a2f7a0c2f4d04631095dbd4e35b09adc7f801ec3e743f7.
//
// Solidity: e InitialStakeWithdrawn(_staker indexed address, _stakeNumber indexed uint256, _amount uint256)
func (_Stake *StakeFilterer) WatchInitialStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *StakeInitialStakeWithdrawn, _staker []common.Address, _stakeNumber []*big.Int) (event.Subscription, error) {

	var _stakerRule []interface{}
	for _, _stakerItem := range _staker {
		_stakerRule = append(_stakerRule, _stakerItem)
	}
	var _stakeNumberRule []interface{}
	for _, _stakeNumberItem := range _stakeNumber {
		_stakeNumberRule = append(_stakeNumberRule, _stakeNumberItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "InitialStakeWithdrawn", _stakerRule, _stakeNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeInitialStakeWithdrawn)
				if err := _Stake.contract.UnpackLog(event, "InitialStakeWithdrawn", log); err != nil {
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

// StakeStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the Stake contract.
type StakeStakeDepositedIterator struct {
	Event *StakeStakeDeposited // Event containing the contract specifics and raw log

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
func (it *StakeStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeStakeDeposited)
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
		it.Event = new(StakeStakeDeposited)
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
func (it *StakeStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeStakeDeposited represents a StakeDeposited event raised by the Stake contract.
type StakeStakeDeposited struct {
	Staker       common.Address
	StakeNum     *big.Int
	CoinsToMint  *big.Int
	ReleaseDate  *big.Int
	ReleaseBlock *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x1a325385f16807e99fb688b597db78b00faee313dcf02e882dd16daab6fc3e1f.
//
// Solidity: e StakeDeposited(_staker indexed address, _stakeNum indexed uint256, _coinsToMint uint256, _releaseDate uint256, _releaseBlock uint256)
func (_Stake *StakeFilterer) FilterStakeDeposited(opts *bind.FilterOpts, _staker []common.Address, _stakeNum []*big.Int) (*StakeStakeDepositedIterator, error) {

	var _stakerRule []interface{}
	for _, _stakerItem := range _staker {
		_stakerRule = append(_stakerRule, _stakerItem)
	}
	var _stakeNumRule []interface{}
	for _, _stakeNumItem := range _stakeNum {
		_stakeNumRule = append(_stakeNumRule, _stakeNumItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "StakeDeposited", _stakerRule, _stakeNumRule)
	if err != nil {
		return nil, err
	}
	return &StakeStakeDepositedIterator{contract: _Stake.contract, event: "StakeDeposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x1a325385f16807e99fb688b597db78b00faee313dcf02e882dd16daab6fc3e1f.
//
// Solidity: e StakeDeposited(_staker indexed address, _stakeNum indexed uint256, _coinsToMint uint256, _releaseDate uint256, _releaseBlock uint256)
func (_Stake *StakeFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *StakeStakeDeposited, _staker []common.Address, _stakeNum []*big.Int) (event.Subscription, error) {

	var _stakerRule []interface{}
	for _, _stakerItem := range _staker {
		_stakerRule = append(_stakerRule, _stakerItem)
	}
	var _stakeNumRule []interface{}
	for _, _stakeNumItem := range _stakeNum {
		_stakeNumRule = append(_stakeNumRule, _stakeNumItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "StakeDeposited", _stakerRule, _stakeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeStakeDeposited)
				if err := _Stake.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
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

// StakeStakeRewardWithdrawnIterator is returned from FilterStakeRewardWithdrawn and is used to iterate over the raw logs and unpacked data for StakeRewardWithdrawn events raised by the Stake contract.
type StakeStakeRewardWithdrawnIterator struct {
	Event *StakeStakeRewardWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakeStakeRewardWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeStakeRewardWithdrawn)
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
		it.Event = new(StakeStakeRewardWithdrawn)
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
func (it *StakeStakeRewardWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeStakeRewardWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeStakeRewardWithdrawn represents a StakeRewardWithdrawn event raised by the Stake contract.
type StakeStakeRewardWithdrawn struct {
	Staker   common.Address
	StakeNum *big.Int
	Reward   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakeRewardWithdrawn is a free log retrieval operation binding the contract event 0x275541ddbc93a3fb1e5e94000231500252d2ba460de93bd1cf285e68563c1a64.
//
// Solidity: e StakeRewardWithdrawn(_staker indexed address, _stakeNum indexed uint256, _reward uint256)
func (_Stake *StakeFilterer) FilterStakeRewardWithdrawn(opts *bind.FilterOpts, _staker []common.Address, _stakeNum []*big.Int) (*StakeStakeRewardWithdrawnIterator, error) {

	var _stakerRule []interface{}
	for _, _stakerItem := range _staker {
		_stakerRule = append(_stakerRule, _stakerItem)
	}
	var _stakeNumRule []interface{}
	for _, _stakeNumItem := range _stakeNum {
		_stakeNumRule = append(_stakeNumRule, _stakeNumItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "StakeRewardWithdrawn", _stakerRule, _stakeNumRule)
	if err != nil {
		return nil, err
	}
	return &StakeStakeRewardWithdrawnIterator{contract: _Stake.contract, event: "StakeRewardWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeRewardWithdrawn is a free log subscription operation binding the contract event 0x275541ddbc93a3fb1e5e94000231500252d2ba460de93bd1cf285e68563c1a64.
//
// Solidity: e StakeRewardWithdrawn(_staker indexed address, _stakeNum indexed uint256, _reward uint256)
func (_Stake *StakeFilterer) WatchStakeRewardWithdrawn(opts *bind.WatchOpts, sink chan<- *StakeStakeRewardWithdrawn, _staker []common.Address, _stakeNum []*big.Int) (event.Subscription, error) {

	var _stakerRule []interface{}
	for _, _stakerItem := range _staker {
		_stakerRule = append(_stakerRule, _stakerItem)
	}
	var _stakeNumRule []interface{}
	for _, _stakeNumItem := range _stakeNum {
		_stakeNumRule = append(_stakeNumRule, _stakeNumItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "StakeRewardWithdrawn", _stakerRule, _stakeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeStakeRewardWithdrawn)
				if err := _Stake.contract.UnpackLog(event, "StakeRewardWithdrawn", log); err != nil {
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

// StakeStakesDisabledIterator is returned from FilterStakesDisabled and is used to iterate over the raw logs and unpacked data for StakesDisabled events raised by the Stake contract.
type StakeStakesDisabledIterator struct {
	Event *StakeStakesDisabled // Event containing the contract specifics and raw log

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
func (it *StakeStakesDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeStakesDisabled)
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
		it.Event = new(StakeStakesDisabled)
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
func (it *StakeStakesDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeStakesDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeStakesDisabled represents a StakesDisabled event raised by the Stake contract.
type StakeStakesDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakesDisabled is a free log retrieval operation binding the contract event 0xfd71816274d74ca17e4eb9ecd097893ef1b3e7549be66c6de9fc95bc55f53324.
//
// Solidity: e StakesDisabled()
func (_Stake *StakeFilterer) FilterStakesDisabled(opts *bind.FilterOpts) (*StakeStakesDisabledIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "StakesDisabled")
	if err != nil {
		return nil, err
	}
	return &StakeStakesDisabledIterator{contract: _Stake.contract, event: "StakesDisabled", logs: logs, sub: sub}, nil
}

// WatchStakesDisabled is a free log subscription operation binding the contract event 0xfd71816274d74ca17e4eb9ecd097893ef1b3e7549be66c6de9fc95bc55f53324.
//
// Solidity: e StakesDisabled()
func (_Stake *StakeFilterer) WatchStakesDisabled(opts *bind.WatchOpts, sink chan<- *StakeStakesDisabled) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "StakesDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeStakesDisabled)
				if err := _Stake.contract.UnpackLog(event, "StakesDisabled", log); err != nil {
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

// StakeStakesEnabledIterator is returned from FilterStakesEnabled and is used to iterate over the raw logs and unpacked data for StakesEnabled events raised by the Stake contract.
type StakeStakesEnabledIterator struct {
	Event *StakeStakesEnabled // Event containing the contract specifics and raw log

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
func (it *StakeStakesEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeStakesEnabled)
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
		it.Event = new(StakeStakesEnabled)
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
func (it *StakeStakesEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeStakesEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeStakesEnabled represents a StakesEnabled event raised by the Stake contract.
type StakeStakesEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakesEnabled is a free log retrieval operation binding the contract event 0xf83de010d0db8e58cb57c6568bd71fe9dcabbad734fd8d37128edffee0820e11.
//
// Solidity: e StakesEnabled()
func (_Stake *StakeFilterer) FilterStakesEnabled(opts *bind.FilterOpts) (*StakeStakesEnabledIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "StakesEnabled")
	if err != nil {
		return nil, err
	}
	return &StakeStakesEnabledIterator{contract: _Stake.contract, event: "StakesEnabled", logs: logs, sub: sub}, nil
}

// WatchStakesEnabled is a free log subscription operation binding the contract event 0xf83de010d0db8e58cb57c6568bd71fe9dcabbad734fd8d37128edffee0820e11.
//
// Solidity: e StakesEnabled()
func (_Stake *StakeFilterer) WatchStakesEnabled(opts *bind.WatchOpts, sink chan<- *StakeStakesEnabled) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "StakesEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeStakesEnabled)
				if err := _Stake.contract.UnpackLog(event, "StakesEnabled", log); err != nil {
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
