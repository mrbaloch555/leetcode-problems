package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	pre := dummy

	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	current := pre.Next
	for i := 0; i < right-left; i++ {
		next := current.Next
		current.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummy.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	node := reverseBetween(head, 2, 4)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
