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
	totalSupply, err := rtc.TotalSupply(nil)
	if err != nil {
		log.Fatal("failed to get total supply", err)
	}

	// approve
	tx, err := rtc.Approve(auth, common.HexToAddress(stakeAddress), totalSupply)
	if err != nil {
		log.Fatal("failed to approve staking contract", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx", err)
	}

	// allowance
	allowance, err := rtc.Allowance(nil, auth.From, common.HexToAddress(stakeAddress))
	if err != nil {
		log.Fatal("failed to get allowance ", err)
	}
	if allowance.Cmp(totalSupply) != 0 {
		log.Fatal("failed to set a proper allownace")
	}

	// set stake
	tx, err = rtc.SetStakeContract(auth, common.HexToAddress(stakeAddress))
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

	// set the RTC address on the stake contract
	tx, err = stake.SetRTI(auth, common.HexToAddress(rtcAddress))
	if err != nil {
		log.Fatal("failed to set RTC address ", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined ", err)
	}

	auth.GasLimit = 275000
	// enable new stakes
	tx, err = stake.AllowNewStakes(auth)
	if err != nil {
		log.Fatal("failed to enable stakes", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined", err)
	}

	stakesAllowed, err := stake.NewStakesAllowed(nil)
	if err != nil {
		log.Fatal("failed to check if stakes allowed ", err)
	}

	if !stakesAllowed {
		log.Fatal("failed to enable new stakes")
	}

	// deposit stakes
	num := new(big.Int).Mul(big.NewInt(1000000000000000000), big.NewInt(3))
	tx, err = stake.DepositStake(auth, num)
	if err != nil {
		log.Fatal("failed to deposit stake", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined")
	}

	stakeStatus, err := stake.Stakes(nil, auth.From, big.NewInt(0))
	if err != nil {
		log.Fatal("failed to get stake status", err)
	}

	fmt.Printf("%+v\n", stakeStatus)

	// mint tokens
	tx, err = stake.Mint(auth, big.NewInt(0))
	if err != nil {
		log.Fatal("failed to mint", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined", err)
	}

	fmt.Println("waiting for blocks and time to pass")
	time.Sleep(time.Minute * 3)

	tx, err = stake.WithdrawInitialStake(auth, big.NewInt(0))
	if err != nil {
		log.Fatal("failed to withdraw initial stake", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined", err)
	}

	stakeStatus, err = stake.Stakes(nil, auth.From, big.NewInt(0))
	if err != nil {
		log.Fatal("failed to get stake status", err)
	}

	if stakeStatus.State != uint8(2) {
		fmt.Printf("%+v\n", stakeStatus)
		log.Fatal("failed to close out stake and withdraw initial funds")
	}

	if stakeStatus.TotalCoinsToMint.Cmp(stakeStatus.CoinsMinted) == 0 {
		log.Fatal("invalid stake detected, coins minted should  not be total coins minted yet")
	}

	tx, err = stake.Mint(auth, big.NewInt(0))
	if err != nil {
		log.Fatal("failed to mint", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined", err)
	}

	stakeStatus, err = stake.Stakes(nil, auth.From, big.NewInt(0))
	if err != nil {
		log.Fatal("failed to get stake status", err)
	}
	if stakeStatus.TotalCoinsToMint.Cmp(stakeStatus.CoinsMinted) != 0 {
		log.Fatal("invalid stake detected, coins minted should  equal total coins mint")
	}
	fmt.Printf("%+v\n", stakeStatus)
}
