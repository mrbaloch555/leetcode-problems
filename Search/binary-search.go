package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 22, 35, 67}
	fmt.Println(binary(nums, 4))
}

func binary(nums []int, target int) int {

	var low, high = 0, len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2
		if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
