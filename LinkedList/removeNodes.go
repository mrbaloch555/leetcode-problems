package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeNodes(head.Next)
	if head.Next != nil && head.Val < head.Next.Val {
		return head.Next
	}
	return head
}
func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}
	removeNodes(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
