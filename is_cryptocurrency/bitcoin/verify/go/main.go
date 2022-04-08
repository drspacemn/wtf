package main

import (
	"fmt"
	"os"
	"math"
	"bytes"
	"strings"
	"math/big"
	"encoding/json"
	"crypto/sha256"
)

// https://www.blockchain.com/btc/block/00000000000000000000dd97d3f8b6198899f6ea21563dc932df76cb5bf00787
// to pull block data: curl https://blockchain.info/rawblock/00000000000000000000dd97d3f8b6198899f6ea21563dc932df76cb5bf00787 > rawBTCHeight730724.json

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
	hash, err := block.HashBlock()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Given Hash: \t\t%s\n", block.Hash)
	fmt.Printf("Calculated Hash: \t%x\n", hash)
	fmt.Println("Match: ", fmt.Sprintf("%x", hash) == block.Hash)
	fmt.Print("\n\n")

	root := block.GetMerkleRoot()
	fmt.Println("Given Merkle Root: ", block.MrklRoot)
	fmt.Printf("Calculated Merkle Root: %x\n", root)
	fmt.Println("Match: ", fmt.Sprintf("%x", root) == block.MrklRoot)
	fmt.Print("\n\n")
}

func (block Block) HashBlock() ([]byte, error) {
	header := block.FmtHeader()

	hashBytes := sha256.Sum256(header[:])
	hashBytes = sha256.Sum256(hashBytes[:])

	return  Reverse(hashBytes[:]), nil
}

func (block Block) GetMerkleRoot() (root *big.Int) {
	var merkleHold []*big.Int
	for _, tx := range block.Tx {
		h1 := sha256.Sum256(HexToBytes(tx.Hash)[:])

		merkleHold = append(merkleHold, new(big.Int).SetBytes(h1[:]))
	}
	if len(merkleHold) % 2 == 1 {
		merkleHold = append(merkleHold, new(big.Int).Set(merkleHold[len(merkleHold)-1]))
	}

	rounds := int(math.Ceil(math.Log2(float64(len(merkleHold)/2))))
	fmt.Println("Num leaves: ", len(merkleHold))
	fmt.Println("Num basenodes: ", len(merkleHold)/2)
	fmt.Println("Rounds: ", rounds)

	var merklePass []*big.Int
	for i := 0; i < rounds; i++ {
		if i % 2 == 0 {
			for j := 0; j <= (len(merkleHold)-2); j += 2 {
				var buf []byte
				buf = append(buf, merkleHold[j].Bytes()...)
				buf = append(buf, merkleHold[j+1].Bytes()...)
				merklePass = append(merklePass, hashFunc(buf))
			}

			fmt.Printf("Round %d: %s %s %s %s\n", i, strings.Repeat("  ", i), merklePass[0].String()[:4], strings.Repeat("....", rounds - i), merklePass[len(merklePass)-1].String()[:4])
			merkleHold = []*big.Int{}
		} else {
			for j := 0; j <= (len(merklePass)-2); j += 2 {
				var buf []byte
				buf = append(buf, merklePass[j].Bytes()...)
				buf = append(buf, merklePass[j+1].Bytes()...)
				merkleHold = append(merkleHold, hashFunc(buf))
			}

			fmt.Printf("Round %d: %s %s %s %s\n", i, strings.Repeat("  ", i), merkleHold[0].String()[:4], strings.Repeat("....", rounds - i), merkleHold[len(merkleHold)-1].String()[:4])
			merklePass = []*big.Int{}
		}
	}
	fmt.Print()

	if len(merkleHold) > 0 {
		bigRoot := Reverse(merkleHold[0].Bytes())
		return new(big.Int).SetBytes(bigRoot)
	} else if len(merklePass) > 0 {
		bigRoot := Reverse(merkleHold[0].Bytes())
		return new(big.Int).SetBytes(bigRoot)
	}
	return root
}

func hashFunc(data []byte) *big.Int {
	hash := sha256.Sum256(data)
	hash = sha256.Sum256(hash[:])
	return new(big.Int).SetBytes(hash[:])
}

func (block Block) FmtHeader() []byte {
	var buf bytes.Buffer
	_, err := buf.Write(Int64ToBytes(block.Ver))
	_, err = buf.Write(HexToBytes(block.PrevBlock))
	_, err = buf.Write(HexToBytes(block.MrklRoot))
	_, err = buf.Write(Int64ToBytes(block.Time))
	_, err = buf.Write(Int64ToBytes(block.Bits))
	_, err = buf.Write(Int64ToBytes(block.Nonce))
	if err != nil {
		return buf.Bytes()
	}
	return buf.Bytes()
}
