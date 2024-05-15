// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {

// 	arr := [][]int{
// 		{
// 			1, 2, 3, 5,
// 		},
// 		{4, 5, 6},
// 	}
// 	fmt.Println(flat(arr))
// }

// func flat(arr [][]interface) interface{} {
// 	flatArray := make([]int, 0)

// 	for i, _ := range arr {
// 		check := reflect.TypeOf(arr[i]).Kind()
// 		if check.String() == "slice" {
// 			flat(arr[i])
// 		}
// 	}

// 	return arr
// }

// func makeFlat([]int)