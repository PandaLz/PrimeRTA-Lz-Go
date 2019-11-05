package main

import (
	"fmt"
	"math/big"
)

func main() {

	//test single number
	var numToTest, numSqrt big.Int
	numToTest.SetInt64(65)
	// or one := big.NewInt(1)
	numSqrt.Sqrt(&numToTest)

	fmt.Printf("Number to test     : %v\n", numToTest)
	fmt.Printf("Number Square root : %v\n", numSqrt)

}
