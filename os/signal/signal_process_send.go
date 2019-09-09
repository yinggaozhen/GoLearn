package main

import (
	"fmt"
	"os"
	"syscall"
)

// 这个主要是往另外一个进程发送信号
// @link https://colobu.com/2015/10/09/Linux-Signals/
func main() {
	// TODO 这个pid需要手动填写
	// @see ./signal_process_recv.go
	pid := 21924

	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
	}
	err = process.Signal(syscall.SIGSTOP)
	if err != nil {
		fmt.Println(err)
	}
}
