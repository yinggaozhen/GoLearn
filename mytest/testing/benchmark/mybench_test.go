package benchmark

import (
	"sync/atomic"
	"testing"
)

// 必须要 Benchmark 开头

// 普通循环递增(进程非安全)
func BenchmarkIncr(t *testing.B) {
	incr()
}

// 用atomic递增(进程安全)
func BenchmarkAtomicIncr(t *testing.B) {
	atomicIncr()
}

func incr() int32 {
	counter := int32(0)
	for i := 0; i < 100000000; i++ {
		counter++
	}

	return counter
}

func atomicIncr() int32 {
	counter := new(int32)
	for i := 0; i < 100000000; i++ {
		atomic.AddInt32(counter, 1)
	}

	return *counter
}