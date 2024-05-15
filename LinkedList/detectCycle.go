package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			break
		}
	}
	fast = head
	for slow != nil && slow.Next != nil && fast != nil && fast.Next != nil {
		if fast == slow {
			return fast
		}
		fast = fast.Next
		slow = slow.Next
	}
	return nil
}

func main() {
	head := &ListNode{Val: 3}
	two := &ListNode{Val: 2}
	zero := &ListNode{Val: 0}
	four := &ListNode{Val: 4}
	head.Next = two
	two.Next = zero
	zero.Next = four
	four.Next = two
	fmt.Println(detectCycle(head))
}
