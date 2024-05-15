package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
    res := []int{}
    nums1Map := map[int]int{}
    nums2Map := map[int]int{}
    for _, val := range nums1 {
        nums1Map[val] = val;
    }
    for _, val := range nums2 {
        nums2Map[val] = val;
    }

    for _, val := range nums1Map {
        if _, ok := nums2Map[val]; ok {
            res = append(res, val)
        }
    }

    return res
}

func main(){
    nums1 := []int{4,9,5}
    nums2 := []int{9,2,9,8,4}
    fmt.Println(intersection(nums1,nums2))
}
