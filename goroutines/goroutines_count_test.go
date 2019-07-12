package array

import (
	"fmt"
	"testing"
	"time"
)

var count int

func goroutinesCount() {
	count++
}

func TestGoRoutinesCount(t *testing.T) {
	count = 0

	// 函数输出是无序的
	fmt.Println("count main start")

	for i := 0; i < 10000; i++ {
		go goroutinesCount()
	}

	// 这里可以更改为sync.waitGroup
	// @link /data/github/GoLearn/sync/wait_test.go
	time.Sleep(10 * time.Millisecond)

	// 这里输出 <= 1000, 类似java线程
	fmt.Println(count)
}