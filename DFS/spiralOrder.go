package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	var result []int
	var visited = make([][]bool, len(matrix))
	for i := range matrix {
		visited[i] = make([]bool, len(matrix[0]))
	}

	var dfs func(x, y, direction int)
	dfs = func(x, y, direction int) {
		if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) || visited[x][y] {
			return
		}

		result = append(result, matrix[x][y])
		visited[x][y] = true

		var nextX, nextY int
		switch direction {
		case 0:
			nextX, nextY = x, y+1
		case 1:
			nextX, nextY = x+1, y
		case 2:
			nextX, nextY = x, y-1
		case 3:
			nextX, nextY = x-1, y
		}

		if nextX < 0 || nextX >= len(matrix) || nextY < 0 || nextY >= len(matrix[0]) || visited[nextX][nextY] {
			direction = (direction + 1) % 4
			switch direction {
			case 0:
				nextX, nextY = x, y+1
			case 1:
				nextX, nextY = x+1, y
			case 2:
				nextX, nextY = x, y-1
			case 3:
				nextX, nextY = x-1, y
			}
		}

		dfs(nextX, nextY, direction)
	}

	dfs(0, 0, 0)

	return result
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix)) // Output: [1 2 3 6 9 8 7 4 5]
}
