package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"./contracts"

	"github.com/DaveAppleton/etherUtils"
	"github.com/DaveAppleton/memorykeys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
)

var baseClient *backends.SimulatedBackend

func getClient() (client *backends.SimulatedBackend, err error) {
	if baseClient != nil {
		return baseClient, nil
	}
	funds, _ := etherUtils.StrToEther("10000.00")
	banker, _ := memorykeys.GetAddress("banker")
	baseClient = backends.NewSimulatedBackend(
		core.GenesisAlloc{
			*banker: {Balance: funds},
		}, 8000000)
	return baseClient, nil
}

func main() {
	banker, _ := memorykeys.GetAddress("banker")
	client, err := getClient()
	if err != nil {
		log.Fatal(err)
	}
	client.AdjustTime(1473046030 * time.Second)
	client.Commit()
	bal, _ := client.BalanceAt(context.Background(), *banker, nil)
	fmt.Println(etherUtils.EtherToStr(bal))
	bankerTx, _ := memorykeys.GetTransactor("banker")
	testContractAddress, tx, testContract, err := contracts.DeployTest(bankerTx, client)
	if err != nil {
		log.Fatal(err)
	}
	client.Commit()
	fmt.Println("Contract deployed at ", testContractAddress.Hex(), "in transaction ", tx.Hash().Hex())

	myVar, err := testContract.MyPublicVariable(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Old Value:", myVar.Uint64())
	tx, err = testContract.SetMyVariable(bankerTx, big.NewInt(42))
	if err != nil {
		log.Fatal(err)
	}
	tx, err = testContract.SetMyVariable(bankerTx, big.NewInt(55))
	if err != nil {
		log.Fatal(err)
	}
	client.Commit() // needed to commit old transactions before adjusting time.
	client.AdjustTime(25 * time.Hour)
	client.Commit()
	tx, err = testContract.SetMyVariable(bankerTx, big.NewInt(48))
	if err != nil {
		log.Fatal(err)
	}
	tx, err = testContract.SetMyVariable(bankerTx, big.NewInt(13))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SetVar ", tx.Hash().Hex())
	client.Commit()
	myChangedVar, err := testContract.MyPublicVariable(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New Value:", myChangedVar.Uint64())
	tf, err := contracts.NewTestFilterer(testContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	filt, err := tf.FilterValueSet(&bind.FilterOpts{})
	if err != nil {
		log.Fatal(err)
	}
	for filt.Next() {
		tm := time.Unix(int64(filt.Event.Date.Uint64()), 0)
		fmt.Println(etherUtils.CoinToStr(filt.Event.NewValue, 0), tm)
	}

}
