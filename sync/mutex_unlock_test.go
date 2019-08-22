package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var m2 sync.Mutex

// test mutex lock
func TestMutexUnlock(t *testing.T) {
	go getLockAndUnlock("thread_1")
	go getLockAndUnlock("thread_2")
	go getLockAndUnlock("thread_3")
	go getLockAndUnlock("thread_4")
	go getLockAndUnlock("thread_5")

	time.Sleep(time.Millisecond * 5)
}

func getLockAndUnlock(threadName string) {
	fmt.Println("Start " + threadName)
	m2.Lock() // 其中一个协程在拿到lock之后，会让协程堵塞，只有在协程解锁之后，其他协程才能继续处理
	defer m2.Unlock() // 建议在lock之后紧跟一个defer unlock，养成一个有借有还的好习惯

	fmt.Println("Lock " + threadName)
	time.Sleep(time.Millisecond * 1)
}
