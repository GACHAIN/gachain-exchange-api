### 1. 构建可执行程序
`go build`


### 2. 执行命令
`./gac-transfer send -ip IPADDR -prikey PRIKEY -to TO -amount AMOUNT -comment COMMENT --发起交易 `

#### 发起转账

    ./gac-transfer send -ip "http://127.0.0.1:7079" -from "./key" -to "1341-7138-3444-6302-0031" -amount "1312312" -comment "第一次转账"

    send : 表示发起交易的命令

    -ip : 主节点IP址和端口号
    -prikey : 用户私钥（私钥存放路径）
    -to : 接收者钱包地址
    -amout : 转账余额
    -comment : 转账备注

##### 返回结果：

**成功**

    result: {"data":{"amount":"1312312000000000000","block_id":"15","comment":"第一 次转账","created_at":"2018-09-19T21:13:17.737353Z","id":"1","recipien:"-5029605729246531585","sender_id":"8732751582640922090","txhash":"oꮲ%!>(MISSING)4\ufffd\ufffd\ufffdVA\ufffdt\u0018\u0000\u001c@;@#-\ufffd6\ufffd\ufffd\ufffd)3m5\ufffd"}}
    -----------交易成功---------


**失败（余额不足）**


    blockid: 0
    errmsg: {"type":"panic","error":"Current balance is not enough"}
    -----------交易失败---------


#### 查询历史

    ./gac-transfer getHistory -ip "http://127.0.0.1:7079" -prikey "./key" -limit "5" -page "1" -searchType "income"

    getHistory : 表示查询交易历史的命令

    -ip : 主节点IP地址和端口号
    -prikey : 用户私钥（私钥存放路径）
    [-limit] : 查询条数 （默认值：20）
    [-page] :  查询第几页 （默认值：1）
    [-searchType] : 查询类型（可选值：income,outcome）【income代表转入，outcome代表转出】


##### 返回结果：
    [
            {
                    "ID": 45,
                    "SenderID": 8732751582640922090,
                    "SenderAdd": "0873-2751-5826-4092-2090",
                    "RecipientID": 8732751582640922090,
                    "RecipientAdd": "0873-2751-5826-4092-2090",
                    "Amount": "30000000",
                    "Comment": "Commission for execution of @1GachainMoneyTransfer contract",
                    "BlockID": 31,
                    "TxHash": "wc7OB97y5Ak58VFS4laWgtnNmBxO5aY8mecUzM0m6xc=",
                    "CreatedAt": "2018-09-20T17:36:19.784945Z",
                    "Money": "0.00003"
            },
            {
                    "ID": 44,
                    "SenderID": 8732751582640922090,
                    "SenderAdd": "0873-2751-5826-4092-2090",
                    "RecipientID": 8732751582640922090,
                    "RecipientAdd": "0873-2751-5826-4092-2090",
                    "Amount": "970000000",
                    "Comment": "Commission for execution of @1GachainMoneyTransfer contract",
                    "BlockID": 31,
                    "TxHash": "wc7OB97y5Ak58VFS4laWgtnNmBxO5aY8mecUzM0m6xc=",
                    "CreatedAt": "2018-09-20T17:36:19.784945Z",
                    "Money": "0.00097"
            }
    ]

#### 查询余额

    ./gac-transfer getBalance -ip "http://127.0.0.1:7079" -prikey "./key" -ecosystem "1"

    getBalance : 表示发起交易的命令

    -ip : 主节点IP地址和端口号
    -prikey : 用户私钥（私钥存放路径）
    [-ecosystem] : 生态系统ID（默认值：1）


##### 返回结果：
    {
            "amount": "480315320000000000000",
            "money": "480315320"
    }



#### bug修复记录：
##### Time: 2018-9-21:
> 1. 修复分页参数无效
> 2. 修复查询历史交易哈希不能正确显示
