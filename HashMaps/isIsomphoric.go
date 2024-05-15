package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if _, exists := sToT[s[i]]; exists {
			if sToT[s[i]] != t[i] {
				return false
			}
		} else {
			sToT[s[i]] = t[i]
		}

		if _, exists := tToS[t[i]]; exists {
			if tToS[t[i]] != s[i] {
				return false
			}
		} else {
			tToS[t[i]] = s[i]
		}
	}

	return true
}

func main() {
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))
}
