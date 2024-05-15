package main

import "fmt"

func main() {
	strs := []string{"", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	prefixer := strs[0]
	longestPrefix := ""
	for i := 0; i < len(prefixer); i++ {
		check := true
		for j := 0; j < len(strs); j++ {
			if len(strs[j]) > i {
				if string(strs[j][i]) != string(prefixer[i]) {
					check = false
					break
				}
			} else {
				check = false
			}
		}
		if !check {
			return longestPrefix
		} else {
			longestPrefix += string(prefixer[i])
		}
	}
	return longestPrefix
}
