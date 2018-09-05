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

var (
//metamask 0xe0e1b1e5d63e0ead6bfefcfc5a9dce543913cc15 7E76C9320595DEAA02A05DE3DE32507BC5C8B680B91E94384ADAA08CBAB0FF56
// 0xaa3f808a9c7bb22bc8d81dd033811b5ee2cb2207 E0EDBBB22B16FB763C9D12F63EFC735495F6716066E720F5C308A5C5B4735923
// selfhost = "http://172.16.5.7:18545"
// // contractAddress = "0x01CAB100BC8878ff0b91FE8eDD9A8B3EB5b353DB"
// // contractAddress = "0x46d26956E1e13F6Fb65427207B6e5194F359a3A0"
// contractAddress = "0x39965D2Ae7D65C21914b7E57d69ad1F39D0aFb60"
// userAddress1="0x469B0E89aE64fF9B90FF93D078dfAA5732Fc0061"
// userAddress2="0x4A605e60d03453F038ff3b49D21053fC63e9f45A"
// userPrivateKey1   = "1a65bef07e8cfedeb61337694b31c71b8b8284dc80aa78d94638ad0202d51e79"
// userPrivateKey2   = "194309de4312874e5ea2b0a083f14d6e3de24cf298d851251a4f60e2dc7ef9a6"

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
func TestWithdrawAndReinvest(t *testing.T) {
	c := NewConnecter(selfhost, contractAddress)
	auth := AuthAccountFromPrivateKey(userPrivateKey1)
	auth.GasLimit = uint64(3000000)
	// auth.Value = big.NewInt(50000000000000000) //0.05eth
	{
		// ret1, err := c.lottery.PlayerWithdraw(auth, big.NewInt(50000000000000000))
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// fmt.Println("PlayerWithdraw: ", ret1.Hash().Hex())
	}

	{
		ret1, err := c.lottery.Reinvest(auth, big.NewInt(50000000000000000), 1)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Reinvest: ", ret1.Hash().Hex())
	}
}
func TestWithdrawFee(t *testing.T) {
	c := NewConnecter(selfhost, contractAddress)
	auth := AuthAccountFromPrivateKey(userPrivateKey1)
	auth.GasLimit = uint64(8000000)

	{
		ret1, err := c.lottery.WithdrawFee(auth, big.NewInt(10000000000000000))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("WithdrawFee: ", ret1.Hash().Hex())
	}
}
func TestGet(t *testing.T) {
	s := NewConnecter(selfhost, contractAddress)
	auth := AuthAccountFromPrivateKey(userPrivateKey1)
	auth.GasLimit = uint64(8000000)
	fmt.Println("########################################################")
	{
		ret1, err := s.lottery.GetCreator(nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("GetCreator: ", ret1.String())
	}
	fmt.Println("########################################################")
	{
		ret1, err := s.lottery.GetCurrentRoundLeft(nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("GetCurrentRoundLeft: ", ret1.String())
	}
	fmt.Println("########################################################")
	{
		ret1, err := s.lottery.GetFee(auth)
		if err != nil {
			fmt.Println(err)
			// return
		}
		fmt.Println("GetFee: ", ret1)
	}
	fmt.Println("########################################################")
	{
		ret1, err := s.lottery.GetEndowmentBalance(nil)
		if err != nil {
			fmt.Println(err)
			// return
		}
		fmt.Println("GetEndowmentBalance: ", ret1)
	}
	fmt.Println("########################################################")
	{
		ret1, _, err := s.lottery.GetAddressArray(nil)
		if err != nil {
			fmt.Println("GetAddressArray:", err)
			// return
		}
		fmt.Println("GetAddressArray: ", ret1.String())
	}
	fmt.Println("########################################################")
	{
		ret1, err := s.lottery.GetAddressLength(nil)
		if err != nil {
			fmt.Println("GetAddressLength:", err)
			// return
		}
		fmt.Println("GetAddressLength: ", ret1)
	}
	fmt.Println("########################################################")
	{
		ret1, err := s.lottery.GetBlock(auth)
		if err != nil {
			fmt.Println("GetBlock:", err)
			// return
		} else {
			fmt.Println("GetBlock:", ret1.Hash().Hex())
		}
	}
	// fmt.Println("########################################################")
	// {
	// 	ret1, err := s.lottery.Reinvest(auth, big.NewInt(1000000000000000000), 0)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println("Reinvest: ", ret1.Hash().Hex())
	// }
	fmt.Println("########################################################")
	{
		// GetPlayerInfoByAddress(opts *bind.CallOpts, addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, uint8, error)
		ret1, ret2, ret3, ret4, ret5, ret6, err := s.lottery.GetPlayerInfoByAddress(nil, common.HexToAddress(userAddress11))
		if err != nil {
			fmt.Println(err)
			// return
		}
		fmt.Println("GetPlayerInfoByAddress: ", userAddress11)
		fmt.Println("currentroundIn0: ", ret1)
		fmt.Println("currentroundIn1: ", ret2)
		fmt.Println("allRoundIn: ", ret3)
		fmt.Println("win: ", ret4)
		fmt.Println("lastwin: ", ret5)
		fmt.Println("withdrawed: ", ret6)
	}
	fmt.Println("########################################################")
	{
		ret1, ret2, ret3, ret4, ret5, ret6, err := s.lottery.GetPlayerInfoByAddress(nil, common.HexToAddress(userAddress22))
		if err != nil {
			fmt.Println(err)
			// return
		}
		fmt.Println("GetPlayerInfoByAddress: ", userAddress22)
		fmt.Println("currentroundIn0: ", ret1)
		fmt.Println("currentroundIn1: ", ret2)
		fmt.Println("allRoundIn: ", ret3)
		fmt.Println("win: ", ret4)
		fmt.Println("lastwin: ", ret5)
		fmt.Println("withdrawed: ", ret6)
	}
	fmt.Println("########################################################")
	{
		// GetRoundInfo(opts *bind.CallOpts, roundId *big.Int) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error)
		ret1, ret2, ret3, ret4, ret5, ret6, err := s.lottery.GetLastRoundInfo(nil)
		if err != nil {
			fmt.Println(err)
			// return
		}
		fmt.Println("GetLastRoundInfo")
		fmt.Println("strt: ", ret1)
		fmt.Println("end: ", ret2)
		fmt.Println("etc0: ", ret3)
		fmt.Println("etc1: ", ret4)
		fmt.Println("ended: ", ret5)
		fmt.Println("winTeam: ", ret6)
	}
	fmt.Println("########################################################")
	{
		// GetCurrentInfo(opts *bind.CallOpts) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error)
		ret1, ret2, ret3, ret4, ret5, ret6, ret7, err := s.lottery.GetCurrentInfo(nil)
		if err != nil {
			fmt.Println(err)
			// return
		}
		fmt.Println("GetCurrentInfo")
		fmt.Println("rid: ", ret1)
		fmt.Println("strt: ", ret2)
		fmt.Println("end: ", ret3)
		fmt.Println("etc0: ", ret4)
		fmt.Println("etc1: ", ret5)
		fmt.Println("ended: ", ret6)
		fmt.Println("winTeam: ", ret7)
	}
	fmt.Println("########################################################")
	{
		// GetTotalInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error)
		ret1, ret2, ret3, ret4, err := s.lottery.GetTotalInfo(nil)
		if err != nil {
			fmt.Println(err)
			// return
		}
		fmt.Println("GetTotalInfo: ")
		fmt.Println("bullTotalIn: ", ret1)
		fmt.Println("bearTotalIn: ", ret2)
		fmt.Println("bullTotalWin: ", ret3)
		fmt.Println("bearTotalWin: ", ret4)
	}

}
