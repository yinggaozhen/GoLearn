package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2,7,4,1,8,1}))
}

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
}

func _quickSort(nums []int) []int {
	var upper []int
	var lower []int
	var result []int

	if len(nums) <= 0 {
		return nums
	}

	sentry := nums[0]

	for num := range nums {
		upper
	}
}