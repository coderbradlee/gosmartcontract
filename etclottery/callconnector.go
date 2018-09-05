package main

import (
	"./ethclient"
	// "github.com/ethereum/go-ethereum/ethclient"
	"context"
	// "crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/core/types"
	// "github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	// "strings"
)

type callConnecter struct {
	ctx             context.Context
	conn            *ethclient.Client
	lottery         *Calletclottery
	contractAddress common.Address
}

func NewCallConnecter(host, addr string) *callConnecter {
	contractAddress := common.HexToAddress(addr)
	conn, err := ethclient.Dial(host)
	if err != nil {
		panic(err)
	}
	l, err := NewCalletclottery(contractAddress, conn)
	if err != nil {
		panic(err)
	}
	return &callConnecter{
		ctx:     context.Background(),
		conn:    conn,
		lottery: l,
	}
}

// NewConnecterWithDeploy 部署合约，并创建一个connecter
func NewCallConnecterWithDeploy(host string, ownerAuth *bind.TransactOpts) *callConnecter {
	conn, err := ethclient.Dial(host)
	if err != nil {
		panic(err)
	}
	_, tx, l, err := DeployCalletclottery(ownerAuth, conn)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	contractAddress, err := bind.WaitDeployed(ctx, conn, tx)
	// fmt.Println("contractAddress:",contractAddress)

	if err != nil {
		panic(err)
	}
	return &callConnecter{
		ctx:             ctx,
		conn:            conn,
		lottery:         l,
		contractAddress: contractAddress,
	}
}

// BlockNumber 当前块高度
func (c *callConnecter) BlockNumber() *big.Int {
	blockNumber, err := c.conn.BlockByNumber(c.ctx, nil)
	if err != nil {
		log.Fatal("Get block number error", err)
		return big.NewInt(0)
	}
	return blockNumber.Number()
}
func (c *callConnecter) Accounts() (ret []string) {
	ret, err := c.conn.GetAccounts(c.ctx)
	if err != nil {
		log.Fatal("GetAccounts error", err)
		return nil
	}
	return ret
}

// BalanceOfEth 以太币余额
func (c *callConnecter) BalanceOfEth(addr common.Address) *big.Int {
	balance, err := c.conn.BalanceAt(c.ctx, addr, nil)
	if err != nil {
		log.Fatal("Get address eth balance error", err)
		return big.NewInt(0)
	}
	return balance
}

func (c *callConnecter) Buy(auth *bind.TransactOpts, _team uint8) {

	auth.GasLimit = uint64(300000)
	auth.Value = big.NewInt(1000000000000000000) //1eth
	ret1, err := c.lottery.Testbuy(auth)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Buy: ", ret1.Hash().Hex())
	auth.Value = big.NewInt(0)
}
func (c *callConnecter) Send(private, to string, amountInt *big.Int) error {

	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		// fmt.Println(err)
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		// fmt.Println("error casting public key to ECDSA")
		return err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		// fmt.Println(err)
		return err
	}

	gasPrice, err := c.conn.SuggestGasPrice(context.Background())
	if err != nil {
		// fmt.Println(err)
		return err
	}
	toAddress := common.HexToAddress(to)
	tx := types.NewTransaction(nonce, toAddress, amountInt, uint64(300000), gasPrice, nil)
	signed, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		// fmt.Println(err)
		return err
	}

	err = c.conn.SendTransaction(context.Background(), signed)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	return nil
}
