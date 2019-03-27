package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	// "time"
	"./erc721"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
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

	addr, _, _, err := erc721.DeployErc721(auth, sim)
	if err != nil {
		log.Fatalf("could not deploy Divies: %v", err)
	}
	fmt.Printf("erc721 contract at address %s\n", addr.String())


	_, debtorPriKey, debtorAddr, err := createAccount()
	if err != nil {
		log.Fatal("Failed to create account.", zap.Error(err))
	}
	fmt.Printf("debtor private:%s,address:%s\n",debtorPriKey,debtorAddr)
	_, creditorPriKey, creditorAddr, err := createAccount()
	if err != nil {
		log.Fatal("Failed to create account.", zap.Error(err))
	}
	fmt.Printf("creditor private:%s,address:%s\n",creditorPriKey,creditorAddr)

}
func createAccount() (public string, private string, addr string, err error) {
	priKey, err := crypto.GenerateKey()
	if err != nil {
		return
	}
	private=priKey.HexString()


	pubKey := priKey.PublicKey
	public=pubKey.HexString()
	pubBytes := crypto.FromECDSAPub(&pubKey)
	address := common.BytesToAddress(crypto.Keccak256(pubBytes[1:])[12:])
	addr=hex.EncodeToString(address[:])

	return
}