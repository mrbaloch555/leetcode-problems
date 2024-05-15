package main

import "fmt"

type Path struct {
	source      string
	dest        string
	travelTime  float64
	customerId  int
	isCompleted bool
}

type UndergroundSystem struct {
	paths           map[string][]Path
	customerTracker map[int]string
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		paths:           map[string][]Path{},
		customerTracker: map[int]string{},
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	// check if source already exists in paths
	this.customerTracker[id] = stationName
	if _, ok := this.paths[stationName]; ok {
		this.paths[stationName] = append(this.paths[stationName], Path{
			customerId:  id,
			source:      stationName,
			dest:        "",
			travelTime:  float64(t),
			isCompleted: false,
		})
	} else {
		this.paths[stationName] = []Path{
			{
				source:      stationName,
				dest:        "",
				travelTime:  float64(t),
				customerId:  id,
				isCompleted: false,
			},
		}
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	origin := this.customerTracker[id]
	source := this.paths[origin]
	index := -1
	for i, val := range source {
		if val.customerId == id && !val.isCompleted {
			index = i
			break
		}
	}
	if index != -1 {
		source[index].dest = stationName
		source[index].travelTime = float64(t) - source[index].travelTime
		source[index].isCompleted = true
	}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	source := this.paths[startStation]
	var sum float64 = 0.0
	var count float64 = 0

	for _, val := range source {
		if val.dest == endStation && val.isCompleted {
			sum += val.travelTime
			count += 1
		}
	}

	return sum / count
}

func main() {
	obj := Constructor()
	obj.CheckIn(45, "Leyton", 3)
	obj.CheckIn(32, "Paradise", 8)
	obj.CheckIn(27, "Leyton", 10)
	obj.CheckOut(45, "Waterloo", 15)
	obj.CheckOut(27, "Waterloo", 20)
	obj.CheckOut(32, "Cambridge", 22)

	fmt.Println(obj.GetAverageTime("Paradise", "Cambridge"))
	fmt.Println(obj.GetAverageTime("Leyton", "Waterloo"))
	obj.CheckIn(10, "Leyton", 24)
	fmt.Println(obj.GetAverageTime("Leyton", "Waterloo"))
	obj.CheckOut(10, "Waterloo", 38)
	fmt.Println(obj.GetAverageTime("Leyton", "Waterloo"))
	obj.CheckIn(32, "Paradise", 8)
	// obj.CheckOut(32, "Cambridge", 22)

	fmt.Println(obj.paths)
}
