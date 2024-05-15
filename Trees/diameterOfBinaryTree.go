package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := depth(root.Left)
	rightDepth := depth(root.Right)

	diameter := leftDepth + rightDepth

	leftDiameter := diameterOfBinaryTree(root.Left)
	rightDiamter := diameterOfBinaryTree(root.Right)

	return max(diameter, max(leftDiameter, rightDiamter))
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := depth(root.Left)
	rightDepth := depth(root.Right)
	return max(leftDepth, rightDepth) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(diameterOfBinaryTree(root))
}
