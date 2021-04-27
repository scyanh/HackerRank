package main

import "fmt"

func main() {

	// Modify Sequence *******
	//arr:=[]int32{5,4,3,4,5}
	//dynamicProgramming.StartThis()

	// RotLeft	*******
	/*times:=4
	arrays.RotLeftNotOptimizedService(times)
	arrays.RotLeftServiceWithDivision(times)
	arrays.RotLeftService(times)*/

	// Swaps Elemens *****
	//arrays.SwapsToOrderAsc()

	// Sum in array
	//arrays.StartThis()

	// Create interfaces
	//interfaces.StartThis()

	// Recursion with interface custom struct
	//recursion.StartThis()

	// Dynamic
	//dynamicProgramming.StartThis()

	// Strings
	//strings.StartThis()

	// Search
	//sorting.StartThis()

	// Graph
	//graph.StartThis()

	// Arrays
	//arrays.StartThis()
	//array := []int{8, 5, 2, 9, 5, 6, 3}
	//fmt.Println(SortAscending(array))

	// Inter
	//inter.Prob1()

	// Inter V
	//inter.V1()
	//inter.V2()

	// Binary tree
	//binary_trees.Start()

	// Stacks
	//stacks.Start()

	arr := [][]int{
		{10, 20},
		{30, 200},
		{400, 50},
		{30, 20},
	}
	distance := minDistance(arr)
	fmt.Println("min distance=",distance)
}

func getMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minDistance(arr [][]int) int {
	sum := 0
	for _, el := range arr {
		sum += getMin(el[0], el[1])
	}
	return sum
}

/*
func helperSwapAsc(array []int, idx int){
	array[idx-1], array[idx]=array[idx], array[idx-1]
}

func SortAscending(items []int) []int{
	isOrdered:=true

	for i,e:=range items{
		if i>0 && e<items[i-1]{
			isOrdered=false
			helperSwapAsc(items, i)
		}
	}
	if isOrdered{
		return items
	}
	return SortAscending(items)
}

*/
