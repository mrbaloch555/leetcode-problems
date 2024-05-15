package main

import "fmt"

func numSubarrayWithSum(nums []int, goal int) int {
    count := 0;
    sumMap := map[int]int{}
    prefixSum := 0;

    for _, num := range nums {
        prefixSum += num;
        if prefixSum == goal {
            count++
        }
        count += sumMap[prefixSum-goal]
        sumMap[prefixSum]++
    }
   return count;
}

func main(){
    nums := []int{1,0,1,0,1}
    fmt.Println(numSubarrayWithSum(nums, 2))
}
