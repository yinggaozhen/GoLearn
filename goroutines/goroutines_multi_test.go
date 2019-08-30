package array

import (
	"fmt"
	"testing"
	"time"
)

func TestGoMulti(t *testing.T) {
	go father()

	time.Sleep(time.Second * 2)
}

func father() {
	fmt.Println("father")

	defer func() {
		fmt.Println("father shunxu")
		if err := recover(); err != nil {
			fmt.Println("father defer" + err.(string)) // son 报出来的错，是捕获不到的
		}
	}()

	time.Sleep(time.Second * 1)

	go son()
}

func son() {
	fmt.Println("son")
	panic("panic from son")
}