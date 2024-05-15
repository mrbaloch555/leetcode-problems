package main

import "fmt"

func countBattleships(board [][]byte) int {

	var dfs func([][]byte, int, int)
	count := 0
	dfs = func(board [][]byte, x int, y int) {
		if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] == '.' {
			return
		}
		board[x][y] = '.'
		dfs(board, x+1, y)
		dfs(board, x-1, y)
		dfs(board, x, y+1)
		dfs(board, x, y-1)

	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				count++
				dfs(board, i, j)
			}
		}
	}
	return count
}

func main() {
	board := [][]byte{
		{'X', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}
	fmt.Println(countBattleships(board))
}
