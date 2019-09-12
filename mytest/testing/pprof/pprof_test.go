package benchmark

import (
	"sync/atomic"
	"testing"
)

// 通过 pprof 分析性能瓶颈
// 执行命令 go test -v github.com/yinggaozhen/GoLearn/mytest/testing/pprof -test.bench=".*" -cpuprofile=/tmp/cpuprofile.cov
func BenchmarkIncr(t *testing.B) {
	incr()
	atomicIncr()
}

func incr() int32 {
	counter := int32(0)
	for i := 0; i < 10000000; i++ {
		counter++
	}

	return counter
}

func atomicIncr() int32 {
	counter := new(int32)
	for i := 0; i < 10000000; i++ {
		atomic.AddInt32(counter, 1)
	}

	return *counter
}