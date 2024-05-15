package main

import (
	"fmt"
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	res := [][]int{}

	for i := 0; i < n; i += 3 {
		t := nums[i : i+3]
		if t[2]-t[0] > k {
			return [][]int{}
		}
		res = append(res, t)
	}

	return res
}

func main() {
	nums := []int{1, 3, 4, 8, 7, 9, 3, 5, 1}
	k := 2
	fmt.Println(divideArray(nums, k))
}
