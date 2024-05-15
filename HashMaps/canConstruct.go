package main

import (
	"fmt"
)

func main() {
	ransomNote := "bg"
	magazine := "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	// for i := 0; i < len(magazine); i++ {
	// 	str := string(magazine[i])
	// 	index := strings.IndexAny(ransomNote, str)
	// 	if index != -1 {
	// 		ransomNote = ransomNote[:index] + ransomNote[index+1:]
	// 	}
	// }
	// return ransomNote == ""
	myMap := map[string]int{}

	for i := 0; i < len(ransomNote); i++ {
		str := string(ransomNote[i])
		_, ok := myMap[str]
		if ok {
			myMap[str] += 1
		} else {
			myMap[str] = 1
		}
	}
	for i := 0; i < len(magazine); i++ {
		str := string(magazine[i])
		_, ok := myMap[str]
		if ok {
			if myMap[str] != 0 {
				myMap[str] -= 1
			}
		}
	}
	for val, _ := range myMap {
		if myMap[val] != 0 {
			return false
		}
	}
	return true
}
