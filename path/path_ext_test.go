package array

import (
	"fmt"
	"path"
	"testing"
)

//返回路径中的扩展名
//如果没有点，返回空
func TestFileExt(t *testing.T) {
	fmt.Println(path.Ext("./a/b/../c.txt"))
	fmt.Println(path.Ext("./a/b/../c") == "")
}