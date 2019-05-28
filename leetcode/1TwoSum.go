package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	fmt.Print(result)
}

func twoSum(nums []int, target int) []int {
	var result []int

	for key1, value1 := range nums {
		for key2, value2 := range nums {
			if key1 == key2 {
				continue
			}

			if value1+value2 == target {
				result = append(result, key1)
				result = append(result, key2)

				return result
			}
		}
	}

	return result
}
