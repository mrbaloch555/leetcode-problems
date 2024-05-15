package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	counter := 0
	dfs(root, &counter)
	return counter
}

func dfs(root *TreeNode, counter *int) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftSum, leftCount := dfs(root.Left, counter)
	rightSum, rightCount := dfs(root.Right, counter)
	sum := leftSum + rightSum + root.Val
	count := leftCount + rightCount + 1

	avg := sum / count

	if avg == root.Val {
		*counter++
	}

	return sum, count
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   8,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(averageOfSubtree(root))
}
