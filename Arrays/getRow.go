// **********************************************************************************
// Problem Number: 119
// Problem Name: Pascal's Triangle II
// Problem Link: https://leetcode.com/problems/pascals-triangle-ii/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

func main() {
	num := 3
	fmt.Println(generate(num))
}

func generate(numRows int) []int {
	res := [][]int{}
	for i := 1; i <= numRows+1; i++ {
		d := []int{}
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				d = append(d, 1)
			} else {
				d = append(d, res[i-2][j-1]+res[i-2][j])
			}
		}
		res = append(res, [][]int{d}...)
	}

	return res[len(res)-1]
}
