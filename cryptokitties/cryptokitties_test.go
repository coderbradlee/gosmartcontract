package cryptokitties

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

	"testing"
	// "time"
	// "math/rand"
)

var (
	//metamask 0xe0e1b1e5d63e0ead6bfefcfc5a9dce543913cc15 7E76C9320595DEAA02A05DE3DE32507BC5C8B680B91E94384ADAA08CBAB0FF56
	// 0xaa3f808a9c7bb22bc8d81dd033811b5ee2cb2207 E0EDBBB22B16FB763C9D12F63EFC735495F6716066E720F5C308A5C5B4735923
	// selfhost               = "http://192.168.1.52:18545"
	selfhost               = "https://kovan.infura.io"
	userAddressmetamask    = "0x6356908ACe09268130DEE2b7de643314BBeb3683"
	userPrivateKeymetamask = "0d4d9b248110257c575ef2e8d93dd53471d9178984482817dcbd6edb607f8cc5"

	userAddress11   = "0xc416d12f3EBA9D10A1Cf21E1E6ea6509Da009ec1"
	userAddress22   = "0x86993A973Cfffc4aCe686492DbA8d718e9C0eC64"
	userPrivateKey1 = "e9348b789c81492178e4e1aedb05b0a33babd07d07efd57da9ff883a00ddbb34"
	userPrivateKey2 = "f82110fa5ea7a7d65d0ad12a8c1576c6aec921b49343fc3ac04c7263333754eb"

	// 	GeneScience: 0x3fFf77AB74E16B33EA4b2356fb8e6b2D95E7BCd5
	// ERC721Metadata: 0x7687D0B5BDA5840128BcE84013D08803736c6edB
	// KittyCore: 0x2A67216e774cB9daC7Ff5c269a131d0d2864b4c8
	// SiringClockAuction: 0xF065432f6f61C62FBf70D77F00F1BA8A2fB214C4
	// SaleClockAuction: 0x47c1d3a0Bfda3Dc8cE4cE2F980F51BD64CeEB8e3

	GeneScienceAddress        = "0x3fFf77AB74E16B33EA4b2356fb8e6b2D95E7BCd5"
	ERC721MetadataAddress     = "0x7687D0B5BDA5840128BcE84013D08803736c6edB"
	KittyCoreAddress          = "0x2A67216e774cB9daC7Ff5c269a131d0d2864b4c8"
	SiringClockAuctionAddress = "0xF065432f6f61C62FBf70D77F00F1BA8A2fB214C4"
	SaleClockAuctionAddress   = "0x47c1d3a0Bfda3Dc8cE4cE2F980F51BD64CeEB8e3"
)

func TestMix(t *testing.T) {

	gene1 := "000063169218f348dc640d171b000208934b5a90189038cb3084624a50f7316c"
	// 63169218f348dc640d171b000208934b5a90189038cb3084624a50f7316c
	gene2 := "00005a13429085339c6521ef0300011c82438c628cc431a63298e3721f772d29"
	// 5a13429085339c6521ef0300011c82438c628cc431a63298e3721f772d29
	// gene3:="00004a52931ce4085c14bdce014a0318846a0c808c60294a6314a34a1295b9ce"
	// fmt.Println(anyToDecimal(gene1, 32))
	// fmt.Println(anyToDecimal(gene2, 32))
	out1 := ConvertTo32(gene1)
	fmt.Println(out1)
	out2 := ConvertTo32(gene2)
	fmt.Println(out2)
	// out=ConvertTo32(gene3)
	// fmt.Println(out)
	// ddca578ka4f7949p4d11535kaeea175h846k2243aa9gfdcd
	// ddca 578k a4f7 949p 4d11 535k aeea 175h 846k 2243 aa9g fdcd
	// c9am 6556 7ff7 b9gg 1d11 3853 9f77 6475 77k4 6784 f9gp fcaa
	// c9am 6556 7ff7 b9gg 1d11 3853 9f77 6475 77k4 6784 f9gp fcaa
	// aaaa 7885 22f2 agff 1661 7755 e979 2441 6667 7664 a9aa cfff
	// aaaa 7885 22f2 agff 1661 7755 e979 2441 6667 7664 a9aa cfff
	mix := MixGenes(out1, out2)
	fmt.Println(mix)
}

func TestDeploy(t *testing.T) {
	// ownerAuth := AuthAccount(userKeystore1, userPassphrase1)
	ownerAuth := AuthAccountFromPrivateKey(userPrivateKeymetamask)
	NewConnecterWithDeploy(selfhost, ownerAuth)

	// fmt.Println("Contract address is: ", c.contractAddress.String())

}
func TestSet(t *testing.T) {
	// ownerAuth := AuthAccount(userKeystore1, userPassphrase1)
	ownerAuth := AuthAccountFromPrivateKey(userPrivateKeymetamask)
	ownerAuth.GasLimit = uint64(3000000)
	connec, err := NewConnecter(selfhost, KittyCoreAddress, "2")
	if err != nil {
		fmt.Println(err)
		return
	}
	kc := connec.KittyCores
	{
		ret, err := kc.CeoAddress(nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ret: ", ret.Hex())
	}
	{
		// 	KittyAuction setSaleAuctionAddress 设置SaleClockAuction地址
		ret, err := kc.SetSaleAuctionAddress(ownerAuth, common.HexToAddress(SaleClockAuctionAddress))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("SetSaleAuctionAddress: ", ret.Hash().Hex())
	}
	{
		// 	KittyAuction setSiringAuctionAddress设置SiringClockAuction地址
		ret, err := kc.SetSiringAuctionAddress(ownerAuth, common.HexToAddress(SiringClockAuctionAddress))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("SetSiringAuctionAddress: ", ret.Hash().Hex())
	}
	{
		// 	KittyAccessControl setCEO设置CEO地址
		ret, err := kc.SetCEO(ownerAuth, common.HexToAddress(userAddressmetamask))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("SetCEO: ", ret.Hash().Hex())
	}
	{
		// 	KittyAccessControl setCFO设置CFO地址
		ret, err := kc.SetCFO(ownerAuth, common.HexToAddress(userAddressmetamask))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("SetCFO: ", ret.Hash().Hex())
	}
	{
		// 	KittyAccessControl setCOO设置COO地址
		ret, err := kc.SetCOO(ownerAuth, common.HexToAddress(userAddressmetamask))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("SetCOO: ", ret.Hash().Hex())
	}
	{
		// 	KittyOwnership setMetadataAddress 设置ERC721Metadata地址
		ret, err := kc.SetMetadataAddress(ownerAuth, common.HexToAddress(ERC721MetadataAddress))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("SetMetadataAddress: ", ret.Hash().Hex())
	}
	{
		// 	KittyBreeding setGeneScienceAddress
		ret, err := kc.SetGeneScienceAddress(ownerAuth, common.HexToAddress(GeneScienceAddress))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("SetGeneScienceAddress: ", ret.Hash().Hex())
	}
	// KittyCore setNewAddress 升级合约
}

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
