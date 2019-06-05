package array

import (
	"fmt"
	"path/filepath"
	"testing"
)

//匹配文件名，完全匹配则返回true
func TestFilepathMatch(t *testing.T) {
	fmt.Println(filepath.Match("*/*/test", "c/d/test"))
	fmt.Println(filepath.Match("*/*/test.*", "c/d/test.txt"))
	fmt.Println(filepath.Match("*", "a/b/c"))
}