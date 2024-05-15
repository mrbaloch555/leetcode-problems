package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {

	res := []string{}
	prev := 0
	for i := 0; i < len(nums); i++ {
		if len(nums[i:]) > 1 {
			if nums[i+1]-nums[i] != 1 {
				res = append(res, strconv.Itoa(prev)+"->"+strconv.Itoa(nums[i]))
				prev = nums[i+1]
			}
		} else {
			res = append(res, strconv.Itoa(nums[i]))
		}
	}
	return res
}

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))
}
