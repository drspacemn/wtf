package main

import (
	"fmt"
	"encoding/hex"
	"crypto/aes"
	"crypto/cipher"
	"math/rand"

)

var count = 0
var stringToEncrypt = "ssssshhhhhhhhhhh. I'm vewwwwwwy vewwwwwwwyyyyy secretttttt"

func main() {
	// generate random 32 byte key
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	wtf_is(bytes)

	// save  in a vewwwwwwwwy safe place
	saveKeyString := hex.EncodeToString(bytes)
	wtf_is(saveKeyString)

	key, _ := hex.DecodeString(saveKeyString)
	plaintext := []byte(stringToEncrypt)
	wtf_is(plaintext)
	
	block, _ := aes.NewCipher(key)
	wtf_is(block)
	
	aesGCM, _ := cipher.NewGCM(block)
	wtf_is(aesGCM)

	nonce := make([]byte, aesGCM.NonceSize())
	wtf_is(nonce)

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	wtf_is(ciphertext)
	wtf_is(plaintext)

	encryptedString := fmt.Sprintf("%x", ciphertext)
	wtf_is(encryptedString)

	// MEANWHILE ON THE SERVER SIDE

	key, _ = hex.DecodeString(saveKeyString)
	enc, _ := hex.DecodeString(encryptedString)
	
	block, _ = aes.NewCipher(key)

	aesGCM, _ = cipher.NewGCM(block)

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext = enc[:nonceSize], enc[nonceSize:]
	wtf_is(nonce, ciphertext)

	plaintext, _ = aesGCM.Open(nil, nonce, ciphertext, nil)
	wtf_is(string(plaintext))
}

func wtf_is(is_this ...interface{}) {
	count++
	fmt.Printf("%d) wtf is this: %v\n\n", count, is_this)
}
