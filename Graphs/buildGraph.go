package main

func BuildGraph(edges [][]int) map[int][]int {
	graph := map[int][]int{}
	for _, val := range edges {
		graph[val[0]] = append(graph[val[0]], val[1])
		graph[val[1]] = append(graph[val[1]], val[0])
	}
	return graph
}
