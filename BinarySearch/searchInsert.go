package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	mid := 0
	for low <= high {
		mid = low + (high-low)/2

		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	if nums[mid] > target {
		return mid
	} else {
		return mid + 1
	}
}

func main() {

	nums := []int{1, 3, 5, 6}
	target := 0
	fmt.Println(searchInsert(nums, target))
}
