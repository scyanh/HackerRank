package main

import (
	"fmt"
	"math"
	"strings"
)

//# Input : K = 3 # Output : 000, 001, 010, 100, 101
//  Input : K = 4 # Output :0000, 0001, 0010, 0100, 0101, 1000, 1001, 1010


func start() {
	Creator(3)
}

func Creator(input int) []string {
	var output []string

	for i := 0.0; i < math.Pow(2.0, float64(input)); i++ {
		var sb strings.Builder
		val := int(i)
		s := fmt.Sprintf("%b", val)
		// eliminate the consecutives ones
		if strings.Contains(s, "11") {
			continue
		}

		// append with 0's
		total := len(s)
		for j := total; j < input; j++ {
			sb.WriteString("0")
		}
		sb.WriteString(s)
		output = append(output, sb.String())
	}

	fmt.Println("output=", output)

	return output
}
