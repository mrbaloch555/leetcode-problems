// **********************************************************************************
// Problem Number: 771
// Problem Name: Jewels and Stones
// Problem Link: https://leetcode.com/problems/jewels-and-stones/description/
// Status: Solved
// **********************************************************************************
package main

import "fmt"

func main() {
	jewls := "aA"
	stones := "aAAbbbb"

	fmt.Println(numJewelsInStones(jewls, stones))
}

func numJewelsInStones(jewels string, stones string) int {
	var count int
	obj := make(map[string]string, 0)
	for i := range jewels {
		obj[string(jewels[i])] = string(jewels[i])
	}

	for i := range stones {
		_, exists := obj[string(stones[i])]
		if exists {
			count++
		}
	}
	return count
}
