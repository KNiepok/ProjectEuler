package problem042

import "math"

const runeOffset = 64

func isWordTriangular(word string) bool {
	sum := 0
	for _, char := range word {
		n := rune(char) - runeOffset
		sum += int(n)
	}

	return isTriangleNumber(sum)
}

// t = ½n(n+1) => n = (sqrt(1+8t) - 1)/2
func isTriangleNumber(t int) bool {
	n := (math.Sqrt(1.0+float64(8*t)) - 1) / 2

	return n == float64(int(n))
}

func calculateWords(words []string) int {
	s := 0
	for _, w := range words {
		if isWordTriangular(w) {
			s++
		}
	}

	return s
}

