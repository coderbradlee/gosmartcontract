package main

import (
	"math/big"
	// etclottery "../etclottery"
	// "context"
	// "crypto/ecdsa"
	"fmt"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "log"
	// "./hdwallet"
	// "crypto/ecdsa"
	// "encoding/hex"
	// "math/big"
	"testing"
	"time"
)

func TestDeployCall(t *testing.T) {
	ownerAuth := AuthAccountFromPrivateKey(userPrivateKey1)
	c := NewCallConnecterWithDeploy(selfhost, ownerAuth)

	fmt.Println("Call Contract address is: ", c.contractAddress.String())

}

func TestCallBuy(t *testing.T) {
	s := NewCallConnecter(selfhost, callcontractAddress)
	// s.Send(userPrivateKey1, callcontractAddress, big.NewInt(5000000000000000000))
	// time.Sleep(time.Second * 10)
	bla := s.BalanceOfEth(common.HexToAddress(callcontractAddress))
	fmt.Println(callcontractAddress, ":", bla)

	{
		ownerAuth1 := AuthAccountFromPrivateKey(userPrivateKey1)
		ret1, err := s.lottery.Setup(ownerAuth1, common.HexToAddress(callcontractAddress))
		if err != nil {
			fmt.Println("Setup:", err)
			// return
		}
		fmt.Println("Setup: ", ret1.Hash().Hex())
		time.Sleep(time.Second * 10)
	}
	{
		ownerAuth1 := AuthAccountFromPrivateKey(userPrivateKey1)
		ret1, err := s.lottery.CallByFun(ownerAuth1, common.HexToAddress(callcontractAddress))
		if err != nil {
			fmt.Println("CallByFun:", err)
			// return
		}
		fmt.Println("CallByFun: ", ret1.Hash().Hex())
	}
	{
		// ownerAuth1 := AuthAccountFromPrivateKey(userPrivateKey1)
		// ret1, err := s.lottery.DelegatecallByFun(ownerAuth1, common.HexToAddress(callcontractAddress))
		// if err != nil {
		// 	fmt.Println("DelegatecallByFun:", err)
		// 	// return
		// }
		// fmt.Println("DelegatecallByFun: ", ret1.Hash().Hex())
	}
	// {
	// 	ownerAuth1 := AuthAccountFromPrivateKey(userPrivateKey1)
	// 	s.Buy(ownerAuth1, 0)
	// 	time.Sleep(time.Second * 10)
	// 	s.Buy(ownerAuth1, 1)
	// 	time.Sleep(time.Second * 10)
	// }
	// {
	// 	ownerAuth1 := AuthAccountFromPrivateKey(userPrivateKey2)
	// 	s.Buy(ownerAuth1, 0)
	// 	time.Sleep(time.Second * 12)
	// 	s.Buy(ownerAuth1, 1)
	// }

	// {
	// 	ret1, err := s.lottery.GetEndowmentBalance(nil)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		// return
	// 	}
	// 	fmt.Println("GetEndowmentBalance: ", ret1)
	// }
}
