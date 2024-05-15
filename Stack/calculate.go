package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculate(s string) int {
	stack := []int{}
	num := 0
	sign := 1

	s = strings.ReplaceAll(s, " ", "")
	for i := 0; i < len(s); i++ {
		char := string(s[i])

		switch char {
		case " ":
		case "+":
			sign = 1
		case "-":
			sign = -1
		case "(":
			stack = append(stack, num, sign)
			num, sign = 0, 1
		case ")":
			prevSign := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prevNum := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num = prevNum + prevSign*num
		default:
			for len(s[i:]) > 1 && !strings.Contains("+-()", string(s[i+1])) {
				char += string(s[i+1])
				i++
			}
			val, _ := strconv.Atoi(char)
			num += sign * val
		}
	}

	return num
}

func main() {
	// Example usage:
	result := calculate("1 + 1")
	fmt.Println(result) // Output: 0
}
