package main

import "fmt"

func isSubsequence(s string, t string) bool {
	s_l := 0
	t_l := 0

	for t_l < len(t) && s_l < len(s) {
		if s[s_l] == t[t_l] {
			s_l += 1
		}
		t_l += 1
	}
	if s_l >= len(s) {
		return true
	}
	return false
}

func main() {
	s := "abx"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
