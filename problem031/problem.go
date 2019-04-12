package problem031

import "fmt"

//In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:
//
//1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
//It is possible to make £2 in the following way:
//
//1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
//How many different ways can £2 be made using any number of coins?
func countWays() {
	coins := []int{
		1, 2, 5, 10, 20, 50, 100, 200,
	}

	result := recursivelyCountWays(200, coins, len(coins))
	fmt.Printf("Number of ways %d\n", result)
}

func recursivelyCountWays(target int, set []int, size int) int {
	if target < 0 {
		return 0
	}

	if target == 0 {
		return 1
	}

	if size == 1 {
		return size
	}

	return recursivelyCountWays(target, set, size-1) + recursivelyCountWays(target-set[size-1], set, size)
}
