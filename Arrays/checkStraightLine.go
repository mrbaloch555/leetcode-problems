package main

import "fmt"

func main() {
	coordinates := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	fmt.Println(checkStraightLine(coordinates))
}

func checkStraightLine(coordinates [][]int) bool {

	for i := 0; i < len(coordinates); i++ {
		if len(coordinates[i:]) >= 2 {
			if coordinates[i+1][0]-coordinates[i][0] != 1 || coordinates[i+1][1]-coordinates[i][1] != 1 {
				return false
			}
		}
	}
	return true
}
