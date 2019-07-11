package array

import (
	"fmt"
	"testing"
)

func goroutinesSync1() {
	fmt.Println("goroutines1")
}

func goroutinesSync2() {
	fmt.Println("goroutines2")
}

func goroutinesSync3() {
	fmt.Println("goroutines3")
}

func goroutinesSync4() {
	fmt.Println("goroutines4")
}

func TestGoSync(t *testing.T) {
	// 函数输出是无序的
	fmt.Println("main start")
	go goroutinesSync1()
	go goroutinesSync2()
	go goroutinesSync3()
	go goroutinesSync4()
	fmt.Println("main end")
}