package main

import (
	"fmt"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/crypto"
)

var indices = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "[17]"}

type Trie struct {
	root Node
}

type Node interface {
	Hash() []byte
}

type LeafNode struct {
	Path []byte
	Value []byte
}

func (l LeafNode) Hash() []byte {
	return crypto.Keccak256(l.Value)
}

type BranchNode struct {
	Children [16]Node
	Value    []byte
}

func (b BranchNode) Hash() []byte {
	return crypto.Keccak256(b.Value)
}

type ExtensionNode struct {
	Path []byte
	Next Node
}

func (e ExtensionNode) Hash() []byte {
	return crypto.Keccak256(e.Path)
}

func main() {
	idx, _ := rlp.EncodeToBytes(uint(0))
	idx1, _ := rlp.EncodeToBytes(uint(1))
	fmt.Printf("ZERO INDEX: %x %x\n", idx, idx1)

	_ = pullBlock("0xA1A489")
	// fmt.Printf("BLOCK: %+v\n", block)

	// fmt.Println("Transaction Root: ", block.MrklRoot)
	// tx := pullTx(block.Txids[0])
	// fmt.Printf("TX: %+v\n", tx)
}

func NewTrie() *Trie {
	return &Trie{}
}

func IsEmptyNode(node Node) bool {
	return node == nil
}

func (t *Trie) Insert(key, value []byte) {
	node := &t.root
	for {
		if IsEmptyNode(*node) {
			leaf := &LeafNode{key, value}
			*node = Node(leaf)
			return
		}

		if leaf, ok := *node.(*LeafNode); ok {
			matchLen := prefixLen(key, leaf.Value)
			
		}
	}
}