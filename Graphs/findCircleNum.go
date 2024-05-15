package main

import "fmt"

func FindCircleNum(isConnected [][]int) int {
	count := 0
	n := len(isConnected)
	visited := make([]bool, n)
	for i := 0; i < len(visited); i++ {
		visited[i] = false
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			dfs(i, visited, isConnected)
		}
	}
	return count
}

func dfs(city int, visited []bool, isConnected [][]int) {
	visited[city] = true
	for i := 0; i < len(isConnected); i++ {
		fmt.Println(visited, city, i)
		if isConnected[city][i] == 1 && !visited[i] {
			dfs(i, visited, isConnected)
		}
	}
}
