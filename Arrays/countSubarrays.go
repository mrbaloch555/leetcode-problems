package main

import "fmt"

func countSubarrays(nums []int, k int64) int64 {
	count := 0
	sum := 0
	l := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum*(i-l+1) >= int(k) {
			sum -= nums[l]
			l++
		}
		count += i - l + 1
	}
	return int64(count)
}

func main() {
	nums := []int{7, 6, 7, 9, 1, 5, 1, 4, 9, 1, 10, 8, 4, 1, 7, 4, 2}
	var k int64 = 65
	fmt.Println(countSubarrays(nums, k))
}
