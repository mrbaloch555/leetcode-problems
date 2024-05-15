package main

import (
	"fmt"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	var explor func(*TreeNode, *TreeNode, int)
	explor = func(left, right *TreeNode, i int) {
		if left == nil {
			return
		}
		if i%2 == 1 {
			val := left.Val
			left.Val = right.Val
			right.Val = val
		}
		explor(left.Left, right.Right, i+1)
		explor(left.Right, right.Left, i+1)
	}
	explor(root.Left, root.Right, 1)
	return root
}

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 8,
			},
			Right: &TreeNode{
				Val: 13,
			},
		}, Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 21,
			},
			Right: &TreeNode{
				Val: 34,
			},
		},
	}

	fmt.Println(reverseOddLevels(root))
	fmt.Println(root.Left, root.Right)
}
