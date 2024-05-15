// **********************************************************************************
// Problem Number: 230
// Problem Name: Kth Smallest Element in a BST
// Problem Link: https://leetcode.com/problems/kth-smallest-element-in-a-bst/
// Status: Solved
// **********************************************************************************

package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	k := 2
	fmt.Println(kthSmallest(root, k))
}

func kthSmallest(root *TreeNode, k int) int {
	values := traverse(root)
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return values[k-1]
}

func traverse(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	leftNodes := traverse(root.Left)
	rightNodes := traverse(root.Right)
	return append(leftNodes, append([]int{root.Val}, rightNodes...)...)
}
