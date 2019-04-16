package problem003

import (
	"fmt"
	"math"
)

//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?
func largestPrimeFactor(value int) int {
	maxPrime := -1

	for value%2 == 0 {
		value >>= 1
	}

	for i := 3; float64(i) <= math.Sqrt(float64(value)); i += 2 {
		for value%i == 0 {
			maxPrime = i
			value = value/i
		}
 	}

	if value > 2 {
		maxPrime = value
	}

	return maxPrime
}


func Result() {
	fmt.Printf("Result for 600851475143 is %d\n", largestPrimeFactor(600851475143))
}
