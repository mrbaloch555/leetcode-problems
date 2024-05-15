package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}

	return head
}
func main() {
	head := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val: 7,
						},
					},
				},
			},
		},
	}
	fmt.Println(removeElements(head, 7))
}
