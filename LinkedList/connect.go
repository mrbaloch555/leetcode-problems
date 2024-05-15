package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		nextLevel := []*Node{}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			if len(queue) > 0 {
				curr.Next = queue[0]
			}

			if curr.Left != nil {
				nextLevel = append(nextLevel, curr.Left)
			}
			if curr.Right != nil {
				nextLevel = append(nextLevel, curr.Right)
			}
		}
		queue = append(queue, nextLevel...)
	}
	return root
}

func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 7},
		},
	}
	root = connect(root)
	stack := []*Node{root}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}

		fmt.Println(curr.Val, curr.Next)
	}
}
