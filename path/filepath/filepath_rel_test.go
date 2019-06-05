package array

import (
	"fmt"
	"path/filepath"
	"testing"
)

//返回以basepath为基准的相对路径
func TestFilepathRel(t *testing.T) {
	fmt.Println(filepath.Rel("/", "/tmp"))
	fmt.Println(filepath.Rel("/guoozhen", "/tmp/aa"))
	fmt.Println(filepath.Rel("/guoozhen", "./"))
	fmt.Println(filepath.Rel("/guoozhen", ""))
}