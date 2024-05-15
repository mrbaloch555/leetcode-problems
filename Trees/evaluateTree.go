package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}}}
	fmt.Println(evaluateTree(root))
}

func evaluateTree(root *TreeNode) bool {
	if root.Val < 2 {
		return root.Val == 1
	}
	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}
	return evaluateTree(root.Left) && evaluateTree(root.Right)
}
