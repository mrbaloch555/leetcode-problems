package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 0, 1, 2}
	moveZeros(nums)
	fmt.Println(nums)
}

func moveZeros(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			j++
			i--
		}
	}
	for i := 0; i < j; i++ {
		nums = append(nums, 0)
	}
}
