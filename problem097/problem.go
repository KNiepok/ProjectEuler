package problem097

import (
	"fmt"
	"math/big"
)

//The first known prime found to exceed one million digits was discovered in 1999, and is a Mersenne prime of the form 26972593−1; it contains exactly 2,098,960 digits. Subsequently other Mersenne primes, of the form 2p−1, have been found which contain more digits.
//
//However, in 2004 there was found a massive non-Mersenne prime which contains 2,357,207 digits: 28433×2^7830457+1.
//Find the last ten digits of this prime number.
func lastDigits() {
	// set big ints
	two, _ := new(big.Int).SetString("2", 10)
	power, _ := new(big.Int).SetString("7830457", 10)
	multi, _ := new(big.Int).SetString("28433", 10)
	one, _ := new(big.Int).SetString("1", 10)
	m, _ := new(big.Int).SetString("10", 10)
	m.Exp(m, m, nil)

	// perform actual calculation
	two.Exp(two, power, nil).Mul(two, multi).Add(two, one).Mod(two, m)

	fmt.Println(two)
}
