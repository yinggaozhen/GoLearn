package array

import (
	"fmt"
	"testing"
	"time"
)

var goSlice []int

func goroutinesSlice(i int) {
	goSlice[i] = i
}

func TestGoRoutinesSlice(t *testing.T) {
	goSlice = make([]int, 10000)
	for i := 0; i < 10000; i++ {
		if i != 1 {
			go goroutinesSlice(i)
		}
	}

	// 这里可以更改为sync.waitGroup
	// @link /data/github/GoLearn/sync/wait_test.go
	time.Sleep(10 * time.Millisecond)

	fmt.Println(goSlice)
}