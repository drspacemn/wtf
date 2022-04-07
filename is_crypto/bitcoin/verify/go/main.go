package main

import (
	"fmt"
	"os"
	"bytes"
	"strings"
	"math/big"
	"encoding/hex"
	"encoding/json"
	"crypto/sha256"
)

// https://www.blockchain.com/btc/block/00000000000000000000dd97d3f8b6198899f6ea21563dc932df76cb5bf00787
// https://www.blockchain.com/btc/block/00000000000000000007456be0bc3b712c4e4343e4be19ff33b95dcacc13b1d9
// to pull block data: curl https://blockchain.info/rawblock/00000000000000000000dd97d3f8b6198899f6ea21563dc932df76cb5bf00787 > rawBTCHeight730724.json
const hash730725 = "00000000000000000007456be0bc3b712c4e4343e4be19ff33b95dcacc13b1d9"

type Block struct {
	Hash         string   `json:"hash"`
	Ver          int64      `json:"ver"`
	PrevBlock    string   `json:"prev_block"`
	MrklRoot     string   `json:"mrkl_root"`
	Time         int64      `json:"time"`
	Bits         int64      `json:"bits"`
	Nonce        int64    `json:"nonce"`
	NTx          int64      `json:"n_tx"`
	Size         int64      `json:"size"`
	BlockIndex   int64      `json:"block_index"`
	MainChain    bool     `json:"main_chain"`
	Height       int64      `json:"height"`
	ReceivedTime int64      `json:"received_time"`
	RelayedBy    string   `json:"relayed_by"`
	Tx           []Transaction `json:"tx"`
}


func main() {
    rawBlockFile, err := os.ReadFile("../rawBTCHeight730724.json")
	if err != nil {
		panic(err.Error())
	}
	var block Block
	err = json.Unmarshal(rawBlockFile, &block)
	if err != nil {
		panic(err.Error())
	}

	// first we will trust the merkle root
	hash, err := block.OptimisticHash()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Given Hash: \t\t%s\n", block.Hash)
	fmt.Printf("Optimistic Hash: \t%x\n", hash)
	fmt.Println("Match: ", fmt.Sprintf("%x", hash) == block.Hash)
	fmt.Print("\n\n")

	hash, err = block.MorePessimisticHash()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Given Hash: \t\t%s\n", block.Hash)
	fmt.Printf("More Pessimistic Hash: \t%x\n", hash)
	fmt.Println("Match: ", fmt.Sprintf("%x", hash) == block.Hash)
	fmt.Print("\n\n")

	// hash, err = block.MostPessimisticHash()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Printf("Given Hash: \t\t%s\n", block.Hash)
	// fmt.Printf("Most Pessimistic Hash: \t%x\n", hash)
	// fmt.Println("Match: ", fmt.Sprintf("%x", hash) == block.Hash)
	// fmt.Print("\n\n")
}

func (block Block) OptimisticHash() ([]byte, error) {
	header := block.FmtHeader(block.MrklRoot)

	hashBytes := sha256.Sum256(header[:])
	hashBytes = sha256.Sum256(hashBytes[:])

	return  Reverse(hashBytes[:]), nil
}

func (block Block) MorePessimisticHash() ([]byte, error) {
	merkleRoot := block.GetMerkleRootMore()
	fmt.Println("MERKLE: ", merkleRoot)
	// header := block.FmtHeader(merkleRoot)

	// hashBytes := sha256.Sum256(header[:])
	// hashBytes = sha256.Sum256(hashBytes[:])

	// return  Reverse(hashBytes[:]), nil
	return []byte{}, nil
}

func (block Block) GetMerkleRootMore() (hash []byte) {
	// check if transaction array is even
	// - if not replicate the last transaction
	if len(block.Tx) % 2 == 1 {
		block.Tx = append(block.Tx, block.Tx[len(block.Tx)-1])
	}

	var merklePass []*big.Int
	var merkleHold []*big.Int
	for _, tx := range block.Tx {
		byHash := HexToBytes(tx.Hash)
		merkleHold = append(merkleHold, new(big.Int).SetBytes(byHash))
	}
	fmt.Println("HOLD: ", merkleHold)

	for j := 0; j < len(block.Tx)/2; j++ {
		if j % 2 == 1 {
			merkleHold = []*big.Int{}
			for i := 1; i < len(merklePass)/2; i += 2 {
				var buf bytes.Buffer
				_, _ = buf.Write(merklePass[i-1].Bytes())
				_, _ = buf.Write(merklePass[i].Bytes())
				hashBytes := sha256.Sum256(buf.Bytes())
				hashBytes = sha256.Sum256(hashBytes[:])
				hashBig := new(big.Int).SetBytes(hashBytes[:])
				merkleHold = append(merkleHold, hashBig)
				// merklePass = append(merklePass, hashBig)
			}
		} else {
			merklePass = []*big.Int{}
			for i := 1; i < len(merklePass)/2; i += 2 {
				var buf bytes.Buffer
				_, _ = buf.Write(merkleHold[i-1].Bytes())
				_, _ = buf.Write(merkleHold[i].Bytes())
				hashBytes := sha256.Sum256(buf.Bytes())
				hashBytes = sha256.Sum256(hashBytes[:])
				hashBig := new(big.Int).SetBytes(hashBytes[:])
				merklePass = append(merklePass, hashBig)
			}
		}

		fmt.Println("MerklePass: ", merklePass)
		fmt.Println("MerkleHold: ", merkleHold)
	}

	fmt.Println("MERKLE PASS: ", merklePass)
	return hash
}

// func (tx Transaction) Hash() ([]byte, error) {
// 	var buf bytes.Buffer
// 	_, err := buf.Write(Int64ToBytes(block.Ver))

// }

func (tx Transaction) FmtTransaction() ([]byte, error) {

	return []byte{}, nil
}


func (block Block) FmtHeader(merkleRoot string) []byte {
	var buf bytes.Buffer
	_, err := buf.Write(Int64ToBytes(block.Ver))
	_, err = buf.Write(HexToBytes(block.PrevBlock))
	_, err = buf.Write(HexToBytes(merkleRoot))
	_, err = buf.Write(Int64ToBytes(block.Time))
	_, err = buf.Write(Int64ToBytes(block.Bits))
	_, err = buf.Write(Int64ToBytes(block.Nonce))
	if err != nil {
		return buf.Bytes()
	}
	return buf.Bytes()
}

type Transaction struct {
	Hash        string `json:"hash"`
	Ver         int64    `json:"ver"`
	VinSz       int64    `json:"vin_sz"`
	VoutSz      int64    `json:"vout_sz"`
	Size        int64    `json:"size"`
	Weight      int64    `json:"weight"`
	Fee         int64    `json:"fee"`
	RelayedBy   string `json:"relayed_by"`
	LockTime    int64    `json:"lock_time"`
	TxIndex     int64  `json:"tx_index"`
	DoubleSpend bool   `json:"double_spend"`
	Time        int64    `json:"time"`
	BlockIndex  int64    `json:"block_index"`
	BlockHeight int64    `json:"block_height"`
	Inputs      []struct {
		Sequence int64  `json:"sequence"`
		Witness  string `json:"witness"`
		Script   string `json:"script"`
		Index    int64    `json:"index"`
		PrevOut  struct {
			TxIndex           int64    `json:"tx_index"`
			Value             int64    `json:"value"`
			N                 int64  `json:"n"`
			Type              int64    `json:"type"`
			Spent             bool   `json:"spent"`
			Script            string `json:"script"`
			SpendingOutpoints []struct {
				TxIndex int64 `json:"tx_index"`
				N       int64   `json:"n"`
			} `json:"spending_outpoints"`
		} `json:"prev_out"`
	} `json:"inputs"`
	Out []struct {
		Type              int64           `json:"type"`
		Spent             bool          `json:"spent"`
		Value             int64           `json:"value"`
		SpendingOutpoints []interface{} `json:"spending_outpoints"`
		N                 int64           `json:"n"`
		TxIndex           int64         `json:"tx_index"`
		Script            string        `json:"script"`
		Addr              string        `json:"addr,omitempty"`
	} `json:"out"`
}

// IntToHex converts an int64 to a byte array
func Int64ToBytes(number int64) []byte {
    big := new(big.Int)
    big.SetInt64(number)
    return Reverse(big.Bytes())
}

func HexToBytes(hexString string) []byte {
	numStr := strings.Replace(hexString, "0x", "", -1)
	if (len(numStr) % 2) != 0 {
		numStr = fmt.Sprintf("%s%s", "0", numStr)
	}

	by, _ := hex.DecodeString(numStr)
	return Reverse(by)
}

func Reverse(numbers []byte) []byte {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
