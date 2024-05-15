package main

import "fmt"

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	arr := this.nums[left : right+1]
	sum := 0
	for _, val := range arr {
		sum += val
	}

	return sum
}

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}
