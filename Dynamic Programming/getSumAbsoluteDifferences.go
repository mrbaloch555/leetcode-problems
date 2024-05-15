package main

import "fmt"

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	totalSum := func(nums []int) int {
		sum := 0
		for _, val := range nums {
			sum += val
		}
		return sum
	}(nums)

	leftSum, rightSum := 0, totalSum
	result := make([]int, 0, n)
	for i, val := range nums {
		result = append(result, (i*val-leftSum)+(rightSum-(n-i)*val))
		leftSum += val
		rightSum -= val
	}
	return result
}

func main() {
	nums := []int{2, 3, 5}
	fmt.Println(getSumAbsoluteDifferences(nums))
}
