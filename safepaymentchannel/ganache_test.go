package main

import (
	// etclottery "../etclottery"
	// "context"
	// "crypto/ecdsa"
	"fmt"
	// "strconv"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "log"
	"./hdwallet"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/miguelmota/go-solidity-sha3"
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
func TestChannel(t *testing.T) {
	c := make(chan int, 1)

	select {
	case c <- 10: // c中放入了10，因为chan的buffer为1
		fmt.Println("10")
	default:
		fmt.Println("default1")
	}

	select {
	case c <- 11:
		fmt.Println("11")
	default: // c中只有10，没有11
		fmt.Println("default2")
	}

	select {
	case v, ok := <-c:
		// 读出来一个，v=10, ok=true
		fmt.Println("v ok", v, ok)
	default:
		fmt.Println("default3")
	}

	select {
	case v, ok := <-c:
		fmt.Println("v ok2", v, ok)
	default: // 没有可读的，走这个分支
		fmt.Println("default4")
	}
	// {
	// 	ch := make(chan int, 1)
	// 	for {
	// 		select {
	// 		case ch <- 0:
	// 		case ch <- 1:
	// 		}
	// 		i := <-ch
	// 		fmt.Println("Value received:", i) // 随机输出0和1
	// 	}
	// }
	// {
	// 	ch1 := make(chan int, 1)
	// 	ch2 := make(chan int, 1)
	// 	ch1 <- 0
	// 	ch2 <- 1
	// 	select {
	// 	case i := <-ch1:
	// 		fmt.Println("ch1 received:", i)
	// 	case i := <-ch2:
	// 		fmt.Println("ch2 received:", i)
	// 	}
	// }
	// {
	// 	ch := make(chan int)
	// 	go func() {
	// 		for {
	// 			select {
	// 			case i := <-ch:
	// 				fmt.Println("Value received:", i)
	// 			}
	// 		}
	// 	}()
	// 	for {
	// 		select {
	// 		case ch <- 0:
	// 			fmt.Println("0")
	// 		case ch <- 1:
	// 			fmt.Println("1")
	// 		}
	// 	}
	// }
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
		time.Sleep(time.Second * 10)
		balance, err := s.lottery.GetEndowmentBalance(nil)
		if err != nil {
			fmt.Println("GetBalance:", err)
			return
		}
		// fmt.Println("GetBalance:",addressStr,":", balance)
		fmt.Printf("GetBalance %s:%s\n", contractAddress, balance.Text(10))
	}
}
func TestSoliditySha3fromethereumjsabi(t *testing.T) {
	// new BN("43989fb883ba8111221e89123897538475893837", 16), 0, 10000, 1448075779
	// * sha3 will return ```0xc3ab5ca31a013757f26a88561f0ff5057a97dfcc33f43d6b479abc3ac2d1d595```
	// first := big.NewInt(0)
	// b, _ := first.SetString("43989fb883ba8111221e89123897538475893837", 16)
	// second := []byte(strconv.Itoa(0))
	// third := []byte(strconv.Itoa(10000))
	// fourth := []byte(strconv.Itoa(1448075779))
	// msg := crypto.Keccak256(b.Bytes(), second, third, fourth)
	// out := hex.EncodeToString(msg)
	// fmt.Println("out:", out)
	hash := solsha3.SoliditySHA3(
		solsha3.Address("0x43989fb883ba8111221e89123897538475893837"),
		solsha3.Address("0x0"),
		solsha3.Uint256(big.NewInt(10000)),
		solsha3.Uint256(big.NewInt(1448075779)))

	fmt.Println("hash:", hex.EncodeToString(hash))
}
func TestClaim(t *testing.T) {
	fmt.Println("len:", len(common.Hash{}))
	key, _ := crypto.HexToECDSA(userPrivateKey1)
	addr := common.HexToAddress(userAddress11)

	//want to send to userAddress22 1eth nonce 10
	// msg0 := crypto.Keccak256([]byte(userAddress22), big.NewInt(1000000000000000000).Bytes(), big.NewInt(10).Bytes(), []byte(contractAddress))

	msg0 := solsha3.SoliditySHA3(
		solsha3.Address(userAddress22),
		solsha3.Uint256(big.NewInt(1000000000000000000)),
		solsha3.Uint256(big.NewInt(10)),
		solsha3.Address(contractAddress),
	)
	// msg := signHash(msg0)
	msg := solsha3.SoliditySHA3WithPrefix(msg0)
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
	fmt.Println("recoveredAddr:", recoveredAddr.String())
	// // should be equal to SigToPub
	// recoveredPub2, err := SigToPub(msg, sig)
	// if err != nil {
	// 	t.Errorf("ECRecover error: %s", err)
	// }
	// recoveredAddr2 := PubkeyToAddress(*recoveredPub2)
	// if addr != recoveredAddr2 {
	// 	t.Errorf("Address mismatch: want: %x have: %x", addr, recoveredAddr2)
	// }

	ownerAuth1 := AuthAccountFromPrivateKey(userPrivateKey2)
	ownerAuth1.GasLimit = uint64(8000000)
	s := NewConnecter(selfhost, contractAddress)
	ret1, err := s.lottery.ClaimPayment(ownerAuth1, big.NewInt(1000000000000000000), big.NewInt(10), sig)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ClaimPayment: ", ret1.Hash().Hex())

}

// signHash is a helper function that calculates a hash for the given message
// that can be safely used to calculate a signature from.
//
// The hash is calulcated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
