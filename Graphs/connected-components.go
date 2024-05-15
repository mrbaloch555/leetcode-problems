package main

func ConnectedComponents(graph map[int][]int) int {
	visited := map[int]bool{}
	count := 0
	for key, _ := range graph {
		if explore(graph, key, visited) {
			count++
		}

	}
	return count
}

func explore(graph map[int][]int, node int, visited map[int]bool) bool {
	if visited[node] {
		return false
	}
	visited[node] = true
	for _, val := range graph[node] {
		explore(graph, val, visited)
	}
	return true
}
