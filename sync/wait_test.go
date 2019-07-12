package test

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup
var count int

func waitGoFunc() {
	count++
	wg.Done()
}

// test struct extends
func TestWait(t *testing.T) {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go waitGoFunc()
	}

	wg.Wait()

	fmt.Println(count)
}
