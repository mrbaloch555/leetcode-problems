package main

import (
	"fmt"
	"math"
)

func scoreOfString(s string) int {
	sum := 0
	for i, _ := range s {
		if len(s[i:]) > 1 {
			sum += int(math.Abs(float64(s[i]) - float64(s[i+1])))
		}
	}
	return sum
}

func main() {
	s := "zaz"
	fmt.Println(scoreOfString(s))
}
