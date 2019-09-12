package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// 执行命令，并且会等结果执行完毕
func main() {
	cmd := exec.Command("echo", "hello")

	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	fmt.Println(err)

	fmt.Println(buf.String() + "world")
}
