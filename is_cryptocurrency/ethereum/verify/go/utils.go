package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/crypto/sha3"
	// "github.com/ethereum/go-ethereum/rlp"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Nibble byte

type Payload struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type Proof map[string][]byte

func (t *Trie) Prove(key []byte) (Proof, bool) {
	proof := make(map[string][]byte)
	node := t.root
	nibbles := FromBytes(key)

	for {
		proof[fmt.Sprintf("%x", Hash(node))] = Serialize(node)

		if IsEmptyNode(node) {
			return nil, false
		}

		switch n := node.(type) {
		case *LeafNode:
			matched := PrefixMatchedLen(n.Rest, nibbles)
			if matched != len(n.Rest) || matched != len(nibbles) {
				return nil, false
			}
			return proof, true

		case *BranchNode:
			if len(nibbles) == 0 {
				return proof, n.HasValue()
			}

			b, remaining := nibbles[0], nibbles[1:]
			nibbles = remaining
			node = n.Branches[b]

		case *ExtensionNode:
			matched := PrefixMatchedLen(n.Shared, nibbles)
			if matched < len(n.Shared) {
				return nil, false
			}

			nibbles = nibbles[matched:]
			node = n.Next

		}
	}
}

func Verify(rootHash, key []byte, proof Proof) ([]byte, error) {
	key = keybytesToHex(key)
	wantHash := rootHash

	for i := 0; ; i++ {
		val, ok := proof[fmt.Sprintf("%x", wantHash)]
		if !ok {
			return nil, fmt.Errorf("superman no home")
		}

		rawNode, err := decodeNode(wantHash[:], val)
		if err != nil {
			return nil, fmt.Errorf("superman no home")
		}

		rest, node := get(rawNode, key, true)

		switch n := node.(type) {
		case nil:
			return nil, nil
		case hashNode:
			key = rest
			copy(wantHash[:], n)
		case valueNode:
			return n, nil
		}
	}
}

func (t *Trie) Walk(nLeaves int) {
	node := t.root
	leaf := [][]Nibble{}
	leaves := 0
	for {
		if leaves >= nLeaves {
			return
		}
		if IsEmptyNode(node) {
			fmt.Println("Empty Node")
			return
		}

		switch n := node.(type) {
		case *LeafNode:
			fmt.Println("SHOULDN'T GET HERE")

		case *BranchNode:
			fmt.Printf("Branch(%v):\n[", leaf)
			for i, elem := range n.Branches {
				if elem != nil {
					if _, ok := elem.(*BranchNode); ok {
						fmt.Printf("%x:   BRANCH   | ", Indeces[i])
					} else {
						fmt.Printf("%x:   LEAF   | ", Indeces[i])
					}
				} else {
					fmt.Printf("%x: - | ", Indeces[i])
				}
			}
			fmt.Println("]")
			for j, elem := range n.Branches {
				if elem != nil {
					if ileaf, ok := elem.(*LeafNode); ok {
						raw := leaf
						raw = append(raw, []Nibble{Nibble(Indeces[j])})
						if len(ileaf.Rest) > 0 {
							raw = append(raw, ileaf.Rest)
						}
						spacer := fmt.Sprintf("%s%s", strings.Repeat("\t", j), strings.Repeat(" ", j/2))
						fmt.Printf("%sLeaf(%x):\n%s%v\n%s--> %x\n\n", spacer, j, spacer, raw, spacer, ileaf.Value[:4])
						leaves++
					}
				}
			}
			for j, elem := range n.Branches {
				if elem != nil {
					if _, ok := elem.(*BranchNode); ok {
						leaf = append(leaf, []Nibble{Nibble(Indeces[j])})
						node = elem
						break
					}
				}
			}

		case *ExtensionNode:
			fmt.Printf("%sExtension: nibbles %v\n", strings.Repeat("\t\t\t\t\t", (nLeaves-leaves)/2), n.Shared)
			leaf = append(leaf, n.Shared)
			node = n.Next
		}
	}
}

func pullBlock(blockNum string) (types.Header, [][]byte) {
	if os.Getenv("INFURA_URL") == "" {
		panic("Please provide valid url and set via: 'export INFURA_URL=<URL>'")
	}
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

	var res [][]byte
	for i := 0; i < int(txNum); i++ {
		var tx *types.Transaction
		if err := client.Call(&tx, "eth_getTransactionByBlockNumberAndIndex", blockNum, fmt.Sprintf("0x%x", i)); err != nil {
			panic(err.Error())
		}

		var buf bytes.Buffer
		err = tx.EncodeRLP(&buf)
		res = append(res, buf.Bytes())
	}

	var header types.Header
	if err := client.Call(&header, "eth_getBlockByNumber", blockNum, false); err != nil {
		panic(err.Error())
	}

	return header, res
}

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

func IsNibble(nibble byte) bool {
	n := int(nibble)
	// 0-9 && a-f
	return n >= 0 && n < 16
}

func FromNibbleByte(n byte) (Nibble, error) {
	if !IsNibble(n) {
		return 0, fmt.Errorf("non-nibble byte: %v", n)
	}
	return Nibble(n), nil
}

// nibbles contain one nibble per byte
func FromNibbleBytes(nibbles []byte) ([]Nibble, error) {
	ns := make([]Nibble, 0, len(nibbles))
	for _, n := range nibbles {
		nibble, err := FromNibbleByte(n)
		if err != nil {
			return nil, fmt.Errorf("contains non-nibble byte: %w", err)
		}
		ns = append(ns, nibble)
	}
	return ns, nil
}

func FromByte(b byte) []Nibble {
	return []Nibble{
		Nibble(byte(b >> 4)),
		Nibble(byte(b % 16)),
	}
}

func FromBytes(bs []byte) []Nibble {
	ns := make([]Nibble, 0, len(bs)*2)
	for _, b := range bs {
		ns = append(ns, FromByte(b)...)
	}
	return ns
}

func FromString(s string) []Nibble {
	return FromBytes([]byte(s))
}

// ToPrefixed add nibble prefix to a slice of nibbles to make its length even
// the prefix indicts whether a node is a leaf node.
func ToPrefixed(ns []Nibble, isLeafNode bool) []Nibble {
	// create prefix
	var prefixBytes []Nibble
	// odd number of nibbles
	if len(ns)%2 > 0 {
		prefixBytes = []Nibble{1}
	} else {
		// even number of nibbles
		prefixBytes = []Nibble{0, 0}
	}

	// append prefix to all nibble bytes
	prefixed := make([]Nibble, 0, len(prefixBytes)+len(ns))
	prefixed = append(prefixed, prefixBytes...)
	for _, n := range ns {
		prefixed = append(prefixed, Nibble(n))
	}

	// update prefix if is leaf node
	if isLeafNode {
		prefixed[0] += 2
	}

	return prefixed
}

// ToBytes converts a slice of nibbles to a byte slice
// assuming the nibble slice has even number of nibbles.
func ToBytes(ns []Nibble) []byte {
	buf := make([]byte, 0, len(ns)/2)

	for i := 0; i < len(ns); i += 2 {
		b := byte(ns[i]<<4) + byte(ns[i+1])
		buf = append(buf, b)
	}

	return buf
}

func PrefixMatchedLen(node1 []Nibble, node2 []Nibble) int {
	matched := 0
	for i := 0; i < len(node1) && i < len(node2); i++ {
		n1, n2 := node1[i], node2[i]
		if n1 == n2 {
			matched++
		} else {
			break
		}
	}

	return matched
}
