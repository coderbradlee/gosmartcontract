package iotexsolidity

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// var (
// 	contractAddress = common.HexToAddress("0x84F70FEa5Ba54323C0EF85c58A47c98E1a2fe2Db")

// 	contractAddressHash = contractAddress.Hash()
// 	selfhost        = "ws://172.16.2.12:18546"
// )

type Connecter struct {
	ctx             context.Context
	conn            *ethclient.Client
	contractAddress common.Address
}
type SpecificContract struct {
	Connecter
	MiniDAO  *MiniDAO
	Attacker *Attacker
}

func NewConnecter(host, addr string) *SpecificContract {
	contractAddress := common.HexToAddress(addr)
	conn, err := ethclient.Dial(host)
	if err != nil {
		panic(err)
	}
	l, err := NewAttacker(contractAddress, conn)
	if err != nil {
		panic(err)
	}
	return &SpecificContract{
		Connecter: Connecter{
			ctx:             context.Background(),
			conn:            conn,
			contractAddress: contractAddress,
		},
		Attacker: l,
	}
}

// NewConnecterWithDeploy 部署合约，并创建一个connecter
func NewConnecterWithDeploy(host string, ownerAuth *bind.TransactOpts) *SpecificContract {
	conn, err := ethclient.Dial(host)
	if err != nil {
		panic(err)
	}
	_, tx, MiniDAO, err := DeployMiniDAO(ownerAuth, conn)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	contractAddress, err := bind.WaitDeployed(ctx, conn, tx)
	_, tx, Attacker, err := DeployAttacker(ownerAuth, conn, contractAddress)
	if err != nil {
		panic(err)
	}
	ctx = context.Background()
	contractAddress, err = bind.WaitDeployed(ctx, conn, tx)
	// fmt.Println("contractAddress:",contractAddress)

	if err != nil {
		panic(err)
	}
	return &SpecificContract{
		Connecter: Connecter{
			ctx:             context.Background(),
			conn:            conn,
			contractAddress: contractAddress,
		},
		MiniDAO:  MiniDAO,
		Attacker: Attacker,
	}
}

// BlockNumber 当前块高度
func (c *Connecter) BlockNumber() *big.Int {
	blockNumber, err := c.conn.BlockByNumber(c.ctx, nil)
	if err != nil {
		log.Fatal("Get block number error", err)
		return big.NewInt(0)
	}
	return blockNumber.Number()
}

// BalanceOfEth 以太币余额
func (c *Connecter) BalanceOfEth(addr common.Address) *big.Int {
	balance, err := c.conn.BalanceAt(c.ctx, addr, nil)
	if err != nil {
		log.Fatal("Get address eth balance error", err)
		return big.NewInt(0)
	}
	return balance
}

// AuthAccount 解锁账户
// 正式使用时候，此处应该限制GasPrice和GasLimit
func AuthAccount(key, passphrase string) *bind.TransactOpts {
	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		log.Fatalf("Auth account error: %v", err)
		return nil
	}
	log.Printf("Gas price is: %v, Gas limit is: %v", auth.GasPrice, auth.GasLimit)
	return auth
}
func AuthAccountFromPrivateKey(private string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		fmt.Println(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)

	log.Printf("Gas price is: %v, Gas limit is: %v", auth.GasPrice, auth.GasLimit)
	return auth
}
