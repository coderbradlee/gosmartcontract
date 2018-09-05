package main

import (
	"math/big"
	// etclottery "../etclottery"
	// "context"
	// "crypto/ecdsa"
	"fmt"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
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
	{
		s := NewConnecter(selfhost, callcontractAddress)
		s.Send(userPrivateKey1, callcontractAddress, big.NewInt(5000000000000000000))
		{
			ownerAuth1 := AuthAccountFromPrivateKey(userPrivateKey1)
			s.Buy(ownerAuth1, 0)
			time.Sleep(time.Second * 10)
			s.Buy(ownerAuth1, 1)
			time.Sleep(time.Second * 10)
		}
		{
			ownerAuth1 := AuthAccountFromPrivateKey(userPrivateKey2)
			s.Buy(ownerAuth1, 0)
			time.Sleep(time.Second * 12)
			s.Buy(ownerAuth1, 1)
		}
	}

}
