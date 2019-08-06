package array

import (
	"fmt"
	"strings"
	"testing"
)

// 字符串切割
func TestStringSpilt(t *testing.T) {
	string := "1:2:3:4"

	array := strings.Split(string, ":")
	fmt.Println(array)
}