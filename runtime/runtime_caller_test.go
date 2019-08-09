package test

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

const (
	SLASH = "/"
	DOT = "."
)

// 往上找调用堆栈
func TestRuntimeCaller(t *testing.T) {
	aStack()
}

func aStack() {
	bStack()
}

func bStack() {
	cStack()
}

func cStack() {
	dStack()
}

func dStack() {
	pc, file, line, ok := runtime.Caller(1)
	fmt.Println(function(pc))
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)
}

func function(pc uintptr) string {
	fn := runtime.FuncForPC(pc).Name()

	if index := strings.LastIndex(fn, SLASH); index >= 0 {
		fn = fn[index+1:]
	}

	if index := strings.LastIndex(fn, DOT); index >= 0 {
		fn = fn[index+1:]
	}

	return fn
}