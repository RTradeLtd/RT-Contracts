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
const StakeABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"disableNewStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MULTIPLIER\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINSTAKE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"allowNewStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stakeNumber\",\"type\":\"uint256\"}],\"name\":\"withdrawInitialStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakes\",\"outputs\":[{\"name\":\"initialStake\",\"type\":\"uint256\"},{\"name\":\"blockLocked\",\"type\":\"uint256\"},{\"name\":\"blockUnlocked\",\"type\":\"uint256\"},{\"name\":\"releaseDate\",\"type\":\"uint256\"},{\"name\":\"totalCoinsToMint\",\"type\":\"uint256\"},{\"name\":\"coinsMinted\",\"type\":\"uint256\"},{\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"name\":\"lastBlockWithdrawn\",\"type\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RTI\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newStakesAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BLOCKHOLDPERIOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stakeNumber\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BLOCKSEC\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKENCONTRACT\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canMint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_numRTC\",\"type\":\"uint256\"}],\"name\":\"depositStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setRTI\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"numberOfStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"activeStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"internalRTCBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakesDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakesEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_stakeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_coinsToMint\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_releaseDate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_releaseBlock\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_stakeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"StakeRewardWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_stakeNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"InitialStakeWithdrawn\",\"type\":\"event\"}]"

// StakeBin is the compiled bytecode used for deploying new contracts.
const StakeBin = `608060405260008054600160a060020a0319908116733fde03720917246b73ba532e4650c656d6020578179182905560018054909116600160a060020a039290921691909117905534801561005357600080fd5b50600054600160a060020a0316151561006b57600080fd5b60038054600160a060020a0319163317905561105f8061008c6000396000f3006080604052600436106101065763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166304b5723d811461010b578063059f8b16146101345780631d0a0e8d1461015b578063211db50d1461017057806329830ccc14610185578063584b62a11461019d57806358e1c1741461021d578063594548d51461024e57806377be337614610263578063a0712d6814610278578063b63e54c314610290578063b792967a146102a5578063beb9716d146102ba578063cb82cc8f146102cf578063deb78ab6146102e7578063dfef667914610308578063ed2f236914610329578063f16108211461033e578063f851a4401461035f575b600080fd5b34801561011757600080fd5b50610120610374565b604080519115158252519081900360200190f35b34801561014057600080fd5b506101496103b2565b60408051918252519081900360200190f35b34801561016757600080fd5b506101496103bd565b34801561017c57600080fd5b506101206103c9565b34801561019157600080fd5b506101206004356104bc565b3480156101a957600080fd5b506101c1600160a060020a03600435166024356106d5565b604051808a815260200189815260200188815260200187815260200186815260200185815260200184815260200183815260200182600281111561020157fe5b60ff168152602001995050505050505050505060405180910390f35b34801561022957600080fd5b5061023261072b565b60408051600160a060020a039092168252519081900360200190f35b34801561025a57600080fd5b5061012061073a565b34801561026f57600080fd5b5061014961075b565b34801561028457600080fd5b50610120600435610760565b34801561029c57600080fd5b506101496109b6565b3480156102b157600080fd5b506102326109bb565b3480156102c657600080fd5b506101206109ca565b3480156102db57600080fd5b50610120600435610a7e565b3480156102f357600080fd5b50610120600160a060020a0360043516610d76565b34801561031457600080fd5b50610149600160a060020a0360043516610dcc565b34801561033557600080fd5b50610149610dde565b34801561034a57600080fd5b50610149600160a060020a0360043516610de4565b34801561036b57600080fd5b50610232610df6565b600354600090600160a060020a0316331461038e57600080fd5b506003805474ff000000000000000000000000000000000000000019169055600190565b662386f26fc1000081565b670de0b6b3a764000081565b600354600090600160a060020a031633146103e357600080fd5b6003805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055600154604080517f1a18622700000000000000000000000000000000000000000000000000000000815290513092600160a060020a031691631a1862279160048083019260209291908290030181600087803b15801561047757600080fd5b505af115801561048b573d6000803e3d6000fd5b505050506040513d60208110156104a157600080fd5b5051600160a060020a0316146104b657600080fd5b50600190565b60008082600133600090815260046020908152604080832085845290915290206008015460ff1660028111156104ee57fe5b146104f857600080fd5b336000908152600460209081526040808320848452909152902060030154421080159061054257503360009081526004602090815260408083208484529091529020600201544310155b151561054d57600080fd5b3360008181526004602090815260408083208584528252808320549383526006909152902054101561057e57600080fd5b336000908152600460209081526040808320878452909152902080546008909101805460ff19166002908117909155549092506105bc906001610e05565b600255336000908152600660205260409020546105df908363ffffffff610e0516565b33600081815260066020908152604091829020939093558051858152905187937f7d252c33d474583922a2f7a0c2f4d04631095dbd4e35b09adc7f801ec3e743f7928290030190a3600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018590529051600160a060020a039092169163a9059cbb916044808201926020929091908290030181600087803b15801561069457600080fd5b505af11580156106a8573d6000803e3d6000fd5b505050506040513d60208110156106be57600080fd5b505115156106cb57600080fd5b5060019392505050565b600460208181526000938452604080852090915291835291208054600182015460028301546003840154948401546005850154600686015460078701546008909701549597949693959293919290919060ff1689565b600154600160a060020a031681565b60035474010000000000000000000000000000000000000000900460ff1681565b600581565b600080828180600133600090815260046020908152604080832087845290915290206008015460ff16600281111561079457fe5b14806107cb5750600233600090815260046020908152604080832087845290915290206008015460ff1660028111156107c957fe5b145b15156107d657600080fd5b336000908152600460208181526040808420878552909152909120908101546005909101541061080557600080fd5b5050336000908152600460209081526040808320848452909152902060070154439080821161083357600080fd5b61083c86610e1a565b3360009081526004602081815260408084208b8552909152909120908101546005909101549195509061086f9086610e97565b111561087a57600080fd5b3360009081526004602090815260408083208984529091529020600501546108a8908563ffffffff610e9716565b3360008181526004602090815260408083208b84528252918290206005810194909455436007909401939093558051878152905189937f275541ddbc93a3fb1e5e94000231500252d2ba460de93bd1cf285e68563c1a64928290030190a3600154604080517f40c10f19000000000000000000000000000000000000000000000000000000008152336004820152602481018790529051600160a060020a03909216916340c10f19916044808201926020929091908290030181600087803b15801561097357600080fd5b505af1158015610987573d6000803e3d6000fd5b505050506040513d602081101561099d57600080fd5b505115156109aa57600080fd5b50600195945050505050565b600f81565b600054600160a060020a031681565b600030600160a060020a0316600160009054906101000a9004600160a060020a0316600160a060020a0316631a1862276040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610a4257600080fd5b505af1158015610a56573d6000803e3d6000fd5b505050506040513d6020811015610a6c57600080fd5b5051600160a060020a0316146104b657fe5b6000806000806000806000610a91610fdc565b88610a9a6109ca565b1515610aa557600080fd5b60035474010000000000000000000000000000000000000000900460ff161515610ace57600080fd5b670de0b6b3a7640000811015610ae357600080fd5b610aec33610eb0565b9750610af78a610ecb565b96509650965096509650610120604051908101604052808b81526020018881526020018781526020018681526020018581526020016000815260200184815260200188815260200160016002811115610b4c57fe5b8152509150816004600033600160a060020a0316600160a060020a0316815260200190815260200160002060008a8152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701556101008201518160080160006101000a81548160ff02191690836002811115610bfd57fe5b02179055505033600090815260056020526040902054610c259150600163ffffffff610e9716565b33600090815260056020908152604080832093909355600690522054610c51908b63ffffffff610e9716565b33600090815260066020526040902055600254610c7590600163ffffffff610e9716565b60025560408051858152602081018790528082018890529051899133917f1a325385f16807e99fb688b597db78b00faee313dcf02e882dd16daab6fc3e1f9181900360600190a3600154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018d90529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b158015610d2f57600080fd5b505af1158015610d43573d6000803e3d6000fd5b505050506040513d6020811015610d5957600080fd5b50511515610d6657600080fd5b5060019998505050505050505050565b600354600090600160a060020a03163314610d9057600080fd5b5060018054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1991821681178355600080549092161790555b919050565b60056020526000908152604090205481565b60025481565b60066020526000908152604090205481565b600354600160a060020a031681565b600082821115610e1457600080fd5b50900390565b600080600080610e2985610f57565b3360009081526004602090815260408083208984529091529020600701549093509150610e5c838363ffffffff610e0516565b336000908152600460209081526040808320898452909152902060060154909150610e8e90829063ffffffff610f9e16565b95945050505050565b600082820183811015610ea957600080fd5b9392505050565b600160a060020a031660009081526005602052604090205490565b436000808080610ee285600563ffffffff610e9716565b9350610f06610ef96005600f63ffffffff610f9e16565b429063ffffffff610e9716565b9250610f1f86662386f26fc1000063ffffffff610f9e16565b9150610f3982670de0b6b3a764000063ffffffff610fc516565b9150610f4c82600563ffffffff610fc516565b905091939590929450565b33600090815260046020908152604080832084845290915290206002015443908110610dc75750336000908152600460209081526040808320938352929052206002015490565b6000828202831580610fba5750828482811515610fb757fe5b04145b1515610ea957600080fd5b6000808284811515610fd357fe5b04949350505050565b6101206040519081016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000600281111561102e57fe5b9052905600a165627a7a723058208c45ae39176c6a22a3aced068a53f806f2b8e6e7388ff53a86555d51a425d5030029`

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

// BLOCKHOLDPERIOD is a free data retrieval call binding the contract method 0x77be3376.
//
// Solidity: function BLOCKHOLDPERIOD() constant returns(uint256)
func (_Stake *StakeCaller) BLOCKHOLDPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "BLOCKHOLDPERIOD")
	return *ret0, err
}

// BLOCKHOLDPERIOD is a free data retrieval call binding the contract method 0x77be3376.
//
// Solidity: function BLOCKHOLDPERIOD() constant returns(uint256)
func (_Stake *StakeSession) BLOCKHOLDPERIOD() (*big.Int, error) {
	return _Stake.Contract.BLOCKHOLDPERIOD(&_Stake.CallOpts)
}

// BLOCKHOLDPERIOD is a free data retrieval call binding the contract method 0x77be3376.
//
// Solidity: function BLOCKHOLDPERIOD() constant returns(uint256)
func (_Stake *StakeCallerSession) BLOCKHOLDPERIOD() (*big.Int, error) {
	return _Stake.Contract.BLOCKHOLDPERIOD(&_Stake.CallOpts)
}

// BLOCKSEC is a free data retrieval call binding the contract method 0xb63e54c3.
//
// Solidity: function BLOCKSEC() constant returns(uint256)
func (_Stake *StakeCaller) BLOCKSEC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "BLOCKSEC")
	return *ret0, err
}

// BLOCKSEC is a free data retrieval call binding the contract method 0xb63e54c3.
//
// Solidity: function BLOCKSEC() constant returns(uint256)
func (_Stake *StakeSession) BLOCKSEC() (*big.Int, error) {
	return _Stake.Contract.BLOCKSEC(&_Stake.CallOpts)
}

// BLOCKSEC is a free data retrieval call binding the contract method 0xb63e54c3.
//
// Solidity: function BLOCKSEC() constant returns(uint256)
func (_Stake *StakeCallerSession) BLOCKSEC() (*big.Int, error) {
	return _Stake.Contract.BLOCKSEC(&_Stake.CallOpts)
}

// MINSTAKE is a free data retrieval call binding the contract method 0x1d0a0e8d.
//
// Solidity: function MINSTAKE() constant returns(uint256)
func (_Stake *StakeCaller) MINSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "MINSTAKE")
	return *ret0, err
}

// MINSTAKE is a free data retrieval call binding the contract method 0x1d0a0e8d.
//
// Solidity: function MINSTAKE() constant returns(uint256)
func (_Stake *StakeSession) MINSTAKE() (*big.Int, error) {
	return _Stake.Contract.MINSTAKE(&_Stake.CallOpts)
}

// MINSTAKE is a free data retrieval call binding the contract method 0x1d0a0e8d.
//
// Solidity: function MINSTAKE() constant returns(uint256)
func (_Stake *StakeCallerSession) MINSTAKE() (*big.Int, error) {
	return _Stake.Contract.MINSTAKE(&_Stake.CallOpts)
}

// MULTIPLIER is a free data retrieval call binding the contract method 0x059f8b16.
//
// Solidity: function MULTIPLIER() constant returns(uint256)
func (_Stake *StakeCaller) MULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "MULTIPLIER")
	return *ret0, err
}

// MULTIPLIER is a free data retrieval call binding the contract method 0x059f8b16.
//
// Solidity: function MULTIPLIER() constant returns(uint256)
func (_Stake *StakeSession) MULTIPLIER() (*big.Int, error) {
	return _Stake.Contract.MULTIPLIER(&_Stake.CallOpts)
}

// MULTIPLIER is a free data retrieval call binding the contract method 0x059f8b16.
//
// Solidity: function MULTIPLIER() constant returns(uint256)
func (_Stake *StakeCallerSession) MULTIPLIER() (*big.Int, error) {
	return _Stake.Contract.MULTIPLIER(&_Stake.CallOpts)
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

// TOKENCONTRACT is a free data retrieval call binding the contract method 0xb792967a.
//
// Solidity: function TOKENCONTRACT() constant returns(address)
func (_Stake *StakeCaller) TOKENCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "TOKENCONTRACT")
	return *ret0, err
}

// TOKENCONTRACT is a free data retrieval call binding the contract method 0xb792967a.
//
// Solidity: function TOKENCONTRACT() constant returns(address)
func (_Stake *StakeSession) TOKENCONTRACT() (common.Address, error) {
	return _Stake.Contract.TOKENCONTRACT(&_Stake.CallOpts)
}

// TOKENCONTRACT is a free data retrieval call binding the contract method 0xb792967a.
//
// Solidity: function TOKENCONTRACT() constant returns(address)
func (_Stake *StakeCallerSession) TOKENCONTRACT() (common.Address, error) {
	return _Stake.Contract.TOKENCONTRACT(&_Stake.CallOpts)
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
