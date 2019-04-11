package problem025

import (
	"fmt"
	"math/big"
)

var fibs map[int]*big.Int

//The Fibonacci sequence is defined by the recurrence relation:
//
//Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
//Hence the first 12 terms will be:
//
//F1 = 1
//F2 = 1
//F3 = 2
//F4 = 3
//F5 = 5
//F6 = 8
//F7 = 13
//F8 = 21
//F9 = 34
//F10 = 55
//F11 = 89
//F12 = 144
//The 12th term, F12, is the first term to contain three digits.
//
//What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
func nDigitFibonacciNumber(n int) int {
	initMap()

	i := 2
	for true {
		i++
		fib := fib(i)

		if len(fib.String()) >= n {
			return i
		}
	}

	return i
}

func initMap() {
	fibs = make(map[int]*big.Int)
	fibs[1] = big.NewInt(1)
	fibs[2] = big.NewInt(1)
}

func fib(position int) *big.Int {
	if val, ok := fibs[position]; ok {
		return val
	}

	fib := new(big.Int).Add(fib(position-2), fib(position-1))
	fibs[position] = fib
	return fib
}

func Result() {
	fmt.Printf("First 1000 digit number is %d\n", nDigitFibonacciNumber(1000))
}
