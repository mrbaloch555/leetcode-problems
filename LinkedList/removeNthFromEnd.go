package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}
	if n >= length {
		return head.Next
	}

	positionToRemove := length - n
	current = head
	for i := 1; i < positionToRemove; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next

	return head
}

func main() {
	head := &ListNode{Val: 1}
	n := 1
	node := removeNthFromEnd(head, n)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
