package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := []int{}
	l1 := 0
	l2 := 0
	for l1 < len(nums1) && l2 < len(nums2) {
		if nums1[l1] < nums2[l2] {
			nums = append(nums, nums1[l1])
			l1++
		} else {
			nums = append(nums, nums2[l2])
			l2++
		}
	}
	if len(nums1[l1:]) > 0 {
		nums = append(nums, nums1[l1:]...)
	} else {
		nums = append(nums, nums2[l2:]...)
	}
	mid := len(nums) / 2
	if len(nums)%2 == 0 {
		return float64(nums[mid-1]+nums[mid]) / 2
	} else {
		return float64(nums[mid])
	}
}
