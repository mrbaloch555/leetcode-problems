package main

import "fmt"

func isPowerOfFour(n int) bool {
	return helper(float64(n))
}

func helper(n float64) bool {
	if n < 4 {
		return n == 1
	}
	return helper(n / 4)
}
func main() {
	n := 256
	fmt.Println(isPowerOfFour(n))
}
