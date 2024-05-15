package main

func BFS(graph map[int][]int, node int) []int {
	visited := map[int]bool{}
	queue := []int{node}
	values := []int{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited[current] = true
		values = append(values, current)
		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
	return values
}
