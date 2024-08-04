package main

import (
	"fmt"
	"math"
)

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	matrix := make([][]int, len(rowSum))
	for i := range matrix {
		matrix[i] = make([]int, len(colSum)) // Correct the second dimension to len(colSum)
	}
	for i := range matrix {
		for j := range matrix[i] {
			min := math.Min(float64(rowSum[i]), float64(colSum[j]))
			rowSum[i] -= int(min)
			colSum[j] -= int(min)
			matrix[i][j] = int(min)
		}
	}
	return matrix
}

func main() {
	rowSum := []int{3, 8}
	colSum := []int{4, 7}
	fmt.Println(restoreMatrix(rowSum, colSum))
}
