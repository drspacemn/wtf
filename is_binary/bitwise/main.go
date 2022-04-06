package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a uint = 60	/* 60 = 0011 1100 */  
	var b uint = 13	/* 13 = 0000 1101 */
	var c uint = 0          
 
	c = a & b       /* 12 = 0000 1100 */ 
	fmt.Printf("Line 1 - Value of c is %d %b\n", c, c)
 
	c = a | b       /* 61 = 0011 1101 */
	fmt.Printf("Line 2 - Value of c is %d %b\n", c, c)
 
	c = a ^ b       /* 49 = 0011 0001 */
	fmt.Printf("Line 3 - Value of c is %d %b\n", c, c)
 
	c = a << 2     /* 240 = 1111 0000 */
	fmt.Printf("Line 4 - Value of c is %d %b\n", c, c)

	pub, _ := new(big.Int).SetString("2754806153357301156380357983574496185342034785016738734224771556919270737441", 10)
	shift := new(big.Int)
	shift = shift.Exp(big.NewInt(2), big.NewInt(128), nil)
	
	highshift := new(big.Int)
	highshift = highshift.Exp(big.NewInt(2), big.NewInt(123), nil)
	highshift = highshift.Mul(highshift, shift)
	highshift = highshift.Sub(highshift, shift)
	fmt.Printf("HIGH: %b %v\n", highshift, highshift.BitLen())

	m := new(big.Int)
	m = m.Sub(shift, big.NewInt(1))

	low := new(big.Int)
	low = low.And(pub, m)

	high := new(big.Int)
	high = high.And(pub, highshift)
	fmt.Println("HIGHSHIFT: ", highshift)
	fmt.Printf("LOW: %b %v\n", low, low)
	fmt.Printf("HIGH: %b %v\n", high, high.Div(high, shift))
}