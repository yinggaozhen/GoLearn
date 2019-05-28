package array

import (
	"fmt"
	"testing"
)

// 反转字符串
func TestStringReverse(t *testing.T) {
	var string1 = "123"

	fmt.Print(Reverse(string1))
}

func Reverse(reverse string) string {
	var result string

	for _ , v := range reverse {
		result = string(v) + result
	}

	return result
}























//
//func Reverse(reverse string) string {
//	var result string
//
//	for _ , v := range reverse {
//		result = string(v) + result
//	}
//
//	return result
//}
