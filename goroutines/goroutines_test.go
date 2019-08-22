package array

import (
	"fmt"
	"testing"
	"time"
)

func goroutines1() {
	fmt.Println("goroutines1")
}

func goroutines2() {
	fmt.Println("goroutines2")
}

func goroutines3() {
	fmt.Println("goroutines3")
}

func goroutines4() {
	fmt.Println("goroutines4")
}

func TestGoRoutines(t *testing.T) {
	// 函数输出是无序的
	fmt.Println("goroutines start")
	go goroutines1()
	go goroutines2()
	go goroutines3()
	go goroutines4()
	fmt.Println("goroutines end")

	// 这里可以更改为sync.waitGroup
	// @link /data/github/GoLearn/sync/wait_test.go
	time.Sleep(10 * time.Millisecond)
}