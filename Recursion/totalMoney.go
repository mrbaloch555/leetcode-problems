package main

import "fmt"

func totalMoney(n int) int {
	return helper(n, 0)
}

func helper(n int, count int) int {
	if n < 7 {
		sum := 0
		for i := count + 1; i <= n+count; i++ {
			sum += i
		}
		for i := 4; i < count+4; i++ {
			sum += 7 * i
		}
		return sum
	}
	count++
	return helper(n-7, count)
}
func main() {
	fmt.Println(totalMoney(30))
}
