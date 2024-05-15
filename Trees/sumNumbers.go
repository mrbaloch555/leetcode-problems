package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 0, Right: &TreeNode{Val: 1}}
	fmt.Println(sumNumbers(root))
}

func sumNumbers(root *TreeNode) int {
	return helper(root, "")
}

func helper(root *TreeNode, path string) int {
	if root == nil {
		return 0
	}
	path += strconv.Itoa(root.Val)
	if root.Right == nil && root.Left == nil {
		num, _ := strconv.ParseInt(path, 10, 32)
		return int(num)
	}

	return helper(root.Left, path) + helper(root.Right, path)
}
