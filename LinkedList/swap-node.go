// **********************************************************************************
// Problem Number: 24
// Problem Name: Swap Nodes in Pairs
// Problem Link: https://leetcode.com/problems/swap-nodes-in-pairs/
// Status: Solved
// **********************************************************************************

package main

type Node struct {
	Val  int
	Next *Node
}

func main() {
	var a = Node{Val: 1}
	var b = Node{Val: 2}
	var c = Node{Val: 3}
	var d = Node{Val: 4}

	a.Next = &b
	b.Next = &c
	c.Next = &d
	swapLinkedList(&a)
}

func swapLinkedList(head *Node) *Node {
	var current = head
	for current != nil && current.Next != nil {
		next := current.Next
		temp := current.Val
		current.Val = next.Val
		next.Val = temp
		current = next.Next
	}
	return head
}
