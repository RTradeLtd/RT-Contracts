package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"

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

	miner, err := bindings.NewMergedMinerValidator(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	err = mine(auth, client, miner)
	if err != nil {
		log.Fatal(err)
	}
}

func mine(auth *bind.TransactOpts, client *ethclient.Client, miner *bindings.MergedMinerValidator) error {
	rtcAwarded := float64(0)
	for {
		var err error
		var gasPrice *big.Int
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			fmt.Println("failed to get gas price due ", err.Error())
			fmt.Println("using default gas price of 25Gwei")
			gasPrice = big.NewInt(25000000000)
		}
		auth.GasPrice = gasPrice
		tx, err := miner.SubmitBlock(auth)
		if err != nil {
			fmt.Println("failed to submit transaction to blockchain ", err.Error())
			fmt.Println("sleeping for 30 seconds")
			time.Sleep(time.Second * 30)
			continue
		}
		rcpt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			fmt.Println("failed to wait for transaction to be mined ", err.Error())
			fmt.Println("sleeping for 30 seconds")
			time.Sleep(time.Second * 30)
			continue
		}
		if len(rcpt.Logs) == 0 {
			fmt.Println("no events emitted, transaction failed to execute properly")
			fmt.Println("sleeping for 30 seconds")
			time.Sleep(time.Second * 30)
			continue
		}
		if rcpt.Status != 1 {
			fmt.Println("transaction status is 0 indicating a failure, however events were emitted")
			fmt.Println("this is an unexpected failure, exiting")
			return errors.New("transaction status is 0 indicating a failure, however events were emitted which is unexpected")
		}
		rtcAwarded = rtcAwarded + rtcAwarded
		fmt.Println("congrats! you managed to submit block information")
		fmt.Println("total rtc mined this session ", rtcAwarded)
		fmt.Println("sleeping for 5 seconds before continuing")
		time.Sleep(time.Second * 5)
	}
}
