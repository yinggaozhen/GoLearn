package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4,5,5,5}))
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}