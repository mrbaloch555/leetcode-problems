package main

import "fmt"

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	result := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		result[i] = make([]int, n-2)
	}
	for j := 0; j < n-2; j += 1 {
		for k := 0; k < n-2; k += 1 {
			max := grid[j][k]

			for x := j; x < j+3; x++ {
				for y := k; y < k+3; y++ {
					if grid[x][y] > max {
						max = grid[x][y]
					}
				}
			}

			result[j][k] = max
		}
	}
	return result
}

func main() {
	grid := [][]int{
		{9, 9, 8, 1},
		{5, 6, 2, 6},
		{8, 2, 6, 4},
		{6, 2, 2, 2},
	}
	fmt.Println(largestLocal(grid))
}
