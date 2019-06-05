package array

import (
	"fmt"
	"path/filepath"
	"testing"
)

//分割路径中的目录与文件
func TestFilepathSplit(t *testing.T) {
	fmt.Println(filepath.Split("./a/b/c/d.jpg"))
	fmt.Println(filepath.Split("./a/b/c"))
}