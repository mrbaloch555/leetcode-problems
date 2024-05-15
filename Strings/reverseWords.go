package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	str := ""
	array := strings.Split(s, " ")
	for i := 0; i < len(array); i++ {
		str += reverse(array[i])
		if i != len(array)-1 {
			str += " "
		}
	}
	return str
}

func reverse(s string) string {
	str := ""
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}
