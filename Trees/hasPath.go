package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	target := 1
	fmt.Println(hasPathSum(root, target))
}

func hasPathSum(root *TreeNode, target int) bool {
	return helper(root, target)
}

func helper(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	target -= root.Val
	if root.Left == nil && root.Right == nil {
		return target == 0
	}
	return helper(root.Left, target) || helper(root.Right, target)
}
