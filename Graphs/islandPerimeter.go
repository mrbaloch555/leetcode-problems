package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	var ans int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				bfs(grid, i, j)
				ans++
			}
		}
	}

	return ans
}

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func bfs(grid [][]byte, r, c int) {
	queue := [][]int{{r, c}}
	grid[r][c] = '2' // Mark '2' as visited.

	for len(queue) > 0 {
		i := queue[0][0]
		j := queue[0][1]
		queue = queue[1:]

		for _, dir := range dirs {
			x := i + dir[0]
			y := j + dir[1]
			if x < 0 || x == len(grid) || y < 0 || y == len(grid[0]) {
				continue
			}
			if grid[x][y] != '1' {
				continue
			}
			queue = append(queue, []int{x, y})
			grid[x][y] = '2' // Mark '2' as visited.
		}
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Println(numIslands(grid)) // Output: 3
}
