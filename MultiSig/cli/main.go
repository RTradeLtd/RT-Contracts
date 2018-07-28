package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var keyPass = "password123"
var key = `{"address":"7e4a2359c745a982a54653128085eac69e446de1","crypto":{"cipher":"aes-128-ctr","ciphertext":"eea2004c17292a9e94217bf53efbc31ff4ae62f3dd57f0938ab61c949a565dc1","cipherparams":{"iv":"6f6a7a89b556604940ac87ab1e78cfd1"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8088e943ac0f37c8b4d01592d8bee96468853b6f1f13ca64d201cd68e7dc7b12"},"mac":"f856d734705f35e2acf854a44eb40796518730bd835ecaec01d1f3e7a7037813"},"id":"99e2cd49-4b51-4f01-b34c-aaa0efd332c3","version":3}`
var ipcFile = "/home/solidity/DevChain/node1/geth.ipc"
var owner1 = "0xf806f47aacdd85d10e406e4c740fe6d08cb79e2a"
var owner2 = "0x6958551613fa29ee9044794a9fdf6a1013c65860"
var ownerArray = []common.Address{common.HexToAddress(owner1), common.HexToAddress(owner2)}
var required = uint8(2)

func main() {
	client, err := ethclient.Dial(ipcFile)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), keyPass)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress, tx, _, err := DeployMultiSig(auth, client, ownerArray, required)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract address is ", contractAddress.String())
	address, err := bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}

	if address.String() != contractAddress.String() {
		log.Fatal("incorrect addresses")
	}
}
