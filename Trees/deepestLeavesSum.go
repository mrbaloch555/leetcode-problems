package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sum := 0
		nextLevel := []*TreeNode{}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			sum += current.Val
			if current.Left != nil {
				nextLevel = append(nextLevel, current.Left)
			}
			if current.Right != nil {
				nextLevel = append(nextLevel, current.Right)
			}
		}
		if len(nextLevel) == 0 {
			return sum
		}
		queue = append(queue, nextLevel...)
	}
	return 0
}
func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 7}},
		Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3,
		Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 8}}}}
	fmt.Println(deepestLeavesSum(root))
}
