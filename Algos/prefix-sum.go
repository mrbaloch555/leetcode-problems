package main

import "fmt"

type PrefixSum struct {
	nums []int
}

func NewPrefixSum(nums []int) PrefixSum {
	sums := []int{}
	sum := 0
	for _, val := range nums {
		sum += val
		sums = append(sums, sum)
	}
	return PrefixSum{
		nums: sums,
	}
}

func (p *PrefixSum) query(left, right int) int {
	if left == 0 {
		return p.nums[right]
	}
	return p.nums[right] - p.nums[left-1]
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	p := NewPrefixSum(nums)
	fmt.Println(p.query(2, 5))
}
