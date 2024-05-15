package main

import "fmt"

type Node struct {
	Val  string
	Next *Node
	Prev *Node
}

type BrowserHistory struct {
	Node *Node
	Tail *Node
}

func Constructor(homepage string) BrowserHistory {
	node := &Node{Val: homepage}
	node.Prev = nil
	return BrowserHistory{
		Node: node,
		Tail: node,
	}
}

func (this *BrowserHistory) Visit(url string) {
	node := &Node{Val: url}
	node.Prev = this.Tail
	this.Tail.Next = node
	this.Tail = node
}

func (this *BrowserHistory) Back(steps int) string {
	curr := this.Tail
	for curr.Prev != nil && steps != 0 {
		steps--
		curr = curr.Prev
	}
	this.Tail = curr
	return curr.Val
}

func (this *BrowserHistory) Forward(steps int) string {
	curr := this.Tail
	for curr.Next != nil && steps != 0 {
		steps--
		curr = curr.Next
	}
	this.Tail = curr
	return curr.Val
}

func main() {
	obj := Constructor("leetcode.com")
	obj.Visit("google.com")
	obj.Visit("facebook.com")
	obj.Visit("youtube.com")
	fmt.Println(obj.Back(1))
	fmt.Println(obj.Back(1))
	fmt.Println(obj.Forward(1))
	obj.Visit("linkedin.com")
	fmt.Println(obj.Forward(2))
	fmt.Println(obj.Back(2))
	fmt.Println(obj.Back(7))
	// fmt.Println(obj.Tail)
}
