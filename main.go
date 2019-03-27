package main

import (
	"fmt"
	"log"
	"math/big"
	"./erc721"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	creditorKey, creditorAddr, _ := createAccount()
	creditorAuth := AuthAccountFromPrivateKey(creditorKey)
	creditorAuth.GasLimit = uint64(6800000)

	alloc := make(core.GenesisAlloc)
	b := new(big.Int)
	b.SetString("1337000000000000000000000", 10)
	alloc[creditorAuth.From] = core.GenesisAccount{Balance: b}
	sim := backends.NewSimulatedBackend(alloc, 8000000)

	contractAddr, _, _, err := erc721.DeployErc721(creditorAuth, sim)
	if err != nil {
		log.Fatalf("could not deploy erc721: %v", err)
	}
	log.Printf("erc721 contract at address %s\n", contractAddr.String())
	sim.Commit()

	debtorPriKey, debtorAddr, err := createAccount()
	if err != nil {
		log.Fatal("Failed to create account.", err)
	}
	log.Printf("debtor private:%s,address:%s\n", debtorPriKey, debtorAddr)

	token, err := erc721.NewErc721(contractAddr, sim)
	{
		bala, err := token.BalanceOf(nil, common.HexToAddress(debtorAddr))
		if err != nil {
			log.Fatal("BalanceOf.", err)
		}
		log.Printf("debtor balance:%v\n", bala.Text(10))
	}
	{
		bala, err := token.BalanceOf(nil, common.HexToAddress(creditorAddr))
		if err != nil {
			log.Fatal("BalanceOf.", err)
		}
		log.Printf("creditor balance:%v\n", bala.Text(10))
	}
	/////mint to creditorAddr
	_, err = token.Mint(creditorAuth, common.HexToAddress(creditorAddr), big.NewInt(11111))
	if err != nil {
		log.Fatal("mint.", err)
	}
	sim.Commit()
	log.Println("after mint:")
	{
		bala, err := token.BalanceOf(nil, common.HexToAddress(creditorAddr))
		if err != nil {
			log.Fatal("BalanceOf.", err)
		}
		log.Printf("creditor balance:%v\n", bala.Text(10))
	}
	_, err = token.SafeTransferFrom(creditorAuth, common.HexToAddress(creditorAddr), common.HexToAddress(debtorAddr), big.NewInt(11111), nil)
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
		log.Printf("debtor balance:%v\n", bala.Text(10))
	}
	{
		bala, err := token.BalanceOf(nil, common.HexToAddress(creditorAddr))
		if err != nil {
			log.Fatal("BalanceOf.", err)
		}
		log.Printf("creditor balance:%v\n", bala.Text(10))
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
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	return auth
}