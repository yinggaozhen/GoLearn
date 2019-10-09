package array

import (
	"fmt"
	"testing"
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