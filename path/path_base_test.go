package array

import (
	"fmt"
	"path"
	"testing"
)

func TestFileBase(t *testing.T) {
	//返回路径的最后一个元素
	fmt.Println(path.Base("./a/b/c"))
	//返回路径的最后一个元素
	fmt.Println(path.Base("./a/b/c.txt"))
	//如果路径为空字符串，返回.
	fmt.Println(path.Base(""))
	//如果路径只有斜线，返回/
	fmt.Println(path.Base("///"))
}