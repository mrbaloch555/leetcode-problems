package main

import "fmt"

func main() {
	s := "){"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := []string{}
	if len(s)%2 != 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		val := string(s[i])
		if val == "(" || val == "{" || val == "[" {
			stack = append(stack, val)
		} else {
			if len(stack) != 0 {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if pop == "(" && val != ")" {
					return false
				} else if pop == "{" && val != "}" {
					return false
				} else if pop == "[" && val != "]" {
					return false
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
