package main

type MinStack struct {
	stack []int
	min   int
	prev  int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   0,
		prev:  0,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.stack) == 1 {
		this.stack = append(this.stack, val)
		this.min = 0
		this.prev = 0
	} else {
        if this.stack[min] < 
	}
}

func (this *MinStack) Pop() {

}

func (this *MinStack) Top() int {

}

func (this *MinStack) GetMin() int {

}

func main() {

}
