package mytest

import (
	"fmt"
)

func Max(a,b int) int{
	if a>b{
		return a
	}else{
		return b
	}
}
func largestRectangleArea(heights []int)int{
	var max int
	var stack = make([]int,0)
	var nums= make([]int,0)

	nums = append(nums, []int{0}...)
	nums = append(nums, heights...)

	for i:=0; i<len(nums); i++{
		for len(stack)>0 && nums[stack[len(stack)-1]]>nums[i]{
			tmp :=stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			max = Max(max, nums[tmp]*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack,i)
	}
	return max
}

// Palindrome
func FizzBuzz() int{
	samples := [][]int{
		{1,1,1,1,1},
		{1,1,1,0,0},
		{1,1,1,0,0},
		{0,0,0,0,0},
		{0,0,1,1,0},
	}
	var max int
	rows:=len(samples)
	if rows==0{
		return max
	}

	cols:=len(samples[0])
	nums:=make([]int,cols)

	for row:=0; row<rows; row++{
		for col:=0; col < cols; col++{
			if samples[row][col]!=0{
				nums[col]=nums[col]+1
			}else{
				nums[col]=0
			}
		}
		maxArea:=largestRectangleArea(nums)
		max=Max(max,maxArea)
	}
	return max
}

func StartThis() {
	fmt.Println(FizzBuzz())
}
