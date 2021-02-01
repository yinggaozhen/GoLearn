package array

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var concurrencyCount = new(int32)
var wg = sync.WaitGroup{}

func goroutinesAtomicCount() {
	atomic.AddInt32(concurrencyCount, 1)
	wg.Done()
}

func TestGoRoutinesAtomicCount(t *testing.T) {
	*concurrencyCount = 0

	// 函数输出是无序的
	fmt.Println("concurrency count main start")

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go goroutinesAtomicCount()
	}

	wg.Wait()

	fmt.Println(*concurrencyCount)
}