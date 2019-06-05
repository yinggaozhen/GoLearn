package array

import (
	"fmt"
	"path/filepath"
	"testing"
)

//返回所有匹配的文件
func TestFilepathGlob(t *testing.T) {
	match, _ := filepath.Glob("./*.go")
	fmt.Println(match)
}