package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {

	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	queue := []*TreeNode{root}
	lastLayer := []*TreeNode{}
	for len(queue) > 0 {
		nextLevel := []*TreeNode{}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			if curr.Left != nil {
				nextLevel = append(nextLevel, curr.Left)
			}
			if curr.Right != nil {
				nextLevel = append(nextLevel, curr.Right)
			}
			if len(nextLevel) > 0 {
				lastLayer = nextLevel
			}
		}
		queue = append(queue, nextLevel...)
	}
	return lastLayer[0].Val
}

func main() {
	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 5,
				Left: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 6}},
	}
	fmt.Println(findBottomLeftValue(root))
}
