package sorting

import (
	"fmt"
)

/*
Having the array of numbers [7,1,3,2,4,5,6] asking us to sort in ascending mode:
Nice to have, is have 5 changes in the array

i   arr                     change (indices)
0   [7, 1, 3, 2, 4, 5, 6]   change (0,3)
1   [2, 1, 3, 7, 4, 5, 6]   change (0,1)
2   [1, 2, 3, 7, 4, 5, 6]   change (3,4)
3   [1, 2, 3, 4, 7, 5, 6]   change (4,5)
4   [1, 2, 3, 4, 5, 7, 6]   change (5,6)
5   [1, 2, 3, 4, 5, 6, 7]

but I found a solution to only have 2 changes in the array

0   [1, 3, 2, 4, 5, 6, 7]
1   [1, 2, 3, 4, 5, 6, 7]

*/

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func moveElement(idx int, arr []int) []int {
	fmt.Println("n=",arr[idx])
	distance := arr[idx+1] - arr[idx]
	if distance < 0 {
		distance *= -1
	}
	distance=Min(distance,len(arr)-idx-1)
	var tmpArr []int

	tmpArr = append(tmpArr, arr[:idx]...)
	tmpArr = append(tmpArr, arr[idx+1:(idx+1+distance)]...)
	tmpArr = append(tmpArr, arr[idx])
	tmpArr = append(tmpArr, arr[(idx+distance+1):]...)

	return tmpArr
}

var moves int

func sortFridayPlus(arr []int) []int {
	changed := false
	for i := 0; i < len(arr); i++ {
		if i < len(arr)-1 && arr[i] > arr[i+1] {
			changed = true
			arr = moveElement(i, arr)
			fmt.Println("arr=", arr)
			moves++
			break
		}
	}
	if changed {
		arr=sortFridayPlus(arr)
	}
	return arr
}

func FridayPlus() {
	//arr := []int{7, 1, 3, 2, 4, 5, 6}
	moves=0
	//arr := []int{7, 3, 2, 6, 5, 1, 20, 10}
	arr := []int{4, 3, 1, 2}
	fmt.Println("arr=", arr)

	arr2 := sortFridayPlus(arr)

	fmt.Println("final arr=", arr2)
	fmt.Println("moves=", moves)
}
