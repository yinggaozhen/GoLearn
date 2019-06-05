package array

import (
	"fmt"
	"path/filepath"
	"testing"
)

//返回路径中的扩展名
//如果没有点，返回空
func TestFilepathExt(t *testing.T) {
	fmt.Println(filepath.Ext("./a/b/../c.txt"))
	fmt.Println(filepath.Ext("./a/b/../c") == "")
}