package array

import (
	"fmt"
	"path/filepath"
	"testing"
)

//将路径使用路径列表分隔符分开，见os.PathListSeparator
//linux下默认为:，windows下为;
func TestFilepathSplitList(t *testing.T) {
	fmt.Println(filepath.SplitList("a:/a/b/c:/tmp"))
}