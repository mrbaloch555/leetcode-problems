package main

import "fmt"

func main(){
    height := []int{1, 1}
    fmt.Println(maxArea(height))
}

func maxArea(height []int) int {

    max1 := height[0]
    max2 := height[0]

    for i := 0; i < len(height); i++ {

        if height[i] > max1 {
            max1 = height[i]
        } else if height[i] > max2 && height[i] < max1 {
            max2 = height[i]
        }
    }
    return max2 * max2
}
