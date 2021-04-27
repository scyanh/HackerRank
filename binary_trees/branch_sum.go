package binary_trees

import "fmt"

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func NewBinaryTree(root int, values ...int) *BinaryTree {
	tree := &BinaryTree{Value: root}
	tree.Insert(values, 0)
	return tree
}
func (tree *BinaryTree) Insert(values []int, i int) *BinaryTree {
	if i >= len(values) {
		return tree
	}
	val := values[i]

	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		var current *BinaryTree
		current, queue = queue[0], queue[1:]
		if current.Left == nil {
			current.Left = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Left)

		if current.Right == nil {
			current.Right = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Right)
	}

	tree.Insert(values, i+1)
	return tree
}

func Start() {
	tree := NewBinaryTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	expected := []int{15, 16, 18, 10, 11}
	output := BranchSums(tree)
	fmt.Println("tree=", tree)
	fmt.Println("expected sum=", expected)
	fmt.Println("output sum=", output)
	fmt.Println("Equal=", Equal(expected, output))
}
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func calculateBranchSums(node *BinaryTree, runningSum int, sums *[]int) {
	if node == nil {
		return
	}

	runningSum += node.Value
	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, runningSum)
		return
	}

	calculateBranchSums(node.Left, runningSum, sums)
	calculateBranchSums(node.Right, runningSum, sums)
}

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	var sums []int
	calculateBranchSums(root, 0, &sums)
	return sums
}
