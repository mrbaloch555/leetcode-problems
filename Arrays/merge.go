package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	n := 3
	m := 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	for i := 0; i < len(nums1); i++ {
		if nums1[i] == 0 {
			if len(nums2) > 0 {
				nums1[i] = nums2[len(nums2)-1]
				nums2 = nums2[:len(nums2)-1]
			}
		}
	}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
}
