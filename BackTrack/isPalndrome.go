package main

import "fmt"

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func backtrack(s string, start int, currentPartition []string, result *[][]string) {
	if start >= len(s) {
		temp := make([]string, len(currentPartition))
		copy(temp, currentPartition)
		*result = append(*result, temp)
		return
	}

	for end := start; end < len(s); end++ {
		if isPalindrome(s, start, end) {
			currentPartition = append(currentPartition, s[start:end+1])
			backtrack(s, end+1, currentPartition, result)
			currentPartition = currentPartition[:len(currentPartition)-1]
		}
	}
}

func partition(s string) [][]string {
	var result [][]string
	backtrack(s, 0, []string{}, &result)
	return result
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
