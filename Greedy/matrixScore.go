package main

import "fmt"

func matrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			for j := 0; j < n; j++ {
				grid[i][j] ^= 1
			}
		}
	}

	for j := 1; j < n; j++ {
		countOnes := 0
		for i := 0; i < m; i++ {
			countOnes += grid[i][j]
		}
		if countOnes < m-countOnes {
			for i := 0; i < m; i++ {
				grid[i][j] ^= 1
			}
		}
	}

	score := 0
	for _, row := range grid {
		num := 0
		for _, bit := range row {
			num = num<<1 | bit
		}
		score += num
	}

	return score
}

func main() {
	grid := [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}
	fmt.Println(matrixScore(grid)) // Output: 39
}
