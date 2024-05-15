package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var cpdomains = []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	fmt.Println(subdomainVisits(cpdomains))
}

func subdomainVisits(cpdomains []string) []string {
	visitied := make(map[string]string, 0)
	count_paired := make([]string, 0)
	for i, _ := range cpdomains {
		times := strings.Split(cpdomains[i], " ")[0]
		splits := strings.Split(cpdomains[i], ".")
		for j := len(splits) - 1; j >= 0; j-- {
			if j == 0 {
				visitied[splits[j]] = strings.Join(splits[j:], ".")
			} else {
				_, ok := visitied[splits[j]]
				if ok {
					value, _ := strconv.ParseInt(strings.Split(visitied[splits[j]], " ")[0], 10, 0)
					timeInt, _ := strconv.ParseInt(times, 10, 0)
					visitied[splits[j]] = strconv.Itoa(int(value)+int(timeInt)) + " " + strings.Join(splits[j:], ".")
				} else {
					visitied[splits[j]] = times + " " + strings.Join(splits[j:], ".")
				}
			}
		}
	}
	for _, val := range visitied {
		count_paired = append(count_paired, val)
	}
	return count_paired
}
