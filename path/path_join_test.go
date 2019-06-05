package array

import (
	"fmt"
	"path"
	"testing"
)

//连接路径，返回已经clean过的路径
func TestFileJoin(t *testing.T) {
	fmt.Println(path.Join("/a", "c", "d"))
	fmt.Println(path.Join("a", "c", "d"))
	fmt.Println(path.Join("./a", "c", "d"))
	fmt.Println(path.Join("./a", "/c", "d"))
	fmt.Println(path.Join("./a", "//c", "../d/"))
	fmt.Println(path.Join("./a", "//c", "../d/e.txt"))
}