package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	// "time"
	"./erc721"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
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

	ownerKey, _, _ := createAccount()
	auth :=AuthAccountFromPrivateKey(ownerKey)
	auth.GasLimit = uint64(6800000)

	alloc := make(core.GenesisAlloc)
	b := new(big.Int)
	b.SetString("1337000000000000000000000", 10)
	fmt.Println(b.Text(10))
	alloc[auth.From] = core.GenesisAccount{Balance: b}
	sim := backends.NewSimulatedBackend(alloc, 8000000)

	contractAddr, _, _, err := erc721.DeployErc721(auth, sim)
	if err != nil {
		log.Fatalf("could not deploy Divies: %v", err)
	}
	fmt.Printf("erc721 contract at address %s\n", contractAddr.String())
	sim.Commit()

	debtorPriKey, debtorAddr, err := createAccount()
	if err != nil {
		log.Fatal("Failed to create account.", err)
	}
	fmt.Printf("debtor private:%s,address:%s\n", debtorPriKey, debtorAddr)
	creditorPriKey, creditorAddr, err := createAccount()
	if err != nil {
		log.Fatal("Failed to create account.", err)
	}
	fmt.Printf("creditor private:%s,address:%s\n", creditorPriKey, creditorAddr)

	token, err := erc721.NewErc721(contractAddr, sim)
	{
		bala, err := token.BalanceOf(nil, common.HexToAddress(debtorAddr))
		if err != nil {
			log.Fatal("BalanceOf.", err)
		}
		fmt.Printf("debtor balance:%v\n", bala.Text(10))
	}
	{
		bala, err := token.BalanceOf(nil, common.HexToAddress(creditorAddr))
		if err != nil {
			log.Fatal("BalanceOf.", err)
		}
		fmt.Printf("creditor balance:%v\n", bala.Text(10))
	}
	/////mint to creditorAddr
	_, err = token.Mint(auth, common.HexToAddress(creditorAddr),big.NewInt(11111))
	if err != nil {
		log.Fatal("mint.", err)
	}
	sim.Commit()
	fmt.Println("after mint:")
	{
		bala, err := token.BalanceOf(nil, common.HexToAddress(creditorAddr))
		if err != nil {
			log.Fatal("BalanceOf.", err)
		}
		fmt.Printf("creditor balance:%v\n", bala.Text(10))
	}
	
	creditorAuth :=AuthAccountFromPrivateKey(creditorPriKey)
	creditorAuth.GasLimit = uint64(6800000)
	_, err = token.SafeTransferFrom(creditorAuth, common.HexToAddress(creditorAddr),common.HexToAddress(debtorAddr),big.NewInt(11111),nil)
	if err != nil {
		log.Fatal("SafeTransferFrom.", err)
	}
	sim.Commit()

	fmt.Println("after transfer:")
	{
		bala, err := token.BalanceOf(nil, common.HexToAddress(debtorAddr))
		if err != nil {
			log.Fatal("BalanceOf.", err)
		}
		fmt.Printf("debtor balance:%v\n", bala.Text(10))
	}
	{
		bala, err := token.BalanceOf(nil, common.HexToAddress(creditorAddr))
		if err != nil {
			log.Fatal("BalanceOf.", err)
		}
		fmt.Printf("creditor balance:%v\n", bala.Text(10))
	}




}
func createAccount() ( private string, addr string, err error) {
	priKey, err := crypto.GenerateKey()
	if err != nil {
		return
	}
	private=fmt.Sprintf("%x", priKey.D.Bytes())
	pubKey := priKey.PublicKey
	pubBytes := crypto.FromECDSAPub(&pubKey)
	address := common.BytesToAddress(crypto.Keccak256(pubBytes[1:])[12:])
	addr=hex.EncodeToString(address[:])
	return
}
func AuthAccountFromPrivateKey(private string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		fmt.Println(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	return auth
}