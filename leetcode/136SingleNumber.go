package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2,2,1}))
	fmt.Println(singleNumber([]int{4,1,2,1,2}))
}

func singleNumber(nums []int) int {
	result := make(map[int]int)

	for _, value := range nums {
		if _, ok := result[value]; !ok {
			result[value] = 1
		} else {
			result[value]++
		}
	}

	for key, value := range result {
		if value == 1 {
			return key
		}
	}

	return 0
}