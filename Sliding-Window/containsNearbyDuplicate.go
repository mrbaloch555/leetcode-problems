package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	numsMap := map[int]int{}

	for i, val := range nums {
		found, ok := numsMap[val]
		if ok && i-found <= k {
			return true
		}
		numsMap[val] = i
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	nums := []int{1, 0, 1, 1}
	k := 1
	fmt.Println(containsNearbyDuplicate(nums, k))
}
