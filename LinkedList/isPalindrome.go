package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	stack := []int{}

	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] != slow.Val {
			return false
		}
		slow = slow.Next
	}

	return true
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	stack := []int{}

	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] != slow.Val {
			return false
		}
		slow = slow.Next
	}

	return true
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	fmt.Println(isPalindrome(head))
}
