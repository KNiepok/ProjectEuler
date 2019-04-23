package problem085

import (
	"fmt"
	"math"
)

//https://projecteuler.net/problem=85

func gridArea() int {
	limit := 2 * 1000 * 1000
	x := 2000
	y := 1

	d := float64(math.MaxInt16)

	area := 0
	for x >= y {
		rects := x * (x + 1) * y * (y + 1) / 4

		if d > math.Abs(float64(rects-limit)) {
			area = x * y
			d = math.Abs(float64(rects - limit))
		}

		if rects > limit {
			x--
		} else {
			y++
		}
	}

	return area
}

func Result() {
	fmt.Printf("Minimum grid area is: %d\n", gridArea())
}
