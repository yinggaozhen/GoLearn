package array

import (
	"fmt"
	"path/filepath"
	"testing"
)

//将路径中的/替换为路径分隔符
func TestFilepathFromSlash(t *testing.T) {
	fmt.Println(filepath.FromSlash("./a/b/c"))
}