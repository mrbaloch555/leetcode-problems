package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLen := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentLen := 1

			for numSet[currentNum+1] {
				currentNum++
				currentLen++
			}

			if currentLen > maxLen {
				maxLen = currentLen
			}
		}
	}

	return maxLen
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
