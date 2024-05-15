package main

import (
	"fmt"
	"sort"
)

// **********************************************************************************
// Problem Number: 414
// Problem Name: Third Maximum Number
// Problem Link: https://leetcode.com/problems/third-maximum-number/
// Status: Solved
// **********************************************************************************

func main() {
	nums := []int{3, 2, 1}
	fmt.Println(thirdMax(nums))
}

func thirdMax(nums []int) int {
	obj := make(map[int]int, 0)
	unique := []int{}
	for _, val := range nums {
		_, ok := obj[val]
		if !ok {
			obj[val] = val
		}
	}
	for _, val := range obj {
		unique = append(unique, val)
	}

	sort.Slice(unique, func(i, j int) bool {
		return unique[i] > unique[j]
	})
	if len(unique) >= 3 {
		return unique[2]
	} else {
		return unique[0]
	}
}
