package array

import (
	"fmt"
	"testing"
	"time"
)

func TestAfter(t *testing.T) {
	fmt.Println("before: " + getNow())

	select {
	case <-time.After(time.Second * 3):
		fmt.Println("delay 3 seconds")
	}

	fmt.Println("after: " + getNow())
}

func getNow() string {
	return time.Unix(time.Now().Unix(), 0).Format(time.RFC3339)
}