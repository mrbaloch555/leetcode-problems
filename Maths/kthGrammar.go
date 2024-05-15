package main

import (
	"fmt"
)

// func kthGrammar(n int, k int) int {
// 	res := []string{"0", "01"}
// 	for i := 2; i <= n; i++ {
// 		lastElem := res[len(res)-1]
// 		if i%2 == 1 {
// 			res = append(res, lastElem+reverseString(lastElem[:len(lastElem)/2])+reverseString(lastElem[len(lastElem)/2:]))
// 		} else {
// 			res = append(res, lastElem+reverseString(lastElem))
// 		}

//		}
//		result, _ := strconv.ParseInt(string(res[n-1][k-1]), 10, 32)
//		return int(result)
//	}
//
//	func reverseString(s string) string {
//		runes := []rune(s)
//		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//			runes[i], runes[j] = runes[j], runes[i]
//		}
//		return string(runes)
//	}

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	mid := 1 << uint(n-2)
	if k <= mid {
		return kthGrammar(n-1, k)
	}
	return 1 - kthGrammar(n-1, k-mid)
}

func main() {

	n := 30
	k := 434991989
	fmt.Println(kthGrammar(n, k))
}
