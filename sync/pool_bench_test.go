package test

import (
	"sync"
	"testing"
)

type Student struct {
	Index int
}

func (s Student) SetIndex(index int) {
	s.Index = index
}

var poolBench sync.Pool
var poolBenchMutex sync.Mutex
var loop = 10000000

// pool 是当没有的值的时候会New
// 被Put一次之后，下次再取就会从pool取出来，然后清空
func BenchmarkSyncPool1(b *testing.B) {
	poolBench.New = func() interface{} {
		return Student{
		}
	}

	for i := 0; i < loop; i++ {
		go func(index int) {
			s := poolBench.Get().(Student)
			s.SetIndex(i)
			poolBench.Put(s)
		}(i)
	}
}

func BenchmarkSyncPool2(b *testing.B) {
	for i := 0; i < loop; i++ {
		go func(index int) {
			poolBenchMutex.Lock()
			defer poolBenchMutex.Unlock()
			s := Student{}
			s.SetIndex(index)
		}(i)
	}
}
