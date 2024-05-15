package main

import "fmt"

func findMaxLength(nums []int) int {
    numsMap := map[int]int{0: -1}
    maxLength := 0;
    count := 0;

    for i, num := range nums {
        if num == 0 {
            count++
        } else {
            count--
        }

        if index, ok := numsMap[count]; ok {
            if i - index > maxLength {
                maxLength = i - index;
            }
        } else {
            numsMap[count] = i;
        }
    }
    return maxLength;
}

func main(){
    nums := []int{0,1,0,0,1}
    fmt.Println(findMaxLength(nums))
}
