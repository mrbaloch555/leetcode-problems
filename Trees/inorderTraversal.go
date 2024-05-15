// **********************************************************************************
// Problem Number: 94
// Problem Name: Binary Tree Inorder Traversal
// Problem Link: https://leetcode.com/problems/binary-tree-inorder-traversal/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	fmt.Println(inorderTraversal(root))
}
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	leftNodes := inorderTraversal(root.Left)
	rightNodes := inorderTraversal(root.Right)
	return append(leftNodes, append([]int{root.Val}, rightNodes...)...)
}
