package main

import (
	"fmt"
)

func fullBloomFlowers(flowers [][]int, people []int) []int {
	maxDay := 0
	for _, interval := range flowers {
		maxDay = max(maxDay, interval[1])
	}

	fenwickTree := make([]int, maxDay+1)
	add := func(index, value int) {
		for i := index; i <= maxDay; i += i & -i {
			fmt.Println(i)
			fenwickTree[i] += value
		}
	}

	query := func(index int) int {
		sum := 0
		for i := index; i > 0; i -= i & -i {
			sum += fenwickTree[i]
		}
		return sum
	}

	for _, interval := range flowers {
		add(interval[0], 1)
		add(interval[1]+1, -1)
	}

	result := make([]int, len(people))
	for i, day := range people {
		result[i] = query(day)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	flowers := [][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}
	people := []int{2, 3, 7, 11}

	fmt.Println(fullBloomFlowers(flowers, people))
}
