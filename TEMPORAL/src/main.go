package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	signer "./signer"
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
		log.Fatalf("Error dialing client %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), keyPass)
	if err != nil {
		log.Fatalf("Error unlocking account %v", err)
	}
	paymentsADDR, tx, contract, err := DeployPayments(auth, client)
	if err != nil {
		log.Fatalf("Error creating contract deployment transaction %v", err)
	}

	_, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("Error waiting for transaction to be mined %v", err)
	}

	sAddr, err := contract.SIGNER(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve signer address %v", err)
	}

	if sAddr != auth.From {
		log.Fatal("signer addr is not equal to the current account")
	}

	fmt.Println("Payments contract address", paymentsADDR.String())
	// read our key
	ps, err := signer.GeneratePaymentSigner(keyFile, keyPass)
	if err != nil {
		log.Fatalf("Error generating payment signer %v", err)
	}
	method := uint8(0)
	number := big.NewInt(0)
	amount := big.NewInt(0)
	msg, err := ps.GenerateSignedPaymentMessageNoPrefix(auth.From, method, number, amount)
	if err != nil {
		log.Fatalf("Error signing payment message %v", err)
	}

	fmt.Printf("%+v\n", msg)
}
