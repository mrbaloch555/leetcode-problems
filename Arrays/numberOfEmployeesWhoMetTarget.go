package main

import "fmt"
func main(){

    hours := []int{0,1,2,3,4}
    target := 2
    fmt.Println(numberOfEmployeesWhoMetTarget(hours, target))
    
}

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
 
    res := 0

    for i := 0; i < len(hours); i++ {
        if(hours[i] >= target){
            res++
        }
    }
    return res
}
