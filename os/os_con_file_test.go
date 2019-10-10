package os

import (
	"fmt"
	"os"
	"testing"
	"time"
)

const ConOFile = "/tmp/hello"

// 切记，file默认是会从第一行覆盖式写
// os.O_CREATE : 没有文件就会创建
// os.O_RDWR : 文件读写
// os.O_TRUNC : 将文件清空再写
// os.O_APPEND : 在文件尾进行追加
func TestConcurrentIoWrite(t *testing.T) {
	handler, err := os.OpenFile(ConOFile, os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 1000; i++ {
		go ConcurrentIoWrite(handler)
	}

	time.Sleep(1 * time.Second)
}

func ConcurrentIoWrite(handler *os.File) {
	newString := fmt.Sprint("Hello World!" + time.Now().Format(time.RFC3339) + "\n")

	i, err := handler.WriteString(newString)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}