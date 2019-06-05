package array

import (
	"fmt"
	"path/filepath"
	"testing"
)

//返回所给路径的绝对路径
func TestFilepathAbs(t *testing.T) {
	fmt.Println(filepath.Abs("./a/b/../c.txt"))
	fmt.Println(filepath.Abs("./"))
	fmt.Println(filepath.Abs("/a"))
	fmt.Println(filepath.Abs(""))
}