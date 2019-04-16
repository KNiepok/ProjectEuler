package problem005

import (
	"fmt"
	"math"
)

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
func smallestNumberDivisible(limit int) int {
	primes := getPrimes(limit)

	result := 1

	for i := 0; i < len(primes); i++ {
		a := int(math.Floor(math.Log(float64(limit)) / math.Log(float64(primes[i]))))
		result = result * int(math.Pow(float64(primes[i]), float64(a)))
	}

	return result
}

func getPrimes(upperLimit int) []int {

	var isPrime bool
	var j int
	primes := []int{2}

	for i := 3; i <= upperLimit; i += 2 {
		j = 0
		isPrime = true

		for primes[j]*primes[j] <= i {
			if i%primes[j] == 0 {
				isPrime = false
				break
			}
			j++
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

func Result() {
	fmt.Printf("smallest nunber divisible: %d\n", smallestNumberDivisible(20))
}
