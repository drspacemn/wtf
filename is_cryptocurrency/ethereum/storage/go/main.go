package main 

import (
	"fmt"
	"sort"
	"bytes"
	"math/big"
	"encoding/binary"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var emptyCode = crypto.Keccak256([]byte("idk"))

var (
	codehashes = []common.Hash{
		crypto.Keccak256Hash([]byte{0}),
		crypto.Keccak256Hash([]byte{1}),
		crypto.Keccak256Hash([]byte{2}),
		crypto.Keccak256Hash([]byte{3}),
		crypto.Keccak256Hash([]byte{4}),
		crypto.Keccak256Hash([]byte{5}),
		crypto.Keccak256Hash([]byte{6}),
		crypto.Keccak256Hash([]byte{7}),
	}
)

func main() {
	_, _, _, _ = makeAccountTrieWithStorage(3, 3000, true)
	// fmt.Println("ABOVE: ", accTrie, entries, storageTries, storageEntries)

}
type kv struct {
	k, v []byte
}

// Some helpers for sorting
type entrySlice []*kv

func (p entrySlice) Len() int           { return len(p) }
func (p entrySlice) Less(i, j int) bool { return bytes.Compare(p[i].k, p[j].k) < 0 }
func (p entrySlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// makeAccountTrieWithStorage spits out a trie, along with the leafs
func makeAccountTrieWithStorage(accounts, slots int, code bool) (*trie.Trie, entrySlice, map[common.Hash]*trie.Trie, map[common.Hash]entrySlice) {
	var (
		db             = trie.NewDatabase(rawdb.NewMemoryDatabase())
		accTrie, _     = trie.New(common.Hash{}, db)
		entries        entrySlice
		storageTries   = make(map[common.Hash]*trie.Trie)
		storageEntries = make(map[common.Hash]entrySlice)
	)
	// Make a storage trie which we reuse for the whole lot
	var (
		stTrie    *trie.Trie
		stEntries entrySlice
	)

	stTrie, stEntries = makeStorageTrieWithSeed(uint64(slots), 6, db)
	
	stRoot := stTrie.Hash()

	// Create n accounts in the trie
	for i := uint64(1); i <= uint64(accounts); i++ {
		key := key32(i)
		codehash := emptyCode[:]
		if code {
			codehash = getCodeHash(i)
		}
		value, _ := rlp.EncodeToBytes(&types.StateAccount{
			Nonce:    i,
			Balance:  big.NewInt(int64(i)),
			Root:     stRoot,
			CodeHash: codehash,
		})
		fmt.Println("VALUE: ", key, stRoot)
		elem := &kv{key, value}
		accTrie.Update(elem.k, elem.v)
		entries = append(entries, elem)
		// we reuse the same one for all accounts
		storageTries[common.BytesToHash(key)] = stTrie
		storageEntries[common.BytesToHash(key)] = stEntries
	}
	sort.Sort(entries)
	stTrie.Commit(nil)
	accTrie.Commit(nil)
	return accTrie, entries, storageTries, storageEntries
}

// low level account storage
func makeStorageTrieWithSeed(n, seed uint64, db *trie.Database) (*trie.Trie, entrySlice) {
	trie, _ := trie.New(common.Hash{}, db)
	var entries entrySlice
	for i := uint64(1); i <= n; i++ {
		// store 'x' at slot 'x'
		slotValue := key32(i + seed)
		rlpSlotValue, _ := rlp.EncodeToBytes(common.TrimLeftZeroes(slotValue[:]))
		
		slotKey := key32(i)
		key := crypto.Keccak256Hash(slotKey[:])
		fmt.Println("SLOT KEY: ", key)
		fmt.Println("SLOT VAL: ", slotValue)

		elem := &kv{key[:], rlpSlotValue}
		trie.Update(elem.k, elem.v)
		entries = append(entries, elem)
	}
	sort.Sort(entries)
	trie.Commit(nil)
	return trie, entries
}

func key32(i uint64) []byte {
	key := make([]byte, 32)
	binary.LittleEndian.PutUint64(key, i)
	return key
}

// getCodeHash returns a pseudo-random code hash
func getCodeHash(i uint64) []byte {
	h := codehashes[int(i)%len(codehashes)]
	return common.CopyBytes(h[:])
}