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

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min := math.Inf(1)
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.Left != nil {
			stack = append(stack, current.Left)
			diff := float64(math.Abs(float64(current.Val) - float64(current.Left.Val)))
			if diff < min {
				min = diff
			}

			if current.Right != nil {
				stack = append(stack, current.Right)
				diff := float64(math.Abs(float64(current.Val) - float64(current.Right.Val)))
				if diff < min {
					min = diff
				}
			}
		}
		return int(min)
	}
}

// func traverse(root *TreeNode) int {

// }
func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 3}}}
	fmt.Println(getMinimumDifference(root))
}
