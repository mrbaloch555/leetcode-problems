package main

import (
	"container/heap"
	"fmt"
)

type SeatManager struct {
	unreserved *SeatHeap
	reserved   map[int]bool
}

type SeatHeap []int

func (h SeatHeap) Len() int            { return len(h) }
func (h SeatHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h SeatHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *SeatHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *SeatHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Constructor(n int) SeatManager {
	unreserved := &SeatHeap{}
	reserved := make(map[int]bool)
	for i := 1; i <= n; i++ {
		*unreserved = append(*unreserved, i)
	}
	heap.Init(unreserved)
	return SeatManager{unreserved, reserved}
}

func (this *SeatManager) Reserve() int {
	seat := heap.Pop(this.unreserved).(int)
	this.reserved[seat] = true
	return seat
}

func (this *SeatManager) Unreserve(seatNumber int) {
	delete(this.reserved, seatNumber)
	heap.Push(this.unreserved, seatNumber)
}

func main() {
	obj := Constructor(3)
	fmt.Println(obj.Reserve())
	fmt.Println(obj.Reserve())
	obj.Unreserve(1)
	obj.Unreserve(2)
	fmt.Println(obj.Reserve())
	obj.Unreserve(1)
	fmt.Println(obj.Reserve())
	obj.Unreserve(1)
	// fmt.Println(obj.Reserve())
	// fmt.Println(obj.Reserve())
	// fmt.Println(obj.Reserve())
	// fmt.Println(obj.Reserve())
	// obj.Unreserve(5)

	fmt.Println(obj)
}
