package main

import "fmt"

func buldCost(nums []int) int {
	cost := 0

	for i, _ := range nums {
		if cost%2 != 0 {
			nums[i] = nums[i] ^ 1
		}

		if nums[i]%2 == 0 {
			cost++
		}
	}

	return cost
}

func main() {

	nums := []int{1, 0, 1}
	fmt.Println(buldCost(nums))
}
