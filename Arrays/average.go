// **********************************************************************************
// Problem Number: 1491
// Problem Name: Average Salary Excluding the Minimum and Maximum Salary
// Problem Link: https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

func main() {
	salary := []int{48000, 59000, 99000, 13000, 78000, 45000, 31000, 17000, 39000, 37000, 93000, 77000, 33000, 28000, 4000, 54000, 67000, 6000, 1000, 11000}
	fmt.Println(average(salary))
}

func average(salary []int) float64 {
	return func() float64 {
		minSalaryIndex := 0
		maxSalaryIndex := 0

		for i, _ := range salary {
			if salary[minSalaryIndex] > salary[i] {
				minSalaryIndex = i
			}
			if salary[maxSalaryIndex] < salary[i] {
				maxSalaryIndex = i
			}
		}

		sum := 0
		for i, _ := range salary {
			if i != minSalaryIndex && i != maxSalaryIndex {
				sum += salary[i]
			}
		}
		sumInt := (sum * 100000) / (len(salary) - 2)
		return float64(sumInt) / 100000
	}()
}
