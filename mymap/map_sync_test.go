package array

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 反面教材 : fatal error: concurrent map writes
//func TestMapErrorConcurrent(t *testing.T) {
//	sm := make(map[int]int)
//
//	for i := 0; i < 1000; i++ {
//		go func() {
//			sm[i] = i
//		}()
//	}
//
//	time.Sleep(2 * time.Second)
//	fmt.Println(sm)
//}

var mux sync.Mutex

func TestMapConcurrent(t *testing.T) {
	sm := sync.Map{}

	for i := 0; i < 1000; i++ {
		go func(number int) {
			sm.Store(number, number)
		}(i)
	}

	time.Sleep(1 * time.Second)

	count := 0
	sm.Range(func (key, value interface{}) bool {
		fmt.Println(key, value)
		count++
		return true
	})
	fmt.Println(count)
}