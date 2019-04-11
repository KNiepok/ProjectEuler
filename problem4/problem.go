package problem4

import (
	"fmt"
	"math"
)

//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//Find the largest palindrome made from the product of two 3-digit numbers.
func largestPalindrome(digits int) int {
	m := int(math.Pow10(digits)) - 1
	n := m

	var palindromes []int
	for i := m; i > int(math.Pow10(digits-1)); i-- {
		for j := n; j > int(math.Pow10(digits-1)); j-- {
			result := j * i
			if isPalindrome(result) {
				palindromes = append(palindromes, result)
			}
		}
	}

	max := 0
	for _, v := range palindromes {
		if v > max {
			max = v
		}
	}

	return max
}

func isPalindrome(number int) bool {
	temp := number
	reverse := 0

	for true {
		remainder := number % 10
		reverse = reverse*10 + remainder
		number /= 10
		if number == 0 {
			break
		}
	}

	return temp == reverse
}

func Result() {
	fmt.Printf("Result for 3 digits is %d\n", largestPalindrome(3))
}
