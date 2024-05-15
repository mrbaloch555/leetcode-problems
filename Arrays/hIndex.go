package main

import (
	"fmt"
	"sort"
)

func main() {

	citations := []int{3, 0, 6, 1, 5}

	fmt.Println(hIndex(citations))
}

func hIndex(citations []int) int {
	h_index := 0
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	for i, val := range citations {
		if i <= val {
			h_index = i
		} else {
			break
		}
	}
	return h_index
}
