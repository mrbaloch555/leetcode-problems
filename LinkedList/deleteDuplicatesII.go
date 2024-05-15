package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy
	curr := head

	for curr != nil {
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr = curr.Next
		}

		if prev.Next != curr {
			prev.Next = curr.Next
		} else {
			prev = prev.Next
		}

		curr = curr.Next
	}

	return dummy.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
	}

	node := deleteDuplicates(head)

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
