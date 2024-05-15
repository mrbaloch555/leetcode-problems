package main

import "fmt"

func findArray(pref []int) []int {

	n := len(pref)
	original_array := make([]int, n+1)
	original_array[0] = pref[0]
	for i := 1; i < n+1; i++ {
		original_array[i] = pref[i-1] ^ original_array[i-1]
	}
	return original_array[1:]
}

func main() {
	pref := []int{5, 2, 0, 3, 1}
	fmt.Println(findArray(pref))
}
