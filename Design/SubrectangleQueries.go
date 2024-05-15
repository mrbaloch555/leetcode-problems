package main

import "fmt"

type SubrectangleQueries struct {
	matrix [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{
		matrix: rectangle,
	}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.matrix[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.matrix[row][col]
}

func main() {
	obj := Constructor([][]int{
		{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1},
	})
	fmt.Println(obj.GetValue(0, 2))
	obj.UpdateSubrectangle(0, 0, 3, 2, 5)
	fmt.Println(obj.GetValue(0, 2))
	fmt.Println(obj.GetValue(3, 1))
	obj.UpdateSubrectangle(3, 0, 3, 2, 10)
}
