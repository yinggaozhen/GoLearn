package array

import (
	"fmt"
	"path"
	"testing"
)

//返回路径最后一个元素的目录
//路径为空则返回.
func TestFileDir(t *testing.T) {
	fmt.Println(path.Dir("./a/b/../"))
	fmt.Println(path.Dir("./a/b/c"))
	fmt.Println(path.Dir("./a/b/c.txt"))
}