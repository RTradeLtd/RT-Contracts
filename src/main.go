package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var keyFile = "/home/solidity/DevChain/node1/UTC--2018-07-10T00-50-38.032362728Z--7e4a2359c745a982a54653128085eac69e446de1"
var keyPass = "password123"
var ipcFile = "/home/solidity/DevChain/node1/geth.ipc"
var key = `{"address":"7e4a2359c745a982a54653128085eac69e446de1","crypto":{"cipher":"aes-128-ctr","ciphertext":"eea2004c17292a9e94217bf53efbc31ff4ae62f3dd57f0938ab61c949a565dc1","cipherparams":{"iv":"6f6a7a89b556604940ac87ab1e78cfd1"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8088e943ac0f37c8b4d01592d8bee96468853b6f1f13ca64d201cd68e7dc7b12"},"mac":"f856d734705f35e2acf854a44eb40796518730bd835ecaec01d1f3e7a7037813"},"id":"99e2cd49-4b51-4f01-b34c-aaa0efd332c3","version":3}`

func main() {
	client, err := ethclient.Dial(ipcFile)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), keyPass)
	if err != nil {
		log.Fatal(err)
	}

	rtcADDR, tx, rtc, err := DeployRTCoin(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitDeployed(context.TODO(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	stakeADDR, tx, stake, err := DeployStake(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitDeployed(context.TODO(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rtcADDR.String())
	fmt.Println(stakeADDR.String())
	tx, err = rtc.SetStakeContract(auth, stakeADDR)
	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.TODO(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	addr, err := rtc.StakeContract(nil)
	if err != nil {
		log.Fatal(err)
	}
	if addr != stakeADDR {
		log.Fatal("recorded stake address is incorrect")
	}

	tx, err = stake.SetRTI(auth, rtcADDR)
	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.TODO(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	tx, err = stake.AllowNewStakes(auth)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bind.WaitMined(context.TODO(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	allowed, err := stake.NewStakesAllowed(nil)
	if err != nil {
		log.Fatal(err)
	}

	if !allowed {
		log.Fatal("stakes not allowed when they should be")
	}
	fmt.Println("stakes allowed yeboi ", allowed)
}
