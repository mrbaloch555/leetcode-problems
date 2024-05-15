package main

import "fmt"

type StockSpanner struct {
	prices []int
	spans  []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		prices: make([]int, 0),
		spans:  make([]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for len(this.prices) > 0 && this.prices[len(this.prices)-1] <= price {
		this.prices = this.prices[:len(this.prices)-1]
		span += this.spans[len(this.spans)-1]
		this.spans = this.spans[:len(this.spans)-1]
	}
	this.prices = append(this.prices, price)
	this.spans = append(this.spans, span)
	return span
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Next(100))
	fmt.Println(obj.Next(80))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(70))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(75))
	fmt.Println(obj.Next(85))
}
