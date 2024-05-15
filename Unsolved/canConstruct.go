package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "annabelle"
	k := 2

	fmt.Println(canConstruct(s, k))
}

func canConstruct(s string, k int) bool {

	for i := 0; i < len(s); i++ {
		comb := []string{}
		for j := i; j < len(s); j++ {
			comb = append(comb, s[i:j+1], "")
		}
		fmt.Println(comb)
	}
	return true
}
