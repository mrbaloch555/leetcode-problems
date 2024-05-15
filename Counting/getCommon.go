package main

import (
    "math"
    "fmt"
)

func getCommon(nums1 []int, nums2 []int) int {

    nums1Map := map[int]int{}
    nums2Map := map[int]int{}

    for _, val := range nums1 {
        nums1Map[val]++
    }
    for _, val := range nums2 {
        nums2Map[val]++
    }

    fmt.Println(nums1Map, nums2Map)
    min := math.MaxInt32
    for key, _ := range nums1Map {
        if _, ok := nums2Map[key]; ok {
            if key < min {
                min = key
            }
        }
    }
    fmt.Println(min)
    if min == math.MaxInt32 {
        return -1;
    }
    return min
 }

 func main(){
     nums1 := []int{1,2,3,6}
     nums2 := []int{2,3,4,5}
     fmt.Println(getCommon(nums1, nums2))
 }
