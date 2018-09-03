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

// Miner is our miner service
type Miner struct {
	Client   *ethclient.Client
	Auth     *bind.TransactOpts
	Contract *bindings.MergedMinerValidator
}

func main() {
	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		log.Fatal("CONFIG_PATH environment variable is empty")
	}
	cfg, err := LoadConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(cfg.Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	fileBytes, err := ioutil.ReadFile(cfg.KeyFilePath)
	if err != nil {
		log.Fatal(err)
	}
	pk, err := keystore.DecryptKey(fileBytes, cfg.KeyFilePass)
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(pk.PrivateKey)
	contract, err := bindings.NewMergedMinerValidator(common.HexToAddress(cfg.ContractAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	miner := Miner{
		Client:   client,
		Auth:     auth,
		Contract: contract,
	}
	err = miner.Mine()
	if err != nil {
		log.Fatal(err)
	}
}

// Mine is used to start our miner process
func (m *Miner) Mine() error {
	totalMined := float64(0)
	for {
		gasPrice := big.NewInt(25000000000)
		fmt.Println("setting gas price")
		m.Auth.GasPrice = gasPrice
		fmt.Println("submitting block to contract")
		tx, err := m.Contract.SubmitBlock(m.Auth)
		if err != nil {
			fmt.Println("failed to submit transaction to blockchain ", err.Error())
			fmt.Println("sleeping for 30 seconds")
			time.Sleep(time.Second * 30)
			continue
		}
		fmt.Println("waiting for transaction to be mined")
		rcpt, err := bind.WaitMined(context.Background(), m.Client, tx)
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
		totalMined = totalMined + 0.5
		fmt.Println("congrats! you managed to submit block information")
		fmt.Println("total rtc mined this session ", totalMined)
		fmt.Println("sleeping for 5 seconds before continuing")
		time.Sleep(time.Second * 5)
	}
}
