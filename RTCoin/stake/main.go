package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var keyFile = "/home/solidity/DevChainPOW/node1/UTC--2018-07-10T00-50-38.032362728Z--7e4a2359c745a982a54653128085eac69e446de1"
var keyPass = "password123"
var ipcFile = "/home/solidity/DevChain/node1/geth.ipc"
var key = `{"address":"7e4a2359c745a982a54653128085eac69e446de1","crypto":{"cipher":"aes-128-ctr","ciphertext":"eea2004c17292a9e94217bf53efbc31ff4ae62f3dd57f0938ab61c949a565dc1","cipherparams":{"iv":"6f6a7a89b556604940ac87ab1e78cfd1"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8088e943ac0f37c8b4d01592d8bee96468853b6f1f13ca64d201cd68e7dc7b12"},"mac":"f856d734705f35e2acf854a44eb40796518730bd835ecaec01d1f3e7a7037813"},"id":"99e2cd49-4b51-4f01-b34c-aaa0efd332c3","version":3}`
var rtcAddress = "0xE9AEc23c620681a59e2111785b0D35a90498128f"
var stakeAddress = "0x321d5bc70eF0AA22D1E33e4328cfb827858F8b21"

var deployStake = true
var deployRTC = true

func main() {
	client, err := ethclient.Dial(ipcFile)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), keyPass)
	if err != nil {
		log.Fatal(err)
	}

	if deployRTC {
		_, tx, _, err := DeployRTCoin(auth, client)
		if err != nil {
			log.Fatal(err)
		}

		addr, err := bind.WaitDeployed(context.Background(), client, tx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("RTC Contract Address is", addr.String())
		rtcAddress = addr.String()
	}

	if deployStake {
		_, tx, _, err := DeployStake(auth, client)
		if err != nil {
			log.Fatal(err)
		}
		addr, err := bind.WaitDeployed(context.Background(), client, tx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Stake Contract Address is", addr.String())
		stakeAddress = addr.String()
	}

	rtc, err := NewRTCoin(common.HexToAddress(rtcAddress), client)
	if err != nil {
		log.Fatal("failed to create rtc handler", err)
	}

	// set stake
	tx, err := rtc.SetStakeContract(auth, common.HexToAddress(stakeAddress))
	if err != nil {
		log.Fatal("failed to set stake contract", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined", err)
	}

	// get the stake address
	stakeContract, err := rtc.StakeContractAddress(nil)
	if err != nil {
		log.Fatal("failed to get stake contract address", err)
	}
	if stakeContract.String() != stakeAddress {
		log.Fatal("stake contract is not properly configured, it is set to ", stakeContract.String())
	}

	fmt.Println("Stake Contract Set", stakeContract.String())
	fmt.Println("stake address on file", stakeAddress)

	// create stake handler
	stake, err := NewStake(common.HexToAddress(stakeAddress), client)
	if err != nil {
		log.Fatal("failed to create stake contract", err)
	}

	totalSupply, err := rtc.TotalSupply(nil)
	if err != nil {
		log.Fatal("failed to get ")
	}

	// approve the stake contract
	tx, err = rtc.Approve(auth, common.HexToAddress(stakeAddress), totalSupply)
	if err != nil {
		log.Fatal("failed to approve ", err)
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined", err)
	}

	// enable token transfers
	tx, err = rtc.ThawTransfers(auth)
	if err != nil {
		log.Fatal("failed to enable token transfers ", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for transaction to be mined ", err)
	}

	transfersFrozen, err := rtc.TransfersFrozen(nil)
	if err != nil {
		log.Fatal("failed to get tx frozen status", err)
	}

	if transfersFrozen {
		log.Fatal("transfers still frozen")
	}

	fmt.Println("transfer frozen status ", transfersFrozen)

	num := new(big.Int).Mul(big.NewInt(1000000000000000000), big.NewInt(100))
	// attempt a transfer
	tx, err = rtc.Transfer(auth, common.HexToAddress(stakeAddress), num)
	if err != nil {
		log.Fatal("failed to transfer tokens ", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx mined ", err)
	}

	// set the RTC address on the stake contract
	tx, err = stake.SetRTI(auth, common.HexToAddress(rtcAddress))
	if err != nil {
		log.Fatal("failed to set RTC address ", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined ", err)
	}

	// make sure we can mint
	canMint, err := stake.CanMint(nil)
	if err != nil {
		log.Fatal("failed to get can mint status ", err)
	}
	if !canMint {
		log.Fatal("can't mint")
	}

	fmt.Println("mint status ", canMint)

	// enable staking
	tx, err = stake.AllowNewStakes(auth)
	if err != nil {
		log.Fatal("failed to enable new stakes ", err)
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined ", err)
	}

	stakesAllowed, err := stake.NewStakesAllowed(nil)
	if err != nil {
		log.Fatal("failed to check if stakes are allowed ", err)
	}
	if !stakesAllowed {
		log.Fatal("failed to enable new stakes")
	}
	fmt.Println("stake enabled status ", stakesAllowed)

	// deposit stakes
	auth.GasLimit = 350000
	tx, err = stake.DepositStake(auth, num)
	if err != nil {
		log.Fatal("failed to deposit stake ", err)
	}
	rcpt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined ", err)
	}
	if len(rcpt.Logs) == 0 {
		log.Fatal("no event logs detected")
	}

	fmt.Println("Deposit stake gas used: ", rcpt.CumulativeGasUsed)
	fmt.Println("waiting for time  and blocks to pass")
	time.Sleep(time.Minute * 3)

	prevWithdrawBalance, err := rtc.BalanceOf(nil, auth.From)
	if err != nil {
		log.Fatal("failed to get balance of ", err)
	}
	tx, err = stake.WithdrawInitialStake(auth, big.NewInt(0))
	if err != nil {
		log.Fatal("failed to withdraw initial stake")
	}
	rcpt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined ", err)
	}
	if len(rcpt.Logs) == 0 {
		log.Fatal("failed to generate event logs")
	}
	fmt.Println("withdraw initial stake gas used ", rcpt.CumulativeGasUsed)
	postWithdrawBalance, err := rtc.BalanceOf(nil, auth.From)
	if err != nil {
		log.Fatal("failed to get balance of ", err)
	}
	fmt.Println("prev initial stake withdraw balance ", prevWithdrawBalance)
	fmt.Println("post initial withdraw       balance ", postWithdrawBalance)
	if postWithdrawBalance.Cmp(prevWithdrawBalance) == 0 {
		log.Fatal("initial stake withdrawal failed")
	}
	stakeStatus, err := stake.Stakes(nil, auth.From, big.NewInt(0))
	if err != nil {
		log.Fatal("failed to get stak estatus ", err)
	}

	if stakeStatus.State != uint8(2) {
		log.Fatal("failed to close out initial stake and withdraw initial stake")
	}

	fmt.Printf("stake successfully closed\n%+v\n", stakeStatus)

	tx, err = stake.Mint(auth, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}

	rcpt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mint gas used ", rcpt.CumulativeGasUsed)
	stakeStatus, err = stake.Stakes(nil, auth.From, big.NewInt(0))
	if err != nil {
		log.Fatal("failed to get stak estatus ", err)
	}

	fmt.Printf("%+v\n", stakeStatus)

	totalSupplyNew, err := rtc.TotalSupply(nil)
	if err != nil {
		log.Fatal("failed to get total supply")
	}
	if totalSupplyNew.Cmp(totalSupply) == 0 {
		log.Fatal("minting process failed to increase token supply")
	}
	fmt.Println("new total supply", totalSupplyNew)
}
