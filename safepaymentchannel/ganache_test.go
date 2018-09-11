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
}
func TestClaim(t *testing.T) {
	{
		key, _ := crypto.HexToECDSA(userPrivateKey1)
		addr := common.HexToAddress(userAddress11)

		//want to send to userAddress22 1eth nonce 10
		msg := crypto.Keccak256([]byte(userAddress22), big.NewInt(1000000000000000000).Bytes(), big.NewInt(10).Bytes(), []byte(contractAddress))
		sig, err := crypto.Sign(msg, key)
		if err != nil {
			fmt.Printf("Sign error: %s\n", err)
		}
		recoveredPub, err := crypto.Ecrecover(msg, sig)
		if err != nil {
			fmt.Printf("ECRecover error: %s\n", err)
		}
		pubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
		recoveredAddr := crypto.PubkeyToAddress(*pubKey)
		if addr != recoveredAddr {
			fmt.Printf("Address mismatch: want: %x have: %x\n", addr, recoveredAddr)
		}
		fmt.Println("recoveredAddr:", recoveredAddr)
		// // should be equal to SigToPub
		// recoveredPub2, err := SigToPub(msg, sig)
		// if err != nil {
		// 	t.Errorf("ECRecover error: %s", err)
		// }
		// recoveredAddr2 := PubkeyToAddress(*recoveredPub2)
		// if addr != recoveredAddr2 {
		// 	t.Errorf("Address mismatch: want: %x have: %x", addr, recoveredAddr2)
		// }
	}
	// ownerAuth1 := AuthAccountFromPrivateKey(userPrivateKey2)
	// s := NewConnecter(selfhost, contractAddress)
	// ret1, err := s.lottery.ClaimPayment(ownerAuth1, big.NewInt(1000000000000000000), big.NewInt(10), signature []byte)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("ClaimPayment: ", ret1.Hash().Hex())

}
