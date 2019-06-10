package array

import (
	"fmt"
	"testing"
)

// 反转字符串
func TestStringGet(t *testing.T) {
	context := "123"

	var c byte

	c = context[0]
	fmt.Println(c)

	c = context[1]
	fmt.Println(c)

	c = context[2]
	fmt.Println(c)

	// 报错
	c = context[3]
	fmt.Println(c)
}