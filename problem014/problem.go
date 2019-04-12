package problem014

import (
	"fmt"
	"math"
)

//The following iterative sequence is defined for the set of positive integers:
//
//n → n/2 (n is even)
//n → 3n + 1 (n is odd)
//
//Using the rule above and starting with 13, we generate the following sequence:
//
//13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
//It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
//
//Which starting number, under one million, produces the longest chain?
//
//NOTE: Once the chain starts the terms are allowed to go above one million.
func longestCollatzChain(maxNumber int) int {
	max := 1
	number := 1
	for i := 1; i < maxNumber; i++ {
		l := collatzChainLength(i)

		if l > max {
			max = l
			number = i
		}
	}

	return number
}

func collatzChainLength(i int) int {
	l := 0

	val := i
	for true {
		l++
		if val == 1 {
			break
		}

		if val%2 == 0 {
			val /= 2
		} else {
			val = 3*val + 1
		}
	}

	return l
}

func Result() {
	fmt.Printf("Number with longest chain for number below 1*10^6 is:  %d\n", longestCollatzChain(1*int(math.Pow10(6))))
}
