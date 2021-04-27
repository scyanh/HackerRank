package dynamicProgramming

import (
	"fmt"
)

func ModifySequence(arr []int32) {
	for _, n := range arr {
		fmt.Println(n)
	}

	fmt.Println("--")

	changes := 0
	lower := int32(0)

	// get the lower number
	for i, n := range arr {
		if i == 0 {
			lower = n
		}
		if n < lower {
			lower = n
		}
	}

	fmt.Println("lower=", lower)

	for _, n := range arr {
		fmt.Println(n)
	}

	for i, _ := range arr {
		if i > 1 && arr[i] <= arr[i-1] {
			// is not asc
			changes = changes + 1
			arr[i] = arr[i-1] + 1
		}
	}

	fmt.Println("changes=", changes)

}

func GetNthFib(n int) int {
	memo := make(map[int]int)
	memo[1] = 0
	memo[2] = 1

	return helper(n, memo)
}

func helper(n int, memo map[int]int) int {
	if memo[n] > 0 || n == 1 {
		return memo[n]
	}
	/*if val,found:= memo[n]; found{
		return val
	}*/

	memo[n] = helper(n-1, memo) + helper(n-2, memo)
	return memo[n]

}

var fibMap map[int]int

func StartThis() {
	//ObjToString()
	fmt.Println(GetNthFib(6))
	fibMap = make(map[int]int)
}

func getFib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	if val, ok:=fibMap[n]; ok{
		return val
	}

	fibMap[n] = getFib(n-1) + getFib(n-2)
	return fibMap[n]
}
