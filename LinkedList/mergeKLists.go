package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode

	for _, val := range lists {
		head = merge(head, val)
	}
	return head
}

func merge(p1, p2 *ListNode) *ListNode {
	if p1 == nil {
		return p2
	}
	if p2 == nil {
		return p1
	}

	if p1.Val < p2.Val {
		p1.Next = merge(p1.Next, p2)
		return p1
	} else {
		p2.Next = merge(p1, p2.Next)
		return p2
	}
}

func main() {
	list := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{Val: 2, Next: &ListNode{Val: 6}},
	}

	head := mergeKLists(list)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
