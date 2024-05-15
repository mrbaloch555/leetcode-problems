// **********************************************************************************
// Problem Number: 100
// Problem Name: Same Tree
// Problem Link: https://leetcode.com/problems/same-tree/
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
	a := TreeNode{}

	// a.Left = &b
	// a.Right = &c
	fmt.Println(isSameTree(&a, &a))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
