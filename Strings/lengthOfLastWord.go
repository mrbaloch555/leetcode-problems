package main

import (
	"fmt"
)

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	arr := make([]string, 0)
	func(s string) {
		word := ""
		for i := 0; i < len(s); i++ {
			if string(s[i]) == " " && word != "" {
				arr = append(arr, word)
				word = ""
			}
			if string(s[i]) != " " {
				word += string(s[i])
			}

			if i == len(s)-1 && word != "" {
				arr = append(arr, word)
			}
		}
	}(s)
	return len(arr[len(arr)-1])
}
