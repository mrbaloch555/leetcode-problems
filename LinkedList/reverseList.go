package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var explore func(head *ListNode, prev *ListNode) *ListNode

	explore = func(head, prev *ListNode) *ListNode {
		if head == nil {
			return prev
		}
		node := explore(head.Next, head)
		head.Next = prev
		prev.Next = nil
		return node
	}
	return explore(head, head)

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

	reverse := reverseList(head)
	current := reverse
	fmt.Println(current)

	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}
