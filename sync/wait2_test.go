package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg2 sync.WaitGroup

func processNeed5Sec() {
	time.Sleep(5*time.Second)
	wg2.Done()
}

func processNeed6Sec() {
	time.Sleep(6*time.Second)
	wg2.Done()
}

// test struct extends
func TestWait2(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	wg2.Add(1)
	go processNeed5Sec()

	wg2.Add(1)
	go processNeed6Sec()

	wg2.Wait()

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
