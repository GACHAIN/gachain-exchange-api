**Choose language [zh_cn](https://github.com/GACHAIN/gachain-exchange-api/blob/master/zh_cn.md),[en_us](https://github.com/GACHAIN/gachain-exchange-api)**

#### bug修复记录：
##### Time: 2018-9-21:

> 1. 修复分页参数无效
> 2. 修复查询历史交易哈希不能正确显示
---
> 3. 将转账成功和失败后的数据返回格式更新为`json`格式并且新增`code`（code:1 转账成功， code:0 转账失败）字段
> 4. 转账成功新增字段`hash`作为转账成功后的`txhash`
> 5. 新增获取地址和`KeyID`和`Address`
---
> 6. 新增`createPriAndPub`命令用于随机产生一对公私钥
> 7. 新增多端二进制程序

##### Time: 2018-9-25:
> 1. 删除govendor中不必要的包

##### Time: 2018-9-26:
> 1. 新增转账增加加急费

### 1. 构建二进制程序
[Reference resources](https://github.com/mitchellh/gox) 

#### 执行测试脚本
```
1. 找到对应需要测试功能的脚本文件（./test-*.sh）
2. 使用 vim ./test-*.sh 查看脚本执行命令
3. 将第一个单词和二进制文件名对应（需对应自己的操作系统）
```


### 2. 执行命令
`./gac-transfer-macos(寻找对应系统版本) send -ip IPADDR -prikey PRIKEY -to TO -amount AMOUNT -comment COMMENT --发起交易 `

#### 发起转账

    ./gac-transfer send -ip "http://127.0.0.1:7079" -prikey "./key" -to "1341-7138-3444-6302-0031" -amount "1312312" -comment "第一次转账"

    send : 表示发起交易的命令

    -ip : 主节点IP址和端口号
    -prikey : 用户私钥（私钥存放路径）
    -to : 接收者钱包地址
    -amout : 转账余额
    [-payover] : 转账加急费（不设置加急费为0）
    -comment : 转账备注

##### 返回结果：

**成功**

    {
        "block_id": 76,
        "code": 1,
        "result": "{\"data\":{\"amount\":\"1312312000000000000\",\"block_id\":\"76\",\"comment\":\"第一次转账\",\"created_at\":\"2018-09-21T11:31:21.6Z\",\"id\":\"49\",\"recipient_id\":\"-5029605729246531585\",\"sender_id\":\"8732751582640922090\",\"txhash\":\"\\ufffdd \\ufffd\\ufffd\\ufffdpˑ\\ufffd\\ufffd\\ufffd\\u0018u\\ufffdM\\n\\u0005 y:\\u000c2#\\ufffd\\ufffdJ7\\ufffd\\ufffd\"}}",
        "txHash": "8264c29eb28eae70cb919ccfe81875b64d0a05c28a793a0c3223bb844a379feb"
    }



**失败（余额不足）**

    {
        "block_id": 0,
        "code": 0,
        "errmsg": {}
        "txHash": "8264c29eb28eae70cb919ccfe81875b64d0a05c28a793a0c3223bb844a379feb"
    }



#### 查询历史

    ./gac-transfer getHistory -ip "http://127.0.0.1:7079" -prikey "./key" -limit "5" -page "1" -searchType "income"

    getHistory : 表示查询交易历史的命令

    -ip : 主节点IP地址和端口号
    -prikey : 用户私钥（私钥存放路径）
    [-limit] : 查询条数 （默认值：20）
    [-page] :  查询第几页 （默认值：1）
    [-searchType] : 查询类型（可选值：income,outcome）【income代表转入，outcome代表转出】（默认值：income）


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
            "TxHash": "67c01308faa5e3716da5b9afaee4f4e6761bd674b649b52ad30f6d14160f4354",
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
            "TxHash": "67c01308faa5e3716da5b9afaee4f4e6761bd674b649b52ad30f6d14160f4354",
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

#### 查询KeyID和地址
    ./gac-transfer getAddress -prikey "./key"

    getAddress : 表示获取KeyId和地址

    -prikey : 用户私钥（私钥存放路径）

##### 返回结果：

    {
        "Address": "0873-2751-5826-4092-2090",
        "KeyId": 8732751582640922090
    }

#### 随机生成一对公私钥
    ./gac-transfer createPriAndPub

##### 返回结果
    {
        "prikey": "a9a11d8dde1f0a3a88bdde433d8b250807c9b19f9431c12f05448326baee3290",
        "pubkey": "38fe6c442912dbb890cb2182e2f814f3f40654764cf9ef92f169ed59a29afb7d5ec80486b28920a8c3ee7bfc59d3ffd62b10e323f0b088ec0eccee6a22ea3198"
    }
