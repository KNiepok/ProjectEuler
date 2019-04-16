package problem092

import (
	"fmt"
	"math"
)

//A number chain is created by continuously adding the square of the digits in a number to form a new number until it has been seen before.
//
//For example,
//
//44 → 32 → 13 → 10 → 1 → 1
//85 → 89 → 145 → 42 → 20 → 4 → 16 → 37 → 58 → 89
//
//Therefore any chain that arrives at 1 or 89 will become stuck in an endless loop. What is most amazing is that EVERY starting number will eventually arrive at 1 or 89.
//
//How many starting numbers below ten million will arrive at 89?
func bruteforce() int {
	// every number will eventually end up being reduced to 1 or 89. 
	// so every member of a chain will eventually be 1 or 89.
	// init map of 1 million items
	const unknownEnding = 0
	const endsWith1 = 1
	const endsWith89 = 2
	const million = 10 * 1000 * 1000

	items := make(map[int]int)
	items[1] = endsWith1
	items[89] = endsWith89

	for i := 1; i <= million; i++ {
		// it ending is already defined, ignore the item
		if val, exists := items[i]; exists && val != unknownEnding {
			continue
		}
		// initialize a slice to store chain values
		slice := []int{i}
		// assign the number to local variable
		j := i

		// generate next numbers from the chain and break it when we encounter one that has known ending
		for true {
			slice = append(slice, j)
			if result, exists := items[j]; exists && result != unknownEnding {
				for _, v := range slice {
					items[v] = result
				}
				break
			}
			j = nextItem(j)
		}
	}

	sum := 0

	for i := 1; i <= million; i++ {
		if items[i] == endsWith89 {
			sum++
		}
	}

	return sum
}

func nextItem(i int) int {
	sum := 0
	for i != 0 {
		a := i % 10
		sum += a * a
		i /= 10
	}
	return sum
}

// knowing that for each sextuple of numbers we eventually get the same result, we only need to create those combinations and determine in how many ways we can shuffle them 
func combinatorics() int {
	result := 0
	target := 10 * 1000 * 1000

	cachesize := int(math.Ceil(math.Log10(float64(target))))

	number := make([]int, cachesize)
	i := cachesize - 1

	for true {
		if i == 0 && number[i] == 9 {
			break
		}

		if i == cachesize-1 && number[i] < 9 {
			number[i]++
			result += checkNumber(number)
		} else if number[i] == 9 {
			i--
		} else {
			number[i]++
			for j := i + 1; j < cachesize; j++ {
				number[j] = number[i]
			}
			i = cachesize - 1
			result += checkNumber(number)
		}
	}

	return result
}

func checkNumber(number []int) int {
	factor := 1
	candidate := 0

	for i := len(number) - 1; i >= 0; i-- {
		candidate += factor * number[i]
		factor *= 10
	}

	for candidate != 89 && candidate != 1 {
		candidate = nextItem(candidate)
	}

	if candidate == 89 {
		//calc the number of each digit
		numDigits := [10]int{}

		for i := 0; i < len(number); i++ {
			numDigits[number[i]]++
		}

		result := getFactor(len(number))
		for i := 0; i < len(numDigits); i++ {
			result /= getFactor(numDigits[i])
		}
		return result
	}

	return 0

}

func getFactor(i int) int {
	if i < 1 {
		return 1
	}

	p := 1

	for j := 1; j <= i; j++ {
		p *= j
	}
	return p
}

func Result() {
	fmt.Printf("Elements ending with 89 count (bruteforce): %d\n", combinatorics())
}
