package main

import (
	"fmt"
)

// Given an array, return true if there are two elements
// within a window of size k that are equal

func closeDuplicatesBruteForce(nums []int, k int) bool {
	for i, _ := range nums {
		for j := i + 1; j < min(len(nums), i+k); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func closeDuplicates(nums []int, k int) bool {
	numMap := map[int]bool{}
	l := 0
	for i, val := range nums {
		if l-i > k {
			delete(numMap, l)
			l += 1
		}
		if _, ok := numMap[val]; ok {
			return true
		} else {
			numMap[val] = true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 2, 3, 3}
	fmt.Println(closeDuplicatesBruteForce(nums, 4))
	fmt.Println(closeDuplicates(nums, 1))
}
