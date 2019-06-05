package array

import (
	"fmt"
	"path"
	"testing"
)

//分割路径中的目录与文件
func TestFileSplit(t *testing.T) {
	fmt.Println(path.Split("./a/b/c/d.jpg"))
	fmt.Println(path.Split("./a/b/c"))
}