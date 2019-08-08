package array

import (
	"fmt"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	fmt.Println("before sleep")
	time.Sleep(3 * time.Second)
	fmt.Println("after sleep")
}