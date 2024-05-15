package main

import "fmt"

func main() {
	haystack, needle := "leetcode", "ass"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		str := haystack[i : i+len(needle)]
		if str == needle {
			return i
		}
	}

	return -1
}
