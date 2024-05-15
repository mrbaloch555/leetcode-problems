package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-1}
	k := 1

	fmt.Println(findMaxAverage(nums, k))
}

func findMaxAverage(nums []int, k int) float64 {
	max := math.Inf(-1)
	sum := 0.0
	for i := 0; i <= len(nums)-k; i++ {
		sub_array := nums[i : i+k]
		if i == 0 {
			func(nums []int) {
				for _, val := range sub_array {
					sum += float64(val)
				}
			}(sub_array)
		} else {
			sum += float64(sub_array[len(sub_array)-1])
		}
		avg := sum / float64(len(sub_array))
		if avg > max {
			max = avg
		}
		sum -= float64(sub_array[0])
	}
	return max
}
