package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{{}}
	n := len(nums)
	for i := 1; i < (1 << n); i++ {
		var subset []int
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				subset = append(subset, nums[j])
			}
		}
		res = append(res, subset)
	}

	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
