package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	traverseFlatten(root)
}

func traverseFlatten(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftNodes := traverseFlatten(root.Left)
	rightNodes := traverseFlatten(root.Right)

	if leftNodes != nil {
		tmp := root.Right
		root.Right = root.Left
		root.Left = nil
		leftNodes.Right = tmp
	}
	if rightNodes != nil {
		return rightNodes
	}
	if leftNodes != nil {
		return leftNodes
	}
	return root
}
func main() {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}},
	}
	flatten(root)
	stack := []*TreeNode{root}
	leftValues := []int{}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(curr.Val)
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			leftValues = append(leftValues, curr.Val)
		}
	}
	fmt.Println(leftValues)
}
