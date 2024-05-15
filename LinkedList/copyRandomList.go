package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = &Node{Val: curr.Val}
		curr.Next.Next = next
		curr = next
	}

	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	curr = head
	newHead := head.Next
	for curr != nil {
		next := curr.Next.Next
		if next != nil {
			curr.Next.Next = next.Next
		}
		curr.Next = next
		curr = next
	}

	return newHead
}

func main() {

	head := &Node{Val: 7}
	sec := &Node{Val: 13}
	third := &Node{Val: 11}
	fourth := &Node{Val: 10}
	fifth := &Node{Val: 1}
	head.Next = sec

	sec.Random = head
	sec.Next = third

	third.Random = fifth
	third.Next = fourth

	fourth.Random = third
	fourth.Next = fifth

	fifth.Random = head

	copiedList := copyRandomList(head)
	fmt.Println(copiedList == head)
}
