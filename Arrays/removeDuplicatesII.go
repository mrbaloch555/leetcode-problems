// **********************************************************************************
// Problem Number: 80
// Problem Name: Remove Duplicates from Sorted Array II
// Problem Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	i := 0
// 	for j := 1; j < len(nums); j++ {
// 		if nums[i] != nums[j] {
// 			i++
// 			nums[i] = nums[j]
// 		}
// 		fmt.Println(nums)
// 	}

//		return i + 1
//	}
func removeDuplicatesII(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	i := 0
	for j := 2; j < len(nums); j++ {
		if nums[j] != nums[i] || nums[j] != nums[i+1] {
			i++
			nums[i+1] = nums[j]
		}
	}

	return i + 2
}

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	var people []Person
	people = append(people, Person{Id: 1, Name: "DK", Age: 20})
	people = append(people, Person{Id: 1, Name: "DK", Age: 31})
	people = append(people, Person{Id: 2, Name: "Sha", Age: 24})
	people = append(people, Person{Id: 2, Name: "Sha", Age: 21})
	people = append(people, Person{Id: 3, Name: "Khan", Age: 50})
	fmt.Println(removeDuplicates(people))
}

func removeDuplicates(people []Person) []Person {
	if len(people) == 0 {
		return []Person{}
	}

	i := 0
	for j := 1; j < len(people); j++ {
		if people[i].Id != people[j].Id {
			i++
			people[i] = people[j]
		}
	}

	return people[:i+1]
}
