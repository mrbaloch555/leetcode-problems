package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 7}}

	// fmt.Println(increasingBST(root))
	node := increasingBST(root)
	for node != nil {
		fmt.Println(node.Right)
		node = node.Right
	}
}

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	leftNode := increasingBST(root.Left)
	rightNode := increasingBST(root.Right)

	root.Right = rightNode
	root.Left = nil

	if leftNode == nil {
		return root
	}

	lastNode := leftNode
	for lastNode.Right != nil {
		lastNode = lastNode.Right
	}

	lastNode.Right = root

	return leftNode
}
