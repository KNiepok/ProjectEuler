package problem007

import "fmt"

//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//What is the 10 001st prime number?
func nthPrime(i int) int {
	p := getPrimes(i)
	return p[i-1]
}

// reuse problem 5 function, with changed loop condition
func getPrimes(upperLimit int) []int {
	var isPrime bool
	var j int
	primes := []int{2}

	for i := 3; len(primes) < upperLimit; i += 2 {
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
	fmt.Printf("10001st prime: %d\n", nthPrime(10001))
}
