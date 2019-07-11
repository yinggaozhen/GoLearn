package array

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

var concurrencyCount = new(int32)

func goroutinesAtomicCount() {
	atomic.AddInt32(concurrencyCount, 1)
}

func TestGoRoutinesAtomicCount(t *testing.T) {
	*concurrencyCount = 0

	// 函数输出是无序的
	fmt.Println("concurrency count main start")

	for i := 0; i < 10000; i++ {
		go goroutinesAtomicCount()
	}

	time.Sleep(15 * time.Millisecond)
	fmt.Println(*concurrencyCount)
}