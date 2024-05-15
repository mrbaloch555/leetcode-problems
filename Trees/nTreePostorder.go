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
	fmt.Println(postorder(root))
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	stack := []*Node{root}
	res := []int{}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{current.Val}, res...)
		for _, child := range current.Children {
			stack = append(stack, child)
		}
	}

	return res
}
