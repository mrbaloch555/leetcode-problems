package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func middleNode(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func hasCycle(head *Node) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func detectCycle(head *Node) *Node {

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow2 := head
	for slow != slow2 {
		slow2 = slow2.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	head := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}
	four := &Node{Val: 4}
	five := &Node{Val: 5}
	head.Next = two
	two.Next = three
	three.Next = four
	four.Next = five
	five.Next = two

	fmt.Println(hasCycle(head))
	fmt.Println(detectCycle(head))
}
