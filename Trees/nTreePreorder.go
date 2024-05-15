package main

import (
	"fmt"
)

type Node struct {
	Val      int
	Children []*Node
}

func main() {

	root := &Node{Val: 1, Children: []*Node{&Node{Val: 3, Children: []*Node{&Node{Val: 5}, &Node{Val: 6}}}, &Node{Val: 2}, &Node{Val: 4}}}
	fmt.Println(preorder(root))
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var result []int
	for _, child := range root.Children {
		result = append(result, preorder(child)...)
	}
	result = append([]int{root.Val}, result...)

	return result
}
