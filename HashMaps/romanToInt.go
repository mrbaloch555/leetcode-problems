package main

import "fmt"

func main() {
	s := "LVIII"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	decimal := 0
	mapping := map[string][]int{
		"I": {0, 1},
		"V": {1, 5},
		"X": {2, 10},
		"L": {3, 50},
		"C": {4, 100},
		"D": {5, 500},
		"M": {6, 1000},
	}
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 {
			comb := s[i : i+2]
			p, _ := mapping[string(comb[0])]
			q, _ := mapping[string(comb[1])]

			if p[0] >= q[0] {
				decimal += p[1]
			} else {
				decimal += q[1] - p[1]
				i++
			}
		} else {
			decimal += mapping[string(s[i])][1]
		}
	}
	return decimal
}
