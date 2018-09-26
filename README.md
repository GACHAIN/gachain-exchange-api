**Choose language [zh_cn](https://github.com/GACHAIN/gachain-exchange-api/blob/master/zh_cn.md),[en_us](https://github.com/GACHAIN/gachain-exchange-api)**

#### Bug fix record：
##### Time: 2018-9-21:

> 1. Fix paging parameter is invalid
> 2. Fix query history transaction hash does not display correctly
---
> 3. Update the data return format after the transfer is successful and failed to the `json` format and add the `code` (code:1 transfer succeeded, code:0 transfer failed) field
> 4. The transfer successfully added the field `hash` as the `txhash` after the successful transfer.
> 5. Add the get address and `KeyID` and `Address`
---
> 6. Added `createPriAndPub` command to randomly generate a pair of public and private keys
> 7. New multi-end binary program, support one-click cross-compilation

##### Time: 2018-9-25:
> 1. Delete unnecessary packages in govendor


##### Time: 2018-9-26:
> 1. Add transfer to increase expedited fee

### 1. Build binary
[Reference resources](https://github.com/mitchellh/gox) 

#### Execute test script
```
1. Find the script file (./test-*.sh) that corresponds to the test function you need.
2. Use vim ./test-*.sh to view script execution commands
3. Correspond to the first word and the binary file name (need to correspond to your own operating system)
```


### 2. Excuting an order
`./gac-transfer-macos(Find the corresponding system version) send -ip IPADDR -prikey PRIKEY -to TO -amount AMOUNT -comment COMMENT --Initiate a transaction`

#### Initiate a transfer

    ./gac-transfer send -ip "http://127.0.0.1:7079" -from "./key" -to "1341-7138-3444-6302-0031" -amount "1312312" -comment "First transfer"

    send : Indicates the order to initiate a transaction

    -ip : Primary node? IP address and port number
    -prikey : User private key (private key? storage path)
    -to : Recipient wallet address
    -amout : Transfer balance
    [-payover]: Transfer expedited fee (do not set expedited fee is 0)
    -comment : Transfer note

##### Result：

**Success**

    {
        "block_id": 76,
        "code": 1,
        "result": "{\"data\":{\"amount\":\"1312312000000000000\",\"block_id\":\"76\",\"comment\":\"Initiate a transaction\",\"created_at\":\"2018-09-21T11:31:21.6Z\",\"id\":\"49\",\"recipient_id\":\"-5029605729246531585\",\"sender_id\":\"8732751582640922090\",\"txhash\":\"\\ufffdd\\ufffd\\ufffd\\ufffdpˑ\\ufffd\\ufffd\\ufffd\\u0018u\\ufffdM\\n\\u0005y:\\u000c2#\\ufffd\\ufffdJ7\\ufffd\\ufffd\"}}",
        "txHash": "8264c29eb28eae70cb919ccfe81875b64d0a05c28a793a0c3223bb844a379feb"
    }



**error(current blance is not energy)**

    {
        "block_id": 0,
        "code": 0,
        "errmsg": {}
    }



#### Query history

    ./gac-transfer getHistory -ip "http://127.0.0.1:7079" -prikey "./key" -limit "5" -page "1" -searchType "income"

    getHistory : indicates the command to query the transaction history

    -ip : Primary node? IP address and port number
    -prikey : User private key (private key? storage path)
    [-limit] : Number of queries (default: 20)
    [-page] :  Query the first page (default: 1)
    [-searchType] : Query type (optional value: income, outcome) [income represents transfer, and outcome represents rollout] (default: income)

##### Result：
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

#### Check balances

    ./gac-transfer getBalance -ip "http://127.0.0.1:7079" -prikey "./key" -ecosystem "1"

    getBalance : Indicates the order to initiate a transaction

    -ip : Primary node IP address and port number
    -prikey : User private key (private key? storage path)
    [-ecosystem] : Ecosystem ID (default: 1)


##### Result：
    {
        "amount": "480315320000000000000",
        "money": "480315320"
    }

#### Query KeyID and address
    ./gac-transfer getAddress -prikey "./key"

    getAddress : Indicates that the KeyId and address are obtained.

    -prikey : User private key (private key storage path)

##### Result：

    {
        "Address": "0873-2751-5826-4092-2090",
        "KeyId": 8732751582640922090
    }

#### Randomly generate a pair of public and private keys
    ./gac-transfer createPriAndPub

##### Result
    {
        "prikey": "a9a11d8dde1f0a3a88bdde433d8b250807c9b19f9431c12f05448326baee3290",
        "pubkey": "38fe6c442912dbb890cb2182e2f814f3f40654764cf9ef92f169ed59a29afb7d5ec80486b28920a8c3ee7bfc59d3ffd62b10e323f0b088ec0eccee6a22ea3198"
    }
