package main

import "fmt"

func findUnsortedSubarray(nums []int) int {
	last := 0
	start := 0
	match := false
	for i := 0; i < len(nums); i++ {
		if len(nums[i:]) > 1 {
			if nums[i] > nums[i+1] {
				match = true
				if start != 0 && match {
					last = i + 2
				} else {
					start = i
				}
			}
		} else {
			if last == 0 && match {
				last += len(nums)
			}
		}
	}
	return len(nums[start:last])
}

func main() {
	nums := []int{5, 4, 3, 2, 1}
	fmt.Println(findUnsortedSubarray(nums))
}
