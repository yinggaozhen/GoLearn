package array

import (
	"fmt"
	"testing"
)

// 字符串切割
func TestStringSub(t *testing.T) {
	string := "0123456"

	// 前闭后开[...)
	fmt.Println(string[0:1]) // "0"
	fmt.Println(string[0:2]) // "01"
	fmt.Println(string[2:2] == "") // ""
	fmt.Println(string[2:3]) // "2"
	fmt.Println(string[2:6]) // "2345"
	fmt.Println(string[2:7]) // "234567" 因为是开区间，所以还是能访问

	// fmt.Println(string[2:8]) // 8 已经越界了，会报错
}