package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func trebuchet(lines string) int {
	// numEn := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	res := 0
	stars := []string{}
	for _, line := range strings.Split(lines, "\n") {
		foundedDigits := []string{}
		for i, val := range line {
			_, err := strconv.ParseInt(string(val), 10, 32)
			if err == nil {
				foundedDigits = append(foundedDigits, string(val))
			} else {
				if len(line[i:]) >= 5 {
					if slices.Contains(numbers, line[i:i+3]) {
					}
					// for j := 0; j <
					// for j := 0; j < len(str); j++ {
					// 	key, ok := numEn[str[j:]]
					// 	fmt.Println(str[j:])
					// 	if ok {
					// 		fmt.Println(key)
					// 		numStr := strconv.Itoa(key)
					// 		foundedDigits = append(foundedDigits, numStr)
					// 		// i += key
					// 	}
					// }

				}
			}
		}
		if len(foundedDigits) >= 1 {
			if len(foundedDigits) == 1 {
				stars = append(stars, foundedDigits[0]+foundedDigits[0])
			} else {
				stars = append(stars, foundedDigits[0]+foundedDigits[len(foundedDigits)-1])
			}
		}
	}
	fmt.Println(stars)
	for _, val := range stars {
		digit, _ := strconv.ParseInt(string(val), 10, 32)
		res += int(digit)
	}
	return res
}

func main() {

	data, _ := os.ReadFile("day1.txt")
	fmt.Println(trebuchet(string(data)))

}
