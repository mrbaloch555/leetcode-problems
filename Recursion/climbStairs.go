// **********************************************************************************
// Problem Number: 70
// Problem Name: Climbing Stairs
// Problem Link: https://leetcode.com/problems/climbing-stairs/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

func main() {

	n := 5
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	return calculate(n, make(map[int]int))
}

func calculate(n int, memo map[int]int) int {
	if n <= 1 {
		return 1
	}
	v, ok := memo[n]

	if ok {
		return v
	}
	result := calculate(n-1, memo) + calculate(n-2, memo)
	memo[n] = result
	return result
}
