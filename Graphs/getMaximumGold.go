package main

import "fmt"

func getMaximumGold(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i, val := range grid {
		visited[i] = make([]bool, len(val))
	}

	var dfs func(x, y int) int

	dfs = func(x, y int) int {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 || visited[x][y] {
			return 0
		}
		visited[x][y] = true
		sum := grid[x][y]
		sum += max(dfs(x+1, y), dfs(x-1, y), dfs(x, y+1), dfs(x, y-1))
		visited[x][y] = false
		return sum
	}
	maxGold := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			maxGold = max(maxGold, dfs(i, j))
		}
	}
	return maxGold
}

func max(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	grid := [][]int{
		{0, 6, 0}, {5, 8, 7}, {0, 9, 0},
	}
	fmt.Println(getMaximumGold(grid))
}
