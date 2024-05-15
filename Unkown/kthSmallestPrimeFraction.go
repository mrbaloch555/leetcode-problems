package main

import (
	"fmt"
	"sort"
)

func kthSmallestPrimeFraction(arr []int, k int) []int {
	fractionsArray := [][]float64{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			frac := float64(arr[i]) / float64(arr[j])
			fractionsArray = append(fractionsArray, []float64{float64(arr[i]), float64(arr[j]), float64(frac)})
		}
	}
	sort.Slice(fractionsArray, func(i, j int) bool {
		return fractionsArray[i][2] < fractionsArray[j][2]
	})
	res := []int{}
	for _, val := range fractionsArray[k-1][:2] {
		res = append(res, int(val))
	}
	return res
}

func main() {
	arr := []int{1, 2, 3, 5}
	k := 3
	fmt.Println(kthSmallestPrimeFraction(arr, k))
}
