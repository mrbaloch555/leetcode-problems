package main

import "fmt"

func sortColors(nums []int) {
	redIndex := -1
	whiteIndex := -1
	blueIndex := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			// red
			nums[redIndex+1] = nums[i]
			redIndex += 1
		} else if nums[i] == 1 {
			// white
			nums[whiteIndex+redIndex+1] = nums[i]
			whiteIndex += redIndex + 1
		} else {
			// blue
			nums[blueIndex+whiteIndex+redIndex+1] = nums[i]
			blueIndex += redIndex + whiteIndex + 1
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
