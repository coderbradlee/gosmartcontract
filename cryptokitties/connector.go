package cryptokitties

import (
	"./ethclient"
	// "github.com/ethereum/go-ethereum/ethclient"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strings"
)

type Connecter struct {
	ctx  context.Context
	conn *ethclient.Client
	// lottery         *NXlottery
	contractAddress common.Address
}

// func NewConnecter(host, addr string) *Connecter {
// 	contractAddress := common.HexToAddress(addr)
// 	conn, err := ethclient.Dial(host)
// 	if err != nil {
// 		panic(err)
// 	}
// 	l, err := NewNXlottery(contractAddress, conn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return &Connecter{
// 		ctx:     context.Background(),
// 		conn:    conn,
// 		lottery: l,
// 	}
// }

// NewConnecterWithDeploy 部署合约，并创建一个connecter
func NewConnecterWithDeploy(host string, ownerAuth *bind.TransactOpts) (ret *Connecter, errs error) {
	ownerAuth.GasLimit = uint64(3000000)
	conn, err := ethclient.Dial(host)
	var kittycorecontractaddress common.Address
	if err != nil {
		// panic(err)
		fmt.Println("Dial err:", err)
		errs = err
		return
	}
	{
		_, tx, _, err := DeployGeneScienceInterface(ownerAuth, conn)
		if err != nil {
			fmt.Println("DeployGeneScienceInterface err:", err)
			errs = err
			return
		}
		ctx := context.Background()
		contractAddress, err := bind.WaitDeployed(ctx, conn, tx)
		fmt.Println("GeneScience:", contractAddress.Hex())
	}
	{
		_, tx, _, err := DeployERC721Metadata(ownerAuth, conn)
		if err != nil {
			fmt.Println("DeployERC721Metadata err:", err)
			errs = err
			return
		}
		ctx := context.Background()
		contractAddress, err := bind.WaitDeployed(ctx, conn, tx)
		fmt.Println("ERC721Metadata:", contractAddress.Hex())
	}
	{
		_, tx, _, err := DeployKittyCore(ownerAuth, conn)
		if err != nil {
			fmt.Println("DeployKittyCore err:", err)
			errs = err
			return
		}
		ctx := context.Background()
		kittycorecontractaddress, err := bind.WaitDeployed(ctx, conn, tx)
		fmt.Println("KittyCore:", kittycorecontractaddress.Hex())
	}
	{
		_, tx, _, err := DeploySiringClockAuction(ownerAuth, conn, kittycorecontractaddress, big.NewInt(375))
		if err != nil {
			fmt.Println("DeploySiringClockAuction err:", err)
			errs = err
			return
		}
		ctx := context.Background()
		contractAddress, err := bind.WaitDeployed(ctx, conn, tx)
		fmt.Println("SiringClockAuction:", contractAddress.Hex())
	}
	{
		_, tx, _, err := DeploySaleClockAuction(ownerAuth, conn, kittycorecontractaddress, big.NewInt(375))
		if err != nil {
			fmt.Println("DeploySaleClockAuction err:", err)
			errs = err
			return
		}
		ctx := context.Background()
		contractAddress, err := bind.WaitDeployed(ctx, conn, tx)
		fmt.Println("SaleClockAuction:", contractAddress.Hex())
	}

	if err != nil {
		// panic(err)
		fmt.Println("last err:", err)
		errs = err
		return
	}

	ret = &Connecter{
		// ctx:  ctx,
		conn: conn,
		// lottery:         l,
		contractAddress: kittycorecontractaddress,
	}
	return
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
func (c *Connecter) Accounts() (ret []string) {
	ret, err := c.conn.GetAccounts(c.ctx)
	if err != nil {
		log.Fatal("GetAccounts error", err)
		return nil
	}
	return ret
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

func (c *Connecter) WatchOnBuys() {
	// buyCh := make(chan *EtclotteryOnBuys)
	// buySub, err := c.lottery.WatchOnBuys(&bind.WatchOpts{}, buyCh)
	// if err != nil {
	// 	panic(err)
	// }

	// withdrawCh := make(chan *EtclotteryOnWithdraw)
	// withdrawSub, err := c.lottery.WatchOnWithdraw(&bind.WatchOpts{}, withdrawCh)
	// if err != nil {
	// 	panic(err)
	// }

	// buydisCh := make(chan *EtclotteryOnBuyAndDistribute)
	// buydisSub, err := c.lottery.WatchOnBuyAndDistribute(&bind.WatchOpts{}, buydisCh)
	// if err != nil {
	// 	panic(err)
	// }

	// for {
	// 	select {
	// 	case t := <-buyCh:
	// 		// Addr   common.Address Amount *big.Int Team   uint8
	// 		log.Printf("buy addr: %s, buy: %s, team: %d", t.Addr.String(), t.Amount.Text(10), t.Team)
	// 	case e := <-buySub.Err():
	// 		panic(e)
	// 	case t := <-withdrawCh: //PlayerAddress common.Address EtcOut *big.Int	TimeStamp     *big.Int
	// 		log.Printf("withdraw add: %s, Amount: %s, TimeStamp: %s", t.PlayerAddress.String(), t.Out.Text(10), t.TimeStamp.Text(10))
	// 	case e := <-withdrawSub.Err():
	// 		panic(e)

	// 	case t := <-buydisCh:
	// 		log.Printf("buy and distribute addr: %s, Strt: %s, End: %s, Etc: %s, Pot: %s, Ended: %d, Etc0: %s, Etc1: %s, Win: %s, TotalIn: %s, BullTotalIn: %s, BearTotalIn: %s, BullTotalWin: %s, BearTotalWin: %s", t.Addr.String(), t.Strt.Text(10), t.End.Text(10), t.Etc.Text(10), t.Pot.Text(10), t.Etc0.Text(10), t.Etc1.Text(10), t.Win.Text(10))
	// 	case e := <-buydisSub.Err():
	// 		panic(e)
	// 	}
	// }
}

func (c *Connecter) Buy(auth *bind.TransactOpts, _team uint8) {
	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	gasPrice, err := client.SuggestGasPrice(context.Background())
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	auth := bind.NewKeyedTransactor(privateKey)
	// 	auth.Nonce = big.NewInt(int64(nonce))
	// 	auth.Value = big.NewInt(1000000000000000)     // in wei,0.001ether
	// 	auth.GasLimit = uint64(300000) // in units
	// 	auth.GasPrice = gasPrice
	// auth.GasLimit = uint64(300000)
	// auth.Value = big.NewInt(1000000000000000000) //0.05eth
	// ret1, err := c.lottery.Buy(auth, _team)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Buy: ", ret1.Hash().Hex())
	// auth.Value = big.NewInt(0)
}
func (c *Connecter) Send(private, to string, amountInt *big.Int) error {
	// privateKey, err := crypto.HexToECDSA(private)
	// if err != nil {
	// 	// fmt.Println(err)
	// 	return err;
	// }
	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := privateKey.(*ecdsa.PublicKey)
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
