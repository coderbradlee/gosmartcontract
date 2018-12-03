package loopring

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
	"time"
)

var (
	//ganache 0x469B0E89aE64fF9B90FF93D078dfAA5732Fc0061 1a65bef07e8cfedeb61337694b31c71b8b8284dc80aa78d94638ad0202d51e79
	// 0x4A605e60d03453F038ff3b49D21053fC63e9f45A 194309de4312874e5ea2b0a083f14d6e3de24cf298d851251a4f60e2dc7ef9a6
	selfhost = "http://192.168.1.52:18545"
	// contractAddress = "0x01CAB100BC8878ff0b91FE8eDD9A8B3EB5b353DB"
	contractAddress = "0x39965D2Ae7D65C21914b7E57d69ad1F39D0aFb60"
	userAddress1    = "0xA6e2d4666b0c4B73bDbb54CA8209bfD0EC89Ea88"
	userAddress2    = "0x1eB5E96518f6DA9FFcf32a401F26e2a8CA9E10b7"
	userPrivateKey1 = "8b01fcf00efb07a877e57c69811b414586ec1e7d531443ab0d88a9093b05f015"
	userPrivateKey2 = "6a7af08e232aacd433158d5586d513be64d97044e83b7f5738bc345917ba6c17"
)

func TestDeploy(t *testing.T) {
	ownerAuth := AuthAccountFromPrivateKey(userPrivateKey1)
	c := NewConnecterWithDeploy(selfhost, ownerAuth)

	fmt.Println("Contract address is: ", c.contractAddress.String())

}
func TestRun(t *testing.T) {
	go func() {
		s := NewConnecter(selfhost, contractAddress)
		ownerAuth1 := AuthAccountFromPrivateKey(userPrivateKey1)
		privateKey, err := crypto.HexToECDSA(userPrivateKey1)
		if err != nil {
			// fmt.Println(err)
			return
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			// fmt.Println("error casting public key to ECDSA")
			return
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := s.conn.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			// fmt.Println(err)
			return
		}

		for i := 0; i < 10; i++ {
			ownerAuth1.Nonce = big.NewInt(int64(nonce + uint64(i)))
			s.Run(ownerAuth1)
		}

	}()
	go func() {
		s := NewConnecter(selfhost, contractAddress)
		// ownerAuth1 := AuthAccountFromPrivateKey(userPrivateKey2)
		// reply,err:=s.rpc.SendTransaction(userAddress1, userAddress2, "30000", "20000000000", "1000000000000000000",false)
		for i := 0; i < 10; i++ {
			err := s.Send(userPrivateKey2, userAddress1, big.NewInt(1000000000000000000))
			if err != nil {
				fmt.Println("sendtransaction:", err)
			}
			fmt.Println("sendtransaction:")
			time.Sleep(12 * time.Second)
		}
	}()
	{
		s := NewConnecter(selfhost, contractAddress)
		bn := s.BlockNumber()
		fmt.Println("blocknumber:", bn)
	}
	// {
	// 	s := NewConnecter(selfhost,contractAddress)
	// 	bn:=s.hl.Del()
	// 	fmt.Println("Del:",bn.Hash().Hex())
	// }
	c := make(chan int)
	<-c
}
