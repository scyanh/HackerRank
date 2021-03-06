package arrays

import (
	"fmt"
	"math"
	"sort"
)

func rotLeftNotOptimized(a []int, d int) []int {
	for i := 0; i < d; i++ {
		finalNumber := a[0]
		for i, _ := range a {
			if i == len(a)-1 {
				a[i] = finalNumber
			} else {
				a[i] = a[i+1]
			}
		}
	}

	return a
}
func rotLeftWithDivision(a []int, d int) []int {
	if d == len(a) || d == 0 {
		return a
	}
	dd := d % len(a)

	fmt.Println("array len=", len(a))
	fmt.Println("d=", d)
	fmt.Println("dd=", dd)

	for i := 0; i < dd; i++ {
		finalNumber := a[0]
		for i := range a {
			if i == len(a)-1 {
				a[i] = finalNumber
			} else {
				a[i] = a[i+1]
			}
		}
	}

	return a
}
func rotLeft(a []int, d int) []int {
	if d == len(a) || d == 0 {
		return a
	}
	dd := d % len(a)

	b1 := a[:dd]
	b2 := a[dd:]

	// sum slices
	b0 := append(b2, b1...)

	fmt.Println("b1=", b1)
	fmt.Println("b2=", b2)
	fmt.Println("b0=", b0)

	return b0
}

func RotLeftNotOptimizedService(times int) {
	myArr := []int{1, 2, 3, 4, 5}
	fmt.Println(rotLeftNotOptimized(myArr, times))
}
func RotLeftServiceWithDivision(times int) {
	myArr := []int{1, 2, 3, 4, 5}
	fmt.Println(rotLeftWithDivision(myArr, times))
}
func RotLeftService(times int) {
	myArr := []int{1, 2, 3, 4, 5}
	fmt.Println(rotLeft(myArr, times))
}

// Swap elements to order Asc
func SwapsToOrderAsc(array []int) []int {
	fmt.Println("array=", array)

	ordered := true
	for i, n := range array {
		if i < len(array)-1 && n > array[i+1] {
			ordered = false
			// do Swap
			array[i], array[i+1] = array[i+1], array[i]
		}
	}
	if ordered {
		return array
	}

	return SwapsToOrderAsc(array)
}

// Find sum in array
func SumInArray(array []int, target int) []int {
	array = []int{3, 5, -4, 8, 11, 1, -1, 6}
	target = 10

	if array == nil || len(array) == 0 {
		return []int{}
	}

	for i, n := range array {
		for j := i + 1; j <= len(array)-1; j++ {
			if n+array[j] == target {
				return []int{n, array[j]}
			}
		}
	}

	return []int{}
}

// Find sum in array with Map
func SumInArray2(array []int, target int) []int {
	array = []int{3, 5, -4, 8, 11, 1, -1, 6}
	target = 10

	if array == nil || len(array) == 0 {
		return []int{}
	}

	nums := make(map[int]bool)

	for _, n := range array {
		potential := target - n
		if nums[potential] {
			return []int{potential, n}
		}
		nums[n] = true
	}

	return []int{}
}

// Check if is sequence
func Sequence() bool {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	sequence := []int{5, -4}

	idSeq := 0

	for _, n := range array {
		if idSeq == len(sequence) {
			break
		}
		if n == sequence[idSeq] {
			idSeq++
		}
	}

	fmt.Println("idSeq=", idSeq)

	return idSeq == len(sequence)
}

// Sum target
func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	tripplets := [][]int{}

	for i, _ := range array {
		sum := 0
		if i < (len(array) - 2) {
			for j, _ := range array {
				if j > i {
					for k, _ := range array {
						if k > j {
							sum = array[i] + array[j] + array[k]
							if sum == target {
								tripplets = append(tripplets, []int{array[i], array[j], array[k]})
							}
						}
					}
				}
			}
		}
	}

	return tripplets
}

// Smallest Difference
func SmallestDifference(array1 []int, array2 []int) []int {
	array := []int{}
	smallestNumber := 0

	for i, e := range array1 {
		for j, e2 := range array2 {
			diff := e - e2
			abs := int(math.Sqrt(float64(diff * diff)))
			if i == 0 && j == 0 {
				smallestNumber = abs
				array = append(array, array1[i])
				array = append(array, array2[j])
			} else if abs < smallestNumber {
				smallestNumber = abs
				array[0] = array1[i]
				array[1] = array2[j]
			}
		}
	}

	return array
}

// Move n to the end
func MoveElementToEnd(array []int, toMove int) []int {
	sort.Ints(array)
	finalArray := []int{}
	toMoveArray := []int{}
	for _, e := range array {
		if e == toMove {
			toMoveArray = append(toMoveArray, e)
		} else {
			finalArray = append(finalArray, e)
		}
	}
	for _, e := range toMoveArray {
		finalArray = append(finalArray, e)
	}

	return finalArray
}

// Is Monotonic
func IsMonotonic(array []int) bool {
	if len(array) < 2 {
		return true
	}

	isAsc := true
	arrayOrdered := make([]int, len(array))
	copy(arrayOrdered, array)

	sort.Ints(arrayOrdered)
	if array[0] != arrayOrdered[0] {
		isAsc = false
	}

	for i, e := range array {
		if isAsc {
			if e != arrayOrdered[i] {
				return false
			}
		} else {
			if e != arrayOrdered[len(array)-1-i] {
				return false
			}
		}
	}

	fmt.Println("array=", array)
	fmt.Println("arrayOrdered=", arrayOrdered)

	return true
}

// Spiral Matrix
func SpiralTraverse() []int {
	matrix := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}

	result := []int{}

	//1 - toRight, 2 - toDown, 3 - toLeft, 4 - toUp
	//spiral:=[]int{1,2,3,4}
	runTo := 1
	idx := 0
	skipLeft := 0
	skipRight := 0
	skipTop := 0
	skipDown := 0

	lenWidth := len(matrix[0])
	lenHeigh := len(matrix)

	for {
		if runTo == 1 { // toRight
			for i, e := range matrix[skipTop] {
				if i >= skipLeft && i <= lenWidth-1-skipRight {
					result = append(result, e)
					idx += 1
					if idx == lenWidth*lenHeigh {
						return result
					}
				}
			}
			runTo += 1
			skipTop += 1
		} else if runTo == 2 { // toDown
			for i, e := range matrix {
				if i > skipDown && i <= lenHeigh-skipTop {
					result = append(result, e[len(matrix[i])-1-skipRight])
					idx += 1
					if idx == lenWidth*lenHeigh {
						return result
					}
				}
			}
			runTo += 1
			skipRight += 1
		} else if runTo == 3 { // toLeft
			for i, _ := range matrix[lenHeigh-1-skipDown] {
				if i >= skipRight && i <= lenWidth-skipLeft {
					result = append(result, matrix[lenHeigh-1-skipDown][lenWidth-1-i])
					idx += 1
					if idx == lenWidth*lenHeigh {
						return result
					}
				}
			}
			runTo += 1
			skipDown += 1
		} else if runTo == 4 { // toTop
			for i, _ := range matrix {
				if i >= skipDown && i < lenHeigh-skipTop {
					result = append(result, matrix[lenHeigh-1-i][skipLeft])
					idx += 1
					if idx == lenWidth*lenHeigh {
						return result
					}
				}
			}
			runTo = 1
			skipLeft += 1
		}
	}

	return result
}

func reverseIntArray(array []int) []int {
	reversed := []int{}
	for i := range array {
		n := array[len(array)-1-i]
		//fmt.Println(n) -- sanity check
		reversed = append(reversed, n)
	}
	return reversed
}

func SpiralTraverse2() []int {
	/*matrix := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}*/
	/*matrix := [][]int{
		{1, 2, 3, 4, 5, 6},
		{20, 21, 22, 23, 24, 7},
		{19, 32, 33, 34, 25, 8},
		{18, 31, 36, 35, 26, 9},
		{17, 30, 29, 28, 27, 10},
		{16, 15, 14, 13, 12, 11},
	}*/
	matrix := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{20, 19, 18, 17, 16, 15, 14, 13, 12, 11},
	}

	totalElements := len(matrix[0]) * len(matrix)

	fmt.Println("total elements=", totalElements)

	result := []int{}

	if totalElements == 1 {
		result = append(result, matrix[0]...)
		return result
	}

	//1 - toRight, 2 - toDown, 3 - toLeft, 4 - toUp
	runTo := 1

	for i := 1; i < totalElements; i++ {
		switch runTo {
		case 1:
			fmt.Println("copy first row and pop it")
			result = append(result, matrix[0]...)
			matrix = matrix[1:]
		case 2:
			fmt.Println("copy and pop last element of each row")
			for j, e := range matrix {
				result = append(result, e[len(e)-1])
				matrix[j] = matrix[j][:len(matrix[j])-1]
			}
		case 3:
			fmt.Println("inverse last row and copy")
			result = append(result, reverseIntArray(matrix[len(matrix)-1])...)
			matrix = matrix[:len(matrix)-1]
		case 4:
			fmt.Println("copy and pop first element of each row")
			for j := range matrix {
				result = append(result, matrix[len(matrix)-j-1][0])
				matrix[len(matrix)-j-1] = matrix[len(matrix)-j-1][1:]
			}
		}

		fmt.Println("len(result)=", len(result))
		if len(result) >= totalElements {
			return result
		}

		runTo++
		if runTo > 4 {
			runTo = 1
		}
	}

	return result
}

func selectAWinner() {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}

	memo := make(map[string]int)

	for i, result := range results {
		if result == 1 {
			fmt.Println(competitions[i][0])
			memo[competitions[i][0]] += 3
		} else {
			fmt.Println(competitions[i][1])
			memo[competitions[i][1]] += 3
		}
	}

	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range memo {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	fmt.Println("winner=", ss[0].Key)
	fmt.Println("winner=", ss[0].Value)
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func LongestPeak() {
	//arr := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	var arr []int

	i := 1
	maxLong := 0
	for i < len(arr)-1 {
		isPeak := false
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			isPeak = true
		}

		if isPeak {
			fmt.Println("peak=", i)
			fmt.Println("peak=", arr[i])

			currentLong := 3

			idxLeft := i - 2
			for idxLeft >= 0 && arr[idxLeft] < arr[idxLeft+1] {
				currentLong++
				idxLeft--
			}
			idxRight := i + 2
			for idxRight <= len(arr)-1 && arr[idxRight] < arr[idxRight-1] {
				currentLong++
				idxRight++
			}

			maxLong = MaxInt(maxLong, currentLong)

		}

		i += 1
	}
	fmt.Print("maxLong=", maxLong)
}

func ArrOfProducts() {
	//arr := []int{5, 1, 4, 2}
	arr := []int{9, 3, 2, 1, 9, 5, 3, 2}

	var sol []int
	for i := range arr {
		multiplies := 1
		for j := range arr {
			if i != j {
				multiplies = multiplies * arr[j]
			}
		}

		sol = append(sol, multiplies)
	}
	fmt.Print("sol=", sol)
}

func FirstDuplicated() {
	arr := []int{2, 1, 5, 3, 3, 2, 4}
	duplicated := -1

	memo := make(map[int]bool)
	for _, el := range arr {
		if memo[el] {
			duplicated = el
			break
		}
		memo[el] = true
	}

	fmt.Print("duplicated=", duplicated)
}

func FourNumberSum() {
	arr := []int{7, 6, 4, -1, 1, 2}
	targetSum := 100

	var quadruplets [][]int
	var finalQuadruplets [][]int

	for i1 := 0; i1 <= len(arr)-4; i1++ {
		for i2 := i1 + 1; i2 <= len(arr)-3; i2++ {
			for i3 := i2 + 1; i3 <= len(arr)-2; i3++ {
				for i4 := i3 + 1; i4 <= len(arr)-1; i4++ {
					quadruplet := []int{arr[i1], arr[i2], arr[i3], arr[i4]}
					quadruplets = append(quadruplets, quadruplet)
				}
			}
		}
	}

	// select sum that is equal to target
	for i :=range quadruplets{
		sum:=0
		for j := range quadruplets[i]{
			sum+=quadruplets[i][j]
		}
		if sum==targetSum{
			finalQuadruplets=append(finalQuadruplets, quadruplets[i])
		}
	}

	fmt.Println(quadruplets)
	fmt.Println(finalQuadruplets)
}

func SubArrSort(){
	arr:=[]int{1,2,4,7,10,11,7,12,6,7,16,18,19}
	arr2:=make([]int,len(arr))
	copy(arr2,arr)
	sort.Ints(arr2)
	fmt.Println(arr)
	fmt.Println(arr2)


	idxMin:=0
	idxMax:=0
	sol:=[]int{idxMin,idxMax}

	for i :=range arr{
		if arr[i]!=arr2[i]{
			idxMin=i
			break
		}
	}
	for i:=range arr{
		if arr[len(arr)-i-1]!=arr2[len(arr)-i-1]{
			idxMax=len(arr)-i-1
			break
		}
	}
	sol[0]=idxMin
	sol[1]=idxMax
	fmt.Println(idxMin)
	fmt.Println(idxMax)
	fmt.Println(sol)
}




func StartThis() {
	//fmt.Println(ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0))
	//fmt.Println(SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17}))
	//fmt.Println(MoveElementToEnd([]int{2, 1, 2, 2, 2, 3, 4, 2}, 2))
	//fmt.Println(IsMonotonic([]int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}))

	//fmt.Println(SwapsToOrderAsc([]int{1, 3, 5, 2, 4, 6, 7}))
	//fmt.Println(SpiralTraverse())
	//selectAWinner()

	//LongestPeak()
	//ArrOfProducts()
	//FirstDuplicated()
	//FourNumberSum()
	SubArrSort()
}
