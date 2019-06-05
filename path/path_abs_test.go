package array

import (
	"fmt"
	"path"
	"testing"
)

//判断路径是不是绝对路径
func TestFileIsAbs(t *testing.T) {
	fmt.Println(path.IsAbs("./a/b/../c.txt"))
	fmt.Println(path.IsAbs("./"))
	fmt.Println(path.IsAbs("/a"))
}