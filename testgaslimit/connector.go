package highlimit

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/types"
	"crypto/ecdsa"
)

// var (
// 	contractAddress = common.HexToAddress("0x84F70FEa5Ba54323C0EF85c58A47c98E1a2fe2Db")
	
// 	contractAddressHash = contractAddress.Hash()
// 	selfhost        = "ws://172.16.2.12:18546"
// )

type Connecter struct {
	ctx  context.Context
	conn *ethclient.Client
	hl *Highlimit
	contractAddress common.Address
	rpc      *RPCClient
}

func NewConnecter(host,addr string) *Connecter {
	contractAddress := common.HexToAddress(addr)
	conn, err := ethclient.Dial(host)
	if err != nil {
		panic(err)
	}
	l, err := NewHighlimit(contractAddress, conn)
	if err != nil {
		panic(err)
	}
	rpc := NewRPCClient("test", host, "20s")
	return &Connecter{
		ctx:  context.Background(),
		conn: conn,
		hl: l,
		rpc:rpc,
	}
	
}

// NewConnecterWithDeploy 部署合约，并创建一个connecter
func NewConnecterWithDeploy(host string,ownerAuth *bind.TransactOpts) *Connecter {
	conn, err := ethclient.Dial(host)
	if err != nil {
		panic(err)
	}
	_, tx, l, err := DeployHighlimit(ownerAuth, conn)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	contractAddress, err := bind.WaitDeployed(ctx, conn, tx)
	// fmt.Println("contractAddress:",contractAddress)

	if err != nil {
		panic(err)
	}
	return &Connecter{
		ctx:  ctx,
		conn: conn,
		hl: l,
		contractAddress:contractAddress,
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
func (c *Connecter) Send(private,to string,amountInt *big.Int)error {

	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		// fmt.Println(err)
		return err;
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		// fmt.Println("error casting public key to ECDSA")
		return err;
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		// fmt.Println(err)
		return err;
	}

	gasPrice, err := c.conn.SuggestGasPrice(context.Background())
	if err != nil {
		// fmt.Println(err)
		return err;
	}
	toAddress:=common.HexToAddress(to)
	
	// auth := bind.NewKeyedTransactor(privateKey)
	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(1000000000000000000)     // in wei,0.001ether
	// auth.GasLimit = uint64(300000) // in units
	// auth.GasPrice = gasPrice
	// nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte
	tx := types.NewTransaction(nonce, toAddress, amountInt, uint64(300000), gasPrice, nil)
	signed, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		// fmt.Println(err)
		return err;
	}

	err=c.conn.SendTransaction(context.Background(),signed)
	if err != nil {
		// fmt.Println(err)
		return err;
	}
	return nil;
}
func (c *Connecter) Run(auth *bind.TransactOpts) {
	
	auth.GasLimit = uint64(8000000)
	// auth.Value = big.NewInt(50000000000000000) //0.05eth
	ret1, err := c.hl.Run(auth)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Run: ", ret1.Hash().Hex())
}
// func (c *Connecter) WatchOnDel() {
// 	buyCh := make(chan *HighlimitOnDel)
// 	buySub, err := c.hl.WatchOnDel(&bind.WatchOpts{}, buyCh)
// 	if err != nil {
// 		panic(err)
// 	}
	
// 	for {
// 		select {
// 		case t := <-buyCh:
// 			// Addr   common.Address Amount *big.Int Team   uint8
// 			fmt.Println(t.A1,",",t.A2,",",t.A3,",",t.A4,",",t.A5,",",t.A6,",",t.A17)
// 		case e := <-buySub.Err():
// 			panic(e)
		
// 		}
// 	}
// }