package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}

func maxProfit(prices []int) int {
	max := 0

	if len(prices) < 2 {
		return max
	}

	minPrice := prices[0]
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price - minPrice > max {
			max = price - minPrice
		}
	}

	return max
}