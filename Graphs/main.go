package main

// func main() {
// visited := map[int]bool{}
// graph := map[int][]int{
// 	1: []int{2, 3},
// 	2: []int{1, 4, 5},
// 	3: []int{1, 6},
// 	4: []int{2},
// 	5: []int{2},
// 	6: []int{3},
// 	7: []int{8},
// 	8: []int{7},
// }
// fmt.Println(BFS(graph, 1))
// fmt.Println(DFS(graph, 1, visited, []int{}))
// fmt.Println(ConnectedComponents(graph))
// fmt.Println(HasPath(graph, 1, 6, map[int]bool{}))
// edges := []int{3, 3, 4, 2, 3}
// graph := map[int][]int{}
// 	graph := [][]int{{1, 2}, {3}, {3}, {}}

// 	// for i, val := range edges {
// 	// 	graph[i] = []int{val}
// 	// }
// 	fmt.Println(AllPathsSourceTarget(graph))
// }

// func buildGraph(edges [][]int) map[int][]int {
// 	graph := map[int][]int{}

// 	for _, val := range edges {
// 		graph[val[0]] = append(graph[val[0]], val[1])
// 		graph[val[1]] = append(graph[val[1]], val[0])
// 	}
// 	return graph
// }
