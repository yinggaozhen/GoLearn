package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(0))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	sum := 0
	first := 1
	next := 2

	for i := 2; i < n; i++ {
		sum = first + next
		first = next
		next = sum
	}

	return sum
}