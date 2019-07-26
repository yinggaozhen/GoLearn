package array

import (
	"fmt"
	"testing"
)

// 交换两个数
func TestChangeNumber(t *testing.T) {
	a, b := 1, 2
	fmt.Println(a, b)
	b, a = a, b
	fmt.Println(a, b)
}