package main

import (
	"fmt"
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int {
    counter := 0;
    maximum := 0;
    sort.Slice(happiness, func(i, j int) bool {
        return happiness[i] > happiness[j]
    })

    for i := 0; i < k; i++ {
        if happiness[i] - counter > 0 {
            maximum += happiness[i] - counter
        }
        counter++;
    }
    return maximum
}

func main(){
    happiness := []int{2,3,4,5}
    k := 1
    fmt.Print(maximumHappinessSum(happiness, k))
}
