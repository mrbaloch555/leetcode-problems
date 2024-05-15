package main

import (
	"fmt"
	"sort"
)

// func twoSum(nums []int, target int) []int {

// 	numMap := map[int]int{}

// 	for i, val := range nums {
// 		compl := target - val
// 		if _, ok := numMap[compl]; ok {
// 			return []int{numMap[compl], i}
// 		} else {
// 			numMap[val] = i
// 		}
// 		fmt.Println(numMap)
// 	}
// 	return []int{}
// }

func twoSum(nums []int, target int) []int {
	right, left := 0, len(nums)-1
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	sortedIndices := make([]int, len(nums))
	for i := range sortedIndices {
		sortedIndices[i] = i
	}

	sort.Slice(sortedIndices, func(i, j int) bool {
		return nums[sortedIndices[i]] < nums[sortedIndices[j]]
	})

	for right < left {
		sum := nums[sortedIndices[right]] + nums[sortedIndices[left]]
		if sum == target {
			return []int{sortedIndices[right], sortedIndices[left]}
		} else if sum > target {
			left--
		} else {
			right++
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
