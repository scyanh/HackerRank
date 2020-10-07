package sorting

import (
	"fmt"
)

// Sort array Bubble
func sortThis(array []int) []int {
	isOrdered:=true

	for i,e:=range array{
		if i>0 && e<array[i-1]{
			isOrdered=false
			helperSwap(array, i)
		}
	}
	if isOrdered{
		return array
	}
	return sortThis(array)
}

func helperSwap(array []int, idx int){
	array[idx-1], array[idx]=array[idx], array[idx-1]
}


// Sort array Desc Bubble
func sortThisDesc(array []int) []int {
	isOrdered:=true

	for i,e:=range array{
		if i>0 && e>array[i-1]{
			isOrdered=false
			helperSwap(array, i)
		}
	}
	if isOrdered{
		return array
	}
	return sortThisDesc(array)
}

func StartThis() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(sortThisDesc(array))
}
