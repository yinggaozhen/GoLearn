package array

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wLock sync.Mutex
var rLock sync.RWMutex

func writeLockBench() {
	wLock.Lock()
	defer wLock.Unlock()
}

func readLockBench() {
	rLock.RLock()
	defer rLock.RUnlock()
}

func BenchmarkWriteLock(b *testing.B) {
	start := time.Now()

	for i := 0; i < 10000000; i++ {
		go writeLockBench()
	}

	fmt.Println(time.Now().Sub(start).Seconds())
}

func BenchmarkReadLock(b *testing.B) {
	start := time.Now()

	for i := 0; i < 10000000; i++ {
		go readLockBench()
	}

	fmt.Println(time.Now().Sub(start).Seconds())
}