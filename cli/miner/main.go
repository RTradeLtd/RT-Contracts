package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/RTradeLtd/RT-Contracts/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var endpoint = "http://127.0.0.1:8545"

func main() {
	epoint := os.Getenv("MAINNET_INFURA")
	if epoint == "" {
		endpoint = "http://127.0.0.1:8545"
	}
	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	if contractAddress == "" {
		log.Fatal("CONTRACT_ADDRESS environment variable is empty")
	}
	keyFile := os.Getenv("KEY_FILE")
	if keyFile == "" {
		log.Fatal("KEY_FILE environment variable is empty")
	}
	keyPass := os.Getenv("KEY_PASS")
	if keyPass == "" {
		log.Fatal("KEY_PASS environment variable is empty")
	}

	fileBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatal(err)
	}
	pk, err := keystore.DecryptKey(fileBytes, keyPass)
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(pk.PrivateKey)
	auth.GasPrice = big.NewInt(50000000000)

	miner, err := bindings.NewMergedMinerValidator(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	addr, err := miner.TOKENADDRESS(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Token Address that miner is using is ", addr.String())
}
