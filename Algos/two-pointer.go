package main

import (
	"fmt"
)

func isPalindrom(nums []int) bool {
	start := 0
	end := len(nums) - 1

	for start < end {
		if nums[start] != nums[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[right] + nums[left]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			right--
		} else {
			left++
		}
	}
	return []int{0, 0}
}

func main() {
	nums := []int{3, 2, 4}
	fmt.Println(isPalindrom(nums))
	fmt.Println(twoSum(nums, 6))
}
