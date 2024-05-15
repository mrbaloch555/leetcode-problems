package main

import "fmt"

func main() {
	s := "Hannah"
	reverseString([]byte(s))
}

func reverseString(s []byte) {
	j := len(s) - 1
	for i := 0; i < len(s) && j >= i; i++ {
		temp := s[j]
		s[j] = s[i]
		s[i] = temp
		j--
	}
	fmt.Println(string(s))
}
