package basics

import (
	"fmt"
	"strconv"
)

// string to int
func stringToInt() int {
	s := "10"
	i, _ := strconv.Atoi(s)
	i64, _ := strconv.ParseInt(s, 10, 64)
	return i + int(i64)
}

// int to string
func intToString() string {
	i := 10
	s := strconv.Itoa(i)

	i64 := int64(10)
	s2 := strconv.FormatInt(i64, 10)
	return s + s2
}

// create strings
func createString() string {
	i := 10
	s := fmt.Sprintf("number %d", i)
	return s
}

//arrays
func createArray() []int {
	arr := [2]int{}
	arr[0] = 1
	arr[1] = 2

	//sl := []int{}
	var sl []int
	for _, el := range arr {
		sl = append(sl, el)
	}

	return sl
}

//matrix
func createMatrix() [][]int {
	m := [][]int{
		{1, 2},
		{1, 2},
	}
	return m
}

// map
func createMap() map[int]int {
	m := make(map[int]int)
	m[1] = 1

	if val, ok := m[1]; ok {
		m[2] = val
	}

	return m
}

// go routine
func createGoRoutine() {
	messages := make(chan string)
	go func() {
		messages <- "hello"
	}()
	msg := <-messages
	fmt.Println(msg)
}
