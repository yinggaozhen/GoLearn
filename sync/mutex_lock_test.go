package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var m sync.Mutex

// test mutex lock
func TestMutexLock(t *testing.T) {
	go getLock("thread_1")
	go getLock("thread_2")
	go getLock("thread_3")
	go getLock("thread_4")
	go getLock("thread_5")

	time.Sleep(time.Second)
}

func getLock(threadName string) {
	fmt.Println("Start " + threadName)
	m.Lock()
	m.Unlock()
	fmt.Println("Lock " + threadName)
}
