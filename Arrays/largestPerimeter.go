package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] < nums[i+1]+nums[i+2] {
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}

	return -1
}
func main() {

	nums := []int{1, 12, 1, 2, 5, 50, 3}
	fmt.Println(largestPerimeter(nums))
}
