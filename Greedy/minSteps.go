package main

import "fmt"

func minSteps(n int) int {
	minStep := 0
	curr := 1
	n -= 1
	for curr < n {
		curr += curr
		fmt.Println(curr)
		if curr > n {
			curr -= (curr / 2)
			break
		}
		minStep++
	}
	minStep += n - curr + 1
	return minStep
}

func main() {
	n := 3
	fmt.Println(minSteps(n))
}
