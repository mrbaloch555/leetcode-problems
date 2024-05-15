// **********************************************************************************
// Problem Number: 118
// Problem Name: Pascal's Triangle
// Problem Link: https://leetcode.com/problems/pascals-triangle/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

func main() {
	num := 30
	fmt.Println(generate(num))
}

func generate(numRows int) [][]int {

	res := [][]int{}
	for i := 1; i <= numRows; i++ {
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

	return res
}
