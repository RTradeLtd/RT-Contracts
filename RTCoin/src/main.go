package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
NOTE: merged mining validation won't work on PoA networks since they don't set the `mined` block header
*/
var keyFile = "/home/solidity/DevChain/node1/UTC--2018-07-10T00-50-38.032362728Z--7e4a2359c745a982a54653128085eac69e446de1"
var keyPass = "password123"
var ipcFile = "/home/solidity/DevChain/node1/geth.ipc"
var key = `{"address":"7e4a2359c745a982a54653128085eac69e446de1","crypto":{"cipher":"aes-128-ctr","ciphertext":"eea2004c17292a9e94217bf53efbc31ff4ae62f3dd57f0938ab61c949a565dc1","cipherparams":{"iv":"6f6a7a89b556604940ac87ab1e78cfd1"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8088e943ac0f37c8b4d01592d8bee96468853b6f1f13ca64d201cd68e7dc7b12"},"mac":"f856d734705f35e2acf854a44eb40796518730bd835ecaec01d1f3e7a7037813"},"id":"99e2cd49-4b51-4f01-b34c-aaa0efd332c3","version":3}`
var rtcAddress = "0xB8fe3B2C83014566733B766a27d94CB9AC167Dc6"
var datBig = new(big.Int).Mul(big.NewInt(100000000000000), big.NewInt(10000000000000))

func main() {
	client, err := ethclient.Dial(ipcFile)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), keyPass)
	if err != nil {
		log.Fatal(err)
	}
	var rtc *RTCoin
	if rtcAddress != "" {
		rtc, err = NewRTCoin(common.HexToAddress(rtcAddress), client)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("deploying vesting")
	vestingADDR, tx, vesting, err := DeployVesting(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("vesting address is", vestingADDR.String())
	_, err = vesting.TOKENADDRESS(nil)
	if err != nil {
		log.Fatal(err)
	}
	tx, err = rtc.Approve(auth, vestingADDR, datBig)
	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}

	dateOne := time.Now().Add(time.Minute * 1).Unix()
	dateTwo := time.Now().Add(time.Minute * 2).Unix()
	dateOneBig := big.NewInt(dateOne)
	dateTwoBig := big.NewInt(dateTwo)
	amountBig := big.NewInt(100000000000)
	totalBig := new(big.Int).Mul(amountBig, big.NewInt(2))
	tx, err = vesting.AddVest(auth, auth.From, totalBig, []*big.Int{dateOneBig, dateTwoBig}, []*big.Int{amountBig, amountBig})
	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	// sleep for 2 minutes, allowing enough time to pass to withdraw vested tokens
	time.Sleep(time.Minute * 3)
	tx, err = vesting.WithdrawVestedTokens(auth, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	vests, err := vesting.Vests(nil, auth.From)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", vests)
}

/* validator
func main() {
	client, err := ethclient.Dial(ipcFile)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), keyPass)
	if err != nil {
		log.Fatal(err)
	}
	var rtc *RTCoin
	if rtcAddress != "" {
		rtc, err = NewRTCoin(common.HexToAddress(rtcAddress), client)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("deploying merged miner validator")
	// deploy the validator
	validatorADDR, tx, validator, err := DeployValidator(auth, client)
	if err != nil {
		fmt.Println("failed to deploy merged miner validator")
		log.Fatal(err)
	}
	fmt.Println("validator address is", validatorADDR.String())
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("setting merged miner validator on RTC")
	tx, err = rtc.SetMergedMinerValidator(auth, validatorADDR)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		fmt.Println("error setting merged miner validator")
		log.Fatal(err)
	}
	addr, err := rtc.MergedMinerValidator(nil)
	if err != nil {
		log.Fatal(err)
	}
	if addr != validatorADDR {
		log.Fatal("set validator address on rtc does not match deployed contract addr")
	}
	fmt.Println("successfully set merged miner validator")
	amt := new(big.Int).Mul(big.NewInt(1000000000000000000), big.NewInt(10))
	bal, err := rtc.BalanceOf(nil, auth.From)
	if err != nil {
		log.Fatal("error determining balance", err)
	}
	fmt.Println("retrieved balance ", bal)
	fmt.Println("retrieved balance ", amt)
	fmt.Println("transfering funds to merged miner validator")
	// we need to deposit coins into the merged mining contract
	tx, err = rtc.Transfer(auth, validatorADDR, amt)
	if err != nil {
		fmt.Println("error sending transfer")
		log.Fatal(err)
	}
	fmt.Println("waiting for transfer transaction to be mined")
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		fmt.Println("error sending tokens")
		log.Fatal(err)
	}
	// now we need to approve the merged miner validator
	tx, err = rtc.Approve(auth, validatorADDR, amt)
	if err != nil {
		log.Fatal("error approving validator contract", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("error waiting for transaction to be mined", err)
	}
	fmt.Println("setting block")
	tx, err = validator.SubmitBlock(auth)
	if err != nil {
		log.Fatal("error setting block", err)
	}
	// wait for tx to be mined
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("error waiting for transaction to be mined", err)
	}

	blocks, err := validator.BlockNumberArray(nil, big.NewInt(0))
	if err != nil {
		log.Fatal("eror retrieving block number array", err)
	}
	fmt.Println("block number array first element", blocks)
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("error retrieving block by number", err)
	}
	fmt.Println("claiming funds")
	prevBlockNum := new(big.Int).Sub(block.Number(), big.NewInt(1))
	tx, err = validator.ClaimReward(auth, prevBlockNum)
	if err != nil {
		log.Fatal("error claiming reward", err)
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("error waiting for block to be mined", err)
	}

}
*/
/*  STAKING
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

	b := big.NewInt(1000000000000000000)
	bigWei := ConvertNumberToBaseWei(b)
	tx, err = rtc.Approve(auth, stakeADDR, bigWei)
	if err != nil {
		fmt.Println("failed to approve")
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.TODO(), client, tx)
	if err != nil {
		fmt.Println("error waiting for approve tx")
		log.Fatal(err)
	}

	tx, err = stake.DepositStake(auth, b)
	if err != nil {
		fmt.Println("failed to deposit stake")
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.TODO(), client, tx)
	if err != nil {
		log.Fatal(err)
	}

	s, err := stake.Stakes(nil, auth.From, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", s)
	time.Sleep(time.Minute * 1)
	fmt.Println("minting coins")
	tx, err = stake.Mint(auth, big.NewInt(0))
	if err != nil {
		fmt.Println("failed to withdraw stake")
		log.Fatal(err)
	}
	_, err = bind.WaitMined(context.TODO(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Minute * 2)
	fmt.Println("withdrawing initial stake")
	tx, err = stake.WithdrawInitialStake(auth, big.NewInt(0))
	if err != nil {
		fmt.Println("error withdrawing initial stake")
		log.Fatal(err)
	}
	fmt.Println("waiting for stake withdrawal transaction to be mined")
	_, err = bind.WaitMined(context.TODO(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
}
*/
// ConvertNumberToBaseWei is used to take a number, and multiply it by 10^18
func ConvertNumberToBaseWei(num *big.Int) *big.Int {
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	baseWei := new(big.Int).Mul(num, exp)
	return baseWei
}
