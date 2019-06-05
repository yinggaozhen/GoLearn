package array

import (
	"fmt"
	"path/filepath"
	"testing"
)

//连接路径，返回已经clean过的路径
func TestFilepathJoin(t *testing.T) {
	fmt.Println(filepath.Join("/a", "c", "d"))
	fmt.Println(filepath.Join("a", "c", "d"))
	fmt.Println(filepath.Join("./a", "c", "d"))
	fmt.Println(filepath.Join("./a", "/c", "d"))
	fmt.Println(filepath.Join("./a", "//c", "../d/"))
	fmt.Println(filepath.Join("./a", "//c", "../d/e.txt"))
}