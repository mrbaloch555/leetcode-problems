package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left := 0
	right := len(people) - 1
	conut := 0
	rem := 0
	for left <= right {
		total := people[left] + people[right]
		if total > limit {
			right--
			rem++
		} else if total <= limit {
			conut++
			left++
			right--
		}
	}
	return conut + rem
}

func main() {
	people := []int{3, 5, 3, 4}
	limit := 4
	fmt.Println(numRescueBoats(people, limit))
}
