package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l <= r && numbers[l]+numbers[r] != target {
		if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}
	return []int{l + 1, r + 1}
}

func main() {
	numbers := []int{2, 3, 4}
	target := 6
	fmt.Println(twoSum(numbers, target))
}
