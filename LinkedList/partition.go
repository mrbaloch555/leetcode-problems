package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	lesserDummy := &ListNode{}
	greaterOrEqualDummy := &ListNode{}
	lesserTail := lesserDummy
	greaterOrEqualTail := greaterOrEqualDummy

	current := head
	for current != nil {
		if current.Val < x {
			lesserTail.Next = current
			lesserTail = lesserTail.Next
		} else {
			greaterOrEqualTail.Next = current
			greaterOrEqualTail = greaterOrEqualTail.Next
		}
		current = current.Next
	}

	lesserTail.Next = greaterOrEqualDummy.Next
	greaterOrEqualTail.Next = nil

	return lesserDummy.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
		},
	}
	node := partition(head, 3)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
