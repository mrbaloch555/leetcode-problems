// **********************************************************************************
// Problem Number: 125
// Problem Name: Valid Palindrome
// Problem Link: https://leetcode.com/problems/valid-palindrome/
// Status: Solved
// **********************************************************************************

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "recer"
	fmt.Println(isPalindrome(s))
}
func isPalindrome(s string) bool {
	s = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	reverse := ""
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse == s
}
