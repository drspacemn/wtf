package main

import (
	"os"
	"fmt"
	"bytes"
	"strconv"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/core/types"
)

type Payload struct {
	JsonRpc string `json:"jsonrpc"`
	Method string `json:"method"`
	Params []interface{} `json:"params"`
	Id int `json:"id"`
}

func pullBlock(blockNum string) types.Header {
	client, err := rpc.DialHTTP(os.Getenv("INFURA_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer client.Close()

	var txNumStr string
	if err := client.Call(&txNumStr, "eth_getBlockTransactionCountByNumber", blockNum); err != nil {
		panic(err.Error())
	}
	txNum, _ := strconv.ParseInt(txNumStr[2:], 16, 64)
	fmt.Println("TXNUM: ", txNum)

	for i := 0; i < int(txNum); i++ {
		var tx *types.Transaction
		if err := client.Call(&tx, "eth_getTransactionByBlockNumberAndIndex", blockNum, fmt.Sprintf("0x%x", i)); err != nil {
			panic(err.Error())
		}

		var buf bytes.Buffer
		err = tx.EncodeRLP(&buf)
		fmt.Printf("BUF: %+x\n", buf.Bytes())
	}

	var header types.Header
	if err := client.Call(&header, "eth_getBlockByNumber", blockNum, false); err != nil {
		panic(err.Error())
	}
	// fmt.Println("HEADER: ", header)

	return header
}

func prefixLen(a, b []byte) int {
	var i, length = 0, len(a)
	if len(b) < length {
		length = len(b)
	}
	for ; i < length; i++ {
		if a[i] != b[i] {
			break
		}
	}
	return i
}

// func pullTx(txHash string) *types.Transaction {
// 	client := &http.Client {}
// 	req, err := http.NewRequest(http.MethodGet, ETH_EXPLORER_URL + "/txs/" + txHash, nil)
// 	if err != nil {
// 	  panic(err.Error())
// 	}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	var tx *types.Transaction
// 	err = tx.UnmarshalBinary(body)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return tx
// }