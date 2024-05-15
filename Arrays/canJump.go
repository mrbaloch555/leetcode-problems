package main

import "fmt"

func main() {

	jums := []int{0, 2, 3}
	fmt.Println(canJump(jums))
}

func canJump(jumps []int) bool {
	n := len(jumps)
	tracker := jumps[0]

	for i := 0; i < n-1; i++ {
		if jumps[i] >= tracker {
			tracker = jumps[i]
		} else if tracker == 0 {
			tracker += jumps[i]
		}
		if tracker == 0 && jumps[i] == 0 {
			return false
		}
		tracker -= 1
	}
	return tracker >= 0
}
