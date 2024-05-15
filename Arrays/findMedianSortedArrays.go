package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, num2 []int) float64 {

	nums1 = append(nums1, num2...)
	fmt.Println(nums1)
	mid := int(math.Floor(float64(len(nums1) / 2)))
	if len(nums1)%2 == 0 {
		return float64(nums1[mid-1] + num2[mid]/2)
	} else {
		return float64(nums1[mid+1])
	}
}
