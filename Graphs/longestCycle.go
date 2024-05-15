package main

import "fmt"

// import "fmt"

// func main() {
// 	edges := []int{3, 3, 4, 2, 3}

// 	graph := map[int][]int{}
// 	for i, val := range edges {
// 		graph[i] = append(graph[i], val)
// 	}

// 	fmt.Println(edges)
// }

func LongestCycle(graph map[int][]int) int {

	visited := map[int]bool{}
	fmt.Println(graph)
	for key, _ := range graph {
		bfs(graph, key, visited)
	}
	fmt.Println(visited)
	return -1
}

func bfs(graph map[int][]int, node int, visited map[int]bool) bool {
	// fmt.Println(node)
	if visited[node] {
		return true
	}
	visited[node] = true
	for _, val := range graph[node] {
		if bfs(graph, val, visited) {
			fmt.Println(node)
		}
	}

	return false
}
