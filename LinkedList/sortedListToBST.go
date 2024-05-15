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

func findMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if prev != nil {
		prev.Next = nil
	}
	return slow
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	mid := findMiddle(head)
	root := &TreeNode{Val: mid.Val}
	if head == mid {
		return root
	}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(mid.Next)
	return root
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Printf("%d ", root.Val)
	inorderTraversal(root.Right)
}
func main() {
	// Example usage
	head := &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
	root := sortedListToBST(head)
	fmt.Println("Inorder traversal of the BST:")
	inorderTraversal(root)
}
