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

// RTCoinABI is the input ABI used to generate the binding from.
const RTCoinABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"freezeTransfers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stake\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setStakeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"transferOutEth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferForeignToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_mergedMinerValidator\",\"type\":\"address\"}],\"name\":\"setMergedMinerValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"thawTransfers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setFailOverStakeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIALSUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeFailOverRestrictionLifted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transfersFrozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mergedMinerValidatorAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"minters\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_transfersFrozen\",\"type\":\"bool\"}],\"name\":\"TransfersFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_transfersThawed\",\"type\":\"bool\"}],\"name\":\"TransfersThawed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ForeignTokenTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"EthTransferOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"MergedMinerValidatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"StakeContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"FailOverStakeContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_stakeContract\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_mintAmount\",\"type\":\"uint256\"}],\"name\":\"CoinsMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RTCoinBin is the compiled bytecode used for deploying new contracts.
const RTCoinBin = `60028054600160a060020a0319908116909155600380548216905560048054909116905560c0604052600660808190527f5254436f696e000000000000000000000000000000000000000000000000000060a090815262000064916005919062000159565b506040805180820190915260038082527f52544300000000000000000000000000000000000000000000000000000000006020909201918252620000ab9160069162000159565b506a32f44eb0f61c61240000006007556008805462ff00001961ff001960ff199092166012179190911661010017169055348015620000e957600080fd5b506000805433600160a060020a0319918216811783556001805490921681179091556007548183526009602090815260408085208390558051928352519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a3620001fe565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200019c57805160ff1916838001178555620001cc565b82800160010185558215620001cc579182015b82811115620001cc578251825591602001919060010190620001af565b50620001da929150620001de565b5090565b620001fb91905b80821115620001da5760008155600101620001e5565b90565b611af6806200020e6000396000f3006080604052600436106101875763ffffffff60e060020a60003504166301502460811461018c57806306fdde03146101b5578063095ea7b31461023f57806318160ddd1461026357806323b872dd1461028a578063272caf69146102b457806327e235e3146102e5578063313ce567146103065780633a4b66f11461033157806340c10f1914610346578063509484d51461036a5780635c6581651461038b57806366188463146103b2578063704b6c02146103d657806370a08231146103f75780638da5cb5b146104185780638f87c84b1461042d57806395d89b41146104425780639e5fea8a14610457578063a9059cbb14610481578063c0da7e69146104a5578063ce8e120a146104c6578063d73dd623146104db578063db6900fa146104ff578063dd62ed3e14610520578063de6ab39c14610547578063dfeb34b61461055c578063e45b813414610571578063e695751414610586578063f2fde38b1461059b578063f46eccc4146105bc578063f851a440146105dd578063ffa1ad74146105f2575b600080fd5b34801561019857600080fd5b506101a1610607565b604080519115158252519081900360200190f35b3480156101c157600080fd5b506101ca610677565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102045781810151838201526020016101ec565b50505050905090810190601f1680156102315780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024b57600080fd5b506101a1600160a060020a0360043516602435610705565b34801561026f57600080fd5b5061027861076b565b60408051918252519081900360200190f35b34801561029657600080fd5b506101a1600160a060020a0360043581169060243516604435610771565b3480156102c057600080fd5b506102c9610a66565b60408051600160a060020a039092168252519081900360200190f35b3480156102f157600080fd5b50610278600160a060020a0360043516610a75565b34801561031257600080fd5b5061031b610a87565b6040805160ff9092168252519081900360200190f35b34801561033d57600080fd5b506102c9610a90565b34801561035257600080fd5b506101a1600160a060020a0360043516602435610a9f565b34801561037657600080fd5b506101a1600160a060020a0360043516610bed565b34801561039757600080fd5b50610278600160a060020a0360043581169060243516610e48565b3480156103be57600080fd5b506101a1600160a060020a0360043516602435610e65565b3480156103e257600080fd5b506101a1600160a060020a0360043516610f54565b34801561040357600080fd5b50610278600160a060020a036004351661103d565b34801561042457600080fd5b506102c9611058565b34801561043957600080fd5b506101a1611067565b34801561044e57600080fd5b506101ca611108565b34801561046357600080fd5b506101a1600160a060020a0360043581169060243516604435611163565b34801561048d57600080fd5b506101a1600160a060020a036004351660243561139d565b3480156104b157600080fd5b506101a1600160a060020a036004351661158d565b3480156104d257600080fd5b506101a16116ab565b3480156104e757600080fd5b506101a1600160a060020a0360043516602435611717565b34801561050b57600080fd5b506101a1600160a060020a03600435166117b0565b34801561052c57600080fd5b50610278600160a060020a03600435811690602435166118e5565b34801561055357600080fd5b50610278611910565b34801561056857600080fd5b506101a161191f565b34801561057d57600080fd5b506101a161192e565b34801561059257600080fd5b506102c961193c565b3480156105a757600080fd5b506101a1600160a060020a036004351661194b565b3480156105c857600080fd5b506101a1600160a060020a0360043516611a21565b3480156105e957600080fd5b506102c9611a36565b3480156105fe57600080fd5b506101ca611a45565b60008054600160a060020a031633148061062b5750600154600160a060020a031633145b151561063657600080fd5b6008805461ff0019166101001790556040516001907fff7ea91c52ebd8c0d8018fdba50cb801e862f6795b1e17eeac882d4288b0934090600090a250600190565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106fd5780601f106106d2576101008083540402835291602001916106fd565b820191906000526020600020905b8154815290600101906020018083116106e057829003601f168201915b505050505081565b336000818152600a60209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60075490565b600854600090610100900460ff16156107d4576040805160e560020a62461bcd02815260206004820152601c60248201527f7472616e7366657273206d757374206e6f742062652066726f7a656e00000000604482015290519081900360640190fd5b82600160a060020a0381161515610823576040805160e560020a62461bcd0281526020600482015260186024820152600080516020611aab833981519152604482015290519081900360640190fd5b600160a060020a0385166000908152600960205260409020548311156108b9576040805160e560020a62461bcd02815260206004820152602160248201527f6f776e657220646f6573206e6f74206861766520656e6f75676820746f6b656e60448201527f7300000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0385166000908152600a6020908152604080832033845290915290205483111561095a576040805160e560020a62461bcd02815260206004820152602560248201527f73656e64657220646f6573206e6f74206861766520656e6f75676820616c6c6f60448201527f77616e6365000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0385166000908152600a6020908152604080832033845290915290205461098e908463ffffffff611a7c16565b600160a060020a0386166000818152600a60209081526040808320338452825280832094909455918152600990915220546109cf908463ffffffff611a7c16565b600160a060020a038087166000908152600960205260408082209390935590861681522054610a04908463ffffffff611a9116565b600160a060020a0380861660008181526009602090815260409182902094909455805187815290519193928916927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3506001949350505050565b600354600160a060020a031681565b60096020526000908152604090205481565b60085460ff1681565b600254600160a060020a031681565b336000908152600b602052604081205460ff161515600114610b0b576040805160e560020a62461bcd02815260206004820152601d60248201527f73656e646572206d75737420626520612076616c6964206d696e746572000000604482015290519081900360640190fd5b600160a060020a038316600090815260096020526040902054610b34908363ffffffff611a9116565b600160a060020a038416600090815260096020526040902055600754610b60908363ffffffff611a9116565b600755604080518381529051600160a060020a038516916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3604080518381529051600160a060020a0385169133917f601ace93afe864dd1288a16daad8ee79b21790c32c80d710b03cbae376e9e65f9181900360200190a350600192915050565b60008054600160a060020a03163314610c0557600080fd5b6000548290600160a060020a03808316911614801590610c335750600154600160a060020a03828116911614155b1515610c89576040805160e560020a62461bcd02815260206004820152601b60248201527f616464722063616e74206265206f776e6572206f722061646d696e0000000000604482015290519081900360640190fd5b600354600160a060020a031615610db857600260009054906101000a9004600160a060020a0316600160a060020a031663ed2f23696040518163ffffffff1660e060020a02815260040160206040518083038186803b158015610ceb57600080fd5b505afa158015610cff573d6000803e3d6000fd5b505050506040513d6020811015610d1557600080fd5b505115610db8576040805160e560020a62461bcd02815260206004820152604b60248201527f7374616b696e6720636f6e747261637420616c726561647920636f6e6669677560448201527f7265642c20746f206368616e6765206974206d7573742068617665203020616360648201527f74697665207374616b6573000000000000000000000000000000000000000000608482015290519081900360a40190fd5b60038054600160a060020a03851673ffffffffffffffffffffffffffffffffffffffff1991821681179092556000828152600b6020908152604091829020805460ff191660011790556002805490931684179092558051928352517fcf229ad20569d02c4a6cd3b3ae6130cb9e6257558e22a670804b6a6eb866b7149281900390910190a1600191505b50919050565b600a60209081526000928352604080842090915290825290205481565b336000908152600a60209081526040808320600160a060020a0386168452909152812054808310610eb957336000908152600a60209081526040808320600160a060020a0388168452909152812055610eee565b610ec9818463ffffffff611a7c16565b336000908152600a60209081526040808320600160a060020a03891684529091529020555b336000818152600a60209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b60008054600160a060020a03163314610f6c57600080fd5b81600160a060020a0381161515610fbb576040805160e560020a62461bcd0281526020600482015260186024820152600080516020611aab833981519152604482015290519081900360640190fd5b600154600160a060020a0384811691161415610fd657600080fd5b60018054600160a060020a03851673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c9181900360200190a150600192915050565b600160a060020a031660009081526009602052604090205490565b600054600160a060020a031681565b600080548190600160a060020a031633148061108d5750600154600160a060020a031633145b151561109857600080fd5b506040513080319133913180156108fc02916000818181858888f193505050501580156110c9573d6000803e3d6000fd5b5060408051828152905133917ffed66b098dae795e8a862bb1a0d1883d488f015acc2cf25cd29091efa8d8fb6b919081900360200190a2600191505090565b6006805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106fd5780601f106106d2576101008083540402835291602001916106fd565b600080548190600160a060020a03163314806111895750600154600160a060020a031633145b151561119457600080fd5b83600160a060020a03811615156111e3576040805160e560020a62461bcd0281526020600482015260186024820152600080516020611aab833981519152604482015290519081900360640190fd5b600160a060020a038616301415611269576040805160e560020a62461bcd028152602060048201526024808201527f746f6b656e20616464726573732063616e2774206265207468697320636f6e7460448201527f7261637400000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b85915081600160a060020a031663a9059cbb86866040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156112cf57600080fd5b505af11580156112e3573d6000803e3d6000fd5b505050506040513d60208110156112f957600080fd5b50511515611351576040805160e560020a62461bcd02815260206004820152601560248201527f746f6b656e207472616e73666572206661696c65640000000000000000000000604482015290519081900360640190fd5b604080518581529051600160a060020a0387169133917f10a46ed575affad8e954ae27853b1f89c6da90d8c35f619fc640f8a21bcb78579181900360200190a350600195945050505050565b600854600090610100900460ff1615611400576040805160e560020a62461bcd02815260206004820152601c60248201527f7472616e7366657273206d757374206e6f742062652066726f7a656e00000000604482015290519081900360640190fd5b82600160a060020a038116151561144f576040805160e560020a62461bcd0281526020600482015260186024820152600080516020611aab833981519152604482015290519081900360640190fd5b336000908152600960205260409020548311156114dc576040805160e560020a62461bcd02815260206004820152602260248201527f73656e64657220646f6573206e6f74206861766520656e6f75676820746f6b6560448201527f6e73000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b336000908152600960205260409020546114fc908463ffffffff611a7c16565b3360009081526009602052604080822092909255600160a060020a0386168152205461152e908463ffffffff611a9116565b600160a060020a0385166000818152600960209081526040918290209390935580518681529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35060019392505050565b60008054600160a060020a031633146115a557600080fd5b6000548290600160a060020a038083169116148015906115d35750600154600160a060020a03828116911614155b1515611629576040805160e560020a62461bcd02815260206004820152601b60248201527f616464722063616e74206265206f776e6572206f722061646d696e0000000000604482015290519081900360640190fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385169081179091556000818152600b6020908152604091829020805460ff19166001179055815192835290517f09eeb152b2546a9b79b2896b1b207bd9d9e94f00c0fad15b068e84478511bd529281900390910190a150600192915050565b60008054600160a060020a03163314806116cf5750600154600160a060020a031633145b15156116da57600080fd5b6008805461ff00191690556040516001907fb36ea4d45a6246e5ea6da988f57a5bf9a9022c85940cc6fe92dd9e45bf632cf690600090a250600190565b336000908152600a60209081526040808320600160a060020a038616845290915281205461174b908363ffffffff611a9116565b336000818152600a60209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b60008054600160a060020a031633146117c857600080fd5b6000548290600160a060020a038083169116148015906117f65750600154600160a060020a03828116911614155b151561184c576040805160e560020a62461bcd02815260206004820152601b60248201527f616464722063616e74206265206f776e6572206f722061646d696e0000000000604482015290519081900360640190fd5b60085462010000900460ff161515611878576008805462ff000019166201000017905560019150610e42565b600160a060020a0383166000818152600b6020908152604091829020805460ff191660011790556008805462ff000019169055815192835290517f540af0fc125e4047c03435dd52febc08726667f13f9c4ac5e8795a451be52f8b9281900390910190a160019150610e42565b600160a060020a039182166000908152600a6020908152604080832093909416825291909152205490565b6a32f44eb0f61c612400000081565b60085462010000900460ff1681565b600854610100900460ff1681565b600454600160a060020a031681565b60008054600160a060020a0316331461196357600080fd5b81600160a060020a03811615156119b2576040805160e560020a62461bcd0281526020600482015260186024820152600080516020611aab833981519152604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03851690811790915560408051338152602081019290925280517f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09281900390910190a150600192915050565b600b6020526000908152604090205460ff1681565b600154600160a060020a031681565b60408051808201909152600a81527f70726f64756374696f6e00000000000000000000000000000000000000000000602082015281565b600082821115611a8b57600080fd5b50900390565b600082820183811015611aa357600080fd5b939250505056006d757374206265206e6f6e207a65726f20616464726573730000000000000000a165627a7a723058208f2f7d91b7c745255cce3442185ec32f0f3d35ba926c0ea86b94f4c13f0204290029`

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

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_RTCoin *RTCoinCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_RTCoin *RTCoinSession) VERSION() (string, error) {
	return _RTCoin.Contract.VERSION(&_RTCoin.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_RTCoin *RTCoinCallerSession) VERSION() (string, error) {
	return _RTCoin.Contract.VERSION(&_RTCoin.CallOpts)
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
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_RTCoin *RTCoinTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_RTCoin *RTCoinSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.Approve(&_RTCoin.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_RTCoin *RTCoinTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.Approve(&_RTCoin.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_RTCoin *RTCoinTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_RTCoin *RTCoinSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.DecreaseApproval(&_RTCoin.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_RTCoin *RTCoinTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.DecreaseApproval(&_RTCoin.TransactOpts, _spender, _subtractedValue)
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

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_RTCoin *RTCoinTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_RTCoin *RTCoinSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.IncreaseApproval(&_RTCoin.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_RTCoin *RTCoinTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.IncreaseApproval(&_RTCoin.TransactOpts, _spender, _addedValue)
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
