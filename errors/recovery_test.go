package array

import (
	"fmt"
	"testing"
	"time"
)

func funcPanic() {
	panic("hello")
}

func funcCanRecovery() {
	// defer 在执行完panic之后再执行
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	funcPanic()
}

func funcCanNotRecovery() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
}

func TestCanRecovery(t *testing.T) {
	funcCanRecovery()
	fmt.Println("after panic")
}


func TestCanNotRecovery(t *testing.T) {
	// 已经执行完defer了
	funcCanNotRecovery()
	funcPanic()
	fmt.Println("after panic")
}

// 协程之间的recovery不会互相影响，A协程抛出异常之后，B协程的recovery是不会起作用的
func TestGoroutineRecovery(t *testing.T) {
	// 这里不会捕获协程抛出的错误，只会捕获当前进程
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	go funcPanic()

	fmt.Println("before goroutine panic")
	time.Sleep(1 * time.Second)
	// 这里就会终止
	fmt.Println("after goroutine panic")
}