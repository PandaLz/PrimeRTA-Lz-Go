package main

import (
	"fmt"
	"math/big"
)

func main() {

	// one := big.NewInt(1)
	//test single number
	var zero, two big.Int
	var numToTest, numSqrt, numSqrtReminder, sqrtRemMod big.Int
	var isPair bool = false

	// set the number to test
	numToTest.SetInt64(85)
	zero.SetInt64(0)
	two.SetInt64(2)
	// or one := big.NewInt(1)

	// get Square root of number to test
	numSqrt.Sqrt(&numToTest)

	// get reminder of number to test / square root
	numSqrtReminder.Mod(&numToTest, &numSqrt)

	// Know if numSqrt is odd
	sqrtRemMod.Mod(&numSqrt, &two)

	if 0 == sqrtRemMod.Cmp(&zero) {
		isPair = true
	}

	fmt.Printf("Number to test       : %v\n", numToTest)
	fmt.Printf("Number Square root   : %v\n", numSqrt)
	fmt.Printf("Square root reminder : %v\n", numSqrtReminder)
	fmt.Printf("Mod of square root reminder  : %v\n", sqrtRemMod)
	fmt.Printf("Square root is pair  : %v\n", isPair)

}
