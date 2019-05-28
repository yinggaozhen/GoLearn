package array

import (
	"fmt"
	"testing"
)

// 数组添加
func TestArrayAppend(t *testing.T) {
	var array []string

	array = append(array, "hello")
	array = append(array, "world")

	fmt.Print(array)
}
