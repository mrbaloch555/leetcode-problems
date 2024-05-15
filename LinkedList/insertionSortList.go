package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	for curr != nil {
		var prev *ListNode
		next := curr.Next
		prev = curr
		for next != nil && curr.Val < next.Val {
			prev = next
			next = next.Next
		}
		curr = curr.Next
	}
	return head
}

func main() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}

	head = insertionSortList(head)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
