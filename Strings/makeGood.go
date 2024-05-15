package main

import (
    "fmt"
    "strings"
)

func makeGood(s string) string {
    var sb strings.Builder

    for _, c := range s {
        if sb.Len() > 0 && isBadPair(sb.String()[sb.Len()-1], byte(c)) {
            // Remove the last character by slicing the string
            sb.WriteString(s[:sb.Len()-1])
        } else {
            sb.WriteRune(c)
        }
    }

    return sb.String()
}

func isBadPair(a, b byte) bool {
    return a != b && strings.ToLower(string(a)) == strings.ToLower(string(b))
}

func main(){
    s := "abABaC"
    fmt.Println(makeGood(s));
}
