package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: -1}}
	fmt.Println(maxPathSum(root))
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32 // Initialize maxSum with the smallest possible integer value
	calculateMaxPathSum(root, &maxSum)
	return maxSum
}

func calculateMaxPathSum(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	leftSum := calculateMaxPathSum(node.Left, maxSum)
	rightSum := calculateMaxPathSum(node.Right, maxSum)

	maxPath := max(max(node.Val, node.Val+leftSum+rightSum), max(node.Val+leftSum, node.Val+rightSum))
	*maxSum = max(*maxSum, maxPath)

	return max(node.Val, max(node.Val+leftSum, node.Val+rightSum))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
