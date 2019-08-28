package array

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestFmtFprintFunc(t *testing.T) {
	w := new(bytes.Buffer)

	// 就是讲内容写入到writer中
	fmt.Fprintf(w, "【1】 Hello %s", "xiaoMing")
	fmt.Println(w.String())

	// 继续往writer中追加
	fmt.Fprintf(w, "【2】  ----- Hello2 %s", "xiaoMing2")
	fmt.Println(w.String())

	// 清空
	w.Reset()
	fmt.Fprintf(w, "【3】 ----- clean %s", "CleanXiaoMing")
	fmt.Println(w.String())
}

func TestFmtFprintFunc2(t *testing.T) {
	w := os.Stdout

	// 就是讲内容写入到writer中
	fmt.Fprintf(w, "【1】 Hello %s\n", "xiaoMing")

	// 继续往writer中追加
	fmt.Fprintf(w, "【2】  ----- Hello2 %s\n", "xiaoMing2")

	// 清空
	fmt.Fprintf(w, "【3】 ----- clean %s\n", "CleanXiaoMing")
}