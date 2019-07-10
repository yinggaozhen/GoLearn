package array

import (
	"fmt"
	"testing"
)

// https://sanyuesha.com/2017/07/23/go-defer/

// recovery
func TestPanic(t *testing.T) {
	defer func() {
		if ok := recover(); ok != nil {
			if ok == "hello panic1" {
				fmt.Println("I'm hello panic1")
			}
		}
	}()

	panic("hello panic1")
	panic("hello panic2")
	panic("hello panic3")
}