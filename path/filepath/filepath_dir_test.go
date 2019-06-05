package array

import (
	"fmt"
	"path/filepath"
	"testing"
)

//返回路径最后一个元素的目录
//路径为空则返回.
func TestFilepathDir(t *testing.T) {
	fmt.Println(filepath.Dir("./a/b/../"))
	fmt.Println(filepath.Dir("./a/b/c"))
	fmt.Println(filepath.Dir("./a/b/c.txt"))
}