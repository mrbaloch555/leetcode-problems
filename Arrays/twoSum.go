package main

import "fmt"

func main() {
	nums := []int{-1, 2, 3, 8, 9, 4}
	target := 7
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	obj := map[int]int{}

	for i := 0; i < len(nums); i++ {
		com := target - nums[i]
		_, ok := obj[com]
		if ok {
			return []int{obj[com], i}
		}
		obj[nums[i]] = i
		fmt.Println(obj)
	}
	return []int{0, 0}
}
