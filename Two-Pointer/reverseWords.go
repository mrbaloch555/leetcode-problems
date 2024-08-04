package main

import (
	"fmt"
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	var reverse func(str []string) string
	re := regexp.MustCompile(`\s+`)

	reverse = func(str []string) string {
		if len(str) == 0 {
			return ""
		}
		reversed := reverse(str[1:])
		if reversed == "" {
			return strings.TrimSpace(re.ReplaceAllString(str[0], ""))
		}
		return reversed + " " + strings.TrimSpace(re.ReplaceAllString(str[0], ""))
	}
	trimmed := strings.TrimSpace(re.ReplaceAllString(s, " "))
	words := strings.Split(trimmed, " ")
	return reverse(words)
}
func main() {
	s := "  hello world  "
	fmt.Println(reverseWords(s))
}
