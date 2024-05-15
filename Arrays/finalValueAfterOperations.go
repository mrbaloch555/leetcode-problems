// **********************************************************************************
// Problem Number: 2011
// Problem Name: Final Value of Variable After Performing Operations
// Problem Link: https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

func main() {
	var oprations = []string{"++X", "++X", "X++"}

	fmt.Println(finalValueAfterOperations(oprations))
}

func finalValueAfterOperations(operations []string) int {
	return func() int {
		counter := 0
		for _, val := range operations {
			if val == "X++" || val == "++X" {
				counter++
			} else {
				counter--
			}
		}
		return counter
	}()
}
