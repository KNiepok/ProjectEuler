package problem063

import (
	"fmt"
	"math"
)

//The 5-digit number, 16807=7^5, is also a fifth power. Similarly, the 9-digit number, 134217728=8^9, is a ninth power.
//
//How many n-digit positive integers exist which are also an nth power?

func howMany() int {
	s := 0
	for i := 1; i < 10; i++ {
		s += int(1 / (1 - math.Log10(float64(i))))
	}

	return s
}

func Result() {
	fmt.Printf("Number of digits %d\n", howMany())
}
