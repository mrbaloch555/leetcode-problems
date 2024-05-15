package main

import "fmt"

func main() {
	arr := []int{3, 5, 1}
	fmt.Println(canMakeArithmeticProgression(arr))
}

func canMakeArithmeticProgression(arr []int) bool {
	fmt.Println(permutations(arr))
	return true
}

func permutations(arr []int) bool {
	var helper func([]int, int)
	helper(arr, len(arr))
	return true
}

func helper(arr []int, n int) bool {
	if n == 1 {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		var def int
		if len(tmp) >= 2 {
			def = tmp[0] - tmp[1]
		}
		for i := 1; i < len(tmp)-1; i++ {
			sub := tmp[i] - tmp[i+1]
			if sub != def {
				return false
			}
		}
	} else {
		for i := 0; i < n; i++ {
			helper(arr, n-1)
			if n%2 == 1 {
				tmp := arr[i]
				arr[i] = arr[n-1]
				arr[n-1] = tmp
			} else {
				tmp := arr[0]
				arr[0] = arr[n-1]
				arr[n-1] = tmp
			}
		}
	}
	return true
}
