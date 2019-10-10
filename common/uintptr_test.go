package array

import (
	"fmt"
	"reflect"
	"testing"
)

// uintptr 就是一个无符号的整形，uintptr值可以进行算术运算，由uintptr变量表示的地址处的数据也可能被GC回收。
func TestUintptr(t *testing.T) {
	var a uintptr
	fmt.Println(reflect.TypeOf(a).String())
	a = addUintptr(a)
	fmt.Println(a)
}

func addUintptr(a uintptr) uintptr {
	return a + 1
}