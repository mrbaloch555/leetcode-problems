package main

import (
	"fmt"
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	numMap := map[int]int{}
	for _, val := range arr {
		numMap[val]++
	}
	uniques := [][]int{}

	for key, val := range numMap {
		uniques = append(uniques, []int{key, val})
	}
	sort.Slice(uniques, func(i, j int) bool {
		return uniques[i][1] < uniques[j][1]
	})

	for k > 0 {
		curr := uniques[0]
		if curr[1] == 1 {
			uniques = uniques[1:]
		} else {
			uniques[0][1]--
		}
		k--
	}

	return len(uniques)
}
func main() {

	arr := []int{4, 3, 1, 1, 3, 3, 2}
	n := 3
	fmt.Println(findLeastNumOfUniqueInts(arr, n))
}
