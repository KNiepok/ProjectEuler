package problem002

import "fmt"

//By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
func evenFibonacciSum(limit int) int {
	values := []int{1, 1}

	for true {
		newValue := values[len(values)-1] + values[len(values)-2]

		if newValue >= limit {
			break
		}
		values = append(values, newValue)
	}

	sum := 0

	for _, val := range values {
		if val%2 == 0 {
			sum += val
		}
	}
	return sum
}

func Result() {
	fmt.Printf("Result for 1000 is %d\n", evenFibonacciSum(4*1000*1000))
}
