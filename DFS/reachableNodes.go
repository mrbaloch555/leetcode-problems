package main

import "fmt"

func reachableNodes(n int, edges [][]int, restricted []int) int {
	resitrictedMap := map[int]int{}
	for _, val := range restricted {
		resitrictedMap[val] = val
	}
	return dfs(0, buildList(edges), map[int]bool{}, resitrictedMap)
}

func buildList(edges [][]int) map[int][]int {
	list := map[int][]int{}

	for _, edge := range edges {
		list[edge[0]] = append(list[edge[0]], edge[1])
		list[edge[1]] = append(list[edge[1]], edge[0])
	}
	return list
}

func dfs(src int, list map[int][]int, visited map[int]bool, restricted map[int]int) int {
	if visited[src] {
		return 0
	}
	if _, ok := restricted[src]; ok {
		return 0
	}
	visited[src] = true
	count := 1

	for _, val := range list[src] {
		if !visited[val] {
			count += dfs(val, list, visited, restricted)
		}
	}
	return count
}

func main() {
	n := 7
	edges := [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}
	restricted := []int{4, 5}
	fmt.Println(reachableNodes(n, edges, restricted))
}
