package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var newHead *ListNode
	curr := newHead
	carry := 0
	for l1 != nil && l2 != nil {
		value := l1.Val + l2.Val + carry
		nodeVal := value % 10
		carry = value / 10

		if newHead == nil {
			newHead = &ListNode{Val: nodeVal}
			curr = newHead
		} else {
			curr.Next = &ListNode{Val: nodeVal}
			curr = curr.Next
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	addNodeToList(l1, &carry, &curr, &newHead)
	addNodeToList(l2, &carry, &curr, &newHead)

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return newHead
}

func addNodeToList(l *ListNode, carry *int, curr **ListNode, newHead **ListNode) {
	for l != nil {
		value := l.Val + *carry
		nodeVal := value % 10
		*carry = value / 10

		if *newHead == nil {
			*newHead = &ListNode{Val: nodeVal}
			*curr = *newHead
		} else {
			(*curr).Next = &ListNode{Val: nodeVal}
			*curr = (*curr).Next
		}

		l = l.Next
	}
}

func main() {
	l1 := ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 1,
			},
		},
	}
	l2 := ListNode{
		Val: 1,
	}

	head := addTwoNumbers(&l1, &l2)
	for head != nil {
		fmt.Println(head)
		head = head.Next
	}
}

// 0  0  0  0  4  0
// 0  0  2  0  0  0
// 0  0  0  0  0  0
// 0  0  0  3  0  0
// 0  0  2  0  0  0
// 0  0  0
