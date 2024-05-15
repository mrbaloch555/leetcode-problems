package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{3, 5, 2, 3}
	fmt.Println(minPairSum(nums))
}

func minPairSum(nums []int) int {
	max := math.Inf(-1)
	min := math.Inf(1)
	for _, val := range nums {
		if int(max) < val {
			max = float64(val)
		}

		if math.Abs(float64(min)) > float64(val) {
			min = float64(val)
		}
	}
	return int(min) + int(max)
}
