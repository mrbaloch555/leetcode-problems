package main

import "fmt"

type MinStack struct {
	main []int
	aux  []int
}

func Constructor() MinStack {

	return MinStack{
		main: make([]int, 0),
		aux:  make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	mN := len(this.main)
	aN := len(this.aux)
	this.main = append(this.main, val)
	if mN == 0 {
		this.aux = append(this.aux, val)
	} else {
		if val < this.aux[aN-1] {
			this.aux = append(this.aux, val)
		} else {
			this.aux = append(this.aux, this.aux[aN-1])
		}
	}
}

func (this *MinStack) Pop() {
	mN := len(this.main)
	aN := len(this.aux)
	this.main = this.main[:mN-1]
	this.aux = this.aux[:aN-1]
}

func (this *MinStack) Top() int {
	return this.main[len(this.main)-1]
}

func (this *MinStack) GetMin() int {
	return this.aux[len(this.aux)-1]
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
