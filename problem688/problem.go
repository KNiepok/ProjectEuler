package problem688

import (
	"math/big"
	"time"
)

var (
	one = new(big.Int).SetInt64(1)
	two = new(big.Int).SetInt64(2)
)

func solve() {
	start := time.Now()
	var columns = new(big.Int).Exp(new(big.Int).SetInt64(10), new(big.Int).SetInt64(16), nil)
	var rows = new(big.Int).Mul(new(big.Int).Sqrt(columns), two)
	modulo := new(big.Int).SetInt64(1000000007)

	rowNumber := new(big.Int)

	c := make(chan *big.Int, 100)
	sumc := make(chan *big.Int, 100)

	for i := 0; i < 8; i++ {
		go func() {
			sum := new(big.Int)
			for rowNumber := range c {
				rowsum := rowSum(columns, rowNumber)
				sum.Add(sum, rowsum)
			}
			sumc <- sum
		}()
	}

	for rowNumber.Cmp(rows) != 1 {
		c <- new(big.Int).Add(rowNumber, one)
		rowNumber.Add(rowNumber, one)
	}

	close(c)

	sum := new(big.Int)
	for i := 0; i < 8; i++ {
		sum.Add(sum, <-sumc)
	}

	println(sum.Mod(sum, modulo).String())
	println(time.Now().Sub(start).String())
}

// counts a sum of a row
func rowSum(n, k *big.Int) *big.Int {
	if k.Uint64() == 0 {
		return new(big.Int)
	}
	if k.Uint64() == 1 {
		return ar(one, n)
	}

	s := new(big.Int)
	i := ar(one, k)
	if i.Cmp(n) == 1 {
		return new(big.Int)
	}

	s.Sub(n, i.Sub(i, one))
	c := new(big.Int)
	r := new(big.Int)
	c, r = c.DivMod(s, k, r)
	res := new(big.Int)
	c2 := new(big.Int)
	res.Mul(k, ar(one, c)).Add(res, c2.Add(c, one).Mul(c2, r))

	return res
}

// sum of integers between k and n
// return (n - k + 1) * (n + k) / 2
func ar(k, n *big.Int) *big.Int {
	var a, b, result big.Int
	a.Sub(n, k).Add(&a, one)
	b.Add(n, k)
	result.Mul(&a, &b).Quo(&result, two)
	return &result
}
