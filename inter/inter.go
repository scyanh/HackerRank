package inter

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)



func StartThis() {
	//var numbers []int32
	numbers :=[]int32{6,2,4,10}
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })

	minDiff:=numbers[0]-numbers[1]
	if minDiff < 0 {
		minDiff=minDiff*-1
	}
	fmt.Println("minDiff=",minDiff)
	for i := range numbers{
		//fmt.Println(numbers[i])
		if i<len(numbers)-1 {
			diff:=numbers[i+1]-numbers[i]
			if diff==minDiff {
				fmt.Println(numbers[i],numbers[i+1])
			}
		}
	}
}

func Prob1(){
	fmt.Println("in Prob1")
	// Declaring some variables
	var inpu1 int
	var input2 int

	// Calling Scanf() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Scanf("%d", &inpu1)
	fmt.Scanf("%d", &input2)

	fmt.Printf("The 1= %d 2= %d ", inpu1, input2)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")

	text:=""
	if t, _ := reader.ReadString('\n'); true {
		text = strings.Trim(t, " \n")
	}
	fmt.Println("text=",text)

	split := strings.Split(text, " ")

	fmt.Println("split",split)

}