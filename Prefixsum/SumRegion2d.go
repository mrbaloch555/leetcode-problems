package main

import "fmt"

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}

	rows, cols := len(matrix), len(matrix[0])
	sumMatrix := make([][]int, rows)

	for i := range sumMatrix {
		sumMatrix[i] = make([]int, cols)
	}

	sumMatrix[0][0] = matrix[0][0]
	for j := 1; j < cols; j++ {
		sumMatrix[0][j] = sumMatrix[0][j-1] + matrix[0][j]
	}

	for i := 1; i < rows; i++ {
		sumMatrix[i][0] = sumMatrix[i-1][0] + matrix[i][0]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			sumMatrix[i][j] = matrix[i][j] + sumMatrix[i-1][j] + sumMatrix[i][j-1] - sumMatrix[i-1][j-1]
		}
	}

	return NumMatrix{
		matrix: sumMatrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.matrix[row2][col2]
	} else if row1 == 0 {
		return this.matrix[row2][col2] - this.matrix[row2][col1-1]
	} else if col1 == 0 {
		return this.matrix[row2][col2] - this.matrix[row1-1][col2]
	} else {
		return this.matrix[row2][col2] - this.matrix[row2][col1-1] - this.matrix[row1-1][col2] + this.matrix[row1-1][col1-1]
	}
}

func main() {

	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	obj := Constructor(matrix)
	fmt.Println(obj)
}
