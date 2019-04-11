package problem25

import (
	"fmt"
	"math/big"
)

var fibs map[int]*big.Int

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
