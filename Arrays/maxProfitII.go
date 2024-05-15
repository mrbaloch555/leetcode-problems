package main

// WRONG VERSION

// import "fmt"

// func main() {

// 	prices := []int{7, 6, 4, 3, 1}
// 	fmt.Println(maxProfit(prices))
// }

// // func maxProfit(prices []int) int {

// // 	if len(prices) == 0 {
// // 		return 0
// // 	}

// // 	min_price := prices[0]
// // 	max_profit := 0
// // 	for i := 0; i < len(prices); i++ {
// // 		if prices[i] < min_price {
// // 			min_price = prices[i]
// // 		}
// // 		if len(prices[i:]) > 2 {
// // 			if prices[i+1] > prices[i+2] && prices[i+1] > min_price {
// // 				max_profit += prices[i+1] - min_price
// // 				min_price = prices[i+2]
// // 			}
// // 		} else if i == len(prices)-1 && min_price < prices[i] {
// // 			max_profit += prices[i] - min_price
// // 		}
// // 	}
// // 	return max_profit
// // }
