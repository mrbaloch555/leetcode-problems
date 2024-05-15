package main

import "fmt"

func rearrangeArray(nums []int) []int {

	pos := filter(nums, func(num int) bool {
		return num >= 0
	})

	neg := filter(nums, func(num int) bool {
		return num < 0
	})

	result := []int{}

	for i := 0; i < len(pos); i++ {
		if len(pos) == 0 || len(neg) == 0 {
			break
		}

		posvalue := pos[0]
		pos = pos[1:]
		negvalue := neg[0]
		neg = neg[1:]

		result = append(result, []int{posvalue, negvalue}...)
		i--
	}

	if len(pos) > 0 {
		result = append(result, pos...)
	} else if len(neg) > 0 {
		result = append(result, neg...)
	}
	return result
}

func filter(nums []int, cond func(num int) bool) []int {
	result := []int{}

	for _, val := range nums {
		if cond(val) {
			result = append(result, val)
		}
	}
	return result
}

func main() {

	nums := []int{3, 1, -2, -5, 2, -4}
	fmt.Println(rearrangeArray(nums))
}
