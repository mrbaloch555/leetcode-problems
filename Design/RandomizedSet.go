package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	dict  map[int]int
	items []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		dict:  make(map[int]int),
		items: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.dict[val]
	if ok {
		return false
	}

	this.items = append(this.items, val)
	this.dict[val] = len(this.items) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.dict[val]
	if !ok {
		return false
	}

	lastIdx := len(this.items) - 1
	lastVal := this.items[lastIdx]

	// Swap the value to be removed with the last value
	this.items[idx] = lastVal
	this.dict[lastVal] = idx

	// Remove the last element
	this.items = this.items[:lastIdx]

	// Delete the value from the dictionary
	delete(this.dict, val)

	return true
}
func (this *RandomizedSet) GetRandom() int {
	if len(this.items) == 0 {
		return -1
	}

	randomIdx := rand.Intn(len(this.items))
	return this.items[randomIdx]
}
func main() {
	obj := Constructor()
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Remove(2))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.GetRandom())
}
