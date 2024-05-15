package main

import "fmt"

func pivotArray(nums []int, pivot int) []int {
    left := []int{}
    right := []int{}
    mid := []int{}

    for _, num := range nums {
        if num == pivot {
            mid = append(mid, num)
        } else if num < pivot {
            left = append(left, num)
        } else {
            right = append(right, num)
        }
    }
    return append(left, append(mid, right...)...)
}

func main(){
    nums := []int{9,12,5,10,14,3,10}
    fmt.Println(pivotArray(nums, 10))
}
