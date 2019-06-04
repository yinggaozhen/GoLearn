package main

import "fmt"

func main() {
	fmt.Println(maxProfit2([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit2([]int{1,2,3,4,5}))
}

func maxProfit2(prices []int) int {
	profit := 0

	if len(prices) == 0 {
		return profit
	}

	prefix := prices[0]
	for _, price := range prices[1:] {
		if price > prefix {
			profit += price - prefix
		}

		prefix = price
	}

	return profit
}