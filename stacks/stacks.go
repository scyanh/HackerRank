package stacks

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() int {
	lastIdx := len(s.items) - 1
	item := s.items[lastIdx]
	s.items = s.items[:lastIdx]
	return item
}

func Start() {
	var stack Stack
	fmt.Println("stack=", stack)
	stack.push(1)
	fmt.Println("stack=", stack)
	stack.push(2)
	fmt.Println("stack=", stack)
	stack.push(3)
	fmt.Println("stack=", stack)

	for range stack.items{
		item := stack.pop()
		fmt.Println("item=", item)
	}

}
