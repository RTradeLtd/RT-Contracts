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

// StakeABI is the input ABI used to generate the binding from.
const StakeABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"disableNewStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"allowNewStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stakeNumber\",\"type\":\"uint256\"}],\"name\":\"withdrawInitialStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKENADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakes\",\"outputs\":[{\"name\":\"initialStake\",\"type\":\"uint256\"},{\"name\":\"blockLocked\",\"type\":\"uint256\"},{\"name\":\"blockUnlocked\",\"type\":\"uint256\"},{\"name\":\"releaseDate\",\"type\":\"uint256\"},{\"name\":\"totalCoinsToMint\",\"type\":\"uint256\"},{\"name\":\"coinsMinted\",\"type\":\"uint256\"},{\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"name\":\"lastBlockWithdrawn\",\"type\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RTI\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newStakesAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferForeignToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stakeNumber\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canMint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_numRTC\",\"type\":\"uint256\"}],\"name\":\"depositStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"numberOfStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"activeStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"internalRTCBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakesDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakesEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_stakeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_coinsToMint\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_releaseDate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_releaseBlock\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_stakeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"StakeRewardWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_stakeNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"InitialStakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ForeignTokenTransfer\",\"type\":\"event\"}]"

// StakeBin is the compiled bytecode used for deploying new contracts.
const StakeBin = `608060405234801561001057600080fd5b506040516020806118cd833981016040525160018054600160a060020a031916600160a060020a039290921691909117905561187c806100516000396000f3006080604052600436106100e55763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166304b5723d81146100ea578063211db50d1461011357806329830ccc14610128578063516f898614610140578063584b62a11461017157806358e1c17414610140578063594548d5146101f15780639e5fea8a14610206578063a0712d6814610230578063beb9716d14610248578063cb82cc8f1461025d578063dfef667914610275578063ed2f2369146102a8578063f1610821146102bd578063f851a440146102de578063ffa1ad74146102f3575b600080fd5b3480156100f657600080fd5b506100ff61037d565b604080519115158252519081900360200190f35b34801561011f57600080fd5b506100ff610404565b34801561013457600080fd5b506100ff6004356105e1565b34801561014c57600080fd5b5061015561099b565b60408051600160a060020a039092168252519081900360200190f35b34801561017d57600080fd5b50610195600160a060020a03600435166024356109b3565b604051808a81526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020018260028111156101d557fe5b60ff168152602001995050505050505050505060405180910390f35b3480156101fd57600080fd5b506100ff610a09565b34801561021257600080fd5b506100ff600160a060020a0360043581169060243516604435610a2a565b34801561023c57600080fd5b506100ff600435610ca0565b34801561025457600080fd5b506100ff61105d565b34801561026957600080fd5b506100ff6004356110d2565b34801561028157600080fd5b50610296600160a060020a036004351661154e565b60408051918252519081900360200190f35b3480156102b457600080fd5b50610296611560565b3480156102c957600080fd5b50610296600160a060020a0360043516611566565b3480156102ea57600080fd5b50610155611578565b3480156102ff57600080fd5b50610308611587565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561034257818101518382015260200161032a565b50505050905090810190601f16801561036f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600154600090600160a060020a031633146103e2576040805160e560020a62461bcd02815260206004820152601360248201527f73656e646572206973206e6f742061646d696e00000000000000000000000000604482015290519081900360640190fd5b506001805474ff00000000000000000000000000000000000000001916815590565b600154600090600160a060020a03163314610469576040805160e560020a62461bcd02815260206004820152601360248201527f73656e646572206973206e6f742061646d696e00000000000000000000000000604482015290519081900360640190fd5b6001805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055604080517f272caf690000000000000000000000000000000000000000000000000000000081529051309173ecc043b92834c1ebde65f2181b59597a6588d6169163272caf6991600480820192602092909190829003018186803b15801561050557600080fd5b505afa158015610519573d6000803e3d6000fd5b505050506040513d602081101561052f57600080fd5b5051600160a060020a0316146105db576040805160e560020a62461bcd02815260206004820152604a60248201527f72746320746f6b656e20636f6e7472616374206973206e6f742073657420746f60448201527f20757365207468697320636f6e747261637420617320746865207374616b696e60648201527f6720636f6e747261637400000000000000000000000000000000000000000000608482015290519081900360a40190fd5b50600190565b6000808260013360009081526002602081815260408084208685529091529091206008015460ff169081111561061357fe5b14610668576040805160e560020a62461bcd02815260206004820152601360248201527f7374616b65206973206e6f742061637469766500000000000000000000000000604482015290519081900360640190fd5b33600090815260026020908152604080832084845290915290206003015442108015906106b1575033600090815260026020818152604080842085855290915290912001544310155b1515610753576040805160e560020a62461bcd02815260206004820152604160248201527f617474656d7074696e6720746f20776974686472617720696e697469616c207360448201527f74616b65206265666f726520756e6c6f636b20626c6f636b20616e642064617460648201527f6500000000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b336000818152600260209081526040808320858452825280832054938352600490915290205410156107cf576040805160e560020a62461bcd02815260206004820152601c60248201527f696e76616c696420696e7465726e616c207274632062616c616e636500000000604482015290519081900360640190fd5b336000908152600260208181526040808420888552909152822080546008909101805460ff1916909217909155905490925061080c9060016115be565b60009081553381526004602052604090205461082e908363ffffffff6115be16565b33600081815260046020908152604091829020939093558051858152905187937f7d252c33d474583922a2f7a0c2f4d04631095dbd4e35b09adc7f801ec3e743f7928290030190a3604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815233600482015260248101849052905173ecc043b92834c1ebde65f2181b59597a6588d6169163a9059cbb9160448083019260209291908290030181600087803b1580156108e957600080fd5b505af11580156108fd573d6000803e3d6000fd5b505050506040513d602081101561091357600080fd5b50511515610991576040805160e560020a62461bcd02815260206004820152603960248201527f756e61626c6520746f207472616e7366657220746f6b656e73206c696b656c7960448201527f2064756520746f20696e636f72726563742062616c616e636500000000000000606482015290519081900360840190fd5b5060019392505050565b73ecc043b92834c1ebde65f2181b59597a6588d61681565b600260208181526000938452604080852090915291835291208054600182015492820154600383015460048401546005850154600686015460078701546008909701549597969495939492939192909160ff1689565b60015474010000000000000000000000000000000000000000900460ff1681565b6001546000908190600160a060020a03163314610a91576040805160e560020a62461bcd02815260206004820152601360248201527f73656e646572206973206e6f742061646d696e00000000000000000000000000604482015290519081900360640190fd5b600160a060020a0384161515610af1576040805160e560020a62461bcd02815260206004820181905260248201527f726563697069656e7420616464726573732063616e277420626520656d707479604482015290519081900360640190fd5b600160a060020a03851673ecc043b92834c1ebde65f2181b59597a6588d6161415610b66576040805160e560020a62461bcd02815260206004820152601260248201527f746f6b656e2063616e2774206265205254430000000000000000000000000000604482015290519081900360640190fd5b50604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038581166004830152602482018590529151869283169163a9059cbb9160448083019260209291908290030181600087803b158015610bd357600080fd5b505af1158015610be7573d6000803e3d6000fd5b505050506040513d6020811015610bfd57600080fd5b50511515610c55576040805160e560020a62461bcd02815260206004820152601560248201527f746f6b656e207472616e73666572206661696c65640000000000000000000000604482015290519081900360640190fd5b604080518481529051600160a060020a0386169133917f10a46ed575affad8e954ae27853b1f89c6da90d8c35f619fc640f8a21bcb78579181900360200190a3506001949350505050565b60008082818060013360009081526002602081815260408084208885529091529091206008015460ff1690811115610cd457fe5b1480610d0957503360009081526002602081815260408084208785529091529091206008015460ff1681811115610d0757fe5b145b1515610d85576040805160e560020a62461bcd02815260206004820152603860248201527f7374616b65206d75737420626520616374697665206f7220696e61637469766560448201527f20696e206f7264657220746f206d696e7420746f6b656e730000000000000000606482015290519081900360840190fd5b3360009081526002602090815260408083208684529091529020600481015460059091015410610e25576040805160e560020a62461bcd02815260206004820152602c60248201527f63757272656e7420636f696e73206d696e746564206d757374206265206c657360448201527f73207468616e20746f74616c0000000000000000000000000000000000000000606482015290519081900360840190fd5b50503360009081526002602090815260408083208484529091529020600701544390808211610ec4576040805160e560020a62461bcd02815260206004820152603560248201527f63757272656e7420626c6f636b206d757374206265206f6e652068696768657260448201527f207468616e206c617374207769746864726177616c0000000000000000000000606482015290519081900360840190fd5b610ecd866115d3565b3360009081526002602090815260408083208a8452909152902060050154909450610efe908563ffffffff6116aa16565b3360008181526002602090815260408083208b84528252918290206005810194909455436007909401939093558051878152905189937f275541ddbc93a3fb1e5e94000231500252d2ba460de93bd1cf285e68563c1a64928290030190a3604080517f40c10f1900000000000000000000000000000000000000000000000000000000815233600482015260248101869052905173ecc043b92834c1ebde65f2181b59597a6588d616916340c10f199160448083019260209291908290030181600087803b158015610fcf57600080fd5b505af1158015610fe3573d6000803e3d6000fd5b505050506040513d6020811015610ff957600080fd5b50511515611051576040805160e560020a62461bcd02815260206004820152601460248201527f746f6b656e206d696e74696e67206661696c6564000000000000000000000000604482015290519081900360640190fd5b50600195945050505050565b600030600160a060020a031673ecc043b92834c1ebde65f2181b59597a6588d616600160a060020a031663272caf696040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561050557600080fd5b60008060008060008060006110e56117f9565b886110ee61105d565b151561116a576040805160e560020a62461bcd02815260206004820152602960248201527f7374616b696e6720636f6e747261637420697320756e61626c6520746f206d6960448201527f6e7420746f6b656e730000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60015474010000000000000000000000000000000000000000900460ff1615156111de576040805160e560020a62461bcd02815260206004820152601a60248201527f6e6577207374616b657320617265206e6f7420616c6c6f776564000000000000604482015290519081900360640190fd5b670de0b6b3a7640000811015611264576040805160e560020a62461bcd02815260206004820152602c60248201527f737065636966696564207374616b65206973206c6f776572207468616e206d6960448201527f6e696d756d20616d6f756e740000000000000000000000000000000000000000606482015290519081900360840190fd5b61126d336116c3565b97506112788a6116e2565b96509650965096509650610120604051908101604052808b815260200188815260200187815260200186815260200185815260200160008152602001848152602001888152602001600160028111156112cd57fe5b90523360009081526002602081815260408084208d855282529283902084518155908401516001808301919091559284015181830155606084015160038201556080840151600482015560a0840151600582015560c0840151600682015560e08401516007820155610100840151600882018054959750879592949193909260ff19169190849081111561135d57fe5b021790555050336000908152600360205260409020546113859150600163ffffffff6116aa16565b336000908152600360209081526040808320939093556004905220546113b1908b63ffffffff6116aa16565b33600090815260046020526040812091909155546113d690600163ffffffff6116aa16565b60005560408051858152602081018790528082018890529051899133917f1a325385f16807e99fb688b597db78b00faee313dcf02e882dd16daab6fc3e1f9181900360600190a3604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018c9052905173ecc043b92834c1ebde65f2181b59597a6588d616916323b872dd9160648083019260209291908290030181600087803b15801561149657600080fd5b505af11580156114aa573d6000803e3d6000fd5b505050506040513d60208110156114c057600080fd5b5051151561153e576040805160e560020a62461bcd02815260206004820152602b60248201527f7472616e736665722066726f6d206661696c65642c206c696b656c79206e656560448201527f647320617070726f76616c000000000000000000000000000000000000000000606482015290519081900360840190fd5b5060019998505050505050505050565b60036020526000908152604090205481565b60005481565b60046020526000908152604090205481565b600154600160a060020a031681565b60408051808201909152600a81527f70726f64756374696f6e00000000000000000000000000000000000000000000602082015281565b6000828211156115cd57600080fd5b50900390565b60008060008060008060006115e788611775565b3360009081526002602090815260408083208c8452909152902060070154909650945061161a868663ffffffff6115be16565b3360009081526002602090815260408083208c845290915290206006015490945061164c90859063ffffffff6117bb16565b3360009081526002602090815260408083208c8452909152902060048101546005909101549198509350915061168282886116aa565b90508281111561169f5761169c818463ffffffff6115be16565b96505b505050505050919050565b6000828201838110156116bc57600080fd5b9392505050565b600160a060020a0381166000908152600360205260409020545b919050565b4360008080806116fb8562201a2063ffffffff6116aa16565b935061172161171462201a20600f63ffffffff6117bb16565b429063ffffffff6116aa16565b925061173b8667016345785d8a000063ffffffff6117bb16565b915061175582670de0b6b3a764000063ffffffff6117e216565b915061176a8262201a2063ffffffff6117e216565b905091939590929450565b3360009081526002602081815260408084208585529091529091200154439081106116dd5750336000908152600260208181526040808420948452939052919020015490565b60008282028315806117d757508284828115156117d457fe5b04145b15156116bc57600080fd5b60008082848115156117f057fe5b04949350505050565b6101206040519081016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000600281111561184b57fe5b9052905600a165627a7a72305820fc5228638a7dc266e5b363c2b71e378399a8cd50ad1fe69209b2d27d7cdb201d0029`

// DeployStake deploys a new Ethereum contract, binding an instance of Stake to it.
func DeployStake(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address) (common.Address, *types.Transaction, *Stake, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakeBin), backend, _admin)
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

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Stake *StakeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Stake.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Stake *StakeSession) VERSION() (string, error) {
	return _Stake.Contract.VERSION(&_Stake.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Stake *StakeCallerSession) VERSION() (string, error) {
	return _Stake.Contract.VERSION(&_Stake.CallOpts)
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

// TransferForeignToken is a paid mutator transaction binding the contract method 0x9e5fea8a.
//
// Solidity: function transferForeignToken(_tokenAddress address, _recipient address, _amount uint256) returns(bool)
func (_Stake *StakeTransactor) TransferForeignToken(opts *bind.TransactOpts, _tokenAddress common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "transferForeignToken", _tokenAddress, _recipient, _amount)
}

// TransferForeignToken is a paid mutator transaction binding the contract method 0x9e5fea8a.
//
// Solidity: function transferForeignToken(_tokenAddress address, _recipient address, _amount uint256) returns(bool)
func (_Stake *StakeSession) TransferForeignToken(_tokenAddress common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.TransferForeignToken(&_Stake.TransactOpts, _tokenAddress, _recipient, _amount)
}

// TransferForeignToken is a paid mutator transaction binding the contract method 0x9e5fea8a.
//
// Solidity: function transferForeignToken(_tokenAddress address, _recipient address, _amount uint256) returns(bool)
func (_Stake *StakeTransactorSession) TransferForeignToken(_tokenAddress common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.TransferForeignToken(&_Stake.TransactOpts, _tokenAddress, _recipient, _amount)
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

// StakeForeignTokenTransferIterator is returned from FilterForeignTokenTransfer and is used to iterate over the raw logs and unpacked data for ForeignTokenTransfer events raised by the Stake contract.
type StakeForeignTokenTransferIterator struct {
	Event *StakeForeignTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StakeForeignTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeForeignTokenTransfer)
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
		it.Event = new(StakeForeignTokenTransfer)
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
func (it *StakeForeignTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeForeignTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeForeignTokenTransfer represents a ForeignTokenTransfer event raised by the Stake contract.
type StakeForeignTokenTransfer struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterForeignTokenTransfer is a free log retrieval operation binding the contract event 0x10a46ed575affad8e954ae27853b1f89c6da90d8c35f619fc640f8a21bcb7857.
//
// Solidity: e ForeignTokenTransfer(_sender indexed address, _recipient indexed address, _amount uint256)
func (_Stake *StakeFilterer) FilterForeignTokenTransfer(opts *bind.FilterOpts, _sender []common.Address, _recipient []common.Address) (*StakeForeignTokenTransferIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "ForeignTokenTransfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return &StakeForeignTokenTransferIterator{contract: _Stake.contract, event: "ForeignTokenTransfer", logs: logs, sub: sub}, nil
}

// WatchForeignTokenTransfer is a free log subscription operation binding the contract event 0x10a46ed575affad8e954ae27853b1f89c6da90d8c35f619fc640f8a21bcb7857.
//
// Solidity: e ForeignTokenTransfer(_sender indexed address, _recipient indexed address, _amount uint256)
func (_Stake *StakeFilterer) WatchForeignTokenTransfer(opts *bind.WatchOpts, sink chan<- *StakeForeignTokenTransfer, _sender []common.Address, _recipient []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "ForeignTokenTransfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeForeignTokenTransfer)
				if err := _Stake.contract.UnpackLog(event, "ForeignTokenTransfer", log); err != nil {
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
