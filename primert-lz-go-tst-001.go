package main

import (
	"fmt"
	"math/big"
)

func main() {

	// one := big.NewInt(1)
	//test single number
	var zero, one, two big.Int
	var numToTest, numSqrt, numSqrtReminder, sqrtRemMod big.Int
	var halfNumSqrt, midRem big.Int
	var isPair bool = false

	// set the number to test
	numToTest.SetInt64(85)

	zero.SetInt64(0)
	one.SetInt64(1)
	two.SetInt64(2)
	// or one := big.NewInt(1)

	// get Square root of number to test
	numSqrt.Sqrt(&numToTest)

	// get reminder of number to test / square root
	// numSqrtReminder.Mod(&numToTest, &numSqrt)
	// missing math/big Sqrt with reminder method
	// calculating from number to test - square root squared...
	numSqrtReminder.Mul(&numSqrt, &numSqrt)
	numSqrtReminder.Sub(&numToTest, &numSqrtReminder)

	// Know if numSqrt is pair
	sqrtRemMod.Mod(&numSqrt, &two)

	if 0 == sqrtRemMod.Cmp(&zero) {
		isPair = true
	}

	// process according to pair test of numSqrt
	if isPair {
		// get the half of the reminder
		halfNumSqrt.Div(&numSqrt, &two)
		// get reminder for middle column
		midRem.Add(&numSqrtReminder, &one)
		midRem.Div(&midRem, &two)
	} else {
		// the correct way would be -1 / 2
		halfNumSqrt.Div(&numSqrt, &two)
		// get reminder for middle column
		midRem.Div(&numSqrtReminder, &two)
	}

	fmt.Printf("------------------------------------\n")
	fmt.Printf("-------Initial data required--------\n")
	fmt.Printf("------------------------------------\n")
	fmt.Printf("Number to test               : %v\n", numToTest.String())
	fmt.Printf("Number Square root           : %v\n", numSqrt.String())
	fmt.Printf("Square root reminder         : %v\n", numSqrtReminder.String())
	fmt.Printf("Mod2 of square root reminder : %v\n", sqrtRemMod.String())
	fmt.Printf("Square root is pair          : %v\n", isPair)
	fmt.Printf("------------------------------------\n")
	fmt.Printf("Half square root             : %v\n", halfNumSqrt.String())
	fmt.Printf("mid reminder                 : %v\n", midRem.String())
	fmt.Printf("------------------------------------\n")

}
