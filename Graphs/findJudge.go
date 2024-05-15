package main

func FindJudge(n int, trust [][]int) int {
	trustCount := make([]int, n+1)

	for _, edge := range trust {
		trustCount[edge[0]]--
		trustCount[edge[1]]++
	}

	for i := 1; i <= n; i++ {
		if trustCount[i] == n-1 {
			return i
		}
	}

	return -1
}
