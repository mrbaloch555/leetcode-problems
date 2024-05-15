package main

import "fmt"

type Node struct {
	Key  int
	Val  int
	Next *Node
}

type LRUCache struct {
	Head     *Node
	Size     int
	Capacity int
	Keys     map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Size:     0,
		Capacity: capacity,
		Keys:     make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	node, exists := this.Keys[key]
	if !exists {
		return -1
	}

	this.moveToFront(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.Keys[key]; exists {
		node.Val = value
		this.moveToFront(node)
		return
	}

	node := &Node{Key: key, Val: value, Next: this.Head}
	this.Keys[key] = node
	this.Head = node
	this.Size++

	if this.Size > this.Capacity {
		this.removeTail()
	}
}

func (this *LRUCache) moveToFront(node *Node) {
	if this.Head == node {
		return
	}

	prev := this.Head
	for prev.Next != node {
		prev = prev.Next
	}

	prev.Next = node.Next
	node.Next = this.Head
	this.Head = node
}

func (this *LRUCache) removeTail() {
	if this.Head == nil {
		return
	}

	if this.Head.Next == nil {
		delete(this.Keys, this.Head.Key)
		this.Head = nil
		this.Size = 0
		return
	}

	prev := this.Head
	for prev.Next.Next != nil {
		prev = prev.Next
	}

	delete(this.Keys, prev.Next.Key)
	prev.Next = nil
	this.Size--
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}
