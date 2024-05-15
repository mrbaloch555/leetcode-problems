// **********************************************************************************
// Problem Number: 347
// Problem Name: Top K Frequent Elements
// Problem Link: https://leetcode.com/problems/top-k-frequent-elements/
// Status: Solved
// **********************************************************************************

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1}
	k := 1
	fmt.Println(topKFrequent(nums, k))
}
func topKFrequent(nums []int, k int) []int {
	obj := make(map[int]int, 0)
	keys := make([]int, 0)
	values := make([]int, 0)
	topK := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		_, ok := obj[nums[i]]
		if ok {
			obj[nums[i]]++
		} else {
			obj[nums[i]] = 1
		}
	}
	for key, val := range obj {
		keys = append(keys, key)
		values = append(values, val)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	for _, val := range values {
		if k != 0 {
			for key, val_ := range obj {
				if val_ == val {
					topK = append(topK, key)
					delete(obj, key)
					k--
					break
				}
			}
		}
	}
	return topK
}
