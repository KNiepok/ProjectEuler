package problem055

import "fmt"

//If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.
//Not all numbers produce palindromes so quickly. For example,
//349 + 943 = 1292,
//1292 + 2921 = 4213
//4213 + 3124 = 7337
//That is, 349 took three iterations to arrive at a palindrome.
//Although no one has proved it yet, it is thought that some numbers, like 196, never produce a palindrome. A number that never forms a palindrome through the reverse and add process is called a Lychrel number. Due to the theoretical nature of these numbers, and for the purpose of this problem, we shall assume that a number is Lychrel until proven otherwise. In addition you are given that for every number below ten-thousand, it will either (i) become a palindrome in less than fifty iterations, or, (ii) no one, with all the computing power that exists, has managed so far to map it to a palindrome. In fact, 10677 is the first number to be shown to require over fifty iterations before producing a palindrome: 4668731596684224866951378664 (53 iterations, 28-digits).
//Surprisingly, there are palindromic numbers that are themselves Lychrel numbers; the first example is 4994.
//How many Lychrel numbers are there below ten-thousand?
func howManyLychrels(limit int) int {
	sum := 0

	for i := 1; i < limit; i++ {
		if isLychrel(i) {
			sum++
		}
	}

	return sum
}

func isLychrel(number int) bool {
	for i := 0; i < 50; i++ {
		number = number + reverseInt(number)

		if isPalindrome(number) {
			return false
		}
	}
	return true
}

func reverseInt(n int) int {
	newInt := 0
	for n > 0 {
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return newInt
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
	fmt.Printf("Lychrels in 10000 first ints: %d\n", howManyLychrels(10000))
}
