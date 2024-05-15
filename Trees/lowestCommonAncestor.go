package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 5,
		Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}},
		Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}},
	}
	p, q := 5, 4
	fmt.Println(lowestCommonAncestor(root, p, q))
}

func lowestCommonAncestor(root *TreeNode, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
