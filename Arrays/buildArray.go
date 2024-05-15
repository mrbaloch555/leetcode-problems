// **********************************************************************************
// Problem Number: 1920
// Problem Name: Build Array from Permutation
// Problem Link: https://leetcode.com/problems/build-array-from-permutation/
// Status: Solved
// **********************************************************************************
package main

import "fmt"

func main() {

	arr := []int{0, 2, 1, 5, 3, 4}
	fmt.Println(buildArray(arr))
}

func buildArray(nums []int) []int {
	var newArr = make([]int, 0)

	for i := 0; i < len(nums); i++ {
		newArr = append(newArr, nums[nums[i]])
	}
	return newArr
}
