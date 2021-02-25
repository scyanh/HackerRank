package stringToInt

import (
	"fmt"
	"strconv"
)

// Convertion
func Convertion() {
	var myInt int
	var myInt64 int64

	myInt = 1
	myInt64 = 2

	// int to string
	var myString string
	myString = strconv.Itoa(myInt)
	myString = strconv.FormatInt(myInt64, 10)
	fmt.Println("myString int64=", myString)

	// string to int
	myInt, _ = strconv.Atoi(myString)
	myInt64, _ = strconv.ParseInt(myString, 10, 64)
}