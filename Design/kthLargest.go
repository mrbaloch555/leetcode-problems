package main

import (
	"fmt"
	"sort"
)

type KthLargest struct {
	stream []int
	k      int
}

func Constructor(k int, nums []int) KthLargest {
	var stream KthLargest
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	stream.stream = nums
	stream.k = k
	return stream
}

func (this *KthLargest) Add(val int) int {
	func(nums []int) {
		push := false
		for i := 0; i < len(nums); i++ {
			if len(nums)-1 > i && val <= nums[i] && val >= nums[i+1] {
				cop := make([]int, len(this.stream[i+1:]))
				copy(cop, this.stream[i+1:])
				this.stream = this.stream[:i+1]
				this.stream = append(this.stream, val)
				this.stream = append(this.stream, cop...)
				push = true
				break
			} else if val >= nums[i] {
				this.stream = append([]int{val}, this.stream...)
				push = true
				break
			}
		}
		if !push {
			this.stream = append(this.stream, val)
		}
	}(this.stream)
	fmt.Println(this.stream)
	return this.stream[this.k-1]
}

func main() {
	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)
	fmt.Println(obj.Add(3), obj.Add(5), obj.Add(10), obj.Add(9), obj.Add(4))
}
