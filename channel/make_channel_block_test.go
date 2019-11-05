package array

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelGoBlock(t *testing.T) {
	c := make(chan string, 2)

	go func() {
		// goroutines block
		for {
			r := <-c
			fmt.Println(r)
		}
	}()

	go func() {
		for {
			time.Sleep(3 * time.Second)
			c <- "hello"
		}
	}()

	select {
	}
}