package search

import (
	"fmt"
	"math"
)

// Search number
func binariSearch() int {
	array := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	target := 33

	for i, e := range array {
		if e == target {
			return i
		}
	}

	return -1
}

// order array
func greatestSearch() int {
	array := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	ordered := []int{math.MinInt32, math.MinInt32, math.MinInt32}

	for _, e := range array {
		if e > ordered[2] {
			shiftArray(ordered, e, 2)
		} else if e > ordered[1] {
			shiftArray(ordered, e, 1)
		} else if e > ordered[0] {
			shiftArray(ordered, e, 0)
		}
	}
	fmt.Println("ordered=", ordered)

	return 0
}

func shiftArray(array []int, number int, idx int) {
	if idx == 2 {
		array[0] = array[1]
		array[1] = array[2]
	} else if idx == 1 {
		array[0] = array[1]
	}

	array[idx] = number
}

func StartThis() {
	fmt.Println(greatestSearch())
}
