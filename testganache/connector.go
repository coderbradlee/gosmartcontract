package lottery

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	etclottery "../etclottery"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
)

// var (
// 	contractAddress = common.HexToAddress("0x84F70FEa5Ba54323C0EF85c58A47c98E1a2fe2Db")
	
// 	contractAddressHash = contractAddress.Hash()
// 	selfhost        = "ws://172.16.2.12:18546"
// )

type Connecter struct {
	ctx  context.Context
	conn *ethclient.Client
	lottery *etclottery.Etclottery
	contractAddress common.Address
}

func NewConnecter(host,addr string) *Connecter {
	contractAddress := common.HexToAddress(addr)
	conn, err := ethclient.Dial(host)
	if err != nil {
		panic(err)
	}
	l, err := etclottery.NewEtclottery(contractAddress, conn)
	if err != nil {
		panic(err)
	}
	return &Connecter{
		ctx:  context.Background(),
		conn: conn,
		lottery: l,
	}
}

// NewConnecterWithDeploy 部署合约，并创建一个connecter
func NewConnecterWithDeploy(host string,ownerAuth *bind.TransactOpts) *Connecter {
	conn, err := ethclient.Dial(host)
	if err != nil {
		panic(err)
	}
	_, tx, l, err := etclottery.DeployEtclottery(ownerAuth, conn)
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
		lottery: l,
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

func (c *Connecter) WatchOnBuys() {
	buyCh := make(chan *etclottery.EtclotteryOnBuys)
	buySub, err := c.lottery.WatchOnBuys(&bind.WatchOpts{}, buyCh)
	if err != nil {
		panic(err)
	}

	withdrawCh := make(chan *etclottery.EtclotteryOnWithdraw)
	withdrawSub, err := c.lottery.WatchOnWithdraw(&bind.WatchOpts{}, withdrawCh)
	if err != nil {
		panic(err)
	}

	reinvestCh := make(chan *etclottery.EtclotteryOnReinvestment)
	reinvestAddress:=[]common.Address{common.HexToAddress("0x7b6d393414d77cce0177ecd6883b08a0b5b881ea"),common.HexToAddress("0xe143c83f4e365928a1bc41034d3dc4ec8e5c47d9")}
	reinvestSub, err := c.lottery.WatchOnReinvestment(&bind.WatchOpts{}, reinvestCh, reinvestAddress)
	if err != nil {
		panic(err)
	}

	buydisCh := make(chan *etclottery.EtclotteryOnBuyAndDistribute)
	buydisSub, err := c.lottery.WatchOnBuyAndDistribute(&bind.WatchOpts{}, buydisCh)
	if err != nil {
		panic(err)
	}
	
	for {
		select {
		case t := <-buyCh:
			// Addr   common.Address Amount *big.Int Team   uint8
			log.Printf("buy addr: %s, buy: %s, team: %d", t.Addr.String(), t.Amount.Text(10), t.Team)
		case e := <-buySub.Err():
			panic(e)
		case t := <-withdrawCh://PlayerAddress common.Address EtcOut *big.Int	TimeStamp     *big.Int
			log.Printf("withdraw add: %s, Amount: %s, TimeStamp: %s", t.PlayerAddress.String(), t.EtcOut.Text(10),t.TimeStamp.Text(10))
		case e := <-withdrawSub.Err():
			panic(e)

		case t := <-reinvestCh://CustomerAddress    common.Address	EthereumReinvested *big.Int Team  uint8
			log.Printf("reinvest add: %s, Amount: %s, team: %d", t.CustomerAddress.String(), t.EthereumReinvested.Text(10),t.Team)
		case e := <-reinvestSub.Err():
			panic(e)

		case t := <-buydisCh:
	// 		Addr         common.Address
	// Strt         *big.Int
	// End          *big.Int
	// Etc          *big.Int
	// Pot          *big.Int
	// Ended        bool
	// Etc0         *big.Int
	// Etc1         *big.Int
	// Win          *big.Int
	// TotalIn      *big.Int
	// BullTotalIn  *big.Int
	// BearTotalIn  *big.Int
	// BullTotalWin *big.Int
	// BearTotalWin *big.Int
			log.Printf("buy and distribute addr: %s, Strt: %s, End: %s, Etc: %s, Pot: %s, Ended: %d, Etc0: %s, Etc1: %s, Win: %s, TotalIn: %s, BullTotalIn: %s, BearTotalIn: %s, BullTotalWin: %s, BearTotalWin: %s", t.Addr.String(), t.Strt.Text(10),t.End.Text(10),t.Etc.Text(10),t.Pot.Text(10),t.Ended,t.Etc0.Text(10),t.Etc1.Text(10),t.Win.Text(10),t.TotalIn.Text(10),t.BullTotalIn.Text(10),t.BearTotalIn.Text(10),t.BullTotalWin.Text(10),t.BearTotalWin.Text(10))
		case e := <-buydisSub.Err():
			panic(e)
		}
	}
}

func (c *Connecter) Buy(auth *bind.TransactOpts,_team uint8) {
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
	auth.GasLimit = uint64(3000000)
	auth.Value = big.NewInt(40000000000000000) //0.05eth
	ret1, err := c.lottery.Buy(auth, _team)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Buy: ", ret1.Hash().Hex())
	auth.Value = big.NewInt(0)
}