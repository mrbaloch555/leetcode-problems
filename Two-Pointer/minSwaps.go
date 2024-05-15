package main

import (
	"fmt"
	"math"
)

func minSwaps(s string) int {
     unmatched := 0

    for _, c := range s {
        if c == '[' {
            unmatched++
        } else if unmatched > 0 {
            unmatched--
        }
    }

    return int(math.Ceil(float64(unmatched) / 2))
}

func main(){
    s := "]]][[[";
    fmt.Println(minSwaps(s))
}
