package main

import "fmt"

func countSubarrays(nums []int, minK int, maxK int) int64 {
    var ans int64
    j := -1
    prevMinKIndex := -1
    prevMaxKIndex := -1

    for i := 0; i < len(nums); i++ {
        if nums[i] < minK || nums[i] > maxK {
            j = i
        }
        if nums[i] == minK {
            prevMinKIndex = i
        }
        if nums[i] == maxK {
            prevMaxKIndex = i
        }
        ans += int64(max(0, min(prevMinKIndex, prevMaxKIndex)-j))
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func main(){
    // nums := []int{1,1,1,1}
    nums := []int{1,1,1,1}
    fmt.Println(countSubarrays(nums, 1, 1))
}
