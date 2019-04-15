package problem016

import (
	"fmt"
	"math/big"
)

func powerDigitSum(power int) int {
	two := new(big.Int).SetUint64(2)
	p := new(big.Int).SetUint64(uint64(power))
	num := new(big.Int).Exp(two, p, nil)

	res := num.String()
	sum := 0

	const runePosition = 48
	for i := 0; i < len(res); i++ {
		sum += int(res[i]) - runePosition // we need to shift rune to get real ints
	}

	return sum
}

func Result() {
	fmt.Printf("2^1000 digits sum:  %d\n", powerDigitSum(1000))
}
