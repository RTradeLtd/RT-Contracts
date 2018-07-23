package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var keyFile = "/home/solidity/DevChainPOW/node1/UTC--2018-07-10T00-50-38.032362728Z--7e4a2359c745a982a54653128085eac69e446de1"
var keyPass = "password123"
var ipcFile = "/home/solidity/DevChainPOW/node1/geth.ipc"
var key = `{"address":"7e4a2359c745a982a54653128085eac69e446de1","crypto":{"cipher":"aes-128-ctr","ciphertext":"eea2004c17292a9e94217bf53efbc31ff4ae62f3dd57f0938ab61c949a565dc1","cipherparams":{"iv":"6f6a7a89b556604940ac87ab1e78cfd1"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8088e943ac0f37c8b4d01592d8bee96468853b6f1f13ca64d201cd68e7dc7b12"},"mac":"f856d734705f35e2acf854a44eb40796518730bd835ecaec01d1f3e7a7037813"},"id":"99e2cd49-4b51-4f01-b34c-aaa0efd332c3","version":3}`
var rtcAddress = "0xE9AEc23c620681a59e2111785b0D35a90498128f"
var mmAddress = "0x7f7AF7AE6e2CE658398e8fb5337ad02a6578D6C8"
var deployRTC = false
var deployMined = true

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
		os.Exit(1)
	}

	if deployMined {
		_, tx, _, err := DeployMergedMinerValidator(auth, client)
		if err != nil {
			log.Fatal(err)
		}
		addr, err := bind.WaitDeployed(context.Background(), client, tx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Merged Miner Validator Contract Address is", addr.String())
		mmAddress = addr.String()
	}

	rtc, err := NewRTCoin(common.HexToAddress(rtcAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	miner, err := NewMergedMinerValidator(common.HexToAddress(mmAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("setting merged miner validator")
	tx, err := rtc.SetMergedMinerValidator(auth, common.HexToAddress(mmAddress))
	if err != nil {
		log.Fatal("failed to set merged miner validator", err)
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx mined", err)
	}

	mAddr, err := rtc.MergedMinerValidatorAddress(nil)
	if err != nil {
		log.Fatal("failed to retrieve merged miner validator", err)
	}

	if mAddr.String() != mmAddress {
		log.Fatal("failed to correctly set the merged miner validator")
	}

	tokenAddress, err := miner.TOKENADDRESS(nil)
	if err != nil {
		log.Fatal("failed to retrieve token address", err)
	}
	fmt.Println(tokenAddress)

	lastBlockSet, err := miner.LastBlockSet(nil)
	if err != nil {
		log.Fatal("failed to retrieve last block set", err)
	}

	intArr := []*big.Int{}
	intArr = append(intArr, big.NewInt(64))
	auth.GasLimit = 275000
	tx, err = miner.BulkClaimReward(auth, intArr)
	if err != nil {
		log.Fatal("faled to claim reward", err)
	}

	rcpt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be miend", err)
	}

	fmt.Printf("%+v\n", rcpt)

	auth.GasLimit = 275000
	tx, err = miner.SubmitBlock(auth)
	if err != nil {
		log.Fatal("unable to submit block", err)
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx mined", err)
	}

	newLastBlockSet, err := miner.LastBlockSet(nil)
	if err != nil {
		log.Fatal("failed to retrieve last block set", err)
	}

	if newLastBlockSet.Cmp(lastBlockSet) == 0 {
		fmt.Println(lastBlockSet)
		fmt.Println(newLastBlockSet)
		log.Fatal("failed to update last block set")
	}

	fmt.Println("succes")
}
