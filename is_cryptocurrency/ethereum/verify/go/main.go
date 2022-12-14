package main

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	trie := NewTrie()
	k1, _ := rlp.EncodeToBytes("a711355")
	trie.Put(k1, []byte("45.0 ETH"))
	k2, _ := rlp.EncodeToBytes("a77d4337")
	trie.Put(k2, []byte("1.00 WEI"))
	k3, _ := rlp.EncodeToBytes("a779365")
	trie.Put(k3, []byte("1.1 ETH"))
	k4, _ := rlp.EncodeToBytes("a77d397")
	trie.Put(k4, []byte("0.12 ETH"))

	proof, ok := trie.Prove(k3)
	if !ok {
		panic(fmt.Errorf("could not prove valid key"))
	}
	val, err := Verify(trie.Hash(), k3, proof)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Validity Proof length: ", len(proof))
	fmt.Printf("Validity Proof Verified: true %v\n", val)

	trie = NewTrie()
	// header, txs := pullBlock("0xA1A489")
	header, txs := pullBlock("0xA1A48A")

	for i, tx := range txs {
		key, err := rlp.EncodeToBytes(uint(i))
		if err != nil {
			panic(err.Error())
		}
		trie.Put(key, tx)
	}

	trie.Walk(len(txs))
	fmt.Printf("Computed: \t0x%x\n", trie.Hash())
	fmt.Printf("Tx Hash: \t0x%x\n", header.TxHash)

	key, _ := rlp.EncodeToBytes(uint(3))
	proof, ok = trie.Prove(key)
	if !ok {
		panic(fmt.Errorf("could not prove valid key"))
	}
	val, err = Verify(trie.Hash(), key, proof)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Validity Proof length: ", len(proof))
	fmt.Printf("Validity Proof Verified: true %v\n", val)
}

func (t *Trie) Get(key []byte) (out []byte, ok bool) {
	node := t.root
	nibbles := FromBytes(key)
	common := 0
	for {
		if IsEmptyNode(node) {
			return nil, false
		}

		switch n := node.(type) {
		case *LeafNode:
			matched := PrefixMatchedLen(n.Rest, nibbles[common:])
			if matched != len(nibbles[common:]) || matched != len(n.Rest) {
				return nil, false
			}
			return n.Value, true

		case *BranchNode:
			idx := bytes.IndexByte(Indeces, byte(nibbles[common]))
			if n.Branches[idx] != nil && len(nibbles[common:]) > 1 {
				common++
				node = n.Branches[idx]
				continue
			}
			return nil, false

		case *ExtensionNode:
			matched := PrefixMatchedLen(n.Shared, nibbles[common:])
			if matched == len(n.Shared) {
				common += len(n.Shared)
				node = n.Next
				continue
			}
			return nil, false

		default:
			panic("unknown node type")
		}
	}
}

func (t *Trie) Put(key, val []byte) {
	node := &t.root
	nibbles := FromBytes(key)
	common := 0
	for {
		if IsEmptyNode(*node) {
			*node = &LeafNode{
				Rest:  nibbles,
				Value: val,
			}
			return
		}

		switch n := (*node).(type) {
		case *LeafNode:
			matched := PrefixMatchedLen(n.Rest, nibbles[common:])
			branch := &BranchNode{}
			idx := bytes.IndexByte(Indeces, byte(nibbles[common+matched]))
			if matched > 0 {
				ext := &ExtensionNode{
					Shared: n.Rest[:matched],
				}
				branch.Branches[idx] = &LeafNode{nibbles[common+matched+1:], val}
				branch.Branches[n.Rest[matched]] = &LeafNode{n.Rest[matched+1:], n.Value}
				ext.Next = branch
				*node = ext
				return
			}

			if len(nibbles[common:]) == 1 {
				branch.Branches[idx] = &LeafNode{[]Nibble{}, val}
			} else {
				branch.Branches[idx] = &LeafNode{nibbles[common+1:], val}
			}
			if len(n.Rest) == 1 {
				branch.Branches[n.Rest[0]] = &LeafNode{[]Nibble{}, n.Value}
			} else {
				branch.Branches[n.Rest[0]] = &LeafNode{n.Rest[1:], n.Value}
			}
			*node = branch
			return
		case *BranchNode:
			idx := bytes.IndexByte(Indeces, byte(nibbles[common]))
			if n.Branches[idx] == nil {
				if len(nibbles[common:]) == 1 {
					n.Branches[idx] = &LeafNode{[]Nibble{}, val}
					return
				}

				n.Branches[idx] = &LeafNode{nibbles[common+1:], val}
				return
			}
			common++
			node = &n.Branches[idx]
		case *ExtensionNode:
			matched := PrefixMatchedLen(n.Shared, nibbles[common:])
			if matched < len(n.Shared) {
				branch := &BranchNode{}
				branch.Branches[nibbles[common]] = &LeafNode{nibbles[common+1:], val}
				branch.Branches[n.Shared[matched]] = n.Next
				*node = branch
				return
			}
			common += len(n.Shared)

			node = &n.Next
		default:
			panic("unknown node type")
		}
	}
}
