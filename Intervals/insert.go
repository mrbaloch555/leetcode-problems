package main

import "fmt"
func insert(intervals [][]int, newInterval []int) [][]int {
    res := [][]int{}
    inserted := false

    for _, interval := range intervals {
        if inserted || newInterval[1] < interval[0] {
            if !inserted {
                res = append(res, newInterval)
                inserted = true
            }
            res = append(res, interval)
        } else if newInterval[0] > interval[1] {
            res = append(res, interval)
        } else {
            newInterval[0] = min(newInterval[0], interval[0])
            newInterval[1] = max(newInterval[1], interval[1])
        }
    }

    if !inserted {
        res = append(res, newInterval)
    }

    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func main(){
    nums := [][]int{{1,3}, {6, 9}}
    newInterval := []int{2,5}
    fmt.Println(insert(nums, newInterval))
}
