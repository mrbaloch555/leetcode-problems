package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	if dfsFindPath(root, head) {
		return true
	}

	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfsFindPath(root *TreeNode, head *ListNode) bool {
	if head == nil {
		return true
	}

	if root == nil || root.Val != head.Val {
		return false
	}

	return dfsFindPath(root.Left, head.Next) || dfsFindPath(root.Right, head.Next)
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6}}},
	}

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 6},
			},
			Right: &TreeNode{
				Val:   8,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
		},
	}

	fmt.Println(isSubPath(head, root))
}
