package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var m3 sync.Mutex

// test mutex lock
func TestMutexUnlock2(t *testing.T) {
	go getLockAndUnlock2("thread_1")
	go getLockAndUnlock2("thread_2")
	go getLockAndUnlock2("thread_3")
	go getLockAndUnlock2("thread_4")
	go getLockAndUnlock2("thread_5")

	time.Sleep(time.Second * 10)
}

func getLockAndUnlock2(threadName string) {
	fmt.Println(">> Access " + threadName + getNow())
	m3.Lock()
	defer m3.Unlock()

	LockAndUnlock2FakeProcess(threadName)
}

func LockAndUnlock2FakeProcess(threadName string) {
	fmt.Println("++ Lock " + threadName + getNow())
	time.Sleep(time.Second * 1)
	fmt.Println("-- UnLock " + threadName + getNow())
}

func getNow2() string {
	return " ||| " + time.Now().Format("04:05")
}