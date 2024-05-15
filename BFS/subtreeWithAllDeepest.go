package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Result struct {
	Node  *TreeNode
	Depth int
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	result := dfs(root, 0)
	return result.Node
}

func dfs(node *TreeNode, depth int) Result {
	if node == nil {
		return Result{nil, depth}
	}

	left := dfs(node.Left, depth+1)
	right := dfs(node.Right, depth+1)

	if left.Depth > right.Depth {
		return left
	} else if left.Depth < right.Depth {
		return right
	} else {
		return Result{node, left.Depth}
	}
}
func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{Val: 8},
		},
	}
	fmt.Println(subtreeWithAllDeepest(root))
}
