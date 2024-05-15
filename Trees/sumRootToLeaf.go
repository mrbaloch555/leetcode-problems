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

	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}}}
	fmt.Println(sumRootToLeaf(root))
}

func sumRootToLeaf(root *TreeNode) int {
	return int(traverse(root, "", 0))
}

func traverse(root *TreeNode, path string, res int64) int64 {
	if root == nil {
		return res
	}
	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		dec, _ := strconv.ParseInt(path, 2, 64)
		res += dec
		return res
	}
	if root.Left != nil {
		res = traverse(root.Left, path, res)
	}
	if root.Right != nil {
		res = traverse(root.Right, path, res)
	}
	return res
}
