package main

import (
	"fmt"
	"sort"
)

type UnionFind struct {
	id   []int
	rank []int
}

func newUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		id:   make([]int, n),
		rank: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.rank[i] = 0
	}
	return uf
}

func (uf *UnionFind) unionByRank(u, v int) {
	i := uf.find(u)
	j := uf.find(v)
	if i == j {
		return
	}
	if uf.rank[i] < uf.rank[j] {
		uf.id[i] = j
	} else if uf.rank[i] > uf.rank[j] {
		uf.id[j] = i
	} else {
		uf.id[i] = j
		uf.rank[j] += 1
	}
}

func (uf *UnionFind) connected(u, v int) bool {
	return uf.find(u) == uf.find(v)
}

func (uf *UnionFind) reset(u int) {
	uf.id[u] = u
}

func (uf *UnionFind) find(u int) int {
	if uf.id[u] != u {
		uf.id[u] = uf.find(uf.id[u])
	}
	return uf.id[u]
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	uf := newUnionFind(n)
	timeToPairs := make(map[int][][2]int)

	uf.unionByRank(0, firstPerson)

	for _, meeting := range meetings {
		time := meeting[2]
		timeToPairs[time] = append(timeToPairs[time], [2]int{meeting[0], meeting[1]})
	}

	var sortedTimes []int
	for time := range timeToPairs {
		sortedTimes = append(sortedTimes, time)
	}
	sort.Ints(sortedTimes)

	for _, time := range sortedTimes {
		pairs := timeToPairs[time]
		peopleUnioned := make(map[int]bool)
		for _, pair := range pairs {
			x, y := pair[0], pair[1]
			uf.unionByRank(x, y)
			peopleUnioned[x] = true
			peopleUnioned[y] = true
		}
		for person := range peopleUnioned {
			if !uf.connected(person, 0) {
				uf.reset(person)
			}
		}
	}

	var result []int
	for i := 0; i < n; i++ {
		if uf.connected(i, 0) {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	n1 := 6
	meetings1 := [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}
	firstPerson1 := 1
	fmt.Println(findAllPeople(n1, meetings1, firstPerson1)) // Output: [0 1 2 3 5]

	n2 := 4
	meetings2 := [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}
	firstPerson2 := 3
	fmt.Println(findAllPeople(n2, meetings2, firstPerson2)) // Output: [0 1 3]

	n3 := 5
	meetings3 := [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}}
	firstPerson3 := 1
	fmt.Println(findAllPeople(n3, meetings3, firstPerson3)) // Output: [0 1 2 3 4]
}
