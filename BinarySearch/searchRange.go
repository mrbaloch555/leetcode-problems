package main

import "fmt"

func searchRange(nums []int, target int) []int {
	left := binarySearchLeft(nums, target)
	right := binarySearchRight(nums, target)

	if left <= right {
		return []int{left, right}
	} else {
		return []int{-1, -1}
	}
}

func binarySearchLeft(nums []int, target int) int {
	low, high := 0, len(nums)-1
	result := -1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			result = mid
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}

func binarySearchRight(nums []int, target int) int {
	low, high := 0, len(nums)-1
	result := -1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			result = mid
			low = mid + 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 6))
}
