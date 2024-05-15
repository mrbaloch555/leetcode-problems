package main

import (
	"fmt"
	"math"
)

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

func main() {
	n := 10
	fmt.Println(bulbSwitch(n))
}
