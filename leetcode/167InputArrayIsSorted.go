package main

import "fmt"

func main() {
	fmt.Println(twoSum2([]int{2,7,11,15}, 9))
}

func twoSum2(numbers []int, target int) []int {
	flip := make(map[int]int)

	for key, value := range numbers {
		flip[value] = key
	}

	for key2, value2 := range numbers {
		sub, err := flip[target - value2]

		if err {
			return []int{key2 + 1, sub + 1}
		}
	}

	return []int{0, 0}
}