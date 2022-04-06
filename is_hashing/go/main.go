package main

import (
	"fmt"
	"time"
	"strings"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"crypto/sha256"
)

type Hash [32]byte

type Block string

type EmptyBlock struct {
}

type Hashable interface {
	hash() Hash
}

type Node struct {
	left Hashable
	right Hashable
}

func (_ EmptyBlock) hash() Hash {
	return [sha256.Size]byte{}
}

func (b Block) hash() Hash {
	return hash([]byte(b)[:])
}

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

func hash(data []byte) Hash {
	return sha256.Sum256(data)
}

func printTree(node Node) {
	printNode(node, 0)
}

func printNode(node Node, level int) {
	fmt.Printf("%d %s %s\n", level, strings.Repeat(" ", level), node.hash())
	if l, ok := node.left.(Node); ok {
		printNode(l, level + 1)
	} else if l, ok := node.left.(Block); ok {
		fmt.Printf("%d %s %s data %s \n", level + 1, strings.Repeat(" ", level + 1), l.hash(), 1)
	}
	if r, ok := node.right.(Node); ok {
		printNode(r, level + 1)
	} else if r, ok := node.right.(Block); ok {
		fmt.Printf("%d %s %s data %s \n", level + 1, strings.Repeat(" ", level + 1), r.hash(), r)
	}
}

func (n Node) hash() Hash {
	var l, r [sha256.Size]byte
	l = n.left.hash()
	r = n.right.hash()
	return (hash(append(l[:], r[:]...)))
}

func buildTree(parts []Hashable) []Hashable {
	var nodes []Hashable
	var i int
	for i = 0; i < len(parts); i += 2 {
		if i + 1 < len(parts) {
			nodes = append(nodes, Node{left: parts[i], right: parts[i+1]})
		} else {
			nodes = append(nodes, Node{left: parts[i], right: EmptyBlock{}})
		}
	}
	if len(nodes) == 1 {
		return nodes
	} else {
		return buildTree(nodes)
	}
}

func main() {
	s := "dontpanic42dontleaveyourtowelll"

	start := time.Now()
	hmd5 := md5.Sum([]byte(s))
	duration := time.Since(start)

	start1 := time.Now()
	hsha1 := sha1.Sum([]byte(s))
	duration1 := time.Since(start1)

	start2 := time.Now()
	hsha2 := sha256.Sum256([]byte(s))
	duration2 := time.Since(start2)

	fmt.Printf("%v   MD5: %x\n", duration, hmd5)
	fmt.Printf("%v  SHA1: %x\n", duration1, hsha1)
	fmt.Printf("%vSHA256: %x\n", duration2, hsha2)

	printTree(buildTree([]Hashable{Block("a"), Block("B"), Block("c"), Block("d"), Block("e"), Block("f")})[0].(Node))
}