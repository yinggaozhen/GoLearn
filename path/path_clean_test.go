package array

import (
	"fmt"
	"path"
	"testing"
)

// 返回等价的最短路径
//1.用一个斜线替换多个斜线
//2.清除当前路径.
//3.清除内部的..和他前面的元素
//4.以/..开头的，变成/
func TestFileClean(t *testing.T) {
	fmt.Println(path.Clean("./a/b/../"))
	fmt.Println(path.Clean("./a/b/../c"))
	fmt.Println(path.Clean("./"))
	fmt.Println(path.Clean("."))
	fmt.Println(path.Clean(""))
	fmt.Println(path.Clean("./a.txt"))
}