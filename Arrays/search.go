package main

import "fmt"

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(arr, target))
}

func search(nums []int, target int) int {

	low, high := 0, len(nums)-1

	for low < high {
		middle := (low + high) / 2
		if nums[middle] == target {
			return middle
		}

		if nums[low] <= nums[middle] {
			if target >= nums[low] && target < nums[middle] {
				high = middle
			} else {
				low = middle + 1
			}
		} else {
			if target > nums[middle] && target <= nums[high] {
				low = middle + 1
			} else {
				high = middle
			}
		}
	}
	if nums[low] == target {
		return low
	} else {
		return -1
	}
}
