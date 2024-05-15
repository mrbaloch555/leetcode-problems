package main

import (
	"fmt"
	"strconv"
)
func isStrictlyPalindromic(n int) bool {
    for i := 2; i <= n-2; i++ {
        base := strconv.FormatInt(int64(n), i)
        if !isPalindrom(base) {
            return false
        }
    }
    return true
}

func isPalindrom(s string) bool {
    if len(s) <= 1 {
        return true
    }
    if s[0] == s[len(s)-1] {
        return isPalindrom(s[1 : len(s)-1])
    }
    return false
}

func main(){
    n := 60;
    fmt.Println(isStrictlyPalindromic(n))
}
