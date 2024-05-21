package main

import (
	"fmt"
	"math"
)

func subsetXORSum(nums []int) int {
	count := 0
	n := len(nums)
	for i := 1; i < (1 << n); i++ {
		var subset []int
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				subset = append(subset, nums[j])
			}
		}
		xor := subset[0]
		for _, val := range subset[1:] {
			xor = xor ^ val
		}
		count += xor
	}

	return count
}

func main() {
	nums := []int{3, 4, 5, 6, 7, 8}
	fmt.Println(subsetXORSum(nums))
	fmt.Println(MAX(nums))
}

func MAX(nums []int) int {
	max := math.Inf(-1)

	for _, val := range nums {
		if val > int(max) {
			max = float64(val)
		}
	}
	return int(max)
}
