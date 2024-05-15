package main

import "fmt"

func majorityElement(nums []int) []int {

	res := []int{}
	countMap := make(map[int]int)
	n := len(nums) / 3
	for _, val := range nums {
		_, ok := countMap[val]
		if ok {
			countMap[val] += 1
		} else {
			countMap[val] = 1
		}
	}

	for key, val := range countMap {
		if val > n {
			res = append(res, key)
		}
	}
	return res
}

func main() {

	num := []int{1, 2}
	fmt.Println(majorityElement(num))
}
