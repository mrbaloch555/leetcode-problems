package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
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
	slow = explore(slow, slow)
	left := head
	right := slow
	max := 0

	for right != nil {
		sum := left.Val + right.Val
		if sum > max {
			max = sum
		}
		right = right.Next
		left = left.Next
	}
	return max
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	node := reverse(head.Next)
	head = node.Next.Next
	return node
}

func main() {

	a := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(pairSum(a))
}
