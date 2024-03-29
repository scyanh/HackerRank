package sorting

import (
	"fmt"
	"sort"
)

func SortPkg() {
	s := []int{4, 2, 3, 1}
	sort.Ints(s)
	fmt.Println(s)
	// reverse
	reverseThis(s)
}

func reverseThis(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}

// Sort array Bubble
func sortThis(array []int) []int {
	isOrdered := true

	for i, e := range array {
		if i > 0 && e < array[i-1] {
			isOrdered = false
			helperSwap(array, i)
			fmt.Println("arr=", array)
		}
	}
	if isOrdered {
		return array
	}
	return sortThis(array)
}

func helperSwap(array []int, idx int) {
	array[idx-1], array[idx] = array[idx], array[idx-1]
}

// Sort array Desc Bubble
func sortThisDesc(array []int) []int {
	isOrdered := true

	for i, e := range array {
		if i > 0 && e > array[i-1] {
			isOrdered = false
			helperSwap(array, i)
			fmt.Println("arr=", array)
		}
	}
	if isOrdered {
		return array
	}
	return sortThisDesc(array)
}

func StartThis() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(sortThisDesc(array))
}
