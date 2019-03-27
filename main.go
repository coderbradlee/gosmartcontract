package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	// "time"
	"./Divies"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	host = "http://47.91.31.224:2301"
	// host="http://13.113.9.175:18545"
)

func main() {
	// total indicates the total amount value of a erc721 token
	var total int64
	// risk indicates the risk amount value of a erc721 token
	var risk int64
	// transfer indicates the transfer amount value of a erc721 token
	var transfer int64

	flag.Int64Var(&total, "total", 10000, "total amount value of a erc721 token")
	flag.Int64Var(&risk, "risk", 2000, "risk amount value of a erc721 token")
	flag.Int64Var(&transfer, "transfer", 1000, "transfer amount value of a erc721 token")
	flag.Parse()
	if risk > total {
		log.Fatal("risk amount cannot be greater than total amount")
	}


	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	alloc := make(core.GenesisAlloc)
	b := new(big.Int)
	b.SetString("1337000000000000000000000", 10)
	fmt.Println(b.Text(10))
	alloc[auth.From] = core.GenesisAccount{Balance: b}
	sim := backends.NewSimulatedBackend(alloc,8000000)

	addrDivies, _, _, err := Divies.DeployDivies(auth, sim)
	if err != nil {
		log.Fatalf("could not deploy Divies: %v", err)
	}
	fmt.Printf("Divies contract at address %s...\n", addrDivies.String())

}
