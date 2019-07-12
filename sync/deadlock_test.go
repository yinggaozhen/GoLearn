package test

import (
	"fmt"
	"sync"
	"testing"
)

var deadLockWg sync.WaitGroup
var deadLockCount int

func waitGoDeadLockFunc() {
	deadLockCount++
	deadLockWg.Done()
}

// add/done必须要成对出现，否则将会报死锁
func TestDeadLock(t *testing.T) {
	for i := 0; i < 100; i++ {
		deadLockWg.Add(1)
		go waitGoDeadLockFunc()
	}

	// fatal error: all goroutines are asleep - deadlock!
	deadLockWg.Add(1)
	deadLockWg.Wait()

	fmt.Println(deadLockCount)
}
