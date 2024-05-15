package main

import "fmt"

func main() {
	nums := [][]int{{1, 3}, {4, 5, 6}}
	fmt.Println(searchMatrix(nums, 3))
}

func searchMatrix(matrix [][]int, target int) bool {

	low, high := 0, len(matrix)-1
	rowIndex := -1
	for low <= high {
		mid := low + (high-low)/2
		if target < matrix[mid][0] {
			high = mid - 1
		} else if len(matrix)-1 > mid && target > matrix[mid][0] && target > matrix[mid+1][0] {
			low = mid + 1
		} else {
			rowIndex = mid
		}
	}
	fmt.Println(rowIndex)
	// for i := 0; i < len(matrix); i++ {
	// 	low := 0
	// 	high := len(matrix[i]) - 1
	// 	for low <= high {
	// 		mid := low + (high-low)/2
	// 		if target < matrix[i][mid] {
	// 			high = mid - 1
	// 		} else if target > matrix[i][mid] {
	// 			low = mid + 1
	// 		} else {
	// 			return true
	// 		}
	// 	}
	// }
	return false
}
