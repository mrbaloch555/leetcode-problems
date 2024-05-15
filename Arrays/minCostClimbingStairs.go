package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	pay := 0
	for i := 0; i < len(cost); i++ {
		if i == 0 {
			if cost[i+1] < cost[i] {
				i += 1
			}
		} else if len(cost[i:]) > 1 {
			if cost[i] <= cost[i+1] {
				pay += cost[i]
			}
		}
	}
	return pay
}

func main() {

	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}
