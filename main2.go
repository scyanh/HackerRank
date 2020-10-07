package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//s:="3"

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")

	if s, _ := reader.ReadString('\n'); true {
		s = strings.Trim(s, " \n")
		n,_:=strconv.Atoi(s)
		for i:=0; i<n; i++{
			fmt.Println("Hello")
		}
	}
}