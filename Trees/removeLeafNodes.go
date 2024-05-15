package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)

	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}

	return root
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}}
	fmt.Println(removeLeafNodes(root, 2))
	stack := []*TreeNode{root}
	stack = []*TreeNode{root}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		fmt.Println(current)
	}

}
