package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	var explore func(head *ListNode, prev *ListNode, left int, right int) *ListNode
	var start ListNode
	var end ListNode
	explore = func(head, prev *ListNode, left int, right int) *ListNode {
		left--
		if head == nil || left == 1 {
			start = *head
			return prev
		}
		node := explore(head.Next, head, left, right)
		for right != 1 {

			right--
		}
		head.Next = prev
		prev.Next = nil
		return node
	}
	return explore(head, head, left, right)

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

	reverse := reverseBetween(head, 2, 4)
	current := reverse
	fmt.Println(current)

	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}
