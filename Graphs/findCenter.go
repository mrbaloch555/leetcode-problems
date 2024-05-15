package main

// func FindCenter(edges [][]int) int {
// 	graph := func(edges [][]int) map[int][]int {
// 		graph := map[int][]int{}
// 		for _, val := range edges {
// 			graph[val[0]] = append(graph[val[0]], val[1])
// 			graph[val[1]] = append(graph[val[1]], val[0])
// 		}
// 		return graph
// 	}(edges)
// 	max := 0

// 	for key, val := range graph {
// 		if len(val) > len(graph[max]) {
// 			max = key
// 		}
// 	}
// 	return max
// }

func FindCenter(edges [][]int) int {
	mapper := make(map[int]struct{}, len(edges)*2)

	for _, edge := range edges {
		if _, ok := mapper[edge[0]]; ok {
			return edge[0]
		}
		if _, ok := mapper[edge[1]]; ok {
			return edge[1]
		}

		mapper[edge[0]] = struct{}{}
		mapper[edge[1]] = struct{}{}

	}

	return 0
}
