// **********************************************************************************
// Problem Number: 1512
// Problem Name: Number of Good Pairs
// Problem Link: https://leetcode.com/problems/number-of-good-pairs/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

func main() {
	var nums = []int{1, 1, 1, 1}

	fmt.Println(numIdenticalPairs(nums))
}

func numIdenticalPairs(nums []int) int {
	return func() int {
		pairs := 0
		for i, _ := range nums {
			for j := i; j < len(nums); j++ {
				if i < j && nums[i] == nums[j] {
					pairs++
				}
			}
		}
		return pairs
	}()
}
