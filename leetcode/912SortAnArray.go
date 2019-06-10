package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(sortArray([]int{5,1,1,2,0,0}))

	var randomArray []int

	for i := 0; i < 10; i++ {
		randomArray = append(randomArray, rand.Intn(1000))
	}
	fmt.Println(sortArray(randomArray))
}

// 1. 快速排序
func sortArray(nums []int) []int {
	var lower []int
	var upper []int
	var result []int

	if len(nums) == 0 {
		return result
	}

	sentry := nums[0]

	for key, num := range nums {
		if key == 0 {
			continue
		}

		if num > sentry {
			upper = append(upper, num)
		} else {
			lower = append(lower, num)
		}
	}

	upper = sortArray(upper)
	lower = sortArray(lower)

	result = append(append(lower, sentry), upper...)

	return result
}

//func sortArray(nums []int) []int {
//	var t int
//
//	for i := 0; i < len(nums); i++ {
//		for j := 0; j < len(nums); j++ {
//			if nums[i] < nums[j] {
//				t = nums[i]
//				nums[i] = nums[j]
//				nums[j] = t
//			}
//		}
//	}
//
//	return nums
//}