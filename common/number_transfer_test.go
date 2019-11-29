package array

import (
	"fmt"
	"testing"
)

// 数字之间的类型转换
func TestNumberTypeTransfer(t *testing.T) {
	var a uint = 1
	var b int = 2

	fmt.Println(int64(a) + int64(b))
}