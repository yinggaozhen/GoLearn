package array

import (
	"fmt"
	"runtime"
	"testing"
)

// 切割数组
func TestCutArray(t *testing.T) {
	var array1 = []int{1, 2, 3, 4, 5}
	// 前闭后开 [...)
	fmt.Println(array1[2:4])
}