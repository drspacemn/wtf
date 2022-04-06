package main

import (
	"fmt"
)

// NewCipher
// NewGCM
// NonceSize
// Open

// aesCipherGCM implements crypto/cipher.gcmAble so that crypto/cipher.NewGCM
// will use the optimised implementation in this file when possible. Instances
// of this type only exist when hasGCMAsm returns true.
type aesCipherGCM struct {
	aesCipherAsm
}

func (c *aesCipherGCM) NewGCM(nonceSize, tagSize int) (cipher.AEAD, error) {
	g := &gcmAsm{ks: c.enc, nonceSize: nonceSize, tagSize: tagSize}
	gcmAesInit(&g.productTable, g.ks)
	return g, nil
}