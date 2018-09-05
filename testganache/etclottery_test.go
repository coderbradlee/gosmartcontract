package lottery

import (
	// etclottery "../etclottery"
	// "context"
	// "crypto/ecdsa"
	"fmt"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "log"
	"math/big"
	"testing"
	"time"
)

var (
	//metamask 0xe0e1b1e5d63e0ead6bfefcfc5a9dce543913cc15 7E76C9320595DEAA02A05DE3DE32507BC5C8B680B91E94384ADAA08CBAB0FF56
	// 0xaa3f808a9c7bb22bc8d81dd033811b5ee2cb2207 E0EDBBB22B16FB763C9D12F63EFC735495F6716066E720F5C308A5C5B4735923
	selfhost        = "http://172.16.5.7:18545"
	// contractAddress = "0x01CAB100BC8878ff0b91FE8eDD9A8B3EB5b353DB"
	// contractAddress = "0x46d26956E1e13F6Fb65427207B6e5194F359a3A0"
	contractAddress = "0x733A9F95cE2D8EB23EdCCcFdB3A2a647aFe1F76c"
	userAddress1="0x469B0E89aE64fF9B90FF93D078dfAA5732Fc0061"
	userAddress2="0x4A605e60d03453F038ff3b49D21053fC63e9f45A"
	userPrivateKey1   = "1a65bef07e8cfedeb61337694b31c71b8b8284dc80aa78d94638ad0202d51e79"
	userPrivateKey2   = "194309de4312874e5ea2b0a083f14d6e3de24cf298d851251a4f60e2dc7ef9a6"
	
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
	c := NewConnecterWithDeploy(selfhost,ownerAuth)

	fmt.Println("Contract address is: ", c.contractAddress.String())
	
}
func TestBuy(t *testing.T){
	s := NewConnecter(selfhost,contractAddress)
	ownerAuth1 := AuthAccountFromPrivateKey(userPrivateKey1)
	s.Buy(ownerAuth1,1)
	time.Sleep(time.Second*10)
	s.Buy(ownerAuth1,0)
	ownerAuth2 := AuthAccountFromPrivateKey(userPrivateKey2)
	s.Buy(ownerAuth2,1)
	time.Sleep(time.Second*10)
	s.Buy(ownerAuth2,0)
}
func TestWithdrawAndReinvest(t *testing.T){
	c := NewConnecter(selfhost,contractAddress)
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
			ret1, err := c.lottery.Reinvest(auth, big.NewInt(50000000000000000),1)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Reinvest: ", ret1.Hash().Hex())
	}
}
func TestWithdrawFee(t *testing.T){
	c := NewConnecter(selfhost,contractAddress)
	auth:=AuthAccountFromPrivateKey(userPrivateKey1)
	auth.GasLimit = uint64(3000000)
	
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
	s := NewConnecter(selfhost,contractAddress)
	auth := AuthAccountFromPrivateKey(userPrivateKey1)
	auth.GasLimit = uint64(3000000)
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
		ret1, err := s.lottery.GetFee(nil)
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
		ret1,err := s.lottery.GetBlock(auth)
		if err != nil {
			fmt.Println("GetBlock:",err)
			// return
		}else{
			fmt.Println("GetBlock:",ret1.Hash().Hex())
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
		ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8, ret9, ret10, ret11,err := s.lottery.GetPlayerInfoByAddress(nil, common.HexToAddress(userAddress1))
		if err != nil {
			fmt.Println(err)
			// return
		}
		fmt.Println("GetPlayerInfoByAddress: ",userAddress1)
		fmt.Println("currentroundIn0: ", ret1)
		fmt.Println("currentroundIn1: ", ret2)
		fmt.Println("lastRoundIn0: ", ret3)
		fmt.Println("lastRoundIn1: ", ret4)
		fmt.Println("allRoundIn: ", ret5)
		fmt.Println("win: ", ret6)
		fmt.Println("lastwin: ", ret7)
		fmt.Println("withdrawed: ", ret8)
		fmt.Println("currRoundId: ", ret9)
		fmt.Println("lrnd: ", ret10)
		fmt.Println("teamId: ", ret11)
	}
	fmt.Println("########################################################")
	{
		ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8,ret9, ret10, ret11, err := s.lottery.GetPlayerInfoByAddress(nil, common.HexToAddress(userAddress2))
		if err != nil {
			fmt.Println(err)
			// return
		}
		fmt.Println("GetPlayerInfoByAddress: ",userAddress2)
		fmt.Println("currentroundIn0: ", ret1)
		fmt.Println("currentroundIn1: ", ret2)
		fmt.Println("lastRoundIn0: ", ret3)
		fmt.Println("lastRoundIn1: ", ret4)
		fmt.Println("allRoundIn: ", ret5)
		fmt.Println("win: ", ret6)
		fmt.Println("lastwin: ", ret7)
		fmt.Println("withdrawed: ", ret8)
		fmt.Println("currRoundId: ", ret9)
		fmt.Println("lrnd: ", ret10)
		fmt.Println("teamId: ", ret11)
	}
	fmt.Println("########################################################")
	{
		// GetRoundInfo(opts *bind.CallOpts, roundId *big.Int) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error)
		ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8, ret9, ret10, err := s.lottery.GetRoundInfo(nil, big.NewInt(1))
		if err != nil {
			fmt.Println(err)
			// return
		}
		fmt.Println("GetRoundInfo(1)")
		fmt.Println("rid: ", ret1)
		fmt.Println("addr: ", ret2.String())
		fmt.Println("strt: ", ret3)
		fmt.Println("end: ", ret4)
		fmt.Println("etc: ", ret5)
		fmt.Println("etc0: ", ret6)
		fmt.Println("etc1: ", ret7)
		fmt.Println("pot: ", ret8)
		fmt.Println("ended: ", ret9)
		fmt.Println("winTeam: ", ret10)
	}
	fmt.Println("########################################################")
	{
		// GetCurrentInfo(opts *bind.CallOpts) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error)
		ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8, ret9, ret10, err := s.lottery.GetCurrentInfo(nil)
		if err != nil {
			fmt.Println(err)
			// return
		}
		fmt.Println("GetRoundInfo")
		fmt.Println("rid: ", ret1)
		fmt.Println("addr: ", ret2.String())
		fmt.Println("strt: ", ret3)
		fmt.Println("end: ", ret4)
		fmt.Println("etc: ", ret5)
		fmt.Println("etc0: ", ret6)
		fmt.Println("etc1: ", ret7)
		fmt.Println("pot: ", ret8)
		fmt.Println("ended: ", ret9)
		fmt.Println("winTeam: ", ret10)
	}
	fmt.Println("########################################################")
	{
		// GetTotalInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error)
		ret1, ret2, ret3, ret4, ret5, err := s.lottery.GetTotalInfo(nil)
		if err != nil {
			fmt.Println(err)
			// return
		}
		fmt.Println("totalIn: ", ret1)
		fmt.Println("bullTotalIn: ", ret2)
		fmt.Println("bearTotalIn: ", ret3)
		fmt.Println("bullTotalWin: ", ret4)
		fmt.Println("bearTotalWin: ", ret5)
	}

}
