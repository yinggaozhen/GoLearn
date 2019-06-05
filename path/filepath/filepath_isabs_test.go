package array

import (
	"fmt"
	"path/filepath"
	"testing"
)

//返回所给路径的绝对路径
func TestFilepathIsAbs(t *testing.T) {
	fmt.Println(filepath.IsAbs("./a/b/../c.txt"))
	fmt.Println(filepath.IsAbs("./"))
	fmt.Println(filepath.IsAbs("/a"))
	fmt.Println(filepath.IsAbs(""))
}