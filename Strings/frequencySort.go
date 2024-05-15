package main

import (
	"fmt"
	"sort"
	"strings"
)

func frequencySort(s string) string {
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	type SortSequence struct {
		key rune
		val int
	}

	var sortSeq []SortSequence
	for key, val := range freq {
		sortSeq = append(sortSeq, SortSequence{key: key, val: val})
	}

	sort.Slice(sortSeq, func(i, j int) bool {
		return sortSeq[i].val > sortSeq[j].val
	})

	var builder strings.Builder
	for _, entry := range sortSeq {
		builder.WriteString(strings.Repeat(string(entry.key), entry.val))
	}

	return builder.String()
}

func main() {
	s := "cccaaa"
	fmt.Println(frequencySort(s))
}
