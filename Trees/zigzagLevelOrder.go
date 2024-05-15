package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := [][]int{}
	iteration := true
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
		if !iteration {
			// Reverse the order of node values for odd levels
			reverse(nodeValues)
		}
		iteration = !iteration
		res = append(res, nodeValues)
		queue = nextLevel
	}
	return res
}

// Helper function to reverse a slice of integers in place
func reverse(slice []int) {
	left, right := 0, len(slice)-1
	for left < right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
}

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7}}}
	fmt.Println(zigzagLevelOrder(root))
}
