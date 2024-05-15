package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type Worker struct {
	wagePerQuality float64
	quality        int
}

type Workers []Worker

func (w Workers) Len() int           { return len(w) }
func (w Workers) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w Workers) Less(i, j int) bool { return w[i].wagePerQuality < w[j].wagePerQuality }

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	ans := math.MaxFloat64
	qualitySum := 0
	workers := make(Workers, len(quality))
	maxHeap := &IntHeap{}
	heap.Init(maxHeap)

	for i := 0; i < len(quality); i++ {
		workers[i] = Worker{float64(wage[i]) / float64(quality[i]), quality[i]}
	}

	sort.Sort(workers)

	for _, w := range workers {
		heap.Push(maxHeap, w.quality)
		qualitySum += w.quality
		if maxHeap.Len() > k {
			qualitySum -= heap.Pop(maxHeap).(int)
		}
		if maxHeap.Len() == k {
			ans = math.Min(ans, float64(qualitySum)*w.wagePerQuality)
		}
	}

	return ans
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	quality := []int{10, 20, 5}
	wage := []int{70, 50, 30}
	k := 2
	fmt.Println(mincostToHireWorkers(quality, wage, k)) // Output: 105.0

	quality2 := []int{3, 1, 10, 10, 1}
	wage2 := []int{4, 8, 2, 2, 7}
	k2 := 3
	fmt.Println(mincostToHireWorkers(quality2, wage2, k2)) // Output: 30.666666666666668
}
