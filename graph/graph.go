package graph

import (
	"fmt"
)

type Node struct {
	Name     string
	Children []*Node
}

func (node *Node) DepthFirstSearch(array []string) []string {
	array = append(array, node.Name)
	for _, e := range node.Children {
		array = e.DepthFirstSearch(array)
	}

	return array
}

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: []*Node{},
	}
}
func (n *Node) AddChildren(names ...string) *Node {
	for _, name := range names {
		child := Node{Name: name}
		n.Children = append(n.Children, &child)
	}
	return n
}
func StartThis() {
	var graph = NewNode("A").AddChildren("B", "C", "D")
	graph.Children[0].AddChildren("E").AddChildren("F")
	graph.Children[2].AddChildren("G").AddChildren("H")
	graph.Children[0].Children[1].AddChildren("I").AddChildren("J")
	graph.Children[2].Children[0].AddChildren("K")
	fmt.Println(graph.DepthFirstSearch([]string{}))

}
