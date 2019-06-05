package array

import (
	"fmt"
	"path/filepath"
	"testing"
)

//将路径分隔符使用/替换
// 主要是为了适配不同操作系统
func TestFilepathToSlash(t *testing.T) {
	fmt.Println(filepath.ToSlash("C:/a/b"))
}