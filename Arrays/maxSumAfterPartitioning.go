package main

import (
	"fmt"
)

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		maxInPartition := 0
		for j := i; j > max(0, i-k); j-- {
			maxInPartition = max(maxInPartition, arr[j-1])
			dp[i] = max(dp[i], dp[j-1]+maxInPartition*(i-j+1))
		}
	}
	fmt.Println(dp)

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}
	fmt.Println(maxSumAfterPartitioning(arr, 4))
}
