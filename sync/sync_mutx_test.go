package test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

// 似乎只有在map的时候会有问题，会报错
// concurrent error
func TestMapConcurrentWithoutMutex(t *testing.T) {
	kv := make(map[string]string)
	for i := 0; i < 10000; i++ {
		go func(key string) {
			kv[key] = key
		}(strconv.Itoa(i))
	}
}

// scalar success
func TestScalarConcurrentWithoutMutex(t *testing.T) {
	var v string
	for i := 0; i < 10000; i++ {
		go func(key string) {
			v = key
		}(strconv.Itoa(i))
	}
	fmt.Println(v)
}

// slice success
func TestSliceConcurrentWithoutMutex(t *testing.T) {
	var array []string

	for i := 0; i < 10000; i++ {
		go func(key string) {
			array = append(array, key)
		}(strconv.Itoa(i))
	}

	time.Sleep(time.Second)
	fmt.Println(len(array))
}

// slice success
func TestSliceConcurrentWithMutex(t *testing.T) {
	var array []string
	var mu sync.Mutex

	for i := 0; i < 10000; i++ {
		go func(key string) {
			mu.Lock()
			defer mu.Unlock()
			array = append(array, key)
		}(strconv.Itoa(i))
	}

	time.Sleep(time.Second)
	fmt.Println(len(array))
}
