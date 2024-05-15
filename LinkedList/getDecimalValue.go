package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	result := 0
	current := head

	for current != nil {
		result = (result << 1) + current.Val
		current = current.Next
	}

	return result
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 0}}},
	}

	fmt.Println(getDecimalValue(head))
}
