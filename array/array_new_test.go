package array

import (
	"fmt"
	"testing"
)

// 新建数组
func TestNewArray(t *testing.T) {
	var array1 = []int{1, 2}
	fmt.Println(array1)
}

// 新建数组
func TestNewArray2(t *testing.T) {
	var array2 = []int{1, 2}
	testAppend(&array2)
	fmt.Println(array2)
}

func testAppend(arr *[]int) {
	tmpArray := append(*arr, 1)
	arr = &tmpArray
}
