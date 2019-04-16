package problem010

import (
	"fmt"
	"math"
)

//The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//Find the sum of all the primes below two million.
func sumOfPrimesBelow(value int) int {
	// init slice with all numbers and boolean -> true if is a prime (except for 0 and 1)
	primes := make([]bool, value+1)
	for i := 0; i < value+1; i++ {
		primes[i] = true
	}

	// determine if number is prime
	for i := 2; i*i <= value; i++ {
		if primes[i] == true {
			for j := i * 2; j <= value; j += i {
				primes[j] = false // cross
			}
		}

	}

	// sum primes
	sum := 0
	for k, v := range primes {
		if v == true && k > 1 {
			sum += k
		}
	}

	return sum
}

func Result() {
	fmt.Printf("Sum of primes below 2*10^6 is:  %d\n", sumOfPrimesBelow(2*int(math.Pow10(6))))
}
