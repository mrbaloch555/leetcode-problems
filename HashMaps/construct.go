package main

import "fmt"

type MyHashMap struct {
	myMap map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		myMap: make(map[int]int, 0),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.myMap[key] = value
}

func (this *MyHashMap) Get(key int) int {
	val, ok := this.myMap[key]
	if ok {
		return val
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	delete(this.myMap, key)
}

func main() {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	param_2 := obj.Get(1)
	obj.Put(2, 1)
	fmt.Println(param_2)
	obj.Remove(1)
}
