package main

func CanVisitAllRooms(graph [][]int) bool {
	visited := make([]bool, len(graph))

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	dfs(0)

	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}
