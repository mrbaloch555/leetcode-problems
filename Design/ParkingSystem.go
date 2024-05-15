package main

import "fmt"

type ParkingSystem struct {
	allocation map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		allocation: map[int]int{
			1: big,
			2: medium,
			3: small,
		},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {

	val, _ := this.allocation[carType]
	if val == 0 {
		return false
	}
	this.allocation[carType] -= 1
	return true
}

func main() {
	obj := Constructor(1, 1, 0)
	fmt.Println(obj.AddCar(1))
	fmt.Println(obj.AddCar(2))
	fmt.Println(obj.AddCar(3))
	fmt.Println(obj.AddCar(1))
}
