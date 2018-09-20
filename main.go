package main

import (
	"fmt"
	"flag"
	"os"

	api "github.com/GACHAIN/gac-transfer/chainapi"
	"github.com/GACHAIN/gac-transfer/crypto"
	"io/ioutil"
	"encoding/hex"
	"net/url"
	"github.com/shopspring/decimal"
	"time"
)

var (
	ipAddr *string
	prikey *string
)

type WalletHistory struct {
	tableName string
	ID int64
	SenderID int64
	SenderAdd string
	RecipientID int64
	RecipientAdd string
	Amount decimal.Decimal
	Comment string
	BlockID int64
	TxHash []byte
	CreatedAt time.Time
	Money string
}

type CLI struct {}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tsend -ip IP -prikey PRIKEY -to TO -amount AMOUNT -comment COMMENT --发起交易 ")
	//fmt.Println("\tgetBalance -address --查询余额")
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
	ipAddr = sendGac.String("ip", "", "节点IP地址")
	prikey = sendGac.String("from", "", "发送者私钥")
	flagTo := sendGac.String("to", "", "转账目的地址")
	lagAmount := sendGac.String("amount", "", "转账金额")
	comment := sendGac.String("comment", "", "转账备注")

	// 查询余额
	getBalancecmd := flag.NewFlagSet("getBalance", flag.ExitOnError)
	getBalanceWithAddress := getBalancecmd.String("address", "", "查询某个地址对应的余额")

	// 查询历史
	getHistorycmd := flag.NewFlagSet("getHistory", flag.ExitOnError)
	ipAddr = getHistorycmd.String("ip", "", "节点IP地址")
	prikey = getHistorycmd.String("prikey", "", "私钥")
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
		if *ipAddr == "" || *prikey == "" || *flagTo == "" || *lagAmount == "" || *comment == "" {
			fmt.Println("转账参数输入错误，请严格按照提示信息输入")
			Exit(1)
		}else {
			// 接收参数
			ip, from, to, amount, comment := *ipAddr, *prikey, *flagTo, *lagAmount, *comment
			// 发送参数执行转账
			cli.Send(ip, from, to, amount, comment)
		}
	}

	if getBalancecmd.Parsed() {
		if (*getBalanceWithAddress == ""){
			// 判断参数的有效性
			fmt.Println("查询余额参数输入错误，请严格按照提示信息输入！")
			Exit(1)
		}else{
			address := *getBalanceWithAddress
			cli.GetBalance(address)
		}
	}

	if getHistorycmd.Parsed() {
		if *ipAddr == "" || *prikey == "" || *limit == "" || *page == "" || *searchType == ""{
			fmt.Println("查询历史参数输入错误，请严格按照提示信息输入")
			Exit(1)
		}else{
			ip, prikey, limit, page, searchType  := *ipAddr, *prikey, *limit, *page, *searchType
			cli.GetHistory(ip, prikey, limit, page, searchType)
		}
	}
}
var (
	err   error
	key []byte
	sign []byte
	pub string
	blockId int64
	result string
)
// 发起交易
func (cli *CLI) Send(ip string, from string, to string, amount string, comment string) {
	api.ApiAddress = ip
	// 1. 登陆
	if api.KeyLogin(from, 1); err != nil {
		fmt.Println("error", err)
		return
	}

	key, err = ioutil.ReadFile(from)
	if err != nil {
		return
	}
	if len(key) > 64 {
		key = key[:64]
	}

	pub, err = api.PrivateToPublicHex(string(key))
	transferData :=
		"Comment:"+comment+",Gac:"+amount+",Recipient:"+to+",payover:0,pubkey:"+pub

	sign, err = crypto.Sign(string(key), transferData)
	if err != nil {
		return
	}
	mysign := hex.EncodeToString(sign)
	pub, err = api.PrivateToPublicHex(string(key))
	if err != nil {
		return
	}
	// 2. 转账
	blockId, result, err = api.PostTxResult(`GachainMoneyTransfer`, &url.Values{
		`Comment`: {comment},
		`Gac`:{amount},
		`Recipient`:{to},
		`payover`:{`0`},
		`pubkey`:{pub},
		`gacsign`:{mysign},
	})

	//1341-7138-3444-6302-0031

	if err != nil {
		fmt.Println("blockid:", blockId)
		fmt.Println("errmsg:", err)
		fmt.Println("-----------交易失败---------")
	}else{
		fmt.Println("blockid:",blockId)
		fmt.Println("result:", result)
		fmt.Println("-----------交易成功---------")
	}
}

// 查询余额
func (cli *CLI) GetBalance(address string) {
	fmt.Println(address)
	fmt.Println("开始查询余额的逻辑")
}

// 查询历史
func (cli *CLI) GetHistory(ip string, prikey string, limit string, page string, searchType string) {
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
	var walletHistories []WalletHistory
	err = api.SendGet("walletHistory", nil, &walletHistories)
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

func main(){
	cli := CLI{}
	//CLI命令行
	cli.Run()
}

