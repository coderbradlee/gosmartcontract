package main

import (
	// etclottery "../etclottery"
	// "context"
	// "crypto/ecdsa"
	"fmt"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "log"
	"./hdwallet"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"testing"
	"time"
)

// func TestDeploy(t *testing.T) {
// 	// deploy()
// }
// func TestWatch(t *testing.T) {
// 	s := NewConnecter(selfhost,contractAddress)

// 	s.WatchOnBuys()

// }
func TestDeploy(t *testing.T) {
	ownerAuth := AuthAccountFromPrivateKey(userPrivateKey1)
	c := NewConnecterWithDeploy(selfhost, ownerAuth)

	fmt.Println("Contract address is: ", c.contractAddress.String())

}
func TestAccounts(t *testing.T) {

	mnemonic := "syringe height fiat cool rims loincloth weavers newt envy egotistic items rugged"
	for i := 0; i < 100; i++ {

		extKey, err := hdwallet.GetPrivateKeyByPath(mnemonic, "ETH", 0, 0, i)
		if err != nil {
			fmt.Println("GetPrivateKeyByPath: ", err)
			return
		}
		priKey, err := extKey.ECPrivKey()
		if err != nil {
			return
		}

		// priKeyec:=priKey.ToECDSA()

		// fmt.Printf("priKeyec:%x\n",priKeyec.D.Bytes())

		privateKeyBytes := priKey.Serialize()
		private_str := hex.EncodeToString(privateKeyBytes)
		fmt.Printf("The %d ETH privateKeyBytes is %s\n", i, private_str)
		{

			publickey, _ := extKey.ECPubKey()
			var p *ecdsa.PublicKey
			p = (*ecdsa.PublicKey)(publickey)
			pubBytes := crypto.FromECDSAPub(p)
			pkPrv := common.BytesToAddress(crypto.Keccak256(pubBytes[1:])[12:])
			pkHash := pkPrv[:]
			addressStr := "0x" + hex.EncodeToString(pkHash)
			// fmt.Println("The ETH address is ", addressStr)
			c := NewConnecter(selfhost, contractAddress)
			c.Send(userPrivateKey1, addressStr, big.NewInt(1000000000000000000))
			time.Sleep(time.Second * 10)
			balance, err := c.conn.GetBalance(c.ctx, addressStr)
			if err != nil {
				fmt.Println("GetBalance:", err)
				return
			}
			// fmt.Println("GetBalance:",addressStr,":", balance)
			fmt.Printf("The %d GetBalance %s:%s\n", i, addressStr, balance.Text(10))
			fmt.Println(" ")
		}
	}

}
func TestBuy(t *testing.T) {
	{
		s := NewConnecter(selfhost, contractAddress)
		{
			// ownerAuth1 := AuthAccountFromPrivateKey(userPrivateKey1)
			s.Send(userPrivateKey1, contractAddress, big.NewInt(1000000000000000000))
		}

	}

	// mnemonic:="syringe height fiat cool rims loincloth weavers newt envy egotistic items rugged"
	// for i:=0;i<100;i++{

	// 	extKey,err:=hdwallet.GetPrivateKeyByPath(mnemonic,"ETH",0,0,i)
	// 	if err!=nil{
	// 		fmt.Println("GetPrivateKeyByPath: ", err)
	// 		return
	// 	}
	// 	priKey,err:=extKey.ECPrivKey()
	// 	if err != nil {
	// 		return
	// 	}

	// 	privateKeyBytes := priKey.Serialize()
	// 	private_str := hex.EncodeToString(privateKeyBytes)
	// 	fmt.Printf("The %d ETH privateKeyBytes is %s\n", i,private_str)
	// 	s := NewConnecter(selfhost, contractAddress)
	// 	ownerAuth1 := AuthAccountFromPrivateKey(private_str)
	// 	s.Buy(ownerAuth1, uint8(i%2))
	// 	time.Sleep(time.Second*10)
	// }
	//
	// time.Sleep(time.Second * 10)
	// s.Buy(ownerAuth1, 0)
	// ownerAuth2 := AuthAccountFromPrivateKey(userPrivateKey2)
	// s.Buy(ownerAuth2, 1)
	// time.Sleep(time.Second * 10)
	// s.Buy(ownerAuth2, 0)
}
