package main

import (
	"context"
	"errors"
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

var endpoint = ""

func main() {
	epoint := os.Getenv("RINKEBY_INFURA")
	if epoint == "" {
		endpoint = "http://127.0.0.1:8545"
	} else {
		endpoint = epoint
	}
	if len(os.Args) < 2 {
		fmt.Println("invalid invocation")
		fmt.Println("cli <mode> <args>")
		fmt.Println("mode: deploy")
		fmt.Println("deploy args: 1 = contract | 2 = keyFile | 3 = keyPass | 4 = multisig wallet")
		os.Exit(1)
	}
	mode := os.Args[1]
	switch mode {
	case "deploy":
		contract := os.Args[2]
		keyFile := os.Args[3]
		keyPass := os.Args[4]
		msigWalletAddress := os.Args[5]
		switch contract {
		case "rtc":
			deployRTC(keyFile, keyPass, msigWalletAddress)
		case "stake":
			deployStake(keyFile, keyPass, msigWalletAddress)
		case "rtceth":
			deployRTCETH(keyFile, keyPass, msigWalletAddress)
		case "vesting":
			deployVesting(keyFile, keyPass, msigWalletAddress)
		case "mergedminer":
			deployMergedMiner(keyFile, keyPass, msigWalletAddress)
		case "rtctest":
			deployRTCTest(keyFile, keyPass)
		case "mergedminertest":
			deployMergedMinerTest(keyFile, keyPass)
		default:
			log.Fatal("invalid contract type")
		}
	default:
		log.Fatal("invalid run mode")
	}
}

func deployMergedMiner(keyFile, keyPass, multisig string) error {
	fileBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return err
	}
	pk, err := keystore.DecryptKey(fileBytes, keyPass)
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(pk.PrivateKey)
	auth.GasPrice = big.NewInt(25000000000)
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		fmt.Println("error dialing ethclient ", err)
		return err
	}
	minerAddress, tx, _, err := bindings.DeployMergedMinerValidator(auth, client, common.HexToAddress(multisig))
	if err != nil {
		fmt.Println("error deploying merged miner validator ", err)
		return err
	}
	_, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		fmt.Println("failed to wait for merged miner validator to be deployed ", err)
		return err
	}

	fmt.Println("Merged Miner Validator Address ", minerAddress.String())
	return nil
}

func deployVesting(keyFile, keyPass, multisig string) error {
	fileBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return err
	}
	pk, err := keystore.DecryptKey(fileBytes, keyPass)
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(pk.PrivateKey)
	auth.GasPrice = big.NewInt(25000000000)
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		fmt.Println("error dialing ethclient ", err)
		return err
	}
	fmt.Println("deploying vesting contract")
	vestingAddr, tx, _, err := bindings.DeployVesting(auth, client, common.HexToAddress(multisig))
	if err != nil {
		fmt.Println("error deploying vesting contract ", err)
		return err
	}

	_, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		fmt.Println("failed to wait for contract deployment ", err)
		return err
	}
	fmt.Println("vesting contract address ", vestingAddr.String())
	return nil
}

func deployRTCETH(keyFile, keyPass, multisig string) error {
	fileBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return err
	}
	pk, err := keystore.DecryptKey(fileBytes, keyPass)
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(pk.PrivateKey)
	auth.GasPrice = big.NewInt(25000000000)
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		fmt.Println("error dialing ethclient ", err)
		return err
	}
	fmt.Println("deploying sale contract")
	saleAddr, tx, sale, err := bindings.DeployRTCETH(auth, client)
	if err != nil {
		fmt.Println("error deploying RTCETH ", err)
		return err
	}

	_, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		fmt.Println("error waiting for deployment ", err)
		return err
	}
	fmt.Println("transferring admin rights to postables")
	// transfer admin rights
	tx, err = sale.SetAdmin(auth, common.HexToAddress("0xc7459562777DDf3A1A7afefBE515E8479Bd3FDBD"))
	if err != nil {
		fmt.Println("failed to set admin rights ", err)
		return err
	}
	rcpt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		fmt.Println("failed to wait for transaction to be miend ", err)
		return err
	}
	if len(rcpt.Logs) == 0 {
		fmt.Println("failed to transfer admin rights logs emitted")
		return errors.New("failed to transfer admin rightss no logs emitted")
	}
	fmt.Println("transferring ownership to multisig")
	// transfer ownership
	tx, err = sale.TransferOwnership(auth, common.HexToAddress(multisig))
	if err != nil {
		fmt.Println("failed to transfer ownership ", err)
		return err
	}
	rcpt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		fmt.Println("failed to wait for transaction to be miend ", err)
		return err
	}
	if len(rcpt.Logs) == 0 {
		fmt.Println("failed to transfer ownership logs emitted")
		return errors.New("failed to transfer ownership no logs emitted")
	}
	fmt.Println("RTC-ETH Contract Address ", saleAddr.String())
	return nil
}

func deployStake(keyFile, keyPass, multisig string) error {
	fileBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return err
	}
	pk, err := keystore.DecryptKey(fileBytes, keyPass)
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(pk.PrivateKey)
	auth.GasPrice = big.NewInt(25000000000)
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		fmt.Println("error dialing ethclient ", err)
		return err
	}
	fmt.Println("deploying staking contract")
	_, tx, _, err := bindings.DeployStake(auth, client, common.HexToAddress(multisig))
	if err != nil {
		fmt.Println("failed to deploy staking contract ", err)
		return err
	}

	addr, err := bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		fmt.Println("error waiting for contract to be deployed ", err)
		return err
	}
	fmt.Println("staking contract deployed")
	fmt.Println("Staking Contract Address ", addr.String())
	return nil
}

func deployRTC(keyFile, keyPass, multisigWalletAddress string) error {
	fileBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return err
	}
	pk, err := keystore.DecryptKey(fileBytes, keyPass)
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(pk.PrivateKey)
	auth.GasPrice = big.NewInt(25000000000)
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		fmt.Println("error dialing ethclient ", err)
		return err
	}
	fmt.Println("deploying RTC token")
	rtcAddress, tx, rtc, err := bindings.DeployRTCoin(auth, client)
	if err != nil {
		fmt.Println("error deploying RTC token ", err)
		return err
	}

	_, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		fmt.Println("failed to wait for token to be deployed ", err)
		return err
	}
	fmt.Println("thawing transfers")
	// enable transfers
	tx, err = rtc.ThawTransfers(auth)
	if err != nil {
		fmt.Println("failed to submit tx thaw transaction ", err)
		return err
	}
	rcpt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		fmt.Println("failed to wait for transaction to be mined ", err)
		return err
	}
	if len(rcpt.Logs) == 0 {
		fmt.Println("failed to thaw transers no logs emitted")
		return errors.New("failed to thaw transfers, no logs emitted")
	}
	// get the total supply
	tSupply, err := rtc.TotalSupply(nil)
	if err != nil {
		fmt.Println("failed to retrieve total supply ", err)
		return err
	}
	fmt.Println("transferring tokens to multisig")
	// transfer tokens to multisig
	tx, err = rtc.Transfer(auth, common.HexToAddress(multisigWalletAddress), tSupply)
	if err != nil {
		fmt.Println("failed to submit transfer tx to blockchain ", err)
		return err
	}
	rcpt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		fmt.Println("failed to wait for transaction to be miend ", err)
		return err
	}
	if len(rcpt.Logs) == 0 {
		fmt.Println("failed to transfer tokens no logs emitted")
		return errors.New("failed to transfer tokens no logs emitted")
	}
	fmt.Println("transferring admin rights to multisig")
	// transfer admin rights
	tx, err = rtc.SetAdmin(auth, common.HexToAddress(multisigWalletAddress))
	if err != nil {
		fmt.Println("failed to set admin rights ", err)
		return err
	}
	rcpt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		fmt.Println("failed to wait for transaction to be miend ", err)
		return err
	}
	if len(rcpt.Logs) == 0 {
		fmt.Println("failed to transfer admin rights logs emitted")
		return errors.New("failed to transfer admin rightss no logs emitted")
	}
	fmt.Println("transferring ownership to multisig")
	// transfer ownership
	tx, err = rtc.TransferOwnership(auth, common.HexToAddress(multisigWalletAddress))
	if err != nil {
		fmt.Println("failed to transfer ownership ", err)
		return err
	}
	rcpt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		fmt.Println("failed to wait for transaction to be miend ", err)
		return err
	}
	if len(rcpt.Logs) == 0 {
		fmt.Println("failed to transfer ownership logs emitted")
		return errors.New("failed to transfer ownership no logs emitted")
	}
	fmt.Println("RTC token deployed and setup")
	fmt.Println("Token address: ", rtcAddress.String())
	return nil
}

func deployRTCTest(keyFile, keyPass string) error {
	fileBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return err
	}
	pk, err := keystore.DecryptKey(fileBytes, keyPass)
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(pk.PrivateKey)
	auth.GasPrice = big.NewInt(25000000000)
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		fmt.Println("error dialing ethclient ", err)
		return err
	}
	fmt.Println("deploying RTC token")
	rtcAddress, tx, rtc, err := bindings.DeployRTCoin(auth, client)
	if err != nil {
		fmt.Println("error deploying RTC token ", err)
		return err
	}

	_, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		fmt.Println("failed to wait for token to be deployed ", err)
		return err
	}
	fmt.Println("thawing transfers")
	// enable transfers
	tx, err = rtc.ThawTransfers(auth)
	if err != nil {
		fmt.Println("failed to submit tx thaw transaction ", err)
		return err
	}
	rcpt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		fmt.Println("failed to wait for transaction to be mined ", err)
		return err
	}
	if len(rcpt.Logs) == 0 {
		fmt.Println("failed to thaw transers no logs emitted")
		return errors.New("failed to thaw transfers, no logs emitted")
	}
	fmt.Println("RTC token deployed and setup")
	fmt.Println("Token address: ", rtcAddress.String())
	return nil
}

func deployMergedMinerTest(keyFile, keyPass string) error {
	fileBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return err
	}
	pk, err := keystore.DecryptKey(fileBytes, keyPass)
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(pk.PrivateKey)
	auth.GasPrice = big.NewInt(25000000000)
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		fmt.Println("error dialing ethclient ", err)
		return err
	}
	minerAddress, tx, _, err := bindings.DeployMergedMinerValidator(auth, client, auth.From)
	if err != nil {
		fmt.Println("error deploying merged miner validator ", err)
		return err
	}
	_, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		fmt.Println("failed to wait for merged miner validator to be deployed ", err)
		return err
	}

	fmt.Println("Merged Miner Validator Address ", minerAddress.String())
	return nil
}
