package main

import "fmt"

func isSameSubArray(nums []int, k int) bool {
	for i, _ := range nums {
		for j := i + 1; j < min(len(nums), i+k); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func closeDuplicate(nums []int, k int) bool {

	windos := map[int]int{}
	L := 0

	for i := 0; i < len(nums); i++ {
		if i-L+1 > k {
			delete(windos, L)
			L++
		}
		if _, ok := windos[nums[L]]; ok {
			return true
		}
		windos[i] = nums[i]
	}
	return false
}

func longestSubArray(nums []int) int {
	max := 0
	L := 0

	for i := 0; i < len(nums); i++ {
		for i < len(nums) && nums[i] == nums[L] {
			i++
		}
		if i-L > max {
			max = i - L
		}
		L++
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	k := 3
	fmt.Println(isSameSubArray(nums, k))
	fmt.Println(closeDuplicate(nums, k))
	fmt.Println(longestSubArray(nums))
}
