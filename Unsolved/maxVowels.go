package main

import (
	"fmt"
)

func main() {
	s, k := "aeiou", 2
	fmt.Println(maxVowels(s, k))
}

func maxVowels(s string, k int) int {
	return func(s string, k int) int {
		max, counter, previous := 0, k, 0
		for i := 0; i < len(s)-k+1; i++ {
			if i == 0 {
				previous += checkNumberOfVolwels(s[i:counter], true)
			} else {
				count := checkNumberOfVolwels(s[i:counter], false)
				if count > 0 {
					previous += count
					if previous > max {
						max = previous
					}
					previous--
				} else {
					previous--
				}
			}
			counter++
		}
		return max
	}(s, k)
}

func checkNumberOfVolwels(s string, check bool) int {
	counter := 0
	vowels := map[string]string{
		"a": "a",
		"e": "e",
		"i": "i",
		"o": "o",
		"u": "u",
	}
	if check {

		for _, val := range s {
			_, ok := vowels[string(val)]

			if ok {
				counter++
			}
		}
	} else {
		_, ok := vowels[string(s[len(s)-1])]
		if ok {
			counter++
		}
	}
	return counter
}
