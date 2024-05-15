package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	head = reverse(head)
	dummyHead := &ListNode{}
	currentNode := dummyHead
	multiplier := 2
	carry := 0

	for head != nil {
		product := head.Val*multiplier + carry
		carry = product / 10
		currentNode.Next = &ListNode{Val: product % 10}
		currentNode = currentNode.Next
		head = head.Next
	}

	if carry > 0 {
		currentNode.Next = &ListNode{Val: carry}
	}

	return reverse(dummyHead.Next)
}

func reverse(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	currentNode := head

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = dummyHead.Next
		dummyHead.Next = currentNode
		currentNode = nextNode
	}

	return dummyHead.Next
}

func main() {

	head := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
		},
	}

	head = doubleIt(head)

	for head != nil {
		fmt.Println(head)
		head = head.Next
	}
}
