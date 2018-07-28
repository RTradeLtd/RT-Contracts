package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var keyFile = "/home/solidity/DevChain/node1/UTC--2018-07-10T00-50-38.032362728Z--7e4a2359c745a982a54653128085eac69e446de1"
var keyPass = "password123"
var ipcFile = "/home/solidity/DevChain/node1/geth.ipc"
var key = `{"address":"7e4a2359c745a982a54653128085eac69e446de1","crypto":{"cipher":"aes-128-ctr","ciphertext":"eea2004c17292a9e94217bf53efbc31ff4ae62f3dd57f0938ab61c949a565dc1","cipherparams":{"iv":"6f6a7a89b556604940ac87ab1e78cfd1"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8088e943ac0f37c8b4d01592d8bee96468853b6f1f13ca64d201cd68e7dc7b12"},"mac":"f856d734705f35e2acf854a44eb40796518730bd835ecaec01d1f3e7a7037813"},"id":"99e2cd49-4b51-4f01-b34c-aaa0efd332c3","version":3}`
var rtcAddress = "0xE9AEc23c620681a59e2111785b0D35a90498128f"
var stakeAddress = "0x321d5bc70eF0AA22D1E33e4328cfb827858F8b21"

func main() {
	client, err := ethclient.Dial(ipcFile)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), keyPass)
	if err != nil {
		log.Fatal(err)
	}

	rtcAddr, tx, rtc, err := DeployRTCoin(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		log.Fatal("error waiting for tx to be mined ", err)
	}
	fmt.Println("RTC contract address is ", rtcAddr.String())

	tx, err = rtc.StartOwnerTransferDelay(auth, common.HexToAddress("0x9f11af43a0c4c959fa3f2489812cfd81e8cb1a9a"))
	if err != nil {
		log.Fatal("error starting owner transfer delay ", err)
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("error waiting for tx to be mined")
	}

	delay, err := rtc.Delay(nil)
	if err != nil {
		log.Fatal("error reading delay ", err)
	}
	fmt.Printf("%+v\n", delay)

	fmt.Println("sleeping for 100 seconds before activating transfer")
	oldOwner, err := rtc.Owner(nil)
	if err != nil {
		log.Fatal("error retrieving owner ", err)
	}
	time.Sleep(time.Second * 100)
	fmt.Println("activating transfer")
	auth.GasLimit = 750000
	tx, err = rtc.TransferOwnership(auth, common.HexToAddress("0x9f11af43a0c4c959fa3f2489812cfd81e8cb1a9a"))
	if err != nil {
		log.Fatal("error transferring owner ", err)
	}

	rcpt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("error waiting for tx to be mined ", err)
	}
	fmt.Printf("%+v\n", rcpt)
	if len(rcpt.Logs) == 0 {
		log.Fatal("failed to transfer ownership, no log events detected")
	}
	newOwner, err := rtc.Owner(nil)
	if err != nil {
		log.Fatal("error reading owner ", err)
	}

	if oldOwner.String() == newOwner.String() {
		log.Fatal("failed to transfer owner")
	}
}
