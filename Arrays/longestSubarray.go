package main

import (
	"fmt"
	"math"
)

func main() {

	nums := []int{0, 0, 0}
	fmt.Println(longestSubarray(nums))
}

func longestSubarray(nums []int) int {
	previous := []int{}
	current_counter := 0
	max := math.Inf(1)
	found_zero := false
	found_non_zero := false
	for i, val := range nums {
		if val == 0 && i != 0 && nums[i-1] == 0 {
			previous = append(previous, 0)
		} else if val == 0 && i != 0 && nums[i-1] != 0 {
			found_zero = true
			if len(previous) >= 1 {
				if float64(current_counter+previous[len(previous)-1]) > max {
					max = float64(current_counter + previous[len(previous)-1])
					previous = append(previous, current_counter)
				}
			} else {
				max = float64(current_counter)
				previous = append(previous, current_counter)
			}
			current_counter = 0
		} else if val != 0 {
			current_counter++
			found_non_zero = true
		}
		if i == len(nums)-1 {
			if len(previous) >= 1 {
				if float64(current_counter+previous[len(previous)-1]) > max {
					max = float64(current_counter + previous[len(previous)-1])
				}
			}
		}
	}
	if found_zero && !found_non_zero {
		return int(max)
	} else if found_non_zero {
		return 0
	} else {
		return len(nums) - 1
	}
}
