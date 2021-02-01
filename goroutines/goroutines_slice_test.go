package array

import (
	"fmt"
	"sync"
	"testing"
)

var goSlice []int
var lock = sync.Mutex{}
var wg = sync.WaitGroup{}

func goroutinesSlice(i int) {
	goSlice = append(goSlice, i)
	wg.Done()
	lock.Unlock()
}

func TestGoRoutinesSlice(t *testing.T) {
	for i := 1; i < 1000; i = i + 2 {
		lock.Lock()
		wg.Add(1)
		go goroutinesSlice(i)
	}

	wg.Wait()

	for k, v := range goSlice {
		fmt.Println(k, v)
	}
}