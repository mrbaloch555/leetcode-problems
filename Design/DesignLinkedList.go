package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Node *Node
	Tail *Node
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	curr := this.Node
	for index != 0 && curr != nil {
		curr = curr.Next
		index--
	}
	if curr == nil {
		return -1
	}
	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.Size++
	node := &Node{Val: val}
	if this.Node == nil {
		this.Node = node
		this.Tail = node
	} else {
		head := this.Node
		this.Node = node
		this.Node.Next = head
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.Size++
	node := &Node{Val: val}
	if this.Tail == nil {
		this.Node = node
		this.Tail = node
	} else {
		this.Tail.Next = node
		this.Tail = this.Tail.Next
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	node := &Node{Val: val}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.Size {
		this.AddAtTail(val)
		return
	}
	curr := this.Node
	for index != 1 {
		curr = curr.Next
		index--
	}
	node.Next = curr.Next
	curr.Next = node
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	if index == 0 {
		this.Node = this.Node.Next
		if this.Size == 1 {
			this.Tail = nil
		}
		this.Size--
		return
	}
	prev := this.Node
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}
	if prev.Next == this.Tail {
		this.Tail = prev
	}
	prev.Next = prev.Next.Next
	if prev.Next == nil {
		this.Tail = prev
	}
	this.Size--
}

func main() {
	list := Constructor()
	list.AddAtHead(1)
	fmt.Println(list.Node)
	list.AddAtTail(3)
	fmt.Println(list.Tail)
	list.AddAtIndex(1, 2)
	fmt.Println(list.Node.Next, list.Tail)
	fmt.Println(list.Get(1))
	list.DeleteAtIndex(2)
	fmt.Println(list.Node, list.Tail)
	// fmt.Println(list.Get(1))
}
