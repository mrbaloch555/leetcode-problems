package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)

	leftProducts := make([]int, n)
	rightProducts := make([]int, n)

	leftProduct := 1
	for i := 0; i < n; i++ {
		leftProducts[i] = leftProduct
		leftProduct *= nums[i]
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		rightProducts[i] = rightProduct
		rightProduct *= nums[i]
	}

	answer := make([]int, n)
	for i := 0; i < n; i++ {
		answer[i] = leftProducts[i] * rightProducts[i]
	}

	return answer
}

func main() {

	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}
