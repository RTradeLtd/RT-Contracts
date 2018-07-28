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

// RTCoinABI is the input ABI used to generate the binding from.
const RTCoinABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"freezeTransfers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"startOwnerTransferDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stake\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setStakeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delay\",\"outputs\":[{\"name\":\"previousOwner\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"},{\"name\":\"delayExpirationBlock\",\"type\":\"uint256\"},{\"name\":\"delayExpirationTime\",\"type\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"transferOutEth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferForeignToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_mergedMinerValidator\",\"type\":\"address\"}],\"name\":\"setMergedMinerValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"thawTransfers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setFailOverStakeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIALSUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeFailOverRestrictionLifted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transfersFrozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mergedMinerValidatorAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"minters\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_transfersFrozen\",\"type\":\"bool\"}],\"name\":\"TransfersFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_transfersThawed\",\"type\":\"bool\"}],\"name\":\"TransfersThawed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ForeignTokenTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"EthTransferOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"MergedMinerValidatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"StakeContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"FailOverStakeContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_stakeContract\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_mintAmount\",\"type\":\"uint256\"}],\"name\":\"CoinsMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_activationBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_activationTime\",\"type\":\"uint256\"}],\"name\":\"OwnerTransferDelayStarted\",\"type\":\"event\"}]"

// RTCoinBin is the compiled bytecode used for deploying new contracts.
const RTCoinBin = `60078054600160a060020a0319908116909155600880548216905560098054909116905560c0604052600660808190527f5254436f696e000000000000000000000000000000000000000000000000000060a09081526200006491600a919062000159565b506040805180820190915260038082527f52544300000000000000000000000000000000000000000000000000000000006020909201918252620000ab91600b9162000159565b506a32f44eb0f61c6124000000600c55600d805462ff00001961ff001960ff199092166012179190911661010017169055348015620000e957600080fd5b506000805433600160a060020a031991821681178355600180549092168117909155600c54818352600e602090815260408085208390558051928352519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a3620001fe565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200019c57805160ff1916838001178555620001cc565b82800160010185558215620001cc579182015b82811115620001cc578251825591602001919060010190620001af565b50620001da929150620001de565b5090565b620001fb91905b80821115620001da5760008155600101620001e5565b90565b6119ed806200020e6000396000f30060806040526004361061017c5763ffffffff60e060020a60003504166301502460811461018157806306fdde03146101aa578063095ea7b3146102345780630ff69b521461025857806318160ddd1461027957806323b872dd146102a0578063272caf69146102ca57806327e235e3146102fb578063313ce5671461031c5780633a4b66f11461034757806340c10f191461035c578063509484d5146103805780635c658165146103a15780636a42b8f8146103c8578063704b6c021461042a57806370a082311461044b5780638da5cb5b1461046c5780638f87c84b1461048157806395d89b41146104965780639e5fea8a146104ab578063a9059cbb146104d5578063c0da7e69146104f9578063ce8e120a1461051a578063db6900fa1461052f578063dd62ed3e14610550578063de6ab39c14610577578063dfeb34b61461058c578063e45b8134146105a1578063e6957514146105b6578063f2fde38b146105cb578063f46eccc4146105ec578063f851a4401461060d575b600080fd5b34801561018d57600080fd5b50610196610622565b604080519115158252519081900360200190f35b3480156101b657600080fd5b506101bf610692565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101f95781810151838201526020016101e1565b50505050905090810190601f1680156102265780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024057600080fd5b50610196600160a060020a0360043516602435610720565b34801561026457600080fd5b50610196600160a060020a0360043516610814565b34801561028557600080fd5b5061028e6109c7565b60408051918252519081900360200190f35b3480156102ac57600080fd5b50610196600160a060020a03600435811690602435166044356109cd565b3480156102d657600080fd5b506102df610c27565b60408051600160a060020a039092168252519081900360200190f35b34801561030757600080fd5b5061028e600160a060020a0360043516610c36565b34801561032857600080fd5b50610331610c48565b6040805160ff9092168252519081900360200190f35b34801561035357600080fd5b506102df610c51565b34801561036857600080fd5b50610196600160a060020a0360043516602435610c60565b34801561038c57600080fd5b50610196600160a060020a0360043516610dae565b3480156103ad57600080fd5b5061028e600160a060020a0360043581169060243516610f83565b3480156103d457600080fd5b506103dd610fa0565b60408051600160a060020a03808816825286166020820152908101849052606081018390526080810182600281111561041257fe5b60ff1681526020019550505050505060405180910390f35b34801561043657600080fd5b50610196600160a060020a0360043516610fc5565b34801561045757600080fd5b5061028e600160a060020a036004351661105e565b34801561047857600080fd5b506102df611079565b34801561048d57600080fd5b50610196611088565b3480156104a257600080fd5b506101bf611129565b3480156104b757600080fd5b50610196600160a060020a0360043581169060243516604435611184565b3480156104e157600080fd5b50610196600160a060020a0360043516602435611477565b34801561050557600080fd5b50610196600160a060020a03600435166115cc565b34801561052657600080fd5b50610196611665565b34801561053b57600080fd5b50610196600160a060020a03600435166116d1565b34801561055c57600080fd5b5061028e600160a060020a0360043581169060243516611780565b34801561058357600080fd5b5061028e6117ab565b34801561059857600080fd5b506101966117ba565b3480156105ad57600080fd5b506101966117c9565b3480156105c257600080fd5b506102df6117d7565b3480156105d757600080fd5b50610196600160a060020a03600435166117e6565b3480156105f857600080fd5b50610196600160a060020a036004351661191c565b34801561061957600080fd5b506102df611931565b60008054600160a060020a03163314806106465750600154600160a060020a031633145b151561065157600080fd5b600d805461ff0019166101001790556040516001907fff7ea91c52ebd8c0d8018fdba50cb801e862f6795b1e17eeac882d4288b0934090600090a250600190565b600a805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156107185780601f106106ed57610100808354040283529160200191610718565b820191906000526020600020905b8154815290600101906020018083116106fb57829003601f168201915b505050505081565b6000808211610779576040805160e560020a62461bcd02815260206004820152601d60248201527f616d6f756e74206d7573742062652067726561746572207468616e2030000000604482015290519081900360640190fd5b336000908152600f60209081526040808320600160a060020a03871684529091529020546107ad908363ffffffff61194016565b336000818152600f60209081526040808320600160a060020a03891680855290835292819020949094558351868152935191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600061081e611995565b600054600160a060020a03163314806108415750600154600160a060020a031633145b151561084c57600080fd5b600060065460ff16600281111561085f57fe5b148061087b5750600260065460ff16600281111561087957fe5b145b151561088657600080fd5b6040805160a081018252338152600160a060020a03851660208201529081016108b643600a63ffffffff61194016565b81526020016108dd6108d06064600163ffffffff61195916565b429063ffffffff61194016565b815260200160019052805160028054600160a060020a0392831673ffffffffffffffffffffffffffffffffffffffff199182161782556020840151600380549190941691161790915560408201516004556060820151600555608082015160068054939450849360ff19166001838581111561095557fe5b02179055505081516020808401516040808601516060808801518351600160a060020a03978816815296909416948601949094528482015291830152517f0fd69fbe3cf6151621edd8a8aa0b65868e1cb2d9515ce4038554dca2802265dd92509081900360800190a150600192915050565b600c5490565b600d54600090610100900460ff16156109e557600080fd5b600160a060020a0384166000908152600e6020526040902054821115610a7b576040805160e560020a62461bcd02815260206004820152602160248201527f6f776e657220646f6573206e6f74206861766520656e6f75676820746f6b656e60448201527f7300000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0384166000908152600f60209081526040808320338452909152902054821115610b1c576040805160e560020a62461bcd02815260206004820152602560248201527f73656e64657220646f6573206e6f74206861766520656e6f75676820616c6c6f60448201527f77616e6365000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0384166000908152600f60209081526040808320338452909152902054610b50908363ffffffff61198016565b600160a060020a0385166000818152600f60209081526040808320338452825280832094909455918152600e9091522054610b91908363ffffffff61198016565b600160a060020a038086166000908152600e60205260408082209390935590851681522054610bc6908363ffffffff61194016565b600160a060020a038085166000818152600e602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35060019392505050565b600854600160a060020a031681565b600e6020526000908152604090205481565b600d5460ff1681565b600754600160a060020a031681565b3360009081526010602052604081205460ff161515600114610ccc576040805160e560020a62461bcd02815260206004820152601d60248201527f73656e646572206d75737420626520612076616c6964206d696e746572000000604482015290519081900360640190fd5b600160a060020a0383166000908152600e6020526040902054610cf5908363ffffffff61194016565b600160a060020a0384166000908152600e6020526040902055600c54610d21908363ffffffff61194016565b600c55604080518381529051600160a060020a038516916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3604080518381529051600160a060020a0385169133917f601ace93afe864dd1288a16daad8ee79b21790c32c80d710b03cbae376e9e65f9181900360200190a350600192915050565b60008054600160a060020a03163314610dc657600080fd5b600854600160a060020a031615610ef557600760009054906101000a9004600160a060020a0316600160a060020a031663ed2f23696040518163ffffffff1660e060020a02815260040160206040518083038186803b158015610e2857600080fd5b505afa158015610e3c573d6000803e3d6000fd5b505050506040513d6020811015610e5257600080fd5b505115610ef5576040805160e560020a62461bcd02815260206004820152604b60248201527f7374616b696e6720636f6e747261637420616c726561647920636f6e6669677560448201527f7265642c20746f206368616e6765206974206d7573742068617665203020616360648201527f74697665207374616b6573000000000000000000000000000000000000000000608482015290519081900360a40190fd5b60088054600160a060020a03841673ffffffffffffffffffffffffffffffffffffffff199182168117909255600082815260106020908152604091829020805460ff191660011790556007805490931684179092558051928352517fcf229ad20569d02c4a6cd3b3ae6130cb9e6257558e22a670804b6a6eb866b7149281900390910190a15060015b919050565b600f60209081526000928352604080842090915290825290205481565b600254600354600454600554600654600160a060020a03948516949093169260ff1685565b60008054600160a060020a03163314610fdd57600080fd5b600154600160a060020a0383811691161415610ff857600080fd5b60018054600160a060020a03841673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c9181900360200190a1506001919050565b600160a060020a03166000908152600e602052604090205490565b600054600160a060020a031681565b600080548190600160a060020a03163314806110ae5750600154600160a060020a031633145b15156110b957600080fd5b506040513080319133913180156108fc02916000818181858888f193505050501580156110ea573d6000803e3d6000fd5b5060408051828152905133917ffed66b098dae795e8a862bb1a0d1883d488f015acc2cf25cd29091efa8d8fb6b919081900360200190a2600191505090565b600b805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156107185780601f106106ed57610100808354040283529160200191610718565b600080548190600160a060020a03163314806111aa5750600154600160a060020a031633145b15156111b557600080fd5b600160a060020a03851630141561123b576040805160e560020a62461bcd028152602060048201526024808201527f746f6b656e20616464726573732063616e2774206265207468697320636f6e7460448201527f7261637400000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b50604080517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152905185918491600160a060020a038416916370a08231916024808301926020929190829003018186803b15801561129f57600080fd5b505afa1580156112b3573d6000803e3d6000fd5b505050506040513d60208110156112c957600080fd5b50511015611347576040805160e560020a62461bcd02815260206004820152603360248201527f617474656d7074696e6720746f2073656e64206d6f726520746f6b656e73207460448201527f68616e2063757272656e742062616c616e636500000000000000000000000000606482015290519081900360840190fd5b80600160a060020a031663a9059cbb85856040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156113aa57600080fd5b505af11580156113be573d6000803e3d6000fd5b505050506040513d60208110156113d457600080fd5b5051151561142c576040805160e560020a62461bcd02815260206004820152601560248201527f746f6b656e207472616e73666572206661696c65640000000000000000000000604482015290519081900360640190fd5b604080518481529051600160a060020a0386169133917f10a46ed575affad8e954ae27853b1f89c6da90d8c35f619fc640f8a21bcb78579181900360200190a3506001949350505050565b600d54600090610100900460ff161561148f57600080fd5b336000908152600e602052604090205482111561151c576040805160e560020a62461bcd02815260206004820152602260248201527f73656e64657220646f6573206e6f74206861766520656e6f75676820746f6b6560448201527f6e73000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b336000908152600e602052604090205461153c908363ffffffff61198016565b336000908152600e602052604080822092909255600160a060020a0385168152205461156e908363ffffffff61194016565b600160a060020a0384166000818152600e60209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b60008054600160a060020a031633146115e457600080fd5b6009805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038416908117909155600081815260106020908152604091829020805460ff19166001179055815192835290517f09eeb152b2546a9b79b2896b1b207bd9d9e94f00c0fad15b068e84478511bd529281900390910190a1506001919050565b60008054600160a060020a03163314806116895750600154600160a060020a031633145b151561169457600080fd5b600d805461ff00191690556040516001907fb36ea4d45a6246e5ea6da988f57a5bf9a9022c85940cc6fe92dd9e45bf632cf690600090a250600190565b60008054600160a060020a031633146116e957600080fd5b600d5462010000900460ff1615156117145750600d805462ff00001916620100001790556001610f7e565b600160a060020a038216600081815260106020908152604091829020805460ff19166001179055600d805462ff000019169055815192835290517f540af0fc125e4047c03435dd52febc08726667f13f9c4ac5e8795a451be52f8b9281900390910190a1506001610f7e565b600160a060020a039182166000908152600f6020908152604080832093909416825291909152205490565b6a32f44eb0f61c612400000081565b600d5462010000900460ff1681565b600d54610100900460ff1681565b600954600160a060020a031681565b60008054600160a060020a031633146117fe57600080fd5b81600160065460ff16600281111561181257fe5b1461181c57600080fd5b600354600160a060020a03828116911614611881576040805160e560020a62461bcd02815260206004820152601360248201527f696e636f7272656374206e6577206f776e657200000000000000000000000000604482015290519081900360640190fd5b60055442101561189057600080fd5b60045443101561189f57600080fd5b6006805460ff1916600217905560008054600160a060020a03851673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051338152602081019290925280517f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09281900390910190a150600192915050565b60106020526000908152604090205460ff1681565b600154600160a060020a031681565b60008282018381101561195257600080fd5b9392505050565b6000828202831580611975575082848281151561197257fe5b04145b151561195257600080fd5b60008282111561198f57600080fd5b50900390565b6040805160a08101825260008082526020820181905291810182905260608101829052906080820152905600a165627a7a723058201a8b66c736f6c7876ad30285618fd4a8db0f29b0d11f0e01a4fddea6961626aa0029`

// DeployRTCoin deploys a new Ethereum contract, binding an instance of RTCoin to it.
func DeployRTCoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RTCoin, error) {
	parsed, err := abi.JSON(strings.NewReader(RTCoinABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RTCoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RTCoin{RTCoinCaller: RTCoinCaller{contract: contract}, RTCoinTransactor: RTCoinTransactor{contract: contract}, RTCoinFilterer: RTCoinFilterer{contract: contract}}, nil
}

// RTCoin is an auto generated Go binding around an Ethereum contract.
type RTCoin struct {
	RTCoinCaller     // Read-only binding to the contract
	RTCoinTransactor // Write-only binding to the contract
	RTCoinFilterer   // Log filterer for contract events
}

// RTCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type RTCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RTCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RTCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RTCoinSession struct {
	Contract     *RTCoin           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RTCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RTCoinCallerSession struct {
	Contract *RTCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RTCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RTCoinTransactorSession struct {
	Contract     *RTCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RTCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type RTCoinRaw struct {
	Contract *RTCoin // Generic contract binding to access the raw methods on
}

// RTCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RTCoinCallerRaw struct {
	Contract *RTCoinCaller // Generic read-only contract binding to access the raw methods on
}

// RTCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RTCoinTransactorRaw struct {
	Contract *RTCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRTCoin creates a new instance of RTCoin, bound to a specific deployed contract.
func NewRTCoin(address common.Address, backend bind.ContractBackend) (*RTCoin, error) {
	contract, err := bindRTCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RTCoin{RTCoinCaller: RTCoinCaller{contract: contract}, RTCoinTransactor: RTCoinTransactor{contract: contract}, RTCoinFilterer: RTCoinFilterer{contract: contract}}, nil
}

// NewRTCoinCaller creates a new read-only instance of RTCoin, bound to a specific deployed contract.
func NewRTCoinCaller(address common.Address, caller bind.ContractCaller) (*RTCoinCaller, error) {
	contract, err := bindRTCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RTCoinCaller{contract: contract}, nil
}

// NewRTCoinTransactor creates a new write-only instance of RTCoin, bound to a specific deployed contract.
func NewRTCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*RTCoinTransactor, error) {
	contract, err := bindRTCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RTCoinTransactor{contract: contract}, nil
}

// NewRTCoinFilterer creates a new log filterer instance of RTCoin, bound to a specific deployed contract.
func NewRTCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*RTCoinFilterer, error) {
	contract, err := bindRTCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RTCoinFilterer{contract: contract}, nil
}

// bindRTCoin binds a generic wrapper to an already deployed contract.
func bindRTCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RTCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RTCoin *RTCoinRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RTCoin.Contract.RTCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RTCoin *RTCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCoin.Contract.RTCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RTCoin *RTCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RTCoin.Contract.RTCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RTCoin *RTCoinCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RTCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RTCoin *RTCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RTCoin *RTCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RTCoin.Contract.contract.Transact(opts, method, params...)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0xde6ab39c.
//
// Solidity: function INITIALSUPPLY() constant returns(uint256)
func (_RTCoin *RTCoinCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "INITIALSUPPLY")
	return *ret0, err
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0xde6ab39c.
//
// Solidity: function INITIALSUPPLY() constant returns(uint256)
func (_RTCoin *RTCoinSession) INITIALSUPPLY() (*big.Int, error) {
	return _RTCoin.Contract.INITIALSUPPLY(&_RTCoin.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0xde6ab39c.
//
// Solidity: function INITIALSUPPLY() constant returns(uint256)
func (_RTCoin *RTCoinCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _RTCoin.Contract.INITIALSUPPLY(&_RTCoin.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_RTCoin *RTCoinCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_RTCoin *RTCoinSession) Admin() (common.Address, error) {
	return _RTCoin.Contract.Admin(&_RTCoin.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_RTCoin *RTCoinCallerSession) Admin() (common.Address, error) {
	return _RTCoin.Contract.Admin(&_RTCoin.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RTCoin *RTCoinCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RTCoin *RTCoinSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RTCoin.Contract.Allowance(&_RTCoin.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RTCoin *RTCoinCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RTCoin.Contract.Allowance(&_RTCoin.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_RTCoin *RTCoinCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "allowed", arg0, arg1)
	return *ret0, err
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_RTCoin *RTCoinSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RTCoin.Contract.Allowed(&_RTCoin.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_RTCoin *RTCoinCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RTCoin.Contract.Allowed(&_RTCoin.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_holder address) constant returns(uint256)
func (_RTCoin *RTCoinCaller) BalanceOf(opts *bind.CallOpts, _holder common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "balanceOf", _holder)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_holder address) constant returns(uint256)
func (_RTCoin *RTCoinSession) BalanceOf(_holder common.Address) (*big.Int, error) {
	return _RTCoin.Contract.BalanceOf(&_RTCoin.CallOpts, _holder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_holder address) constant returns(uint256)
func (_RTCoin *RTCoinCallerSession) BalanceOf(_holder common.Address) (*big.Int, error) {
	return _RTCoin.Contract.BalanceOf(&_RTCoin.CallOpts, _holder)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_RTCoin *RTCoinCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_RTCoin *RTCoinSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RTCoin.Contract.Balances(&_RTCoin.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_RTCoin *RTCoinCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RTCoin.Contract.Balances(&_RTCoin.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RTCoin *RTCoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RTCoin *RTCoinSession) Decimals() (uint8, error) {
	return _RTCoin.Contract.Decimals(&_RTCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RTCoin *RTCoinCallerSession) Decimals() (uint8, error) {
	return _RTCoin.Contract.Decimals(&_RTCoin.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() constant returns(previousOwner address, newOwner address, delayExpirationBlock uint256, delayExpirationTime uint256, state uint8)
func (_RTCoin *RTCoinCaller) Delay(opts *bind.CallOpts) (struct {
	PreviousOwner        common.Address
	NewOwner             common.Address
	DelayExpirationBlock *big.Int
	DelayExpirationTime  *big.Int
	State                uint8
}, error) {
	ret := new(struct {
		PreviousOwner        common.Address
		NewOwner             common.Address
		DelayExpirationBlock *big.Int
		DelayExpirationTime  *big.Int
		State                uint8
	})
	out := ret
	err := _RTCoin.contract.Call(opts, out, "delay")
	return *ret, err
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() constant returns(previousOwner address, newOwner address, delayExpirationBlock uint256, delayExpirationTime uint256, state uint8)
func (_RTCoin *RTCoinSession) Delay() (struct {
	PreviousOwner        common.Address
	NewOwner             common.Address
	DelayExpirationBlock *big.Int
	DelayExpirationTime  *big.Int
	State                uint8
}, error) {
	return _RTCoin.Contract.Delay(&_RTCoin.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() constant returns(previousOwner address, newOwner address, delayExpirationBlock uint256, delayExpirationTime uint256, state uint8)
func (_RTCoin *RTCoinCallerSession) Delay() (struct {
	PreviousOwner        common.Address
	NewOwner             common.Address
	DelayExpirationBlock *big.Int
	DelayExpirationTime  *big.Int
	State                uint8
}, error) {
	return _RTCoin.Contract.Delay(&_RTCoin.CallOpts)
}

// MergedMinerValidatorAddress is a free data retrieval call binding the contract method 0xe6957514.
//
// Solidity: function mergedMinerValidatorAddress() constant returns(address)
func (_RTCoin *RTCoinCaller) MergedMinerValidatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "mergedMinerValidatorAddress")
	return *ret0, err
}

// MergedMinerValidatorAddress is a free data retrieval call binding the contract method 0xe6957514.
//
// Solidity: function mergedMinerValidatorAddress() constant returns(address)
func (_RTCoin *RTCoinSession) MergedMinerValidatorAddress() (common.Address, error) {
	return _RTCoin.Contract.MergedMinerValidatorAddress(&_RTCoin.CallOpts)
}

// MergedMinerValidatorAddress is a free data retrieval call binding the contract method 0xe6957514.
//
// Solidity: function mergedMinerValidatorAddress() constant returns(address)
func (_RTCoin *RTCoinCallerSession) MergedMinerValidatorAddress() (common.Address, error) {
	return _RTCoin.Contract.MergedMinerValidatorAddress(&_RTCoin.CallOpts)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters( address) constant returns(bool)
func (_RTCoin *RTCoinCaller) Minters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "minters", arg0)
	return *ret0, err
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters( address) constant returns(bool)
func (_RTCoin *RTCoinSession) Minters(arg0 common.Address) (bool, error) {
	return _RTCoin.Contract.Minters(&_RTCoin.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters( address) constant returns(bool)
func (_RTCoin *RTCoinCallerSession) Minters(arg0 common.Address) (bool, error) {
	return _RTCoin.Contract.Minters(&_RTCoin.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RTCoin *RTCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RTCoin *RTCoinSession) Name() (string, error) {
	return _RTCoin.Contract.Name(&_RTCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RTCoin *RTCoinCallerSession) Name() (string, error) {
	return _RTCoin.Contract.Name(&_RTCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RTCoin *RTCoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RTCoin *RTCoinSession) Owner() (common.Address, error) {
	return _RTCoin.Contract.Owner(&_RTCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RTCoin *RTCoinCallerSession) Owner() (common.Address, error) {
	return _RTCoin.Contract.Owner(&_RTCoin.CallOpts)
}

// Stake is a free data retrieval call binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() constant returns(address)
func (_RTCoin *RTCoinCaller) Stake(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "stake")
	return *ret0, err
}

// Stake is a free data retrieval call binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() constant returns(address)
func (_RTCoin *RTCoinSession) Stake() (common.Address, error) {
	return _RTCoin.Contract.Stake(&_RTCoin.CallOpts)
}

// Stake is a free data retrieval call binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() constant returns(address)
func (_RTCoin *RTCoinCallerSession) Stake() (common.Address, error) {
	return _RTCoin.Contract.Stake(&_RTCoin.CallOpts)
}

// StakeContractAddress is a free data retrieval call binding the contract method 0x272caf69.
//
// Solidity: function stakeContractAddress() constant returns(address)
func (_RTCoin *RTCoinCaller) StakeContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "stakeContractAddress")
	return *ret0, err
}

// StakeContractAddress is a free data retrieval call binding the contract method 0x272caf69.
//
// Solidity: function stakeContractAddress() constant returns(address)
func (_RTCoin *RTCoinSession) StakeContractAddress() (common.Address, error) {
	return _RTCoin.Contract.StakeContractAddress(&_RTCoin.CallOpts)
}

// StakeContractAddress is a free data retrieval call binding the contract method 0x272caf69.
//
// Solidity: function stakeContractAddress() constant returns(address)
func (_RTCoin *RTCoinCallerSession) StakeContractAddress() (common.Address, error) {
	return _RTCoin.Contract.StakeContractAddress(&_RTCoin.CallOpts)
}

// StakeFailOverRestrictionLifted is a free data retrieval call binding the contract method 0xdfeb34b6.
//
// Solidity: function stakeFailOverRestrictionLifted() constant returns(bool)
func (_RTCoin *RTCoinCaller) StakeFailOverRestrictionLifted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "stakeFailOverRestrictionLifted")
	return *ret0, err
}

// StakeFailOverRestrictionLifted is a free data retrieval call binding the contract method 0xdfeb34b6.
//
// Solidity: function stakeFailOverRestrictionLifted() constant returns(bool)
func (_RTCoin *RTCoinSession) StakeFailOverRestrictionLifted() (bool, error) {
	return _RTCoin.Contract.StakeFailOverRestrictionLifted(&_RTCoin.CallOpts)
}

// StakeFailOverRestrictionLifted is a free data retrieval call binding the contract method 0xdfeb34b6.
//
// Solidity: function stakeFailOverRestrictionLifted() constant returns(bool)
func (_RTCoin *RTCoinCallerSession) StakeFailOverRestrictionLifted() (bool, error) {
	return _RTCoin.Contract.StakeFailOverRestrictionLifted(&_RTCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RTCoin *RTCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RTCoin *RTCoinSession) Symbol() (string, error) {
	return _RTCoin.Contract.Symbol(&_RTCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RTCoin *RTCoinCallerSession) Symbol() (string, error) {
	return _RTCoin.Contract.Symbol(&_RTCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RTCoin *RTCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RTCoin *RTCoinSession) TotalSupply() (*big.Int, error) {
	return _RTCoin.Contract.TotalSupply(&_RTCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RTCoin *RTCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _RTCoin.Contract.TotalSupply(&_RTCoin.CallOpts)
}

// TransfersFrozen is a free data retrieval call binding the contract method 0xe45b8134.
//
// Solidity: function transfersFrozen() constant returns(bool)
func (_RTCoin *RTCoinCaller) TransfersFrozen(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "transfersFrozen")
	return *ret0, err
}

// TransfersFrozen is a free data retrieval call binding the contract method 0xe45b8134.
//
// Solidity: function transfersFrozen() constant returns(bool)
func (_RTCoin *RTCoinSession) TransfersFrozen() (bool, error) {
	return _RTCoin.Contract.TransfersFrozen(&_RTCoin.CallOpts)
}

// TransfersFrozen is a free data retrieval call binding the contract method 0xe45b8134.
//
// Solidity: function transfersFrozen() constant returns(bool)
func (_RTCoin *RTCoinCallerSession) TransfersFrozen() (bool, error) {
	return _RTCoin.Contract.TransfersFrozen(&_RTCoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.Approve(&_RTCoin.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.Approve(&_RTCoin.TransactOpts, _spender, _amount)
}

// FreezeTransfers is a paid mutator transaction binding the contract method 0x01502460.
//
// Solidity: function freezeTransfers() returns(bool)
func (_RTCoin *RTCoinTransactor) FreezeTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "freezeTransfers")
}

// FreezeTransfers is a paid mutator transaction binding the contract method 0x01502460.
//
// Solidity: function freezeTransfers() returns(bool)
func (_RTCoin *RTCoinSession) FreezeTransfers() (*types.Transaction, error) {
	return _RTCoin.Contract.FreezeTransfers(&_RTCoin.TransactOpts)
}

// FreezeTransfers is a paid mutator transaction binding the contract method 0x01502460.
//
// Solidity: function freezeTransfers() returns(bool)
func (_RTCoin *RTCoinTransactorSession) FreezeTransfers() (*types.Transaction, error) {
	return _RTCoin.Contract.FreezeTransfers(&_RTCoin.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_recipient address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinTransactor) Mint(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "mint", _recipient, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_recipient address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinSession) Mint(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.Mint(&_RTCoin.TransactOpts, _recipient, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_recipient address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinTransactorSession) Mint(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.Mint(&_RTCoin.TransactOpts, _recipient, _amount)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_RTCoin *RTCoinTransactor) SetAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "setAdmin", _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_RTCoin *RTCoinSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.SetAdmin(&_RTCoin.TransactOpts, _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_RTCoin *RTCoinTransactorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.SetAdmin(&_RTCoin.TransactOpts, _newAdmin)
}

// SetFailOverStakeContract is a paid mutator transaction binding the contract method 0xdb6900fa.
//
// Solidity: function setFailOverStakeContract(_contractAddress address) returns(bool)
func (_RTCoin *RTCoinTransactor) SetFailOverStakeContract(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "setFailOverStakeContract", _contractAddress)
}

// SetFailOverStakeContract is a paid mutator transaction binding the contract method 0xdb6900fa.
//
// Solidity: function setFailOverStakeContract(_contractAddress address) returns(bool)
func (_RTCoin *RTCoinSession) SetFailOverStakeContract(_contractAddress common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.SetFailOverStakeContract(&_RTCoin.TransactOpts, _contractAddress)
}

// SetFailOverStakeContract is a paid mutator transaction binding the contract method 0xdb6900fa.
//
// Solidity: function setFailOverStakeContract(_contractAddress address) returns(bool)
func (_RTCoin *RTCoinTransactorSession) SetFailOverStakeContract(_contractAddress common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.SetFailOverStakeContract(&_RTCoin.TransactOpts, _contractAddress)
}

// SetMergedMinerValidator is a paid mutator transaction binding the contract method 0xc0da7e69.
//
// Solidity: function setMergedMinerValidator(_mergedMinerValidator address) returns(bool)
func (_RTCoin *RTCoinTransactor) SetMergedMinerValidator(opts *bind.TransactOpts, _mergedMinerValidator common.Address) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "setMergedMinerValidator", _mergedMinerValidator)
}

// SetMergedMinerValidator is a paid mutator transaction binding the contract method 0xc0da7e69.
//
// Solidity: function setMergedMinerValidator(_mergedMinerValidator address) returns(bool)
func (_RTCoin *RTCoinSession) SetMergedMinerValidator(_mergedMinerValidator common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.SetMergedMinerValidator(&_RTCoin.TransactOpts, _mergedMinerValidator)
}

// SetMergedMinerValidator is a paid mutator transaction binding the contract method 0xc0da7e69.
//
// Solidity: function setMergedMinerValidator(_mergedMinerValidator address) returns(bool)
func (_RTCoin *RTCoinTransactorSession) SetMergedMinerValidator(_mergedMinerValidator common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.SetMergedMinerValidator(&_RTCoin.TransactOpts, _mergedMinerValidator)
}

// SetStakeContract is a paid mutator transaction binding the contract method 0x509484d5.
//
// Solidity: function setStakeContract(_contractAddress address) returns(bool)
func (_RTCoin *RTCoinTransactor) SetStakeContract(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "setStakeContract", _contractAddress)
}

// SetStakeContract is a paid mutator transaction binding the contract method 0x509484d5.
//
// Solidity: function setStakeContract(_contractAddress address) returns(bool)
func (_RTCoin *RTCoinSession) SetStakeContract(_contractAddress common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.SetStakeContract(&_RTCoin.TransactOpts, _contractAddress)
}

// SetStakeContract is a paid mutator transaction binding the contract method 0x509484d5.
//
// Solidity: function setStakeContract(_contractAddress address) returns(bool)
func (_RTCoin *RTCoinTransactorSession) SetStakeContract(_contractAddress common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.SetStakeContract(&_RTCoin.TransactOpts, _contractAddress)
}

// StartOwnerTransferDelay is a paid mutator transaction binding the contract method 0x0ff69b52.
//
// Solidity: function startOwnerTransferDelay(_newOwner address) returns(bool)
func (_RTCoin *RTCoinTransactor) StartOwnerTransferDelay(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "startOwnerTransferDelay", _newOwner)
}

// StartOwnerTransferDelay is a paid mutator transaction binding the contract method 0x0ff69b52.
//
// Solidity: function startOwnerTransferDelay(_newOwner address) returns(bool)
func (_RTCoin *RTCoinSession) StartOwnerTransferDelay(_newOwner common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.StartOwnerTransferDelay(&_RTCoin.TransactOpts, _newOwner)
}

// StartOwnerTransferDelay is a paid mutator transaction binding the contract method 0x0ff69b52.
//
// Solidity: function startOwnerTransferDelay(_newOwner address) returns(bool)
func (_RTCoin *RTCoinTransactorSession) StartOwnerTransferDelay(_newOwner common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.StartOwnerTransferDelay(&_RTCoin.TransactOpts, _newOwner)
}

// ThawTransfers is a paid mutator transaction binding the contract method 0xce8e120a.
//
// Solidity: function thawTransfers() returns(bool)
func (_RTCoin *RTCoinTransactor) ThawTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "thawTransfers")
}

// ThawTransfers is a paid mutator transaction binding the contract method 0xce8e120a.
//
// Solidity: function thawTransfers() returns(bool)
func (_RTCoin *RTCoinSession) ThawTransfers() (*types.Transaction, error) {
	return _RTCoin.Contract.ThawTransfers(&_RTCoin.TransactOpts)
}

// ThawTransfers is a paid mutator transaction binding the contract method 0xce8e120a.
//
// Solidity: function thawTransfers() returns(bool)
func (_RTCoin *RTCoinTransactorSession) ThawTransfers() (*types.Transaction, error) {
	return _RTCoin.Contract.ThawTransfers(&_RTCoin.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_recipient address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_recipient address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.Transfer(&_RTCoin.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_recipient address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.Transfer(&_RTCoin.TransactOpts, _recipient, _amount)
}

// TransferForeignToken is a paid mutator transaction binding the contract method 0x9e5fea8a.
//
// Solidity: function transferForeignToken(_tokenAddress address, _recipient address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinTransactor) TransferForeignToken(opts *bind.TransactOpts, _tokenAddress common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "transferForeignToken", _tokenAddress, _recipient, _amount)
}

// TransferForeignToken is a paid mutator transaction binding the contract method 0x9e5fea8a.
//
// Solidity: function transferForeignToken(_tokenAddress address, _recipient address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinSession) TransferForeignToken(_tokenAddress common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.TransferForeignToken(&_RTCoin.TransactOpts, _tokenAddress, _recipient, _amount)
}

// TransferForeignToken is a paid mutator transaction binding the contract method 0x9e5fea8a.
//
// Solidity: function transferForeignToken(_tokenAddress address, _recipient address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinTransactorSession) TransferForeignToken(_tokenAddress common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.TransferForeignToken(&_RTCoin.TransactOpts, _tokenAddress, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_owner address, _recipient address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinTransactor) TransferFrom(opts *bind.TransactOpts, _owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "transferFrom", _owner, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_owner address, _recipient address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinSession) TransferFrom(_owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.TransferFrom(&_RTCoin.TransactOpts, _owner, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_owner address, _recipient address, _amount uint256) returns(bool)
func (_RTCoin *RTCoinTransactorSession) TransferFrom(_owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.TransferFrom(&_RTCoin.TransactOpts, _owner, _recipient, _amount)
}

// TransferOutEth is a paid mutator transaction binding the contract method 0x8f87c84b.
//
// Solidity: function transferOutEth() returns(bool)
func (_RTCoin *RTCoinTransactor) TransferOutEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "transferOutEth")
}

// TransferOutEth is a paid mutator transaction binding the contract method 0x8f87c84b.
//
// Solidity: function transferOutEth() returns(bool)
func (_RTCoin *RTCoinSession) TransferOutEth() (*types.Transaction, error) {
	return _RTCoin.Contract.TransferOutEth(&_RTCoin.TransactOpts)
}

// TransferOutEth is a paid mutator transaction binding the contract method 0x8f87c84b.
//
// Solidity: function transferOutEth() returns(bool)
func (_RTCoin *RTCoinTransactorSession) TransferOutEth() (*types.Transaction, error) {
	return _RTCoin.Contract.TransferOutEth(&_RTCoin.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_RTCoin *RTCoinTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_RTCoin *RTCoinSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.TransferOwnership(&_RTCoin.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_RTCoin *RTCoinTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.TransferOwnership(&_RTCoin.TransactOpts, _newOwner)
}

// RTCoinAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the RTCoin contract.
type RTCoinAdminSetIterator struct {
	Event *RTCoinAdminSet // Event containing the contract specifics and raw log

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
func (it *RTCoinAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinAdminSet)
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
		it.Event = new(RTCoinAdminSet)
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
func (it *RTCoinAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinAdminSet represents a AdminSet event raised by the RTCoin contract.
type RTCoinAdminSet struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: e AdminSet(_admin address)
func (_RTCoin *RTCoinFilterer) FilterAdminSet(opts *bind.FilterOpts) (*RTCoinAdminSetIterator, error) {

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &RTCoinAdminSetIterator{contract: _RTCoin.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: e AdminSet(_admin address)
func (_RTCoin *RTCoinFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *RTCoinAdminSet) (event.Subscription, error) {

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinAdminSet)
				if err := _RTCoin.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// RTCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RTCoin contract.
type RTCoinApprovalIterator struct {
	Event *RTCoinApproval // Event containing the contract specifics and raw log

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
func (it *RTCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinApproval)
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
		it.Event = new(RTCoinApproval)
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
func (it *RTCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinApproval represents a Approval event raised by the RTCoin contract.
type RTCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _amount uint256)
func (_RTCoin *RTCoinFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*RTCoinApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &RTCoinApprovalIterator{contract: _RTCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _amount uint256)
func (_RTCoin *RTCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RTCoinApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinApproval)
				if err := _RTCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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

// RTCoinCoinsMintedIterator is returned from FilterCoinsMinted and is used to iterate over the raw logs and unpacked data for CoinsMinted events raised by the RTCoin contract.
type RTCoinCoinsMintedIterator struct {
	Event *RTCoinCoinsMinted // Event containing the contract specifics and raw log

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
func (it *RTCoinCoinsMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinCoinsMinted)
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
		it.Event = new(RTCoinCoinsMinted)
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
func (it *RTCoinCoinsMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinCoinsMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinCoinsMinted represents a CoinsMinted event raised by the RTCoin contract.
type RTCoinCoinsMinted struct {
	StakeContract common.Address
	Recipient     common.Address
	MintAmount    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCoinsMinted is a free log retrieval operation binding the contract event 0x601ace93afe864dd1288a16daad8ee79b21790c32c80d710b03cbae376e9e65f.
//
// Solidity: e CoinsMinted(_stakeContract indexed address, _recipient indexed address, _mintAmount uint256)
func (_RTCoin *RTCoinFilterer) FilterCoinsMinted(opts *bind.FilterOpts, _stakeContract []common.Address, _recipient []common.Address) (*RTCoinCoinsMintedIterator, error) {

	var _stakeContractRule []interface{}
	for _, _stakeContractItem := range _stakeContract {
		_stakeContractRule = append(_stakeContractRule, _stakeContractItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "CoinsMinted", _stakeContractRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return &RTCoinCoinsMintedIterator{contract: _RTCoin.contract, event: "CoinsMinted", logs: logs, sub: sub}, nil
}

// WatchCoinsMinted is a free log subscription operation binding the contract event 0x601ace93afe864dd1288a16daad8ee79b21790c32c80d710b03cbae376e9e65f.
//
// Solidity: e CoinsMinted(_stakeContract indexed address, _recipient indexed address, _mintAmount uint256)
func (_RTCoin *RTCoinFilterer) WatchCoinsMinted(opts *bind.WatchOpts, sink chan<- *RTCoinCoinsMinted, _stakeContract []common.Address, _recipient []common.Address) (event.Subscription, error) {

	var _stakeContractRule []interface{}
	for _, _stakeContractItem := range _stakeContract {
		_stakeContractRule = append(_stakeContractRule, _stakeContractItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "CoinsMinted", _stakeContractRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinCoinsMinted)
				if err := _RTCoin.contract.UnpackLog(event, "CoinsMinted", log); err != nil {
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

// RTCoinEthTransferOutIterator is returned from FilterEthTransferOut and is used to iterate over the raw logs and unpacked data for EthTransferOut events raised by the RTCoin contract.
type RTCoinEthTransferOutIterator struct {
	Event *RTCoinEthTransferOut // Event containing the contract specifics and raw log

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
func (it *RTCoinEthTransferOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinEthTransferOut)
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
		it.Event = new(RTCoinEthTransferOut)
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
func (it *RTCoinEthTransferOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinEthTransferOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinEthTransferOut represents a EthTransferOut event raised by the RTCoin contract.
type RTCoinEthTransferOut struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthTransferOut is a free log retrieval operation binding the contract event 0xfed66b098dae795e8a862bb1a0d1883d488f015acc2cf25cd29091efa8d8fb6b.
//
// Solidity: e EthTransferOut(_recipient indexed address, _amount uint256)
func (_RTCoin *RTCoinFilterer) FilterEthTransferOut(opts *bind.FilterOpts, _recipient []common.Address) (*RTCoinEthTransferOutIterator, error) {

	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "EthTransferOut", _recipientRule)
	if err != nil {
		return nil, err
	}
	return &RTCoinEthTransferOutIterator{contract: _RTCoin.contract, event: "EthTransferOut", logs: logs, sub: sub}, nil
}

// WatchEthTransferOut is a free log subscription operation binding the contract event 0xfed66b098dae795e8a862bb1a0d1883d488f015acc2cf25cd29091efa8d8fb6b.
//
// Solidity: e EthTransferOut(_recipient indexed address, _amount uint256)
func (_RTCoin *RTCoinFilterer) WatchEthTransferOut(opts *bind.WatchOpts, sink chan<- *RTCoinEthTransferOut, _recipient []common.Address) (event.Subscription, error) {

	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "EthTransferOut", _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinEthTransferOut)
				if err := _RTCoin.contract.UnpackLog(event, "EthTransferOut", log); err != nil {
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

// RTCoinFailOverStakeContractSetIterator is returned from FilterFailOverStakeContractSet and is used to iterate over the raw logs and unpacked data for FailOverStakeContractSet events raised by the RTCoin contract.
type RTCoinFailOverStakeContractSetIterator struct {
	Event *RTCoinFailOverStakeContractSet // Event containing the contract specifics and raw log

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
func (it *RTCoinFailOverStakeContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinFailOverStakeContractSet)
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
		it.Event = new(RTCoinFailOverStakeContractSet)
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
func (it *RTCoinFailOverStakeContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinFailOverStakeContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinFailOverStakeContractSet represents a FailOverStakeContractSet event raised by the RTCoin contract.
type RTCoinFailOverStakeContractSet struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFailOverStakeContractSet is a free log retrieval operation binding the contract event 0x540af0fc125e4047c03435dd52febc08726667f13f9c4ac5e8795a451be52f8b.
//
// Solidity: e FailOverStakeContractSet(_contractAddress address)
func (_RTCoin *RTCoinFilterer) FilterFailOverStakeContractSet(opts *bind.FilterOpts) (*RTCoinFailOverStakeContractSetIterator, error) {

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "FailOverStakeContractSet")
	if err != nil {
		return nil, err
	}
	return &RTCoinFailOverStakeContractSetIterator{contract: _RTCoin.contract, event: "FailOverStakeContractSet", logs: logs, sub: sub}, nil
}

// WatchFailOverStakeContractSet is a free log subscription operation binding the contract event 0x540af0fc125e4047c03435dd52febc08726667f13f9c4ac5e8795a451be52f8b.
//
// Solidity: e FailOverStakeContractSet(_contractAddress address)
func (_RTCoin *RTCoinFilterer) WatchFailOverStakeContractSet(opts *bind.WatchOpts, sink chan<- *RTCoinFailOverStakeContractSet) (event.Subscription, error) {

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "FailOverStakeContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinFailOverStakeContractSet)
				if err := _RTCoin.contract.UnpackLog(event, "FailOverStakeContractSet", log); err != nil {
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

// RTCoinForeignTokenTransferIterator is returned from FilterForeignTokenTransfer and is used to iterate over the raw logs and unpacked data for ForeignTokenTransfer events raised by the RTCoin contract.
type RTCoinForeignTokenTransferIterator struct {
	Event *RTCoinForeignTokenTransfer // Event containing the contract specifics and raw log

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
func (it *RTCoinForeignTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinForeignTokenTransfer)
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
		it.Event = new(RTCoinForeignTokenTransfer)
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
func (it *RTCoinForeignTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinForeignTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinForeignTokenTransfer represents a ForeignTokenTransfer event raised by the RTCoin contract.
type RTCoinForeignTokenTransfer struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterForeignTokenTransfer is a free log retrieval operation binding the contract event 0x10a46ed575affad8e954ae27853b1f89c6da90d8c35f619fc640f8a21bcb7857.
//
// Solidity: e ForeignTokenTransfer(_sender indexed address, _recipient indexed address, _amount uint256)
func (_RTCoin *RTCoinFilterer) FilterForeignTokenTransfer(opts *bind.FilterOpts, _sender []common.Address, _recipient []common.Address) (*RTCoinForeignTokenTransferIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "ForeignTokenTransfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return &RTCoinForeignTokenTransferIterator{contract: _RTCoin.contract, event: "ForeignTokenTransfer", logs: logs, sub: sub}, nil
}

// WatchForeignTokenTransfer is a free log subscription operation binding the contract event 0x10a46ed575affad8e954ae27853b1f89c6da90d8c35f619fc640f8a21bcb7857.
//
// Solidity: e ForeignTokenTransfer(_sender indexed address, _recipient indexed address, _amount uint256)
func (_RTCoin *RTCoinFilterer) WatchForeignTokenTransfer(opts *bind.WatchOpts, sink chan<- *RTCoinForeignTokenTransfer, _sender []common.Address, _recipient []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "ForeignTokenTransfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinForeignTokenTransfer)
				if err := _RTCoin.contract.UnpackLog(event, "ForeignTokenTransfer", log); err != nil {
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

// RTCoinMergedMinerValidatorSetIterator is returned from FilterMergedMinerValidatorSet and is used to iterate over the raw logs and unpacked data for MergedMinerValidatorSet events raised by the RTCoin contract.
type RTCoinMergedMinerValidatorSetIterator struct {
	Event *RTCoinMergedMinerValidatorSet // Event containing the contract specifics and raw log

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
func (it *RTCoinMergedMinerValidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinMergedMinerValidatorSet)
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
		it.Event = new(RTCoinMergedMinerValidatorSet)
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
func (it *RTCoinMergedMinerValidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinMergedMinerValidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinMergedMinerValidatorSet represents a MergedMinerValidatorSet event raised by the RTCoin contract.
type RTCoinMergedMinerValidatorSet struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMergedMinerValidatorSet is a free log retrieval operation binding the contract event 0x09eeb152b2546a9b79b2896b1b207bd9d9e94f00c0fad15b068e84478511bd52.
//
// Solidity: e MergedMinerValidatorSet(_contractAddress address)
func (_RTCoin *RTCoinFilterer) FilterMergedMinerValidatorSet(opts *bind.FilterOpts) (*RTCoinMergedMinerValidatorSetIterator, error) {

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "MergedMinerValidatorSet")
	if err != nil {
		return nil, err
	}
	return &RTCoinMergedMinerValidatorSetIterator{contract: _RTCoin.contract, event: "MergedMinerValidatorSet", logs: logs, sub: sub}, nil
}

// WatchMergedMinerValidatorSet is a free log subscription operation binding the contract event 0x09eeb152b2546a9b79b2896b1b207bd9d9e94f00c0fad15b068e84478511bd52.
//
// Solidity: e MergedMinerValidatorSet(_contractAddress address)
func (_RTCoin *RTCoinFilterer) WatchMergedMinerValidatorSet(opts *bind.WatchOpts, sink chan<- *RTCoinMergedMinerValidatorSet) (event.Subscription, error) {

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "MergedMinerValidatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinMergedMinerValidatorSet)
				if err := _RTCoin.contract.UnpackLog(event, "MergedMinerValidatorSet", log); err != nil {
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

// RTCoinOwnerTransferDelayStartedIterator is returned from FilterOwnerTransferDelayStarted and is used to iterate over the raw logs and unpacked data for OwnerTransferDelayStarted events raised by the RTCoin contract.
type RTCoinOwnerTransferDelayStartedIterator struct {
	Event *RTCoinOwnerTransferDelayStarted // Event containing the contract specifics and raw log

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
func (it *RTCoinOwnerTransferDelayStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinOwnerTransferDelayStarted)
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
		it.Event = new(RTCoinOwnerTransferDelayStarted)
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
func (it *RTCoinOwnerTransferDelayStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinOwnerTransferDelayStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinOwnerTransferDelayStarted represents a OwnerTransferDelayStarted event raised by the RTCoin contract.
type RTCoinOwnerTransferDelayStarted struct {
	PreviousOwner   common.Address
	NewOwner        common.Address
	ActivationBlock *big.Int
	ActivationTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferDelayStarted is a free log retrieval operation binding the contract event 0x0fd69fbe3cf6151621edd8a8aa0b65868e1cb2d9515ce4038554dca2802265dd.
//
// Solidity: e OwnerTransferDelayStarted(_previousOwner address, _newOwner address, _activationBlock uint256, _activationTime uint256)
func (_RTCoin *RTCoinFilterer) FilterOwnerTransferDelayStarted(opts *bind.FilterOpts) (*RTCoinOwnerTransferDelayStartedIterator, error) {

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "OwnerTransferDelayStarted")
	if err != nil {
		return nil, err
	}
	return &RTCoinOwnerTransferDelayStartedIterator{contract: _RTCoin.contract, event: "OwnerTransferDelayStarted", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferDelayStarted is a free log subscription operation binding the contract event 0x0fd69fbe3cf6151621edd8a8aa0b65868e1cb2d9515ce4038554dca2802265dd.
//
// Solidity: e OwnerTransferDelayStarted(_previousOwner address, _newOwner address, _activationBlock uint256, _activationTime uint256)
func (_RTCoin *RTCoinFilterer) WatchOwnerTransferDelayStarted(opts *bind.WatchOpts, sink chan<- *RTCoinOwnerTransferDelayStarted) (event.Subscription, error) {

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "OwnerTransferDelayStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinOwnerTransferDelayStarted)
				if err := _RTCoin.contract.UnpackLog(event, "OwnerTransferDelayStarted", log); err != nil {
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

// RTCoinOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RTCoin contract.
type RTCoinOwnershipTransferredIterator struct {
	Event *RTCoinOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RTCoinOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinOwnershipTransferred)
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
		it.Event = new(RTCoinOwnershipTransferred)
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
func (it *RTCoinOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinOwnershipTransferred represents a OwnershipTransferred event raised by the RTCoin contract.
type RTCoinOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_previousOwner address, _newOwner address)
func (_RTCoin *RTCoinFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts) (*RTCoinOwnershipTransferredIterator, error) {

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &RTCoinOwnershipTransferredIterator{contract: _RTCoin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_previousOwner address, _newOwner address)
func (_RTCoin *RTCoinFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RTCoinOwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinOwnershipTransferred)
				if err := _RTCoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RTCoinStakeContractSetIterator is returned from FilterStakeContractSet and is used to iterate over the raw logs and unpacked data for StakeContractSet events raised by the RTCoin contract.
type RTCoinStakeContractSetIterator struct {
	Event *RTCoinStakeContractSet // Event containing the contract specifics and raw log

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
func (it *RTCoinStakeContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinStakeContractSet)
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
		it.Event = new(RTCoinStakeContractSet)
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
func (it *RTCoinStakeContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinStakeContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinStakeContractSet represents a StakeContractSet event raised by the RTCoin contract.
type RTCoinStakeContractSet struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeContractSet is a free log retrieval operation binding the contract event 0xcf229ad20569d02c4a6cd3b3ae6130cb9e6257558e22a670804b6a6eb866b714.
//
// Solidity: e StakeContractSet(_contractAddress address)
func (_RTCoin *RTCoinFilterer) FilterStakeContractSet(opts *bind.FilterOpts) (*RTCoinStakeContractSetIterator, error) {

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "StakeContractSet")
	if err != nil {
		return nil, err
	}
	return &RTCoinStakeContractSetIterator{contract: _RTCoin.contract, event: "StakeContractSet", logs: logs, sub: sub}, nil
}

// WatchStakeContractSet is a free log subscription operation binding the contract event 0xcf229ad20569d02c4a6cd3b3ae6130cb9e6257558e22a670804b6a6eb866b714.
//
// Solidity: e StakeContractSet(_contractAddress address)
func (_RTCoin *RTCoinFilterer) WatchStakeContractSet(opts *bind.WatchOpts, sink chan<- *RTCoinStakeContractSet) (event.Subscription, error) {

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "StakeContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinStakeContractSet)
				if err := _RTCoin.contract.UnpackLog(event, "StakeContractSet", log); err != nil {
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

// RTCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RTCoin contract.
type RTCoinTransferIterator struct {
	Event *RTCoinTransfer // Event containing the contract specifics and raw log

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
func (it *RTCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinTransfer)
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
		it.Event = new(RTCoinTransfer)
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
func (it *RTCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinTransfer represents a Transfer event raised by the RTCoin contract.
type RTCoinTransfer struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_sender indexed address, _recipient indexed address, _amount uint256)
func (_RTCoin *RTCoinFilterer) FilterTransfer(opts *bind.FilterOpts, _sender []common.Address, _recipient []common.Address) (*RTCoinTransferIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "Transfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return &RTCoinTransferIterator{contract: _RTCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_sender indexed address, _recipient indexed address, _amount uint256)
func (_RTCoin *RTCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RTCoinTransfer, _sender []common.Address, _recipient []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "Transfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinTransfer)
				if err := _RTCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// RTCoinTransfersFrozenIterator is returned from FilterTransfersFrozen and is used to iterate over the raw logs and unpacked data for TransfersFrozen events raised by the RTCoin contract.
type RTCoinTransfersFrozenIterator struct {
	Event *RTCoinTransfersFrozen // Event containing the contract specifics and raw log

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
func (it *RTCoinTransfersFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinTransfersFrozen)
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
		it.Event = new(RTCoinTransfersFrozen)
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
func (it *RTCoinTransfersFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinTransfersFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinTransfersFrozen represents a TransfersFrozen event raised by the RTCoin contract.
type RTCoinTransfersFrozen struct {
	TransfersFrozen bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransfersFrozen is a free log retrieval operation binding the contract event 0xff7ea91c52ebd8c0d8018fdba50cb801e862f6795b1e17eeac882d4288b09340.
//
// Solidity: e TransfersFrozen(_transfersFrozen indexed bool)
func (_RTCoin *RTCoinFilterer) FilterTransfersFrozen(opts *bind.FilterOpts, _transfersFrozen []bool) (*RTCoinTransfersFrozenIterator, error) {

	var _transfersFrozenRule []interface{}
	for _, _transfersFrozenItem := range _transfersFrozen {
		_transfersFrozenRule = append(_transfersFrozenRule, _transfersFrozenItem)
	}

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "TransfersFrozen", _transfersFrozenRule)
	if err != nil {
		return nil, err
	}
	return &RTCoinTransfersFrozenIterator{contract: _RTCoin.contract, event: "TransfersFrozen", logs: logs, sub: sub}, nil
}

// WatchTransfersFrozen is a free log subscription operation binding the contract event 0xff7ea91c52ebd8c0d8018fdba50cb801e862f6795b1e17eeac882d4288b09340.
//
// Solidity: e TransfersFrozen(_transfersFrozen indexed bool)
func (_RTCoin *RTCoinFilterer) WatchTransfersFrozen(opts *bind.WatchOpts, sink chan<- *RTCoinTransfersFrozen, _transfersFrozen []bool) (event.Subscription, error) {

	var _transfersFrozenRule []interface{}
	for _, _transfersFrozenItem := range _transfersFrozen {
		_transfersFrozenRule = append(_transfersFrozenRule, _transfersFrozenItem)
	}

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "TransfersFrozen", _transfersFrozenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinTransfersFrozen)
				if err := _RTCoin.contract.UnpackLog(event, "TransfersFrozen", log); err != nil {
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

// RTCoinTransfersThawedIterator is returned from FilterTransfersThawed and is used to iterate over the raw logs and unpacked data for TransfersThawed events raised by the RTCoin contract.
type RTCoinTransfersThawedIterator struct {
	Event *RTCoinTransfersThawed // Event containing the contract specifics and raw log

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
func (it *RTCoinTransfersThawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinTransfersThawed)
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
		it.Event = new(RTCoinTransfersThawed)
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
func (it *RTCoinTransfersThawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinTransfersThawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinTransfersThawed represents a TransfersThawed event raised by the RTCoin contract.
type RTCoinTransfersThawed struct {
	TransfersThawed bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransfersThawed is a free log retrieval operation binding the contract event 0xb36ea4d45a6246e5ea6da988f57a5bf9a9022c85940cc6fe92dd9e45bf632cf6.
//
// Solidity: e TransfersThawed(_transfersThawed indexed bool)
func (_RTCoin *RTCoinFilterer) FilterTransfersThawed(opts *bind.FilterOpts, _transfersThawed []bool) (*RTCoinTransfersThawedIterator, error) {

	var _transfersThawedRule []interface{}
	for _, _transfersThawedItem := range _transfersThawed {
		_transfersThawedRule = append(_transfersThawedRule, _transfersThawedItem)
	}

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "TransfersThawed", _transfersThawedRule)
	if err != nil {
		return nil, err
	}
	return &RTCoinTransfersThawedIterator{contract: _RTCoin.contract, event: "TransfersThawed", logs: logs, sub: sub}, nil
}

// WatchTransfersThawed is a free log subscription operation binding the contract event 0xb36ea4d45a6246e5ea6da988f57a5bf9a9022c85940cc6fe92dd9e45bf632cf6.
//
// Solidity: e TransfersThawed(_transfersThawed indexed bool)
func (_RTCoin *RTCoinFilterer) WatchTransfersThawed(opts *bind.WatchOpts, sink chan<- *RTCoinTransfersThawed, _transfersThawed []bool) (event.Subscription, error) {

	var _transfersThawedRule []interface{}
	for _, _transfersThawedItem := range _transfersThawed {
		_transfersThawedRule = append(_transfersThawedRule, _transfersThawedItem)
	}

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "TransfersThawed", _transfersThawedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinTransfersThawed)
				if err := _RTCoin.contract.UnpackLog(event, "TransfersThawed", log); err != nil {
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
