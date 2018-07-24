package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
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
var deployRTC = true
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
		rtcAddress = addr.String()
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
		log.Fatal("failed to generate rtc contract handler ", err)
	}

	miner, err := NewMergedMinerValidator(common.HexToAddress(mmAddress), client)
	if err != nil {
		log.Fatal("failed to generate merged miner contract handler ", err)
	}

	totalSupply, err := rtc.TotalSupply(nil)
	if err != nil {
		log.Fatal("failed to get total supply ", err)
	}

	// enable transfers
	tx, err := rtc.ThawTransfers(auth)
	if err != nil {
		log.Fatal("failed to thaw transfers ", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined")
	}

	transferStatus, err := rtc.TransfersFrozen(nil)
	if err != nil {
		log.Fatal("failed to get transfer status ", err)
	}
	if transferStatus {
		log.Fatal("failed to enable transfer status")
	}
	fmt.Println("transactions frozen ", transferStatus)
	//set merged miner
	tx, err = rtc.SetMergedMinerValidator(auth, common.HexToAddress(mmAddress))
	if err != nil {
		log.Fatal("failed to set merged miner validator ", err)
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for tx to be mined ", err)
	}

	mAddr, err := rtc.MergedMinerValidatorAddress(nil)
	if err != nil {
		log.Fatal("failed to get merged mienr validator address ", err)
	}

	if mAddr.String() != mmAddress {
		log.Fatal("failed to set merged miner validator")
	}

	tx, err = miner.SetRTI(auth, common.HexToAddress(rtcAddress))
	if err != nil {
		log.Fatal("failed to set rtc token interface ", err)
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for transaction to be mined")
	}

	lastBlockSubmitted, err := miner.LastBlockSet(nil)
	if err != nil {
		log.Fatal("failed to get last block submitted ", err)
	}
	fmt.Println("last block submitted ", lastBlockSubmitted)

	intAr := []*big.Int{}
	intAr = append(intAr, lastBlockSubmitted)

	for i := 0; i < 9; i++ {
		tx, err = miner.SubmitBlock(auth)
		if err != nil {
			log.Fatal("failed to submit block information ", err)
		}

		rcpt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			log.Fatal("failed to watch for tx to be mined")
		}
		if len(rcpt.Logs) == 0 {
			log.Fatal("failed to emit any events")
		}
		fmt.Println("gas used ", rcpt.CumulativeGasUsed)
		lastBlock, err := miner.LastBlockSet(nil)
		if err != nil {
			log.Fatal("failed to get last block set", err)
		}
		intAr = append(intAr, lastBlock)
	}

	tx, err = miner.BulkClaimReward(auth, intAr)
	if err != nil {
		log.Fatal("failed to submit reward claim ", err)
	}
	rcpt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("failed to wait for transaction to be miend ", err)
	}
	fmt.Printf("reward claim receipt %+v\n", rcpt)
	fmt.Println("number of claims", len(intAr))
	newTotalSupply, err := rtc.TotalSupply(nil)
	if err != nil {
		log.Fatal("failed to get total supply ", err)
	}

	if newTotalSupply.Cmp(totalSupply) == 0 {
		log.Fatal("failed to increase total supply")
	}

	fmt.Println("new total supply ", newTotalSupply)
}
