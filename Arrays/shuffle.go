// **********************************************************************************
// Problem Number: 1470
// Problem Name: Shuffle the Array
// Problem Link: https://leetcode.com/problems/shuffle-the-array/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

func main() {

	arr := []int{2, 5, 1, 3, 4, 7}
	fmt.Println(shuffle(arr, 3))
}

func shuffle(nums []int, n int) []int {
	newArr := make([]int, 0)
	for i := 0; i < len(nums)/2; i++ {
		newArr = append(newArr, nums[i])
		newArr = append(newArr, nums[n])
		n++
	}
	return newArr
}
