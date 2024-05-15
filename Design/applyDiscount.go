package main

import "fmt"

type Cashier struct {
	products   map[int]int
	discount   int
	minimum    int
	orderCount int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	return Cashier{
		products: func(products, prices []int) map[int]int {
			t := map[int]int{}
			for i, val := range products {
				t[val] = prices[i]
			}
			return t
		}(products, prices),
		discount: discount,
		minimum:  n,
	}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	this.orderCount++
	var total float64 = 0
	for i, val := range product {
		total += float64(amount[i] * this.products[val])
	}
	if this.orderCount == this.minimum {
		var dis float64 = float64(100-this.discount) / 100
		total = total * dis
		this.orderCount = 0
	}
	return float64(total)
}

func main() {
	obj := Constructor(3, 50, []int{1, 2, 3, 4, 5, 6, 7}, []int{100, 200, 300, 400, 300, 200, 100})
	fmt.Println(obj.GetBill([]int{1, 2}, []int{1, 2}))
	fmt.Println(obj.GetBill([]int{3, 7}, []int{10, 10}))
	fmt.Println(obj.GetBill([]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(obj.GetBill([]int{4}, []int{10}))
	fmt.Println(obj.GetBill([]int{7, 3}, []int{10, 10}))
	fmt.Println(obj.GetBill([]int{7, 5, 3, 1, 6, 4, 2}, []int{10, 10, 10, 9, 9, 9, 7}))
	fmt.Println(obj.GetBill([]int{2, 3, 5}, []int{5, 3, 2}))
}
