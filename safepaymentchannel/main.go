package main

import (
	"./hdwallet"
	"./xlog"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/robfig/cron"
	"log"
	"math/big"
	"time"
)

type Conf struct {
	Host            string `json:"host"`
	ContractAddress string `json:"contractAddress"`
	Interval        string `json:"interval"`
	Keystore        string `json:"keystore"`
	Password        string `json:"password"`
	Gaslimit        int64  `json:"gaslimit"`
	Gaspricemin     int64  `json:"gaspricemin"`
	Gaspricemax     int64  `json:"gaspricemax"`
}

var cfg Conf //proxy.Config

var lastGetTime = time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

func watch() {
	// fmt.Println("start watching...")
	interval := cfg.Interval
	log.Printf("config is :%v", interval)
	c := cron.New()
	c.AddFunc(interval, func() {
		runOnce()
	})
	c.Start()
}
func runOnce() {
	log.Println("run once:")
	s := NewConnecter(cfg.Host, cfg.ContractAddress)
	auth := AuthAccount(cfg.Keystore, cfg.Password)
	// auth:= AuthAccountFromPrivateKey(userPrivateKey1)
	auth.GasLimit = uint64(cfg.Gaslimit)
	log.Println("GasLimit: ", auth.GasLimit)
	gasPrice, err := s.conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
	}
	log.Println("GasPrice from chain: ", gasPrice.Text(10))
	gasmin := big.NewInt(cfg.Gaspricemin)
	gasmax := big.NewInt(cfg.Gaspricemax)
	auth.GasPrice = gasPrice
	if gasPrice.Cmp(gasmin) < 0 {
		auth.GasPrice = gasmin
	}
	if gasPrice.Cmp(gasmax) > 0 {
		auth.GasPrice = gasmax
	}
	// balanceWei := auth.GasPrice.Mul(auth.GasPrice, big.NewInt(cfg.Gaslimit))
	tempPrice := big.NewInt(0)
	tempPrice.Set(auth.GasPrice)
	balanceWei := tempPrice.Mul(tempPrice, big.NewInt(cfg.Gaslimit))
	log.Println("GasPrice: ", auth.GasPrice.Text(10))
	log.Println("need balance: ", balanceWei.Text(10))
	left, err := s.lottery.GetCurrentRoundLeft(nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("GetCurrentRoundLeft: ", left.String())
	// _, _, _, _, _, _, _, ret8, _, _, err := s.lottery.GetCurrentInfo(nil)
	// _, _, _, ret4, ret5, _, _, err := s.lottery.GetCurrentInfo(nil)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// a := ret4.Add(ret4, ret5)
	// m := a.Mul(a, big.NewInt(8))
	// pot := m.Div(m, big.NewInt(100))
	// log.Println("GetRoundInfo pot: ", pot)
	if left.Cmp(big.NewInt(0)) == 0 {

		if time.Now().Sub(lastGetTime).Seconds() < 300 {
			log.Println("GetBlock<300s")
			return
		}

		clearingHash, err := s.lottery.GetBlock(auth)
		if err != nil {
			log.Println("GetBlock:", err)
			// return
		} else {
			log.Println("GetBlock:", clearingHash.Hash().Hex())
		}
		lastGetTime = time.Now()
	}
}
func generateAccounts() {

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
func testBuy() {
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

		privateKeyBytes := priKey.Serialize()
		private_str := hex.EncodeToString(privateKeyBytes)
		fmt.Printf("The %d ETH privateKeyBytes is %s\n", i, private_str)
		s := NewConnecter(selfhost, contractAddress)
		ownerAuth1 := AuthAccountFromPrivateKey(private_str)
		s.Buy(ownerAuth1, uint8(i%2))
		time.Sleep(time.Second * 10)
	}
}
func main() {
	xlog.XX()
	if !LoadConfig("etclottery.toml", &cfg) {
		return
	}

	log.Println(cfg)
	// generateAccounts()
	// testBuy()
	go watch()

	quit := make(chan bool)
	<-quit
}
