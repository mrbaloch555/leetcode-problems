package main

import "fmt"

func minimumDeletions(s string) int {
	n := len(s)
	bCount := 0
	deletions := 0

	for i := 0; i < n; i++ {
		if s[i] == 'b' {
			bCount++
		} else {
			if bCount > 0 {
				deletions++
				bCount--
			}
		}
	}

	return deletions
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "bbaaaaabb"
	fmt.Println(minimumDeletions(s))
}
