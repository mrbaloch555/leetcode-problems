// **********************************************************************************
// Problem Number: 104
// Problem Name: Maximum Depth of Binary Tree
// Problem Link: https://leetcode.com/problems/maximum-depth-of-binary-tree/
// Status: Solved
// **********************************************************************************

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	counts := make(map[string]int, 0)
	paths := traverse(root, "", counts)
	max := 0
	for _, val := range paths {
		if val > max {
			max = val
		}
	}
	return max
}

func traverse(root *TreeNode, path string, counts map[string]int) map[string]int {
	if root == nil {
		return map[string]int{}
	}

	val := strconv.Itoa(root.Val)
	path += val + "->"
	if root.Left != nil {
		traverse(root.Left, path, counts)
	}
	if root.Right != nil {
		traverse(root.Right, path, counts)
	}
	counts[path] = len(strings.Split(path, "->")) - 1
	return counts
}
