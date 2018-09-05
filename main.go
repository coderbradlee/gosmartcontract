package main

import (
	"fmt"
	// "log"
	// "math/big"
	// // "time"
	// "./Divies"
	// "crypto/ecdsa"
	// // "./PlayerBook"
	// fomo3d "./FomoETC"
	// "./PlayerBook"
	// // "./testinterface"
	// "context"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// // "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	// // "github.com/ethereum/go-ethereum/core"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/ethclient"
	// // "golang.org/x/net/websocket"
	// "./test"
	"./testgaslimit"
)

func testwatch(){
	fmt.Println("start watching...")
	selfhost        := "ws://172.16.2.12:18546"
	contractAddress := "0x46d26956E1e13F6Fb65427207B6e5194F359a3A0"
	s := highlimit.NewConnecter(selfhost,contractAddress)

	s.WatchOnDel()
}
func main() {
	testwatch()
	
}
