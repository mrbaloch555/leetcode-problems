package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	obj1 := map[int]int{}
	obj2 := map[int]int{}

	for _, val := range arr1 {
		obj1[val] += 1
	}
	for _, val := range arr2 {
		obj2[val] = 1
	}
	absent := []int{}
	result := []int{}
	for _, val := range arr2 {
		for i := 0; i < obj1[val]; i++ {
			result = append(result, val)
		}
		delete(obj1, val)
	}
	for key, val := range obj1 {
		for i := 0; i < val; i++ {
			absent = append(absent, key)
		}
	}
	sort.Ints(absent)
	result = append(result, absent...)
	return result
}

func main() {
	arr1 := []int{2, 21, 43, 38, 0, 42, 33, 7, 24, 13, 12, 27, 12, 24, 5, 23, 29, 48, 30, 31}
	arr2 := []int{2, 42, 38, 0, 43, 21}

	fmt.Println(relativeSortArray(arr1, arr2))
}
