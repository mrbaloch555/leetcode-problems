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

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	res := []int{}
	for len(queue) > 0 {
		nextLevel := []*TreeNode{}
		count := 0
		max := math.Inf(-1)
		for i := 0; i < len(queue); i++ {
			count++
			current := queue[0]
			queue = queue[1:]
			if current.Left != nil {
				nextLevel = append(nextLevel, current.Left)
			}
			if current.Right != nil {
				nextLevel = append(nextLevel, current.Right)
			}
			if max < float64(current.Val) {
				max = float64(current.Val)
			}
			i--
		}
		res = append(res, int(max))
		queue = append(queue, nextLevel...)
	}
	return res
}

func main() {
	root := &TreeNode{Val: 0, Left: &TreeNode{Val: -1}}
	fmt.Println(largestValues(root))
}
