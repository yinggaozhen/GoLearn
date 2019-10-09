package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"io/ioutil"
	"path/filepath"
	"time"
)

// go build -o /tmp/hello index.go && /tmp/hello
// 最终获取到的结果还是不变，读文件也是

func main() {
	fmt.Println(filepath.Abs("."))
	fmt.Println(filepath.Abs("../"))

	result, _ := ioutil.ReadFile("../hello")
	fmt.Println(string(result))

	n := now.New(time.Now())
	fmt.Println(n.String())

	g := now.Guozhen{}
	g.Echo()
}
