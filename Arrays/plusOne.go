package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	digits := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	strs := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(digits)), ""), "[]")
	fmt.Println(strs)
	parsedInt, _ := strconv.ParseInt(strs, 10, 64)
	parsedInt = parsedInt + 1
	fmt.Println(parsedInt)
	re_string := strings.Split(strconv.Itoa(int(parsedInt)), "")

	return func(str []string) []int {
		digits = []int{}
		for _, val := range re_string {
			newInt, _ := strconv.ParseInt(val, 10, 0)
			digits = append(digits, int(newInt))
		}
		return digits
	}(re_string)
}
