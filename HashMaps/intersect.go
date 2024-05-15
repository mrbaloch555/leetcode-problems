package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{4, 7, 9, 7, 6, 7}
	nums2 := []int{5, 0, 0, 6, 1, 6, 2, 2, 4}

	fmt.Println(intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
	obj := map[int][]int{}
	res := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		_, ok := obj[nums1[i]]

		if !ok {
			obj[nums1[i]] = append(obj[nums1[i]], 1)
		} else {
			obj[nums1[i]][0]++
		}
	}
	for i := 0; i < len(nums2); i++ {
		val, ok := obj[nums2[i]]

		if ok {
			if len(val) <= 1 {
				obj[nums2[i]] = append(obj[nums2[i]], 1)
			} else {
				obj[nums2[i]][1] += 1
			}
		}
	}
	for key, val := range obj {
		if len(val) >= 2 {
			min := math.Min(float64(val[0]), float64(val[1]))

			for i := 0; i < int(min); i++ {
				res = append(res, key)
			}
		}
	}
	return res
}
