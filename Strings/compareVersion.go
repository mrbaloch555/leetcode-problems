package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	levels1 := strings.Split(version1, ".")
	levels2 := strings.Split(version2, ".")
	length := max(len(levels1), len(levels2))

	for i := 0; i < length; i++ {
		v1 := 0
		v2 := 0
		if i < len(levels1) {
			v1, _ = strconv.Atoi(levels1[i])
		}
		if i < len(levels2) {
			v2, _ = strconv.Atoi(levels2[i])
		}
		if v1 < v2 {
			return -1
		}
		if v1 > v2 {
			return 1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	ver1 := "1.1"
	ver2 := "1.10"
	fmt.Println(compareVersion(ver1, ver2))
}
