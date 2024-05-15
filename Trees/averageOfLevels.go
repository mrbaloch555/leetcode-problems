package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	queue := []*TreeNode{root}
	res := []float64{}
	for len(queue) > 0 {
		nextLevel := []*TreeNode{}
		sum := 0
		count := 0
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
			sum += current.Val
			i--
		}
		res = append(res, float64(sum)/float64(count))
		queue = append(queue, nextLevel...)
	}
	return res
}

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7}}}
	fmt.Println(averageOfLevels(root))
}
