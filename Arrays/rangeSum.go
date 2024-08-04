package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func rangeSum(nums []int, n int, left int, right int) int {
	const MOD int = 1000000007
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for i := 0; i < n; i++ {
		curr := 0
		for j := i; j < n; j++ {
			curr += nums[j]
			heap.Push(minHeap, curr)
		}
	}

	var result int
	for i := 1; i <= right; i++ {
		val := heap.Pop(minHeap).(int)
		if i >= left {
			result = (result + val) % MOD
		}
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	n := 4
	left := 1
	right := 5
	fmt.Println(rangeSum(nums, n, left, right))
}
