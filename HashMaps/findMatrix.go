package main

import "fmt"

func findMatrix(nums []int) [][]int {
	hashTable := map[int]int{}

	for _, val := range nums {
		hashTable[val]++
	}
	res := [][]int{}
	for len(hashTable) > 0 {
		subset := []int{}
		for key, _ := range hashTable {
			subset = append(subset, key)
			hashTable[key]--
			if hashTable[key] == 0 {
				delete(hashTable, key)
			}
		}
		res = append(res, subset)
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(findMatrix(nums))
}
