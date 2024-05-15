package main

import "fmt"

func main() {
	grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
	fmt.Println(equalPairs(grid))
}

func equalPairs(grid [][]int) int {

	for i := 0; i < len(grid); i++ {
		// fmt.Println(grid[i])
		col := []int{}
		for j := 0; j < len(grid[i]); j++ {
			col = append(col, grid[j][i])
		}
		fmt.Println(col, grid[i])
	}
	return 0
}
