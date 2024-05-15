package main

import "fmt"

func maxSubarrayLength(nums []int, k int) int {
    ans := 0
    count := make(map[int]int)

    for l, r := 0, 0; r < len(nums); r++ {
        count[nums[r]]++
        for count[nums[r]] == k+1 {
            count[nums[l]]--
            l++
        }
        ans = max(ans, r-l+1)
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main(){
    nums := []int{1,11}
    k := 1;

    fmt.Println(maxSubarrayLength(nums, k))
}
