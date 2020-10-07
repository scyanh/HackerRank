package recursion

import (
	"fmt"
	"strconv"
	"strings"
)

type ArrInterface []interface{}

/*func handlerSum(e interface{}, multi int) int {

	switch rt.Kind() {
	case reflect.Slice:
		fmt.Println("is a slice")
		handlerSum(e., multi+1)
	case reflect.Array:
		fmt.Println("is an array")
	case reflect.Int:
		fmt.Println("is an int")
		return e.(int)*multi
	default:
		fmt.Println("is something else")
	}
	return 0
}*/

func sumNumber(array ArrInterface, multi int) int {
	sum := 0
	for _, e := range array {
		if cast, ok := e.(ArrInterface); ok {
			sum += sumNumber(cast, multi+1)
		} else if number, ok := e.(int); ok {
			sum += number
		}
	}

	return sum * multi
}

func SuperDigit(s2 string) int {
	n := 0
	k := 3 // times
	s2 = "148"
	s := ""

	for i := 0; i < k; i++ {
		s += s2
	}

	array := strings.Split(s, "")
	fmt.Println(s)

	n = sumNumberArray(array)

	return n
}

func sumNumberArray(array []string) int {
	sum := 0
	for _, e := range array {
		n, _ := strconv.Atoi(e)
		sum += n
	}

	if sum < 10 {
		return sum
	}

	stringSum := strconv.Itoa(sum)
	return sumNumberArray(strings.Split(stringSum, ""))
}

func Permutations(array []int) [][]int {
	result := [][]int{}

	if len(array) == 0 || len(array) == 1 {
		return result
	}

	/*for i, n := range array {

	}*/

	return result
}

func StartThis() {
	/*array := ArrInterface{
		5, 2,
		ArrInterface{7, -1},
		3,
		ArrInterface{
			6,
			ArrInterface{
				-13, 8,
			},
			4,
		},
	}
	fmt.Println(sumNumber(array, 1))*/
	//fmt.Println(Permutations([]int{1, 2, 3}))
	//fmt.Println(SuperDigit(""))


}
