package main

import "fmt"

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		pings: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.pings = append(this.pings, t)
	diff := []int{t - 3000, t}
	count := 0
	for i := 0; i < len(this.pings); i++ {
		if this.pings[i] >= diff[0] && this.pings[i] <= diff[1] {
			count++
		}
	}

	return count
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))
	fmt.Println(obj.Ping(100))
	fmt.Println(obj.Ping(3001))
	fmt.Println(obj.Ping(3002))
}
