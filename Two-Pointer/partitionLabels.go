package main

import "fmt"

func partitionLabels(s string) []int {
    lastOccurence := make(map[byte]int)
    for i, ch := range s {
        lastOccurence[byte(ch)] = i;
    }

    partitions := []int{}
    start, end := 0, 0
    for i, ch := range s {
        if lastOccurence[byte(ch)] > end {
            end = lastOccurence[byte(ch)]
        }

        if i == end {
            partitions = append(partitions, end-start+1)
            start = i + 1;
        }
    }
    return partitions
}

func main(){
    s := "ababcbacadefegdehijhklij"
    fmt.Println(partitionLabels(s))
}
