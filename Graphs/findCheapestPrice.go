package main

import (
	"fmt"
	"math"
)

type Node struct {
	to   int
	cost int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := buildGraph(flights)
	visited := make(map[int]bool)
	memo := make(map[string]int)
	min := ex(graph, src, dst, visited, k+1, n, memo)
	return min
}

func ex(graph map[int][]Node, src, dst int, visited map[int]bool, k, n int, memo map[string]int) int {
	if src == dst {
		return 0
	}

	if k == 0 {
	}

	visited[src] = true
	memoKey := fmt.Sprintf("%d-%d", src, k)

	if val, ok := memo[memoKey]; ok {
		visited[src] = false
		return val
	}

	minCost := math.MaxInt32

	for _, neighbor := range graph[src] {
		if !visited[neighbor.to] {
			cost := ex(graph, neighbor.to, dst, visited, k-1, n, memo)
			fmt.Println(cost)
			if cost != -1 {
				minCost = min(minCost, neighbor.cost+cost)
			}
		}
	}

	visited[src] = false

	if minCost == math.MaxInt32 {
		memo[memoKey] = -1
		return -1
	}

	memo[memoKey] = minCost
	return minCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func buildGraph(flights [][]int) map[int][]Node {
	graph := make(map[int][]Node)
	for _, flight := range flights {
		src, dst, cost := flight[0], flight[1], flight[2]
		graph[src] = append(graph[src], Node{to: dst, cost: cost})
	}
	return graph
}

// func main() {
// 	n := 5
// 	flights := [][]int{{0, 1, 100}, {0, 2, 100}, {0, 3, 10}, {1, 2, 100}, {1, 4, 10}, {2, 1, 10}, {2, 3, 100}, {2, 4, 100}, {3, 2, 10}, {3, 4, 100}}
// 	src := 0
// 	dst := 4
// 	k := 3
// 	fmt.Println(findCheapestPrice(n, flights, src, dst, k))
// }
