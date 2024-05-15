package main

import "fmt"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := map[int][][]int{}
	for i := 0; i < len(flights); i++ {

		_, ok := graph[flights[i][0]]
		if ok {
			graph[flights[i][0]] = append(graph[flights[i][0]], []int{flights[i][1], flights[i][2]})
		} else {
			graph[flights[i][0]] = [][]int{{flights[i][1], flights[i][2]}}
		}
	}
	for _, val := range graph {
		fmt.Println("-----------------", graph)
		if len(val) > 0 {
			fmt.Println(helper(graph, flights, val[0], src, dst, k+1, 0))
		}
	}
	return n
}

func helper(graph map[int][][]int, flights [][]int, next []int, src int, dst int, k int, sum int) int {
	if k == 0 {
		return sum
	}
	if len(graph[next[0]]) == 0 {
		return sum
	}
	fmt.Println("NEXT =>", next)
	sum += next[1]
	_, okay := graph[next[0]]
	if okay {
		k -= 1
		next = graph[next[0]][0]
		graph[next[1]] = graph[next[1]][:]
		helper(graph, flights, next, src, dst, k, sum)
	}

	return sum
}
func main() {

	n := 4
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	src := 0
	dst := 3
	k := 1

	fmt.Println(findCheapestPrice(n, flights, src, dst, k))
}
