package problem015

import (
	"fmt"
)

//Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.
//How many such routes are there through a 20×20 grid?

// introduce an int tuple
type tuple struct {
	x, y int
}

type tupleMap map[tuple]int

// since these tuples will have association so l(x,y) == l(y,x) we will abstract comparision to a method on map
func (tm tupleMap) getTupleLength(t tuple) (result int, exists bool) {
	if val, exists := tm[tuple{t.x, t.y}]; exists {
		return val, exists
	}
	if val, exists := tm[tuple{t.y, t.x}]; exists {
		return val, exists
	}

	return 0, false
}

var cached tupleMap

func initMap() {
	cached = make(map[tuple]int)
	cached[tuple{1, 1}] = 2
}

func letticePaths(x, y int) int {
	// return if already evaluated
	if val, ok := cached.getTupleLength(tuple{x, y}); ok {
		return val
	}

	//base case: if we have grid size (x,1) or (y,1), we have two ways (right->down or down->right)
	if x == 0 || y == 0 {
		return 1
	}

	// calculate tuple path length from Pascals triangle
	val := letticePaths(x-1, y) + letticePaths(x, y-1)
	cached[tuple{y, x}] = val
	return val
}

func Result() {
	initMap()
	fmt.Printf("Number of lettice paths:  %d\n", letticePaths(20, 20))
}
