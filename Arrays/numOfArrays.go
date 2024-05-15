package main

import "fmt"

func numOfArrays(n int, m int, k int) int {
	const kMod = int(1e9) + 7
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}

	for j := 1; j <= m; j++ {
		dp[1][j][1] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for cost := 1; cost <= k; cost++ {
				dp[i][j][cost] = (dp[i-1][j][cost] * j) % kMod
				for prevMax := 1; prevMax < j; prevMax++ {
					dp[i][j][cost] = (dp[i][j][cost] + dp[i-1][prevMax][cost-1]) % kMod
				}
			}
		}
	}

	ans := 0
	for j := 1; j <= m; j++ {
		ans = (ans + dp[n][j][k]) % kMod
	}

	return ans
}

func main() {
	n := 2
	m := 3
	k := 1
	fmt.Println(numOfArrays(n, m, k))
}
