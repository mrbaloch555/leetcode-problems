package main

import (
	"fmt"
)

type Ride struct {
	Start, End, Tip int
}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	startToEndAndEarns := make([][]Ride, n)
	dp := make([]int, n+1)

	for _, ride := range rides {
		start, end, tip := ride[0], ride[1], ride[2]
		earn := end - start + tip
		startToEndAndEarns[start] = append(startToEndAndEarns[start], Ride{End: end, Tip: earn})
	}

	for i := n - 1; i > 0; i-- {
		dp[i] = dp[i+1]
		for _, ride := range startToEndAndEarns[i] {
			dp[i] = max(dp[i], dp[ride.End]+ride.Tip)
		}
	}
	return int64(dp[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := 20
	rides := [][]int{
		{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1},
	}

	fmt.Println(maxTaxiEarnings(n, rides))
}
