package main

func AllPathsSourceTarget(graph [][]int) [][]int {
	result := [][]int{}
	path := []int{}

	allPathsExplore(graph, 0, &result, path)
	return result
}

func allPathsExplore(graph [][]int, node int, result *[][]int, path []int) {
	path = append(path, node)
	if node == len(graph)-1 {
		*result = append(*result, append([]int{}, path...))
	} else {
		for _, val := range graph[node] {
			allPathsExplore(graph, val, result, path)
		}
	}
	path = path[:len(path)-1]
}
