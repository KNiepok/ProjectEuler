package problem006

import "fmt"

//The sum of the squares of the first ten natural numbers is,
//12 + 22 + ... + 102 = 385
//The square of the sum of the first ten natural numbers is,
//(1 + 2 + ... + 10)2 = 552 = 3025
//Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
func sumSquareDifference(numbers int) int {
	return squareSum(numbers) - sumOfSquares(numbers)
}

func sumOfSquares(numbers int) int {
	sum := 0

	for i := 1; i <= numbers; i++ {
		sum += i * i
	}

	return sum
}

func squareSum(numbers int) int {
	sum := 0

	for i := 1; i <= numbers; i++ {
		sum += i
	}

	return sum * sum
}

func Result() {
	fmt.Printf("difference: %d\n", sumSquareDifference(100))
}

