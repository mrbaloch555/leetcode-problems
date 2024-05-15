package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {

	res := [][]string{}
	obj := map[string][]string{}

	for i := 0; i < len(strs); i++ {
		reverse := func(str string) string {
			newStr := []byte(str)
			sort.Slice(newStr, func(k, l int) bool {
				return newStr[k] < newStr[l]
			})
			return string(newStr)
		}(strs[i])
		_, ok := obj[reverse]
		if ok {
			obj[reverse] = append(obj[reverse], strs[i])
		} else {
			obj[reverse] = []string{strs[i]}
		}
	}
	for _, val := range obj {
		res = append(res, val)
	}
	return res
}

func main() {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
