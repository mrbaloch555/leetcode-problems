package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	if node == nil || node.Next == nil {
		return
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
	deleteNode(head.Next)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
