// **********************************************************************************
// Problem Number: 1929
// Problem Name: Concatenation of Array
// Problem Link: https://leetcode.com/problems/concatenation-of-array/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

func main() {

	var arr = []int{1, 2, 1}
	fmt.Println(getConcatenation(arr))
}

func getConcatenation(arr []int) []int {
	return append(arr, arr...)
}
