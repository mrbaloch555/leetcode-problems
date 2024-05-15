package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 5}}
	fmt.Println(rightSideView(root))
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Right == nil && root.Left != nil {
		return rightSideView(root.Left)
	}
	return append([]int{root.Val}, rightSideView(root.Right)...)
}
