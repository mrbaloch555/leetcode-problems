package main

import "fmt"

func main() {

	nums := []int{8}
	k := 10000

	fmt.Println(getAverages(nums, k))
}

func getAverages(nums []int, k int) []int {

	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if len(nums[0:i]) >= k && len(nums[i+1:]) >= k {
			arr := nums[i-k : i+k+1]
			sum := 0
			for j := 0; j < len(arr); j++ {
				sum += arr[j]
			}
			res = append(res, sum/len(arr))
		} else {
			res = append(res, -1)
		}
	}
	return res
}
