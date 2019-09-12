package main

import (
	"fmt"
	"os"
	"os/exec"
)

// 执行命令，但是不等待结果
func main() {
	cmd := exec.Command("sleep", "2")

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	err := cmd.Start()
	fmt.Println(err)

	// 需要加入wait才能等待命令执行完毕
	cmd.Wait()

	fmt.Println("finish")
}
