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

var key = "..."
var keyPass = "password123"
var endpoint = "..."
var owner1 = "0xf806f47aacdd85d10e406e4c740fe6d08cb79e2a"
var owner2 = "0x6958551613fa29ee9044794a9fdf6a1013c65860"
var ownerArray = []common.Address{common.HexToAddress(owner1), common.HexToAddress(owner2)}
var required = uint8(2)

func main() {
	client, err := ethclient.Dial(endpoint)
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
