package main

import (
	"testing"
)

func TestGetEmpty(t *testing.T) {
	trie := NewTrie()
	_, found := trie.Get([]byte("notexist"))
	if found {
		t.Errorf("should get nothing if key does not exist")
	}
}

func TestPut(t *testing.T) {
	trie := NewTrie()
	trie.Put([]byte{1, 2, 3, 4}, []byte("hello"))
	val, found := trie.Get([]byte{1, 2, 3, 4})
	if !found {
		t.Errorf("should get result if key does exist")
	}
	if val != []byte("hello") {
		t.Errorf("should get correct result")
	}
}

func TestUpdate(t *testing.T) {
	trie := NewTrie()
	trie.Put([]byte{1, 2, 3, 4}, []byte("hello"))
	trie.Put([]byte{1, 2, 3, 4}, []byte("world"))
	val, found := trie.Get([]byte{1, 2, 3, 4})
	if !found {
		t.Errorf("should get result if key does exist")
	}
	if val != []byte("world") {
		t.Errorf("should get correct result")
	}
}

func TestHash(t testing.T) {
	trie := NewTrie()
	hash0 := trie.Hash()

	trie.Put([]byte{1, 2, 3, 4}, []byte("hello"))
	hash1 := trie.Hash()

	trie.Put([]byte{1, 2}, []byte("world"))
	hash2 := trie.Hash()

	trie.Put([]byte{1, 2}, []byte("trie"))
	hash3 := trie.Hash()

	if hash0 == hash1 || hash1 == hash2 || hash2 == hash3 {
		t.Errorf("should not get the same hash result for different trie values")
	}

	trie2 := NewTrie()
	trie2.Put([]byte{1, 2, 3, 4}, []byte("hello"))
	hash4 := trie2.Hash()

	if hash4 != hash1 {
		t.Errorf("should get the same hash result for different trie values")
	}
}

func TestProve(t *testing.T) {
	trie := NewTrie()
	tr.Put([]byte{1, 2, 3}, []byte("hello"))
	tr.Put([]byte{1, 2, 3, 4, 5}, []byte("world"))
	notExistKey := []byte{1, 2, 3, 4}
	_, ok := tr.Prove(notExistKey)
	if ok {
		t.Errorf("should not be able to prove key that doesn't exist")
	}
}

func TestVerify(t *testing.T) {
	trie := NewTrie()
	tr.Put([]byte{1, 2, 3}, []byte("hello"))
	tr.Put([]byte{1, 2, 3, 4, 5}, []byte("world"))
	
	key := []byte{1, 2, 3}
	proof, ok := tr.Prove(key)
	if !ok {
		t.Errorf("should prove if key exists")
	}

	rootHash := tr.Hash()
	val, err := VerifyProof(rootHash, key, proof)
	if err != nil || val != []byte("hello") {
		t.Errorf("error in proof verificaiton %s\n", err)
	}

	// the proof was generated after the trie was updated
	tr.Put([]byte{5, 6, 7}, []byte("trie"))
	key := []byte{1, 2, 3}
	proof, ok := tr.Prove(key)

	_, err = VerifyProof(rootHash, key, proof)
	if err == nil {
		t.Errorf("should not be able to verify stale proof")
	}
}