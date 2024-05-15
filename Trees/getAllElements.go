package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var result []int
	stack1, stack2 := []*TreeNode{}, []*TreeNode{}

	// Helper function to push left nodes to stack
	pushLeft := func(node *TreeNode, stack *[]*TreeNode) {
		for node != nil {
			*stack = append(*stack, node)
			node = node.Left
		}
	}

	pushLeft(root1, &stack1)
	pushLeft(root2, &stack2)

	for len(stack1) > 0 || len(stack2) > 0 {
		var n1, n2 *TreeNode
		if len(stack1) > 0 {
			n1 = stack1[len(stack1)-1]
		}
		if len(stack2) > 0 {
			n2 = stack2[len(stack2)-1]
		}

		if n1 != nil && (n2 == nil || n1.Val <= n2.Val) {
			result = append(result, n1.Val)
			stack1 = stack1[:len(stack1)-1]
			pushLeft(n1.Right, &stack1)
		} else if n2 != nil && (n1 == nil || n2.Val < n1.Val) {
			result = append(result, n2.Val)
			stack2 = stack2[:len(stack2)-1]
			pushLeft(n2.Right, &stack2)
		}
	}

	return result
}
func main() {
	l1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	l2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(getAllElements(l1, l2))
}
