package main

import "fmt"

func validPath(n int, edges [][]int, start int, end int) bool {
	neighbors := make(map[int][]int)
	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		neighbors[n1] = append(neighbors[n1], n2)
		neighbors[n2] = append(neighbors[n2], n1)
	}

	seen := make(map[int]bool)
	return dfs(neighbors, start, end, seen)
}

func dfs(neighbors map[int][]int, node int, end int, seen map[int]bool) bool {
	if node == end {
		return true
	}
	if seen[node] {
		return false
	}

	seen[node] = true
	for _, n := range neighbors[node] {
		if dfs(neighbors, n, end, seen) {
			return true
		}
	}
	return false
}
func main() {
	edges := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	fmt.Println(validPath(len(edges), edges, 0, 5))
}
