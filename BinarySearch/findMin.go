package main

import "fmt"

func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1

	for low < high {
		mid := low + (high-low)/2

		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return nums[low]
}

func main() {
	nums := []int{11, 13, 15, 17}
	fmt.Println(findMin(nums))
}
