package main

import "fmt"

func main() {
	arr := []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}
	fmt.Println(arraySign(arr))
}

func arraySign(nums []int) int {
	return func() int {
		var sum = int(1)
		for _, val := range nums {
			sum = val * sum
		}
		fmt.Println(sum)
		if sum < 0 {
			return -1
		} else if sum > 0 {
			return 1
		} else {
			return 0
		}
	}()
}
