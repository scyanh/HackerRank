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

	// Inter
	//inter.Prob1()

	// Inter V
	//inter.V1()
	//inter.V2()
	array := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(SortAscending(array))
}

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
