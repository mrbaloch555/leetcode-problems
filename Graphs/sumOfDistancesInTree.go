package main

import "fmt"

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := buildAdjacencyList(edges)

	count := make([]int, n)
	sums := make([]int, n)

	dfsCount(graph, 0, -1, count, sums)

	dfsSum(graph, 0, -1, count, sums)

	return sums
}

func buildAdjacencyList(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	return graph
}

func dfsCount(graph map[int][]int, src, parent int, count, sums []int) {
	count[src] = 1
	for _, neighbor := range graph[src] {
		if neighbor != parent {
			dfsCount(graph, neighbor, src, count, sums)
			count[src] += count[neighbor]
			sums[src] += sums[neighbor] + count[neighbor]
		}
	}
}

func dfsSum(graph map[int][]int, src, parent int, count, sums []int) {
	for _, neighbor := range graph[src] {
		if neighbor != parent {
			sums[neighbor] = sums[src] - count[neighbor] + len(count) - count[neighbor]
			dfsSum(graph, neighbor, src, count, sums)
		}
	}
}

func main() {
	n := 6
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}
	result := sumOfDistancesInTree(n, edges)
	fmt.Println(result) // Output: [8 12 6 10 10 10]
}
