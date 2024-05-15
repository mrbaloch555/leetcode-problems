package main

import "fmt"

type CustomStack struct {
	elements []int
	capicity int
	size     int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		elements: make([]int, maxSize),
		size:     0,
		capicity: maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if this.size < this.capicity {
		this.elements[this.size] = x
		this.size++
	}
}

func (this *CustomStack) Pop() int {
	if this.size == 0 {
		return -1
	}
	val := this.elements[this.size-1]
	this.elements[this.size-1] = 0
	this.size--
	return val
}

func (this *CustomStack) Increment(k int, val int) {
	if this.size < k {
		k = this.size
	}
	for i, _ := range this.elements[:k] {
		this.elements[i] += val
	}
}

func main() {
	stack := Constructor(3)
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop())
	stack.Push(2)
	stack.Push(3)
	stack.Increment(5, 100)
	stack.Increment(2, 100)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	// fmt.Println(stack.Pop())
	// stack.Push(3)
	// stack.Increment(2, 100)
	fmt.Println(stack)
}
