package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	var explore func(*TreeNode, *TreeNode, *TreeNode) int
	explore = func(root, parent, grandparent *TreeNode) int {
		if root == nil {
			return 0
		}

		sum := 0
		if grandparent != nil && grandparent.Val%2 == 0 {
			sum += root.Val
		}

		sum += explore(root.Left, root, parent) + explore(root.Right, root, parent)

		return sum
	}
	return explore(root, nil, nil)
}

func main() {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 9},
			},
			Right: &TreeNode{
				Val:   7,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 5},
			},
		},
	}
	fmt.Println(sumEvenGrandparent(root))
}
