// **********************************************************************************
// Problem Number: 169
// Problem Name:  Majority Element
// Problem Link: https://leetcode.com/problems/majority-element/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}

type maxStruct struct {
	val   int
	count int
}

func majorityElement(nums []int) int {
	obj := make(map[int]int)

	for _, val := range nums {
		_, ok := obj[val]
		if ok {
			obj[val]++
		} else {
			obj[val] = 1
		}
	}
	max := maxStruct{val: 0, count: 0}

	for key, val := range obj {
		if val > max.count {
			max.val = key
			max.count = val
		}
	}
	return max.val
}
