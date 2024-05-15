package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		count++
	}
	return count
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2,
		Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5},
	}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}

	fmt.Println(countNodes(root))
}
