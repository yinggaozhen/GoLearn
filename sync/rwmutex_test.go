package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wr sync.RWMutex
var number int

// test mutex read lock
// @link https://fivezh.github.io/2019/04/09/sync_mutex_translation/
func TestReadMutexLock(t *testing.T) {
	for i := 1; i < 1000; i++ {
		go getReadLock()
	}
	wr.Lock()
	number = 1
	time.Sleep(time.Second * 3)
	wr.Unlock()
}

func getReadLock() {
	// 添加读锁之后，当其他协程发生lock之后，读锁也会被lock，否则读锁不会堵塞
	wr.RLock()
	defer wr.RUnlock()

	if number > 0 {
		fmt.Println(number)
	}
}