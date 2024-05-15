// **********************************************************************************
// Problem Number: 101
// Problem Name: Symmetric Tree
// Problem Link: https://leetcode.com/problems/symmetric-tree/
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
	root := &TreeNode{Val: 1, Left: &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 4},
	},
		Right: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3}},
	}
	fmt.Println(isSymmetric(root))
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkTree(root.Left, root.Right)
}

func checkTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) || (p != nil && q != nil && p.Val != q.Val) {
		return false
	}
	if p != nil && q != nil {
		return checkTree(p.Left, q.Right) && checkTree(p.Right, q.Left)
	}
	return true
}
