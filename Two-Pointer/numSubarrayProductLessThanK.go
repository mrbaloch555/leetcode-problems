package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
    n := len(nums) - 1;
    count := 0;
    for i := 0; i <= n; i++ {
        prod := 1;
        l := i;
        for prod < k && l <= n {
            prod *= nums[l]
            if prod < k {
                count++
            }
            l++;
        }
    }
    return count;
}

func main(){
    nums := []int{1,2,3}
    k := 0;
    fmt.Println(numSubarrayProductLessThanK(nums, k))
}
