package test

import (
	"fmt"
	"log"
	"os"
	"testing"
)

//Ldate         = 1 << iota     //日期示例： 2009/01/23
//Ltime                         //时间示例: 01:23:23
//Lmicroseconds                 //毫秒示例: 01:23:23.123123.
//Llongfile                     //绝对路径和行号: /a/b/c/d.go:23
//Lshortfile                    //文件和行号: d.go:23.
//LUTC                          //日期时间转为0时区的
//LstdFlags     = Ldate | Ltime //Go提供的标准抬头信息

// 参考资料 : {@link https://www.flysnow.org/2017/05/06/go-in-action-go-log.html}
func TestLogger(t *testing.T) {
	logger := log.New(os.Stdout, "[prefix] ", log.Ldate)
	logger.Println("hello1")
	logger.Println("hello2")
	logger.Println("hello3")

	logger = log.New(os.Stdout, "[prefix2] ", log.Ldate | log.Llongfile)
	logger.Println("hello1")
	logger.Println("hello2")
	logger.Println("hello3")

	logger = log.New(os.Stdout, "[prefix2] ", log.LstdFlags | log.Llongfile)
	logger.Println("hello1")
	logger.Println("hello2")
	logger.Println("hello3")
}

func TestFatalLogger(t *testing.T) {
	logger := log.New(os.Stdout, "[FatalPrefix] ", log.LstdFlags | log.Llongfile)

	// 输出message之后，在调用os.Exit(1)
	logger.Fatal("fatal message")

	fmt.Println("如果你看到我，我就让你嘿嘿嘿")
}

func TestPanicLogger(t *testing.T) {
	logger := log.New(os.Stdout, "[PanicPrefix] ", log.LstdFlags | log.Llongfile)

	defer func() {
		if ok := recover(); ok != nil {
			fmt.Println("[Recover Message] " + ok.(string))
		}
	}()

	// 输出message之后，抛出panic
	logger.Panic("panic message")

	fmt.Println("如果你看到我，我就让你嘿嘿嘿")
}