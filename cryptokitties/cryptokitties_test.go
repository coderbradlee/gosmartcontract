package cryptokitties

import (
	// etclottery "../etclottery"
	// "context"
	// "crypto/ecdsa"
	"fmt"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "log"
	// "math/big"
	"testing"
	// "time"
	// "math/rand"
)

var (
	//metamask 0xe0e1b1e5d63e0ead6bfefcfc5a9dce543913cc15 7E76C9320595DEAA02A05DE3DE32507BC5C8B680B91E94384ADAA08CBAB0FF56
	// 0xaa3f808a9c7bb22bc8d81dd033811b5ee2cb2207 E0EDBBB22B16FB763C9D12F63EFC735495F6716066E720F5C308A5C5B4735923
	selfhost        = "ws://172.16.2.12:18546"
	// contractAddress = "0x01CAB100BC8878ff0b91FE8eDD9A8B3EB5b353DB"
	// contractAddress = "0x46d26956E1e13F6Fb65427207B6e5194F359a3A0"
	contractAddress = "0xCE8667E5a2cf69Cb9929C1faD2088486bE81Ad71"
	userAddress1="0x7b6d393414d77cce0177ecd6883b08a0b5b881ea"
	userAddress2="0xe143c83f4e365928a1bc41034d3dc4ec8e5c47d9"
	userKeystore1   = `{"id":"16f2aa71-435d-66a9-8dc9-4c4889103772","version":3,"crypto":{"cipher":"aes-128-ctr","cipherparams":{"iv":"11d6de57754581d92c701fcda0312e22"},"ciphertext":"8c5593f40055c9d686b5901036ef180c8f2e856f7b608baa56c45ce9f2371e5d","kdf":"pbkdf2","kdfparams":{"c":10240,"dklen":32,"prf":"hmac-sha256","salt":"552e3fe91c2c6ac4241ca7ca1d44979081cc5a7890f4647d8e38b899816c08bc"},"mac":"c6fa315e304b7625d4bd2b28b26d8f3ca1774c0e2e3461b8ae9faec1156b068b"},"address":"7b6d393414d77cce0177ecd6883b08a0b5b881ea","name":"","meta":"{}"}`
	userPassphrase1 = "123456789"
	userKeystore2   = `{"id":"954ae627-bd39-6a38-bb0b-8f290dfe8dc9","version":3,"crypto":{"cipher":"aes-128-ctr","cipherparams":{"iv":"92c2dd85471721fd7e0997ab9f5703c4"},"ciphertext":"cddc315e72d89a50228d4244e1e390413e12f39ed0fbd348312ebc39b1108175","kdf":"pbkdf2","kdfparams":{"c":10240,"dklen":32,"prf":"hmac-sha256","salt":"d8afc1723758d985273eb08979d2c9542757bd57f26e841e2372b507a80ec15f"},"mac":"fed2fdb5b538e65c9ee11c47b7d6f31f17eab1abb6776abcae1fe8d9573e959c"},"address":"e143c83f4e365928a1bc41034d3dc4ec8e5c47d9","name":"","meta":"{}"}`
	userPassphrase2 = "123456789"
)
// func TestDeploy(t *testing.T) {
// 	// deploy()
// }
// func TestWatch(t *testing.T) {
// 	s := NewConnecter(selfhost,contractAddress)

// 	s.WatchOnBuys()
	
// }
// func mix(genes1,genes2 string)string{
// 	// mGenes[48], sGenes[48], babyGenes[48]

// 	for i:=0;i<12;i++{
// 		index:=4*i
// 		for j:=3;j>0;j--{
// 			if rand.Float64()< 0.25{
// 				// swap(mGenes, index+j, index+j-1)
// 				temp:=mGenes[index+j]
// 				mGenes[index+j]=mGenes[index+j-1]
// 				mGenes[index+j-1]=temp
// 			}
// 			if rand.Float64()< 0.25{
// 				// swap(sGenes, index+j, index+j-1)
// 				temp:=sGenes[index+j]
// 				sGenes[index+j]=sGenes[index+j-1]
// 				sGenes[index+j-1]=temp
// 			}
// 		}
// 	}
// 	for i:=0;i<48;i++{
// 		mutation := 0
// 		if i % 4 == 0{
// 			gene1 := mGene[i]
// 			gene2 := sGene[i]
// 			if gene1 > gene2{
// 				gene1, gene2 = gene2, gene1
// 			}
// 			if (gene2 - gene1 == 1) && iseven(gene1){
// 				probability = 0.25
// 				if gene1 > 23{
// 					probability /= 2
// 				}
// 				if rand.Float64() < probability{
// 					mutation = (gene1 / 2) + 16
// 				}
// 			}
// 		}
//       	if mutation{
// 			baby[i] = mutation
// 		}else{
// 			if rand.Float64() < 0.5{
// 				babyGenes[i] = mGene[i]
// 			}else{
// 				babyGenes[i] = sGene[i]
// 			}	
// 		}
// 	}
// }

func TestMix(t *testing.T){
	gene1:="0x000063169218f348dc640d171b000208934b5a90189038cb3084624a50f7316c"
	gene2:="0x000063169218f348dc640d171b000208934b5a90189038cb3084624a50f7316c"
	// fmt.Println(anyToDecimal(gene1, 32))
	// fmt.Println(anyToDecimal(gene2, 32))
	g1:=big.NewInt(0)
	g1.SetString(gene1,16)
	fmt.Println(g1.Text(16))

	g1.SetString(gene2,16)
	fmt.Println(g1.Text(16))
}
// func TestDeploy(t *testing.T) {
// 	ownerAuth := AuthAccount(userKeystore1,userPassphrase1)
// 	c := NewConnecterWithDeploy(selfhost,ownerAuth)

// 	fmt.Println("Contract address is: ", c.contractAddress.String())
	
// }
// func TestBuy(t *testing.T){
// 	s := NewConnecter(selfhost,contractAddress)
// 	ownerAuth1 := AuthAccount(userKeystore1,userPassphrase1)
// 	s.Buy(ownerAuth1,1)
// 	time.Sleep(time.Second*10)
// 	s.Buy(ownerAuth1,0)
// 	// ownerAuth2 := AuthAccount(userKeystore2,userPassphrase2)
// 	// s.Buy(ownerAuth2,0)
// }
// func TestWithdrawAndReinvest(t *testing.T){
// 	c := NewConnecter(selfhost,contractAddress)
// 	auth := AuthAccount(userKeystore1,userPassphrase1)
// 	auth.GasLimit = uint64(3000000)
// 	// auth.Value = big.NewInt(50000000000000000) //0.05eth
// 	{
// 		// ret1, err := c.lottery.PlayerWithdraw(auth, big.NewInt(50000000000000000))
// 		// if err != nil {
// 		// 	fmt.Println(err)
// 		// 	return
// 		// }
// 		// fmt.Println("PlayerWithdraw: ", ret1.Hash().Hex())
// 	}
	
// 	{
// 			ret1, err := c.lottery.Reinvest(auth, big.NewInt(50000000000000000),1)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			fmt.Println("Reinvest: ", ret1.Hash().Hex())
// 	}
// }
// func TestWithdrawFee(t *testing.T){
// 	c := NewConnecter(selfhost,contractAddress)
// 	auth := AuthAccount(userKeystore1,userPassphrase1)
// 	auth.GasLimit = uint64(3000000)
	
// 	{
// 		ret1, err := c.lottery.WithdrawFee(auth, big.NewInt(10000000000000000))
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Println("WithdrawFee: ", ret1.Hash().Hex())
// 	}
// }
// func TestGet(t *testing.T) {
// 	s := NewConnecter(selfhost,contractAddress)
// 	auth := AuthAccount(userKeystore1,userPassphrase1)
// 	// auth.GasLimit = uint64(3000000)
// 	fmt.Println("########################################################")
// 	{
// 		ret1, err := s.lottery.GetCreator(nil)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Println("GetCreator: ", ret1.String())
// 	}
// 	fmt.Println("########################################################")
// 	{
// 		ret1, err := s.lottery.GetCurrentRoundLeft(nil)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Println("GetCurrentRoundLeft: ", ret1.String())
// 	}
// 	fmt.Println("########################################################")
// 	{
// 		ret1, err := s.lottery.GetFee(nil)
// 		if err != nil {
// 			fmt.Println(err)
// 			// return
// 		}
// 		fmt.Println("GetFee: ", ret1)
// 	}
// 	fmt.Println("########################################################")
// 	{
// 		ret1, err := s.lottery.GetEndowmentBalance(nil)
// 		if err != nil {
// 			fmt.Println(err)
// 			// return
// 		}
// 		fmt.Println("GetEndowmentBalance: ", ret1)
// 	}
// 	fmt.Println("########################################################")
// 	{
// 		ret1,err := s.lottery.GetBlock(auth)
// 		if err != nil {
// 			fmt.Println("GetBlock:",err)
// 			// return
// 		}else{
// 			fmt.Println("GetBlock:",ret1.Hash().Hex())
// 		}
// 	}
// 	// fmt.Println("########################################################")
// 	// {
// 	// 	ret1, err := s.lottery.Reinvest(auth, big.NewInt(1000000000000000000), 0)
// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 		return
// 	// 	}
// 	// 	fmt.Println("Reinvest: ", ret1.Hash().Hex())
// 	// }
// 	fmt.Println("########################################################")
// 	{
// 		// GetPlayerInfoByAddress(opts *bind.CallOpts, addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, uint8, error) 
// 		ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8, ret9, ret10, ret11,err := s.lottery.GetPlayerInfoByAddress(nil, common.HexToAddress(userAddress1))
// 		if err != nil {
// 			fmt.Println(err)
// 			// return
// 		}
// 		fmt.Println("GetPlayerInfoByAddress: ",userAddress1)
// 		fmt.Println("currentroundIn0: ", ret1)
// 		fmt.Println("currentroundIn1: ", ret2)
// 		fmt.Println("lastRoundIn0: ", ret3)
// 		fmt.Println("lastRoundIn1: ", ret4)
// 		fmt.Println("allRoundIn: ", ret5)
// 		fmt.Println("win: ", ret6)
// 		fmt.Println("lastwin: ", ret7)
// 		fmt.Println("withdrawed: ", ret8)
// 		fmt.Println("currRoundId: ", ret9)
// 		fmt.Println("lrnd: ", ret10)
// 		fmt.Println("teamId: ", ret11)
// 	}
// 	fmt.Println("########################################################")
// 	{
// 		ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8,ret9, ret10, ret11, err := s.lottery.GetPlayerInfoByAddress(nil, common.HexToAddress(userAddress2))
// 		if err != nil {
// 			fmt.Println(err)
// 			// return
// 		}
// 		fmt.Println("GetPlayerInfoByAddress: ",userAddress1)
// 		fmt.Println("currentroundIn0: ", ret1)
// 		fmt.Println("currentroundIn1: ", ret2)
// 		fmt.Println("lastRoundIn0: ", ret3)
// 		fmt.Println("lastRoundIn1: ", ret4)
// 		fmt.Println("allRoundIn: ", ret5)
// 		fmt.Println("win: ", ret6)
// 		fmt.Println("lastwin: ", ret7)
// 		fmt.Println("withdrawed: ", ret8)
// 		fmt.Println("currRoundId: ", ret9)
// 		fmt.Println("lrnd: ", ret10)
// 		fmt.Println("teamId: ", ret11)
// 	}
// 	fmt.Println("########################################################")
// 	{
// 		// GetRoundInfo(opts *bind.CallOpts, roundId *big.Int) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error)
// 		ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8, ret9, ret10, err := s.lottery.GetRoundInfo(nil, big.NewInt(1))
// 		if err != nil {
// 			fmt.Println(err)
// 			// return
// 		}
// 		fmt.Println("GetRoundInfo(1)")
// 		fmt.Println("rid: ", ret1)
// 		fmt.Println("addr: ", ret2.String())
// 		fmt.Println("strt: ", ret3)
// 		fmt.Println("end: ", ret4)
// 		fmt.Println("etc: ", ret5)
// 		fmt.Println("etc0: ", ret6)
// 		fmt.Println("etc1: ", ret7)
// 		fmt.Println("pot: ", ret8)
// 		fmt.Println("ended: ", ret9)
// 		fmt.Println("winTeam: ", ret10)
// 	}
// 	fmt.Println("########################################################")
// 	{
// 		// GetCurrentInfo(opts *bind.CallOpts) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error)
// 		ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8, ret9, ret10, err := s.lottery.GetCurrentInfo(nil)
// 		if err != nil {
// 			fmt.Println(err)
// 			// return
// 		}
// 		fmt.Println("GetRoundInfo")
// 		fmt.Println("rid: ", ret1)
// 		fmt.Println("addr: ", ret2.String())
// 		fmt.Println("strt: ", ret3)
// 		fmt.Println("end: ", ret4)
// 		fmt.Println("etc: ", ret5)
// 		fmt.Println("etc0: ", ret6)
// 		fmt.Println("etc1: ", ret7)
// 		fmt.Println("pot: ", ret8)
// 		fmt.Println("ended: ", ret9)
// 		fmt.Println("winTeam: ", ret10)
// 	}
// 	fmt.Println("########################################################")
// 	{
// 		// GetTotalInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error)
// 		ret1, ret2, ret3, ret4, ret5, err := s.lottery.GetTotalInfo(nil)
// 		if err != nil {
// 			fmt.Println(err)
// 			// return
// 		}
// 		fmt.Println("totalIn: ", ret1)
// 		fmt.Println("bullTotalIn: ", ret2)
// 		fmt.Println("bearTotalIn: ", ret3)
// 		fmt.Println("bullTotalWin: ", ret4)
// 		fmt.Println("bearTotalWin: ", ret5)
// 	}

// }
