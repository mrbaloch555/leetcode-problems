// **********************************************************************************
// Problem Number: 229
// Problem Name:  Majority Element II
// Problem Link: https://leetcode.com/problems/majority-element-ii/description/
// Status: Solved
// **********************************************************************************

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2}
	fmt.Println(majorityElement(nums))
}
func majorityElement(nums []int) []int {
	obj := make(map[int]int)

	for _, val := range nums {
		_, ok := obj[val]
		if ok {
			obj[val]++
		} else {
			obj[val] = 1
		}
	}
	res := []int{}
	n := int(math.Floor(float64(len(nums) / 3)))

	for key, val := range obj {
		if val > n {
			res = append(res, key)
		}
	}
	return res
}
