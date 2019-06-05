package array

import (
	"fmt"
	"path"
	"testing"
)

//匹配文件名，完全匹配则返回true
func TestFileMatch(t *testing.T) {
	fmt.Println(path.Match("*/*/test", "c/d/test"))
	fmt.Println(path.Match("*/*/test.*", "c/d/test.txt"))
	fmt.Println(path.Match("*", "a/b/c"))
}