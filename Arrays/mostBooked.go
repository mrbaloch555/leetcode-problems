package main

import (
	"fmt"
	"math"
)

func mostBooked(n int, meetings [][]int) int {
	minStartTime, maxEndTime := findMinMix(meetings)
	roomMap := map[int][]int{}

	for j := 0; j < n; j++ {
		room := meetings[j]
		roomMap[room[0]] = []int{j, room[1], 1}
	}

	for i := minStartTime; i < maxEndTime+1; i++ {
		for _, val := range roomMap {
			if val[1] <= i {
				if len(meetings) > 0 {
					lowestRoom := findLowestRoom(roomMap, i)
					roomMap[lowestRoom][1] = val[1] + i
					roomMap[lowestRoom][2]++
					meetings = meetings[1:]
				}
			}
		}
	}

	maxBookings, mostBookedRoom := 0, math.MaxInt32
	fmt.Println(roomMap)
	for k, v := range roomMap {
		if v[2] >= maxBookings || (v[2] == maxBookings && k < mostBookedRoom) {
			maxBookings = v[2]
			mostBookedRoom = k
		}
	}

	return mostBookedRoom
}

func findMinMix(meetings [][]int) (int, int) {
	min := math.MaxInt32
	max := 0
	for _, val := range meetings {
		if val[1] > max {
			max = val[1]
		}
		if val[0] < min {
			min = val[0]
		}
	}
	return min, max
}

func findLowestRoom(roomMap map[int][]int, t int) int {
	key := math.MaxInt32
	for k, val := range roomMap {
		if val[1] == t {
			if k < key {
				key = k
			}
		}
	}
	return key
}
func main() {
	meetings := [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}
	n := 2
	fmt.Println(mostBooked(n, meetings))
}
