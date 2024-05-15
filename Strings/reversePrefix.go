package main

import (
	"fmt"
	"strings"
)

func reversePrefix(word string, ch byte) string {
	index := strings.IndexByte(word, ch)
	if index == -1 {
		return word
	}
	return reverse(word[:index+1]) + word[index+1:]
}

func reverse(str string) string {
	if str == "" {
		return str
	}
	return reverse(str[1:]) + string(str[0])
}

func main() {
	word := "abcdefd"
	ch := byte('d')
	fmt.Println(reversePrefix(word, ch))
}
