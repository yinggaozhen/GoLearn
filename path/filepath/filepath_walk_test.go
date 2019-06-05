package array

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

//遍历指定目录下所有文件
func TestFilepathWalk(t *testing.T) {
	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	});
}