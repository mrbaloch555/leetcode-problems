package main

import "fmt"

func main() {

	jums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(jums))
}
func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	jumps := 0
	currentMax := 0
	farthest := 0

	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == currentMax {
			jumps++
			currentMax = farthest
		}
	}

	return jumps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
