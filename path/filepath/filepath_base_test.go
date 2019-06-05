package array

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestFilepathBase(t *testing.T) {
	//返回路径的最后一个元素
	fmt.Println(filepath.Base("./a/b/c"))
	//返回路径的最后一个元素
	fmt.Println(filepath.Base("./a/b/c.txt"))
	//如果路径为空字符串，返回.
	fmt.Println(filepath.Base(""))
	//如果路径只有斜线，返回/
	fmt.Println(filepath.Base("///"))
}