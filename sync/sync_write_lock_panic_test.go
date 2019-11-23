package test

import (
	"fmt"
	"sync"
	"testing"
)

var mu sync.Mutex

func TestWriteLockPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("writeLockPanic start")
			fmt.Println(err)

			// mu.Unlock() 如果不解锁的话，直接在次进行加锁，会导致错误

			mu.Lock()
			fmt.Println("是否能到这里呢??")
			mu.Unlock()
		}
	}()

	mu.Lock()
	fmt.Println("hello")
	panic("writeLockPanic")
	mu.Unlock()
}

// 采用defer unlock 来取代 unlock
// defer unlock 的顺序在 recovery 后面
// 结果 : 可行，因为defer顺序是FILO，所以先1.执行解锁，2.然后再执行recovery
func TestWriteLockPanic2(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("writeLockPanic start")
			fmt.Println(err)

			// mu.Unlock() 如果不解锁的话，直接在次进行加锁，会导致错误

			mu.Lock()
			fmt.Println("是否能到这里呢??")
			mu.Unlock()
		}
	}()

	mu.Lock()
	defer mu.Unlock()

	fmt.Println("hello")
	panic("writeLockPanic")
}


// 采用defer unlock 来取代 unlock
// recovery 的顺序在 defer unlock 后面
// 结果 : 不可行。因为defer顺序是FILO，所以先1.执行recovery，2.然后再执行解锁，但是执行1时，锁还未被释放
func TestWriteLockPanic3(t *testing.T) {
	mu.Lock()
	defer mu.Unlock()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("writeLockPanic start")
			fmt.Println(err)

			// mu.Unlock() 如果不解锁的话，直接在次进行加锁，会导致错误

			mu.Lock()
			fmt.Println("是否能到这里呢??")
			mu.Unlock()
		}
	}()

	fmt.Println("hello")
	panic("writeLockPanic")
}