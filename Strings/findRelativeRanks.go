package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	copiedArray := make([]int, len(score))
	for i, v := range score {
		copiedArray[i] = v
	}
	sort.Slice(copiedArray, func(i, j int) bool {
		return copiedArray[i] > copiedArray[j]
	})
	hashMap := map[int]string{}
	res := make([]string, len(score))
	for i, val := range copiedArray {
		if i < 3 {
			switch i {
			case 0:
				hashMap[val] = "Gold Medal"
			case 1:
				hashMap[val] = "Silver Medal"
			case 2:
				hashMap[val] = "Bronze Medal"
			}
		} else {
			str := strconv.Itoa(i + 1)
			hashMap[val] = str
		}
	}
	for i, val := range score {
		res[i] = hashMap[val]
	}
	return res
}

func main() {
	score := []int{10, 3, 8, 9, 4}
	fmt.Println(findRelativeRanks(score))
}
