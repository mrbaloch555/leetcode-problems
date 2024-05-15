package main

func HasPath(graph map[int][]int, src, dst int, visited map[int]bool) bool {
	if src == dst {
		return true
	}
	if visited[src] {
		return false
	}
	visited[src] = true

	for _, neighbor := range graph[src] {
		if HasPath(graph, neighbor, dst, visited) {
			return true
		}
	}
	return false
}
