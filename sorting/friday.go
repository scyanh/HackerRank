package sorting

import "fmt"

/*
Teniendo el array de numeros [7,1,3,2,4,5,6] se nos pide intercambiar la posición
de los elementos unos con otros hasta lograr ordenarlos de menor a mayor.
Resolver:
Cantidad mínima de intercambios necesarios para ordenar el arreglo

i   arr                     intercambio (indices)
0   [7, 1, 3, 2, 4, 5, 6]   intercambio (0,3)
1   [2, 1, 3, 7, 4, 5, 6]   intercambio (0,1)
2   [1, 2, 3, 7, 4, 5, 6]   intercambio (3,4)
3   [1, 2, 3, 4, 7, 5, 6]   intercambio (4,5)
4   [1, 2, 3, 4, 5, 7, 6]   intercambio (5,6)
5   [1, 2, 3, 4, 5, 6, 7]
*/

func switchElements(idx1 int, idx2 int, arr []int) {
	arr[idx1], arr[idx2] = arr[idx2], arr[idx1]
}


func sortFriday(arr []int) []int {
	changed := false
	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] < arr[i-1] {
			changed = true
			switchElements(i, i-1, arr)
			fmt.Println("arr=", arr)
			moves++
			fmt.Println("moves=", moves)
		}
	}
	if changed {
		sortFriday(arr)
	}
	return arr
}

func Friday() {
	//arr := []int{7, 1, 3, 2, 4, 5, 6}

	arr := []int{7, 3, 2, 6, 5, 1, 20, 10}
	fmt.Println("arr=", arr)
	moves=0
	arr = sortFriday(arr)
	//arr=sortThis(arr)

	fmt.Println("final arr=", arr)
}
