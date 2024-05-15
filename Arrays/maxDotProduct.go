package main

import "fmt"

func maxDotProduct(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]+nums1[i-1]*nums2[j-1], nums1[i-1]*nums2[j-1])
		}
	}
	return dp[m][n]
}

func max(a ...int) int {
	fmt.Println(a)
	result := a[0]
	for _, val := range a {
		if val > result {
			result = val
		}
	}
	return result
}
func main() {
	nums1 := []int{1, 1}
	nums2 := []int{-1, -1}
	fmt.Println(maxDotProduct(nums1, nums2))
}
