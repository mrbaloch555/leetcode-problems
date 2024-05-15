package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	n := 1
	current := head
	for current.Next != nil {
		current = current.Next
		n++
	}

	k = k % n

	if k == 0 {
		return head
	}

	current.Next = head

	for i := 0; i < n-k; i++ {
		current = current.Next
	}
	newHead := current.Next
	current.Next = nil

	return newHead
}

func main() {

	head := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	k := 4
	node := rotateRight(head, k)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
