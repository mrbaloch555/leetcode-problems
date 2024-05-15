package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(isEven(nums, []int{}))
}

func isEven(nums []int, even []int) []int {
	if len(nums) <= 0 {
		return even
	}
	if nums[0]%2 == 0 {
		even = append(even, nums[0])
	}
	nums = append(nums[:0], nums[0+1:]...)
	return isEven(nums, even)
}
