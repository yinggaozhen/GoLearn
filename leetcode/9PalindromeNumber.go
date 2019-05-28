package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	rx := Reverse9(strconv.Itoa(x))

	var result bool = false
	if strings.Compare(strconv.Itoa(x), rx) == 0 {
		result = true
	}

	return result
}

func Reverse9(reverse string) string {
	var result string

	for _, v := range reverse {
		result = string(v) + result
	}

	return result
}
