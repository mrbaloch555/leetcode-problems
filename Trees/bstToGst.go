package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	var explore func(*TreeNode, int) int
	explore = func(node *TreeNode, sum int) int {
		if node == nil {
			return sum
		}
		sum = explore(node.Right, sum)
		sum += node.Val
		node.Val = sum
		sum = explore(node.Left, sum)
		return sum
	}

	explore(root, 0)
	return root
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
		},
		Right: &TreeNode{
			Val:  6,
			Left: &TreeNode{Val: 5},
			Right: &TreeNode{
				Val:   7,
				Right: &TreeNode{Val: 8},
			},
		},
	}
	root = bstToGst(root)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(curr)
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}
}
