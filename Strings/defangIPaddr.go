// **********************************************************************************
// Problem Number: 1108
// Problem Name: Defanging an IP Address
// Problem Link: https://leetcode.com/problems/defanging-an-ip-address/
// Status: Solved
// **********************************************************************************

package main

import (
	"fmt"
)

func main() {

	ip := "255.100.50.0"
	fmt.Println(defangIPaddr(ip))
}

func defangIPaddr(address string) string {
	// runes := []rune(address)
	var res string
	for i := 0; i < len(address); i++ {
		if string(address[i]) != "." {
			res = res + string(address[i])
		} else {
			res = res + "[.]"
		}
	}

	return res
}
