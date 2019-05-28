package main

import "fmt"

func main() {
	// fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4,5,5,5}))
	fmt.Println(searchInsert([]int{1,3,5,6}, 7))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}

	return len(nums)
}