package iotexsolidity

import (
	// etclottery "../etclottery"
	// "context"
	// "crypto/ecdsa"

	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var (
	//metamask 0xe0e1b1e5d63e0ead6bfefcfc5a9dce543913cc15 7E76C9320595DEAA02A05DE3DE32507BC5C8B680B91E94384ADAA08CBAB0FF56
	// 0xaa3f808a9c7bb22bc8d81dd033811b5ee2cb2207 E0EDBBB22B16FB763C9D12F63EFC735495F6716066E720F5C308A5C5B4735923
	selfhost = "http://10.1.195.63:18545"
	// contractAddress = "0x01CAB100BC8878ff0b91FE8eDD9A8B3EB5b353DB"
	// contractAddress = "0x46d26956E1e13F6Fb65427207B6e5194F359a3A0"
	contractAddress = "0x3fFf77AB74E16B33EA4b2356fb8e6b2D95E7BCd5"
	userAddress1    = "0xc416d12f3EBA9D10A1Cf21E1E6ea6509Da009ec1"
	userAddress2    = "0x86993A973Cfffc4aCe686492DbA8d718e9C0eC64"
	userPrivateKey1 = "e9348b789c81492178e4e1aedb05b0a33babd07d07efd57da9ff883a00ddbb34"
	userPrivateKey2 = "f82110fa5ea7a7d65d0ad12a8c1576c6aec921b49343fc3ac04c7263333754eb"
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
	ownerAuth.GasLimit = uint64(3000000)
	ownerAuth.Value = big.NewInt(1000000000000000000)
	ownerAuth.GasPrice = big.NewInt(10000000000)
	c := NewConnecterWithDeploy(selfhost, ownerAuth)

	fmt.Println("Contract address is: ", c.contractAddress.String())
	ret1 := c.BalanceOfEth(common.HexToAddress(c.contractAddress.String()))

	fmt.Println("balance of contract: ", ret1.Text(10))

				ret1, err := c.l.WithdrawFee(auth, big.NewInt(10000000000000000))
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("WithdrawFee: ", ret1.Hash().Hex())
			}

	// transfer to contract address,not work,need call method with payable
	//nonceUint64, err := c.conn.NonceAt(c.ctx, common.HexToAddress(userAddress1), nil)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//toAddress := common.HexToAddress(c.contractAddress.String())
	//amountInt := big.NewInt(1000000000000000000)
	//gasLimitInt := uint64(3000000)
	//gasPriceInt := big.NewInt(10000000000)
	//
	//tx := types.NewTransaction(nonceUint64, toAddress, amountInt, gasLimitInt, gasPriceInt, nil)
	//privateKey, err := crypto.HexToECDSA(userPrivateKey1)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//signed, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//err = c.conn.SendTransaction(c.ctx, signed)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//time.Sleep(time.Second * 30)
	//
	//ret2 := c.BalanceOfEth(common.HexToAddress(c.contractAddress.String()))
	//fmt.Println("balance of contract: ", ret2.Text(10))
}

//func TestWithdrawFee(t *testing.T){
//	c := NewConnecter(selfhost,contractAddress)
//	auth:=AuthAccountFromPrivateKey(userPrivateKey1)
//	auth.GasLimit = uint64(3000000)
//
//	{
//		ret1, err := c.lottery.WithdrawFee(auth, big.NewInt(10000000000000000))
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("WithdrawFee: ", ret1.Hash().Hex())
//	}
//}
//func TestGet(t *testing.T) {
//	s := NewConnecter(selfhost,contractAddress)
//	auth := AuthAccountFromPrivateKey(userPrivateKey1)
//	auth.GasLimit = uint64(3000000)
//	fmt.Println("########################################################")
//	{
//		ret1, err := s.lottery.GetCreator(nil)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("GetCreator: ", ret1.String())
//	}
//	fmt.Println("########################################################")
//	{
//		ret1, err := s.lottery.GetCurrentRoundLeft(nil)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("GetCurrentRoundLeft: ", ret1.String())
//	}
//	fmt.Println("########################################################")
//	{
//		ret1, err := s.lottery.GetFee(nil)
//		if err != nil {
//			fmt.Println(err)
//			// return
//		}
//		fmt.Println("GetFee: ", ret1)
//	}
//	fmt.Println("########################################################")
//	{
//		ret1, err := s.lottery.GetEndowmentBalance(nil)
//		if err != nil {
//			fmt.Println(err)
//			// return
//		}
//		fmt.Println("GetEndowmentBalance: ", ret1)
//	}
//	fmt.Println("########################################################")
//	{
//		ret1,err := s.lottery.GetBlock(auth)
//		if err != nil {
//			fmt.Println("GetBlock:",err)
//			// return
//		}else{
//			fmt.Println("GetBlock:",ret1.Hash().Hex())
//		}
//	}
//	// fmt.Println("########################################################")
//	// {
//	// 	ret1, err := s.lottery.Reinvest(auth, big.NewInt(1000000000000000000), 0)
//	// 	if err != nil {
//	// 		fmt.Println(err)
//	// 		return
//	// 	}
//	// 	fmt.Println("Reinvest: ", ret1.Hash().Hex())
//	// }
//	fmt.Println("########################################################")
//	{
//		// GetPlayerInfoByAddress(opts *bind.CallOpts, addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, uint8, error)
//		ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8, ret9, ret10, ret11,err := s.lottery.GetPlayerInfoByAddress(nil, common.HexToAddress(userAddress1))
//		if err != nil {
//			fmt.Println(err)
//			// return
//		}
//		fmt.Println("GetPlayerInfoByAddress: ",userAddress1)
//		fmt.Println("currentroundIn0: ", ret1)
//		fmt.Println("currentroundIn1: ", ret2)
//		fmt.Println("lastRoundIn0: ", ret3)
//		fmt.Println("lastRoundIn1: ", ret4)
//		fmt.Println("allRoundIn: ", ret5)
//		fmt.Println("win: ", ret6)
//		fmt.Println("lastwin: ", ret7)
//		fmt.Println("withdrawed: ", ret8)
//		fmt.Println("currRoundId: ", ret9)
//		fmt.Println("lrnd: ", ret10)
//		fmt.Println("teamId: ", ret11)
//	}
//	fmt.Println("########################################################")
//	{
//		ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8,ret9, ret10, ret11, err := s.lottery.GetPlayerInfoByAddress(nil, common.HexToAddress(userAddress2))
//		if err != nil {
//			fmt.Println(err)
//			// return
//		}
//		fmt.Println("GetPlayerInfoByAddress: ",userAddress2)
//		fmt.Println("currentroundIn0: ", ret1)
//		fmt.Println("currentroundIn1: ", ret2)
//		fmt.Println("lastRoundIn0: ", ret3)
//		fmt.Println("lastRoundIn1: ", ret4)
//		fmt.Println("allRoundIn: ", ret5)
//		fmt.Println("win: ", ret6)
//		fmt.Println("lastwin: ", ret7)
//		fmt.Println("withdrawed: ", ret8)
//		fmt.Println("currRoundId: ", ret9)
//		fmt.Println("lrnd: ", ret10)
//		fmt.Println("teamId: ", ret11)
//	}
//	fmt.Println("########################################################")
//	{
//		// GetRoundInfo(opts *bind.CallOpts, roundId *big.Int) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error)
//		ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8, ret9, ret10, err := s.lottery.GetRoundInfo(nil, big.NewInt(1))
//		if err != nil {
//			fmt.Println(err)
//			// return
//		}
//		fmt.Println("GetRoundInfo(1)")
//		fmt.Println("rid: ", ret1)
//		fmt.Println("addr: ", ret2.String())
//		fmt.Println("strt: ", ret3)
//		fmt.Println("end: ", ret4)
//		fmt.Println("etc: ", ret5)
//		fmt.Println("etc0: ", ret6)
//		fmt.Println("etc1: ", ret7)
//		fmt.Println("pot: ", ret8)
//		fmt.Println("ended: ", ret9)
//		fmt.Println("winTeam: ", ret10)
//	}
//	fmt.Println("########################################################")
//	{
//		// GetCurrentInfo(opts *bind.CallOpts) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error)
//		ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8, ret9, ret10, err := s.lottery.GetCurrentInfo(nil)
//		if err != nil {
//			fmt.Println(err)
//			// return
//		}
//		fmt.Println("GetRoundInfo")
//		fmt.Println("rid: ", ret1)
//		fmt.Println("addr: ", ret2.String())
//		fmt.Println("strt: ", ret3)
//		fmt.Println("end: ", ret4)
//		fmt.Println("etc: ", ret5)
//		fmt.Println("etc0: ", ret6)
//		fmt.Println("etc1: ", ret7)
//		fmt.Println("pot: ", ret8)
//		fmt.Println("ended: ", ret9)
//		fmt.Println("winTeam: ", ret10)
//	}
//	fmt.Println("########################################################")
//	{
//		// GetTotalInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error)
//		ret1, ret2, ret3, ret4, ret5, err := s.lottery.GetTotalInfo(nil)
//		if err != nil {
//			fmt.Println(err)
//			// return
//		}
//		fmt.Println("totalIn: ", ret1)
//		fmt.Println("bullTotalIn: ", ret2)
//		fmt.Println("bearTotalIn: ", ret3)
//		fmt.Println("bullTotalWin: ", ret4)
//		fmt.Println("bearTotalWin: ", ret5)
//	}
//
//}
