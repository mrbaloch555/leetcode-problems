package main

import "fmt"

func main() {
	nums := []int{1, 100}
	fmt.Println(distinctAverages(nums))
}
func distinctAverages(nums []int) int {
	var average []float64
	for i := 0; i < len(nums); i++ {
		maxIndex := findMax(nums)
		maxValue := nums[maxIndex]
		nums[maxIndex] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		minIndex := findMin(nums)
		minValue := nums[minIndex]
		nums[minIndex] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		average = append(average, float64((float64(minValue+maxValue))/2))
		fmt.Println(average)
		i -= 1
	}
	return len(removeDuplicateInt(average))
}

func findMax(nums []int) int {
	type maximum struct {
		val   int
		index int
	}

	var max maximum

	max.val = nums[0]
	max.index = 0
	for i := 0; i < len(nums); i++ {
		if max.val < nums[i] {
			max.val = nums[i]
			max.index = i
		}
	}
	return max.index
}

func findMin(nums []int) int {
	type minimum struct {
		val   int
		index int
	}
	var min minimum
	min.val = nums[0]
	min.index = 0
	for i := 0; i < len(nums); i++ {
		if min.val > nums[i] {
			min.val = nums[i]
			min.index = i
		}
	}
	return min.index
}

func removeDuplicateInt(intSlice []float64) []float64 {
	allKeys := make(map[float64]bool)
	list := []float64{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
