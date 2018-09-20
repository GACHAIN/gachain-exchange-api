package main

import (
	"flag"
	"fmt"
	"os"

	"encoding/hex"
	api "github.com/GACHAIN/gac-transfer/chainapi"
	"github.com/GACHAIN/gac-transfer/crypto"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/url"
	"time"
)

var (
	err     error
	key     []byte
	sign    []byte
	pub     string
	blockId int64
	result  string
)

type WalletHistory struct {
	tableName    string
	ID           int64
	SenderID     int64
	SenderAdd    string
	RecipientID  int64
	RecipientAdd string
	Amount       decimal.Decimal
	Comment      string
	BlockID      int64
	TxHash       []byte
	CreatedAt    time.Time
	Money        string
}

type myBalanceResult struct {
	Amount string `json:"amount"`
	Money  string `json:"money"`
}

type CLI struct{}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tsend -ip IP -prikey PRIKEY -to TO -amount AMOUNT -comment COMMENT --发起交易 ")
	fmt.Println("\tgetBalance -ip IP -prikey PRIKEY -ecosystem ECOSYSTEMID  --查询余额")
	fmt.Println("\tgetHistory -ip IP -prikey PRIKEY -limit LIMIT -page PAGE -searchType SEARCHTYPE  - --查询交易历史")
}

func Exit(n int) {
	printUsage()
	os.Exit(n)
}

func (cli *CLI) Run() {
	// 输入参数小于2验证
	IsvalidArgs()
	// 发起交易
	sendGac := flag.NewFlagSet("send", flag.ExitOnError)
	sendGacbyIp := sendGac.String("ip", "", "节点IP地址")
	sendGacByPrikey := sendGac.String("prikey", "", "发送者私钥")
	flagTo := sendGac.String("to", "", "转账目的地址")
	lagAmount := sendGac.String("amount", "", "转账金额")
	comment := sendGac.String("comment", "", "转账备注")

	// 查询余额
	getBalancecmd := flag.NewFlagSet("getBalance", flag.ExitOnError)
	getBalanceByIp := getBalancecmd.String("ip", "", "节点IP地址")
	getBalanceByPrikey := getBalancecmd.String("prikey", "", "查询者私钥")
	getBalanceByEcosystem := getBalancecmd.String("ecosystem", "1", "生态系统ID")

	// 查询历史
	getHistorycmd := flag.NewFlagSet("getHistory", flag.ExitOnError)
	getHistoryByIp := getHistorycmd.String("ip", "", "节点IP地址")
	getHistoryByPrikey := getHistorycmd.String("prikey", "", "私钥")
	limit := getHistorycmd.String("limit", "", "查询条数")
	page := getHistorycmd.String("page", "", "查询分页")
	searchType := getHistorycmd.String("searchType", "", "查询历史类型")

	switch os.Args[1] {
	case "send":
		err := sendGac.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
		}

	case "getBalance":
		err := getBalancecmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
		}
	case "getHistory":
		err := getHistorycmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
		}
	default:
		Exit(1)
	}

	if sendGac.Parsed() {
		// 判断参数的有效性
		if *sendGacbyIp == "" || *sendGacByPrikey == "" || *flagTo == "" || *lagAmount == "" || *comment == "" {
			fmt.Println("转账参数输入错误，请严格按照提示信息输入")
			Exit(1)
		} else {
			// 接收参数
			ip, prikey, to, amount, comment := *sendGacbyIp, *sendGacByPrikey, *flagTo, *lagAmount, *comment
			// 发送参数执行转账
			cli.Send(ip, prikey, to, amount, comment)
		}
	}

	if getBalancecmd.Parsed() {
		if *getBalanceByIp == "" || *getBalanceByPrikey == "" || *getBalanceByEcosystem == "" {
			// 判断参数的有效性
			fmt.Println("查询余额参数输入错误，请严格按照提示信息输入！")
			Exit(1)
		} else {
			ip, prikey, ecosystem  := *getBalanceByIp, *getBalanceByPrikey, *getBalanceByEcosystem
			cli.GetBalance(ip, prikey, ecosystem)
		}
	}

	if getHistorycmd.Parsed() {
		if *getHistoryByIp == "" || *getHistoryByPrikey == "" || *limit == "" || *page == "" || *searchType == "" {
			fmt.Println("查询历史参数输入错误，请严格按照提示信息输入")
			Exit(1)
		} else {
			ip, prikey, limit, page, searchType := *getHistoryByIp, *getHistoryByPrikey, *limit, *page, *searchType
			cli.GetHistory(ip, prikey, limit, page, searchType)
		}
	}
}

func (cli *CLI) Send(ip string, prikey string, to string, amount string, comment string) {
	api.ApiAddress = ip
	if api.KeyLogin(prikey, 1); err != nil {
		fmt.Println("error", err)
		return
	}

	key, err = ioutil.ReadFile(prikey)
	if err != nil {
		return
	}
	if len(key) > 64 {
		key = key[:64]
	}

	pub, err = api.PrivateToPublicHex(string(key))
	transferData :=
		"Comment:" + comment + ",Gac:" + amount + ",Recipient:" + to + ",payover:0,pubkey:" + pub

	sign, err = crypto.Sign(string(key), transferData)
	if err != nil {
		return
	}
	mysign := hex.EncodeToString(sign)
	pub, err = api.PrivateToPublicHex(string(key))
	if err != nil {
		return
	}

	params := &url.Values{
		`Comment`:   {comment},
		`Gac`:       {amount},
		`Recipient`: {to},
		`payover`:   {`0`},
		`pubkey`:    {pub},
		`gacsign`:   {mysign},
	}
	blockId, result, err = api.PostTxResult(`GachainMoneyTransfer`, params)
	if err != nil {
		fmt.Println("blockid:", blockId)
		fmt.Println("errmsg:", err)
		fmt.Println("-----------交易失败---------")
	} else {
		fmt.Println("blockid:", blockId)
		fmt.Println("result:", result)
		fmt.Println("-----------交易成功---------")
	}
}

// 查询余额
func (cli *CLI) GetBalance(ip string, prikey string, ecosystem string) {
	var balanceresult myBalanceResult
	api.ApiAddress = ip
	// 1. 登陆
	if api.KeyLogin(prikey, 1); err != nil {
		fmt.Println("error:", err)
		return
	}
	key, err = ioutil.ReadFile(prikey)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	if len(key) > 64 {
		key = key[:64]
	}
	params := &url.Values{
		"ecosystem":      {`1`},
	}
	err = api.SendGet("myBalance", params, &balanceresult)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("balanceresult: ", balanceresult)
}

// 查询历史
func (cli *CLI) GetHistory(ip string, prikey string, limit string, page string, searchType string) {
	var (
		walletHistories []WalletHistory
	)
	api.ApiAddress = ip
	// 1. 登陆
	if api.KeyLogin(prikey, 1); err != nil {
		fmt.Println("error:", err)
		return
	}
	key, err = ioutil.ReadFile(prikey)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	if len(key) > 64 {
		key = key[:64]
	}
	params := &url.Values{
		"limit":      {limit},
		"page":       {page},
		`searchType`: {searchType},
	}
	err = api.SendGet("walletHistory", params, &walletHistories)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("historyResult: ", walletHistories)
}

// 参数错误校验
func IsvalidArgs() {
	if len(os.Args) < 2 {
		Exit(1)
	}
}

func main() {
	cli := CLI{}
	//CLI命令行
	cli.Run()
}
