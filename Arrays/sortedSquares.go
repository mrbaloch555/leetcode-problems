package main

import "fmt"

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	left, right, index := 0, n-1, n-1

	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			result[index] = leftSquare
			left++
		} else {
			result[index] = rightSquare
			right--
		}

		index--
	}

	return result
}

func main() {
	nums := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums))
}
