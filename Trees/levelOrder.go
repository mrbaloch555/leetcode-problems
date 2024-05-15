package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	res := [][]int{}

	for len(queue) > 0 {
		nextLevel := []*TreeNode{}
		nodeValues := []int{}

		for _, current := range queue {
			nodeValues = append(nodeValues, current.Val)

			if current.Left != nil {
				nextLevel = append(nextLevel, current.Left)
			}
			if current.Right != nil {
				nextLevel = append(nextLevel, current.Right)
			}
		}

		res = append(res, nodeValues)
		queue = nextLevel
	}

	return res
}

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7}}}
	fmt.Println(levelOrder(root))
}
