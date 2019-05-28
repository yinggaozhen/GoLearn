package main

import (
	"fmt"
)

func main() {
	// fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4,5,5,5}))
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]

	for _, num := range nums[1:] {
		if sum < 0 {
			sum = num
		} else {
			sum += num
		}

		if max < sum {
			max = sum
		}
	}

	return max
}