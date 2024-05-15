package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	Node   *TreeNode
	Prev   *TreeNode
	Parent *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	queue := []*Queue{{Node: root, Parent: nil, Prev: root}}
	for len(queue) > 0 {
		nextLevlel := []*Queue{}
		sum := 0
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			sum += curr.Prev.Val
			if curr.Prev.Left != nil {
				nextLevlel = append(nextLevlel,
					&Queue{Node: curr.Node.Left, Parent: curr.Node, Prev: curr.Prev.Left},
				)
			}
			if curr.Prev.Right != nil {
				nextLevlel = append(nextLevlel,
					&Queue{Node: curr.Node.Right, Parent: curr.Node, Prev: curr.Prev.Right},
				)
			}
		}
		fmt.Println("-----------------------")
		for i := 0; i < len(nextLevlel); i++ {
			fmt.Println("Level Node => ", nextLevlel[i].Prev)
			if len(nextLevlel[i:]) > 1 {
			} else {
				nextLevlel[i].Prev.Val = 0
			}
		}
		queue = append(queue, nextLevlel...)
	}
	return root
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 10},
		},
		Right: &TreeNode{
			Val: 9,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	stack := []*TreeNode{replaceValueInTree(root)}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Println(curr.Val)
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}
}
