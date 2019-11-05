package main

import (
	"fmt"
	"math/big"
)

func main() {

	//test single number
	var numToTest big.Int
	numToTest.SetInt64(65)
	// or one := big.NewInt(1)

	fmt.Printf("Number to test : %v\n", numToTest)

}
