package main

func DFS(graph map[int][]int, node int, visited map[int]bool, values []int) []int {
	if visited[node] {
		return values
	}
	visited[node] = true
	values = append(values, node)
	for _, val := range graph[node] {
		values = DFS(graph, val, visited, values)
	}
	return values
}
