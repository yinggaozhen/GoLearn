package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, S int) int {
	return _findTargetSumWays(nums, S, 0)
}

func _findTargetSumWays(nums []int, S int, step int) int {
	if step == len(nums) {
		if sumArray(nums) == S {
			return 1
		}

		return 0
	}

	count1 := _findTargetSumWays(nums, S, step + 1)

	nums[step] *= -1
	count2 := _findTargetSumWays(nums, S, step + 1)

	return count1 + count2
}

func sumArray(nums []int) int {
	result := 0

	for _, value := range nums {
		result += value
	}

	return result
}